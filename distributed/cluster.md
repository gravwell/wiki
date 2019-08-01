# Gravwell Clusters

While Gravwell will happily run entirely contained on a single node, environments with large quantities of data may be best served with a cluster. In a Gravwell cluster, one or more webservers control one or more indexers, all on separate servers.

The Gravwell architecture is described in detail in [this document](#!architecture/architecture.md); this article deals more concretely with details of configuration for clustering.

The following are all valid Gravwell configurations:

* One webserver and one indexer, both running on the same system. This is the default installation; we do not consider it a cluster because the webserver and indexer are on the same machine.
* One webserver and one indexer, each on its own server. This is the simplest possible cluster.
* One webserver and multiple indexers, each on their own servers.
* [Multiple webservers](#!distributed/frontend.md) and multiple indexers.

We will discuss configuring a cluster with a single webserver and one or more indexers. If you want multiple webservers, we recommend following the steps outlined here to first configure a single-webserver cluster, then using the information in [this document](#!distributed/frontend.md) to add additional webservers later.

## Planning Your Cluster

There are several things to keep in mind while planning your cluster:

* Your Gravwell license -- how many nodes does it allow?
* Type and number of servers available
* Network interconnect
* Quantity of data you wish to ingest

The last point is tricky, because the number of indexers required to store a given amount of data depends heavily on the hardware available (how much RAM per node, NVME vs SSD vs spinning disk, etc.) and how much querying you intend to do. We recommend contacting [Gravwell support](mailto:support@gravwell.io) for help in planning your cluster.

Before beginning, it is useful to know the specific IP addresses or hostnames which will be used for the webserver and the indexers. Also, review the ports in [this document](#!configuration/networking.md) to ensure that your network is configured to allow necessary connections between Gravwell components; in brief, you'll want to make sure the webserver can reach port 9404 on the indexers, that users can reach port 80 and 443 on the webserver, and that your ingesters will be able to reach ports 4023 and 4024 on the indexers.

**Do I need multiple webservers?** Most likely, no. Multiple webservers make your cluster more complex. In general, we recommend setting up a single webserver, then adding more if the load is too high in use. 

### Planning Tags and Wells

If you know you will be ingesting netflow, packet capture, syslogs, Apache logs, and Windows logs, you should consider putting each data class into its own storage well. By putting syslog into a separate well from high-volume sources such as raw packets, you'll improve the speed of your syslog queries. See [the advanced configuration article](#!configuration/configuration.md#Tags_and_Wells) for a more detailed discussion of tags and wells.

In the planning stage, it is useful to make a list of each data source, which tag you'll apply to it, and which well you intend to use for storing it. If you're not yet sure what data you'll be ingesting, you can use the default well configuration and make adjustments once real data begins flowing in.

Also consider [ageout](#!configuration/ageout.md) needs at this time. How long do you want to retain system logs? Packet captures? Netflow records? For example, if you're ingesting 100 gigabytes of packets per day and your cluster has 10 TB of storage in total, you could store about 100 days of uncompressed packet data, but any data older than that will need to be aged out.

## Installing the Webserver

We recommend starting your installation with the webserver. This will generate a `gravwell.conf` you can use as a starting point for configuring the indexers. Follow the instructions below to install the webserver using one of the two available methods.

### Installing webserver via shell installer

To install only the webserver and search agent using the [standalone shell installer](#!quickstart/downloads.md), run the following command (the version will likely differ for you):

```
root@headnode# bash gravwell_3.2.0.sh --no-indexer
```

### Installing webserver via Debian package

Setting up a webserver-only node from the Debian package requires a few additional steps. Set up the repository and install the Gravwell package as described in [the quickstart document](#!quickstart/quickstart.md#Debian_repository). You can allow the installer to auto-generate secret tokens for you when prompted.

After installing the `gravwell` package, you'll need to disable the indexer:

```
root@headnode# systemctl stop gravwell_indexer.service
root@headnode# systemctl disable gravwell_indexer.service
```

## Building the Config File

Now that the webserver is installed, we will build the configuration file to be used by the indexers. Copy `/opt/gravwell/etc/gravwell.conf` from the webserver to a local directory, then open it in an editor. By using the `gravwell.conf` file from the webserver, we will ensure that all indexers share the exact same set of authentication tokens.

Delete any lines which begin with `Indexer-UUID`. If you installed the webserver using the shell installer, there may not be any Indexer-UUID line; this is fine.

Next, define your storage wells. At the time of writing, the default configuration creates separate wells for Linux syslogs, Windows event logs, web server logs (Apache, nginx, etc.), netflow/ipfix records, and "raw" data such as packet capture. You should feel free to delete or modify any of the `Storage-Well` definitions to match your planned tag/well configuration, but please note that the configuration **must** contain a `Default-Well` definition, so don't delete it! Now is a good time to define [data retention and ageout policies](#!configuration/ageout.md) for each well, per the guidelines you chose during the planning stage.

If you intend to use [data replication](#!configuration/replication.md), in which indexers replicate each other's stored entries to prevent data loss in case of hardware failure, this is also the time to add it to the configuration. See the [replication article](#!configuration/replication.md) for instructions on how to configure replication.

## Installing the Indexers

Having written the indexer config file, you should now copy it to the servers you will use as indexers. We will assume you have put the file in `/tmp/indexer.conf` on every indexer server.

### Installing indexer via shell installer

To install only the indexer using the [standalone shell installer](#!quickstart/downloads.md), run the following command on every indexer server:

```
root@indexer0# bash gravwell_3.2.0.sh  --no-webserver --no-searchagent --use-config /tmp/indexer.conf
```

This will copy your config file to its proper location in /opt/gravwell/etc/gravwell.conf, so you can delete /tmp/indexer.conf after installation if desired.

### Installing indexer via Debian package

Setting up an indexer-only node from the Debian package requires a few additional steps. Follow these steps on every indexer server.

Set up the repository and install the Gravwell package as described in [the quickstart document](#!quickstart/quickstart.md#Debian_repository). You can allow the installer to auto-generate secret tokens for you when prompted, because we will be replacing the config file with our own.

After installing the `gravwell` package, we want to stop all Gravwell services:

```
root@indexer0# systemctl stop gravwell_indexer.service
root@indexer0# systemctl stop gravwell_webserver.service
root@indexer0# systemctl stop gravwell_searchagent.service
```

Then install the customized config file:

```
root@indexer0# cp /tmp/indexer.conf /opt/gravwell/etc/gravwell.conf
```

Disable the webserver and search agent components:

```
root@indexer0# systemctl disable gravwell_webserver.service
root@indexer0# systemctl disable gravwell_searchagent.service
```

And finally restart the indexer service:

```
root@indexer0# systemctl start gravwell_indexer.service
```

## Final Webserver Configuration

At this point, you should have the webserver running on one node and indexer processes running on several other nodes. The final step is to inform the webserver of the newly-configured indexers.

Connect to the webserver and edit `/opt/gravwell/etc/gravwell.conf`. Find the line beginning with `Remote-Indexers`; it will probably look like `Remote-Indexers=net:127.0.0.1:9404`. Delete that, then add one Remote-Indexers line *per indexer*; for example, if your indexers are at 10.0.1.1 through 10.0.1.5, your webserver's gravwell.conf should contain the following:

```
Remote-Indexers=net:10.0.1.1:9404
Remote-Indexers=net:10.0.1.2:9404
Remote-Indexers=net:10.0.1.3:9404
Remote-Indexers=net:10.0.1.4:9404
Remote-Indexers=net:10.0.1.5:9404
```

Note: You can use hostnames instead of IP addresses if you wish.

Once the webserver's gravwell.conf is updated, restart the webserver process:

```
root@headnode# systemctl restart gravwell_webserver.service
```

You can now point your web browser at the webserver and upload a license file when prompted. The webserver will automatically distribute the license file to the indexers.

## Administration

For the most part, administration of a Gravwell cluster is the same as administration of a single-node Gravwell instance. If an indexer is down, a high-priority notification will be shown in the Gravwell UI. If any indexer's disk becomes overly full, it will send a notification message.

## Caveats

Be aware that if an indexer is down, you can still run searches, but the results of the search will lack any entries which are stored on that indexer. If your ingesters distribute entries equally across 10 indexers, but one indexer goes down, your searches will only contain 90% of the results which can lead to inaccurate statistics.
