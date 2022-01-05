# インジェスター

このセクションには、Gravwellインジェスターを設定して実行するためのより詳細な説明が記載されています。

Gravwellが作成したインジェスターは、BSDオープンソースライセンスで公開されており、[Github](https://github.com/gravwell/gravwell/tree/master/ingesters)で見ることができます。インジェストAPIもオープンソースなので、独自のデータソース、追加の正規化や前処理など、様々な方法で独自のインジェスターを作成することができます。インジェストAPIのコードは[ここ](https://github.com/gravwell/gravwell/tree/master/ingest)にあります。

一般的に、インジェスターがGravwellにデータを送信するためには、インジェスターの認証のためにGravwellインスタンスの "Ingest Secret"を知る必要があります。これは、Gravwellサーバー上の`/opt/gravwell/etc/gravwell.conf`ファイル内の、`Ingest-Auth`エントリに設定されています。インジェスターがGravwell自身と同じシステム上で動作している場合、インストーラーは通常、この値を検出して自動的に設定してくれます。

Gravwell GUIには、(システムメニューカテゴリの下に) インジェスターページがあり、どのリモートインジェスターがアクティブに接続されているか及び、それらが接続されている時間と、プッシュされたデータ量を簡単に識別するために使用できます。

![](remote-ingesters.png)

注意： [レプリケーションシステム](#!configuration/replication.md)は、999MBを超えるエントリをレプリケーションしません。より大きなエントリは通常通りインジェストして検索できますが、レプリケーションからは省略されます。このページで詳しく説明しているインジェスターはすべて数キロバイト以下のエントリを作成する傾向があるので、これは99.9%のユースケースでは問題になりません。

## インジェスター

| インゲスター | 説明 |
|----------|-------------|
| [Amazon SQS](#!ingesters/sqs.md) | Amazon SQSのキューからサブスクライブしてインジェストします。|
| [collectd](#!ingesters/collectd.md) | collectdのサンプルを取り込みます。|
| [Disk Monitor](#!ingesters/disk.md) | ディスクのアクティビティを定期的にサンプリングします。|
| [File Follower](#!ingesters/file_follow.md) | ログなど、ディスク上のファイルを監視して取り込みます。|
| [GCP PubSub](#!ingesters/pubsub.md) | Google Compute Platform PubSub Streamsからエントリをフェッチしてインジェストします。|
| [HTTP](#!ingesters/http.md) | 複数のURLパスにHTTPリスナーを作成します。|
| [IPMI](#!ingesters/ipmi.md) | IPMIデバイスからSDRおよびSELレコードを定期的に収集します。|
| [Kafka](#!ingesters/kafka.md) | GravwellにインジェストするKafkaコンシューマーを作成します。Gravwell Kafka Federatorとの連携も可能です。|
| [Kinesis](#!ingesters/kinesis.md) | Amazonの[Kinesis Data Streams](https://aws.amazon.com/kinesis/data-streams/)サービスからインジェストする。|
| [Mass File](#!ingesters/massfile.md) | 大量の静的ファイルをインジェストすることができます。|
| [Microsoft Graph API](#!ingesters/msg.md) | MicrosoftのGraph APIからインジェストします。|
| [Netflow](#!ingesters/netflow.md) | NetflowおよびIPFIXレコードを収集します。|
| [Network Capture](#!ingesters/pcap.md) | ワイヤー上のPCAPをインジェストします。|
| [Office 365](#!ingesters/o365.md) | Microsoft o365 Logsを取り込みます。|
| [Packetfleet](#!ingesters/packetfleet.md) | Google Stenographerからクエリを発行してデータを取り込みます。|
| [Session](#!ingesters/session.md) | 大規模なレコードを1つのエントリに取り込みます。|
| [Simple Relay](#!ingesters/simple_relay.md) | TCP/UDPやsyslogなど、あらゆるテキストをインジェストします。|
| [Windowsイベント](#!ingesters/winevent.md) | Windowsイベントを収集します。|

## タグ

タグはGravwellの重要なコンセプトです。すべてのエントリーには1つのタグが関連付けられており、これらのタグによってデータを基本的なレベルで分離・分類することができる。例えば、Linuxシステムのログファイルから読み取ったエントリーには「syslog」タグを適用し、Windowsのログには「winlog」を適用し、生のネットワークパケットには「pcap」を適用するといった具合だ。どのタグをエントリに適用するかは、インジェスターが決定します。

ユーザーから見ると、タグは「syslog」、「pcap-router」、「default」などの文字列です。タグ名には以下の文字は使用できません。

```
!@#$%^&*()=+<>,.:;"'{}[]|\
```

また、タグ名を選択する際に、非印刷文字やタイプしにくい文字を使用するのは控えるべきです。☺という名前のタグにインジェストすることは*可能*ですが、それが良いアイデアというわけではありません。

### タグのワイルドカード

タグ名を選択する際には、Gravwellではタグ名を指定してクエリを実行する際にワイルドカードを使用できることに注意してください。タグ名を慎重に選択することで、後のクエリを容易にすることができます。

例えば、5台のサーバからシステムログを収集しており、そのうち2台がHTTPサーバ、2台がファイルサーバ、1台がメールサーバである場合、以下のようなタグを使用することができます。

* syslog-http-server1
* syslog-http-server2
* syslog-file-server1
* syslog-file-server2
* syslog-email-server1

これにより、[クエリ](#!search/search.md)でログをより柔軟に選択できるようになります。`tag=syslog-*`を指定すれば、すべてのシステムログを検索することができます。`tag=syslog-http-*`を指定すれば、すべてのHTTPサーバーのログを検索できますし、`tag=syslog-http-server1`と言えば、1つのサーバーを選択できます。また、`tag=syslog-http-*,syslog-email-*`のように複数のワイルドカードグループを選択することもできます。

### タグの内部事情

Gravwellを使う上でこのセクションを読むことは必須ではありませんが、タグが内部でどのように管理されているかを理解するのに役立つかもしれません。

内部的には、Gravwellの*インデクサ*はタグを16ビットの整数として格納している。各インデクサはタグ名とタグ番号のマッピングを管理しており、その情報は `/opt/gravwell/etc/tags.dat` にある。Gravwellのサポートから明確な指示がない限り、絶対にこのファイルを変更したり削除したりしないでください。

*インゲスター*がインデクサーに接続する際には、使用したいタグ名のリストを送信します。インデクサーは、タグ名とタグ番号のマッピングを返信します。インゲスターがそのインデクサーにエントリーを送るときはいつでも、適切な*タグ番号*をエントリーに追加します。

## グローバル設定パラメータ

ほとんどのコア・インゲスターは、共通のグローバル・コンフィギュレーション・パラメーターのセットをサポートしています。 共有のグローバル設定パラメータは、[ingest config](https://godoc.org/github.com/gravwell/ingest/config#IngestConfig)パッケージを使用して実装されます。 Global設定パラメータは、各Gravwellインゲスター設定ファイルのGlobalセクションで指定する必要があります。 以下のGlobal ingesterパラメータがあります。

* Ingest-Secret
* Connection-Timeout
* Rate-Limit
* Enable-Compression
* Insecure-Skip-TLS-Verify
* Cleartext-Backend-Target
* Encrypted-Backend-Target
* Pipe-Backend-Target
* Ingest-Cache-Path
* Max-Ingest-Cache
* Cache-Depth
* Cache-Mode
* Log-Level
* Log-File
* Source-Override
* Log-Source-Override

### Ingest-Secret

Ingest-Secretパラメータは、インジェスト認証に使用するトークンを指定します。 ここで指定したトークンは、GravwellインデクサーのIngest-Authパラメータと一致しなければなりません（MUST）。

### 接続タイムアウト

Connection-Timeoutパラメータは、インデクサへの接続をあきらめるまでの待ち時間を指定します。 タイムアウトが空だと、インゲスターが起動するまで永遠に待つことになります。 タイムアウトは、分、秒、時間のいずれかで指定します。

#### 例
```
Connection-Timeout=30s
Connection-Timeout=5m
Connection-Timeout=1h
```

### Insecure-Skip-TLS-Verify

Insecure-Skip-TLS-Verifyトークンは、暗号化されたTLSトンネルで接続する際に、不正な証明書を無視するようインゲスターに指示します。その名が示すように、TLSが提供するすべての認証は窓から投げ出され、攻撃者は簡単にTLS接続の中間者となることができます。 インジェスト接続は暗号化されますが、接続は決して安全ではありません。 デフォルトでは、TLS証明書は検証され、証明書の検証に失敗すると、接続は失敗します。

#### 例
```
Insecure-Skip-TLS-Verify=true
Insecure-Skip-TLS-Verify=false
```

### Rate-Limit

Rate-Limitパラメータは、インゲスターが消費できる最大の帯域幅を設定します。これは、低速な接続でインデクサと通信する「バースト型」インゲスタを設定する際に便利で、インゲスタが大量のデータを送信しようとする際に利用可能な帯域幅をすべて占有しないようにします。

引数は、数値の後にオプションのレートサフィックスを付ける必要があります。例えば、`1048576`や`10Mbit`などです。次のような接尾辞があります。

* **kbit, kbps, Kbit, Kbps**: "kilobits per second"
* **KBps**: "kilobytes per second"
* **mbit, mbps, Mbit, Mbps**: "megabits per second"
* **MBps**: "megabytes per second"
* **gbit, gbps, Gbit, Gbps**: "gigabits per second"
* **GBps**: "gigabytes per second"

#### 例

```
Rate-Limit=1Mbit
Rate-Limit=2048Kbps
Rate-Limit=3MBps
```

### Enable-Compression

インジェストシステムは、インジェスターとインデクサーの間を流れるデータを圧縮する、透過的な圧縮システムをサポートしています。この透過的な圧縮は非常に高速で、低速なリンクの負荷を軽減することができます。各インジェスターは、グローバル設定ブロックの`Enable-Compression`パラメータを`true`に設定することで、すべての接続に対して圧縮されたアップリンクを要求することができます。

圧縮システムは日和見主義で、インジェスターが圧縮を要求しても、圧縮を有効にするかどうかの最終決定権は上流のリンクにあります。上流のエンドポイントが圧縮をサポートしていないか、あるいは圧縮を許可しないように設定されている場合、リンクは圧縮されません。

圧縮は、インジェスターのCPUおよびメモリ要求を増加させます。インジェスターが最小限のCPUおよびメモリでエンドポイント上で動作している場合、圧縮はスループットを低下させる可能性があります。圧縮は、WAN接続に最も適しており、Unixの名前付きパイプで圧縮を有効にしても、CPUとメモリのオーバーヘッドが発生するだけで、何の付加価値もありません。


#### 例

```
Enable-Compression=true
```

### Cleartext-Backend-Target

Cleartext-Backend-Targetは、Gravwellインデクサーのホストとポートを指定します。 インゲスターは、クリアテキストのTCP接続を使用してインデクサに接続します。 ポートが指定されていない場合は、デフォルトの4023ポートが使用されます。 クリアテキスト接続は、IPv6とIPv4の両方の宛先をサポートします。 **複数のCleartext-Backend-Targetsを指定して、1つのインゲスターを複数のインデクサーにまたがってロードバランスすることができます。**

#### 例
```
Cleartext-Backend-Target=192.168.1.1
Cleartext-Backend-Target=192.168.1.1:4023
Cleartext-Backend-Target=DEAD::BEEF
Cleartext-Backend-Target=[DEAD::BEEF]:4023
```

### Encrypted-Backend-Target

Encrypted-Backend-Targetは、Gravwell Indexerのホストとポートを指定します。インゲスターはTCP経由でインデクサーに接続し、TLSのフルハンドシェイク/証明書の検証を行います。 ポートが指定されていない場合は、デフォルトの4024ポートが使用されます。 暗号化された接続は、IPv6とIPv4の両方の宛先をサポートしています。 **1つのインゲスターを複数のインデクサーでロードバランスするために、Multiple Encrypted-Backend-Targetsを指定することができます。**

#### 例
```
Encrypted-Backend-Target=192.168.1.1
Encrypted-Backend-Target=192.168.1.1:4023
Encrypted-Backend-Target=DEAD::BEEF
Encrypted-Backend-Target=[DEAD::BEEF]:4023
```

### Pipe-Backend-Target

Pip-Backend-Targetは、フルパスを介してUnixの名前付きソケットを指定します。 Unixの名前付きソケットは、非常に高速でオーバーヘッドがほとんどないため、インデクサーと共存するインジェスターに最適です。 ingesterごとにサポートされるPipe-Backend-Targetは1つだけですが、パイプはクリアテキストおよび暗号化された接続と一緒に多重化できます。

#### 例
```
Pipe-Backend-Target=/opt/gravwell/comms/pipe
Pipe-Backend-Target=/tmp/gravwellpipe
```

### Ingest-Cache-Path

Ingest-Cache-Pathは、インジェスト・データのためのローカル・キャッシュを有効にします。 これを有効にすると、インデクサーにエントリを転送できないときに、インジェストはローカルにキャッシュすることができます。 インジェストキャッシュは、リンクがダウンしたときや、Gravwellクラスタを一時的にオフラインにする必要があるときに、データを失わないようにするのに役立ちます。 Max-Ingest-Cacheの値を必ず指定して、長期間のネットワーク障害でインゲスターがホストディスクを満杯にすることがないようにしてください。 ローカルのインジェストキャッシュは、インデクサーに直接インジェストするほど高速ではありませんので、インデクサーのようにインジェストキャッシュが毎秒200万エントリを処理できるとは思わないでください。

注意：File Followerインゲスターでは、インジェストキャッシュを有効にしてはいけません。このインゲスターは、ディスク上のファイルから直接読み取り、各ファイル内の位置を追跡するため、キャッシュを必要としません。

#### 例
```
Ingest-Cache-Path=/opt/gravwell/cache/simplerelay.cache
Ingest-Cache-Path=/mnt/storage/networklog.cache
```

### Max-Ingest-Cache

Max-Ingest-Cacheは、キャッシュが作動したときにインゲスターが消費するストレージ容量を制限します。 キャッシュの最大値はメガバイト単位で指定され、1024という値は、インゲスターが1GBのストレージを消費してから新しいエントリの受け入れを停止することを意味します。 キャッシュシステムは、キャッシュがいっぱいになっても古いエントリを上書きしません。これは、攻撃者がネットワーク接続を中断して、中断した時点で潜在的に重要なデータをインゲスターに上書きさせることがないように設計されています。

#### 例
```
Max-Ingest-Cache=32
Max-Ingest-Cache=1024
Max-Ingest-Cache=10240
```

### Cache-Depth

Cache-Depthは、インメモリ・バッファに保持するエントリの数を設定します。既定値は128で、Ingest-Cache-Pathが無効になっていても、インメモリ・バッファは常に有効になっています。Cache-Depthを大きな値に設定すると、メモリ消費量を犠牲にしても、インジェスターのバースト動作を吸収することができます。

#### 例
```
Cache-Depth=256
```

### Cache-Mode

Cache-Modeは、実行時のバッキングキャッシュ（Ingest-Cache-Pathの設定で有効）の動作を設定します。利用可能なモードは、"always "と "fail "です。always "モードでは、キャッシュは常に有効で、（Cache-Depthで設定された）インメモリ・バッファがいっぱいになったときに、インジェスターがディスクにエントリを書き込むことができます。これは、インデクサの接続が切れているか遅い場合、またはインゲスターがインデクサとの接続で可能な範囲を超えてデータをプッシュしようとしている場合に起こります。always "モードを使用することで、インゲスターがいつでもエントリーをドロップしたり、データの取り込みをブロックしたりしないことが保証されます。Cache-Modeを "fail "に設定すると、すべてのインデクサの接続がダウンしたときにのみキャッシュの動作が有効になります。

#### 例
```
Cache-Mode=always
Cache-Mode=fail
```

### Log-File

インジェスターは、エラーやデバッグ情報をログファイルに記録し、インストールや設定の問題をデバッグするのに役立ちます。Log-Fileパラメータを空にすると、ファイルのロギングが無効になります。

#### 例
```
Log-File=/opt/gravwell/log/ingester.log
```

### Log-Level

Log-Levelパラメータは、"gravwell "タグでインデクサに送信されるログ・ファイルとメタデータの両方について、各インゲスタのロギング・システムを制御します。 ログレベルをINFOに設定すると、File Followerが新しいファイルを追跡したときや、Simple Relayが新しいTCP接続を受信したときなど、詳細なログを記録するようにインゲスターに指示します。一方、レベルをERRORに設定すると、最も重要なエラーのみがログに記録されます。ほとんどの場合、WARNレベルが適切です。以下のレベルがサポートされています。

* OFF
* INFO
* WARN
* ERROR

#### 例
```
Log-Level=Off
Log-Level=INFO
Log-Level=WARN
Log-Level=ERROR
```

### Source-Override

Source-Overrideパラメータは、各エントリに添付されているSRCデータ項目を上書きします。 SRCアイテムは、IPv6またはIPv4アドレスで、通常はインジェスターが動作しているマシンの外部IPアドレスです。

#### 例
```
Source-Override=10.0.0.1
Source-Override=0.0.0.0
Source-Override=DEAD:BEEF::FEED:FEBE
```

### Log-Source-Override

多くのインゲスターは、監査、健全性とステータス、および一般的なインジェストインフラストラクチャのロギングを目的として、`gravwell`タグにエントリを発行することができます。 一般的に、これらのエントリは、SRCフィールドにインデクサから見たインゲスタのソースIPアドレスを使用します。 しかし、インゲスターによって実際に生成されたエントリーだけにソースIPフィールドをオーバーライドすることは有用です。 例えば、Gravwell Federator で `Log-Source-Override` を使用して、ヘルスやステータスのエントリの SRC フィールドを変更するが、Federator を通過するすべてのエントリは変更しないといった例があります。

`Log-Source-Override`構成パラメータには、IPv4またはIPv6の値をパラメータとして指定する必要があります。

#### 例
```
Log-Source-Override=10.0.0.1
Log-Source-Override=0.0.0.0
Log-Source-Override=DEAD:BEEF::FEED:FEBE
Log-Source-Override=::1
```

## データコンシューマーの設定

グローバル設定オプションの他に、設定ファイルを使用する各インゲスターは、少なくとも1つの*データ・コンシューマー*を定義する必要があります。データ・コンシューマーは、インゲスターに以下を伝えるコンフィグ定義です。

* どこでデータを取得するか
* データにどのようなタグを使用するか
* 特別なタイムスタンプ処理ルール
* SRCフィールドなどのフィールドのオーバーライド

Simple RelayインゲスターとHTTPインゲスターは「Listener」を定義し、File Followは「Followers」を使用し、netflowインゲスターは「Collectors」を定義しています。以下の各インゲスターのセクションでは、インゲスターの特定のデータ・コンシューマー・タイプと、それらが必要とする独自の設定について説明します。次の例は、File Followerインゲスターが、特定のディレクトリからデータを読み取るために「Follower」データコンシューマーを定義する方法を示しています。

```
[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
```

データソースの指定（`Base-Directory`と`File-Filter`のルール）、使用するタグの指定（`Tag-Name`のルール）、受信データのタイムスタンプを解析するための追加ルール（`Assume-Local-Timezone`のルール）に注目してください。

## タイムパーシングのオーバーライド

ほとんどのインジェスターは、データからタイムスタンプを抽出することで、各エントリにタイムスタンプを適用しようとします。このタイムスタンプの抽出を微調整するために、各 *data consumer* に適用できるオプションがいくつかあります。

* `Ignore-Timestamps` (boolean): `Ignore-Timestamps=true` を設定すると、インジェスターはタイムスタンプを抽出しようとせずに、各エントリに現在の時間を適用します。これは、非常に支離滅裂な入力データがある場合に、データをインジェストするための唯一のオプションになります。
* `Assume-Local-Timezone` (boolean)：デフォルトでは、タイムスタンプにタイムゾーンが含まれていない場合、インジェスターはそれがUTCタイムスタンプであると仮定します。Assume-Local-Timezone=true`を設定すると、インゲスターは代わりにローカルコンピュータのタイムゾーンを想定します。これはTimezone-Overrideオプションとは相互に排他的です。
* Timezone-Override` (string): Timezone-Override`を設定すると、タイムゾーンを含まないタイムスタンプは指定されたタイムゾーンで解析されるべきだとインゲスターに伝えます。したがって、`Timezone-Override=US/Pacific`とすると、インゲスターは受信したタイムスタンプをアメリカの太平洋時間であるかのように扱うようになります。受け入れ可能なタイムゾーン名（「TZデータベース名」欄）の完全なリストについては、[このページ](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)を参照してください。Assume-Local-Timezoneとは相互に排他的です。
* `Timestamp-Format-Override` (string): 例えば、`Timestamp-Format-Override="RFC822"`のようになります。使用可能なオーバーライドの全リストと例については、[タイムグラインダーのドキュメント](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder)を参照してください。

KinesisとGoogle Pub/Subのインジェスターは`Ignore-Timestamps`オプションを提供していません。KinesisとPub/Subは各エントリに到着タイムスタンプを含んでいます。デフォルトでは、インゲスターはそれをGravwellのタイムスタンプとして使用します。データコンシューマの定義で`Parse-Time=true`が指定されている場合、インゲスターは代わりにメッセージボディからタイムスタンプを抽出しようとします。追加情報については、これらのインゲスターの各セクションを参照してください。

カスタムタイムスタンプフォーマットは多くのインゲスターでサポートされています。詳細は[Custom Time Formats](#!ingesters/customtime/customtime.md)を参照してください。

## ソースオーバーライド

Source-Override」パラメータは、コンシューマがデータのソースを無視してハードコードされた値を適用することを指示します。 データソースを整理したりグループ化したりするために、入力データのソース値をハードコードすることが望ましい場合があります。 「Source-Override」の値は、IPv4またはIPv6の値です。

```
Source-Override=192.168.1.1
Source-Override=127.0.0.1
Source-Override=[fe80::899:b3ff:feb7:2dc6]
```

## Gravwell フェデレータ

インジェスタはフェデレータに接続してエントリを送信し、フェデレータはそれらのエントリをインデクサーに渡します。 フェデレータは信頼境界として機能し、インジェストの秘密を公開したり、信頼されていないノードが許可されていないタグのデータを送信したりすることなく、ネットワークセグメント間でエントリを安全に中継します。 フェデレータのアップストリーム接続は、他のインジェスターと同様に構成され、多重化、ローカルキャッシング、暗号化などを可能にします。

![](federatorDiagram.png)

### IngestListener の例

```
[IngestListener "enclaveA"]
	Ingest-Secret = CustomSecrets
	TLS-Bind = 0.0.0.0:4024
	TLS-Certfile = /opt/gravwell/etc/cert.pem
	TLS-Keyfile = /opt/gravwell/etc/key.pem
	Tags=windows
	Tags=syslog-*

[IngestListener "enclaveB"]
	Ingest-Secret = OtherIngestSecrets
	Cleartext-Bind = 0.0.0.0:4023
	Tags=apache
	Tags=bash
```


### ユースケース

 * 堅牢な接続性がない場合、地理的に多様な地域にまたがるデータのインジェスト
 * ネットワークセグメント間の認証バリアの提供
 * インデクサへの接続数の削減
 * データソースグループが提供できるタグの制御

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-federator
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラをダウンロードします。Gravwell サーバー上のターミナルを使って、スーパーユーザーとして(例 : `sudo` コマンドを使って)以下のコマンドを実行してフェデレータをインストールしてください:

```
root@gravserver ~ # bash gravwell_federator_installer.sh
```

フェデレータは、ほぼ確実にあなた独自のセットアップのための設定が必要になります。詳細については以下のセクションを参照してください。設定ファイルは `/opt/gravwell/etc/federator.conf` にあります。

### 設定例

以下の構成例では、*保護された*ネットワークセグメント内の2つのアップストリームインデクサーに接続し、2つの*信頼されていない*ネットワークセグメント上でインジェストサービスを提供しています。 各信頼されていないインジェストポイントには固有のIngest-Secretがあり、1つは特定の証明書とキーペアでTLSを提供しています。設定ファイルはローカルキャッシュも有効にし、フェデレータをGravwellインデクサーと信頼されていないネットワークセグメントの間でフォールトトレラントバッファとして機能させます。

```
[Global]
	Ingest-Secret = SuperSecretUpstreamIndexerSecret
	Connection-Timeout = 0
	Insecure-Skip-TLS-Verify = false
	Encrypted-Backend-target=172.20.232.105:4024
	Encrypted-Backend-target=172.20.232.106:4024
	Ingest-Cache-Path=/opt/gravwell/cache/federator.cache
	Max-Ingest-Cache=1024 #1GB
	Log-Level=INFO

[IngestListener "BusinessOps"]
        Ingest-Secret = CustomBusinessSecret
        Cleartext-Bind = 10.0.0.121:4023
        Tags=windows
        Tags=syslog

[IngestListener "DMZ"]
       Ingest-Secret = OtherRandomSecret
       TLS-Bind = 192.168.220.105:4024
       TLS-Certfile = /opt/gravwell/etc/cert.pem
       TLS-Keyfile = /opt/gravwell/etc/key.pem
       Tags=apache
       Tags=nginx
```

DMZ内のインジェスターは、TLS暗号化を使用して192.168.220.105:4024でフェデレータに接続できます。これらのインジェスタは、`apache` および `nginx` タグでタグ付けされたエントリの送信**のみ**許可されています。ビジネスネットワークセグメントのインジェスターは、クリアテキストを使って 10.0.0.0.121:4023 に接続し、`windows` と `syslog` タグのついたエントリを送信できます。誤ってタグ付けされたエントリはフェデレータによって拒否されますが、受け入れ可能なエントリはグローバルセクションで指定された2つのインデクサーに渡されます。

### トラブルシューティング

フェデレータの一般的な設定エラーには、次のようなものがあります:

* グローバル設定での不正なインジェストシークレット
* 誤ったバックエンドとターゲットの仕様
* バインド仕様が無効または既に取得されている場合
* アップストリームのインデクサやフェデレータが、信頼された証明書局が署名した証明書を持っていない場合の、強制的な証明書の検証(`Insecure-Skip-TLS-Verify` オプションを参照)
* 下流のインジェスターとのIngest-Secretの不一致


## インジェスト API

GravwellのインジェストAPIとコアインジェスターは、BSD 2-Clauseライセンスの下、完全にオープンソースです。 つまり、独自のインジェスターを作成して、Gravwellのエントリー生成を自社の製品やサービスに統合することができます。 コアのインジェストAPIはGoで書かれていますが、利用可能なAPI言語のリストは現在拡張中です。

[APIコード](https://github.com/gravwell/ingest)

[APIドキュメント](https://godoc.org/github.com/gravwell/ingest)

ファイルを監視し、ファイルに書き込まれた行をGravwellクラスタに送信する非常に基本的なインゲスターの例（コードは100行以下）[ここ](https://www.godoc.org/github.com/gravwell/ingest#example-package)で見ることができます

GravwellのGitHubページをチェックしてみてください。チームはインジェストAPIを継続的に改善し、他の言語に移植しています。コミュニティ開発は完全にサポートされているので、もしマージリクエストや言語の移植、オープンソース化した素晴らしい新しいインゲスターがあれば、Gravwellに知らせてください。 Gravwellチームは、あなたの努力をインゲスター・ハイライト・シリーズで紹介したいと思っています。
