# Syslog

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
    Preprocessor, [Syslogrouter Preprocessor](/ingesters/preprocessors/syslogrouter.md)
    Kit, [Syslog Kit](https://github.com/gravwell/kits/tree/main/syslog)  <br /> [Linux Syslog Kit](https://github.com/gravwell/kits/tree/main/linux_syslog)
:::

## Syslog Configuration

This page provides assistance if you have a system that is not currently listed on the integrations page but supports syslog. Each system is going to have its own configuration file to send syslog remotely. There will generally be three settings that need to be configured:
* Transport/Protocol: UDP/TCP 
    * Configure Gravwell to use the same protocol. The syslog default is UDP.
* Hostname/Remote IP
    * Set to the IP address of your Gravwell ingester
* Remote Port: 514
    * Configure Gravwell to use the same port. The default syslog port is 514, and the default for TLS is 6514. The sample below listens on UDP 514 and TCP 601.

## Gravwell Configuration

### Gravwell Storage Well Configuration
**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/syslog-well.conf`
```ini
[Storage-Well "syslog"]
    Location=/opt/gravwell/storage/syslog
    Tags=syslog*
```

### Gravwell Ingester Configuration: Simple Relay
**Sample Syslog config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/syslog.conf`
```ini
[Listener "syslogtcp"]
    Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
    Reader-Type=rfc5424
    Tag-Name=syslog
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
    Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry

[Listener "syslogudp"]
    Bind-String="udp://0.0.0.0:514" #standard UDP based RFC5424 syslog
    Reader-Type=rfc5424
    Tag-Name=syslog
    Timezone-Override="America/Chicago"
    Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```