# File Follower

The File Follower ingester is the best way to ingest files on the local filesystem in situations where those files may be updated on the fly. It ingests each line of the files as a single entry.

The most common use case for File Follower is monitoring a directory containing log files which are actively being updated, such as /var/log. It intelligently handles log rotation, detecting when `logfile` has been moved to `logfile.1` and so on. It can be configured to ingest files matching a specific pattern in a directory, optionally recursively descending into the subdirectories of that top-level directory.

Attention: On RHEL/CentOS, `/var/log` belongs to the "root" group, not "adm" as we assume. File Follower runs in the adm group by default, so if you want it to read `/var/log` you need to `chgrp -R adm /var/log` OR change the group in the systemd unit file.

## Kernel parameter tuning

On Linux, File Follower uses the inotify kernel subsystem. If you are seeing filesystem notification errors in File Follow, consider increasing the depth of the inotify event queue. You can do this by setting the depth with sysctl: `sysctl -w fs.inotify.max_queued_events=32768`. Additionally, the maximum number of files allowed to be watched is controlled by the kernel. You can increase this value with sysctl: `sysctl -w fs.inotify.max_user_watches=32768`. 

## Startup considerations

At startup, file follower will ingest any existing data in the tracked paths before servicing inotify events. If you are starting file follower to ingest a large amount of existing data, it is recommended that you do not write to the tracked paths until the initial ingest is complete. Writing to tracked paths during initial ingest could result in overflowing the kernel inotify buffer.

## Basic Configuration

The File Follower configuration file is by default located in `/opt/gravwell/etc/file_follow.conf` on Linux and `C:\Program Files\gravwell\file_follow.cfg` on Windows.

The File Follower ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, File Follower supports multiple upstream indexers, TLS, cleartext, and named pipe connections, and local logging.

Note: We recommend strongly against using a file cache with the File Follower ingester, since it is already tracking its position within the source files.

An example configuration for the File Follower ingester, configured to watch several different types of log files in /var/log and recursively follow files under /tmp/incoming:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Cleartext-Backend-target=172.20.0.1:4023 #example of adding a cleartext connection
Cleartext-Backend-target=172.20.0.2:4023 #example of adding another cleartext connection
State-Store-Location=/opt/gravwell/etc/file_follow.state
Log-Level=ERROR #options are OFF INFO WARN ERROR
Max-Files-Watched=64

