# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.7.sh) (MD5: ce603315391c3c5b3527b9e07f43dc48)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.8 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |dae04aaf6c97beb2e50135a046b6cba6| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.8.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |baf555096ca1b427d6acf7fa9874df77| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.8.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |08ad336a199895172965d3285a24ba37| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.8.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a3b694401df452f99c141e8cedca9d0b| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.8.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |1f73a4363bd022790159a9ebf3162254| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.8.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |80a82ddbf716aebf144c58f7487ab42a| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.8.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |201abb09e74d2b955884c27cc49feddc| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.8.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2fe060dfb711bedcdaa19f3891100115| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |61a856fcaa0d7eb4dab56bc3ff0794b7| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.8.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |26f5901b9fb4d3977bb954672782c9c1| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.8.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |6fd44a40264e538b529795e8d805b3a6| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.8.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |3a2dd074914c0098db81a0dd458ece2b| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.8.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |ae3e11e9798eca95ef689603aa89ecb9| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.7.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |9c32a9da5ef7b94dc46df490d7661a14| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.7.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |9c32a9da5ef7b94dc46df490d7661a14| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_3.3.7.sh) |
