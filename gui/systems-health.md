# Systems & Health Menu

The Systems & Health sub-menu contains pages which describe the current state of the Gravwell cluster.

![](systems-and-health.png)

## Storage, Indexers, & Wells

This page shows information about the data stored in the indexers of the Gravwell system.

![](storage.png)

The Storage section shows a summary of how much data is in the system, with separate stats for hot and cold storage. The Ingest chart shows the rate at which new data has been entering the system.

At the bottom of the page, the Search Agent section shows information about the search agent component and when it last "checked in".

The Indexer Status section shows how much data is on each indexer and how quickly each is ingesting new data. If you see that one indexer has much less data than the others, you may need to investigate your ingester configs to make sure they are configured to use *all* indexers. Clicking on an indexer in this section, or clicking on it in the left-hand menu, will open a page which displays more detailed information specific to that indexer:

![](ingester-stats.png)

## Ingesters & Federators

This page shows information about ingesters. The ingester list is searchable and sortable. Ingesters which have connected via Federators will appear in this page, as will the Federators themselves; be aware that entry/byte counts for Federators are the sum of counts from all ingesters connected to them.

![](ingesters-page.png)

If an ingester gets disconnected, it will be displayed at the top of the page in the "Missing Ingesters" section:

![](missing-ingesters.png)

Each indexer keeps track of the ingesters it has seen. It stores the most-recently-seen ingester state in `/opt/gravwell/etc/ingester_states.json`. If you decide to "retire" some ingesters and no longer want to see them in the Missing Ingesters section, you can stop the indexer, remove that file, and restart:

```
systemctl stop gravwell_indexer.service
rm /opt/gravwell/etc/ingester_states.json 
systemctl start gravwell_indexer.service
```

```{note}
Removing the ingester_states.json file means that *all* currently-missing ingesters will be forgotten.
```

## Hardware

The Hardware page shows information about the individual computers which make up the Gravwell cluster. At the top of the page is information about cluster-wide CPU and memory usage, ingest rates, etc.; below are individual "cards" for each indexer (be1, be2, be3):

![](hardware.png)

Each card has several different display options, selected via the links in the upper-right corner of each card. "Health" shows uptime, CPU and memory usage, and network/disk read & write stats. "Ingestion" shows the rate at which new entries are being ingested into that particular indexer. "Specifications" shows system specs for the hardware. "Disks" shows information about the storage on the system, but in general that information is more conveniently viewed on the Disks page.

## Disks

The Disks page contains information about disk storage on the cluster. Only disks which contain Gravwell data will be displayed, to avoid clutter.

![](disks.png)


## Topology

The Topology page shows how indexers and ingesters are connected.

![](topology.png)

Note how both indexer1 and indexer2 connect to the same set of wells. This means that the same wells are defined on each indexer. Note also that the "flow" ingester is connected directly to the ingesters, while the others connect via a Federator.
