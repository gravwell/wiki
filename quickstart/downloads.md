# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.1/installers/gravwell_4.2.1.sh) (SHA256: 57aca710dd18f09c12056f9942ad8ace0baa0d5cd61137ceb2cf6a5725e455d8)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.1/installers/gravwell_simple_relay_installer_4.2.1.sh) | ``4842ee96eef6862f0b17bda1e0ea7c1d99f7e235caed00893840adc26a0d9c6f`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.1/installers/gravwell_file_follow_installer_4.2.1.sh) | ``992c4ca3047f6e56dd5d62e41a9ef47202a492190c2eb8a102ca86a3fe119a21`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.1/installers/gravwell_http_ingester_installer_4.2.1.sh) | ``dbf94b122003c848f8669cf77f780491b9ac7ccd0e2a9bd50362ead376f5d462`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.1/installers/gravwell_ipmi_installer_4.2.1.sh) | ``81f9d003186618c1e943cda019f68e9035cdb9d4b134c1e4a6433bee8432fac4`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.1/installers/gravwell_netflow_capture_installer_4.2.1.sh) | ``ee7b8ada55f27c15857f60c339ff5d015212fa0851edc7e76db8b3b9a40654bc`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.1/installers/gravwell_network_capture_installer_4.2.1.sh) | ``78ae94460c4d0cc1b8ee91e51ad664f9bc77ace3cea5ba5a79f5b5663bef1b48`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.1/installers/gravwell_collectd_installer_4.2.1.sh) | ``95712d0ea7606d3b5cd160954865cb855aaf3fb44218b773e7b3aa793087e0b3`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.1/installers/gravwell_federator_installer_4.2.1.sh) | ``2bf1434efb6a25727b91f0725ac87c48cb9e1305edfeb6a8a4cb851914d4753f`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.1/installers/gravwell_win_events_4.2.1.msi) | ``88e6835f85cbcb1c4531c21965b5d095a28a831c0757205da143cc3e5b12f26b`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.1/installers/gravwell_file_follow_4.2.1.msi) | ``cb4bc86b7492656c7bad0dbe3f810bbe7c18f01cc118f93ecea6c7528a676484`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.1/installers/gravwell_kafka_installer_4.2.1.sh) | ``2d338e8957e11980bea95ff97a2252f77f22fa7f623c61775cb4dec2cba20730`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.1/installers/gravwell_kinesis_ingest_installer_4.2.1.sh) | ``92877f6fc7cfd222fa353ebf0efda02bd44585fb75df9117837605fde08cd7c1`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.1/installers/gravwell_pubsub_ingest_installer_4.2.1.sh) | ``393d4854fa545a80c11b5a87ae2047a5628cfc62d93fa57d04d915ccfd7f6b43`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.1/installers/gravwell_o365_installer_4.2.1.sh) | ``df31d36d5df2eb36f2602448420d7ff1a377934c8ff0a09d7daca45de1760da9`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.1/installers/gravwell_msgraph_installer_4.2.1.sh) | ``07c58cfa89137a8576dd23798b07ec493c920ddda00619fb95dd75da71f14ed2`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.1/installers/gravwell_datastore_installer_4.2.1.sh) | ``621a3df56a67029a91dd0c960eb988ae136101cfaa52310437ce4da84d4b7570`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.1/installers/gravwell_offline_replication_installer_4.2.1.sh) | ``6ef514f577ad7f44b9ff18061216c678cbddcd2ec01db934f22e8d55fa4caaa4`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.1/installers/gravwell_loadbalancer_installer_4.2.1.sh) | ``9bfa6efb4725588063dfe30544eae68bf48c4723b9e5b3f1a4a6815d95c29c01`` | |
