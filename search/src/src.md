## Src

The source module is used for filtering entries based on source, which is a universal metadata item that all entries have.  The module is useful for looking at entries emanating from a specific location.  Src can filter on IP and subnet.

The source module can translate/handle values specified as IPs, subnets, integers, base 16 integers, and 16 byte hashes.  The entry source field is meant to be extremely flexible.

Note: The source field can be used by the acceleration/indexing system to help speed up queries.  However, only direct equality matches invoke the acceleration system.  Filtering by subnet or using negation does not engage the accelerator.

### Example Usage

Eliminate entries coming from a specific source:

```
tag=syslog,apache,pcap src != 192.168.1.1 | count by TAG | chart count by TAG
```

Select only those entries coming from a specific subnet:

```
tag=syslog,apache,pcap src ~ 192.168.1.0/24 | count by SRC | chart count by SRC
```

Eliminate entries from a specific subnet

```
tag=syslog src !~ 192.168.0.0/16
```

Select only entries with a src representing an integer ID

```
tag=syslog src == 1
```

Eliminate entries with a src representing a hex encoded ID

```
tag=syslog src != 0xfeadbeef
```

Search for entries with src as a hex string

```
tag=syslog src == 1234567890ABCDEF0011223344556677
```
