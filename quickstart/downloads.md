# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.0/installers/gravwell_5.0.0.sh) (SHA256: bad49a90284054a5382e4ea7e7c707307a6b65d477b500d02139f950ed6c5ce2)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.0/installers/gravwell_simple_relay_installer_5.0.0.sh) | ``1d9cd07cee507a0907b339293525fa8970d6eab0b7de86283a379a388561a8d7`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.0/installers/gravwell_file_follow_installer_5.0.0.sh) | ``5b6a7489a089b558ccbec53b604363ff5c2b57aee73e6ad828ff987009f29cf5`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.0/installers/gravwell_http_ingester_installer_5.0.0.sh) | ``9aa5659308ab864e3d859129d4fa44a25113a86774636b436cd84026376093e9`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.0/installers/gravwell_ipmi_installer_5.0.0.sh) | ``e956952a0ee8c57c05f8f3383e59e99b33c4ad3cd1178f2b82f5c44314db346a`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.0/installers/gravwell_netflow_capture_installer_5.0.0.sh) | ``83b4969387544ddaaefd0b69b9ad1fcfb2bf5e8c8f34766684791a38930b718d`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.0/installers/gravwell_network_capture_installer_5.0.0.sh) | ``29abeefe3790952e945ddea84cb101e0298c8ba5e4dc030e0f5ea0150759ed0f`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.0/installers/gravwell_collectd_installer_5.0.0.sh) | ``ed1c46f722b5c74694f3ef5ca00a3ee57fb55f38160630f040aabdc1a2f5f755`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.0/installers/gravwell_federator_installer_5.0.0.sh) | ``6f17e9210f56719a89945821304aea999e6eeca37227d013ff7110fdc81efb53`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.0/installers/gravwell_win_events_5.0.0.msi) | ``5f928c6232199cd5ba6dcacef888991b193f69f94c6685c3e5cae91221bfec65`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.0/installers/gravwell_file_follow_5.0.0.msi) | ``89c3e5fd85d005b1e10ab253bd422dd47528264aec459f32774b6b31014a7cc8`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.0/installers/gravwell_kafka_installer_5.0.0.sh) | ``4826498ea0e77e7c263c1e4f189e22f96bfa104d61584618c8a1c9720d24e6b8`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.0/installers/gravwell_kinesis_ingest_installer_5.0.0.sh) | ``ecce79b796f8ce4522e6e77ad769f9af0b5ee15ddc5aec0289acaac9808f1448`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.0/installers/gravwell_pubsub_ingest_installer_5.0.0.sh) | ``ba648a81753c806effa7dbb724b08136ddf05e077c224cf4b3342943448d0c70`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.0/installers/gravwell_o365_installer_5.0.0.sh) | ``a56de3f176afac0c705ac90ef90c0f6e1228ccaab7ec043017e920d4d5e4bb02`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.0/installers/gravwell_msgraph_installer_5.0.0.sh) | ``6983556f7cf4a1c32858573b89b0fd9d956abd43987faa98dd5b3f45f3b14395`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.0/installers/gravwell_datastore_installer_5.0.0.sh) | ``074641c269a5c786b736708690d5b60eaa7bb3d72c76181f472ffcc6032bb3a5`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.0/installers/gravwell_offline_replication_installer_5.0.0.sh) | ``9978e278859e0aace8a5b76ce3a9a2f5153561b8528dea5501f79653161b3104`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.0/installers/gravwell_loadbalancer_installer_5.0.0.sh) | ``2126b842ce2af588771014b3d019abcd304fc5cead41ab89b661cd2b0f8e6e88`` | |
