## JSON

The json module is used to extract and filter data from search entries into enumerated values for later use.  JSON is an excellent data format for daynmic exploration as the data is self-describing.  The JSON module can extract items and rename them, or filter based on the extracted value.  Filtering directly within the JSON module provides a very high speed and intuitive way to select data of a specific format 

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.
* `-s`: The “-s” option informs the json module that we are in strict mode, meaning that if any item isn't found, drop the entire entry.

### Filtering Operators

The JSON module allows for a filtering based on equality.  If a filter is enabled that specifies equality ("equal" or "not equal") any entry that fails the filter specification will be dropped entirely.  If a field is specified as not equal "!=" and the field does not exist, the field is not extracted but the entry won't be dropped entirely.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

### Examples
To find the most prolific reddit posters, the following search extracts the "Author" field from each reddit post into a new enumerated value, then counts the occurrence of each author and puts it into a table:

```
tag=reddit json Author | count by Author | table Author count
```

The module can also descend multiple layers into the JSON entry. For example, in the Shodan data we ingest for testing, we can extract the "region code" from entries to discover where the endpoint resides. If we want to learn which states have the most AT&T U-verse customers, we can issue the following search:

```
tag=shodan grep "AT&T U-verse" | json location.region_code | count by region_code | table region_code count
```

It can also operate on enumerated values rather than the full entry data if desired; for instance, if an XML entry contains json within it:

```
<System><Data>{ "domain": "gravwell.io" }</Data></System>
```

We can use the following command to extract the JSON from within the XML as an enumerated value named "Data", then apply the json module to parse out the domain value into another enumerated value named "domain":

```
xml System.Data | json -e Data domain
```
Enumerated value names are derived by the last name in a JSON specification, in the earlier example which extracted the region_code field the output is populated in the "region_code" enumerated value.  Output enumerated value names can be overridden with an "as" argument.  The following example extracts the domain member from the Data enumerated value and assignes it into a new enumerated value named "dd":

```
json -e Data domain as dd
```
Using the filter operator we can extract the Data field, but only when the domain is not the value "google.com." Filters can be combined with renaming.

```
json -e Data domain != "google.com" as dd
```

The JSON format is extremely liberal and allows names of all types, including characters Gravwell usually treats as separators such as '.' and "-". In cases where the JSON name contains such characters, wrap the individual field in double-quotes to parse it as a single token. For example, this JSON string contains a dot character in a field name:

```
{ "subfield.op": "stuff", "subfield.type": "int", "subfield.value": 99}
```

An example json module argument to extract the subfield.op member would be:

```
json "subfield.op" as sop
```

Similarly, consider the following nested structure:

```
{ "fields": { "search-id": 1234, "search-type": "background" } }
```

Because search-id and search-type contain a dash character, they should be wrapped in quotes when used:

```
json fields."search-id" fields."search-type" as type | count by "search-id",type | table "search-id" type count
```
