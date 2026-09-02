# APT-Cacher NG

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
:::

## APT-Cacher NG Configuration

APT-Cacher NG defines its log location in `acng.conf` via the `LogDir` parameter (for example, `LogDir: /var/log/apt-cacher-ng`). Install File Follower on your APT-Cacher NG host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/aptcacher.conf`
```ini
[Storage-Well "aptcacher"]
    Location=/opt/gravwell/storage/aptcacher
    Tags=aptcacher*
```

### Gravwell Ingester Configuration: File Follower
**Sample APT-Cacher NG config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/aptcacher.conf`
```ini
[Follower "aptcacher"]
    Base-Directory="/var/log/apt-cacher-ng/"
    File-Filter="apt-cacher.log,apt-cacher.[0-9]"
    Tag-Name=aptcacher
    Ignore-Timestamps=true

[Follower "aptcacher_err"]
    Base-Directory="/var/log/apt-cacher-ng/"
    File-Filter="apt-cacher.err,apt-cacher.err.[0-9]"
    Tag-Name=aptcacher-err
    Ignore-Timestamps=true
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```