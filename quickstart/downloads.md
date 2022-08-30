# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.6/installers/gravwell_5.0.6.sh) (SHA256: 8d75cef55bc965a80381b2f684dc1d5679373060ba0ecf7000e009ae3e50eb71)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.6/installers/gravwell_simple_relay_installer_5.0.6.sh) | ``394130386d2a1fac26387bd1e50122c582fdb20d23a483405091f6cee6ec29a6`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.6/installers/gravwell_file_follow_installer_5.0.6.sh) | ``25ab8a69c380c47427634f3fd54c0f37e07b48a4c0661a87c25e6d41abf27778`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.6/installers/gravwell_http_ingester_installer_5.0.6.sh) | ``cbf1ae655ef583aacfb448a3bae1f9df506be4c3546a060f31a3ae2c7899c983`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.6/installers/gravwell_ipmi_installer_5.0.6.sh) | ``4aafb70f5b991af2ff2601d9dae165f9e9b99c1d3ae929673c9a57653b38ba24`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.6/installers/gravwell_netflow_capture_installer_5.0.6.sh) | ``eaa0b4bfdc8922419c38938c602eb0f74d9a2490ffd6e434ca61e0a86594727e`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.6/installers/gravwell_network_capture_installer_5.0.6.sh) | ``23d4f088493cc2b7e9e13c815ea9161bb34ac4399e8d9eb20e65a4aaf9fe95cf`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.6/installers/gravwell_collectd_installer_5.0.6.sh) | ``bc22b62010eef2c65d40539dd811bc7f35abc075849722f7480a23333afa10e1`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.6/installers/gravwell_federator_installer_5.0.6.sh) | ``61ec9797eee8dad4053d85791143319bbe24d8df0f0327639211ac57f332094e`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.6/installers/gravwell_win_events_5.0.6.msi) | ``9d024e59635491bb594a6e61126b9e77124cbde6b8ea5d9f039d829ae2889de8`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.6/installers/gravwell_file_follow_5.0.6.msi) | ``e52ee5e45c6935a3b266ba5cfe627cb2c473501c0cf071882dfb08d72fe1741e`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.6/installers/gravwell_kafka_installer_5.0.6.sh) | ``c6af0aa2737c0cf86a12c0e4b3a3580028e7f535cffa55e12e2da170fdc132c1`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.6/installers/gravwell_kinesis_ingest_installer_5.0.6.sh) | ``be68c6453732e5c856e30be8143afea363769695f66bb44f1b79af0276b02ac1`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.6/installers/gravwell_pubsub_ingest_installer_5.0.6.sh) | ``9bff98f55fbed92cf44110fadc4ecc029cc43b1ba492e37de67d96bea3a20fc5`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.6/installers/gravwell_o365_installer_5.0.6.sh) | ``5c705c90c7211d3494695bd8f1f0f53125c46e3ddf408cb17e4bbad9af0cc7cb`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.6/installers/gravwell_msgraph_installer_5.0.6.sh) | ``606cb158b5c07deef53f7e22b6b8eb5ef71236342590edb063038ca7c0218233`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.6/installers/gravwell_datastore_installer_5.0.6.sh) | ``39eb2b3346e5daad2ec6db899a7f8d97de4c77df03a365add3c1cab0f518cdd3`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.6/installers/gravwell_offline_replication_installer_5.0.6.sh) | ``3c7533f0da9c59789f514dc97636cfd75e22f6f9307fedcbf78e19ba49a73323`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.6/installers/gravwell_loadbalancer_installer_5.0.6.sh) | ``15f1c763289e8d059ee465da4409e9a3973145f74b371fed5f148d7a4729d3ce`` | |
