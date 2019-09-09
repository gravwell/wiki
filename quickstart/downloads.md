# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.2.1-10.sh) (MD5: 146c61265d8688763ef08877c3f7572f)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.2.1 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |b9ce589134451af1800cf776dca700e8| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.2.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |84040d0b052d877bdd63cbb2d2f594c1| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.2.1.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |b4d9caf724a10c2aba003202794f23f4| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.2.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |d8ffc8e9765bc397f9fded8c34084426| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.2.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |35a3c315dfe868e3c58d79368c6f4833| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.2.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f1ba837222083faee4a0c4813e0f9d8d| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.2.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |2fdea520a41b1e131ee2ad6a9765ceaa| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.2.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent ingester uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |548b3b1de5230b1f83606bf9b0eabc7f| [Download](https://update.gravwell.io/files/gravwell_win_events_3.2.1.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |f2322fd64d12b87973056c7c0cc6cdab| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.2.1.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |9d5698de3166ce24d822841d7950663f| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.2.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |9d2976c0e42accdb6b1be3045a14202a| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.2.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |fdf0e34592b945fbc4a7cbaf149dd7ef| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.2.1.sh) |
