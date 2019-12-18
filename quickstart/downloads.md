# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.2.sh) (MD5: 35133dd3b564732c59969d384a63b148)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.2 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |2261a9323eb6dbcce735966f0a5a6bcd| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.2.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |d1c494d33dcd765219596248cd192c6a| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.2.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |fb9ccb870b173eed2d114e4d3b6cad5a| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.2.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a93103ef14eb5c94bb32d981778367d9| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.2.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |f6c9bc0b6a446eef4694150e38ddec41| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.2.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |cdad076b07b305fd6e81f77d040c83a9| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.2.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |3a6d6f3fdb8c34c5507bbb307b2139af| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.2.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |1b21e815bb1c35243e7b2c5a036caead| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |b9e535482b30e198766e883e3c0cd117| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.2.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |a7c8c577ca4412fe96646acbd9757526| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.2.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |fe977e1f7878669886439279cd94f562| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.2.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |fe09df3cd5010b8a85ffe1fce1f01071| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.2.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |302d512716376ef9d93507b7b7ba3719| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.2.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |b778e9bc7dc306f5f33fb50afb758194| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.2.sh) |
