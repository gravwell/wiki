# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_4.2.8.sh) (SHA256: 1a3338cc664c0dac926b415ecf4cfc7c1e8281945ff7ed2320467c41968d5750)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.8/installers/gravwell_simple_relay_installer_4.2.8.sh) | ``befa815047f1fa70c086854fcdaeea2811259e1a688bba23bc981cd2cbad224d`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_installer_4.2.8.sh) | ``5097648dd63b58e56e0cee3b6bb40a64d5780a2c987fe8e002844da29490994c`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_http_ingester_installer_4.2.8.sh) | ``e5452624528f28971de877cb36728fcac6748f31b38d7b7496d49884c8bd1357`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_ipmi_installer_4.2.8.sh) | ``0d94e010783345523ff42221467a3a234f09dabc5acad6eb9d2a931e1d2b2bf7`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.8/installers/gravwell_netflow_capture_installer_4.2.8.sh) | ``6ec462226fe7f5b60a165da98c1debac82b1b886e5d0ed9c00982392075477ea`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.8/installers/gravwell_network_capture_installer_4.2.8.sh) | ``0b91f9ca0b82e003308835e90584d48bb9366d5146b30b620d5349262ae0e03f`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.8/installers/gravwell_collectd_installer_4.2.8.sh) | ``c6f8b89f03f46acaffac227d9db1e19da132473bedb1d0b14ab340487ed26ea6`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_federator_installer_4.2.8.sh) | ``fe383c2e0a2529d77d6ba1bd043c22128cb061f709ab81b4612fb42f16974581`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.8/installers/gravwell_win_events_4.2.8.msi) | ``1a795884390e89c742092ea8642527746e738059aa8bcc8888eae0860a8979a3`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_4.2.8.msi) | ``ba4e980762ea4701c067ab3ca28b72d5a08cc649bee5b13afbe3c48678de8004`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kafka_installer_4.2.8.sh) | ``ba64f0f32aecc52e7a59aa4755ee9f13ff6118657fce69eb02c9732d4cc2e38b`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kinesis_ingest_installer_4.2.8.sh) | ``58e4e85bd69c4642397df3e4a3d548c98643e234de328fc845205bb6a819fa1d`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.8/installers/gravwell_pubsub_ingest_installer_4.2.8.sh) | ``7ee53543044c63ae5413b91a3b04da1f49c08b2e7ffbe65a696218983163bf9d`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.8/installers/gravwell_o365_installer_4.2.8.sh) | ``20a47fab01763df6deac6ab95a1b6e293725f600fa1dabe8043125c1fc30e68a`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.8/installers/gravwell_msgraph_installer_4.2.8.sh) | ``78d300bfd7d641127f2abd056de798094d77073a5a74005e15ce4c679f55a52d`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.8/installers/gravwell_datastore_installer_4.2.8.sh) | ``d3b2d0664affe74434183422c3f8ef0ba5d8e61521f91e295bf8d9f0475dcf24`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_offline_replication_installer_4.2.8.sh) | ``ff91a02e05631859badb3290bab89e35cc33964a4f4b8ca659f42d8bf63ea3d2`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_loadbalancer_installer_4.2.8.sh) | ``04b07415948f2db4156a781c20a4432e8b6493f3965ba04f04daba881b99a47a`` | |
