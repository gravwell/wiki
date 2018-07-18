# Downloads

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).


### Version 2.1.1 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |a04d8e954498cb2099d306d5ddc518d6| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.1.1.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |1b8c76557667feb91f9c1f273d93f021| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.1.1.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |16d731bae9bfffb496fe06f263a922a8| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.1.1.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |16d731bae9bfffb496fe06f263a922a8| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.1.1.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |cfa649bc9d449e38f0f9a168e48c41a5| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.1.1.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |91904091f9e05cb59b450e82386375a8| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.1.1.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2ffe88a57ae960a8ce069bcabe830fb9| [Download](https://update.gravwell.io/files/gravwell_win_events_2.1.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |59db3ed3368f0c90fa927b483aefb557| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.1.0.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |f8513b3a9d4452defed4a7d3a7eef871| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.1.1.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |5cee83b39a20a2537fbef59765dd138f| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.1.1.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |5c7a56aedf3dc108e1e7e09a43353ccc| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.1.1.tar.bz2) |
