# Windows イベントサービス

Gravwell Windows イベントインジェスターは、Windows マシン上のサービスとして動作し、Windows イベントを Gravwell インデクサーに送信します。 インジェスターはデフォルトで `System`、`Application`、`Setup`、`Security` の各チャンネルを消費します。 各チャンネルは、特定のイベントやプロバイダのセットから消費するように設定できます。

## 基本的な構成

Windows イベントインジェスターは、[インジェスターのセクション](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバルコンフィギュレーションブロックを使用します。 他の多くの Gravwell インジェスターと同様に、Windows イベントインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## EventChannel の例

```
[EventChannel "system"]
	Tag-Name=windows
	Channel=System #pull from the system channel

[EventChannel "sysmon"]
	Tag-Name=sysmon
	Channel="Microsoft-Windows-Sysmon/Operational"
	Max-Reachback=24h  #reachback must be expressed in hours (h), minutes (m), or seconds(s)

[EventChannel "Application"]
	Channel=Application #pull from the application channel
	Tag-Name=winApp #Apply a new tag name
	Provider=Windows System #Only look for the provider "Windows System"
	EventID=1000-4000 #Only look for event IDs 1000 through 4000
	Level=verbose #Only look for verbose entries
	Max-Reachback=72h #start looking for logs up to 72 hours in the past
	Request_Buffer=16 #use a large 16MB buffer for high throughput
	Request_Size=1024 #Request up to 1024 entries per API call for high throughput

[EventChannel "System Critical and Error"]
	Channel=System #pull from the system channel
	Tag-Name=winSysCrit #Apply a new tag name
	Level=critical #look for critical entries
	Level=error #AND for error entries
	Max-Reachback=96h #start looking for logs up to 96 hours in the past

[EventChannel "Security prune"]
	Channel=Security #pull from the security channel
	Tag-Name=winSec #Apply a new tag name
	EventID=-400 #ignore event ID 400
	EventID=-401 #AND ignore event ID 401
```

## インストール

[ダウンロードページ](#!quickstart/downloads.md)から Gravwell Windows インジェスターのインストーラーをダウンロードします。

.msi のインストールウィザードを実行し、Gravwell events serviceをインストールします。 初回のインストール時には、インデクサのエンドポイントとインジェストシークレットの設定を求めるプロンプトが表示されます。 その後のインストールやアップグレードでは、常駐する設定ファイルが特定され、プロンプトは表示されません。

インジェスターの設定は、`%PROGRAMDATA%\gravwell\eventlog\config.cfg` にある `config.cfg` ファイルで行います。 設定ファイルは他の Gravwell インジェスターと同じ形式で、インデクサの接続を設定する`[Global]`セクションと、複数の`EventChannel`定義があります。

インデクサの接続を変更したり、複数のインデクサを指定するには、接続 IP アドレスを Gravwell サーバの IP に変更し、Ingest-Secret の値を設定します。 この例では、暗号化トランスポートを設定しています。

```
Ingest-Secret=YourSecretGoesHere
Encrypted-Backend-target=ip.addr.goes.here:port
```

一度設定したこのファイルは、イベントを収集したい他の Windows システムにコピーすることができます。

### サイレントインストール

Windows イベントインジェスターは、自動化された展開と互換性があるように設計されています。 これは、ドメインコントローラがインストーラをクライアントにプッシュして、ユーザーの操作なしにインストールを起動できることを意味します。 サイレントインストールを強制的に行うには、[msiexec](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/msiexec) で `/quiet` 引数を指定して、管理者権限でインストーラを実行します。 このインストール方法では、デフォルトの構成がインストールされ、サービスが開始されます。

特定のパラメータを設定するには、変更した設定ファイルを `%PROGRAMDATA%gravwell\eventlog\config.cfg` にプッシュしてサービスを再起動するか、`CONFIGFILE` 引数に `config.cfg` ファイルへの完全修飾パスを指定する必要があります。

なお、`%PROGRAMDATA%%gravwell\eventlog` のパスを作成する必要がある場合があります。

グループポリシープッシュの完全な実行順序は次のようになります：

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet
xcopy \\share\gravwell_config.cfg %PROGRAMDATA%\gravwell\eventlog\config.cfg
sc stop "GravwellEvent Service"
sc start "GravwellEvent Service"
```

または、

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet CONFIGFILE=\\share\gravwell_config.cfg
```

## オプションの Sysmon の統合

sysinternals suite の一部である Sysmon ユーティリティは、Windows システムを監視するための効果的で人気のあるツールです。Sysmon の優れた設定ファイルの例を紹介した資料は数多くあります。Gravwell では、infosec Twitter のパーソナリティである @InfosecTaylorSwift 氏が作成した設定ファイルを好んで使用しています。

`%PROGRAMDATA%%gravwell\eventlog\config.cfg` にある Gravwell Windows Agent の設定ファイルを編集し、以下の行を追加します。

```
[EventChannel "Sysmon"]
        Tag-Name=sysmon #Apply a new tag name
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

[SwiftOnSecurity による優れた sysmon 設定ファイルのダウンロード](https://raw.githubusercontent.com/SwiftOnSecurity/sysmon-config/master/sysmonconfig-export.xml)

[sysmon のダウンロード](https://technet.microsoft.com/en-us/sysinternals/sysmon)

管理者シェル（Powershellも可）を使って、以下のコマンドを実行し、設定した内容で `sysmon` をインストールします。

```
sysmon.exe -accepteula -i sysmonconfig-export.xml
```

Windows 標準のサービス管理で Gravwell サービスを再起動します。

### Sysmon の設定例

```
[EventChannel "system"]
        Tag-Name=windows
        #no Provider means accept from all providers
        #no EventID means accept all event ids
        #no Level means pull all levels
        #no Max-Reachback means look for logs starting from now
        Channel=System #pull from the system channel

[EventChannel "application"]
        Tag-Name=windows
        Channel=Application #pull from the system channel

[EventChannel "security"]
        Tag-Name=windows
        Channel=Security #pull from the system channel

[EventChannel "setup"]
        Tag-Name=windows
        Channel=Setup #pull from the system channel

[EventChannel "sysmon"]
        Tag-Name=windows
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

## トラブルシューティング

Web インターフェースの Ingester ページに移動することで、Windows インジェスターの接続性を確認することができます。 Windows のインジェスターが存在しない場合は、Windows の GUI を使用するか、コマンドラインで `sc query GravwellEvents` を実行してサービスの状態を確認してください。

![](querystatus.png)

![](querystatusgui.png)

## Windows 検索の例

デフォルトのタグ名が使用されていると仮定して、すべての sysmon エントリを全体的に見るには、以下の検索を実行します。

```
tag=sysmon
```

すべての Windows イベントを確認するには、以下を実行してください。

```
tag=windows
```

以下の検索では、`winlog` 検索モジュールを使って、特定のイベントやフィールドをフィルタリングして抽出することができます。 非標準的なプロセスによるすべてのネットワーク作成を確認します。

```
tag=sysmon regex winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
table TIMESTAMP SourceHostname Image DestinationIP DestinationPort
```

送信元ホストのネットワーク作成状況をグラフ化します。

```
tag=sysmon regex winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
count by SourceHostname |
chart count by SourceHostname limit 10
```

疑わしいファイルの作成を確認するには次のようにします。

```
tag=sysmon winlog EventID==11 Image TargetFilename |
count by TargetFilename |
chart count by TargetFilename
```
