# Migrating Data

A common task when first standing up Gravwell is getting data from an old system or log archive into Gravwell.  Migration may involve pulling historical data out of an existing system, scraping a large database, or just ingesting 5 years of old syslog from flat files.  Gravwell provides a series of ingesters designed to stream data in from a variety of sources in near real-time, but when you need to get massive amounts of historical data in they are typically not the right choice.

This section will explore the available tools to migrate existing data and give some tips on how to efficiently migrate potentially hundreds of TB of existing logs into a new Gravwell instance.  We will examine a few scenarios where using migration tools will provide much better migration performance and storage efficiency as opposed to just tossing something at a typical ingester.  Most of our migration tools are open source, so we can also provide links to source code and deep dive examinations of functionality.  Most data migrations are a one time occurrence which means a one off tool that is specialized for the task is almost always the right answer.  

### Migration Caveats

Most Gravwell licenses are unlimited, meaning that when it is time to migrate massive quantities of data in you are only limited by the resources available to accept and index that data.  However, the one exception is the free Community Edition license which has hard ingest limits.  All other license types are either unlimited or allow for bursting to accommodate data migration.


## One Shot File

The [singleFile](https://github.com/gravwell/gravwell/tree/master/ingesters/singleFile) tool is one of the most simplistic ingesters in the Gravwell arsenal.  It is designed to very easily ingest a single line delimited file to a specific tag in one go.  It can transparently decompress files and has some limited parsing ability.  If you have a single large Apache access log or just need to script up some one off file ingestion, [singleFile](https://github.com/gravwell/gravwell/tree/master/ingesters/singleFile) has your back.

The [singleFile](https://github.com/gravwell/gravwell/tree/master/ingesters/singleFile) ingester is a standalone ingester that is designed to operate using flags rather than a config file.  This means that it lacks some of the additional functionality of other ingesters such as custom timestamp definitions and preprocessor support.  The following flags are supported:


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
| -timeout | | Timeout in seconds the ingester should use when attempting to connect to an indexer. | -timeout=10 |
| -ignore-ts | | Do not attempt to resolve timestamps out of entries, everything is ingested at the time "NOW" | -ignore-ts |
| -tls-remote-verify | | Enable or disable TLS certificate validation when using TLS connections, default is true. | -tls-remote-verify=false |
| -timezone-override | | Override the timezone applied if the ingester cannot resolve a timezone in the timestamp itself. | -timezone-override=UTC |
| -source-override | | Apply a SRC value instead of allowing the indexer to choose one for you.  This allows for manually setting source values. | -source-override="192.168.1.2" |
| -timestamp-override | | Specify a specific timestamp format to use. | -timestamp-override="rfc3339" |
| -utc |          | Assume discovered timestamps are in the UTC timezone | |
| -block-size |   | Ingest data in batches, this can increase throughput. | -block-size 256 |
| -clean-quotes | | When an entry is surrounded in quotes, remove the quotes before ingesting | |
| -quotable-lines | | Lines may be quotable, meaning a newlines encapsulated in quotes do not delimit entries |
| -ignore-prefix | | Specify a string that when found at the start of a line causes the line to be ignored, useful for ignoring headers on CSV files | -ignore-prefix="#" |

[singleFile](https://github.com/gravwell/gravwell/tree/master/ingesters/singleFile) is part of the Gravwell monorepo and is licensed using the BSD 2-Clause license.  Clone it, fork it, modify it, it's yours.  If you have the [Go](https://go.dev/dl/) tool-chain and git installed, you can build and install the singleFile binary using the following commands:

```
git clone https://github.com/gravwell/gravwell.git
cd gravwell/ingesters/singleFile
CGO_ENABLED=0 go build
```

Notice the `CGO_ENABLED=0` environment variable passed to `go build`, this causes the Go tool-chain to build a static binary which can then be run on pretty much any Linux with a 3.2 or better kernel.  singleFile can be built for Linux, Windows, or MacOS in any architecture that [Go supports](https://go.dev/doc/install/source).


## Ingesting Many Flat Files

Migrating flat files is the common case when Gravwell first enters an organization that has no log monitoring at all or is transitioning off of a product where there was significant data that needed to be kept but could not be ingested into their current product.  Common data sources here are Apache access logs, Linux system logs, application logs, and any other application that can drop logs to a file.  These flat files are typically very numerous and potentially compressed, while the file follower ingester *CAN* hand many of these scenarios, its not optimized to do so.  Enter the [migrate](https://github.com/gravwell/gravwell/tree/master/ingesters/migrate) tool.

The migrate tool is designed to allow for defining multiple sources of data and then in a single shot consume from all of them linearly, while keeping track of progress.  This tool is every useful when migrating extremely large sets of data where it the actual migration process may take hours, days, weeks.  Migrate uses a configuration file so that it can support some of the more sophisticated configuration options like custom timestamp formats, recursive file scanning, and selective filename matching.  Migrate can also consume from compressed files, this means that if you have lots of Apache access logs and many of them are compressed using gzip, migrate can consume them without manually decompressing first.  If you have a few large directories of significant log files, migrate is the appropriate tool.

The [migrate](https://github.com/gravwell/gravwell/tree/master/ingesters/migrate) tool combines some functionality out of `singleFile` tool while also enabling some of the functionality of the [file follower](#!ingesters/file_follow.md) ingester.  Specifically, the migrate tool can handle compressed files, this means that if you have old system logs that are compressed with gzip it will transparently decompress and consume those logs; the file follower ingester cannot do that.  However, what this means is that the `migrate` tool must operate at a file level granularity, which means you cannot being migrating a flat file, stop, add more data to the file and pick up where you left off.  The `migrate` tool is ONLY for static flat files that will NOT be changing during your migration.  The `migrate` tool CAN be stopped mid-migration and then continue on where it left off, but it will finish the work on a given file before stopping.  The `state` of of migration is maintained in a file pointed to by the `State-Store-Location` parameter in the global configuration block.

Because the `migrate` tool uses a configuration file it can handle some additional scenarios where the `singleFile` ingester would struggle, namely custom timestamps and preprocessing requirements.  This means that you can define your own custom timestamps and use preprocessors to fix-up, filter, or otherwise modify the data before ingest.  The `migrate` configuration file is broken into two sections, a `[Global]` section containing upstream indexer configuration parameters and several `[Files]` sections which define locations to consume files from.  The `[Global]` section is identical to all other ingesters with the sole addition of the `State-Store-Location` parameter.

Here is an example `migrate.conf``[Global]` section which connects to two indexers over a cleartext connection:

```
[Global]
	Ingest-Secret = IngestSecrets
	Connection-Timeout = 0
	Cleartext-Backend-Target=192.168.1.1:4023 #example of adding a cleartext connection
	Cleartext-Backend-Target=192.168.1.2:4023 #example of adding a cleartext connection
	State-Store-Location=/tmp/migrate.state
	Log-Level=INFO #options are OFF INFO WARN ERROR
	Log-File=/tmp/migrate.log

```


The Migrate tool will use the embedded `gravwell` tag to send status updates and log information about each file it consumes if the `Log-Level` value is set to `INFO`.  This means you can run the query `tag=gravwell syslog Appname==migrate` and see status updates about what the `migrate` tool is did at a later time.

### Migrate Files Configuration

The `migrate` tool can have multiple `Files` configurations which allows the migrate tool to consume from many different directories, applying tags, custom timestamps, and even custom preprocessors to different batches of files.  If you are performing a very large migration from many different file sources you setup a big config file that points at all your sources, fire up the migration and head home for the weekend.  The `migrate` tool will walk each directory and consume each file according to the given configuration.

A `Files` configuration target is defined using the `Files` specifier with a unique name, here is an example looking for files in `/tmp/logs`:

```
[Files "auth"]
	Base-Directory="/tmp/logs/"
	File-Filter="auth.log,auth.log.[0-9]" #we are looking for all authorization log files
	Tag-Name=auth
```

The `Files` configuration element can contain the following elements:


| Config Parameter | Required | Description | Example |
|------------------|----------|-------------|---------|
| Base-Directory   |   X      | A full path to the directory containing flat files to be ingested. | `Base-Directory=/var/log/auth` |
| File-Filter      |   X      | A comma separated list of file glob patterns that specify which files to consume. | `File-Filter="*.log,*.log.gz,file.txt,file.?.txt"` |
| Tag-Name         |   X      | The tag name that the migrate tool should ingest all files into.  This must be a valid tag name. | `Tag-Name=auth` |
| Ignore-Timestamps |         | A Boolean value indicating if the migrate tool should not attempt to resolve timestamps, but should instead the timestamp of now. The default value is false. | `Ignore-Timestamps=true` |
| Assume-Local-Timezone |     | A Boolean indicating that if the resolved timestamps do not have a timezone, use the local timezone. The default is false, meaning use UTC. | `Assume-Local-Timezone=true` |
| Timezone-Override |         | A string indicating that if the resolved timestamps do not have a timezone, use the given timezone. The timezone must be specified in IANA format. | `Timezone-Override="America/New_York"` |
| Recursive |                 | A Boolean indicating that the tool should recurse into any sub-directories found in the `Base-Directory` and attempt to match and consume files. Default is false. | `Recursive=true` |
| Ignore-Line-Prefix |        | A string which indicates that a line should be ignored.  This can be specified multiple times and is useful for dealing with headers on files. | `Ignore-Line-Prefix="#" |
| Preprocessor |              | Specify preprocessors to be applied to entries as they are consumed from flat files.  More than one preprocessor can be specified and they are executed in order. | `Preprocessor="logins"` |


Here is an example config snippet which shows the full range of config options for a directory:

```
[Files "testing"]
	Base-Directory="/tmp/testlogs/"
	File-Filter="app.log,app.log.[0-9],host.log*"
	Tag-Name=apps
	Assume-Local-Timezone=true #Default for assume localtime is false
	Ignore-Timestamps=false
	Recursive=true
	Ignore-Line-Prefix="#"
	Ignore-Line-Prefix="//"
	Preprocessor="loginapps"

[Preprocessor "loginapp"]
	Type=regexextract
	Regex="\\S+ (?P<host>\\S+) \\d+ \\S+ \\S+ (?P<data>\\{.+\\})$"
	Template="{\"host\": \"${host}\", \"data\": ${data}}"
```

### Performing a Migrate

The migrate tool contains the following flags:

| Flag | Required | Description | Example |
|------|----------|-------------|---------|
| -h   |          | Print options and help and exit | |
| -version |      | Print version information and exit | |
| -v |      | Be very verbose in operation, this means printing every entry as it is ingested as well as step-by-step status updates | |
| -status | Suppress log information and show a real-time status display showing ingest rate, total ingested entries, etc... | |
| -config-file | X | Specify the full path to a configuration file that drives the `migrate` tool. | `-config-file=/tmp/migrate.conf` |


The `migrate` tool is an excellent way to rapidly ingest huge quantities of old data into a new Gravwell system, it is designed to optimize ingest speed as well as query and compression performance.  When you are first standing up Gravwell and preparing to bring in old data, the migrate tool should be your primary tool.
