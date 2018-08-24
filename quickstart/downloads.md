# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.3.tar.bz2) (MD5: b545b8aa2c40c3367c91d71b6042ea6f)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.3 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |9007c88814b0ee2b66f175e11b5e5e30| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.3.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |87453c64fa7f5f99c3f12b866df2c37c| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.3.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |0495a4ccff6615d75fb169284d7d49e7| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.3.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |3eb21c70793632f1be9e42b15d67a910| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.3.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |e3d4f5a6db9bbdec43f4a6eeca3ff2d7| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.3.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |f8b9fb01e5ae4b66410a19bcc1b55d59| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.3.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e47c9b5b2c9ed5744653faa48688d999| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |6d744d21d1c119e896b4e3524f6c8cef| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.2.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |790b9441424c34ac6cdfa2c486650b52| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.3.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |b6c2a9b38c9f342747ed88ac1bdb3d2e| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.3.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |b34ec28a19600e99d64db566837b6ece| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.3.tar.bz2) |
