---
myst:
  substitutions:
    package: "gravwell-netflow-capture"
    standalone: "gravwell_netflow_capture"
    dockername: "netflow_capture"
---
# Netflow Ingester

The Netflow ingester acts as a Netflow collector (see [the Wikipedia article](https://en.wikipedia.org/wiki/NetFlow) for a full description of Netflow roles), gathering records created by Netflow exporters and capturing them as Gravwell entries for later analysis. These entries can then be analyzed using the [netflow](/search/netflow/netflow) search module.

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

The Netflow ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, the Netflow ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## Collector Examples

```
[Collector "netflow v5"]
	Bind-String="0.0.0.0:2055" #we are binding to all interfaces
	Tag-Name=netflow
	Assume-Local-Timezone=true
	Session-Dump-Enabled=true

[Collector "ipfix"]
	Tag-Name=ipfix
	Bind-String="0.0.0.0:6343"
	Flow-Type=ipfix
```

## Installation

If you're using the Gravwell Debian repository, installation is just a single apt command:

```
apt-get install gravwell-netflow-capture
```

Otherwise, download the installer from the [Downloads page](/quickstart/downloads). To install the Netflow ingester, simply run the installer as root (the actual file name will typically include a version number):

```console
root@gravserver ~ # bash gravwell_netflow_capture_installer.sh
```

If there is no Gravwell indexer on the local machine, the installer will prompt for an Ingest-Secret value and an IP address for an indexer (or a Federator). Otherwise, it will pull the appropriate values from the existing Gravwell configuration. In any case, review the configuration file in `/opt/gravwell/etc/netflow_capture.conf` after installation. A straightforward example which listens on UDP port 2055 might look like this:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO

[Collector "netflow v5"]
	Bind-String="0.0.0.0:2055" #we are binding to all interfaces
	Tag-Name=netflow
```

Note that this configuration sends entries to a local indexer via `/opt/gravwell/comms/pipe`. Entries are tagged 'netflow'.

You can configure any number of `Collector` entries listening on different ports with different tags; this can help organize the data more clearly.
