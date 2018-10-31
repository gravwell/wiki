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

## Data Ageout

Gravwell supports an ageout system whereby data management policies can be applied to individual wells.  The ageout policies control data retention, storage well utilization, and compression.  For more information about configuration data ageout see the [Data Ageout](ageout.md). section

## Well Replication

A Gravwell cluster with multiple indexer nodes can be configured so that nodes replicate their data to one another in case of disk failure or accidental deletion. See the [replication documentation](replication.md) for information on configuring replication.

## Query Acceleration

Gravwell supports the notion of "accelerators" for individual wells, which allow you apply parsers to data at ingest to generate optimization blocks.  Accelerators are just as felexible as query modules and are transparently engaged when performing queries.  Accelerators are extremely useful for needle-in-haystack style queries, where you need to zero in on data that has specific field values very quickly.  See the [Accelerators](accelerators.md) section for more information and configuration techniques.
