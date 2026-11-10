# Proxmox VE

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay) <br /> [HTTP Ingester](/ingesters/http)
:::

This guide collects system logs, the API access log and metrics from Proxmox Virtual Environment (Proxmox VE) nodes. They record different things, and each can be set up on its own:

* **System logs**: what happened on each node and who did it. They record who logged in to the web interface, API or SSH, including failed attempts. They record tasks such as starting, stopping or backing up a guest, opening a console and changing a guest's configuration, with the user who ran each one. They also record commands run with `sudo` and node health events such as disk (SMART) and ZFS warnings. `rsyslog`, installed on each node, forwards them to Simple Relay. Gravwell stores them under the `proxmox-system` tag.
* **API access log**: every request to the web interface and API, with the client's address, the user or API token and the response code. It is the only log that shows who changed users, permissions or API tokens, and from which address each session and API call came. `rsyslog` reads the log file and forwards it to Simple Relay. Gravwell stores it under the `proxmox-access` tag.
* **Metrics** (Proxmox VE 9.0 and later): how each node, virtual machine, container and storage is performing, such as CPU, memory, disk, network and storage use. They add performance and capacity data alongside the system logs. Proxmox VE sends them to the HTTP Ingester over OpenTelemetry, with no additional installation. Gravwell stores them under the `proxmox-metrics` tag.

The steps in this guide were tested on Proxmox VE 9.2.24. The system log steps were also tested on Proxmox VE 8.4.

## Proxmox Configuration

### System Logs

Proxmox VE writes its logs to the systemd journal. `rsyslog` forwards them to Gravwell Simple Relay, one entry per log message. Fresh installs of Proxmox VE 8.0 and later do not include `rsyslog`; nodes upgraded from earlier releases usually still have it. Run the following as `root` on each Proxmox VE node; every node sends its own logs.

1. Install `rsyslog`. If it is already installed, this changes nothing:

    ```bash
    apt update && apt install rsyslog
    ```

2. Check for existing forwarding rules. Run this before creating the file in the next step:

    ```bash
    grep -snE '^[^#]*(@@?[^ ]|omfwd)' /etc/rsyslog.conf /etc/rsyslog.d/*.conf
    ```

    No output means nothing is forwarded today. A line that sends to your Gravwell host (for example `*.* @gravwell.example.com:514`) would make the node send every message twice; remove or comment it out. Leave forwards to other destinations alone. For a multi-line `action(...)` block, the matching line is `type="omfwd"`; its `target` is in the same block.

    Before you edit a file, back it up. Use a new backup name each time:

    ```bash
    cp /etc/rsyslog.conf /etc/rsyslog.conf.bak-$(date +%Y%m%d-%H%M%S)
    ```

    Replace `/etc/rsyslog.conf` with the file the check printed.

3. Create `/etc/rsyslog.d/40-gravwell.conf` with the following, replacing `SIMPLE-RELAY-HOST` with the address of your Simple Relay:

    ```text
    # Forward all syslog messages to Gravwell Simple Relay over TCP, RFC 5424 format.
    # Messages queue on disk (up to 1 GB) while Simple Relay is unreachable.
    *.* action(
        type="omfwd"
        target="SIMPLE-RELAY-HOST"
        port="15160"
        protocol="tcp"
        template="RSYSLOG_SyslogProtocol23Format"
        action.resumeRetryCount="-1"
        queue.type="LinkedList"
        queue.filename="gravwell_fwd"
        queue.maxDiskSpace="1g"
        queue.saveOnShutdown="on"
    )
    ```

    Keep the `RSYSLOG_SyslogProtocol23Format` template: it sends timestamps with the node's UTC offset, so no timezone setting is needed.

    Keep the `40-` prefix. `rsyslog` reads the files in `/etc/rsyslog.d` in name order, and some packages add files that stop their messages after writing them locally. For example, the `frr` package (the routing suite used by Proxmox VE SDN) adds `45-frr.conf`. A file read after it would not forward FRR messages.

4. Check the configuration. The output ends with `End of config validation run. Bye.` and shows no error lines:

    ```bash
    rsyslogd -N1
    ```

5. Restart `rsyslog` to apply the configuration:

    ```bash
    systemctl restart rsyslog
    ```

