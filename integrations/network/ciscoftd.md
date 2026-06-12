# Cisco FTD

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay.md)
         Kit, [Cisco FTD Kit](https://github.com/gravwell/kits/tree/main/ciscoftd)
:::

## Cisco FTD Configuration

Configure log forwarding as described in [Cisco FTD documentation](https://www.cisco.com/c/en/us/support/docs/security/firepower-ngfw/200479-Configure-Logging-on-FTD-via-FMC.html) 

Things to note as you follow the logging setup:
* Enable EMBLEM format
* Set IP address and port

```{warning}
If using TCP for syslog you probably want to check the `Allow user traffic to pass when TCP syslog server is down` check box otherwise if the FTD is unable to connect to the Gravwell ingester it will block All new connections.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/cisco-ftd-well.conf`
```ini
[Storage-Well "ciscoftd"]
    Location=/opt/gravwell/storage/cisco-ftd
    Tags=cisco-ftd*
```
### Gravwell Ingester Configuration: Simple Relay
**Sample Cisco FTD config:**  
Create or edit: `/opt/gravwell/etc/simple_relay/cisco-ftd.conf`
```ini
[Listener "syslogtcp_cisco_ftd"]
    Bind-String="tcp://0.0.0.0:6901"
    Reader-Type=rfc5424
    Tag-Name=cisco-ftd-events
    Assume-Local-Timezone=true
    Preprocessor="Cisco FTD 43000X Router"
    Preprocessor="Cisco FTD Class Router"

# Route 43000X security-event syslogs 
[preprocessor "Cisco FTD 43000X Router"]
    Type=regexrouter
    Drop-Misses=false
    Regex=`%FTD-[0-7]-(?P<msgid>43000[0-9]):`
    Route-Extraction=msgid
    Route=430001:cisco-ftd-intrusion
    Route=430002:cisco-ftd-connection
    Route=430003:cisco-ftd-connection
    Route=430004:cisco-ftd-file
    Route=430005:cisco-ftd-malware

# Route non-43000X messages by 3-digit class prefix
[preprocessor "Cisco FTD Class Router"]
    Type=regexrouter
    Drop-Misses=false
    # Match any FTD message id EXCEPT 43000X (handled above).
    Regex=`%FTD-[0-7]-(?P<class>(?!43000)\d{3})\d{3}:`
    Route-Extraction=class

    # auth
    Route=109:cisco-ftd-auth
    Route=113:cisco-ftd-auth

    # config
    Route=111:cisco-ftd-config
    Route=112:cisco-ftd-config
    Route=208:cisco-ftd-config
    Route=308:cisco-ftd-config

    # vpn
    Route=213:cisco-ftd-vpn
    Route=316:cisco-ftd-vpn
    Route=320:cisco-ftd-vpn
    Route=402:cisco-ftd-vpn
    Route=403:cisco-ftd-vpn
    Route=404:cisco-ftd-vpn
    Route=501:cisco-ftd-vpn
    Route=602:cisco-ftd-vpn
    Route=603:cisco-ftd-vpn
    Route=611:cisco-ftd-vpn
    Route=702:cisco-ftd-vpn
    Route=713:cisco-ftd-vpn
    Route=714:cisco-ftd-vpn
    Route=715:cisco-ftd-vpn
    Route=716:cisco-ftd-vpn
    Route=718:cisco-ftd-vpn
    Route=720:cisco-ftd-vpn
    Route=722:cisco-ftd-vpn

    # traffic
    Route=106:cisco-ftd-traffic
    Route=108:cisco-ftd-traffic
    Route=201:cisco-ftd-traffic
    Route=202:cisco-ftd-traffic
    Route=204:cisco-ftd-traffic
    Route=302:cisco-ftd-traffic
    Route=303:cisco-ftd-traffic
    Route=304:cisco-ftd-traffic
    Route=305:cisco-ftd-traffic
    Route=314:cisco-ftd-traffic
    Route=405:cisco-ftd-traffic
    Route=406:cisco-ftd-traffic
    Route=407:cisco-ftd-traffic
    Route=500:cisco-ftd-traffic
    Route=502:cisco-ftd-traffic
    Route=607:cisco-ftd-traffic
    Route=608:cisco-ftd-traffic
    Route=609:cisco-ftd-traffic
    Route=616:cisco-ftd-traffic
    Route=620:cisco-ftd-traffic
    Route=703:cisco-ftd-traffic
    Route=710:cisco-ftd-traffic

    # threat
    Route=400:cisco-ftd-threat
    Route=401:cisco-ftd-threat
    Route=420:cisco-ftd-threat
    Route=733:cisco-ftd-threat

    # system
    Route=101:cisco-ftd-system
    Route=102:cisco-ftd-system
    Route=103:cisco-ftd-system
    Route=104:cisco-ftd-system
    Route=105:cisco-ftd-system
    Route=199:cisco-ftd-system
    Route=210:cisco-ftd-system
    Route=211:cisco-ftd-system
    Route=214:cisco-ftd-system
    Route=216:cisco-ftd-system
    Route=306:cisco-ftd-system
    Route=307:cisco-ftd-system
    Route=311:cisco-ftd-system
    Route=315:cisco-ftd-system
    Route=414:cisco-ftd-system
    Route=604:cisco-ftd-system
    Route=605:cisco-ftd-system
    Route=606:cisco-ftd-system
    Route=610:cisco-ftd-system
    Route=612:cisco-ftd-system
    Route=614:cisco-ftd-system
    Route=615:cisco-ftd-system
    Route=701:cisco-ftd-system
    Route=709:cisco-ftd-system
    Route=711:cisco-ftd-system
    Route=741:cisco-ftd-system
```


```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```
