# Downloads

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_2.2.5.sh) (MD5: 7fd044ad530ca0d24b42289716fee7e7)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 2.2.5 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |37aa5801a48ac8166de49cf58f1970c5| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.2.5.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |c066ddf5a873b6d4815fd0d967f0d07c| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.2.5.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |23776c4bf4344f73ffb8da6728425d57| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.2.5.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |c42ba6e9e5740c981aaf0362970ae613| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.2.5.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |e9bdebddb33ee5c813675a6878ea46a0| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_2.2.5.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |dc7f922616f3072fa0e48220e0be69a8| [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.2.5.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |f987f35c129b52c50be660a4a967dfa3| [Download](https://update.gravwell.io/files/gravwell_win_events_2.2.4.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |6c2c15c7c648a337bcddc2cd5702a7bb| [Download](https://update.gravwell.io/files/gravwell_file_follow_2.2.4.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |55fc2f1ad973de170a89a1007a804a95| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.2.5.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |71f94f12feabdc95440be702cf8b3785| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.2.5.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |6f66f4404beac8ff56e7f46fe3ec2e29| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_2.2.5.sh) |
