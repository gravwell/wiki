# Keycloak

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, • [File Follower](/ingesters/file_follow) <br /> • [Simple Relay](/ingesters/simple_relay)
:::

## Keycloak Configuration

Keycloak writes three kinds of entries to its log: user events (sign-ins, token requests, account changes), admin events (changes made in the admin console or through the Admin API), and server messages. Gravwell can read Keycloak's log file (Option 1, recommended) or receive the log through Keycloak's built-in syslog handler (Option 2). See [Keycloak Server Guide: Logging](https://www.keycloak.org/server/logging).

Whichever option you use, Keycloak needs these settings:

* `--spi-events-listener-jboss-logging-success-level=info` (environment variable `KC_SPI_EVENTS_LISTENER_JBOSS_LOGGING_SUCCESS_LEVEL=info`). Keycloak logs successful events at `debug` by default, so without this only failed events reach Gravwell.
* JSON output for the handler that Gravwell reads (`--log-file-output=json` or `--log-syslog-output=json`), using the default JSON format, not `ecs`.
* `--proxy-headers=xforwarded` or `--proxy-headers=forwarded` if Keycloak is behind a reverse proxy or load balancer. Without it, every event's IP address is the proxy's address. See the Reverse Proxy section below.

```{note}
Events reach the log as long as `jboss-logging` stays in each realm's **Event listeners** list under **Realm Settings > Events**, which is the default. The **Save events** and **Admin events** toggles on that page only store a copy in Keycloak's database for the admin console. Gravwell doesn't need them.
```

### [Option 1] Keycloak Log File to File Follower (recommended)

Start (or reconfigure) Keycloak with a JSON log file:

```bash
bin/kc.sh start --log="file" \
  --log-file-output=json \
  --spi-events-listener-jboss-logging-success-level=info \
  --proxy-headers=xforwarded
```

* Keycloak writes the file to `data/log/keycloak.log` under its installation directory. Use `--log-file=<path>` to change it.
* Don't include `console` in `--log` when Keycloak runs under systemd. Console output goes to the host's journal and syslog, so a host that already sends syslog to Gravwell would deliver every Keycloak line twice, in two tags.
* The file holds events while Gravwell, the network or File Follower is unavailable, and File Follower catches up when it's back. Keycloak renames its log file each time it restarts (`keycloak.log` becomes `keycloak.log.1`, and so on), and the File Follower configuration below reads only the live file. If Keycloak restarts while File Follower is stopped, lines written to the old file before the restart may be missed, so avoid restarting Keycloak while File Follower is down.

Install File Follower on your Keycloak host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

### [Option 2] Keycloak Syslog Handler to Simple Relay

Use this option if you can't install an ingester on the Keycloak host. Start (or reconfigure) Keycloak with the syslog handler pointing at your Gravwell Simple Relay ingester:

```bash
bin/kc.sh start --log="file,syslog" \
  --log-syslog-endpoint=<gravwell-host>:15060 \
  --log-syslog-protocol=tcp \
  --log-syslog-type=rfc5424 \
  --log-syslog-output=json \
  --log-syslog-max-length=65536 \
  --spi-events-listener-jboss-logging-success-level=info \
  --proxy-headers=xforwarded
```

```{warning}
Keycloak's syslog handler doesn't queue messages. Events logged while Simple Relay is unreachable (restarts, network outages) are never delivered, and the first message after the connection returns can be lost too. Keycloak reconnects on its own but reports no error. The local `file` handler in the command above keeps a copy that you can import afterwards.
```

* As in Option 1, leave out `console` when Keycloak runs under systemd. `file` keeps a local copy and is optional.
* Port `15060` matches the listener in the Gravwell Configuration section. Use a dedicated port: a stock Simple Relay install already has a `syslogtcp` listener on tcp/601, tagged `syslog`.
* `--log-syslog-max-length` defaults to 2048 bytes, header included, and longer messages are cut off. Events fit, but stack traces can run past 20 KB and would arrive as broken JSON. 65536 keeps them whole. Lower it if you'd rather trim large messages.
* Use `--log-syslog-protocol=tcp`, as shown. Keycloak also accepts `ssl-tcp` (the listener then needs a TLS `Bind-String`) and `udp`. Don't use `udp` with the maximum length above: a UDP message can carry at most about 65,507 bytes, and messages that don't arrive are dropped without notice.
* Message framing depends on your Keycloak version, and the Simple Relay `Reader-Type` must match it:
  * **Keycloak 26.3 and later:** `--log-syslog-counting-framing` defaults to `protocol-dependent`, which prefixes each message sent over `tcp` or `ssl-tcp` with its length (RFC 6587 octet counting). Use `Reader-Type=rfc6587`. If you set `--log-syslog-counting-framing=false`, use `Reader-Type=rfc5424`.
  * **Keycloak 26.2 and earlier:** the option doesn't exist and messages are newline-delimited. Use `Reader-Type=rfc5424`.

  A mismatched `Reader-Type` fails silently: the connection stays up but nothing is indexed. After setup, confirm that data is arriving in the `keycloak` tag.

### Running Keycloak in Docker

Use the same Keycloak settings as above, passed as environment variables, so Keycloak writes or sends its log for Gravwell itself. Keep `console` in `KC_LOG` so `docker logs` still works; it doesn't reach Gravwell. Set `--hostname`: Keycloak records the hostname in every entry, and without it the hostname is the container ID, which changes each time the container is re-created.

Option 1 (log file on the host):

```bash
docker run ... --hostname keycloak-1 \
  -v /var/log/keycloak:/opt/keycloak/data/log \
  -e KC_LOG=console,file \
  -e KC_LOG_FILE_OUTPUT=json \
  -e KC_SPI_EVENTS_LISTENER_JBOSS_LOGGING_SUCCESS_LEVEL=info \
  -e KC_PROXY_HEADERS=xforwarded \
  quay.io/keycloak/keycloak:<version> start ...
```

Point File Follower's `Base-Directory` at `/var/log/keycloak`. The Keycloak image runs as user ID 1000, which must be able to write that directory (for example `chown 1000 /var/log/keycloak`).

Option 2 (syslog):

```bash
docker run ... --hostname keycloak-1 \
  -e KC_LOG=console,syslog \
  -e KC_LOG_CONSOLE_OUTPUT=json \
  -e KC_LOG_SYSLOG_ENDPOINT=<gravwell-host>:15060 \
  -e KC_LOG_SYSLOG_PROTOCOL=tcp \
  -e KC_LOG_SYSLOG_TYPE=rfc5424 \
  -e KC_LOG_SYSLOG_OUTPUT=json \
  -e KC_LOG_SYSLOG_MAX_LENGTH=65536 \
  -e KC_SPI_EVENTS_LISTENER_JBOSS_LOGGING_SUCCESS_LEVEL=info \
  -e KC_PROXY_HEADERS=xforwarded \
  quay.io/keycloak/keycloak:<version> start ...
```

```{note}
Don't collect Keycloak's logs from Docker's own log files (`/var/lib/docker/containers/*/*-json.log`). Docker's default `json-file` driver stores each Keycloak line as an escaped string inside Docker's own JSON (`{"log":"{\"timestamp\"...`), which the Keycloak kit's queries don't read. Docker also splits lines longer than 16 KB into separate entries, so long stack traces, and sign-in attempts with very long usernames, arrive broken.
```

### Reverse Proxy

With `--proxy-headers=xforwarded`, Keycloak takes each event's IP address from the `X-Forwarded-For` header. Keycloak's [reverse proxy guide](https://www.keycloak.org/server/reverseproxy) requires the proxy to **overwrite** that header: if the proxy appends to a client-supplied value, a client can choose the IP address that Keycloak logs. For nginx, set this in the location block that proxies to Keycloak:

```nginx
proxy_set_header X-Forwarded-For $remote_addr;
```

### Log Format

With JSON output, each line Keycloak writes is a JSON object with fields such as `timestamp`, `level`, `loggerName`, `hostName` and `message`. With Option 1, each Gravwell entry is exactly that JSON object. With Option 2, each entry is stored as Keycloak sent it: an RFC 5424 syslog header, a byte-order mark, then the JSON object. Running `json` on the whole entry returns nothing, so extract the JSON object first (the Keycloak kit's queries do this).

User and admin events have `loggerName` set to `org.keycloak.events`, and their details (`type`, `username`, `ipAddress`, `error`, `operationType`, `resourceType`, `resourcePath` and so on) are a `key="value", ...` string inside `message`, not separate JSON fields. Admin events record which object changed, but not its new values.

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

### [Option 1] Gravwell File Follower Configuration

**Sample Keycloak config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/keycloak.conf`
```ini
[Follower "keycloak"]
    Base-Directory="/opt/keycloak/data/log" #Keycloak's data/log directory, or the mounted directory for Docker
    File-Filter="keycloak.log"
    Tag-Name=keycloak
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```

### [Option 2] Gravwell Simple Relay Configuration

**Sample Keycloak config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/keycloak.conf`
```ini
[Listener "keycloak"]
    Bind-String="tcp://0.0.0.0:15060"
    Reader-Type=rfc5424 #Keycloak 26.2 and earlier; use rfc6587 for 26.3 and later (see above)
    Tag-Name=keycloak
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```

## Additional Resources
* [Keycloak Server Guide: Logging](https://www.keycloak.org/server/logging)
* [Keycloak Server Guide: Configuring a reverse proxy](https://www.keycloak.org/server/reverseproxy)
* [Keycloak Admin REST API](https://www.keycloak.org/docs-api/latest/rest-api/index.html)
* [Keycloak Server Administration Guide](https://www.keycloak.org/docs/latest/server_admin/)
