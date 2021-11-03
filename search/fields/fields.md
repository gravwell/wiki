## Fields

The fields module is used to extract and filter data from search entries into enumerated values for later use.  The fields module is designed to be extremely flexible in capturing and filtering data where data items are delimited by a constant set of bytes.  Formats that are comma delimited (CSV), tab delimited, or space delimited are easily processed using the fields module.  More complicated structures which multi-byte delimiters and/or binary formats with known fields separators can use the fields system to extract on arbitrary field boundaries.  Example data producers that can benefit from the fields module are [bro](https://www.bro.org/) with its tab delimited format or the CSV output format from snort.

### Specifying Extraction Fields

Fields are extracted by specifying an index into data from a base of zero.  An index is specified using a positive integer surrounded by square brackets.  Multiple fields can be extracted by providing multiple directives.  Field extraction indexes do not need be be specified in order.

Extracted index fields can be renamed by appending the directive `as <name>` immediately after a field index value.  For example, to extract the 6th field from a piece of data into an enumerated value with the name "uri" the extraction directive would be `[5] as uri`.  If no rename directive is provided the extracted values are given the name that matches the index.  Extracted fields also support filters which allows for quickly filtering entries based on equality or contained values.  Filters must be specified before the renaming statement.  An example fields directive which only allows entries to pass by where the 1st field is the value "stuff" would be `[0]=="stuff"`.  To only allow entries where the 1st field does not equal the value "stuff" and rename the 1st field to "things" the directive would be `[0] != "stuff" as things`.

Attention: Field extraction indexes can be specified as base 10, base 8, or base 16.  The default name applied is the original text value of the index.  An extraction directive of [0xA] will extract the 11th field with the name "0xA", while [010] will extract the 9th field and apply the name "010".

Attention: To specify filter values and or extraction names which contain special characters like "-", ".", or spaces surround the value in double quotes.

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.
* `-d <arg>` : The “-d” option specifies the delimiter used to extract fields.  A delimiter can be any string of bytes.  The default is a comma: ",".
* `-s` : The “-s” option specifies that the fields module should operate in a strict mode.  If any field specification cannot be met, the entry is dropped.  For example if you want the 0th, 1st, and 2nd field but an entry only has 2 fields the strict flag will cause the entry to be dropped.
* `-q` : The “-q” option specifies that the fields can be quoted.  This is useful when dealing with delimiters which might show up in fields.  For example, if the field delimiter is a space, columns may need to contain a space and will be quoted.  If the "-q" argument is specified, any delimiter that is surrounded boy double quotes will be ignored and included in the field.  Delimiters cannot contain double quotes when using the "-q" flag.
* `-clean` : The “-clean” flag specifies that the fields module should remove all surrounding whitespace from extracted fields.  Data formats like CSV which may have trailing whitespace can use the "-clean" flag to remove the unwanted whitespace.  If the "-q" flag is specified with "-clean" double quotes will be removed from quoted fields.

### Filtering Operators

The fields module allows for a filtering based on equality.  If a filter is enabled that specifies equality ("equal", "not equal", "contains", "not contains") any entry that fails the filter specification will be dropped entirely.  If a field is specified as not equal "!=" and the field does not exist, the field is not extracted but the entry won't be dropped entirely.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

### Examples

Extract the URL field from a tab delimited bro http.log feed and name it "url".

```
tag=brohttp fields -d "\t" [9] as url
```

Extract the URL and requester field from a tab delimited bro http.log feed and filter for only entries where the URL contains a space and outputting the results in a table.

```
tag=brohttp fields -d "\t" [9] ~ " " as url [2] as requester | table url requester
```

Extract the 4th, 5th, and 6th fields using a delimiter of "|" and clean white space from extracted fields.

```
tag=default fields -clean -d "|" [3] [4] [5] | table 3 4 5
```

Extract a URI from a bro http log stream and separate the URI into a path and PUT arguments components then calculate the entropy of the args for each path and chart the results.

```
tag=brohttp fields -d "\t" [9] ~ "?" as uri |  regex -e uri "^(?P<path>[^\?;]+)\?(?P<args>.+)" | entropy args by path | chart entropy by path
```

Find HTTP packets with JPEG structures that have more than one stream (e.g. main image and thumbnail)

```
tag=pcap packet tcp.Port==80 tcp.Payload | fields -s -e Payload -d "\xd8\xff" [1]~"JFIF"  [2]~"JFIF" | slice 1[0:10] 2[0:10] | table 1, 2
```
