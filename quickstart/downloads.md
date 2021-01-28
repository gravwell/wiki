# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.2/installers/gravwell_4.1.2.sh) (SHA256: eb5a4bfe9eead0a85cb47a4249322cf45e191c69e633552a104c6a16ca0e9533)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |54f49ee1e5c0f150f01830b793b8ca19f120ca935bb1633fff726717b829f8cb| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_simple_relay_installer_4.1.2.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |5c58a18e059288d48daa00943a728c0afe55431c8ae1d68a797f2cac34fbd661| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_file_follow_installer_4.1.2.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |16e008479d200ec1eecd908a96f4ecbf7a94f32f41df68d08c91e52b3f78203f| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_http_ingester_installer_4.1.2.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |7a4d3b1aa6019ddf228a4f83b63d7994e7e388724a9035bdb30ed8b4bdd2631d| [Download](http://update.gravwell.io/archive/4.1.2/installers/gravwell_netflow_capture_installer_4.1.2.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |84f5eea402f0aa052ef91f4657458de8f68308e722a5d0672b7db78eb24c939e| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_network_capture_installer_4.1.2.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |eed842c0adf13079745e82998e4f03b5c071ef71b579783df25766cf51655d1b| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_collectd_installer_4.1.2.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |7d3d18c2bf6f36f285a7388221b46f5d40ffedb084b534f673f53fc4ac7276b2| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_federator_installer_4.1.2.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2d0ab737cd7877e9513ec43e6e6544e6bf9105c65d76a431f43027a254a062ae| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_win_events_4.1.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |d78fd269868f1e5182b29510f1ffe2ed0dcb7c6edd03e3c89c0779838f83fa22| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_file_follow_4.1.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |3f0a9b06a8b44bc958c1cd4dc1de648704a518cfd7d812c138f6cbd953ebbc42| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_kafka_installer_4.1.2.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |2fa0895f1c5a7e66a7492e55ffb4b388087ca9c3eca05635ea97e04633e9f4c2| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_kinesis_ingest_installer_4.1.2.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |2b6eac801f3f015e364c38eb97f96b034f2a6e1af8ff86adb38d81945d13026e| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_pubsub_ingest_installer_4.1.2.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |dde7babda483046da7fcd46fc5d627601a139c3ffb4ef7e93008404770f54335| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_o365_installer_4.1.2.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |f01819094.1.24acb19c120f2f42062953ddc96a7c26968d29460662e427712e| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_msgraph_installer_4.1.2.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.2/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |8c4a0dd822aec74af92c448783cbe13e74c3f237da6fe4ec3aa3bd03983ccc9d| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_datastore_installer_4.1.2.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |baf423e7bc61242da98aade4c40731c08f09d8c637cc958d092a3d1c469bbc63| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_offline_replication_installer_4.1.2.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |1d0d43c896dda7db5e092e1555cbd1fa61ca01d0535af1a03358b019dd305942| [Download](https://update.gravwell.io/archive/4.1.2/installers/gravwell_loadbalancer_installer_4.1.2.sh) |
