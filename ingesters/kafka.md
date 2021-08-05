# Kafka

Kafkaインジェスターは、[Apache Kafka](https://kafka.apache.org/)のコンシューマーとして動作するように設計されており、GravwellはKafkaクラスタにアタッチしてデータを取得することができます。 Kafkaは、Gravwellにとって高可用性のある[データブローカー](https://kafka.apache.org/uses#uses_logs)として機能します。 KafkaはGravwell Federatorが提供する役割の一部を担ったり、Gravwellを既存のデータフローに統合する際の負担を軽減することができます。 すでにデータがKafkaに流れている場合は、Gravwellの統合は `apt-get` で簡単にできます。

Gravwell Kafkaインジェスターは、単一のインデクサーのための同一ロケーションのインジェストポイントとして最適です。 KafkaクラスタとGravwellクラスタを運用している場合、GravwellのインジェストレイヤーでKafkaのロードバランシング特性を重複させない方が良いでしょう。 KafkaインジェスターをGravwellインデクサーと同じマシンにインストールし、Unixの名前付きパイプ接続を使用します。各インデクサーにはそれぞれのKafkaインジェスターを設定することで、Kafkaクラスターがロードバランシングを管理することができます。

ほとんどのKafka構成ではデータの耐久性を保証しており、コンシューマーが消費できないときにデータが不揮発性ストレージに保存されます。そのため、GravwellのインジェストキャッシュをKafkaインジェスターで有効にすることは推奨しません。

## 基本設定

Kafkaインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、Kafkaインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## コンシューマーの例

```
[Consumer "default"]
	Leader="127.0.0.1"
	Default-Tag=default   #send bad tag names to default tag
	Tags=*                #allow all tags
	Topic=default
	Tag-Header=TAG        #look for the tag in the Kafka TAG header
	Source-Header=SRC     #look for the source in the Kafka SRC header

[Consumer "test"]
	Leader="127.0.0.1:9092"
	Tag-Name=test
	Topic=test
	Consumer-Group=mygroup
	Synchronous=true
	Key-As-Source=true #A custom feeder is putting its source IP in the message key value
	Header-As-Source="TS" #look for a header key named TS and treat that as a source
	Source-As-Text=true #the source value is going to come in as a text representation
	Batch-Size=256 #get up to 256 messages before consuming and pushing
	Rebalance-Strategy=roundrobin
```

## インストール

Kafkaインジェスターは、GravwellのDebianリポジトリにDebianパッケージとして、またシェルインストーラーとして[ダウンロード](#!quickstart/downloads.md)に用意されています。 リポジトリからのインストールは `apt` を使って行います:

```
apt-get install gravwell-kafka
```

シェルインストーラーは、Arch、Redhat、Gentoo、Fedoraなど、SystemDを使用するDebian以外のシステムをサポートしています。

```
root@gravserver ~ # bash gravwell_kafka_installer.sh
```

## 設定

Gravwell Kafkaインジェスターは、複数のトピック、さらには複数のKafkaクラスターにサブスクライブすることができます。 各コンシューマーは、いくつかの重要な設定値を持つコンシューマーブロックを定義します。


| パラメータ | タイプ | 説明 | 必須 |
|-----------|------|--------------| -------- |
| Tag-Name  | string | データの送信先となるGravwellタグです。 | YES |
| Leader    | host:port | Kafkaクラスターのリーダー/ブローカー。 ポートが指定されていない場合は、デフォルトのポートである9092が付加されます。 | YES |
| Topic     | string | コンシューマが読み取るKafkaトピック | YES |
| Consumer-Group | string | インジェスターが所属しているKafkaコンシューマーグループ | NO (デフォルトは `gravwell`) |
| Source-Override | IPv4 or IPv6 | すべてのエントリーのSRCとして使用するIPアドレス | NO |
| Rebalance-Strategy | string | Kafkaからの読み込み時に使用するリバランシングストラテジー | NO (デフォルトは `roundrobin`。また`sticky`, `range` もオプション) |
| Key-As-Source | boolean | Gravwellのプロデューサーは、データソースのアドレスをメッセージキーに入れることが多く、設定されていれば、インジェスターはメッセージキーをソースアドレスとして解釈しようとします。 キー構造が正しくない場合、インジェスターはオーバーライド（設定されている場合）またはデフォルトのソースを適用します。 | NO (デフォルトはfalse) |
| Synchronous | boolean | インジェスターは、Kafkaのバッチが書き込まれるたびに、インジェスト接続で同期を実行します。 | NO (デフォルトはfalse) |
| Batch-Size | integer | インジェストコネクションへの書き込みを強制する前にKafkaから読み取るエントリーの数 | NO (デフォルトは512) |

警告: コンシューマーを同期型に設定すると、そのコンシューマーはインジェストパイプラインを継続的に同期します。 これは、すべてのコンシューマーのパフォーマンスに大きな影響を与えます。

注: `Synchronous=true` を使用する際に大きな `Batch-Size` を設定すると、高負荷時のパフォーマンスに役立ちます。

### 設定例

ここでは、2つの異なるコンシューマ・グループを使用して2つの異なるトピックを購読する構成例を示します。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-target=/opt/gravwell/comms/pipe
Log-Level=INFO
Log-File=/opt/gravwell/log/kafka.log

[Consumer "default"]
	Leader="tasks.kafka.internal"
	Tag-Name=default
	Topic=default
	Consumer-Group=gravwell1
	Key-As-Source=true
	Batch-Size=256


[Consumer "test"]
	Leader="tasks.testcluster.internal:9092"
	Tag-Name=test
	Topic=test
	Consumer-Group=testgroup
	Source-Override="192.168.1.1"
	Rebalance-Strategy=range
	Batch-Size=4096
```
