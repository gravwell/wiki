# Microsoft Graph API インジェスター

Gravwellは、MicrosoftのGraph APIからセキュリティ情報を取得するインジェスターを提供しています。インジェスターを設定するには、Azure Active Directory管理ポータルで新しい *アプリケーション* を登録する必要があります。これにより、ログにアクセスするためのキーセットが生成されます。その際、以下の情報が必要となります。

* クライアントID: クライアントID：Azure管理コンソールでアプリケーション用に生成されたUUID。
* クライアントシークレット: Azureコンソールでアプリケーション用に生成されたシークレットトークン
* テナントドメイン: Azure ドメインのドメイン(例："mycorp.onmicrosoft.com")

## 基本設定

MS Graphインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、MS Graphインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## コンテンツタイプの例

```
[ContentType "alerts"]
	Content-Type="alerts"
	Tag-Name="graph-alerts"

[ContentType "scores"]
	Content-Type="secureScores"
	Tag-Name="graph-scores"
	Ignore-Timestamps=true

[ContentType "profiles"]
	Content-Type="controlProfiles"
	Tag-Name="graph-profiles"
```

## インストールと設定

まず、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードして、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_msgraph_installer.sh
```

Gravwellのサービスが同一マシン上に存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメータを抽出し、適切に設定します。次に、`/opt/gravwell/etc/msgraph_ingest.conf`設定ファイルを開いて、アプリケーション用に設定し、プレースホルダーフィールドを置き換えたり、必要に応じてタグを修正する必要があります。以下のように設定を変更したら、コマンド `systemctl start gravwell_msgraph_ingest.service` でサービスを開始します。

デフォルトでは、インジェスターはセキュリティアラートが到着するとそれを取り込みます。また、新しいセキュリティスコアの結果を定期的に照会し（通常、毎日発行されます）、それらのセキュリティスコアの結果を構築するために使用される関連する制御プロファイルを取り込みます。これら3つのデータソースは、デフォルトでは、それぞれ`graph-alerts`、`graph-scores`、`graph-profiles`というタグに取り込まれます。

以下の例は、ローカルマシン上のインデクサーに接続し（`Pipe-Backend-target`の設定に注意）、サポートされているすべてのタイプからのログをフィードするサンプル構成を示しています。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR
State-Store-Location=/opt/gravwell/etc/o365_ingest.state

Client-ID=79fb8690-109f-11ea-a253-2b12a0d35073
Client-Secret="<secret>"
Tenant-Domain=mycorp.onmicrosoft.com

[ContentType "alerts"]
	Content-Type="alerts"
	Tag-Name="graph-alerts"

[ContentType "scores"]
	Content-Type="secureScores"
	Tag-Name="graph-scores"

[ContentType "profiles"]
	Content-Type="controlProfiles"
	Tag-Name="graph-profiles"
```
