# Kinesis

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Kinesis Ingester](https://docs.gravwell.io/ingesters/kinesis.html)
:::

## Kinesis Configuration

In order to configure the Kinesis ingester for an existing stream you will need:
* An AWS access key (ID number & secret key)
* The region in which your stream resides
* The name of the stream itself

Once the stream is configured, each record in the Kinesis stream will be stored as a single entry in Gravwell.

## Gravwell Configuration

### Gravwell Storage Well Configuration
**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/kinesis.well`
```ini
[Storage-Well "kinesis"]
    Location=/opt/gravwell/storage/kinesis
    Tags=kinesis*
```
### Gravwell Ingester Configuration
**Sample Kinesis config:**  
Create or edit: `/opt/gravwell/etc/kinesis_ingest.conf`
```ini
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
