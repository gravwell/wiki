# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.13.sh) (SHA256: cb15bc187c1643e4c06bffb61512a63c52cca3600e1d7b36f6366fb1ee996e7b)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.13 Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |12793a9ce39362bd45fb4b87d92c5fa97c71732b63164752325fd2fa18129220| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.13.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |0747daf6c0b76030cf79f3f31fcf32c2e54057865b6f130da1ec385a4c408332| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.13.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |a44984c52bace63688478465b6343852047038a1d30908cb29727b718d0ec7f0| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.13.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |6934e64628b2c46f61626d10c92ac7db61611a837ce65d75a7b656055c945056| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.13.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |66f23031043da579e3fe365b5dd7664594cb758a54179aab0c1f17001cd6f088| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.13.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f8242931dc48865551f5f26086bfc5128f0cf4b4855b9fffbe5828d24b4a3306| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.13.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |80e9dea0b5441ece65d76280ad203bc65cd7ba3f7f931ebb50ebd74e5178e46f| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.13.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |52a6dc6a5e6f3fe3b61339711a2ca6ce18e02331632299b6f8be916d945c0256| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |ae644a1e770e8651f1039b3d6219ab455309f6c2b5798330f5a8e7d3531bec9c| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.13.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |36cfcc1c0c2a4db64211a4b355995df79befa87f8010340b5fc39292788f59ab| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.13.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |739f52aef63b060ae7c3ed2b8be54da039acfbf17127ef3bfad57b44dfe93c07| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.13.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |8e22a3cf074fce5e6eba53a46598f21ef3228e360a5115128f67e7e911c02b07| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.13.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |7b249a5252f61900bbec066bd60375d6afaf4b7a17ea18390f1ed949ab659629| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.13.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |58e56e13035e661eec651206b7b1c7247213365e2435b275e1e5fcca3743f96b| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.13.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |649e416dfc6f0285694deb3b00e0a0435aec51af127b89c4e0b82d1b71331002| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_3.3.13.sh) |
