# Advanced Gravwell Configuration

This document describes some more advanced configuration options for Gravwell installations, including information on configuring storage wells, data ageout, and multi-node clusters.

Gravwell is optionally a distributed system, allowing for multiple indexers which comprise a Gravwell cluster.  The default installation will install both the webserver and indexer on the same machine, but with an appropriate license configuration, running a cluster is just as simple as running a single instance.

## Installer Options

The Gravwell installer supports several flags to make automated installation or deployment easier.  The following flags are supported:

| Flag | Description |
|------|-------------|
| `--help` | Display installer help menu |
| `--no-certs` | Installer will not generate self-signed certificates
| `--no-questions` | Assume all defaults and automatically accept EULA
| `--no-random-passwords` | Do not generate random ingest and control secrets
| `--no-indexer` | Do not install Gravwell indexer component
| `--no-webserver` | Do not install the Gravwell webserver component
| `--no-searchagent` | Do not install the Gravwell search agent (typically used with `--no-webserver`)
| `--no-start` | Do not start the components after installation
| `--no-crash-report` | Do not install the automated debug report component
| `--use-config` | Use a specific config file

### Common use-case examples of advanced installation requirements

If you are deploying Gravwell to a cluster with multiple indexers, you would not want to install the webserver component on your indexer nodes.

If you are using an automated deployment tool you don’t want the installer stopping and asking a questions.

If you already have your list of indexers with ingest and control shared secrets, specifying a configuration file at install time can greatly speed up the process.

An example argument list for installing the indexer component without installing the webserver or randomizing passwords is:

```
root@gravserver# bash gravwell_8675309_0.2.sh --no-questions --no-random-passwords --no-webserver
```

If you choose to randomize passwords, you will need to go back through your indexers and webserver and ensure the Control-Auth parameter in the gravwell.conf file matches for the webserver and each indexer.

## General Configuration

Configuration of a Gravwell cluster is designed to be simple and efficient right from the start.  However, there are knobs to twist that can allow the system to better take advantage of extremely large systems or smaller embedded and industrial devices with memory constraints.  The core configuration file is designed to be shared by both the webserver and indexer, and is located by default at `/opt/gravwell/etc/gravwell.conf`

For a detailed listing of configuration options see [this page](parameters.md)

For a complete example indexer configuration see our [example default config](indexer-default-config.md)

The most important items in the configuration file are the `Ingest-Auth`, `Control-Auth`, and `Search-Agent-Auth` configuration parameters.  The `Control-Auth` parameter is the shared secret that the webserver and indexers use to authenticate each other. If an attacker can communicate with your indexers and has the `Control-Auth` token, he has total access to the data they store.  The `Ingest-Auth` token is used to validate ingesters, and restricts the ability to create tags and push data into Gravwell.  Gravwell prides itself on speed, which means an attacker with access to your `Ingest-Auth` token can push a tremendous amount of data into Gravwell in a very short amount of time.  The `Search-Agent-Auth` token allows Gravwell's Search Agent utility to automatically connect to the webserver and issue searches on the behalf of users. These tokens are important and you should protect them carefully.

Attention: In clustered Gravwell installations, it is essential that all nodes are configured with the same `Ingest-Auth` and `Control-Auth` values to enable proper intercommunication.

## Webserver Configuration

The webserver acts as the focusing point for all searches, and provides an interactive interface into Gravwell.  While the webserver does not require significant storage, it can benefit from small pools of very fast storage so that even when a search hands back large amounts of data, users can fluidly navigate their results.  The webserver also participates in the search pipeline and often performs some of the filtering, metadata extraction, and rendering of data.  When speccing a webserver, we recommend a reasonably sized solid state disk (NVME if possible), a memory pool of 16GB of RAM or more, and at least 4 physical cores.  Gravwell is built to be extremely concurrent, so more CPU cores and additional memory will only increase its performance.  An Intel E5 or AMD Epic chip with 32GB of memory or more is a good choice, and more is always better.

