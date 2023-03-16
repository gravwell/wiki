# Source Router Preprocessor

The source router preprocessor can route entries to different tags based on the SRC field of the entry. Typically the SRC field will be the IP address of the entry's origination point, e.g. the system which created the syslog message sent to Simple Relay.

The source router preprocessor Type is `srcrouter`.

## Supported Options

* `Route` (string, optional): `Route` defines a mapping of SRC field value to tag, separated by a colon. For instance, `Route=192.168.0.1:server-logs` will send all entries with SRC=192.168.0.1 to the "server-logs" tag. Multiple `Route` parameters can be specified. Leaving the tag blank (`Route=192.168.0.1:`) tells the preprocessor to drop all matching entries instead.
* `Route-File` (string, optional): `Route-File` should contain a path to a file containing newline-separated route specifications, e.g. `192.168.0.1:server-logs`.
* `Drop-Misses` (boolean, optional): By default, entries which do not match any of the defined routes will be passed through unmodified. Setting `Drop-Misses` to true will instead drop any entries which do not explicitly match a route definition.

At least one `Route` definition is required, unless `Route-File` is used.

A route can be either a single IP address or a properly formed CIDR specification. Both IPv4 and IPv6 are supported.

## Example: Inline Route Definitions

The snippet below shows part of a Simple Relay ingester configuration that uses the source router preprocessor with routes defined inline. Recall that Simple Relay applies a SRC field corresponding to the remote IP which has connected. Entries originating from 10.0.0.1 will be tagged "internal-syslog", entries originating from 7.82.33.4 will be tagged "external-syslog", and all other entries will retain the default tag "syslog". Any entries with SRC=3.3.3.3 will be dropped. There are also two IPv6 routes defined.

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=srcroute

[preprocessor "srcroute"]
        Type = srcrouter
        Route=10.0.0.0/24:internal-syslog
        Route=7.82.33.4:external-syslog
        Route=3.3.3.3:
        Route=DEAD::BEEF:external-syslog
        Route=FEED:FEBE::0/64:external-syslog
```

## Example: File-based Definitions

The snippet below shows part of a Simple Relay ingester configuration that uses the source router preprocessor with routes defined in a file:

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=srcroute

[preprocessor "srcroute"]
        Type = srcrouter
        Route-File=/opt/gravwell/etc/syslog-routes
```

The following is written to `/opt/gravwell/etc/syslog-routes`:

```
10.0.0.0/24:internal-syslog
7.82.33.4:external-syslog
3.3.3.3:
```


