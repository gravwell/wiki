## Src

The source module is used for filtering entries based on source, which is a universal metadata item that all entries have.  The module is useful for looking at entries emanating from a specific location.  Src can filter on IP and subnet.

### Example Usage

Eliminate entries coming from a specific source:

```
tag=syslog,apache,pcap src != 192.168.1.1 | count by TAG | chart count by TAG
```

Select only those entries coming from a specific subnet:

```
tag=syslog,apache,pcap src == 192.168.1.0/24 | count by SRC | chart count by SRC
```