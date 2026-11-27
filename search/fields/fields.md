# Fields

The fields module extracts and filters fields from data split by a fixed delimiter, such as comma-separated (CSV), tab-separated, or space-separated text. The delimiter can be any string of bytes, so the module also handles multi-byte delimiters and binary formats with known separators. Examples include the tab-separated logs from [bro](https://www.bro.org/) and the CSV output from snort.

## Specifying Extraction Fields

Specify each field by its zero-based index in square brackets, such as `[0]` for the first field. You can extract several fields, in any order, but each index only once.

By default, the enumerated value is named after the index text, so `[5]` produces an enumerated value named `5`. To rename it, put `as <name>` after the index: `[5] as uri` extracts the 6th field into `uri`. A filter goes between the index and the rename. `[0]=="stuff"` only passes entries whose first field is `stuff`, and `[0] != "stuff" as things` passes entries whose first field is not `stuff` and names that field `things`.

```{attention}
Indexes are decimal; an index that isn't a valid decimal number is read as hexadecimal, with or without a `0x` prefix. The default name is the index exactly as written. `[0xA]` extracts the 11th field and names it `0xA`. Leading zeros don't mean octal: `[010]` is decimal 10, so it also extracts the 11th field, named `010`.
```

```{attention}
Wrap filter values and names that contain special characters, such as `-`, `.`, or spaces, in double quotes.
```

### Naming fields by position

Instead of indexes, you can list names for the fields in order, separated by commas or spaces. Each name applies to the field at that position. With commas, leave a name empty to skip a field. This form doesn't support filters or `as`.

```gravwell
tag=default fields -d " " ts,host,,status
```

This extracts the 1st, 2nd, and 4th fields into `ts`, `host`, and `status`, and skips the 3rd.

## Supported Options

* `-e <arg>`: Operate on an enumerated value instead of the entire entry.
* `-d <arg>`: The delimiter between fields. It can be any non-empty string of bytes. Default is a comma (`,`).
* `-s`: Strict mode. If any extraction fails, drop the entry. For example, if you extract `[0] [1] [2]` but an entry only has 2 fields, the entry is dropped.
* `-q`: Allow quoted fields. Delimiters inside double quotes are treated as part of the field, and the surrounding quotes are removed from the extracted value. Use this when fields may contain the delimiter, such as space-delimited data with quoted strings. Delimiters cannot contain double quotes when using `-q`.
* `-clean`: Strip leading and trailing spaces, tabs, and newlines from extracted fields. Useful for CSV data with padding around values. With `-q`, whitespace outside the quotes is stripped before the quotes are removed.

## Filtering Operators

The fields module can filter on equality and substring matching. Entries that fail a filter are dropped. If a field filtered with `!=` or `!~` doesn't exist, nothing is extracted but the entry isn't dropped.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

## Examples

Extract the URL field from a tab-delimited bro http.log feed and name it `url`:

```gravwell
tag=brohttp fields -d "\t" [9] as url
```

Extract the URL and requester fields from a bro http.log feed, keep only entries where the URL contains a space, and show the results in a table:

```gravwell
tag=brohttp fields -d "\t" [9] ~ " " as url [2] as requester | table url requester
```

Extract the 4th, 5th, and 6th fields using `|` as the delimiter, and strip whitespace from the extracted values:

```gravwell
tag=default fields -clean -d "|" [3] [4] [5] | table 3 4 5
```

Extract URIs that contain a query string from a bro http log, split each into a path and its arguments, calculate the entropy of the arguments for each path, and chart the results:

```gravwell
tag=brohttp fields -d "\t" [9] ~ "?" as uri |  regex -e uri "^(?P<path>[^\?;]+)\?(?P<args>.+)" | entropy args by path | chart entropy by path
```

Find HTTP packets with JPEG data that has more than one image stream (for example, a main image and a thumbnail):

```gravwell
tag=pcap packet tcp.Port==80 tcp.Payload | fields -s -e Payload -d "\xd8\xff" [1]~"JFIF"  [2]~"JFIF" | slice 1[0:10] 2[0:10] | table 1, 2
```

(namedfields-module)=
## namedfields

The `namedfields` module is `fields` with the indexes looked up by name from a resource. Use it when one delimited format is queried often and remembering column numbers is a chore.

```
namedfields -r <resource> -g <group> [flags] <name> [<name> ...]
```

* `-r <resource>`: Required. The resource holding the name-to-index mappings.
* `-g <group>`: Required. The group (`Set` entry) in the resource to use.
* `-e`, `-s`, `-q`, `-clean`: Same as for `fields`.

At least one name is required. Filters and `as` work as they do in `fields`. The delimiter comes from the group, not a flag.

The resource is JSON. Each group sets a delimiter and maps names to zero-based field indexes:

```json
{
  "Version": 1,
  "Set": [
    {
      "Name": "conn",
      "Delim": "\t",
      "Subs": [
        {"Name": "ts", "Index": 0},
        {"Name": "uid", "Index": 1},
        {"Name": "orig_h", "Index": 2},
        {"Name": "resp_h", "Index": 4}
      ]
    }
  ]
}
```

With that resource saved as `zeek`, this query extracts `orig_h` and `resp_h` by name and counts connections to one host:

```gravwell
tag=zeekconn namedfields -r zeek -g conn orig_h resp_h=="10.0.0.1" | count by orig_h | table orig_h count
```

`DataType` entries in a group are accepted but ignored; every value is extracted as a string.
