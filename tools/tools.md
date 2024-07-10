# Gravwell Tools

The Gravwell tools package (available in the Redhat and Debian repositories and [the downloads page](/quickstart/downloads)) contains small utilities which may come in handy when using Gravwell. The programs will be installed in `/usr/local/sbin` when the tools package is installed.

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
gravwell_timetester -h
Usage of gravwell_timetester:
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

```{note}
Terminals capable of handling ANSI color codes will highlight the timestamp location in the log in green.
```

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

The Splunk migration tool is [fully documented here](/migrate/migrate).

##  Account Unlock

The account unlock tool is a last-resort method for unlocking accounts, resetting passwords, or clearing MFA configurations.  While an admin user can perform these same functions via the User Administration screens in the UI or via the Gravwell CLI, if the *admin's* account is the one in trouble, this tool can fix it.

The account unlock tool is included in the installers for the core Gravwell system (webserver/indexer/searchagent) and for the datastore. It will be installed to `/opt/gravwell/bin/account_unlock`.

```{warning}
Before using the account unlock tool, you *must* shut down the Gravwell webserver! The webserver maintains a lock on the user database, so the tool cannot make any changes while the webserver is running.
```

The tool takes the following options:

* `-clear-mfa`: Clear MFA configuration, potentially forcing user to re-configure on next login
* `-lock`: If set, specified accounts will be locked rather than unlocked
* `-reset-password <password>`: If set, user password will be reset to this value
* `-override-userdb-path <path>`: Override the path to the webstore file

By default, the tool will always unlock any specified accounts, as well as performing additional actions (clearing MFA or resetting password) specified by the flags. If the `-lock` flag is used, the specified accounts will be *locked* rather than unlocked, in addition to any other actions specified.

To unlock the admin account and reset its password:

```
sudo systemctl stop gravwell_webserver
sudo /opt/gravwell/bin/account_unlock -reset-password changeme admin
sudo systemctl restart gravwell_webserver
```

To reset a user's MFA configuration and lock the account:

```
sudo systemctl stop gravwell_webserver
sudo /opt/gravwell/bin/account_unlock -clear-mfa -lock jsmith
sudo systemctl restart gravwell_webserver
```

### Account Unlock with Distributed Webservers

When you have [distributed webservers](/distributed/frontend), you cannot run the account unlock tool on the webserver because any changes will be overwritten immediately from the datastore component, which is considered the source of truth for user accounts. Instead, run the command *on the datastore system*, being sure to stop the datastore process itself first:

```
sudo systemctl stop gravwell_datastore
sudo /opt/gravwell/bin/account_unlock -reset-password changeme admin
sudo systemctl start gravwell_datastore
```

The changes should be rapidly propagated to the webservers from the datastore.

## Export

The export tool (`/usr/local/sbin/gravwell_export`) outputs the entries in a given indexer or well within an indexer to one or more compressed JSON archives of entries. The output files can natively be re-ingested to Gravwell. The DATA portion of an entry is base64 encoded, while the TAG, SRC, and Timestamp are strings.  Intrinsic EVs attached to data at ingest are exported under the `Enumerated` field.

### Usage

To export data, run the export tool with the server address and other necessary configuration options, as shown below. The tool will prompt for login credentials, and begin to export data. 

```
Usage of gravwell_export:
  -insecure
    	Do NOT enforce webserver certificates, TLS operates in insecure mode
  -insecure-no-https
    	Use insecure HTTP connection, passwords are shipped plaintext
  -max-duration string
    	maximum duration in the past to export data
  -output string
    	Output directory
  -s string
    	Address and port of Gravwell webserver
  -well string
    	limit export to a specific well
```

For example:

```
/usr/local/sbin/gravwell_export -output tmp/ -s 10.0.0.1
Username:  admin
Password:  
processing well default to tmp/default containing 3 tags and 2 shards
Total Data Processed: 8.00 MB                                     
DONE

ls -l tmp/default 
total 8196
-rw-r--r-- 1 gravwell gravwell 4891411 Jun 28 14:31 2023-06-26-21:54:08.json
-rw-r--r-- 1 gravwell gravwell 3494587 Jun 28 14:31 2023-06-28-10:18:40.json
```

### Output Data Structure

The export tool outputs data in a JSON structure which is easy to work with using a variety of tools; the native golang types are documented on our open source repo [here](https://pkg.go.dev/github.com/gravwell/gravwell/v3@v3.8.24/client/types#SearchEntry) and [here](https://pkg.go.dev/github.com/gravwell/gravwell/v3@v3.8.24/client/types#EnumeratedPair):

The basic types are:

```
type SearchEntry struct {
	TS         entry.Timestamp
	SRC        net.IP
	Tag        entry.EntryTag
	Data       []byte
	Enumerated []EnumeratedPair
}

type EnumeratedPair struct {
	Name     string
	Value    string             `json:"ValueStr"`
	RawValue RawEnumeratedValue `json:"Value"`
}

type RawEnumeratedValue struct {
	Type uint16
	Data []byte
}
```

An example JSON encoded entry from the export tool is:

```
{
  "TS": "2023-11-06T13:32:33.736216872Z",
  "Tag": "testdata",
  "SRC": "172.19.100.1",
  "Data": "Q2hlY2tFVnMsIG5vdGhpbmcgaW50ZXJlc3RpbmcgaGVyZQ==",
  "Enumerated": [
    {
      "Name": "_type",
      "ValueStr": "intrinsic evs",
      "Value": {
        "Type": 13,
        "Data": "aW50cmluc2ljIGV2cw=="
      }
    },
    {
      "Name": "ts",
      "ValueStr": "2023-11-06T13:32:33.736216872Z",
      "Value": {
        "Type": 16,
        "Data": "8eHa3A4AAAAox+Er"
      }
    },
    {
      "Name": "user",
      "ValueStr": "jaydenwilliams474",
      "Value": {
        "Type": 13,
        "Data": "amF5ZGVud2lsbGlhbXM0NzQ="
      }
    },
    {
      "Name": "name",
      "ValueStr": "William Robinson",
      "Value": {
        "Type": 13,
        "Data": "V2lsbGlhbSBSb2JpbnNvbg=="
      }
    }
  ]
}
```
