# 設定パラメータ

設定ファイル `gravwell.conf` は、インデクサ、ウェブサーバ、データストアの設定のために使用されます。設定ファイルには*セクション*が含まれ、それらは角括弧内で定義されます(例: `[global]`); 各セクションにはパラメータが含まれます。以下の例では、`[global]`全体セクション、`[Replication]`レプリケーションセクション、`[Default-Well]`デフォルトウェルセクション、および`[Storage-Well]`ストレージウェルセクションが含まれています:

```
[global]
### Authentication tokens
Ingest-Auth=IngestSecrets
Control-Auth=ControlSecrets
Search-Agent-Auth=SearchAgentSecrets

### Web server HTTP/HTTPS settings
Web-Port=80
Insecure-Disable-HTTPS=true

### Other web server settings
Remote-Indexers=net:127.0.0.1:9404
Persist-Web-Logins=True
Session-Timeout-Minutes=1440
Login-Fail-Lock-Count=4
Login-Fail-Lock-Duration=5

### Ingester settings
Ingest-Port=4023
Control-Port=9404
Search-Pipeline-Buffer-Size=4

### Other settings
Log-Level=WARN

### Paths
Pipe-Ingest-Path=/opt/gravwell/comms/pipe
Log-Location=/opt/gravwell/log
Web-Log-Location=/opt/gravwell/log/web
Render-Store=/opt/gravwell/render
Saved-Store=/opt/gravwell/saved
Search-Scratch=/opt/gravwell/scratch
Web-Files-Path=/opt/gravwell/www
License-Location=/opt/gravwell/etc/license
User-DB-Path=/opt/gravwell/etc/users.db
Web-Store-Path=/opt/gravwell/etc/webstore.db

[Replication]
	Peer=10.0.0.1
	Storage-Location=/opt/gravwell/replication_storage
	Insecure-Skip-TLS-Verify=true
	Disable-TLS=true
	Connect-Wait-Timeout=20

[Default-Well]
	Location=/opt/gravwell/storage/default/
	Accelerator-Name=fulltext #fulltext is the most resilient to varying data types
	Accelerator-Engine-Override=bloom #The bloom engine is effective and fast with minimal disk overhead
	Disable-Compression=true

[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Tags=syslog
	Tags=kernel
	Tags=dmesg
	Accelerator-Name=fulltext #fulltext is the most resilient to varying data types
	Accelerator-Args="-ignoreTS" #tell the fulltext accelerator to not index timestamps, syslog entries are easy to ID
```

各パラメータにはデフォルト値があり、パラメータが空であるか、設定ファイルで指定されていない場合に適用されます。

## Global（全体）のパラメータ

このセクションのパラメーターは、設定ファイルの`[Global]`見出しの下にあります。これは通常、ファイルの先頭にあります。

**Indexer-UUID**
適用対象:インデクサー
デフォルト値:[設定がない時には、値がランダム生成される]
設定例:`Indexer-UUID="ecdeeff8-8382-48f1-a24f-79f83af00e97"`
設定内容:インデクサ毎に一意の識別子を設定してください。二つのインデクサが同じ UUID にならないようにしなければいけません。このパラメータが設定されていない場合、インデクサは UUID を生成して設定ファイルに書き込みます。 インデクサが致命的な障害を起こしてレプリケーションから再構築するという場合でない限り、通常はこれが最良の選択です ( [レプリケーションドキュメント](configuration/replication.md) を参照してください)。このパラメータを変更しようとする時には熟慮が必要です。

**Webserver-UUID**
適用対象:ウェブサーバ－
デフォルト値:[設定がなされてない時には、値がランダム生成されます]
設定例:`Webserver-UUID="ecdeeff8-8382-48f1-a24f-79f83af00e97"`
設定内容:ウェブサーバ毎に一意の識別子を設定してください。2つのウェブサーバが同じUUIDにならないようにしなければいけません。このパラメータが設定されていない場合、ウェブサーバは UUID を生成して*設定ファイルに書き込みます*。通常はこれが最良の選択です。このパラメータを変更しようとする時には熟慮が必要です。

**License-Location**
適用対象:インデクサーとウェブサーバー
デフォルト値:`/opt/gravwell/etc/license`
設定例:`License-Location=/opt/gravwell/etc/my_license`
設定内容:gravwell ライセンスファイルへのパスを設定してください。パスは、gravwell のユーザーとグループが読めるようにする必要があります。

**Config-Location**
適用対象:インデクサーとウェブサーバー
デフォルト値:`/opt/gravwell/etc`
設定例:`Config-Location=/tmp/path/to/etc`
設定内容:`Config-Location`では、Config-Location以外の全ての設定パラメータ記述を収められる代替場所を指定することができます。 代替の Config-Location を指定すると、他のすべてのパラメータを代替パスで指定しなくても、1 つのパラメータを設定するだけで済みます。

**Web-Port**
適用対象:ウェブサーバー
デフォルト値:`443`
設定例:`Web-Port=80`
設定内容:Web サーバのリッスンポートを指定します。Web-Port パラメータを 80 に設定しても、ウェブサーバは HTTP のみのモードにはならないことに注意してください。HTTP のみのモードにするには、Insecure-Disable-HTTPSパラメータを設定してください。

**Disable-HTTP-Redirector**
適用対象:ウェブサーバー
デフォルト値:`False`
設定例:`Disable-HTTP-Redirector=true`
設定内容:デフォルトでは、GravwellはHTTPリダイレクタを起動し、プレーンテキスト通信のHTTPポータルへのアクセスをしようとするウェブブラウザに対して、暗号化されたHTTPSポータルへリダイレクトします。`Disable-HTTP-Redirector=true`を設定すると、このリダイレクトはなされません。

**Insecure-Disable-HTTPS**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Insecure-Disable-HTTPS=true`
設定内容:デフォルトでは、GravwellはHTTPSモードで動作します。`Insecure-Disable-HTTPS=true`を設定すると、Gravwellは代わりにプレーンテキスト通信のHTTPを使用し、`Web-Port`をリッスンするようになります。

**Webserver-Domain**
適用対象:ウェブサーバ－
デフォルト値:0
設定例:`Webserver-Domain=17`
設定内容: `Webserver-Domain` パラメーターによってウェブサーバが属する [resources](#!resources/resources.md) ドメインを制御できます。2つのウェブサーバーが同じドメインで設定されていて、異なるリソースセットを持ち、同じインデクサに接続されていて、データストア経由で同期化されて*いない*場合、インデクサーは2つのリソースセットの間でバタバタしてしまいます。ウェブサーバーを別のドメインに置くことで、リソースの競合なしに同じインデクサーを使うことができます。

**Control-Listen-Address**
適用対象:インデクサー
デフォルト値:
設定例:`Control-Listen-Address=”10.0.0.1”`
設定内容:`Control-Listen-Address`パラメータは、インデクサがウェブサーバからの制御コマンドをリッスンするIPアドレスにバインドできます。 デュアルホームマシン、または高速データネットワークと低速制御ネットワークを持つマシンへのGravwellインストールでは、トラフィックが適切にルーティングされることを確実にするために、インデクサーを特定のアドレスにバインドすることが推奨されます。

**Control-Port**
適用対象:インデクサーとウェブサーバー
デフォルト値:`9404`
設定例:`Control-Port=12345`
設定内容:`Control-Port`パラメーターでは、インデクサーがウェブサーバーからの制御コマンドをリッスンするポートを選択します。インデクサーとウェブサーバーが通信するためには、この設定値は、インデクサーとウェブサーバーとで同じでなければなりません。インストーラーはデフォルトでインデクサーにポートのバインド機能を付与しないので、ポートは 1024 より大きい値に設定する必要があります。 複数のインデクサーが単一のマシン上で実行されている環境や、別のアプリケーションがポート9404にバインドされている場合は、制御ポートを調整する必要があるでしょう。

**Datastore-Listen-Address**
適用対象:データストア
デフォルト値:
設定例:`Datastore-Listen-Address=10.0.0.1`
設定内容:`Datastore-Listen-Address`パラメータは、特定のアドレスのみをリッスンするようにデータストアに指示します。デフォルトでは、データストアはシステム上のすべてのアドレスをリッスンします。

**Datastore-Port**
適用対象:データストア
デフォルト値:`9405`
設定例:`Datastore-Port=7888`
設定内容:`Datastore-Port`パラメータは、データストアが通信するポートを選択します。ポートは 1024 より大きくする必要があります。デフォルト値の9405は、通常、ほとんどのインストールに適しています。

**Datastore**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`Datastore=10.0.0.1:9405`
設定内容:`Datastore`パラメータでは、ウェブサーバがダッシュボード、リソース、ユーザー設定、検索履歴を同期させるために必要なデータストアへの接続を指定します。これを設定すると、[分散型ウェブサーバ](distributed/frontend.md)を使用することができますが、複数のウェブサーバを用いなければならない場合にのみ設定してください。デフォルトでは、ウェブサーバはデータストアに接続しません。

**Datastore-Update-Interval**
適用対象:ウェブサーバ－
デフォルト値:`10`
設定例:`Datastore-Update-Interval=5`
設定内容:`Datastore-Update-Interval` パラメーターは、ウェブサーバーがデータストアの更新をチェックするまでにどれくらいの時間 (秒単位で) 待つべきかを決定します。デフォルト値の10秒が一般的に適しています。

**Datastore-Insecure-Disable-TLS**
適用対象::データストアとデータストア
デフォルト値:`false`
設定例:`Datastore-Insecure-Disable-TLS=true`
設定内容:`Datastore-Insecure-Disable-TLS` パラメータを設定すると、ウェブサーバーとデータストアの両方で使用されます。デフォルトでは、データストアはウェブサーバーからの HTTPS 接続の着信をリッスンします。このパラメーターを true に設定すると、データストアはプレーンテキストの HTTP を期待し、ウェブサーバーに HTTP を使用するように指示します。

**Datastore-Insecure-Skip-TLS-Verify**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Datastore-Insecure-Skip-TLS-Verify` パラメーターは、データストアへの接続時にTLS証明書が無効であってもそのことを無視するようにウェブサーバーに指示します。この設定は自己署名証明書を使用する場合に必要ですが、可能な限り避けるべきです。

