# FleetDM

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
:::

## FleetDM Configuration

Fleet's server-side osquery logging plugin controls where **result logs** (query/live-query output) and **status logs** (osquery daemon health/diagnostics) go, set independently via `--osquery_result_log_plugin` and `--osquery_status_log_plugin`. See [Fleet server configuration](https://fleetdm.com/docs/configuration/fleet-server-configuration) and [Log destinations](https://fleetdm.com/docs/using-fleet/log-destinations).

Fleet documents several plugin values this guide uses the `filesystem` plugin, Fleet's default, because it covers both log streams the same way.

With the default `filesystem` plugin, Fleet writes newline-delimited JSON to:
* `filesystem_result_log_file` (default `/tmp/osquery_result`)
* `filesystem_status_log_file` (default `/tmp/osquery_status`)

Rotation is available via `filesystem_enable_log_rotation` (default max size 500MB / max age 28 days / max backups 3).

```{note}
`/tmp` is cleared on reboot, and a Fleet service started with systemd `PrivateTmp=true` writes to a private `/tmp` that File Follower cannot see. For a production deployment, set `filesystem_result_log_file` and `filesystem_status_log_file` to a persistent directory such as `/var/log/fleet/`, and use that directory as `Base-Directory` below.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fleetdm-well.conf`
```ini
[Storage-Well "fleetdm"]
    Location=/opt/gravwell/storage/fleetdm
    Tags=fleetdm*
```

### Gravwell Ingester Configuration: File Follower

Install File Follower on the Fleet server host by following the instructions in [File Follower](/ingesters/file_follow), then point it at the configured log file locations. The filters below match the exact file names, so the rotated backups that Fleet creates when log rotation is enabled are not ingested a second time.

**Sample FleetDM config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/fleetdm.conf`
```ini
[Follower "fleetdm-result"]
    Base-Directory="/tmp/"
    File-Filter="osquery_result"
    Tag-Name=fleetdm-result

[Follower "fleetdm-status"]
    Base-Directory="/tmp/"
    File-Filter="osquery_status"
    Tag-Name=fleetdm-status
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```