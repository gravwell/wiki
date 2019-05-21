## Packet

The Packet pipeline module extracts fields from Ethernet, IPv4, IPv6, TCP, and UDP packets.  Each field supports an operator which can selectively filter the packet, if no operator is provided the field is extracted if available.

The packet module is useful both for filtering traffic down to specific protocols and for extracting specific fields from packets for analysis--see the examples for more.

Some field modules allow for flexible section where it is desirable to filter on a field that may have a source and destination.  To accomdate selection on IPs, Ports, MACs where there are both a source and destination, the special fields Port, IP, MAC are available.  If either source or destination matches an enumerated value with the field will be populated with the component that matched.  For example, tcp.Port==80 will match whenever either tcp.SrcPort or tcp.DstPort are equal to 80; tpc.Port != 80 will ensure that if either the source or destination ports are 80 the packet is filtered.

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record. For example, the packet processing engine can operate on extracted values such as analyzing layer 2 tunnels.`

### Packet Processing Operators

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

### Packet Processing submodules
The packet processor supports a growing list of submodules which allow for breaking out specific fields in a packet.  Each submodule and field supports a set of operators that allow the packet processor to also filter events based on the subfields.  The following sub modules are available:

| Submodule | Description |
|-----------|-------------|
| eth | Ethernet frames |
| ipv4 | IP Version 4 packets |
| ipv6 | IP Version 6 packets |
| tcp | TCP packets |
| udp | UDP packets |
| icmpv4 | ICMP packets |
| dot1q | VLAN tagged frames |
| dot11 | 802.11 Wireless packets |
| modbus | modbus/TCP packets |

### Packet Processing Submodules

#### Ethernet

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| eth | SrcMAC | == != | eth.SrcMAC==DE:AD:BE:EF:11:22 
| eth | DstMAC | == != | eth.DstMAC != DE:AD:BE:EF:11:22 
| eth | MAC | == != | eth.MAC == DE:AD:BE:EF:11:22 
| eth | Len | > < <= >= == != | eth.Len > 0 
| eth | Type | < > <= >= == != | eth.Type < 5 
| eth | Payload | | eth.Payload

#### VLAN dot1q

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| dot1q | VLANID | > < <= >= == != | dot1q.VLANID > 1024
| dot1q | Priority | > < <= >= == != | dot1q.Priority < 2
| dot1q | Type | > < <= >= == != | dot1q.Type == 2
| dot1q | DropEligible |  == != | dot1q.DropEligible == true

The dot1q packet submodule is designed to enable parsing of VLAN tagged packets.

##### Example Search

An example search that shows all mac addresses routing multiple IPv4 addresses with VLAN tagged packets:

```
tag=pcap packet dot1q.Drop==false eth.SrcMAC ipv4.SrcIP | unique SrcMAC SrcIP | count by SrcMAC | eval count > 1 | table SrcMAC count
```

#### 802.11 Wireless

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| dot11 | Address1 | == != | dot11.Address1==DE:AD:BE:EF:11:22 
| dot11 | Address2 | == != | dot11.Address2 != DE:AD:BE:EF:11:22 
| dot11 | Address3 | == != | dot11.Address3 
| dot11 | Address4 | == != | dot11.Address4 
| dot11 | Type | < > <= >= == ! | dot11.Type == 1
| dot11 | ToDS | == ! | dot11.ToDS == true
| dot11 | FromDS | == ! | dot11.FromDS != false
| dot11 | Payload | | dot11.Payload

#### IPv4

| Packet type | Field | Operators | Example 
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

| Packet type | Field | Operators | Example 
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

| Packet type | Field | Operators | Example 
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

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| udp | SrcPort | == != < > <= >= | udp.SrcPort > 0xfff 
| udp | DstPort | == != < > <= >= | udp.DstPort < 1024 
| udp | Port | == != < > <= >= | udp.Port == 53 
| udp | Length | == != < > <= >= | udp.Length > 100 
| udp | Checksum | == != < > <= >= | udp.Checksum != 0x1234 
| udp | Payload | ~ !~ | udp.Payload


#### ICMP V4

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| icmpv4 | Type | == != < > <= >= | icmpv4.Type < 0x10 
| icmpv4 | Code | == != < > <= >= | icmpv4.Code ==0x2 
| icmpv4 | Checksum | == != < > <= >= | icmpv4.Checksum == 1024 
| icmpv4 | ID | == != < > <= >= | icmpv4.ID == 4 
| icmpv4 | Seq | == != < > <= >= | icmpv4.Seq > 100
| icmpv4 | Payload | == != ~ !~ | icmpv4.Payload

#### ICMP V6

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| icmpv6 | Type | == != < > <= >= | icmpv6.Type < 0x10 
| icmpv6 | Code | == != < > <= >= | icmpv6.Code != 0x2
| icmpv6 | Checksum | == != < > <= >= | icmpv6.Checksum == 1024 
| icmpv6 | Payload | == != ~ !~ | icmpv6.Payload

#### Modbus

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| modbus | Transaction | == != < > <= >= | modbus.Transaction==0x120
| modbus | Protocol | == != < > <= >= | modbus.Protocol==1
| modbus | Length | == != < > <= >= | modbus.Length > 0
| modbus | Unit | == != < > <= >= | modbus.Unit == 2
| modbus | Function | == != < > <= >= | modbus.Function == 0x05
| modbus | Exception | == != | modbus.Exception == false
| modbus | ReqResp | | modbus.ReqResp
| modbus | Payload | | modbus.Payload

For example, the following command will find all DNS queries for tumblr:

```
tag=pcap packet udp.DstPort==53 udp.Payload | grep -e Payload "tumblr" | text
```

The `udp.DstPort==53` component specifies that we should only match on packets destined for UDP port 53, while the `udp.Payload` component specifies that the payload portion of each packet should be extracted into an enumerated value. We then use the `grep` module to search the payload for the word “tumblr” and send the results to the `text` renderer for display.

#### MPLS

The packet search module can decode MPLS headers and allows for selective filtering.  The following MPLS fields are available.

| Packet type | Field | Operators | Example 
|-----|-------|-----------|---------
| mpls | Label | == != < > <= >= | mpls.Label==0x10
| mpls | TrafficClass | == != < > <= >= | mpls.TrafficClass==4
| mpls | StackBottom | == != | mpls.StackBottom==true
| mpls | TTL | == != < > <= >= | mpls.TTL>1
| mpls | Payload | == != ~ !~ | mpls.Payload~foo

For example, the following command will filter all traffic which contains MPLS headers and a traffic Label of 5

```
tag=pcap packet mpls.Label==5 mpls.TrafficClass mpls.Payload | grep -e Payload "HTTP" | count by TrafficClass | table TrafficClass count
```

Note: The MPLS package module will only look at the first MPLS layer, if there are multiple layers you will need to use the [packetlayer](#!search/packetlayer/packetlayer.md) module to decode the additional layers by referencing the Payload enumerated value.

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
