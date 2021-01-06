# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.1.0.sh) (SHA256: 7eb940e64d0e484487322c5e31b2049f029930d1f2533574486d86391cfbc6af)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |326381d3130ff1f4598a09c53d54f2e34101838aa55f97b155b57498f50c4324| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.1.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |a86458b4f8cb1ac9234aeb6a1e823b736616a3e5ae7e97308ee9bd1955a8eb02| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.1.0.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |9221730929c26dfc47b1d0e8143257067e1e655ecb9823aa202b49619355c0e7| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.1.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |d463c64cdd30a1cc220b3dfb0a24ccd0776d71cc9e4feb743dc0bcb90af07a9d| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.1.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |ceb19983d7fb979a0eb2dafb68b6aa3d178e803d919a6c50ff2f333c6f8af0e5| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.1.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |8cbd256e1453f5c675d0646136cdb90c2dd46f9b1cc603ccd621cf4d7f527194| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.1.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |a43524994e43827612e73bc93c8dd413c8c47f141833a76b533f35028cb2f6c4| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.1.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |0044bc4b2a323ad2652400dd8b489acb18d5ae2519a3add7cdb8e62258219d04| [Download](https://update.gravwell.io/files/gravwell_win_events_4.1.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |7ab90e5c573b27a5d677ef35ecf3141b88a0b5745b0bfd08291960f8c14397ff| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.1.0.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |3e70a2f12782435dfc0314823d9826aad37aa90bba25dfe821309e0869d97c50| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.1.0.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |4427093e3c6f6c0e0caa1c6dbbc7310ddc574d60ac26b9a8bffea6b060d2baa9| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.1.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |0dd2e15debaa55deffacaed806da8adacb20e5c97206af1054c593b553992bcc| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.1.0.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |6a8c9cee0e48898b848c1339becad72993462022d2debbdffc7e8958a5b9ebfd| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.1.0.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |4dbdbe3fed9ca10460ce3e2d5087681161513e743aa386f86d07b75cca15a218| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.1.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |7f26b1a37747c0a66aeaa2b87871dcedd2e50fe233f62dbdfc7abfc0fdb87056| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.1.0.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |bcf6755b740b6632b2d2f8f12f52915ac4218e722d8ef696cc8accd0cc34af31| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.1.0.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |90ec7d1a85fb436dfdb47259bcf4d3c35726e52be73401be2201b48abf1c2e4f| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.1.0.sh) |
