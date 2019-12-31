# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.3.sh) (MD5: 35133dd3b564732c59969d384a63b148)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.3 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |7ddaae4ffb68eff102393b26e97e002a| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.3.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |e854f8973e0977d3808863218e36032a| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.3.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |8c6a3c54ec2aab064ef3848140540ffb| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.3.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |1190dd4d14e7c98ce8e9e70de7a93ecb| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.3.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |e2c1d37e38152f72f38999dbf7a82f7c| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.3.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |5f45427b42f761ae65780a8ea10f3a1f| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.3.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |4a366e118e26f6f65bb164e7b18c8bb3| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.3.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |1b21e815bb1c35243e7b2c5a036caead| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |192b340158aa4d1d83549329796f30ae| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.3.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |301aabf71b1f11b6dda84db108557560| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.3.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |f0214e5927df9a3a86aa981d2acf58f3| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.3.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |843e68ad0f4e6edc971ec739ae11e41f| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.3.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |fefbd9c9f4f80ac9ea9c6665ded57e8a| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.3.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |b14331ddfb21559b8cc10bbef2534def| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.3.sh) |
