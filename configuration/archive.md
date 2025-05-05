# Cloud Archive

Gravwell supports an ageout mechanism called Cloud Archive.  When Cloud Archive is enabled, indexers will upload shards to the Cloud Archive server before deleting them from storage.  Gravwell Cloud Archive is an excellent method for long term archival storage for data that does not need to be actively searchable but must be retained.  The Cloud Archive service can be hosted on a variety of storage platforms and is designed to provide a remote, low-cost storage platform.  Cloud Archive configuration can be enabled on a per-well basis, which means you can decide which data sets warrant long term archival.

```{attention}
Indexers will not delete data until they have successfully uploaded it to the archive server.  If the indexer cannot upload due to connectivity issues, misconfigurations, or poor network throughput they will not delete data.  The inability to delete data may cause indexers to run out of storage and cease ingesting new data.  If a Cloud Archive upload fails to complete, the user interface will display a notification with the failure.
```

```{attention}
The Cloud Archive system compresses data while in transit, which requires some CPU resources when uploading.  Pushing data to a remote system also takes time, depending on available bandwidth and CPU.  Be sure to leave yourself a little headroom when defining ageout parameters to account for additional time consumed by Cloud Archive -- if you are ingesting at 1Gbps but only have a 500Mbps uplink, you may not be able to archive shards as fast as new data comes in!
```

## Configuring Indexers

Every indexer can define a single Cloud Archive configuration block which specifies the remote archive server and authentication token. The configuration block is specified using the `[Cloud-Archive]` section header.  To enable Cloud Archive on a well, add the "Archive-Deleted-Shards=true" directive within that well.

Here is an example configuration with three wells:

```
[global]
Web-Port=443
Control-Port=9404
Ingest-Port=4023

[Cloud-Archive]
	Archive-Server=test.archive.gravwell.io:8886
	Archive-Shared-Secret="password"

[Default-Well]
	Location=/opt/gravwell/storage/default/
	Cold-Location=/opt/gravwell/cold_storage/default/
	Hot-Duration=1d
	Cold-Duration=30d
	Delete-Frozen-Data=true
	Archive-Deleted-Shards=true

[Storage-Well "netflow"]
	Location=/opt/gravwell/storage/netflow/
	Hot-Duration=7d
	Delete-Cold-Data=true
	Archive-Deleted-Shards=true
	Tags=netflow

[Storage-Well "raw"]
	Location=/opt/gravwell/storage/raw/
	Hot-Duration=7d
	Delete-Cold-Data=true
	Tags=pcap
```

The above example has 3 configured wells (default, netflow, and raw).  The default well uses both a hot and cold storage tier which means that data is archived when it would normally roll out of the cold storage tier.  The netflow well contains only a hot storage tier, its data will be uploaded when it would normally be deleted after 7 days.  The raw well does not have Cloud Archive enabled (the default is Archive-Deleted-Shards=false), so its data will not be uploaded.

## Hosting Cloud Archive