Two configuration options inform the webserver which indexers it should use for searching. The `Remote-Indexers` option specifies the IPs of the indexers, and the `Control-Auth` option gives a shared key used by the webserver to authenticate to the indexers. A webserver connecting to three indexers might contain the following in its `gravwell.conf`:

```
Control-Auth=MySuperSecureControlToken
Remote-Indexers=net:10.0.1.1:9404
Remote-Indexers=net:10.0.1.2:9404
Remote-Indexers=net:10.0.1.3:9404
```

Note: The indexers listed above are listening for control connections on port 9404, the default. This port is set by the `Control-Port` option in the indexer's gravwell.conf file.

### Webserver Configuration Pitfalls

* Missing or misconfigured Remote-Indexers
* Missing or mismatched Control-Auth tokens
* Mismatched licenses on webserver and backend
  * Both the webserver and indexer must have compatible licenses
* Poor network connectivity between the webserver and indexers
  * High latency, low bandwidth, or misconfigured MTU sizes.
* Firewalls blocking access to indexer or webserver ports
  * The default is 9404

## Indexer Configuration

Indexers are the storage centers of Gravwell and are responsible for storing, retrieving, and processing data.  Indexers perform the first heavy lifting when executing a query, first finding the data then pushing it into the search pipeline.  The search pipeline will distribute as much of a query as is possible to ensure that the indexers can do as much work in parallel as possible.  Indexers benefit from high speed low latency storage and as much RAM as possible.  Gravwell can take advantage of file system caches, which means that as you are running multiple queries over the same data it won’t even have to go to the disks.  We have seen Gravwell operate at over 5GB/s per node on well-cached data.  The more memory, the more data can be cached.  When searching over large pools that exceed the memory capacity of even the largest machines, high speed RAID arrays can help increase throughput.

We recommend indexers have at least 32GB of memory with 8 CPU cores.  If possible, Gravwell also recommends a very high speed NVME solid state disk that can act as a hot well, holding just a few days of of the most recent data and aging out to the slower spinning disk pools.  The hot well enables very fast access to the most recent data, while enabling Gravwell to organize and consolidate older data so that he can be searched as efficiently as possible.

There are a few key configuration options in an indexer's gravwell.conf which affect its general behavior:

* `Control-Port` sets the port on which the indexer will listen for incoming connections from a webserver. Default 9404.
* `Control-Auth` sets a shared secret which webservers use to authenticate. Defaults to a randomly-generated string.
* `Ingest-Port` and `TLS-Ingest-Port` specify which ports to listen on for unencrypted and encrypted ingest traffic, respectively.
* `Ingest-Auth` sets a shared secret used by data ingesters to authenticate to the indexer. Defaults to a randomly-generated string.

Indexers store their data in _wells_. Each well stores some number of tags. If a well contains 100GB of data tagged "pcap" and 10MB of data tagged "syslog", searching for syslog data means the indexer also has to read the pcap data from the disk, slowing down the search. For this reason we strongly suggest creating separate wells for tags you anticipate will contain a lot of data. See the 'Tags and Wells' section for more information.

## Tags and Wells

**Tags** are used as a method to logically separate data of different types.  Tags are applied at ingest time by the ingesters (SimpleRelay, NetworkCapture, etc).  For example, it is useful to apply unique tags to syslog logs, Apache logs, network packets, video streams, audio streams, etc.  **Wells** are the storage groupings which actually organize and store the ingested data. Although users typically do not interact with them, the wells store data on-disk in **shards**, with each shard containing approximately 1.5 days of data.

Tags can be assigned to wells so that data streams can be routed to faster or larger storage pools. For example, a raw pcap stream from a high bandwidth link may need to be assigned to a faster storage pool while relatively low-volume log entries from syslog or a webserver do not require fast storage. A tag-to-well mapping is a one-to-one mapping; a single tag cannot be assigned to multiple wells, although a well can contain multiple tags.  Logically and physically separating data streams allows different rules to be applied to different data.  For example, it may be desirable to expire or compress high bandwidth streams, like network traffic, every 15 days while keeping low bandwidth streams for much longer.  The logical separation also greatly increases search performance as the system intelligently queries the appropriate well based on tag (e.g. when searching syslog entries located in the well named default, Gravwell will not engage any other wells).

