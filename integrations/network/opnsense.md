# OPNsense

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
:::

This guide sends the system logs of an OPNsense firewall to Gravwell. The logs record firewall rule decisions (connections that a rule passed, blocked or translated), DHCP leases, web GUI logins, and configuration changes along with the user who made each change. OPNsense sends the logs to Simple Relay over syslog, so the firewall needs no additional software. Gravwell stores them under the `opnsense` tag.

The steps in this guide were tested on OPNsense 25.7.11 and 26.7.8.

## OPNsense Configuration

Add the Simple Relay listener first (see "Gravwell Ingester Configuration: Simple Relay" below), so OPNsense can connect as soon as you apply the changes.

### Remote Logging

In the OPNsense web GUI, go to **System** > **Settings** > **Logging**, open the **Remote** tab and click **+** below the list of destinations.

![OPNsense System > Settings > Logging menu](images/opnsense_menu_logging.png)

![OPNsense Remote logging tab](images/opnsense_logging_remote.png)

Set the following options:

* **Enabled**: Checked.
* **Transport**: `TCP(4)`, or `TCP(6)` if you reach Simple Relay over IPv6.
* **Applications**: Leave as **Nothing selected**. An empty list sends every application, including applications that plugins add later. A list with selections sends only the selected applications.
* **Levels**: Leave as **Nothing selected**. OPNsense sends only messages at exactly the levels you select, not that level and above. Firewall rule decisions are logged at the informational level.
* **Facilities**: Leave as **Nothing selected**. Applications log to different facilities, and some use more than one, so a facility list drops part of their messages.
* **Hostname**: The address of your Simple Relay.
* **Port**: `15170`, the port of the Simple Relay listener below.
* **RFC5424**: Checked. Without it, OPNsense sends timestamps with no year and no UTC offset. Gravwell reads them as UTC, so entries from a firewall that is not set to UTC are stored with the wrong time. With RFC 5424, timestamps include the offset.
* **Suppress duplicate messages** (shown only in newer OPNsense versions): Unchecked. It replaces repeated messages with a summary line.
* **Description**: Use an identifiable name, for example `Gravwell Relay`.

Click **Save**, then **Apply**.

```{note}
Firewall rule decisions usually make up most of the log volume. OPNsense logs packets that match any rule with logging enabled. It also logs packets that match its default block and default pass rules, unless you turn that logging off. To change what the firewall logs, use the Logging section of **Firewall** > **Settings** > **Advanced** and the logging option of each rule, instead of the filters above. This keeps the firewall's other logs arriving.
```

### Web GUI Access Log

The web GUI access log records each request to the OPNsense web GUI and API, with the client address, the page or API path and the result. It is the only record of successful API calls made with an API key and of configuration backup downloads. It is off by default. This guide recommends turning it on.

Go to **System** > **Settings** > **Administration**, and under **Access log** check **Enable access log**. Click **Save**. The web GUI reloads after you save.

![OPNsense System > Settings > Administration menu](images/opnsense_menu_administration.png)

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/opnsense-well.conf`
```ini
[Storage-Well "opnsense"]
    Location=/opt/gravwell/storage/opnsense
    Tags=opnsense
```

If the `opnsense` tag already has data in another well, see "Tag Restrictions and Gotchas" in [Advanced Gravwell Configuration](/configuration/configuration) before adding it here.

### Gravwell Ingester Configuration: Simple Relay

This guide assumes Simple Relay is already installed and configured as described in [Simple Relay](/ingesters/simple_relay). Add the following to a new overlay file.

**Sample OPNsense Simple Relay config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/opnsense.conf`
```ini
[Listener "opnsense"]
    Bind-String="tcp://0.0.0.0:15170"
    Reader-Type=rfc5424
    Tag-Name=opnsense
```

This listener accepts unencrypted TCP. Simple Relay also supports TLS listeners (see "Bind-String" in [Simple Relay](/ingesters/simple_relay)), and OPNsense can send over TLS, but this guide does not cover that setup.

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```

To confirm data is arriving, run `tag=opnsense count`. Gravwell reports an unknown tag until the first entry arrives.

## Additional Resources
* [OPNsense: Settings, Logging](https://docs.opnsense.org/manual/settingsmenu.html#logging)
* [Simple Relay](/ingesters/simple_relay)
