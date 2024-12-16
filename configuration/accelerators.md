# Gravwell Accelerators

Gravwell can process entries as they are *ingested* in order to perform field extraction.  The extracted fields are then recorded in acceleration blocks which accompany each shard.  Using the accelerators can enable dramatic speedups in throughput with minimal storage overhead.  Accelerators are specified on a per-well basis and are designed to be as unobtrusive and flexible as possible.  If data enters a well that does not match the acceleration directive, or is missing the specified fields, Gravwell processes it just like any other entry.  Acceleration will engage when it can.

We refer to "accelerators" and "acceleration" rather than "indexers" and "indexing" for two reasons. First, Gravwell already has a very important component called an "indexer". Second, acceleration can be done by direct indexing **or** with a bloom filter, so describing about an "index" is not necessarily accurate.

## Acceleration Basics

Gravwell accelerators use a filtering technique that works best when data is relatively unique.  If a field value is extremely common, or present in almost every entry, it doesn't make much sense to include it in the accelerator specification.  Specifying and filtering on multiple fields can also improve accuracy, which improves query speed.  Fields which make good candidates for acceleration are fields that users will query for directly.  Examples include process names, usernames, IP addresses, module names, or any other field that will be used in a needle-in-the-haystack type query.

Tags are always included in the acceleration, regardless of the extraction module in use.  Even when the query does not specify inline filters, the acceleration system can help narrow down and accelerate queries when there are multiple tags in a single well.

Most acceleration modules incur about a 1-1.5% storage overhead when using the bloom engine, but extremely low-throughput wells may consume more storage.  If a well typically sees about 1-10 entries per second, acceleration may incur a 5-10% storage penalty, where a well with 10-15 thousand entries per second may see as little as 0.5% storage overhead.  Gravwell accelerators also allow user-specified collision rate adjustments.  If you can spare the storage, a lower collision rate may increase accuracy and speed up queries while increasing storage overhead.  Reducing the accuracy reduces the storage penalty but decreases accuracy and reduces the effectiveness of the accelerator.  The index engine will consume significantly more space depending on the number of fields extracted and the variability of the extracted data.  For example, full text indexing may cause the accelerator files to consume as much space as the stored data files.

Accelerators must operate on the direct data portion of an entry (with the exception of the src accelerator which directly operates on the SRC field).

## Acceleration Engines

The engine is the system that actually stores the extracted acceleration data.  Gravwell supports two acceleration engines. Each engine provide different benefits depending on desired ingest rates, disk overhead, search performance, and data volumes.  The acceleration engine is entirely independent from the accelerator extractor itself (as specified with the Accelerator-Name configuration option).

The default engine is the "bloom" engine.  The bloom engine uses bloom filters to indicate whether or not a piece of data exists in a given block.  The bloom engine typically has very little disk overhead and works well with needle-in-haystack style queries, for example finding logs where a specific IP showed up.  The bloom engine performs poorly on filters where filtered entries occur regularly.

The "index" engine is a full indexing system designed to be fast across all query types.  The index engine typically consumes considerably more disk space than the bloom engine but is significantly faster when operating on very large data volumes or queries that may touch a significant portion of the total data.  It is not uncommon for the index engine to consume as much space as the compressed data in heavily-indexed systems.



### Optimizing the Index Engine

The "index" uses a file-backed data structure to store and query key data. The file backing is performed using memory maps, which can be pretty abusive when the kernel is too eager to write back dirty pages.  It is highly recommended that you tune the kernel's dirty page parameters to reduce the frequency that the kernel writes back dirty pages.  This is done via the "/proc" interface and can be made permanent using the "/etc/sysctl.conf" configuration file.  The following script will set some efficient parameters and ensure they stick across reboots.

```
#!/bin/bash
user=$(whoami)
if [ "$user" != "root" ]; then
	echo "must run as root"
fi

echo 70 > /proc/sys/vm/dirty_ratio
echo 60 > /proc/sys/vm/dirty_background_ratio
echo 2000 > /proc/sys/vm/dirty_writeback_centisecs
echo 3000 > /proc/sys/vm/dirty_expire_centisecs

echo "vm.dirty_ratio = 70" >> /etc/sysctl.conf
echo "vm.dirty_background_ratio = 60" >> /etc/sysctl.conf
echo "vm.dirty_writeback_centisecs = 2000" >> /etc/sysctl.conf
echo "vm.dirty_expire_centisecs = 3000" >> /etc/sysctl.conf

```

## Configuring Acceleration

