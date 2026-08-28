# Suricata

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25

**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
    Kit, [Suricata Kit](https://github.com/gravwell/kits/tree/main/suricata)
:::

## Suricata Configuration
Suricata by default stores its log files in `/var/log/suricata/`. The log file used by the Gravwell kit is `eve.json`. Additional log files can be monitored by adding additional stanzas to the File Follower config.

Install File Follower on your Suricata host by following the instructions in [File Follower](/ingesters/file_follow), then configure it to point to your Gravwell environment.

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

### Gravwell Ingester Configuration

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