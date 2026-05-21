# NetFlow

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [NetFlow Ingester](/ingesters/netflow.md)
         Kit, [NetFlow v5 Kit](https://github.com/gravwell/kits/tree/main/netflowv5)
:::

## NetFlow Configuration

Each system is going to have its own configuration file to send netflow remotely. There will generally be three settings that need to be configured:
* Interfaces
    * Specify the interfaces that you want captured by NetFlow
    * Some devices will allow you to set WAN interfaces to avoid duplicating traffic
* Version
    * v5 
    * v9 see: [IPFIX](ipfix.md)
* Destination
    * Set to the IP address of your Gravwell ingester

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/netflow-well.conf`
```ini
[Storage-Well "netflow"]
    Location=/opt/gravwell/storage/netflow
    Tags=netflow*
```

### Gravwell Ingester Configuration
**Sample NetFlow config:**  
Create or edit: `/opt/gravwell/etc/netflow_capture.conf.d/netflow.conf`
```ini
[Collector "netflow v5"]
    Bind-String="0.0.0.0:2055" #we are binding to all interfaces
    Tag-Name=netflow
    Assume-Local-Timezone=true
    Session-Dump-Enabled=true
```