Tag-to-well mappings are defined in the `/opt/gravwell/etc/gravwell.conf` configuration file. By default, only a `Default-Well` will be configured, which accepts all tags. An example configuration snippet for an indexer with multiple wells associated tags might look like this:

```
[Default-Well]
	Location=/opt/gravwell/storage/default/

[storage-well "raw"]
	Location=/opt/gravwell/storage/raw/
	tags=pcap
	tags=video
```

The well named "raw" is thus used to store data tagged "pcap" and "video", which we could reasonably assume will consume a significant amount of storage.

### Tag Restrictions and Gotchas

Tag names can only contain alpha numeric values; dashes, underscores, special characters, etc are not allowed in tag names.  Tags should be simple names like "syslog" or "apache" that are easy to type and reflect the type of data in use.

The Default well receives all entries with tags that have not been explicitely assigned to other wells.  For example, if you have one well named Syslog which has been assigned the tags "syslog" and "apache" then all other tags will go to the Default well.  Ingesters can still produce entries with tag names that are not explicitely defined in the gravwell.conf file; the entries will just be co-mingled with all other unassigned tags in the default well.

When reassigning tags between wells, the system will NOT move the data.  If you ingest data under the tag "syslog" without pinning the tag to a non-default well, then change the config file to define a new well or assign the syslog tag to an existing well, all data that exists in the default well under the syslog tag is no longer searchable.  Contact support@gravwell.io for access to a standalone tool for well and tag migration that can recover the entries, or for help reingesting old wells into an optimized/alternate configuration.

### Well Replication

A Gravwell cluster with multiple indexer nodes can be configured so that nodes replicate their data to one another in case of disk failure or accidental deletion. See the [replication documentation](replication.md) for information on configuring replication.

## Data Ageout

Gravwell supports an ageout system whereby data management policies can be applied to individual wells.  The ageout policies control data retention, storage well utilization, and compression.  Each well supports a hot and cold storage location with a set of parameters which determine how data is moved from one storage system to the other.  An ideal Gravwell storage architecture is comprised of relatively small pools of high-speed storage that is tolerant to random accesses and a high volume/low cost storage pool to be used for longer term storage.  NVME-based flash and/or XPoint drives make a great hot well while magnetic RAID arrays, NAS, or SAN pools work well for cold pools.  Searching is not impeded during ageout, nor is ingestion.  

Ageout policies can be defined via three parameters:

* Time
* Total Storage
* Storage Availability

The time component of storage allows for specifying data retention policies to adhere to policies, contractual agreements, or legal requirements. For instance, corporate policy may state that web proxy logs should be stored for no more than 30 days.

 The total storage parameter specifies an upper storage bound for a well, instructing Gravwell to only ageout or discard data when the amount of stored data exceeds the storage bounds. 

The storage availability parameter instructs Gravwell to maintain at least a certain minimum amount of available storage, aging out data as needed to free up space.  Storage availability constraints are useful for when you want to allow Gravwell to use free storage on a device, but discard data if the device ever drops below some availability threshhold.

Multiple constraints can be added to a single well, allowing for an intersection of rules. For instance, you might specify that data should be kept for up to 90 days, unless the total exceeds 100GB.

The ageout system is designed to optimize data storage as it transfers entries from hot to cold pools.  The optimization localizes entries of the same timerange and tag, reducing head movement on traditional spinning disks.  Combined with compression, the optimization phase can significantly improve storage utilization and search performance on cold data.

The ageout system can be configured to delete old data, so it is critically important that the configuration be vetted prior to engaging well ageout.

