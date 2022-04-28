# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.1/installers/gravwell_5.0.1.sh) (SHA256: 75184c50cedfdf8d6f809f10ada175fc8b842b95404eaef37c3f3d92dee87cc0)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.1/installers/gravwell_simple_relay_installer_5.0.1.sh) | ``ce79ff83a2bd636310a6a6f73c5c8f5c61ad0796871f344d17f42511876ec772`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.1/installers/gravwell_file_follow_installer_5.0.1.sh) | ``442a9f4241010bc541797aaf07c924f590432d8988c36a1d0ffd8f912dd9ee99`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.1/installers/gravwell_http_ingester_installer_5.0.1.sh) | ``2ae7f1ed990f25ca20e375bf8c289b80e0c2134cfbb7671ae747940314429e05`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.1/installers/gravwell_ipmi_installer_5.0.1.sh) | ``b0d464cec45af26d1aff432c13e396e625b9a4a3f295e9e0d2b17cd69f0ec3f6`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.1/installers/gravwell_netflow_capture_installer_5.0.1.sh) | ``bfeb43e69c6bdc4c11d8d8ccffbf6130c79cc8b7905005acead2c7e15e165f5c`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.1/installers/gravwell_network_capture_installer_5.0.1.sh) | ``9e6dce0cd581935cdfc5e38116f6e90c3788bcb6512b4c7f325001f24e34b9ba`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.1/installers/gravwell_collectd_installer_5.0.1.sh) | ``3e7efec40c9dd96d16354a2d68d3e947f154fd7127b8499e0e9e4a1dbcab9a94`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.1/installers/gravwell_federator_installer_5.0.1.sh) | ``97dadceacfe82c9e4692abf7e658fda1aa0d6d35ab364d599694be651070d125`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.1/installers/gravwell_win_events_5.0.1.msi) | ``0a099e06948e7d7131190de6f5216c77288cc854da383d89df2ad7f213679c82`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.1/installers/gravwell_file_follow_5.0.1.msi) | ``3475fe81d8b6bc039981db52d9e8f4dbcb317c8cb9d1e810a79fcedb1600fa6d`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.1/installers/gravwell_kafka_installer_5.0.1.sh) | ``f63d1a2a2edbf7b10cd55d3c263af0bc09a8b703529909aa03f82e4e51e48055`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.1/installers/gravwell_kinesis_ingest_installer_5.0.1.sh) | ``b9cddef566f999bfe0be596c124bd94a78aff5c99bacdb2dd52fb677f8b7a82f`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.1/installers/gravwell_pubsub_ingest_installer_5.0.1.sh) | ``feb35e374f4b2c48874ad4f37647360a5dbc510e4a2269f8f9720cf4164c15f6`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.1/installers/gravwell_o365_installer_5.0.1.sh) | ``9f29b413daa67b9126d79c8d330c51c7708cb6f1bb6d5fa77d041b74f4a3f92e`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.1/installers/gravwell_msgraph_installer_5.0.1.sh) | ``a569c0037fb519cc9e6ea70bf1e8a59b22936637603f11dcf27eab6c21d9a327`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.1/installers/gravwell_datastore_installer_5.0.1.sh) | ``39245f50755a005a315402f3608d88d8ce84fbab8db1dae094503d15ed64a2ec`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.1/installers/gravwell_offline_replication_installer_5.0.1.sh) | ``3130c47a10855ca7aff6e5f45fc0622601440bf078410b77d1a326422c385645`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.1/installers/gravwell_loadbalancer_installer_5.0.1.sh) | ``ba7f0e0500101292538f6b99982dc71a6517deaef9265ff755a8d3118052a887`` | |
