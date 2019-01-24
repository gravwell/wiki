# Join

The join module makes it easier to join two or more enumerated values into a single enumerated value. All enumerated value types are converted to strings and concatenated, except for byte slices which can only be joined to byte slices and will remain byte slices.

The following search will extract the destination IP and port from netflow records and join them with a semicolon as a separator, placing the result in an enumerated value named `dialstring`:

```
tag=netflow netflow Dst DstPort | join -s : -t dialstring Dst DstPort | table Dst DstPort dialstring
```

Any number of enumerated values can be specified. The `-t` flag specifies a "target" enumerated value; if not specified, the first-listed enumerated value will be overwritten.

## Supported Options

* `-s <separator>`: Place the given separator string between the value of each enumerated value in the resulting string. If not specified, no separator will be used. Ignored for byte slices.
* `-t <target>`: Store the result in an enumerated value with the given name rather than overwriting the first enumerated value.

## Example

```
tag=pcap packet ipv4.SrcIP ~ 192.168.0.0/16 tcp.SrcPort | join -s : -t dialstring SrcIP SrcPort | unique SrcIP,SrcPort | table SrcIP SrcPort dialstring
```
![](join.png)
