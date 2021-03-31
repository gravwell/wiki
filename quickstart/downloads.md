# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.5/installers/gravwell_4.1.5.sh) (SHA256: 2a5805ebe8dbe3b0be132fc9952380360060b92128ded9534aa3c1a99394f2f9)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |576b7795ef28889399ec4a011850033abb9d6f5452f1a199d60794feffe5211e| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_simple_relay_installer_4.1.5.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |d6dc45038bb7f39306c0a4d6b54617277ff42558586810d53c0f0fa0c56fe4dc| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_file_follow_installer_4.1.5.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |dd0df3ca1dd7dcca07f7e642f65c5150ee0a2f0cb816a5fab6eb4f0643b63614| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_http_ingester_installer_4.1.5.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |9d8de23c0e7ca8358533a0da1d6192a22ef5352ca71b6569a81eb9ce356cc518| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_ipmi_installer_4.1.5.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |eb2b904ee507795d1cfefb194246b7197390a7747d1bf7b82fcd13f32a8faa59| [Download](http://update.gravwell.io/archive/4.1.5/installers/gravwell_netflow_capture_installer_4.1.5.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |015c83b6a3550032e5106838af1b12526dec6ccd201ea39c40cfd6643a48c3a6| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_network_capture_installer_4.1.5.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |c7810ee03abddac4a3d4c2fe1316e50f60599f42776744dbc8f0655af4bc9229| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_collectd_installer_4.1.5.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |a52cecececd34efe26ad4e5e4221502f1822a39eea4dd0666937c74b5eb58c57| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_federator_installer_4.1.5.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |315d2eaf5e7ef3f39d18dfbcc35970b0f6758774693d2fe76a5adc5fea6cee1c| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_win_events_4.1.5.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |f6cba02481708f75c1b9fcd3d3f0c259ac97906338d55439bf1a3a96a5cf52a9| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_file_follow_4.1.5.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |0847178530b85bd6dbb79b8e69ec0fa50b819ba3ef23a21b163a6d5e3e31f799| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_kafka_installer_4.1.5.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |a313699f25ef248bbab8935c8f631141358ac7f88b11ae518fb44af46a6d012f| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_kinesis_ingest_installer_4.1.5.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |6bdc84e5a750ffaf5cdb3a12d9831cb4baf026e8be4eae397d06dfa39afb0aa4| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_pubsub_ingest_installer_4.1.5.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |b3088ba6c04cf50779a8ea054ed8cd5d400abecb0493a5c8443fc5b5370efb16| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_o365_installer_4.1.5.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |d5483e4a1dc78beba78c48621b63d4d74ba914f4c45721a7edfeece4a4a43ad4| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_msgraph_installer_4.1.5.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.5/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |458dcb585e092e396974a40e9c0c743328686cfe69c5f6f9ac01ac6c24d0af1d| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_datastore_installer_4.1.5.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |7615a6434731b3bf4b92db3c370db3dff2ec10caeed14acf6582f2f2383cb797| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_offline_replication_installer_4.1.5.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |5bc3f949e0dc1963a883db5bdaf9a07139b4b1868e7e8343d86a5027e6430c52| [Download](https://update.gravwell.io/archive/4.1.5/installers/gravwell_loadbalancer_installer_4.1.5.sh) |