**External-Addr**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`External-Addr=10.0.0.1:443`
設定内容:`External-Address` パラメーターは、他のウェブサーバ－がこのウェブサーバ－に接続するために使用するアドレスを指定します。このパラメーターは、あるウェブサーバ－のユーザーが別のウェブサーバ－で実行された検索の結果をロードできるようにするため、データストアを使用する場合は**必須**です。


**Search-Forwarding-Insecure-Skip-TLS-Verify**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Search-Forwarding-Insecure-Skip-TLS-Verify=true`
設定内容:このパラメーターは、データストアを使用して分散モードで複数のウェブサーバ－を操作する場合にのみ役立ちます。ウェブサーバ－に自己署名証明書がある場合、このパラメーターがtrueに設定されていない限り、ユーザーはリモートウェブサーバ－からの検索にアクセスできません。

**Ingest-Port**
適用対象:インデクサー
デフォルト値:`4023`
設定例:`Ingest-Port=14023`
設定内容:`Ingest-Port` パラメータは、インデクサーがインジェスターとの接続のためにリッスンするポートを制御します。 `Ingest-Port` パラメーターを変更すると、単一のマシンで複数のインデクサーを実行している場合や、別のアプリケーションがすでにデフォルトのポート 4023 にバインドされている場合に便利です。

**TLS-Ingest-Port**
適用対象:インデクサー
デフォルト値:`4024`
設定例:`TLS-Ingest-Port=14024`
設定内容:`TLS-Ingest-Port` パラメーターは、インデクサーがインジェスターとの接続のためにリッスンするポートを制御します。 `TLS-Ingest-Port` パラメーターを変更すると、1 台のマシンで複数のインデクサーを実行している場合や、別のアプリケーションがすでにデフォルトのポート 4024 にバインドされている場合に便利です。 デフォルトでは、TLS トランスポートを使用するすべてのインジェスターがリモート証明書を検証します。 デプロイの際に自動生成された証明書を使用している場合、インジェスターは証明書を信頼済みとしてインストールしておくか、証明書の検証を無効にする必要があります(これは TLS トランスポートが提供する保護の効果を損ないます)。

**Pipe-Ingest-Path**
適用対象:インデクサー
デフォルト値:`/opt/gravwell/comms/pipe`
設定例:`Pipe-Ingest-Path=/tmp/path/to/pipe`
設定内容:`Pipe-Ingest-Path` は Unix の名前付きパイプへのフルパスを指定します。 インデクサーが名前付きパイプを作成し、同じノードにあるインジェスターがそのパイプに接続して、非常に高速で低遅延のトランスポートとして使用することができます。 名前付きパイプは、1ギガビット以上で動作するネットワークパケットインジェスターのように、非常に高い性能を必要とするインジェスターに最適です。 ネームドパイプは、通常とは異なるネットワークトランスポートや、超高速の非IPベースの相互接続を介したトランスポートを容易にするためにも使用できます。

**Log-Location**
適用対象:インデクサー
デフォルト値:`/opt/gravwell/log`
設定例:`Log-Location=/tmp/path/to/logs`
設定内容:`Log-Location` パラメータは、Gravwell インフラストラクチャが自身のログを置く場所を制御します。 Gravwellは自身のログを直接インデクサーに供給せず、代わりにファイルに書き込みます(Gravwellのログもインジェストしたい場合は、file followerインジェスターを使用してください)。 このパラメータは、それらのログがどこに置かれるかを指定します。

**Web-Log-Location**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/log/web`
設定例:`Web-Log-Location=/tmp/path/to/logs/web`
設定内容:`Web-Log-Location` パラメーターは、ウェブサーバーのログが保存される場所を制御します。 Gravwellは自分のログを直接インデクサーに供給せず、代わりにファイルに書き込みます(Gravwellのログもインジェストしたい場合はfile followerインジェスターを使用してください)。 このパラメータはログの保存先を指定します。

**Datastore-Log-Location**
適用対象:データストア
デフォルト値:`/opt/gravwell/log/datastore`
設定例:`Datastore-Log-Location=/tmp/path/to/logs/datastore`
設定内容:`Datastore-Log-Location` パラメーターは、データストアのログが保存される場所を制御します。

**Log-Level**
適用対象:インデクサーとデータストアとウェブサーバー
デフォルト値:`INFO`
設定例:`Log-Level=ERROR`
設定内容:`Log-Level` パラメーターは、gravwell インフラストラクチャからのログの詳細レベルを制御します。 Log-Levelには3種類の引数値、INFO、WARN、ERRORのいずれかを設定できます。 INFO は最も詳細で、ERROR は最も簡便です。 ロギングシステムは、詳細レベルごとにログファイルを生成し、syslogデーモンと同様の方法でそれらをローテーションさせます。

**Disable-Access-Log**
適用対象:ウェブサーバー
デフォルト値:`false`
設定例:`Disable-Access-Log=true`
設定内容:`Disable-Access-Log` パラメータは、ウェブサーバーのアクセスログ生成を無効にするために使用されます。 アクセスログ機能は、個々のページアクセスを記録します。これらのアクセスログを持つことは、Gravwellアクセスを監査し、潜在的な問題をデバッグするために上で一般的には価値がありますが、多くのユーザーがいる環境ではアクセスログがあまりに大きくなる可能性があるので、それらを無効にすることが望ましいこともあるでしょう。

**Disable-Self-Ingest**
適用対象：ウェブサーバー、インデクサー
デフォルト値：`false`
設定例: `Disable-Self-Ingest=true`
設定内容：Disable-Self-Ingest パラメーターは、ウェブサーバーとインデクサーがログを`gravwell`タグに取り込まないようにします。

**Persist-Web-Logins**
適用対象:ウェブサーバー
デフォルト値:`true`
設定例:`Persist-Web-Logins=false`
設定内容:`Persist-Web-Logins` パラメータは、シャットダウン時にユーザーセッションを不揮発性ストレージに保存することをウェブサーバーに通知するために使用されます。 デフォルトでは、ウェブサーバーがシャットダウンまたは再起動された場合も、クライアントセッションが保持されます。 `Persist-Web-Logins` をFalseに設定すると、ウェブサーバが再起動される度にセッションが無効になります。

**Session-Timeout-Minutes**
適用対象:ウェブサーバー
デフォルト値:`60`
設定例:`Session-Timeout-Minutes=1440`
設定内容:`Session-Timeout-Minutes` パラメータは、ウェブクライアントから何も操作がなされないままでもウェブサーバがセッションを破棄しないで待つ時間を制御します。 例えば、クライアントがログアウトせずにブラウザを閉じた場合、システムは指定された時間だけ待ってからセッションを無効にします。 （何も指定が無い場合の）デフォルト値は60分ですが、インストーラは特に指定しないでもこの値を 1 日に設定ます。

**Key-File**
適用対象:インデクサーとデータストアとウェブサーバー
デフォルト値:`/opt/gravwell/etc/key.pem`
設定例:`Key-File=/opt/gravwell/etc/privkey.pem`
設定内容:`Key-File` パラメーターは、ウェブサーバー、データストア、インデクサーに対してどのファイルを秘密鍵として使用するかを指定します。 秘密鍵/公開鍵は PEM 形式でエンコードされていなければなりません。 秘密鍵は秘密保持されていなければならず、もしも漏洩した場合は破棄して再発行する必要があります。 詳細は http://www.tldp.org/HOWTO/SSL-Certificates-HOWTO/x64.html を参照してください。

**Certificate-File**
適用対象:インデクサーとデータストアとウェブサーバー
デフォルト値:`/opt/gravwell/etc/cert.pem`
設定例:`Certificate-File=/opt/gravwell/etc/cert.pem`
設定内容:`Certificate-File` パラメータは、TLS 通信に使用される公開鍵/秘密鍵ペアの公開鍵コンポーネントを指定します。 公開鍵は、機密情報ではなく、すべてのインジェスターとウェブクライアントに配信されます。 Gravwellは、ここで指定された公開鍵がPEM形式でエンコードされ、公開鍵部分の情報のみ含まれていることを前提にしています。

