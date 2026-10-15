# Microsoft IIS

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Windows File Follower](/ingesters/win_file_follow)
:::

## Microsoft IIS Configuration

IIS writes its logs in the W3C Extended log file format. It can store them locally or [push them to a remote share](https://learn.microsoft.com/en-us/iis/manage/provisioning-and-managing-iis/managing-iis-log-file-storage). See [Configure Logging in IIS](https://learn.microsoft.com/en-us/iis/manage/provisioning-and-managing-iis/configure-logging-in-iis) and [W3C Extended Log File Format](https://learn.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2003/cc786596(v=ws.10)).

* Configured per-site or globally via **IIS Manager > [Server or Site] > Logging**.
* Default log directory: `%SystemDrive%\inetpub\logs\LogFiles\W3SVC<siteID>\`.
* Default filename pattern: `u_exYYMMDD.log` (e.g. `u_ex260921.log`), one file per day by default.
* Default format is W3C Extended (space-delimited fields such as `date`, `time`, `c-ip`, `cs-username`, `s-ip`, `cs-method`, `cs-uri-stem`, `sc-status`, `time-taken`).

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/iis-well.conf`
```ini
[Storage-Well "iis"]
    Location=/opt/gravwell/storage/iis
    Tags=iis*
```

### Gravwell Ingester Configuration: Windows File Follower

Install the Windows File Follower on the IIS host, following [Windows File Follower](/ingesters/win_file_follow). Point it at the site's log directory; adjust `W3SVC1` to match the site ID you want to follow (visible in IIS Manager, or as the folder name under `LogFiles`).

**Sample IIS config:**  
Create or edit: `%PROGRAMDATA%\gravwell\filefollow\file_follow.cfg`
```ini
[Follower "iis"]
    Base-Directory="C:\\inetpub\\logs\\LogFiles\\W3SVC1\\"
    File-Filter="u_ex*.log"
    Tag-Name=iis
    Ignore-Line-Prefix="#"
```

```{note}
IIS writes W3C Extended timestamps in UTC without a timezone label (for example, `2026-09-21 14:03:11`), and the File Follower parses unlabeled timestamps as UTC by default, so no timezone setting is needed. The `Ignore-Line-Prefix` line drops the `#Software:`, `#Version:`, `#Date:` and `#Fields:` header lines that IIS writes at the top of each file.
```

```{note}
Remember to restart the Gravwell File Follow service via standard Windows service management to apply the new config.
```
