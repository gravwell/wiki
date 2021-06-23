# Windows Event Service

The Gravwell Windows events ingester runs as a service on a Windows machine and sends Windows events to the Gravwell indexer.  The ingester consumes from the `System`, `Application`, `Setup`, and `Security` channels by default.  Each channel can be configured to consume from a specific set of events or providers.

## Basic Configuration

The Windows Event ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the Windows Evenet ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## EventChannel Examples

```
[EventChannel "system"]
	Tag-Name=windows
	Channel=System #pull from the system channel

[EventChannel "sysmon"]
	Tag-Name=sysmon
	Channel="Microsoft-Windows-Sysmon/Operational"
	Max-Reachback=24h  #reachback must be expressed in hours (h), minutes (m), or seconds(s)

[EventChannel "Application"]
	Channel=Application #pull from the application channel
	Tag-Name=winApp #Apply a new tag name
	Provider=Windows System #Only look for the provider "Windows System"
	EventID=1000-4000 #Only look for event IDs 1000 through 4000
	Level=verbose #Only look for verbose entries
	Max-Reachback=72h #start looking for logs up to 72 hours in the past
	Request_Buffer=16 #use a large 16MB buffer for high throughput
	Request_Size=1024 #Request up to 1024 entries per API call for high throughput

[EventChannel "System Critical and Error"]
	Channel=System #pull from the system channel
	Tag-Name=winSysCrit #Apply a new tag name
	Level=critical #look for critical entries
	Level=error #AND for error entries
	Max-Reachback=96h #start looking for logs up to 96 hours in the past

[EventChannel "Security prune"]
	Channel=Security #pull from the security channel
	Tag-Name=winSec #Apply a new tag name
	EventID=-400 #ignore event ID 400
	EventID=-401 #AND ignore event ID 401
```

## Installation

Download the Gravwell Windows ingester installer from the [Downloads page](#!quickstart/downloads.md).

Run the .msi installation wizard to install the Gravwell events service.  On first installation the installation wizard will prompt to configure the indexer endpoint and ingest secret.  Subsequent installations and/or upgrades will identify a resident configuration file and will not prompt.

The ingester is configured with the `config.cfg` file located at `%PROGRAMDATA%\gravwell\eventlog\config.cfg`.  The configuration file follows the same form as other Gravwell ingesters with a `[Global]` section configuring the indexer connections and multiple `EventChannel` definitions.

To modify the indexer connection or specify multiple indexers, change the connection IP address to the IP of your Gravwell server and set the Ingest-Secret value.  This example shows configuring an encrypted transport:

```
Ingest-Secret=YourSecretGoesHere
Encrypted-Backend-target=ip.addr.goes.here:port
```

Once configured, this file can be copied to any other Windows system from which you would like to collect events.

### Silent Installation

The Windows event ingester is designed to be compatible with an automated deployment.  This means that a domain controller can push the installer to clients and invoke installation without user interaction.  To force a silent installation execute the installer with administrative privileges via [msiexec](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/msiexec) with the `/quiet` argument.  This installation method will install the default configuration and start the service.

To configure your specific parameters you will then need to either push a modified configuration file to `%PROGRAMDATA%\gravwell\eventlog\config.cfg` and restart the service, or also provide the `CONFIGFILE` argument with the fully qualified path to the `config.cfg` file.

Note that you may need to create the `%PROGRAMDATA%\gravwell\eventlog` path.

A complete execution sequence for a Group Policy push might look like:

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet
xcopy \\share\gravwell_config.cfg %PROGRAMDATA%\gravwell\eventlog\config.cfg
sc stop "GravwellEvent Service"
sc start "GravwellEvent Service"
```

Or

```
msiexec.exe /i gravwell_win_events_3.3.12.msi /quiet CONFIGFILE=\\share\gravwell_config.cfg
```

## Optional Sysmon Integration

The Sysmon utility, part of the sysinternals suite, is an effective and popular tool for monitoring Windows systems. There are plenty of resources with examples of good sysmon configuration files. At Gravwell, we like to use the config created by infosec Twitter personality @InfosecTaylorSwift.

Edit the Gravwell Windows agent config file located at `%PROGRAMDATA%\gravwell\eventlog\config.cfg` and add the following lines:

```
[EventChannel "Sysmon"]
        Tag-Name=sysmon #Apply a new tag name
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

[Download the excellent sysmon configuration file by SwiftOnSecurity](https://raw.githubusercontent.com/SwiftOnSecurity/sysmon-config/master/sysmonconfig-export.xml)

[Download sysmon](https://technet.microsoft.com/en-us/sysinternals/sysmon)

Install `sysmon` with your configuration using an administrator shell (Powershell works too) by running the following command:

```
sysmon.exe -accepteula -i sysmonconfig-export.xml
```

Restart the Gravwell service via standard windows service management.

### Example Configuration with Sysmon

```
[EventChannel "system"]
        Tag-Name=windows
        #no Provider means accept from all providers
        #no EventID means accept all event ids
        #no Level means pull all levels
        #no Max-Reachback means look for logs starting from now
        Channel=System #pull from the system channel

[EventChannel "application"]
        Tag-Name=windows
        Channel=Application #pull from the system channel

[EventChannel "security"]
        Tag-Name=windows
        Channel=Security #pull from the system channel

[EventChannel "setup"]
        Tag-Name=windows
        Channel=Setup #pull from the system channel

[EventChannel "sysmon"]
        Tag-Name=windows
        Provider=Microsoft-Windows-Sysmon #Only look for the provider
        Channel=Microsoft-Windows-Sysmon/Operational
```

## Troubleshooting

You can verify the Windows ingester connectivity by navigating to the Ingester page on the web interface.  If the Windows ingester is not present, check the status of the service either via the windows GUI or by running `sc query GravwellEvents` at the command line.

![](querystatus.png)

![](querystatusgui.png)

## Example Windows Searches

Assuming the default tag names are used, to see ALL sysmon entries in their entirety run this search:

```
tag=sysmon
```

To see ALL Windows events in their entirety run:

```
tag=windows
```

For the following searches we can use the `winlog` search module to filter and extract specific events and fields.  To see all network creation by non-standard processes:

```
tag=sysmon regex winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
table TIMESTAMP SourceHostname Image DestinationIP DestinationPort
```

To chart network creation by source host:

```
tag=sysmon regex winlog EventID==3 Image SourceHostname DestinationIp DestinationPort |
count by SourceHostname |
chart count by SourceHostname limit 10
```

To see suspicious file creation:

```
tag=sysmon winlog EventID==11 Image TargetFilename |
count by TargetFilename |
chart count by TargetFilename
```
