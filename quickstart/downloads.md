# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.4/installers/gravwell_4.2.4.sh) (SHA256: 073913ed531e6e8195f96ee5d801b288c28d6d7e7687ba9444ecddb7f9aaeba7)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.4/installers/gravwell_simple_relay_installer_4.2.4.sh) | ``27cfb87da48e6491928d45171947a435f866476b49869e6b5d45dc63cd5dc27a`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.4/installers/gravwell_file_follow_installer_4.2.4.sh) | ``328d8a1e2d081f96d571b0573b603e5deb508e266fb438adff90c5a1e381e120`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.4/installers/gravwell_http_ingester_installer_4.2.4.sh) | ``51549d7c803f0c871985c363eb9fdde56723e1b68a6fefb46d07f98a74731c57`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.4/installers/gravwell_ipmi_installer_4.2.4.sh) | ``88c8ab940efadf7319f5e7a258d8bf48eca0b6312cfe7cd5a87973fd13b504dc`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.4/installers/gravwell_netflow_capture_installer_4.2.4.sh) | ``0f72d2b0c18943ad5967e32362f4ad4b380ed6e61e264d72a59d40a272521bf2`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.4/installers/gravwell_network_capture_installer_4.2.4.sh) | ``e5c36c689f6fa43975c3ddc6fd7c2e6e58458b69fbc9234fd8d5a7471d422a82`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.4/installers/gravwell_collectd_installer_4.2.4.sh) | ``aa038da57bb4c0fae2cf4514650d692fbdbcdffdcfeb703aa394fbca9e85d174`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.4/installers/gravwell_federator_installer_4.2.4.sh) | ``e304296175d3cf3ea80b01d1e32865ce4ca8b3706c363568e82c727e9508c190`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.4/installers/gravwell_win_events_4.2.4.msi) | ``bf2435a938656b68a396d5feca8b79e8488c7dc14e75e0eea6e83a7e6e2dd7ca`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.4/installers/gravwell_file_follow_4.2.4.msi) | ``e71c10c71d0433bae1a3126e1da7b0a21b365f65877c0a9ef2f94e0e813c3cb8`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.4/installers/gravwell_kafka_installer_4.2.4.sh) | ``a50620b222bd36e81a2389d726dfb407363988af39790cf36d3d7361a5d8f39e`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.4/installers/gravwell_kinesis_ingest_installer_4.2.4.sh) | ``31843526fbf2259951a9a947966656d49e7349033dc3b9d6730ff8de2a5d8d9e`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.4/installers/gravwell_pubsub_ingest_installer_4.2.4.sh) | ``82b9c8543021d41f9291454a1d545cb2da67679fc1850ab3dca8cb4641232a88`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.4/installers/gravwell_o365_installer_4.2.4.sh) | ``e24a3e630e849ee6e0ed5a47ab61612d0a1973e43cf4f926b48bba846c59067e`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.4/installers/gravwell_msgraph_installer_4.2.4.sh) | ``4cca1f44976084fa85f12c6b6de0904bb5231973712141ed297f8acfb8b5a399`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.4/installers/gravwell_datastore_installer_4.2.4.sh) | ``01dc900723b91a9d136cfe70b530e8dcec4554b95c7e0b535df1b97117f3962d`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.4/installers/gravwell_offline_replication_installer_4.2.4.sh) | ``8d76dfd4af9a4b7013ebe029232d4c62e9b70d6820a150a47bcc8ec4ee506683`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.4/installers/gravwell_loadbalancer_installer_4.2.4.sh) | ``1175af6d43b095d1b18c8ab9ff9f53b4e810ecee483f1a68ca78a7192d67bf9f`` | |
