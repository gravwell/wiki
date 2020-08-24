# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.1.sh) (SHA256: 73ad3379a0990379674048ed769f131523e87b31ed04da24149f30a799487912)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 4.0.1 Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |05b7f54c0f425efc4b8e7d1519ad4470b2d2424fe292851138153e2ba84f3012| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |71acb6f16892411d70598eb68f81987e48563f715a9ab4f82173d5e1567e392b| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |f40ca3fcf6287a64ae4cf1f7bb20cd1a905483426f942c60353aa7f122d81e92| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |401fe46c91da5cde25a03c90e5bc3854a9e3f8033ff05ae1d9601027a3a2546d| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |13a890382f95e5fce9b2d61207d9b8995d65476eb9cadfc5b4f393d9df271332| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |920fc91cbcf6352b1048b41c1c95f1757cf9c2c74f180f97d6526128878f1ae2| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |4c6fa2222c90f9b71c97803672f9a510a878f0289ae97ba69895dc19ac2baec3| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |68245c70c250f84f12759337cf4609316bff51f2e4000e69621ea8b918c54a2a| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.1.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |9ef68b515dd2906dc5c5ba3c605de9e6a4e0530cdbe745bad41ee16be5eaab62| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |d6a03f61db218ef56b112301192c2e03cbe1425dd09535c92a4899fb7d8c1de9| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.1.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |5799957b0907e192e675ed18303e29619370b3e7ddf216d4cba870d44e6fa392| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |e722b2998cf9ee80d695b0d5dd66e2814b167698b8dc2bf72a46643559115014| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.1.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |637805421aa10a212e388012da5bf40a0b56f42226375193acc65fdd4f459ad7| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.1.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |2aea9cdaee884071cf86e26c8f40556a8829dcdd3e32caca28599cf4d6443646| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.1.sh) |
