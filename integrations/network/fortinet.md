# Fortinet

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay.md)
         Kit, [Gravell Fortinet](https://github.com/gravwell/kits/tree/main/fortinet)
:::

## Fortinet Configuration

To get logs flowing from your Fortinet FortiGate/FortiOS 7.6.6 device configure remote syslog logging as described in the Fortinet documentation (Log setting and target)[https://docs.fortinet.com/document/fortigate/7.6.6/administration-guide/250999/log-settings-and-targets#Remote_logging]. 

Recommended FortiGate syslog settings:

* Use mode reliable (RFC6587 over TCP). (config log syslogd setting)
* Use format rfc5424 (best timestamp framing) or format default (simple key=value).
* Point the syslog server to the Simple Relay host on port 6701.

Example FortiGate CLI config:
```
config log syslogd setting
    set status enable
    set server <SIMPLE_RELAY_IP>
    set mode reliable
    set port 6701
    set format rfc5424
end
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fortinet-well.conf`
```ini
[Storage-Well "fortinet"]
    Location=/opt/gravwell/storage/fortinet
    Tags=fortinet*
```
### Gravwell Ingester Configuration
**Sample Fortinet config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/fortinet.conf`
```ini
    [Listener "syslogtcp_fortinet"]
        Bind-String="tcp://0.0.0.0:6701"
        Reader-Type=rfc5424
        Tag-Name=fortinet-events
        Assume-Local-Timezone=true
        Preprocessor="Fortinet Type Router"
        Preprocessor="Fortinet System Router"

    [preprocessor "Fortinet Type Router"]
        Type=regexrouter
        Drop-Misses=false
        Regex=`\btype="(?P<type>traffic|utm|event?)?\"`
        #Regex=`^[^.]+\s[^.]+\s[^.]+\stype\=\"(?<type>.+?)?\"`
        Route-Extraction=type
        Route=traffic:fortinet-traffic
        Route=utm:fortinet-utm
        # event -> stays on default tag fortinet-events

    [preprocessor "Fortinet System Router"]
        Type=regexrouter
        Drop-Misses=false
        Regex=`\bsubtype="(?P<subtype>.+?)?\"`
        Route-Extraction=subtype
        Route=system:fortinet-system
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```