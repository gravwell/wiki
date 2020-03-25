# Ingest Preprocessors

Sometimes, ingested data needs some additional massaging before we send it to the indexer. Maybe you're getting JSON data sent over syslog and would like to strip out the syslog headers. Maybe you're getting gzip-compressed data from an Apache Kafka stream. Maybe you'd like to be able to route entries to different tags based on the contents of the entries. Ingest preprocessors make this possible by inserting one or more processing steps before the entry is sent up to the indexer.

## Preprocessor Data Flow

An ingester reads raw data from some source (a file, a network connection, an Amazon Kinesis stream, etc.) and splits that incoming data stream out into individual entries. Before those entries are sent to a Gravwell indexer, they may optionally be passed through an arbitrary number of preprocessors as shown in the diagram below.

![](arch.png)

Each preprocessor will have the opportunity to modify the entries. The preprocessors will always be applied in the same order, meaning you could e.g. uncompress the entry's data, then modify the entry tag based on the uncompressed data.

## Configuring Preprocessors

Preprocessors can be used with the following ingesters:

* Simple Relay
* Kinesis
* Kafka
* Google Pub/Sub
* Office 365
* HTTP

The other Gravwell ingesters will receive preprocessor support soon.

Preprocessors are configured in the ingester's config file. Consider the following example for the Simple Relay ingester:

```
[Global]
Ingester-UUID="e985bc57-8da7-4bd9-aaeb-cc8c7d489b42"
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=true
Cleartext-Backend-target=127.0.0.1:4023 #example of adding a cleartext connection
Log-Level=INFO

[Listener "default"]
        Bind-String="0.0.0.0:7777" #we are binding to all interfaces, with TCP implied
        Tag-Name=default
        Preprocessor=timestamp

[Listener "syslog"]
	Bind-String="0.0.0.0:601" # TCP syslog
	Tag-Name=syslog

[preprocessor "timestamp"]
        type = regextimestamp
	Regex ="(?P<badtimestamp>.+) MSG (?P<goodtimestamp>.+) END"
	TS-Match-Name=goodtimestamp
	Timezone-Override=US/Pacific
```

This configuration defines two data consumers (Simple Relay calls them "Listeners") named "default" and "syslog". It also defines a preprocessor named "timestamp". Note how the "default" listener includes the option `Preprocessor=timestamp`. This specifies that entries coming from that listener on port 7777 should be sent to the "timestamp" preprocessor. Because the "syslog" listener does not set any `Preprocessor` option, entries coming in on port 601 will not go through any preprocessors.

## gzip Preprocessor

The gzip preprocessor can uncompress entries which have been compressed with the GNU 'gzip' algorithm.

### Supported Options

* `Passthrough-Non-Gzip` (boolean, optional): if set to true, the preprocessor will pass through any entries whose contents cannot be uncompressed with gzip. By default, the preprocessor will drop any entries which are not gzip-compressed.

### Common Use Cases

Many cloud data bus providers will ship entries and/or package in a compressed form.  This preprocessor can decompress the data stream in the ingester rather than routing through a cloud lambda function can incur costs.


### Example: Decompressing compressed entries

Example config:

```
[Preprocessor "gz"]
	Type=gzip
	Passthrough-Non-Gzip=true
```

## JSON Extraction Preprocessor

The JSON extraction preprocessor can parse the contents of an entry as JSON, extract one or more fields from the JSON, and replace the entry contents with those fields. This is a useful way to simplify overly-complex messages into more concise entries containing only the information of interest.

If only a single field extraction is specified, the result will contain purely the contents of that field; if multiple fields are specified, the preprocessor will generate valid JSON containing those fields.

### Supported Options

