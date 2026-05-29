# IPMI

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [IPMI Ingester](/ingesters/ipmi.md)
         Kit, [IPMI Kit](https://github.com/gravwell/kits/tree/main/)
:::

## IPMI Configuration
No setup required. The IPMI ingester requires username and password.

```{note} 
We recommend referring to your device's configuration for best practices on setting up a read-only or restricted user account.
```

## Gravwell Configuration

Install the gravwell IPMI Ingester which collects Sensor Data Record (SDR) and System Event Log (SEL) records from any number of IPMI devices. The configuration file provides a simple host/port, username, and password field for connecting to each IPMI device. 

SEL and SDR records are ingested in a JSON-encoded schema. For example:
```json
{
    "Type": "SDR",
    "Target": "10.10.10.10:623",
    "Data": {
        "+3.3VSB": {
            "Type": "Voltage",
            "Reading": "3.26",
            "Units": "Volts",
            "Status": "ok"
        },
        "+5VSB": {},
        "12V": {}
    }
}

{
    "Target": "10.10.10.10:623",
    "Type": "SEL",
    "Data": {
        "RecordID": 25,
        "RecordType": 2,
        "Timestamp": {
            "Value": 1506550240
        },
        "GeneratorID": 32,
        "EvMRev": 4,
        "SensorType": 5,
        "SensorNumber": 81,
        "EventType": 111,
        "EventDir": 0,
        "EventData1": 240,
        "EventData2": 255,
        "EventData3": 255
    }
}
```


### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/ipmi-well.conf`
```ini
[Storage-Well "IPMI"]
    Location=/opt/gravwell/storage/ipmi
    Tags=ipmi*
```
### Gravwell Ingester Configuration
**Sample IPMI config:**  
Create or edit: `/opt/gravwell/etc/ipmi.conf.d/ipmi.conf`
```ini
[IPMI "Server 1"]
    Target="1.2.3.4:623"
    Username="user"
    Password="pass"
    Tag-Name=ipmi
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_ipmi.service`
```