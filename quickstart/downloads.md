# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.2/installers/gravwell_4.2.2.sh) (SHA256: 9afe707b4d43ab65810016d9737b9a0b0464c73164afb7bac8b9bf6443062ad3)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.2/installers/gravwell_simple_relay_installer_4.2.2.sh) | ``2a17d747bf84f6051cefb0c80566366fc64dbf76028ed3b87b6a377cca4a8cef`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.2/installers/gravwell_file_follow_installer_4.2.2.sh) | ``77d7c0cbbcfe536c475a683eee3d63ac48d7e9db8886c392c981c1c801e07bd6`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.2/installers/gravwell_http_ingester_installer_4.2.2.sh) | ``6b8667b97bf854d9411f8b4eef67c07d9679bad773bda36d7a8a08ddc403df01`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.2/installers/gravwell_ipmi_installer_4.2.2.sh) | ``baf8b00c45077de40f4e6ab72326744b4c77c597be4ff93db1c9dc792b1d9c76`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.2/installers/gravwell_netflow_capture_installer_4.2.2.sh) | ``62f57a5e9469b0c8d27579676e1f4d435c96df83c9506ea1447040addd19a413`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.2/installers/gravwell_network_capture_installer_4.2.2.sh) | ``72a8229ab6c6b64ed2ac39f910b94eeb6a8c9655ebe84655d76573a53bca1ded`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.2/installers/gravwell_collectd_installer_4.2.2.sh) | ``dd3a92cbe983f82762796ce05a4016c8a41f10a847b2ab8c32c0478cc2734128`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.2/installers/gravwell_federator_installer_4.2.2.sh) | ``0e3a0a76e48b901bc555405f78928b2475c44c479b3235dea120abe2bf443b17`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.2/installers/gravwell_win_events_4.2.2.msi) | ``fc6a4583b8e53367c6f97267e675299a34d97cad93e04cb537ba9ff630a6de46`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.2/installers/gravwell_file_follow_4.2.2.msi) | ``acb73efe09f1e499a7ad8ef21c6b5501967f499d8d253056070bd48550c06039`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.2/installers/gravwell_kafka_installer_4.2.2.sh) | ``2485ae26659ea3d9616c55200c520620e10aa7fa788fbc85e899b4236126058d`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.2/installers/gravwell_kinesis_ingest_installer_4.2.2.sh) | ``73964ad33f5615634aa30371f9989d388200a00dfb5be1b2c3098e4eb8c2ea29`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.2/installers/gravwell_pubsub_ingest_installer_4.2.2.sh) | ``7c5d86f90f3c63f8bb888c2131635ea25c45eb3292cf83c4d4b953329373b43d`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.2/installers/gravwell_o365_installer_4.2.2.sh) | ``0cf46595811cf15bdf12d25d08e3cbc8300aebbe740c8511d7d74a75c374bd08`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.2/installers/gravwell_msgraph_installer_4.2.2.sh) | ``724afcd06298efa73e8a2fd15479dd4d06b5550a2cd3e5557228257c6267ac00`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.2/installers/gravwell_datastore_installer_4.2.2.sh) | ``8ef11cc8a9872274a774ca3e853c72d634ce3fbf45f5ad8bf07aa6f617aca1e6`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.2/installers/gravwell_offline_replication_installer_4.2.2.sh) | ``f451b3abb9209552898093e39bb8753f0578ad0525df653ee688437d0e21abfe`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.2/installers/gravwell_loadbalancer_installer_4.2.2.sh) | ``54dc533b285b9847b71ce325781c43592c03aeeea88a22caead7af7392e9fb1a`` | |
