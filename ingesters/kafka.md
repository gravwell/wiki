---
myst:
  substitutions:
    package: "gravwell-kafka"
    standalone: "gravwell_kafka"
    dockername: "kafka_consumer"
---
# Kafka

The Kafka ingester designed to act as a consumer for [Apache Kafka](https://kafka.apache.org/) so that data Gravwell can attach to a Kafka cluster and consume data.  Kafka can act as a high availability [data broker](https://kafka.apache.org/uses#uses_logs) to Gravwell.  Kafka can take on some of the roles provided by the Gravwell Federator, or ease the burden of integrating Gravwell into an existing data flow.  If your data is already flowing to Kafka, integrating Gravwell is just an `apt-get` away.

The Gravwell Kafka ingester is best suited as a co-located ingest point for a single indexer.  If you are operating a Kafka cluster and a Gravwell cluster, it is best not to duplicate the load balancing characteristics of Kafka at the Gravwell ingest layer.  Install the Kafka ingester on the same machine as the Gravwell indexer and use the Unix named pipe connection.  Each indexer should be configured with its own Kafka ingester, this way the Kafka cluster can manage load balancing.

Most Kafka configurations enforce a data durability guarantee, which means data is stored in non-volatile storage when consumers are not available to consume it.  As a result we do not recommend that the Gravwell ingest cache be enabled on Kafka ingester, instead let Kafka provide the data durability.

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

The Kafka ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, the Kafka Ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/kafka.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/kafka.conf.d`).

## Consumer Examples

```
[Consumer "default"]
	Leader="127.0.0.1"
	Default-Tag=default   #send bad tag names to default tag
	Tags=*                #allow all tags
	Topic=default
	Tag-Header=TAG        #look for the tag in the Kafka TAG header
	Source-Header=SRC     #look for the source in the Kafka SRC header

[Consumer "test"]
	Leader="127.0.0.1:9092"
	Tag-Name=test
	Topic=test
	Consumer-Group=mygroup
	Synchronous=true
	Key-As-Source=true #A custom feeder is putting its source IP in the message key value
	Header-As-Source="TS" #look for a header key named TS and treat that as a source
	Source-As-Text=true #the source value is going to come in as a text representation
	Batch-Size=256 #get up to 256 messages before consuming and pushing
	Rebalance-Strategy=roundrobin
```

## Installation

The Kafka ingester is available in the Gravwell Debian repository as a Debian package as well as a shell installer on our [Downloads page](/quickstart/downloads).  Installation via the repository is performed using `apt`:

```
apt-get install gravwell-kafka
```

The shell installer provides support for any non-Debian system that uses systemd, including Arch, Redhat, Gentoo, and Fedora.

```console
root@gravserver ~ # bash gravwell_kafka_installer.sh
```

## Configuration

The Gravwell Kafka ingester can subscribe to multiple topics and even multiple Kafka clusters.  Each consumer defines a consumer block with a few key configuration values.


| Parameter | Type | Descriptions | Required |
|-----------|------|--------------| -------- |
| Tag-Name  | string | The Gravwell tag that data should be sent to.  | YES |
| Leader    | host:port | The Kafka cluster leader/broker.  This should be an IP or hostname, if no port is specified the default port of 9092 is appended | YES |
| Topic     | string | The Kafka topic this consumer will read from | YES |
| Consumer-Group | string | The Kafka consumer group this ingester is a member of | NO - default is `gravwell` |
| Source-Override | IPv4 or IPv6 | An IP address to use as the SRC for all entries | NO |
| Rebalance-Strategy | string | The re-balancing strategy to use when reading from Kafka | NO - default is `roundrobin`.  `sticky`, and `range` are also options |
| Key-As-Source | boolean | Gravwell producers will often put the data source address in a message key, if set the ingester will attempt to interpret the message key as a Source address.  If the key structure is not correct the ingester will apply the override (if set) or the default source. | NO - default is false |
| Synchronous | boolean | The ingester will perform a sync on the ingest connection every time a Kafka batch is written. | NO - default is false |
| Batch-Size | integer | The number of entries to read from Kafka before forcing a write to the ingest connection | NO - default is 512 |
| Auth-Type | string | Enable SASL authentiation and specify mechanism |
| Username | string | Specify username for SASL authentication |
| Password | string | Specify password for SASL authentication |

```{warning}
Setting any consumer as synchronous causes that consumer to continually Sync the ingest pipeline.  It will have significant performance implications for ALL consumers.
```

```{note}
Setting a large `Batch-Size` when using `Synchronous=true` can help with performance under heavy load.
```

### Authentication

The Kafka ingester supports the following SASL authentication mechanisms: `PLAIN`, `SCRAMSHA256`, and `SCRAMSHA512`.  To enable authentication set the `Auth-Type` configuration parameter to the desired authentication type and provide a `Username` and `Password`.  Each consumer must set its own authentication.

The following is a valid configuration snippet from a listener named `tester` using plaintext authentication:

```
[Consumer "tester"]
	Leader="127.0.0.1"
	Default-Tag=default   #send bad tag names to default tag
	Tags=*                #allow all tags
	Topic=default
	Auth-Type=plain
	Username=TheDude
	Password="a super secret password"
```

Valid `Auth-Type` options are `plain`, `scramsha256`, and `scramsha512`.



### Example Configuration

Here is an example configuration that is subscribing to two different topics using two different consumer groups.

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-target=/opt/gravwell/comms/pipe
Log-Level=INFO
Log-File=/opt/gravwell/log/kafka.log

[Consumer "default"]
	Leader="tasks.kafka.internal"
	Tag-Name=default
	Topic=default
	Consumer-Group=gravwell1
	Key-As-Source=true
	Batch-Size=256


[Consumer "test"]
	Leader="tasks.testcluster.internal:9092"
	Tag-Name=test
	Topic=test
	Consumer-Group=testgroup
	Source-Override="192.168.1.1"
	Rebalance-Strategy=range
	Batch-Size=4096
```
