#Gravwell Accelerators

Gravwell enables for processing entries as they are ingested in order to perform field extraction.  The extracted fields are then processed and placed in acceleration blocks which accompany each shard.  Using the accelerators can enable dramatic speedups in throughput with minimal storage overhead.  Accelerators are specified on a per well basis and are designed to be as unobtrusive and flexible as possible.  If data enters a well that does not match the acceleration directive, or is missing the specified fields, Gravwell processes it just like any other entry.  Acceleration will engage when it can.

## Acceleration Basics

Gravwell accelerators are based on a filtering technique that operates best when data is relatively unique.  If a field value is extremely common, or present in almost every entry, it doesn't make much sense to include it in the accelerator specification.  Specifying and filtering on multiple fields can also improve accuracy which improves query speed.  Fields make good candidates for acceleration are fields that will be directly queried for.  Examples include process names, usernames, IP addresses, module names, or any other field that will be used in a needle-in-the-haystack type query.

Most acceleration modules incur about a 1-1.5% storage overhead when using the bloom engine, but extremely low throughput wells may be much higher.  If a well typically sees about 1-10 entries per second acceleration may incur a 5-10% storage penalty, where a well with 10-15 thousand entries per second will see as little as 0.5% storage overhead.  Gravwell accelerators also allow for user specified collision rate adjustments.  If you can spare the storage a lower collision rate may increase accuracy and speed up queries while increasing storage overhead.  Reducing the accuracy reduces the storage penaly but decreases accuracy and reduces the effectiveness of the accelerator.  The index engine will consume significantly more space depending on the number of fields extracted and the variability of the extracted data.  For example, full text indexing may cause the accelerator files to consume as much space as the stored data files.

Accelerators must operate on the direct data portion of an entry (with the exception of the src accelerator which directly operates on the SRC field).

## Acceleration Engines

Gravwell supports two acceleration engines; the engine is the system that actually stores the extracted acceleration data.  Each engine provide different benefits depending on designed ingest rates, disk overhead, search performance, and data volumes.  The acceleration engine is entirely independent from the accelerator-name (extraction system).

The default engine is the "bloom" engine.  The bloom engine uses bloom filters to provide an indication of whether or not a piece of data exists in a given block.  The bloom engine typically has very little disk overhead and works well with needle-in-haystack style queries, an example might be finding logs where a specific IP showed up.  The bloom engine performs poorly on filters where filtered entries occur regularly.  The bloom engine is a poor choice when combined with the fulltext accelerator.

The "index" engine is a full indexing system designed to be fast across all query types.  The index engine typically consumes considerably more disk space than the bloom engine but is significantly faster when operating on very large data volumes or queries that may touch a significant portion of the total data.  It is not uncommon for the index engine to consume as much space as the compressed data in heavily indexed systems.

### Optimizing the Index Engine

The "index" uses a file-backed data structure to store and query key data, the file-backing is performed using memory maps which can be pretty abusive when the kernel is too eager to write back dirty pages.  It is highly reccomended that you tune the kernel dirty page parameters to reduce the frequency that the kernel writes back dirty pages.  This is done via the "/proc" interface and can be made permanent using the "/etc/sysctl.conf" configuration file.  The following script will set some efficient parameters and ensure they stick across reboots.

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

Accelerators are configured on a per well basis.  Each well can specify an acceleration module, fields for extraction, a collision rate, and the option to include the entry source field.  If it is commonplace to filter on specific sources (e.g. only look at syslog entries coming from a specific device) including the source field provides an effective way to boost accelerator accuracy independent of the fields being extracted.

| Acceleration Parameter | Description | Example |
|----------|------|-------------|
| Accelerator-Name  | Specifies the field extraction module to use at ingest | Accelerator-Name="json" |
| Accelerator-Args  | Specifies arguments for the acceleration module, often the fields to extract | Accelerator-Args="username hostname appname" |
| Collision-Rate | Controls the accuracy for the acceleration modules using the bloom engine.  Must be between 0.1 and 0.000001. Defaults to 0.001. |
| Accelerate-On-Source | Specifies that the SRC field of each module should be included.  This allows combining a module like CEF with SRC. |
| Accelerate-Engine-Override | Specifies the engine to use for indexing.  By default the bloom engine is used. |

### Supported Extraction Modules

* [CSV](search/csv/csv.md)
* [Fields](search/fields/fields.md)
* [Syslog](search/syslog/syslog.md)
* [JSON](search/json/json.md)
* [CEF](search/cef/cef.md)
* [Regex](search/regex/regex.md)
* [Winlog](search/winlog/winlog.md)
* [Slice](search/slice/slice.md)

### Example Configuration

Below is an example configuration which extracts the 2nd, 4th, and 5th field in a tab delimited data stream like bro.  In this example we are extracting and accelerating on the source ip, destination ip, and destination port from each bro log.  All entries which enter "bro" well (which is only the tag bro for this example) will pass through the extraction module during ingest.  If a piece of data does not conform to the 

```
[Storage-Well "bro"]
	Location=/opt/gravwell/storage/bro
	Tags=bro
	Accelerator-Name="fields"
	Accelerator-Args="-d \"\t\" [2] [4] [5]"
	Accelerate-On-Source=true
	Collision-Rate=0.0001
```