To enable data ageout each well must be provided a hot storage location and a cold storage location.  The hot location is specified via the ["Location" directive](parameters.md) in the gravwell.conf file.  The cold storage location is specified via a "Cold-Location" directive which specifies an absolute path to a directory to be used for cold storage.  Cold storage locations may not overlap with any other well storage location, hot or cold.  In addition to a cold storage location, at least one ageout constraint must be specified to direct when and how data is moved from the hot pool to the cold pool.

Attention: Ageout configurations are on a per well basis.  Each well operates independently and asynchronously from all others.  If two wells are sharing the same volume, enabling ageout directives based on storage reserve can cause one well to agressively migrate and/or delete data due to the incursion by another.

Note: If data is actively coming into a storage shard that is marked for ageout or is actively being queried, the ageout system will defer aging out the shard to a later time.

### Time-Based Ageout

Time-based ageout manages data based on time retention requirements.  For example, an organization may have requirements that all logs be kept for 90 days.  The time-based ageout constraint is best used on cold data pools where policy and/or legal requirements dictate log retention times.  Time based ageout durations can be specified in days and weeks using the following case-insensitive abbreviations:

* "d"     - days
* "days"  - days
* "w"     - weeks
* "weeks" - weeks

An example well configuration using only a hot pool of data and deleting data that is more than 30 days old:

```
[Storage-Well "syslog"]
	Location=/mnt/xpoint/gravwell/syslog
	Tags=pcap
	Hot-Duration=30D
	Delete-Cold-Data=true
```

An example configuration in which data is moved from the hot pool to the cold pool after 7 days, and deleted from the cold pool after 90 days:

```
[Storage-Well "syslog"]
	Location=/mnt/xpoint/gravwell/syslog
	Cold-Location=/mnt/storage/gravwell_cold/syslog
	Tags=pcap
	Hot-Duration=7D
	Cold-Duration=90D
	Delete-Frozen-Data=true
```

Note: In the above configuration, data will be deleted permanently when it is 97 days old, having spent 7 days in the hot pool and 90 days in the cold pool.

The Time based ageout is invoked once per day, sweeping each pool for shards that can be aged out.  By default the sweep happens at midnight UTC, but the execution time can be overridden in the well configuration with the Ageout-Time-Override directive.  The override directive is specified in 24 hour UTC time.

An example configuration that overrides the ageout time checks to occur at 7PM UTC:
```
[Storage-Well "syslog"]
	Location=/mnt/xpoint/gravwell/syslog
	Cold-Location=/mnt/storage/gravwell_cold/syslog
	Tags=syslog
	Tags=switchlogs
	Hot-Duration=7D
	Cold-Duration=90D
	Delete-Frozen-Data=true
	Ageout-Time-Override=19:00
```

### Total Storage-Based Ageout

Total storage constraints allots a specific amount of storage in a volume regardless of time spans.  Storage constraints allow a Gravwell indexer to make agressive and full use of high speed storage pools which may be of limited size (such as NVME flash).  The indexer will keep entries in the storage pool, as long as the well isn't consuming more than allowed.  Storage constraints allow for unexpected bursts of ingest without disrupting data storage.  For example, if an indexer has 1TB of high speed flash storage which typically handles 7 days of hot storage but an unexpected data event causes 600GB of ingest in a single day, the indexer can age out the older data to the cold pool without disrupting the hot pool's ability to take in new data.  Shards are prioritized by time; the oldest shards are aged out first for both hot and cold pools.

An example well configuration that keeps up to 500GB of data in a hot pool, deleting old data when the 500GB limit is exceeded:

```
[Storage-Well "windows"]
	Location=/mnt/xpoint/gravwell/windows
	Tags=windows
	Tags=sysmon
	Max-Hot-Storage-GB=500
	Delete-Cold-Data=true
```

An example well configuration where the hot pool keeps approximately 50GB and then ages out into a cold pool which keeps up to 10TB:

