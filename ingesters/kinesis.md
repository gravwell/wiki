# Kinesis Ingester

Gravwell provides an ingester capable of fetching entries from Amazon's [Kinesis Data Streams](https://aws.amazon.com/kinesis/data-streams/) service. The ingester can process multiple Kinesis streams at a time, with each stream composed of many individual shards. The process of setting up a Kinesis stream is outside the scope of this document, but in order to configure the Kinesis ingester for an existing stream you will need:

* An AWS access key (ID number & secret key)
* The region in which your stream resides
* The name of the stream itself

Once the stream is configured, each record in the Kinesis stream will be stored as a single entry in Gravwell.

## Basic Configuration

The Kinesis ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the Kinesis ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## KinesisStream Examples

```
[KinesisStream "stream1"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as created in AWS
	Iterator-Type=TRIM_HORIZON
	Parse-Time=false
	Assume-Local-Timezone=true

[KinesisStream "stream2"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as created in AWS
	Iterator-Type=TRIM_HORIZON
	Metrics-Interval=60
	JSON-Metric=true
```

## Installation and configuration

First, download the installer from the [Downloads page](#!quickstart/downloads.md), then install the ingester:

```
root@gravserver ~# bash gravwell_kinesis_ingest_installer.sh
```

If the Gravwell services are present on the same machine, the installation script should automatically extract and configure the `Ingest-Auth` parameter and set it appropriately. You will now need to open the `/opt/gravwell/etc/kinesis_ingest.conf` configuration file and set it up for your Kinesis stream. Once you have modified the configuration as described below, start the service with the command `systemctl start gravwell_kinesis_ingest.service`

The example below shows a sample configuration which connects to an indexer on the local machine (note the `Pipe-Backend-target` setting) and feeds it from a single Kinesis stream named "MyKinesisStreamName" in the us-west-1 region.

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR
State-Store-Location=/opt/gravwell/etc/kinesis_ingest.state

# This is the access key *ID* to access the AWS account
AWS-Access-Key-ID=REPLACEMEWITHYOURKEYID
# This is the secret key which is only displayed once, when the key is created
#   Note: This option is not required if running in an AWS instance (the AWS
#         the AWS SDK handles that)
AWS-Secret-Access-Key=REPLACEMEWITHYOURKEY

[KinesisStream "stream1"]
	Region="us-west-1"
	Tag-Name=kinesis
	Stream-Name=MyKinesisStreamName	# should be the stream name as AWS knows it
	Iterator-Type=TRIM_HORIZON
	Parse-Time=false
	Assume-Localtime=true
```

Note the `State-Store-Location` option. This sets the location of a state file which will track the ingester's position in the streams, to prevent re-ingesting entries which have already been seen.

You will need to set at least the following fields before starting the ingester:

* `AWS-Access-Key-ID` - this is the ID of the AWS access key you wish to use
* `AWS-Secret-Access-Key` - this is the secret access key itself
* `Region` - the region in which the kinesis stream resides
* `Stream-Name` - the name of the kinesis stream

You can configure multiple `KinesisStream` sections to support multiple different Kinesis streams.

You can test the config by running `/opt/gravwell/bin/gravwell_kinesis_ingester -v` by hand; if it does not print out errors, the configuration is probably acceptable.

Most of the fields are self-explanatory, but the `Iterator-Type` setting deserves a note. This setting selects where the ingester starts reading data **if it does not have a state file entry** for the stream/shard. The default is "LATEST", which means the ingester will ignore all existing records and only read records created after the ingester starts. By setting it to TRIM_HORIZON, the ingester will start reading records from the oldest available. In most situations we recommend setting it to TRIM_HORIZON so you can fetch older data; on further runs of the ingester, the state file will maintain the sequence number and prevent duplicate ingestion.

The Kinesis ingester does not provide the `Ignore-Timestamps` option found in many other ingesters. Kinesis messages include an arrival timestamp; by default, the ingester will use that as the Gravwell timestamp. If `Parse-Time=true` is specified in the data consumer definition, the ingester will instead attempt to extract a timestamp from the message body.
