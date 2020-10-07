# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.3.sh) (SHA256: 7eb940e64d0e484487322c5e31b2049f029930d1f2533574486d86391cfbc6af)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |faa2c50a080cba0420f0aaecea170ea0a028eaa80e77529cc15c7c6845e2a2f2| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.3.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |f0edf678f85d76d894082e3ec20fe9cf1f1c4cdfd1f60e126276f6a58dcca099| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.3.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |788286684fd94970acd82b23f5a2e48a4adb4c405b6aef6bdabc1e8a7aef671c| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.3.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a915562df294e0010866524e7abe0d57a1c5e5add2fa60886b2b6ca0cc9118d2| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.3.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |e4f1b85a61343c449675daaa57fef7c67336b8632ddf1e0610c6f36b4d64ccf0| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.3.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f1c3bcb200c78dbb03ab42f7731ee31203d1d3512e1586f2977841c527ea844e| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.3.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |115bf449f3f86d5d4cef176087ece383d33caf5d0f48b14a360c60ed2a540866| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.3.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |412fd43c4ce96a80b9532d350beebd1532631d46362dfc3f79ecc7dbd795c62b| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.3.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |befe68036d039a07daf8405eef351685a591acac7a90c457967dd95183d2f015| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.3.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |749b78312e15a8afbec4e74a6e4cf7627e907ff0fd75b91d836713f18b8063a8| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.3.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |1ac5707f46e684d15b2947226bcdd35a61c382d8ffef6d9d81b3ad7a45ab5d52| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.3.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |2c43101e9e4514f0248be18fc9499d9b12895a80136975c9651c7f1c1cd85448| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.0.3.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |3624fc91cebf8ef11e7687225abbbed97f3560ad79f1ccda405b3f25193f1e1d| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.3.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |6fc00b565fc1263c49154a641b6e19aad628c61120c74f565b27aefc31e80250| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.3.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |ecdef85d18c0025daca2b54e4414094e1dc4e1ded78cbf5c4bba2b932136ce1b| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.3.sh) |
