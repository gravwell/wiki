# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_4.2.9.sh) (SHA256: 36db8f7ac2c5eb726669eaf532f991da2312ce527c0cbc9aef1f3fc2c86f077e)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.9/installers/gravwell_simple_relay_installer_4.2.9.sh) | ``514d3c0e3754c44b38b366901fcad3e4da92ad7b7372cb3c7fda3f57b851e632`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_installer_4.2.9.sh) | ``8b871a3bbef0b9b0d341b97d24201f03564500ab91c7aef234f26e46119e02ee`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_http_ingester_installer_4.2.9.sh) | ``461da6b1ff6b7aaf856eed3159df167ca1b7125190f8fdeab62664078146dd90`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_ipmi_installer_4.2.9.sh) | ``96c4978fbfc4a37d4862cf3208c2d6e0a30b421f29f9cd7554c1f2db643d61e7`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.9/installers/gravwell_netflow_capture_installer_4.2.9.sh) | ``970a60e641d172bb54cf8717d420a03e85597a19cc7f79ae39855bc848f93546`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.9/installers/gravwell_network_capture_installer_4.2.9.sh) | ``26b375e837e267f6737b020ec5229225ac540a808c752e934b69b12921df1b53`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.9/installers/gravwell_collectd_installer_4.2.9.sh) | ``bd78993c780f71f587f1c4069c78d202e701108e7424c97c57f0863fc58415c1`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_federator_installer_4.2.9.sh) | ``4871c4010c216aa077d7e3733fcb631bbcd98cc80a752d52d4b2a222f130800e`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.9/installers/gravwell_win_events_4.2.9.msi) | ``7e2f7bcd46cbcc3675730f45853a80ef42082845896699d5d0e4c0757afce91b`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_4.2.9.msi) | ``d0de790e198d87e4c0177692dc1c68b25ab8539ede977820bd2192b383c8ae80`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kafka_installer_4.2.9.sh) | ``19e29430b7af40f09dc5e90d0b9c1d5d2ba5f13727b87095a3e92b3e274a389c`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kinesis_ingest_installer_4.2.9.sh) | ``1aef0afa5d2becfa4ccd4b21e9c6059c2157f9100fa7a7c0a867e255e016de27`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.9/installers/gravwell_pubsub_ingest_installer_4.2.9.sh) | ``dfe68089c921b8aea53922a84a4e3c52f4de38c062bfd003fa3f4fe24311f0f6`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.9/installers/gravwell_o365_installer_4.2.9.sh) | ``ae5bd54a9ebcb1f331e0c22327422a213fe75a117dd5b67aacc48e699b36e9df`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.9/installers/gravwell_msgraph_installer_4.2.9.sh) | ``af8a927668c5efcb03453d9756475a8c0ddd86e2efe6e0388af66489016d71d1`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.9/installers/gravwell_datastore_installer_4.2.9.sh) | ``2777c91b9848b721e9b3c06d3481b7c4d6f05d321a58f95a5d2e1bbb7d7bfca4`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_offline_replication_installer_4.2.9.sh) | ``fc2fd8bb5215ba4a83dd677a96f968ee4fbd67462c2fd607b2f9b111b30ded22`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_loadbalancer_installer_4.2.9.sh) | ``98f19f96c2ba6f65678f1a4b779203d109f60e072a31f29fff54809eca6f73fe`` | |
