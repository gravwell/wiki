## Regex Extraction Preprocessor

It is common for transport buses to wrap data streams with additional metadata that may not be pertinent to the actual event.  Syslog is an excellent example, where the syslog header may not provide value to the underlying data, duplicates fields in the underlying data, or simply complicates the analysis of the data.  The regexextractor preprocessor uses a regular expression to extract multiple fields and reform them into a new structure for ingestion.

The regexextraction preprocessor uses named regular expression extraction fields and a template to extract data and then reform it into an output record.  Output templates can contain static values and completely reform the output data if needed.

Templates reference extracted values by name using field definitions similar to bash.  For example, if your regex extracted a field named `foo` you could insert that extraction in the template with `${foo}`. The templates also support the following special keys:

* `${_SRC_}`, which will be replaced by the SRC field of the current entry.
* `${_DATA_}`, which will be replaced by the string-formatted Data field of the current entry.
* `${_TS_}`, which will be replaced by the string-formatted TS (timestamp) field of the current entry.

The Regex Extraction preprocessor Type is `regexextract`.

### Supported Options

* Drop-Misses (boolean, optional): This parameter specifies whether the preprocessor should drop the entry if the regular expression does not match. By default the entry is passed through.
* Regex (string, required): This parameter defines the regular expression for extraction
* Template (string, required): This parameter defines the output form of the record.

### Common Use Cases

The regexpreprocessor is most commonly used for stripping un-needed headers from data streams, however it can also be used to reform data into easier to process formats.

#### Example: Stripping Syslog Headers

Given the following record, we want to remove the syslog header and ship just the JSON blob. We will slightly restructure the JSON, though, to include the application name from the syslog header.

```
<30>1 2020-03-20T15:35:20Z webserver.example.com loginapp 4961 - - {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}
```

The syslog message contains a well structured JSON blob but the syslog transport adds additional metadata that does not necessarily enhance the record.  We can use the Regex extractor to pull out the data we want and reform it into an easy to use record.

We will use the regex extractor to pull out the data fields and the hostname, then use the template to build a new JSON blob with the host inserted:


```
[Listener "logins"]
	Bind-String="0.0.0.0:7777"
	Preprocessor=loginapp

[Preprocessor "loginapp"]
	Type=regexextract
	Regex=`\S+ (?P<host>\S+) \d+ \S+ \S+ (?P<data>\{.+\})$`
	Template=`{"host": "${host}", "data": ${data}}`
```

NOTE: Observe the use of backticks in the Regex and Template fields. This eliminates the need to escape each backslash within the value.

The result is:

```
{"host": "loginapp", "data": {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}}
```

NOTE: Templates can specify multiple fields constant values.  Extracted fields can be inserted multiple times.

