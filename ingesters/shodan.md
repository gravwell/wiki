# Shodan Ingester

The Shodan Ingester collects data the Shodan Streaming API. It supports both the "full firehose" banners API, which includes all data that Shodan collects, and the "alerts" API, which is a filtered form of the banners API. 

More information on the structure of the Shodan Streaming API can be found [here](https://developer.shodan.io/api/stream). 

A Shodan entry is a JSON-formatted record containing the elements for a single Shodan collection event. For example (taken from the Shodan API website):

```
{
    "_shodan": {
        "id": "7383056c-d513-4b43-8734-b82d897888e6",
        "options": {},
        "ptr": true,
        "module": "dns-udp",
        "crawler": "9d8ac08f91f51fa9017965712c8fdabb4211dee4"
    },
    "hash": -553166942,
    "os": null,
    "opts": {
        "raw": "34ef818200010000000000000776657273696f6e0462696e640000100003"
    },
    "ip": 134744072,
    "isp": "Google",
    "port": 53,
    "hostnames": [
        "dns.google"
    ],
    "location": {
        "city": null,
        "region_code": null,
        "area_code": null,
        "longitude": -97.822,
        "country_code3": null,
        "country_name": "United States",
        "postal_code": null,
        "dma_code": null,
        "country_code": "US",
        "latitude": 37.751
    },
    "dns": {
        "resolver_hostname": null,
        "recursive": true,
        "resolver_id": null,
        "software": null
    },
    "timestamp": "2021-01-28T07:21:33.444507",
    "domains": [
        "dns.google"
    ],
    "org": "Google",
    "data": "\nRecursion: enabled",
    "asn": "AS15169",
    "transport": "udp",
    "ip_str": "8.8.8.8"
}
```

More information about the Shodan Record format is available at the [Shodan datapedia site](https://datapedia.shodan.io).

## Basic Configuration

The Shodan ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the Shodan ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The Shodan Ingester requires a [Shodan API key](https://developer.shodan.io/api/requirements). Using the API key, you can configure one or more Shodan readers. For example:

```
[ShodanAccount "shodan"]
	API-Key=YOUR-KEY-HERE
	Tag-Name=shodan
	Module-Tags-Prefix=shodan-	# modules extracted separately will be tagged `shodan-<module>`
	Extracted-Modules=http
	Extracted-Modules=https
	Extracted-Modules=ssh
	Extract-All-Modules=false
	Full-Firehose=true		# consume the banners API instead of the alerts API.
```

The ShodanAccount stanza requires the `API-Key` and either (or both) the `Tag-Name` and `Module-Tags-Prefix` fields. All others are optional. The `Module-Tags-Prefix` field sets the prefix for tags to be created for each type of Shodan module (for example shodan-http, shodan-dns). Additionally, you can specify which modules to extract into separate tags with one or more `Extracted-Modules` fields, or extract all modules with the `Extract-All-Modules` field.

By default, the Shodan ingester reads from the "alerts" API, which must first be [setup using your API key](https://developer.shodan.io/api/stream) to provide a filtered set of events to ingest. If you instead want to ingest the full firehose of Shodan events, set the `Full-Firehose` field to true. This will cause the Shodan Ingester to instead read from the "banners" API. 
