# Downloads

Attention: The debian repository is more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.8/installers/gravwell_4.1.8.sh) (SHA256: e31009ef3c016937b736d92bf25d6815dada4b45d74e1098d96a3ff1466d8ca6)

## Ingesters

The core suite of ingesters are available for download as an installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at [Github](https://github.com/gravwell/gravwell/tree/master/ingesters).

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.1.8/installers/gravwell_simple_relay_installer_4.1.8.sh) | ``aa35233b114e1838f7744f29827c7b08227304d92b1235058bc8ae658d813b17`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_installer_4.1.8.sh) | ``3bb8c0e53aed4b143e11fc9986f355bf6fbd1c925e92580c84e25f8f27fe8e19`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.1.8/installers/gravwell_http_ingester_installer_4.1.8.sh) | ``ea0e4ccff6f274532e22ac2d65b6c6c56112965b49f5a053b8f95078c6a358da`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.1.8/installers/gravwell_ipmi_installer_4.1.8.sh) | ``cd2e9d08da1d14e4b9e752df1ba6f6e3a8bb29378b7075eab34752a9bff64288`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.1.8/installers/gravwell_netflow_capture_installer_4.1.8.sh) | ``a5078e250c1fcbf4bde46697cf788278958ca49f5ef2692000e077b77924115a`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.1.8/installers/gravwell_network_capture_installer_4.1.8.sh) | ``6263eaa86a5b18075f2029fc4a819e3aa96ced016cf90d1708dccd8a864bdddf`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.1.8/installers/gravwell_collectd_installer_4.1.8.sh) | ``8943a74361bc36a0bde9720faaf1761a472a98016a8a336a17716f6c9510ddea`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.1.8/installers/gravwell_federator_installer_4.1.8.sh) | ``20dc4d5a98f1da3c851f36bd3e42a23bdde6d1735e730c459dd44459a283fcb4`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.1.8/installers/gravwell_win_events_4.1.8.msi) | ``eea6ac62cdcbe50e31a783ae6a3becb7e3cab97f842086a94b18de876c070e52`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.1.8/installers/gravwell_file_follow_4.1.8.msi) | ``c34aca1492ad9544140ff943c218116d4ddb735a853f91d508fcde873f33b455`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kafka_installer_4.1.8.sh) | ``a60c6a1ca15974eb4cbd1c5515dccf6cd33c301e0d7c319dedb20c23b480be09`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.1.8/installers/gravwell_kinesis_ingest_installer_4.1.8.sh) | ``a00373d067fed6c5d427408404bd9067394fc953f38900d66301a6ed8b83ee5f`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.1.8/installers/gravwell_pubsub_ingest_installer_4.1.8.sh) | ``b05e1159c3046f5526a50ce5183763d2577b7e115ee3109c9a26e8a0cd1cb484`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.1.8/installers/gravwell_o365_installer_4.1.8.sh) | ``2678b4d8ecc7b61ca975ca927e9683d3522957268d60f0d45b1c56e8313cf523`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.1.8/installers/gravwell_msgraph_installer_4.1.8.sh) | ``8c402031b931ee5ac6b3136aa978e81719e4daa5e049251c02258ac56cea6f56`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.1.8/installers/gravwell_datastore_installer_4.1.8.sh) | ``c4a9993134e3766107820fda16876658dd7ab750dc2af4f296980cbb0b0b5b04`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.1.8/installers/gravwell_offline_replication_installer_4.1.8.sh) | ``82b04e4469204e6e8109fc508b92e3ab40edbedcc4320617b9def254606cfe26`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.1.8/installers/gravwell_loadbalancer_installer_4.1.8.sh) | ``555af582521c535a8f0495b5622627723a7c82ab629961a0b2c4f3f07c9355ae`` | |
