# Ingester Custom Time Formats

Many ingesters can support the inclusion of custom time formats that can extend the capability of the gravwell TimeGrinder time resolution system.  The [TimeGrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder) has a wide array of timestamp formats that it can automatically identify and resolve.  However, in the real world with real developers there is no telling what time format a system may decide to use.  That is why we enable users to specify custom time formats for inclusion in the TimeGrinder system.

## Supported Ingesters

Not all ingesters support the inclusion of custom time formats.  One-off or standalone ingesters such as [singlefile](https://github.com/gravwell/gravwell/blob/v3.7.0/ingesters/singleFile/main.go) are applications meant to be invoked by hand and do not have a configuration file.  Dedicated ingesters like [netflow](#!ingesters/ingesters.md#Netflow_Ingester) don't need to resolve timestamps, so there is no need for custom formats.

The following ingesters support the inclusion of custom time formats:

* [Simple Relay](#!ingesters/ingesters.md#Simple_Relay)
* [File Follower](#!ingesters/ingesters.md#File_Follower)
* [HTTP Ingester](#!ingesters/ingesters.md#HTTP)
* [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester)
* [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)
* [Office 365](#!ingesters/ingesters.md#Office_365_Log_Ingester)
* [Kafka](#!ingesters/ingesters.md#Kafka)

## Defining a Custom Format

A custom format requires three items to function:

* Name
* Regular Expression
* Format

The given name for a custom time format must be unique across other custom time formats and the included timegrinder formats.  For a complete up-to-date listing of included time formats and their names, check out our [timegrinder documentation)[https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder#pkg-constants].

Custom time formats are declared in the configuration files for supported ingesters by specifying a named `TimeFormat` INI block.  Here is an example format which handles timestamps that are delimited using underscores:

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

Note: The custom timestamp format names can be used in [Timestamp-Format-Override](#!ingesters/ingesters.md#Time_Parsing_Overrides) values.

### Time Formats

The `Format` component uses the [golang standard time format specification](https://golang.org/pkg/time/#pkg-constants).  Long story short, you must describe the date `Mon Jan 2 15:04:05 MST 2006` using whatever format you choose.

Time formats can ommit the date component.  When the custom format system identifies that a custom time format does not include a date component it will automatically update the extracted timestamps date to `today`.

Warning: All custom time formats will attempt to operate in UTC unless otherwise indicated using the Format directive.  This means that if you have a time format without a date component you must pay special attention to the timezone.  If an application emits a timestamp of `12:00:00` in MST and there is no timezone component or timezone overrides, timegrinder will interpret the timestamp as UTC and the extracted date will be 7 hours in the past.

## Examples

Here is an example [File Follower](#!ingesters/file_follow.md) configuration:

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
