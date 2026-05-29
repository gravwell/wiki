# Duo

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Fetcher](https://github.com/gravwell/gravwell/blob/main/experiments/gravwell_fetcher/README.md)
         Kit, [Duo Kit](https://github.com/gravwell/kits/tree/main/duo)
:::

## Duo Configuration

You will need to collect a Domain, API Key, and secret from Duo. These can be gathered by following Duo's documentation: [Duo Admin API](https://duo.com/docs/adminapi#overview)

## Gravwell Configuration

The Gravwell Fetcher provides a lightweight Go-based fetcher that polls external APIs and ingests events into Gravwell. 
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

5. Edit `_/etc/gravwell/gravwell\_fetcher.conf_` and replace the duo stanzas (see example below).  

6. Run the fetcher (from the built binary).  
    Typical invocation (binary + config file):  
    `./gravwell\_fetcher -config /etc/gravwell/gravwell\_fetcher.conf`

```{attention}
The canonical example config shipped with the experiment is gravwell_fetcher.conf.example — copy it and update the values for Duo.
```

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/duo-well.conf`
```ini
[Storage-Well "duo"]
    Location=/opt/gravwell/storage/duo
    Tags=duo*
```
### Gravwell Fetcher Configuration

Setup the fetcher configuration file.

**Sample Duo config:**  
Create or edit: `/opt/gravwell/etc/gravwell_fetcher.conf.d/duo.conf`
```ini
[DuoConf "duo-admin"]
    StartTime="2025-01-01T00:00:01.000Z"  # Initial fetch time
    Domain=""                             # Duo domain
    Key=""                                # Duo API key
    Secret=""                             # Duo API secret
    DuoAPI="admin"                        # API type: admin, authentication, activity
    Tag-Name="duo-admin"                  # Tag for Gravwell

[DuoConf "duo-auth"]
    StartTime="2025-01-01T00:00:01.000Z"
    Domain=""
    Key=""
    Secret=""
    DuoAPI="authentication"
    Tag-Name="duo-auth"

[DuoConf "duo-activity"]
    StartTime="2025-01-01T00:00:01.000Z"
    Domain=""
    Key=""
    Secret=""
    DuoAPI="activity"
    Tag-Name="duo-activity"

[DuoConf "duo-account"]
    StartTime="2025-01-01T00:00:01.000Z"
    Domain=""
    Key=""
    Secret=""
    DuoAPI="activity"
    Tag-Name="duo-account"
```
