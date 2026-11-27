# Packet

The packet module extracts fields from raw packets, such as Ethernet, IP, TCP, and UDP headers. The supported layers are listed below. Give a field an operator to filter on it. Without an operator, the field is extracted if present. Packets that don't contain a requested layer, or don't match a filter, are dropped.

Use it to filter traffic down to specific protocols, or to extract fields from packets for analysis.

The special fields `Port`, `IP`, and `MAC` match either the source or the destination value, and the extracted enumerated value holds whichever one matched. For example, `tcp.Port==80` matches when either `tcp.SrcPort` or `tcp.DstPort` is 80, while `tcp.Port != 80` drops the packet if either port is 80. Without a filter, `IP` and `MAC` extract the source value and `Port` extracts the lower of the two ports.

## Supported Options

* `-e <arg>`: Read the packet from an enumerated value instead of the entry data, for example to analyze a layer 2 tunnel extracted by an earlier module. Entries without the enumerated value are dropped.

## Packet Processing Operators

| Operator | Name | Description
|----------|------|-------------
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| < | Less than | Field must be less than
| > | Greater than | Field must be greater than
| <= | Less than or equal | Field must be less than or equal to
| >= | Greater than or equal | Field must be greater than or equal to
| ~ | Subset | Field must be a member of
| !~ | Not subset | Field must not be a member of

The operators a field supports depend on its type. Numeric values can be decimal, hex (`0x1f`), or binary (`0b101`); a leading zero means octal. IP fields take a single address for `==` and `!=`, and a CIDR subnet or the keyword `PRIVATE` or `MULTICAST` for `~` and `!~`. For `Payload` fields, `~` and `!~` check whether the payload contains the given bytes.

## Packet Processing Submodules

Each submodule breaks out the fields of one protocol layer. Specify a field as `<submodule>.<field>`, e.g. `ipv4.SrcIP`. The following submodules are available:

| Submodule | Description                    |
|-----------|--------------------------------|
| eth       | Ethernet frames                |
| ipv4      | IP Version 4 packets           |
| ipv6      | IP Version 6 packets           |
| tcp       | TCP packets                    |
| udp       | UDP packets                    |
| icmpv4    | ICMP packets                   |
| icmpv6    | ICMPv6 packets                |
| dot1q     | VLAN tagged frames             |
| dot11     | 802.11 Wireless packets        |
| dot11info | 802.11 information elements    |
| modbus    | modbus/TCP packets             |
| mpls      | Multi-protocol Label Switching |

## Renaming Extractions

An extracted field's enumerated value is named after the field, so `ipv4.SrcIP` produces an enumerated value named `SrcIP`. Use `as` to choose a different name. For example, to extract `ipv4.SrcIP` as `foo`:

```gravwell
tag=pcap packet ipv4.SrcIP as foo | table foo
```

Renamed enumerated values can also be used with filters:

```gravwell
tag=pcap packet ipv4.SrcIP==8.8.8.8 as foo | table foo
```

## Examples

The following search finds DNS queries for Tumblr:

```gravwell
tag=pcap packet udp.DstPort==53 udp.Payload | grep -e Payload "tumblr" | text
```

`udp.DstPort==53` keeps only packets sent to UDP port 53, and `udp.Payload` extracts each packet's payload into an enumerated value. The `grep` module then searches the payload for "tumblr" and the `text` renderer displays the results.

## List of Packet Processing Submodules

### Ethernet

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| eth | SrcMAC | == != | eth.SrcMAC==DE:AD:BE:EF:11:22 
| eth | DstMAC | == != | eth.DstMAC != DE:AD:BE:EF:11:22 
| eth | MAC | == != | eth.MAC == DE:AD:BE:EF:11:22 
| eth | Len | > < <= >= == != | eth.Len > 0 
| eth | Type | < > <= >= == != | eth.Type < 5 
| eth | Payload | == != ~ !~ | eth.Payload

### VLAN dot1q

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| dot1q | VLANID | > < <= >= == != | dot1q.VLANID > 1024
| dot1q | Priority | > < <= >= == != | dot1q.Priority < 2
| dot1q | Type | > < <= >= == != | dot1q.Type == 2
| dot1q | DropEligible |  == != | dot1q.DropEligible == true

The dot1q submodule parses VLAN tagged frames.

#### Example Search

The following search shows MAC addresses that send VLAN tagged packets from more than one IPv4 address:

```gravwell
tag=pcap packet dot1q.DropEligible==false eth.SrcMAC ipv4.SrcIP | unique SrcMAC SrcIP | count by SrcMAC | eval count > 1 | table SrcMAC count
```

### 802.11 Wireless

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| dot11 | Address1 | == != | dot11.Address1==DE:AD:BE:EF:11:22 
| dot11 | Address2 | == != | dot11.Address2 != DE:AD:BE:EF:11:22 
| dot11 | Address3 | == != | dot11.Address3 
| dot11 | Address4 | == != | dot11.Address4 
| dot11 | Type | < > <= >= == != | dot11.Type == 1
| dot11 | ToDS | == != | dot11.ToDS == true
| dot11 | FromDS | == != | dot11.FromDS != false
| dot11 | Payload | == != ~ !~ | dot11.Payload