Accelerators are configured on a per-well basis.  Each well can specify an acceleration module, fields for extraction, a collision rate, and the option to include the entry source field.  If you commonly filter on specific sources (e.g. only look at syslog entries coming from a specific device) including the source field provides an effective way to boost accelerator accuracy independent of the fields being extracted.

| Acceleration Parameter | Description | Example |
|----------|------|-------------|
| Accelerator-Name  | Specifies the field extraction module to use at ingest. | Accelerator-Name="json" |
| Accelerator-Args  | Specifies arguments for the acceleration module, typically the fields to extract. | Accelerator-Args="username hostname appname" |
| Collision-Rate | Controls the accuracy for the acceleration modules using the bloom engine.  Must be between 0.1 and 0.000001. Defaults to 0.001. | Collision-Rate=0.01
| Accelerator-Engine-Override | Specifies the engine to use for indexing.  By default the index engine is used. | Accelerator-Engine-Override=index

### Supported Extraction Modules

* [CSV](/search/csv/csv)
* [Fields](/search/fields/fields)
* [Syslog](/search/syslog/syslog)
* [JSON](/search/json/json)
* [CEF](/search/cef/cef)
* [Regex](/search/regex/regex)
* [Winlog](/search/winlog/winlog)
* [Slice](/search/slice/slice)
* [Netflow](/search/netflow/netflow)
* [IPFIX](/search/ipfix/ipfix)
* [Packet](/search/packet/packet)
* Fulltext

### Example Configuration

Below is an example configuration which extracts the 2nd, 4th, and 5th field in a tab-delimited entry, for example a line from a bro log file.  In this example we are extracting and accelerating on the source ip, destination ip, and destination port from each bro connection log.  All entries which enter the "bro" well (which contains only the tag "bro" for this example) will pass through the extraction module during ingest.  If a piece of data does not conform to the acceleration specification, it will be stored but not accelerated; it will be included in the query, but if many nonconforming entries are in the well, queries will be much slower.

```
[Storage-Well "bro"]
	Location=/opt/gravwell/storage/bro
	Tags=bro
	Accelerator-Name="fields"
	Accelerator-Args="-d \"\t\" [2] [4] [5]"
	Collision-Rate=0.0001
```

## Acceleration Basics

Each acceleration module uses the same syntax as their companion search module for basic field extraction.  Accelerators do not support renaming, filtering, or operating on enumerated values.  They are the first-level filter.  Acceleration modules are transparently invoked whenever the corresponding search module operates and performs an equality filter.

For example, consider the following well configuration which uses the JSON accelerator:

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname app.field1 app.field2"
```

If we were to issue the following query:

```gravwell
tag=app json username==admin app.field1=="login event" app.field2 != "failure" | count by hostname | table hostname count
```

The json search module will transparently invoke the acceleration framework and provide a first-level filter on the "username" and "app.field1" extracted values.  The "app.field2" field is NOT accelerated in this query because it does not use a direct equality filter.  Filters that exclude, compare, or check for subsets are not eligible for acceleration.

(accelerating_specific_tags)=
### Accelerating Specific Tags

The acceleration system allows acceleration at the well or tag levels. This allows you to specify a basic acceleration scheme on a well then specify specific accelerator configurations for specific tags or groups of tags.

Per tag acceleration is enabled by including one or more `Tag-Accelerator-Definitions` in the `[global]` configuration block in your `gravwell.conf`.  The `Tag-Accelerator-Definitions` configuration parameter should point to a file containing `Tag-Accelerator` blocks.  The `Tag-Accelerator` blocks are used to specify a set of tags and an accelerator configuration for those specific tags.

For example, lets look at a definition where a well has a default Acceleration schema (or none at all) and several tags are singled out.  In this example we are going to define two wells in addition to the default well.  We will then include an accelerator definition file that will specify specific accelerators for tags.

Note - Multiple `Tag-Accelerator-Definitions` may be specified to include multiple files.

#### gravwell.conf

```
[global]
	Ingest-Auth=foo
	Control-Auth=bar
	Ingest-Port=4023
	TLS-Ingest-Port=0
	Log-Level=ERROR

	Tag-Accelerator-Definitions="/opt/gravwell/etc/accel.defs"

[Default-Well]
	Location="/opt/gravwell/storage/default"

[Storage-Well "test"]
	Location="/opt/gravwell/storage/test"
	Tags=test*
	Accelerator-Name="fulltext"
	Accelerator-Args="-ignoreTS -ignoreUUID"

