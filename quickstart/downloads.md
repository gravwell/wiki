# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_4.2.8.sh) (SHA256: 53d7797c6db205650c034574ff2c82d90810ea615ca4a4b5eb0bde87eb1e82f2)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.8/installers/gravwell_simple_relay_installer_4.2.8.sh) | ``c733022480fb2d14beb812af00f00427465715cd5b2a0ff483537f8d1e87adbb`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_installer_4.2.8.sh) | ``5a1ff2f4abc67f0e1417876a2254a645c9a5fda2c607ee315bf0267dfb970df4`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_http_ingester_installer_4.2.8.sh) | ``27938cf5b3517a3606e4c243c475d52d310cdaa179d9ca569ae4f313a8580617`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_ipmi_installer_4.2.8.sh) | ``8bf88c35b9fee7cc208d0f90ac242502f12ba5c9227d3776f9a7ee6b803da88f`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.8/installers/gravwell_netflow_capture_installer_4.2.8.sh) | ``5b10c81e324524f1b278e180ff3bd60ac838f1a7bc8d4f17ac1f5f30a8ceb3ee`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.8/installers/gravwell_network_capture_installer_4.2.8.sh) | ``0ea91f46614f2ff5c59f1fb5775e8dfb2951a0870943bf0200fa384e3ba5330c`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.8/installers/gravwell_collectd_installer_4.2.8.sh) | ``c2ef0c36c2232e50959081ae746c3ababa7e2d38dfad0ef2ebdbaf6163898916`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_federator_installer_4.2.8.sh) | ``1b067f3862a20500d25b1ee218f9ec6b4bb9c666df9187e018186116e25bf439`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.8/installers/gravwell_win_events_4.2.8.msi) | ``283069a20431894ad2f96e54d642cb75b93be83cd14a98e9617224af9b51e7b9`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_4.2.8.msi) | ``d13d3bb60724eea2e1b73860b7929ff11d80972ea103c01471c79b12147b718a`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kafka_installer_4.2.8.sh) | ``6718436e10fe3bcfcd5e1cf836c3ee95d9feb7c967033e8d43c23dea5f2f345d`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kinesis_ingest_installer_4.2.8.sh) | ``1f707c6146227316863f4c9445a113a4aed7ffe38352f513fbec604ff3f197b1`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.8/installers/gravwell_pubsub_ingest_installer_4.2.8.sh) | ``b5521a2d45c5721771902ff7d82118801d79ccf345e59507cbb2422e37771036`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.8/installers/gravwell_o365_installer_4.2.8.sh) | ``7bcc8c2deb358844a73413bb671ece5343d0aae8585179b86a9c7a663b48dcc8`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.8/installers/gravwell_msgraph_installer_4.2.8.sh) | ``e26bf191d0317d892b38d47dbf927daf2100aec3d5398c03a59b9bb90b3c99f8`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.8/installers/gravwell_datastore_installer_4.2.8.sh) | ``505b1b0c9a3b6e39231d0646c80818ddb60afa7ee962ed1b6262c52d506017ed`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_offline_replication_installer_4.2.8.sh) | ``c52a9ef4a322acc3c0e85b1d06299667ee3f40a635662a49911a8b34925cf5eb`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_loadbalancer_installer_4.2.8.sh) | ``c676e251d19cb2e2718699c66fb7bde79e6de89cd38c2c39ab3f258eb15608cb`` | |