* `Extractions` (string, required): This specifies the field or fields (comma-separated) to be extracted from the JSON. Given an input of `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}`, you could specify `Extractions=foo`, `Extractions=foo,bar`, `Extractions=baz.frog,foo`, etc.
* `Force-JSON-Object` (boolean, optional): By default, if a single extraction is specified the preprocessor will replace the entry contents with the contents of that extension; thus selecting `Extraction=foo` will change an entry containing `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}` to simply contain `a`. If this option is set, the preprocessor will always output a full JSON structure, e.g. `{"foo":"a"}`.
* `Passthrough-Misses` (boolean, optional): If set to true, the preprocessor will pass along entries for which it was unable to extract the requested fields. By default, these entries are dropped.
* `Strict-Extraction` (boolean, optional): By default, the preprocessor will pass along an entry if at least one of the extractions succeeds. If this parameter is set to true, it will require that all extractions succeed.

### Common Use Cases

Many data sources may provide additional metadata related to transport and/or storage that are not part of the actual log stream.  The jsonextract preprocessor can downselect fields to reduce storage costs.

### Example: Condensing JSON Data Records

```
[Preprocessor "json"]
	Type=jsonextract
	Extractions=IP,Alert.ID,Message
	Passthrough-Misses=true
```

## JSON Array Split Preprocessor

This preprocessor can split an array in a JSON object into individual entries. For example, given an entry which contains an array of names, the preprocessor will instead emit one entry for each name. Thus this:

```
{"IP": "10.10.4.2", "Users": ["bob", "alice"]}
```

Becomes two entries, one containing "bob" and one containing "alice".

### Supported Options

* `Extraction` (string, required): specifies the JSON field containing a struct which should be split, e.g. `Extraction=Users`, `Extraction=foo.bar`.
* `Passthrough-Misses` (boolean, optional): If set to true, the preprocessor will pass along entries for which it was unable to extract the requested field. By default, these entries are dropped.
* `Force-JSON-Object` (boolean, optional): By default, the preprocessor will emit entries with each containing one item in the list and nothing else; thus extracting `foo` from `{"foo": ["a", "b"]}` would result in two entries containing "a" and "b" respectively. If this option is set, that same entry would result in two entries containg `{"foo": "a"}` and `{"foo": "b"}`.

### Common Use Cases

Many data providers may pack multiple events into a single entry which can degrade the atomic nature of an event and increase the complexity of analysis.  Splitting a single message that contains multiple events into individual entries can simplify working with the events.


### Example: Splitting Multiple Messages In a Single Record

```
[preprocessor "json"]
	Type=jsonarraysplit
	Extraction=Alerts
	Force-JSON-Object=true
```


## JSON Field Filtering Preprocessor

This preprocessor will parse entry data as a JSON object, then extract specified fields and compare them against lists of acceptable values. The lists of acceptable values are specified in files on the disk, one value per line.

It can be configured to either *pass* only those entries whose fields match the lists, or to *drop* those entries which match the lists--whitelisting, or blacklisting. It can be set up to filter against multiple fields, requiring either that *all* fields must match (logical AND) or that *at least one* field must match (logical OR).

This preprocessor is particularly useful to narrow down a firehose of general data before sending it across a slow network link.

### Supported Options

* `Field-Filter` (string, required): This specifies two things: the name of the JSON field of interest, and the path to a file which contains values to match against. For example, one might specify `Field-Filter=ComputerName,/opt/gravwell/etc/computernames.txt` in order to extract a field named "ComputerName" and compare it against values in `/opt/gravwell/etc/computernames.txt`. The `Field-Filter` option may be specified multiple times in order to filter against multiple fields.
* `Match-Logic` (string, optional): This parameter specifies the logic operation to use when filtering against multiple fields. If set to "and", an entry is only considered a match when *all* specified fields match against the given lists. If set to "or", an entry is considered a match when *any* field matches.
* `Match-Action` (string, optional): This specifies the option which should be take for entries whose fields match the provided lists. It may be set to "pass" or "drop"; if omitted, the default is "pass". If set to "pass", entries which match will be allowed to pass to the indexer (whitelisting). If set to "drop", entries which match will be dropped (blacklisting).

The `Match-Logic` parameter is only necessary when more than one `Field-Filter` has been specified.

