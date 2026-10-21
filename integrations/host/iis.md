# Microsoft IIS

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Windows File Follower](/ingesters/win_file_follow)
    Ingester, [Windows Event Ingester](/ingesters/winevent)
:::

## Microsoft IIS Configuration

IIS writes one log file per site per day in the W3C Extended log file format, and HTTP.sys writes a separate error log (see HTTP.sys Error Log below). Gravwell's Windows File Follower reads both on the IIS server. See [Configure Logging in IIS](https://learn.microsoft.com/en-us/iis/manage/provisioning-and-managing-iis/configure-logging-in-iis) and the [`<logFile>` reference](https://learn.microsoft.com/en-us/iis/configuration/system.applicationhost/sites/sitedefaults/logfile/).

* Log directory: `%SystemDrive%\inetpub\logs\LogFiles\W3SVC<site ID>\`, one folder per site.
* File name: `u_ex<YYMMDD>.log`, one file per day by default.
* Keep the log format **W3C**, the default.

### Logging Fields

W3C logs are space-delimited, and IIS writes only the fields you select, always in a fixed order. What a server logs today depends on its history:

* If its field list has never been changed, it logs the 15 IIS default fields.
* On Windows Server 2019 and later (and Windows 11), a Windows update from February 2026 or later adds Bytes Sent and Bytes Received to those defaults, for 17 fields.
* If fields were chosen in IIS Manager (**Select Fields**) or by a script, it logs that selection.

Whichever applies, add the following fields to the defaults. IIS doesn't log them by default (except Bytes Sent and Bytes Received after the update above), and Gravwell's per-site, per-server, host name and bandwidth queries need them:

| IIS Manager name | `logExtFileFlags` value | Log field | What it adds |
|---|---|---|---|
| Site Name | `SiteName` | `s-sitename` | Which site served the request, as the site ID (`W3SVC1`) |
| Server Name | `ComputerName` | `s-computername` | Which server, when logs from several servers are collected |
| Protocol Version | `ProtocolVersion` | `cs-version` | The HTTP version of the request |
| Host | `Host` | `cs-host` | The host name the client requested |
| Bytes Sent | `BytesSent` | `sc-bytes` | Response size |
| Bytes Received | `BytesRecv` | `cs-bytes` | Request size, including uploads |

```{warning}
Enable all six fields, or none of them. Gravwell reads IIS logs by position. Some partial selections shift values into the wrong fields without any error: Bytes Sent without Bytes Received, for example, makes response sizes appear as response times. Don't add other standard fields, such as Cookie, for the same reason.
```

The `X-Forwarded-For` field in Behind a Reverse Proxy is a custom field. IIS writes custom fields after the standard ones, so it doesn't shift them. Don't add other custom fields.

#### Set the Recommended Fields

Run these PowerShell commands as an administrator on the IIS server. They set the same 21 fields whatever the server logs today.

1. Back up the IIS configuration. The backup name must be new: if a backup with that name already exists, the command fails and keeps the old one.

   ```powershell
   Import-Module WebAdministration
   Backup-WebConfiguration -Name 'before-gravwell-logging'
   ```

2. Record the current field lists. The first command prints the server-wide list. The second prints each site's list in full.

   ```powershell
   Get-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/logFile' -Name logExtFileFlags
   Get-ChildItem IIS:\Sites | Select-Object Name, @{ n = 'logExtFileFlags'; e = { $_.logFile.logExtFileFlags } } | Format-List
   ```

3. Set the recommended fields as the server-wide list:

   ```powershell
   Set-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/logFile' -Name logExtFileFlags `
     -Value 'Date,Time,ClientIP,UserName,ServerIP,Method,UriStem,UriQuery,TimeTaken,HttpStatus,Win32Status,ServerPort,UserAgent,HttpSubStatus,Referer,SiteName,ComputerName,ProtocolVersion,Host,BytesSent,BytesRecv'
   ```

4. List the sites that log a different field list from the server-wide one:

   ```powershell
   $d = Get-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/logFile' -Name logExtFileFlags
   Get-ChildItem IIS:\Sites | Where-Object { $_.logFile.logExtFileFlags -ne $d } | Select-Object Name
   ```

   No output means every site uses the server-wide list. A listed site has its own field list. Remove it so the site uses the server-wide list:

   ```powershell
   Clear-ItemProperty 'IIS:\Sites\<site name>' -Name logFile.logExtFileFlags
   ```

5. Browse each site once, then check the fields IIS writes to each site's newest log file:

   ```powershell
   Get-ChildItem C:\inetpub\logs\LogFiles -Directory -Filter W3SVC* | ForEach-Object {
     $f = Get-ChildItem $_.FullName -Filter u_ex*.log | Sort-Object LastWriteTime | Select-Object -Last 1
     [pscustomobject]@{ Site = $_.Name; Fields = (Select-String -Path $f.FullName -Pattern '^#Fields:' | Select-Object -Last 1).Line.Split(' ').Count - 1 }
   }
   ```

   Every site should show `21` (`22` once `X-Forwarded-For` is added). IIS writes the new field list only with a site's next request, so a site that hasn't been browsed still shows its old count. A folder for a site that no longer exists keeps its last count; ignore it. The command assumes the default log directory.

* The server-wide list applies to every site that doesn't have its own list, including sites created later. Step 3 replaces the whole list, and it's safe to run again.
* The change applies immediately, without restarting IIS.
* In IIS Manager instead: select the server node, open **Logging**, choose **Select Fields**, and select the default fields plus the six above.
* If the server is behind a reverse proxy or load balancer, also add `X-Forwarded-For` (see Behind a Reverse Proxy).

The resulting layout is:

```text
#Fields: date time s-sitename s-computername s-ip cs-method cs-uri-stem cs-uri-query s-port cs-username c-ip cs-version cs(User-Agent) cs(Referer) cs-host sc-status sc-substatus sc-win32-status sc-bytes cs-bytes time-taken
```

To go back to only the previous field list, run step 3 with the list step 2 printed. On a server whose fields were never changed, this saves the default list as an explicit setting; the fields logged are the same. Or run `Restore-WebConfiguration -Name 'before-gravwell-logging'`, which restores the whole IIS configuration to the backup from step 1, including anything else changed since.

### Behind a Reverse Proxy

Behind a reverse proxy or load balancer, `c-ip` holds the proxy's address for every request. To keep the client's address, log the `X-Forwarded-For` request header as a custom field (IIS 8.5 and later). Run these commands once, after setting the recommended fields:

1. Back up the IIS configuration, using a new backup name:

   ```powershell
   Import-Module WebAdministration
   Backup-WebConfiguration -Name 'before-gravwell-xff'
   ```

2. List any custom fields the server already logs. No output means it logs none.

   ```powershell
   Get-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/logFile/customFields' -Name Collection |
     Select-Object logFieldName, sourceName, sourceType
   ```

3. Add `X-Forwarded-For`:

   ```powershell
   Add-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/logFile/customFields' -Name '.' `
     -Value @{ logFieldName = 'X-Forwarded-For'; sourceName = 'X-Forwarded-For'; sourceType = 'RequestHeader' }
   ```

In IIS Manager: server node > **Logging** > **Select Fields** > **Add Field**, with **Source type** set to **Request Header** and **Source** set to `X-Forwarded-For`.

* If step 3 reports `Cannot add duplicate collection entry`, the field is already configured. Nothing needs to change.
* IIS writes custom fields after the standard fields, so `X-Forwarded-For` is the last field. Requests that don't come through the proxy log `-`.
* `c-ip` always stays the address IIS received the connection from. `X-Forwarded-For` is set by the client or the proxy, so treat it as context, not proof of where a request came from.

To remove only `X-Forwarded-For` again (safe to run more than once):

```powershell
Remove-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
  -Filter 'system.applicationHost/sites/siteDefaults/logFile/customFields' -Name '.' `
  -AtElement @{ logFieldName = 'X-Forwarded-For' }
```

