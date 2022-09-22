## Regex Timestamp Extraction Preprocessor

Ingesters will typically attempt to extract a timestamp from an entry by looking for the first thing which appears to be a valid timestamp and parsing it. In combination with additional ingester configuration rules for parsing timestamps (specifying a specific timestamp format to look for, etc.) this is usually sufficient to properly extract the appropriate timestamp, but some data sources may defy these straightforward methods. Consider a situation where a network device may send CSV-formatted event logs wrapped in syslog--a situation we have seen at Gravwell!

The Regex Timestamp Extraction preprocessor Type is `regextimestamp`.

### Supported Options

* `Regex` (string, required): This parameter specifies the regular expression to be applied to the incoming entries. It must contain at least one [named capturing group](https://www.regular-expressions.info/named.html), e.g. `(?P<timestamp>.+)` which will be used with the `TS-Match-Name` parameter.
* `TS-Match-Name` (string, required): This parameter gives the name of the named capturing group from the `Regex` parameter which will contain the extracted timestamp.
* `Timestamp-Format-Override` (string, optional): This can be used to specify an alternate timestamp parsing format. Available time formats are:
	- AnsiC
	- Unix
	- Ruby
	- RFC822
	- RFC822Z
	- RFC850
	- RFC1123
	- RFC1123Z
	- RFC3339
	- RFC3339Nano
	- Apache
	- ApacheNoTz
	- Syslog
	- SyslogFile
	- SyslogFileTZ
	- DPKG
	- NGINX
	- UnixMilli
	- ZonelessRFC3339
	- SyslogVariant
	- UnpaddedDateTime
	- UnpaddedMilliDateTime
	- UK
	- Gravwell
	- LDAP
	- UnixSeconds
	- UnixMs
	- UnixNano

	Some timestamp formats have values that overlap (for example LDAP and UnixNano can produce timestamps with the same number of digits). If `Timestamp-Format-Override` is not used, the preprocessor will attempt to derive the timestamp in the order listed above. Always use `Timestamp-Format-Override` if using a timestamp format that can conflict with others in this list.

* `Timezone-Override` (string, optional): If the extracted timestamp doesn't contain a timezone, the timezone specified here will be applied. Example: `US/Pacific`, `Europe/Rome`, `Cuba`.
* `Assume-Local-Timezone` (boolean, optional): This option tells the preprocessor to assume the timestamp is in the local timezone if no timezone is included. This is mutually exclusive with the `Timezone-Override` parameter.


### Common Use Cases

Many data streams may have multiple timestamps or values that can easily be interpreted as timestamps.  The regextimestamp preprocessor allows you to force timegrinder to examine a specific timestamp within a log stream.  A good example is a log stream that is transported via syslog using an application that includes it's own timestamp but does not relay that timestamp to the syslog API.  The syslog wrapper will have a well formed timestamp which, but the actual data stream may need to use some internal field for the accurate timestamp.

### Example: Wrapped Syslog Data

```
Nov 25 15:09:17 webserver alerts[1923]: Nov 25 14:55:34,GET,10.1.3.4,/info.php
```

We would like to extract the inner timestamp, "Nov 25 14:55:34", for the TS field on the ingested entry. Because it uses the same format as the syslog timestamp at the beginning of the line, we cannot extract it with clever timestamp format rules. However, the regex timestamp preprocessor can be used to extract it. By specifying a regular expression which captures the desired timestamp in a named sub-match, we can extract timestamps from anywhere in an entry. For this entry, the regex `\S+\s+\S+\[\d+\]: (?<timestamp>.+),` should be sufficient to properly extract the desired timestamp.

This config could be used to extract the timestamp shown in the example above:

```
[Preprocessor "ts"]
	Type=regextimestamp
	Regex="\S+\s+\S+\[\d+\]: (?P<timestamp>.+),"
	TS-Match-Name=timestamp
	Timezone-Override=US/Pacific
```
