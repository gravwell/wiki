# Ingesting Entries via the Webserver API

The webserver provides an API for ingesting entries directly. This is useful when running automated scripts or to allow users to import files.

## Ingesting JSON-formatted entries

Automated scripts construct their own entries and ship them as a JSON array of structs. This has the advantage of allowing arbitrary bytes in the Data fields. The essential fields of each entry are:

* TS: a timestamp (e.g. "2018-02-12T11:06:44.215431364-07:00")
* Tag: the string tag to be used (e.g. "syslog")
* Data: base64-encoded bytes (e.g. "Zm9vCg==")

Thus to ship two entries, one containing "foo" and one containing "test", to the tag `mytag`, do a PUT to `/api/ingest/json` with the following data:

```
[ { "TS": "2018-02-12T11:06:44.215431364-07:00", "Tag": "mytag", "Data": "Zm9v" }, { "TS": "2018-02-12T11:06:45.215431364-07:00", "Tag": "mytag", "Data": "dGVzdA==" } ]
```

The server will return the number of entries ingested.

## Ingesting line-delimited files

To ingest user-provided data, the line-delimited file API is the simplest option. It consists of doing a multipart POST to `/api/ingest/lines` with the following form parts:

* A file part named `file` containing the user-provided file
* A field named `tag` containing the desired ingest tag
* An optional field `source`; setting this will override the source value on entries.  This field must be a properly formed IPv4 or IPv6 address.
* An optional field `noparsetimestamp`; setting this to "true" will force entries to be ingested with the current timestamp rather than attempting to parse one from each entry.
* An optional field `assumelocaltimezone`; setting this to "true" means timestamps extracted from entries will assume to be in the local timezone (instead of UTC) if the timezone is not explicitly specified.

The uploaded file will be split by newlines. Each line will be ingested as an individual entry.
 
