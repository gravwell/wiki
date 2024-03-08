---
myst:
  substitutions:
    package: "gravwell-collectd"
    standalone: "gravwell_collectd"
    dockername: "collectd"
---
# collectd Ingester

The collectd ingester is a fully standalone [collectd](https://collectd.org/) collection agent which can directly ship collectd samples to Gravwell.  The ingester supports multiple collectors which can be configured with different tags, security controls, and plugin-to-tag overrides.

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

The collectd ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, the collectd ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/collectd.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/collectd.conf.d`).

The Gravwell collectd ingester is designed to accept the native binary collectd data formats as exported by the `network` plugin which uses the UDP transport.  A basic `network` plugin definition which ships data to a Gravwell ingester might look like so:

```
<Plugin network>
	<Server "10.0.0.70" "25826">
		SecurityLevel Encrypt
		Username "user"
		Password "secret"
		ResolveInterval 14400
	</Server>
	CacheFlush 1800
	MaxPacketSize 1452
	ReportStats false
	TimeToLive 128
</Plugin>
```

## Collector Examples

```
[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	Security-Level=encrypt
	User=user
	Password=secret
	Encoder=json

[Collector "example"]
	Bind-String=10.0.0.1:9999 #default is "0.0.0.0:25826
	Tag-Name=collectdext
	Tag-Plugin-Override=cpu:collectdcpu
	Tag-Plugin-Override=swap:collectdswap
```

## Installation
If you're using the Gravwell Debian repository, installation is just a single apt command:

```
apt-get install gravwell-collectd
```

Otherwise, download the installer from the [Downloads page](/quickstart/downloads). Using a terminal on the Gravwell server, issue the following command as a superuser (e.g. via the `sudo` command) to install the ingester:

```console
root@gravserver ~ # bash gravwell_collectd_installer.sh
```

If the Gravwell services are present on the same machine, the installation script will automatically extract and configure the `Ingest-Auth` parameter and set it appropriately.  However, if your ingester is not resident on the same machine as a pre-existing Gravwell backend, the installer will prompt for the authentication token and the IP address of the Gravwell indexer. You can set these values during installation or leave them blank and modify the configuration file in `/opt/gravwell/etc/collectd.conf` manually.

## Configuration

The collectd ingester relies on the same Global configuration system as all other ingesters.  The Global section is used for defining indexer connections, authentication, and local cache controls.

Collector configuration blocks are used to define listening collectors which can accept collectd samples.  Each collector configuration can have a unique Security-Level, authentication, tag, source override, network bind, and tag overrides.  Using multiple collector configurations, a single collectd ingester can listen on multiple interfaces and apply unique tags to collectd samples coming in from multiple network enclaves.

By default the collectd ingester reads a configuration file located at _/opt/gravwell/etc/collectd.conf_.

### Example Configuration

```
[Global]
	Ingest-Secret = SuperSecretKey
	Connection-Timeout = 0
	Cleartext-Backend-target=192.168.122.100:4023
	Log-Level=INFO

[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	User=user
	Password=secret

[Collector "localhost"]
	Bind-String=[fe80::1]:25827
	Tag-Name=collectdlocal
	Security-Level=none
	Source-Override=[fe80::beef:1000]
	Tag-Plugin-Override=cpu:collectdcpu
```

### Collector Configuration Options

Each Collector block must contain a unique name and non-overlapping Bind-Strings.  You cannot have multiple Collectors that are bound to the same interface on the same port.

| Parameter            | Type         | Required | Default Value | Description  |
|----------------------|--------------|----------|---------------|--------------|
| Bind-String          | string       | YES      |               |              |
| Tag-Name             | string       | NO       |               | Tag to be assigned to data ingested on this listener |
| Source-Override      | string       | NO       |               | Override the source IP assigned to entries ingested on the listener |
| Security-Level       | string       | YES      |               | Collectd data transport security encoding, must match the value in the network plugin. |
| User                 | string       | YES      |               | Collectd data transport username, must match the value in the network Plugin. |
| Password             | string       | YES      |               | Collectd data transport password, must match the value in network Plugin. |
| Encoder              | string       | NO       | json          | Output data format, default is JSON and published Gravwell kits expect JSON. |
| Tag-Plugin-Override  | string array | NO       |               | Optional set of plugin to tag mappings. |
| Preprocessor         | string array | NO       |               | Set of preprocessors to apply to entries. |



#### Bind-String

Bind-String controls the address and port which the Collector uses to listen for incoming collectd samples.  A valid Bind-String must contain either an IPv4 or IPv6 address and a port.  To listen on all interfaces use the "0.0.0.0" wildcard address.

##### Example Bind-String
```
Bind-String=0.0.0.0:25826
Bind-String=127.0.0.1:25826
Bind-String=127.0.0.1:12345
Bind-String=[fe80::1]:25826
```

#### Tag-Name

Tag-Name defines the tag that collectd samples will be assigned unless a Tag-Plugin-Override applies.

#### Source-Override

The Source-Override directive is used to override the source value applied to entries when they are sent to Gravwell.  By default the ingester applies the Source of the ingester, but it may be desirable to apply a specific source value to a Collector block in order to apply segmentation or filtering at search time.  A Source-Override is any valid IPv4 or IPv6 address.

#### Example Source-Override
```
Source-Override=192.168.1.1
Source-Override=[DEAD::BEEF]
Source-Override=[fe80::1:1]
```

#### Security-Level

The Security-Level directive controls how the Collector authenticates collectd packets.  Available options are: encrypt, sign, none.  By default a Collector uses the "encrypt" Security-Level and requires that both a User and Password are specified.  If "none" is used, no User or Password is required.

#### Example Security-Level
```
Security-Level=none
Security-Level=encrypt
Security-Level = sign
Security-Level = SIGN
```

#### User and Password

When the Security-Level is set as "sign" or "encrypt" a username and password must be provided that match the values set in endpoints.  The default values are "user" and "secret" to match the default values shipped with collectd.  These values should be changed when collectd data might contain sensitive information.

##### User and Password Examples
```
User=username
Password=password
User = "username with spaces in it"
Password = "Password with spaces and other characters @$@#@()*$#W)("
```

#### Encoder

The default collectd encoder is JSON, but a simple text encoder is also available.  Options are "JSON" or "text"

An example entry using the JSON encoder:

```
{"host":"build","plugin":"memory","type":"memory","type_instance":"used","value":727789568,"dsname":"value","time":"2018-07-10T16:37:47.034562831-06:00","interval":10000000000}
```

## Tag Plugin Overrides

Each Collector block supports N number of Tag-Plugin-Override declarations which are used to apply a unique tag to a collectd sample based on the plugin that generated it.  Tag-Plugin-Overrides can be useful when you want to store data coming from different plugins in different wells and apply different ageout rules.  For example, it may be valuable to store collectd records about disk usage for 9 months, but CPU usage records can expire out at 14 days.  The Tag-Plugin-Override system makes this easy.

The Tag-Plugin-Override format is comprised of two strings separated by the ":" character.  The string on the left represents the name of the plugin and the string on the right represents the name of the desired tag.  All the usual rules about tags apply.  A single plugin cannot be mapped to multiple tags, but multiple plugins CAN be mapped to the same tag.

### Example Tag Plugin Overrides
```
Tag-Plugin-Override=cpu:collectdcpu # Map CPU plugin data to the "collectdcpu" tag.
Tag-Plugin-Override=memory:memstats # Map the memory plugin data to the "memstats" tag.
Tag-Plugin-Override= df : diskdata  # Map the df plugin data to the "diskdata" tag.
Tag-Plugin-Override = disk : diskdata  # Map the disk plugin data to the "diskdata" tag.
```


## Example Collect Configuration

The Collectd system is a plugin based system instrumentation and testing framework, there is no standard Collectd deployment and every plugin can send a unique set of fields and structures.   The only hard requirement for configuring the Collectd system with Gravwell is a proper `network` Plugin definition with matching username, password, and Security-Level.  Here are two basic configurations that will collect some reasonable metrics and send them to Gravwell:


### /etc/collectd/collectd.conf

```
Hostname "server.example.com"
FQDNLookup false

AutoLoadPlugin true

CollectInternalStats false

Interval 10

<Plugin network>
	<Server "10.0.0.70" "25826">
		SecurityLevel Encrypt
		Username "user"
		Password "secret"
		ResolveInterval 14400
	</Server>
	CacheFlush 1800
	MaxPacketSize 1452
	ReportStats false
	TimeToLive 128
</Plugin>

<Plugin cpu>
	ReportByCpu true
	ReportByState true
	ValuesPercentage false
	ReportNumCpu false
	ReportGuestState false
	SubtractGuestState true
</Plugin>

<Plugin df>
	# ignore rootfs; else, the root file-system would appear twice, causing
	# one of the updates to fail and spam the log
	FSType rootfs
	# ignore the usual virtual / temporary file-systems
	FSType sysfs
	FSType proc
	FSType devtmpfs
	FSType devpts
	FSType tmpfs
	FSType fusectl
	FSType cgroup
	IgnoreSelected true

	ValuesPercentage true
</Plugin>

<Plugin ethstat>
	Interface "eno1"
	Map "rx_csum_offload_errors" "if_rx_errors" "checksum_offload"
	Map "multicast" "if_multicast"
	MappedOnly false
</Plugin>

<Plugin load>
	ReportRelative true
</Plugin>

<Plugin memory>
	ValuesAbsolute false
	ValuesPercentage true
</Plugin>
```

### /opt/gravwell/etc/collectd.conf

```
[Global]
Ingester-UUID="5cb70d8d-3800-4044-bae1-308f00b6f7b5"
Ingest-Secret = "SuperHardSecrets"
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Cleartext-Backend-Target=10.0.0.42 #example of adding a cleartext connection
Ingest-Cache-Path=/opt/gravwell/cache/collectd.cache
Max-Ingest-Cache=1024 #Number of MB to store, localcache will only store 1GB before stopping.  This is a safety net
Log-Level=INFO
Log-File=/opt/gravwell/log/collectd.log

[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	Security-Level=encrypt
	User=user
	Password=secret
```
