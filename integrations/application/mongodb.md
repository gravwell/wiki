# MongoDB

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, • [File Follower](/ingesters/file_follow) <br /> • [Simple Relay](/ingesters/simple_relay)
:::

## MongoDB Configuration

### [Option 1] Logging to File Follower
Verify `systemLog` stanza in `/etc/mongod.conf` matches the following:
```
systemLog:
  destination: file
  logAppend: true
  path: /var/log/mongodb/mongod.log
```

```{note}
By default mongdb.log only has user `rw` permissions. For file follower to read the file these will need to be modified.
```

### [Option 2] Logging with Rsyslog to Simple Relay

Verify `systemLog` in `/etc/mongod.conf` matches the following:
```
systemLog:
  destination: syslog
  syslogFacility: daemon
  verbosity: 0
```

Create or edit `/etc/rsyslog.d/50-mongodb.conf`
```
:programname, isequal, "mongod" @@192.168.0.10
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/mongodb-well.conf`
```ini
[Storage-Well "mongodb"]
    Location=/opt/gravwell/storage/mongodb
    Tags=mongodb*
```
### [Option 1] Gravwell File Follower Configuration
**Sample MongoDB config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/mongodb.conf`
```ini
[Follower "mongodb"]
    Base-Directory="/var/log/mongodb/"
    File-Filter="mongod.log,mongod.log.[0-9]"
    Tag-Name=mongodb
    Assume-Local-Timezone=true #Default for assume localtime is false
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```

### [Option 2] Gravwell Simple Relay
**Sample MongoDB config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/mongodb.conf`
```ini
[Listener "mongodb"]
    Bind-String="tcp://0.0.0.0:519"
    Reader-Type=rfc5424
    Tag-Name=mongodb
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```