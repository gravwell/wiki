# IPFIXとNetflow V9

ipfixプロセッサは、IPFIXやNetflow V9の未加工のデータフレームを抽出してフィルタリングするためのもので、ネットワークフローの識別、ポートのフィルタリング、フローの集合体の挙動の監視などを素早く行うことができます。GravwellにはネイティブのIPFIX + Netflowインジェスターがあり、オープンソースで、https://github.com/gravwell/ingesters、または[クイックスタートセクション](/#!quickstart/downloads.md)にあるインストーラーとして利用できます。

## テンプレートについての注意点

Gravwellはエントリーの順番にも柔軟に対応しています。古いエントリーから新しいエントリー、新しいエントリーから古いエントリーの順にクエリを実行することもできますし、まったく別の基準でエントリーをソートすることもできます。

残念ながら、この柔軟性は IPFIX や Netflow V9 とはうまく相互作用しません。これらのプロトコルは、データレコードの形式を記述する *templates* を定義しています。テンプレートは滅多に送信されません。任意の IPFIX メッセージがテンプレートレコードを含むことはほとんどありませんが、以前に定義されたテンプレートを参照しているデータレコードはあります。適切なテンプレートにアクセスできなければ、データレコードは無意味です。これは、IPFIXメッセージを新しいものから古いものへと解析する(あるいは他の方法で検索する)ことが、追加のステップを**踏まない限り**、適切に動作しないことを意味します。

最も簡単な方法は、IPFIX/Netflow V9の受信メッセージを、データ・レコード用の適切なテンプレートを使って、*インジェスター*にリパックさせることだと判断しました。これにより、保存される各メッセージのサイズは大きくなりますが、メッセージには通常数十個のデータレコードが含まれているため、増加分は大きくありません。また、すべてのメッセージ（エントリ）は完全に独立しており、解析のために以前のメッセージに依存していないことを意味しています。

## サポートされているオプション

* `-e <name>`: "-e"オプションは、ipfixモジュールが列挙値を操作することを指定します。 列挙値を操作することは、上流のモジュールを使ってipfixフレームを抽出した場合に便利です。 
* `-s`: "-s"オプションは、ipfixモジュールを厳密モードにします。厳密モードでは、指定した抽出が*必ず*成功しなければ、そのエントリーは削除されます。`ipfix protocolIdentifier==1 srcPort` とすると、すべてのエントリが削除されます。なぜなら、ICMPメッセージ（IPプロトコル1）にはソースポートフィールドが含まれないからです。
* `-f`: "-f"フラグは、IPFIXデータの各レコードから利用可能なフィールドのリストを抽出し、人間が読める形式で "FIELDS "と呼ばれる列挙値にドロップするようモジュールに指示します。これは、フィールド抽出とは相互に排他的です。-fを指定した場合、抽出するフィールドを指定することはできません。

## 処理演算子

各IPFIXフィールドは、高速フィルタとして機能する一連の演算子をサポートしています。 各演算子がサポートするフィルタは、フィールドのデータ型によって決まります。数値は*サブセット演算子を除く*すべてをサポートし、IPアドレスは*サブセット演算子のみ*をサポートします。

| 演算子 | 名前 | 説明
|----------|------|-------------
| == | 等しい | フィールドは等しい
| != | 等しくない | フィールドは等しくない
| < | 未満 | フィールドはその値より小さい
| > | より大きい | フィールドはその値より大きい
| <= | 以下 | フィールドはその値以下
| >= | 以上 | フィールドはその値以上
| ~ | 含む | フィールドはその値を含む
| !~ | 含まない | フィールドはその値を含まない


## データ項目

ipfixの検索モジュールは、生のIPFIXおよびNetflow v9メッセージを処理するように設計されています。 1つのIPFIX（またはNFv9）メッセージは、ヘッダとN個のデータレコードで構成されています。IPFIX/NFv9とNetflow v5の本質的な違いは、Netflow v5ではすべてのフィールドがあらかじめ定義されているのに対し、IPFIX/NFv9のデータレコードは、生成デバイスが指定したテンプレートに準拠し、以前のメッセージで送信されていることです。そのため、あるIPFIX/NFv9ジェネレーターはフローの送信元と送信先のIPとポートを送信し、一方、スイッチはパケット数を含むレコードを送信するだけかもしれません。

IPFIX/NFv9 メッセージ*ヘッダ* のすべての要素は、データレコード自体の要素と同様にフィルタリングに使用できます。ヘッダデータ項目でフィルタリングする場合、フィルタはメッセージの*すべての*レコードに適用されます。 ヘッダデータ項目が最初に処理され、ヘッダフィルタがフレームを落とさない場合にのみ、個々のレコードが処理されます。

ipfixプロセッサは拡張モジュールであり、拡張モジュールは入力エントリを複数の出力エントリに分割します。このモジュールは、IPFIXメッセージ全体に対応する個々のエントリを受け取ります。このモジュールは、IPFIXメッセージ全体に対応する個々のエントリを受け取り、メッセージ内の各データレコードを処理して、*各レコードごとに*新しいエントリを出力します。

