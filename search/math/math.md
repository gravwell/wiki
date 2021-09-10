# Math Modules

Math modules operate on the pipeline to perform statistical analysis. They are also important when information is condensed over a timeline. For example, if the temperature is measured 10 times per second but the user requests it to be displayed by the second, a math module is used to condense that data.

## Sum

The sum module adds the value of the records. This is the default behavior and likely would not be invoked directly.

Example search summing the data that a MAC address has sent on the network:

```
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

## Mean

The mean module returns the average value of the records.
Example search charting vehicle RPM:

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | mean RPM | chart mean
```

## Count

The count module counts instances of records. It does not conduct any arithmetic on the data within a record.

Example search counting sudo commands from a Linux machine:

```
grep sudo | regex "COMMAND=(?P<command>\S+)" | count by command | chart count by command
```

This is an example search showing how many packets were sent by a MAC address over the network (which is different than the size of each packet as shown in the sum module example):

```
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

## Max

The max module returns the maximum value seen.

Example search showing a table of the maximum RPM for each vehicle in an entire fleet:

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | max RPM by SRC | table SRC max
```

## Min

The min module returns the minimum value seen.

Example search showing a table of the minimum RPM for each vehicle in an entire fleet:

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | min RPM by SRC | table SRC min
```

## Variance

The variance module returns the variance information of a record. This is useful for highlighting the rate of change.

Example search charting the variance of throttle data on a Toyota vehicle.

```
tag=CAN canbus ID=0x2C1 | slice byte(data[6:7]) as throttle | variance throttle | chart variance
```

## Stddev

Standard Deviation

The stddev module returns the standard deviation information of a record. This is useful for highlighting anomalous events.

Example search charting RPM signals that are outliers:

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | stddev RPM | chart stddev
```

## Unique

The unique module eliminates duplicate entries in the query data. Simply specifying `unique` will check for duplicate entries based on the hash of each entry's data. Specifying one or more enumerated value names will cause unique to filter on the enumerated values alone. The difference can be illustrated thus:

```
tag=pcap packet tcp.DstPort | unique
```

```
tag=pcap packet tcp.DstPort | unique DstPort
```

The first query will filter out duplicate packets by looking at the entire contents of the packet; because packets typically include things like sequence numbers, it will not accomplish much. The second query uses the extracted DstPort enumerated value to test uniqueness; this means that e.g. the first packet destined for TCP port 80 will pass through, but all further packets destined for port 80 will be dropped.

Specifying multiple arguments means unique will look for each unique combination of those arguments.

```
tag=pcap packet tcp.DstPort tcp.DstIP | eval DstPort < 1024 | unique DstPort DstIP | table DstIP DstPort
```

The search above will output every unique combination of IP + port, provided the port is less than 1024. This is a useful way to find servers on a network, for instance.
