## Forwarding Preprocessor

The forwarding preprocessor is used to split a log stream and forward it to another endpoint.  This can be extremely useful when standing up additional logging tools or when feeding data to external archive processors.  The Forwarding preprocessor is a forking preprocessor.  This means that it will not alter the data stream; it only forwards a copy of the data on to additional endpoints.

By default the forwarding preprocessor is blocking, which means that if you specify a forwarding endpoint using a stateful connection like TCP or TLS and the endpoint is not up or is unable to accept data it will block ingestion.  This behavior can be altered using the `Non-Blocking` flag or by using the UDP protocol.

The forwarding preprocessor also supports several filter mechanisms to cut down or specify exactly which data streams are forwarded.  Streams can be filtered using entry tags, sources, or regular expressions which operate on the actual log data.  Each filter specification can be specified multiple times to create an OR pattern.

Multiple forwarding preprocessors can be specified, allowing for a specific log stream to be forwarded to multiple endpoints.

The Forwarding preprocessor Type is `forwarder`.

### Supported Options

* `Target` (string, required): Endpoint for forwarded data.  Should be a host:port pair unless using the `unix` protocol.
* `Protocol` (string, required): Protocol to use when forwarding data.  Options are `tcp`, `udp`, `unix`, and `tls`.
* `Delimiter` (string, optional): Optional delimiter to use when sending data via the `raw` output format.
* `Format` (string, optional): Output format to send data.  Options are `raw`, `json`, and `syslog`.
* `Tag` (string, optional, multiple allowed): Tags used to filter events.  Multiple specifications implies an OR.
* `Regex` (string, optional, multiple allowed): Regular expressions used to filter events.  Multiple specifications implies an OR.
* `Source` (string, optional, multiple allowed): IP or CIDR specifications used to filter events.  Multiple specifications implies an OR.
* `Timeout` (unsigned int, optional, specified in seconds):  Timeout on connection and write attempts for the forwarder.
* `Buffer` (unsigned int, optional): Specifies the number of events that the forwarder can buffer when attempting to send data.
* `Non-Blocking` (boolean, optional): True specifies that the forwarder will make a best effort to forward data, but will not block ingestion.
* `Insecure-Skip-TLS-Verify` (boolean, optional): Specifies that TLS-based connections can ignore invalid TLS certificates.

### Example: Forwarding syslog from a specific set of hosts

For this example we are using the SimpleRelay ingester to ingest syslog messages and forward them in their raw form to another endpoint.  We are using the `Source` filter in the `forward` preprocessor to *only* forward logs that have the source field from either the `192.168.0.1` IP or within the `192.168.1.0/24` subnet.  The logs are forwarded on in their original format with a newline between each:

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=forwardprivnet

[Preprocessor "forwardprivnet"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false
```

### Example: Forwarding Specific Windows Event Logs

For this example we are using the Federator to forward data streams from potentially many downstream ingesters.  We are using both the `Tag` and `Regex` filters to capture a specific set of entries and forward them on.  Note that we are using the `syslog` format, which means that we will send data to the endpoint with an RFC5424 header and the data in the body of the syslog message.  Forwarded data using the syslog format uses `gravwell` as the Hostname and the entry TAG as the Appname.

The `Tag` filters specify that we only want to forward entries that are using the `windows` or `sysmon` tags.

The `Regex` filters are used so that we only get event data from specific Channel/EventID combinations, namely login events from the security provider and execution events from the sysmon provider.

```
[IngestListener "enclaveB"]
	Ingest-Secret = IngestSuperSecrets
	Cleartext-Bind = 0.0.0.0:4123
	Tags=win*
	Preprocessor=forwardloginsysmon

[Preprocessor "forwardloginsysmon"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="syslog"
	Buffer=128
	Tag=windows
	Tag=sysmon
	Regex="Microsoft-Windows-Sysmon.+>(1|22)</EventID>"
	Regex="Microsoft-Windows-Security-Auditing.+>(4624|4625|4626)</EventID>"
	Non-Blocking=false
```


### Example: Forwarding logs to multiple hosts

For this example we are using the Gravwell Federator to forward subsets of logs to different endpoints using different formats.  Because the forwarder preprocessor can be stacked the same way as any other preprocessor, we can specify multiple forwarding preprocessors with their own filters, endpoints, and formats.

```
[IngestListener "enclaveA"]
	Ingest-Secret = IngestSuperSecrets
	Cleartext-Bind = 0.0.0.0:4023
	Tags=win*
	Preprocessor=forwardloginsysmon
	Preprocessor=forwardprivnet

[Preprocessor "forwardloginsysmon"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="syslog"
	Buffer=128
	Tag=windows
	Tag=sysmon
	Regex="Microsoft-Windows-Sysmon.+>(1|22)</EventID>"
	Regex="Microsoft-Windows-Security-Auditing.+>(4624|4625|4626)</EventID>"
	Non-Blocking=false

[Preprocessor "forwardsyslog"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Tag=syslog
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false
```
