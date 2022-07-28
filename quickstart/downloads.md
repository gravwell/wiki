# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.4/installers/gravwell_5.0.4.sh) (SHA256: 23c452ca11984eec13695ed1c913ffa92761b86f91d35c6047e390d45b9b5ecb)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.4/installers/gravwell_simple_relay_installer_5.0.4.sh) | ``85a791fbf56459198f6452db8a570d38a1a2615cad2b246d4aff34e2f2dd458f`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.4/installers/gravwell_file_follow_installer_5.0.4.sh) | ``df2a90bdeb005c9867c675ff57cda0ab07b5ad474f85c129aa8a1ebca28ec612`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.4/installers/gravwell_http_ingester_installer_5.0.4.sh) | ``feb03f59cf577d5dcfc3e44f0e257658b40fe1e9b7344601e2f06a00d010c9de`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.4/installers/gravwell_ipmi_installer_5.0.4.sh) | ``1a35aa0a0f4634b3fc2a4acd0c64468afc96767a41b9628541b735a7f53a087e`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.4/installers/gravwell_netflow_capture_installer_5.0.4.sh) | ``432d9e6317ae2742e815f84fa3128699f91ce230465a40506c663b760ebe66eb`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.4/installers/gravwell_network_capture_installer_5.0.4.sh) | ``2ff005306b1d73963da3f0dacde76848e04f1f97cc68082a928bd3407b935a24`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.4/installers/gravwell_collectd_installer_5.0.4.sh) | ``83df97a0b952a14fcfd1d759d9eb5856137d582cab09de6284eb5437b13c2c27`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.4/installers/gravwell_federator_installer_5.0.4.sh) | ``b7f09f1c4d478ff6604abe9cb4adfcf0cf3b21d073d1051a8e869ba7386a0e5e`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.4/installers/gravwell_win_events_5.0.4.msi) | ``fee19c2af84585baef1461c5d35cabdd06f3ec73497bf0c62ef4cd5b481b47d2`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.4/installers/gravwell_file_follow_5.0.4.msi) | ``0aec2923b19b3701767bb39dba1df617474c4366ed60fd89f95c14b0dedfa1a2`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.4/installers/gravwell_kafka_installer_5.0.4.sh) | ``aeac15b3c42a5e9e53e0828aaeb504cebbea6789a6add56bf4dbe0ee5ea1f84a`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.4/installers/gravwell_kinesis_ingest_installer_5.0.4.sh) | ``5a08f91680d343ddcefb0394fbcb928d0afc786cf332ee92b90ef2d8a97d7823`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.4/installers/gravwell_pubsub_ingest_installer_5.0.4.sh) | ``b81d5a6b103af778a59ec98d3bd39d2d88a0e225f75d8f8baf5e1622d5bc4fcf`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.4/installers/gravwell_o365_installer_5.0.4.sh) | ``a089d56113713bec42dfa66c8e5058e38fbb5548f0833e48eb7340732b475dfb`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.4/installers/gravwell_msgraph_installer_5.0.4.sh) | ``f5b2de24c5c91628344564890365ccf5ee22fcc91d72e6a7e434a62cc782be83`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.4/installers/gravwell_datastore_installer_5.0.4.sh) | ``ee8e32d75106a9ee921d63de2705a57dc730d9ce99947e49551aad8411d2b452`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.4/installers/gravwell_offline_replication_installer_5.0.4.sh) | ``23bd97ffbf071a57d4acd791bf7a736c5a1f6f555afb3c32f0ce753eb0fbc0ba`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.4/installers/gravwell_loadbalancer_installer_5.0.4.sh) | ``915c7e2d3dcd4055ac99d1c02dcd1c575ff009f1d63b496e22d8cf170e93ca69`` | |
