# Auditd

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower ingester](/ingesters/file_follow.md)
         Kit, [Auditd Kit](https://github.com/gravwell/kits/tree/main/auditd)
:::

## Auditd Configuration

The standard method to collect Auditd logs is by installing the [gravwell-file-follow](file_follow_installation) package which can be installed through your [configured](/quickstart/quickstart.md) package manager or via a [standalone shell installer](/quickstart/downloads.md).

**Sample File Follower Configuration pointing to Gravwell Environment:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf`
```ini
Ingest-Secret = IngestSecrets
Insecure-Skip-TLS-Verify = false
Cleartext-Backend-Target=172.20.0.1:4023 #example of adding a cleartext connection
State-Store-Location=/opt/gravwell/etc/file_follow.state
Max-Files-Watched=64
```

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

### Gravwell File Follower Ingester Configuration

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