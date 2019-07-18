# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.1.5.sh) (MD5: 93be92a15b33d812c182e7ab22d99253)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.1.5 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |63788ec0b6301fd3f3f3b00e6e20daea| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.1.5.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |fe21b34f8f423f948b414a7168c63bcc| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.1.5.sh) |
| [HTTP Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP ingester allows for hosting a simple webserver that takes HTTP requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP ingester is perfectly suited to support. |fd38bbcfd8093d75d6a07fb0e914fa78| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.1.5.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |d1da781baa286910fbfda4c3bb99428a| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.1.5.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |c191fd882f7e3d9cf54048e1c3b5613f| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.1.5.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |f2d44a5d6c269b1062a7de6044d42136| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.1.5.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |d802ad0b607e1767e5676713bde2cc84| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.1.5.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |49b0fc2e4da9145e289538691a7a6e71| [Download](https://update.gravwell.io/files/gravwell_win_events_3.1.5.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |491578e7a6be13a6bc9ed39754918469| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.1.5.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |078c2580f920aca5e94bf775bf2dd824| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.1.5.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |8ef94013c286ede11a9e2694fcf2262c| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.1.5.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |5994ff2c2d052d8a4b0fb5ac1c51c483| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.1.5.sh) |
