# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.4/installers/gravwell_4.1.4.sh) (SHA256: f9ad5eadea580f9e571abfea612e14e9756f6a92c83b713e917a0fb127aaca9b)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |3b454e9dd76907762a7147d39bfb39bce871672fe370f50a54387ecd1bc74483| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_simple_relay_installer_4.1.4.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |9cde4defed89670d07f43441af42a42f6fcb5ed578894cda8d9321803f7f4ec5| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_installer_4.1.4.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |e96dcaff8429fc8379036483370f9b6d1568aee4e18ad6fe5248f0897ad864d6| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_http_ingester_installer_4.1.4.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |4ce5f967c9936b4393369c86722acd9b880bb668c0f094ac227758572551e003| [Download](http://update.gravwell.io/archive/4.1.4/installers/gravwell_netflow_capture_installer_4.1.4.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |03b29c185de9ec1d292e3cf5f6ef3310a699d0a1b8dd56ac3a6a368d2c2619c3| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_network_capture_installer_4.1.4.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |01f57719a532cef23f4f45ead9ee894e63bf53c5a95552af9793ec8d68637307| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_collectd_installer_4.1.4.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |d95aebf2940014daf820e573c6ea2b02a7fefacbd0ea368810efceca01110f90| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_federator_installer_4.1.4.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |aab40ad4b7a88a91ea3a4d1d231169c15e41d059fc2fd21f190918e21c549896| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_win_events_4.1.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |b1d228df2e3282a9596aabd52ba915f60953f237fb62d4f09c764cc41de8b4cf| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_4.1.4.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |e4d5e81229eaa232b1b30a49f051f47804e132c1d74c471fe9516af7933f7746| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kafka_installer_4.1.4.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |3faec458de971c06e5f295e98ffadb068da39508b4a31ef901efa454276d8f23| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kinesis_ingest_installer_4.1.4.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |830505626ef82d5c59b4ecac0a47ee6b7b7d8a6dcfdb5fbb5d5f2c56b984e73c| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_pubsub_ingest_installer_4.1.4.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |346fcccdfda2acabfb0650808d588530f1d5ecacf8a734df737b829c44ff6fb9| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_o365_installer_4.1.4.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |f01819094.1.24acb19c120f2f42062953ddc96a7c26968d29460662e427712e| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_msgraph_installer_4.1.4.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.4/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |1c79f81f2743a0f2bd78666282399dd35ac38614bb7336936ff81e5303c75a1f| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_datastore_installer_4.1.4.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |6b1db29bd241421e0b3670cf335f5fc61fe53a2755d9b25475ff090f6a3d94a1| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_offline_replication_installer_4.1.4.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |fd0f6d7014d8ec1f906058d9da01548cda59a205ebaf583f13dd80262f509065| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_loadbalancer_installer_4.1.4.sh) |
