# Downloads

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).


### Version 2.2.1 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |25acadf682f1c7857a5fe9907bf53cbf| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.1.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |a4b1d159eb7c96eb6c638085571c2a55| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.1.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |4c33d866f9820d928a8c4895e58838a8| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.1.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |0f7f37a88376c311bd3e8d2864af53f5| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.1.tar.bz2) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |24705c67234f31f2c64ff0b6cd3ddb97| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.1.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |50499b2dd2021061f93c558af563e35c| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.1.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |c157f4fd2e38b642921733c766f3b61e| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.1.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |8263f62d7f1a355750b8330197d0cbaa| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.1.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |aee4dd15cc16be803c3896862c0d3b66| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.1.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |2e83be70e994762c65b4dddb60282a94| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.1.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |e11d23cff8fa970554c4d802cfe98cd7| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.0.tar.bz2) |
