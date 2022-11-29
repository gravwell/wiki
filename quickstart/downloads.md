# Downloads

**Attention:** The Debian and RHEL repositories are more easily maintained than these standalone installers and are the recommended methods of installation. See the [quickstart instructions](quickstart).

## Gravwell Core

The Gravwell core installer contains the indexer and webserver frontend. You'll need a license; either get a Community Edition free license, or contact info@gravwell.io for commercial options.

<<<<<<< HEAD
[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_5.1.2.sh) (SHA256: 170faade301ebd0c67d1af76073822e0179f72f1c3ae6a1538a4eb6bed0fbbfa)
=======
[Download Gravwell Core Installer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_5.1.2.sh) (SHA256: b56f1ee6b45d2176a70e5cf0b51181986730a6617d7ff1fb97c4602ccb1cc0a9)
>>>>>>> upstream/dev

## Ingesters

The core suite of ingesters are available for download as installable packages.  Ingesters designed to operate on Linux machines are typically self contained, statically linked executables that are agnostic to the hosts package management system (with the exception of the NetworkCapture ingester).  Windows based ingesters are distributed as executable MSI packages.  Source code for many ingesters can be found at the [Gravwell Github](https://github.com/gravwell/gravwell/tree/master/ingesters) repository.

### Current Ingester Releases
| Ingester | SHA256 | More Info |
|:--------:|-------:|----------:|
| [Simple Relay](https://update.gravwell.io/archive/5.1.2/installers/gravwell_simple_relay_installer_5.1.2.sh) | ``c5e24190a0333724630360dce287f17b8cdc2dc023eb54b83e5993d5814eadc2`` | [Documentation](../ingesters/simple_relay)|
| [File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_installer_5.1.2.sh) | ``7916d444d883c4fd453771e04599f262c20aff5cc4e935fd0072169454f3f3ba`` | [Documentation](../ingesters/file_follow) |
| [HTTP Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_http_ingester_installer_5.1.2.sh) | ``37b6756337b510378cb61bfd6fce815da6ab87783b405c5d7793302273731cdb`` | [Documentation](../ingesters/http) |
| [IPMI Ingester](https://update.gravwell.io/archive/5.1.2/installers/gravwell_ipmi_installer_5.1.2.sh) | ``34551139829f5a1cd21a04a4eab812ae490026f94d7e83e89cca0c7c67df9a91`` | [Documentation](../ingesters/ipmi)|
| [Netflow Capture](http://update.gravwell.io/archive/5.1.2/installers/gravwell_netflow_capture_installer_5.1.2.sh) | ``a0ebba48a3f4888a97d9d3bf29744fe49cebd979fd37ab35f57b8d0d2a894307`` | [Documentation](../ingesters/netflow) |
| [Network Capture](https://update.gravwell.io/archive/5.1.2/installers/gravwell_network_capture_installer_5.1.2.sh) | ``6dad58d94ae5224174fdb356722f6854207094cf60055e124c655b6211ac965a`` | [Documentation](../ingesters/pcap) |
| [Collectd Collector](https://update.gravwell.io/archive/5.1.2/installers/gravwell_collectd_installer_5.1.2.sh) | ``dd855ecab99fdc80e94564e272f91e5b0c579169ad39299728e6d7231ed71bf9`` | [Documentation](../ingesters/collectd) |
| [Ingest Federator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_federator_installer_5.1.2.sh) | ``4c9274a5c680bc7d2ea088f01a672f4335d3a7318fb3b0cbbb2715e21c9a778d`` | [Documentation](../ingesters/federators/federator) |
| [Windows Events](https://update.gravwell.io/archive/5.1.2/installers/gravwell_win_events_5.1.2.msi) | ``eff3d858322863cb0b8603ed755565eb383903c958941fcf5b85cbb36a37b7f2`` | [Documentation](../ingesters/winevent) |
| [Windows File Follower](https://update.gravwell.io/archive/5.1.2/installers/gravwell_file_follow_5.1.2.msi) | ``6fb8babe0159423d865f2b05d32a92a64f4628c7d91f9ec06fad2773a6fd901e`` | [Documentation](../ingesters/win_file_follow) |
| [Apache Kafka](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kafka_installer_5.1.2.sh) | ``cf6d34e97eb87655890283a33740695c8a8aa84d234eca2a9568792ddf330761`` | [Documentation](../ingesters/kafka)|
| [Amazon Kinesis](https://update.gravwell.io/archive/5.1.2/installers/gravwell_kinesis_ingest_installer_5.1.2.sh) | ``ad14319d9886fd7988a92b2a67852b40413cb8731ce9a95e34e0b9c5f41fd4ab`` | [Documentation](../ingesters/kinesis)|
| [Google PubSub](https://update.gravwell.io/archive/5.1.2/installers/gravwell_pubsub_ingest_installer_5.1.2.sh) | ``c5a3d2e32999bc4544735afced3f471360dc5c562f8f45b810f881f543b3fa18`` | [Documentation](../ingesters/pubsub)|
| [Office 365 Logs](https://update.gravwell.io/archive/5.1.2/installers/gravwell_o365_installer_5.1.2.sh) | ``58cbabf158c004b5b5f0cf94dbf8579166c895cf5742a35d6606d6db7ca04acc`` | [Documentation](../ingesters/o365)|
| [Microsoft Graph API](https://update.gravwell.io/archive/5.1.2/installers/gravwell_msgraph_installer_5.1.2.sh) | ``60a3de7a6296410fcf14fb877a4eff760f97ec3216d9780c5bccef8dbae258ab`` | [Documentation](../ingesters/msg)|

## Other downloads

Some Gravwell components are distributed as optional additional installers, such as the search agent and the datastore.

| Component | SHA256 | More Info |
|:---------:|:------:|----------:|
| [Datastore](https://update.gravwell.io/archive/5.1.2/installers/gravwell_datastore_installer_5.1.2.sh) | ``a3e2678cada7dfd14fd9d20737a5c50c51732b7eea23ac4d91b6695d14c99f83`` | [Documentation](../distributed/frontend) |
| [Offline Replicator](https://update.gravwell.io/archive/5.1.2/installers/gravwell_offline_replication_installer_5.1.2.sh) | ``1deab6156d79918421d7f39625af7328812386bf6434de062e69718da3966473`` | [Documentation](../configuration/replication.md) |
| [Load Balancer](https://update.gravwell.io/archive/5.1.2/installers/gravwell_loadbalancer_installer_5.1.2.sh) | ``0356fb62554b18f1c5a4810ed1c0540242011a91bbb40699ea7bb68d7827ad39`` | [Documentation](../distributed/loadbalancer.md) |
| [Gravwell Tools](https://update.gravwell.io/archive/5.1.2/installers/gravwell_tools_5.1.2.sh) | ``98f0127e67845d1f4f07a3d997788ee9b379f56f0ce551b918628f9372c89c11`` | [Documentation](../tools/tools.md)|
