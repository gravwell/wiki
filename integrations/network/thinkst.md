# Thinkst

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Fetcher](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/README.md)
    Kit, [Thinkst Kit](https://github.com/gravwell/kits/tree/main/thinkstcanary)
:::

## Thinkst Configuration

Collect an API token and domain from your Thinkst Canary console. These can be gathered by following Canary's documentation: [How does the API work?](https://help.canary.tools/hc/en-gb/articles/360012727537-How-does-the-API-work)

## Gravwell Configuration

The Gravwell Fetcher provides a lightweight Go-based fetcher that polls external APIs (including Thinkst Canary endpoints) and ingests events into Gravwell. 
The Fetcher includes an [example configuration file](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/gravwell_fetcher.conf.example) which you need to copy and adapt for your environment prior to running the fetcher. See the [README](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/README.md) for further information. 

### Basic Installation Steps (Example)

1. Clone the Gravwell repo (or just the experiment):  
    `git clone https://github.com/gravwell/gravwell.git`

2. Change directory to the fetcher experiment:  
    `cd gravwell/experiments/gravwell_fetcher`

3. Build the fetcher binary (standard Go build):  
    `go build -o gravwell_fetcher`

4. Copy the example config to a location you will edit  
    for example `/opt/gravwell/etc/gravwell_fetcher.conf`:  
    `cp gravwell_fetcher.conf.example /opt/gravwell/etc/gravwell_fetcher.conf`

5. Edit `/opt/gravwell/etc/gravwell_fetcher.conf` and replace the Thinkst Canary Domain and Token values (see example below).  

6. Run the fetcher (from the built binary).  
    Typical invocation (binary + config file):  
    `./gravwell_fetcher -config /opt/gravwell/etc/gravwell_fetcher.conf`

```{attention}
The canonical example config shipped with the experiment is gravwell_fetcher.conf.example — copy it and update the values for Thinkst.
```

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/thinkst-well.conf`
```ini
[Storage-Well "thinkst"]
    Location=/opt/gravwell/storage/thinkst
    Tags=thinkst*
```

### Gravwell Ingester Configuration: Fetcher
Set up the Fetcher configuration file. Add the following stanzas to the same configuration file you edited above, since the Fetcher reads a single config file.

**Sample Thinkst config:**  
Edit: `/opt/gravwell/etc/gravwell_fetcher.conf`
```ini
[ThinkstConf "thinkst-audit"]
    ThinkstAPI="audit"                    # API type: audit, incident
    Token=""                              # Thinkst API token
    Domain="XXXXXXXX.canary.tools"        # Your Thinkst domain
    StartTime="2025-01-01T00:00:01.000Z"  # Initial fetch time
    Tag-Name="thinkst-audit"              # Tag for Gravwell

[ThinkstConf "thinkst-incident"]
    ThinkstAPI="incident"
    Token=""
    Domain="XXXXXXXX.canary.tools"
    StartTime="2025-01-01T00:00:01.000Z"
    Tag-Name="thinkst-incident"
```

```{note}
The Fetcher is run manually (see the steps above), so there is no service to restart. Stop the running Fetcher and start it again to apply the new config.
```
