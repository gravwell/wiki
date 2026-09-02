# Auditd

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
         Kit, [Auditd Kit](https://github.com/gravwell/kits/tree/main/auditd)
:::

## Auditd Configuration

Auditd defines its log location in `/etc/audit/auditd.conf` via the `log_file` parameter (for example, `log_file = /var/log/audit/audit.log`). Install File Follower on your Auditd host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/auditd-well.conf`
```ini
[Storage-Well "auditd"]
    Location=/opt/gravwell/storage/auditd
    Tags=auditd*
    # Hot-Duration=30d
    # Cold-Duration=90D
    # Max-Hot-Storage-GB=20
    # Delete-Frozen-Data=true
```

### Gravwell Ingester Configuration: File Follower
Setup the file follower configuration file.

**Sample File Follower configuration:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/auditd.conf`
```ini
[Follower "auditd"]
    Base-Directory = "/var/log/audit"
    File-Filter    = "audit.log"
    Tag-Name       = auditd
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```