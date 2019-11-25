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

The gzip preprocessor supports the following options in its configuration block:

* `Passthrough-Non-Gzip` (boolean, optional): if set to true, the preprocessor will pass through any entries whose contents cannot be uncompressed with gzip. By default, the preprocessor will drop any entries which are not gzip-compressed.

Example config:

```
[Preprocessor "gz"]
	Type=gzip
	Passthrough-Non-Gzip=true
```

## JSON Extraction Preprocessor

The JSON extraction preprocessor can parse the contents of an entry as JSON, extract one or more fields from the JSON, and replace the entry contents with those fields. This is a useful way to simplify overly-complex messages into more concise entries containing only the information of interest.

If only a single field extraction is specified, the result will contain purely the contents of that field; if multiple fields are specified, the preprocessor will generate valid JSON containing those fields.

The following configuration options are available:

* `Extractions` (string, required): This specifies the field or fields (comma-separated) to be extracted from the JSON. Given an input of `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}`, you could specify `Extractions=foo`, `Extractions=foo,bar`, `Extractions=baz.frog,foo`, etc.
* `Force-JSON-Object` (boolean, optional): By default, if a single extraction is specified the preprocessor will replace the entry contents with the contents of that extension; thus selecting `Extraction=foo` will change an entry containing `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}` to simply contain `a`. If this option is set, the preprocessor will always output a full JSON structure, e.g. `{"foo":"a"}`.
* `Passthrough-Misses` (boolean, optional): If set to true, the preprocessor will pass along entries for which it was unable to extract the requested fields. By default, these entries are dropped.
* `Strict-Extraction` (boolean, optional): By default, the preprocessor will pass along an entry if at least one of the extractions succeeds. If this parameter is set to true, it will require that all extractions succeed.

Example config:

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

The following configuration options are recognized:

* `Extraction` (string, required): specifies the JSON field containing a struct which should be split, e.g. `Extraction=Users`, `Extraction=foo.bar`.
* `Passthrough-Misses` (boolean, optional): If set to true, the preprocessor will pass along entries for which it was unable to extract the requested field. By default, these entries are dropped.
* `Force-JSON-Object` (boolean, optional): By default, the preprocessor will emit entries with each containing one item in the list and nothing else; thus extracting `foo` from `{"foo": ["a", "b"]}` would result in two entries containing "a" and "b" respectively. If this option is set, that same entry would result in two entries containg `{"foo": "a"}` and `{"foo": "b"}`.

Example config:

```
[preprocessor "json"]
		Type=jsonarraysplit
		Extraction=Alerts
		Force-JSON-Object=true
```

## Regex Timestamp Extraction Preprocessor

Ingesters will typically attempt to extract a timestamp from an entry by looking for the first thing which appears to be a valid timestamp and parsing it. In combination with additional ingester configuration rules for parsing timestamps (specifying a specific timestamp format to look for, etc.) this is usually sufficient to properly extract the appropriate timestamp, but some data sources may defy these straightforward methods. Consider a situation where a network device may send CSV-formatted event logs wrapped in syslog--a situation we have seen at Gravwell!

```
Nov 25 15:09:17 webserver alerts[1923]: Nov 25 14:55:34,GET,10.1.3.4,/info.php
```

We would like to extract the inner timestamp, "Nov 25 14:55:34", for the TS field on the ingested entry. Because it uses the same format as the syslog timestamp at the beginning of the line, we cannot extract it with clever timestamp format rules. However, the regex timestamp preprocessor can be used to extract it. By specifying a regular expression which captures the desired timestamp in a named submatch, we can extract timestamps from anywhere in an entry. For this entry, the regex `\S+\s+\S+\[\d+\]: (?<timestamp>.+),` should be sufficient to properly extract the desired timestamp.

The following configuration options are supported:

* `Regex` (string, required): This parameter specifies the regular expression to be applied to the incoming entries. It must contain at least one [named capturing group](https://www.regular-expressions.info/named.html), e.g. `(?P<timestamp>.+)` which will be used with the `TS-Match-Name` parameter.
* `TS-Match-Name` (string, required): This parameter gives the name of the named capturing group from the `Regex` parameter which will contain the extracted timestamp.
* `Timestamp-Format-Override` (string, optional): This can be used to specify an alternate timestamp parsing format. Refer to the [Go time package's reference time format](https://golang.org/pkg/time/#pkg-constants) for information on how to specify this.
* `Timezone-Override` (string, optional): If the extracted timestamp doesn't contain a timezone, the timezone specified here will be applied. Example: `US/Pacific`, `Europe/Rome`, `Cuba`.
* `Assume-Local-Timezone` (boolean, optional): This option tells the preprocessor to assume the timestamp is in the local timezone if no timezone is included. This is mutually exclusive with the `Timezone-Override` parameter.

This config could be used to extract the timestamp shown in the example above:

```
[Preprocessor "ts"]
        Type=regextimestamp
		Regex="\S+\s+\S+\[\d+\]: (?<timestamp>.+),"
		TS-Match-Name=timestamp
		Timezone-Override=US/Pacific
```
