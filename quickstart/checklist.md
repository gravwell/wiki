# インストールチェックリスト

# # シングルノードチェックリスト

このチェックリストは、シングルノード、スタンドアロンのGravwellインスタンスを設定するための一般的な操作手順を示します。手順の詳細は[クイックスタート](quickstart.md)を参照してください。

自己解凍型インストーラ、Debian/Redhatパッケージ、またはDockerコンテナでGravwellをインストールしてください（[クイックスタート](quickstart.md)参照）。

□ ファイアウォールで[Gravwellが使用するポート](#!configuration/networking.md)への着信を許可することを確認してください。

□ Webブラウザで新しいGravwellにアクセスし、`http://gravwell.example.org/`のように表示されたらライセンスファイルをアップロードしてください。

□ 管理者（admin）とパスワード（changeme）でログインします。デフォルトのパスワードは右上のユーザーアイコンから変更できます。

□ 必要に応じて追加のストレージウェルを設定します。[Gravwell configuration](#!configuration/configuration.md) および [detail configuration parameters](#!configuration/parameters.md) を参照してください。

□ ディスク容量が不足しないように、ウェルに[ageout](#!configuration/ageout.md)を設定してください。

□ オプションです。ユーザーのアクセスとインゲスターの接続に[TLSの設定](#!configuration/certificates.md)を行います。

□ [インジェスターの設定](#!ingesters/ingesters.md)でデータをGravwellに取り込みます。


<!-- TODO: これは複雑なプロセスで、使用しているかどうかわからないオプションがたくさんあるため、直線的なチェックリストで捉えるのは困難です。少なくともいくつかのステップを集めているので、ここに残しておきます。
## クラスターチェックリスト

### 準備

□ どのノードがインデクサーになり、どのノードがウェブサーバーになるかを決定します。複数の Web サーバを導入する場合は、検索エージェントを実行する Web サーバを 1 つ選択します。

□ 分散フロントエンド](#!distributed/frontend.md)を使用する場合は、*データストア*用に追加のシステムを用意してください。データストアは、インデクサやウェブサーバプロセスと同居することはできないので注意が必要です。

□ Webサーバとインデクサの各ノードにGravwellをインストールしてください（[クイックスタート](quickstart.md)参照）。

□ 必要に応じてデータストアをインストールしてください。これはコアシェルインストーラーに含まれているが、Debian と Redhat では別のパッケージになっています。

□ 必要に応じてロードバランサーをインストールしてください。

□ TLS 証明書を Web サーバ、データストア、ロードバランサーに適宜導入してください。証明書を `/opt/gravwell/etc/cert.pem` に、秘密鍵を `/opt/gravwell/etc/key.pem` にコピーすることをお勧めします。

### 構成

□ 設定のベースとなるノードの `gravwell.conf` ファイルをコピーします。Webserver-UUID "行や "Indexer-UUID "行を削除してください。

#### Indexer Config

□ インデクサーに使用するコンフィグのコピーを作成してください。

□ インデクサーのコンフィグに必要なWellを定義してください。([このドキュメント](#!configuration/configuration.md)を参照してください）

□ 各ウェルに[ageout configuration](#!configuration/ageout.md)を設定してください。

#### ウェブサーバの設定

ウェブサーバに使用するベースコンフィグのコピーを作成します。

□ `Remote-Indexers` パラメータを設定して、予定されているすべてのインデクサーをリストアップします。
```
Remote-Indexers=ネット:インデクサー0.example.net:9404
Remote-Indexers=ネット:インデクサー1.example.net:9404
2.Remote-Indexers=net:indexer2.example.net:9404
```

□ データストアを使用する場合は、[distributed frontends](#!distributed/frontend.md)に記載されている通り、gravwell.confの`Datastore`と`External-Addr`オプションを設定してください。

□ [TLS](#!configuration/certificates.md)の`Certificate-File`と`Key-File`フィールドを設定してください。

### デプロイメント

□ systemd を使用して必要のない Gravwell プロセスを無効にしてください。インデクサでは Web サーバと searchagent を、Web サーバでは indexer を無効にしてください。インデクサでは webserver と searchagent を、web サーバでは indexer を無効にしてください。searchagent プロセスは 1 つの web サーバでのみ有効にしてください。

□ インデクサーの設定をインデクサーに、ウェブサーバーの設定をウェブサーバーにコピーしてください。

□ 全ノードの gravwell プロセスを再起動してください。
-->