# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.2/installers/gravwell_5.0.2.sh) (SHA256: 44e3ba4b319358bfd4e2b419f0f00987457303ac46b6cf436706b0193ec177a1)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.2/installers/gravwell_simple_relay_installer_5.0.2.sh) | ``f1a1e95f680317ac58fbe6536763147257296edb23a02a4e892d203379105b15`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.2/installers/gravwell_file_follow_installer_5.0.2.sh) | ``1d3205b1ca3aa3de3ff3208f7beb0bc25870b9774199a79c9f15e41ea0a81c6d`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.2/installers/gravwell_http_ingester_installer_5.0.2.sh) | ``71e9c4c68bfad5810bacc3ec1e7a20758b1d2a5d04c3d27bd3ec2232cd44e04a`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.2/installers/gravwell_ipmi_installer_5.0.2.sh) | ``f9bfed78130234665eebee7f28eb0beccbd844811cc8c4e0abd81cc58158a8fd`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.2/installers/gravwell_netflow_capture_installer_5.0.2.sh) | ``fceebaac0eb9c5e53a71e28243b0ebef20ab9e5b7ef4008bfc3c497496897314`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.2/installers/gravwell_network_capture_installer_5.0.2.sh) | ``44f4649ad3454880d6c23d10ce765ac37d250cc9f88fa345b215330b8675a73c`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.2/installers/gravwell_collectd_installer_5.0.2.sh) | ``1afacadfca7d9f382d51d86b022624dce6f399e7a8ae40c332028c905de00f8a`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.2/installers/gravwell_federator_installer_5.0.2.sh) | ``b26aaf1456563111d48e0c7fffc98314e8ae18aadecd9348515d7ced6857ee78`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.2/installers/gravwell_win_events_5.0.2.msi) | ``57564845d010dd1756f5b706f31176bf12961375d0ea1d2c430411b6e74db3d4`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.2/installers/gravwell_file_follow_5.0.2.msi) | ``691db21112893e223398935d0d27c31f6fc76805798aa99bb24da109c3b18272`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.2/installers/gravwell_kafka_installer_5.0.2.sh) | ``22b08e225e58450856d7606878fd1afc7d3eb3972cfc66940589b6a5f8e1c352`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.2/installers/gravwell_kinesis_ingest_installer_5.0.2.sh) | ``540da644da5b79062c834c9b92bb56abd04b354a313484a2d6f6b8771db856f9`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.2/installers/gravwell_pubsub_ingest_installer_5.0.2.sh) | ``84a560692db4ab591c6dd6a25ae85ca64f60bffa5bdd1a1e55bb7effffa27b20`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.2/installers/gravwell_o365_installer_5.0.2.sh) | ``2f0e3570e79f672226a2d45e7dcf7e3450a42cbe2c0e17671fb829caf4582ec7`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.2/installers/gravwell_msgraph_installer_5.0.2.sh) | ``1c21d52c4ed5371682632d0d52bbac997457666fe77820bf5f63c7679b5d08be`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.2/installers/gravwell_datastore_installer_5.0.2.sh) | ``cfbf16f99df8f50eebe6d5f02df0ecd29ef30fb3e927a3b0ed413bf8e32ff3ef`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.2/installers/gravwell_offline_replication_installer_5.0.2.sh) | ``17d36a3c559663262db2baf4e7fe21bd9a4d811df2c6e6edfad0bcac0d8916b9`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.2/installers/gravwell_loadbalancer_installer_5.0.2.sh) | ``1edc316a7635e75e702dc734b5387b43700a45c353418d15f22d1f22f555781d`` | |
