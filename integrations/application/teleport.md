# Teleport

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
    Kit, [Teleport Kit](https://github.com/gravwell/kits/tree/main/teleport)
:::

## Teleport Configuration

For self-hosted clusters, Teleport defines its audit log location in `teleport.yaml` via the `audit_events_uri` parameter (for example, `audit_events_uri: ['file:///var/lib/teleport/log']`). Install File Follower on your Teleport host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

Logs can also be exported using Fluentd. Follow the Teleport guide here: [Export Audit Events: Fluentd](https://goteleport.com/docs/zero-trust-access/export-audit-events/fluentd/), then follow the Gravwell [Fluentd integration guide](fluentd.md). With the following modifications:

Changes necessary to: `/etc/fluent/fluentd.conf`
* `<match **>`: Change to match the pattern of your Teleport input
* `endpoint http://path.to.gravwell:port/fluentd`: Change to `endpoint http://path.to.gravwell:port/teleport`

Changes necessary to: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/fluentd.conf`
```ini
[Listener "teleport"]
    URL="/teleport"
    Tag-Name="teleport-audit"
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/teleport-well.conf`
```ini
[Storage-Well "teleport"]
    Location=/opt/gravwell/storage/teleport
    Tags=teleport*
```

### Gravwell Ingester Configuration

**Sample File Follower config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/teleport.conf`
```ini
[Follower "teleport"]
    Base-Directory = "/var/lib/teleport/log"
    File-Filter    = "*.log"
    Recursive      = true
    Tag-Name       = "teleport-audit"
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```