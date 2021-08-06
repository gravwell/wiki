# DockerでのSoliton NKのデプロイ

Docker Hubで利用可能なビルド済みのDockerイメージを使えば、実験や長期利用のためにDocker内にSoliton NKをデプロイすることが非常に簡単にできます。このドキュメントでは、Docker内でSoliton NK環境を設定する方法を紹介します。

Soliton NKの正規ユーザーで、DockerでSoliton NKをデプロイしたい場合は、[サポート](https://www.soliton.co.jp/support/) を通じて連絡してください。

Soliton NKをセットアップしたら、[クイックスタート](#!quickstart/quickstart.md)をチェックして、*Soliton NKを使う上での*スタートポイントを確認してください。

注：MacOS上でDockerを実行しているユーザーは、[Dockerの解説ページ](https://docs.docker.com/docker-for-mac/networking/)で説明されているように、MacOSホストからコンテナに対して直接IPでアクセスできないことに注意する必要があります。ホストからコンテナのネットワークサービスにアクセスする必要がある場合は、追加のポートを転送する準備をしておきましょう。

## Dockerネットワークの作成

Soliton NKコンテナを他のコンテナから分離しておくために、`gravnet`というDockerネットワークを作成します:

	docker network create gravnet

## インデクサーとウェブサーバーのデプロイ

Soliton NKのコンテナイメージは利便性を考慮して、インデクサーとウェブサーバーのフロントエンドにSimple Relayインジェスターを加えたものを１つのDockerイメージ（[soliton/solitonnk](https://hub.docker.com/r/soliton/solitonnk/)）として出荷しています。ウェブサーバーへアクセスするために、ホストのポート8080版をコンテナの80番に転送して起動します:

	docker run --net gravnet -p 8080:80 -p 4023:4023 -p 4024:4024 -d -e GRAVWELL_INGEST_SECRET=MyIngestSecret -e GRAVWELL_INGEST_AUTH=MyIngestSecret -e GRAVWELL_CONTROL_AUTH=MyControlSecret -e GRAVWELL_SEARCHAGENT_AUTH=MySearchAgentAuth --name gravwell soliton/solitonnk:latest

新しいコンテナは`gravwell`という名前になっていることに注意してください。この名前は、インジェスターにデータを送る先のインデクサーを指定するときに使用します。

テストに使える環境変数がいくつかあります。この環境変数に、Soliton NKのコンポーネント間の通信に使用される共有鍵を設定します。通常は[設定ファイル](#!configuration/parameters.md)に設定しますが、Dockerフレンドリーな方法で手早く設定するには、[環境変数](#!configuration/environment-variables.md)を使います。後でインジェスターにも `GRAVWELL_INGEST_SECRET=MyIngestSecret` の値を設定します。ここで設定された環境変数の意味は次の通りです:

* `GRAVWELL_INGEST_AUTH=MyIngestSecret` ：インジェスターの認証にMyIngestSecretを使用するように*インデクサー*に指示します。
* `GRAVWELL_INGEST_SECRET=MyIngestSecret` ：インデクサーの認証にMyIngestSecretを使用するように*Simple Relay インジェスター*に指示します。この値は、**必ず**GRAVWELL_INGEST_AUTHの値と一致しなければなりません！
* `GRAVWELL_CONTROL_AUTH=MyControlSecret`：*フロントエンド*と*インデクサー*にMyControlSecretを使用して相互に認証するように指示します。
* `GRAVWELL_SEARCHAGENT_AUTH=MySearchAgentAuth` ：検索エージェントの認証に MySearchAgentAuth を使用するように*フロントエンド*に指示します。

注意：長期的に運用する予定の場合、特に何らかの方法でインターネットに公開する場合には、これらの認証鍵をデフォルトとは異る値に設定することを**強く**お勧めします。

注意：GRAVWELL_INGEST_AUTH の鍵は GRAVWELL_INGEST_SECRET の鍵と必ず一致しなければなりません。

### 永続ストレージの設定

デフォルトの Soliton NK docker デプロイでは、すべてのストレージにベースコンテナを使用します。Dockerには、バインドやボリュームなどを動作させているベースコンテナから独立した、永続的ストレージを設定するためのオプションがいくつか用意されています。本番環境でSoliton NKをデプロイする場合、コンポーネントに応じていくつかのディレクトリを永続的ストレージに保持したいと思うでしょう。永続ストレージの詳細については、[Docker Volumes](https://docs.docker.com/storage/volumes/)のドキュメントを参照してください。

#### インデクサーの永続ストレージ

Soliton NKインデクサーは2つの重要なデータセット、保存されたデータの束と `tags.dat` ファイルを保持しています。インデクサーの他のほとんどのコンポーネントはデータを失うことなく復旧できますが、通常の操作ではいくつかのディレクトリは永続的なストレージにバインドされていなければいけません。重要なデータは `storage`, `resources`, `log`, `etc` ディレクトリに存在します。各ディレクトリはそれぞれ別のボリュームにマウントすることもできますし、`gravwell.conf` ファイルでの記述によって単一の永続ストレージのディレクトリを指すように設定することもできます。dockerのデプロイ用に設計された `gravwell.conf` の例では、各データディレクトリのストレージパスを変更して、`/opt/gravwell` だけではなく `/opt/gravwell/persistent` の中の別のパスを指定して永続ストレージを用いることができるようになっています。すべての `gravwell.conf` 設定パラメーターに関する完全なドキュメントは、[詳細な設定](parameters.md)ページにあります。

#### ウェブサーバーの永続ストレージ

Soliton NK ウェブサーバーには、設定データや検索結果を失わないようにするために保守すべきいくつかのディレクトリがあります。`etc`, `resources`, `saved` ディレクトリには、コンテナデプロイ全体で維持すべき重要なファイルが保存されています。`saved` ディレクトリには、ユーザーが保存することを選択した検索結果が保存されています。`etc` ディレクトリには、ユーザーデータベース、ウェブストア、`tags.dat` ファイルが保存されています。これらすべては Soliton NK の適切な運用に不可欠なものです。

#### インジェスターの永続ストレージ

Soliton NKインジェスターはデータを中継するように設計されており、通常は永続的なストレージを必要としません。例外はキャッシュシステムです。Soliton NK ingest APIには統合されたキャッシュシステムが含まれているので、インデクサーへのアップリンクに問題が発生した場合、インジェスターはデータをローカルの永続的なストアにキャッシュして、データが失われることがないようにします。ほとんどのインジェスターはデフォルトではキャッシュをデプロイしませんが、一般的なキャッシュストレージの場所は `/opt/gravwell/cache` です。`cache`ディレクトリを永続的なストレージにバインドすれば、インジェスターが状態を維持し、コンテナの再起動や更新でデータを失わないようにすることができます。

## ライセンスのアップロードとログイン

Soliton NKを起動後に、http://localhost:8080 に Web ブラウザでアクセスしてください。ライセンスのアップロードを求める表示が出るはずです。

![](license-upload-docker.png)

注: 評価用ライセンスをご希望の方は、[コミュニティサイト](https://www.solitonnk.com/)にユーザ登録をして、評価ラインセンスをダウンロードしてください。

ライセンスをアップロードして検証が済むと、ログイン画面が表示されます:

![](docker-login.png)

デフォルトのログイン情報 **admin** / **changeme** でログインしてください。これでいよいよ Soliton NK に入ります! Soliton NKを動かし続けるつもりなら、パスワードを変更した方がいいでしょう（右上のユーザーアイコンをクリックしてパスワードを変更してください）。

## テスト用のデータ追加

soliton/solitonnk から得られるDockerイメージには、[Simple Relay ingester](#!ingesters/ingesters.md)がプリインストールされていて、以下のポートを開けてリッスンしています:

* TCP 7777 行区切りのログデータ用 ('default'タグ)
* TCP 601 syslog メッセージ用 ('syslog'タグ)
* UDP 514 syslog メッセージ用 ('syslog'タグ)

Soliton NKがデータを取り込めるようになったかを確かめるために、netcatを使ってポート7777にラインを書き込んでみましょう。とはいえ、VMを起動したときにはこれらのポートをホストに転送していなかったはずです。そのような場合、`docker inspect`を使って、Soliton NKコンテナに割り当てられたIPアドレスを取得することができます:

	docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' gravwell

今回例としてSoliton NKコンテナのIPアドレスは、**172.19.0.2**だったとして話を進めます。次に、netcatを使って次のような行を書き、Ctrl-Dを押してそれらを送信します。:

	$ netcat 172.19.0.2 7777
	this is a test
	this is another test

注意：MacOSでは、コンテナは実際にはLinux VM内で実行されているため、IPアドレスを指定して直接コンテナにアクセスすることはできません。Dockerコンテナ内でnetcatを使用するか(同じコンテナでも新しいコンテナでも)、Soliton NKコンテナを起動する際にポート7777をホストに転送する設定をしてください。

そして、ブラウザから時間範囲を「直近１時間」に指定して検索してください。データが入っているかどうか、Soliton NKが正常に動作しているかどうかを確認することができます:

![](docker-search.png)

## インジェスターのセットアップ

soliton/solitonkイメージに同梱されているSimple Relayインジェスターの他に、Soliton NKの開発元であるGravwellから、現在、複数の構築済みスタンドアロンインジェスターイメージを提供しています:

* [gravwell/netflow_capture](https://hub.docker.com/r/gravwell/netflow_capture/)  は Netflow コレクターで、ポート 2055 で Netflow v5 レコードを、ポート 6343 で IPFIX レコードを受信するように構成されています。
* [gravwell/collectd](https://hub.docker.com/r/gravwell/collectd/) は、ポート 25826 で collectd の取得ポイントからのハードウェア統計情報を受信します。
* [gravwell/simple_relay](https://hub.docker.com/r/gravwell/simple_relay/) は、コアイメージにプリインストールされているのと同じSimple Relayインジェスターですが、Simple Relayインジェスターを別の所にデプロイしたい場合に使えます。

Netflow インジェスターの起動方法を以下に示します。同じコマンドを （名前とポートを変更して）他のインジェスターにも使用できます:

	docker run -d --net gravnet -p 2055:2055/udp --name netflow -e GRAVWELL_CLEARTEXT_TARGETS=gravwell -e GRAVWELL_INGEST_SECRET=MyIngestSecret gravwell/netflow_capture

環境変数を設定するために `-e` フラグを使用していることに注意してください。これにより、インジェストのために'gravwell'という名前のコンテナに接続するようにインジェスターを指示し(GRAVWELL_CLEARTEXT_TARGETS=gravwell)、インジェスト共有鍵を'IngestSecrets'に設定する(GRAVWELL_INGEST_SECRET=IngestSecrets)ことで、インジェスターを動的に設定することができるようになります。

`-p 2055:2055/udp` オプションは、UDP ポート 2055 (Netflow v5 のインジェストポート) をコンテナからホストに転送します。これにより、Netflow レコードをインジェストコンテナに送るのが簡単になるはずです。

注: netflow インジェスターは、ポート 6343 の UDP 上で IPFIX レコードを受け入れるようにデフォルトで設定されています。IPFIX レコードもインジェストしたい場合は、上のコマンドラインに `-p 6343:6343/udp` を追加してください。

メニューの「システム」→「インジェスター」をクリックして、どのインジェスターがアクティブであるかを確認できます:

![](netflow_ingest.png)

これで、Netflow ジェネレーターをホストのポート 2055 に向けてレコードを送信するように設定できます。Netflowのデータはコンテナに渡され、Soliton NK にインジェストされます。

## サービスのカスタマイズ

公式のSoliton NK dockerコンテナには、コンテナ内の複数のサービスの起動と制御を非常に簡単にするサービス管理システムが含まれています。サービス管理システムでは、サービスの再起動、エラー報告、バックオフ制御を管理操作できます。Soliton NKの開発元であるGravwellは、BSD 3-Clauseライセンスのもと、[github](https://github.com/gravwell)上の[manager](https://github.com/gravwell/manager)アプリケーションをオープンソース化しています。ですから、もしあなたが非常に小さくて簡単に設定できるSystemDのようなサービスマネージャをdockerコンテナ用に使いたいのであれば、ぜひ使ってみてください。

公式のSoliton NK Dockerイメージには、Simple Relayインジェスターだけでなく、フルSoliton NKスタック(インデクサーとウェブサーバー)も含まれています。デフォルトのマネージャ設定は次の通りです:

```
[Global]
	Log-File=/opt/gravwell/log/manager.log
	Log-Level=INFO

[Error-Handler]
	Exec=/opt/gravwell/bin/crashReport

[Process "indexer"]
	Exec="/opt/gravwell/bin/gravwell_indexer -stderr indexer"
	Working-Dir=/opt/gravwell
	Max-Restarts=3 #three attempts before cooling down
	CoolDown-Period=60 #1 hour
	Restart-Period=10 #10 minutes

[Process "webserver"]
	Exec="/opt/gravwell/bin/gravwell_webserver -stderr webserver"
	Working-Dir=/opt/gravwell
	Max-Restarts=3 #three attempts before cooling down
	CoolDown-Period=30 #30 minutes
	Restart-Period=10 #10 minutes

[Process "searchagent"]
	Exec="/opt/gravwell/bin/gravwell_searchagent -stderr searchagent"
	Working-Dir=/opt/gravwell
	Max-Restarts=3 #three attempts before cooling down
	CoolDown-Period=10 #10 minutes
	Restart-Period=10 #10 minutes

[Process "simple_relay"]
	Exec="/opt/gravwell/bin/gravwell_simple_relay -stderr simple_relay"
	Working-Dir=/opt/gravwell
	Max-Restarts=3 #three attempts before cooling down
	CoolDown-Period=10 #10 minutes
	Restart-Period=10 #10 minutes
```

このマネージャアプリケーションのデフォルト設定では、バグの特定と修正に役立つエラー報告システムを有効にしています。サービスがゼロ以外の終了コードで終了した場合、エラーレポートを取得します。エラー報告システムを無効にするには、"[Error-Handler]" セクションを削除するか、環境変数 "DISABLE_ERROR_REPORTING" に "TRUE" を指定してください。

サービスを個別に無効にする場合は、該当サービス名をすべて大文字にしてその前に"DISABLE_"を付けた環境変数に、"TRUE"の値を指定したものを、起動コマンドのオプションに追加してください。

例えば、エラー報告をせずに Soliton NK docker コンテナを起動するには、"-e DISABLE_ERROR_REPORTING=true" オプションを指定して起動します。

インデクサーは起動するけれども、統合されたSimpleRelayインジェスターを無効にしたい場合には、"-e DISABLE_SIMPLE_RELAY=TRUE "を以下のように起動オプションに追加してください。:

```
docker run --name gravwell -e GRAVWELL_INGEST_SECRET=MyIngestSecret -e DISABLE_SIMPLE_RELAY=TRUE -e DISABLE_WEBSERVER=TRUE -e DISABLE_SEARCHAGENT=TRUE soliton/solitonnk:latest
```

サービスマネージャの詳細については、[githubのページ](https://github.com/gravwell/manager)を参照してください。

### インジェスターコンテナのカスタマイズ

インジェスターコンテナを起動した後に、デフォルト設定を多少変更したくなることもあるかもしれません。たとえば、Netflow インジェスターを別のポートで実行させようと考えたとしましょう。

起動した Netflow インジェスターコンテナに変更を加えるには、コンテナ内でシェルを起動します:

	docker exec -it netflow sh

次に vi を使って `/opt/gravwell/etc/netflow_capture.conf` を [インジェスターのドキュメント](#!ingesters/ingesters.md) で説明されているように編集することができます。変更を加え終えたら、コンテナ全体を再起動するだけです:

	docker restart netflow

## （Docker上ではない）外部のインジェスターの設定

`soliton/solitonnk` イメージを起動するのに使ったオリジナルのコマンドをもう一度見直すと、ポート4023と4024をホストに転送したことに気づくでしょう。これらはそれぞれインデクサーの平文とTLS暗号文のインジェスト受信ポートです。別のシステムでインジェスターを実行している場合(おそらくどこかのLinuxサーバーでログファイルを収集したりしてるでしょう)、インジェスター設定ファイルの `Cleartext-Backend-target` または `Encrypted-Backend-target` フィールドにDockerホストを指定すれば、そのDockerホストで動いてるSoliton NKインスタンスにデータをインジェストすることができます。

インジェスターの設定の詳細については、[インジェスターのドキュメント](#!ingesters/ingesters.md)を参照してください。

## セキュリティ上の考慮事項

転送されたコンテナポートをインターネットに公開する場合は、以下で挙げているの設定値を安全な値に設定することが**極めて重要**です:

* 'admin' のパスワードはデフォルトの 'changeme' から変更する必要があります。
* 環境変数のGRAVWELL_INGEST_SECRET、GRAVWELL_INGEST_AUTH、GRAVWELL_CONTROL_AUTH、およびGRAVWELL_SEARCHAGENT_AUTHの値は、インデクサー＆ウェブサーバー（上記参照）の起動時に複雑な文字列に設定しなければなりません。

## クラッシュレポートとメトリクス

Soliton NKソフトウェアには、自動化されたクラッシュレポートとメトリクスレポートが組み込まれています。どこにどのような情報が送られるか、および、オプトアウトする方法の詳細については、[クラッシュレポートとメトリクスのページ](#!metrics.md)を参照してください。

## その他の情報源

Soliton NKを使っていて、さらに詳しい使い方を知りたい場合には[他のドキュメント](#!index.md)を探してみてください。

Soliton NKの正規ユーザーで、DockerでSoliton NKをデプロイしたい場合は、[サポート](https://www.soliton.co.jp/support/) を通じて連絡してください。
