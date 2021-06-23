# The Gravwell Load Balancer

To make setting up your environment as easy as possible, Gravwell provides a custom load balancer specifically designed for use with Gravwell webservers. It can automatically discover Gravwell webservers, meaning you don't need to reconfigure the loadbalancer every time you add or remove a webserver--and if a webserver goes down, it will automatically direct users to another server.

## Load Balancer Architecture

The load balancer is an HTTP(S) proxy which automatically directs clients to one of the Gravwell webservers. It sets a cookie on the user's browser to maintain a level of "stickiness", so one session's requests all go to the same webserver.

The load balancer discovers Gravwell webservers by communicating with the Gravwell datastore, which provides a list of active webservers. See [the distributed webserver documentation](frontend.md) for more info about the datastore.

Once installed and configured, users should access Gravwell through the load balancer. We recommend setting a hostname such as `gravwell.example.org` to point at the load balancer while naming webservers something like `web1.example.org`; encourage users to visit `gravwell.example.org` instead of accessing the webservers directly.  Users do not need direct access to Gravwell webservers when using the load balancer--the webservers may be privately addressed or otherwise inaccessible to the wider world.

## Deploying the Load Balancer

The load balancer component is distributed through all the same channels as the main Gravwell installer:

* Self-extracting shell installer is available [on the downloads page](https://docs.gravwell.io/#!quickstart/downloads.md)
* In the Debian and RedHat repositories as a package named `gravwell-loadbalancer`.
* On DockerHub as [gravwell/loadbalancer](https://hub.docker.com/r/gravwell/loadbalancer)

The Debian installer will prompt for basic configuration options and should need no further setup after you've installed. For other installation methods, you will need to edit `/opt/gravwell/etc/loadbalancer.conf` as detailed below. If you are using Docker, you can also configure the container purely through environment variables, as described in the "Docker Environment Variables" section below.

## Config File Settings

Configuration is managed in `/opt/gravwell/etc/loadbalancer.conf`. Here is a simple configuration for an HTTP-only configuration:

```
[Global]
Disable-HTTP-Redirector=true
Insecure-Disable-HTTPS=true
Update-Interval=10
Session-Timeout=10
Log-Dir=/opt/gravwell/log/loadbalancer
Log-Level=info
Enable-Access-Log=true

Control-Secret=ControlSecrets
Datastore=datastore.example.org
Datastore-Insecure-Disable-TLS=true
```

The Disable-HTTP-Redirector and Insecure-Disable-HTTPS settings make the load balancer listen for incoming connections on HTTP only. At the bottom of the file, the Datastore parameter tells the load balancer where the Gravwell datastore may be found; the Control-Secret parameter gives the authentication token for communicating with the datastore, while Datastore-Insecure-Disable-TLS sets us to talk to the datastore over an unencrypted connection.

If we want to use HTTPS instead, we need to provide the load balancer with a valid TLS certificate & key pair (see [the TLS documentation](#!configuration/certificates.md) for more information on setting up TLS in Gravwell). Here's an example configuration that listens on HTTPS and communicates with the datastore over an encrypted channel:

```
[Global]
Key-File=/opt/gravwell/etc/key.pem
Certificate-File=/opt/gravwell/etc/cert.pem
Update-Interval=10
Session-Timeout=10
Log-Dir=/opt/gravwell/log/loadbalancer
Log-Level=info
Enable-Access-Log=true

#Control-Secret=ControlSecrets
#Datastore=172.19.0.2
#Datastore-Insecure-Disable-TLS=true
#Datastore-Insecure-Skip-TLS-Verify=true
```

### Global Configuration Parameters

This section lists every configuration parameter available in the `[Global]` section of the configuration file.

**Disable-HTTP-Redirector**
Default Value:	`false`
Description:	When this parameter is set to `false`, the load balancer will redirect incoming HTTP connections to the HTTPS port. Set this to `true` if you are not using TLS.

**Insecure-Disable-HTTPS**
Default Value:	`false`
Description:	When this parameter is set to `true`, the load balancer will listen for incoming HTTP connections instead of HTTPS. If this is set to `false`, you must also set the `Key-File` and `Certificate-File` parameters.

**Web-Port**
Default Value:	443 (80 if `Insecure-Disable-HTTPS=true` is set)
Description:	This parameter sets the port on which to listen for incoming connections. The default is typically acceptable.

**Bind-Addr**
Default Value:	`0.0.0.0`
Description:	This parameter sets the IP address on which the load balancer will listen for incoming connections. By default, it listens on all interfaces.

**Key-File**
Default Value:	(empty)
Description:	Sets the location of a TLS secret key. The key must correspond to a valid certificate for the load balancer's hostname.

**Certificate-File**
Default Value:	(empty)
Description:	Sets the location of a TLS certificate file. The certificate must be valid for the load balancer's hostname.

**Insecure-Skip-TLS-Verify**
Default Value:	`false`
Description:	If this parameter is set to `true`, the load balancer will not verify TLS certificates on Gravwell webservers when proxying connections. This setting is ignored if `Insecure-Disable-HTTPS=true` is set.

**Update-Interval**
Default Value:	30
Description:	This parameter sets, in seconds, how frequently the load balancer should check for new or failed webservers.

**Session-Timeout**
Default Value:	10
Description:	This parameter sets, in minutes, how long each load balancer session lasts. Note that users will not notice when these sessions expire; Gravwell webservers synchronize their user login sessions, so even though the load balancer starts sending requests to a different webserver, they will still work. The default value should be fine.

**Datastore**
Default Value:	(empty)
Description:	This parameter points to the Gravwell datastore component. It should be a hostname or IP address, e.g. `Datastore=datastore.example.org` or `Datastore=192.168.0.11`. If the datastore is listening on a non-standard port (instead of 9405), you may include the port: `Datastore=datastore.example.org:9999`.

**Control-Secret**
Default Value:	`ControlSecrets`
Description:	This parameter gives the authentication token for communicating with the datastore. You will almost certainly need to change the default.

**Disable-Datastore**
Default Value:	`false`
Description:	If set to true, the load balancer will not attempt to communicate with any datastore. Instead, it will use the webservers listed in the `[Override]` config stanzas (see following section).

**Datastore-Insecure-Disable-TLS*
Default Value:	`false`
Description:	If set to true, the load balancer will connect to the datastore via an unencrypted channel.

**Datastore-Insecure-Skip-TLS-Verify*
Default Value:	`false`
Description:	If set to true, the load balancer will not verify the datastore's TLS certificates.

**Log-Dir**
Default Value:	`/opt/gravwell/log/loadbalancer`
Description:	Sets the directory where the loadbalancer will keep its log files.

**Log-Level**
Default Value:	`error`
Description:	Sets the severity level for logs. By default, only errors are logged. Valid levels are, in order of decreasing severity: `error`, `warn`, `info`. Setting it to `off`, `none`, or `disabled` will turn off logging completely.

**Enable-Access-Log**
Default Value:	`false`
Description:	If this parameter is set to true, the load balancer will log every URL requested by clients and the response code from the webserver.

### Override Stanzas

Most systems will use the datastore to automatically get a list of Gravwell webservers. However, sometimes the load balancer cannot communicate directly with the datastore--frequently this is a result of corporate network security rules. In this case, you may add `[Override]` configuration blocks at the end of the configuration file to manually list webservers:

```
[Global]
Disable-HTTP-Redirector=true
Insecure-Disable-HTTPS=true
Disable-Datastore=true

[Overrides "example1"]
	Webserver=172.19.0.100
	Insecure-Disable-HTTPS=true

[Overrides "example2"]
	Webserver=172.19.0.101
	Insecure-Disable-HTTPS=true
```

In the example above, we manually specify Override stanzas for two webservers.

Each override may use the following parameters:

**Webserver**
Default Value:	(empty)
Description:	This is the hostname/IP and (optional) port for the webserver, e.g. `Webserver=192.168.0.1:8080` or `Webserver=web1.example.org`. If no port is specified, the appropriate default will be set (80 if HTTPS is disabled, 443 otherwise).

**Insecure-Disable-HTTPS**
Default Value:	`false`
Description:	If set to true, the load balancer will communicate with the webserver over insecure HTTP.

**Insecure-Skip-TLS-Verify**
Default Value:	`false`
Description:	If set to true, the load balancer will not validate the webserver's TLS certificates.

## Docker Environment Variables

When deploying in Docker, it is frequently easier to configure components via Docker environment variables instead of modifying a config file. All basic parameters of the load balancer can be configured through environment variables:

* `GRAVWELL_INSECURE_DISABLE_HTTPS` - Equivalent to the `Insecure-Disable-HTTPS` config parameter
* `GRAVWELL_INSECURE_SKIP_TLS_VERIFY` - Equivalent to the `Insecure-Skip-TLS-Verify` config parameter
* `GRAVWELL_DISABLE_HTTP_REDIRECTOR` - Equivalent to the `Disable-HTTP-Redirector` config parameter
* `GRAVWELL_WEB_PORT` - Equivalent to the `Web-Port` config parameter
* `GRAVWELL_BIND_ADDR` - Equivalent to the `Bind-Addr` config parameter
* `GRAVWELL_DATASTORE` - Equivalent to the `Datastore` config parameter
* `GRAVWELL_CONTROL_SECRET` - Equivalent to the `Control-Secret` config parameter
* `GRAVWELL_DATASTORE_INSECURE_DISABLE_TLS` - Equivalent to the `Datastore-Insecure-Disable-TLS` config parameter
* `GRAVWELL_DATASTORE_INSECURE_SKIP_TLS_VERIFY` - Equivalent to the `Datastore-Insecure-Skip-TLS-Verify` config parameter
* `GRAVWELL_LOG_LEVEL` - Equivalent to the `Log-Level` config parameter
* `GRAVWELL_LOG_DIR` - Equivalent to the `Log-Dir` config parameter
* `GRAVWELL_LOG_ENABLED_ACCESS_LOG` - Equivalent to the `Enable-Access-Log` config parameter

Thus for example you might invoke:

```
docker create --name loadbalancer \
	-e GRAVWELL_CONTROL_SECRET=ControlSecrets \
	-e GRAVWELL_DATASTORE=datastore.example.org \
	-e GRAVWELL_INSECURE_DISABLE_HTTPS=TRUE \
	-e GRAVWELL_LOG_DIR=/tmp -e GRAVWELL_LOG_LEVEL=INFO \
	-e GRAVWELL_LOG_ENABLE_ACCESS_LOG=TRUE \
	-e GRAVWELL_DATASTORE_INSECURE_DISABLE_TLS=TRUE \
	gravwell/loadbalancer
```