```
[Storage-Well "windows"]
	Location=/mnt/xpoint/gravwell/windows
	Cold-Location=/mnt/storage/gravwell_cold/windows
	Tags=windows
	Tags=sysmon
	Max-Hot-Storage-GB=50
	Max-Cold-Storage-GB=10000
	Delete-Frozen-Data=true
```

Attention: Storage-based constraints are not an instant hard limit.  Be sure to leave a little room so that when a storage allotment is exceeded, the indexer can ageout data while still ingesting.  For example, if a hot storage device can hold 512GB and the system typically ingests 100GB per day, setting the storage limit to 490GB should provide enough headroom so that the hotpool won't completely fill up while the indexer is migrating data.

### Storage Available-Based Ageout

Well storage constraints can also be applied based on availability of storage.  Some wells may be low priority, consuming storage only when it is available.  Using the storage reserve directives allows you to specify a well which is free to consume as much space as it wants, so long as some ceiling on available storage is maintained on the volume.  A storage reserve paradigm essentially allows the well to act as a "second class" well, consuming only when no one else is.  Specifying a "Hot-Storage-Reserve=5" ensures that should the hot storage volume drop below 5% free space the well will begin migrating or deleting its oldest shards.  The reserve directives apply to the underlying volume hosting the storage location, meaning that if the volume is also hosting other wells or other arbitrary file storage, the well can pull back its storage usage as needed.

An example well configuration which will use the hot location as long as there is 30% free space on the volume, and will use the cold volume as long as there is 10% free space:

```
[Storage-Well "doorlogs"]
	Location=/mnt/xpoint/gravwell/doorlogs
	Cold-Location=/mnt/storage/gravwell_cold/doorlogs
	Tags=badgeaccess
	Hot-Storage-Reserve=30
	Cold-Storage-Reserve=10
	Delete-Frozen-Data=true
```

Attention: The Gravwell ageout system which operates on storage reserves is operating entirely orthogonal to outside influences, if a well is configured to respect a 50% storage ceiling and an outside application fills the volume to 60%, Gravwell will delete all entries outside the active shard.  Wells configured with storage reserved should be treated as expendable.

### Caveats and Important Notes

Ageout contraints are applied to entire shards, if a single shard grows beyond the size of a data constraint the shard will age out in its entirety once the shard is idle.

Time-based constraints require that the entire shard fall outside the specified time window.  As such, time contraits that are less than 1 day have no meaning, and hot pools must be able to hold at least 2 days worth of data.

Take care when combining time-based constraints with total storage constraints. If `Hot-Duration=7D` and `Cold-Duration=90D` are specified, data will be deleted after 97 days. However, if `Max-Hot-Storage-GB=2` and `Cold-Duration=90D` are specified, data will move from the hot well to the cold well when the hot well exceeds 2GB, and **data will be deleted from the cold well when it is 90 days old**.

#### Deletion Safety Nets

Ageout policies can and will **delete data**.  It is critically important that administrators vet configurations and hardware resources to ensure that both hot and cold storage pools can accommodate the requested ageout constraints.  Well configurations require that deletion be explicitely authorized via a "Delete-X-Data" directive (`Delete-Cold-Data` for hot pools and `Delete-Frozen-Data` for cold pools).

The Delete-X-Data directives are used as safety nets to ensure administrators think about data deletion during configuration.  For example, if a well is configured to have a hot pool, no cold pool, and a hot retention of 7 days, Gravwell will attempt to delete data from the hot pool that is older than 7 days.  However, if the well does not have a configuration directive of "Delete-Cold-Data=true" the ageout system WILL NOT delete the data; instead it will complain in the logs and essentially disable ageout.

Cold pools contain the same directive requirement, meaning that Gravwell will not delete data from the cold pool as it ages out of the cold pool unless the "Delete-Frozen-Data=true" directive is set.  Note that if data that is 120 days old is ingested into the hot pool and the cold pool retention configuration is 90 days, **and the Delete-Cold-Data directive is set to true** the ageout system will bypass the cold pool entirely and directly delete the data.

