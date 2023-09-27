# Math Modules

Math modules operate on the pipeline to perform statistical analysis. They are also important when information is condensed over a timeline. For example, if the temperature is measured 10 times per second but the user requests it to be displayed by the second, a math module is used to condense that data.

(sum_module)=

## Sum

The sum module adds the value of the records. This is the default behavior and likely would not be invoked directly.

Example search summing the data that a MAC address has sent on the network:

```gravwell
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

(mean_module)=

## Mean

The mean module returns the average value of the records.
Example search charting vehicle RPM:

```gravwell
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | mean RPM | chart mean
```

(count_module)=

## Count

The count module counts instances of records. It does not conduct any arithmetic on the data within a record.

Example search counting sudo commands from a Linux machine:

```
grep sudo | regex "COMMAND=(?P<command>\S+)" | count by command | chart count by command
```

This is an example search showing how many packets were sent by a MAC address over the network (which is different than the size of each packet as shown in the sum module example):

```gravwell
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

(max_module)=

## Max

The max module returns the maximum value seen.

Example search showing a table of the maximum RPM for each vehicle in an entire fleet:

```gravwell
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | max RPM by SRC | table SRC max
```

(min_module)=

## Min

The min module returns the minimum value seen.

Example search showing a table of the minimum RPM for each vehicle in an entire fleet:

```gravwell
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | min RPM by SRC | table SRC min
```

(variance_module)=

## Variance

The variance module returns the variance information of a record. This is useful for highlighting the rate of change.

Example search charting the variance of throttle data on a Toyota vehicle.

```gravwell
tag=CAN canbus ID=0x2C1 | slice byte(data[6:7]) as throttle | variance throttle | chart variance
```

(stddev_module)=

## Stddev

Standard Deviation

The stddev module returns the standard deviation information of a record. This is useful for highlighting anomalous events.

Example search charting RPM signals that are outliers:

```gravwell
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | stddev RPM | chart stddev
```

