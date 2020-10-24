# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.3.sh) (SHA256: dc516958b50c08b6a7b3f398bcaa11f30e62cf11cd8251963dcfad9ff8cc0a69)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |97d6e0f451ebf87456133e29d409c8706842e0a4537ab90852216b3fb30c5423| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.3.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |e0e4d148cf29c5b92499e12ee8ea732a0adc5106a136c653cd91d0432080264a| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.3.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |163da99bf10ee337f88c6e9e5dd6152c02f40002beb6582a71f3536e978b29da| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.3.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |849d7d99e3fc4e14eb7e9e979b4725c28305ad945f05f5de564dec76ae1f90db| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.3.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |eb08c60daf77f3308f2f8abd8b984f7fd46d994e019fd7a8fc96c0928a895edc| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.3.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |02c352e0d0104d95c24050b8cca658485b517e7691b5910a1eb2fe51a5986e20| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.3.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |e962f46b74ad93035dd1b7c3451d912cdb288dd6b84498267a035a0e156f4fa3| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.3.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |570ff54992c8adac171bc4ca0dd724571d2e91d4af0c1a8b8e1f627f7c19ef2f| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.0.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |0f1be7cd487d5dd9fb92de83fef11f36ca4118f58bdf704ec22b89570d9ab075| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.3.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |ecdca90f636d519c86abe66741b335f5c3a2c75d38970ace2074f207b219a060| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.3.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |ef02022666450ac674277cd0e8fedb0a72dfbf86a0ee5d9d3e89d1286041fd85| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.3.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |33608b987dcf8ab3ceed198d3e0091922181e0190170c68ffe3498e2423a353b| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.3.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |148a740cfdc5c25fc3fadf3bea635d59fd062d5c2a9848dcc0cfd5c7c83d8fd2| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.0.3.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |af9b6c83c4ab7b3e760db5bbe78a3f58f7d22645d2653655a3429cf511a720e5| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.3.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |d197fa67c37ac84a9a39feb2db09c4dd35a2ccfd8561b69cb1f3775f5a880f58| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.3.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |6c3e5cfd192c801cc57e539c89ef97222d3c3e2cf82b0e56a37fe8a11b853661| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.3.sh) |
