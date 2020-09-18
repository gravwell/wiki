# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.2-2.sh) (SHA256: cef2e8c370c0271eab5afcb379477a8ed27cb03119b511023da5baab880bb29c)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |dc7c77fc3b3efbaf4c39260d70d34e8885c1fafb004a58db7235d06c4e9d1e50| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.2.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |f451fd11bbc6a950d9df13e2f17ea4d8dae9db6348526395701c20cbe634cb50| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.2.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |05fd887865d59c98924326d90ea7d04e9a81835b27174c4a96727e7a51d483ad| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.2.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |3aacb7deaf6c955b4988d80cbdf2a04b20d2610899daf4eb658b1a9bddbc60d6| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.2.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |11069d989e1d9023f9396575279395410de307531b204755afce1a309c0b23f2| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.2.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |4af078b2bab39caf28c460bb439b3bae9e49e09564787765141cd7d82da34f15| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.2.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |4c4c231f62ba9585f11c1710703143d0c7e2f987c81bfd47e05a96047ccc0b36| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.2.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |182bf634afc248ca0870c367d25b9a90d58924d8c07bcd25bcd2ca52a4c85ae0| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.2.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |4a3822a9386030fcae1a58df0ec39d2071cb2348e2f3cdc5a8ae30babe5a6727| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.2.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |257b300a04cd5fe565052ac417eb317cda68e2955040f320a6162f70c5c26bde| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.2.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |9db7a2cee8d2e5b99e7b0df7819e47f144ef14556551ac29c4b6f4336d0d8969| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.2.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |e9ba56cb59c180e8908809b14909484c15c019fbcc6b1876bad7c1b6030b408a| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.0.2.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |5f3d7646e316a2f672bbd22d48c9bdfacaf288805e8cf3b88484945ed3f03fe5| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.2.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |aa7936c718de5f49f9df1cd0d21093103b03eae6a6e4e07c59b3b2e36559308d| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.2.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |0b6a2103dff851c520d4b6f4a90dd721b5fba91a2553906929b9b40836226128| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.2.sh) |
