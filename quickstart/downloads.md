# Downloads

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/ingesters).


## Version 2.0.5 Ingester Releases
| Ingester | Description | MD5 | More Info |
|:--------:|-------------|:---:|----------:|
| [Simple Relay](#!ingesters/ingesters.md#Simple_Relay) | An ingester capable of accepting syslog or line brokend data sent over the network. | 86b2450b98d3bef635dd98ed0fe85d0a | [Download](https://update.gravwell.io/files/gravwell_simple_relay_installer_2.0.5.tar.bz2)|
| [File Follower](#!ingesters/ingesters.md#File_Follower) | The standard file following ingester designed to look for line broken log entries in files.  Useful for ingesting logs from systems that can only log to files. | 76abfded4ad6689ac9cf0e7462076083 | [Download](https://update.gravwell.io/files/gravwell_file_follow_installer_2.0.5.tar.bz2) |
| [Network Capture](#!ingesters/ingesters.md#Network_Ingester) | The Network Catpure ingester is a passive network sniffing ingester which can bind to multiple network taps and send raw network traffic to Gravwell. | b02db73d82dc38d1bad736954d203c22 | [Download](https://update.gravwell.io/files/gravwell_network_capture_installer_2.0.5.tar.bz2) |
| [Ingest Federator](#!ingesters/ingesters.md#Federator_Ingester) | The Federator ingester is designed to aggregate multiple downstream ingesters and relay entries to upstream ingestion points.  The Federator is useful for crossing trust boundaries, aggregating entry flows, and insulating Gravwell indexers from potentially untrusted downstream entry generators. | 8cdc375f9e20f9fd114bc01d929ae8d3 | [Download](https://update.gravwell.io/files/gravwell_federator_installer_2.0.5.tar.bz2) |
| [Windows Events](#!ingesters/ingesters.md#Windows_Event_Service) | The Winevent uses the Windows events subsystem to acquire windows events and ship them to gravwell.  The Winevent ingester can be placed on a single Windows machine acting as a log collector, or on multiple endpoints. | 598aab27009d460d11352086104fd64d | [Download](https://update.gravwell.io/files/gravwell_win_events_2.0.5.msi) |
| [Windows File Follower](#!ingesters/ingesters.md#File_Follower) | The Windows file follower is identical to the File Follower ingester, but for Windows. | 38e2ffd9ca33680dcc475cb414a7ca24 | [Download](https://update.gravwell.io/files/gravwell_win_filefollow_2.0.5.msi) |
| [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester) | The Amazon Web Services Kinesis ingester can attach to the Kinesis stream and dramatically simplify logging a cloud deployment | 8a34e0c470728101a9e199655e761ec4 | [Download](https://update.gravwell.io/files/gravwell_kinesis_ingest_installer_2.0.5.tar.bz2)|
| [Google PubSub](#!ingesters/ingesters.md#GCP_PubSub) | The Google Cloud Platform PubSub Ingester can subscribe to exhausts on the GCP PubSub system, easing integration with GCP. | 5d1171c1920dbb6fa7c4a0524d0080ba | [Download](https://update.gravwell.io/files/gravwell_pubsub_ingest_installer_2.0.5.tar.bz2)|

[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)
[//]: <> (| [](#!ingesters/ingesters.md#) | | | [Download](https://update.gravwell.io/files/) |)