Time-based ageout relies on the system clock to determine when data should be aged out, meaning that an improper system clock could cause data to be aged out prematurely.  Administrators must ensure that system clocks are well maintained and if NTP is in use, it is a trusted source.  If an attacker can gain control over the system clock of a Gravwell indexer for reasonable periods of time, they can cause ageouts.  Time is the anchor in Gravwell and it must be protected with the same vigor as the storage pools.

### Compression

The cold storage locations are compressed by default, which helps shift some load from storage to CPU.  Gravwell is a highly asynchronous system built for modern hardware in a scale wide paradigm.  Modern systems are typically overprovisioned on CPU with mass storage lagging behind.  By employing compression in the cold pool Gravwell can reduce stress on storage links while employing excess CPU cycles to asynchronously decompress data.  The result is that searches can be faster when employing compression on slower mass storage devices.

A notable exception is data that will not compress much (if at all). In this situation, attempting to compress the data burns CPU time with no actual improvement in storage space or speed. Raw network traffic is a good example where encryption and high entropy prevent effective compression.  To disable compression in a cold storage pool, add the "Disable-Compression=true" directive.

An example storage well with compression disabled and the hot pool constrained by total storage:

```
[Storage-Well "network"]
	Location=/mnt/xpoint/gravwell/network
	Cold-Location=/mnt/storage/gravwell_cold/network
	Tags=pcap
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Disable-Compression=true
```

### Ageout Rule Interactions

Ageout rules can be stacked and combined to provide robust control over storage resources.  For example, Gravwell highly reccomends a small pool of high speed storage for use in the hot pool and a large low cost array of disks for the cold pool.  Storage constraints can be combined to utilize the small high speed pool at will, while still adhering to data retention policies.

Data ageout constraints operate independently of one another in a first come first serve basis.  If a time constraint of 7 days is applied to a well in addition to a storage constraint of 100GB, shards will be aged out at 7 days or when the pool exceeds 100GB, whichever comes first.  While useful to combine multiple constraints for a single pool in a well, be aware that an overly agressive ageout configuration can cause unexpected behavior.  For example if a well has a 7 day retention, but also a Hot-Storage-Reserve of 5% Gravwell will attempt to meet the 5% retention independently of the 7 day retention.  Each ageout directive acts entirely independently, setting a retention size or storage reserve can and will override any time based directives, and vice versa.

Example well configuration which uses up to 100GB of hot storage and retains data for 90 days:

```
[Storage-Well "weblogs"]
	Location=/mnt/xpoint/gravwell/web
	Cold-Location=/mnt/storage/gravwell_cold/web
	Tags=apache
	Tags=nginx
	Tags=iis
	Max-Hot-Data-GB=100
	Cold-Duration=90D
	Delete-Frozen-Data=true
```

Example well configuration which tries to keep 10% free on the hot storage volume while retaining data for 30 days:

```
[Storage-Well "weblogs"]
	Location=/mnt/xpoint/gravwell/web
	Cold-Location=/mnt/storage/gravwell_cold/web
	Tags=apache
	Tags=nginx
	Tags=iis
	Hot-Storage-Reserve=10
	Cold-Duration=30D
	Delete-Frozen-Data=true
```

Example well configuration which keeps 7 days or 100GB in the hot pool and keeps all data in the cold pool as long as 20% is available:

```
[Storage-Well "weblogs"]
	Location=/mnt/xpoint/gravwell/web
	Cold-Location=/mnt/storage/gravwell_cold/web
	Tags=apache
	Tags=nginx
	Tags=iis
	Hot-Storage-Reserve=100
	Hot-Duration=7D
	Cold-Storage-Reserve=20
	Delete-Frozen-Data=true
```

<<<<<<< Updated upstream
=======
### Well Replication

A Gravwell cluster with multiple indexer nodes can be configured so that nodes replicate their data to one another in case of disk failure or accidental deletion. See the [replication documentation](replication.md) for information on configuring replication.
>>>>>>> Stashed changes
