# Entropy

The `entropy` module calculates the byte entropy of data, normalized to a value between 0 and 1, and stores it in an enumerated value named `entropy`. With no arguments, it calculates the entropy of all entries' DATA fields across the search range. It can instead operate on an enumerated value, group by enumerated values, and supports temporal search mode, so entropy can be charted over time.

Syntax:

```
entropy [-p] [enumerated value] [by ...] [over <duration>]
```

The optional enumerated value names the data to measure; without it, `entropy` uses the DATA field. Entries that lack the named enumerated value are ignored. To group results, list one or more enumerated values after the `by` keyword. For example, to calculate entropy on the enumerated value `foo`, grouped by `bar` and `baz`:

```gravwell
tag=gravwell entropy foo by bar baz
```

The `over` keyword groups results into time windows of the given duration:

```gravwell
tag=gravwell entropy over 10m
```

All arguments are optional.

## Supported Options

* `-p`: Calculate entropy for each entry individually rather than across many entries. Each entry passes through with its own `entropy` value attached; entries that lack the named enumerated value are dropped. `by` is ignored in this mode.

## Examples

This query calculates and charts the entropy of TCP packet payloads based on port:

```gravwell
tag=pcap packet tcp.Port tcp.Payload | entropy Payload by Port | chart entropy by Port
```

This query calculates the entropy of URLs by host and lists the results:

```gravwell
tag=pcap packet tcp.Port==80 ipv4.IP !~ 10.0.0.0/8 tcp.Payload | grep -e Payload GET PUT HEAD POST | regex -e Payload "[A-Z]+\s(?P<url>\S+)\sHTTP\/" | entropy url by IP | table IP entropy
```

This query attaches the entropy of each entry's DATA field and keeps only high-entropy entries:

```gravwell
tag=gravwell entropy -p | eval entropy > 0.7
```