**Ingest-Auth**
適用対象:インデクサー
デフォルト値:`IngestSecrets`
設定例:`Ingest-Auth=abcdefghijklmnopqrstuvwxyzABCD`
設定内容:`Ingest-Auth` パラメーターは、インジェスターをインデクサーに対して認証させるために使用する共有鍵トークンを指定します。 このトークンは任意の長さにすることが可能ですが、Gravwellでは、少なくとも24文字以上の高エントロピートークンを推奨しています。 デフォルトでは、インストーラーはトークンをランダム生成します。

**Control-Auth**
適用対象:インデクサーとウェブサーバー
デフォルト値:`ControlSecrets`
設定例:`Control-Auth=abcdefghijklmnopqrstuvwxyzABCD`
設定内容:`Control-Auth` パラメータは、インジェスターとウェブサーバーの相互の認証に使用される共有鍵トークンを指定します。 このトークンは任意の長さにすることが可能ですが、Gravwellでは、少なくとも24文字以上の高エントロピートークンを推奨しています。 デフォルトでは、インストーラーはトークンをランダム生成します。

**Search-Agent-Auth**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`Search-Agent-Auth=abcdefghijklmnopqrstuvwxyzABCD`
設定内容:`Search-Agent-Auth` パラメータは、ウェブサーバ－に対して検索エージェントを認証させるために使用する共有鍵トークンを指定します。デフォルトでは、インストーラーはトークンをランダム生成します。

**Web-Files-Path**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/www`
設定例:`Web-Files-Path=/tmp/path/to/www`
設定内容:`Web-Files-Path` は、ウェブサーバ－によって提供されるフロントエンド GUI ファイルを含むパスを指定します。 このウェブサーバ用ファイルには、Webページの表示とWebブラウザを介したGravwellシステムとのやりとりを担当するすべてのGravwellコードが含まれています。

**Tag-DB-Path**
適用対象:インデクサー
デフォルト値:`/opt/gravwell/etc/tags.db`
設定例:`Tag-DB-Path=/tmp/path/to/tags.db`
設定内容:`Tag-DB-Path` パラメータは、タグデータベースの場所を指定します。このファイルでは、インデクサの数値タグIDがタグ名文字列にマッピングされています。

**User-DB-Path**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/etc/users.db`
設定例:`User-DB-Path=/tmp/path/to/users.db`
設定内容:`User-DB-Path` パラメーターは、ユーザーデータベースファイルの場所を指定します。 ユーザーデータベースファイルには、ユーザーとグループの設定が含まれています。 ユーザーデータベースはパスワードの保存と検証に bcrypt ハッシュアルゴリズムを使用します。 bcryptによるハッシュシステムは非常に堅牢であると考えられますが、users.dbファイルはそれでもなお秘密保持されている必要があります。デフォルトでは、インストーラーはユーザーデータベースファイルのファイルシステムパーミッションを設定して、Gravwell のユーザーとグループのみが読めるようにします。


**Datastore-User-DB-Path**
適用対象:データストア
デフォルト値:`/opt/gravwell/etc/datastore-users.db`
設定例:`Datastore-User-DB-Path=/tmp/path/to/datastore-users.db`
設定内容:`Datastore-User-DB-Path` パラメーターは、データストア・コンポーネントによって管理されるユーザー・データベース・ファイルの場所を指定します。これは、User-DB-Path パラメータで指定されたパスと**同じパスであってはなりません**！

**Web-Store-Path**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/etc/webstore.db`
設定例:`Web-Store-Path=/tmp/path/to/webstore.db`
設定内容:`Web-Store-Path` は、検索履歴、ダッシュボード、ユーザー設定、ユーザーセッション、およびその他の雑多なユーザーデータを保存するために使用されるデータベースファイルを指定します。 ウェブストアデータベースファイルにはユーザーの資格情報は含まれていませんが、ユーザーセッションクッキーと CSRF トークンは*含まれています*。 Gravwell はクッキーと CSRF トークンをアクセス元に関連付けているので、攻撃者が盗んだクッキーやトークンを再利用するリスクは低いものの、それでもデータストアは保護する必要があります。 インストーラーは、ファイルシステムのパーミッションを設定して、Gravwell ユーザによる読み取り/書き込みのみを許可するようにします。

**Datastore-Web-Store-Path**
適用対象:データストア
デフォルト値:`/opt/gravwell/etc/datastore-webstore.db`
設定例:`Datastore-Web-Store-Path=/tmp/path/to/datastore-webstore.db`
設定内容:`Datastore-Web-Store-Path` パラメータは、検索履歴、ダッシュボード、およびユーザー設定を保存するためにデータストアが使用するデータベース・ファイルを指定します。これは、Web-Store-Pathパラメータで指定されたパスと**同じパスであってはなりません**！

**Web-Listen-Address**
適用対象:ウェブサーバー
デフォルト値:
設定例:`Web-Listen-Address=10.0.0.1`
設定内容: `Web-Listen-Address` パラメーターは、ウェブサーバがバインドしてサービスを提供するアドレスを指定します。 デフォルトでは、このパラメータは空で、ウェブサーバはすべてのインターフェイスとアドレスとバインドします。

**Login-Fail-Lock-Count**
適用対象:ウェブサーバー
デフォルト値:`5`
設定例:`Login-Fail-Lock-Count=10`
設定内容:`Login-Fail-Lock-Count` パラメーターでは、ブルートフォース保護が有効になるまでに許される、ユーザーアカウントに対する連続ログイン失敗の回数を指定します。 例えば、値が4に設定されていて、ユーザーが不正なパスワードを連続して4回入力した場合、その後はログイン試行に対する応答完了までに時間がかかるようになり、攻撃者の攻撃再試行により時間が要するようになります。
注: Gravwellは以前は、特定の失敗回数の後にアカウントをロックしていましたが、現在はそこまで過激ではないブルートフォース保護が有効になっています。とはいうものの、レガシーな理由から設定パラメータには「Lock」という名前が残っています。

**Login-Fail-Lock-Duration**
適用対象:ウェブサーバー
デフォルト値:`5`
設定例:`Login-Fail-Lock-Duration=10`
設定内容:`Login-Fail-Lock-Duration` パラメータは、Login-Fail-Lock-Countの保護期間が終わるまでの、保護時間幅(分単位)を指定します。
注: Gravwellは以前は、特定の失敗回数の後にアカウントをロックしていましたが、現在はそこまで過激ではないブルートフォース保護が有効になっています。とはいうものの、レガシーな理由から設定パラメータには「Lock」という名前が残っています。

**Remote-Indexers**
適用対象:ウェブサーバー
デフォルト値:`net:10.0.0.1:9404`
設定例:`Remote-Indexers=net:10.0.0.1:9404`
設定内容:`Remote-Indexers` パラメーターは、ウェブサーバーが接続して制御するリモートインデクサーのアドレスとポートを指定します。 `Remote-Indexers` は列挙可能パラメータであり、つまり、複数のリモートインデクサーを提供するために何度も指定できます。Gravwell Cluster 版では、クラスタ内の各インデクサーを指定する必要があります。 値の先頭に“net:”がプレフィックスとして付いている場合、リモートインデクサがネットワークトランスポート経由でアクセス可能であることを示します。Gravwellの特別な版は他の通信方式を使うことができますが、ほとんどの商用版では“net:”を使うことが前提となっています。

**Search-Scratch**
適用対象:インデクサーとウェブサーバー
デフォルト値:`/opt/gravwell/scratch`
設定例:`Search-Scratch=/tmp/path/to/scratch`
設定内容:`Search-Scratch` パラメーターは、検索モジュールがアクティブな検索中に一時ストレージとして使用できるストレージの場所を指定します。 検索モジュールによっては、メモリの制約により一時的なストレージを使用する必要がある場合があります。 例えば、ソートモジュールは5GBのデータをソートする必要があるかもしれませんが、そんな場合でも物理マシンには4GBの物理RAMしかないような場合もあるでしょう。 モジュールが、ホストのスワップ（これはソートの機能だけでなく、すべてのモジュールにペナルティをもたらす）を呼び出すことなく、大規模なデータセットをソートするために、スクラッチスペースをインテリジェントに使用することができます。 各検索が終われば、このスクラッチスペースは破棄されます。

**Render-Store**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/render`
設定例:`Render-Store=/tmp/path/to/render`
設定内容:`Render-Store` パラメーターは、レンダラーモジュールが検索結果を保存する場所を指定します Render-Storeの場所は一時的な保存場所であり、通常は適度に小さいデータセットを表します。検索がアクティブに実行されているか、クライアントとの対話で待機状態の場合、レンダラーはRender-Storeにデータセットを保存したり、そこからデータセットを取得したりします。Render-Storeは、フラッシュベースまたはXPointSSDなどの高速ストレージに配置する必要があります。検索が中止されると、レンダーストアは削除されます（検索が保存されていない場合）。


