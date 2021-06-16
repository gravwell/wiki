# Downloads

Attention: The Debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.8/installers/gravwell_4.1.8.sh) (SHA256: e31009ef3c016937b736d92bf25d6815dada4b45d74e1098d96a3ff1466d8ca6)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line broken data sent over the network. |aa35233b114e1838f7744f29827c7b08227304d92b1235058bc8ae658d813b17| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_simple_relay_installer_4.1.8.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |3bb8c0e53aed4b143e11fc9986f355bf6fbd1c925e92580c84e25f8f27fe8e19| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_installer_4.1.8.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |ea0e4ccff6f274532e22ac2d65b6c6c56112965b49f5a053b8f95078c6a358da| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_http_ingester_installer_4.1.8.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |cd2e9d08da1d14e4b9e752df1ba6f6e3a8bb29378b7075eab34752a9bff64288| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_ipmi_installer_4.1.8.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |a5078e250c1fcbf4bde46697cf788278958ca49f5ef2692000e077b77924115a| [Download](http://update.gravwell.io/archive/4.1.8/installers/gravwell_netflow_capture_installer_4.1.8.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |6263eaa86a5b18075f2029fc4a819e3aa96ced016cf90d1708dccd8a864bdddf| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_network_capture_installer_4.1.8.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |8943a74361bc36a0bde9720faaf1761a472a98016a8a336a17716f6c9510ddea| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_collectd_installer_4.1.8.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |20dc4d5a98f1da3c851f36bd3e42a23bdde6d1735e730c459dd44459a283fcb4| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_federator_installer_4.1.8.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |eea6ac62cdcbe50e31a783ae6a3becb7e3cab97f842086a94b18de876c070e52| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_win_events_4.1.8.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |c34aca1492ad9544140ff943c218116d4ddb735a853f91d508fcde873f33b455| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_4.1.8.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |a60c6a1ca15974eb4cbd1c5515dccf6cd33c301e0d7c319dedb20c23b480be09| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kafka_installer_4.1.8.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |a00373d067fed6c5d427408404bd9067394fc953f38900d66301a6ed8b83ee5f| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kinesis_ingest_installer_4.1.8.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |b05e1159c3046f5526a50ce5183763d2577b7e115ee3109c9a26e8a0cd1cb484| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_pubsub_ingest_installer_4.1.8.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |2678b4d8ecc7b61ca975ca927e9683d3522957268d60f0d45b1c56e8313cf523| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_o365_installer_4.1.8.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |8c402031b931ee5ac6b3136aa978e81719e4daa5e049251c02258ac56cea6f56| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_msgraph_installer_4.1.8.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.8/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |c4a9993134e3766107820fda16876658dd7ab750dc2af4f296980cbb0b0b5b04| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_datastore_installer_4.1.8.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |82b04e4469204e6e8109fc508b92e3ab40edbedcc4320617b9def254606cfe26| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_offline_replication_installer_4.1.8.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |555af582521c535a8f0495b5622627723a7c82ab629961a0b2c4f3f07c9355ae| [Download](https://update.gravwell.io/archive/4.1.8/installers/gravwell_loadbalancer_installer_4.1.8.sh) |