6. Send a test message:

    ```bash
    logger -t gravwell-test "Gravwell forwarding test"
    ```

    Search for it in Gravwell with `tag=proxmox-system grep gravwell-test | table TIMESTAMP DATA`. TIMESTAMP is shown in UTC (or your Gravwell display timezone) and must match the time in the message header, which carries the node's offset.

`rsyslog` forwards new messages only; it does not send the journal history from before it started. While Simple Relay is unreachable, messages queue on disk under `/var/spool/rsyslog` and are sent when it is back, including across an `rsyslog` restart.

If you previously followed an earlier version of this guide that used `systemd-journal-upload`, stop it so the node does not send its logs twice:

```bash
systemctl disable --now systemd-journal-upload.service
```

To stop forwarding, remove the file and restart `rsyslog`:

```bash
rm /etc/rsyslog.d/40-gravwell.conf
systemctl restart rsyslog
```

If you removed a forward in step 2, restore it from your backup before the restart.

### API Access Log

The Proxmox VE API proxy, `pveproxy`, writes every web interface and API request to `/var/log/pveproxy/access.log`, not to the systemd journal. `rsyslog` reads the file and forwards each line to its own Simple Relay listener. Run the following as `root` on each Proxmox VE node. If you did not set up system logs, install `rsyslog` first (step 1 in "System Logs").

1. Create `/etc/rsyslog.d/41-gravwell-access.conf` with the following, replacing `SIMPLE-RELAY-HOST` with the address of your Simple Relay:

    ```text
    # Forward the Proxmox VE API access log to Gravwell Simple Relay over TCP, RFC 5424 format.
    module(load="imfile")
    input(type="imfile"
          File="/var/log/pveproxy/access.log"
          Tag="pveproxy-access"
          Severity="info"
          Facility="local0"
          freshStartTail="on"
          Ruleset="gravwell_access")
    ruleset(name="gravwell_access") {
        action(
            type="omfwd"
            target="SIMPLE-RELAY-HOST"
            port="15163"
            protocol="tcp"
            template="RSYSLOG_SyslogProtocol23Format"
            action.resumeRetryCount="-1"
            queue.type="LinkedList"
            queue.filename="gravwell_access_fwd"
            queue.maxDiskSpace="1g"
            queue.saveOnShutdown="on"
        )
    }
    ```

    The `ruleset` sends these lines only to this listener, not with the system logs.

2. Check the configuration. The output ends with `End of config validation run. Bye.` and shows no error lines:

    ```bash
    rsyslogd -N1
    ```

3. Restart `rsyslog` to apply the configuration:

    ```bash
    systemctl restart rsyslog
    ```

4. Send a test request:

    ```bash
    curl -sk -o /dev/null https://localhost:8006/api2/json/gravwell-test
    ```

    Search for it in Gravwell with `tag=proxmox-access grep gravwell-test | table TIMESTAMP DATA`. The line ends in `401`, because the request carries no login.

The first time `rsyslog` reads the file, it starts at the end, so earlier contents are not sent. After that it remembers its position: if you remove this file and later restore it, the lines written in between are sent when `rsyslog` restarts. It follows the file when Proxmox VE rotates it.

To stop forwarding, remove the file and restart `rsyslog`:

```bash
rm /etc/rsyslog.d/41-gravwell-access.conf
systemctl restart rsyslog
```

#### Behind a Reverse Proxy

If users reach the web interface through a reverse proxy, every line records the proxy's address instead of the client's. To record the client's address, set the proxy to replace the `X-Forwarded-For` header with the client's address, then on each node:

1. If `/etc/default/pveproxy` exists, back it up. Use a new backup name each time:

    ```bash
    cp /etc/default/pveproxy /etc/default/pveproxy.bak-$(date +%Y%m%d-%H%M%S)
    ```

2. Add the following to `/etc/default/pveproxy`, creating the file if needed. Replace `PROXY-ADDRESS` with the address your proxy connects from:

    ```text
    PROXY_REAL_IP_HEADER="X-Forwarded-For"
    PROXY_REAL_IP_ALLOW_FROM="PROXY-ADDRESS"
    ```

    Keep `PROXY_REAL_IP_ALLOW_FROM`. Without it, `pveproxy` logs the header from any client, so anyone can write a false address into the log.

3. Restart `pveproxy` to apply the configuration:

    ```bash
    systemctl restart pveproxy
    ```

