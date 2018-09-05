# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.4.tar.bz2) (MD5: 81c8e3d71ae473592503b20bb586ad87)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.4 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |b9a1f347ae04de4cefd7135f8245c328| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.4.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |6aa181e20366cf8a63fdf896f205fb25| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.4.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |5b2769dc2649f8f5dcdb6ae535360f0b| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.4.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |cec64dbdb48ff3d423884eeb70a1dd00| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.4.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |b53355c85d601169774f394d0ee27769| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.4.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |34d24dbee9dd1442eacf16f2ff4f2b0e| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.4.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |f987f35c129b52c50be660a4a967dfa3| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |6c2c15c7c648a337bcddc2cd5702a7bb| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.4.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |79f1c048058097dad0a5a606890734aa| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.4.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |0d97e1b06225ddc19451a806e4304c7b| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.4.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6a9bd3c62e7397c63436acca65f0a5a0| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.4.tar.bz2) |
