# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.1/installers/gravwell_5.1.1.sh) (SHA256: 30de40e4382b09190e388874103fd7161a7341a70a7182ac77666d0a7db59087)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.1/installers/gravwell_simple_relay_installer_5.1.1.sh) | ``6e897bcfecd826417d8582aad52ea91b862cce6c39bead6245c9bc7ded56484e`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.1/installers/gravwell_file_follow_installer_5.1.1.sh) | ``9899ab68c1f380356302c86b378fab8dcb8fa5b918ce3f4d0ee214c611c6a0b6`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.1/installers/gravwell_http_ingester_installer_5.1.1.sh) | ``d7f05042955607727bae5b6697a0e24125df32ee1130860a19f73fee463ad55d`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.1/installers/gravwell_ipmi_installer_5.1.1.sh) | ``9d888177f34bb5392c434e70ee03aef8436ced48be19aee893e426f72d9c8370`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.1/installers/gravwell_netflow_capture_installer_5.1.1.sh) | ``f911baa3c4fdae66ce68cb8c025a057cd8a667eac0be0f853e1fc5a45b877280`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.1/installers/gravwell_network_capture_installer_5.1.1.sh) | ``e690c26ea711a91759b106bc35055f54ffa81f23736ce89c3c01d9540e94531e`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.1/installers/gravwell_collectd_installer_5.1.1.sh) | ``beac4c198e8c2e02a0ef87fa80cf8693d48f5a0c785c9695d3effa29f4e87a8a`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.1/installers/gravwell_federator_installer_5.1.1.sh) | ``d9e77f93a8624fe2296c7b43725a64cd0633c254272ad0bf0e98f60df0338563`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.1/installers/gravwell_win_events_5.1.1.msi) | ``179205e5106de76486eac488115c91ddbe1a3600f26991cbf72173aac5a7de6a`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.1/installers/gravwell_file_follow_5.1.1.msi) | ``def9b9c7d2aa2c2386a36a2a48dd94ab47ce422a57f5d572725a5bca80046cf0`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.1/installers/gravwell_kafka_installer_5.1.1.sh) | ``254e083f76438262c550a74f264e4763b5d05783cd0086f06d5a2188c1a38d3a`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.1/installers/gravwell_kinesis_ingest_installer_5.1.1.sh) | ``4df2ef7821963bd75c9c6984fb5a1a1c5361a48c46471f4bbc23883414835d65`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.1/installers/gravwell_pubsub_ingest_installer_5.1.1.sh) | ``b87d8e7394af3678cab7dbf96de20ff506c8f3d74588224d94680de305cc1e5e`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.1/installers/gravwell_o365_installer_5.1.1.sh) | ``9ed06bd3a6808012d5a373c8267e37e64282456b3f3ec05d1310f8b2e731b5c3`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.1/installers/gravwell_msgraph_installer_5.1.1.sh) | ``b6582706f2bdceff05e50d25a615d733c8e14ff9d3110b428acbc199cb7f5725`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.1/installers/gravwell_datastore_installer_5.1.1.sh) | ``8cac97d7bf439c37752c3050cbdc474a269863bac4fb1e40afa010452dc393af`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.1/installers/gravwell_offline_replication_installer_5.1.1.sh) | ``af72740792fefc0072837aac500548d43475b880883056356390beaa5b737847`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.1/installers/gravwell_loadbalancer_installer_5.1.1.sh) | ``16d1a66b7c3a5bcfe1a11a828a5a3a60d5e8038210726e60e2e3c063ef25f088`` | |