**Saved-Store**
適用対象:ウェブサーバー
デフォルト値:`/opt/gravwell/saved`
設定例:`Saved-Store=/path/to/saved/searches`
設定内容:Saved-Storeパラメーターは、保存された検索が保存される場所を指定します。保存された検索は検索の出力状態を表し、ユーザーが検索を再開せずに後で検索結果を再度参照できるようにしたい監査や状況に役立ちます。保存された検索は明示的に削除する必要があり、データはシャードエージングアウトポリシーの対象ではありません。保存された検索は完全にアトミックです。つまり、保存された検索の基になるデータは完全に古くなり、削除されても、ユーザーは保存された検索を再度開いて調べることができます。保存された検索は共有することもできます。つまり、ユーザーは保存された検索をまとめて、Gravwellの他のインスタンスと共有できます。

**Search-Pipeline-Buffer-Size**
適用対象:インデクサーとウェブサーバー
デフォルト値:`2`
設定例:`Search-Pipeline-Buffer-Size=8`
設定内容:Search-Pipeline-Buffer-Sizeは、検索中に各モジュール間で転送できるブロックの数を指定します。 サイズが大きいほどバッファリングが良くなり、常駐メモリの使用量を犠牲にしても高スループットの検索が可能になります。 インデクサはパイプラインのサイズに敏感ですが、システムが自由にメモリを退避・再設定できる共有メモリ技術を使用しています。ウェブサーバは通常、パイプラインを移動する際にすべてのエントリを常駐させ、モジュールを凝縮してメモリ負荷を軽減します。 回転ディスクのような待ち時間の長いストレージシステムを使用している場合は、このバッファサイズを大きくすることが有利になります。
このパラメータを大きくすると、検索のパフォーマンスが向上しますが、システムが一度に処理できる実行中の検索数に直接影響を与えます。 ビデオフレーム、PE実行ファイル、オーディオファイルなどの非常に大きなエントリを保存している場合は、バッファサイズを小さくして常駐メモリの使用量を制限する必要があります。ホストカーネルがOut Of Memory (OOM)を起動してGravwellプロセスを強制終了させた場合、これが最初に回すべきツマミとなる。

**Search-Relay-Buffer-Size**
適用対象:ウェブサーバ－
デフォルト値:`4`
設定例:`Search-Relay-Buffer-Size=8`
設定内容:Search-Relay-Buffer-Sizeパラメーターは、別のインデクサーからの未処理のブロックを待機している間に、ウェブサーバ－が各インデクサーから受け入れるエントリブロックの数を制御します。検索エントリが一時的に流入するため、1つのインデクサがまだ古いエントリを処理している一方で、別のインデクサがより新しいエントリに進んでいる可能性があります。ウェブサーバ－はエントリを時間順に処理する必要があるため、遅いインデクサーが追いつくのを待つ間、「先行」しているインデクサーからのエントリをバッファリングします。一般に、デフォルト値は、許容可能なパフォーマンスを提供しながら、メモリの問題を防ぐのに役立ちます。大量のメモリを搭載したシステムでは、この値を増やすと便利な場合があります。

**Max-Search-History**
適用対象:ウェブサーバー
デフォルト値:`100`
設定例:`Max-Search-History=256`
設定内容:Max-Search-Historyパラメーターは、ユーザーに対して保持される検索の数を制御します。検索履歴は、戻って古い検索を調べたり、グループ内の他のユーザーが検索しているものを確認したりするのに役立ちます。履歴を大きくすると、古い検索文字列の末尾を大きくすることができますが、履歴に保持される検索が多すぎると、GUIを操作するときに速度が低下する可能性があります。

**Prebuff-Block-Hint**
適用対象:インデクサー
デフォルト値:`32`
設定例:`Prebuff-Block-Hint=8`
設定内容:Prebuff-Block-Hintは、データのブロックを格納するときにインデクサーが狙うべきソフトターゲットをメガバイト単位で指定します。非常に高スループットのシステムでは、この値を少し高くしたい場合がありますが、メモリに制約のあるシステムでは、この値を低くしたい場合があります。この値はソフトターゲットであり、インデクサーは通常、摂取が高率で発生している場合にのみこの値を使用します。

**Prebuff-Max-Size**
適用対象:インデクサー
デフォルト値:`32`
設定例:`Prebuff-Max-Size=128`
設定内容:Prebuff-Max-Sizeパラメーターは、エントリをディスクに強制する前にプリバッファーが保持する最大データサイズをメガバイト単位で制御します。プリバッファは、ソースクロックが十分に同期されていない可能性がある場合に、エントリのストレージを最適化するために使用されます。プリバッファが大きいということは、インデクサーが、非常に順序が狂っている値を提供しているインジェスターをより適切に最適化できることを意味します。各ウェルには独自のプリバッファーがあるため、インストールで4つのウェルが定義され、Prebuff-Max-Sizeが256の場合、インデクサーはデータを保持する最大1GBのメモリを消費できます。 プリバッファは定期的にエントリを削除し、それらを常にストレージメディアにプッシュするため、プリバッファの最大サイズは通常、高スループットシステムにのみ関与します。これは、ホストシステムのOOMキラーがGravwellプロセスを終了している場合に、（Search-Pipeline-Buffer-Sizeの後に）回す2番目のノブです。

**Prebuff-Max-Set**
適用対象:ウェブサーバー
デフォルト値:`256`
設定例:`Prebuff-Max-Set=256`
設定内容:Prebuff-Max-Setは、最適化のためにプリバッファーに保持できる1秒のブロックの数を指定します。インジェスターによって提供されたエントリのタイムスタンプが同期していないほど、このセットを大きくする必要があります。たとえば、タイムスタンプが2時間も変動する可能性のあるソースから消費している場合は、この値を7200に設定できますが、データが通常非常に厳しいタイムスタンプ許容値で到着する場合は、この値を低くすることができます。 Prebuff-Max-Sizeコントロールは引き続きプリバッファの削除を有効にして強制するため、この値を高く設定しすぎると、低く設定しすぎるよりも害が少なくなります。

**Prebuff-Tick-Interval**
適用対象:ウェブサーバー
デフォルト値:`3`
設定例:`Prebuff-Tick-Interval=4`
設定内容:Prebuff-Tick-Intervalパラメーターは、プリバッファーがプリバッファーにあるエントリーの人為的な排除を行う頻度を秒単位で指定します。プリバッファは、アクティブな取り込みがある場合、常に値を永続ストレージに削除しますが、非常に低スループットのシステムでは、この値を使用して、エントリが永続ストレージに強制的にプッシュされるようにすることができます。 Gravwellは、データが役立つ場合、データが失われることを決して許しません。インデクサーを正常にシャットダウンすると、プリバッファーはすべてのエントリが永続ストレージに確実に到達するようにします。ただし、ホストの安定性にあまり自信がない場合は、この間隔を2に近づけて、システム障害や怒っている管理者がインデクサーの下から敷物を引き出せないようにすることができます。

**Prebuff-Sort-On-Consume**
適用対象:インデクサー
デフォルト値:`false`
設定例:`Prebuff-Sort-On-Consume=true`
設定内容:Prebuff-Sort-On-Consumeパラメーターは、データをディスクにプッシュする前に、データのロックをソートするようにプリバッファーに指示します。並べ替えプロセスは個々のブロックにのみ適用され、パイプラインに入るときにデータが並べ替えられることを保証するものではありません。ストレージの前にブロックを並べ替えると、取り込み時にパフォーマンスが大幅に低下します。ほとんどすべてのインストールでは、この値をfalseのままにしておく必要があります。

**Max-Block-Size**
適用対象:インデクサー
デフォルト値:`4`
設定例:`Max-Block-Size=8`
設定内容:Max-Block-Sizeはメガバイト単位の値を指定し、エントリをパイプラインにプッシュするときに生成できる最大ブロックサイズをインデクサーに通知するためのヒントとして使用されます。ブロックを大きくすると、パイプラインへのプレッシャーは軽減されますが、メモリプレッシャーは増加します。大容量のメモリと高スループットのシステムではこの値を増やしてスループットを上げることができ、小さいメモリシステムではこのサイズを減らしてメモリの負荷を減らすことができます。 Prebuff-Block-HintパラメーターとMax-Block-Sizeパラメーターが交差して、取り込みと検索のスループットを調整する2つのノブを提供します。Gravwellでは、128GBノードで、次のことが達成されます。クリーンな1GB /秒の検索スループット。Max-Block-Sizeが16の場合、1秒あたり125万エントリが取り込まれます。そして、8のプリバフブロックヒントが達成されます

**Render-Store-Limit**
適用対象:ウェブサーバ－
デフォルト値:1024
設定例:`Render-Store-Limit=512`
設定内容:Render-Store-Limitパラメーターは、検索レンダラーが保存できるメガバイト数を指定します。

**Search-Control-Script**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`Search-Control-Script=/opt/gravwell/etc/authscripts/limits.grv`
設定内容:Search-Control-Scriptパラメーターは、検索時に適用されるスクリプトを指定できるリストパラメーターです。リストパラメータであるため、複数回指定して複数のスクリプトを指定できます。これらのスクリプトは、ユーザーが実行する検索に追加の制限を適用できます。すべてのスクリプトは、検索ごとに実行されます。検索制御スクリプトの詳細については、Gravwellにお問い合わせください。

