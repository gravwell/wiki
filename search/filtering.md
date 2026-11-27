# Inline Filtering

You will often want to filter entries on some criteria: drop HTTPS flows from Netflow data, look only at traffic from a certain subnet, or match a particular username in Windows logs. *Inline filtering* does this efficiently for many types of data.

Most extraction modules can *filter* items as they *extract* them. This query extracts the IPv4 destination IP and the TCP destination port from packets:

```gravwell
tag=pcap packet ipv4.DstIP tcp.DstPort
```

Add filters to show only packets destined for port 22 on IPs in the 10.0.0.0/8 subnet:

```gravwell
tag=pcap packet ipv4.DstIP ~ 10.0.0.0/8 tcp.DstPort == 22
```

Any entry whose DstIP and DstPort do not match the filters is **dropped**.

## Filtering Operations & Data Types

Enumerated values have *types*, such as strings, integers, or IP addresses, and not every operator makes sense for every type: asking whether one IP address is "less than" another isn't useful. Gravwell supports these filter operators:

| Operator | Name |
|----------|------|
| == | Equal |
| != | Not equal |
| < | Less than |
| > | Greater than |
| <= | Less than or equal |
| >= | Greater than or equal |
| ~ | Subset |
| !~ | Not subset |

The subset operators need some explanation. For strings and byte slices, `~` means "the value contains the argument" (a case-sensitive substring match). For IP addresses, it means "the address is within the given subnet". So `json domainName ~ "gravwell.io"` passes only entries with a JSON field `domainName` that contains `gravwell.io`, and `packet ipv4.DstIP ~ 10.0.0.0/8` passes only entries whose IPv4 destination address is in 10.0.0.0/8. `!~` is the inverse of each.

A single `=` is not a filter operator; use `==`.

Each enumerated value type supports these operators:

| Enumerated Value Type | Compatible Operators |
|-----------------------|----------------------|
| string | ==, !=, ~, !~
| byte slice | ==, !=, ~, !~
| MAC address | ==, !=
| IP address | ==, !=, ~, !~
| integer | ==, !=, <, >, <=, >=
| floating point | ==, !=, <, >, <=, >=
| boolean | ==, !=
| duration | ==, !=, <, >, <=, >=

If an extraction module can't find the value in an entry at all, `==`, `~`, `<`, `>`, `<=`, and `>=` filters drop the entry, while `!=` and `!~` filters pass it. For example, `json user != "root"` passes entries that have no `user` field.

## Filtering & Acceleration

If your data is in an [accelerated well](/configuration/accelerators), inline filters can speed up queries. Only the equality operator (`==`) engages acceleration, because it lets the accelerator look up entries that match the field.

The well must be configured to accelerate on the filtered field. For instance, if pcap data is indexed on `tcp.DstPort` and `tcp.SrcPort`, `tag=pcap packet tcp.DstPort==22` uses the index but `tag=pcap packet ipv4.SrcIP==10.0.0.1` does not.

## Built-in Keywords

When filtering IP addresses by subnet, you can use the keywords `PRIVATE` and `MULTICAST` in place of a subnet, e.g. `packet ipv4.DstIP ~ PRIVATE`. They map to these subnets:

* PRIVATE: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, 127.0.0.0/8, 224.0.0.0/24, 169.254.0.0/16, fd00::/8, fe80::/10, ::1/128
* MULTICAST: 224.0.0.0/4, ff00::/8

## Filtering Examples

Regex:

```gravwell
tag=syslog regex "sshd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)" user==root ip ~ "192.168"
```

AX:

```gravwell
tag=testregex,testcsv ax dstport==80 | table app src dst
```

CSV:

```gravwell
tag=default csv [5]!=stuff [4] [3]~"things" | table 3 4 5
```

IPFIX:

```gravwell
tag=ipfix ipfix destinationIPv4Address as Dst destinationTransportPort==443 | count by Dst | chart count by Dst
```

IP module:

```gravwell
tag=csv csv [2] as srcip | ip srcip ~ PRIVATE
```
