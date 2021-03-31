# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.4/installers/gravwell_4.1.4.sh) (SHA256: b23aebd5098a1010d6c29b5dcf5cfa73209326cf37a2f263c2cf88baa653f4be)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |574dfbc294b5c7c149144923ab3b03bff263732924185e84a5fbf33349b4ce3f| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_simple_relay_installer_4.1.4.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |ec10d847a3fd251d514e65e6b9a4e40d830987f1b0be0f064823b8b3c3e08da9| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_installer_4.1.4.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |3dd620f21563f497c18f54263d950e119f1e513b4ec8b685f15110337bd3f333| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_http_ingester_installer_4.1.4.sh) |
| [IPMI Ingester](#!ingesters/ingesters.md#IPMI_Ingester) | Collect SDR and SEL records from IPMI endpoints. |ca89104a25641f8cd3568c91b97fe74add52e25477e248dd2e92e27d72d93ade| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_ipmi_installer_4.1.4.sh)|
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |ab2887d22c662ad62a50ba77917e3b39c6b3fecd083e7aa86400c9a5ae81f9ca| [Download](http://update.gravwell.io/archive/4.1.4/installers/gravwell_netflow_capture_installer_4.1.4.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |7b33c0414d66506fff7458b00d175b3a732348a3a6a03cf76515bc61c150f4d7| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_network_capture_installer_4.1.4.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |fa2f52e79a4981383842c1edb7c91a9bd085866795c7f604cd22ed61dcd1e942| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_collectd_installer_4.1.4.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |aee3c81d0901bc5b2b48f64d98ab36cd0f4f0b40f4fa7d0633becbbe3d455056| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_federator_installer_4.1.4.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |57489db05d48cd50dd5f9a370612ffc0ae1c12182f3c548e6ffd32ca254205b5| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_win_events_4.1.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |4bdf85c5a8060196f468571645680152f8d50115317c553be3193e773e9f463a| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_file_follow_4.1.4.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |a8ae2f698e51153c977512ad890061dcc9800e8a9c1d49be4477821217434768| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kafka_installer_4.1.4.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |32f003bc18a5890f652a3226c3dde3ee24debaaa0dcbbcba69559c344074f16d| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_kinesis_ingest_installer_4.1.4.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |b4240f1a7305b34678a1d0e6f00f6218fc519a3c90df5c73cc7cb00cc7f64a81| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_pubsub_ingest_installer_4.1.4.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |50860ae54416bec1dab8dd09c47eecaec329d143e3f769ea33fecc2a17c2020a| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_o365_installer_4.1.4.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |79905c31e24861f32e2aad06054b3c3a7c90e4547ca9634f5cf39bc1ac82333b| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_msgraph_installer_4.1.4.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/archive/4.1.4/installers/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |ae25d83f6704ff011b72e0138507bcb76e9efab4ded3efea804298b9b25d18e2| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_datastore_installer_4.1.4.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |d6d46e1cbe5a4a64dd52910ae27e65c71e68f1aafb9303b404df161ee765ddaa| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_offline_replication_installer_4.1.4.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |aeda5ba995ecba0c4596a010ba44103573aea41898ceb032f13c190d6a85bcb3| [Download](https://update.gravwell.io/archive/4.1.4/installers/gravwell_loadbalancer_installer_4.1.4.sh) |
