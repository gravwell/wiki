# HTTP

HTTPインジェスターは、1つまたは複数のパスにHTTPリスナーを設定します。これらのパスのいずれかにHTTPリクエストが送信されると、リクエストのBodyが1つのエントリとして取り込まれます。

`curl`コマンドを使えば、標準入力をボディにしたPOSTリクエストを簡単に行うことができるので、スクリプトによるデータ取り込みには非常に便利な方法です。

## 基本設定

HTTPインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、HTTPインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## リスナーの例

すべてのインジェスターで使用される普遍的な構成パラメータに加えて、HTTP POSTインジェスターには、組み込みウェブサーバの動作を制御する2つの追加グローバル設定パラメータがあります。 1つ目の設定パラメータは`Bind`オプションで、ウェブサーバがリッスンするインターフェースとポートを指定します。 2つ目は、`Max-Body`パラメータで、ウェブサーバが許容するPOSTの大きさを制御するものです。Max-Bodyパラメータは、不正なプロセスが非常に大きなファイルを1つのエントリとしてGravwellインスタンスにアップロードしようとするのを防ぐための安全策として有効です。Gravwellは1つのエントリーとして1GBまでサポートしていますが、それを推奨はしません。

複数のリスナーを定義することで、特定のURLから特定のタグにエントリーを送信することができます。 設定例では、天気予報のIoTデバイスとスマートサーモスタットからのデータを受け付ける2つのリスナーを定義しています。

```
 Example using basic authentication
[Listener "basicAuthExample"]
	URL="/basic"
	Tag-Name=basic
	AuthType=basic
	Username=user1
	Password=pass1

[Listener "jwtAuthExample"]
	URL="/jwt"
	Tag-Name=jwt
	AuthType=jwt
	LoginURL="/jwt/login"
	Username=user1
	Password=pass1
	Method=PUT # alternate method, data is still expected in the body of the request

[Listener "cookieAuthExample"]
	URL="/cookie"
	Tag-Name=cookie
	AuthType=cookie
	LoginURL="/cookie/login"
	Username=user1
	Password=pass1
	Method=PUT # alternate method, data is still expected in the body of the request

[Listener "presharedTokenAuthExample"]
	URL="/preshared/token"
	Tag-Name=pretoken
	AuthType="preshared-token"
	TokenName=Gravwell
	TokenValue=Secret

[Listener "presharedTokenAuthExample"]
	URL="/preshared/param"
	Tag-Name=preparam
	AuthType="preshared-parameter"
	TokenName=Gravwell
	TokenValue=Secret
```

## Splunk HEC 互換性

HTTPインジェスターは、Splunk HTTPイベントコレクター(HEC)とAPI互換性のあるリスナーブロックをサポートしています。 この特別なリスナーブロックにより、Splunk HECにデータを送信できるエンドポイントであれば、Gravwell HTTPインジェスターにもデータを送信できるように設定を簡略化することができます。 HEC互換の設定ブロックは以下のようになります。

```
[HEC-Compatible-Listener "testing"]
	URL="/services/collector/event"
	TokenValue="thisisyourtoken"
	Tag-Name=HECStuff

```

`HEC-Compatible-Listener`ブロックは、`TokenValue`と`Tag-Name`の設定を必要とし、`URL`の設定が省略された場合は、`/services/collector/event`がデフォルトとなります。

`Listener`と`HEC-Compatible-Listener`の両設定ブロックは、同じHTTPインジェスターに指定することができます。

## ヘルスチェック

一部のシステム（AWSロードバランサーなど）では、プローブして生存確認できる認証されていないURLを必要とします。HTTPインジェスターは、任意のメソッド、ボディ、クエリパラメータでアクセスした場合に、常に200 OKを返すようなURLを提供するように設定できます。 このヘルスチェックエンドポイントを有効にするには、グローバル設定ブロックに`Health-Check-URL`スタンザを追加します。

