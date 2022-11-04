# Migrating Data

When you first stand up Gravwell, one of the first tasks is typically getting data from an old system or log archive into Gravwell.  Migration may involve pulling historical data out of an existing system like Splunk, scraping a large database, or just ingesting 5 years of old syslog from flat files.  Gravwell provides a series of "normal" ingesters designed to stream data in from a variety of sources in near real-time, but for one-time migration of existing data, there are specialized tools which may be a better choice.

This section will explore the available tools to migrate existing data and give some tips on how to efficiently migrate potentially hundreds of TB of existing logs into a new Gravwell instance.  We will examine a few scenarios where using migration tools will provide much better migration performance and storage efficiency as opposed to just tossing something at a typical ingester.  Most of our migration tools are open source, so we can also provide links to source code and deep dive examinations of functionality.  Most data migrations are a one time occurrence, which means a one-off tool that is specialized for the task is usually the right answer.  

### Migration Caveats

Most Gravwell licenses are unlimited, meaning that when it is time to migrate massive quantities of data in you are only limited by the resources available to accept and index that data.  However, the one exception is the free Community Edition license which has hard ingest limits.  All other license types are either unlimited or allow for bursting to accommodate data migration.

## Splunk Migration

Gravwell provides an interactive migration tool which can import data from Splunk. See the [migrate tool](/migrate/migrate) documentation.

## Importing Many Files

Gravwell provides an interactive migration tool which can pull data from static files on disk efficiently. See the [migrate tool](/migrate/migrate) documentation.

Note: Only use the migrate tool if the files on disk are not expected to change. To ingest files which are still being added to, use the [File Follow](/ingesters/file_follow) ingester.

## Importing One File

The single file ingester is one of the most simplistic ingesters in the Gravwell arsenal.  It is designed to ingest a single line-delimited file to a specific tag.  It can transparently decompress files and has some limited parsing ability.  If you have a single large Apache access log or just need to script up some one off file ingestion, it can be the simplest option. The ingester is included in the `gravwell-tools` package for Debian and Redhat, and is also available as a standalone shell installer on [our downloads page](/quickstart/downloads). Once installed, the program is located at `/usr/local/sbin/gravwell_single_file_ingester`.

The ingester is a standalone ingester that is designed to operate using flags rather than a config file.  This means that it lacks some of the additional functionality of other ingesters such as custom timestamp definitions and preprocessor support.  The following flags are supported:


| Flag | Required | Description | Example |
|------|----------|-------------|---------|
| -h   |          | Print options and help and exit | |
| -version |      | Print version information and exit | |
| -verbose |      | Be very verbose in operation, this means printing every entry as it is ingested as well as step-by-step status updates | |
| -i | X          | Specify the input file, setting this to "-" means read from stdin | -i /tmp/access.log |
| -clear-conns | X | Specify an IP or IP:Port for ingest, typically an indexer.  Specify multiple indexers using a comma delimited list | -clear-conns=127.0.0.1,172.16.0.1:4023,192.168.1.1:4423 |
| -tls-conns | X | Specify an IP or IP:Port for ingest using a TLS connection.  Specify multiple indexers using a comma delimited list | -clear-conns=127.0.0.1,172.16.0.1:4024,192.168.1.1:4424 |
| -pipe-conns | X | Specify a Unix named pipe for ingest. | -pipe-conns=/opt/gravwell/comms/pipe |
| -ingest-secret | X | Specify your ingest authentication token. | -ingest-secret="ASuperSecretString" |
| -tag-name | X | The tag that all entries will be ingested to. | -tag-name=apacheaccess |
| -tls-remote-verify | | Enable or disable TLS certificate validation when using TLS connections, default is true. | -tls-remote-verify=false |
| -timeout | | Timeout in seconds the ingester should use when attempting to connect to an indexer. | -timeout=10 |
| -ignore-ts | | Do not attempt to resolve timestamps out of entries, everything is ingested at the time "NOW" | -ignore-ts |
| -timestamp-override | | Specify a specific timestamp format to use. | -timestamp-override="rfc3339" |
| -timezone-override | | Override the timezone applied if the ingester cannot resolve a timezone in the timestamp itself. | -timezone-override=UTC |
| -utc |          | Assume discovered timestamps are in the UTC timezone | |
| -source-override | | Apply a SRC value instead of allowing the indexer to choose one for you.  This allows for manually setting source values. | -source-override="192.168.1.2" |
| -block-size |   | Ingest data in batches, this can increase throughput. | -block-size 256 |
| -clean-quotes | | When an entry is surrounded in quotes, remove the quotes before ingesting | |
| -quotable-lines | | Lines may be quotable, meaning a newlines encapsulated in quotes do not delimit entries |
| -ignore-prefix | | Specify a string that when found at the start of a line causes the line to be ignored, useful for ignoring headers on CSV files | -ignore-prefix="#" |
| -status | | Output ingest stats as we go | |

