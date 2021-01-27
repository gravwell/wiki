# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.1.1.sh) (SHA256: 85a27d7f67c508a0ca1130670f3cad0d0c92cc32710013c3f6fbe3373ad801a6)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Current Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |e170c21c171ae2080a94d6bbb761c9d56eac07bc587487d869b3073c561c4c5e| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.1.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |d3e55d972f2f0d8d332ac5ab09d281eb046427faf65ef3e62cde0a29be9a7b8e| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.1.1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |958d5b9312925921a23bffb0fc9b91fcfdce40defa610e20341eedb74b3a2bca| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.1.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |4af7cad550766853dbc11b8bbb126674cd1493e14c5445a539a0e6c92504459d| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.1.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |1c1c46bd4daf9e8080c3b2c781eaf2d76690d35948bdf08c7e2260c012686198| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.1.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |2237e907d6e54c41eaa130a4d8e3cb3d6ece7a6172c8f3ae3b6a9f0b4492a407| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.1.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |d42d9dc355dd35233703ef53a8cabe75a6535a0c346860689a5bbdd121081ee3| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.1.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |2d0ab737cd7877e9513ec43e6e6544e6bf9105c65d76a431f43027a254a062ae| [Download](https://update.gravwell.io/files/gravwell_win_events_4.1.1.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |d78fd269868f1e5182b29510f1ffe2ed0dcb7c6edd03e3c89c0779838f83fa22| [Download](https://update.gravwell.io/files/gravwell_file_follow_4.1.1.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |f2465d1e8441302e88f81d0f0c4c79cbf3d1cdd2a34d1fd6ae6322568e66fc37| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.1.1.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |2e96052aad81ec98ec1f2a6dde41f67d0ed212613778c5cf6f80f1cfb02d7636| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.1.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |736b659d64965b474106d476e7ccdb15e24cc3eef7082ce0de1b9341e7e12e01| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.1.1.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |52e81e1aba813d5f634b26469a511b0729e48d4d8f5e166867a4a2c6a4322792| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.1.1.sh)|
| [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester) | The MS Graph API ingester can fetch security information from the Microsoft Graph API. |c135d08bbf1f63e87b9108da39f4e848140bfe3a8f12a6bd88971d71a9af65eb| [Download](https://update.gravwell.io/files/gravwell_msgraph_installer_4.1.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |f13c85086a0fbfc053c69280299fdfb9997510d13ee1851baa34d92a6b9be410| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.1.1.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |3eba439fa6229e46a000c869a9af42ca04a6ad418e66bab93574270f1c42099b| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.1.1.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |96c61c8ab6415deefeb10f40b3ed59c6805e5fe8618823fb408050ca85547212| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.1.1.sh) |
