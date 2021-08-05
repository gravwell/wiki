# Netflowインジェスター

NetflowインジェスターはNetflowコレクター（Netflowの役割の詳細については[Wikipedia記事](https://en.wikipedia.org/wiki/NetFlow)を参照）として機能し、Netflowエクスポーターによって作成されたレコードを収集し、後の分析のためにGravwellのエントリとして取り込むことができます。これらのエントリは、[netflow](#!search/netflow/netflow.md)検索モジュールを使用して分析することができます。

## 基本設定

Netflowインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、Netflowインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## コレクターの例

```
[Collector "netflow v5"]
	Bind-String="0.0.0.0:2055" #we are binding to all interfaces
	Tag-Name=netflow
	Assume-Local-Timezone=true
	Session-Dump-Enabled=true

[Collector "ipfix"]
	Tag-Name=ipfix
	Bind-String="0.0.0.0:6343"
	Flow-Type=ipfix
```

## インストール

GravwellのDebianリポジトリを使用している場合、インストールはaptコマンド1つで完了します:

```
apt-get install gravwell-netflow-capture
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードしてください。Netflowインジェスターをインストールするには、rootとしてインストーラーを実行するだけです（実際のファイル名には通常、バージョン番号が含まれています）:

```
root@gravserver ~ # bash gravwell_netflow_capture_installer.sh
```

ローカルマシンにGravwellインデクサーがない場合、インストーラはIngest-Secretの値とインデクサー（またはフェデレーター）のIPアドレスの入力を求めます。そうでない場合は、既存のGravwellの設定から適切な値を読み込みます。いずれにしても、インストール後は `/opt/gravwell/etc/netflow_capture.conf` にある設定ファイルを確認してください。UDPポート2055をリッスンする簡単な例は以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO

[Collector "netflow v5"]
	Bind-String="0.0.0.0:2055" #we are binding to all interfaces
	Tag-Name=netflow
```

この設定では、`/opt/gravwell/comms/pipe`を介してローカルのインデクサにエントリを送信することに注意してください。エントリには「netflow」というタグが付けられます。

異なるポートをリッスンする`Collector`エントリを、異なるタグでいくつでも構成することができます。これにより、データをより明確に整理することができます。

注：現時点では、インジェスターはNetflow v5しかサポートしていません。Netflowエクスポーターを構成する際には、この点に注意してください。

