# Proxmox

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
:::

## Proxmox Configuration

There are two primary methods to monitor Proxmox environments.
* **OpenTelemetry (Simplest, metrics only):** Proxmox provides metrics about the environment and individual virtual machines through OpenTelemetry. This method requires no additional installation and is useful for monitoring things like system, storage, and network usage. 
* **systemd logs (Proxmox logs):** systemd logs include Proxmox commands, such as VM power-on and restart. This method requires installing a package on Proxmox from the command line.

If you use both methods, send them to different URLs and tags.

### [Option 1] Exporting Proxmox Metrics via OpenTelemetry

OpenTelemetry collects metrics about each of the nodes and individual VMs. In the Proxmox web interface, go to **Datacenter** > **Metric Server** > **Add** > **OpenTelemetry** to send them to the HTTP Ingester, and set:

* **Name**: Use an identifiable name.
* **Server**: Set to the address of your HTTP Ingester.
* **Port**: Set to the port your HTTP Ingester listens on (its `Bind` setting, see [HTTP Ingester Setup](/integrations/generic/http_ingester)).
* **Protocol**: Set to the protocol used by your HTTP Ingester.
* **Path**: Set to the `URL` value in your HTTP Ingester listener (`/proxmox/v1/metrics` in the sample below).

![image](images/proxmox_opentelemetry.png)

### [Option 2] Exporting Proxmox Systemd Logs

Install `systemd-journal-remote` to export logs to Gravwell:
```
apt update && apt install systemd-journal-remote
```

Create or edit `/etc/systemd/journal-upload.conf` with the following:
```
[Upload]
# Set to the address and port of your HTTP Ingester, plus the listener path (/upload in the sample below)
URL=http://ingesterIP:port/upload
# ServerKeyFile=/etc/ssl/private/journal-upload.pem
# ServerCertificateFile=/etc/ssl/certs/journal-upload.pem
# TrustedCertificateFile=/etc/ssl/ca/trusted.pem
```

Enable and restart the service:
```
sudo systemctl enable systemd-journal-upload.service
sudo systemctl restart systemd-journal-upload.service
sudo systemctl status systemd-journal-upload.service
```

The first run uploads the entire journal, which may exceed the `Max-Body` limit of the Gravwell HTTP Ingester. You may need to increase `Max-Body` in the HTTP Ingester configuration:
* Find the current file size with `journalctl --disk-usage`
* Increase `Max-Body` in the `[Global]` block of the HTTP Ingester's base configuration file to at least that size, then restart the ingester. See [HTTP Ingester Setup](/integrations/generic/http_ingester).

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/proxmox-well.conf`
```ini
[Storage-Well "proxmox"]
    Location=/opt/gravwell/storage/proxmox
    Tags=prox-*
    Tags=otel-prox-*
```

### Gravwell Ingester Configuration: HTTP Ingester

This guide assumes the HTTP Ingester is already installed and configured as described in [HTTP Ingester Setup](/integrations/generic/http_ingester). Add the following to a new overlay file; you do not need a separate HTTP Ingester for this integration.

**Sample Proxmox HTTP config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/proxmox.conf`
```ini
# [Option 1] Open Telemetry Metrics Listener
[OpenTelemetry-Metrics-Listener "otel-metrics"]
    URL="/proxmox/v1/metrics"        # OTLP metrics endpoint with a Proxmox-specific prefix
    Tag-Name="otel-prox-metrics"     # Tag for ingested metrics
    Ignore-Timestamps=false          # Use timestamps from OTLP data
    Debug-Posts=true                 # Log debug info about requests
    Encode-As-JSON=true              # Encode metrics as JSON and include in the entry DATA, (doubles storage requirements)
#    Preprocessor="otel-processor"   # Optional preprocessor

# [Option 2] Exporting Proxmox Systemd logs
[Listener "systemd"]
    URL="/upload"
    Tag-Name=prox-systemd
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```
