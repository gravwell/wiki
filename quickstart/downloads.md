# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and is the recommended method of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.1.9/installers/gravwell_4.1.9.sh) (SHA256: 92fb647d6adfc3ed0329631b3be4be5533f416a97444e880ce89dec472ef1652)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.1.9/installers/gravwell_simple_relay_installer_4.1.9.sh) | ``a30e40ef9c662f52aed0dc778931cdb9e539ed78ac19fa643e24063aa866d9dd`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.1.9/installers/gravwell_file_follow_installer_4.1.9.sh) | ``ee6ee1e8b57d83ea6d18730d34fabd96eadbb75aaa174582322a2710f8c489aa`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.1.9/installers/gravwell_http_ingester_installer_4.1.9.sh) | ``c88e3748dedc03b57fc6f88f8f4ba5f85908b1c5c93dfc9ca8a78cdc56a752f8`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.1.9/installers/gravwell_ipmi_installer_4.1.9.sh) | ``1bb0205a60b7f5c8ec2de6d606af0bebff95065b0b74d3a57354c5b1d156dd06`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.1.9/installers/gravwell_netflow_capture_installer_4.1.9.sh) | ``97b53c8daf5c910b239d249ebf4a93318d8296518e82e687e637fe3fd98fbe3b`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.1.9/installers/gravwell_network_capture_installer_4.1.9.sh) | ``24585a12550a93d303416fc954eb2dcb10fcdc3862a215b3993603f3d36519d4`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.1.9/installers/gravwell_collectd_installer_4.1.9.sh) | ``13cfb75462ea428b276637f8b7f2504aae42f3074f6ca9ac534a616d4e31564b`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.1.9/installers/gravwell_federator_installer_4.1.9.sh) | ``1fecd4b1f3fbe904c780a19fa000dae10521a7d31f9c1e21b962582fd5b4abf3`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.1.9/installers/gravwell_win_events_4.1.9.msi) | ``37227530d67eca34d8e98f975cd325d523f6f9a29d0c2247b059169855c609ee`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.1.9/installers/gravwell_file_follow_4.1.9.msi) | ``8040d8f925bedb09500defcf1b71017243e35651805a38ebd16dad1443a8ffc5`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.1.9/installers/gravwell_kafka_installer_4.1.9.sh) | ``7c121df8258b75c0ae9dcbd23abc7127dcd4086c964f7a483f78231ef785449f`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.1.9/installers/gravwell_kinesis_ingest_installer_4.1.9.sh) | ``aa690a87a72f05e3fa5251babcb288bdf6117cd70ab6267b4ba04fd0a6ca5e14`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.1.9/installers/gravwell_pubsub_ingest_installer_4.1.9.sh) | ``711aafcbc2fc217acfbfe443f0d0fbab00cad20ea123ab972adeb79e244799b3`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.1.9/installers/gravwell_o365_installer_4.1.9.sh) | ``7c90c6d7390cc769fbb98d20b09a90ddf5bd9bae0a60411c635a9c313e3ed75c`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.1.9/installers/gravwell_msgraph_installer_4.1.9.sh) | ``d160278ac418c02704813821a511c59d8f698224bef58f901a42872349061f48`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.1.9/installers/gravwell_datastore_installer_4.1.9.sh) | ``6360a936569b709bd9afc6b919b892cd308c7cc30a4092c44aacddc3410b5fab`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.1.9/installers/gravwell_offline_replication_installer_4.1.9.sh) | ``74a410b9bbb7637e385cd559ed89a9091b08030fce97594fb898728e02011f66`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.1.9/installers/gravwell_loadbalancer_installer_4.1.9.sh) | ``381afa4d9bad0005517d287f6feb386473c523418b6f5c54c6232cb1ae9008e1`` | |
