# Thinkst

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Gravwell Fetcher](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/README.md)
         Kit, [Thinkst Kit](https://github.com/gravwell/kits/tree/main/thinkst-canary)
:::

## Thinkst Configuration

You will need to collect an API Key and Domain from your Canary device. These can be gathered by following Canary's documentation: [How does the API work?](https://help.canary.tools/hc/en-gb/articles/360012727537-How-does-the-API-work)

## Gravwell Configuration

The Gravwell Fetcher provides a lightweight Go-based fetcher that polls external APIs (including Thinkst Canary endpoints) and ingests events into Gravwell. 
The Fetcher includes an [example configuration file](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/gravwell_fetcher.conf.example) which you need to copy and adapt for your environment prior to running the fetcher. See the [README](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/README.md) for further information. 

### Basic installation steps (example)

1. Clone the Gravwell repo (or just the experiment):  
    `git clone https://github.com/gravwell/gravwell.git`

2. Change directory to the fetcher experiment:  
    `cd gravwell/experiments/gravwell\_fetcher`

3. Build the fetcher binary (standard Go build):  
    `go build -o gravwell\_fetcher`

4. Copy the example config to a location you will edit  
    e.g. /etc/gravwell/gravwell\_fetcher.conf or /opt/gravwell/etc/gravwell\_fetcher.conf:  
    `cp gravwell\_fetcher.conf.example /etc/gravwell/gravwell\_fetcher.conf`

5. Edit `/etc/gravwell/gravwell\_fetcher.conf` and replace the Thinkst Canary Domain and Token values (see example below).  

6. Run the fetcher (from the built binary).  
    Typical invocation (binary + config file):  
    `./gravwell\_fetcher -config /etc/gravwell/gravwell\_fetcher.conf`

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
### Gravwell Fetcher Configuration

Setup the fetcher configuration file.

**Sample Thinkst config:**  
Create or edit: `/opt/gravwell/etc/gravwell_fetcher.conf.d/thinkst.conf`
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
    Tag-Name="thinks-incident"
```
