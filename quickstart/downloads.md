# Downloads

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).


## Version 2.0.5 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. | 671feb6d123e37ab606ca6b4baeafb00 | [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.0.6.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. | d1a9e601d881c98d0a9196598d963935 | [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.0.6.tar.bz2) |
| [Netflow Capture](#!ingesters/ingesters.md#Netflow_Ingester) | The Netflow Capture ingester acts as a Netflow v5 collector, ingesting Netflow records as Gravwell entries. | 9e23de9c315dffa341bbdf973c23c333 | [Download](http://update.gravwell.io/files/gravwell_netflow_capture_installer_2.0.6.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Capture ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. | 1dbf8982545ba0d70f7d903fc9ba73ef | [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.0.6.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. | b89d17df7331ed670c7f738ff94801e1 | [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.0.6.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. | 22eeb2dd812ead1e18aaeeb7e57460c2 | [Download](https://update.gravwell.io/files/gravwell_win_events_2.0.6.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. | 0f25d510c9e3464e576240e2bdc6a647 | [Download](https://update.gravwell.io/files/gravwell_win_filefollow_2.0.6.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment | f90a8716fa4aa3a5d53160cc96590fa7 | [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.0.6.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. | ac994107e4c80e975e903be0f9c37e40 | [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.0.6.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

