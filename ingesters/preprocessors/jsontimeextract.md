# JSON Time Extraction Preprocessor

The JSON time extraction preprocessor is designed to parse a JSON object and perform time resolution on a specific field.  This preprocessor is particularly useful when a JSON record has multiple timestamps in it and a user needs to specify a particular field for use as the entry timestamp.

If the specified extraction is not present in an entry or the extracted field cannot be processed as a timestamp, the entry timestamp is left unchanged.

The JSON Time Extraction preprocessor Type is `jsontimeextract`.

## Supported Options

* `Path` (string, required): This specifies the field to be extracted from the JSON. Given an input of `{"foo":"a", "bar":2, "baz":{"frog": "2022-12-31T12:00:00Z"}}`, you could specify `Path=baz.frog` to extract and process the timestamp `2022-12-31T12:00:00Z`.
* `Assume-Local-Timezone` (Boolean, optional): By default, the timestamp processing code will assume UTC time zones if a timestamp does not contain a timezone.  This option can force the processing module to use the local machines timezone when no timezone is present in the timestamp.
* `Timestamp-Override` (string, optional): Manually specify a timestamp format to use when processing extracted timestamps.
* `Timestamp-Max-Past-Delta` (string, optional): The maximum age of an extracted timestamp, as a duration such as `720h`. Timestamps older than this are ignored. Defaults to `8760h` (one year).
* `Timestamp-Max-Future-Delta` (string, optional): How far into the future an extracted timestamp may be, as a duration such as `1h`. Timestamps further ahead than this are ignored. Defaults to `24h`.

## Timestamp Bounds

The preprocessor only applies an extracted timestamp that falls within a window around the current time, set by `Timestamp-Max-Past-Delta` and `Timestamp-Max-Future-Delta`. If the extracted timestamp falls outside the window, the entry timestamp is left unchanged. Timestamps before 1970-01-01 are always ignored, whatever the window. Both durations must be non-zero; negative values are treated as positive.

```{warning}
The default window rejects timestamps more than a year old. If you replay archives or backfill older data through this preprocessor, set `Timestamp-Max-Past-Delta` large enough to cover it, or those entries will keep their original timestamp.
```

## Common Use Cases

Many JSON data sources may provide a few timestamps or contain raw data with timestamps in it.  The timestamp processor TimeGrinder will attempt to find the first timestamp it can in a data record and then lock on to that timestamp and/or format.  Depending on the data schema this may not be the correct timestamp.  With this preprocessor, you can tell the ingester to treat the data record as JSON and then manually specify the field containing a timestamp.

## Example: Correcting Timestamp extraction on JSON Data Records

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

## Example: Backfilling Older Data

This configuration accepts extracted timestamps up to five years old:

```
[Preprocessor "json-timestamp"]
	Type=jsontimeextract
	Path="record.created"
	Timestamp-Max-Past-Delta=43800h
```