The Cloud Archive service is designed to be self-hosted and potentially integrated into other larger infrastructures. The code is open-sourced and available at [github.com/gravwell/cloudarchive](https://github.com/gravwell/cloudarchive). It is also packaged for Debian, Redhat, and as a shell installer.

### Installing the Server
To install on Debian:

```
apt install gravwell-cloudarchive-server
```

On Redhat:

```
yum install gravwell-cloudarchive-server
```

As a standalone shell installer (downloaded from [the downloads page](/quickstart/downloads)):

```
sh gravwell_cloudarchive_server_installer_x.x.x.sh
```

These commands will *install* the server, but not *configure* it. Read on for instructions on configuration.

### Password Database

Use the `gravwell_cloudarchive_usertool` command to set up the password database with an entry for your customer number:

```
sudo su gravwell -c "/opt/gravwell/bin/gravwell_cloudarchive_usertool -action useradd -id 11111 -passfile /opt/gravwell/etc/cloud.passwd"
```

The tool will prompt for the passphrase to use for the specified customer number. You can find your customer number on the License page of the Gravwell UI.

```{note}
Indexers will authenticate to the cloud archive service using the customer license number *on the indexer*. In an [overwatch](/distributed/overwatch) configuration, this number may be different from the license number deployed on the *webservers*.
```

### Server Configuration

A default configuration will be installed at `/opt/gravwell/etc/cloudarchive_server.conf`. This configuration stores archived shards on the local disk, in `/opt/gravwell/cloudarchive/`. It listens for clients on port 8886, using the specified TLS cert/key pair for encryption. The `Password-File` parameter points at the password database set up earlier.

```
[Global]
Listen-Address="0.0.0.0:8886"
Cert-File=/opt/gravwell/etc/cert.pem
Key-File=/opt/gravwell/etc/key.pem
Password-File=/opt/gravwell/etc/cloud.passwd
Log-Level=INFO
Log-File=/opt/gravwell/log/cloudarchive.log
Storage-Directory=/opt/gravwell/cloudarchive/
```

The following config archives incoming data shards to an FTP server instead of the local disk. Note the specification of the `FTP-Server`; the `FTP-Username` and `FTP-Password` fields should be for a valid account on that FTP server. The `Storage-Directory` parameter is still required; this directory will be used as temporary storage for archive operations.

```
[Global]
Listen-Address="0.0.0.0:8886"
Cert-File=/opt/gravwell/etc/cert.pem
Key-File=/opt/gravwell/etc/key.pem
Password-File=/opt/gravwell/etc/cloud.passwd
Log-Level=INFO
Log-File=/opt/gravwell/log/cloudarchive.log
Storage-Directory=/opt/gravwell/cloudarchive/
Backend-Type=ftp
FTP-Server=ftp.example.org:21
FTP-Username=cloudarchiveuser
FTP-Password=ca_secret_password
```

```{note}
If you don't want to set up certificates for TLS, you can put the server into plaintext mode by setting `Disable-TLS=true`. Be aware that this is horribly insecure and a terrible idea unless your Cloud Archive server and your indexers are on the same trusted network!
```

### Configure Gravwell

Configure your Gravwell indexers as above, setting the `Cloud-Archive` stanza to point at your server. The `Archive-Shared-Secret` value should match the password you entered when running gravwell_cloudarchive_usertool.

```
[Cloud-Archive]
	Archive-Server=archive.example.org:8886
	Archive-Shared-Secret="mysecrettoken"

[Default-Well]
	Location=/opt/gravwell/storage/default/
	Cold-Location=/opt/gravwell/cold_storage/default/
	Hot-Duration=1d
	Cold-Duration=30d
	Delete-Frozen-Data=true
	Archive-Deleted-Shards=true
```

If you disabled TLS on the server, set `Insecure-Disable-TLS=true` in the `Cloud-Archive` stanza. If you are using self-signed certs, set `Insecure-Skip-TLS-Verify=true`.

These are the available parameters for the `Cloud-Archive` configuration block.

* `Archive-Server` (required): specifies the destination host and port for the Cloud-Archive server.
* `Archive-Shared-Secret` (required): specifies the user secret/password for authentication to the Cloud-Archive server.
* `Archive-User-ID` (optional): specifies the user id for authentication to the Cloud-Archive server, defaults to zero.
* `Insecure-Skip-TLS-Verify` (optional): specifies that the client will ignore invalid TLS certificates when connecting to the Cloud-Archive server, defaults to false.
* `Insecure-Disable-TLS` (optional): specifies that the client will use a cleartext connection when communicating with the Cloud-Archive server, defaults to false.

### Backend Processors

#### File Backend Type

The File backend type is the default Cloud Archive backend plugin, if no `Backend-Type` is specified in the `[Global]` section then the `file` type is assumed.  The `file` driver is designed to store raw Gravwell shards organized in an way that allows rapid re-import or direct binding by a Gravwell instance.  Use the file type if your storage array can support seek operations and you wish to be able to directly bind a Gravwell instance to the storage location and perform queries without thawing or re-ingesting the data.

##### Required Configuration Parameters

* `Storage-Directory`: Directory must be writable and POSIX compliant.

#### FTP Backend Type

The FTP backend type is an optional variation of the `file` backend type that directly stores shards on a remote FTP server rather than local file storage.  This backend type can be useful as a method to access some tape storage systems that are frontended by FTP protocols.  The FTP backend type does not modify shards and could potentially enable direct binding to data without thawing depending on the underlying storage architecture.

##### Required Configuration Parameters

* `Backend-Type`: Must be configured as "ftp" to enable the FTP backend.
* `FTP-Server`
* `FTP-User`: May be "anonymous" for unauthenticated access.
* `FTP-Password`: May be empty when using "anonymous" for unauthenticated access.

#### S3 Backend Type

The S3 backend type is designed for system agnostic, low cost, long term storage of archive data. S3 backed storage is re-encoded and compressed into smaller chunks and then uploaded to S3-compatible storage, shards sent to S3 cannot be directly bound to by a Gravwell instance and must be re-thawed for query.  The data formats are open and designed to support alternate query data access methods like AWS S3 Select.

##### Required Configuration Parameters

* `Backend-Type`: Must be configured as "ftp" to enable the FTP backend.
* `Storage-Directory`: Directory must be writable and POSIX compliant, uploaded shards are temporarily stored until workers can process them.

##### S3 Specific Configuration Block

The S3 backend requires additional configuration in an `[s3]` configuration block.

| Parameter | Required | Example | Description |
|:----------|:---------|:--------|------------:|
| Checkpoint-File | TRUE | Checkpoint-File="/opt/gravwell/etc/cloudarchive.db" | Designate a local database path to track incoming shards and work progress. |
| ID        | TRUE | ID="s3-token-id" | S3 access token. |
| Secret    | TRUE | Secret="randoms3accessstring" | S3 access secret. |
| Region    | TRUE | Region="us-west-01" | S3 region. |
| Endpoint  | TRUE | Endpoint="http://foo.bar.baz/mybucket" | S3 endpoint URL. |
| Bucket    | TRUE | Bucket="mybucket" | S3 Bucket name. |
| Bucket-Path | FALSE | Bucket-Path="/data/shards" | Optional path to be prepended to shard chunk paths. |
| Encoder   | TRUE | Encoder="json" | Shard chunk encoder, options are `json`, `gob`, `bson`. |
| Compressor | TRUE | Compressor="zstd" | Shard chunk compressor, options are `zstd`, `zstd-fast`, `zstd-best`, `gz`, `gz-best`, `snappy`. |
| Target-Chunk-Size | FALSE | Target-Chunk-Size=67108864 | Shard chunk target size for S3 storage, defaults to 64MB. |
| Max-In-Memory-Size | TRUE | Max-In-Memory-Size=1073741824 | Maximum in memory size of shard chunks per worker, used to control memory usage with many workers. |
| Max-Retries | FALSE | Max-Retries=3 | Optional retry count for S3 uploads, many S3 providers can fail randomly, this option controls how many times the process will retry before giving up and aborting. |
| Worker-Count | FALSE | Worker-Count=5 | Optional number of concurrent workers that process shards, defaults to 1.  Worker-Count does not apply to uploads, at least 2 cores should be reserved for handling uploads from Gravwell indexers. |


##### Example S3 Cloud Archive Configuration

```
[Global]
Listen-Address="0.0.0.0:8886"
Cert-File=/opt/gravwell/etc/cert.pem
Key-File=/opt/gravwell/etc/key.pem
Password-File=/opt/gravwell/etc/cloud.passwd
Log-Level=INFO
Log-File=/opt/gravwell/log/cloudarchive.log
Storage-Directory=/opt/gravwell/cloudarchive/
Backend-Type=s3

[S3]
	ID="mytokenid"
	Secret="mytokensecret"
	Region="uswest0004"
	Endpoint="http://foo.bar.baz/mybucket"
	Bucket="test-bucket"
	Bucket-Path="/data/shards"
	Encoder="json"
	Compressor="zstd"
	Target-Chunk-Size=67108864 # 64MB
	Max-In-Memory-Size=1073741824 #1GB
	Min-Storage-GB-Available=64
	Max-Retries=3
	Worker-Count=4
```
