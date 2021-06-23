# HTTP

The HTTP ingester sets up HTTP listeners on one or more paths. If an HTTP request is sent to one of those paths, the request's Body will be ingested as a single entry.

This is an extremely convenient method for scriptable data ingest, since the `curl` command makes it easy to do a POST request using standard input as the body.

## Basic Configuration

The HTTP ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the HTTP ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## Listener Examples

In addition to the universal configuration parameters used by all ingesters, the HTTP POST ingester has two additional global configuration parameters that control the behavior of the embedded webserver.  The first configuration parameter is the `Bind` option, which specifies the interface and port that the webserver listens on.  The second is the `Max-Body` parameter, which controls how large of a POST the webserver will allow.  The Max-Body parameter is a good safety net to prevent rogue processes from attempting to upload very large files into your Gravwell instance as a single entry.  Gravwell can support up to 1GB as a single entry, but we wouldn't recommend it.

Multiple "Listener" definitions can be defined allowing specific URLs to send entries to specific tags.  In the example configuration we define two listeners which accept data from a weather IOT device and a smart thermostat.

```
 Example using basic authentication
[Listener "basicAuthExample"]
	URL="/basic"
	Tag-Name=basic
	AuthType=basic
	Username=user1
	Password=pass1

[Listener "jwtAuthExample"]
	URL="/jwt"
	Tag-Name=jwt
	AuthType=jwt
	LoginURL="/jwt/login"
	Username=user1
	Password=pass1
	Method=PUT # alternate method, data is still expected in the body of the request

[Listener "cookieAuthExample"]
	URL="/cookie"
	Tag-Name=cookie
	AuthType=cookie
	LoginURL="/cookie/login"
	Username=user1
	Password=pass1
	Method=PUT # alternate method, data is still expected in the body of the request

[Listener "presharedTokenAuthExample"]
	URL="/preshared/token"
	Tag-Name=pretoken
	AuthType="preshared-token"
	TokenName=Gravwell
	TokenValue=Secret

[Listener "presharedTokenAuthExample"]
	URL="/preshared/param"
	Tag-Name=preparam
	AuthType="preshared-parameter"
	TokenName=Gravwell
	TokenValue=Secret
```

## Splunk HEC Compatibility

The HTTP ingester supports a listener block that is API compatible with the Splunk HTTP Event Collector.  This special listener block enables a simplified configuration so that any endpoint that can send data to the Splunk HEC can also send to the Gravwell HTTP Ingester.  The HEC compatible configuration block looks like so:

```
[HEC-Compatible-Listener "testing"]
	URL="/services/collector/event"
	TokenValue="thisisyourtoken"
	Tag-Name=HECStuff

```

The `HEC-Compatible-Listener` block requires the `TokenValue` and `Tag-Name` configuration items, if the `URL` configuration item is omitted it will default to `/services/collector/event`.

Both `Listener` and `HEC-Compatible-Listener` configuration blocks can be specified on the same HTTP ingester.

## Health Checks

Some systems (such as AWS load balancers) require an unauthenticated URL that can be probed and interpreted as "proof of life".  The HTTP ingester can be configured to provide an a URL which when accessed with any method, body, and/or query parameters will always return a 200 OK.  To enable this health check endpoint add the `Health-Check-URL` stanza to the Global configuration block.

Here is a minimal example configuration snippet with the health check URL `/logbot/are/you/alive`:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO #options are OFF INFO WARN ERROR
Bind=":8080"
Max-Body=4096000 #about 4MB
Log-File="/opt/gravwell/log/http_ingester.log"
Health-Check-URL="/logbot/are/you/alive"

```

## Installation

If you're using the Gravwell Debian repository, installation is just a single apt command:

```
apt-get install gravwell-http-ingester
```

Otherwise, download the installer from the [Downloads page](#!quickstart/downloads.md). Using a terminal on the Gravwell server, issue the following command as a superuser (e.g. via the `sudo` command) to install the ingester:

```
root@gravserver ~ # bash gravwell_http_ingester_installer_3.0.0.sh
```

If the Gravwell services are present on the same machine, the installation script will automatically extract and configure the `Ingest-Auth` parameter and set it appropriately. However, if your ingester is not resident on the same machine as a pre-existing Gravwell backend, the installer will prompt for the authentication token and the IP address of the Gravwell indexer. You can set these values during installation or leave them blank and modify the configuration file in `/opt/gravwell/etc/gravwell_http_ingester.conf` manually.

## Configuring HTTPS

By default the HTTP Ingester runs a cleartext HTTP server, but it can be configured to run an HTTPS server using x509 TLS certificates.  To configure the HTTP Ingester as an HTTPS server provide a certificate and key PEM files in the Global configuration space using the `TLS-Certificate-File` and `TLS-Key-File` parameters.

An example global configuration with HTTPS enabled might look like the following:

```
[Global]
	TLS-Certificate-File=/opt/gravwell/etc/cert.pem
	TLS-Key-File=/opt/gravwell/etc/key.pem
