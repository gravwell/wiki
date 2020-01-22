# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.4-1.sh) (MD5: 0e4697e55493fea176e1cd8d7824e344)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.4 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |dc804809b8665afe6f73167d092ecf5a| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.4-1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |2b901acc6a90e174217dbf69fd53ae3b| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.4-1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |a99689b4607e1f21d5315e9b160f728b| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.4-1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |16609732fc36bffc668b543234306e12| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.4-1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |01d66074e9708bbb2b0dc7f118464350| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.4-1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |e583bd95dcca28137980978182eea805| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.4-1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |cb1bfece177f3b57c94c7de6986b7160| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.4-1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |1b21e815bb1c35243e7b2c5a036caead| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |388516cb6c2067028ba9250f4da5e179| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.4-1.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |942d38a0b8dc1972d8be22496be4d5d1| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.4-1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |02ae16ec8c5614510b6d2bf79f9eb7ec| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.4-1.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |38ec2ba73467ef467a352d545a37f714| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.4-1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |8b13306e5f307432eecae56d1e02583a| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.4-1.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |296b62c3eed24026e9bdd477acd69c62| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.4-1.sh) |
