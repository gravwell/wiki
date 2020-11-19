# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.4.sh) (SHA256: b81788f2e76d269964db06fc3995aac42dc1c0303dd0079ce5ff6b35aaf129d4)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |79012c9f2353f560bb454d94594fd63b741df9d31b24786a2328248938ca6135| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.4.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |a57d940b3420f0fee367ddd2ec15738dafa83d82cb191f8ad0ed4f5d6549b04b| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.4.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |72911853d6ec2e94a90087a8a6025d1862655c9839e15efcbc8424885096412d| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.4.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |e494a2ebd03418513e555170b5ff7647db424a4d825557e87d64b6a55cfd464f| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.4.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |2e321de2248d58468caf251e13bc08ace9f25dc4361f6f8934e1e1ad4a829f75| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.4.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |5780ec8f6060abd60ae138288c520502217e8d1b303f5723cdfe4a89c881944c| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.4.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |aaf49e0b9b35d7fd335a4cbe749d652035975e96d35efdf547a1b14e93cf8739| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.4.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |ab7cc3e1b943862ea3c6a18f17aac608beb2dbeb0445861b07f7c7838dc1759b| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |de80cda70933f5e237d91f718da2305c614da85fac203f6ce3773dd2c1488447| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.0.4.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |aa60e2c88986c9c53798438d8bc4ac7b29729727ff42ae0299763799127945f7| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.4.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |e97129bb0f8b4e08138993de2495d6357bf37c546c13bd4191c499a2ad98212f| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.4.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |66c7987b85cfc13ed8ba685e89f9a60ff7a39004861e6fce36f67ac1c163b0b2| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.4.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |ca07b4e334c9ddeb8714b64889d82284b401af2b16640f54f8fe43d73f81df83| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.4.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |4f9b069ab81a5bd0d1fac4b45b3de5b4c30e0ede6b6b77bb5a2cd239b7d4c066| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.0.4.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |61033c13a75ce1a7f10dced889658fe9aca3dbae72310249d45467e1040cf9b5| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.4.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |e2cdf7b542b41949a83f83054baad96cbb00b86804d571d61405df2f9b957d4c| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.4.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |8fc65812dde0261e5ea85b5caa5599da77a0c3c688967c09c5d6cc0a2e5453cf| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.4.sh) |
