# Networking Considerations for Gravwell

Gravwell uses several network ports for communication between distributed components. This article describes which ports are used for which purposes.

## Indexer Control Port: TCP 9404

This port, set by the `Control-Port` option in gravwell.conf, is used by the webserver to communicate with the indexers. Ensure that any firewalls on the *indexers* allow incoming connections on this port from the *webserver*, and that no network infrastructure blocks this port between the webserver and the indexers.

## Webserver Port: TCP 80/443

This port is how Gravwell users access the Gravwell webserver. The default configuration uses unencrypted HTTP on port 80, specified with the `Web-Port` option in gravwell.conf. This can be changed to another value, e.g. 8080 if desired. We recommend changing the port to 443 if you [install TLS certificates](#!configuration/certificates.md).

## Cleartext Ingest Port: TCP 4023

This port is used by ingesters to connect to indexers and upload entries via unencrypted communications. The default port is TCP 4023, but it can be changed using the `Ingest-Port` option in gravwell.conf. Because ingesters and indexers are often on entirely different networks, it is essential that firewalls are configured such that the *ingesters* are allowed to connect to this port on the *indexers*

## TLS Ingest Port: TCP 4024

This port is used by ingesters to connect to indexers and upload entries via TLS-encrypted communications. The default port is TCP 4024, but it can be changed using the `TLS-Ingest-Port` option in gravwell.conf. Because ingesters and indexers are often on entirely different networks, it is essential that firewalls are configured such that the *ingesters* are allowed to connect to this port on the *indexers*

## Indexer Replication Port: TCP 9606

This port is used by indexers to communicate with each other for [replication](#!configuration/replication.md). The default port is 9606, if not otherwise specified in the `Peer` and `Listen-Address` options of the Replication portion of gravwell.conf. Only indexers use this port.

## Datastore Port: TCP 9405

This port is used when a Gravwell cluster has [multiple webservers](#!distributed/frontend.md) configured. The *datastore* component listens on this port (specified using the `Datastore-Port` option) for incoming connections from *webservers*.