**Webserver-Resource-Store**
適用対象:ウェブサーバ－
デフォルト値:`/opt/gravwell/resources/webserver`
設定例:`Webserver-Resource-Store=/tmp/path/to/resources/webserver`
設定内容:Webserver-Resource-Storeパラメーターは、ウェブサーバ－がリソースを格納する場所を指定します。このディレクトリは、他のプロセスで使用されていない必要があり、インデクサーまたはデータストアのリソースの場所として指定することはできません。

**Indexer-Resource-Store**
適用対象:インデクサー
デフォルト値:`/opt/gravwell/resources/indexer`
設定例:`Indexer-Resource-Store=/tmp/path/to/resources/indexer`
設定内容:Indexer-Resource-Storeパラメーターは、インデクサーがリソースを格納する場所を指定します。このディレクトリは他のプロセスで使用されていない必要があり、ウェブサーバ－またはデータストアのリソースの場所として指定することはできません。

**Datastore-Resource-Store**
適用対象:データストア
デフォルト値:`/opt/gravwell/resources/datastore`
設定例:`Datastore-Resource-Store=/tmp/path/to/resources/datastore`
設定内容:Indexer-Resource-Storeパラメーターは、インデクサーがリソースを格納する場所を指定します。このディレクトリは他のプロセスで使用されていない必要があり、ウェブサーバ－またはデータストアのリソースの場所として指定することはできません。

**Resource-Max-Size**
適用対象:ウェブサーバ－, Datastore, and Indexer
デフォルト値:`134217728`
設定例:`Resource-Max-Size=1000000000`
設定内容:Resource-Max-Sizeパラメーターは、リソースの最大サイズをバイト単位で指定します。

