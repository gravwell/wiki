# Require

Require drops entries that lack specific enumerated values. Use it as a second-level filter when an upstream module extracts fields but doesn't always populate every one you need.

By default, an entry passes if it has **at least one** of the listed enumerated values.

## Supported Options

* `-s`: Strict. The entry must have *all* listed enumerated values. This turns the default OR into an AND.
* `-v`: Invert. Drop entries that have *any* of the listed enumerated values, so only entries with none of them pass. Use it to find entries where an upstream module did not extract a field.

## Example Usage

Drop packets without a "SrcPort" enumerated value, then count each source port. This eliminates packets that are not TCP:

```gravwell
tag=pcap packet tcp.SrcPort | require SrcPort | count by SrcPort | table SrcPort count
```

Find DNS requests with no IPv4 answer by dropping entries that have an "A" enumerated value:

```gravwell
tag=dns json Question.Hdr.Name Question.A | require -v A | count by Name | table Name count
```

Pass any packet that has *either* a TCP or a UDP source port:

```gravwell
tag=pcap packet tcp.SrcPort as tsp udp.SrcPort as usp | require tsp usp | table tsp usp
```

Drop any packet that does not have *both* an IPv4 source IP *and* a TCP source port:

```gravwell
tag=pcap packet tcp.SrcPort ipv4.SrcIP | require -s SrcPort SrcIP | table SrcPort SrcIP
```