注意：抽出するヘッダ項目*のみ*を指定した場合（例：`ipfix Version Sequence`）、IPFIXメッセージには1つのヘッダしか存在しないため、エントリは*展開されません*。各受信エントリは、1つ以上の送信エントリにはなりません。ヘッダ項目とデータ項目の組み合わせを指定すると、展開が行われます。

### IPFIXヘッダーデータ項目

| フィールド |       説明        | サポートされている演算子 | 例 |
|-------|--------------------------|---------------------|---------|
| Version | 使用中のNetflowのバージョン。10 は IPFIX、9 は Netflow v9 を意味する | > < <= >= == != | Version != 0xa
| Length | IPFIX：この IPFIX メッセージの全長。Netflow v9：このメッセージに含まれるレコードの数 | > < <= >= == != | Length > 1000
| Sec | センサーデバイスのUnixタイム | > < <= >= == != | Sec == 1526511023
| Uptime | 検出デバイスがオンラインになったミリ秒数 (Netflow v9 のみ) | > < <= >= == != | Sec == 3516293216
| Sequence | センシングデバイスが送信した総メッセージのシーケンスカウンタ  | > < <= >= == != | Sequence == 1
| Domain | レコードの元となった観測領域 | > < <= >= == != | Domain == 0x1A


### データレコード項目

IPFIX と Netflow v9 では、データレコードを構成するフィールドを定義しています。Netflow では、数十個のフィールドが定義されており、RFC3954 に記載されています（https://tools.ietf.org/html/rfc3954#section-8）。Netflow v9 のフィールドは、フィールド ID（正の整数）、データタイプ（例：16 ビットの符号なし整数など）、名前（例：L4_SRC_PORT）で定義されます。

