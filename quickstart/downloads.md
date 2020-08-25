# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.1.sh) (SHA256: 9cce033ef7c50937b3169934e5c3c95e10d97be059ac1b212d8960f8e4deeff5)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 4.0.1 Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |c60e51650a01a82e133ee3589595136c1750fcb279c6fd6db3689cf73dd58cfc| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |ef54c24628e382802b0ce73cd393db5777eb3ffed8283e1745c3fac1ebdccfdb| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |879c94fc34477f469c6d441da0c2a6ec5721fe6ccca47cf0582bd32cc8f8ca99| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |cd35beb977a241789953bc18b443d24b086ea08cde335ade4e813a2a3dab0b92| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |1cd224f48b578a7df69b77385c16e3e045b3d7fd44eee39271dcf7cde60e4d30| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |7ed5db840b4758792d0b5c8fba97df100191ad5b89820468bf371e863bb6672e| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |88a170af00fa95016732e693c3f0dc9167bf3bfdb2c70b2e98c8d2064db23cb2| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |9007ea963f18211023a82f44e526867da1de9f7add013e731bd973a68ee4753d| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.1.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |a01d42aa8bae5373261ce5cb99a0729f764ecd75c00b54d93f5f966cae519b7f| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |28dda55ae5489928209673b93306adf1573e5306ee90c35359ae135c6cedf916| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.1.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |6d7d8a720a07eeeb5d5d00e70ea30d21e0108491604d3df9e562e3678642c3cd| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |d9b912efe8dae44ef0c0deb3deed7c9cb070aa43744b884f6f8b5730034a8157| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.1.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |e11db58f0ad836bf7eb3c0c02e84068514d6186cbcbf4b5bd72d91f9476563fa| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.1.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |6816f1635a646e6abf60d6a572e57059b24f087d4340757cfb7aed7bcadfd097| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.1.sh) |