Or run `Restore-WebConfiguration -Name 'before-gravwell-xff'`, which also rolls back every other IIS change made since the backup.

```{note}
Named backups stay in `%windir%\System32\inetsrv\backup\<name>\` until you delete the folder, for example `Remove-Item -Recurse "$env:windir\System32\inetsrv\backup\before-gravwell-logging"`. The backup list shows every creation time as midnight, so tell same-day backups apart by name.
```

### Log Format

* `date` and `time` are UTC, whatever the server's time zone.
* IIS writes `+` in place of spaces in `cs(User-Agent)` and `cs(Referer)`, and `-` for any empty field.
* `time-taken` is in milliseconds.
* `s-sitename` is the site ID (`W3SVC1`), not the site's name. IIS Manager lists each site's ID under **Sites**.
* IIS writes header lines (`#Software`, `#Version`, `#Date`, `#Fields`) at the start of every file, and again whenever IIS restarts or the logging settings change.

### HTTP.sys Error Log (HTTPERR)

HTTP.sys, the Windows component that receives requests before IIS, rejects some requests and connections itself, including some path traversal attempts. They never reach the IIS logs above and are written only to the HTTP.sys error log, so collect it as a second stream. See [Error Logging in the HTTP Server API](https://learn.microsoft.com/en-us/windows/win32/http/error-logging-in-the-http-server-api).

* Error logging is on by default. Nothing needs to change in IIS or Windows.
* Log directory: `%SystemRoot%\System32\LogFiles\HTTPERR\`, one log per server, not per site.
* Files are named `httperr<N>.log`. HTTP.sys starts the next file when the current one reaches 1 MB. Each file starts with its own header lines.

The layout is fixed by Windows:

```text
#Fields: date time c-ip c-port s-ip s-port cs-version cs-method cs-uri streamid streamid_ex sc-status s-siteid s-reason s-queuename transport
```

The reason for each entry is `s-reason`, the third field from the end.

### IIS FTP Service Logs

If the server runs the IIS FTP service, IIS writes its logs to `FTPSVC<n>` folders under the same `LogFiles` directory, with the same `u_ex*.log` file names and a different layout (13 fields). The File Follower's `File-Filter` matches file names only, so the `iis` follower in Gravwell Configuration below reads FTP logs into the `iis` tag too. To keep them out, move FTP logging to its own directory. Servers without the FTP service can skip this section.

Run these PowerShell commands as an administrator on the IIS server:

1. Back up the IIS configuration, using a new backup name:

   ```powershell
   Import-Module WebAdministration
   Backup-WebConfiguration -Name 'before-gravwell-ftp'
   ```

2. Check where FTP logs go today. The first command prints the server-wide FTP log directory, `C:\inetpub\logs\LogFiles` on a server where it was never changed. The second lists each FTP site and the directory it actually logs to; a site with its own directory shows a different path. The third lists FTP log folders in the directory the `iis` follower reads; no output means there are none.

   ```powershell
   (Get-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/ftpServer/logFile' -Name directory).Value
   Get-ChildItem IIS:\Sites | Where-Object { $_.Bindings.Collection.protocol -contains 'ftp' } |
     Select-Object Name, Id, @{n='LogDirectory';e={$_.ftpServer.logFile.directory}}
   Get-ChildItem C:\inetpub\logs\LogFiles -Directory -Filter FTPSVC* | Select-Object Name
   ```

3. Create the new directory and make it the server-wide FTP log directory:

   ```powershell
   New-Item -ItemType Directory C:\inetpub\logs\FTP -Force
   Set-WebConfigurationProperty -PSPath 'MACHINE/WEBROOT/APPHOST' `
     -Filter 'system.applicationHost/sites/siteDefaults/ftpServer/logFile' -Name directory -Value 'C:\inetpub\logs\FTP'
   ```

4. Run the second and third commands from step 2 again. Every FTP site should show `C:\inetpub\logs\FTP`. Each folder the third command lists holds FTP logs written before the move, and the `iis` follower keeps reading them. Copy them elsewhere if you need them, then delete each one:

   ```powershell
   Remove-Item -Recurse 'C:\inetpub\logs\LogFiles\FTPSVC<n>'
   ```

* The change applies immediately, without restarting IIS, to running FTP sites and to FTP sites created later. A site with its own log directory keeps it.
* Step 3 changes only `applicationHost.config`, and it's safe to run again.

To undo the move (safe to run more than once), clear the server-wide setting. New FTP rows go back to `LogFiles` at once, where the `iis` follower reads them:

```powershell
Clear-WebConfiguration -PSPath 'MACHINE/WEBROOT/APPHOST' `
  -Filter 'system.applicationHost/sites/siteDefaults/ftpServer/logFile/@directory'
```

Or run `Restore-WebConfiguration -Name 'before-gravwell-ftp'`, which also rolls back every other IIS change made since the backup.

### IIS Configuration Event Log

Windows can record changes to the IIS configuration in the `Microsoft-IIS-Configuration/Operational` event log: the setting that changed, its old and new values, the account that made the change, and the ID of the process that saved it. Changes made with PowerShell or `appcmd` are recorded. Changes made by editing a configuration file directly are not. The log is off by default. Gravwell collects it with the Windows Event Ingester (see Gravwell Configuration).

Run these commands as an administrator on the IIS server:

1. Check whether the log is on. `enabled: false` means it is off.

   ```powershell
   wevtutil gl Microsoft-IIS-Configuration/Operational
   ```

2. Turn it on. The command prints nothing, and it's safe to run again.

   ```powershell
   wevtutil sl Microsoft-IIS-Configuration/Operational /e:true
   ```

3. Run step 1 again. It should show `enabled: true`.

* Changes made while the log is off are never recorded, and turning it on doesn't add them later. Turn it on before making the logging changes in this guide, and they are recorded too. Undoing the FTP log move is the exception: it is recorded only as a save to the server configuration, without the setting or its old value.
* The log keeps 1 MB and then overwrites its oldest events. Installing an IIS feature can fill most of it, so the copy in Gravwell is the lasting record.

To turn the log off again (safe to run more than once; events already in the log stay):

```powershell
wevtutil sl Microsoft-IIS-Configuration/Operational /e:false
```

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

Install the Windows File Follower on the IIS server by following the instructions in [Windows File Follower](/ingesters/win_file_follow).

**Sample IIS config:**  
Edit: `%PROGRAMDATA%\gravwell\filefollow\file_follow.cfg`
```ini
[Follower "iis"]
    Base-Directory=`C:\inetpub\logs\LogFiles`
    File-Filter=`u_ex*.log`
    Tag-Name=iis
    Recursive=true

[Follower "iis-httperr"]
    Base-Directory=`C:\Windows\System32\LogFiles\HTTPERR`
    File-Filter=`httperr*.log`
    Tag-Name=iis-httperr
```

* `Recursive=true` reads every site's folder under `LogFiles`. On a server that runs the IIS FTP service, move the FTP logs first (see IIS FTP Service Logs), or the `iis` follower reads them too.
* `u_ex*.log` also matches `u_ex<YYMMDD>_x.log`, the name IIS uses once a custom field such as `X-Forwarded-For` is added. Don't narrow it.
* `httperr*.log` follows HTTP.sys from one numbered file to the next.
* The storage well above (`Tags=iis*`) covers both tags, and `iis-ftp` below.
* No time zone setting is needed. IIS and HTTP.sys timestamps are UTC, which is what File Follower assumes for timestamps without a time zone.

To collect the FTP logs as well, add a third follower for the directory they were moved to. Keep its header lines, because the FTP `#Fields` line differs from the web layout:

```ini
[Follower "iis-ftp"]
    Base-Directory=`C:\inetpub\logs\FTP`
    File-Filter=`u_ex*.log`
    Tag-Name=iis-ftp
    Recursive=true
```

```{note}
The followers keep the header lines. The `#Fields` lines are the only record in the data of which fields a server logs, which differs between servers (15, 17 or 21 or more fields, depending on their history) and changes whenever the logging settings do. To save ingest on the `iis` follower, add `Ignore-Line-Prefix="#"`: it drops the header lines, including those IIS writes mid-file after a restart or a logging change, and keeps every data line. Gravwell's IIS queries don't need the header lines. Keep them on `iis-httperr`: they are few.
```

```{note}
Remember to restart the Gravwell File Follow service via standard Windows service management to apply the new config, for example `Restart-Service 'GravwellFileFollower Service'` in an administrator PowerShell.
```

After the restart, browse one of the sites, then search `tag=iis count` in Gravwell: it should return a count. `tag=iis-httperr count` may be empty on a quiet server, because HTTP.sys only writes an entry when it rejects something. IIS buffers FTP log writes, so FTP rows can reach the file several minutes after the session (row timestamps are the real event times); search `tag=iis-ftp count` after an FTP session and that delay.

### Gravwell Ingester Configuration: Windows Event

Install the Windows Event Ingester on the IIS server by following the instructions in [Windows Event Ingester](/ingesters/winevent), then add a channel for the IIS configuration event log.

**Sample Windows Event config:**  
Create or edit: `%PROGRAMDATA%\gravwell\eventlog\config.cfg`
```ini
[EventChannel "iis-config"]
    Tag-Name=iis-config
    Channel=Microsoft-IIS-Configuration/Operational
    Max-Reachback=24h
```

* The storage well above (`Tags=iis*`) covers `iis-config`.
* Collect the System, Application and Security logs from the IIS server too, as described in [Windows Event](/integrations/host/windowsevent). They record application pool failures, ASP.NET events (on servers with ASP.NET installed) and sign-in failures on the web server.

```{note}
Remember to restart the Gravwell Windows Event service via standard Windows service management to apply the new config, for example `Restart-Service 'GravwellEvents Service'` in an administrator PowerShell.
```

After the restart, search `tag=iis-config count` in Gravwell. It stays empty until the next change to the IIS configuration.

## Additional Resources
* [Configure Logging in IIS](https://learn.microsoft.com/en-us/iis/manage/provisioning-and-managing-iis/configure-logging-in-iis)
* [Default Log File Settings for Web Sites `<logFile>`](https://learn.microsoft.com/en-us/iis/configuration/system.applicationhost/sites/sitedefaults/logfile/)
* [Custom Fields for Log Files `<customFields>`](https://learn.microsoft.com/en-us/iis/configuration/system.applicationhost/sites/site/logfile/customfields/)
* [Default FTP Log File Settings `<logFile>`](https://learn.microsoft.com/en-us/iis/configuration/system.applicationhost/sites/sitedefaults/ftpserver/logfile/)
* [W3C Logging](https://learn.microsoft.com/en-us/windows/win32/http/w3c-logging)
* [Error Logging in the HTTP Server API](https://learn.microsoft.com/en-us/windows/win32/http/error-logging-in-the-http-server-api)
* [Types of Errors Logged by the HTTP Server API](https://learn.microsoft.com/en-us/windows/win32/http/types-of-errors-logged-by-the-http-server-api)
