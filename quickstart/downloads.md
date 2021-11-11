# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.6/installers/gravwell_4.2.6.sh) (SHA256: 181cd97d90f5d85e571c15002a8bdf0dd9e3cd4b2c073a20e0d09bd2ff550057)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.6/installers/gravwell_simple_relay_installer_4.2.6.sh) | ``717ab01dff2e5970ab2c5e4ab860492a47a75bcc753f4a8efc48dc66955e33d3`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.6/installers/gravwell_file_follow_installer_4.2.6.sh) | ``adec2877b01650ae5003d47297a65e3040e1605e0d8c70a9967e4be2639ede6f`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.6/installers/gravwell_http_ingester_installer_4.2.6.sh) | ``332ced7aa2ea1a12490682b86e11167a1bc86d752ec32819dc18aebbad74c50d`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.6/installers/gravwell_ipmi_installer_4.2.6.sh) | ``61891fc5914d96ce70ba950c62d422837e2029b7fbae1cdc15eec127a5018082`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.6/installers/gravwell_netflow_capture_installer_4.2.6.sh) | ``81cb91cacb1e03a095804b61104e8d5f0fa1095bfd606578a866cd7383c1aeef`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.6/installers/gravwell_network_capture_installer_4.2.6.sh) | ``b22aecd047b3aad3e96d3d12a824a849d1218a138ca4682c7c9b7d14e11bfc4e`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.6/installers/gravwell_collectd_installer_4.2.6.sh) | ``04dee5d73b58875083b0bbddb7598a1452ce1a60954e7bd97b945a69eacd86d1`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.6/installers/gravwell_federator_installer_4.2.6.sh) | ``f79a4bf5db70467b665375a3964b93aa19a163069f4cca2aca0d65e3de3b2c60`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.6/installers/gravwell_win_events_4.2.6.msi) | ``0ee34b034230caaabf94985e7f7ec6826ef5def2ffd92204d56fe4535ebb6a85`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.6/installers/gravwell_file_follow_4.2.6.msi) | ``8f14bfdd384ca422663919e8dc4cb78d19e0f42a92d078fbda5cf66103e06a5e`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.6/installers/gravwell_kafka_installer_4.2.6.sh) | ``0e9526d6e22d9b8e77c5c83797b2f1689c96073877e360c23d9ce271341aa757`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.6/installers/gravwell_kinesis_ingest_installer_4.2.6.sh) | ``38cc5dd73efd148e769d0c9cf4c7b4648fc0d628ffc6f65745a9ef952ddf3808`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.6/installers/gravwell_pubsub_ingest_installer_4.2.6.sh) | ``e13596a4684efa3f79e556c6dfb48003f410c3c52d3f23160efa8174ea6465fa`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.6/installers/gravwell_o365_installer_4.2.6.sh) | ``80fc6550bdeb788bae3bfa7e9955b04f69a42b7081fbc1effc844a343a580136`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.6/installers/gravwell_msgraph_installer_4.2.6.sh) | ``a32f87bf244f91d47f57f56de11923f09791cdaaac880652ffd9827e1b186e80`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.6/installers/gravwell_datastore_installer_4.2.6.sh) | ``89cae75413feea4f5d7f4949f5d7cac5930526d5863cf476bb4121b190065706`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.6/installers/gravwell_offline_replication_installer_4.2.6.sh) | ``a007f517cc7d25ff39ed85881181323585ba40feba4bf286116986daca204157`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.6/installers/gravwell_loadbalancer_installer_4.2.6.sh) | ``1d5437b36dccf87dc29de49f4274942c735e1815e0367711149980053e576a23`` | |
