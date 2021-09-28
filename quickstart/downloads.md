# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.3/installers/gravwell_4.2.3.sh) (SHA256: 008b5bb2c86626b6906bcb5ef603a2c62f4addf30d87bbe4842638de2e83acf4)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.3/installers/gravwell_simple_relay_installer_4.2.3.sh) | ``b0b4780a8cb8254dff8bee3c30e2cee837fcce07676e71d1f64f6efa4e9d8839`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.3/installers/gravwell_file_follow_installer_4.2.3.sh) | ``67f21b4dc13bd7587e6b98585522bd34d4190ec43109f7abf39492245d58a3ee`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.3/installers/gravwell_http_ingester_installer_4.2.3.sh) | ``9e8f88d64db4dd262779add3f41104f51fb5aa6b866fb4e669ab90483585e15c`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.3/installers/gravwell_ipmi_installer_4.2.3.sh) | ``c1c0befd91f6cd039dd800949cd8b0a64d5f26764fa9ff5e7a2356fee3cb4dc3`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.3/installers/gravwell_netflow_capture_installer_4.2.3.sh) | ``5f20ce11f015e6cbc85ebc5a4e4f839331bae89ea33d50594af36dc917f3d963`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.3/installers/gravwell_network_capture_installer_4.2.3.sh) | ``51ef52b3692882d62b4caae9e80c4af905f6f8b01e6ea7407f1764ede98aaff8`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.3/installers/gravwell_collectd_installer_4.2.3.sh) | ``e07e8a5738c436cce8c9266266f6e71adb989c7b13c510211938832894a806c7`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.3/installers/gravwell_federator_installer_4.2.3.sh) | ``608ac0cc797c7face1fe719659ec1164559079f2dc92d29b24ace4c592a214e8`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.3/installers/gravwell_win_events_4.2.3.msi) | ``3dd14c6b717c662e52f80bcbfbc76303ba23acf8494d6ff266b1f509baf1dcb6`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.3/installers/gravwell_file_follow_4.2.3.msi) | ``255e34a779f0bacc6e519996c4024b9824751de40eb318862636cb7a29923c23`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.3/installers/gravwell_kafka_installer_4.2.3.sh) | ``211ff2446b48fc6f3b21cd74526356caa5a2b82c215b98ef74fc5b71ae9870c8`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.3/installers/gravwell_kinesis_ingest_installer_4.2.3.sh) | ``ed538f66352b65ef65843efe31b7fed4837a36d9d6a1b21c4343cfecac69c8ee`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.3/installers/gravwell_pubsub_ingest_installer_4.2.3.sh) | ``f58c0699f4409a13969a66ba44f8b6e916aa58cd49458f872d26596af5c20ad0`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.3/installers/gravwell_o365_installer_4.2.3.sh) | ``e4287b07c723d1ac58b13c3a07f396a15876c1c4e79cc12977b357089a6bc99e`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.3/installers/gravwell_msgraph_installer_4.2.3.sh) | ``4898860d329cc58c3bc6b1b435ab8214ebc4d6382eb8f9a83a344d01bb58a8a5`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.3/installers/gravwell_datastore_installer_4.2.3.sh) | ``190c6414bc9aadd0792117016b75276ff2ee9944713bdacf576d221eab3163d4`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.3/installers/gravwell_offline_replication_installer_4.2.3.sh) | ``c4911f4299feb0b4115fb60a071417119956682531e2411fbd8a6e619492b1a8`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.3/installers/gravwell_loadbalancer_installer_4.2.3.sh) | ``4bb58fe5ed4988e8042f3eefbac0e0677977a613e53da4877b2b177a65d74aef`` | |
