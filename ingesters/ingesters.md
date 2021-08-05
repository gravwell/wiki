# インジェスター

このセクションには、Gravwellインジェスターを設定して実行するためのより詳細な説明が記載されています。

Gravwellで作成されたインジェスターは、BSDオープンソースライセンスでリリースされており、[Github](https://github.com/gravwell/ingesters)で見ることができます。インジェストAPIもオープンソースなので、独自のデータソース用のインジェスターを作成したり、追加の正規化や前処理を行ったり、その他あらゆる方法でインジェスターを作成できます。インジェストAPIのコードは[ここ](https://github.com/gravwell/ingest)にあります。

一般的に、インジェスターがGravwellにデータを送信するためには、インジェスターの認証のためにGravwellインスタンスの "Ingest Secret"を知る必要があります。これは、Gravwellサーバー上の`/opt/gravwell/etc/gravwell.conf`ファイル内の、`Ingest-Auth`エントリに設定されています。インジェスターがGravwell自身と同じシステム上で動作している場合、インストーラーは通常、この値を検出して自動的に設定してくれます。

Gravwell GUIには、(システムメニューカテゴリの下に) インジェスターページがあり、どのリモートインジェスターがアクティブに接続されているか及び、それらが接続されている時間と、プッシュされたデータ量を簡単に識別するために使用できます。

![](remote-ingesters.png)

注意： [レプリケーションシステム](#!configuration/replication.md)は、999MBを超えるエントリをレプリケーションしません。より大きなエントリは通常通りインジェストして検索できますが、レプリケーションからは省略されます。このページで詳しく説明しているインジェスターはすべて数キロバイト以下のエントリを作成する傾向があるので、これは99.9%のユースケースでは問題になりません。

## タグ

タグはGravwellの重要なコンセプトです。すべてのエントリは、それに関連付けられた単一のタグを持っています。これらのタグによって、基本的なレベルでデータを分離し、分類できます。例えば、Linuxシステムのログファイルから読み込んだエントリに "syslog"タグを適用したり、Windowsログに "winlog"を適用したり、生のネットワークパケットに "pcap"を適用できます。インジェスターは、どのタグをエントリに適用するかを決定します。

ユーザーから見れば、タグは "syslog"、"pcap-router"、または "default "などの文字列です。タグ名には以下の文字は使用できません:

```
!@#$%^&*()=+<>,.:;"'{}[]|\
```

また、タグ名を選択する際に、非印刷文字やタイプしにくい文字を使用するのは控えるべきです。☺という名前のタグにインジェストすることは*可能*ですが、それが良いアイデアというわけではありません。

### タグのワイルドカード

タグ名を選択する際には、Gravwellでは検索するタグ名を指定する際にワイルドカードを許可していることを覚えておいてください。タグ名を慎重に選択することで、後の検索を容易にできます。

例えば、5台のサーバーからシステムログを収集していて、そのうち2台がHTTPサーバー、2台がファイルサーバー、1台がEメールサーバーの場合、以下のタグの使用を選択できます:

* syslog-http-server1
* syslog-http-server2
* syslog-file-server1
* syslog-file-server2
* syslog-email-server1

これにより、[検索](#!search/search.md)でログを選択する際の柔軟性が高まります。`tag=syslog-*` を指定することで、すべてのシステムログを検索できます。また、`tag=syslog-http-*` と指定することですべてのHTTPサーバのログを検索することもできますし、`tag=syslog-http-server1` と指定することで単一のサーバを選択することもできます。また、`tag=syslog-http-*,syslog-email-*`のように、複数のワイルドカードグループを選択することもできます。

### タグの内部

このセクションを読むことは、Gravwellを使用するために必要ではありませんが、タグが内部的にどのように管理されているかを理解するのに役立つかもしれません。

内部的には、Gravwellの*インデクサー*はタグを16ビットの整数として格納します。各インデクサーはタグ名とタグ番号のマッピングを独自に管理しており、それは `/opt/gravwell/etc/tags.dat` にあります。Gravwellサポートから明示的に指示がない限り、このファイルを変更したり削除したりしないでください!

*インジェスター*がインデクサーに接続すると、使用したいタグ名のリストを送信します。インデクサーはタグ名とタグ番号の対応付けで応答します。インジェスターがそのインデクサーにエントリを送るときはいつでも、適切な*タグ番号*をエントリに追加します。

## グローバル設定パラメータ

ほとんどのコアインジェスターは、共通のグローバル設定パラメータのセットをサポートしています。共有されたグローバル構成パラメータは、[ingest config](https://godoc.org/github.com/gravwell/ingest/config#IngestConfig) パッケージを使用して実装されます。グローバル構成パラメータは、各Gravwellインジェスター設定ファイルのグローバルセクションで指定する必要があります。以下のグローバルインジェスターパラメータが利用可能です:

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

Ingest-Secretパラメータは、インジェスト認証に使用するトークンを指定します。 ここで指定されたトークンは、GravwellインデクサーのIngest-Authパラメータと一致しなければなりません。

### Connection-Timeout

Connection-Timeoutパラメータは、インデクサへの接続をあきらめるまでの待ち時間を指定します。 空のタイムアウトは、インジェスターが開始するまで永遠に待つことを意味します。タイムアウトは、分、秒、または時間の長さで指定する必要があります。

#### 例
```
Connection-Timeout=30s
Connection-Timeout=5m
Connection-Timeout=1h
```

### Insecure-Skip-TLS-Verify

Insecure-Skip-TLS-Verify トークンは、暗号化された TLS トンネルを介して接続する際に、不正な証明書を無視するようにインジェスターに指示します。その名が示すように、TLS によって提供されるすべての認証は窓から投げ出され、攻撃者は簡単に中間者による TLS 接続を行うことができます。 インジェスト接続は依然として暗号化されていますが、接続は決して安全ではありません。 デフォルトでは、TLS 証明書は検証され、証明書の検証に失敗すると接続は失敗します。

#### 例
```
Insecure-Skip-TLS-Verify=true
Insecure-Skip-TLS-Verify=false
```

### Rate-Limit

Rate-Limit パラメータは、インジェスターが消費できる最大帯域幅を設定します。これは、遅い接続でインデクサーと通信する"バースト型"インジェスターを設定する場合に便利です。

引数は数値の後にオプションのレートサフィックス、例えば `1048576` や `10Mbit` を指定します。以下のサフィックスが存在します。

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

Cleartext-Backend-Target は、Gravwell インデクサーのホストとポートを指定します。 インジェスターはクリアテキストTCP接続を使用してインデクサーに接続します。 ポートが指定されていない場合は、デフォルトのポート4023が使用されます。 Cleartext 接続は、IPv6 と IPv4 の両方の宛先をサポートします。 **複数の Cleartext-Backend-Targets を指定して、複数のインデクサー間でインジェスターの負荷分散を行うことができます。**

#### 例
```
Cleartext-Backend-Target=192.168.1.1
Cleartext-Backend-Target=192.168.1.1:4023
Cleartext-Backend-Target=DEAD::BEEF
Cleartext-Backend-Target=[DEAD::BEEF]:4023
```

### Encrypted-Backend-Target

Encrypted-Backend-Target は、Gravwell インデクサーのホストとポートを指定します。インジェスターはTCPを介してインデクサーに接続し、完全なTLSハンドシェイク/証明書の検証を実行します。 ポートが指定されていない場合は、デフォルトのポート4024が使用されます。 暗号化された接続は、IPv6 と IPv4 の両方の宛先をサポートします。 **複数のインデクサー間でインジェスターの負荷分散を行うために、複数の暗号化バックエンドターゲットを指定できます。**

#### 例
```
Encrypted-Backend-Target=192.168.1.1
Encrypted-Backend-Target=192.168.1.1:4023
Encrypted-Backend-Target=DEAD::BEEF
Encrypted-Backend-Target=[DEAD::BEEF]:4023
```

### Pipe-Backend-Target

Pipe-Backend-Target はフルパスで Unix 名前付きソケットを指定します。 Unixの名前付きソケットは、非常に高速でオーバーヘッドが少ないため、インデクサーと同居しているインジェスターに最適です。 インジェスターごとに1つのPipe-Backend-Targetしかサポートされていませんが、パイプはクリアテキストや暗号化された接続と並行して多重化できます。

#### 例
```
Pipe-Backend-Target=/opt/gravwell/comms/pipe
Pipe-Backend-Target=/tmp/gravwellpipe
```

### Ingest-Cache-Path

Ingest-Cache-Path は、インジェストされたデータのローカルキャッシュを有効にします。 有効にすると、インデクサーにエントリを転送できない場合、インジェスターはローカルにキャッシュできます。 インジェストキャッシュは、リンクがダウンしたときや、Gravwellクラスタを一時的にオフラインにする必要がある場合に、データを失わないようにするのに役立ちます。 長期的なネットワーク障害が発生してもインジェスターがホストディスクを埋め尽くすことがないように、Max-Ingest-Cache値を必ず指定してください。 ローカルインジェストキャッシュはインデクサーに直接インジェストするほど高速ではないので、インデクサーのようにインジェストキャッシュが毎秒200万エントリを処理できるとは期待しないでください。

注: ファイル・フォロワー・インジェスターでは、インジェスト・キャッシュは**有効にすべきではありません**。このインジェスターはディスク上のファイルから直接読み込み、各ファイル内の位置を追跡するため、キャッシュは必要ありません。

#### 例
```
Ingest-Cache-Path=/opt/gravwell/cache/simplerelay.cache
Ingest-Cache-Path=/mnt/storage/networklog.cache
```

### Max-Ingest-Cache

Max-Ingest-Cache は、キャッシュが使用されている時にインジェスターが消費するストレージ容量を制限します。 キャッシュの最大値はメガバイトで指定されます。1024の値はインジェスターが新しいエントリの受け入れを停止する前に1GBのストレージを消費できることを意味します。 キャッシュが一杯になってもキャッシュシステムは古いエントリを上書きしません。これは、攻撃者がネットワーク接続を混乱させ、インジェスターが混乱が起こった時点で潜在的に重要なデータを上書きさせることができないように設計されています。

#### 例
```
Max-Ingest-Cache=32
Max-Ingest-Cache=1024
Max-Ingest-Cache=10240
```

### Cache-Depth

Cache-Depth は、インメモリ・バッファに保持するエントリ数を設定します。既定値は 128 で、Ingest-Cache-Path が無効になっている場合でも、インメモリ・バッファは常に有効になります。Cache-Depth を大きな値に設定すると、より多くメモリを消費する代わりにインジェスターのバースト動作を吸収できます。

#### 例
```
Cache-Depth=256
```

### Cache-Mode

Cache-Modeは、実行時のバッキングキャッシュの動作を設定します(Ingest-Cache-Pathを設定することで有効になります)。利用可能なモードは、"always" と "fail" です。"always"モードでは、キャッシュは常に有効で、インジェスターはインメモリバッファ(Cache-Depthで設定)が一杯になるといつでもディスクにエントリを書き込めるようになります。これは、インデクサーとの接続がダウンしている場合や遅い場合、あるいはインジェスターがインデクサーとの接続で可能な量を超えるデータをプッシュしようとしている場合に発生する可能性があります。"always"モードを使用することで、インジェスターがエントリを削除したり、データのインジェストをブロックしたりすることはありません。Cache-Mode を "fail" に設定すると、すべてのインデクサ接続がダウンしているときにのみ有効になるようにキャッシュの動作が変更されます。

#### 例
```
Cache-Mode=always
Cache-Mode=fail
```

### Log-File

インジェスターは、インストールや構成の問題をデバッグする際に役立つように、エラーやデバッグ情報をログファイルに記録できます。 空の Log-File パラメータは、ファイルのロギングを無効にします。

#### 例
```
Log-File=/opt/gravwell/log/ingester.log
```

### Log-Level

Log-Levelパラメータは、ログファイルと "gravwell"タグの下でインデクサに送られるメタデータの両方について、各インジェスターのロギングシステムを制御します。 ログレベルを INFO に設定すると、ファイルフォロワーが新しいファイルを追跡したときやシンプルリレーが新しい TCP 接続を受信したときなどに、インジェスターに詳細なログを記録するように指示します。一方、レベルを ERROR に設定すると、最も重要なエラーのみがログに記録されます。ほとんどの場合、WARN レベルが適切です。以下のレベルがサポートされています:

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

Source-Override パラメータは、各エントリにアタッチされている SRC データ項目をオーバーライドします。 SRC 項目は IPv6 または IPv4 アドレスで、通常はインジェスターが実行されているマシンの外部IPアドレスです。

#### 例
```
Source-Override=10.0.0.1
Source-Override=0.0.0.0
Source-Override=DEAD:BEEF::FEED:FEBE
```

### Log-Source-Override

多くのインジェスターは、監査、健全性とステータス、一般的なインジェスト・インフラストラクチャのロギングを目的として、`gravwell`タグにエントリを出すことができます。 一般的に、これらのエントリは、インデクサーから見たインジェスターのソースIPアドレスをSRCフィールドに使用します。そのため、インジェスターによって実際に生成されたエントリのみソースIPフィールドをオーバーライドする機能が役立ちます。良い例としては、Gravwellフェデレータの `Log-Source-Override` を使用して、フェデレータを通過するすべてのエントリではなく、ヘルスとステータスのエントリの SRC フィールドだけを変更することが挙げられます。

`Log-Source-Override` 設定パラメータには、IPv4またはIPv6の値をパラメータとして指定する必要があります。

#### 例
```
Log-Source-Override=10.0.0.1
Log-Source-Override=0.0.0.0
Log-Source-Override=DEAD:BEEF::FEED:FEBE
Log-Source-Override=::1
```

## データコンシューマー設定

グローバル設定オプションの他に、設定ファイルを使用する各インジェスターは、少なくとも1つの*データコンシューマー*を定義する必要があります。データコンシューマーは、インジェスターに伝える以下の設定の定義です:

* データ取得先
* データに使用するタグ
* 特別なタイムスタンプ処理ルール
* SRC フィールドなどのフィールドのオーバーライド

シンプルリレーインジェスターと HTTP インジェスターは「リスナー」を定義し、ファイルフォロワーは「フォロワー」を使用し、netflowインジェスターは「コレクター」を定義します。以下の個々のインジェスターのセクションでは、インジェスターの特定のデータコンシューマータイプと、それらが必要とする可能性のある固有の構成について説明します。次の例では、ファイルフォロワーインジェスターが「フォロワー」データコンシューマーを定義して特定のディレクトリからデータを読み取る方法を示しています:

```
[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
```

データソース(`Base-Directory` と `File-Filter` ルールによる)、使用するタグ(`Tag-Name` による)、そして受信データのタイムスタンプを解析するための追加ルール(`Assume-Local-Timezone`)を指定していることに注意してください。

## Time Parsing Overrides

ほとんどのインジェスタは、データからタイムスタンプを抽出することで、各エントリにタイムスタンプを適用しようとします。このタイムスタンプの抽出を微調整するために、各*データコンシューマー*に適用できるオプションがいくつかあります:

* `Ignore-Timestamps` (boolean): `Ignore-Timestamps=true` を設定すると、インジェスターはタイムスタンプを抽出しようとするのではなく、各エントリに現在の時刻を適用するようになります。これは、非常に支離滅裂なデータが入力されている場合にデータをインジェストするための唯一のオプションになります。
* `Assume-Local-Timezone` (boolean): デフォルトでは、タイムスタンプにタイムゾーンが含まれていない場合、インジェスターはそれをUTCタイムスタンプとみなします。`Assume-Local-Timezone=true`を設定すると、インジェスターはローカルコンピュータのタイムゾーンが何であれ、それを仮定するようになります。これはTimezone-Overrideオプションとは相互に排他的です。
* `Timezone-Override` (string): `Timezone-Override`を設定すると、タイムゾーンを含まないタイムスタンプは指定されたタイムゾーンで解析されるべきであることをインジェスターに伝えます。例えば、`Timezone-Override=US/Pacific`を設定すると、インジェスターは受信したタイムスタンプをあたかも米国太平洋時間であるかのように扱うように指示します。許容可能なタイムゾーン名の完全なリストについては、[このページ](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)を参照してください( 'TZデータベース名'列)。Assume-Local-Timezoneとは相互に排他的です。
* `Timestamp-Format-Override` (string): このパラメータは、インジェスターにデータの中から特定のタイムスタンプフォーマットを探すように指示します。例を挙げたオーバーライドの完全なリストについては、[the timegrinder documentation](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder)を参照してください。

KinesisとGoogle Pub/Subインジェスターには、`Ignore-Timestamps`オプションはありません。KinesisとGoogle Pub/Subのインジェスターには、`Ignore-Timestamps`オプションはありません。KinesisとPub/Subは、すべてのエントリに到着タイムスタンプを含みます。データコンシューマ定義で `Parse-Time=true` が指定された場合、インジェスターは代わりにメッセージ本文からタイムスタンプを抽出しようとします。追加情報については、これらのインジェスターのそれぞれのセクションを参照のこと。

カスタムタイムスタンプフォーマットは多くのインジェスターでサポートされています。詳細は[カスタムタイムフォーマット](#!ingesters/customtime/customtime.md)を参照してください。

## シンプルリレー

[完全な設定とドキュメント](#!ingesters/simple_relay.md).

Simple Relay は複数の TCP ポートや UDP ポートをリッスンすることができるテキストインジェスターです。 各ポートにはタグとインジェスト標準 (例: RFC5424 のパースや単純な改行区切りエントリ) を割り当てることができます。 Simple Relay は、リモートの syslog エントリをインジェストしたり、ネットワーク接続を介してテキストログを投げることができるあらゆるデータソースからデータを取得したりするために最適なインジェスターです。

### リスナーの例

```
[Listener "default"]
	Bind-String="0.0.0.0:7777" # Bind to all interfaces, with TCP implied
	Assume-Local-Timezone=false # Do not assume the local timezone if not provided
	Tag-Name=foo # Set the tag to ingest into
	Source-Override="DEAD::BEEF" # Override the source for just this listener
	Ignore-Timestamps=true

[Listener "syslogtcp"]
	Bind-String="tcp://0.0.0.0:601" # standard RFC5424 reliable syslog
	Reader-Type=rfc5424 # Syslog reader
	Tag-Name=syslog
	Keep-Priority=true # Leave the <nnn> priority value at the beginning of the syslog message
	Assume-Local-Timezone=true # If a time format does not have a timezone, assume local time

[Listener "syslogudp"]
	Bind-String="udp://0.0.0.0:514" #standard UDP based RFC5424 syslog
	Reader-Type=rfc5424
	Tag-Name=syslog
	Timezone-Override=America/Denver

[JSONListener "json"]
    Bind-String=[2600:1f18:63ef:e802:355f:aede:dbba:2c03]:901
    Extractor="field1"
    Default-Tag=json
    Tag-Match=test1:tag1
    Tag-Match=test2:tag2
    Tag-Match=test3:tag3

[Listener "tls"]
	Bind-String="tls://0.0.0.0:514" # TLS over port 514
	Reader-Type=rfc5424
	Cert-File=/opt/gravwell/cert.pem
	Key-File=/opt/gravwell/key.pem
	Tag-Name=syslog
	Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
```

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-simple-relay
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラをダウンロードします。インジェスターをインストールするには、スーパーユーザとして (例えば `sudo` コマンドで) 以下のコマンドを実行してください:

```
root@gravserver ~ # bash gravwell_simple_relay_installer.sh
```

Gravwellサービスが同じマシン上に存在する場合、インストールスクリプトは自動的に `Ingest-Auth` パラメータを抽出して設定し、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラーは認証トークンとGravwellインデクサーのIPアドレスを要求します。インストール中にこれらの値を設定するか、空欄のままにして、`/opt/gravwell/etc/simple_relay.conf`の設定ファイルを手動で修正できます。

## ファイルフォロワー

ファイルフォロワーインジェスターは、ローカルシステム上のファイルを監視するように設計されており、Gravwellとネイティブに統合できなかったり、ネットワーク接続を介してログを送信できないソースからログをキャプチャします。ファイルフォロワーは、LinuxとWindowsの両方のフレーバーがあり、任意の行で区切られたテキストファイルを追跡できます。 ファイルローテーションと互換性があり、強力なパターンマッチングシステムを採用しており、ログファイルに一貫性のない名前を付けるアプリケーションに対応しています。

### フォロワーの例

```
[Follower "auth"]
	Base-Directory="/var/log/"
	File-Filter="auth.log,auth.log.[0-9]" # Look for all authorization log files
	Tag-Name=auth
	Assume-Local-Timezone=true 
	Ignore-Timestamps=true

[Follower "test"]
	Base-Directory="/tmp/testing/"
	File-Filter="*"
	Tag-Name=default
	Recursive=true
	Ignore-Line-Prefix="#" # ignore lines beginning with #
	Ignore-Line-Prefix="//"
	Regex-Delimiter="###"

[Follower "timevoodoo"]
	Base-Directory="/tmp/time/"
	File-Filter="*"
	Tag-Name=default
	Timestamp-Delimited=true
	Timestamp-Regex=`[JFMASOND][anebriyunlgpctov]+\s+\S{1,2},\s+\d{4}\s+\d{1,2}:\d\d:\d\d,\d+\s+\S{2}\s+\S+`
	Timestamp-Format-String="Jan _2, 2006 3:04:05,999 PM MST"
```

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-file-follow
```

それ以外の場合は、[ダウンロード](#!Quickstart/downloads.md)からインストーラをダウンロードしてください。Windows システムでは、ダウンロードした実行ファイルを実行し、インストーラのプロンプトに従ってください。Linuxでは、スーパーユーザとして(例: `sudo` コマンドで)以下のコマンドを実行してインジェスターをインストールしてください:

```
root@gravserver ~ # bash gravwell_file_follow_installer.sh
```

Gravwellのサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメータを抽出し、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラは認証トークンとGravwellインデクサーのIPアドレスを要求します。これらの値はインストール時に設定するか、あるいは空欄にして、`/opt/gravwell/etc/file_follow.conf`の設定ファイルを手動で修正することができます。

### 設定例

ファイルフォロワーの設定は、Windows と Linux の両方のバージョンでほぼ同じです。より詳細な設定情報は [ファイルフォロワー](file_follow.md)にあります。

#### Windows

Windows 構成ファイルは、デフォルトでは `C:\Program Files\gravwell\file_follow.cfg` にあります。 Windows File Follower は Windows サービスとして動作します。 そのステータスは、コマンドプロンプトで `sc query GravwellFileFollow` を発行することで問い合わせることができます。 Windows CBSログファイルを追跡する構成例は以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Cleartext-Backend-target=172.20.0.2:4023 #example of adding a cleartext connection
State-Store-Location="c:\\Program Files\\gravwell\\file_follow.state"
Ingest-Cache-Path="c:\\Program Files\\gravwell\\file_follow.cache"
Log-Level=ERROR #options are OFF INFO WARN ERROR
#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Follower "cbs"]
        Base-Directory="C:\\Windows\\Logs\\CBS"
        File-Filter="CBS.log" #we are looking for just the CBS log
        Tag-Name=auth
        Assume-Local-Timezone=true
```

#### Linux

linux設定ファイルはデフォルトでは `/opt/gravwell/etc/file_follow.conf` にあります。 カーネル、dmesg、debian のインストールログを監視する設定例は以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Cleartext-Backend-target=172.20.0.1:4023 #example of adding a cleartext connection
Cleartext-Backend-target=172.20.0.2:4023 #example of adding another cleartext connection
#Encrypted-Backend-target=127.1.1.1:4024 #example of adding an encrypted connection
#Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
State-Store-Location=/opt/gravwell/etc/file_follow.state
Log-Level=ERROR #options are OFF INFO WARN ERROR
Max-Files-Watched=64
#Ingest-Cache-Path=/opt/gravwell/cache/file_follow.cache #allows for ingested entries to be cached when indexer is not available
#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
[Follower "auth"]
        Base-Directory="/var/log/"
        File-Filter="auth.log,auth.log.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
[Follower "packages"]
        Base-Directory="/var/log"
        File-Filter="dpkg.log,dpkg.log.[0-9]" #we are looking for all dpkg files
        Tag-Name=dpkg
        Ignore-Timestamps=true
[Follower "kernel"]
        Base-Directory="/var/log"
        File-Filter="dmesg,dmesg.[0-9]"
        Tag-Name=kernel
        Ignore-Timestamps=true
[Follower "kernel2"]
        Base-Directory="/var/log"
        File-Filter="kern.log,kern.log.[0-9]"
        Tag-Name=kernel
        Ignore-Timestamps=true
```

## HTTP

HTTP インジェスターは、1 つ以上のパスに HTTP リスナーを設定します。HTTP リクエストがこれらのパスのいずれかに送信されると、リクエストの Body が単一のエントリとしてインジェストされます。

これはスクリプト可能なデータのインジェストには非常に便利な方法で、`curl`コマンドを使うと標準入力をボディとして使ったPOSTリクエストが簡単にできるからです。

### リスナーの例

すべてのインジェスタで使用される汎用設定パラメータに加えて、HTTP POSTインジェスタには、組み込みウェブサーバの動作を制御する2つのグローバル設定パラメータが追加されています。 最初の設定パラメータは `Bind` オプションで、ウェブサーバがリッスンするインターフェイスとポートを指定します。 2つ目は `Max-Body` パラメータで、ウェブサーバが許可するPOSTの大きさを制御します。 Max-Bodyパラメータは、不正なプロセスが単一のエントリとして非常に大きなファイルをGravwellインスタンスにアップロードしようとするのを防ぐための良いセーフティネットです。Gravwellは1つのエントリとして1GBまでサポートできますが、我々はそれを推奨しません。

複数の「リスナー」定義を定義することで、特定のURLから特定のタグにエントリーを送信できます。 例の構成では、天気IoTデバイスとスマートサーモスタットからデータを受け取る2つのリスナーを定義しています。

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

### Splunk HEC 互換性

HTTPインジェスターは、Splunk HTTPイベントコレクターとAPI互換性のあるリスナーブロックをサポートしています。 この特別なリスナーブロックにより、Splunk HECにデータを送信できるエンドポイントであれば、Gravwell HTTP インジェスターにもデータを送信できるように構成を簡略化することができます。 HEC互換の設定ブロックは以下のようになります。

```
[HEC-Compatible-Listener "testing"]
	URL="/services/collector/event"
	TokenValue="thisisyourtoken"
	Tag-Name=HECStuff

```

`HEC-Comptabile-Listener`ブロックには、`TokenValue`と`Tag-Name`の設定項目が必要です。`URL`の設定項目が省略された場合は、`/services/collector/event`がデフォルトになります。

`Listener`と`HEC-Compatible-Listener`の両方の構成ブロックを同じHTTPインジェスターに指定することができます。

### ヘルスチェック

いくつかのシステム(AWSロードバランサーなど)では、プロービングされ、「命の証明」として解釈できる認証されていないURLを必要とします。 HTTPインジェスターはURLを提供するように設定することができ、任意のメソッド、ボディ、クエリパラメータでアクセスされた場合、常に200 OKを返すようになります。このヘルスチェックのエンドポイントを有効にするには、`Health-Check-URL` スタンザをグローバル設定ブロックに追加します。

ここでは、ヘルスチェックURL `/logbot/are/you/alive` を使用した最小限の設定例を示します：

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

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-http-ingester
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラをダウンロードします。Gravwell サーバー上のターミナルを使用して、スーパーユーザーとして(例: `sudo` コマンドで)以下のコマンドを実行し、インジェスターをインストールします:

```
root@gravserver ~ # bash gravwell_http_ingester_installer_3.0.0.sh
```

Gravwellサービスが同じマシン上に存在する場合、インストールスクリプトは自動的に `Ingest-Auth` パラメータを抽出して、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラーは認証トークンとGravwellインデクサーのIPアドレスを要求します。インストール中にこれらの値を設定するか、空欄のままにして、`/opt/gravwell/etc/gravwell_http_ingester.conf`内の設定ファイルを手動で修正します。

### HTTPS の設定

デフォルトでは、HTTPインジェスターはクリアテキスト HTTP サーバを実行しますが、x509 TLS 証明書を使用して HTTPS サーバを実行するように設定することもできます。 HTTPインジェスターを HTTPS サーバとして設定するには、`TLS-Certificate-File` と `TLS-Key-File` パラメータを使用して、グローバル設定スペースで証明書と鍵の PEM ファイルを提供します。

HTTPS を有効にしたグローバル設定の例は、以下のようになります:

```
[Global]
	TLS-Certificate-File=/opt/gravwell/etc/cert.pem
	TLS-Key-File=/opt/gravwell/etc/key.pem
```

#### リスナーの認証

各 HTTPインジェスターリスナーは、認証を強制するように設定できます。サポートされている認証方法は以下の通りです:

* none
* basic
* jwt
* cookie
* preshared-token
* preshared-parameter

none以外の認証システムを指定する場合は、認証情報を指定しなければなりません。`jwt`、`cookie`、クッキー認証システムはユーザ名とパスワードを必要とし、`preshared-token`、`preshared-parameter` はトークンの値とオプションのトークン名を指定しなければなりません。

警告: 他のウェブページと同様に、認証はクリアテキスト接続では安全ではなく、トラフィックを盗み見できる攻撃者はトークンやクッキーをキャプチャできます。

#### 認証なし

デフォルトの認証方法は「なし」で、インジェスターに到達できる人なら誰でもエントリをプッシュできるようになっています。`Basic`認証メカニズムは HTTP Basic 認証を利用します。

以下に基本認証システムを使ったリスナーの例を示します:

```
[Listener "basicauth"]
	URL="/basic/data"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

基本的な認証でエントリを送信するための curl コマンドの例は次のようになります:

```
curl -d "only i can say hi" --user secretuser:secretpassword -X POST http://10.0.0.1:8080/basic/data
```

#### JWT認証

JWT認証システムでは、認証に暗号署名済みのトークンを使用します。jwt認証を使用する際には、クライアントが認証してトークンを受け取るログインURLを指定する必要があります。jwtトークンの有効期限は48時間です。認証はログインURLに `POST` リクエストを送信し、フォームフィールドに `username` と `password` を入力することで行われます。

jwt認証を使用してHTTPインジェスターで認証するには2つのステップがあり、追加の設定パラメータが必要です。 以下に設定例を示します:

```
[Listener "jwtauth"]
	URL="/jwt/data"
	LoginURL="/jwt/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

エントリを送信するには、エンドポイントが最初に認証してトークンを取得する必要があり、トークンはその後最大48時間まで再利用できます。リクエストが 401 レスポンスを受信した場合、クライアントは再認証を行う必要があります。ここでは、curl を使用して認証を行い、データをプッシュする例を示します。

```
x=$(curl -X POST -d "username=user1&password=pass1" http://127.0.0.1:8080/jwt/login) #grab the token and stuff it into a variable
curl -X POST -H "Authorization: Bearer $x" -d "this is a test using JWT auth" http://127.0.0.1:8080/jwt/data #send the request with the token
```

#### Cookie認証

cookie認証の仕組みは、状態を制御する方法以外は JWT 認証と実質的に同じです。cookie認証を使用するリスナーは、ログインページで設定されたクッキーを取得するために、クライアントがユーザ名とパスワードでログインする必要があります。インジェストURLへのそれ以降のリクエストは、各リクエストでcookieを提供しなければなりません。

以下に設定ブロックの例を示します:

```
[Listener "cookieauth"]
	URL="/cookie/data"
	LoginURL="/cookie/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

いくつかのデータをインジェストする前にログインしてcookieを取得する curl コマンドの例は次のようになります:

```
curl -c /tmp/cookie.txt -d "username=user1&password=pass1" localhost:8080/cookie/login
curl -X POST -c /tmp/cookie.txt -b /tmp/cookie.txt -d "this is a cookie data" localhost:8080/cookie/data
```

#### 事前共有トークン

プレシェアドトークン認証メカニズムは、ログインメカニズムではなくプレシェアドシークレットを使用します。Preshared secret は、Authorization ヘッダで各リクエストと一緒に送信されることが期待されています。 多くのHTTPフレームワークはこのタイプのインジェストを期待しており、Splunk HECやサポートしているAWS KinesisやLambdaインフラストラクチャなどがそうです。 事前共有トークンリスナーを使用することで、Splunk HEC のプラグインの代替となるキャプチャシステムを定義できます。

注: `TokenName` の値を定義しない場合、デフォルトの `Bearer` が使用されます。

事前共有トークンを定義する設定例:

```
[Listener "presharedtoken"]
	URL="/preshared/token/data"
	Tag-name=token
	AuthType="preshared-token"
	TokenName=foo
	TokenValue=barbaz
```

事前に共有された秘密を使ってデータを送信する curl コマンドの例:

```
curl -X POST -H "Authorization: foo barbaz" -d "this is a preshared token" localhost:8080/preshared/token/data
```

#### 事前共有パラメータ

事前共有パラメータ認証メカニズムは、クエリパラメータとして提供されるプレシェアードシークレットを使用します。 `Preshared-parameter` システムは、認証トークンを URL に埋め込むことで、通常は認証をサポートしていないデータプロデューサをスクリプトで作成したり、使用する際に有用です。

注: URL に認証トークンを埋め込むということは、プロキシや HTTP ロギングインフラストラクチャが認証トークンをキャプチャしてログに記録することを意味します。

事前共有パラメータを定義する設定例:

```
[Listener "presharedtoken"]
	URL="/preshared/parameter/data"
	Tag-name=token
	AuthType="preshared-parameter"
	TokenName=foo
	TokenValue=barbaz
```

事前に共有された秘密を使ってデータを送信する curl コマンドの例:

```
curl -X POST -d "this is a preshared parameter" localhost:8080/preshared/parameter/data?foo=barbaz
```

### リスナーの方法

HTTPインジェスターは実質的に任意のメソッドを使用するように設定できますが、データは常にリクエストの本文にあることが期待されます。

例えば、以下は PUT メソッドを期待するリスナーの設定です:

```
[Listener "test"]
	URL="/data"
	Method=PUT
	Tag-Name=stuff
```

対応する curl コマンドは次のようになります:

```
curl -X PUT -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```

HTTP インジェスターは、特殊文字を含まないほとんどすべての ASCII 文字列を受け入れ、メソッドの仕様外になることがあります。

```
[Listener "test"]
	URL="/data"
	Method=SUPER_SECRET_METHOD
	Tag-Name=stuff
```

```
curl -X SUPER_SECRET_METHOD -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```

## マスファイルインジェスター

マスファイルインジェスターは、多くのソースから多くのログのアーカイブをインジェストするための、非常に強力ではありますが専門的なツールです。

### 使用例

Gravwellユーザーは、潜在的なネットワーク侵害を調査する際にこのツールを使用したことがあります。ユーザーは50以上の異なるサーバーからのApacheログを持っていて、それらをすべて検索する必要がありました。それらを次々にインジェストすると、時間的なインデックス作成のパフォーマンスが低下します。このツールは、ログエントリの一時的な性質を維持し、確実なパフォーマンスを確保しながらファイルをインジェストするために作成されました。マスファイルインジェスターは、インジェストマシンがインジェスト前にソースログを最適化するのに十分なスペース（ストレージとメモリ）を持っている場合に最適に動作します。 最適化フェーズは、インジェスト時と検索時にGravwellストレージシステムにかかる圧力を軽減するのに役立ち、インシデント対応者が迅速に移動し、ログデータに短時間でパフォーマンスの高いアクセスを得られるようにします。

### 注意事項

この大量ファイルインジェスターはコマンドラインパラメータを介して動作し、サービスとして動作するようには設計されていません。 コードは [Github](https://github.com/gravwell/ingesters) にあります。

```
Usage of ./massFile:
  -clear-conns string
        comma seperated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -no-ingest
        Optimize logs but do not perform ingest
  -pipe-conn string
        Path to pipe connection
  -s string
        Source directory containing log files
  -skip-op
        Assume working directory already has optimized logs
  -tag-name string
        Tag name for ingested data (default "default")
  -timeout int
        Connection timeout in seconds (default 1)
  -tls-conns string
        comma seperated server:port list of TLS connections
  -tls-private-key string
        Path to TLS private key
  -tls-public-key string
        Path to TLS public key
  -tls-remote-verify string
        Path to remote public key to verify against
  -w string
        Working directory for optimization
```

## Windowsイベントサービス

Gravwell Windows イベントインジェスターは Windows マシン上でサービスとして動作し、Windows イベントを Gravwell インデクサーに送信します。 インジェスターはデフォルトでは `System`、`Application`、`Setup`、`Security` チャンネルからイベントを消費します。 各チャンネルは、特定のイベントやプロバイダのセットから消費するように設定できます。

### イベントチャンネルの例

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

### インストール

[ダウンロード](#!Quickstart/downloads.md)から、Gravwell Windowsインジェスターインストーラーをダウンロードします。

.msi インストールウィザードを実行して、Gravwell イベントサービスをインストールします。最初のインストールでは、インストールウィザードは、インデクサーエンドポイントとインジェストシークレットを設定するように促されます。その後のインストールやアップグレードでは、常駐の設定ファイルを識別し、プロンプトは表示されません。

インジェスターは、`%PROGRAMDATA%\gravwell\eventlogconfig.cfg`にある`config.cfg`ファイルで構成されます。 設定ファイルは他のGravwellインジェスターと同じ形式で、インデクサ接続を構成する`[Global]`セクションと複数の`EventChannel`定義を持ちます。

インデクサ接続を変更したり、複数のインデクサを指定するには、接続IPアドレスをGravwellサーバーのIPに変更し、Ingest-Secret値を設定します。 この例では、暗号化されたトランスポートを構成しています:

```
Ingest-Secret=YourSecretGoesHere
Encrypted-Backend-target=ip.addr.goes.here:port
```

一度設定したこのファイルは、イベントを収集する他のWindowsシステムにコピーできます。

#### サイレントインストール

Windows イベントインジェスターは、自動展開と互換性があるように設計されています。 これは、ドメインコントローラがインストーラをクライアントにプッシュして、ユーザとの対話なしにインストールを起動できることを意味します。 サイレントインストールを強制するには、管理者権限で [msiexec](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/msiexec) で `/quiet` 引数を指定してインストーラを実行します。 このインストール方法では、デフォルトの設定がインストールされ、サービスが開始されます。

特定のパラメータを設定するには、変更した設定ファイルを `%PROGRAMDATA%\gravwell\eventlog\config.cfg` にプッシュしてサービスを再起動するか、`CONFIGFILE` 引数に `config.cfg` ファイルへの完全修飾パスを指定する必要があります。

あなたは `%PROGRAMDATA%\gravwell\eventlog` パスを作成する必要があるかもしれないことに注意してください。

グループポリシーのプッシュのための完全な実行シーケンスは以下のようになります:

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet
xcopy \\share\gravwell_config.cfg %PROGRAMDATA%\gravwell\eventlog\config.cfg
sc stop "GravwellEvent Service"
sc start "GravwellEvent Service"
```

または

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet CONFIGFILE=\\share\gravwell_config.cfg
```

### オプションのSysmon統合

sysinternals スイートの一部である Sysmon ユーティリティは、Windows システムを監視するための効果的で人気のあるツールです。優れた sysmon 設定ファイルの例が掲載されているリソースはたくさんあります。Gravwellでは、infosec Twitterのパーソナリティである@InfosecTaylorSwiftが作成した設定を好んで使用しています。

Gravwell Windows agent configファイルを編集して、`%PROGRAMDATA%\gravwell\eventlog\config.cfg`にある以下の行を追加してください:

```
[EventChannel "Sysmon"]
        Tag-Name=sysmon #Apply a new tag name
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

SwiftOnSecurityによる優れたsysmon設定ファイルの[ダウンロード](https://raw.githubusercontent.com/SwiftOnSecurity/sysmon-config/master/sysmonconfig-export.xml)

sysmon の[ダウンロード](https://technet.microsoft.com/en-us/sysinternals/sysmon)

以下のコマンドを実行して、管理者シェルを使って `sysmon` を設定してインストールします(Powershellでも動作します):

```
sysmon.exe -accepteula -i sysmonconfig-export.xml
```

標準のWindowsサービス管理からGravwellサービスを再起動します。

#### Sysmonでの設定例

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

### トラブルシューティング

Windowsインジェスターの接続性は、Webインターフェースのインジェスターページに移動することで確認できます。Windowsインジェスターが存在しない場合は、windows GUIからサービスの状態を確認するか、コマンドラインで `sc query GravwellEvents` を実行してください。

![](querystatus.png)

![](querystatusgui.png)

### Windowsでの検索例

デフォルトのタグ名が使用されていると仮定して、sysmonの全エントリを表示するには、以下の検索を実行してください:

```
tag=sysmon
```

すべてのWindowsイベントを完全に表示するには、以下を実行してください:

```
tag=windows
```

以下の検索では、`winlog`検索モジュールを使用して、特定のイベントやフィールドをフィルタリングして抽出できます。 非標準プロセスによるすべてのネットワーク作成を見るには:

```
tag=sysmon winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
table TIMESTAMP SourceHostname Image DestinationIP DestinationPort
```

ソースホスト別にネットワーク作成のチャートを作成する:

```
tag=sysmon winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
count by SourceHostname |
chart count by SourceHostname limit 10
```

不審なファイル作成を見る:

```
tag=sysmon winlog EventID==11 Image TargetFilename |
count by TargetFilename |
chart count by TargetFilename
```

## Netflow インジェスター

Netflow インジェスターは Netflow コレクターとして動作します（Netflow の役割の完全な説明は [wikipedia の記事](https://en.wikipedia.org/wiki/NetFlow) を参照してください）。これらの項目はその後、[netflow](#!search/netflow/netflow.md) 検索モジュールを使って分析できます。

### コレクターの例

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

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-netflow-capture
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md) からインストーラーをダウンロードします。Netflow インジェスターをインストールするには、単に root としてインストーラーを実行します（実際のファイル名には通常、バージョン番号が含まれます）:

```
root@gravserver ~ # bash gravwell_netflow_capture_installer.sh
```

ローカルマシン上に Gravwell インデクサーがない場合、インストーラーは Ingest-Secret 値とインデクサー(またはフェデター)の IP アドレスを要求します。さもなければ、既存の Gravwell 設定から適切な値を引き出します。いずれにせよ、インストール後に `/opt/gravwell/etc/netflow_capture.conf` の設定ファイルを確認してください。UDP ポート 2055 をリッスンする簡単な例は以下のようになります:

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

この設定では、`/opt/gravwell/comms/pipe`を介してローカルのインデクサーにエントリを送ることに注意してください。エントリには 'netflow' というタグが付けられています。

異なるポートで異なるタグを使用して listen している `Collector` エントリをいくつでも設定できます。

注: 現時点では、インジェスターは Netflow v5 のみをサポートしています。

## Networkインジェスター

Gravwellの第一の強みは、バイナリデータをインジェストできることです。これは、単にネットフローや他の凝縮されたトラフィック情報を保存するよりもはるかに優れた柔軟性を提供します。

### スニッファーの例

```
[Sniffer "spy1"]
	Interface="p1p1" #sniffing from interface p1p1
	Tag-Name="pcap"  #assigning tag  fo pcap
	Snap-Len=0xffff  #maximum capture size
	BPF-Filter="not port 4023" #do not sniff any traffic on our backend connection
	Promisc=true

[Sniffer "spy2"]
	Interface="p5p2"
	Source-Override=10.0.0.1
```

### インストール

Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install libpcap0.8 gravwell-network-capture
```

それ以外の場合は、[ダウンロード](#!Quickstart/downloads.md)からインストーラをダウンロードします。ネットワークインジェスターをインストールするには、単に root でインストーラを実行してください (ファイル名が若干異なる場合があります):

```
root@gravserver ~ # bash gravwell_network_capture_installer.sh
```

注: インジェスターを動作させるためには、libpcapがインストールされている必要があります。

可能であれば、network インジェスターをインデクサーと一緒に配置し、`clear-conn` や `tls-conn` リンクではなく、`pipe-conn` リンクを使用してデータを送信することを強くお勧めします。 networkインジェスターがエントリをプッシュするために使用しているリンクと同じリンクからキャプチャしている場合、リンクを急速に飽和させるフィードバックループを作成できます(例えば、eth0からキャプチャしながらeth0経由でインジェスターにエントリを送信するなど)。これを緩和するために、`BPF-フィルター`オプションを使用できます。

Gravwellバックエンドがインストールされているマシンにインジェスターがある場合、インストーラは自動的に正しい`Ingest-Secrets`の値を拾い、設定ファイルに入力します。そうでない場合は、インデクサのIPアドレスとインジェスト・シークレットの入力を求められます。いずれにしても、実行する前に `/opt/gravwell/etc/network_capture.conf` にある設定ファイルを確認してください。eth0からのトラフィックをキャプチャする例は以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
#Cleartext-Backend-target=127.1.0.1:4023 #example of adding a cleartext connection
#Encrypted-Backend-target=127.1.1.1:4023 #example of adding an encrypted connection
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO #options are OFF INFO WARN ERROR
Ingest-Cache-Path=/opt/gravwell/cache/network_capture.cache

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Sniffer "spy1"]
	Interface="eth0" #sniffing from interface eth0
	Tag-Name="pcap"  #assigning tag  fo pcap
	Snap-Len=0xffff  #maximum capture size
	#BPF-Filter="not port 4023" #do not sniff any traffic on our backend connection
	#Promisc=true
```

異なるインタフェースからキャプチャするために、`Sniffer` エントリを任意の数だけ設定できます。

ディスク容量が気になる場合は、`Snap-Len` パラメータを変更してパケットのメタデータのみをキャプチャするようにするとよいでしょう。通常、ヘッダのみをキャプチャするには96の値で十分です。

非常に高い帯域幅のリンクの可能性があるため、ネットワークキャプチャデータを独自のウェルに割り当てることもお勧めします。

これには、インデクサでパケットキャプチャタグ用の別のウェルを定義する設定が必要です。NetworkCapture インジェスターは、libpcap 構文に準拠した `BPF-Filter` パラメータを使用したネイティブ BPF フィルタリングもサポートしています。 ポート 22 上のすべてのトラフィックを無視するには、以下のようにスニッファーを設定できます:

```
[Sniffer "no-ssh"]
	Interface="eth0"
	Tag-Name="pcap"
	Snap-Len=0xffff
	BPF-Filter="not port 22"
```

インジェスターがインデクサーとは別のシステムにあり、エントリーがインジェストされるためにネットワークを通過しなければならない場合は、`BPF-Filter`を "not port 4023" (cleartextを使用している場合) または "not port 4024" (TLSを使用している場合)に設定する必要があります。

### ネットワーク検索の例

以下の検索では、IP 10.0.0.0.0/24クラスCサブネットから発信されていないRSTフラグが設定されたTCPパケットを検索し、IPごとにグラフ化します。このクエリを使用して、ネットワークからのアウトバウンドポートスキャンを迅速に特定できます。

```
tag=pcap packet tcp.RST==TRUE ipv4.SrcIP !~ 10.0.0.0/24 | count by SrcIP | chart count by SrcIP limit 10
```

![](portscan.png)

次の検索では、IPv6 トラフィックを探して FlowLabel を抽出し、これを数学演算に渡します。 これにより、パケットの長さを合計してチャートレンダラに渡すことで、フローごとのトラフィックのアカウンティングが可能になります。

```
tag=pcap packet ipv6.Length ipv6.FlowLabel | sum Length by FlowLabel | chart sum by FlowLabel limit 10
```

TCPペイロードで使用されている言語を識別するために、ネットワークデータをフィルタリングしてlangfindモジュールに渡すことができます。このクエリは、アウトバウンド HTTP リクエストを探し、TCP ペイロードのデータを langfind モジュールに渡します。これにより、アウトバウンド HTTP クエリで使用される人間の言語のチャートが生成されます。

```
tag=pcap packet ipv4.DstIP != 10.0.0.100 tcp.DstPort == 80 tcp.Payload | langfind -e Payload | count by lang | chart count by lang
```

![](langfind.png)

トラフィックアカウンティングは、レイヤ2でも実行できます。これは、イーサネットヘッダからパケット長を抽出し、その長さを宛先MACアドレスで合計し、トラフィックカウントでソートすることで達成されます。これにより、イーサネットネットワーク上の特におしゃべりな物理デバイスを迅速に特定できます:

```
tag=pcap packet eth.DstMAC eth.Length > 0 | sum Length by DstMAC | sort by sum desc | table DstMAC sum
```

同様のクエリでは、パケット数を介しておしゃべりなデバイスを識別できます。例えば、あるデバイスがスイッチにストレスを与えているが、大量のトラフィックにはならない小さなイーサネットパケットを積極的にブロードキャストしているかもしれません。

```
tag=pcap packet eth.DstMAC eth.Length > 0 | count by DstMAC | sort by count desc | table DstMAC count
```

非標準の HTTP ポートで動作する HTTP トラフィックを識別することが望ましいかもしれません。 これは、フィルタリングオプションを行使し、ペイロードを他のモジュールに渡すことで実現できます。 例えば、TCPポート80ではなく、特定のサブネットから発信されているアウトバウンドトラフィックを探し、ppacketペイロード内のHTTPリクエストを探すことで、異常なHTTPトラフィックを特定できます:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP tcp.DstPort !=80 ipv4.SrcIP ~ 10.0.0.0/24 tcp.Payload | regex -e Payload "(?P<method>[A-Z]+)\s+(?P<url>[^\s]+)\s+HTTP/\d.\d" | table method url SrcIP DstIP DstPort
```

![](nonstandardhttp.png)

## Kafka

Kafkaインジェスターは、[Apache Kafka](https://kafka.apache.org/)のコンシューマとして動作するように設計されており、GravwellがKafkaクラスタにアタッチしてデータを消費できるようになっています。Kafkaは、Gravwellに対して高可用性の[data broker](https://kafka.apache.org/uses#uses_logs)として機能することができます。Kafkaは、Gravwellフェデレータによって提供される役割の一部を引き受けたり、Gravwellを既存のデータフローに統合する負担を軽減できます。データがすでにKafkaに流れている場合、Gravwellを統合するのはapt-getするだけです。

Gravwell Kafkaインジェスターは、単一のインデクサーのコロケーションされたインジェストポイントとして最適です。 KafkaクラスタとGravwellクラスタを運用している場合は、Gravwellのインジェスト層でKafkaのロードバランシング特性を重複させない方が良いでしょう。KafkaインジェスターをGravwellインデクサーと同じマシンにインストールし、Unixの名前付きパイプ接続を使用します。各インデクサーにそれぞれKafkaインジェスターを設定することで、Kafkaクラスターがロードバランシングを管理できるようになります。

ほとんどのKafkaの設定では、データの耐久性を保証するために、消費者が利用できないときは、データは非揮発性のストレージに保存されます。そのため、KafkaインジェスターでGravwellインジェストキャッシュを有効にすることはお勧めしませんが、代わりにKafkaがデータの耐久性を提供してくれます。

### コンシューマーの例

```
[Consumer "default"]
	Leader="127.0.0.1"
	Default-Tag=default   #send bad tag names to default tag
	Tags=*                #allow all tags
	Topic=default
	Tag-Header=TAG        #look for the tag in the kafka TAG header
	Source-Header=SRC     #look for the source in the kafka SRC header

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

### インストール

Kafka インジェスターは、Gravwell の debian リポジトリで debian パッケージとして、またシェルインストーラとして [ダウンロード](#!quickstart/downloads.md) で入手できます。リポジトリからのインストールは `apt` を用いて行います:

```
apt-get install gravwell-kafka
```

シェルインストーラは Arch、Redhat、Gentoo、Fedora を含む SystemD を使用する Debian 以外のシステムをサポートしています。

```
root@gravserver ~ # bash gravwell_kafka_installer.sh
```

### 設定

Gravwell Kafkaインジェスターは、複数のトピック、さらには複数のKafkaクラスタをサブスクライブできます。 各コンシューマは、いくつかのキー設定値を持つコンシューマブロックを定義します。


| パラメーター | タイプ | 説明 | 必須 |
|-----------|------|--------------| -------- |
| Tag-Name  | string | データが送信されるべきGravwellタグ。  | YES |
| Leader    | host:port | Kafka クラスタのリーダー/ブローカー。 ポートが指定されていない場合は、デフォルトのポート 9092 が追加されます。 | YES |
| Topic     | string | コンシューマーが読むであろうKafka topic | YES |
| Consumer-Group | string | このインジェスターが属している Kafka コンシューマーグループ | NO - デフォルトは `gravwell` です。 |
| Source-Override | IPv4 or IPv6 | すべてのエントリの SRC として使用する IP アドレス | NO |
| Rebalance-Strategy | string | Kafkaを読むときに使うリバランスストラテジー | NO - デフォルトは `roundrobin` です。 sticky`, `range` もオプションです。 |
| Key-As-Source | boolean | Gravwell プロデューサはデータソースアドレスをメッセージキーに入れることが多く、設定されている場合、インジェスターはメッセージキーをソースアドレスとして解釈しようとします。 キー構造が正しくない場合、インジェスターはオーバーライド(設定されている場合)またはデフォルトのソースを適用します。 | NO - デフォルトはfalse。 |
| Synchronous | boolean | インジェスターは、kafka バッチが書き込まれるたびに、インジェスト接続上で同期を行います。 | NO - デフォルトはfalse。 |
| Batch-Size | integer | インジェスト接続への書き込みを強制する前に Kafka から読み込むエントリ数。 | NO - デフォルトは512。 |

警告: コンシューマを同期に設定すると、そのコンシューマはインジェスト・パイプラインを継続的に同期させます。 これは、すべてのコンシューマに重大なパフォーマンスの影響を与えます。

注: `Synchronous=true`を使用しているときに大きな`Batch-Size`を設定すると、高負荷時のパフォーマンスを向上させることができます。

#### 設定例

ここでは、2つの異なるコンシューマーグループを使用して2つの異なるトピックをサブスクライブしている構成例を示します。

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

## collectdインジェスター

collectdインジェスターは完全にスタンドアロンの[collectd](https://collectd.org/)コレクションエージェントで、Collectdサンプルを直接Gravwellで取得できます。インジェスターは複数のコレクターをサポートしており、異なるタグ、セキュリティコントロール、およびプラグインからタグへのオーバーライドで設定できます。

### Collectorの例

```
[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	Security-Level=encrypt
	User=user
	Password=secret
	Encoder=json

[Collector "example"]
	Bind-String=10.0.0.1:9999 #default is "0.0.0.0:25826
	Tag-Name=collectdext
	Tag-Plugin-Override=cpu:collectdcpu
	Tag-Plugin-Override=swap:collectdswap
```

### インストール
Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-collectd
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラをダウンロードします。Gravwell サーバー上のターミナルを使用して、スーパーユーザーとして(例: `sudo` コマンドで)以下のコマンドを実行し、インジェスターをインストールします:

```
root@gravserver ~ # bash gravwell_collectd_installer.sh
```

Gravwellのサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメータを抽出し、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラは認証トークンとGravwellインデクサーのIPアドレスを要求します。これらの値はインストール時に設定するか、あるいは空欄にして、`/opt/gravwell/etc/collectd.conf`の設定ファイルを手動で修正することができます。

### 設定

collectdインジェスターは、他のすべてのインジェスターと同じグローバル設定システムに依存します。グローバルセクションはインデクサ接続、認証、ローカルキャッシュ制御の定義に使用されます。

コレクター設定ブロックは、Collectd サンプルを受け入れることができるリッスンコレクターを定義するために使用されます。各コレクター構成は、固有のセキュリティレベル、認証、タグ、ソースオーバーライド、ネットワークバインド、およびタグオーバーライドを持つことができます。複数のコレクター設定を使用することで、1つのコレクトードインジェスターは複数のインターフェイスをリッスンし、複数のネットワークエンクレーブからのCollectdサンプルに固有のタグを適用できます。

デフォルトでは、Collectdインジェスターは _/opt/gravwell/etc/collectd.conf_ にある設定ファイルを読み込みます。

#### 設定例

```
[Global]
	Ingest-Secret = SuperSecretKey
	Connection-Timeout = 0
	Cleartext-Backend-target=192.168.122.100:4023
	Log-Level=INFO

[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	User=user
	Password=secret

[Collector "localhost"]
	Bind-String=[fe80::1]:25827
	Tag-Name=collectdlocal
	Security-Level=none
	Source-Override=[fe80::beef:1000]
	Tag-Plugin-Override=cpu:collectdcpu
```

#### コレクター設定オプション

各 Collector ブロックには、一意の名前と重複しないバインド文字列を含める必要があります。同じポート上の同じインターフェイスにバインドされた複数のコレクターを持つことはできません。

##### Bind-String

Bind-String は、コレクターが受信するコレクトサンプルをリッスンするために使用するアドレスとポートを制御します。有効なバインド文字列には、IPv4 または IPv6 アドレスとポートのいずれかを含める必要があります。すべてのインターフェースをリッスンするには、ワイルドカードアドレス「0.0.0.0.0」を使用します。

###### Bind-Stringの例
```
Bind-String=0.0.0.0:25826
Bind-String=127.0.0.1:25826
Bind-String=127.0.0.1:12345
Bind-String=[fe80::1]:25826
```

##### Tag-Name

Tag-Name は、Tag-Plugin-Override が適用されない限り、Collectd サンプルが割り当てられるタグを定義します。

##### Source-Override

Source-Override ディレクティブは、エントリが Gravwell に送信されるときに適用されるソース値を上書きするために使用されます。デフォルトではインジェスターはインジェスターのソースを適用しますが、検索時にセグメンテーションやフィルタリングを適用するために、特定のソース値をコレクターブロックに適用することが望ましい場合があります。ソースオーバーライドは、任意の有効なIPv4またはIPv6アドレスです。

##### Source-Overrideの例
```
Source-Override=192.168.1.1
Source-Override=[DEAD::BEEF]
Source-Override=[fe80::1:1]
```

##### Security-Level

Security-Level ディレクティブは、コレクターが collectd パケットをどのように認証するかを制御します。 利用可能なオプションは encrypt、sign、none です。デフォルトでは、コレクターは "encrypt" セキュリティレベルを使用し、ユーザとパスワードの両方を指定する必要があります。"none" を使用した場合、ユーザとパスワードは必要ありません。

##### Security-Levelの例
```
Security-Level=none
Security-Level=encrypt
Security-Level = sign
Security-Level = SIGN
```

##### ユーザーとパスワード

セキュリティレベルが "sign "または "encrypt "に設定されている場合、エンドポイントで設定されている値と一致するユーザ名とパスワードを指定する必要があります。デフォルト値は "user "と "secret "で、collectdに同梱されているデフォルト値と一致しています。collectd のデータにセンシティブな情報が含まれている可能性がある場合は、これらの値を変更してください。

###### ユーザーとパスワードの例
```
User=username
Password=password
User = "username with spaces in it"
Password = "Password with spaces and other characters @$@#@()*$#W)("
```

##### エンコーダ

デフォルトのcollectd encoderはJSONですが、シンプルなテキストエンコーダも利用できます。オプションは "JSON" または "text" です。

JSON エンコーダを使用したエントリの例:

```
{"host":"build","plugin":"memory","type":"memory","type_instance":"used","value":727789568,"dsname":"value","time":"2018-07-10T16:37:47.034562831-06:00","interval":10000000000}
```

### Tag PluginのOverrides

各 Collector ブロックはN個の Tag-Plugin-Override 宣言をサポートしており、生成したプラグインに基づいてサンプルに一意のタグを適用するために使用します。タグプラグインオーバーライドは、異なるプラグインからのデータを異なるウェルに保存し、異なるエイジアウトルールを適用したい場合に便利です。例えば、ディスク使用量に関するコレクトレコードを9ヶ月間保存しておくことは価値があるかもしれませんが、CPU使用量のレコードは14日で期限切れになってしまうことがあります。Tag-Plugin-Overrideシステムはこれを簡単にします。

Tag-Plugin-Overrideのフォーマットは、":"で区切られた2つの文字列で構成されています。左側の文字列はプラグインの名前を表し、右側の文字列は希望するタグの名前を表します。タグに関するすべての通常のルールが適用されます。単一のプラグインを複数のタグにマッピングすることはできませんが、複数のプラグインを同じタグにマッピングすることはできます。

#### Tag Plugin Overridesの例
```
Tag-Plugin-Override=cpu:collectdcpu # Map CPU plugin data to the "collectdcpu" tag.
Tag-Plugin-Override=memory:memstats # Map the memory plugin data to the "memstats" tag.
Tag-Plugin-Override= df : diskdata  # Map the df plugin data to the "diskdata" tag.
Tag-Plugin-Override = disk : diskdata  # Map the disk plugin data to the "diskdata" tag.
```

## Kinesis インジェスター

Gravwellは、Amazonの[Kinesisデータストリーム](https://aws.amazon.com/kinesis/data-streams/)サービスからエントリを取得できるインジェスターを提供しています。インジェスターは一度に複数の Kinesis ストリームを処理することができ、各ストリームは多数の個別のシャードで構成されています。Kinesis ストリームを設定するプロセスはこのドキュメントの範囲外ですが、既存のストリームに対して Kinesisインジェスターを設定するには、次のものが必要です:

* AWSのアクセスキー(ID番号と秘密鍵)
* ストリームが存在する地域
* ストリームの名前

ストリームが設定されると、Kinesis ストリームの各レコードは、Gravwell の 1 つのエントリとして保存されます。

### Kinesisストリームの例

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

### インストールと設定

まず、[ダウンロード](#!Quickstart/downloads.md)からインストーラーをダウンロードし、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_kinesis_ingest_installer.sh
```

Gravwellのサービスが同一マシン上に存在する場合は、インストールスクリプトが自動的に`Ingest-Auth`パラメータを抽出して適切に設定してくれるはずです。ここで、`/opt/gravwell/etc/kinesis_ingest.conf`という設定ファイルを開き、Kinesisストリーム用に設定を追加する必要があります。以下のように設定を変更したら、コマンド `systemctl start gravwell_kinesis_ingest.service` でサービスを開始します。

以下の例は、ローカルマシン上のインデクサに接続し（`Pipe-Backend-target` の設定に注意）、US-West-1 リージョンにある "MyKinesisStreamName" という名前の Kinesis ストリームからインデクサに供給する設定のサンプルです。

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

オプション `State-Store-Location` に注意してください。これは、既にインジェストされたエントリの再インジェストを防ぐために、ストリーム中のインジェスターの位置を追跡するステートファイルの場所を設定します。

インジェスターを開始する前に、少なくとも以下のフィールドを設定する必要があります:

* `AWS-Access-Key-ID` - これは利用したいAWSのアクセスキーのIDです。
* `AWS-Secret-Access-Key` - これは秘密のアクセスキーです。
* `Region` - kinesisストリームが存在する領域
* `Stream-Name` - kinesisストリームの名前

複数の異なるKinesisストリームをサポートするために、複数の `KinesisStream` セクションを設定できます。

設定をテストするには、`/opt/gravwell/bin/gravwell_kinesis_ingester -v` を手で実行してください。

ほとんどのフィールドは自明ですが、`Iterator-Type`の設定には注意が必要です。この設定では、インジェスターがデータの読み込みを開始する場所を選択します。デフォルトは "LATEST" で、インジェスターは既存のレコードをすべて無視し、インジェスターの開始後に作成されたレコードのみを読み込みます。これを TRIM_HORIZON に設定すると、インジェスターは利用可能な最も古いレコードからレコードの読み込みを開始します。ほとんどの状況では、古いデータを取得できるように TRIM_HORIZON に設定することをお勧めします。

Kinesisインジェスターは、他の多くのインジェスターに見られる `Ignore-Timestamps` オプションを提供していません。Kinesisメッセージには到着タイムスタンプが含まれます。デフォルトでは、インジェスターはそれをGravwellタイムスタンプとして使用します。データ消費者定義で `Parse-Time=true` が指定されている場合、インジェスターは代わりにメッセージ本文からタイムスタンプを抽出しようとします。

## GCP PubSubインジェスター

Gravwellは、Google Compute Platformの[PubSubストリーム](https://cloud.google.com/pubsub/)サービスからエントリを取得できるインジェスターを提供します。このインジェスターは、1つのGCPプロジェクト内で複数のPubSubストリームを処理できます。PubSubストリームを設定するプロセスはこのドキュメントの範囲外ですが、既存のストリームのためにPubSubインジェスターを設定するには、以下のものが必要です:

* GoogleプロジェクトID
* GCPサービスアカウントの資格情報を含むファイル([サービスアカウントの作成](https://cloud.google.comauthentication/getting-started))のドキュメントを参照してください。
* PubSubトピックの名前

ストリームを構成すると、PubSub ストリームトピックの各レコードは、Gravwell の 1 つのエントリとして保存されます。

### PubSubの例

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

### インストールと設定

まず、[ダウンロード](#!Quickstart/downloads.md)からインストーラーをダウンロードし、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_pubsub_ingest_installer.sh
```

Gravwell サービスが同じマシン上に存在する場合、インストールスクリプトは自動的に `Ingest-Auth` パラメータを抽出して設定し、適切に設定する必要があります。ここで、`/opt/gravwell/etc/pubsub_ingest.conf` 設定ファイルを開き、PubSub トピック用に設定する必要があります。以下のように設定を変更したら、`systemctl start gravwell_pubsubub_ingest.service` コマンドでサービスを起動します。

以下の例は、ローカルマシン上のインデクサに接続し(`Pipe-Backend-target`の設定に注意)、「myproject-127400」GCPプロジェクトの一部である「mytopic」という名前の単一のPubSubトピックからインデクサをフィードする設定のサンプルを示しています。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR

# The GCP project ID to use
Project-ID="myproject-127400"
Google-Credentials-Path=/opt/gravwell/etc/google-compute-credentials.json

[PubSub "gravwell"]
	Topic-Name=mytopic	# the pubsub topic you want to ingest
	Tag-Name=gcp
	Parse-Time=false
	Assume-Localtime=true
```

以下の必須項目に注意してください:

* `Project-ID` - GCPプロジェクトのプロジェクトID文字列
* `Google-Credentials-Path` - JSON形式のGCPサービスアカウントの資格情報を含むファイルへのパス
* `Topic-Name` - 指定されたGCPプロジェクト内のPubSubトピックの名前

1つのGCPプロジェクト内で複数の異なるPubSubトピックをサポートするために、複数の `PubSub` セクションを設定できます。

エラーが表示されなければ、おそらく設定は問題ありません。

PubSub インジェスターは、他の多くのインジェスターに見られる `Ignore-Timestamps` オプションを提供していません。PubSubメッセージには到着タイムスタンプが含まれています。デフォルトでは、インジェスターはそれをGravwellタイムスタンプとして使用します。データ消費者定義で `Parse-Time=true` が指定されている場合、インジェスターは代わりにメッセージ本文からタイムスタンプを抽出しようとします。

## Office 365ログ インジェスター

Gravwellは、Microsoft Office 365ログ用のインジェスターを提供します。インジェスターは、サポートされているすべてのログタイプを処理できます。インジェスターを設定するには、Azure Active Directory管理ポータル内で新しい*アプリケーション*を登録する必要があります。以下の情報が必要です:

* Client ID: Azure管理コンソールを介してアプリケーション用に生成されたUUID
* Client secret: Azureコンソールを介してアプリケーション用に生成されたシークレットトークン
* Azure Directory ID: Active Directoryインスタンスを表すUUIDで、Azure Active Directoryダッシュボードに表示されます。
* Tenant Domain: Office 365ドメインのドメイン（例："mycorp.onmicrosoft.com"）

### ContentTypeの例

```
[ContentType "azureAD"]
	Content-Type="Audit.AzureActiveDirectory"
	Tag-Name="365-azure"

[ContentType "exchange"]
	Content-Type="Audit.Exchange"
	Tag-Name="365-exchange"

[ContentType "sharepoint"]
	Content-Type="Audit.SharePoint"
	Tag-Name="365-sharepoint"

[ContentType "general"]
	Content-Type="Audit.General"
	Tag-Name="365-general"

[ContentType "dlp"]
	Content-Type="DLP.All"
	Tag-Name="365-dlp"
```

### インストールと設定

まず、[ダウンロード](#!Quickstart/downloads.md)からインストーラーをダウンロードし、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_o365_installer.sh
```

Gravwellサービスが同じマシン上に存在する場合、インストールスクリプトは自動的に `Ingest-Auth` パラメータを抽出して設定し、適切に設定する必要があります。次に、`/opt/gravwell/etc/o365_ingest.conf`構成ファイルを開き、必要に応じてプレースホルダフィールドを置き換えたり、タグを修正したりして、Office 365アカウント用に設定する必要があります。以下のように設定を変更したら、`systemctl start gravwell_o365_ingest.service` コマンドでサービスを起動します。

以下の例では、ローカルマシン上のインデクサに接続して (`Pipe-Backend-target` の設定に注意してください)、サポートされているすべてのログタイプのログをフィードする設定のサンプルを示しています:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR
State-Store-Location=/opt/gravwell/etc/o365_ingest.state

Client-ID=79fb8690-109f-11ea-a253-2b12a0d35073
Client-Secret="<secret>"
Directory-ID=e8b7895e-109f-11ea-9dcc-93fb14b5dab5
Tenant-Domain=mycorp.onmicrosoft.com

[ContentType "azureAD"]
	Content-Type="Audit.AzureActiveDirectory"
	Tag-Name="365-azure"

[ContentType "exchange"]
	Content-Type="Audit.Exchange"
	Tag-Name="365-exchange"

[ContentType "sharepoint"]
	Content-Type="Audit.SharePoint"
	Tag-Name="365-sharepoint"

[ContentType "general"]
	Content-Type="Audit.General"
	Tag-Name="365-general"

[ContentType "dlp"]
	Content-Type="DLP.All"
	Tag-Name="365-dlp"
```

## Microsoft Graph API インジェスター

Gravwellは、MicrosoftのGraph APIからセキュリティ情報を引き出すことができるインジェスターを提供します。インジェスターを設定するには、Azure Active Directory管理ポータル内で新しい*application*を登録する必要があります。以下の情報が必要です:

* クライアントID: Azure管理コンソールを介してアプリケーション用に生成されたUUID
* クライアントシークレット: Azureコンソールを介してアプリケーションのために生成されたシークレットトークン
* テナントドメイン: 例えば"mycorp.onmicrosoft.com"といった、あなたのAzureドメインのドメイン

### ContentTypeの例

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

### インストールと設定

まず、[ダウンロード](#!Quickstart/downloads.md)からインストーラーをダウンロードし、インジェスターをインストールします:

```
root@gravserver ~# bash gravwell_msgraph_installer.sh
```

Gravwellサービスが同じマシン上に存在する場合、インストールスクリプトは自動的に `Ingest-Auth` パラメータを抽出して、適切に設定します。ここで、`/opt/gravwell/etc/msgraph_ingest.conf`設定ファイルを開き、プレースホルダフィールドを置き換えたり、必要に応じてタグを修正したりして、アプリケーション用に設定する必要があります。以下のように設定を変更したら、`systemctl start gravwell_msgraph_ingest.service`コマンドでサービスを起動します。

デフォルトでは、インジェスターはセキュリティアラートが到着するたびにインジェストします。また、定期的に新しいセキュリティスコアの結果 (通常は毎日発行されます) をクエリし、それらのセキュリティスコアの結果を構築するために使用される関連する制御プロファイルをインジェストします。これら3つのデータソースは、デフォルトでそれぞれ `graph-alerts`, `graph-scores`, `graph-profiles` というタグにインジェストされます。

以下の例は、ローカルマシン上のインデクサに接続し（`Pipe-Backend-target` の設定に注意してください）、サポートされているすべてのタイプのログをフィードする設定のサンプルを示しています:

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


## ディスクモニター

ディスクモニターインジェスターは、ディスクアクティビティの定期的なサンプルを採取し、そのサンプルを gravwell に出荷するように設計されています。 ディスクモニタは、ストレージのレイテンシの問題、迫り来るディスク障害、およびその他の潜在的なストレージの問題を特定するのに非常に有用です。 Gravwell では、クエリーがどのように動作しているかを研究し、ストレージインフラストラクチャの動作が悪 い場合を特定するために、ディスクモニタを使用して積極的にストレージインフラストラクチャを監視しています。 RAID コントローラが診断ログで言及しなかった場合でも、レイテンシプロットを介して書き込みスルーモードに移行した RAID アレイを特定することができました。

ディスクモニターインジェスターは [github](https://github.com/gravwell/ingesters) で公開されています。

![diskmonitor](diskmonitor.png)

## セッションインジェスター

セッションインジェスターは、より大きな単一のレコードをインジェストするために使用される特殊なツールです。インジェスターは指定されたポートをリッスンし、クライアントからの接続を受信すると、受信したすべてのデータを単一のエントリに集約します。

これにより、すべてのWindows実行ファイルをインデックス化するような動作が可能になります:

```
for i in `ls /path/to/windows/exes`; do cat $i | nc 192.168.1.1 7777 ; done
```

セッションインジェスターは、永続的な設定ファイルではなく、コマンドラインパラメータを介して駆動されます。

```
Usage of ./session:
  -bind string
        Bind string specifying optional IP and port to listen on (default "0.0.0.0:7777")
  -clear-conns string
        comma seperated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -max-session-mb int
        Maximum MBs a single session will accept (default 8)
  -pipe-conns string
        comma seperated list of paths for named pie connection
  -tag-name string
        Tag name for ingested data (default "default")
  -timeout int
        Connection timeout in seconds (default 1)
  -tls-conns string
        comma seperated server:port list of TLS connections
  -tls-private-key string
        Path to TLS private key
  -tls-public-key string
        Path to TLS public key
  -tls-remote-verify string
        Path to remote public key to verify against
```

### 注意事項

セッションインジェスターは正式にはサポートされておらず、インストーラーもありません。 セッションインジェスターのソースコードは [github](https://github.com/gravwell/ingesters) にあります。

## Amazon SQS インジェスター

Amazon SQS インジェスター (sqsIngester)は、インジェスト用の標準SQSキューとFIFO SQSキューの両方をサブスクライブできるシンプルなインジェスターです。Amazon SQSは、メッセージの配信保証、メッセージの「ソフト」順序付け、メッセージの「アットリーストワンス」配信をサポートする大容量メッセージキューサービスです。

Gravwell の場合、「アットリーストワンス」配信は重要な注意点です - SQS インジェスターは、同一のタイムスタンプを持つ重複したメッセージを受信する可能性があります (設定によっては)。また、他の接続サービスとの SQS ワークフローの展開方法によっては、SQS インジェスターが一部のメッセージを見ない可能性もあります。詳細については、[Amazon SQS](https://aws.amazon.com/sqs/)を参照してください。

### キューの例

```
[Queue "default"]
	Region="us-east-2"
	Queue-URL="https://us-east-2.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
	Assume-Local-Timezone=false #Default for assume localtime is false
	Source-Override="DEAD::BEEF" #override the source for just this Queue 

[Queue "default"]
	Region="us-west-1"
	Queue-URL="https://us-west-1.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
```

### インストール
Gravwell の Debian リポジトリを使用している場合、インストールは apt コマンドひとつで済みます:

```
apt-get install gravwell-sqs
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md) からインストーラーをダウンロードします。Netflow インジェスターをインストー ルするには、単に root としてインストーラーを実行します(実際のファイル名には通常バージョン番号が含まれます):

```
root@gravserver ~ # bash gravwell_sqs.sh
```

ローカルマシン上に Gravwell インデクサーがない場合、インストーラーは Ingest-Secret 値とインデクサー(またはフェデレーター)の IP アドレスを要求します。そうでなければ、既存の Gravwell 設定から適切な値を読み取ります。いずれにせよ、インストール後に `/opt/gravwell/etc/sqs.conf` の設定ファイルを確認してください。典型的な設定は以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Pipe-Backend-Target=/opt/gravwell/comms/pipe 
Log-Level=INFO
Log-File=/opt/gravwell/log/sqs.log

# A Queue pulls from a specific SQS queue with a given AKID and Secret. See
# https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
# for information about obtaining an AKID/Secret for your user.
[Queue "default"]
	Region="us-east-2"
	Queue-URL="https://us-east-2.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
```

この設定では、`/opt/gravwell/comms/pipe`を経由してローカルのインデクサにエントリを送ることに注意してください。エントリには 'sqs' というタグが付けられます。

SQS キューごとに1つずつ、任意の数の `Queue` エントリを設定することができ、それぞれに固有の認証やタグ名などを指定できます。

## パケットフリートインジェスター

パケットフリートインジェスターは、Google Stenographerインスタンスにクエリを発行し、その結果をパケットごとにGravwellにインジェストする仕組みを提供します。

各Stenographerインジェスターは指定されたポート（```Listen-Address```）をリッスンし、HTTP POSTとしてStenographerのクエリ（下記のクエリ構文を参照）を受け取ります。クエリを受け取ると、インジェスターは整数のジョブIDを返し、非同期的にStenographerインスタンスにクエリを実行し、返されたPCAPのインジェストを開始します。複数のインフライトクエリを同時に実行できます。ジョブのステータスは、"/status "でHTTP GETを発行することで確認することができ、JSONでエンコードされたインフライトジョブIDの配列を返します。

指定されたインジェスター・ポートを参照することで、ジョブ・ステータスを送信して表示するためのシンプルなウェブ・インターフェースも利用できます。

### Stenographerの例

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	Assume-Local-Timezone=false #Default for assume localtime is false
	Source-Override="DEAD::BEEF" #override the source for just this Queue 

[Stenographer "Region 2"]
	URL="https://my.url:1234"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
```

### 設定オプション

パケットフリートでは、いくつかのグローバル設定とStenographerごとの設定オプションが必要です。グローバル設定には、以下に示すように、TLSの設定（該当する場合）とWebインターフェイスのリスナーアドレスが含まれます:

```
Use-TLS=true
Listen-Address=":9002"
Server-Cert="server.cert"
Server-Key="server.key"
```

各Stenographerインスタンスには、以下のスタンザが必要です。ここでの例の名前 `Region 1` は、ウェブインタフェースがステノグラファーのインスタンスを一覧表示するために使用します。

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	#Assume-Local-Timezone=false #Default for assume localtime is false
	#Source-Override="DEAD::BEEF" #override the source for just this Queue 
```

### クエリ言語 ###

ユーザーは、非常にシンプルなクエリ言語でパケットを指定することで、Stenographerにパケットを要求します。 この言語はBPFのシンプルなサブセットであり、プリミティブです:



    host 8.8.8.8          # Single IP address (hostnames not allowed)
    net 1.0.0.0/8         # Network with CIDR
    net 1.0.0.0 mask 255.255.255.0  # Network with mask
    port 80               # Port number (UDP or TCP)
    ip proto 6            # IP protocol number 6
    icmp                  # equivalent to 'ip proto 1'
    tcp                   # equivalent to 'ip proto 6'
    udp                   # equivalent to 'ip proto 17'

    # Stenographer-specific time additions:
    before 2012-11-03T11:05:00Z      # Packets before a specific time (UTC)
    after 2012-11-03T11:05:00-07:00  # Packets after a specific time (with TZ)
    before 45m ago        # Packets before a relative time
    before 3h ago         # Packets after a relative time

**注意** : 相対時間は、上で示したように、時間または分の整数値で測定する必要があります。


プリミティブは、/&&& と or/|| を使用して組み合わせることができます。これらは同じ優先順位を持ち、左から右へ評価します。丸括弧"()"はグループ化にも使えます。


    (udp and port 514) or (tcp and port 8080)

**注** : 本項は[Google Stenographer](https://github.com/google/stenographer/blob/master/README.md)からの出典です。

## IPMIインジェスター

IPMIインジェスターは、任意の数のIPMIデバイスからセンサーデータレコード（SDR）とシステムイベントログ（SEL）のレコードを収集します。

設定ファイルには、各IPMIデバイスに接続するためのシンプルなホスト/ポート、ユーザー名、パスワードのフィールドが用意されています。SELとSDRのレコードは、JSONにエンコードされたスキーマで取り込まれます。例えば、以下のようになります:

```
{
    "Type": "SDR",
    "Target": "10.10.10.10:623",
    "Data": {
        "+3.3VSB": {
            "Type": "Voltage",
            "Reading": "3.26",
            "Units": "Volts",
            "Status": "ok"
        },
        "+5VSB": {...},
        "12V": {...}
    }
}

{
    "Target": "10.10.10.10:623",
    "Type": "SEL",
    "Data": {
        "RecordID": 25,
        "RecordType": 2,
        "Timestamp": {
            "Value": 1506550240
        },
        "GeneratorID": 32,
        "EvMRev": 4,
        "SensorType": 5,
        "SensorNumber": 81,
        "EventType": 111,
        "EventDir": 0,
        "EventData1": 240,
        "EventData2": 255,
        "EventData3": 255
    }
}
```

### 設定オプション ###

IPMIでは、グローバル設定オプションのデフォルトセットを使用します。IPMIデバイスは「IPMI」スタンザで構成され、各スタンザは同じ認証情報を共有する複数のIPMIデバイスをサポートできます。例えば、以下のようになります:

```
[IPMI "Server 1"]
	Target="127.0.0.1:623"
	Target="1.2.3.4:623"
	Username="user"
	Password="pass"
	Tag-Name=ipmi
	Rate=60
	Source-Override="DEAD::BEEF" 
```

IPMIスタンザはシンプルで、1つまたは複数のターゲット（IPMIデバイスのIP:PORT）、ユーザー名、パスワード、タグ、ポールレート（秒）を指定するだけです。デフォルトのポーリングレートは60秒です。オプションで、ソースオーバーライドを設定して、取り込まれたすべてのエントリーのSRCフィールドを別のIPに強制的に変更することができます。デフォルトでは、SRCフィールドはIPMIデバイスのIPに設定されています。

さらに、すべてのIPMIスタンザは、[こちら](https://docs.gravwell.io/#!ingesters/preprocessors/preprocessors.md)に記載されているように、「Preprocessor」オプションを使用することができます。

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

## インジェストAPI

GravwellインジェストAPIとコアインジェスタは、BSD 2-Clauseライセンスのもとで完全にオープンソースとなっています。これは、独自のインジェスタを書き、Gravwellエントリー生成を独自の製品やサービスに統合できることを意味します。コアインジェストAPIはGoで書かれていますが、利用可能なAPI言語のリストは積極的に拡張中です。

[API code](https://github.com/gravwell/ingest)

[API documentation](https://godoc.org/github.com/gravwell/ingest)

ファイルを監視し、ファイルに書き込まれたすべての行をGravwellクラスタに送信する、非常に基本的なインジェスターの例（100行以下のコード）は、[ここ](https://www.godoc.org/github.com/gravwell/ingest#example-package)で見ることができます。

チームは継続的にインジェストAPIを改善し、追加言語への移植を行っているので、Gravwell Githubページをチェックし続けてください。コミュニティ開発は完全にサポートされていますので、マージリクエスト、言語移植、またはオープンソースの素晴らしい新しいインジェスターがあれば、Gravwellに知らせてください!  Gravwellチームは、あなたの頑張りをインジェスターハイライトシリーズで取り上げたいと思っています。
