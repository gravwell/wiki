## CSV

The csv module is designed to extract and filter data from comma separated values.  The csv module can be thought of as an enhanced fields module that is designed to accommodate some additional rules related to csv values such as the ability to quote columns that may contain commas, or escape commas, or surround columns with white space.

### Specifying Extraction Fields

Fields are extracted by specifying an index into data from a base of zero.  An index is specified using a positive integer surrounded by square brackets.  Multiple columns can be extracted by providing multiple directives.  Column extraction indexes do not need be be specified in order.

Extracted columns can be renamed by appending the directive `as <name>` immediately after a field index value.  For example, to extract the 6th column from a piece of data into an enumerated value with the name "uri" the extraction directive would be `[5] as uri`.  If no rename directive is provided the extracted values are given the name that matches the index.  Extracted columns also support filters which allows for quickly filtering entries based on equality or contained values.  Filters must be specified before the renaming statement.  An example column directive which only allows entries to pass by where the 1st column is the value "stuff" would be `[0]=="stuff"`.  To only allow entries where the 1st column does not equal the value "stuff" and rename the 1st field to "things" the directive would be `[0] != "stuff" as things`.

Attention: Column extraction indexes can be specified as base 10, base 8, or base 16.  The default name applied is the original text value of the index.  An extraction directive of [0xA] will extract the 11th column with the name "0xA", while [010] will extract the 9th column and apply the name "010".

Attention: To specify filter values and or extraction names which contain special characters like "-", ".", or spaces, surround the value in double quotes.

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.
* `-s` : The “-s” option specifies that the csv module operate in a strict mode.  If any column specification cannot be met, the entry is dropped.  For example if you want the 0th, 1st, and 2nd field but an entry only has 2 columns the strict flag will cause the entry to be dropped.

### Filtering Operators

The csv module allows for a filtering based on equality.  If a filter is enabled that specifies equality ("equal", "not equal", "contains", "not contains") any entry that fails the filter specification will be dropped entirely.  If a field is specified as not equal "!=" and the field does not exist, the field is not extracted but the entry won't be dropped entirely.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

#### Filtering Examples

```
csv [0] == "foo" [2] != "bar" [3] ~ baz as ID
```

### Data Extraction

The CSV module will always clean out surrounding whitespace and double quotes from extracted column data.  For example, consider the following entry:

```
2018-11-01T12:46:01.764386-06:00,daffodil, 15554,9870f7cd-b7d3-4bb6-8160-f5f146ebc764 , "OK, what sort of language would we have the world speak?
Isabella", "CL", Lucien
```

The 3rd, 4th, and 5th columns contain surrounding whitespace, and the 5th column contains double quotes that encapsulate commas and new lines.  If we executed the CSV module with the following arguments:

```
csv [2] [3] [4] | table 2 3 4
```

The output would look as follows (notice the lack of quotes or surrounding whitespace:

| 2 | 3 | 4 |
|----------|------|-------------|
| 15554 | 9870f7cd-b7d3-4bb6-8160-f5f146ebc764 | OK, what sort of language would we have the world speak?<br>Isabella |


### Example queries

Extract a URL column from a CSV http.log feed and name it "url".

```
tag=brohttp csv [9] as url
```


Extract the URL and requester field from a CSV bro http.log feed and filter for only entries where the URL contains a space and outputting the results in a table.

```
tag=brohttp csv [9] ~ " " as url [2] as requester | table url requester
```


Extract the 4th, 5th, and 6th columns where the 6th column must not be "stuff" and the 4th column must contain "things.

```
tag=default csv [5]!=stuff [4] [3]~"things" | table 3 4 5
```


Extract a URI from an apache access log and then parse the "DATA" parameter as a CSV.

```
tag=apache tag=apache regex "GET\s(?P<base>[^\s\?]+)\?(?P<params>\S+)\s" | regex -e params "DATA\=(?P<dataparam>[^\s&]+)" | csv -e dataparam [0] [1] [2] | table 0 1 2
```
