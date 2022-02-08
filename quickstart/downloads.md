# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_4.2.8.sh) (SHA256: 54ef8726117a9c82ecb0ce310b90766df86c19c15a7594199c5aa1b56dc9103f)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.8/installers/gravwell_simple_relay_installer_4.2.8.sh) | ``33fba507831503829dfb8bbb3ced740953001967fdf820ef980bf1625e209dab`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_installer_4.2.8.sh) | ``571116aae80d2b024022288de89c8fcd852205718a933c5ac809c428369eee52`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_http_ingester_installer_4.2.8.sh) | ``8930aa2ce74a598b8c9735bad98b1288095ef68414caf4e41153f4f4391d3535`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.8/installers/gravwell_ipmi_installer_4.2.8.sh) | ``a5c4726746b24cf822a70d7b447d020894e00caf636d2a226da53be5dfa76eb2`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.8/installers/gravwell_netflow_capture_installer_4.2.8.sh) | ``ac91dbe2d94f50d9c667b42d66f3e3177bc7c1c06f455cdf8645dbc4aa970758`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.8/installers/gravwell_network_capture_installer_4.2.8.sh) | ``33030de6386496131bf131b40bfb0de907e0351eebb2c2cb1cb917c5180e6fa7`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.8/installers/gravwell_collectd_installer_4.2.8.sh) | ``16ec045a0dea3b2689954fb3fde5584619ea0e751516af6d7de88b2218f45df4`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_federator_installer_4.2.8.sh) | ``c3bfbcd28731c6a69e8af0aacb313a5b8445a8d42b2e7ad49b4c4ed47a2b4128`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.8/installers/gravwell_win_events_4.2.8.msi) | ``8a44636579d748f354c7078b08e1f9b5e08928492eae7a32620ac785f0143f85`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.8/installers/gravwell_file_follow_4.2.8.msi) | ``0fea740797cb44768b1e8fb831466f7f3a945e4731467a9983e2c4245fe3a439`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kafka_installer_4.2.8.sh) | ``b444fdab5275f09e3a5435d0965a9dd811a69fae85adbb3e0dfac48b0c07d6fb`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.8/installers/gravwell_kinesis_ingest_installer_4.2.8.sh) | ``314a69d0ae13f92ac9f14dc674e9a352670b391560047e8b22def9b09ac6a7b1`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.8/installers/gravwell_pubsub_ingest_installer_4.2.8.sh) | ``a08c86d13b62d7ee068b8e44745d44fe7e58f98c15d3393a2e47f26ee7e4670d`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.8/installers/gravwell_o365_installer_4.2.8.sh) | ``2424e26b227b0a8b68905c5616b35c46dd442221f8bd7a944f4f3580e1081a8a`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.8/installers/gravwell_msgraph_installer_4.2.8.sh) | ``fc7ddc2fa140de555d8abb3bca287598d490ab474b82f0cf5ea946070e2cdda6`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.8/installers/gravwell_datastore_installer_4.2.8.sh) | ``cd5a129f477b41933f21f9ba9ace59b1d83b1d3a2a2319def5ff9a72d7f3269b`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.8/installers/gravwell_offline_replication_installer_4.2.8.sh) | ``8985459e31532b1f799dba49de73a826e739178dca47a0c0af79c55eb52b774d`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.8/installers/gravwell_loadbalancer_installer_4.2.8.sh) | ``25f015fedc42f4e10b1f2b4aba62c85c20ae5816d1d5fc0be283ee496d7a3a59`` | |
