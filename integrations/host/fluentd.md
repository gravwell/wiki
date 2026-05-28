# Fluentd

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
         Kit, [Fluentd Kit](https://github.com/gravwell/kits/tree/main/fluentd)
:::

## Fluentd Configuration

In `/etc/fluent/fluentd.conf` a stanza will need to be added. Remember to change the endpoint to point to your Gravwell Server
* `endpoint` will point to the http ingester

**Sample Fluentd Configuration pointing to Gravwell Environment**
Create or `/etc/fluent/fluentd.conf`
```
<match **>
    @type http
    endpoint http://172.20.0.1:8080/fluentd
    open_timeout 2

    <format>
        @type json
    </format>
    <buffer>
        flush_interval 10s
    </buffer>
</match>
```

For additional fluentd options, e.g. using authentication see: [fluentd.org](https://docs.fluentd.org/output/http)

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fluentd.conf`
```ini
# Fortinet 
[Storage-Well "fortinet"]
    Location=/opt/gravwell/storage/fortinet
    Tags=fortinet*
    Accelerator-Name=fulltext #fulltext is the most resilent to varying data types
    Accelerator-Args="-ignoreFloat" #tell the fulltext accelerator to not index timestamps, syslog entries are easy to ID
    Accelerator-Engine-Override=bloom #The bloom engine is effective and fast with minimal disk overhead
    #this well to delete old data when the disk reaches 90% full
    Hot-Storage-Reserve=10 # adapt this for your environment's requirements
    Delete-Cold-Data=true # adapt this for your environment's requirements
```
### Gravwell Ingester Configuration
**Sample HTTP config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/fluentd.conf`
```ini
[Listener "fluentd"]
    URL="/fluentd"
    Tag-Name=fluentd
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```