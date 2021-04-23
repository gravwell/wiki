# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.7/installers/gravwell_4.1.7.sh) (SHA256: 0b025c0f3723f451ed877970c47bc9c1278851c21047bba0a3c3ed35677097e8)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line broken data sent over the network. |95b8498437fdbb89bf15f54f2fc756126ab267b1c91984dd0b3632121099084d| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_simple_relay_installer_4.1.7.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |cf7402cf4076983cc4c4a35d3f85c6fad8e526a02162ef4e1be83e723a5ffc70| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_file_follow_installer_4.1.7.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |691c47f839b425cf9f79bef28cd56e93660e827044becad47e2b776cbb45f129| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_http_ingester_installer_4.1.7.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |a46c54b8ffdbe867452d663306ae2ea7453dd75840710ac1aa6bea1510778d9d| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_ipmi_installer_4.1.7.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |456208028fe57007f15349842a36975dc985a153cd7c291fe3abb596be24289c| [Download](http://update.gravwell.io/archive/4.1.7/installers/gravwell_netflow_capture_installer_4.1.7.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |cc76d76dada517ae5fcbaf76ed637737384377b1ec3feb16706ea25a4400a755| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_network_capture_installer_4.1.7.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |4e1a7ac0476b469964c7ca6c08250e491a8731eb82a206c599293b5d05c316c3| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_collectd_installer_4.1.7.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |14deedc91dfa4653481a669410c852ba84c90082f9479738edd4cb9bfb25973e| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_federator_installer_4.1.7.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |f5b2bc50daf5b224c9767003657a0fbff102057739f296145698368c8bfa2e54| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_win_events_4.1.7.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |640192894e3095899ff7752b4d4d4fabcd6a328715ff179e98c0009fba73f26f| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_file_follow_4.1.7.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |18c63cf131ed55648fe38645cd8653fb7de3bc808ccbc923dda497ad5e1f71e2| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_kafka_installer_4.1.7.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |9e0486398947206c7f9aebc899222626f32e1c010eded5cf6bc6ee4914741872| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_kinesis_ingest_installer_4.1.7.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |ee424208aa60bb3be4f975b890a102666bcc53186bb99b3f28f5e7c5320409b4| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_pubsub_ingest_installer_4.1.7.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |29ed18f17570aa3f8c54c7065196cc757a89e1997b4d24e696957dd49bfb0d95| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_o365_installer_4.1.7.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |4a1b5caae11eb5e322ad2cbcfa5524e58a25f32c8cf344870467bb0007baa3f4| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_msgraph_installer_4.1.7.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.7/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |3a82467b8a337e33388638ce8cbcf35a4efa630362d774b6c1513402c72cfe0e| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_datastore_installer_4.1.7.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |1e7f926c11f3a5d471136dd828daf45a887ce7717cd6ae68eb0136363a8e5c85| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_offline_replication_installer_4.1.7.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |89747432ae0cacbd3d38d61104fd873939dde25f5ee5398c05a9beeb2be8a000| [Download](https://update.gravwell.io/archive/4.1.7/installers/gravwell_loadbalancer_installer_4.1.7.sh) |
