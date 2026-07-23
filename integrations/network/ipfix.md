# IPFIX

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [NetFlow Ingester](/ingesters/netflow)
         Kit, [IPFIX Kit](https://github.com/gravwell/kits/tree/main/ipfix)
:::

## IPFIX Configuration

Each system is going to have its own configuration file to send ipfix remotely. There will generally be three settings that need to be configured:
* Interfaces
    * Specify the interfaces that you want captured by IPFix
    * Some devices will allow you to set WAN interfaces to avoid duplicating traffic
* Version
    * v5 see: [Netflow](netflow.md)
    * v9
* Destination
    * Set to the IP address of your Gravwell ingester

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/ipfix.well`
```ini
[Storage-Well "ipfix"]
    Location=/opt/gravwell/storage/ipfix
    Tags=ipfix*
```

### Gravwell Ingester Configuration: Netflow
**Sample IPFIX config:**  
Create or edit: `/opt/gravwell/etc/netflow_capture.conf.d/ipfix.conf`
```ini
[Collector "ipfix"]
    Tag-Name=ipfix
    Bind-String="0.0.0.0:4739"
    Flow-Type=ipfix
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_netflow_capture.service`
```