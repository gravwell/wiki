# Simple Relay

Simple Relayは、IPv4またはIPv6での平文TCP、暗号化TCPまたは平文UDPネットワーク接続を介して配信できる、テキストベースのデータソースのためのインジェスターです。

Simple Relayの一般的な使用例は次のとおりです:

* リモートsyslog収集
* ネットワーク経由のDevopログ収集
* Broセンサーのログ収集
* ネットワーク経由で配信可能なテキストソースとの簡単な統合

## 基本設定

Simple Relayは、インジェスター設定で説明されている統一されたグローバル構成ブロックを使用します。 他のほとんどのGravwellインジェスターと同様に、Simple Relayは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。  

Simple Relayインジェスターの設定例として、複数のポートをリッスンし、それぞれに固有のタグを適用するように設定した場合は、以下のようになります:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Cleartext-Backend-target=127.0.0.1:4023 #example of a cleartext connection
Cleartext-Backend-target=127.1.0.1:4023 #example of second cleartext connection
Encrypted-Backend-target=127.1.1.1:4024 #example of encrypted connection
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection
Ingest-Cache-Path=/opt/gravwell/cache/simple_relay.cache #local storage cache when uplinks fail
Max-Ingest-Cache=1024 #Number of MB to store, localcache will only store 1GB before stopping
Log-Level=INFO
Log-File=/opt/gravwell/log/simple_relay.log

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Listener "default"]
	Bind-String="0.0.0.0:7777" #bind to all interfaces, with TCP implied
	#Lack of "Reader-Type" implies line break delimited logs
	#Lack of "Tag-Name" implies the "default" tag
	#Assume-Local-Timezone=false #Default for assume localtime is false
	#Source-Override="DEAD::BEEF" #override the source for just this listener

[Listener "syslogtcp"]
	Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
	Reader-Type=rfc5424
	Tag-Name=syslog
	Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
	Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry

[Listener "syslogudp"]
	Bind-String="udp://0.0.0.0:514" #standard UDP based RFC5424 syslog
	Reader-Type=rfc5424
	Tag-Name=syslog
	Timezone-Override="America/Chicago"
	Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry
```

注：[syslog検索モジュール](#!search/syslog/syslog.md)を使ってsyslogエントリを解析する場合は、`Keep-Priority`フィールドが必要です。

## リスナー

各リスナーには、リスナーがラインリーダー、RFC5424リーダー、JSONリスナーのいずれであるかに関わらず、普遍的な設定値のセットが含まれています。

## ユニバーサルリスナーの設定パラメーター

リスナーはいくつかの設定パラメータをサポートしており、プロトコル、リスニングインターフェース、ポートの指定や、インジェストの動作の微調整が可能です。

### Bind-String

Bind-String パラメータは、リスナーがどのインターフェイスとポートにバインドするかを制御します。リスナーは、平文/暗号化TCPポートまたは平文UDPポート、特定のアドレス、特定のポートにバインドできます。 IPv4 と IPv6 がサポートされています。

```
#bind to all interfaces on TCP port 7777
Bind-String=0.0.0.0:7777

#bind to all interfaces on UDP port 514
Bind-String=udp://0.0.0.0:514

#bind to port 1234 on link local IPv6 address on interface p1p
Bind-String=[fe80::4ecc:6aff:fef9:48a3%p1p1]:1234

#bind to IPv6 globally routable address on TCP port 901
Bind-String=[2600:1f18:63ef:e802:355f:aede:dbba:2c03]:901

#listen for TLS connections on port 9999
Bind-String=tls://0.0.0.0:9999
```

### Cert-File

TLSの`Bind-String`を使う場合は、`Cert-File`も指定する必要があります。この値は、有効なTLS証明書を含むファイルへのパスでなければなりません:

```
Cert-File=/opt/gravwell/etc/cert.pem
``` 

### Key-File

TLSの`Bind-String`を使用する場合は、`Key-File`も指定する必要があります。この値は、有効なTLSキーを含むファイルへのパスでなければなりません:

```
Key-File=/opt/gravwell/etc/key.pem
``` 

### リスナーリーダーの種類と構成

シンプルリレーでは、さまざまな場面で活躍する以下の種類のベーシックリーダーに対応しています。

* ラインリーダー
* RFC5424

基本的なリスナーは、"Listener "ブロックで指定され、行区切りやRFC5424のリーダータイプをサポートします。  各リスナーは、一意の名前と一意のバインドポートを持つ必要があります（2つの異なるリスナーが同じプロトコル、アドレス、ポートにバインドすることはできません）。  ベーシック・リスナーのリーダー・タイプは、"Reader-Type"パラメーターで制御されます。  現在、2種類のリスナー(ラインとRFC5424）があります。Reader-Typeが指定されていない場合、ラインリーダータイプが想定されます。

基本リスナーでは、各リスナーが受信データに適用するタグを"Tag-Name"パラメータで指定する必要があります。 "Tag-Name"パラメータが省略された場合、"default"タグが適用されます。 TCPポート5555で改行されたデータを受信し、"testing"というタグを適用する "test"という名前の最も基本的なリスナーは、以下のような構成仕様になっています:

```
[Listener "test"]
	Bind-String=0.0.0.0:5555
	Tag-Name=testing
