# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.0/installers/gravwell_5.0.0.sh) (SHA256: 8ac1052a0f405d60a40cc37155cc50f4894076c3ec0e9227e9a777cfeae22297)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.0/installers/gravwell_simple_relay_installer_5.0.0.sh) | ``58e6558fa97c3db8019eb55d6232a11860f5633505fcebb1d243b02dbdf937b9`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.0/installers/gravwell_file_follow_installer_5.0.0.sh) | ``5b994d9c233fae3995f5dbf180b5ff3557c36f221655e18651a43fdc8e98f0dc`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.0/installers/gravwell_http_ingester_installer_5.0.0.sh) | ``4b329483702ed10a02c9a3a1c50bfce65c0a3980847a9658a93437f41faf844a`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.0/installers/gravwell_ipmi_installer_5.0.0.sh) | ``00be0bde51795dbfc8dc7b80b8213f0cf540dae45312017c7b4a9a907a5f3be6`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.0/installers/gravwell_netflow_capture_installer_5.0.0.sh) | ``68e6b1147e455ae6fcd53be534b9b6ff7aac9e9dd9829f8f9740d14147bc898b`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.0/installers/gravwell_network_capture_installer_5.0.0.sh) | ``824249df75acf188e34d831e620d2487c1dc5fc238be127778677f3d6da6ba09`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.0/installers/gravwell_collectd_installer_5.0.0.sh) | ``24028e32ad07cb8045d1bca611d2430345d3f5cc570f07a686af4c842362fdf7`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.0/installers/gravwell_federator_installer_5.0.0.sh) | ``c23049d86ef7c8e937d3123ae28ad492b2758dbd2f79aa18f3f0465b7d6061cf`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.0/installers/gravwell_win_events_5.0.0.msi) | ``3ba79e7369002d755f7d74af86411530c79668abd488a119585dc7de23f95363`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.0/installers/gravwell_file_follow_5.0.0.msi) | ``ddfe2e50baec045b21e70eb9306fd1f0e92ab610fb92c367683fd110412dda28`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.0/installers/gravwell_kafka_installer_5.0.0.sh) | ``6e1295baa672b5fa7e17a44c3b45edfaf220a636c475c1ca90cf64f7dad116f4`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.0/installers/gravwell_kinesis_ingest_installer_5.0.0.sh) | ``96163789e8cac19801cbd46eb38b237d1377ae918a2dad86cf782cc7066e4e1d`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.0/installers/gravwell_pubsub_ingest_installer_5.0.0.sh) | ``542002fc9ff0192913705e54e9cce79df64f73d374cd74badce5698457237898`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.0/installers/gravwell_o365_installer_5.0.0.sh) | ``a86a3906ebdbaad02c29770b9f05f18e16698084c7728bdbf4a5989fe37513b1`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.0/installers/gravwell_msgraph_installer_5.0.0.sh) | ``3fff2208ed48073098c57a3cbd440f94eb2f30e9a9d719b6186ef3a4441d8c6c`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.0/installers/gravwell_datastore_installer_5.0.0.sh) | ``e76daf679a888c9df8d8a7c336f0dae02859725a2374078a6330fa8f4709fe17`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.0/installers/gravwell_offline_replication_installer_5.0.0.sh) | ``620e1140304c5508bcd2696d89089df7067526f41edd11189da8414a7ffff9b1`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.0/installers/gravwell_loadbalancer_installer_5.0.0.sh) | ``ab0986bd27c90ac26f01ec1ddbc8ef68b4014adea328027d28cabd960040643b`` | |
