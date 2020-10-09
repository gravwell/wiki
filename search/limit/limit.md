## Limit

The limit module allows a specified number of entries through and no more. This may be especially useful during the process of building a query, for example while testing regular expressions; by inserting a `limit 50` into the pipeline, the results displayed will be less overwhelming.

The syntax is simple: `limit <n>`, where `n` is the maximum number of entries to allow through, or `limit <n> <m>`, which allows the Nth to the Mth entries through.

Specifically, `limit X Y` will pass entries `[X,Y)`. That is limit is inclusive of the first term, and exclusive of the second. Terms are also zero-indexed, meaning given a set `[a,b,c,d,e,f]`, `limit 2 5` will return `[c d e]`.

For example, to look at the payload of 10 packets:

```
tag=pcap packet tcp.Payload | limit 10 | table Payload
```

To look at packets 5 to 10:

```
tag=pcap packet tcp.Payload | limit 5 10 | table Payload
```

