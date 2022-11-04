# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_5.1.2.sh) (SHA256: 170faade301ebd0c67d1af76073822e0179f72f1c3ae6a1538a4eb6bed0fbbfa)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.2/installers/gravwell_simple_relay_installer_5.1.2.sh) | ``61bb5606d3840cfe2705c2f377cf77490170903a1058523debfffadd5065bc35`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_installer_5.1.2.sh) | ``43f629bbfe17d03ca0024723ab16a804b011d1722b860ae32897408605639589`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_http_ingester_installer_5.1.2.sh) | ``33756d68895947f8eafca2801a5731dc9fa22608347ab66311fb04e53b941a0d`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_ipmi_installer_5.1.2.sh) | ``9bf925f9a560c5961f77a0e83e9f795a0676113e991682ed494dde3f48207ff1`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.2/installers/gravwell_netflow_capture_installer_5.1.2.sh) | ``751710d87db329f4fa8cb06b9f9f2f5241e777c5e151aa602d02b96c316322fd`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.2/installers/gravwell_network_capture_installer_5.1.2.sh) | ``2be1eaa2a3f4813e7d4d39ef0e477a06b5d593a994325b673a5631e19e1aaef5`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.2/installers/gravwell_collectd_installer_5.1.2.sh) | ``58ce9ef24a43d1e49bf0fbcc6fc2d43cfee672e82af6bdb764af3af9d8d1fc01`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_federator_installer_5.1.2.sh) | ``4cb4812b769b502ef0a113bfcbc9a6e3cdaba8f4057ac7bdab064221b5b77725`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.2/installers/gravwell_win_events_5.1.2.msi) | ``7a6d270884a3ded4e7ab5ad13d28b977404461908bec030fe3d77d73e84383e4`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_5.1.2.msi) | ``cf30c9a4c8f598b347994490b051a168ea4e8f90aaec7f32972780dcdf0b7ea7`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kafka_installer_5.1.2.sh) | ``b4b971c974b514919d1fb5f865dfc1d2a865f149da89368d68ee830f0f159de0`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kinesis_ingest_installer_5.1.2.sh) | ``5a888b2d147684f0e2d0d7984ab6a8ca70028e6f23d8807860ff70ad95343c15`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.2/installers/gravwell_pubsub_ingest_installer_5.1.2.sh) | ``764b4776f71a9d41fad04727ae867b181f3203e6288d3e8ef27afc39bfee8d99`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.2/installers/gravwell_o365_installer_5.1.2.sh) | ``5fbaf81c7b8cdb5add00c3890d701f43d41dc061023ea53f46914a987888ab05`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.2/installers/gravwell_msgraph_installer_5.1.2.sh) | ``6225d0871c1bf1dfe1adba87a689c317e1e69c0959cf2c914b7c12b8d7935e1b`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.2/installers/gravwell_datastore_installer_5.1.2.sh) | ``423bf8c2f5968408c3f91e523cc0a51d7d572e35788c3dd9b07551655e76ac27`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_offline_replication_installer_5.1.2.sh) | ``0d10b9d92806ff68be5595e3983e71b3f0f58e3abe143a4b445f3700b07207b8`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_loadbalancer_installer_5.1.2.sh) | ``f0a99cfd16e2948c255ca4a70a54033a7c8704901191e807c080346bb93a59a1`` | |
