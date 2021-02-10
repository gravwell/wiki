# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.3/installers/gravwell_4.1.3.sh) (SHA256: 41ada3842d9ab9a09e4841f2d3a15c8d7adaa0c89775347e95f3cb9885c035d9)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |a5699012807f3775b5f0ca533a6306959b9871a88bc67e532d59ee09a7ef7e6d| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_simple_relay_installer_4.1.3.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |a33903077842979a250f9644abd184709e1c0c1180624c2a28b4238bcd9fe9dc| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_file_follow_installer_4.1.3.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |cf9c420bedaf7f310393d55330d67dde8d4a75adcdac21338b4aa73faeff3700| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_http_ingester_installer_4.1.3.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |908fa597473b9adf6a8b19818b906dcfc7d32791851c158564499395a77917ca| [Download](http://update.gravwell.io/archive/4.1.3/installers/gravwell_netflow_capture_installer_4.1.3.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |08a52091fa6660aca987075924648cd662a1a1593bc9df2baa0c15cfabdab593| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_network_capture_installer_4.1.3.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |20884e8c26b98830679f838969509963497c876b6797146bc9ea9883c5442062| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_collectd_installer_4.1.3.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |e58e3fec759e869cc8e4ea2b0dc9d280436d214554652a8bdd9eee7a36e06be4| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_federator_installer_4.1.3.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |26ba3363ead1f91fc5721fd088a36c4b21a9de17923aedc1fbecdde8356a6cc4| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_win_events_4.1.3.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |8fc5fdc1f0a7dd7a1ec87d625d508cd4329931c53934016f757e0ac2a0af936a| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_file_follow_4.1.3.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |610bac8ceb7f69a017621fe5eb2956d3b1beda54dcff220f1b56b032b8f20bef| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_kafka_installer_4.1.3.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |03c964e52695dbb893b36e3395a69cf09265d684522746679ae524652d30f73e| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_kinesis_ingest_installer_4.1.3.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |3038047c833844532743c4442d9a3b7d1eec931218080c4cb627f5e9d0fead5a| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_pubsub_ingest_installer_4.1.3.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |fbb9b989ebf605c040193b6a5db02df7424b9dec2002560b0a133f951a8f65d3| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_o365_installer_4.1.3.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |f01819094.1.24acb19c120f2f42062953ddc96a7c26968d29460662e427712e| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_msgraph_installer_4.1.3.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.3/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |4f6666d63534e2c42c48f48b68214b1b22b70126ce7e6f6435b35e2367564075| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_datastore_installer_4.1.3.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |5c15d1e08e59e2cbabc128118e7ff77081ba1828021c12899a8538268d04bb6f| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_offline_replication_installer_4.1.3.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |66dc89ab3a6f1903f40b8abe191a3e67b6627107b2648452bb911fc634d3e7f4| [Download](https://update.gravwell.io/archive/4.1.3/installers/gravwell_loadbalancer_installer_4.1.3.sh) |
