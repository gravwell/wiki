# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.3/installers/gravwell_5.1.3.sh) (SHA256: a15be7fe76bb784a1868f73b8dd5299fde2a3429dcd3f56d5d799bf3ab704c1e)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.3/installers/gravwell_simple_relay_installer_5.1.3.sh) | ``4da36015ad974147b1cc7e9b6e995cd1bd8894de625fa6131eac97797b2ce97d`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.3/installers/gravwell_file_follow_installer_5.1.3.sh) | ``af9e7196858d9a79c57fd108393923c6b09210c625e85e48f4c17893a638737d`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.3/installers/gravwell_http_ingester_installer_5.1.3.sh) | ``de5a4d77a31a7df249016411b9f292b94e9aaa3d73a154d495d48385eafe42cc`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.3/installers/gravwell_ipmi_installer_5.1.3.sh) | ``f74284f893759f8532d719f6a61a52274dc677d65aa614f0dff8fbec20852c16`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.3/installers/gravwell_netflow_capture_installer_5.1.3.sh) | ``e33591dbcc34da457b31461efadce90358c0550570909cf7eaaad60bc24c2520`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.3/installers/gravwell_network_capture_installer_5.1.3.sh) | ``f14ecf5fb4b6efeb0c6237c8dcb075885988a3bdae3af8c3c58f982ab7d3ad89`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.3/installers/gravwell_collectd_installer_5.1.3.sh) | ``b2bdb6a696ddd8c0b7103743a9d3b70c07525f4e0ca732e35cf106a90fa1f0e8`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.3/installers/gravwell_federator_installer_5.1.3.sh) | ``6602266a3890be1e14d263472a9ae9adf4c016f5b1fdba7ccbca566e4d5e2a7a`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.3/installers/gravwell_win_events_5.1.3.msi) | ``867b4c451b52c7eadd8957a5ec60a45e6b2d672f932832350eb8b3dec9bc7f72`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.3/installers/gravwell_file_follow_5.1.3.msi) | ``c66946721d1ee3108fbb49c202865471886f02ed5be2b165bbadd9597b3e31fd`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.3/installers/gravwell_kafka_installer_5.1.3.sh) | ``3d2bfaae2a2afa0d67c9044d73fbcee70df0ce445be1e23dee2ab569af67624b`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.3/installers/gravwell_kinesis_ingest_installer_5.1.3.sh) | ``1f8cc445b6faf5a9132738c94c49e0a81c3f81eaf132b3f2b3207b7df98ecd2f`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.3/installers/gravwell_pubsub_ingest_installer_5.1.3.sh) | ``7e7d0499168bbdd43f27955e956b8f0ce2c401d41a70a852b498808d07571882`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.3/installers/gravwell_o365_installer_5.1.3.sh) | ``fecabfed42599fc8a227a7113b7aab9440a303a271609e0b940dc9bc4293c7ad`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.3/installers/gravwell_msgraph_installer_5.1.3.sh) | ``f8970aedc4ad38d10e07a994d6491f7c04b56c2cc394e59ba2d57ef689dfe41e`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.3/installers/gravwell_datastore_installer_5.1.3.sh) | ``b9997f8eb61447ce4df57816093cb2755ea99e4f79135245fffb8f065a1408c6`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.3/installers/gravwell_offline_replication_installer_5.1.3.sh) | ``dc6d72204cc4403555ec3766dffc0be51b80341a79bed043b45fbdb0ff68f90d`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.3/installers/gravwell_loadbalancer_installer_5.1.3.sh) | ``3703243ab796b361c87cfbdca08b42e464229ce438659ee05eef83febdd98941`` | |
| [Gravwell Tools](https://update.gravwell.io/archive/5.1.3/installers/gravwell_tools_5.1.3.sh) | ``8b2c9e045a8dcadef22d067233f880ac4f0c7188c4b2ff2c20c880f15426a542`` | [Documentation](#!/tools/tools.md)|

