# Gravwellロードバランサー

Gravwellは、お客様の環境をできるだけ簡単に設定するために、Gravwellウェブサーバ用に特別に設計されたカスタムロードバランサーを提供しています。ロードバランサーは、Gravwellのウェブサーバを自動的に検出することができるので、ウェブサーバを追加・削除するたびにロードバランサーを再設定する必要はなく、ウェブサーバがダウンした場合は、自動的に別のサーバにユーザーを誘導することができます。

## ロードバランサーの構成

ロードバランサーはHTTP(S)プロキシであり、クライアントを自動的にGravwellのウェブサーバに誘導します。ロードバランサーは、ユーザーのブラウザにクッキーを設定して"スティッキネス"のレベルを維持し、1つのセッションのリクエストがすべて同じウェブサーバに送られるようにします。

ロードバランサーはGravwellデータストアと通信することでGravwellウェブサーバを発見し、アクティブなウェブサーバのリストを提供します。データストアについての詳細は[分散ウェブサーバ](frontend.md)を参照してください。

インストールと設定が完了したら、ユーザーはロードバランサーを介してGravwellにアクセスします。ロードバランサーを指すホスト名として`gravwell.example.org`などを設定し、ウェブサーバの名前を`web1.example.org`などとすることをお勧めします。ユーザーはウェブサーバに直接アクセスするのではなく、`gravwell.example.org`にアクセスするようにしてください。ロードバランサーを使用する場合、ユーザーはGravwellのウェブサーバーに直接アクセスする必要はありません。ウェブサーバはプライベートアドレスになっていたり、一般にはアクセスできなかったりします。

## ロードバランサーの導入

ロードバランサーコンポーネントは、Gravwellのメインインストーラーと同じルートで配布されます。

