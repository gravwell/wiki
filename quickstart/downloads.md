# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.3.2.sh) (MD5: 01be9d471eeca3a90909f0b77af7708f)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.3.2 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |2a6a74a83f2570cadf09ac10bc084b4f| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.3.2.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |316c77e4c9a54817f5272c79b58c0d0d| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.3.2.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |e96dd67139f2b4068ac15c94c08eb0f9| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.3.2.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |62cc79b4cfc728b750dc566f199c79a5| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.3.2.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |9079476133bee7a941ea33225efe25d8| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.3.2.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |1b7188cb6e75818a1c96fc493225c62d| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.3.2.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |7e106bd72166b67a6bd3a0bdc133d97f| [Download](https://update.gravwell.io/files/gravwell_win_events_2.3.1.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |8a1536bea2f3d8cc6888c5d535a529ca| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.3.1.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |40e5485efb8a16d87abe08aee1c2e260| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.3.2.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |d3dcb4c0b92d38dc716e5a2d43e549be| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.3.2.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |706f6660ce9b5eec26c66228473857d4| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.3.2.sh) |
