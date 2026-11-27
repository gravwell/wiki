# Packet Layer

The [packet](/search/packet/packet) module expects complete, well-formed packets. To reach a requested layer it decodes every parent layer first: asking for a TCP port means decoding Ethernet, any VLAN tags, and IPv4 or IPv6 on the way. When the data is malformed, wrapped in an unknown transport, or starts partway into a packet, use the packetlayer module instead. It decodes only the layers you specify, starting at the first byte of the data.

The packetlayer module takes the same `<submodule>.<field>` arguments and operators as the packet module, and supports the same submodules. With `-m`, it keeps decoding the specified layers until the data runs out, so it can handle a blob of back-to-back messages. Entries that fail to decode or don't match the filters are dropped.

## Supported Options

* `-e <arg>`: Read the data from an enumerated value instead of the entry data, for example a payload extracted by the packet module. Entries without the enumerated value are dropped.
* `-m`: Multi-extract. After decoding the specified layers, start again on the remaining payload, emitting one entry per message, until no data is left or a message fails to decode or match the filters. Useful for layer 4+ messages.

## Layer 4+ Protocol Extraction

Protocols above layer 3 are often stream based and rely on lower layers for transport, so a single TCP payload may hold several layer 4+ messages. The following search uses the packet module to extract the TCP payload, then the packetlayer module to extract each Modbus message in it.

```gravwell
tag=pcap packet tcp.Port == 502 tcp.Payload | packetlayer -m -e Payload modbus.Transaction != 0 modbus.Unit | count by Unit | chart count by Unit
```

## Extracting Wrapped Packet Transports

Some transports are undocumented or hand-rolled. In the following search, a custom UDP protocol carries a 4-byte ID followed by a Modbus message. The `slice` module splits off the ID, and the packetlayer module decodes the rest.

```gravwell
tag=pcap packet ipv6.SrcIP udp.Port == 31337 udp.Payload | slice Payload[0:4] as id Payload[4:] as payload2 | packetlayer -e payload2 modbus.Transaction != 0 modbus.Unit | count by Unit | chart count by Unit
```
