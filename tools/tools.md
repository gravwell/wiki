# Gravwell Tools

The Gravwell tools package (available in the Redhat and Debian repositories and [the downloads page](#!quickstart/downloads.md)) contains small utilities which may come in handy when using Gravwell. The programs will be installed in `/usr/local/sbin` when the tools package is installed.

## Common Arguments for Ingester Tools

Several of the tools described here can ingest entries to Gravwell indexers. They use a common set of flags to specify how to communicate with the indexer. These flags are fully described in each tool's help output (use the `-h` flag), but the most important ones are described here too:

* `-clear-conns`, `-tls-conns`, `-pipe-conns`: specifies the indexer(s) to which the entries should be sent. For example, `-clear-conns 10.0.0.5:4023,10.0.0.6:4023`, or `-tls-conns idx1.example.org:4024`, or `-pipe-conns /opt/gravwell/comms/pipe`.
* `-ingest-secret`: specifies the authentication token used to ingest entries (the `Ingest-Auth` parameter defined in the indexer's gravwell.conf file)
* `-tag-name`: specifies the tag to use for the entries.

## Entry Generator

The generator program (`/usr/local/sbin/gravwell_generator`) is capable of creating & ingesting artificial entries in a variety of formats. This can be very useful for testing if you don't have any appropriate "real" data on hand. There are many command-line options which can be used to adjust the parameters of the entries created or the way they are ingested; use the `-h` flag to see a complete list. The most important options are:

* `-type`: specifies the kind of entries to generate, e.g. syslog or json or Zeek. See below for a complete list.
* `-entry-count`: specifies how many entries to generate.
* `-duration`: specifies how much time the entries should cover, e.g. `-duration 1h` generates entries across the last hour, while `-duration 30d` generates entries across the last month.

The `-type` flag supports the following choices:

* `bind`: BIND DNS server logs.
* `zeekconn`: Zeek connection logs.
* `csv`: Comma-separated values; column names are `ts, app, id, uuid, src, srcport, dst, dstport, data, name, country, city, hash`
* `dnsmasq`: DNSMASQ server logs.
* `fields`: entries which can be parsed using the fields module; columns are: `ts, app, src, srcport, dst, dstport, data` (note that the `-fields-delim-override` flag can be used to select the delimiter character, default is tab).
* `json`: JSON objects.
* `binary`: entries which can be parsed using the slice module as follows: `slice int16([0:2]) as LE16 int16be([2:4]) as BE16 int32([4:8]) as LE32 int32be([8:12]) as BE32 int64([12:20]) as LE64 int64be([20:28]) as BE64 float32([28:32]) as F32 float64be([32:40]) as BEF64 ipv4([40:44]) as IP string([44:]) as STR`
* `regex`: entries which can be parsed using the regex module as follows: `regex "(?P<ts>\S+)\s\[(?P<app>\S+)\]\s<(?P<uuid>\S+)>\s(?P<src>\S+)\s(?P<srcport>\d+)\s(?P<dst>\S+)\s(?P<dstport>\d+)\s(?P<path>\S+)\s(?P<useragent>.+)\s\{(?P<email>\S+)\}$"`

For example, to generate random 1000 JSON entries (an excellent choice for testing) spread over the last day:

```
/usr/local/sbin/gravwell_generator -clear-conns 10.0.0.5:4023 -ingest-secret xyzzy1234 -entry-count 1000 -duration 24h -tag-name json -type json
```

## Single File Ingester

The single file ingester (`/usr/local/sbin/gravwell_single_file_ingester`) is a convenient way to ingest a text file into Gravwell. Each line in the file will become one entry. By default, it will attempt to extract timestamps from the entries, unless the `-ignore-ts` or `-timestamp-override` flags are used.

To ingest a system log file on Linux into the "syslog" tag:

```
/usr/local/sbin/gravwell_single_file_ingester -clear-conns 10.0.0.5:4023 -ingest-secret xyzzy1234 -tag-name syslog -i /var/log/messages
```

## PCAP Ingester

The PCAP ingester (`/usr/local/sbin/gravwell_pcap_file_ingester`) can ingest the contents of a [pcap file](https://en.wikipedia.org/wiki/Pcap), with each packet becoming its own entry. By default, each packet will be tagged with the correct timestamp as extracted from the file; if the `-ts-override` flag is set, the timestamps will be "shifted" such that the oldest packet will be tagged with the current time and further packets will be in the future.

To ingest a pcap file into the "packet" tag:

```
/usr/local/sbin/gravwell_pcap_file_ingester -clear-conns 10.0.0.5:4023 -ingest-secret xyzzy1234 -tag-name packet -pcap-file /tmp/http.pcap
```

## Timestamp Testing

The purpose of this tool is to test [TimeGrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder) against log files.  The Time Tester tool can operate in one of two ways:

* Basic Mode
* Custom Timestamp Mode


### Basic Mode
Basic mode simply shows which timestamp extraction will match a given log line and where in the log line it matched.

It will show each log line with the timestamp capture location highlighted in red, the extracted timestamp, and the extraction name that hit.

### Custom Timestamp Mode

The custom timestamp mode operates the same as the basic mode but also accepts a path to custom timestamp declarations which allows you to test custom timestamps and also see how collisions or misses affect the TimeGrinder.

### Usage

Time Tester will walk each string provided on the command line and attempt to process it as if it were a timestamp.

The tester can set the timegrinder config values and define custom timestamp formats in the same way that Gravwell ingesters can.

```
./timetester -h
Usage of ./timetester:
  -custom string
        Path to custom time format configuration file
  -enable-left-most-seed
        Activate EnableLeftMostSeed config option
  -format-override string
        Enable FormatOverride config option
```

Here is an example of a custom time format that adds a year to the Syslog format:
```
[TimeFormat "syslog2"]
        Format=`Jan _2 15:04:05 2006`
        Regex=`[JFMASOND][anebriyunlgpctov]+\s+\d+\s+\d\d:\d\d:\d\d\s\d{4}`
```

#### Single Log Entry Test
Here is an example invocation using the basic mode and testing a single entry:

```
timetester "2022/05/27 12:43:45 server.go:233: the wombat hit a curb"
```


Results:

```
2022/05/27 12:43:46 server.go:233: the wombat hit a curb
	2022-05-27 12:43:46 +0000 UTC	NGINX
```

**NOTE:** Terminals capable of handling ANSI color codes will highlight the timestamp location in the log in green.

#### Multiple Log Entry Test
Here is an example that tests 3 entries in succession showing how different extractors operated.
First we use a custom time format from a custom application, then a Zeek connection log, then back to the custom time format.  This shows that timegrinder will fail on the existing format then find a new format then go back to the old format.

```
timetester "2022/05/27 12:43:45 server.go:141: the wombat can't jump" \
	"1654543200.411042	CUgreS31Jc2Elmtph5	1.1.1.1	38030	2.2.2.2	23" \
	"2022/05/27 12:43:46 server.go:233: the wombat hit a curb"
```

Results:

```
2022/05/27 12:43:46 server.go:233: the wombat hit a curb
	2022-05-27 12:43:46 +0000 UTC	NGINX
1654543200.411042      CUgreS31Jc2Elmtph5      1.1.1.1 38030   2.2.2.2 23
	2022-06-06 19:20:00.411041975 +0000 UTC	UnixMilli
2022/05/27 12:43:46 server.go:233: the wombat hit a curb
	2022-05-27 12:43:46 +0000 UTC	NGINX
```

### Caveats

The TimeGrinder object is re-used across each test, in order to simulate a single TimeGrinder object that is being used to process a string of logs on a given listener.  The Timegrinder system is designed to "lock on to" a given timestamp format and continue re-using it.  This means that if a log format misses but another hits, the TimeGrinder will continue using the format that "hit".

For example, if you input the following three values:

```
1984-10-26 12:22:31 T-800 "I'm a friend of Sarah Connor. I was told she was here. Could I see her please?"
1991-6-1 12:22:31.453213 1991 T-1000 "Are you the legal guardian of John Connor?"
2004-7-25 22:18:24 Connor "It's not everyday you find out that you're responsible for three billion deaths. He took it pretty well."
```

The system will correctly interpret the first timestamp and lock onto the `UnpaddedDateTime` format, it would then see that the second line (minus the millisecond) also matches the UnpaddedDateTime format and use it, ignoring the millisecond components.  This is an artifact of the how the timegrinder optimizes its throughput by assuming that contiguous entries will be of the same format.


## Splunk Migration

The Splunk migration tool is [fully documented here](https://docs.gravwell.io/#!migrate/migrate.md).
