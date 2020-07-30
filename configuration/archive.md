# Cloud Archive

Gravwell supports an ageout mechanism called Cloud Archive.  Cloud archive is a remote service where data can be remotely archived prior to deleting.  Gravwell Cloud Archive is an excellent method for long term archival storage for data that does not need to be actively searchable but must be retained.  The Cloud Archive service can be hosted on a variety of storage platforms and is designed to provide a remote, low cost storage platform.  Cloud Archive configuration can be enabled on a per-well basis, which means you can decide which data sets warrant long term archival.

The archive system ensures that data is successfully uploaded to the archive server before it is deleted during normal ageout.

Attention: Indexers will not delete data until they have successfully uploaded it to the archive server.  If the indexer cannot upload due to connectivity issues, misconfigurations, or poor network throughput they will not delete data.  The innability to delete data may cause indexers to run out of storage and cease ingesing new data.  If a Cloud Archive upload fails to complete the user interface will display a notification with the failure.

Attention: The Cloud Archive system compresses data while in transit which requires some CPU resources when uploading.  Pushing data to a remote system also requires time, depending on available bandwidth and CPU.  Be sure to leave yourself a little headroom when defining ageout parameters to account for additional time consumed by Cloud Archive.

## Configuring Indexers

Every indexer has a global Cloud Archive configuration block which specifies the remote archive server and authentication token. The configuration block is specified using the "[cloud-archive]" header in the global section.  To enable Cloud Archive on a well, add the "Archive-Deleted-Shards=true" directive within the well.

Here is an example configuration with three wells:

```
[global]
Web-Port=443
Control-Port=9404
Ingest-Port=4023

[Cloud-Archive]
	Archive-Server=test.archive.gravwell.io
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

The above example has 3 configured wells (default, netflow, and raw).  The default well uses both a hot and cold storage tier which means that data is archived when it would normally roll out of the cold storage tier.  The netflow well contains only a hot storage tier, its data will be uploaded when it would normally be deleted after 7 days.  The raw well does not have Cloud Archive enabled (Archive-Deleted-Shards=false), its data will not be uploaded.

## Hosting Cloud Archive

The Cloud Archive service is a module service designed to be self-hosted and potentially integrated into other larger infrastructures.  If you are interested in hosting your own Cloud Archive service or would like to remotely archive your data, contact sales@gravwell.io.

Note: Indexers will authenticate to the cloud archive service using the customer license number *on the indexer*. In an [overwatch](#!distributed/overwatch.md) configuration, this number may be different from the license number deployed on the *webservers*.