# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.0.sh) (MD5: 45aa60a25da1ce466f8975144cd8d008)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.0 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |0fa217925153df7ccf66d9e4aae1f110| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |4948b36400da47b44a7346d678e76520| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.0.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |16f672044dbb2d85df0b9c79540a78c1| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |a48a672b7be1cc1cb5762e3e648df15f| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |c391988cf5405accca6ce04c682b7dbf| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |5cee2f091f65b06f2c04d5acf3680611| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |242dbf78cb9f65797a113013f5457368| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |1b21e815bb1c35243e7b2c5a036caead| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |4025247f3f867b889f3f6588edf354c5| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.0.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |82d0b46d5bdfcbd4ca4d0fbe91aa81a0| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |1701e42d213cc88a80476dfaafdb906a| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |962164d1a358f7a2647933df454e387f| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.0.sh) |