```

## ラインリーダーリスナー

ラインリーダーリスナーは、TCPまたはUDPストリームから改行されたデータストリームを読み取るように設計されています。  ネットワークを介して単純な改行済データを配信できるアプリケーションは、このタイプのリーダーを利用して非常に簡単かつ簡単にSolitonNKと統合できます。ラインリーダーリスナーは、ログファイルをリスニングポートに送信するだけで、単純なログファイルの配信にも使用できます。

例えば、netcatとSimple Relayを使って既存のログファイルをSolitonNKに取り込むことができます。
```
nc -q 1 10.0.0.1 7777 < /var/log/syslog
```

### ラインリーダーリスナーの例

最も基本的なリスナーが必要とするのは、リスナーがどのポートをリッスンするかを示す "Bind-String "という引数だけです。 

```
[Listener "default"]
	Bind-String="0.0.0.0:7777" #bind to all interfaces, with TCP implied
```

## RFC5424 リスナー

RFC5424またはRFC3164に基づいて構造化されたsyslogメッセージを受け入れるように設計されたリスナーにより、Simple Relayはsyslogのアグリゲーションポイントとして機能します。  ポート601で信頼性の高いTCP接続を使用してsyslogメッセージを受信するリスナーを有効にするには、"Reader-Type"を "RFC5424"に設定します。

```
[Listener "syslog"]
	Bind-String=0.0.0.0:601
	Reader-Type=RFC5424
```

ポート514を介してステートレスUDP上のsyslogメッセージを受け付ける場合、リスナーは以下のようになります:

```
[Listener "syslog"]
	Bind-String=udp://0.0.0.0:514
	Reader-Type=RFC524
```

RFC5424のリーダータイプは、"Keep-Priority"というパラメータもサポートしており、デフォルトではtrueに設定されています。  一般的なsyslogメッセージには優先度の識別子が付加されていますが、ユーザによっては保存されたメッセージから優先度を破棄したい場合もあります。  これは、RFC5424ベースのリスナーに"Keep-Priority=false"を追加することで実現します。 LINEベースのリスナーは、"Keep-Priority"パラメータを無視します。

優先度を付けたsyslogメッセージの例:

```
<30>Sep 11 17:04:14 router dhcpd[9987]: DHCPREQUEST for 10.10.10.82 from e8:c7:4f:04:e1:af (Chromecast) via insecure
```

エントリーからプライオリティタグを削除するリスナー仕様の例:

```
[Listener "syslog"]
	Bind-String=udp://0.0.0.0:514
	Reader-Type=RFC524
	Keep-Priority=false
