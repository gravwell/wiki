# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_4.2.9.sh) (SHA256: 8444c603bb7809b14ad6f71a4695981c920954c28878ae150557603accd23fbc)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.9/installers/gravwell_simple_relay_installer_4.2.9.sh) | ``025bbf149fdce36e802d97510ed2fec6c1b9019c59088c93c4ad9085c8c354be`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_installer_4.2.9.sh) | ``231112cdb6b81420df80de2ec7231973d990c7276c5ef0508bae1c0b6d592225`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_http_ingester_installer_4.2.9.sh) | ``cbd6ba99dd0d8e17e3c3bbc718d2da8b32f4e532a693b2a603bb36cf8fe78205`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_ipmi_installer_4.2.9.sh) | ``e72ae470f18b17d0f09300589f96067b7871962af28bae19b8f04fc65485ab00`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.9/installers/gravwell_netflow_capture_installer_4.2.9.sh) | ``0883a753fc21872f2be6d276d37fbf989e533a18fab7b1cd60e4b7a4d310e91f`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.9/installers/gravwell_network_capture_installer_4.2.9.sh) | ``5ded3c80cb6ea2c5c0520ac7c0805ccb02260477e2fb4be71d004d4794d5b509`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.9/installers/gravwell_collectd_installer_4.2.9.sh) | ``32b930e92362744ee2b7ad252c8e964754b2a4ba8aae9d2ae38077fba6a417ee`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_federator_installer_4.2.9.sh) | ``6f741af72f89a3b9886aba7ace9473b3ed2bdf4abc46c854661ec5308c1403d2`` | [Documentation](#!ingesters/federator.md) |
| [Windows Events](https://update.gravwell.io/archive/4.2.9/installers/gravwell_win_events_4.2.9.msi) | ``ec644aa905942cc6267ba79713a6d661202df61e614cf1f051446a69fc30a7ae`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_4.2.9.msi) | ``70fe962f168fae582463058d52301699c45c86a48b8e8219e92ac81a83e89ed2`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kafka_installer_4.2.9.sh) | ``b686dc32f8dee536d85fefb955877f0ab9bdc920e4eccb31a43eae6d389842f9`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kinesis_ingest_installer_4.2.9.sh) | ``9e69b02c63d1b27de40ee340a1d5df25e82bf0f44ef0c39da4b0b35db52cb99a`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.9/installers/gravwell_pubsub_ingest_installer_4.2.9.sh) | ``78ad39ecd4e40d1b3c548e1f35566b05eae5b824a1e830a0b7ba242f92fa5b6b`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.9/installers/gravwell_o365_installer_4.2.9.sh) | ``7f6cd052e4f5c19e3aa0d3b0575a6b49639ee175347d8115782878d462c880d8`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.9/installers/gravwell_msgraph_installer_4.2.9.sh) | ``8ccab2da92fa6b6cf0658416a707c70b4932596a0d76a8d0e92d494a89e1aa93`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.9/installers/gravwell_datastore_installer_4.2.9.sh) | ``167d3034421b84c14b4d777d546f848887effa5306f9bd292942e5c532e85035`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_offline_replication_installer_4.2.9.sh) | ``ff5f6ba42407218b952419ecbfb5a19c06f9f69ce2226788414f60c9d91e0078`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_loadbalancer_installer_4.2.9.sh) | ``11c3bb833235170a5ec9936e48ad1dea3db089f0edcf243c91b19a52ac038e6c`` | |