[Storage-Well "raw"]
	Location="/opt/gravwell/storage/raw"
	tags=raw*
	tags=pcap*
```

#### accel.defs

```
[Tag-Accelerator "csv things"]
	Accelerator-Name=csv
	Accelerator-Args="[0] [1] [2] [3] [4] [5] [6]"
	Tags=test1
	Tags=testspecial*
	Tags=firewall

[Tag-Accelerator "packet stuff"]
	Accelerator-Name=packet
	Accelerator-Args="ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort"
	Tags=pcap
```

The `Tag-Accelerator` definitions support tag wildcards in the same way as a well, but a `Tag-Accelerator` specification does NOT have be the same as a well specification.

Lets build out a table to see how specific tags would be mapped to specific wells and accelerators:

| Tag      | Well Assignment | Accelerator                   |
|----------|-----------------|-------------------------------|
| default  | `default`       | NONE                          |
| firewall | `default`       | `csv`                         |
| test1    | `test`          | `csv`                         |
| testfoo  | `test`          | `fulltext`                    |
| pcap     | `raw`           | `packet`                      |


### Accelerator Tag Match Rules

Accelerator tag definitions CAN overlap with some specific rules.  This is in contrast to the well assignment rules where a single tag may not match two different specifications.  Tags are matched against an accelerator definition using either a `hard` match or a `soft` match.  A `hard` match occurs when the tag is directly specified without any wildcards, a `soft` match occurs when a tag is matched using a wildcard globbing pattern.  For example the tag `foobar` would match the tag specification `foo*` via a soft match, if we directly specified the tag `foobar` then it would be a hard match.

The rules for matching tags are that a tag may not match multiple `soft` specifications or have multiple `hard` matches.  Let's look at a specification which would violate these matching rules and cause undefined behavior:

```
[Tag-Accelerator "csv things"]
	Accelerator-Name=csv
	Accelerator-Args="[0] [1] [2] [3] [4] [5] [6]"
	Tags=test1
	Tags=foobar*

[Tag-Accelerator "packet stuff"]
	Accelerator-Name=packet
	Accelerator-Args="ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort"
	Tags=foo*
```

Notice that the two accelerator definitions have both a `foobar*` and a `foo*` tag matching pattern, this means that if we established the tag `foobarbaz` it could match both specifications (two `soft` matches).  Gravwell will send a notification indicating that the tag matches overlap.

Now lets look at a specification where there is an allowed overlap:

```
#zeekconn is by far the most noisy, set a special accelerator for it
[Tag-Accelerator "zeek conn"]
	Accelerator-Name=fields
	Accelerator-Args=`-d "\t" [1] [2] [3] [4] [5] [6] [7] [11] [15]`
	Tags=zeekconn

# All other zeek logs can use a tuned fulltext accelerator
[Tag-Accelerator "zeek"]
	Accelerator-Name=fulltext
	Accelerator-Args=`-ignoreFloat -min 2`
	Tags=zeek* #apply to all zeek prefixed tags
```

Note that the tag `zeekconn` can be matched against both accelerators, however the accelerator definition `zeek conn` has a specific `hard` match where the more generic `zeek` definition would match with a `soft` match.  Because the tag has a hard and soft match, the specifications are legal and the `zeekconn` tag will be assigned the appropriate accelerator.  The `hard` and `soft` matching rules make it convenient to specify and acceleration for subset of tags and then tailor accelerators for specific high volume tags as is the case for Zeek data.


`Tag-Accelerator` definitions can be included as external files or directly in the `gravwell.conf` file.  Here is a valid `gravwell.conf` that specifies a few accelerator deviations as well as two external definition files:

```
[global]
	Ingest-Auth=foo
	Control-Auth=bar
	Ingest-Port=4023
	TLS-Ingest-Port=0
	Log-Level=ERROR

	Tag-Accelerator-Definitions="/opt/gravwell/etc/syslog.defs"
	Tag-Accelerator-Definitions="/opt/gravwell/etc/zeek.defs"

[Default-Well]
	Location="/opt/gravwell/storage/default"

[Storage-Well "test"]
	Location="/opt/gravwell/storage/test"
	Tags=test*
	Accelerator-Name="fulltext"
	Accelerator-Args="-ignoreTS -ignoreUUID"

[Tag-Accelerator "csv things"]
	Accelerator-Name=csv
	Accelerator-Args="[0] [1] [2] [3] [4] [5] [6]"
	Tags=test1
	Tags=foobar*

