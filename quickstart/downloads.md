# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.3.11.sh) (MD5: 8675bae7bd16308e00adf5d2dacb9d10)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.3.11 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |c1458ab06b40bfcdd18d93f558bde245| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.3.11.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |7944c4cc5b323df307f692799091f08b| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.3.11.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |8009430ff5fbd1e2742a8df963b76909| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.3.11.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |eccb7164c04917ca7a85160d7d3134e4| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.3.11.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |473ba093c9b8aea7cc851f7bb2aa63f7| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.3.11.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |ad27e0612bbeabe56959c779a957ed28| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.3.11.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |e2c6c39f5f120dbb9d47b2292e1e929a| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.3.11.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2fe060dfb711bedcdaa19f3891100115| [Download](https://update.gravwell.io/files/gravwell_win_events_3.3.11.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |9a6bc89d62f813e7e679522d5a359e14| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.3.11.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |a38aba7dd967926a0febc668a7aaa46a| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.3.11.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |da06661d1b9ba4bc1ab4e965570627d0| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.3.11.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |65e3300278d8d8b6a02f7ef1b6619112| [Download](https://update.gravwell.io/files/gravwell_o365_installer_3.3.11.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |8acfb13804f2f223c7d29e522a1cc346| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.3.11.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |3ff79f581b9800c1e822c5cf8fe505f2| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_3.3.11.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |cacdb3fbfed54711acf8b78b1f950986| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_3.3.11.sh) |
