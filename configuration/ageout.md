# Data Ageout

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


