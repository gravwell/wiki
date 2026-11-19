# Fluentd

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
:::

## Fluentd Configuration

In `/etc/fluent/fluentd.conf` a stanza will need to be added. Remember to change the endpoint to point to your Gravwell Server:
* `endpoint` will point to the http ingester

**Sample Fluentd config:**  
Create or edit: `/etc/fluent/fluentd.conf`
```
<match **>
    @type http
    endpoint http://path.to.gravwell:port/fluentd
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
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fluentd-well.conf`
```ini
[Storage-Well "fluentd"]
    Location=/opt/gravwell/storage/fluentd
    Tags=fluentd*
    Accelerator-Name=fulltext #fulltext is the most resilient to varying data types
    Accelerator-Args="-ignoreFloat" #tell the fulltext accelerator not to index floating-point numbers
    Accelerator-Engine-Override=bloom #The bloom engine is effective and fast with minimal disk overhead
    #configure this well to delete old data when the disk reaches 90% full
    Hot-Storage-Reserve=10 # adapt this for your environment's requirements
    Delete-Cold-Data=true # adapt this for your environment's requirements
```

### Gravwell Ingester Configuration: HTTP Ingester

This guide assumes the HTTP Ingester is already installed and configured as described in [HTTP Ingester Setup](/integrations/generic/http_ingester). Add the following to a new overlay file; you do not need a separate HTTP Ingester for this integration.

**Sample Fluentd HTTP config:**  
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