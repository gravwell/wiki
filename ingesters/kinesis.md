# Kinesisインジェスター

Gravwellは、Amazonの[Kinesisデータストリーム](https://aws.amazon.com/kinesis/data-streams/)サービスからエントリを取得できるインジェスターを提供します。インジェスターは一度に複数のKinesisストリームを処理することができ、各ストリームは多数の個別シャードで構成されています。Kinesisストリームを設定するプロセスはこのドキュメントの範囲外ですが、既存のストリームに対してKinesisインジェスターを設定するためには、以下が必要になります:

* AWSアクセスキー (IDナンバーと秘密鍵)
* ストリームが存在する地域
* ストリーム自身の名前

ストリームが設定されると、Kinesisストリームの各レコードはGravwellの1つのエントリーとして保存されます。

## 基本設定

Kinesisインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用しています。 他の多くのGravwellインジェスターと同様に、Kinesisインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## Kinesisストリームの例

```
[KinesisStream "stream1"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as created in AWS
	Iterator-Type=TRIM_HORIZON
	Parse-Time=false
	Assume-Local-Timezone=true

[KinesisStream "stream2"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as created in AWS
	Iterator-Type=TRIM_HORIZON
	Metrics-Interval=60
	JSON-Metric=true
```

## インストールと設定

まず、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードして、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_kinesis_ingest_installer.sh
```

Gravwellのサービスが同一マシン上に存在する場合は、インストールスクリプトが自動的に`Ingest-Auth`パラメータを抽出して適切に設定してくれるはずです。ここで、`/opt/gravwell/etc/kinesis_ingest.conf`という設定ファイルを開き、Kinesisストリーム用に設定を追加する必要があります。以下のように設定を変更したら、コマンド `systemctl start gravwell_kinesis_ingest.service` でサービスを開始します。

以下の例では、ローカルマシン上のインデクサに接続し（`Pipe-Backend-target`の設定に注意）、us-west-1リージョンにある "MyKinesisStreamName" という名前の単一のKinesisストリームからフィードするサンプル構成を示しています。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR
State-Store-Location=/opt/gravwell/etc/kinesis_ingest.state

# This is the access key *ID* to access the AWS account
AWS-Access-Key-ID=REPLACEMEWITHYOURKEYID
# This is the secret key which is only displayed once, when the key is created
#   Note: This option is not required if running in an AWS instance (the AWS
#         the AWS SDK handles that)
AWS-Secret-Access-Key=REPLACEMEWITHYOURKEY

[KinesisStream "stream1"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as AWS knows it
	Iterator-Type=TRIM_HORIZON
	Parse-Time=false
	Assume-Localtime=true
```

`State-Store-Location`オプションに注目してください。これは、既にインジェストされたエントリーの再インジェストを防ぐために、ストリーム内のインジェスターの位置を追跡するステートファイルの場所を設定します。

インジェスターを起動する前に、少なくとも以下のフィールドを設定する必要があります。:

* `AWS-Access-Key-ID` - これは、使用したいAWSアクセスキーのIDです。
* `AWS-Secret-Access-Key` - これは、秘密のアクセスキーそのものです。
* `Region` - kinesisストリームが存在する地域
* `Stream-Name` - kinesisストリームの名前

複数の異なるKinesisストリームをサポートするために、複数の`KinesisStream`セクションを構成することができます。

この設定をテストするには、`/opt/gravwell/bin/gravwell_kinesis_ingester -v`を手動実行します。

ほとんどのフィールドは説明不要ですが、`Iterator-Type`の設定については注意が必要です。この設定は、インジェスターがストリーム/シャードの **状態ファイルエントリを持っていない場合** 、どこからデータを読み始めるかを選択します。デフォルトは "LATEST "で、これはインジェスターが既存のレコードをすべて無視して、インジェスターが開始した後に作成されたレコードのみを読み取ることを意味します。TRIM_HORIZONに設定すると、インジェスターは利用可能な最も古いものからレコードを読み始めます。ほとんどの状況では、古いデータを取得できるようにTRIM_HORIZONに設定することをお勧めします。インジェスターのさらなる実行時に、ステートファイルがシーケンス番号を維持し、重複した取り込みを防止します。

Kinesisインジェスターは、他の多くのインジェスターで見られる`Ignore-Timestamps`オプションを提供しません。Kinesisのメッセージには到着タイムスタンプが含まれており、デフォルトでは、インジェスターはそれをGravwellのタイムスタンプとして使用します。データコンシューマの定義で`Parse-Time=true`が指定されている場合、インジェスターは代わりにメッセージボディからタイムスタンプを抽出しようとします。
