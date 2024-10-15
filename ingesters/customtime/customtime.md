# Ingester Custom Time Formats

Many ingesters can support the inclusion of custom time formats that can extend the capability of the Gravwell timegrinder time resolution system.  The [timegrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder) has a wide array of timestamp formats that it can automatically identify and resolve.  However, in the real world with real developers there is no telling what time format a system may decide to use.  That is why we enable users to specify custom time formats for inclusion in the timegrinder system.

Custom time formats are a fallback when the usual timestamp extraction fails; refer to the [main ingesters page](ingesters_time) for more general information on timestamp extraction.

## Supported Ingesters

Not all ingesters support custom time formats.  One-off or standalone ingesters such as [singlefile](https://github.com/gravwell/gravwell/blob/v3.7.0/ingesters/singleFile/main.go) are applications meant to be invoked by hand and do not have a configuration file.  Dedicated ingesters like [netflow](ingesters_list) don't need to resolve timestamps, so there is no need for custom formats.

The following ingesters support the inclusion of custom time formats:

* [Simple Relay](/ingesters/simple_relay)
* [File Follower](/ingesters/file_follow)
* [HTTP Ingester](/ingesters/http)
* [Amazon Kinesis](/ingesters/kinesis)
* [Microsoft Graph API](/ingesters/msg)
* [Office 365](/ingesters/o365)
* [Kafka](/ingesters/kafka)

## Defining a Custom Format

A custom format requires three items to function:

* Name
* Regular Expression
* Format

The given name for a custom time format must be unique across other custom time formats and the included timegrinder formats.  For a complete up-to-date listing of included time formats and their names, check out our [timegrinder documentation](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder#pkg-constants).

Custom time formats are declared in the configuration files for supported ingesters by specifying a named `TimeFormat` INI block.  Here is an example format named "foo" which handles timestamps that are delimited using underscores:

```
[TimeFormat "foo"]
	Format="2006_01_02_15_04_05"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}`
```

This format would properly handle the timestamps in the following logs:

```
2021_02_05_09_00_00 and my id is 1
2021_02_05_09_00_00 and my id is 2
2021_02_05_09_00_00 and my id is 3
2021_02_05_09_00_00 and my id is 4
2021_02_05_09_00_00 and my id is 5
2021_02_05_09_00_00 and my id is 6
```

Here is another format that handles logs with only a timestamp:

```
[TimeFormat "foo2"]
	Format="15^04^05"
	Regex=`\d{1,2}\^\d{1,2}\^\d{1,2}`
```

This format would handle the following logs, appropriately applying the current date to each extracted timestamp:

```
09^00^00 and my id is 1
09^00^00 and my id is 2
09^00^00 and my id is 3
09^00^00 and my id is 4
09^00^00 and my id is 5
09^00^00 and my id is 6
```

```{note}
The custom timestamp format names can be used in [Timestamp-Format-Override](time_parsing_overrides) values.  For example we can force the timestamp format to our custom format using `Timestamp-Format-Override="foo"`.
```

### Pre-extraction

Sometimes, incoming data contains multiple timestamps or timestamp-like fields in the same format. For instance, consider the following log entry:

```
{ "message": "task completed", "start_time": "2023_11_28_09_00_00", "end_time": "2023_11_28_09_32_01" }
```

There are two fields (`start_time` and `end_time`) in the data, both using the same time format. If we want to be sure to extract the `end_time` field, we can add a *pre-extractor* to our TimeFormat definition with the `Extraction-Regex` option:

```
[TimeFormat "foo"]
	Format="2006_01_02_15_04_05"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}`
	Extraction-Regex=`"end_time":\s*"(?P<ts>[^"]+)"`
```

A pre-extractor is a regular expression which will match the desired timestamp and which contains a single *named capture group* (in this example, the group is named `ts`). The pre-extractor is evaluated first, then the contents of the named capture group will be evaluated against the `Format` and `Regex` arguments as normal.

In our example, `Extraction-Regex` looks for the string `"end_time":`, followed by a space, followed by a quoted string; the capture group is the contents of the quoted string. This means that the time parser will be operating on just the substring `2023_11_28_09_32_01` instead of the entire entry.

```{note}
Exactly one named capture group in the Extraction-Regex must be defined.  If no named capture groups are contained in the regex the configuration will be rejected.
```

#### Pre-extractions With Named Time Formats

Incoming data may also contain timestamps in formats that have already been defined in [timegrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder#Format) but that cannot be reliably extracted without first performing a pre-extraction. This is often the case with embedded `unix`, `unixmilli`, and `unixnano` timestamps.  The Pre-extraction `Extraction-Regex` can be combined with a named format so that timestamp formats that are already defined in [Timegrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder#Format). For example, consider the following entry:

```
[task completed] tss:1701200161 tse:1701200161.1234 value:1700000000
```

There are two Unix timestamps in the JSON data and one other value that would probably match a Unix timestamp. We can use a Pre-Extraction to grab a specific field and then pass it to the already defined UnixMilli timestamp processor.  An example definition which extracts the timestamp from the `tse` field and treats it as a `UnixSeconds` timestamp would look like this:

```
[TimeFormat "tseextractor"]
	Format="UnixSeconds"
	Extraction-Regex=`\s+tse:(?P<ts>\d+)`
```

```{note}
Notice that a Regex is not defined because we are using an already-defined timestamp extraction format.
```

### Time Formats

The `Format` component uses the [Go standard time format specification](https://golang.org/pkg/time/#pkg-constants).  Long story short, you must describe the date `Mon Jan 2 15:04:05 MST 2006` using whatever format you choose.

Time formats can omit the date component.  When the custom format system identifies that a custom time format does not include a date component, it will automatically update the extracted timestamp's date to `today`.

### Time Zones

All custom time formats will attempt to operate in UTC unless otherwise indicated using the `Format` directive.  This means that if you have a time format without a date component you must pay special attention to the timezone.  If an application emits a timestamp of `12:00:00` in MST and there is no timezone component or timezone overrides, timegrinder will interpret the timestamp as UTC and the extracted date will be 7 hours in the past.

If your timestamp does contain a timezone you must include that in your `Format` directive so that the timegrinder system knows to interpret the timestamp in the correct time zone.  For example here is the previously described "foo" custom format but with a timezone component:

```
[TimeFormat "foo"]
	Format="2006_01_02_15_04_05_MST"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\S+`
```

This example will properly handle timestamps in their respective time zones and apply the correct timestamp on extraction.

## Examples

Here is an example [File Follower](/ingesters/file_follow) configuration which adds two custom time formats:

```
[Global]
Ingester-UUID="463c1889-2954-40a0-a3b4-705ea66459f6"
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
State-Store-Location=/opt/gravwell/etc/file_follow.state
Log-Level=INFO #options are OFF INFO WARN ERROR
Log-File=/opt/gravwell/log/file_follow.log
Max-Files-Watched=64 # Maximum number of files to watch before rotating out old ones, this can be bumped but will need sysctl flags adjusted

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Follower "auth"]
	Base-Directory="/tmp/logs/"
	File-Filter="*.log" #we are looking for all authorization log files
	Tag-Name=test
	Assume-Local-Timezone=true #Default for assume localtime is false

[TimeFormat "foo"]
	Format="2006_01_02_15_04_05"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}`

[TimeFormat "foo2"]
	Format="15!04!05"
	Regex=`\d{1,2}!\d{1,2}!\d{1,2}`

```

The file follower will handle timestamps that are specified as `2021_02_14_12_33_52` and `15!05!22` properly due to the additional custom time formats.
