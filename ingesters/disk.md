# Disk Monitor

The diskmonitor ingester is designed to take periodic samples of disk activity and ship the samples to gravwell.  The disk monitor is extremely useful in identifying storage latency issues, looming disk failures, and other potential storage problems.  We at Gravwell actively monitor our own storage infrastructure with the disk monitor to study how queries are operating and to identify when the storage infrastructure is behaving badly.  We were able to identify a RAID array that transitioned to write-through mode via a latency plot even when the RAID controller failed to mention it in the diagnostic logs.

The disk monitor ingester is available on [github](https://github.com/gravwell/ingesters)

![diskmonitor](diskmonitor.png)

## Basic Configuration

The Disk Monitor ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, Disk Monitor supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## Session Ingester

The session ingester is a specialized tool used to ingest larger, single records. The ingester listens on a given port and upon receiving a connection from a client it will aggregate any data received into a single entry.

This enables behavior such as indexing all of your Windows executable files:

```
for i in `ls /path/to/windows/exes`; do cat $i | nc 192.168.1.1 7777 ; done
```

The session ingester is driven via command line parameters rather than a persistent configuration file.

```
Usage of ./session:
  -bind string
        Bind string specifying optional IP and port to listen on (default "0.0.0.0:7777")
  -clear-conns string
        comma separated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -max-session-mb int
        Maximum MBs a single session will accept (default 8)
  -pipe-conns string
        comma separated list of paths for named pie connection
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
```

## Notes

The session ingester is not formally supported, nor is there an installer available.  The source code for the session ingester is available on [github](https://github.com/gravwell/ingesters).

