# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.4/installers/gravwell_4.2.4.sh) (SHA256: a74afa2ea4196302f07fa8e3bb0f0c99ee20c3c13fe9f5a0692904d9153fef25)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.4/installers/gravwell_simple_relay_installer_4.2.4.sh) | ``a65e32ae9b8c6ae1e9095a99f6381ef8446145042c4136c8342998bfff66f5b0`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.4/installers/gravwell_file_follow_installer_4.2.4.sh) | ``40027d4a7096e1e3febed286841bc6226bbfeb1d7b7430ed0119ca23840209cf`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.4/installers/gravwell_http_ingester_installer_4.2.4.sh) | ``387509643b24f7b4dccfc21a31e823fd98e944284f02bd0114532ce1638e11ab`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.4/installers/gravwell_ipmi_installer_4.2.4.sh) | ``3ec842485835803bf9862aa55779a5181751a6e7d2a27444ea55508f156ee7b7`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.4/installers/gravwell_netflow_capture_installer_4.2.4.sh) | ``3416fb2275c770202230ef3f430b9d5a365d22bfee867aeda06aff91dcbaa1c6`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.4/installers/gravwell_network_capture_installer_4.2.4.sh) | ``1ccadeb91127966480cdfc8a4e578b4b0d4316026da23236cbeff6fd272c5017`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.4/installers/gravwell_collectd_installer_4.2.4.sh) | ``41fc45c35895b43cdc35e81ba5d6752dcb2bbfc0439d09781fb2864bb9fd3c34`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.4/installers/gravwell_federator_installer_4.2.4.sh) | ``119842929c143ab0d8ffae5338edc6dff27192ec0da66d5076b95427d4d5bda1`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.4/installers/gravwell_win_events_4.2.4.msi) | ``ada9d4348a7e3845d91d140390298e0ebd16f4e178f70d8b757e0b31284ffbc7`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.4/installers/gravwell_file_follow_4.2.4.msi) | ``5c47d356a99fb2f5c2d33d5e0f3d2f34b3a0bd96508c9b3b9455e87dcb7b9f28`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.4/installers/gravwell_kafka_installer_4.2.4.sh) | ``27e2206b5821047d990d1e631de382db0e5a244f808f47ffa3df97b59e2e7148`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.4/installers/gravwell_kinesis_ingest_installer_4.2.4.sh) | ``5a6251f4200c2ea3fcad90fc262383cf1ac3b4c1329aca803dd6e2560f070be0`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.4/installers/gravwell_pubsub_ingest_installer_4.2.4.sh) | ``f7634111e55c5e2bf1d3cc9a8df367effb887c8ef3100b816802d9c2c385a6a0`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.4/installers/gravwell_o365_installer_4.2.4.sh) | ``f8877d1ce5c81538449d415206ed9f1370b000555bf3ca59fd246462c5738ea3`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.4/installers/gravwell_msgraph_installer_4.2.4.sh) | ``cc7973159d2db340ee0df49041b3cb4b4fe07d732a877a90aa8a8949d0c10506`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.4/installers/gravwell_datastore_installer_4.2.4.sh) | ``406c2670f3c6d3f84b4bfda4f5418646acea935feb9fb48c6aa04eb55bbde5c4`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.4/installers/gravwell_offline_replication_installer_4.2.4.sh) | ``9e7b2ba8761f1fc47442751ec8256967061a1f793bfd6706591ac912e5bd43b1`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.4/installers/gravwell_loadbalancer_installer_4.2.4.sh) | ``e526893cf72f4519e78a14e344c54a72eb409e19e7c953f2c592feb6b68cfdb3`` | |
