# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_4.0.0.sh) (SHA256: c4dbe793583a1319108975c029361d1af6d7aa26cec4208529af66136365f7b9)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 4.0.0 Ingester Releases
| Ingester | Description | SHA256 | More Info |
|:--------:|-------------|:------:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |e4485a476c57f5f14c909fdbcd4c2fc70c011a29312ff317610059052a4e729c| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_4.0.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |d1123d235d306c6a52cfa67499e0eeb967ed1d65a296b0debcfdec533d75433b| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_4.0.0.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |a6a51df12023c5f8bf75ac2a8fb0c9f220356a701fe3b9c19b6b58b927968cdc| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_4.0.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5, v9, and ipfix collector, ingesting Netflow records as Gravwell entries. |b25ff9d0d052a007b3ca961f442b342134efe1412e532b56ef2c000adb32fb94| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_4.0.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |fee1d99dce1b5863e51be2d490206dee557016fef5a02a5b9313d6773ee367f8| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_4.0.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f778d9b166256e15a2fa5fa49a368504c7d31d8f955bb437608028536277fba6| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_4.0.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |a80ab489ff8036c873c0e35680d289813063142062f8896a2fe9c00a40773254| [Download](https://update.gravwell.io/files/gravwell_federator_installer_4.0.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |e92965d04e76b3496fcfde4a90d52a9f8b80227a730f35ad2353db9c604cbf70| [Download](https://update.gravwell.io/files/gravwell_win_events_4.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |79487abdf0147fe7084fd5fdd51ad43b26e1ba47675988a08536e2508b9b37d8| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.2.msi) |
| [Apache Kafka](#!ingesters/ingesters.md#Kafka) | The Apache Kafka ingester can attach to one or many Kafka clusters and read topics. It can simplify massive deployments. |d6cb1bed7e84c6df4d96528aa95e5d84eb31278c9b655e0e087a41e67d792c43| [Download](https://update.gravwell.io/files/gravwell_kafka_installer_4.0.0.sh)|
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |0ed77c44b3a360e9c1542170307243dfe18d0f56ce30b213b55925ce0b9bf940| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_4.0.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |9a2600038ca66a98028c9a94f4a4d526c7972e2bc689a778229106a65fd0db46| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_4.0.0.sh)|
| [Office 365 Logs](#!ingesters/ingesters.md#Office_365_Log_Ingester) | The Office 365 log ingester can fetch log events from Microsoft Office 365. |46eb372bf4714cc8d274682dbe6d8ebfc0f8fd245790fc8e6c0ded7fbc9ce169| [Download](https://update.gravwell.io/files/gravwell_o365_installer_4.0.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | SHA256 | More Info |
|:---------:|-------------|:------:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6e81cf4f02baab41c2604deb2323564b23364d0ee2801cab9ca4f4cdb819447f| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_4.0.0.sh) |
| [Offline Replicator](#!configuration/replication.md) | The offline replication server acts as a standalone replication peer, it will not participate in queries and is best paired with single indexer Gravwell installations |5648e4632f14cc0e589955081137ddfc6d988252da36554dba743300382d3498| [Download](https://update.gravwell.io/files/gravwell_offline_replication_installer_4.0.0.sh) |
| Load Balancer | The load balancer provides a Gravwell-specific HTTP load balancing solution for fronting multiple Gravwell webservers. It connects to the datastore in order to get a list of Gravwell webservers. |3c0dc7fd5d4022585ab73c771aa5fea76b6667bea3a5d6ebaf27a8a558122d27| [Download](https://update.gravwell.io/files/gravwell_loadbalancer_installer_4.0.0.sh) |
