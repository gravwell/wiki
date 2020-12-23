# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.5.sh) (SHA256: 7b9082b4e2194910200923df9d590f22ba6e06eadeb862b4362a86a7174b7687)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |8d180c0d3be8d07ede3d0eed94ea110efd3648ccc3c82bbf0a0a7e1e7f41043c| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.5.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |90f8b768cd896515486a49e850d7d556a228dd2be809a99c5d00c9d1404056ff| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.5.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |b8585d1439598a41270867914cf4268ea8f77d3bf1a93c71db9159548c56a9d9| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.5.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |72e37944d5916465d623dea35ac9efc48c27ee9e3148e842c6d46ce5b68e83a7| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.5.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |5e4b8549439027545e3b09bba32f83c93eae1406fa22cfa1734e0f2e24b0c4bd| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.5.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |7c9cac5e69e72871b5772278fa739c69665ecbaaf10fb5d140c6f69cddf29f79| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.5.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |cd0119bdabfaa7af6a23cfe6dca472dc75c871b4b4e9769571932202e58493e7| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.5.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |ba1add5d245ffa15922bcca989beb686eca978c4612499f81b3fb9096dd9e927| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.5.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |b3d6d371e428ec46aadf8157e4660973985a1beb4302e6b84ffdc723339bb01d| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.0.5.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |9c16680a3b07e323ff7b62de3eccdde3c14a0f37598be140b8be536505f97e41| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.5.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |6f3f16458585f065a8068f8c194134a925cfbf86e022b9956f3ca7cf362f634e| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.5.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |92e2b6cd152721aeb0ad8d97e5628e0f8813ca696c621c07b2aa79ef2f1f84ce| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.5.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |d99cd586eeedd5e9aad7bfa73596e99e4b44dec906c984a39f97992bb9d150bd| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.5.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |670c22b3a53ceebb42cd5d907172490d371845a2d0570edb25f63a41c9c76f29| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.0.5.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |8a21e27e5747804d555994f10407fa73d6790ee8214106d13ce0a9666bcfcec1| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.5.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |6c32004a3e1ac94bfd9a2d3678605865a637745df905c51d356af1891e05c107| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.5.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |c6e74741e513d58b6af8934b934343642bc4da8691b8f4f8c01ceb745efcb72a| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.5.sh) |
