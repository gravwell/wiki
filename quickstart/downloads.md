# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.0.sh) (SHA256: 0a1a4dd9da16861b9a4888131519c408eeeb7ec02b06ef0e8e5f389f45748e7e)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 4.0.0 Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |e38fba469b10a9a22c08d23feb2d876ac24db35c1d2ba89a21e2a833d3c676e4| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |55dccf663c3d46a72cb7ea741fa373b8967abfa77056dfe0e26148a1064ff3c4| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.0.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |609d29fe31cb8e5845cde60a087ea9005fa61549147e36954ef73807f50af049| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |9fd6219e89d4b1140bccb8f6967276f411bc8d21a497d4edc1d21ef4bd119cf3| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |37d3799e4460bc8ead59d0aafbb02c04d3fa08d5a6f7b3d37da418c4757fa6c8| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |2086b61386b42b3f5b5a66e2dbdb4d2fe65477b4be8af5ec3e4cd514a9996b4a| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |09d5bba5329a1e52fa08f3f5be7e86c4af3bac01b35642ef1b3ee61e3ea571f1| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |52a6dc6a5e6f3fe3b61339711a2ca6ce18e02331632299b6f8be916d945c0256| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |1ec01730a0728c30be26ef75eabb71a5855d0d89c1b6aace33a057c9ad156f27| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.0.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |b48449034e1170b380d3ef72a181a8674864083b2eea597246b67f46a2807dc6| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |4a538bf3e70ab15d02565b4938d7d657dcfd5a0a468a14e9e537149104aa471e| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.0.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |810d56826ed84bad0928e5e9d61cf5ea8eef0962f4cbb96bca33fb5a87a9d675| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |3e71be4c8ee1c92aa550bf898c1c363b107047a281f8d899fb871816e3f31b00| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.0.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |73189421546aafaa5de5a1772b01cd14ff0f031c73616436d24d6b9e507de149| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.0.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |e4b93d10c8aaac10706a5d62ca4c1e43e4beb44769c57276d23e36b8546f1480| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.0.sh) |
