## Langfind

The langfind module is a basic human language analysis module that searches data for text and attempts to classify that text as a human language.

### Example Search

The following search will build a table of the most common languages used in Reddit comments, in descending order from most popular to least.

```
tag=reddit json Body | langfind -e Body | count by lang | sort by count desc | table lang count
```

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record. For example, a pipeline that performed language analysis on HTTP payloads would be `tag=pcap ipv4.DstPort==80 tcp.Payload | langfind -e Payload`.
* By default, the output is generated in the enumerated value "lang". Optionally, you can specify the enumerated value name as the last argument. For example, to generate the output in the enumerated value "foo":

```
tag=reddit json Body | langfind -e Body foo | table foo
```
