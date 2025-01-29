---
myst:
  substitutions:
    package: "gravwell-kafka"
    standalone: "gravwell_kafka"
    dockername: "kafka_consumer"
---
# Kafka

The Kafka ingester is designed to act as a consumer for [Apache Kafka](https://kafka.apache.org/) so that Gravwell can attach to a Kafka cluster and consume data.  Kafka can act as a high-availability [data broker](https://kafka.apache.org/uses#uses_logs) to Gravwell.  Kafka can take on some of the roles provided by the Gravwell Federator, or ease the burden of integrating Gravwell into an existing data flow.  If your data is already flowing to Kafka, integrating Gravwell is just an `apt-get` away.

The Gravwell Kafka ingester is best suited as a co-located ingest point for a single indexer.  If you are operating a Kafka cluster and a Gravwell cluster, it is best not to duplicate the load balancing characteristics of Kafka at the Gravwell ingest layer.  Each indexer should be configured with its own Kafka ingester, allowing the Kafka cluster to manage load balancing; install the Kafka ingester on the same machine as the Gravwell indexer and use the Unix named pipe connection for communication with the indexer.

Most Kafka configurations enforce a data durability guarantee, which means data is stored in non-volatile storage when consumers are not available to consume it.  As a result we do not recommend that the Gravwell ingest cache be enabled on Kafka ingester; instead, let Kafka provide the data durability.

## Installation

```{include} installation_instructions_template 
```

## Configuration

The Kafka ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, the Kafka Ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/kafka.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/kafka.conf.d`).

### Consumer Configurations

The Gravwell Kafka ingester can subscribe to multiple topics and even multiple Kafka clusters.  Each consumer defines a consumer block with a few key configuration values.

The following parameters configure the connection to the Kafka cluster:

| Parameter | Type | Descriptions | Required |
|-----------|------|--------------| -------- |
| Leader    | slice of host:port | The set of Kafka cluster leader/broker.  This should be an IP or hostname; if no port is specified the default port of 9092 is appended. Multiple can be specified. | YES |
| Topic     | string | The Kafka topic this consumer will read from | YES |
| Consumer-Group | string | The Kafka consumer group this ingester is a member of; default is `gravwell`. |
| Rebalance-Strategy | slice of string | The re-balancing strategy to use when reading from Kafka. Options are `roundrobin`, `sticky`, and `range`. |
| Auth-Type | string | Enable SASL authentiation and specify mechanism. |
| Username | string | Specify username for SASL authentication. |
| Password | string | Specify password for SASL authentication. |
| Use-TLS | boolean | If set, the ingester will connect to the Kafka cluster using TLS. |
| Insecure-Skip-TLS-Verify | boolean | If TLS is in use, setting this parameter will make the ingester ignore invalid TLS certificates. |

These parameters configure how the ingester handles incoming data from Kafka:

| Parameter | Type | Descriptions | Required |
|-----------|------|--------------| -------- |
| Default-Tag | string | Entries which do not receive a tag from the `Tag-Header` will be assigned this default tag. | YES |
| Tag-Header | string | If set, the ingester will look at the specified header to determine into which tag the entry should be ingested. If the header is not set on the message, the `Default-Tag` will be used. By default, `Tag-Header` is set to "TAG". |
| Tags | string | Specifies a list of allowable tags (or wildcard patterns) for the `Tag-Header`, e.g. `Tags=gravwell,foo,b*r`. Any entry with a tag which does not match one of the patterns will instead be assigned the `Default-Tag`. |
| Source-Header | string | Gravwell producers will often put the data source address in a message header, if set the ingester will attempt to interpret the given header as a Source address.  If the header is not correct the ingester will apply the source override (if set) or the default source. |
| Source-As-Binary | boolean | If set, the ingester will assume that the contents of the `Source-Header` are in binary format, rather than a string. |
| Synchronous | boolean | If set, the ingester will perform a sync on the ingest connection every time a Kafka batch is written. |
| Batch-Size | integer | The number of entries to read from Kafka before forcing a write to the ingest connection; the default is 512. |

These parameters give some standard Gravwell ingester configuration options related to timestamps, timezones, and the source field. See the [general ingester configuration page](/ingesters/ingesters) for more information about these parameters.

| Parameter | Type | Descriptions | Required |
|-----------|------|--------------| -------- |
| Source-Override | IPv4 or IPv6 | An IP address to use as the SRC for all entries. |
| Ignore-Timestamps | boolean | If set, the ingester will apply the current timestamp to all received entries, ignoring Kafka timestamps. |
| Extract-Timestamps | boolean | If set, the ingester will ignore the Kafka timestamps and attempt to extract a timestamp from the entry's contents. |
| Assume-Local-Timezone | boolean | If set, when extracting timestamps from entries the timezone will be assumed to be local, if not explicitly set. |
| Timezone-Override | string | If set, timestamps will be parsed in the given timezone, e.g. "America/New_York". |
| Timestamp_Format_Override | string | Specifies a timestamp format, e.g. "RFC822", to use when parsing timestamps. |

As with most ingesters, each consumer may also specify [preprocessors](/ingesters/preprocessors/preprocessors) if needed.

### Consumer Examples

```
[Consumer "default"]
	Leader="127.0.0.1"
	Default-Tag=default   #send bad tag names to default tag
	Tags=*                #allow all tags
	Topic=default
	Tag-Header=TAG        #look for the tag in the Kafka TAG header
	Source-Header=SRC     #look for the source in the Kafka SRC header

# This consumer does not specify a Tags parameter, so all entries will get the Default-Tag
[Consumer "test"]
	Leader="kafka1.example.org" #leader one
	Leader="kafka2.example.org" #leader two
	Default-Tag=test
	Topic=test
	Consumer-Group=mygroup
	Synchronous=true
	Source-Header=SRC #A custom feeder is putting its source IP in the header named "SRC"
	Batch-Size=256 #get up to 256 messages before consuming and pushing
        # this config sets sticky as the preferred strategy but also offers roundrobin
	Rebalance-Strategy=sticky
	Rebalance-Strategy=roundrobin
```

```{warning}
Setting any consumer as synchronous causes that consumer to continually sync the ingest pipeline.  It will have significant performance implications for ALL consumers.
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
	Default-Tag=default
	Tags=*
	Topic=default
	Consumer-Group=gravwell1
	Batch-Size=256


[Consumer "test"]
	Leader="tasks.testcluster.internal:9092"
	Default-Tag=test
	Topic=test
	Consumer-Group=testgroup
	Source-Override="192.168.1.1"
	Rebalance-Strategy=range
	Batch-Size=4096
```
