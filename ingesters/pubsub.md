# GCP PubSubインジェスター

Gravwellでは、Google Compute Platformの[PubSub stream](https://cloud.google.com/pubsub/)サービスからエントリを取得するインゲスターを提供しています。インゲスターは、1つのGCPプロジェクト内で複数のPubSubストリームを処理することができます。PubSubストリームを設定するプロセスはこのドキュメントの範囲外ですが、既存のストリームに対してPubSubインゲスターを設定するためには以下が必要です。

* GoogleプロジェクトID
* GCP サービスアカウントの認証情報を含むファイル（[Creating a service account](https://cloud.google.comauthentication/getting-started)のドキュメントを参照）。
* PubSub トピックの名前

ストリームの設定が完了すると、PubSubのストリームトピックの各レコードは、Gravwellに1つのエントリーとして保存されます。

## 基本構成

PubSubインゲスターは、[インゲスターセクション](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバルコンフィギュレーションブロックを使用しています。 他の多くのGravwellインゲスターと同様に、PubSubは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## PubSubの例

```
[PubSub "gravwell"]
	Topic-Name=mytopic	# the pubsub topic you want to ingest
	Tag-Name=gcp
	Parse-Time=false
	Assume-Local-Timezone=true

[PubSub "my_other_topic"]
	Topic-Name=foo # the pubsub topic you want to ingest
	Tag-Name=gcp
	Assume-Local-Timezone=false
```

## インストールと設定

まず、[Downloads page](#!quickstart/downloads.md)からインストーラーをダウンロードして、インゲスターをインストールします。

```
root@gravserver ~# bash gravwell_pubsub_ingest_installer.sh
```

Gravwellのサービスが同一マシン上に存在する場合は、インストールスクリプトが自動的に`Ingest-Auth`パラメータを抽出し、適切に設定してくれるはずです。次に、`/opt/gravwell/etc/pubsub_ingest.conf`という設定ファイルを開き、PubSubのトピックに設定する必要があります。以下のように設定を変更したら、コマンド `systemctl start gravwell_pubsub_ingest.service` でサービスを開始します。

下の例では、ローカルマシン上のインデクサに接続し（`Pipe-Backend-target`の設定に注意）、GCPプロジェクト "myproject-127400 "の一部である "mytopic "という名前の単一のPubSubトピックからフィードする設定例を示しています。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR

# 使用するGCPプロジェクトID
Project-ID="myproject-127400"
Google-Credentials-Path=/opt/gravwell/etc/google-compute-credentials.json

[PubSub "gravwell"]
	Topic-Name=mytopic	# the pubsub topic you want to ingest
	Tag-Name=gcp
	Parse-Time=false
	Assume-Localtime=true
```

以下の必須フィールドに注意してください：

* `Project-ID` - GCPプロジェクトのプロジェクトID文字列
* `Google-Credentials-Path` - GCPサービスアカウントの認証情報をJSON形式で格納したファイルのパスです。
* `Topic-Name` - 指定したGCPプロジェクト内のPubSubトピックの名前です。

複数の`PubSub`セクションを構成して、1つのGCPプロジェクト内の複数の異なるPubSubトピックをサポートすることができます。

設定をテストするには、`/opt/gravwell/bin/gravwell_pubsub_ingester -v` を手動で実行します。エラーが出力されなければ、設定はおそらく許容範囲内です。

PubSub インジェスターは、他の多くのインジェスターにある `Ignore-Timestamps` オプションを提供しません。PubSubメッセージには到着タイムスタンプが含まれています。デフォルトでは、インゲスターはそれをGravwellタイムスタンプとして使用します。データコンシューマの定義で`Parse-Time=true`が指定されている場合、インジェスターは代わりにメッセージボディからタイムスタンプを抽出しようとします。


