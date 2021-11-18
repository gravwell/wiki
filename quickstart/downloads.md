# Downloads

Attention: The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](#!quickstart/quickstart.md).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

[Download Gravwell Core Installer](https://update.gravwell.io/archive/4.2.7/installers/gravwell_4.2.7.sh) (SHA256: 66f2fb2bc20476990253c745411cff774badd9a22f157fa1a710ef4176a8823f)

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/4.2.7/installers/gravwell_simple_relay_installer_4.2.7.sh) | ``8732517a54a556bfc58c734b2acd5f15d1635a9affc7b05b5283a2cf93567196`` | [Documentation](#!ingesters/ingesters.md#Simple_Relay)|
| [File Follower](https://update.gravwell.io/archive/4.2.7/installers/gravwell_file_follow_installer_4.2.7.sh) | ``3c1d3cf07b81ee71421e14f7da7d4307384c33ecb03087ab2e6d46c5b641a1fa`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [HTTP Ingester](https://update.gravwell.io/archive/4.2.7/installers/gravwell_http_ingester_installer_4.2.7.sh) | ``0e4e6dc7a18fd4f2563ab0c32f6010035ca5bba6a36751c7e60266318fcdcd97`` | [Documentation](#!ingesters/ingesters.md#HTTP_POST) |
| [IPMI Ingester](https://update.gravwell.io/archive/4.2.7/installers/gravwell_ipmi_installer_4.2.7.sh) | ``6687e184bf4451a217141f1780c1f11b7ae9bed04ee108174e452fe42a948647`` | [Documentation](#!ingesters/ingesters.md#IPMI_Ingester)|
| [Netflow Capture](http://update.gravwell.io/archive/4.2.7/installers/gravwell_netflow_capture_installer_4.2.7.sh) | ``773c460eeda8970306e54d602641dc2dea7363d248c6bacd6718ed9dd04c9f73`` | [Documentation](#!ingesters/ingesters.md#Netflow_Ingester) |
| [Network Capture](https://update.gravwell.io/archive/4.2.7/installers/gravwell_network_capture_installer_4.2.7.sh) | ``0ba057cc8ee470592e6d521b31ceabcd9298e06905a138d211668d86da8e45b6`` | [Documentation](#!ingesters/ingesters.md#Network_Ingester) |
| [Collectd Collector](https://update.gravwell.io/archive/4.2.7/installers/gravwell_collectd_installer_4.2.7.sh) | ``d16edb54949b7fc65b020555d9b22ba438637c52ebd767d3cf8618d89e1256fc`` | [Documentation](#!ingesters/ingesters.md#collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/4.2.7/installers/gravwell_federator_installer_4.2.7.sh) | ``72bff906afd6d8a6242750b575a1ce87205ddc5a0930e0edda53dff1e392fc99`` | [Documentation](#!ingesters/ingesters.md#Federator_Ingester) |
| [Windows Events](https://update.gravwell.io/archive/4.2.7/installers/gravwell_win_events_4.2.7.msi) | ``7cdda07eae68ba8abe59746ab1e5cd04bb8b9ebfdbf22fce02e9149553b8bafe`` | [Documentation](#!ingesters/ingesters.md#Windows_Event_Service) |
| [Windows File Follower](https://update.gravwell.io/archive/4.2.7/installers/gravwell_file_follow_4.2.7.msi) | ``9309264c27a8a6e3b50e5ec612cc5c0ed1d345f77651f6ac028b96b084e4ae83`` | [Documentation](#!ingesters/ingesters.md#File_Follower) |
| [Apache Kafka](https://update.gravwell.io/archive/4.2.7/installers/gravwell_kafka_installer_4.2.7.sh) | ``f5af2b3571ae86bc06557b4ee2d229a63b0b9fe754235ae8bb5223c263773e0a`` | [Documentation](#!ingesters/ingesters.md#Kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/4.2.7/installers/gravwell_kinesis_ingest_installer_4.2.7.sh) | ``b80cbe98efff7b9ba43699c5ea091f3a16f7c946b262ec549eddc14f5648604d`` | [Documentation](#!ingesters/ingesters.md#Kinesis_Ingester)|
| [Google PubSub](https://update.gravwell.io/archive/4.2.7/installers/gravwell_pubsub_ingest_installer_4.2.7.sh) | ``b0c87108c71ef9ca131d27727a9a894964bac28d72652c2655cf9e09faef5b46`` | [Documentation](#!ingesters/ingesters.md#GCP_PubSub)|
| [Office 365 Logs](https://update.gravwell.io/archive/4.2.7/installers/gravwell_o365_installer_4.2.7.sh) | ``8b295fa76f60f6750e202edbec354d829ed7155f60c25d64f74fb0519dc5cd6d`` | [Documentation](#!ingesters/ingesters.md#Office_365_Log_Ingester)|
| [Microsoft Graph API](https://update.gravwell.io/archive/4.2.7/installers/gravwell_msgraph_installer_4.2.7.sh) | ``552c06fec5e55bdbb380c0a499461c56d8ca13c54077176432464a79766ecbc5`` | [Documentation](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/4.2.7/installers/gravwell_datastore_installer_4.2.7.sh) | ``fd0222a73fc169c86e927d484f1fdfa016335c740e597bf7ca787fada84e3876`` | [Documentation](#!distributed/frontend.md) |
| [Offline Replicator](https://update.gravwell.io/archive/4.2.7/installers/gravwell_offline_replication_installer_4.2.7.sh) | ``7ad20d890c11d540d56a5bca104dcc90df52ebc18822cf0dfc59cfc13c1e778a`` | [Documentation](#!configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/4.2.7/installers/gravwell_loadbalancer_installer_4.2.7.sh) | ``93c730d9d10ceccccf8ee35bc08ca558bef9a6383c6692ca6311965c4ac55f4f`` | |
