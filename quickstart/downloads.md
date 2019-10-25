# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.2.5-1.sh) (MD5: 57842cc01d6902fb2dc0306a78cba7a2)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Ingester Releases
| Ingester | Current Version | Description | MD5 | Download Link |
|:--------:|:----------------|------------:|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | 3.2.4 | An ingester capable of accepting syslog or line brokend data sent over the network. |d420befe41f300ed4f8bc2c75e5666f1| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.2.4.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | 3.2.4 | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |24b022543d32c92648ee10eeb968ffed| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.2.4.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | 3.2.4 | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |98a6d05d94a3faa189549ac8994eca0a| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.2.4.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | 3.2.4 |The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |fadb3894e3937423e491d9d8ce4c0028| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.2.4.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | 3.2.4 | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |3d0f7a720b0868734d28461b9a67a2f0| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.2.4.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | 3.2.4 | The collectd ingester acts as a standalone collectd collector.  |86f95558485edfa272f78a0c6eb9bac4| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.2.4.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | 3.2.4 | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |399543b21a1e774087fe9f6d6c1264ae| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.2.4.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | 3.2.5 | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2310f316af89d1002ef6fbef33b9a465| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.5.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | 3.2.4 | The Windows file follower is identical to the File Follower ingester, but for Windows. |aa8bd348e847fd593c41ca4e0ff679a6| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | 3.2.4 | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |1898644d2886801f35ce4aa044fe4657| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_3.2.5.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | 3.2.4 | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |2c3192386b73d97e017bac917a1d8f6b| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.2.4.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | 3.2.4 | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |6f3c5c21e0dbbe946213fba92b4a4d16| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.2.4.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Current Version | Description | MD5 | Download Link |
|:---------:|:----------------|:-----------:|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | 3.2.4 | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |1e1988c328eabfdd57e28e85a199da07| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.2.4.sh) |
