# Suricata

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
:::

## Suricata Configuration

Suricata defines its log location in `suricata.yaml` via the `default-log-dir` parameter (for example, `default-log-dir: /var/log/suricata/`). The log file used by the Gravwell kit is `eve.json`; additional log files can be monitored by adding additional stanzas to the File Follower configuration. Install File Follower on your Suricata host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration on your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/suricata-well.conf`
```ini
[Storage-Well "suricata"]
    Location=/opt/gravwell/storage/suricata
    Tags=suricata*
```

### Gravwell Ingester Configuration: File Follower

**Sample File Follower config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/suricata.conf`
```ini
[Follower "suricata"]
    Base-Directory = "/var/log/suricata"
    File-Filter    = "eve.json"
    Tag-Name       = "suricata"
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```
