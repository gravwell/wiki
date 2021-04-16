# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.6/installers/gravwell_4.1.6.sh) (SHA256: 132d969e9ea45083e576b9f6b3f7dcdcc03102a06619131e94049228792427c7)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line broken data sent over the network. |6372c0e90d15b35f7383709c9cada2f5f1871e138dbf249e4a0149f087fd47fa| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_simple_relay_installer_4.1.6.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |17d9de0e6424a1ebec12886c9fab6208f801b1c6e1d6316beba992f9404b34dc| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_file_follow_installer_4.1.6.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |9c7f0610b99e45aedc7e4d64b215997ef081c1f27ec5bbffaad1e074ed601cca| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_http_ingester_installer_4.1.6.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |b01b647d2e064a0ecee3a06751050e8ef4962567d4822f16563dede186cd9258| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_ipmi_installer_4.1.6.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |39365f21ffbf6f4af98401e443d8620c3852cd8e37fdcc0cc4fc4fb5a4d36977| [Download](http://update.gravwell.io/archive/4.1.6/installers/gravwell_netflow_capture_installer_4.1.6.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |91a679bb0e72935c7193515985c4f3534d06ae2f5d82b79794ffa496726c730a| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_network_capture_installer_4.1.6.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |0517ecbdea7267697ee24c5f5811c1e23036e305f77b7d1f13e25af9a2fdcfb4| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_collectd_installer_4.1.6.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |96829d64837b34abd3e4d5c63a369533e897263296374ac2d5ad78eb14609cf3| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_federator_installer_4.1.6.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |75e5a96aa7bd5fba6ac837fb2e8189575f13653731f703b903259d7bb6e73400| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_win_events_4.1.6.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |5c413f534dca6d97a5309beb8de6e8f7b62bab9b30005df6b9cc88446be1e178| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_file_follow_4.1.6.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |60bf8b8ffe41e2cf5063b098b4ae7cb171628f7eefa32fb5d06295fae3222f58| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_kafka_installer_4.1.6.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |f29c79c9d0960a506ae66b11de5652eadec0e39b001a7a27d35d2c80c74df2af| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_kinesis_ingest_installer_4.1.6.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |d6410f437b8a1d07c8ecaa0c501fcf19826d9c3a2ccdd2aa0d2aeee2672b66c0| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_pubsub_ingest_installer_4.1.6.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |99f09887f0ab075a79d79a095465ef0e6076dee1e934f66641ff44d03a912866| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_o365_installer_4.1.6.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |ed9bc76f1c99a4b1f811332940d5bde8027e54124ba33029ac36a7ba264f6dbf| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_msgraph_installer_4.1.6.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.6/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6fa2cb2ff09b957a34985fdde0d8e7512ed9bad9c42018648cb12d8f4199bc7f| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_datastore_installer_4.1.6.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |48b4a6bb4edccacf0e03b0fcab412890a80682a1fd01d043c565e1e635b980fc| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_offline_replication_installer_4.1.6.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |0d6db35340d61109a4ecb624424f8ae640afbb944c7b89dac889932bf48ac4ab| [Download](https://update.gravwell.io/archive/4.1.6/installers/gravwell_loadbalancer_installer_4.1.6.sh) |