### Example

The following command ingests the contents of `/tmp/my-logs.txt`, one entry per line. It ignores (does not ingest) any lines starting with the `#` character. It will attempt to extract timestamps from the log entries (the default) but if no timezone is specified in the log entry, it will assume America/Chicago. The entries will be ingested to the indexer at `10.0.0.50`.

```
/usr/local/sbin/gravwell_single_file_ingester -clear-conns 10.0.0.50 -ingest-secret xyzzy123 -timezone-override "America/Chicago" -ignore-prefix "#" -status -i /tmp/my-logs.txt
```

## Importing PCAP Files

If you have existing PCAP files (from Wireshark or tcpdump or some other packet capture tool), you can ingest them using the PCAP file ingester.  The ingester is included in the `gravwell-tools` package for Debian and Redhat, and is also available as a standalone shell installer on [our downloads page](/quickstart/downloads). Once installed, the program is located at `/usr/local/sbin/gravwell_pcap_file_ingester`.

The ingester is a standalone ingester that is designed to operate using flags rather than a config file.  This means that it lacks some of the additional functionality of other ingesters such as custom timestamp definitions and preprocessor support.  The following flags are supported:


| Flag | Required | Description | Example |
|------|----------|-------------|---------|
| -h   |          | Print options and help and exit | |
| -version |      | Print version information and exit | |
| -verbose |      | Be very verbose in operation, this means printing every entry as it is ingested as well as step-by-step status updates | |
| -pcap-file | X          | Specify the input file | -i /tmp/netflow.pcap |
| -clear-conns | X | Specify an IP or IP:Port for ingest, typically an indexer.  Specify multiple indexers using a comma delimited list | -clear-conns=127.0.0.1,172.16.0.1:4023,192.168.1.1:4423 |
| -tls-conns | X | Specify an IP or IP:Port for ingest using a TLS connection.  Specify multiple indexers using a comma delimited list | -clear-conns=127.0.0.1,172.16.0.1:4024,192.168.1.1:4424 |
| -pipe-conns | X | Specify a Unix named pipe for ingest. | -pipe-conns=/opt/gravwell/comms/pipe |
| -ingest-secret | X | Specify your ingest authentication token. | -ingest-secret="ASuperSecretString" |
| -tag-name | X | The tag that all entries will be ingested to. | -tag-name=pcap |
| -timeout | | Timeout in seconds the ingester should use when attempting to connect to an indexer. | -timeout=10 |
| -ts-override | | Do not attempt to resolve timestamps out of entries, everything is ingested at the time "NOW" | -ts-override |
| -source-override | | Apply a SRC value instead of allowing the indexer to choose one for you.  This allows for manually setting source values. | -source-override="192.168.1.2" |
| -tls-remote-verify | | Enable or disable TLS certificate validation when using TLS connections, default is true. | -tls-remote-verify=false |

### Example

The following ingests the packets in /tmp/netflow-capture.pcap to an indexer on the local machine. Each packet is ingested as a single entry.

```
/usr/local/sbin/gravwell_pcap_file_ingester -pipe-conns /opt/gravwell/comms/pipe -ingest-secret MyIngestSecret -pcap-file /tmp/netflow-capture.pcap
```
