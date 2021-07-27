## Packet Layer

The packet module is provides a very robust interface into packet processing but has one significant hindrance, it expects sanity.  When invoking various operations the packet processor inherits necessary parent layers in order to accurately deconstruct a packet.  When a packet query asks for a tcp port, the packet processor invokes the necessary sublayers to get to the tcp layer (whether it be Ethernet, 802.11q, IPv4, IPv6).  In the case where packet layer processing is required, but the packets may be malformed, wrapped, or just straight corrupt the "Packet Layer" can be used to handle JUST the layers specified.

The invocation for the "packetlayer" module is identical to the "packet" module, but it expects the lowest level layer to be the only layers present.  Below are are some example queries where this may prove useful.  The packetlayer parser will also continue to extract packet layers as long as there is data available in an entry or enumerated value, making it possible to handle very large blobs of packet layers.  Do you have a flat binary file of netflow packets?  Packetlayer can handle that.

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record. For example, the packet processing engine can operate on extracted values such as analyzing layer 2 tunnels.`
* `-m : The “-m” option tells the packetlayer processor to continue extracting the specified layers until no data is available (multi-extract).  This option is useful for extracting Layer 4+ messages`

### Layer 4+ Protocol Extraction

Protocol layers above layer 3 are often stream based, relying on lower layers to handle transport, which often means that multiple layer 4+ messages may be in a single layer 3 payload.  An example query leveraging the packet layer processor to extract these messages uses the packet module to get to the tcp layer payload and the packetlayer module to extract multiple Modbus payloads.

```
tag=pcap tcp.Port == 502 tcp.Payload | packetlayer -m -e Payload modbus.Transaction != 0 modbus.Unit | count by Unit | chart count by Unit
```

### Extracting Wrapped Packet Transports

We occasionally encounter transport protocols that are not documented or hand rolled by angry developers.  Using the packetlayer module we can still maintain sanity.

```
tag=pcap ipv6.SrcIP udp.Port == 31337 udp.Payload | slice Payload[0:4] as id Payload[4:] as payload2 | packetlayer -e payload2 modbus.Transaction != 0 modbus.Unit | count by Unit | chart count by Unit
```