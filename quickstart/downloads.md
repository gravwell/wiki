# Downloads

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).


### Version 2.1.0 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |137b039902d16c9253bf6fc336514025| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.1.0.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |2c958b597e6e0ae09e6d77caa3e8d2ae| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.1.0.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |eb10ff15722de33c577f0d81d0ca36f5| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.1.0.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |eb10ff15722de33c577f0d81d0ca36f5| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.1.0.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |a8c4fa6913ee4463c2225a6a1ffeff55| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.1.0.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |a355c02d00e26df25a98b622986faa96| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.1.0.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |3955f499dd098c9b27282eb464917d02| [Download](https://update.gravwell.io/files/gravwell_win_events_2.0.9.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |72a3348ad559f34ff00d5f3966c59cf1| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.0.9.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |eb909bef170fd74403bb5a0c7bf66eb0| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.1.0.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |9d630cc96fa325910b420ea10b19dc6e| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.1.0.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |77c28ec9ec2a174b32076d4cf3ef294c| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.1.0.tar.bz2) |
| [Search Agent](#!scripting/scheduledsearch.md) | The search agent runs user-defined searches and scripts on a specified schedule. |81b30432481da44d30be2cd171542370| [Download](https://update.gravwell.io/files/gravwell_searchagent_installer_2.0.8.tar.bz2) |
