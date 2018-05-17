## Limit

The limit module allows a specified number of entries through and no more. This may be especially useful during the process of building a query, for example while testing regular expressions; by inserting a `limit 50` into the pipeline, the results displayed will be less overwhelming.

The syntax is simple: `limit <n>`, where `n` is the maximum number of entries to allow through. To look at the payload of 10 packets:

```
tag=pcap packet tcp.Payload | limit 10 | table Payload
```