## パケット

パケットパイプラインモジュールは、イーサネット、IPv4、IPv6、TCP、UDP のパケットからフィールドを抽出します。 各フィールドは、パケットを選択的にフィルタリングすることができる演算子をサポートしています。演算子が提供されていない場合は、利用可能なフィールドが抽出されます。

パケットモジュールは、トラフィックを特定のプロトコルに絞り込んでフィルタリングする場合と、パケットから特定のフィールドを抽出して解析する場合の両方に有効です。詳細はサンプルをご覧ください。

一部のフィールド・モジュールでは、ソースとデスティネーションを持つフィールドでフィルタリングすることが望ましい、柔軟なセクションが可能です。 ソースとデスティネーションの両方がある IP、ポート、MAC の選択に対応するために、Port、IP、MAC という特殊なフィールドが用意されています。 送信元または送信先のいずれかが列挙値に一致した場合、フィールドには一致したコンポーネントが入力されます。 例えば、tcp.Port==80 は、tcp.SrcPort または tcp.DstPort のいずれかが 80 に等しい場合にマッチします。tpc.Port !=80 は、ソースポートまたはデスティネーションポートのいずれかが 80 の場合、パケットがフィルタリングされることを保証します。

### サポートされているオプション

* `-e <arg>`: 「-e」オプションは、レコード全体ではなく、列挙値を操作します。例えば、パケット処理エンジンは、レイヤー 2 のトンネルの分析など、抽出された値に対して操作することができます。

### パケット処理演算子

| 演算子 | 名称 | 意味
|----------|------|-------------
| == | 等しい | フィールドは等しい
| != | 等しくない | フィールドは等しくない
| < | 小なり | フィールドはその値より小さい
| > | 大なり | フィールドはその値より大きい
| <= | 小なりイコール | フィールドはその値以下である
| >= | 大なりイコール | フィールドはその値以上である
| ~ | 含まれる | フィールドはそれに含まれる
| !~ | 含まれない | フィールドはそれに含まれない

### パケット処理のサブモジュール

パケットプロセッサは、パケット内の特定のフィールドを分割することができるサブモジュールをサポートしています。各サブモジュールとフィールドは、パケットプロセッサがサブフィールドに基づいてイベントをフィルタリングすることを可能にする演算子のセットをサポートしています。 以下のサブモジュールがあります。

| サブモジュール | 意味 |
|-----------|-------------|
| eth | イーサネットフレーム |
| ipv4 | IP バージョン 4 のパケット |
| ipv6 | IP バージョン 6 のパケット |
| tcp | TCP パケット |
| udp | UDP パケット |
| icmpv4 | ICMP パケット |
| dot1q | VLAN タグ付きフレーム |
| dot11 | 802.11 ワイヤレスパケット |
| dot11info | 802.11 の情報要素 |
| modbus | modbus/TCP パケット |
| MPLS | マルチプロトコル・ラベル・スイッチング |

### 抽出物の名称変更

列挙値の名前は、サブモジュールの仕様書の最後の名前によって導き出されます。たとえば、「ipv4.SrcIP」という仕様では、「SrcIP」という列挙値が生成されます。列挙型の値の名前は、as の引数で上書きすることができます。例えば、ipv4.SrcIP を foo として抽出するには、次のようにします。

```
tag=pcap packet ipv4.SrcIP as foo | table foo
```

名前を変更した列挙型の値は、フィルターでも使用できます。

```
tag=pcap packet ipv4.SrcIP="8.8.8.8" as foo | table foo
```

### パケット処理のサブモジュールの一覧

#### イーサネット

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| eth | SrcMAC | == != | eth.SrcMAC==DE:AD:BE:EF:11:22
| eth | DstMAC | == != | eth.DstMAC != DE:AD:BE:EF:11:22
| eth | MAC | == != | eth.MAC == DE:AD:BE:EF:11:22
| eth | Len | > < <= >= == != | eth.Len > 0
| eth | Type | < > <= >= == != | eth.Type < 5
| eth | Payload | | eth.Payload

#### VLAN dot1q

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| dot1q | VLANID | > < <= >= == != | dot1q.VLANID > 1024
| dot1q | Priority | > < <= >= == != | dot1q.Priority < 2
| dot1q | Type | > < <= >= == != | dot1q.Type == 2
| dot1q | DropEligible |  == != | dot1q.DropEligible == true

dot1q パケットサブモジュールは、VLAN タグ付きパケットの解析を可能にするためのものです。

##### 検索例

VLAN タグ付きパケットで複数の IPv4 アドレスをルーティングするすべての mac アドレスを表示する検索例です。

```
tag=pcap packet dot1q.Drop==false eth.SrcMAC ipv4.SrcIP | unique SrcMAC SrcIP | count by SrcMAC | eval count > 1 | table SrcMAC count
```

