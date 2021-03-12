# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.4/installers/gravwell_4.1.4.sh) (SHA256: ea6f2ccc7fa36fcdb0700b150bdac7f567f4793fe5a1b3a507ad2ab3fcd9955e)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |04388cbdcd97b646d3c5c089d7e0237a1a60a460676a4a46c9297b08ce575d59| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_simple_relay_installer_4.1.4.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |3c9a2d105aae4848219db3285384cb4945a0d6b64f7786fb77d2e04cf2b354b8| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_installer_4.1.4.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |7bd167b98127bdd3b16cbae11c450df09a51070de50c271091923b4cebb20099| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_http_ingester_installer_4.1.4.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a5a79257692342bdd670795f633bd29d4614e18bfddd4262854b74a5b625958f| [Download](http://update.gravwell.io/archive/4.1.4/installers/gravwell_netflow_capture_installer_4.1.4.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |ff5ddab63bd27cda8ab3eca139ef62ad17ae32ff4ec884edbd7c615fcdf13da5| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_network_capture_installer_4.1.4.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f747bbb0bd52820466d511f727ceb8a7367042e31727aebb05901bb4d262e9f7| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_collectd_installer_4.1.4.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |8dd6ade73f6a5bba632cf5d5af70afa91e30e70494508a283c4f1ecddfd56df0| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_federator_installer_4.1.4.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |5abfefeea74bfa24ecdd838908d6e86513dec82a486114ef112c6627376485c0| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_win_events_4.1.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |5a61cfa52115f69a6e99e528b5d5eba264d1217152d43884fe0e14e05b45ade8| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_4.1.4.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |751b519a94c0f4cabc298df13a181b417c04d5e1ef46f30d0c433f6cbe61c2a6| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kafka_installer_4.1.4.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |13232d9795a404a74b9f50b38644227aba339b93325203dc51799502926c0790| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kinesis_ingest_installer_4.1.4.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |de1203d6c8ad3950a8bb943e1b080c56281629c04c0913cd26063069ad4982bf| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_pubsub_ingest_installer_4.1.4.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |9ce1efe4fcd161c2e022710828a2a570b6fc263d0001358266c57823e1207f93| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_o365_installer_4.1.4.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |fdd1f2eb8df146da8e029cc2224a1b45a3317c73117044925203f051555181dd| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_msgraph_installer_4.1.4.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.4/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |efe7fbb4f3f921db89be2ffc5f1046b35f1ca4bcdaa7f062f767d3a928e34236| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_datastore_installer_4.1.4.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |71477e1ab38be08e842ff6c8f7939e76193dfc33d923f4eb4d1cfc82c5e48235| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_offline_replication_installer_4.1.4.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |56d36c0b896356e97bd9e4fde0995e26ed8f795134b01fc3e65be42c7bb22f27| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_loadbalancer_installer_4.1.4.sh) |