ここでは、ヘルスチェックURLを `/logbot/are/you/alive` とした最小限の設定例を紹介します:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO #options are OFF INFO WARN ERROR
Bind=":8080"
Max-Body=4096000 #about 4MB
Log-File="/opt/gravwell/log/http_ingester.log"
Health-Check-URL="/logbot/are/you/alive"

```

## インストール

GravwellのDebianリポジトリを使用している場合、インストールはaptコマンド1つで完了します:

```
apt-get install gravwell-http-ingester
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードしてください。Gravwellサーバのターミナルを使って、スーパーユーザーとして（`sudo`コマンドなどで）以下のコマンドを実行し、インジェスターをインストールします。:

```
root@gravserver ~ # bash gravwell_http_ingester_installer_3.0.0.sh
```

Gravwellのサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメータを抽出し、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラは認証トークンとGravwellインデクサーのIPアドレスを要求します。これらの値はインストール時に設定するか、あるいは空欄にして、`/opt/gravwell/etc/gravwell_http_ingester.conf`の設定ファイルを手動で変更することができます。

## HTTPS設定

デフォルトでは、HTTPインジェスターはクリアテキストのHTTPサーバを実行しますが、x509 TLS証明書を使ってHTTPSサーバを実行するように構成することができます。 HTTPインジェスターをHTTPSサーバとして構成するには、グローバル設定で、`TLS-Certificate-File`と`TLS-Key-File`パラメータを使い、証明書と鍵のPEMファイルを提供します。

HTTPSを有効にしたグローバル設定の例は以下のようになります。:

```
[Global]
	TLS-Certificate-File=/opt/gravwell/etc/cert.pem
	TLS-Key-File=/opt/gravwell/etc/key.pem
```

### リスナーの認証

各HTTPインジェスターリスナーは、認証を実施するように設定できます。 サポートされている認証方法は以下の通りです:

* none
* basic
* jwt
* cookie
* preshared-token
* preshared-parameter

none以外の認証システムを指定する場合は、認証情報を提供する必要があります。`jwt`や`cookie`、Cookie認証システムでは、ユーザー名とパスワードが必要です。一方、`preshared-token`や`preshared-parameter`では、トークンの値とオプションのトークン名を提供する必要があります。

警告: 他のウェブページと同様に、認証は平文での接続では安全ではなく、トラフィックを盗聴できる攻撃者はトークンやクッキーをキャプチャすることができます。

### 認証なし

デフォルトの認証方法はなしで、インジェスターに到達できる人なら誰でもエントリーをプッシュできます。 `basic`認証は、HTTP Basic認証を使用しており、ユーザー名とパスワードをbase64でエンコードして、リクエストごとに送信します。

以下は、basic認証システムを使用したリスナーの例です:

```
[Listener "basicauth"]
	URL="/basic/data"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

basic認証でエントリーを送信するcurlコマンドの例は次のようになります:

```
curl -d "only i can say hi" --user secretuser:secretpassword -X POST http://10.0.0.1:8080/basic/data
```

### JWT認証

JWT認証システムは、暗号化されたトークンを使って認証を行います。 JWT認証を使用する場合は、クライアントが認証を行うログインURLを指定し、リクエストごとに送信する必要のあるトークンを受け取る必要があります。JWTトークンは48時間で失効します。 認証は、`username`と`password`のフォームフィールドを入力した`POST`リクエストをログインURLに送信することで行われます。

HTTPインジェスターでJWT認証を使って認証するには、2つのステップがあり、追加の設定パラメータが必要です。 以下に設定例を示します:

```
[Listener "jwtauth"]
	URL="/jwt/data"
	LoginURL="/jwt/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

エントリを送信するには、エンドポイントが最初に認証を行ってトークンを取得する必要があり、トークンは最大48時間まで再利用できます。リクエストが401レスポンスを受け取った場合、クライアントは再度認証を行う必要があります。ここでは、curlを使って認証を行い、データをプッシュする例を示します。

