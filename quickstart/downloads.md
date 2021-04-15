# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.6/installers/gravwell_4.1.6.sh) (SHA256: 24822521e90dfbee1ea759c5de8194849bc3cb396925d50563acbab549555a4c)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line broken data sent over the network. |b6caec6d5f5bc7d3b75d18d241a05e2935488b52c8b2e877a0fdcfbf7207c63c| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_simple_relay_installer_4.1.6.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |63dedf2138ddba4c5db9c8b7cd4d6a46f1a8c08adcc7d8fa720080d100db767c| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_file_follow_installer_4.1.6.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |8f35c445e56d3eff8fe3e8d7db1424e134b22b7b90e73447ec145686387f303b| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_http_ingester_installer_4.1.6.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |f92143e4a2d8ab97b2a3defbac0a69ee59fbac4a7084616a62e23ba7c2e5cba7| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_ipmi_installer_4.1.6.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a4e580a42784f7db5af897c45308186b117e718fb0ed8f14e801f3a4e6114b8a| [Download](http://update.gravwell.io/archive/4.1.6/installers/gravwell_netflow_capture_installer_4.1.6.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |9134d2ccc9fb0055c91dffd85df1ba7978e565cdd8e4fc8975ec81f912e6ba26| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_network_capture_installer_4.1.6.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |eafd451ea9cd1523af492dafa2b1d22e97cb0ef3329c55f5622dee8f82274657| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_collectd_installer_4.1.6.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |f737b7a43a75e746e6e9b1b73be4a119a5cce4949a1a6611663fceea2b0674b3| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_federator_installer_4.1.6.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |3b0fa04d7d6f55c7083d334cb689cb8dabfa11655f8fba64ca19d9f8bf383c1c| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_win_events_4.1.6.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |e6bf94772dcfbc9555e0d9c43fc7d9f9dbb21b8292bab980679d337dfb68aaab| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_file_follow_4.1.6.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |c966e6615525b1f6d5060afa641dba33564bee650d37f13fc48e6474a8a45e32| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_kafka_installer_4.1.6.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |2ca075869cdf9ecaa58e144ac607678689f4c902c8fe34314122d311da7c2817| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_kinesis_ingest_installer_4.1.6.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |04cbf0ba77d665c7c31fa1390d519d8b234b77dc94eb166cdabae78a43a14b66| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_pubsub_ingest_installer_4.1.6.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |f7b8645ffcaa10aa4772a429672e81cf5b2d7386f68a3b0b54c4708463324a5f| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_o365_installer_4.1.6.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |4cc05579128601acf7a006decd51fe59650e3ec142128adb693ec8c4be7f898f| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_msgraph_installer_4.1.6.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.6/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |482feea461e56895d0f7b172a7ed1a81fd064f5fb51913450a333d3b6bd7f4b6| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_datastore_installer_4.1.6.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |841054b6525dcfca6723e9ad3f721a56a3ee1b1474a057b13c7651ebb157e9ea| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_offline_replication_installer_4.1.6.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |7f39ad9c4115fadabf0a129f3b0c0c5818c431f8a791bdbe27959db1fe4f09c0| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_loadbalancer_installer_4.1.6.sh) |
