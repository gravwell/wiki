# Cisco ASA

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay.md)
         Kit, [Cisco ASA Kit](https://github.com/gravwell/kits/tree/main/ciscoasa)
:::

## Cisco ASA Configuration

Configure log forwarding as described in [Cisco ASA documentation](https://www.cisco.com/c/en/us/support/docs/security/pix-500-series-security-appliances/63884-config-asa-00.html#toc-hId-68106104) 

Example Cisco ASA config:
```
logging host interface_name simple_relay_ip udp/514 format emblem
logging trap severity_level
logging facility number
```

```{warning}
If using TCP for syslog you probably want to set `logging permit-hostdown` otherwise if the ASA is unable to connect to the Gravwell ingester it will block All new connections.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/cisco-asa-well.conf`
```ini
[Storage-Well "ciscoasa"]
    Location=/opt/gravwell/storage/cisco-asa
    Tags=cisco-asa*
```
### Gravwell Ingester Configuration
**Sample Cisco ASA config:**  
Create or edit: `/opt/gravwell/etc/simple_relay/cisco-asa.conf`
```ini
[Listener "syslogtcp_cisco_asa"]
    Bind-String="tcp://0.0.0.0:6801"
    Reader-Type=rfc5424
    Keep-Priority=true
    Tag-Name=cisco-asa-events
    Assume-Local-Timezone=true
    Preprocessor="Cisco ASA Class Router"

# ASA: Route by 3-digit class prefix from the 6-digit message number
# Example: %ASA-6-302013: ...  -> class=302
[preprocessor "Cisco ASA Class Router"]
    Type=regexrouter
    Drop-Misses=false
    Regex=`%ASA-[0-7]-(?P<class>\d{3})\d{3}:`
    Route-Extraction=class

    # auth
    Route=109:cisco-asa-auth
    Route=113:cisco-asa-auth

    # config
    Route=111:cisco-asa-config
    Route=112:cisco-asa-config
    Route=208:cisco-asa-config
    Route=308:cisco-asa-config

    # vpn
    Route=213:cisco-asa-vpn
    Route=316:cisco-asa-vpn
    Route=320:cisco-asa-vpn
    Route=402:cisco-asa-vpn
    Route=403:cisco-asa-vpn
    Route=404:cisco-asa-vpn
    Route=501:cisco-asa-vpn
    Route=602:cisco-asa-vpn
    Route=603:cisco-asa-vpn
    Route=611:cisco-asa-vpn
    Route=702:cisco-asa-vpn
    Route=713:cisco-asa-vpn
    Route=714:cisco-asa-vpn
    Route=715:cisco-asa-vpn
    Route=716:cisco-asa-vpn
    Route=718:cisco-asa-vpn
    Route=720:cisco-asa-vpn
    Route=722:cisco-asa-vpn

    # traffic
    Route=106:cisco-asa-traffic
    Route=108:cisco-asa-traffic
    Route=201:cisco-asa-traffic
    Route=202:cisco-asa-traffic
    Route=204:cisco-asa-traffic
    Route=302:cisco-asa-traffic
    Route=303:cisco-asa-traffic
    Route=304:cisco-asa-traffic
    Route=305:cisco-asa-traffic
    Route=314:cisco-asa-traffic
    Route=405:cisco-asa-traffic
    Route=406:cisco-asa-traffic
    Route=407:cisco-asa-traffic
    Route=500:cisco-asa-traffic
    Route=502:cisco-asa-traffic
    Route=607:cisco-asa-traffic
    Route=608:cisco-asa-traffic
    Route=609:cisco-asa-traffic
    Route=616:cisco-asa-traffic
    Route=620:cisco-asa-traffic
    Route=703:cisco-asa-traffic
    Route=710:cisco-asa-traffic

    # threat
    Route=400:cisco-asa-threat
    Route=401:cisco-asa-threat
    Route=420:cisco-asa-threat
    Route=733:cisco-asa-threat

    # system
    Route=101:cisco-asa-system
    Route=102:cisco-asa-system
    Route=103:cisco-asa-system
    Route=104:cisco-asa-system
    Route=105:cisco-asa-system
    Route=199:cisco-asa-system
    Route=210:cisco-asa-system
    Route=211:cisco-asa-system
    Route=214:cisco-asa-system
    Route=216:cisco-asa-system
    Route=306:cisco-asa-system
    Route=307:cisco-asa-system
    Route=311:cisco-asa-system
    Route=315:cisco-asa-system
    Route=414:cisco-asa-system
    Route=604:cisco-asa-system
    Route=605:cisco-asa-system
    Route=606:cisco-asa-system
    Route=610:cisco-asa-system
    Route=612:cisco-asa-system
    Route=614:cisco-asa-system
    Route=615:cisco-asa-system
    Route=701:cisco-asa-system
    Route=709:cisco-asa-system
    Route=711:cisco-asa-system
    Route=741:cisco-asa-system
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```