Note: If a field is specified in the configuration but is not present on an entry, the preprocessor will treat the entry *as if the field existed but did not match anything*. Thus, if you have configured the preprocessor to only pass those entries whose fields match your whitelist, an entry which lacks one of the fields will be dropped.

### Common Use Cases

The json field filtering preprocessor can downselect entries based on fields within the entries.  This allows for building blacklists and whitelists on data flows to ensure that data either does or does not make it to storage.

### Example: Simple Whitelisting

Suppose we have an endpoint monitoring solution which is sending thousands of events per second detailing things which are occurring across the enterprise. Due to the high event volume, we may decide we only want to index events with a certain severity. Luckily, the events include a Severity field:

```
{ "EventID": 1337, "Severity": 8, "System": "email-server-01.example.org", [...] }
```

We know the Severity field goes from 0 to 9, and we decide we want to only pass events with a severity of 6 or higher. We would therefore add the following to our ingester configuration file:

```
[preprocessor "severity"]
	Type=jsonfilter
	Match-Action=pass
	Field-Filter=Severity,/opt/gravwell/etc/severity-list.txt
```

and set `Preprocessor=severity` on the appropriate data input, for instance if we were using Simple Relay:

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
```

Finally, we create `/opt/gravwell/etc/severity-list.txt` and populate it with a list of acceptable Severity values, one per line:

```
6
7
8
9
```

After restarting the ingester, it will extract the `Severity` field from each entry and compare the resulting value against those listed in the file. If the value matches a line in the file, the entry will be sent to the indexer. Otherwise, it will be dropped.

### Example: Blacklisting

Building on the previous example, we may find that that our endpoint monitoring system is generating a *lot* of high-severity false positives from certain systems. We may determine that events with the `EventID` field set to 219, 220, or 1338 and the `System` field set to "webserver-prod.example.org" and "webserver-dev.example.org" are always false positives. We can define another preprocessor to get rid of these entries before they are sent to the indexer:

```
[preprocessor "falsepositives"]
	Type=jsonfilter
	Match-Action=drop
	Match-Logic=and
	Field-Filter=EventID,/opt/gravwell/etc/eventID-blacklist.txt
	Field-Filter=System,/opt/gravwell/etc/system-blacklist.txt
```

If we now add this preprocessor to the data input configuration *after* the existing one, the ingester will apply the two filters in order:

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
	Preprocessor=falsepositives
```

Last, we create `/opt/gravwell/etc/eventID-blacklist.txt`:

```
219
220
1338
```

and `/opt/gravwell/etc/system-blacklist.txt`:

```
webserver-prod.example.org
webserver-dev.example.org
```

This new preprocessor extracts the `EventID` and `System` fields from every entry which makes it past the first filter. It then compares them against the values in the files. Because we set `Match-Logic=and`, it considers an entry a match if *both* field values are found in the files. Because we set `Match-Action=drop`, any entry which matches on both fields will be dropped. Thus, an entry with EventID=220 and System=webserver-dev.example.org is dropped, while one with EventID=220 and System=email-server-01.example.org will *not* be dropped.

## Regex Router Preprocessor

