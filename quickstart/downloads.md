# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.1.0.sh) (MD5: 6562286d971b89cfd2c20061ffbe6ccd)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.1.0 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |bc3668fb265b7de28cc04a8e25980974| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.1.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |54285c81de06720a40763b12200f84c3| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.1.0.sh) |
| [HTTP POST Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP POST ingester allows for hosting a simple webserver that takes HTTP POST requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP POST ingester is perfectly suited to suport. |33c082dac32e08d337373b0addba791a| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.1.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |7d3d567177951687715293a2af2a0ce7| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.1.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |dc148f854ddb8d2eb27d08415e075dd6| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.1.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |be1a9e5a168e996cfd5330fc9e7289ef| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.1.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |13fef5e7087e1d8e3262c59dcc780daf| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.1.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |7e60095bbab80169338bea21a6a9e9d6| [Download](https://update.gravwell.io/files/gravwell_win_events_3.0.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |7d548ccc9cf479b727e5bc87c6b118ad| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.0.2.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |ac8b4709daaf8d7a5cffd5ca41d6c1e7| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.1.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |f39ac03dd8acc2bb8b8c448a3cae9671| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.1.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |435a8fb01926560a782da5f5a90feba6| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.1.0.sh) |
