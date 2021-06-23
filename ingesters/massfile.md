# Mass File Ingester

The Mass File ingester is a very powerful but specialized tool for ingesting an archive of many logs from many sources.

## Basic Configuration

The Mass File ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the Mass File ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## Example use case

Gravwell users have used this tool when investigating a potential network breach. The user had Apache logs from over 50 different servers and needed to search them all. Ingesting them one after another causes poor temporal indexing performance. This tool was created to ingest the files while preserving the temporal nature of the log entries and ensuring solid performance.  The massfile ingester works best when the ingesting machine has enough space (storage and memory) to optimized the source logs prior to ingesting.  The optimization phase helps relieve pressure on the Gravwell storage system at ingest and during search, ensuring that incident responders can move quickly and get performant access to their log data in short order.

## Notes

The mass file ingester is driven via command line parameters and is not designed to run as a service.  The code is available on [Github](https://github.com/gravwell/ingesters).

```
Usage of ./massFile:
  -clear-conns string
        comma separated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -no-ingest
        Optimize logs but do not perform ingest
  -pipe-conn string
        Path to pipe connection
  -s string
        Source directory containing log files
  -skip-op
        Assume working directory already has optimized logs
  -tag-name string
        Tag name for ingested data (default "default")
  -timeout int
        Connection timeout in seconds (default 1)
  -tls-conns string
        comma separated server:port list of TLS connections
  -tls-private-key string
        Path to TLS private key
  -tls-public-key string
        Path to TLS public key
  -tls-remote-verify string
        Path to remote public key to verify against
  -w string
        Working directory for optimization
```
