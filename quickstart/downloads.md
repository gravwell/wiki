# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.10/installers/gravwell_4.2.10.sh) (SHA256: 4357a05c475250a760e439ccf9492a63c09ceaabe6f3b5c98f55c0de65ea013d)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.10/installers/gravwell_simple_relay_installer_4.2.10.sh) | ``f5bdf927e3c8b952f13593a19a7f54d13e3279e03685aa36447f7fa990ee500c`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.10/installers/gravwell_file_follow_installer_4.2.10.sh) | ``7113a3587f0640a81f45057a3ab925cbda79f9f43f8f07e645a46955f85ea441`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.10/installers/gravwell_http_ingester_installer_4.2.10.sh) | ``d0d5b1ccc1b0d2320b79d3a87b8eb49c34b8d3a5d48adb856e97e37212620f37`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.10/installers/gravwell_ipmi_installer_4.2.10.sh) | ``e8d2cf5c29f7fecbb11a0a2b5ab8a2c3415b2616e00183a380d29b0c48322147`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.10/installers/gravwell_netflow_capture_installer_4.2.10.sh) | ``68fd1bd649325fe4d06044baf8cc5a91892cc04e8f4b0e23144ae447e26ca77b`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.10/installers/gravwell_network_capture_installer_4.2.10.sh) | ``fa1d9882c33a9994e43f81a7d183604d5ae3b4cd7d778ef1f95e0ba5a87fdcd3`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.10/installers/gravwell_collectd_installer_4.2.10.sh) | ``ec76e8ed9b2e9b220a5891b70755a34a868b59a776bee6bd49ce46f2b423f697`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.10/installers/gravwell_federator_installer_4.2.10.sh) | ``7f2c92264541a607e0ff8c79b2b2858b3e4d1046282c966307a8cb171d5202af`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.10/installers/gravwell_win_events_4.2.10.msi) | ``bba4cf5e60ead3141fb1a22b39b78f25fb09ed9578bd253428b88f4a0175294a`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.10/installers/gravwell_file_follow_4.2.10.msi) | ``3ca76e0c56c5297776f018405b25135c045dc045834713fc38f0ffa1154d09d3`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.10/installers/gravwell_kafka_installer_4.2.10.sh) | ``fd3491ab58c0bdc1e148e6acbe1618a6d17d6f7c2c6335981073dfd1983277db`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.10/installers/gravwell_kinesis_ingest_installer_4.2.10.sh) | ``c3441ed6b34913917a653c1ab925f8827a44b9e6f66b09133926123cc6fbdc05`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.10/installers/gravwell_pubsub_ingest_installer_4.2.10.sh) | ``58cb82e354425dfff477e0a508054b22b2eb9f8abd1252dec841ef8f7a1b256e`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.10/installers/gravwell_o365_installer_4.2.10.sh) | ``7b46fda75d22fa9fa455aada050b49e56fce8cd676b07134f4fdd02f9e53a321`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.10/installers/gravwell_msgraph_installer_4.2.10.sh) | ``17a1868007a243f4e195b29f4b55b4fedc61b168cbfb2e003bcab9e537f92f19`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.10/installers/gravwell_datastore_installer_4.2.10.sh) | ``b4f05e79d9315ebbe30ad2eafe91f47f4f2aef5ec57c7c9335bff479e321f7ff`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.10/installers/gravwell_offline_replication_installer_4.2.10.sh) | ``620fdc0d1bd04d7f90bd128304e766088e20c9999619bca470a86811061ada24`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.10/installers/gravwell_loadbalancer_installer_4.2.10.sh) | ``b5093049eacec58db30218d6a00e29e51123ee11b7462f73998c0179756fda22`` | |
