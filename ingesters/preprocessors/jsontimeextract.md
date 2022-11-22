## JSON Time Extraction Preprocessor

The JSON time extraction preprocessor is designed to parse a JSON object and perform time resolution on a specific field.  This preprocessor is particularly useful when a JSON record has multiple timestamps in it and a user needs to specific a specific field for use as the entry timestamp.

If the specified extraction is not present in an entry or the extracted field cannot be processed as a timestamp the entries timestamp is left unchanged.

The JSON Time Extraction preprocessor Type is `jsontimeextract`.

### Supported Options

* `Path` (string, required): This specifies the field to be extracted from the JSON. Given an input of `{"foo":"a", "bar":2, "baz":{"frog": "2022-12-31T12:00:00Z"}}`, you could specify `Path=baz.frog` to extract and process the timestamp `2022-12-31T12:00:00Z`.
* `Assume-Local-Timezone` (Boolean, optional): By default, the timestamp processing code will assume UTC time zones if a timestamp does not contain a timezone.  This option can force the processing module to use the local machines timezone when no timezone is present in the timestamp.
* Timestamp-Override (string, optional): Allows for manually specifying a timestamp format to use when processing extracted timestamps.

### Common Use Cases

Many JSON data sources may provide a few timestamps or contain raw data with timestamps in it.  The timestamp processor TimeGrinder will attempt to find the first timestamp it can in a data record and then lock on to that timestamp and/or format.  Depending on the data schema this may not be the correct timestamp.  This preprocessor allows for treating the data record as JSON and then manually specifying the field containing a timestamp.

### Example: Correcting Timestamp extraction on JSON Data Records

Assuming the following data in an event:

```
{"transmitted":"2022-12-31T12:13:14.12345Z","ID":12345,"record":{"user":"bob","created":"2022-12-30T11:54:21Z"}}
```

The TimeGrinder would most likely zero in on the timestamp `2022-12-31T12:13:14.12345Z` but we probably actually want the timestamp `2022-12-30T11:54:21Z`.

The following preprocessor configuration would extract the created field using the path `record.created` and process it using TimeGrinder, correcting the entry timestamp.

```
[Preprocessor "json-timestamp"]
	Type=jsontimeextract
	Path="record.created"
```
