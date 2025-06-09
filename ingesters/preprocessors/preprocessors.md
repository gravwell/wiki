# Ingest Preprocessors

Sometimes, ingested data needs some additional massaging before we send it to the indexer. Maybe you're getting JSON data sent over syslog and would like to strip out the syslog headers. Maybe you're getting gzip-compressed data from an Apache Kafka stream. Maybe you'd like to be able to route entries to different tags based on the contents of the entries. Ingest preprocessors make this possible by inserting one or more processing steps before the entry is sent up to the indexer.

## Preprocessor Data Flow

An ingester reads raw data from some source (a file, a network connection, an Amazon Kinesis stream, etc.) and splits that incoming data stream out into individual entries. Before those entries are sent to a Gravwell indexer, they may optionally be passed through an arbitrary number of preprocessors as shown in the diagram below.

![](arch.png)

Each preprocessor will have the opportunity to modify the entries. The preprocessors will always be applied in the same order, meaning you could e.g. uncompress the entry's data, then modify the entry tag based on the uncompressed data.

## Configuring Preprocessors

Preprocessors are supported on all packaged ingesters.  One-off and unsupported ingesters may not support preprocessors.

Preprocessors are configured in the ingester's config file using the `preprocessor` configuration stanza.  Each Preprocessor stanza must declare the preprocessor module in use via the `Type` configuration parameter, followed by the preprocessor's specific configuration parameters. Consider the following example for the Simple Relay ingester:

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
	Type = regextimestamp
	Regex ="(?P<badtimestamp>.+) MSG (?P<goodtimestamp>.+) END"
	TS-Match-Name=goodtimestamp
	Timezone-Override=US/Pacific
```

This configuration defines two data consumers (Simple Relay calls them "Listeners") named "default" and "syslog". It also defines a preprocessor named "timestamp". Note how the "default" listener includes the option `Preprocessor=timestamp`. This specifies that entries coming from that listener on port 7777 should be sent to the "timestamp" preprocessor. Because the "syslog" listener does not set any `Preprocessor` option, entries coming in on port 601 will not go through any preprocessors.

(ingest_preprocessors_list)=
## Available Preprocessors

```{toctree}
---
maxdepth: 1
hidden: true
---
gzip <gzip>
jsonextract <jsonextract>
jsonarraysplit <jsonarraysplit>
jsonfilter <jsonfilter>
jsontimeextract <jsontimeextract>
syslogrouter <syslogrouter>
csvrouter <csvrouter>
regexrouter <regexrouter>
srcrouter <srcrouter>
tagrouter <tagRouter>
regextimestamp <regextimestamp>
regexextract <regexextract>
forwarder <forwarder>
gravwellforwarder <gravwellforwarder>
drop <drop>
ciscoise <ciscoise>
corelight <corelight>
plugin <plugin>
```

| Preprocessor | Purpose |
| -------------| -------- |
| [gzip](gzip) | Decompress gzipped data in entries |
| [jsonextract](jsonextract) | Parse and extract elements in JSON data |
| [jsonarraysplit](jsonarraysplit) | Parse JSON array data and split the array into individual entries |
| [jsonfilter](jsonfilter) | Parse JSON data and filter based on field values |
| [jsontimeextract](jsontimeextract) | Parse JSON data and extract timestamps from specific field values |
| [syslogrouter](syslogrouter) | Parse Syslog (RFC5424/RFC3164) data and route to specific tags based on field values |
| [csvrouter](csvrouter) | Parse CSV data and route to specific tags based on column values |
| [regexrouter](regexrouter) | Route entries to specific tags based on regular expression matches |
| [srcrouter](srcrouter) | Route entries to specific tags based on source IP or value |
| [tagrouter](tagRouter) | Route entries to specific tags based on the tag, or a combination of tag and either IP address or network |
| [regextimestamp](regextimestamp) | Perform complex timestamp processing using regular expressions |
| [regexextract](regexextract) | Perform data extractions and repacking using regular expressions |
| [forwarder](forwarder) | Forward entries using TCP or UDP connections |
| [gravwellforwarder](gravwellforwarder) | Forward entries using a Gravwell ingest connection |
| [drop](drop) | Simple dropping preprocessor, it stops all entries from moving through the preprocessor chain |
| [ciscoise](ciscoise) | Cisco ISE multi-message reconstruction preprocessor |
| [corelight](corelight) | Preprocessor to adapt Corelight JSON logs to Zeek TSV data |
| [plugin](plugin) | Preprocessor that loads interpretted code to perform custom preprocessing actions |
