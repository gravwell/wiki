# Gravwell サーチエージェント

サーチエージェントは、[automations](schedulesearch.md)を実行するコンポーネントです。サーチエージェントはGravwellの主要なインストールパッケージに含まれており、デフォルトでインストールされます。ウェブサーバコンポーネントを `--no-webserver` フラグで無効にしたり、`--no-searchagent` フラグを設定するとサーチエージェントはインストールされなくなります。サーチエージェントはGravwellのDebianパッケージによって自動的にインストールされます。

サーチエージェントが動作しているかどうかは以下のコマンドで確認できます。

```
$ ps aux | grep gravwell_searchagent
```

## サーチエージェントを無効にします

検索エージェントはデフォルトでインストールされていますが、必要に応じて以下を実行して無効にすることができます：

```
systemctl stop gravwell_searchagent.service
systemctl disable gravwell_searchagent.service
```

## サーチエージェントの設定

サーチエージェントの設定は、`/opt/gravwell/etc/searchagent.conf`で行います。設定例を以下に示します：
```
[global]
Webserver-Address=127.0.0.1:80
Insecure-Skip-TLS-Verify=true
Insecure-Use-HTTP=true
Search-Agent-Auth=SearchAgentSecrets
Max-Script-Run-Time=10	# Minutes, set to 0 for unlimited
Log-File=/opt/gravwell/log/searchagent.log
Log-Level=INFO
```

この設定は、ウェブサーバがHTTPSではなくHTTPを使用するように設定されている場合に、ウェブサーバと同じノードで検索エージェントを実行する場合に適しています。ウェブサーバはループバックインターフェース（127.0.0.0.1）上にあり、HTTPが明示的に有効になっていることに注意してください。

サーチエージェント構成ファイルで利用可能な個々の構成オプションを以下に説明します。

**Webserver-Address**

`Webserver-Address`オプションは、検索エージェントがウェブサーバに接続するために使用するIPアドレスまたはホスト名とポートを指定します。このオプションは複数回指定することができます。複数のウェブサーバが定義されている場合(以下に示すように)、検索エージェントはそれらのウェブサーバ間で検索のロードバラン スを行います。

```
Webserver-Address=gravwell1.example.org:443
Webserver-Address=gravwell2.example.org:443
```

注: [datastore](#!distributed/frontend.md)を使用して同期している場合を除き、複数のWebサーバを指定しないでください。

**Search-Agent-Auth**

`Search-Agent-Auth`パラメータはウェブサーバとの認証に使われる認証トークンを設定します。これはインストールプロセス中に自動的に設定されます。ターゲットのウェブサーバの `/opt/gravwell/etc/gravwell.conf`にある `Search-Agent-Auth` の値と一致しなければなりません。

**Insecure-Skip-TLS-Verify**

`Insecure-Skip-TLS-Verify`がtrueに設定されている場合、HTTPS対応のGravwellウェブサーバに接続する際に、検索エージェントはTLS証明書の有効性を*検証しません*。このオプションは注意して使用し、詳細については [certificates documentation](#!configuration/certificates.md) を参照してください。

**Insecure-Use-HTTP**

`Insecure-Use-HTTP`がtrueに設定されている場合、検索エージェントは、デフォルトのHTTPSではなく、平文のHTTPを使用してGravwellウェブサーバとの通信を試みます。このオプションは、[GravwellはHTTPSを有効にするには手動での設定が必要です](#!configuration/certificates.md)ので、デフォルトの設定ファイルではtrueに設定されています。

**Disable-Network-Script-Functions**

デフォルトでは、検索エージェントによって実行されるスケジュールされたスクリプトは、httpライブラリ、 sftp、sshなどのネットワークユーティリティの使用を許可されています。オプション `Disable-Network-Script-Functions=true` を設定すると、これを無効にすることができます。

**HTTP-Proxy**

`HTTP-Proxy`パラメータでは、スケジュールされたスクリプト*が使用するHTTPプロキシを定義する*ことができます。`HTTP-Proxy=https://proxy.example.com:3128`を設定すると、スケジュールされたスクリプトからのHTTPリクエストはすべてこのプロキシを経由してルーティングされます。

**Max-Script-Run-Time**

`Max-Script-Run-Time` パラメータは、スケジュールされたスクリプトが実行できるウォールクロックの最大時間を分単位で設定します。スクリプトが制限時間を超えた場合、スクリプトは直ちに終了します。このパラメータを0に設定するとスクリプトの実行時間は無制限になりますが、ある程度の最大時間を設定することをお勧めします。デフォルトの設定ファイルでは最大時間を10分に設定していますが、これは多くの目的に適しています。

**Log-File**

`Log-File`パラメータは、検索エージェントがログを出力する場所を指定します。

**Log-Level**

`Log-Level`パラメータは、検索エージェントがログに記録すべき深刻度の最小レベルを指定します。オプションはINFO、WARN、ERROR、またはOFFである。WARNを選択すると、深刻度WARNまたはERRORのログが記録されることを意味します。INFO を選択すると、すべてのログが記録されます。OFF を選択すると何も記録されません。

### 環境変数の設定

「searchagent.conf」の設定変数の多くは、実行時の環境変数で提供することができます。 環境変数を使った設定は、dockerのデプロイメントで便利です。

環境変数は、その設定値が設定ファイルに全く設定されていない場合にのみ使用できます。 つまり、`searchagent.conf`ファイルで設定された設定値は、環境変数で提供された値よりも優先されるということです。

以下に、環境変数を使って設定可能なコンフィグレーション変数の一覧を示します
| 環境変数名 | 設定パラメータ | 注意事項 | 
|---------------------------|-------------------------|-------|
| GRAVWELL_SEARCHAGENT_UUID | Searchagent-UUID | |
| GRAVWELL_SEARCHAGENT_AUTH | Search-Agent-Auth | |
| GRAVWELL_WEBSERVER_ADDRESS | ウェブサーバ・アドレス | 複数のアドレスをカンマで区切って指定することができます。|
| GRAVWELL_SEARCHAGENT_DISABLE_NETWORK_SCRIPTS | Disable-Network-Script-Functions | ブール値|
| GRAVWELL_SEARCHAGENT_HTTP_PROXY | HTTP-Proxy |  | 
| GRAVWELL_SEARCHAGENT_INSECURE_SKIP_TLS_VERIFY | Insecure-Skip-TLS-Verify | 真偽値|
| GRAVWELL_SEARCHAGENT_INSECURE_USE_HTTP | Insecure-Use-HTTP | 真偽値 | 