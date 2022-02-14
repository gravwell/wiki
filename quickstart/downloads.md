# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_4.2.9.sh) (SHA256: 6e72987f5985c47a2804fe4eca19a6a9a50bffa6d84fc2ddda33b0f163ca01f3)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.9/installers/gravwell_simple_relay_installer_4.2.9.sh) | ``63c0281c91c2f10e0ba4cac2fd21661da47ad5eaad70cda0da639de636cfd15d`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_installer_4.2.9.sh) | ``72bac508d34f7e396546969b80bbe986956c252606c68730cf1593cacbda9368`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_http_ingester_installer_4.2.9.sh) | ``053918a6c0657b2697b6ed55aac9c44501060ea18e18d0498ce00865a26aa404`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.9/installers/gravwell_ipmi_installer_4.2.9.sh) | ``d1bef4596ec5b71a6813669cec417b6b907b87e3ce2b075dec1dcfaf52bb82bf`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.9/installers/gravwell_netflow_capture_installer_4.2.9.sh) | ``7bf296230727a4ef392e5a275b52017cc9d80a08d5005e8ee800bc8b4d6494aa`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.9/installers/gravwell_network_capture_installer_4.2.9.sh) | ``55218bb368b60f282941e306beaa93b8abcd76d84e65449e8283ca4c3c7aeb3c`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.9/installers/gravwell_collectd_installer_4.2.9.sh) | ``4a46bac024d73278510ce61f8a6be75b1b763eca2533e6573d84086389525118`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_federator_installer_4.2.9.sh) | ``3e728dca56b60e522ec063f71577cc4fa924e5cf21d3b659e07248a064681b01`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.9/installers/gravwell_win_events_4.2.9.msi) | ``6991295b3057a69b004f6c2362d53b142c54311c2082bcdaee97870d05abf2fd`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.9/installers/gravwell_file_follow_4.2.9.msi) | ``686ef8ed15f8d9805cd9239989214b0e7ecd7ab0f56dedd5bd24c040af823861`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kafka_installer_4.2.9.sh) | ``48ee6d99574d43d1167baf07e4519c2c9edcb3b89c153224c0ba257bedb00d08`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.9/installers/gravwell_kinesis_ingest_installer_4.2.9.sh) | ``2a863ce68dfbcbefd82306ef9844948853f2d68f7f70a6153aee007f955fb7c2`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.9/installers/gravwell_pubsub_ingest_installer_4.2.9.sh) | ``5d2568246da7f52e6a2e0fe2b718adc97d157440e7d1518f580f6477ae7bd878`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.9/installers/gravwell_o365_installer_4.2.9.sh) | ``49552a4c6ef9c20a1f3985fedecd635bf5c439ca902fd481fa833d986c8159d9`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.9/installers/gravwell_msgraph_installer_4.2.9.sh) | ``c52006d0d891f281516a960f7a501efb51884b40e6fea4daa8f4194a2674b921`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.9/installers/gravwell_datastore_installer_4.2.9.sh) | ``40cf5d0fd8a144781e9dfd7076945513d7ab375590c66cbaf734f04252b45be4`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.9/installers/gravwell_offline_replication_installer_4.2.9.sh) | ``591ed084075eef53deeb154e3c92a14e8769371fbcb4cc76c5ca804dcf97b10c`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.9/installers/gravwell_loadbalancer_installer_4.2.9.sh) | ``19e4416b960a188fe385dcba93880ac897ff9e6c991f6e73934b02bcaa6393d3`` | |
