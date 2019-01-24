# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.0.0.sh) (MD5: 661405c4bd70f52fbbba9e12a1e95519)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.0.0 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |4d6db72665136653390e586bcf3aef50| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.0.0.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |0aab22ed7ea3917185b27402a01c4ef6| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.0.0.sh) |
| [HTTP POST Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP POST ingester allows for hosting a simple webserver that takes HTTP POST requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP POST ingester is perfectly suited to suport. |d64a03487f8230d7cb1944b0533d8850| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.0.0.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |4adb90118adc406f8daa7f76c8c88b3e| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.0.0.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |dc219a73c6e4d886a728860a57eca657| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.0.0.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |b5b8c8cc945df3c99a44f0494ead0a5c| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.0.0.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |c3cec1df1ee0f38542107ed040c48169| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.0.0.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |580e129ed3ee44077dc6723a4db829ad| [Download](https://update.gravwell.io/files/gravwell_win_events_3.0.0.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |56efb87fcfe15c7ab71a2775ac0c9fca| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.0.0.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |437d9be28bc3c126044d8199e44f5aae| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.0.0.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |a0b6e72250fea32652c8a5beacf0895c| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.0.0.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |fb1a09a8db3c8a5dc85128d798f55c16| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.0.0.sh) |
