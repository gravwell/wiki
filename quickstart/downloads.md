# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_5.1.2.sh) (SHA256: bff0ffa5645f7e84f8f9590389adb8f2dcd62b5b228455b1c50b909272da68e8)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.2/installers/gravwell_simple_relay_installer_5.1.2.sh) | ``8b89bfcc5267479fad4c13c8b3cd9c4e33b99964ade71d2fbe1924722722c4dc`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_installer_5.1.2.sh) | ``42d9ba7fcf7e84035ae5beba0c82569563a07b8cead915b44bbcc27a22f54da5`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_http_ingester_installer_5.1.2.sh) | ``68196f5b717130bb42a2fb51cf3dbc007b3719428b32f2f82102907a61164c66`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_ipmi_installer_5.1.2.sh) | ``cc1c407332b6dd941b3b2652c71e6875baa3ac33eeffb7dfe1dfaa5618b8022b`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.2/installers/gravwell_netflow_capture_installer_5.1.2.sh) | ``6c1d418fe283d5ee757ae5fcf520c5134acd7954ae279d84ea55eaf9d34e6580`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.2/installers/gravwell_network_capture_installer_5.1.2.sh) | ``e2607aafa40294a7daeab19ac99ebaf557dbc12d3f33bbded0ce6af37392cca6`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.2/installers/gravwell_collectd_installer_5.1.2.sh) | ``47db6fd8e8d33dd0b4ebbb8749f112aeaf85e7482d66b290258a7b3161edaac2`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_federator_installer_5.1.2.sh) | ``d1dc9ad73f4c162ec80d1d4fe91ed27eb240eb01533cd98ccb3a89cd35e02a28`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.2/installers/gravwell_win_events_5.1.2.msi) | ``99e736e33e25251d6e35d9a921c1f50c2ae6ae8cbacde87ba99a093e9eacdc3c`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_5.1.2.msi) | ``3b3a104e8c950eb93d38ca20b87a7ab8ce138fb1496d34a2559d46214402dff9`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kafka_installer_5.1.2.sh) | ``6695920af1cc581682182626b17963f765ef3c1cd0e4c81a0637b4d3108e7b78`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kinesis_ingest_installer_5.1.2.sh) | ``a072a1ad2900050588c0f9547d8dc21f8b4af1f417a0b803c15766953f48f0b8`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.2/installers/gravwell_pubsub_ingest_installer_5.1.2.sh) | ``aa0c68cac4f4af65cd055abb6a968c99068f61b8f28fb1cd4a8a9e5407bf2edc`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.2/installers/gravwell_o365_installer_5.1.2.sh) | ``b271c218d2f34e584959ba3cf597d98355d4083a33ea640645cbf56961a4bf9e`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.2/installers/gravwell_msgraph_installer_5.1.2.sh) | ``b26b3cdd41bfb6f1468f840aa13b615b14fbbe442d6d52fe745962127db38494`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.2/installers/gravwell_datastore_installer_5.1.2.sh) | ``be7f8444788b40650b8d0756ed75ca38e270798b68a199cc23401790e3585e69`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_offline_replication_installer_5.1.2.sh) | ``72bc5d316ca41e1cdc2d6d23c627220911d00f00a717e08321db0901070e5953`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_loadbalancer_installer_5.1.2.sh) | ``99571fc325abe46960bce86bb34919f7c2d727b2acf23fda1e1100e8154ce1f9`` | |
