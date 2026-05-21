# Sysmon

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Windows Event Ingester](winevent_optional-sysmon-integration)
         Kit, [Windows Sysmon Kit](https://github.com/gravwell/kits/tree/main/)
:::

## Sysmon Configuration

The Sysmon utility, part of the sysinternals suite, is an effective and popular tool for monitoring Windows systems. There are plenty of resources with examples of good sysmon configuration files. At Gravwell, we like to use the modular sysmon config on github from [olafhartong](https://github.com/olafhartong/sysmon-modular).

[Download the default sysmon configuration file](https://raw.githubusercontent.com/olafhartong/sysmon-modular/master/sysmonconfig.xml)

[Download sysmon](https://technet.microsoft.com/en-us/sysinternals/sysmon)

Install sysmon with your configuration using an administrator shell (Powershell works too) by running the following command:

```powershell
sysmon.exe -accepteula -i sysmonconfig-export.xml
```
Restart the Gravwell service via standard windows service management.


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
### Gravwell Ingester Configuration
**Sample Sysmon config:**  
Create or edit: `%PROGRAMDATA%\gravwell\eventlog\config.cfg`
```ini
[EventChannel "Sysmon"]
        Tag-Name=sysmon
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```