## Acceleration Basics

Each acceleration module uses the same syntax as their companion search module for basic field extraction.  Accelerators do not support renaming, filtering, or operating on enumerated values.  They are the first level filter.  Acceleration modules are transparently invoked whenever the corresponding search module operates and performs an equality filter.

For example, consider the following well configuration which uses the JSON accelerator.

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname app.field1 app.field2"
```

If we were to issue the following query:

```
tag=app json username==admin app.field1=="login event" app.field2 != "failure" | count by hostname | table hostname count
```

The json search module will transparently invoke the acceleration framework and provide a first level filter on teh username and "app.field1" extracted values.  The "app.field2" field is NOT accelerated on because it is not a direct equality filter.  Filters that exclude, compare, or check for subsets are not eligable for acceleration.

### JSON

The JSON accelerator module is specified using via the accelerator name "json" and uses the exact same syntax for picking fields as the JSON modules.  See the [JSON search module](/search/json/json.md) section for more information on field extraction.

#### Example Well Configuration

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname \"strange-field.with.specials\".subfield"
```

### Syslog

The Syslog accelerator is designed to operate on conformant RFC5424 Syslog messages.  See the [Syslog search module](/search/syslog/syslog.md) section for more information on field extraction.

#### Example Well Configuration

```
[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Tags=syslog
	Accelerator-Name="syslog"
	Accelerator-Args="Hostname Appname MsgID valueA valueB"
```

### CEF

The CEF accelerator is designed to operate on CEF log messages and is just as flexible as the search module.  See the [CEF search module](/search/cef/cef.md) section for more information on field extraction.

#### Example Well Configuration

```
[Storage-Well "ceflogs"]
	Location=/opt/gravwell/storage/cef
	Tags=app1
	Accelerator-Name="cef"
	Accelerator-Args="DeviceVendor DeviceProduct Version Ext.Version"
```

### Fields

The fields accelerator can operate on any delimited data format, whether it be CSV, TSV, or any other delimiter.  The Fields accelerator supports specifying the delimiter the same way as the search module.  See the [Fields search module](#!search/fields/fields.md) secion for more informaton on field extraction.

#### Example Well Configuration

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
```

### CSV

The CSV accelerator is designed to operate on comma seperated value data, automatically removing surrounding whitespace and double quotes from data.  See the [CSV search module](#!search/csv/csv.md) secion for more informaton on column extraction.

#### Example Well Configuration

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="csv"
	Accelerator-Args="[1] [2] [5] [3]"
```

### Regex

The regex accelerator allows for specifying complicated extractions at ingest time in order to handle non-standard data formats.  Regular expressions are one of the slower extraction formats, so accelerating on specific fields can greatly increase query performance.

#### Example Well Configuration

```
[Storage-Well "webapp"]
	Location=/opt/gravwell/storage/webapp
	Tags=webapp
	Accelerator-Name="regex"
	Accelerator-Args="^\\S+\\s\\[(?P<app>\\w+)\\]\\s<(?P<uuid>[\\dabcdef\\-]+)>\\s(?P<src>\\S+)\\s(?P<srcport>\\d+)\\s(?P<dst>\\S+)\\s(?P<dstport>\\d+)\\s(?P<path>\\S+)\\s"
```

Attention: Remember to escape backslashes '\\' when specifying regular expressions in the gravwell.conf file.  The regex argument '\\w' will become '\\\\w'

### Winlog

The winlog module is one of if not the slowest extraction module.  The complexity of XML data combined with the Windows log schema means that the extraction module has to be extremely verbose, resulting in pretty poor extraction performance.  As a result, accelerating windows data may be the single most important performance optimization, as processing millions or billions of entries with the winlog module will be excruciatingly slow.  The accelerators help you narrow down the specific log entries you want without invoking the winlog module on every piece of data.  However, the slow extraction rate means that ingest of windows logs will be impacted, so don't expect Gravwell's typical ingest rate of hundreds of thousands of entries per second when ingesting into a winlog accelerated well.

#### Example Well Configuration

```
[Storage-Well "windows"]
	Location=/opt/gravwell/storage/windows
	Tags=windows
	Accelerator-Name="winlog"
	Accelerator-Args="EventID Provider Computer TargetUserName SubjectUserName"
```

Attention: The winlog accelerator is permissive ('-or' flag is implied).  So specify any field you plan on using to filter searches on, even if two of the fields would not be present in the same entry.


### SRC

The SRC accelerator can be used when only the SRC field should be accelerated.  However, its essentially possible to combine the SRC accelerator with other accelerators by enabling the "Accelerate-On-Source" flag and also adding a the src search module.  See the [SRC search module](#!search/src/src.md) for more information on filtering.

#### Example Well Configuration

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="src"
```

#### Example Well Configuration and Query Combining SRC

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
	Accelerate-On-Source=true
```

The following query invokes both the fields accelerator and the SRC accelerator to specify specific log types coming from specific sources.

```
tag=app src dead::beef | fields -d "," [1]=="security" [2]="process" [5]="domain" [3] as processname | count by processname | table processname count
```
