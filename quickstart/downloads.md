# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.4.tar.bz2) (MD5: 81c8e3d71ae473592503b20bb586ad87)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.4 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |2ec02eafe4f8161a9c867396383800eb| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.4.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |fbecbd54e0b2bd0c876b58db659d199c| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.4.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |d67fc4e511c007e8ab9e81c5af650191| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.4.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |e67abf49c9fe4bc99e8effc59500c917| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.4.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |9d081b308f72d2d0dc07c3016b85454e| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.4.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |e81900bd536fe1b6486662cd2be7b92a| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.4.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e47c9b5b2c9ed5744653faa48688d999| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |6d744d21d1c119e896b4e3524f6c8cef| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.2.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |d747472bedfdc5e44d4d5c92567bf44f| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.4.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |96f8ac4d7747d7943ba32eb87218b0fc| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.4.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6a9bd3c62e7397c63436acca65f0a5a0| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.4.tar.bz2) |
