# Synology

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
:::

## Synology Configuration

Synology DSM has remote syslog forwarding via Log Center. See [Log Sending | Log Center](https://kb.synology.com/en-global/DSM/help/LogCenter/logcenter_client?version=7).

1. Open **Log Center > Log Sending** and check **Send logs to a syslog server**.
2. Set **Server** to your Gravwell Simple Relay address.
3. Set **Transfer protocol** to UDP or TCP.
4. Set **Port** to match the Gravwell listener for that protocol: `15110` for TCP or `15112` for UDP in the sample below. DSM's own default is `514`, so change it.
5. Set **Log format** to IETF (RFC 5424).
6. If using TCP, optionally enable **secure connection (TLS/SSL)** and import a certificate. This requires the TLS listener shown in the sample below, and the port you set in step 4 must be that listener's port (`15111`).
7. Use the **Filter** tab to restrict which log categories/severities are forwarded, and **Send test log** to verify connectivity.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/synology-well.conf`
```ini
[Storage-Well "synology"]
    Location=/opt/gravwell/storage/synology
    Tags=synology*
```

### Gravwell Ingester Configuration: Simple Relay

The `rfc5424` reader below parses both log formats you can choose in DSM, BSD (RFC 3164) and IETF (RFC 5424).

**Sample Synology config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/synology.conf`
```ini
[Listener "synologytcp"]
    Bind-String="tcp://0.0.0.0:15110" #plain TCP, DSM transfer protocol TCP
    Reader-Type=rfc5424
    Tag-Name=synology
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
    Keep-Priority=true # leave the <nnn> priority tag at the start of each syslog entry

[Listener "synologyudp"]
    Bind-String="udp://0.0.0.0:15112" #UDP, DSM transfer protocol UDP
    Reader-Type=rfc5424
    Tag-Name=synology
    Keep-Priority=true

# Only needed if you enable TLS in DSM. Replace the certificate and key paths with your own.
[Listener "synologytls"]
    Bind-String="tls://0.0.0.0:15111"
    Cert-File=/opt/gravwell/etc/cert.pem
    Key-File=/opt/gravwell/etc/key.pem
    Reader-Type=rfc5424
    Tag-Name=synology
    Keep-Priority=true
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```