```

注：syslogメッセージの優先度部分はRFC仕様で成されている。 優先度を削除すると、Gravwellの[syslog](#!search/syslog/syslog.md)検索モジュールが適切に値を解析できなくなります。 Gravwellの有料ライセンスはすべて無制限なので、syslogメッセージに優先度フィールドを残すことをお勧めします。 また、syslog検索モジュールは、正規表現を使って手作業でsyslogメッセージを解析しようとするよりも劇的に高速です。

## JSON リスナー

JSON リスナータイプは、インジェスト時にマイルドなJSON処理を可能にします。 JSONリーダーの目的は、JSONエントリーのフィールドの値に基づいて、エントリーに固有のタグを適用することです。  多くのアプリケーションは、JSONのフォーマットを示すフィールドを持つJSONデータをエクスポートしますが、処理効率の観点から、異なるフォーマットに特定のタグを付けることは有益です。

優れた使用例として、多くのBroセンサ・アプライアンスに見られるJSON over TCPデータ・エクスポート機能があります。 このアプライアンスは、単一のTCPストリーム上ですべてのBroログ・データをエクスポートしますが、ストリーム内には異なるモジュールによって構築された複数のデータ・タイプがあります。 JSONリスナーを使用して、モジュール・フィールドからデータ・タイプを導き出し、一意のタグを適用することができます。これにより、Broの接続ログを1つのデータに、BroのDNSログを別のデータに、その他すべてのBroのログをさらに別のデータに保管するようなことが可能になります。 その結果、異なるタグでデータタイプを区別することができ、複数のJSONデータタイプが1つのストリームを介して入ってくるときにGravwell Wellsを活用することができます。

### JSON リスナーの設定パラメーター

JSONリスナー・ブロックは、上で説明したユニバーサル・リスナー・タイプを実装しています。 追加のパラメータで、タグを定義するためにピボットしたいフィールドを指定できます。

#### 抽出器パラメーター

"Extractor"パラメータは、JSON エントリからフィールドを引き出すために使用される JSON 抽出文字列を指定します。 抽出文字列は、Gravwell [json](#!search/json/json.md)検索モジュールからインラインフィルタリングを除いたものと同じ構文に従います。

次のようなJSONが与えられます:

```
{
  "time": "2018-09-12T12:25:33.503294982-06:00",
  "class": 5.1041415140005e+18,
  "data": "Have I come to Utopia to hear this sort of thing?",
  "identity": {
    "user": "alexanderdavis605",
    "name": "Noah White",
    "email": "alexanderdavis605@test.org",
    "phone": "+52 27 83 68 75069 2"
  },
  "location": {
    "address": "43 Wilson Pkwy,\nBury, AL, 66232",
    "state": "PW",
    "country": "Pakistan"
  },
  "group": "carp",
  "useragent": "Mozilla\/5.0 (X11; Fedora; Linux x86_64) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/52.0.2743.116 Safari\/537.36",
  "ip": "8.83.94.200"
}
```

次のExtractionパラメータを使用して、場所と状態の値を抽出し、どの状態の略語を見つけたかに基づいてタグを適用することができます:

```
Extractor=location.state
```

**Tag-Match**

各JSONListenerは、複数のフィールド値からタグへのマッチ指定をサポートしています。 タグへの値の割り当ては、"Tag-Match"パラメータの引数として、<field value>:<tag name>の形式で指定します。

例えば、"foo"という値を持つフィールドを抽出し、それを"bar"というタグに割り当てたい場合、JSONリスナーの設定ブロックに以下のように追加します。:

```
Tag-Match=foo:bar
```

フィールド抽出値には":"文字を含めることができます。":"文字を含むフィールド値を指定するには、その値を二重引用符で囲みます。

例えば、抽出された値"foo:bar"にタグ"baz"を割り当てたい場合、"Tag-Match"パラメータは以下のようになります:

```
Tag-Match="foo:bar":baz
```

抽出値とタグの対応付けは、多対1にすることができます。つまり、複数の抽出値を同じタグに対応付けることができます。 例えば、以下のパラメータは、"foo"と "bar"の両方の抽出値を "baz"というタグにマッピングします:

```
Tag-Match=foo:baz
Tag-Match=bar:baz
```

ただし、1つの抽出値を複数のタグにマッピングすることはできません。 次のような場合は無効です:

```
Tag-Match=foo:baz
Tag-Match=foo:bar
```

**Default-Tag**

フィールドを抽出してタグを適用する際、マッチするTag-Matchが指定されていない場合、JSON リスナーはデフォルトのタグを適用します。

### JSON リスナーの動作例

以下のように構成されたJSONリスナーを想定します:

```
[JSONListener "testing"]
	Bind-String=0.0.0.0:7777
	Extractor="field1"
	Default-Tag=json
	Tag-Match=test1:tag1
	Tag-Match=test2:tag2
	Tag-Match=test3:tag3
```

JSONデータと結果のタグの例:

#### フィールドが一致

```
{ "field1": "test1", "field2": "test2" }
```

このエントリーは、フィールド "field1"が "Tag-Match=test1:tag1"にマッチしたため、タグ "tag1"を取得します。

#### フィールドが一致しない

```
{ "field1": "foobar", "field2": "test2" }
```

フィールド "field1"がどの "Tag-Match"パラメータにもマッチしなかったため、エントリーには "json"というタグが付けられました。

##### 抽出フィールドが見つからない

```
{ "fieldfoo": "test1", "fieldbar": "test2" }
```

エントリーに"json"というタグが付くのは、抽出器がフィールド"field1"を見つけられなかったからです。
