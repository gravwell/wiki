# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.3/installers/gravwell_5.0.3.sh) (SHA256: 42482d65d9a573c8165847975c6ee3299e1cb9173ba334844d37a354c5f0c2ef)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.3/installers/gravwell_simple_relay_installer_5.0.3.sh) | ``288ae88b569e735b357c3e3819073a7446ee3cb2152c5fbddc29691b6a41ca14`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.3/installers/gravwell_file_follow_installer_5.0.3.sh) | ``257549399c27d7a8e7e51f627aed0f6181928ba45a7318d3a51d7edfaa4ea42d`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.3/installers/gravwell_http_ingester_installer_5.0.3.sh) | ``24025dcc1355f8cefdfa122ab6b692879d537359a6913bb64e329beb456bae08`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.3/installers/gravwell_ipmi_installer_5.0.3.sh) | ``a585a44234bf73035f8b62c79ac5c5975fb1b8fc8390b7a0877c38ed414a8772`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.3/installers/gravwell_netflow_capture_installer_5.0.3.sh) | ``3bd445ca60c9a50535ef840818a90ddc58bf97f0b40f7211a56c2c5cee39f172`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.3/installers/gravwell_network_capture_installer_5.0.3.sh) | ``9aa543341edf71d83094790b80e7d80313f7056752f2bc6f1033717bcbb11d96`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.3/installers/gravwell_collectd_installer_5.0.3.sh) | ``5dca82c0f62940ba45a51fdaa30a4f1fb7724a4c0cf766db3b5aedb2cc5f5350`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.3/installers/gravwell_federator_installer_5.0.3.sh) | ``ad416d2e09086abfb72e94f8056da1f8d06486af29066015a85a6d327e146f83`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.3/installers/gravwell_win_events_5.0.3.msi) | ``666228fdd41ea2b386dffdfa5ce9a970ec393bca4f7e9d49a9b192708c8a65e2`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.3/installers/gravwell_file_follow_5.0.3.msi) | ``de2de32115499e9d50770769f9b00382e001535d85139574dd9c574ab2342fb8`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.3/installers/gravwell_kafka_installer_5.0.3.sh) | ``f202afee29e1307ebc32c1c8ea576732c21412a023b40769a49e3465ddaf2a94`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.3/installers/gravwell_kinesis_ingest_installer_5.0.3.sh) | ``32657a90ac5ea74ed0a7f4318c5c1c59a658fb356695d408a32d8df485d03765`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.3/installers/gravwell_pubsub_ingest_installer_5.0.3.sh) | ``bd7e3429e3df8d12330b93c8e7b2b87b72619cd70fe56b54fcf38f4e59d9d35f`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.3/installers/gravwell_o365_installer_5.0.3.sh) | ``3a4e59963b5d65273dc48a179d4292116f026269aecd951d09bc0c3e7cd212c6`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.3/installers/gravwell_msgraph_installer_5.0.3.sh) | ``77f50cf0dcf84ff05818139f088399e72cb9217b5573916426c4517e1ef9ca99`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.3/installers/gravwell_datastore_installer_5.0.3.sh) | ``fa27a66313b56aec4e824f16c6a3c2ff9db9ad5a8fb3e659d70b3c0900aead93`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.3/installers/gravwell_offline_replication_installer_5.0.3.sh) | ``5f4a9838ca956228c06e0433a4d074d680341a0651a06e1f6df450ca7ad177db`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.3/installers/gravwell_loadbalancer_installer_5.0.3.sh) | ``73791b3025c86ea7057f89ebb89f9527cd56847c1d99e4f2f40a5d35a6bcf9ce`` | |
