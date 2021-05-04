# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.8/installers/gravwell_4.1.8.sh) (SHA256: e52a5c59ceba6bf34e4a280913e25d5a44323d2d8671c5ec18cf8464ec946515)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line broken data sent over the network. |295f2e899328a1ac6e5ab2323ef3e37397253079fc217635a26cb443f9a94b0b| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_simple_relay_installer_4.1.8.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |96e7f3323197cc0e0298ba6b2b5a97e4760a9e90f1510ac8f939133bb5404151| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_installer_4.1.8.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |ed7b58eb1da2bda8537220d88b7e3ba0379927dcf7c579e5ad1b20aaf17c5ec6| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_http_ingester_installer_4.1.8.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |73b5f0b10d5e9a0acb458b09c09a80b96e701327d1b90e04abd4fb2615dc4676| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_ipmi_installer_4.1.8.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |74ddf564803797fa7e88078b2e7f9fd7575b52e089e8baff26e01d6a53f9e267| [Download](http://update.gravwell.io/archive/4.1.8/installers/gravwell_netflow_capture_installer_4.1.8.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |8035b316247a373359c2e5a700a7967acbf08ba6b128b54f977dd93d4329da0b| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_network_capture_installer_4.1.8.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |2eb98f841fde008ced239a021748e9cae82a4ff857a28969ff0577b263082040| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_collectd_installer_4.1.8.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |6e13357256f5ee0c809b131b7aab6ba6e6b3f02414978eca7e141319f662cc0a| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_federator_installer_4.1.8.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |fe8b3d87140047c71993f1f5e528d7f1e515a7e27ebfc2a5569232bbb3f284f4| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_win_events_4.1.8.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |e9b0217b38bbdb92f902a960e4bab080802211966e6827aa5c336dc3825d28f7| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_4.1.8.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |569e2b1864e48402a98be09412618908ff0e3da6f7ec4d4e293fceb7b27e3c90| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kafka_installer_4.1.8.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |4b06367dd25f5e54e28010537aacc76ec91fdfe02da41bca6f99bafa64b8f23e| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kinesis_ingest_installer_4.1.8.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |1c7fd23ed5c186acf3b73ab5ffbec51381b7dd0d3b43a20efeb8425c42fbbd47| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_pubsub_ingest_installer_4.1.8.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |723602364a00104e2030a138e61657382c7cf86c1f51890f46ec2658e6048b8c| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_o365_installer_4.1.8.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |14ff4f8194f219e032bb0a911c013ddc26bf818cc689d1204ecde4c73b721e9e| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_msgraph_installer_4.1.8.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.8/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |cba9752a394a66e2485cbe9f19bab6a9bad829d2c2e3f2bb2c87637f333e1b4b| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_datastore_installer_4.1.8.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |1e599ea9e3efe451987759d556d959f6389a590c3e37f6bd200d844472631fde| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_offline_replication_installer_4.1.8.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |94f04a834f2f2c137597329717360b48d8cbb4051934f985f54b80bd3ec319a2| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_loadbalancer_installer_4.1.8.sh) |
