# Regex Extraction Preprocessor

It is common for transport buses to wrap data streams with additional metadata that may not be pertinent to the actual event.  Syslog is an excellent example, where the syslog header may not provide value to the underlying data, duplicates fields in the underlying data, or simply complicates the analysis of the data.  The regexextract preprocessor uses a regular expression to extract multiple fields and reform them into a new structure for ingestion.

The regexextract preprocessor uses named regular expression extraction fields and a template to extract data and then reform it into an output record.  Output templates can contain static values and completely reform the output data if needed.

Templates reference extracted values by name using field definitions similar to bash.  For example, if your regex extracted a field named `foo` you could insert that extraction in the template with `${foo}`. The templates also support the following special keys:

* `${_SRC_}`, which will be replaced by the SRC field of the current entry.
* `${_DATA_}`, which will be replaced by the string-formatted Data field of the current entry.
* `${_TS_}`, which will be replaced by the string-formatted TS (timestamp) field of the current entry.

The Regex Extraction preprocessor Type is `regexextract`.

## Supported Options

* Drop-Misses (boolean, optional): This parameter specifies whether the preprocessor should drop the entry if the regular expression does not match. By default the entry is passed through.
* Regex (string, required): This parameter defines the regular expression for extraction
* Template (string, required): This parameter defines the output form of the record.
* Attach (string, optional): Name extracted value that will be attached as an intrinsic value.

## Common Use Cases

The regexextract preprocessor is most commonly used for stripping un-needed headers from data streams, however it can also be used to reform data into easier to process formats.

### Example: Stripping Syslog Headers

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

```{note}
Observe the use of backticks in the Regex and Template fields. This eliminates the need to escape each backslash within the value.
```

The result is:

```
{"host": "loginapp", "data": {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}}
```

```{note}
Templates can specify multiple fields constant values.  Extracted fields can be inserted multiple times.
```
### Example: Removing and Attaching Headers

Given the following record, we want to remove the host and category headers and leave only the data field; the host and category will be attached as values to the entry:

```
hostname => [testing.gravwell.io] category => [test values] data => [A very important data item]
```

This log message contains several items and significant formatting. We can use the `regexextract` preprocessor to extract the data field and put it into the body of an entry while extracting the hostname and category and placing them into attached values.

```
[Listener "magic app"]
	Bind-String="0.0.0.0:7777"
	Preprocessor=customapp

[Preprocessor "customapp"]
	Type=regexextract
	Regex=`hostname\s=>\s\[(?P<host>[^\]]+)\]\scategory\s=>\s\[(?P<cat>[^\]]+)\]\sdata\s=>\s\[(?P<data>[^\]]+)\]`
	Template=`$(data)`
	Attach=host
	Attach=cat
```

The result is:

```
A very important data item
```

Two attached values named `host` and `cat` will also be attached to the entry with the values `testing.gravwell.io` and `test values` respectively.

```{note}
Multiple attach directives can be specified, but the specified attach names must be named extractions in the regular expression.
```