### 802.11 Information Elements

| Packet Type | Field | Operators | Example
|-----|-------|-----------|---------
| dot11info | SSID | == != ~ !~ | dot11info.SSID != xfinitywifi

### IPv4

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| ipv4 | Version | == != < > <= >= | ipv4.Version != 0b11 
| ipv4 | IHL | == != < > <= >= | ipv4.IHL > 5 
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
| ipv4 | Payload | == != ~ !~ | ipv4.Payload

### IPv6

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| ipv6 | Version | == != < > <= >= | ipv6.Version == 6 
| ipv6 | TrafficClass | == != < > <= >= | ipv6.TrafficClass != 20 
| ipv6 | FlowLabel | == != < > <= >= | ipv6.FlowLabel == 0xDEADBEEF 
| ipv6 | Length | == != < > <= >= | ipv6.Length >= 100 
| ipv6 | NextHeader | == != < > <= >= | ipv6.NextHeader == 6 
| ipv6 | HopLimit | == != < > <= >= | ipv6.HopLimit < 10 
| ipv6 | SrcIP | == != ~ !~ | ipv6.SrcIP != FF02::1 
| ipv6 | DstIP | == != ~ !~ | ipv6.DstIP !~ FE80::1/64 
| ipv6 | IP | == != ~ !~ | ipv6.IP ~ FE80::/64 
| ipv6 | Payload | == != ~ !~ | ipv6.Payload

### TCP

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| tcp | SrcPort | == != < > <= >= | tcp.SrcPort > 1024 
| tcp | DstPort | == != < > <= >= | tcp.DstPort <= 1024 
| tcp | Port | == != < > <= >= | tcp.Port == 80
| tcp | SeqNum | == != < > <= >= | tcp.SeqNum > 0xffff 
| tcp | AckNum | == != < > <= >= | tcp.AckNum < 112345 
| tcp | Window | == != < > <= >= | tcp.Window < 1024 
| tcp | [SYN/ACK/FIN/RST/PSH/URG/ECE/CWR/NS] | == != | tcp.SYN == true 
| tcp | Checksum | == != < > <= >= | tcp.Checksum != 0x1234 
| tcp | Urgent | == != < > <= >= | tcp.Urgent==0b111010101010101 
| tcp | DataOffset | == != < > <= >= | tcp.DataOffset > 5 
| tcp | Payload | == != ~ !~ | tcp.Payload ~ "HTTP"

### UDP

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| udp | SrcPort | == != < > <= >= | udp.SrcPort > 0xfff 
| udp | DstPort | == != < > <= >= | udp.DstPort < 1024 
| udp | Port | == != < > <= >= | udp.Port == 53 
| udp | Length | == != < > <= >= | udp.Length > 100 
| udp | Checksum | == != < > <= >= | udp.Checksum != 0x1234 
| udp | Payload | == != ~ !~ | udp.Payload


### ICMP V4

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| icmpv4 | Type | == != < > <= >= | icmpv4.Type < 0x10 
| icmpv4 | Code | == != < > <= >= | icmpv4.Code ==0x2 
| icmpv4 | Checksum | == != < > <= >= | icmpv4.Checksum == 1024 
| icmpv4 | ID | == != < > <= >= | icmpv4.ID == 4 
| icmpv4 | Seq | == != < > <= >= | icmpv4.Seq > 100

### ICMP V6

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| icmpv6 | Type | == != < > <= >= | icmpv6.Type < 0x10 
| icmpv6 | Code | == != < > <= >= | icmpv6.Code != 0x2
| icmpv6 | Checksum | == != < > <= >= | icmpv6.Checksum == 1024 

### Modbus

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| modbus | Transaction | == != < > <= >= | modbus.Transaction==0x120
| modbus | Protocol | == != < > <= >= | modbus.Protocol==1
| modbus | Length | == != < > <= >= | modbus.Length > 0
| modbus | Unit | == != < > <= >= | modbus.Unit == 2
| modbus | Function | == != < > <= >= | modbus.Function == 0x05
| modbus | Exception | == != | modbus.Exception == false
| modbus | ReqResp | == != ~ !~ | modbus.ReqResp
| modbus | Payload | == != ~ !~ | modbus.Payload

```{note}
The Modbus layer is typically the last layer, so its Payload value is almost always empty. The ReqResp value holds the read/write command addresses and values.
```

### MPLS

The following MPLS fields are available:

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| mpls | Label | == != < > <= >= | mpls.Label==0x10
| mpls | TrafficClass | == != < > <= >= | mpls.TrafficClass==4
| mpls | StackBottom | == != | mpls.StackBottom==true
| mpls | TTL | == != < > <= >= | mpls.TTL>1
| mpls | Payload | == != ~ !~ | mpls.Payload~foo

The following search keeps traffic with an MPLS label of 5 whose payload contains "HTTP", and counts it by traffic class:

```gravwell
tag=pcap packet mpls.Label==5 mpls.TrafficClass mpls.Payload | grep -e Payload "HTTP" | count by TrafficClass | table TrafficClass count
```

```{note}
The mpls submodule only decodes the first MPLS layer. To decode additional layers, pass the Payload enumerated value to the [packetlayer](/search/packetlayer/packetlayer) module.
```

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
