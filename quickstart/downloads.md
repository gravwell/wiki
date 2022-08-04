# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.5/installers/gravwell_5.0.5.sh) (SHA256: 1e4b56c6ecc9669212e3d773cff8361ae130bdc7d598b1c4ba406767007a1812)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.5/installers/gravwell_simple_relay_installer_5.0.5.sh) | ``c2faf3ae9d3467317db8fc95815a6daecb82e89065c5d0b37cc56d6e0553dc78`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.5/installers/gravwell_file_follow_installer_5.0.5.sh) | ``35a15002453fb4e139699937e38abb0bcb77720c7dee6a3df5ad23d334b830fc`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.5/installers/gravwell_http_ingester_installer_5.0.5.sh) | ``33234be738c35979bfaa37bddde196ab40262e5c5abe07e40404a4dabc007462`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.5/installers/gravwell_ipmi_installer_5.0.5.sh) | ``dc5b4172ebe9f46ce08b74afa156c1b362e905e15d66bd9cbf68a90863cbff96`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.5/installers/gravwell_netflow_capture_installer_5.0.5.sh) | ``6519d5a43f258e26e8936f62f476f8cf1191fe06f111c984e165bb6eb9e47e5e`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.5/installers/gravwell_network_capture_installer_5.0.5.sh) | ``6ef14337975d0f8ba7083e7b7b0c2b225f35bb44f1fd07135fcaa456d4f0ac97`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.5/installers/gravwell_collectd_installer_5.0.5.sh) | ``e2433774526a6f15f81b1f983ecd7f1586b56b41afd82e7cce9733a10969b5be`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.5/installers/gravwell_federator_installer_5.0.5.sh) | ``a7a79db0a9f5baedeefc89d07d94d3ffdbe83ac0c211fddb45327a2f3e361633`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.5/installers/gravwell_win_events_5.0.5.msi) | ``ad0973c7453877df908c22e8d54943bfeb03d3e0a9dfcc688d374cfb2be2ac3a`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.5/installers/gravwell_file_follow_5.0.5.msi) | ``06699ce2dca5390735c3ade1967800a988551162432f0205b8fc7afb36afc8b2`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.5/installers/gravwell_kafka_installer_5.0.5.sh) | ``85a9455f540ea7c55855a69f7e5cc5443142ac45d5364a9302d49401f1b220ce`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.5/installers/gravwell_kinesis_ingest_installer_5.0.5.sh) | ``860964327ca3499f24b6ac189103ee96b3dd48961d88c4acb82b8ae8477734c5`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.5/installers/gravwell_pubsub_ingest_installer_5.0.5.sh) | ``de9fbe4ef04d3bfd6f8709645e92a17451fa6b5eb8c2fb0cbc42fc6f55f55f3f`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.5/installers/gravwell_o365_installer_5.0.5.sh) | ``6c56de40922b738df08f5eb3320f5de234378ff1a102ab0347b19d4c1fa31155`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.5/installers/gravwell_msgraph_installer_5.0.5.sh) | ``91ddf6738ebb6885e9fe4a5c6e75d28fca785e30b8eb0f8eebd429ce51f328f1`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.5/installers/gravwell_datastore_installer_5.0.5.sh) | ``256f90f3636479ad9838ca1d2968b90eab476715c545e1a7100d9e046547395f`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.5/installers/gravwell_offline_replication_installer_5.0.5.sh) | ``b7c3d818cfe6398c45172ebb462d127f518f9ff4c2c0b52f1dfc5883486d374c`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.5/installers/gravwell_loadbalancer_installer_5.0.5.sh) | ``6a4339988b5ceee96d14877eb535807b718e6deb9ba322a0403cb3292dec38b9`` | |