```

### Listener Authentication

Each HTTP Ingester listener can be configured to enforce authentication.  The supported authentication methods are:

* none
* basic
* jwt
* cookie
* preshared-token
* preshared-parameter

When specifying an authentication system other than none credentials must be provided.  The `jwt` and `cookie` and cookie authentication systems require a username and password while the `preshared-token` and `preshared-parameter` must provide a token value and optional token name.

Warning: Like any other webpage, authentication is NOT SECURE over cleartext connections and attackers that can sniff traffic can capture tokens and cookies.

### No Authentication

The default authentication method is none, allowing anyone that can reach the ingester to push entries.  The `basic` authentication mechanism uses HTTP Basic authentication, where a username and password is base64 encoded and sent with every request.

Here is an example listener using the basic authentication system:

```
[Listener "basicauth"]
	URL="/basic/data"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

An example curl command to send an entry with basic authentication might look like:

```
curl -d "only i can say hi" --user secretuser:secretpassword -X POST http://10.0.0.1:8080/basic/data
```

### JWT Authentication

The JWT authentication system uses a cryptographically signed token for authentication.  When using jwt authentication you must specify an Login URL where clients will authenticate and receive a token which must then be sent with each request.  The jwt tokens expire after 48 hours.  Authentication is performed by sending a `POST` request to the login URL with the `username` and `password` form fields populated.

Authenticating with the HTTP ingester using jwt authentication is a two step process and requires an additional configuration parameter.  Here is an example configuration:

```
[Listener "jwtauth"]
	URL="/jwt/data"
	LoginURL="/jwt/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

Sending entries requires that endpoints first authenticate to obtain a token, the token can then be reused for up to 48 hours.  If a request receives a 401 response, clients should re-authenticate.  Here is an example using curl to authenticate and then push data.

```
x=$(curl -X POST -d "username=user1&password=pass1" http://127.0.0.1:8080/jwt/login) #grab the token and stuff it into a variable
curl -X POST -H "Authorization: Bearer $x" -d "this is a test using JWT auth" http://127.0.0.1:8080/jwt/data #send the request with the token
```

### Cookie Authentication

The cookie authentication system is virtually identical to JWT authentication other than the method of controlling state.  Listeners that use cookie authentication require that a client login with a username and password to acquire a cookie which is set by the login page.  Subsequent requests to the ingest URL must provide the cookie in each request.

Here is an example configuration block:

```
[Listener "cookieauth"]
	URL="/cookie/data"
	LoginURL="/cookie/login"
	Tag-Name=stuff
	AuthType=basic
	Username=secretuser
	Password=secretpassword
```

An example set of curl commands that login and retrieve the cookie before ingesting some data might look like:

```
curl -c /tmp/cookie.txt -d "username=user1&password=pass1" localhost:8080/cookie/login
curl -X POST -c /tmp/cookie.txt -b /tmp/cookie.txt -d "this is a cookie data" localhost:8080/cookie/data
```

### Preshared Token

The Preshared token authentication mechanism uses a preshared secret rather than a login mechanism.  The preshared secret is expected to be sent with each request in an Authorization header.  Many HTTP frameworks expect this type of ingest, such as the Splunk HEC and supporting AWS Kinesis and Lambda infrastructure.  Using a preshared token listener we can define a capture system that is a plugin replacement for Splunk HEC.

Note: If you do not define a `TokenName` value, the default value of `Bearer` will be used.

An example configuration which defines a preshared token:

```
[Listener "presharedtoken"]
	URL="/preshared/token/data"
	Tag-name=token
	AuthType="preshared-token"
	TokenName=foo
	TokenValue=barbaz
```

An example curl command the sends data using the preshared secret:

```
curl -X POST -H "Authorization: foo barbaz" -d "this is a preshared token" localhost:8080/preshared/token/data
```

### Preshared Parameter

The Preshared Parameter authentication mechanism uses a preshared secret that is provided as a query parameter.  The `preshared-parameter` system can be useful when scripting or using data producers that typically do not support authentication by embedding the authentication token into the URL.

Note: Embedding the authentication token into the URL means the proxies and HTTP logging infrastructure may capture and log authentication tokens.

An example configuration which defines a preshared parameter:

```
[Listener "presharedtoken"]
	URL="/preshared/parameter/data"
	Tag-name=token
	AuthType="preshared-parameter"
	TokenName=foo
	TokenValue=barbaz
```

An example curl command the sends data using the preshared secret:

```
curl -X POST -d "this is a preshared parameter" localhost:8080/preshared/parameter/data?foo=barbaz
```

## Listener Methods

The HTTP Ingester can be configured to use virtually any method, but data is always expected to be in the body of the request.

For example, here is a Listener configuration that expects the PUT method:

```
[Listener "test"]
	URL="/data"
	Method=PUT
	Tag-Name=stuff
```

The corresponding curl command would be:

```
curl -X PUT -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```

The HTTP Ingester can go out of spec on methods, accepting almost any ASCII string that does not contain special characters.

```
[Listener "test"]
	URL="/data"
	Method=SUPER_SECRET_METHOD
	Tag-Name=stuff
```

```
curl -X SUPER_SECRET_METHOD -d "this is a test 2 using basic auth" http://127.0.0.1:8080/data
```
