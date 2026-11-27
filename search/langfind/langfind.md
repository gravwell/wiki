# Langfind

The langfind module attempts to identify the human language of text. It stores the language as a three-letter ISO 639-3 code, such as `eng`, `spa`, or `cmn`, in an enumerated value named `lang`. To use a different name, give it as the last argument.

Entries are dropped if no language can be detected, or if the text to analyze is empty or missing.

## Supported Options

* `-e <arg>`: Analyze the named enumerated value instead of the DATA field. For example, to analyze HTTP payloads: `tag=pcap packet tcp.DstPort==80 tcp.Payload | langfind -e Payload`.

## Examples

The following search builds a table of the most common languages used in Reddit comments, from most to least popular:

```gravwell
tag=reddit json Body | langfind -e Body | count by lang | sort by count desc | table lang count
```

To store the language in the enumerated value `foo` instead:

```gravwell
tag=reddit json Body | langfind -e Body foo | table foo
```
