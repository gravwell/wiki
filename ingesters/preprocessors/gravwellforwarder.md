## Gravwell Forwarding Preprocessor

The Gravwell forwarding processor can duplicate entries to multiple instances of Gravwell.  This preprocessor can be useful for testing or in situations where a specific data stream needs to be duplicated to an additional set of Gravwell indexers.  The Gravwell forwarding preprocessor utilizes the same configuration structure to specify indexers, ingest secrets, and even cache controls as the packaged ingesters.  The Gravwell forwarding preprocessor is a blocking preprocessor, this means that if you do not enable a local cache it can block the ingest pipeline if the preprocessor cannot forward entries to the specified indexers.

The Gravwell Forwarding preprocessor Type is `gravwellforwarder`.

### Supported Options

See the [Global Configuration Parameters](#!ingesters/ingesters.md#Global_Configuration_Parameters) section for full details on all the Gravwell ingester options.  Most global ingester configuration options are supported by the Gravwell Forwarder preprocessor.

### Example: Duplicating Data In a Federator

For this example we are going to specify a complete Federator configuration that will duplicate all entries to a second cluster. Incoming entries will be sent to both of the following Gravwell systems:

* The cluster composed of 172.19.0.4 and 172.19.0.5
* The single-node system 172.19.0.2

NOTE: We are enabling an `always` cache on the forwarding preprocessor so that it won't ever block the normal ingest path.

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Verify-Remote-Certificates = true
Cleartext-Backend-Target=172.19.0.2:4023 #example of adding a cleartext connection
Log-Level=INFO

[IngestListener "enclaveA"]
	Ingest-Secret = CustomSecrets
	Cleartext-Bind = 0.0.0.0:4423
	Tags=windows
	Tags=syslog-*
	Preprocessor=dup

[Preprocessor "dup"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets
	Connection-Timeout = 0
	Cleartext-Backend-Target=172.19.0.4:4023 #indexer1
	Cleartext-Backend-Target=172.19.0.5:4023 #indexer2 (cluster config)
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup.cache # must be a unique path
	Max-Ingest-Cache=1024 #Limit forwarder disk usage
```

### Example: Stacking Duplicate Forwarders

For this example we are going to specify a complete Federator configuration and multiple Gravwell preprocessors so that we can duplicate our single stream of entries to multiple Gravwell clusters. Each entry will be sent to the following separate Gravwell systems:

* 172.19.0.101
* 172.19.0.102
* 172.19.0.103
* 172.19.0.2

NOTE: The preprocessor control logic does NOT check whether or not you are not forwarding to the same cluster multiple times. This may actually be desired; by using the forwarder preprocessor in combination with other preprocessors (such as the regex tag routing preprorcessor), the same entry may be ingested into multiple different tags on the same indexer.

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Verify-Remote-Certificates = true
Cleartext-Backend-Target=172.19.0.2:4023 #example of adding a cleartext connection
Log-Level=INFO

[IngestListener "enclaveA"]
	Ingest-Secret = CustomSecrets
	Cleartext-Bind = 0.0.0.0:4423
	Tags=windows
	Tags=syslog-*
	Preprocessor=dup1
	Preprocessor=dup2
	Preprocessor=dup3

[Preprocessor "dup1"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets1
	Cleartext-Backend-Target=172.19.0.101:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup1.cache
	Max-Ingest-Cache=1024

[Preprocessor "dup2"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets2
	Cleartext-Backend-Target=172.19.0.102:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup2.cache
	Max-Ingest-Cache=1024

[Preprocessor "dup3"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets3
	Cleartext-Backend-Target=172.19.0.103:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup3.cache
	Max-Ingest-Cache=1024
```


