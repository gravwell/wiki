(unique_module)=
# Unique

The unique module drops duplicate entries. The first entry with a given value passes through, and later entries with the same value are dropped.

## Supported Options

* `-maxtracked <arg>`: sets the maximum number of unique keys to track, e.g. `unique -maxtracked 5000 DstIP`. This guards against memory exhaustion, for example when there are millions of IPv6 addresses in the data. If the limit is exceeded, the search fails with an error suggesting you increase it. Defaults to 1000000. See the [stats module documentation](/search/stats/stats) for more about maxtracked.

## Usage

With no arguments, `unique` compares a hash of each entry's `DATA` field. Given one or more enumerated value names, it compares only those values. Compare these two queries:

```gravwell
tag=pcap packet tcp.DstPort | unique
```

```gravwell
tag=pcap packet tcp.DstPort | unique DstPort
```

The first query compares entire packets. Packets usually differ in fields such as sequence numbers, so it drops very little. The second query compares only the extracted `DstPort`: the first packet to TCP port 80 passes through, and every later packet to port 80 is dropped.

With multiple arguments, unique passes the first entry for each combination of values. Entries that have none of the named values are dropped.

```gravwell
tag=pcap packet ipv4.DstIP tcp.DstPort | eval DstPort < 1024 | unique DstPort DstIP | table DstIP DstPort
```

This outputs every unique combination of IP and port where the port is below 1024, which is one way to find servers on a network.

Unique also supports the `over` keyword, which finds unique values within each time window, similar to the stats module. The duration must be a whole number of seconds, minimum one second. For example:

```gravwell
tag=pcap packet tcp.DstPort | unique DstPort over 1h | table DstPort
```

This finds unique destination ports in each 1-hour window.
