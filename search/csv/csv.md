# CSV

The `csv` module extracts and filters columns from comma-separated values. It works like an enhanced [fields](/search/fields/fields) module that understands CSV rules: columns can be wrapped in double quotes to contain commas or newlines, and surrounding whitespace is ignored.

## Specifying Extraction Fields

Columns are selected by zero-based index in square brackets. `[0]` is the first column, `[5]` the sixth. List as many columns as you need, in any order.

By default, an extracted column is named after its index, so `[5]` produces an enumerated value named `5`. To rename it, add `as <name>` after the index: `[5] as uri`.

Each column can also take a filter (see Filtering Operators below), which must come before `as`. For example, `[0]=="stuff"` passes only entries whose first column is "stuff", and `[0] != "stuff" as things` drops entries whose first column is "stuff" and names the column `things`.

```{attention}
Indexes are decimal. Leading zeros are allowed, and the default name keeps the original text: `[010]` extracts the 11th column into an enumerated value named `010`.
```

```{attention}
To specify filter values or extraction names that contain spaces or other special characters, surround the value in double quotes.
```

### Naming Columns by Position

Instead of indexes, you can list names for the columns in order, separated by commas or spaces. Leave a position empty to skip that column. For example, the following extracts the first column as `ts`, skips the second and third, and extracts the fourth as `status`:

```
csv ts,,,status
```

This form does not support filters.

## Supported Options

* `-e <arg>`: Operate on the enumerated value `<arg>` instead of the entry data.
* `-s`: Strict mode. Drop the entry if any requested column is missing. For example, if you request `[0] [1] [2]` but an entry has only 2 columns, the entry is dropped. Without `-s`, missing columns are skipped and the entry passes.

## Filtering Operators

An entry that fails any filter is dropped. If a column is missing, its filter is not evaluated and the entry passes; use `-s` to drop such entries.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

### Filtering Examples

```
csv [0] == "foo" [2] != "bar" [3] ~ baz as ID
```

## Data Extraction

The `csv` module always strips surrounding whitespace and double quotes from extracted columns. A backslash does not escape a comma; quote the column instead. For example, consider the following entry:

```
2018-11-01T12:46:01.764386-06:00,daffodil, 15554,9870f7cd-b7d3-4bb6-8160-f5f146ebc764 , "OK, what sort of language would we have the world speak?
Isabella", "CL", Lucien
```

The 3rd, 4th, and 5th columns have surrounding whitespace, and the 5th column is quoted because it contains a comma and a newline. The following query:

```
csv [2] [3] [4] | table 2 3 4
```

produces the following output, without the quotes or surrounding whitespace:

| 2 | 3 | 4 |
|----------|------|-------------|
| 15554 | 9870f7cd-b7d3-4bb6-8160-f5f146ebc764 | OK, what sort of language would we have the world speak?<br>Isabella |


## Example queries

Extract a URL column from a CSV http.log feed and name it "url".

```gravwell
tag=brohttp csv [9] as url
```


Extract the URL and requester fields from a CSV bro http.log feed, keep only entries where the URL contains a space, and show the results in a table.

```gravwell
tag=brohttp csv [9] ~ " " as url [2] as requester | table url requester
```


Extract the 4th, 5th, and 6th columns where the 6th column must not be "stuff" and the 4th column must contain "things".

```gravwell
tag=default csv [5]!=stuff [4] [3]~"things" | table 3 4 5
```


Extract a URI from an apache access log and then parse the "DATA" parameter as a CSV.

```gravwell
tag=apache regex "GET\s(?P<base>[^\s\?]+)\?(?P<params>\S+)\s" | regex -e params "DATA\=(?P<dataparam>[^\s&]+)" | csv -e dataparam [0] [1] [2] | table 0 1 2
```
