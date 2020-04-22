# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.9.sh) (MD5: 106cf61d86f2e30e314c7d38332ece0e)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.9 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |5faffb871a2253f667d047b8f2bf43ae| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.9.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |e3cf2960050c6af88de5b8f3c2bf4ea8| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.9.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |da1a41d0c46f4962cd9fa05cd1cad8f0| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.9.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |68a3eebdfb0126ab779fe6053cc30dfa| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.9.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |f990daef4d956fa276a17ee2e162bda0| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.9.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |afa41db1e81c60735d27cdbf4a6ba456| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.9.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |2e76c047f982cd7646f987d0cc6bed8b| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.9.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2fe060dfb711bedcdaa19f3891100115| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |3d555efb2d99f36c71612b2d6d82a2a5| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.9.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |04bec490080fdc99d5d98883a3ce96b2| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.9.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |eec62fbb6ecdcd5eb3a512cfdfb24f8c| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.9.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |14d750fea30911adfb958e2ecf87c779| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.9.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |fe9031501467473de0d76931b3ee413f| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.9.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |eaabb8c470ab43a0984d506d6689b019| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.9.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |eaabb8c470ab43a0984d506d6689b019| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_3.3.9.sh) |
