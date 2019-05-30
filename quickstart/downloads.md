# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/files/gravwell_3.1.1.sh) (MD5: 2233c505947dd81fa6f1ef172c1f18e5)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).

### Version 3.1.1 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. |bca565c499dca27aab8082f90bff6e88| [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_3.1.1.sh)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. |3b113b03ca16df62279075abc040b571| [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_3.1.1.sh) |
| [HTTP POST Ingester](#!ingesters/ingesters.md#HTTP_POST) | The HTTP POST ingester allows for hosting a simple webserver that takes HTTP POST requests in as events.  SOHO and IOT devices often support webhook functionality which the HTTP POST ingester is perfectly suited to suport. |efa216c3f184dce9be05abfe4e18cc3a| [Download](https://update.gravwell.io/files/gravwell_http_ingester_installer_3.1.1.sh) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. |77e19e28055b7a6f2c60af6a3d22216c| [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_3.1.1.sh) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. |51ede2ae93534e7fb0b9b9cae86bbee4| [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_3.1.1.sh) |
| [Collectd Collector](#!ingesters/ingesters.md#collectd) | The collectd ingester acts as a standalone collectd collector.  |8860a5268eb706d04cf5092c588b59ff| [Download](https://update.gravwell.io/files/gravwell_collectd_installer_3.1.1.sh) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. |38bc950d2a7e4142ca8aff05536ec6f6| [Download](https://update.gravwell.io/files/gravwell_federator_installer_3.1.1.sh) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. |7e60095bbab80169338bea21a6a9e9d6| [Download](https://update.gravwell.io/files/gravwell_win_events_3.0.2.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. |7d548ccc9cf479b727e5bc87c6b118ad| [Download](https://update.gravwell.io/files/gravwell_file_follow_3.0.2.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment |1222100230cb7d354d566deeca266f1b| [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_3.1.1.sh)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. |d84fa50b2bebdcc698558957edd2a4b9| [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_3.1.1.sh)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | Description | MD5 | More Info |
|:---------:|-------------|:---:|----------:|
| [Datastore](#!distributed/frontend.md) | The datastore keeps multiple Gravwell webservers in sync, enabling load balancing |9ad80ab965984680d091fd266e44b3ff| [Download](https://update.gravwell.io/files/gravwell_datastore_installer_3.1.1.sh) |
