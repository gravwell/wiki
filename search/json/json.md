## JSON

The json module will extract data from search entries into enumerated values for later use. For example, to find the most prolific reddit posters, the following search extracts the "Author" field from each reddit post into a new enumerated value, then counts the occurrence of each author and puts it into a table:

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

The JSON format is extremely liberal and allows names of all types, in cases where the json name may contain dot "." character it may be desirable to treat the dot as part of the name rather than as a specification for submembers.  For example, this JSON string contains a dot character in a field name:

```
{ "subfield.op": "stuff", "subfield.type": "int", "subfield.value": 99}
```

An example json module argument to extract the subfield.op member would be:

```
json "subfield.op"
```

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.