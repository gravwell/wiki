# HTTP Ingester Setup

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
:::

## HTTP Ingester Configuration

Many integration guides send data to the Gravwell HTTP Ingester. Set the ingester up once, as described on this page, and every integration that uses the HTTP Ingester can share it. Each integration guide only requires a listener, so you are not required to install a separate HTTP Ingester for each individual integration.

1. Install the HTTP Ingester by following the Installation section of [HTTP Ingester](/ingesters/http).
2. Choose the address and port the ingester listens on (`Bind`). Every integration can use this one port, and the integration guides refer to it as `<PORT>`.
3. If the sending systems can reach the ingester over an untrusted network, enable HTTPS by providing a certificate and key in the `[Global]` block, or put the ingester behind an HTTPS reverse proxy. By default the ingester serves cleartext HTTP.
4. Optionally, enable a health check URL for load balancers and monitoring.
5. Raise `Max-Body` if an integration sends bodies larger than the 4 MB default.

These settings apply to the whole ingester, so they belong in the base configuration file and not in an integration's overlay file. The ingester cannot reload `[Global]` changes while it is running, so restart it after you change them.

## Gravwell Configuration

### Gravwell Ingester Configuration: HTTP Ingester

**Sample HTTP Ingester config:**  
Edit: `/opt/gravwell/etc/gravwell_http_ingester.conf`
```ini
[Global]
    Bind=":8080"
    #TLS-Certificate-File=/opt/gravwell/etc/cert.pem
    #TLS-Key-File=/opt/gravwell/etc/key.pem
    Max-Body=4194304 #the default, in bytes (4 MB)
    #Health-Check-URL="/health/check"
```

The base file also holds the connection settings for your Gravwell backend, which the [HTTP Ingester](/ingesters/http) page describes. The settings above are the ones integration guides depend on:

| Setting | What it does |
|---|---|
| `Bind` | Address and port the HTTP server listens on. Required |
| `TLS-Certificate-File`, `TLS-Key-File` | Certificate and key (PEM) that turn on HTTPS |
| `Max-Body` | Largest request body the ingester accepts, in bytes. Default 4 MB |
| `Max-Connections`, `Max-Concurrent-Requests` | Limits on connections and in-flight requests. Defaults 10240 and 16384; `-1` removes a limit |
| `Health-Check-URL` | Unauthenticated URL that always returns `200 OK`, for load balancers and monitoring |

### Adding a Listener for an Integration

Integration guides add their listeners as overlay files. The ingester merges every `.conf` file in `/opt/gravwell/etc/gravwell_http_ingester.conf.d/` into its base configuration.

Each integration guide follows the same conventions:

* **File:** one overlay file per integration, named after the guide, for example `/opt/gravwell/etc/gravwell_http_ingester.conf.d/auth0.conf`.
* **URL:** a unique path per integration, named after the guide: `/auth0`, `/fluentd`, `/teleport`. A guide with several listeners uses `/<guide>/<purpose>`, for example `/ubiquiti/network-alarms`. Unique paths keep each source on its own tag and let you retire one integration without touching another.
* **Tag:** `Tag-Name` is the guide name, or `<guide>-<purpose>` when a guide has several listeners.
* **Authentication:** give each integration its own token and set `AuthType`, `TokenName` and `TokenValue` on the listener. Do not reuse the Gravwell ingest secret. A `preshared-token` listener expects the token in the `Authorization` header. A `preshared-parameter` listener expects it as a URL query parameter, so proxies and HTTP logs can capture it.
* **No `[Global]` blocks:** an integration's overlay file contains listeners only. A guide that needs a global setting, such as a larger `Max-Body` or a health check URL, states the requirement and links to this page.

**Sample listener:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/<name>.conf`
```ini
[Listener "<name>"]
    URL="/<name>"
    Tag-Name=<name>
    AuthType="preshared-token"
    TokenName="Bearer"
    TokenValue="your-token"
```

The URL that the sending system uses is `http://<HTTP-INGESTER-HOST>:<PORT>/<name>`, or `https://` when you enable HTTPS.

### Splunk HEC Listeners

Systems that stream to Splunk send to a fixed HEC path, `/services/collector`, and often cannot change it. Keep that path on the `HEC-Compatible-Listener` and separate sources another way. A `Routed-Token-Value` routes data to a tag by the token the sender uses, so several sources can share the HEC URL. `Tag-Match` routes by the sourcetype in the event, and a `tag` URL parameter sets a default tag for a request.

### OpenTelemetry Listeners

The `OpenTelemetry-Logs-Listener` and `OpenTelemetry-Metrics-Listener` default to `/v1/logs` and `/v1/metrics`. The HTTP Ingester has no listener for traces. The [OpenTelemetry](/integrations/generic/opentelemetry) guide uses these default paths.

The `URL` setting accepts any path, so other integrations that send OpenTelemetry data give each source its own prefix, for example `URL="/<name>/v1/metrics"`. An OpenTelemetry Collector `otlphttp` exporter appends `/v1/logs` or `/v1/metrics` to its `endpoint`, so an endpoint of `http://<HTTP-INGESTER-HOST>:<PORT>/<name>` reaches that listener. Without distinct paths, every source sending OpenTelemetry data to the same URL ends up on the same tag.

### Applying Changes

The HTTP Ingester reloads `Listener`, HEC and Amazon Firehose blocks without restarting when it receives the `SIGHUP` signal:

```bash
sudo killall -SIGHUP gravwell_http_ingester
```

Integration guides use the restart command, which always works, and a restart is required for `[Global]` changes.

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```

## Additional Resources
* [HTTP Ingester](/ingesters/http)
* [OpenTelemetry Collector otlphttp exporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlphttpexporter)
