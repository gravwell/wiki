# インラインフィルタリング

非常に頻繁に、いくつかの基準に基づいてクエリ内のエントリをフィルタリングしたいと思うことがあるでしょう。*Inline filtering* は、Gravwellデータの多くの異なるタイプをフィルタリングする効率的な方法です。

Gravwell抽出モジュールは、通常、抽出時に*extracted*項目を*filtered*することを可能にします。パケットからIPv4宛先IPとTCP宛先ポートを抽出する以下のクエリを考えてみましょう。

```
tag=pcap packet ipv4.DstIP tcp.DstPort
```

例えば、10.0.0.0.0/8サブネット内のIPのポート22向けのパケットのみを表示するようにフィルタを追加することができます。

```
tag=pcap packet ipv4.DstIP ~ 10.0.0.0/8 tcp.DstPort == 22
```

DstIPとDstPortが指定されたフィルタと一致しないエントリは、**dropped**されます。

以下のモジュールがフィルタリングをサポートしています。

* [ax](ax/ax.md)
* [canbus](canbus/canbus.md)
* [cef](cef/cef.md)
* [csv](csv/csv.md)
* [fields](fields/fields.md)
* [grok](grok/grok.md)
* [ip](ip/ip.md)
* [ipfix](ipfix/ipfix.md)
* [j1939](j1939/j1939.md)
* [json](json/json.md)
* [kv](kv/kv.md)
* [namedfields](namedfields/namedfields.md)
* [netflow](netflow/netflow.md)
* [packet](packet/packet.md)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [slice](slice/slice.md)
* [subnet](subnet/subnet.md)
* [syslog](syslog/syslog.md)
* [winlog](winlog/winlog.md)
* [xml](xml/xml.md)

## フィルタリング操作とデータタイプ

Gravwell検索パイプラインの中では、列挙された値は、文字列、整数、IPアドレスなど、様々な異なる*types*になる可能性があります。あるIPアドレスが他のIPアドレスより "less than" かどうかを尋ねるのは特に有用ではありません! Gravwellがサポートするフィルタリング操作は以下の通りです。

| オペレーター | 名前 |
|----------|------|
| == | Equal |
| != | Not equal |
| < | Less than |
| > | Greater than |
| <= | Less than or equal |
| >= | Greater than or equal |
| ~ | Subset |
| !~ | Not subset |

これらの操作のほとんどは自明ですが、サブセット操作については特に言及する必要があります。サブセット操作(~)は文字列とIPアドレスに適用されます。文字列の場合は「列挙された値に引数が含まれている」ことを意味し、IPアドレスの場合は「IPアドレスが指定されたサブネット内にある」ことを意味します。したがって、`json domainName ~ "gravwell.io"`は、文字列 "gravwell.io "を含む'domainName'という名前のJSONフィールドを持つエントリのみを渡すことになります。同様に、`packet ipv4.DstIP ~ 10.0.0.0/8`は、IPv4の宛先IPアドレスが10.0.0.0/8サブネットにあるエントリのみを渡します。

各列挙値型は、いくつかのフィルタと互換性がありますが、他のフィルタとは互換性がありません。

| 列挙値タイプ | 互換性のあるオペレータ |
|-----------------------|----------------------|
| string | ==, !=, ~, !~
| byte slice | ==, !=, ~, !~
| MAC address | ==, !=
| IP address | ==, !=, ~, !~
| integer | ==, !=, <, >, <=, >=
| floating point | ==, !=, <, >, <=, >=
| boolean | ==, !=
| duration | ==, !=, <, >, <=, >=

## フィルタリングと加速

データが[アクセラレーションウェル](#!configuration/accelerators.md)にある場合は、インラインフィルターを使用してクエリを高速化できます。 等しい演算子（==）のみが加速を行います。 等式のフィルタリングにより、アクセラレーションエンジンは目的のフィールドに一致するエントリを検索できます。

アクセラレーションを有効にするには、指定したフィールドでアクセラレーションするようにデータを構成する必要があります。 たとえば、tcp.DstPortフィールドとtcp.SrcPortフィールドでpcapデータにインデックスを付ける場合、クエリ `tag = pcap packet tcp.DstPort == 22`はインデックスを使用しますが、` tag = pcap packet ipv4.SrcIP == 10.0.0.1`はしません。

## ビルトインキーワード

Gravwellはフィルタリングのためにいくつかの特別なショートカットを実装しています。IPアドレスをサブネットでフィルタリングする場合、単一のサブネットを指定する代わりに、PRIVATEとMULTICASTというキーワードを指定することができます(例: `packet ipv4.DstIP ~ PRIVATE`)。キーワードは以下のサブネットにマッピングされます。

* PRIVATE: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, 127.0.0.0/8, 224.0.0.0/24, 169.254.0.0/16, fd00::/8, fe80::/10
* MULTICAST: 224.0.0.0/4, ff00::/8

## フィルタリングの例

Regex:

```
tag=syslog regex *shd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)" user==root ip ~ "192.168"
```

AX:

```
tag=testregex,testcsv ax dstport==80 | table app src dst
```

CSV:

```
tag=default csv [5]!=stuff [4] [3]~"things" | table 3 4 5
```

IPFIX:

```
tag=ipfix ipfix destinationIPv4Address as Dst destinationTransportPort==443 | count by Dst | chart count by Dst
```

IP module:

```
tag=csv csv [2] as srcip | ip srcip ~ PRIVATE
```