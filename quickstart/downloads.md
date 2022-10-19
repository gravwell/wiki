# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.1/installers/gravwell_5.1.1.sh) (SHA256: c5d979fae3808895ec025ddd3e9466ab8e193916d15322de20088a449bcb8e45)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.1/installers/gravwell_simple_relay_installer_5.1.1.sh) | ``71254f02a840e1565dec971f372fed15f3850e396998ff9039452d61fe65443c`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.1/installers/gravwell_file_follow_installer_5.1.1.sh) | ``111a4abfe46518087ca0dbf2f9288f393e4067efab7993848e6d72c22c16749f`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.1/installers/gravwell_http_ingester_installer_5.1.1.sh) | ``3c7f22e61dc9cccfdc2595811c6b023e10c3a4892f6876472bb886d98e7ee75f`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.1/installers/gravwell_ipmi_installer_5.1.1.sh) | ``229ae18664e19a282c25ec069070933573903227f4f4c8fa6640983f065e1f0f`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.1/installers/gravwell_netflow_capture_installer_5.1.1.sh) | ``9c701889fa4fc2504dfc79b1a3e35f312f6429c2aa51823d8a8b74f5f9fc6702`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.1/installers/gravwell_network_capture_installer_5.1.1.sh) | ``92e87e0f86ba6c3dd3d0553787696c0bc09d46ab59c63eeb49d95562ecd70f19`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.1/installers/gravwell_collectd_installer_5.1.1.sh) | ``6983e91ebbf0f3ee39b763def5c21116ed3f5630f47e6d86a0aad1c142fe4581`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.1/installers/gravwell_federator_installer_5.1.1.sh) | ``4c1f88a65e6d6dbe0fa748bc5d9ca50a3eff2e7f78ec735757708817fbe576c4`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.1/installers/gravwell_win_events_5.1.1.msi) | ``e8ba1896721844578cc1b1fa310c9be410239a9cda8a0ec102d7c341654b5fa5`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.1/installers/gravwell_file_follow_5.1.1.msi) | ``15c387fe5649e5904a2b259b72881a19fc116cad3f29b69015e0dcaea4265b4f`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.1/installers/gravwell_kafka_installer_5.1.1.sh) | ``4d4931f211fa2af3e2378b31556f11a6ce7f95b76971fc927154b0f6621f052b`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.1/installers/gravwell_kinesis_ingest_installer_5.1.1.sh) | ``65ca8206262cbbd4226962cede1ad19157c7163581eebf0e61012c98ae2d707d`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.1/installers/gravwell_pubsub_ingest_installer_5.1.1.sh) | ``7363d252bbbfaf5b9bd837c49d6284718d431555e626e082ba1f1ce621c5df4d`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.1/installers/gravwell_o365_installer_5.1.1.sh) | ``c3d727178d365fc4c90e2b5a9f8e43f430b121dbf893424fce37d0b7371597d9`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.1/installers/gravwell_msgraph_installer_5.1.1.sh) | ``19e0d20c4db9e754f75bf22aadc82f4944947d6a9f66a0320cfda964291d1598`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.1/installers/gravwell_datastore_installer_5.1.1.sh) | ``e54e06091dff9af3b9f35fd1e7925d837500aaff11fc47a0243e3745e04bbec9`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.1/installers/gravwell_offline_replication_installer_5.1.1.sh) | ``d6bbc5af34ade99cc8bad25d25e7e2c5cec73bf7534959261d4141324e1cc672`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.1/installers/gravwell_loadbalancer_installer_5.1.1.sh) | ``1341abafcaf2a610aaf6b000c72c7255a5a8c3b55816bade6be8f4bf37c8288e`` | |
