# Keycloak

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
:::

## Keycloak Configuration

Keycloak has a native syslog log handler. See [Keycloak Server Guide: Logging](https://www.keycloak.org/server/logging).

Start (or reconfigure) Keycloak with the syslog handler enabled, pointing at your Gravwell Simple Relay ingester:

```bash
bin/kc.sh start --log="console,file,syslog" \
  --log-syslog-endpoint=<gravwell-host>:601 \
  --log-syslog-protocol=tcp \
  --log-syslog-type=rfc5424 \
  --log-syslog-app-name=keycloak \
  --log-syslog-output=json \
  --spi-events-listener--jboss-logging--success-level=info
```

* `log-syslog-endpoint` defaults to `localhost:514` if unset. Replace `<gravwell-host>` with your Simple Relay address. Port `601` matches the listener in the sample below, so change it if you bind a different port.
* `log-syslog-protocol` can be `tcp`, `udp`, or `ssl-tcp` (default `tcp`).
* `log-syslog-type` can be `rfc5424` (default) or `rfc3164`.
* `log-syslog-counting-framing` defaults to `protocol-dependent`, which means Keycloak prefixes each message with its length when it sends over `tcp` or `ssl-tcp`. This is RFC 6587 octet counting, so the sample listener below uses `Reader-Type=rfc6587`. If you set `--log-syslog-counting-framing=false`, use `Reader-Type=rfc5424` instead.
* `log-syslog-output=json` frames the syslog payload as a JSON object rather than Keycloak's default text pattern; this is optional but makes downstream parsing easier.
* `spi-events-listener--jboss-logging--success-level=info` raises successful audit events from `debug` to `info`. Without it, the server's default `info` log level filters them out and only error events reach Gravwell.

```{note}
Keycloak's built-in `JBossLoggingEventListenerProvider` writes user and admin audit events (logins, admin actions) to the `org.keycloak.events` logging category by default, at `debug` level for successful events and `warn` level for errors. The command above raises the successful level to `info`. Enable events per realm first, under **Realm Settings > Events > User events** and **Admin events**. Otherwise only server and system log lines reach the syslog handler.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/keycloak-well.conf`
```ini
[Storage-Well "keycloak"]
    Location=/opt/gravwell/storage/keycloak
    Tags=keycloak*
```

### Gravwell Ingester Configuration: Simple Relay

**Sample Keycloak config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/keycloak.conf`
```ini
[Listener "keycloaktcp"]
    Bind-String="tcp://0.0.0.0:601" #octet-counted syslog over TCP, Keycloak's default
    Reader-Type=rfc6587
    Tag-Name=keycloak
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```

## Additional Resources
* [Keycloak Server Guide: Logging](https://www.keycloak.org/server/logging)
* [Keycloak Admin REST API](https://www.keycloak.org/docs-api/latest/rest-api/index.html)
* [Keycloak Server Administration Guide](https://www.keycloak.org/docs/latest/server_admin/)
