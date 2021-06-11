# Installation Checklists

## Single-Node Checklist

This checklist gives a general order of operation for configuring a single-node, standalone Gravwell instance. Refer to the [Quickstart](quickstart.md) for additional step-by-step instructions.

□ Install Gravwell via self-extracting installer, Debian/Redhat package, or Docker container (see [Quickstart](quickstart.md)).

□ Verify firewall rules allow incoming traffic on [ports used by Gravwell](#!configuration/networking.md).

□ Use web browser to access the new Gravwell instance, e.g. `http://gravwell.example.org/`, and upload your license file when prompted.

□ Login as "admin", password "changeme". Change default admin password by clicking user icon in upper right.

□ Configure any additional storage wells if desired. (See [Gravwell configuration](#!configuration/configuration.md) and [detailed configuration parameters](#!configuration/parameters.md) documentation)

□ Set up [ageout](#!configuration/ageout.md) on your wells to avoid running out of disk space.

□ Optional: [configure TLS](#!configuration/certificates.md) for user access and ingester connections.

□ [Configure ingesters](#!ingesters/ingesters.md) to bring data into Gravwell.


<!-- TODO: this is a complex process and 
## Cluster Checklist

### Preparation

□ Determine which nodes will be indexers and which will be webservers. If you intend to deploy more than one webserver, select one webserver to run the search agent.

□ If you intend to use [distributed frontends](#!distributed/frontend.md), provision an additional system for the *datastore*. Note that the datastore cannot be co-resident with an indexer or webserver process.

□ Install Gravwell on each of the webserver and indexer nodes (see [Quickstart](quickstart.md)).

□ Install the datastore if desired. This is included in the core shell installer, but is in a separate package for Debian and Redhat.

□ Install the loadbalancer if desired.

□ Deploy TLS certificates to webservers, datastore, and loadbalancer as appropriate. We recommend copying the certificate to `/opt/gravwell/etc/cert.pem` and the secret key to `/opt/gravwell/etc/key.pem`.

### Configuration

□ Copy one node's `gravwell.conf` file out to serve as the base for configurations. Remove any `Webserver-UUID` lines or `Indexer-UUID` lines.

#### Indexer Config

□ Make a copy of the config to be used for the indexers.

□ Define desired wells in the indexer config (see [this document](#!configuration/configuration.md).

□ Set [ageout configuration](#!configuration/ageout.md) for each well.

#### Webserver Config

□ Make a copy of the base config to be used for the webservers.

□ Set `Remote-Indexers` parameters to list all planned indexers, e.g.:
```
Remote-Indexers=net:indexer0.example.net:9404
Remote-Indexers=net:indexer1.example.net:9404
Remote-Indexers=net:indexer2.example.net:9404
```

□ If using a datastore, set the `Datastore` and `External-Addr` options in gravwell.conf as described in the [distributed frontends](#!distributed/frontend.md) document.

□ Set up [TLS](#!configuration/certificates.md) by setting the `Certificate-File` and `Key-File` fields.

### Deployment

□ Use systemd to disable un-needed Gravwell processes: disable webserver & searchagent on indexers, indexer on webservers. Make sure the searchagent process is only enabled on one webserver.

□ Copy indexer config to indexers, webserver config to webservers.

□ Restart gravwell processes on all nodes
-->