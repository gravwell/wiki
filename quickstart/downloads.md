# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.0.5/installers/gravwell_5.0.5.sh) (SHA256: 42580849b677ea2c22614551b1bc88a0696c5bac2c8d123f16b3cd30be767b5d)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.0.5/installers/gravwell_simple_relay_installer_5.0.5.sh) | ``12debd5c17a3da7bd182b12a719e86a8f70d46caf940a9a8585988f971a2f85d`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.0.5/installers/gravwell_file_follow_installer_5.0.5.sh) | ``3d5e08a651dc0ad127a24d0ae72281ad9d623970668b09d3da753f844148d032`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.0.5/installers/gravwell_http_ingester_installer_5.0.5.sh) | ``33f713d15657adc35ecca574c1ff1a54ccad3df1252b09e0d7827092902091b6`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.0.5/installers/gravwell_ipmi_installer_5.0.5.sh) | ``2c528e2e5d93c61f1651748ce8c6db816c5dd730cfb89dcffc6d36de2715c744`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.0.5/installers/gravwell_netflow_capture_installer_5.0.5.sh) | ``b4dda96b80d367750958afd5aee6f781a49b32ead7906c909a71b89a8ed85546`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.0.5/installers/gravwell_network_capture_installer_5.0.5.sh) | ``7497a84960995f949eb5adc9ffb6d345bba0346f38031c1e7a702a0408130b51`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.0.5/installers/gravwell_collectd_installer_5.0.5.sh) | ``6c9cae227e33f11aa0ab3eda2b3d1971f54c115e342e7229a7453445f43310a5`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.0.5/installers/gravwell_federator_installer_5.0.5.sh) | ``966be988afa06bf88f4164dd4dcdad046b452f22ac43446ca32e0f99b7561ef8`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.0.5/installers/gravwell_win_events_5.0.5.msi) | ``a7612e0188af28249f6afad18e63673bd0c3fd987959ab972549dadc8846f180`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.0.5/installers/gravwell_file_follow_5.0.5.msi) | ``0ed5466ec92bb778621eea3364ca83a18fa43dc096761d0ddad1b9260fff28cd`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.0.5/installers/gravwell_kafka_installer_5.0.5.sh) | ``ba77083818a630b181dae70fd640a5923eff7e9d67c686c5f6737b155ef7a4cf`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.0.5/installers/gravwell_kinesis_ingest_installer_5.0.5.sh) | ``f32c8659e14542101b9ce859fc8bacd34c9cbed6e68dcf8d0447ce0003190684`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.0.5/installers/gravwell_pubsub_ingest_installer_5.0.5.sh) | ``28a949454ac292cca194f519a7852a4087b7fb248deb670fa59c7ff89d8cc45f`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.0.5/installers/gravwell_o365_installer_5.0.5.sh) | ``d4b31bda509aaab46f41f8879e89ebbdcd8bc1bc7b83e0882eac5ca7a6a69435`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.0.5/installers/gravwell_msgraph_installer_5.0.5.sh) | ``d612deae0b583f08c866c56a5ec45e27260a9c1928a17e59708dc128e3014d58`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.0.5/installers/gravwell_datastore_installer_5.0.5.sh) | ``34a1baea2cfd9cf2f50180ea39655f817372cbad6806adf771dc4394f25edba7`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.0.5/installers/gravwell_offline_replication_installer_5.0.5.sh) | ``5baad707fec2e556f3768b222c7e57b9175b221cf074a74d4cb4bdf64dc6ac77`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.0.5/installers/gravwell_loadbalancer_installer_5.0.5.sh) | ``d7d7cc11a47c47957bfbddb5efa27b581c3668e8f317a1896c26ef685b15b1f9`` | |