**Docker-Secrets**
適用対象:ウェブサーバ－, Datastore, and Indexer
デフォルト値:`false`
設定例:`Docker-Secrets=true`
設定内容:Docker-Secretsパラメーターは、[Docker secrets](https://docs.docker.com/engine/swarm/secrets/)からエージェントシークレットの取り込み、制御、検索を試行するようにGravwellに指示します。 シークレットには、それぞれ`ingest_secret`、`control_secret`、および `search_agent_secret`という名前が付けられている必要があり、VM内の`/run/secrets/`ディレクトリからアクセスできる必要があります。

**HTTP-Proxy**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`HTTP-Proxy=wwwproxy.example.com:8080`
設定内容:HTTP-Proxyパラメーターは、ウェブサーバ－によるHTTPおよびHTTP要求に使用されるプロキシーを構成します。これは、環境変数$ http_proxyを設定することと実質的に同等であり、同じ構文を許可します。指定されたプロキシ値は、`HTTP`リクエストと`HTTPS`リクエストの両方に使用されます。

**Webserver-Ingest-Groups**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`Webserver-Ingest-Groups=ingestUsers`
設定内容:Webserver-Ingest-Groupsパラメーターは、ユーザーがGravwell WebAPIを介して直接エントリを取り込むことを許可されているグループを指定するリストパラメーターです。リストパラメータとして、複数回指定して、複数のグループがWebAPIを介して取り込むことができるようにすることができます。

**Disable-Update-Notification**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Disable-Update-Notification=false`
設定内容:Disable-Update-Notificationがtrueに設定されている場合、Gravwellの新しいバージョンが利用可能になったときにWebUIは通知を表示しません。

**Disable-Feature-Popups**
適用対象：ウェブサーバ
デフォルト値：`false`
設定例：`Disable-Feature-Popups=true`
設定内容：Disable-Feature-Popupsがtrueに設定されている場合、Gravwell UIは新機能が追加されたときにポップアップ通知を表示しなくなります。

**Disable-Stats-Report**
適用対象: Webserver
デフォルト値: false
設定例: `Disable-Stats-Report=true`
設定内容:このパラメーターをtrueに設定すると、ウェブサーバ－の[metrics reporting routine](#!metrics.md)に、ライセンスに関する最小限の情報のみを送信し、より広範なシステム統計を省略します。

**Temp-Dir**
適用対象:ウェブサーバ－
デフォルト値:`/opt/gravwell/tmp`
設定例:`Temp-Dir=/tmp/gravtmp`
設定内容:Temp-Dirパラメーターは、他のプロセスからの干渉のリスクなしに一時Gravwellファイルに使用できるディレクトリーを指定します。アップロードされたキットをインストール前に保存するために使用されます。

**Insecure-User-Unsigned-Kits-Allowed**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Insecure-User-Unsigned-Kits-Allowed=true`
設定内容:このパラメーターが設定されている場合、すべてのユーザーが署名されていないキットをインストールできます。このオプションを有効にしないことを強くお勧めします。

**Disable-Search-Agent-Notifications**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Disable-Search-Agent-Notifications=true`
設定内容:trueに設定すると、このパラメーターは、検索エージェントがチェックインに失敗した場合にWeb UIが通知を表示しないようにします。これは、検索エージェントを無効にして通知を表示したくない場合に役立ちます。

**Indexer-Storage-Notification-Threshold**
適用対象:インデクサー
デフォルト値:`90`
設定例:`Indexer-Storage-Notification-Threshold=98`
設定内容:ストレージ使用量について警告するタイミングを決定するパーセンテージ値。値が0より大きい場合、インデクサーによって使用されるストレージデバイスが指定されたストレージパーセンテージを超えて使用するたびに通知がスローされます。値は0から99の間でなければなりません。

**Disable-Network-Script-Functions**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Disable-Network-Script-Functions=true`
設定内容:デフォルトでは、パイプライン内のankoスクリプトは、net/httpライブラリやssh/sftpユーティリティなどのネットワーク関数を使用できます。これを「true」に設定すると、これらの機能が無効になります。

**Webserver-Enable-Frame-Embedding**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Webserver-Enable-Frame-Embedding=true`
設定内容:デフォルトでは、ウェブサーバ－はヘッダーX-Frame-Options：denyを設定することにより、Gravwellページがフレーム内でレンダリングされることを禁止しています。この構成パラメーターを「true」に設定すると、そのヘッダーが削除され、ページをフレーム内に埋め込むことができます。

**Webserver-Content-Security-Policy**
適用対象:ウェブサーバ－
デフォルト値:``
設定例:`Webserver-Content-Security-Policy="default-src https:"`
設定内容:このパラメーターを使用すると、管理者は、すべてのGravwellページで送信されるContent-Security-Policyヘッダーを定義できます。これは重要なセキュリティオプションであり、httpsのみを要求するなど、展開要件に基づいて組織に設定する必要があります。

**Default-Language**
適用対象:ウェブサーバ－
デフォルト値:`en-US`
設定例:`Default-Language=en-US`
設定内容:Default-Languageパラメーターの設定は、認証されていないAPIの /api/language で提供されるものを制御し、複数の言語を使用するデプロイメントでどの言語をデフォルトにするかを決定するためにGUIによって使用されます。これは、ユーザーが言語を選択しておらず、ブラウザが `window.navigator.language`を介して優先言語を提供していない場合のフォールバックです。

**Disable-Map-Tile-Server-Proxy**
適用対象:ウェブサーバ－
デフォルト値:`false`
設定例:`Disable-Map-Tile-Server-Proxy=true`
設定内容:このパラメーターは、Gravwellの組み込みマッププロキシを制御します。マップサーバーに過度の負荷がかかるのを防ぐために、Gravwellウェブサーバ－はマップタイルをキャッシュします。ただし、プロキシを使用するということは、実際のマップサーバーに送信される要求が、ユーザーのWebブラウザーではなくGravwellウェブサーバ－から発信されることを意味します。 Gravwellのインストールがロックダウンされたネットワーク上にある場合、発信HTTPが無効になっているため、これが失敗する可能性があります。 `Disable-Map-Tile-Server-Proxy`をtrueに設定すると、組み込みプロキシが無効になり、GUIが直接マップリクエストを行うようになります。プロキシが無効になっていて、 `Map-Tile-Server`パラメータも設定されている場合、GUIはそのサーバーにリクエストを送信します。

**Map-Tile-Server**
適用対象:ウェブサーバ－
デフォルト値:``
設定例:`Map-Tile-Server=https://maps.example.com/osm/`
設定内容:Map-Tile-Serverパラメーターを使用すると、管理者はマップタイルの別のソースを定義できます。デフォルトでは、GravwellはGravwellマップサーバーからタイルをフェッチし、必要に応じてOpenStreetMapサーバーにフォールバックします。このパラメータを設定すると、Gravwellは指定されたサーバーのみを使用するようになります。指定するURLは、[ここ](https://wiki.openstreetmap.org/wiki/Tile_servers)で定義されている標準のOpenStreetMapタイルサーバー形式のプレフィックスであり、z/x/y座標パラメーターは省略されている必要があります。たとえば、タイルに「https://maps.wikimedia.org/osm-intl/${z}/${x}/${y}.png」でアクセスできる場合、たとえばhttps://maps.wikimedia.org/osm-intl/0/1/2.pngの場合、 `Map-Tile-Server = https//maps.wikimedia.org/osm-intl/`を設定できます。

**Gravwell-Tile-Server-Cooldown-Minutes**
適用対象:ウェブサーバ－
デフォルト値:5
設定例:`Gravwell-Tile-Server-Cooldown-Minutes=1`
設定内容:Gravwellタイルプロキシが通常モードで動作している場合（無効になっていない、`Map-Tile-Server`パラメータが設定されていない）、Gravwellが動作するサーバーからマップタイルをフェッチしようとします。そのサーバーへのリクエストの送信が失敗した場合、プロキシはクールダウンの間、代わりにopenstreetmap.orgサーバーにフォールバックします。このパラメーターを0に設定すると、クールダウンが無効になります。

**Gravwell-Tile-Server-Cache-MB**
適用対象:ウェブサーバ－
デフォルト値:4
設定例:`Gravwell-Tile-Server-Cache-MB=32`
設定内容:Gravwellタイルプロキシは、最近アクセスしたタイルのキャッシュを維持して、マップのレンダリングを高速化します。このパラメーターは、キャッシュが使用できるストレージのメガバイト数を制御します。

**Gravwell-Tile-Server-Cache-Timeout-Days**
適用対象:ウェブサーバ－
デフォルト値:7
設定例:`Gravwell-Tile-Server-Cache-Timeout-Days=2`
設定内容:Gravwellタイルプロキシは、最近アクセスしたタイルのキャッシュを維持して、マップのレンダリングを高速化します。このパラメーターは、キャッシュされたタイルが有効であると見なされる最大日数を制御します。その時間が経過すると、タイルはパージされ、アップストリームサーバーから再フェッチされます。

**Disable-Single-Indexer-Optimization**
適用対象:ウェブサーバ－
デフォルト値:false
設定例:`Disable-Single-Indexer-Optimization=true`
設定内容:Gravwellを単一のインデクサーで使用すると、デフォルトで*インデクサー*ですべてのモジュール（レンダリングモジュールを除く）が実行され、インデクサーからウェブサーバ－に転送されるデータの量が削減されます。このオプションは、その最適化を無効にします。Gravwellのサポートからの指示がない限り、このオプションを「false」に設定したままにしておくことを強くお勧めします。

**Library-Dir**
適用対象:ウェブサーバ－
デフォルト値:`/opt/gravwell/libs`
設定例:`Library-Dir=/scratch/libs`
設定内容:スケジュールされたスクリプトは、`include`関数を使用して追加のライブラリをインポートできます。これらのライブラリは外部リポジトリからフェッチされ、ローカルにキャッシュされます。この構成オプションは、キャッシュされたライブラリーが保管されるディレクトリーを設定します。

**Library-Repository**
適用対象:ウェブサーバ－
デフォルト値:`https://github.com/gravwell/libs`
設定例:`Library-Repository=https://github.com/example/gravwell-libs`
設定内容:スケジュールされたスクリプトは、`include`関数を使用して追加のライブラリをインポートできます。これらのライブラリは、このパラメータで指定されたリポジトリにあるファイルからロードされます。デフォルトでは、便利なライブラリのGravwellが管理するリポジトリを指します。独自のライブラリセットを提供する場合は、このパラメーターを、制御するgitリポジトリーを指すように設定します。

**Library-Commit**
適用対象:ウェブサーバ－
デフォルト値:
設定例:`Library-Commit=19b13a3a8eb877259a06760e1ee35fae2669db73`
設定内容:スケジュールされたスクリプトは、`include`関数を使用して追加のライブラリをインポートできます。これらのライブラリは、 `Library-Repository`オプションで指定されたリポジトリにあるファイルからロードされます。デフォルトでは、Gravwellは最新バージョンを使用します。git commit stringが指定されている場合、Gravwellは代わりに指定されたバージョンのリポジトリを使用しようとします。

**Disable-Library-Repository**
適用対象:ウェブサーバ－
デフォルト値:false
設定例:`Disable-Library-Repository=true`
設定内容:スケジュールされたスクリプトは、`include`関数を使用して追加のライブラリをインポートできます。`Disable-Library-Repository`をtrueに設定すると、この機能が無効になります。

**Gravwell-Kit-Server**
適用対象:ウェブサーバ－
デフォルト値:https://kits.gravwell.io/kits
設定例:`Gravwell-Kit-Server=http://internal.mycompany.io/gravwell/kits`
設定内容:Gravwellキットサーバーホストの設定を変更できます。これは、Gravwellキットサーバーのミラーをホストするエアギャップまたはセグメント化されたデプロイメントで役立ちます。この値を空の文字列に設定すると、リモートキットサーバーへのアクセスが完全に無効になります。
例：
`` `
Gravwell-Kit-Server = "" #gravwellキットサーバーへのリモートアクセスを無効にする
Gravwell-Kit-Server = "http://gravwell.mycompany.com/kits" #内部ミラーを使用することに変更します。
`` `

**Kit-Verification-Key**
適用対象: ウェブサーバ－
デフォルト値:
設定例: `Kit-Verification-Key=/opt/gravwell/etc/kits-pub.pem`
設定内容:キットサーバーからキットを検証するときに使用する公開鍵を含むファイルを指定します。代替のGravwell-Kit-Serverを指定した場合は、この値を設定します。Gravwellの公式キットサーバーを使用する場合は必要ありません。キットの署名に適したキーは、[gencert](https://github.com/gravwell/gencert)ユーティリティを使用して生成できます。

**Disable-User-Ingester-Config-Reporting**
適用対象: ウェブサーバ－
デフォルト値: false
設定例: `Disable-User-Ingester-Config-Reporting=true`
設定内容:インジェスターの状態更新の設定フィールドについて(管理者ではない)一般ユーザーが受信すべきではないことをウェブサーバに通知します。設定にはインジェストシークレットやその他の「機密」項目は含まれていませんが、それでも一般ユーザーには設定全体を秘密にしておきたいという場合もあるでしょう。そのような時にこのオプション設定を使ってください。

**Disable-Ingester-Config-Reporting**
適用対象: ウェブサーバ－
デフォルト値: false
設定例: `Disable-Ingester-Config-Reporting=true`
設定内容:インジェスターの状態更新の設定フィールドについて全ユーザーが受信すべきではないことをウェブサーバに通知します。設定にはインジェストシークレットやその他の「機密」項目は含まれていませんが、それでもユーザーには設定全体を秘密にしておきたいという場合もあるでしょう。そのような時にこのオプション設定を使ってください。

**Disable-Indexer-Overload-Warning**
適用対象: Indexer
デフォルト値：false
設定例：`Disable-Indexer-Overload-Warning=true`（インデクサーのオーバーロードを無効にする）。
設定内容：このパラメータが設定されていると、インデクサーは自身が「オーバーロード」であると判断したときに通知を送信しません。

## パスワード制御

`[Password-Control]`設定セクションを使用して、ユーザーの作成時またはパスワードの変更時にパスワードの複雑さのルールを適用できます。このブロックで設定されたオプションは、ウェブサーバ－にのみ適用されます。これらの複雑なパスワードルールは、シングルサインオンを使用する場合には適用されません。

注：`Password-Control`セクションは複数回宣言しないでください。

```
[Password-Control]
	Min-Length=8
	Require-Uppercase=true
	Require-Lowercase=true
	Require-Special=true
	Require-Special=true
```

**MinLength**
デフォルト値:0
設定例:`MinLength=8`
設定内容:`MinLength`は、パスワードの最小文字長を指定します。

**Require-Uppercase**
デフォルト値:false
設定例:`Require-Uppercase=true`
設定内容:If `Require-Uppercase`が設定されている場合、パスワードには少なくとも1つの大文字が含まれている必要があります。

**Require-Lowercase**
デフォルト値:false
設定例:`Require-Lowercase=true`
設定内容:`Require-Lowercase`が設定されている場合、パスワードには少なくとも1つの小文字が含まれている必要があります。

**Require-Number**
デフォルト値:false
設定例:`Require-Number=true`
設定内容:「Require-Number」が設定されている場合、パスワードには少なくとも1桁の数字が含まれている必要があります。


**Require-Special**
デフォルト値:false
設定例:`Require-Special=true`
設定内容:「Require-Special」が設定されている場合、パスワードには少なくとも1つの「特殊」文字が含まれている必要があります。特殊文字のセットには、数字ではないすべてのUnicode文字が含まれます。

## Well の設定

このセクションのパラメータは、`Default-Well`の設定を含め、ウェルの設定に適用されます。ウェル設定は*インデクサー*にのみ適用されます。つまり、これらのパラメータ設定はウェブサーバ－からは無視されます。以下は、`Default-Well`と"pcap"という名前の2つのウェル構成に対する設定サンプルです。

```
[Default-Well]
	Location=/opt/gravwell/storage/default/
	Cold-Location=/opt/gravwell/cold-storage/default
	Accelerator-Name=fulltext
	Accelerator-Engine-Override=bloom
	Max-Hot-Storage-GB=20

[Storage-Well "pcap"]
	Location=/opt/gravwell/storage/pcap
	Cold-Location=/opt/gravwell/cold-storage/pcap
	Hot-Duration=1D
	Cold-Duration=12W
	Delete-Frozen-Data=true
	Max-Hot-Storage-GB=20
	Disable-Compression=true
	Tags=pcap
```

設定ファイルには、`Default-Well`セクションが1つだけ含まれている必要があります。また、1つ以上の`Storage-Well`セクションが含まれる場合もあります。

ウェルがホットストレージ、コールドストレージ、およびアーカイブストレージ間でエントリを移動する方法の詳細については、[ageoutドキュメント](ageout.md)を参照してください。

注：`Default-Well`に`Tags=`の設定を含めることはできません。デフォルトウェルでは、*他のウェルでは設定されていない全てのタグ*について扱われます。

**Location**
デフォルト値:`/opt/gravwell/storage/default` for `Default-Well`, none for `Storage-Well`
設定例:`Location=/opt/gravwell/storage/foo`
設定内容:このパラメーターは、ウェルが「ホット」データを格納する場所を指定します。2つのウェルに対して同一ディレクトリを指定することはできません！

### エイジアウトオプション

**Hot-Duration**
デフォルト値:
設定例:`Hot-Duration=1w`
設定内容:このパラメーターは、データが「ホット」ストレージの場所に保持される期間を決定します。値は、数値とそれに続く接尾辞（「日」を表す「d」または「週」を表す「w」）で指定する必要があります。したがって、 `Hot-Duration = 30d`は、データを30日間保持する必要があることを示します。`Cold-Location`パラメータが指定されているか、`Delete-Cold-Data`がtrueに設定されていない限り、データは実際にはホットストレージから移動されないことに注意してください。

**Cold-Location**
デフォルト値:
設定例:`Cold-Location=/opt/gravwell/cold_storage/foo`
設定内容:このパラメーターは、コールドデータ（`Location`で指定されたホットストアから移動されたデータ）の保存場所を設定します。

**Cold-Duration**
デフォルト値:
設定例:`Cold-Duration=365d`
設定内容:このパラメーターは、データが「コールド」保管場所に保持される期間を決定します。`Delete-Frozen-Data`がtrueに設定されていない限り、データは実際にはコールドストレージから移動されません。


**Max-Hot-Storage-GB**
デフォルト値:
設定例:`Max-Hot-Storage-GB=100`
設定内容:このパラメーターは、特定のウェルのホットストレージの最大ディスク消費量をギガバイト単位で設定します。この数を超えると、最も古いデータがコールドストレージに移行されるか（可能な場合）、クラウドストレージに送信されるか（設定されている場合）、または削除されます（許可されている場合）。


**Max-Cold-Storage-GB**
デフォルト値:
設定例:`Max-Cold-Storage-GB=100`
設定内容:このパラメーターは、特定のウェルのコールドストレージの最大ディスク消費量をギガバイト単位で設定します。この数を超えると、最も古いデータがクラウドストレージに送信され（設定されている場合）、削除されます（許可されている場合）。

**Hot-Storage-Reserve**
デフォルト値:
設定例:`Hot-Storage-Reserve=10`
設定内容:このパラメーターは、ディスク上に少なくとも特定のパーセンテージを空けておく必要があることをウェルに通知します。したがって、 `Hot-Storage-Reserve = 10`が設定されている場合、ディスクの使用率が90％に達すると、ウェルはホットストレージからデータをエージングアウトしようとします。

**Cold-Storage-Reserve**
デフォルト値:
設定例:`Cold-Storage-Reserve=10`
設定内容:このパラメーターは、ディスク上に少なくとも特定のパーセンテージを空けておく必要があることをウェルに通知します。したがって、 `Cold-Storage-Reserve = 10`が設定されている場合、ディスクの使用率が90％に達すると、ウェルはコールドストレージからデータをエージングアウトしようとします。

**Delete-Cold-Data**
デフォルト値:false
設定例:`Delete-Cold-Data=true`
設定内容:このパラメーターをtrueに設定すると、エージングアウト基準の1つが満たされたときに、ホットストレージの場所からデータを削除できることを意味します。

**Delete-Frozen-Data**
デフォルト値:false
設定例:`Delete-Frozen-Data=true`
設定内容:このパラメーターをtrueに設定すると、エージングアウト基準の1つが満たされたときに、データをコールドストレージから削除できることを意味します。

**Archive-Deleted-Shards**
デフォルト値:false
設定例:`Archive-Deleted-Shards=true`
設定内容:このオプションが設定されている場合、ウェルはシャードを削除する前に外部アーカイブサーバーにアップロードしようとします。これは、`[Cloud-Archive]`セクションが設定されている場合にのみ機能することに注意してください。

**Disable-Compression, Disable-Hot-Compression, Disable-Cold-Compression**
デフォルト値:false
設定例:`Disable-Compression=true`
設定内容:これらのパラメーターは、ウェル内のデータのユーザーモード圧縮を制御します。デフォルトでは、Gravwellはウェル内のデータを圧縮します。 `Disable-Hot-Compression`または` Disable-Cold-Compression`を設定すると、それぞれホットストレージまたはコールドストレージで無効になります。`Disable-Compression`を設定すると、両方で無効になります。

**Enable-Transparent-Compression, Enable-Hot-Transparent-Compression, Enable-Cold-Transparent-Compression**
デフォルト値:false
設定例:`Enable-Transparent-Compression=true`
設定内容:これらのパラメーターは、ウェル内のデータのカーネルレベルの透過的な圧縮を制御します。有効にすると、Gravwellは`btrfs`ファイルシステムにデータを透過的に圧縮するように指示できます。これは、ユーザーモードの圧縮よりも効率的です。`Enable-Transparent-Compression`をtrueに設定すると、ユーザーモードの圧縮が自動的にオフになります。 `Disable-Compression = true`を設定すると、透過圧縮が**無効**になることに注意してください。

**Ageout-Time-Override**
デフォルト値:
設定例:`Ageout-Time-Override="3:00AM"`
設定内容:このパラメーターを使用すると、エイジアウトルーチンを実行する特定の時間を指定できます。通常、この設定は必要ありません。

### 加速オプション

**Accelerator-Name**
デフォルト値:
設定例:`Accelerator-Name=json`
設定内容:`Accelerator-Name`パラメーター（および` Accelerator-Args`パラメーター）を設定すると、ウェルでの加速が有効になります。詳細については、[アクセラレータのドキュメント](#!configuration/accelerators.md)を参照してください。

**Accelerator-Args**
デフォルト値:
設定例:`Accelerator-Args`パラメーター（および`Accelerator-Name`パラメーター）を設定すると、ウェルでの加速が有効になります。詳細については、[アクセラレータのドキュメント](#!configuration/accelerators.md)を参照してください。

**Accelerate-On-Source**
デフォルト値:false
設定例:`Accelerate-On-Source=true`
設定内容:各モジュールのSRCフィールドを含める必要があることを指定します。これにより、CEFなどのモジュールをSRCと組み合わせることができます。

**Accelerator-Engine-Override**
デフォルト値:"index"
設定例:`Accelerator-Engine-Override=bloom`
設定内容:使用する加速エンジンを選択します。デフォルトでは、インデックス作成アクセラレータが使用されます。このパラメーターを"bloom" に設定すると、代わりにブルームフィルターが選択されます。

**Collision-Rate**
デフォルト値:0.001
設定例:`Collision-Rate=0.01`
設定内容:ブルームフィルター加速エンジンの精度を設定します。0.1〜0.000001の値である必要があります。

### 一般オプション

**Disable-Replication**
デフォルト値:false
設定例:`Disable-Replication=true`
設定内容:設定されている場合、このウェルの内容は複製されません。

**Enable-Quarantine-Corrupted-Shards**
デフォルト値:false
設定例:`Enable-Quarantine-Corrupted-Shards=true`
設定内容:設定されている場合、回復できない破損したシャードは、後で分析するために検疫場所にコピーされます。デフォルトでは、ひどく破損したシャードが削除される場合があります。

## レプリケーションの設定

`[Replication]`セクションでは、 [Gravwellのレプリケーション機能](#!configuration/replication.md)を設定します。 設定例は次のようになります:

```
[Replication]
	Disable-Server=true
	Peer=10.0.01
	Storage-Location=/opt/gravwell/replication_storage
```

レプリケーション設定ブロックの内容は、インデクサーにのみ適用されます。

**Peer**
デフォルト値:
設定例:`Peer=10.0.0.1:9406`
設定内容:`Peer`パラメーターはレプリケーションピアを指定します。IPまたはホスト名を取り、最後にオプションのポートを付けます。ポートが指定されていない場合は、デフォルトのポート（9406）が使用されます。`Peer`は複数回指定できます。

**Listen-Address**
デフォルト値:":9406"
設定例:`Listen-Address=192.168.1.1:9406`
設定内容:このパラメーターは、GravwellがどのIPとポート*から*のレプリケーション接続をリッスンするかというを指定します。デフォルトでは、全インターフェイスのポート9406の接続に対してリッスンします。

**Storage-Location**
デフォルト値:
設定例:`Storage-Location=/opt/gravwell/replication`
設定内容:他のGravwellインデクサーから複製されたデータを保存する場所を設定します。

**Max-Replicated-Data-GB**
デフォルト値:
設定例:`Max-Replicated-Data-GB=100`
設定内容:保存する複製データの最大量をギガバイト単位で設定します。これを超えると、インデクサーはレプリケートされたデータのクリーンアップを順次開始します。最初に元のインデクサーで既に削除されているシャードを削除し、次に最も古いシャードの削除を開始します。ストレージサイズが制限を下回ると、削除は停止します。

**Replication-Secret-Override**
デフォルト値:
設定例:`Replication-Secret-Override=MyReplicationSecret`
設定内容:デフォルトでは、Gravwellは `Control-Auth`トークンを使用してレプリケーションの認証を行います。このパラメーターを設定すると、代わりにカスタムレプリケーション認証トークンが定義されます。

**Disable-TLS**
デフォルト値:false
設定例:`Disable-TLS=true`
設定内容:このパラメーターをtrueに設定すると、レプリケーションのTLSが無効になります。インデクサーは暗号化されていない着信接続をリッスンし、暗号化されていない接続を使用してピアと通信します。

**Key-File**
デフォルト値:(`[Global]`セクションの`Key-File`の値)
設定例:`Key-File=/opt/gravwell/etc/replication-key.pem`
設定内容:このパラメーターを使用すると、グローバルに定義されたキーではなく、TLS接続に個別のキーを使用できます。

**Certificate-File**
デフォルト値: (`[Global]`セクションの`Certificate-File`の値)
設定例:`Certificate-File=/opt/gravwell/etc/replication-cert.pem`
設定内容:このパラメーターを使用すると、グローバルに定義された証明書ではなく、TLS接続に個別の証明書を使用できます。

**Insecure-Skip-TLS-Verify**
デフォルト値:false
設定例:`Insecure-Skip-TLS-Verify=false`
設定内容:このパラメーターをtrueに設定すると、レプリケーションピアに接続するときにTLS証明書の検証が無効になります。

**Connect-Wait-Timeout**
デフォルト値:30
設定例:`Connect-Wait-Timeout=60`
設定内容:レプリケーションピアに接続するときに使用されるタイムアウトを秒単位で設定します。

**Disable-Server**
デフォルト値:false
設定例:`Disable-Server=true`
設定内容:レプリケーション*サーバー*機能を無効にします。設定されている場合、インデクサーは自身のデータをレプリケーションピアにプッシュしますが、他のインデクサーがそのデータにプッシュすることはできません。

**Disable-Compression**
デフォルト値:false
設定例:`Disable-Compression=true`
設定内容:レプリケートされたデータの圧縮を制御します。デフォルトでは、レプリケートされたデータはディスクに圧縮されます。

**Enable-Transparent-Compression**
デフォルト値:false
設定例:`Enable-Transparent-Compression=true`
設定内容:このパラメーターがtrueに設定されている場合、Gravwellはレプリケートされたデータに対してbtrfs透過圧縮を使用しようとします。 `Disable-Compression = true`を設定すると、これが無効になります。

## シングルサインオンの設定

`[SSO]`設定セクションは、Gravwellウェブサーバ－のシングルサインオンオプションを指定します。サンプルのセクションは、次のように単純なものです。

```
[SSO]
	Gravwell-Server-URL=https://10.10.254.1:8080
	Provider-Metadata-URL=https://sso.gravwell.io/FederationMetadata/2007-06/FederationMetadata.xml
```

しかしながら往々にして、追加的な設定が必要になります:

```
[SSO]
	Gravwell-Server-URL=https://10.10.254.1:8080
	Provider-Metadata-URL=https://sso.gravwell.io/FederationMetadata/2007-06/FederationMetadata.xml
	Groups-Attribute=http://schemas.xmlsoap.org/claims/Group
	Group-Mapping=Gravwell:gravwell-users
	Group-Mapping=TestGroup:testgroup
	Username-Attribute = "uid"
	Common-Name-Attribute = "cn"
	Given-Name-Attribute  = "givenName"
	Surname-Attribute = "sn"
	Email-Attribute = "mail"
```

詳細については、[SSO構成ドキュメント](sso.md)を参照してください。

**Gravwell-Server-URL**
デフォルト値:
設定例:`Gravwell-Server-URL=https://gravwell.example.org/`
設定内容:SSOサーバーがユーザーを認証した後にユーザーがリダイレクトされるURLを指定します。これは、Gravwellサーバーのユーザー向けのホスト名またはIPアドレスである必要があります。このパラメーターは必須です。

**Provider-Metadata-URL**
デフォルト値:
設定例: `Provider-Metadata-URL=https://sso.example.org/FederationMetadata/2007-06/FederationMetadata.xml`
設定内容:SSOサーバーのXMLメタデータのURLを指定します。上記のパス(`/FederationMetadata/2007-06/FederationMetadata.xml`)はAD FSサーバーでは機能するはずですが、他のSSOプロバイダーでは調整する必要がある場合があります。このパラメーターは必須です。

**Insecure-Skip-TLS-Verify**
デフォルト値:false
設定例:`Insecure-Skip-TLS-Verify=true`
設定内容:trueに設定されている場合、このパラメーターは、SSOサーバーと通信するときに無効なTLS証明書を無視するようにGravwellに指示します。このオプションは注意して設定してください。

**Username-Attribute**
デフォルト値:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"
設定例:`Username-Attribute = "uid"`
設定内容:ユーザー名を含むSAML属性を定義します。Shibbolethサーバーでは、代わりにこれを"uid"に設定する必要があります。

**Common-Name-Attribute**
デフォルト値:"http://schemas.xmlsoap.org/claims/CommonName"
設定例:`Common-Name-Attribute="cn"`
設定内容:ユーザーの「共通名」を含むSAML属性を定義します。 Shibbolethサーバーでは、代わりにこれを「cn」に設定する必要があります。

**Given-Name-Attribute**
デフォルト値:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname"
設定例:`Given-Name-Attribute="givenName"`
設定内容:ユーザーの名を含むSAML属性を定義します。Shibbolethサーバーでは、代わりにこれを"givenName"に設定する必要があります。

**Surname-Attribute**
デフォルト値:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"
設定例:`Surname-Attribute=sn`
設定内容:ユーザーの姓を含むSAML属性を定義します。Shibbolethサーバーでは、代わりにこれを"sn"に設定する必要があります。

**Email-Attribute**
デフォルト値:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
設定例:`Email-Attribute="mail"`
設定内容:ユーザーの電子メールアドレスを含むSAML属性を定義します。 Shibbolethサーバーでは、代わりに「メール」に設定する必要があります。

**Groups-Attribute**
デフォルト値:"http://schemas.microsoft.com/ws/2008/06/identity/claims/groups"
設定例:`Groups-Attribute="groups"`
設定内容:ユーザーが所属するグループのリストを含むSAML属性を定義します。通常、グループリストを送信するようにSSOプロバイダを明示的に設定する必要があります。

**Group-Mapping**
デフォルト値:
設定例:`Group-Mapping=Gravwell:gravwell-users`
設定内容:ユーザーのグループメンバーシップにリストされている場合に自動的に作成される可能性のあるグループの1つを定義します。これは、複数のグループを許可するために複数回指定できます。引数は、コロンで区切られた2つの名前で構成する必要があります。1つ目はグループのSSOサーバー側の名前（通常はAD FSの名前、AzureのUUIDなど）で、2つ目はGravwellが使用する名前です。したがって、`Group-Mapping = Gravwell Users：gravwell-users`を定義する場合、グループ「Gravwell Users」のメンバーであるユーザーのログイントークンを受け取ると、「gravwell-users」という名前のローカルグループが作成されます。 "そしてそれにユーザーを追加します。

## クラウドアーカイブの設定

Gravwellは、個々のウェルの「Archive-Deleted-Shards」パラメーターを使用して、データシャードを削除する前にリモートのクラウドアーカイブサーバーにアーカイブするように設定できます。`[Cloud-Archive]`構成セクションは、これを有効にするためのクラウドアーカイブサーバーに関する情報を指定します。

```
[Cloud-Archive]
	Archive-Server=10.0.0.2:443
	Archive-Shared-Secret=MyArchiveSecret
```

クラウドアーカイブの設定内容は、インデクサーにのみ適用されます。

**Archive-Server**
デフォルト値:
設定例:`Archive-Server=cloudarchive.example.org:443`
設定内容:このパラメーターは、クラウドアーカイブサーバーのIP/ホスト名とオプションでポートを指定します。ポートが指定されていない場合、デフォルト（443）が使用されます。

**Archive-Shared-Secret**
デフォルト値:
設定例:`Archive-Shared-Secret=MyArchiveSecret`
設定内容:クラウドアーカイブサーバーへの認証時に使用する共有シークレットを設定します。インデクサーは、認証プロセスの残りの半分としてライセンスの顧客ID番号を使用します。

**Insecure-Skip-TLS-Verify**
デフォルト値:false
設定例:`Insecure-Skip-TLS-Verify=true`
設定内容:trueに設定すると、インデクサーは接続時にCloudArchiveサーバーのTLS証明書を検証しません。
