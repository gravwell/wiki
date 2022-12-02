# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.3/installers/gravwell_5.1.3.sh) (SHA256: e08647c0d1bda9a7e71907eb5cfe839170f99c82fa97bda02e4eca9f7e054af7)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.3/installers/gravwell_simple_relay_installer_5.1.3.sh) | ``7255e981b9956ac7bd01f0c4b95c47f99e824a6915ba84d94ccf9665c0e9f365`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.3/installers/gravwell_file_follow_installer_5.1.3.sh) | ``ffeadccb15c58120c60bf0ab17d17643536ed1302d6e58d037b7d678d89a3db4`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.3/installers/gravwell_http_ingester_installer_5.1.3.sh) | ``c7555f49e6638a9caa00699c39b42bb256c93e49522ee297a11b7ba077472334`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.3/installers/gravwell_ipmi_installer_5.1.3.sh) | ``d77d1525e947ba3f598885c5f9497909a950be9eeb4d9eeffe6ac50cd32c24a8`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.3/installers/gravwell_netflow_capture_installer_5.1.3.sh) | ``e2e005523640859ad54abe20319a0166255e7a1fef0a875398e9f672fb4dadda`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/5.1.3/installers/gravwell_network_capture_installer_5.1.3.sh) | ``ffb1382f1fe75a813f93e1ba9528c0502d3c6c3bd10fa24a00e46b8804b4109c`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.3/installers/gravwell_collectd_installer_5.1.3.sh) | ``a86c13c39d595100e20df58d1f6c7141bddeb02f20c696cb8a44d71f81ae7cf8`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.3/installers/gravwell_federator_installer_5.1.3.sh) | ``06ba9e327ff587c4b504cd74d3a18e010ceee7482960a24248d14609604e8133`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/5.1.3/installers/gravwell_win_events_5.1.3.msi) | ``f04a7a51d4e3bcf6de239e4a8752c2f96a160eab72fbfd57292d57af1ce07548`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.3/installers/gravwell_file_follow_5.1.3.msi) | ``5a19d78ec08cad4ad2f43741378f9c14c1d53d6aa24e97b4daae232f4d507c88`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.3/installers/gravwell_kafka_installer_5.1.3.sh) | ``7eecfa8e259ab0732806be2db8bdfc760c03506e3463501c36435abf0dae470c`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.3/installers/gravwell_kinesis_ingest_installer_5.1.3.sh) | ``75e0546755a305792b5c3139a2ceb5dfe20d19902165d96c728fc8cdd13f5c19`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.3/installers/gravwell_pubsub_ingest_installer_5.1.3.sh) | ``9653c145bfc02562dd5e9d52588640893e17dd8e0e9c3ce70e1e20fa406b3845`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.3/installers/gravwell_o365_installer_5.1.3.sh) | ``1b2f1f2ead2905b7eae936495a96fbb91d0621ef2c27282c75613e7336d3fd5f`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.3/installers/gravwell_msgraph_installer_5.1.3.sh) | ``a19b828285de8872603010f541bce3451df2971bb832b1897b647d318d698b52`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.3/installers/gravwell_datastore_installer_5.1.3.sh) | ``5f20f5672212642e00188117fa5419478a1257032044e5e09271cbdcf79cc317`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.3/installers/gravwell_offline_replication_installer_5.1.3.sh) | ``d8f75680cba553f7a2b57a0c1effaa3273e5aec2a84b3f58ad45afe29d4a0271`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.3/installers/gravwell_loadbalancer_installer_5.1.3.sh) | ``96d3239cbe581ba8fafafc96b5f7cd82206f440125f5a421265128df4b957728`` | |
| [Gravwell Tools](https://update.gravwell.io/archive/5.1.3/installers/gravwell_tools_5.1.3.sh) | ``718e57cf3d8fc1ff3d0fa6a80c4fb058e68dcf54ee999879564155dafe21b210`` | [Documentation](#!/tools/tools.md)|