[Tag-Accelerator "packet stuff"]
	Accelerator-Name=packet
	Accelerator-Args="ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort"
	Tags=foo*
```

(intrinsic-acceleration-target)=
## Acceleration with Intrinsic Enumerated Values

When acceleration is enabled, [intrinsic enumerated values](#attach-target) will always be accelerated with the fulltext engine. This enables queries using the [intrinsic](/search/intrinsic/intrinsic) module to be accelerated. No specific configuration is required for acceleration with intrinsic EVs other than having acceleration enabled.

## Fulltext

The fulltext accelerator is designed to index words within text logs and is considered the most flexible acceleration option.  Many of the other search modules support invoking the fulltext accelerator when executing queries.  However, the primary search module for engaging with the fulltext accelerator is the [grep](/search/grep/grep) module with the `-w` flag.  Much like the Unix grep utility, `grep -w` specifies that the provided filter is expected to a word, rather than a subset of bytes.  Running a search with `words foo bar baz` will look for the words foo, bar, and baz and engage the fulltext accelerator.

While the fulltext accelerator may be the most flexible, it is also the most costly.  Using the fulltext accelerator can significantly reduce the ingest performance of Gravwell and can consume significant storage space, this is due to the fact that the fulltext accelerator is indexing on virtually every component of every entry.

### Fulltext Arguments

The fulltext accelerator supports a few options for refining the types of data that are indexed and removing fields that incur significant storage overhead but may not help much at query time.

| Argument | Description | Example | Default State |
|----------|-------------|---------|---------------|
| -acceptTS | By default the fulltext accelerator attempts to identify and ignore timestamps in the data. This flag disables that behavior and allows timestamps to be indexed. | `-acceptTS` | DISABLED |
| -acceptFloat | By default, the fulltext accelerator attempts to identify and ignore floating-point numbers, since they typically vary widely and are not explicitly queried. Setting this flag disables that behavior and allows floating-point numbers to be indexed. | `-acceptFloat` | DISABLED |
| -min | Require that extracted tokens be at least X bytes long.  This can help prevent indexing on very small words such as "is" or "I". | `-min 3` | DISABLED |
| -max | Require that extracted tokens be less than X bytes long.  This can help prevent indexing on very large "blobs" within logs that will never be feasibly queried. | `-max 256` | DISABLED |
| -ignoreUUID | Enable a filter to ignore UUID/GUID values.  Some logs will generate a UUID for every entry, which incurs significant indexing overhead and provides very little value. | `-ignoreUUID` | DISABLED |
| -ignoreTS | Identify and ignore timestamps during acceleration. Because timestamps change so frequently, they can be a significant source of bloat. The fulltext accelerator ignores timestamps by default. | `-ignoreTS` | ENABLED |
| -ignoreFloat | Ignore floating point numbers. Logs where floating point numbers are used for filters can make use of `-accptFloat`. | `-acceptFloat` | ENABLED |
| -maxInt | Enables a filter that will only index integers below a certain size.  This can be valuable when indexing data that such as HTTP access logs.  You want to index the return codes, but not the response times and data sizes. | `-maxInt 1000` | DISABLED |

```{note}
Make sure you understand your data before enabling the `-acceptTS` and `-acceptFloat` flags as these can dramatically bloat the index when using the index engine.  The Bloom engine is less impacted by orthogonal data such as timestamps and floating point numbers.
```

### Fulltext word extraction

The Fulltext accelerator indexes words within text logs. It does this by extracting any word that is surrounded by any of the below non-word characters, or the beginning or end of the text. For example, the message `foo%bar` will extract "foo" and "bar" in the same way as `foo bar`, since `%` is a split character. 

The following table lists all the split characters used by the Fulltext accelerator.

| Character | Unicode escape code |
| --------- | ------------------- |
| Unprintable ASCII control characters (ie NUL, backspace) | \u0000 - \u001F |
| Space | \u0020 |
| ! | \u0021 | 
| " | \u0022 |
| # | \u0023 |
| $ | \u0024 |
| % | \u0025 |
| & | \u0026 |
| ' | \u0027 |
| ( | \u0028 |
| ) | \u0029 |
| * | \u002A |
| + | \u002B |
| , | \u002C |
| / | \u002F |
| ; | \u003B |
| < | \u003C |
| = | \u003D |
| > | \u003E |
| ? | \u003F |
| [ | \u005B |
| \ | \u005C |
| ] | \u005D |
| ^ | \u005E |
| ` | \u0060 |
| { | \u007B |
| &vert; | \u007C |
| } | \u007D |
| ~ | \u007E |
| DEL | \u007F |
| NSBP | \u00A0 |
| × | \u00D7 |
| Unprintable Unicode whitespace characters | \u2000 - \u200a, \u1680, \u2028, \u2029, \u202f, \u205f, \u3000, \u0085 |


