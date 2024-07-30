(unique_module)=
# Unique

The unique module eliminates duplicate entries in the query data. 

## Supported Options

* `-maxtracked`: sets the maximum number of unique keys to track per operation, e.g. `unique -maxtracked 5000 DstIP`. This is used to help avoid memory exhaustion if there are millions of IPv6 addresses in the data. If the maxtracked value is exceeded, the search will terminate with an error suggesting you should increase the max value. Defaults to 100000000. Refer to the [stats module documentation](/search/stats/stats) for more information about maxtracked.
* `-maxsize <arg>`: sets the maximum amount of memory in megabytes to hold when tracking using keys.

## Usage

Simply specifying `unique` will check for duplicate entries based on the hash of each entry's DATA field. Specifying one or more enumerated value names will cause unique to filter on the enumerated values alone. The difference can be illustrated thus:

```gravwell
tag=pcap packet tcp.DstPort | unique
```

```gravwell
tag=pcap packet tcp.DstPort | unique DstPort
```

The first query will filter out duplicate packets by looking at the entire contents of the packet; because packets typically include things like sequence numbers, it will not accomplish much. The second query uses the extracted DstPort enumerated value to test uniqueness; this means that e.g. the first packet destined for TCP port 80 will pass through, but all further packets destined for port 80 will be dropped.

Specifying multiple arguments means unique will look for each unique combination of those arguments.

```gravwell
tag=pcap packet tcp.DstPort tcp.DstIP | eval DstPort < 1024 | unique DstPort DstIP | table DstIP DstPort
```

The search above will output every unique combination of IP + port, provided the port is less than 1024. This is a useful way to find servers on a network, for instance.

In addition, unique supports the "over" operator, allowing for finding unique values over a given time window, similar to the stats module. For example:

```gravwell
tag=pcap packet tcp.DstPort | unique DstPort over 1h | table DstPort
```

This query will find unique destination ports, split into 1 hour windows.