[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true
[Follower "auth"]
        Base-Directory="/var/log/"
        File-Filter="auth.log,auth.log.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
[Follower "packages"]
        Base-Directory="/var/log"
        File-Filter="dpkg.log,dpkg.log.[0-9]" #we are looking for all dpkg files
        Tag-Name=dpkg
        Ignore-Timestamps=true
[Follower "external"]
        Base-Directory="/tmp/incoming"
		Recursive=true
        File-Filter="*.log"
        Tag-Name=external
		Timezone-Override="America/Los_Angeles"
```

In this example, the "syslog" follower reads `/var/log/syslog` and its rotations, ingesting lines to the syslog tag and assuming dates to be in the local timezone. Similarly, the "auth" follower also uses the syslog tag to ingest `/var/log/auth.log`. The "packages" follower ingests Debian's package management logs to the dpkg tag; for the purposes of illustration, it ignores the timestamps and marks each entry with the time it was read.

Finally, the "external" follower reads all files ending in `.log` from the directory `/tmp/incoming`, descending recursively into directories. It parses timestamps as though they were in the Pacific time zone. This follower illustrates a configuration that would be useful if, for example, several servers on the US west coast periodically uploaded their log files to this system.

The configuration parameters used above are explained in greater detail in the following sections

## Additional Global Parameters

### Max-Files-Watched

The Max-Files-Watched parameter prevents the File Follower from maintaining too many open file descriptors. If `Max-Files-Watched=64` is specified, the File Follower will actively watch up to 64 log files. When a new file is created, the File Follower will stop actively watching the oldest existing file in order to watch the new one. However, if the old file is later updated, it will return to the top of the queue.

We recommend leaving this setting at 64 in most cases; configuring the limit too high can run into limits set by the kernel.

## Follower Configuration

The File Follower configuration file contains one or more "Follower" directives:

```
[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
```

Each follower specifies at minimum a base directory and a filename filtering pattern. This section describes possible configuration parameters which can be set per follower.

###	Base-Directory

The Base-Directory parameter specifies the directory which will contain the files to be ingested. It should be an absolute path and contain no wildcards.

### File-Filter

The File-Filter parameter defines the filenames which should be ingested. It can be as simple as a single file name:

```
File-Filter="foo.log"
```

Or it can contain multiple patterns:

```
File-Filter="kern*.log,kern*.log.[0-9]"
```

which will match any filename beginning with "kern" and ending with ".log", or beginning with "kern" and ending with ".log.0" through ".log.9".

The full matching syntax, as defined in [https://golang.org/pkg/path/filepath/#Match](https://golang.org/pkg/path/filepath/#Match):

```
pattern:
	{ term }
term:
	'*'         matches any sequence of non-Separator characters
	'?'         matches any single non-Separator character
	'[' [ '^' ] { character-range } ']'
	            character class (must be non-empty)
	c           matches character c (c != '*', '?', '\\', '[')
	'\\' c      matches character c

character-range:
	c           matches character c (c != '\\', '-', ']')
	'\\' c      matches character c
	lo '-' hi   matches character c for lo <= c <= hi
```

### Recursive

The recursive parameter directs the File Follower to ingest files matching the File-Filter recursively under the Base-Directory.

By default, the ingester will only ingest those files matching the File-Filter under the top level of the Base-Directory; the following would ingest `/tmp/incoming/foo.log` but not `/tmp/incoming/system1/foo.log`:

```
Base-Directory="/tmp/incoming"
File-Filter="foo.log"
Recursive=false
```

By setting Recursive=true, the configuration will ingest **any** file named foo.log at any directory depth under `/tmp/incoming`.

### Tag-Name

The Tag-Name parameter specifies the tag to apply to entries ingested by this follower.

### Ignore-Line-Prefix

The ingester will drop (not ingest) any lines beginning with the string passed to Ignore-Line-Prefix. This is useful when ingesting log files which contain comments, such as Bro logs. The Ignore-Line-Prefix parameter may be specified multiple times.

The following indicates that lines beginning with `#` or `//` should not be ingested:

```
Ignore-Line-Prefix="#"
Ignore-Line-Prefix="//"
```

### Ignore-Glob

The ingester will drop (not ingest) any lines that match the given glob pattern. Globs are text patterns containing wildcards (such as `*`). For example, to drop any line that contains the word "foo" anywhere in the line:

```
Ignore-Glob="*foo*"
```

Ignore-Glob can be specified multiple times and supports the same wildcard syntax as the File-Filter option.

### Regex-Delimiter

The `Regex-Delimiter` option allows the user to specify a regular expression which will be used to split entries, rather than newlines. Thus, if your input files look like this:

```
####This is the first entry
additional data
####This is the second entry
```

You could add the following line to your follower definition:

```
Regex-Delimiter="####"
```

This will parse the preceding file into two entries:

```
####This is the first entry
additional data
```

and

```
####This is the second entry
```

Note: `Timestamp-Delimited` overrides `Regex-Delimiter`; set one or the other.

### Timestamp-Delimited

The Timestamp-Delimited parameter is a boolean specifying that each occurrence of a time stamp should be considered the start of a new entry. This is useful when log entries may span multiple lines. When specifying Timestamp-Delimited, the Timestamp-Format-Override parameter must also be set.

If a log file looks like this:

```
2012-11-01T22:08:41+00:00 Line 1 of the first entry
Line 2 of the first entry
2012-11-01T22:08:43+00:00 Line 1 of the second entry
Line 2 of the second entry
Line 3 of the second entry
```

Provided the follower is configured with `Timestamp-Delimited=true` and `Timestamp-Format-Override=RFC3339`, it will generate the following two entries:

```
2012-11-01T22:08:41+00:00 Line 1 of the first entry
Line 2 of the first entry
```
```
2012-11-01T22:08:43+00:00 Line 1 of the second entry
Line 2 of the second entry
Line 3 of the second entry
```

Note: `Timestamp-Delimited` overrides `Regex-Delimiter`; set one or the other.

### Ignore-Timestamps

The Ignore-Timestamps parameter indicates that the follower should not attempt to extract a timestamp from each line of the file, but rather tag each line with the current time.

### Assume-Local-Timezone

Assume-Local-Timezone is a boolean setting which directs the ingester to parse timestamps which lack timezone specifications as though they were in the local time zone rather than the default UTC.

Assume-Local-Timezone and Timezone-Override are mutually exclusive.

### Timezone-Override

The Timezone-Override parameter directs the ingester to parse timestamps which lack timezone specifications as though they were in the given time zone rather than the default UTC. The timezone should be specified in IANA database string format as shown in [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones); for example, US Central Time should be specified as follows:

```
Timezone-Override="America/Chicago"
```

Assume-Local-Timezone and Timezone-Override are mutually exclusive. `Timezone-Override="Local"` is functionally equivalent to `Assume-Local-Timezone=true`

### Timestamp-Format-Override

Data values may contain multiple timestamps which can cause some confusion when attempting to derive timestamps out of the data.  Normally, the followers will grab the left most timestamp that can be derived, but if there are multiple timestamps in an entry it may be useful to specify a format to try first.  "Timestamp-Format-Override" tells the follower to try a specific format first.  The following timestamp formats are available:

* AnsiC
* Unix
* Ruby
* RFC822
* RFC822Z
* RFC850
* RFC1123
* RFC1123Z
* RFC3339
* RFC3339Nano
* Apache
* ApacheNoTz
* Syslog
* SyslogFile
* SyslogFileTZ
* DPKG
* Custom1Milli
* NGINX
* UnixMilli
* ZonelessRFC3339
* SyslogVariant
* UnpaddedDateTime

Refer to [the timegrinder documentation](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder) for a full list of possible overrides, with examples.

To force the follower to first look for timestamps that match the RFC3339 specification, add `Timestamp-Format-Override=RFC3339` to the follower. Note that if it is unable to find an RFC3339 timestamp, it will attempt to match against other formats too.

### Timestamp-Regex and Timestamp-Format-String

The `Timestamp-Regex` and `Timestamp-Format-String` options may be used in tandem to specify an additional timestamp format for use when parsing timestamps for this follower. For example, if you are ingesting logs containing Oracle WebLogic timestamps (e.g. "Sep 18, 2020 12:26:48,992 AM EDT"), you would add the following to your configuration:

```
	Timestamp-Regex=`[JFMASOND][anebriyunlgpctov]+\s+\S{1,2},\s+\d{4}\s+\d{1,2}:\d\d:\d\d,\d+\s+\S{2}\s+\S+`
	Timestamp-Format-String="Jan _2, 2006 3:04:05,999 PM MST"
```

The `Timestamp-Format-String` parameter should be a Go-style timestamp format as defined [in this document](https://golang.org/pkg/time/). The `Timestamp-Regex` parameter should be a regular expression which can match the timestamps you wish to extract; note that it must also be able to match the `Timestamp-Format-String` and will return an error if it does not match.

The format defined using these options will be inserted at the top of the list of formats used by timegrinder, meaning it will be checked first, but if a valid timestamp is not found with the user-defined format, the rest of the timegrinder formats will also be tried.