### Example Well Configuration

The following well configuration performs fulltext acceleration using the `index` engine.  We are attempting to identify and ignore timestamps, UUIDs, and require that all tokens be at least 2 bytes in length.

```
[Default-Well]
	Location=/opt/gravwell/storage/default
	Accelerator-Name="fulltext"
	Accelerator-Args="-ignoreTS -ignoreUUID -min 2"
```

## JSON

The JSON accelerator module is specified using the accelerator name "json" and uses the exact same syntax for picking fields as the JSON modules.  See the [JSON search module](/search/json/json) section for more information on field extraction.

### Example Well Configuration

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname \"strange-field.with.specials\".subfield"
```

## Syslog

The syslog accelerator is designed to operate on conforming RFC5424 syslog messages.  See the [syslog search module](/search/syslog/syslog) section for more information on field extraction.

### Example Well Configuration

```
[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Tags=syslog
	Accelerator-Name="syslog"
	Accelerator-Args="Hostname Appname MsgID valueA valueB"
```

## CEF

The CEF accelerator is designed to operate on CEF log messages and is just as flexible as the search module.  See the [CEF search module](/search/cef/cef) section for more information on field extraction.

### Example Well Configuration

```
[Storage-Well "ceflogs"]
	Location=/opt/gravwell/storage/cef
	Tags=app1
	Accelerator-Name="cef"
	Accelerator-Args="DeviceVendor DeviceProduct Version Ext.Version"
```

## Fields

The fields accelerator can operate on any delimited data format, whether it be CSV, TSV, or some other delimiter.  The fields accelerator allows you to specify the delimiter the same way as the search module.  See the [fields search module](/search/fields/fields) section for more information on field extraction.

### Example Well Configuration

This configuration extracts four fields from a comma-separated entry. Note the use of the `-d` flag to specify the delimiter.

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
```

## CSV

The CSV accelerator is designed to operate on comma-separated value data, automatically removing surrounding white space and double quotes from data.  See the [CSV search module](/search/csv/csv) section for more information on column extraction.

### Example Well Configuration

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="csv"
	Accelerator-Args="[1] [2] [5] [3]"
```

## Regex

The regex accelerator allows complicated extractions at ingest time in order to handle non-standard data formats.  Regular expressions are one of the slower extraction formats, so accelerating on specific fields can greatly increase query performance.

### Example Well Configuration

```
[Storage-Well "webapp"]
	Location=/opt/gravwell/storage/webapp
	Tags=webapp
	Accelerator-Name="regex"
	Accelerator-Args="^\\S+\\s\\[(?P<app>\\w+)\\]\\s<(?P<uuid>[\\dabcdef\\-]+)>\\s(?P<src>\\S+)\\s(?P<srcport>\\d+)\\s(?P<dst>\\S+)\\s(?P<dstport>\\d+)\\s(?P<path>\\S+)\\s"
```

```{note}
Remember to escape backslashes '\\' when specifying regular expressions in the gravwell.conf file.  The regex argument '\\w' will become '\\\\w'.
```

```{warning}
The regex accelerator requires very precise matches during query, this means that regular expression specified in an accelerator must match the regular expression provided in a query bit-for-bit.  It is not possible to determine if two regular expressions are equivalent, so Gravwell checks that the regular expression strings are exact matches before engaging acceleration.  The regex accelerator is rarely what you want.
```

## Winlog

The winlog module is perhaps *the* slowest search module.  The complexity of XML data combined with the Windows log schema means that the module has to be extremely verbose, resulting in pretty poor performance.  This means that accelerating Windows log data may be your single most important performance optimization, as processing millions or billions of unaccelerated entries with the winlog module will be excruciatingly slow.  The accelerators help you narrow down the specific log entries you want without invoking the winlog search module on every piece of data.  However, accelerating winlog data simply shifts processing from search time to ingest time, meaning that ingest of Windows logs will be slower when acceleration is enabled, so don't expect Gravwell's typical ingest rate of hundreds of thousands of entries per second when ingesting into a winlog-accelerated well.

### Example Well Configuration

```
[Storage-Well "windows"]
	Location=/opt/gravwell/storage/windows
	Tags=windows
	Accelerator-Name="winlog"
	Accelerator-Args="EventID Provider Computer TargetUserName SubjectUserName"
```

```
The winlog accelerator is permissive ('-or' flag is implied).  So specify any field you plan on using to filter searches on, even if two of the fields would not be present in the same entry.
```

## Netflow

The [netflow](/search/netflow/netflow) module accelerates on Netflow V5 fields and speeding up queries on large amounts of netflow data.  While the `netflow` module is very fast and the data is extremely compact, it can still be beneficial to engage acceleration if you have very large Netflow data volumes.  The `netflow` module can use any of the direct Netflow fields, but cannot use the pivot helper fields.  This means that you must specify `Src` or `Dst` and not `IP`.  The `IP` and `Port` fields cannot be specified in the acceleration arguments.

```{note}
The helper extractions `Timestamp` and `Duration` cannot be used in accelerators.
```

### Example Well Configuration

This example configuration uses the `bloom` engine and is accelerating on the source and destination IP/Port pairs as well as the protocol.

```
[Storage-Well "netflow"]
	Location=/opt/gravwell/storage/netflow
	Tags=netflow
	Accelerator-Name="netflow"
	Accelerator-Args="Src Dst SrcPort DstPort Protocol"
	Accelerator-Engine-Override="bloom"
```

## IPFIX

The [ipfix](/search/ipfix/ipfix) module can accelerate queries on IPFIX-formatted records. This module can accelerate on any of the 'normal' IPFIX fields, but not pivot helper fields. This means you must specify `sourceTransportPort` or `destinationTransportPort` rather than `port`, or `src`/`dst` rather than `ip`.

### Example Well Configuration

This example configuration uses the `index` engine to accelerate on source/destination ip/port pairs as well as the IP protocol of the flow, comparable to the example shown in the netflow section.

```
[Storage-Well "ipfix"]
	Location=/opt/gravwell/storage/ipfix
	Tags=ipfix
	Accelerator-Name="ipfix"
	Accelerator-Args="src dst sourceTransportPort destinationTransportPort protocolIdentifier"
	Accelerator-Engine-Override=index
```

## Packet

The [packet](/search/packet/packet) module can accelerate raw packet captures using the same syntax as the search module of the same name.  There is a subtle but important difference in how the packet accelerator is applied as compared to the search modules; the accelerator can use overlapping layers.  This means that you can specify both UDP and TCP items and extract the right field depending on the packet being processed.

A well configuration can be configured to accelerate IPv4, IPv6, TCP, UDP, ICMP, etc... all at the same time.  The packet accelerator does not treat specified fields as implied filters.

The packet accelerator also requires direct fields, this means you cannot use the convenience extractors like `IP` and `Port`.  You must specify exactly what you want to accelerate on.

### Example Well Configuration

```
[Storage-Well "packets"]
	Location=/opt/gravwell/storage/pcap
	Tags=pcap
	Accelerator-Name="packet"
	Accelerator-Args="ipv4.SrcIP ipv4.DstIP ipv6.SrcIP ipv6.DstIP tcp.SrcPort tcp.DstPort udp.SrcPort udp.DstPort"
```

## SRC

The src accelerator can be used when only the entry's source field should be accelerated. See the [src search module](/search/src/src) for more information on filtering.

### Example Well Configuration

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="src"
```

### Example Well Configuration and Query Combining SRC

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
```

The following query invokes both the fields accelerator and the src accelerator to specify specific log types coming from specific sources.

```gravwell
tag=app src dead::beef | fields -d "," [1]=="security" [2]="process" [5]="domain" [3] as processname | count by processname | table processname count
```

## Acceleration Performance and Benchmarking

To understand the benefits and drawbacks of acceleration it is best to see how it impacts storage use, ingest performance, and query performance.  We will use some Apache combined access logs that are generated using a generator available on [github](https://github.com/kiritbasu/Fake-Apache-Log-Generator).  Our data set is 10 million Apache combined access logs that are spread out over approximately 24 hours; the total data is 2.1GB.  We will define 4 wells with 4 different configurations.  We will be taking a fairly naive approach to indexing this data, as there are many parameters that don't make a lot of sense to index on, such as the number of returned bytes.



| Well | Extractor | Engine | Description |
|------|-----------|--------|-------------|
| raw  | None | None | A completely raw well with no acceleration at all |
| fulltext | fulltext | index | A fulltext accelerated well that uses the index engine and will perform fulltext acceleration on every word |
| regexindex | regex | index | A well accelerated with the regex extractor and using the index engine.  Each parameter is extracted and indexed |
| regexbloom | regex | bloom | A well with the same extractor as the regexindex well but with bloom engine.  Each parameter is extracted and added to the bloom filter |

The well configurations are:

```
[Storage-Well "raw"]
	Location=/opt/gravwell/storage/raw
	Tags=raw
	Enable-Transparent-Compression=true

[Storage-Well "fulltextbloom"]
	Location=/opt/gravwell/storage/fulltextbloom
	Tags=fulltextbloom
	Enable-Transparent-Compression=true
	Accelerator-Name=fulltext
	Accelerator-Args="-ignoreTS -min 2"
	Accelerator-Engine-Override=bloom

[Storage-Well "fulltextindex"]
	Location=/opt/gravwell/storage/fulltextindex
	Tags=fulltextindex
	Enable-Transparent-Compression=true
	Accelerator-Name=fulltext
	Accelerator-Args="-ignoreTS -min 2"
	Accelerator-Engine-Override=index

[Storage-Well "regexindex"]
	Location=/opt/gravwell/storage/regexindex
	Tags=regexindex
	Enable-Transparent-Compression=true
	Accelerator-Name=regex
	Accelerator-Engine-Override=index
	Accelerator-Args=`^(?P<ip>\S+) (?P<ident>\S+) (?P<username>\S+) \[([\w:/]+\s[+\-]\d{4})\] \"(?P<method>\S+)\s?(?P<url>\S+)?\s?(?P<proto>\S+)?\" (?P<resp>\d{3}|-) (?P<bytes>\d+|-)\s?\"?(?P<referer>[^\"]*)\"?\s?"?(?P<useragent>[^\"]*)?\"?$`

[Storage-Well "regexbloom"]
	Location=/opt/gravwell/storage/regexbloom
	Tags=regexbloom
	Enable-Transparent-Compression=true
	Accelerator-Name=regex
	Accelerator-Engine-Override=bloom
	Accelerator-Args=`^(?P<ip>\S+) (?P<ident>\S+) (?P<username>\S+) \[([\w:/]+\s[+\-]\d{4})\] \"(?P<method>\S+)\s?(?P<url>\S+)?\s?(?P<proto>\S+)?\" (?P<resp>\d{3}|-) (?P<bytes>\d+|-)\s?\"?(?P<referer>[^\"]*)\"?\s?"?(?P<useragent>[^\"]*)?\"?$`
```

### Test Machine

Query, ingest, and storage performance characteristics will vary with each data set and hardware platform, but for this test we are using the following hardware:

| Component | Description |
|-----------|-------------|
| CPU       | AMD Ryzen 3700X |
| Memory    | 16GB DDR4-3200 |
| Disk      | Samsung 960 EVO NVME |
| Filesystem | BTRFS with zstd transparent compression at level 10 |

These tests were conducted using Gravwell version `5.1.0` on a Linux 5.13.5 kernel.

### Ingest Performance

For ingest we will use the singleFile ingester.  The singleFile ingester is designed to ingest a single newline delimited file, deriving timestamps as it goes.  Because the ingester is deriving timestamps, it requires some CPU resources.  The singleFile ingester is available on our [GitHub page](https://github.com/gravwell/ingesters/). The exact invocation of the singleFile ingester is:

```
./singleFile -i apache_fake_access_10M.log -clear-conns 172.17.0.2 -block-size 1024 -timeout=10 -tag-name=fulltext
```

|  Well      | Entries Per Second | Data Rate  |
|------------|--------------------|------------|
| raw             | 342.65 KE/s   | 72.06 MB/s |
| regexbloom      | 157.38 KE/s   | 33.10 MB/s |
| regexindex      | 81.92  KE/s   | 17.23 MB/s |
| fulltextbloom   | 246.99  KE/s  |  51.94 MB/s |
| fulltextindex   | 95.72  KE/s   |  20.13 MB/s |

We can see from the ingest performance that the index acceleration engine reduces ingest performance, this is due to the increased complexity of indexing and additional storage usage.  However the index engine can often be much faster at query time for needle-in-haystack type queries.

### Storage Usage

Outside of ingest performance and some additional memory requirements, the main penalty to enabling acceleration is storage usage.  We can see that the index engine for each extraction methodology consumed over 50% more storage, while the bloom engine consumed an additional 4%.  The storage usage is very dependent on data consumed, but on average the indexing system will consume significantly more storage.

|  Well      | Used Storage | Acceleration Storage Overhead |
|------------|--------------|---------------|
| raw             | 2.43GB       | 0%       |
| regexbloom      | 2.45GB       | 4.6%     |
| fulltextbloom   | 2.47GB       | 8.7%      |
| regexindex      | 3.79GB       | 60%      |
| fulltextindex   | 4.04GB       | 70%      |

### Query Performance

To demonstrate the differences in query performance we will execute two queries which can be categorized as sparse and dense.  The sparse query will look for a specific IP in the data set, returning just a handful of entries.  The dense query will look for a specific URL that is reasonably common in the data set.  To simplify the queries we will install an ax module for the regexindex and regexbloom tags that matches the acceleration system.  The dense query will retrieve roughly 12% of the entries in the data set, while the sparse query will retrieve approximately 0.01%.

The sparse and dense queries are:

```
ax ip=="106.218.21.57"
ax url=="/list" method | count by method | chart count by method
```

Prior to executing each query, we will drop the system caches by executing the following command as root:

```
echo 3 > /proc/sys/vm/drop_caches
```

|  Well      | Query Type | Query Time | Processed Entries | Speedup |
|------------|------------|------------|------------------------------|---------|
| raw             | sparse     | 51.8s |  10M              | 0X      |
| regexbloom      | sparse     | 0.259s |                  | 208     |
| regexindex      | sparse     | 0.144s | 1                | 359X    |
| fulltextindex   | sparse     | 0.182s  | 1               | 284X    |
| fulltextbloom   | sparse     | 0.339s | 90               | 152X    |
| raw             | dense      | 52.9s | 10M               | 0X      |
| regexbloom      | dense      | 52.6s | 10M               | 0X   |
| regexindex      | dense      | 17.7s | 1.3M              | 2.92X   |
| fulltextindex   | dense      | 30.2s | 3M                | 1.71X   |
| fulltextbloom   | dense      | 20.5s | 10M               | 2.52X   |

You can see that the dense query did not benefit from the bloom filter at all due to the high collision rate and regular occurrence of the filtered data where the index system was able to make modest gains by limiting the amount of data that had to be extracted from the disk.

```{note}
The regex search module/autoextractor is not fully compatible with the fulltext accelerator, so we have to modify the queries slightly to engage the accelerators.  They are ```words "106.218.21.57"``` and ```words list | ax url=="/list" method | count by method | chart count by method```
```

#### Fulltext

The above benchmarks make it very apparent that the fulltext accelerator can have significant ingest and storage penalties and the dense example query does not appear to justify those expenses, especially when paired with the bloom engine.  This is because the word "list" occurs so regularly that the index and bloom engines do very little to help.  However, if your data has complex components the fulltext accelerator can do things the other accelerators cannot.  We have been using Apache combined access logs, lets look at a query that allows the fulltext accelerator to really shine.

We are going to look at sub-components of the URL and get a chart of users that are browsing our the `/apps` sub directory using a PowerPC Macintosh computer.  The regular expressions in the above examples index on full fields within the Apache log.  They cannot drill down and use parts of those fields for acceleration, fulltext can.

We will optimized the query for both the fulltext indexer and the others so that we can be fair, however both queries will work on either of the data sets.

The fulltext accelerator optimized query:
```
words apps Macintosh PPC | ax url~"/apps" useragent~"Macintosh; U; PPC" | count | chart count
```

The query optimized for non-fulltext:
```
ax url~"/apps" useragent~"Macintosh; U; PPC" | count | chart count
```

The results show why fulltext may often be worth the storage and ingest penalty:

|  Well      | Query Time | Speedup |
|------------|------------|---------|
| raw        | 58s      | 0X      |
| regexbloom | 57s      | ~0X     |
| regexindex | 57s       | ~0X     |
| fulltextindex   | 2.99s      | 12.49X  |
| fulltextbloom   | 3.40s      | 12.49X  |

#### Query AX modules

The AX definition file for all four tags is below, see the [AX](/configuration/autoextractors) documentation for more information:

```
[[extraction]]
  tag = 'regexindex'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apacheindex'
  desc = 'apache index'

[[extraction]]
  tag = 'regexbloom'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apachebloom'
  desc = 'apache bloom'

[[extraction]]
  tag = 'fulltext'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apachefulltext'
  desc = 'apache fulltext'

[[extraction]]
  tag = 'raw'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apacheraw'
  desc = 'apache raw'
```