IANAは、IPFIXについても同様のフィールドを定義しており、その内容は[こちら](https://www.iana.org/assignments/ipfix/ipfix.xhtml#ipfix-information-elements)に記載されています。IPFIXのフィールドは、Netflow v9のフィールドと似ていますが（フィールドID、タイプ、名前を持っています）、IPFIXではNetflow v9で使われているフィールドIDに加えて、 _エンタープライズID_ という概念を導入しています。 エンタープライズID を使うことで、あらかじめ定義されているフィールドに加えて、ユーザーが独自のフィールドを定義することができます。これらのフィールドはすべて enterorise ID が0です。

注意: IPFIX と Netflow v9 はテンプレートベースのため、任意のデータレコードには以下に説明するフィールドが含まれている場合と含まれていない場合があります。例えば"sourceIPv4PrefixLength"などを抽出しようとしても結果が得られない場合は、IPFIX レコードにそのフィールドが含まれていない可能性があります。

便利なので、IPFIXとNetflow v9の最も一般的なフィールドを以下に挙げます。サポートされている名前の完全なリストについては、上記リンク先のドキュメントを参照してください。


| IPFIX Name | Netflow v9 Name |      説明        | サポートされている演算子 | 例 |
|---|---|---|---|---|
| octetDeltaCount | IN_BYTES | 観測点での本フローの受信パケットの前回レポート以降のオクテット数（ある場合）。オクテット数には、IPヘッダとIPペイロードが含まれる | > < <= >= == != | octetDeltaCount == 80 |
| packetDeltaCount | IN_PKTS | 観測点における本フローの前回レポート以降の受信パケット数（ある場合） | > < <= >= == != | packetDeltaCount == 80 |
| deltaFlowCount | FLOWS | この集約フローに寄与している元フローの保守的な数であり、valueDistributionMethod情報要素で表されるいずれかの方法で分配される | > < <= >= == != | packetDeltaCount == 80 |
| protocolIdentifier | PROTOCOL | Protocol フローを表すナンバー (TCP = 6, UDP = 17 | > < <= >= == != | protocolIdentifier == 17 |
| ipClassOfService | TOS | IPv4パケットの場合、IPv4パケットヘッダのTOSフィールドの値、 IPv6パケットの場合、IPv6パケットヘッダのTraffic Classフィールドの値 | > < <= >= == != | ipClassOfService != 0 |
| tcpControlBits | TCP_FLAGS | 本フローのパケットで観測されたTCP制御ビット | > < <= >= == != | tcpControlBits != 0x0004 |
| sourceTransportPort | L4_SRC_PORT | フローのソースポート。 プロトコルがポートを持っていない場合、値はゼロになる | > < <= >= == != | sourceTransportPort != 0 |
| sourceIPv4Address | IPV4_SRC_ADDR| フローのIPv4ソースアドレス | ~ !~ == != | sourceIPv4Address ~ 10.0.0.0/24 |
| sourceIPv4PrefixLength | SRC_MASK | IPv4ソースアドレスのプレフィックスの長さ | > < <= >= == != | sourceIPv4PrefixLength < 24 |
| sourceIPv6Address | IPV6_SRC_ADDR | フローのIPv6ソースアドレス | ~ !~ == != | sourceIPv6Address == ::1 |
| sourceIPv6PrefixLength | IPV6_SRC_MASK | IPv6ソースアドレスのプレフィックスの長さ | > < <= >= == != | sourceIPv6PrefixLength < 64 |
| destinationTransportPort | L4_DST_PORT | フローの宛先ポート。 プロトコルがポートを持たない場合、値はゼロになる | > < <= >= == != | destinationTransportPort != 0 |
| destinationIPv4Address | IPV4_DST_ADDR | フローのIPv4宛先アドレス | ~ !~ == != | destinationIPv4Address ~ 10.0.0.0/24 |
| destinationIPv4PrefixLength | DST_MASK | IPv4の宛先アドレスのプレフィックスの長さ | > < <= >= == != | destinationIPv4PrefixLength < 24 |
| destinationIPv6Address | IPV6_DST_ADDR | フローのIPv6宛先アドレス | ~ !~ == != | destinationIPv6Address == ::1 |
| destinationIPv6PrefixLength | IPV6_DST_MASK | IPv6の宛先アドレスのプレフィックスの長さ | > < <= >= == != | destinationIPv6PrefixLength < 64 |

注意：一般的に、IPFIXメッセージを抽出する際にNetflow v9のフィールド名を指定することができ、その逆も可能です。ただし、2つのプロトコル間でデータタイプが若干異なるため、IPFIXを処理する場合はIPFIXのフィールド名を、Netflowを処理する場合はNetflowの名前を使用するのが最も安全です。

また、このモジュールには、便利な"ショートカット"が用意されています:

| フィールド |       説明        | サポートされているオペレータ | 例 |
|-------|--------------------------|---------------------|---------|
| src | このフローのソースアドレス（IPv4またはIPv6） | ~ !~ == != | src == ::1
| dst | このフローの宛先アドレス (IPv4またはIPv6) | ~ !~ == != | dst !~ PRIVATE
| srcPort | このフローのソースポート | > < <= >= == != | srcPort == 80
| dstPort | このフローの宛先ポート | > < <= >= == != | dstPort == 80
| ip | フィルタにマッチする最初の IP を抽出。 フィルタが指定されていない場合は Src が使用される | ~ !~ == != | ip ~ 10.0.0.0/24
| port | フィルタに一致する最初のポートを抽出する。 フィルタが指定されていない場合は、より低い値が使用される | > < <= >= == != | port == 80
| vlan | フィルタに一致する最初の VLAN を抽出する。VLAN は vlanId または dot1qVlanId フィールドのいずれかから描画される | > < <= >= == != | vlan == 100
| srcMac | フローのソースMACアドレスを抽出する | == != | srcMac==01:23:45:67:89:00
| dstMac | フローの宛先MACアドレスを抽出する | == != | dstMac==01:23:45:67:89:00
| bytes | `octetDeltaCount` と `postOctetDeltaCount` の値を足したもの | > < <= >= == != | bytes <= 10000
| packets | `packetDeltaCount` と `postPacketDeltaCount` の値を足したもの | > < <= >= == != | packets > 0xffffff 
| flowStart | フローの開始タイムスタンプを指定する。`flowStartSeconds`, `flowStartMilliseconds`, `flowStartMicroseconds`, `flowStartNanoseconds` のいずれかのフィールドを使用し、適切なタイムスタンプを出力する | |  flowStart
| flowEnd | フローの終了タイムスタンプ。`flowEndSeconds`, `flowEndMilliseconds`, `flowSEndMicroseconds`, `flowEndNanoseconds` のいずれかのフィールドを使用し、適切なタイムスタンプを出力する | |flowStart
| flowDuration | `flowStart`と`flowEnd`から計算される継続時間の値 | == != < > <= >= | flowDuration > 1m

#### 他のフィールドでのフィルタリング

また、"0x1ad7:0x15"のように、エンタープライズIDとフィールドIDをコロンで区切ってフィールドを指定することもできます。この場合は、`ipfix 0x1ad7:0x15 as foo`のように、より便利な名前で抽出することをお勧めします。この方法でNetflow v9のフィールドを抽出する場合は、エンタープライズIDを0とし、`ipfix 0:7 as srcport`とすれば、ソースポートを抽出することができます。

## 例

### 送信元IP別のHTTPSフロー数の推移

```
tag=ipfix ipfix destinationIPv4Address as Dst destinationTransportPort==443 | count by Dst | chart count by Dst
```

![Number of flows by ip](flowcount.png)

### どのIPがポート80を使用しているかを確認

```
tag=ipfix ipfix port==80 ip ~ PRIVATE | unique ip | table ip
```

### Netflow v9のレコードで最も一般的なプロトコルを検索

この例では、Netflow v9のレコードのみに絞り込んでいます。その際、プロトコルを抽出し、それぞれに何本のフローが現れたかをカウントします。

```
tag=v9 ipfix Version==9 PROTOCOL | count by PROTOCOL | table PROTOCOL count
```

### 各データレコードのフィールドのリストを抽出

```
tag=ipfix ipfix -f | table FIELDS
```

![](fields.png)
