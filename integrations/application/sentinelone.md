# SentinelOne

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
:::

## SentinelOne Configuration

SentinelOne supports native syslog/CEF forwarding of threat and alert events:

1. Go to **Settings > Integrations > Syslog**.
2. Enter the destination IP address of your Gravwell Simple Relay ingester and the port of its listener (`601` in the sample below).
3. Choose **CEF2** as the format.
4. Optionally enable TLS. If you do, use the `tls://` Bind-String shown in the sample below.
5. Under **Notifications**, enable the event types you want forwarded.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/sentinelone-well.conf`
```ini
[Storage-Well "sentinelone"]
    Location=/opt/gravwell/storage/sentinelone
    Tags=sentinelone*
```

### Gravwell Ingester Configuration: Simple Relay

**Sample SentinelOne config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/sentinelone.conf`
```ini
[Listener "sentinelone"]
    Bind-String="tcp://0.0.0.0:601"
    # Add the following for TLS
    #Bind-String="tls://0.0.0.0:601"
    #Cert-File=/opt/gravwell/etc/cert.pem
    #Key-File=/opt/gravwell/etc/key.pem
    Tag-Name=sentinelone
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
    Keep-Priority=true # leave the <nnn> priority tag at the start of each syslog entry
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```
