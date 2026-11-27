# Limit

The limit module passes a set number of entries and drops the rest. It is handy while building a query, for example when testing a regular expression: adding `limit 50` keeps the results manageable.

The syntax is `limit <n>`, which passes the first `n` entries, or `limit <x> <y>`, which passes a range of entries.

`limit x y` passes entries `[x,y)`: the range includes `x` and excludes `y`. Entries are numbered from zero, so given `[a,b,c,d,e,f]`, `limit 2 5` returns `[c d e]`. `y` must be greater than `x`.

For example, to look at the payload of 10 packets:

```gravwell
tag=pcap packet tcp.Payload | limit 10 | table Payload
```

To skip the first 5 packets and look at the next 5:

```gravwell
tag=pcap packet tcp.Payload | limit 5 10 | table Payload
```

Limit can also key on enumerated values. In that case, the limit or range applies separately to each value. To allow 5 entries for each value of the enumerated value `foo`:

```gravwell
tag=default ax | limit 5 by foo
```

You can specify any number of keys. To allow 5 entries for each combination of `foo`, `bar`, and `baz`:

```gravwell
tag=default ax | limit 5 by foo bar baz
```

When keying, entries that have none of the named values are dropped.