#### 802.11 ワイヤレス

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| dot11 | Address1 | == != | dot11.Address1==DE:AD:BE:EF:11:22
| dot11 | Address2 | == != | dot11.Address2 != DE:AD:BE:EF:11:22
| dot11 | Address3 | == != | dot11.Address3
| dot11 | Address4 | == != | dot11.Address4
| dot11 | Type | < > <= >= == ! | dot11.Type == 1
| dot11 | ToDS | == ! | dot11.ToDS == true
| dot11 | FromDS | == ! | dot11.FromDS != false
| dot11 | Payload | | dot11.Payload

#### 802.11 の情報要素

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| dot11info | SSID | == != | dot11.SSID != xfinitywifi

#### IPv4

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| ipv4 | Version | == != < > <= >= | ipv4.Version != 0b11
| ipv4 | IHL | == != < > <= >= | ipv4.IHL == 08
| ipv4 | TOS | == != < > <= >= | ipv4.TOS < 10
| ipv4 | Length | == != < > <= >= | ipv4.Length > 0xff
| ipv4 | ID | == != < > <= >= | ipv4.ID == 0x5
| ipv4 | Flag | == != < > <= >= | ipv4.Flag == 0b1101
| ipv4 | FragOffset | == != < > <= >= | ipv4.FragOffset > 3
| ipv4 | TTL | == != < > <= >= | ipv4.TTL < 2
| ipv4 | Protocol | == != < > <= >= | ipv4.Protocol != 0x08
| ipv4 | Checksum | == != < > <= >= | ipv4.Checksum <= 0x1234
| ipv4 | SrcIP | == != ~ !~ | ipv4.SrcIP ~ 192.168.1.1/16
| ipv4 | DstIP | == != ~ !~ | ipv4.DstIP !~ 10.10.10.1/8
| ipv4 | IP | == != ~ !~ | ipv4.IP ~ 192.168.1.0/14
| ipv4 | Payload | | ipv4.Payload

#### IPv6

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| ipv6 | Version | == != < > <= >= | ipv6.Version == 0x08
| ipv6 | TrafficClass | == != < > <= >= | ipv6.TrafficClass != 20
| ipv6 | FlowLabel | == != < > <= >= | ipv6.FlowLabel == 0xDEADBEEF
| ipv6 | Length | == != < > <= >= | ipv6.Length >= 100
| ipv6 | NextHeader | == != < > <= >= | ipv6.NextHeader == 0x0800
| ipv6 | HopLimit | == != < > <= >= | ipv6.HopLimit < 10
| ipv6 | SrcIP | == != ~ !~ | ipv6.SrcIP != FF02::1
| ipv6 | DstIP | == != ~ !~ | ipv6.DstIP !~ FE80::1/64
| ipv6 | IP | == != ~ !~ | ipv6.IP == FE80::1/64
| ipv6 | Payload | | ipv6.Payload

#### TCP

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| tcp | SrcPort | == != < > <= >= | tcp.SrcPort > 1024
| tcp | DstPort | == != < > <= >= | tcp.DstPort <= 1024
| tcp | Port | == != < > <= >= | tcp.Port == 80
| tcp | SeqNum | == != < > <= >= | tcp.SeqNum > 0xffff
| tcp | AckNum | == != < > <= >= | tcp.AckNum < 112345
| tcp | Window | == != < > <= >= | tcp.Window < 1024
| tcp | [SYN/ACK/FIN/RST/PSH/URG/ECE/CWR/NS] | ==true, != true | tcp.SYN == true
| tcp | Checksum | == != < > <= >= | tcp.Checksum != 0x1234
| tcp | Urgent | == != < > <= >= | tcp.Urgent==0b111010101010101
| tcp | DataOffset | == != < > <= >= | tcp.DataOffset > 96
| tcp | Payload | ~ !~ | tcp.Payload ~ "HTTP"

#### UDP

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| udp | SrcPort | == != < > <= >= | udp.SrcPort > 0xfff
| udp | DstPort | == != < > <= >= | udp.DstPort < 1024
| udp | Port | == != < > <= >= | udp.Port == 53
| udp | Length | == != < > <= >= | udp.Length > 100
| udp | Checksum | == != < > <= >= | udp.Checksum != 0x1234
| udp | Payload | ~ !~ | udp.Payload


#### ICMP V4

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| icmpv4 | Type | == != < > <= >= | icmpv4.Type < 0x10
| icmpv4 | Code | == != < > <= >= | icmpv4.Code ==0x2
| icmpv4 | Checksum | == != < > <= >= | icmpv4.Checksum == 1024
| icmpv4 | ID | == != < > <= >= | icmpv4.ID == 4
| icmpv4 | Seq | == != < > <= >= | icmpv4.Seq > 100
| icmpv4 | Payload | == != ~ !~ | icmpv4.Payload

