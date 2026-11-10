# Sysmon

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Windows Event Ingester](winevent_optional-sysmon-integration)
    Kit, [Windows Sysmon Kit](https://github.com/gravwell/kits/tree/main/sysmon)
:::

## Sysmon Configuration

The Sysmon utility, part of the Sysinternals suite, is an effective and popular tool for monitoring Windows systems. There are plenty of resources with examples of good Sysmon configuration files. Gravwell typically uses the modular Sysmon config on GitHub from [olafhartong](https://github.com/olafhartong/sysmon-modular).

[Download the default sysmon configuration file](https://raw.githubusercontent.com/olafhartong/sysmon-modular/master/sysmonconfig.xml)

[Download Sysmon](https://technet.microsoft.com/en-us/sysinternals/sysmon)

Install Sysmon with your configuration using an administrator shell (PowerShell works too) by running the following command:

```powershell
sysmon.exe -accepteula -i sysmonconfig.xml
```
## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/sysmon-well.conf`
```ini
[Storage-Well "sysmon"]
    Location=/opt/gravwell/storage/sysmon
    Tags=sysmon*
```

### Gravwell Ingester Configuration: Windows Event
**Sample Sysmon config:**  
Create or edit: `%PROGRAMDATA%\gravwell\eventlog\config.cfg`
```ini
[EventChannel "Sysmon"]
        Tag-Name=sysmon
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

```{note}
Remember to restart the Gravwell service via standard Windows service management to apply the new config.
```