```
x=$(curl -X POST -d "username=user1&password=pass1" http://127.0.0.1:8080/jwt/login) #grab the token and stuff it into a variable
curl -X POST -H "Authorization: Bearer $x" -d "this is a test using JWT auth" http://127.0.0.1:8080/jwt/data #send the request with the token
```

### Cookie認証

cookie認証システムは、状態を制御する方法以外はJWT認証とほぼ同じです。 クッキー認証を使用するリスナーは、クライアントがユーザー名とパスワードでログインして、ログインページで設定されたクッキーを取得する必要があります。 インジェストURLへの後続のリクエストは、各リクエストでクッキーを提供しなければなりません。

ここでは、設定ブロックの例を示します:

```
[Listener "cookieauth"]
	URL="/cookie/data"
	LoginURL="/cookie/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

ログインして、データを取り込む前にクッキーを取得するcurlコマンドの例は次のようになります:

```
curl -c /tmp/cookie.txt -d "username=user1&password=pass1" localhost:8080/cookie/login
curl -X POST -c /tmp/cookie.txt -b /tmp/cookie.txt -d "this is a cookie data" localhost:8080/cookie/data
```

### プリシェアド・トークン

プリシェアド・トークン認証メカニズムは、ログインメカニズムではなく、事前に共有されたsecretを使用します。事前共有secretは、認証ヘッダーで各リクエストとともに送信されることが期待されます。多くのHTTPフレームワークは、Splunk HECやサポートするAWS KinesisやLambdaインフラストラクチャのように、このタイプのインジェストを期待しています。事前共有トークンのリスナーを使用すると、Splunk HEC のプラグイン代替となるキャプチャシステムを定義することができます。

注：`TokenName`の値を定義しない場合、デフォルト値の`Bearer`が使用されます。

プリシェアド・トークンを定義する設定例です。:

```
[Listener "presharedtoken"]
	URL="/preshared/token/data"
	Tag-name=token
	AuthType="preshared-token"
	TokenName=foo
	TokenValue=barbaz
```

プリシェアド・シークレットを使ってデータを送信するcurlコマンドの例です:

```
curl -X POST -H "Authorization: foo barbaz" -d "this is a preshared token" localhost:8080/preshared/token/data
```

### プリシェアド・パラメータ

プリシェアド・パラメータ認証機構は、クエリパラメータとして提供される事前共有secretを使用します。`Preshared-parameter`システムは、スクリプトを書いたり、認証トークンをURLに埋め込んで通常は認証をサポートしていないデータプロデューサを使用する場合に便利です。

注：認証トークンをURLに埋め込むことは、プロキシやHTTPロギングインフラが認証トークンを捕捉してログに残すことを意味します。

プリシェアド・パラメーターを定義した設定例です:

```
[Listener "presharedtoken"]
	URL="/preshared/parameter/data"
	Tag-name=token
	AuthType="preshared-parameter"
	TokenName=foo
	TokenValue=barbaz
```

プリシェアド・シークレットを使ってデータを送信するcurlコマンドの例です:

```
curl -X POST -d "this is a preshared parameter" localhost:8080/preshared/parameter/data?foo=barbaz
```

## リスナー・メソッド

HTTPインジェスターはほぼすべてのメソッドを使用するように設定できますが、データは常にリクエストのボディにあることが期待されます。

例えば、PUTメソッドを想定したリスナーの設定を以下に示します:

```
[Listener "test"]
	URL="/data"
	Method=PUT
	Tag-Name=stuff
```

対応するcurlコマンドは次のようになります:

```
curl -X PUT -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```

HTTPインジェスターは、特殊文字を含まないほとんどすべてのASCII文字列を受け入れることができ、メソッドの仕様を変更することができます。

```
[Listener "test"]
	URL="/data"
	Method=SUPER_SECRET_METHOD
	Tag-Name=stuff
```

```
curl -X SUPER_SECRET_METHOD -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```
