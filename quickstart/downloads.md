# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.9.sh) (MD5: c774d32bb88ec142ed06db3aa5460713)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.7 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |d2c948305e6f85d28ec9dbcfbdc4616d| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.9.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |ce3e643fe49627c62e3de12f377c6c04| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.9.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |054f8865528083f0558d5ea2787a54fe| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.9.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |04f6a83683f234e35cda8a5b632e558e| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.9.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |47de0d0b08a31058c633aba1339ca9e9| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.9.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |b8825a9baa3d762294a10a256038285a| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.9.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |cbcdd3c3781ad7396854a0482c19f16c| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.9.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |1d91e68533d3b5b564144dec96ff2ab9| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.9.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |8ec98954f5c09ed926850d056733f49e| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.9.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |9fc99835f770af59c50d6ca1dcde1eb1| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.9.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6422a28fdaf9c3b5514bd7766c219957| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.9.sh) |
