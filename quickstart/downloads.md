# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.0/installers/gravwell_4.2.0.sh) (SHA256: f7bb97ab2f00d3913c72c4b8a9577efe8803309a32d36fa43bc1faccdd4f2ab7)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.0/installers/gravwell_simple_relay_installer_4.2.0.sh) | ``1f55369e091937639c1b2b167e0bfff41757f83faed7313ae367ad3c5b4c4e88`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.0/installers/gravwell_file_follow_installer_4.2.0.sh) | ``cd8e37c8627d08a3edddbfb2e6e5d9aa7cace90f507400e964076aa077d3a067`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.0/installers/gravwell_http_ingester_installer_4.2.0.sh) | ``3dd2a3c8981520a5cce4108a2829890a07c9e75aca4a393718f006f32b60d03d`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.0/installers/gravwell_ipmi_installer_4.2.0.sh) | ``7a7349358a8a5c7dbdda1313fa81303f91ea0b6a8c3e35eaf642378611c3e113`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.0/installers/gravwell_netflow_capture_installer_4.2.0.sh) | ``88147eb36b8327b25b28e06dbfcf7cd38e0176ea99e88bbfefd69e739b879595`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.0/installers/gravwell_network_capture_installer_4.2.0.sh) | ``7116fe8fc347ce60d0a2bab74ebe5c110c23a6affe845fd83a5c28076801c659`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.0/installers/gravwell_collectd_installer_4.2.0.sh) | ``c832812a4220f7a5b8187246c5c881bb171ccbef56c42ad74d797b7ac311cac5`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.0/installers/gravwell_federator_installer_4.2.0.sh) | ``7629d309e31d18d21abb5627be3331b5142c127d7bc8c641dc61a92648a1ab37`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.0/installers/gravwell_win_events_4.2.0.msi) | ``be7c19c0b7b3e580ee5a91b136b134a48462d9c7109124d9b0bc8925db4b6fe0`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.0/installers/gravwell_file_follow_4.2.0.msi) | ``60ccf25c5bd15c4a49343052e76b82097858602d9adf023972303a7dbef87255`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.0/installers/gravwell_kafka_installer_4.2.0.sh) | ``c5af7dbb10f7c0db4c1eccab3471a64444be90cf3b82c29d0696ed9ef6493dff`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.0/installers/gravwell_kinesis_ingest_installer_4.2.0.sh) | ``8849ee5624a33e41e8c9e6d645722d8e797c9c2b4044385046c517812d4e92e8`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.0/installers/gravwell_pubsub_ingest_installer_4.2.0.sh) | ``d2376cdb7c32c67dc21465de618956514af2a4a4b2e090ee606d5b9dd7e12f5a`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.0/installers/gravwell_o365_installer_4.2.0.sh) | ``fff9bced9dda5c31fcdcc3279c149642ba8a0d98730a16e0a56229b542af31fc`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.0/installers/gravwell_msgraph_installer_4.2.0.sh) | ``5bfbf6924459c0a2024b99e6e5d8fa75665937d7689a1df905e8d05595cba184`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.0/installers/gravwell_datastore_installer_4.2.0.sh) | ``339e868a1023143e325ee2f07d121d271d679a05ca13fb6b8cfadc3ec5a41f94`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.0/installers/gravwell_offline_replication_installer_4.2.0.sh) | ``be402ae19caa9867f7f58630a2bb45ea992505acec5a888e64742aad52c256f4`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.0/installers/gravwell_loadbalancer_installer_4.2.0.sh) | ``8d3a819dc9ac0b5c72a01aac9b78931ff97af977cea7db22d1bb5252e333b574`` | |
