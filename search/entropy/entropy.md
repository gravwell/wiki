## Entropy

The `entropy` module calculates the entropy of field values over time. Specifying `entropy` without any arguments will generate the entropy of all entries DATA fields across the search range. The `entropy` module supports temporal search mode allowing for charting of entropy over time. `entropy` can also operate on enumerated values and group by enumerated values. Output values are between 0 and 1.

Syntax: 

```
entropy [enumerated value] [by ...] [over <duration>]
```

The `entropy` module syntax allows for specifying an enumerated value to calculate entropy over. If not specified, `entropy` will calculate entropy over the entire DATA field. The module also supports specifying one or more arguments to group by, using the `by` keyword. For example, to calcaulte entropy on the enumerated value `foo`, grouped by `bar` and `baz`:

```
tag=gravwell entropy foo by bar baz
```

Queries can be temporally grouped over arbitrary time windows using the `over` keyword:

```
tag=gravwell entropy over 10m
```

All arguments are optional.

### Supported Options

`entropy` has no flags.

### Examples

This query calculates and charts the entropy of TCP packet payloads based on port:

```
tag=pcap packet tcp.Port tcp.Payload | entropy Payload by Port | chart entropy by Port
```

An example query which calculates the entropy of URLS by host and sorts the list based on highest entropy value:

```
tag=pcap packet tcp.Port==80 ipv4.IP !~ 10.0.0.0/8 tcp.Payload | grep -e Payload GET PUT HEAD POST | regex -e Payload "[A-Z]+\s(?P<url>\S+)\sHTTP\/" | entropy url by IP | table IP entropy
```