* 自己解凍型のシェルインストーラが[ダウンロード](https://docs.gravwell.io/#!quickstart/downloads.md)にあります。
* DebianやRedHatのリポジトリでは、`gravwell-loadbalancer`という名前のパッケージとして提供されています。
* DockerHubでは[gravwell/loadbalancer](https://hub.docker.com/r/gravwell/loadbalancer)です。

Debian のインストーラーは、基本的な設定オプションを求めてきますので、インストール後にそれ以上の設定は必要ありません。他のインストール方法の場合は、以下のように `/opt/gravwell/etc/loadbalancer.conf` を編集する必要があります。また、Dockerを使用している場合は、後述の「Dockerの環境変数」の項で説明するように、環境変数のみでコンテナを設定することもできます。

## 設定ファイル

設定は、`/opt/gravwell/etc/loadbalancer.conf`で管理します。ここでは、HTTPオンリーの簡単な設定を紹介します:

```
[Global]
Disable-HTTP-Redirector=true
Insecure-Disable-HTTPS=true
Update-Interval=10
Session-Timeout=10
Log-Dir=/opt/gravwell/log/loadbalancer
Log-Level=info
Enable-Access-Log=true

Control-Secret=ControlSecrets
Datastore=datastore.example.org
Datastore-Insecure-Disable-TLS=true
```

Disable-HTTP-RedirectorとInsecure-Disable-HTTPSの設定により、ロードバランサーがHTTPのみで着信接続を待ち受けるようになります。ファイルの一番下にあるDatastoreパラメータは、ロードバランサーにGravwellデータストアの場所を伝えます。Control-Secret パラメーターは、暗号化されていない接続でデータストアと通信するようDatastore-Insecure-Disable-TLSが設定されている場合に、データストアと通信するための認証トークンを与えます。

代わりにHTTPSを使いたい場合は、ロードバランサーに有効なTLS証明書とキーのペアを提供する必要があります（GravwellでのTLSの設定については、[TLS](#!configuration/certificates.md)を参照してください）。以下はHTTPSでリッスンし、データストアと暗号化されたチャネルで通信する設定例です:

```
[Global]
Key-File=/opt/gravwell/etc/key.pem
Certificate-File=/opt/gravwell/etc/cert.pem
Update-Interval=10
Session-Timeout=10
Log-Dir=/opt/gravwell/log/loadbalancer
Log-Level=info
Enable-Access-Log=true

#Control-Secret=ControlSecrets
#Datastore=172.19.0.2
#Datastore-Insecure-Disable-TLS=true
#Datastore-Insecure-Skip-TLS-Verify=true
```

### グローバル設定パラメータ

ここでは、設定ファイルの`[Global]`セクションで利用可能なすべてのコンフィギュレーションパラメータをリストアップしています。

**Disable-HTTP-Redirector**
デフォルト値:	`false`
説明:	このパラメータが`false`に設定されていると、ロードバランサーは受信したHTTP接続をHTTPSポートにリダイレクトします。TLSを使用しない場合は、`true`に設定してください。

**Insecure-Disable-HTTPS**
デフォルト値:	`false`
説明:	このパラメータが`true`に設定されていると、ロードバランサーはHTTPSではなくHTTP接続の着信を待ち受けます。これが`false`に設定されると、`Key-File`と`Certificate-File`パラメータも設定しなければなりません。

**Web-Port**
デフォルト値:	443 (`Insecure-Disable-HTTPS=true`が設定されている場合は80)
説明:	このパラメータでは、受信した接続をリッスンするポートを設定します。一般的には、デフォルトの設定で問題ありません。

**Bind-Addr**
デフォルト値:	`0.0.0.0`
説明:	このパラメーターは、ロードバランサーが着信接続をリッスンするIPアドレスを設定します。デフォルトでは、すべてのインターフェイスをリッスンします。

**Key-File**
デフォルト値:	(空欄)
説明:	TLSの秘密鍵の場所を設定します。この鍵は、ロードバランサーのホスト名に対する有効な証明書に対応していなければなりません。

**Certificate-File**
デフォルト値:	(空欄)
説明:	TLS 証明書ファイルの場所を設定します。証明書は、ロードバランサーのホスト名に対して有効でなければなりません。

**Insecure-Skip-TLS-Verify**
デフォルト値:	`false`
説明:	このパラメータが`true`に設定されている場合、ロードバランサーはプロキシ接続する際にGravwell WebサーバのTLS証明書を検証しません。この設定は、`Insecure-Disable-HTTPS=true`が設定されている場合は無視されます。

**Update-Interval**
デフォルト値:	30
説明:	このパラメーターは、ロードバランサーが新規または故障したウェブサーバーをチェックする頻度を秒単位で設定します。

**Session-Timeout**
デフォルト値:	10
説明:	このパラメータは、各ロードバランサーのセッションが継続する時間を分単位で設定します。Gravwellのウェブサーバーはユーザーのログインセッションを同期しているので、ロードバランサーが別のウェブサーバーにリクエストを送るようになっても、ユーザーはセッションが切れることに気づかないことに注意してください。デフォルト値で問題ありません。

**Datastore**
デフォルト値:	(空欄)
説明:	このパラメータはGravwellのデータストアコンポーネントを指します。これはホスト名またはIPアドレスで、例えば`Datastore=datastore.example.org`または`Datastore=192.168.0.11`のようになります。データストアが（9405ではなく）非標準のポートでリッスンしている場合は、そのポートを`Datastore=datastore.example.org:9999`と含めることができます。

**Control-Secret**
デフォルト値:	`ControlSecrets`
説明:	このパラメータは、データストアと通信するための認証トークンを与えます。ほとんどの場合、デフォルトの設定を変更する必要があります。

**Disable-Datastore**
デフォルト値:	`false`
説明:	trueに設定すると、ロードバランサーはどのデータストアとも通信しようとしません。代わりに、`[Override]`設定スタンザ(次のセクションを参照)にリストされたウェブサーバを使用します。

**Datastore-Insecure-Disable-TLS**
デフォルト値:	`false`
説明:	trueに設定すると、ロードバランサーは暗号化されていないチャンネルでデータストアに接続します。

**Datastore-Insecure-Skip-TLS-Verify**
デフォルト値:	`false`
説明:	trueに設定すると、ロードバランサーはデータストアのTLS証明書を検証しません。

**Log-Dir**
デフォルト値:	`/opt/gravwell/log/loadbalancer`
説明:	ロードバランサーがログファイルを保管するディレクトリを設定します。

**Log-Level**
デフォルト値:	`error`
説明:	ログの重大度を設定します。デフォルトでは、エラーのみが記録されます。有効なレベルは、重大度が低い順に、`error`、`warn`、`info`です。これを `off`, `none`, `disabled` に設定すると、ログを完全にオフにします。

**Enable-Access-Log**
デフォルト値:	`false`
説明:	このパラメーターがtrueに設定されていると、ロードバランサーは、クライアントがリクエストしたすべてのURLと、ウェブサーバーからのレスポンスコードを記録します。

### オーバーライドスタンザ

ほとんどのシステムでは、データストアを使用してGravwellウェブサーバーのリストを自動的に取得します。しかし、ロードバランサーがデータストアと直接通信できない場合があります。これは、企業のネットワークセキュリティルールの結果であることがよくあります。このような場合には、設定ファイルの最後に`[Override]`という設定ブロックを追加して、手動でウェブサーバの一覧を取得することができます。

```
[Global]
Disable-HTTP-Redirector=true
Insecure-Disable-HTTPS=true
Disable-Datastore=true

[Overrides "example1"]
	Webserver=172.19.0.100
	Insecure-Disable-HTTPS=true

[Overrides "example2"]
	Webserver=172.19.0.101
	Insecure-Disable-HTTPS=true
```

上の例では、2つのウェブサーバーのオーバーライドスタンザを手動で指定しています。

各オーバーライドには、以下のパラメータを使用できます。

**Webserver**
デフォルト値:	(空欄)
説明:	例えば、`Webserver=192.168.0.1:8080` や `Webserver=web1.example.org` のように、ウェブサーバのホスト名/IP と (オプションの) ポートを指定します。ポートが指定されていない場合は、適切なデフォルト値が設定されます（HTTPSが無効な場合は80、それ以外は443）。

**Insecure-Disable-HTTPS**
デフォルト値:	`false`
説明:	trueに設定すると、ロードバランサーは安全でないHTTPでウェブサーバーと通信します。

**Insecure-Skip-TLS-Verify**
デフォルト値:	`false`
説明:	trueに設定すると、ロードバランサーはウェブサーバーのTLS証明書を検証しません。

## Docker環境変数

Dockerでデプロイする場合、設定ファイルを変更する代わりに、Dockerの環境変数を使ってコンポーネントを設定する方が簡単なことがよくあります。ロードバランサーの基本的なパラメータはすべて環境変数で設定できます:

* `GRAVWELL_INSECURE_DISABLE_HTTPS` - Equivalent to the `Insecure-Disable-HTTPS` config parameter
* `GRAVWELL_INSECURE_SKIP_TLS_VERIFY` - Equivalent to the `Insecure-Skip-TLS-Verify` config parameter
* `GRAVWELL_DISABLE_HTTP_REDIRECTOR` - Equivalent to the `Disable-HTTP-Redirector` config parameter
* `GRAVWELL_WEB_PORT` - Equivalent to the `Web-Port` config parameter
* `GRAVWELL_BIND_ADDR` - Equivalent to the `Bind-Addr` config parameter
* `GRAVWELL_DATASTORE` - Equivalent to the `Datastore` config parameter
* `GRAVWELL_CONTROL_SECRET` - Equivalent to the `Control-Secret` config parameter
* `GRAVWELL_DATASTORE_INSECURE_DISABLE_TLS` - Equivalent to the `Datastore-Insecure-Disable-TLS` config parameter
* `GRAVWELL_DATASTORE_INSECURE_SKIP_TLS_VERIFY` - Equivalent to the `Datastore-Insecure-Skip-TLS-Verify` config parameter
* `GRAVWELL_LOG_LEVEL` - Equivalent to the `Log-Level` config parameter
* `GRAVWELL_LOG_DIR` - Equivalent to the `Log-Dir` config parameter
* `GRAVWELL_LOG_ENABLED_ACCESS_LOG` - Equivalent to the `Enable-Access-Log` config parameter

例えば、以下のように呼び出します:

```
docker create --name loadbalancer \
	-e GRAVWELL_CONTROL_SECRET=ControlSecrets \
	-e GRAVWELL_DATASTORE=datastore.example.org \
	-e GRAVWELL_INSECURE_DISABLE_HTTPS=TRUE \
	-e GRAVWELL_LOG_DIR=/tmp -e GRAVWELL_LOG_LEVEL=INFO \
	-e GRAVWELL_LOG_ENABLE_ACCESS_LOG=TRUE \
	-e GRAVWELL_DATASTORE_INSECURE_DISABLE_TLS=TRUE \
	gravwell/loadbalancer
```