#### ICMP V6

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| icmpv6 | Type | == != < > <= >= | icmpv6.Type < 0x10
| icmpv6 | Code | == != < > <= >= | icmpv6.Code != 0x2
| icmpv6 | Checksum | == != < > <= >= | icmpv6.Checksum == 1024
| icmpv6 | Payload | == != ~ !~ | icmpv6.Payload

#### Modbus

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| modbus | Transaction | == != < > <= >= | modbus.Transaction==0x120
| modbus | Protocol | == != < > <= >= | modbus.Protocol==1
| modbus | Length | == != < > <= >= | modbus.Length > 0
| modbus | Unit | == != < > <= >= | modbus.Unit == 2
| modbus | Function | == != < > <= >= | modbus.Function == 0x05
| modbus | Exception | == != | modbus.Exception == false
| modbus | ReqResp | | modbus.ReqResp
| modbus | Payload | | modbus.Payload

例えば、次のコマンドは、tumblr に対するすべての DNS クエリを検索します。

```
tag=pcap packet udp.DstPort==53 udp.Payload | grep -e Payload "tumblr" | text
```

また、`udp.DstPort==53` というコンポーネントは、UDP ポート 53 宛のパケットにのみマッチさせることを指定し、`udp.Payload` というコンポーネントは、各パケットのペイロード部分を列挙型の値に抽出することを指定しています。次に、`grep` モジュールを使って、ペイロードの中から「tumblr」という単語を検索し、その結果を `text` レンダラーに送って表示させます。

#### MPLS

パケット検索モジュールは、MPLS ヘッダーをデコードすることができ、選択的なフィルタリングが可能です。 以下の MPLS フィールドがあります。

| パケットタイプ | フィールド | 演算子 | 例
|-----|-------|-----------|---------
| mpls | Label | == != < > <= >= | mpls.Label==0x10
| mpls | TrafficClass | == != < > <= >= | mpls.TrafficClass==4
| mpls | StackBottom | == != | mpls.StackBottom==true
| mpls | TTL | == != < > <= >= | mpls.TTL>1
| mpls | Payload | == != ~ !~ | mpls.Payload~foo

例えば、次のコマンドは、MPLS ヘッダーを含み、トラフィックの Label が 5 であるすべてのトラフィックをフィルタリングします。

```
tag=pcap packet mpls.Label==5 mpls.TrafficClass mpls.Payload | grep -e Payload "HTTP" | count by TrafficClass | table TrafficClass count
```

注意：MPLS package モジュールは、最初の MPLS レイヤーのみを見ます。複数のレイヤーがある場合は、[packetlayer](#!search/packetlayer/packetlayer.md) モジュールを使用して、Payload 列挙値を参照して追加レイヤをデコードする必要があります。

<!---
### ICS-specific protocols

Gravwell includes basic protocol crackers for Modbus, Ethernet/IP, and CIP. Due to the complexity of Ethernet/IP and CIP, only basic decoding is available, but this can still help establish baselines and detect anomalies.

| Packet type | Field | Operators | Example
|-----|-------|-----------|---------
| modbus | Transaction | == != < > <= >= | modbus.Transaction != 0
| modbus | Protocol | == != < > <= >= | modbus.Protocol != 0
| modbus | Length | == != < > <= >= | modbus.Length > 1
| modbus | Unit | == != < > <= >= | modbus.Unit != 255
| modbus | Function | == != < > <= >= | modbus.Function == 0x0f
| modbus | Exception | ==true, !=true | modbus.Exception == true
| modbus | ReqResp | | modbus.ReqResp
| modbus | Payload | | modbus.Payload
#| enip | Command | == != < > <= >= | enip.Command > 0
#| enip | Length | == != < > <= >= | enip.Length > 5
#| enip | SessionHandle | == != < > <= >= | enip.SessionHandle != 0
#| enip | Status | == != < > <= >= | enip.Status != 0
#| enip | Options | == != < > <= >= | enip.Options == 0x02
#| enip | CommandSpecific | | enip.CommandSpecific
#| enip | Payload | | enip.Payload
#| enip | SenderContext | | enip.SenderContext
#| cip | Response | ==true, !=true | cip.Response == true
#| cip | Service | == != < > <= >= | cip.Service == 0x02
#| cip | ClassID | == != < > <= >= | cip.ClassID == 0x00
#| cip | InstanceID | == != < > <= >= | cip.InstanceID == 0x01
#| cip | Status | == != < > <= >= | cip.Status != 0
#| cip | AdditionalStatus | | cip.AdditionalStatus
#| cip | Data | | cip.Data
-->
