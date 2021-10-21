# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.5/installers/gravwell_4.2.5.sh) (SHA256: ce19966d194ddf0fbe69cd7b5a4d6b606641e052fe392b27de1914aa14d1220e)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.5/installers/gravwell_simple_relay_installer_4.2.5.sh) | ``bba544d1ea36995507f84ef920318dfa1a8d9c738ac5c2455e953a94785ef3b6`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.5/installers/gravwell_file_follow_installer_4.2.5.sh) | ``cc98c54d7c3027beeb5d67e8a05a5dd79b7c04b3ad6f95c271ba8a42c0de0c2f`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.5/installers/gravwell_http_ingester_installer_4.2.5.sh) | ``a06b9b81b98a1e7be283e1a0f70eac143a61a9c9e790442bf926910b5e0d36d0`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.5/installers/gravwell_ipmi_installer_4.2.5.sh) | ``0512135760151d2b42f7ff472e0a02723b3822b1ec81773af30783b346df0cc5`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.5/installers/gravwell_netflow_capture_installer_4.2.5.sh) | ``8dda34863480db2ef424684692b12b1a8b3f0860b524a25e335fde37067a4b02`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.5/installers/gravwell_network_capture_installer_4.2.5.sh) | ``41a863a78d94fb890df6f184d24ac05b6f2cad638ac99da7b34e3b9d9d42af63`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.5/installers/gravwell_collectd_installer_4.2.5.sh) | ``c95a996a6eadf784f7dce85b28a0391a405936d79e0e3fcc144c4dcaaa593f9a`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.5/installers/gravwell_federator_installer_4.2.5.sh) | ``4879436357b389e48303819f15be35973b6f63e823a725f51a223bb0b66d8cad`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.5/installers/gravwell_win_events_4.2.5.msi) | ``a4acb14730290a39bfe97034f9861a6138954190585523a5763401d5229738b2`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.5/installers/gravwell_file_follow_4.2.5.msi) | ``2609436520c450393267af77b8cf3775fe653c6c4870dac399e65ded3da84c83`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.5/installers/gravwell_kafka_installer_4.2.5.sh) | ``cd11b10f469c12edf98020783d65c8546dd6fa8874d3790a915755280edb3a09`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.5/installers/gravwell_kinesis_ingest_installer_4.2.5.sh) | ``b1a94162500945ebf686b9df84cd8e738e22a3419a542f86b660904a04788711`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.5/installers/gravwell_pubsub_ingest_installer_4.2.5.sh) | ``c7a25b8ba6d1fb5bff279e30924d0e59b9ed81c5440bf6128a88cb8a0d8d92f3`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.5/installers/gravwell_o365_installer_4.2.5.sh) | ``5093e850c8eba92c748c711539a0ac475fb4a45d36c325d21e3205fe66cbb572`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.5/installers/gravwell_msgraph_installer_4.2.5.sh) | ``ef29f55978798baf16c6cc8ffc8b3aef35798a05cc751c5c543522dfcb28e30a`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.5/installers/gravwell_datastore_installer_4.2.5.sh) | ``9d4bc47fd8352b4be16fedd67ad1fad9345dab64db6c10a55d645fb5b6382e00`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.5/installers/gravwell_offline_replication_installer_4.2.5.sh) | ``04a9830564cbf4c1251e12aa1801139412b271bfb5dd1afa14915f6c513db56e`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.5/installers/gravwell_loadbalancer_installer_4.2.5.sh) | ``92c83b61ad649fd41b0a7e1978bbafd9256f4901a1df6819e82e94eddcd00a22`` | |
