# Environment Variables

The indexer, webserver, and ingester components support configuring some parameters via environment variables rather than config files. This helps make more generic configuration files for bigger deployments. Configuration variables that can contain multiple directives are configured in environment variables using a comma separated list. For instance, to specify the ingest secret when launching the Federator:

```
GRAVWELL_INGEST_SECRET=MyIngestSecret /opt/gravwell/bin/gravwell_federator
```

If "_FILE" is added to the end of the environment variable name, Gravwell assumes the variable contains the path to a file which in turn contains the desired data. This is particularly useful in combination with [Docker's "secrets" feature](https://docs.docker.com/engine/swarm/secrets/).

```
GRAVWELL_INGEST_AUTH_FILE=/run/secrets/ingest_secret /opt/gravwell/bin/gravwell_indexer
```

Note: Environment variable values are **only** used when the corresponding field is not explicitly set in the appropriate config file (gravwell.conf or an ingester's config file).

## Indexer and Webserver

The table below shows which `gravwell.conf` parameters can be set via environment variables for the indexer and the webserver. Note that these variables are only used if the parameter is **not** configured in `gravwell.conf`.

| gravwell.conf variable | Environment Variable | Example |
|:------|:----|:---|----:|
| Ingest-Auth | GRAVWELL_INGEST_AUTH | GRAVWELL_INGEST_AUTH=CE58DD3F22422C2E348FCE56FABA131A |
| Control-Auth | GRAVWELL_CONTROL_AUTH | GRAVWELL_CONTROL_AUTH=C2018569D613932A6BBD62A03A101E84 |
| Indexer-UUID | GRAVWELL_INDEXER_UUID | GRAVWELL_INDEXER_UUID=a6bb4386-3433-11e8-bc0b-b7a5a01a3120 |
| Webserver-UUID | GRAVWELL_WEBSERVER_UUID | GRAVWELL_WEBSERVER_UUID=b3191f54-3433-11e8-a0c2-afbff4695836 |
| Remote-Indexers | GRAVWELL_REMOTE_INDEXERS | GRAVWELL_REMOTE_INDEXERS=172.20.0.1:9404,172.20.0.2:9404|
| Replication-Peers | GRAVWELL_REPLICATION_PEERS | GRAVWELL_REPLICATION_PEERS=172.20.0.1:9406,172.20.0.2:9406 |
| Datastore | GRAVWELL_DATASTORE | GRAVWELL_DATASTORE=172.20.0.10:9405 |

## Ingesters

Ingesters can also accept some parameters as environment variables rather than setting them explicitly in the configuration file.

| Config file variable | Environment Variable | Example |
|:------|:----|:---|
| Ingest-Secret | GRAVWELL_INGEST_SECRET | GRAVWELL_INGEST_SECRET=CE58DD3F22422C2E348FCE56FABA131A |
| Log-Level | GRAVWELL_LOG_LEVEL | GRAVWELL_LOG_LEVEL=DEBUG |
| Cleartext-Backend-target | GRAVWELL_CLEARTEXT_TARGETS | GRAVWELL_CLEARTEXT_TARGETS=172.20.0.1:4023,172.20.0.2:4023 |
| Encrypted-Backend-target | GRAVWELL_ENCRYPTED_TARGETS | GRAVWELL_ENCRYPTED_TARGETS=172.20.0.1:4024,172.20.0.2:4024 |
| Pipe-Backend-target | GRAVWELL_PIPE_TARGETS | GRAVWELL_PIPE_TARGETS=/opt/gravwell/comms/pipe |


### Federator-specific variables

Because the federator may run many listeners, each with a different ingest secret associated with it, it recognizes a special set of environment variables to configure those listener secrets at runtime.

Each listener has a name. In the example below, the listener is named "base":

```
[IngestListener "base"]
	Cleartext-Bind = 0.0.0.0:4023
	Tags=syslog
```

In order to specify an ingest secret for that listener at runtime, we use the variable `FEDERATOR_base_INGEST_SECRET`:

```
FEDERATOR_base_INGEST_SECRET=SuperSecret /opt/gravwell/bin/gravwell_federator
```

Or we can specify a file as with other environment variables:

```
FEDERATOR_base_INGEST_SECRET_FILE=/run/secrets/federator_base_secret /opt/gravwell/bin/gravwell_federator
```

### Datastore-specific variables

The [Datastore](#!distributed/frontend.md) can be configured at run-time by environment variables:

| gravwell.conf variable | Environment variable | Example |
|------------------------|----------------------|---------|
| Datastore-Listen-Address | GRAVWELL_DATASTORE_LISTEN_ADDRESS | GRAVWELL_DATASTORE_LISTEN_ADDRESS=192.168.1.100 |
| Datastore-Port | GRAVWELL_DATASTORE_LISTEN_PORT | GRAVWELL_DATASTORE_LISTEN_PORT=9995 |
