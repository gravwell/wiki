# Math Modules

The math modules compute statistics over enumerated values: `sum`, `mean`, `count`, `max`, `min`, `variance`, and `stddev`. They are often used to condense data over a timeline. For example, if a temperature is measured 10 times per second but charted per second, a math module condenses the readings into one value per second.

Each math module is shorthand for the [stats module](/search/stats/stats) with a single operation. The first argument names the source enumerated value; everything after it is passed to `stats` unchanged. These pairs are equivalent:

```
sum Length by SrcMAC
stats sum(Length) by SrcMAC
```

```
count by command
stats count by command
```

The output is an enumerated value named after the operation (`sum`, `count`, and so on). `by`, `as`, and `over` work as described in the [stats documentation](/search/stats/stats), with `as` placed after the source (`sum Bytes as total by SrcIP`). If no source is given, `count` counts entries and the other modules operate on the entry body (`DATA`).

The `total` and `unique_count` operations have no shorthand form; use `stats` for those.

## Supported Options

The math modules do not accept flags. To set `-maxtracked`, `-maxsize`, or `-b`, use the `stats` form instead, e.g. `stats -maxtracked 5000 count by DstIP`. Refer to the [stats module documentation](/search/stats/stats) for details on these flags.

(sum_module)=

## Sum

The sum module adds up the values.

Example search summing the bytes each MAC address has sent on the network:

```gravwell
tag=pcap packet eth.SrcMAC eth.Len | sum Len by SrcMAC | chart sum by SrcMAC
```

(mean_module)=

## Mean

The mean module returns the average value.

Example search charting vehicle RPM:

```gravwell
tag=CAN canbus ID==0x2C4 data | slice uint16BE(data[0:2]) as RPM | mean RPM | chart mean
```

(count_module)=

## Count

The count module counts entries. It does not do any arithmetic on the data within an entry. If a source is given (`count Bytes`), only entries that have that enumerated value are counted.

Example search counting sudo commands from a Linux machine:

```
grep sudo | regex "COMMAND=(?P<command>\S+)" | count by command | chart count by command
```

Example search counting how many packets each MAC address sent (compare with the byte totals in the sum example):

```gravwell
tag=pcap packet eth.SrcMAC | count by SrcMAC | chart count by SrcMAC
```

(max_module)=

## Max

The max module returns the maximum value seen.

Example search showing a table of the maximum RPM for each vehicle in a fleet:

```gravwell
tag=CAN canbus ID==0x2C4 data | slice uint16BE(data[0:2]) as RPM | max RPM by SRC | table SRC max
```

(min_module)=

## Min

The min module returns the minimum value seen.

Example search showing a table of the minimum RPM for each vehicle in a fleet:

```gravwell
tag=CAN canbus ID==0x2C4 data | slice uint16BE(data[0:2]) as RPM | min RPM by SRC | table SRC min
```

(variance_module)=

## Variance

The variance module returns the variance of the values, a measure of how spread out they are.

Example search charting the variance of throttle data on a Toyota vehicle:

```gravwell
tag=CAN canbus ID==0x2C1 data | slice byte(data[6:7]) as throttle | variance throttle | chart variance
```

(stddev_module)=

## Stddev

The stddev module returns the standard deviation of the values. This is useful for spotting anomalous events.

Example search charting the standard deviation of RPM signals:

```gravwell
tag=CAN canbus ID==0x2C4 data | slice uint16BE(data[0:2]) as RPM | stddev RPM | chart stddev
```
