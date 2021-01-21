# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.1.1.sh) (SHA256: 7eb940e64d0e484487322c5e31b2049f029930d1f2533574486d86391cfbc6af)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |7ca6fd0cdb4967438213a9ec2be24c263761b39bc3ca72b28570ce9a0ccfa85e| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.1.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |cd613bdee81997c2f5a920ca410b3bcb03310544747816fdd3233595981c3ea2| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.1.1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |994388d3735daa5d6278a56704745423781c0049f65fceb776558dff600ffee4| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.1.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |f9c7cccef2538bb78b8f7d98f28011ce3f5f3c93dc2b4c1e78082837fb140956| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.1.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |d4964a600fe8c87e65a5d947c2a600ccafa294497568d2571e0f6668b4eebbc3| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.1.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |54d0fed8dbdf717c08d89102afe3045097b327d5c93fcc75b93892a5eadc7fbf| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.1.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |a8a0d951acc02055dc58748f510ee0cad54926b682dc54c0f883f04045421ba7| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.1.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |0044bc4b2a323ad2652400dd8b489acb18d5ae2519a3add7cdb8e62258219d04| [Download](https://update.gravwell.io/files/gravwell_win_events_4.1.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |7ab90e5c573b27a5d677ef35ecf3141b88a0b5745b0bfd08291960f8c14397ff| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.1.0.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |fdd480e7c117209ec6d042cfa59a888ff0e568fd23e40c73fd37e84d2c8472e6| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.1.1.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |4b70935210b4b64324d0670d8d9421c5a76f4f6725351bd5b6aef578f7f0cefc| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.1.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |c10cc6c50d1bd3caf2568478405c5a4a954899ed4ec0b549071eec324063795c| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.1.1.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |f210be1b4cd30e8a84e71adc29ca0defd6a366646cd8c47f93bd06888aefa475| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.1.1.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |74283c3ff9d8b1f7c2c89c1d00e671fd88ed80e9c27809de06d782eaca3d1ed7| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.1.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |e965db9fae2fb945300054db641f3f3827a161089f5dc5c2d40a4832128e58dc| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.1.1.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |bc4c0a26f3b2d0cacab90e76fdebb173685cd90055276f5025a82e2c5272183d| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.1.1.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |6ddd702a231f766c1baf64ed75bd1e204ba7a81c394fa10b69bfdd8fce3b03a9| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.1.1.sh) |
