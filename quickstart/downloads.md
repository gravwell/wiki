# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_5.1.2.sh) (SHA256: 324560dc3efc3dcc86df75c24470bdbc80c460feef260a8ba5839fc27225928a)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.2/installers/gravwell_simple_relay_installer_5.1.2.sh) | ``473198e497557b292f0fab593975e5fc054ed2173580f2b1ae841eb8c36dd64a`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_installer_5.1.2.sh) | ``5f348a9d8384a5a55ebef771df79c533215c71413e6e3fed3d502d83311815bd`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_http_ingester_installer_5.1.2.sh) | ``107f951b5a8eb2e112270548e92a00b516f78f093bcf54fdd7fcd266e06314a8`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_ipmi_installer_5.1.2.sh) | ``c92ed3166d4f5458caa2fcacb2e7a3073904a4edac7fb15aa664499859a0dbdd`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.2/installers/gravwell_netflow_capture_installer_5.1.2.sh) | ``3489428e4a10631eb75f1980336fae7dc4d62db49350aa94f2ab88fdd890a84c`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.2/installers/gravwell_network_capture_installer_5.1.2.sh) | ``12ff669758f02a6181053bdfcc2ee7aee46990506fa4e3456467e44cbc4594bc`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.2/installers/gravwell_collectd_installer_5.1.2.sh) | ``974134601f02f65ab0f4853b49f98f69007ca2ecc6b03d44a2e3c70414d5985a`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_federator_installer_5.1.2.sh) | ``a3d32ac2616238f40f44a5771f9850e88f40f7f2cf36d047558e9a207eafd98f`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.2/installers/gravwell_win_events_5.1.2.msi) | ``5a4ab1b3b6e68c306cbf8142099bcf482f0803daec77a723e99a6c3d1e4ba12f`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_5.1.2.msi) | ``c15f47e8f360d2b897cdd6902d64190df006520b9b7b287dd155e334e7df0fc7`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kafka_installer_5.1.2.sh) | ``c23206d4067bc6727b3d6b611e03321f068a6bcc0236a684c10a9e87d96f216b`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kinesis_ingest_installer_5.1.2.sh) | ``cf799cc1525d6fb2aab09294c93ce5585f8b72a5c389c84b9d42eba42e548dcd`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.2/installers/gravwell_pubsub_ingest_installer_5.1.2.sh) | ``cda30134b132e39566a510fefbf262b5d872a6e587dae9a41eabe5568fabfb04`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.2/installers/gravwell_o365_installer_5.1.2.sh) | ``eb649de9c05b04b99ee112e1e9640b713c5eea55d704d8c1e5e796dad3f841c3`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.2/installers/gravwell_msgraph_installer_5.1.2.sh) | ``80597f7fdd514ec6f5bc3211168bf309e873eb0a2f586bb1d5824c9628a6db50`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.2/installers/gravwell_datastore_installer_5.1.2.sh) | ``9f7a832a14a204a5c9564219ec57696915b578db38274be68f16b1962c092660`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_offline_replication_installer_5.1.2.sh) | ``250c036e1c2e9a9777ff22f3d34ae4d79e4ced9c31dbd98bf884ac03aeb7d856`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_loadbalancer_installer_5.1.2.sh) | ``3c9769549cf6909b353f0326f3abd2ee5057b7c1afed260621ce87e4761c5f2b`` | |
