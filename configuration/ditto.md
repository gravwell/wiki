# Ditto Indexer Mirroring

Ditto is Gravwell's system for mirroring data from one cluster to another. Unlike replication, it copies data on the source system directly to the destination (or "target") system's *live* storage. This is useful when migrating to an entirely new cluster, or as a means to duplicate some data from one cluster to another.

Data is duplicated at the well level. A given well can be duplicated to one or more destinations. Entries are read from the source well and shipped to the destination, which files them into the appropriate well or wells depending on its own well configuration.

The following terminology will be used in this document:

* **Source cluster**: The indexer or indexers from which data will be cloned.
* **Target cluster**: The indexer or indexers to which data will be cloned.

In order to configure Ditto, you must first define one or more targets, then configure the desired wells to clone their data to the targets.

```{warning}
Entries from the source system will be incorporated directly into the destination/target system's live storage. Once the entries arrive on the destination, they will be indistinguishable from entries ingested to that system in the usual fashion, and it is essentially not possible to "undo" a ditto cloning without deleting the entire well on the destination side.
```

## Target Configuration

Ditto targets are defined in `/opt/gravwell/etc/gravwell.conf` (or a file in `/opt/gravwell/etc/gravwell.conf.d`). A Ditto target uses the same basic configuration block as an [ingester](ingesters_global_configuration_parameters), specifying indexer targets and ingest secrets. There are a few Ditto-specific options, too:

* **Start-Time**: If set to a timestamp (we recommend Unix epoch timestamps or RFC3339 format), this Ditto target will only be sent data from after that timestamp. Specifically, we will find the shard containing that timestamp and duplicate that shard and all following shards.
* **Unresolvable-Tag-Destination**: In some rare cases, the Ditto system may find entries in a shard whose tags do not correspond to any known tag (this can happen if you manually edited `tags.dat`, which is highly discouraged!). By default, these entries are dropped, but if `Unresolvable-Tag-Destination` is set, they will instead be re-tagged with the specified tag.

Here's an example of a simple target definition:

```
[Ditto-Target "new-cluster"]
	Encrypted-Backend-Target=newidx1.example.org
	Encrypted-Backend-Target=newidx2.example.org
	Ingest-Secret=xyzzy
	Start-Time="2024-01-01T00:00:00"
```
```{note}
If you are cloning data across a public network or any metered network connection we highly suggest enabling transport compression by setting `Enable-Compression=true` inside the `Ditto-Target` configuration block.
```

## Well Configuration

To enable Ditto duplication for a given well, add the `Ditto-Target` parameter to the well's config block, e.g.:

```
[Default-Well]
        Location=/opt/gravwell/storage/default/
        Cold-Location=/opt/gravwell/cold_storage/default/
        Hot-Duration=7d
        Ditto-Target="new-cluster"
```

## Worker Configuration

By default, Ditto will only work on one well at a time. If you wish to duplicate multiple wells in parallel, set the `Ditto-Max-Workers` parameter in the `[Global]` section of your `gravwell.conf`. For example, to duplicate up to 4 wells at a time:

```
[Global]
Ditto-Max-Workers=4
```

## Ditto Stats

The Ditto subsystem will periodically emit stats messages into the `gravwell` tag. You can find these stats by running the following query:

```
tag=gravwell syslog Message=="ditto client stats"
```

Each message contains statistics about data transferred for a particular well to a particular Ditto target cluster. The following fields are populated:

* `well`: The well to which the stats apply.
* `entries`: The number of entries transferred for this well since the last stats update.
* `bytes`: The number of bytes transferred for this well since the last stats update.
* `duration`: The elapsed time since the last stats update.
* `Bps`: The approximate transfer rate, in *bytes* per second, over the duration.
* `target-name`: The target cluster which received the data.

You can monitor your transfer rates with a query like this:

```
tag=gravwell syslog Message=="ditto client stats" Bps!="0" well "target-name" as target 
| stats mean(Bps) by well target 
| chart mean by well target
```