Requests sent straight to port 8006 still record the connecting address. To undo, remove the two lines (or restore your backup) and restart `pveproxy`.

### Metrics

The OpenTelemetry metric server is available in Proxmox VE 9.0 and later. Add the HTTP Ingester listener first (see "Gravwell Ingester Configuration: HTTP Ingester" below): Proxmox VE tests the connection when you save the metric server, and the save fails until the listener exists.

In the Proxmox web interface, go to **Datacenter** > **Metric Server** > **Add** > **OpenTelemetry** to send metrics to the HTTP Ingester, and set:

* **Name**: Use an identifiable name.
* **Server**: Set to the address of your HTTP Ingester.
* **Port**: Set to the port your HTTP Ingester listens on (its `Bind` setting, see [HTTP Ingester Setup](/integrations/generic/http_ingester)).
* **Protocol**: Set to the protocol used by your HTTP Ingester. Proxmox VE defaults to `https`; choose `http` if your HTTP Ingester does not use TLS.
* **Path**: Set to the `URL` value in your HTTP Ingester listener (`/proxmox/v1/metrics` in the sample below).

![image](images/proxmox_opentelemetry.png)

Proxmox VE sends every metric every 10 seconds, and Gravwell stores each data point as one entry, so the volume grows with the number of virtual machines, containers and storages. For example, a node with 45 VMs, 22 containers and 4 storages sends about 3,300 data points every 10 seconds. That is about 29 million entries and about 1.3 GB on disk per day. Set the well's retention with this in mind.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/proxmox-well.conf`
```ini
[Storage-Well "proxmox"]
    Location=/opt/gravwell/storage/proxmox
    Tags=proxmox-system
    Tags=proxmox-access
    Tags=proxmox-metrics
```

If any of these tags already has data in another well, see "Tag Restrictions and Gotchas" in [Advanced Gravwell Configuration](/configuration/configuration) before adding it here.

### Gravwell Ingester Configuration: Simple Relay

Used for system logs and the API access log. This guide assumes Simple Relay is already installed and configured as described in [Simple Relay](/ingesters/simple_relay). Add the following to a new overlay file.

**Sample Proxmox Simple Relay config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/proxmox.conf`
```ini
[Listener "proxmox"]
    Bind-String="tcp://0.0.0.0:15160"
    Reader-Type=rfc5424
    Tag-Name=proxmox-system

[Listener "proxmox-access"]
    Bind-String="tcp://0.0.0.0:15163"
    Reader-Type=rfc5424
    Tag-Name=proxmox-access
```

These listeners accept unencrypted TCP. Simple Relay also supports TLS listeners (see "Bind-String" in [Simple Relay](/ingesters/simple_relay)), and `rsyslog` can forward over TLS, but this guide does not cover that setup.

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```

### Gravwell Ingester Configuration: HTTP Ingester

Used for metrics. This guide assumes the HTTP Ingester is already installed and configured as described in [HTTP Ingester Setup](/integrations/generic/http_ingester). Add the following to a new overlay file; you do not need a separate HTTP Ingester for this integration.

**Sample Proxmox HTTP config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/proxmox.conf`
```ini
[OpenTelemetry-Metrics-Listener "proxmox"]
    URL="/proxmox/v1/metrics"        # OTLP metrics endpoint with a Proxmox-specific prefix
    Tag-Name="proxmox-metrics"       # Tag for ingested metrics
    Ignore-Timestamps=false          # Use timestamps from OTLP data
    Encode-As-JSON=true              # Encode metrics as JSON and include in the entry DATA
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```

To confirm data is arriving, run `tag=proxmox-system count` for system logs, `tag=proxmox-access count` for the API access log or `tag=proxmox-metrics count` for metrics. Gravwell reports an unknown tag until the first entry arrives.

## Additional Resources
* [rsyslog: omfwd (Forwarding Output Module)](https://www.rsyslog.com/doc/configuration/modules/omfwd.html)
* [rsyslog: imfile (Text File Input Module)](https://www.rsyslog.com/doc/configuration/modules/imfile.html)
* [Proxmox VE: pveproxy, Real Client IP Logging](https://pve.proxmox.com/pve-docs/pveproxy.8.html)
* [Proxmox VE: External Metric Server](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#external_metric_server)
* [Simple Relay](/ingesters/simple_relay)
* [HTTP Ingester](/ingesters/http)