The regex router preprocessor is a flexible tool for routing entries to different tags based on the contents of the entries. The configuration specifies a regular expression containing a [named capturing group](https://www.regular-expressions.info/named.html), the contents of which are then tested against user-defined routing rules.

### Supported Options

* `Regex` (string, required): This parameter specifies the regular expression to be applied to the incoming entries. It must contain at least one [named capturing group](https://www.regular-expressions.info/named.html), e.g. `(?P<app>.+)` which will be used with the `Route-Extraction` parameter.
* `Route-Extraction` (string, required): This parameter specifies the name of the named capturing group from the `Regex` parameter which will contain the string used to compare against routes.
* `Route` (string, required): At least one `Route` definition is required. This consists of two strings separated by a colon, e.g. `Route=sshd:sshlogtag`. The first string ('sshd') is matched against the value extracted via regex, and the second string defines the name of the tag to which matching entries should be routed. If the second string is left blank, entries matching the first string *will be dropped*.
* `Drop-Misses` (boolean, optional): By default, entries which do not match the regular expression will be dropped. Setting `Drop-Misses` to true will make the ingester pass along those entries.

### Example: Routing to Tag Based on App Field Value

To illustrate the use of this preprocessor, consider a situation where many systems are sending syslog entries to a Simple Relay ingester. We would like to separate out the sshd logs to a separate tag named `sshlog`. Incoming sshd logs are in old-style BSD syslog format (RFC3164):

```
<29>1 Nov 26 11:26:36 localhost sshd[11358]: Failed password for invalid user administrator from 202.198.122.184 port 49828 ssh2
```

By experimenting with regular expressions, we find that the following is a reasonable regular expression to extract the application name (e.g. sshd) from RFC3164 logs into a capturing group named "app":

```
^(<\d+>)?\d?\s?\S+ \d+ \S+ \S+ (?P<app>[^\s\[]+)(\[\d+\])?:
```

We can apply that regular expression to a preprocessor definition, as shown below:

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
	# Regex: <pri>version Month Day Time Host App[pid]
	Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog
```

Note that the preprocessor defines the regular expression, then calls out the capturing group "app" in the `Route-Extraction` parameter. It then uses the `Route=ssh:sshlog` definition to specify that those entries whose application name matches "sshd" should be routed to the tag "sshlog". We could define additional `Route` parameters as needed, e.g. `Route=apache:apachelog`.

With the above configuration, logs from sshd will be sent to the "sshlog" tag, while all other logs will go straight to the "syslog" tag. We could extract other applications from similarly-formatted syslog entries by adding additional `Route` specifications, but suppose we had some intermingled logs in RFC 5424 format, as shown below?

```
<101>1 2019-11-26T13:24:56.632535-07:00 web01.example.org webservice 21581 - [useragent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3191.0 Safari/537.36"] GET /
```

The regular expression we already have won't extract the application name ("webservice") properly, but we can define a *second* preprocessor and put it in the preprocessor chain after the existing one:

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter
	Preprocessor=rfc5424router

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
	# Regex: <pri>version Month Day Time Host App[pid]
	Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog

[preprocessor "rfc5424router"]
	Type=regexrouter
	Drop-Misses=false
	# Regex: <pri>version Date Host App
	Regex="^<\\d+>\\d? \\S+ \\S+ (?P<app>\\S+)"
	Route-Extraction=app
	Route=webservice:weblog
	Route=apache:weblog
	Route=postfix:		# drop
```

Note that this new preprocessor definition defines routes for the applications named "webservice" and "apache", sending both to the "weblog" tag. Note also that it specifies that logs from the "postfix" application should be *dropped*, perhaps because those logs are already being ingested from another source.

## Regex Timestamp Extraction Preprocessor

Ingesters will typically attempt to extract a timestamp from an entry by looking for the first thing which appears to be a valid timestamp and parsing it. In combination with additional ingester configuration rules for parsing timestamps (specifying a specific timestamp format to look for, etc.) this is usually sufficient to properly extract the appropriate timestamp, but some data sources may defy these straightforward methods. Consider a situation where a network device may send CSV-formatted event logs wrapped in syslog--a situation we have seen at Gravwell!

### Supported Options

* `Regex` (string, required): This parameter specifies the regular expression to be applied to the incoming entries. It must contain at least one [named capturing group](https://www.regular-expressions.info/named.html), e.g. `(?P<timestamp>.+)` which will be used with the `TS-Match-Name` parameter.
* `TS-Match-Name` (string, required): This parameter gives the name of the named capturing group from the `Regex` parameter which will contain the extracted timestamp.
* `Timestamp-Format-Override` (string, optional): This can be used to specify an alternate timestamp parsing format. Refer to the [Go time package's reference time format](https://golang.org/pkg/time/#pkg-constants) for information on how to specify this.
* `Timezone-Override` (string, optional): If the extracted timestamp doesn't contain a timezone, the timezone specified here will be applied. Example: `US/Pacific`, `Europe/Rome`, `Cuba`.
* `Assume-Local-Timezone` (boolean, optional): This option tells the preprocessor to assume the timestamp is in the local timezone if no timezone is included. This is mutually exclusive with the `Timezone-Override` parameter.


### Common Use Cases

Many data streams may have multiple timestamps or values that can easily be interpretted as timestamps.  The regextimestamp preprocessor allows you to force timegrinder to examine a specific timestamp within a log stream.  A good example is a log stream that is transported via syslog using an application that includes it's own timestamp but does not relay that timestamp to the syslog API.  The syslog wrapper will have a well formed timestamp which, but the actual data stream may need to use some internal field for the accurate timestamp.

### Example: Wrapped Syslog Data

```
Nov 25 15:09:17 webserver alerts[1923]: Nov 25 14:55:34,GET,10.1.3.4,/info.php
```

We would like to extract the inner timestamp, "Nov 25 14:55:34", for the TS field on the ingested entry. Because it uses the same format as the syslog timestamp at the beginning of the line, we cannot extract it with clever timestamp format rules. However, the regex timestamp preprocessor can be used to extract it. By specifying a regular expression which captures the desired timestamp in a named submatch, we can extract timestamps from anywhere in an entry. For this entry, the regex `\S+\s+\S+\[\d+\]: (?<timestamp>.+),` should be sufficient to properly extract the desired timestamp.

This config could be used to extract the timestamp shown in the example above:

```
[Preprocessor "ts"]
        Type=regextimestamp
	Regex="\S+\s+\S+\[\d+\]: (?<timestamp>.+),"
	TS-Match-Name=timestamp
	Timezone-Override=US/Pacific
```

## Regex Extraction Preprocessor

It is highly common for transport busses to wrap data streams with additional metadata that may not be pertinent to the actual event.  Syslog is an excellent example where the Syslog header may not provide value to the underlying data and/or may simply complicate the analysis of the data.  The regexextractor preprocessor allows for declaring a regular expression that can extract multiple fields and reform them into a new structure for format.

The regexextraction preprocessor uses named regular expression extraction fields and a template to extract data and then reform it into an output record.  Output templates can contain static values and completely reform the output data if needed.

Templates reference extracted values by name using field definitions similar to bash.  For example, if your regex extracted a field named `foo` you could insert that extraction in the template with `${foo}`.

### Supported Options

* Passthrough-Misses (boolean, optional): This parameter specifies whether the preprocessor should pass the record through unchanged if the regular expression does not match.
* Regex (string, required): This parameter defines the regular expression for extraction
* Template (string, required): This parameter defines the output form of the record.


### Common Use Cases

The regexpreprocessor is most commonly used for stripping un-needed headers from data streams, however it can be used to reform data into easier to process formats.

#### Example: Stripping Syslog Headers

Given the following record, we want to remove the syslog header and ship just the JSON blob.

```
<30>1 2020-03-20T15:35:20Z webserver.example.com loginapp 4961 - - {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}
```

The syslog message contains a well structured JSON blob but the syslog transport adds additional metadata that does not nessarily enhance the record.  We can use the Regex extractor to pull out the data we want and reform it into an easy to use record.

We will use the regex extractor to pull out the data fields and the hostname, we will then use the template to build a new JSON blob with the host inserted.


```
[Listener "logins"]
	Bind-String="0.0.0.0:7777"
	Preprocessor=loginapp

[Preprocessor "loginapp"]
	Type=regexextract
	Regex="\\S+ (?P<host>\\S+) \\d+ \\S+ \\S+ (?P<data>\\{.+\\})$"
	Template="{\"host\": \"${host}\", \"data\": ${data}}"
```

NOTE: Regular expressions often have backslashes to describe character sets, those backslashes must be escaped!

The resulting data is:

```
{"host": "loginapp", "data": {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}}
```

NOTE: Templates can specify multiple fieds constant values.  Extracted fields can be inserted multiple times.
