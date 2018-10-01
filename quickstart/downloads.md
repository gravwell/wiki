# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.8.sh) (MD5: a6af0946ae844e40b3741838af324e45)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.8 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |2aa6738aba6210f6f60ab146fe57bb29| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.8.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |e1e1b721f23823e52a57ceb3e7c03b4c| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.8.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |1efd4e6d11add1e12f80140fc7cdbe67| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.8.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |acda85f2aba67514b15abd18b006496f| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.8.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |7d16e7c315bf9a194efac38e740f9b23| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.8.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |fa633a4e0cd312b1358ee5a707b70bf9| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.8.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |d6cd360660e1d3a863953ac9eaff95c6| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.8.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |d2c3600df5a559cce4b6e14cd19c2e35| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.8.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |0484bd530c8435ea323e5c66413051ea| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.8.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |256aec3390e6dc85f9c7416cc8adbe90| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.8.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |de4a5414b41284fe73061dee893a9bfa| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.8.sh) |
