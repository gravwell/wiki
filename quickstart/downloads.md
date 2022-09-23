# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.0/installers/gravwell_5.1.0.sh) (SHA256: 58ea90b544a446b63d2e4e5145eabeddd87bad52ba636055cac655dae7b1318f)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.0/installers/gravwell_simple_relay_installer_5.1.0.sh) | ``7c0de212b270b0e47542b4e51155298de6fbdbfa844f6b1a0aa51368785e0f0f`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.0/installers/gravwell_file_follow_installer_5.1.0.sh) | ``1dc36ea89c695d790aa0538260c041460173e000dc191776816d2e860469b1d0`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.0/installers/gravwell_http_ingester_installer_5.1.0.sh) | ``cb02704401152f6cd9a5e55155fca0b5e463e767166779d60898fbba1d1b1030`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.0/installers/gravwell_ipmi_installer_5.1.0.sh) | ``354694c1cbc37cd58591c4dac384b9a920ff37520aaf7f62f641e22f89a8760d`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.0/installers/gravwell_netflow_capture_installer_5.1.0.sh) | ``fb09cfdb0f0caffe8879bdeb7c460956bc588e0f53b67f3a8c44a0b2e093179b`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.0/installers/gravwell_network_capture_installer_5.1.0.sh) | ``fad0197df531fc7ea22c8216a76f423b0389d359962441587bc742089ffff752`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.0/installers/gravwell_collectd_installer_5.1.0.sh) | ``fc6ddfeaafc8cc89fa1149be5356d2bb7249162602f1f344d59085c0c0c47ed4`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.0/installers/gravwell_federator_installer_5.1.0.sh) | ``53490c00e2a673e47ab508c44f012d17486bcaa03615f67d793041390c7917c4`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.0/installers/gravwell_win_events_5.1.0.msi) | ``bf848e800de45b90766c67d2250d7324c9350bbbee1bcc85cbbf2e2b0d7dc13f`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.0/installers/gravwell_file_follow_5.1.0.msi) | ``7ee098529bbc912cf3129ae0fd839142b12f5d03be1d08c31385f3b08525bb26`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.0/installers/gravwell_kafka_installer_5.1.0.sh) | ``10c885eede336fcac3d0c377a052cfb8ed37dc32756e552de33d5091255f98f0`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.0/installers/gravwell_kinesis_ingest_installer_5.1.0.sh) | ``3b0cb659db331604fe5634531f40f91688e94077733b7a03ce8931de52ceb35f`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.0/installers/gravwell_pubsub_ingest_installer_5.1.0.sh) | ``c593904b7a6cc5bd8cb87e46669c496258cd3f1e847daddac3556ceb95f7702d`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.0/installers/gravwell_o365_installer_5.1.0.sh) | ``4f98897afa2f7dbe565fe9744936a46bf035236a6ceb4ddbc6c3408102b4dae1`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.0/installers/gravwell_msgraph_installer_5.1.0.sh) | ``a7a4caa9a8d1cf9b57a2f73caba59304ec47467adf6e4c2ee1a0491fdb1eed2d`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.0/installers/gravwell_datastore_installer_5.1.0.sh) | ``42c105ca3147262cd51a66fe5ce647eccca7dd000320fc2e6abedb87d789b0ac`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.0/installers/gravwell_offline_replication_installer_5.1.0.sh) | ``ad27f0b3a186ffaf33ff1b884180b8cb7602874111ff2d9c0e3041855afc9b04`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.0/installers/gravwell_loadbalancer_installer_5.1.0.sh) | ``5fce3ce11018ce7e280f333e6b7c0ac7deb1987063bb66e120ca664ae0b1454f`` | |
