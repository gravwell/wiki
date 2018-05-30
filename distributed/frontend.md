# The distributed Gravwell frontend

Just as Gravwell is designed to have multiple indexers operating at once, it can also have multiple frontends operating at once, pointing to the same set of indexers. Having multiple frontends allows for load balancing and high availability.

Once configured, distributed frontends will synchronize resources, users, dashboards, user preferences, and user search histories.

## The datastore server

Gravwell uses a separate server process called the datastore to keep frontends in sync. It can run on its own machine or it can share a server with a frontend. Fetch the datastore installer from [the downloads page](#!quickstart/downloads.md), then run it on the machine which will contain the datastore.

### Configuring the datastore server

The datastore server should be ready to run without any changes to `gravwell.conf`. To enable the datastore server at boot-time and start the service, run the following commands:

```
systemctl enable gravwell_datastore.service
systemctl start gravwell_datastore.service
```

#### Advanced datastore config

By default, the datastore server will listen on all interfaces over port 9405. If for some reason you need to change this uncomment and set the following line in your `/opt/gravwell/etc/gravwell.conf` file:

```
Datastore-Listen-Address=10.0.0.5	# listen only on 10.0.0.5
Datastore-Port=9555					# listen on port 9555 instead of 9405
```

## Configuring frontends for distributed operation

To tell a frontend to start communicating with a datastore, set the `Datastore` field in the "global" section of the webserver's `/opt/gravwell/etc/gravwell.conf`. For example, if the datastore server was was running on the machine with IP 10.0.0.5 and the default datastore port, the entry would look like this:

```
Datastore=10.0.0.5:9405
```

Note: By default, the frontend will check in with the datastore every 10 seconds. This can be modified by setting the `Datastore-Update-Interval` field to the desired number of seconds. Be warned that waiting too long between updates will make changes propagate very slowly between frontends, while overly-frequent updates may cause undue system load. 5 to 10 seconds is a good choice.

## Disaster recovery

Due to the synchronization techniques used by the datastore and frontends, care must be taken if the datastore server is re-initialized or replaced. Once a frontend has synchronized with a datastore, it considers that datastore the ground truth on all topics; if a resource does not exist on the datastore, but the frontend had previously synchronized that resource with the datastore, the frontend will delete the resource.

The datastore stores data in the following locations:

* `/opt/gravwell/etc/datastore-users.db` (user database)
* `/opt/gravwell/etc/datastore-webstore.db` (dashboards, user preferences, search history)
* `/opt/gravwell/etc/resources/datastore/` (resources)

If any of these locations are accidentally lost or deleted, they should be restored from one of the frontend systems before restarting the datastore. Assuming the datastore is on the same machine as one of the frontends, use the following commands:

```
cp /opt/gravwell/etc/users.db /opt/gravwell/etc/datastore-users.db
cp /opt/gravwell/etc/webstore.db /opt/gravwell/etc/webstore.db
cp -r /opt/gravwell/etc/resources/webserver/* /opt/gravwell/etc/resources/datastore/
```

If the datastore is on a separate machine, use `scp` or another file transfer method to copy those files from a frontend server.

## Load-balancing

Although distributed frontends *allow* load-balancing for users, we require the use of an external tool to perform the actual load-balancing / reverse proxy. If your network already uses Nginx+, Apache, or another tool for load-balancing, simply configure it to load balance between your frontends, being sure to enable persistent or "sticky" sessions.

If on the other hand you do not already have a reverse proxy configured for load balancing, it is easy to set up [Træfik](https://traefik.io) as a standalone load balancer for Gravwell frontends.

We recommend putting Traefik on its own machine, or at least not on the same server as a frontend.

First, grab the latest release of Traefik from [the Traefik releases page](https://github.com/containous/traefik/releases) or compile it yourself.

You'll also need a certificate for SSL. Either import a valid cert, or generate a self-signed certificate with the following command:

```
openssl req -newkey rsa:4096 -nodes -sha512 -x509 -days 3650 -nodes -out traefik.crt -keyout traefik.key
```

Next, save the following config file as `traefik.toml`:

```
defaultEntryPoints = ["http", "https"]

InsecureSkipVerify = true

[file]

[entryPoints]
        [entryPoints.http]
        address = ":80"
        [entryPoints.https]
        address = ":443"
                [entryPoints.https.tls]
                        [[entryPoints.https.tls.certificates]]
                        certFile = "traefik.crt"
                        keyFile = "traefik.key"

[frontends]
        [frontends.frontend1]
                backend = "backend1"
        [frontends.frontend1.headers]
                SSLRedirect = true
                SSLTemporaryRedirect = true

[backends]
        [backends.backend1]
                [backends.backend1.loadbalancer.stickiness]
                [backends.backend1.servers.server1]
                        url="https://10.0.0.1"
                [backends.backend1.servers.server2]
                        url="https://10.0.0.2"
```

Note: Traefik has its own concept of "frontends" and "backends" which do not correspond to Gravwell frontends. In this config file, a Traefik "backend" points to Gravwell frontend servers.

Modify the configuration file as needed; note especially the `[backends]` section where we define two servers, 10.0.0.1 and 10.0.0.2. These should be the IPs of your frontends; add more server sections if you have more than two Gravwell frontends.

If your frontend servers use valid SSL certificates, you can remove the line `InsecureSkipVerify = true`.

After modifying the config file, run the following command to start the reverse proxy:

```
./traefik -c traefik.toml -file traefik.toml
```

You can then use a web browser to access the Traefik server's hostname or IP; Traefik will select one of the Gravwell frontends and direct your traffic there, setting a cookie in the browser to ensure all traffic for that session goes to the same server.
