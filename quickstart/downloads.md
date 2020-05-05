# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.10.sh) (MD5: 8af2e1267a404f4ffd0a7984a15a0bc3)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.10 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |a41c7a08eb9c4a27a0feb6c5c56d815f| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.10.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |f7bf10f2b8617e352900a8188ebf0b8e| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.10.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |7f5dd3528f3b610321b48b3e377d5e6a| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.10.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |8efe3a20966f6c1ebbe9fd32f9b1a2eb| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.10.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |0cd9122a4f1bf53e48e9c8879bdd3d1b| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.10.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |e81b49682625914a09125c881bef5204| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.10.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |81645dca6dba657cb5603d89f6722856| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.10.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2fe060dfb711bedcdaa19f3891100115| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |4855766033f8231b24a2a15b80f6747d| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.10.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |f397bfe2fb3e6d1f4a2d28ff83ab7dca| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.10.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |45af7e48f66d4ec512ec4b71ba614974| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.10.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |1c05d9098472992a355f200205454b64| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.10.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |65791b6cb6ae0c7cf563c2292a220786| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.10.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |ba247e171bf6a0e207de963c8871971b| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.10.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |ba247e171bf6a0e207de963c8871971b| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_3.3.10.sh) |
