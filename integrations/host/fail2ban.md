# Fail2ban

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
:::

## Fail2ban Configuration

Gravwell collects Fail2ban's own log with File Follower on each host that runs Fail2ban. Each line records one action: a failure a jail counted, an address it ignored, a ban or unban, or a server message. This log is separate from the logs the jails watch. See [fail2ban.conf](https://github.com/fail2ban/fail2ban/blob/master/config/fail2ban.conf) for the logging settings.

The steps in this guide were tested on Ubuntu 24.04 LTS with Fail2ban 1.0.2 and on Ubuntu 22.04 LTS with Fail2ban 0.11.2. Other Linux distributions can use different package commands, user names and file locations. Check each step against your system before you run it.

Fail2ban needs these two settings, which are the defaults in the `fail2ban.conf` it ships:

* `logtarget = /var/log/fail2ban.log`. Fail2ban can log to `SYSLOG`, `STDOUT` or `STDERR` instead, and from version 1.0.1 to `SYSTEMD-JOURNAL`. With any of these, the file this guide reads doesn't exist. If Fail2ban runs in a container, set its log target to a file on a mounted volume and point File Follower at that file instead of `/var/log`.
* `loglevel = INFO`. This is required. Bans and unbans are logged at `NOTICE`, but each failure a jail counts towards a ban, and each address a jail ignores, is logged at `INFO`. At `NOTICE` or above, Gravwell sees bans but not the activity that led to them, sources that stayed under the ban threshold, ignored addresses, or when jails start, stop or change settings.

To check and set them:

1. Print the current values:
   ```bash
   sudo fail2ban-client get logtarget
   sudo fail2ban-client get loglevel
   ```
   If they show `/var/log/fail2ban.log` and `INFO`, skip to the File Follower step below.
2. Create `/etc/fail2ban/fail2ban.d/gravwell.local` with these lines. Fail2ban reads files in `fail2ban.d` ending in `.local` after `fail2ban.conf` and `fail2ban.local`, so these settings take effect without editing either file, and package updates leave the new file alone.
   ```ini
   [Definition]
   loglevel = INFO
   logtarget = /var/log/fail2ban.log
   ```
3. Check the configuration: `sudo fail2ban-client -t`. It should print `OK: configuration test is successful`. If it reports an error, fix it before restarting, or Fail2ban won't start.
4. Restart Fail2ban: `sudo systemctl restart fail2ban`
5. Run the two commands from step 1 again. They should print `/var/log/fail2ban.log` and `INFO`. If they don't, another `.local` file in `fail2ban.d` that sorts after `gravwell.local` sets a different value; change the setting there instead.

To undo, delete `/etc/fail2ban/fail2ban.d/gravwell.local` and restart Fail2ban.

```{note}
Fail2ban writes its timestamps in the host's local time zone, without a time zone offset. The File Follower configuration below sets `Assume-Local-Timezone=true` so Gravwell reads them in the host's time zone. Without it, Gravwell treats them as UTC.
```

```{note}
On Debian and Ubuntu, `/var/log/fail2ban.log` is readable only by root and the `adm` group, and logrotate renames it weekly to `fail2ban.log.1`, then compresses older copies. File Follower must run as root or as a member of `adm`; the File Follower package's service already runs in the `adm` group. Follow only the live file, as shown below.
```

Install File Follower on your Fail2ban host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

### Log Format

Each line has a timestamp with milliseconds, the Fail2ban component, the process ID, the level and the message. Most messages start with the jail name in brackets:

```text
2026-10-18 01:33:13,102 fail2ban.filter         [116]: INFO    [nginx-abuse] Found 203.0.113.7 - 2026-10-18 01:33:13
2026-10-18 01:34:02,555 fail2ban.actions        [116]: NOTICE  [nginx-abuse] Ban 203.0.113.7
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fail2ban-well.conf`
```ini
[Storage-Well "fail2ban"]
    Location=/opt/gravwell/storage/fail2ban
    Tags=fail2ban
```

```{note}
Restart the indexer to apply the new well: `sudo systemctl restart gravwell_indexer.service`
```

### Gravwell Ingester Configuration: File Follower

**Sample Fail2ban config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/fail2ban.conf`
```ini
[Follower "fail2ban"]
    Base-Directory="/var/log"
    File-Filter="fail2ban.log"
    Tag-Name=fail2ban
    Assume-Local-Timezone=true #Fail2ban timestamps have no time zone offset
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```

To confirm that data is arriving, search `tag=fail2ban count`. When File Follower first starts, it reads the whole current `fail2ban.log`, so entries back to the last log rotation appear at once, with their original timestamps.

## Additional Resources
* [Fail2ban on GitHub](https://github.com/fail2ban/fail2ban)
* [fail2ban.conf (logging settings)](https://github.com/fail2ban/fail2ban/blob/master/config/fail2ban.conf)
* [jail.conf(5) manual page](https://manpages.debian.org/bookworm/fail2ban/jail.conf.5.en.html)
* [fail2ban-client(1) manual page](https://manpages.debian.org/bookworm/fail2ban/fail2ban-client.1.en.html)
