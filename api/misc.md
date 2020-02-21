# Miscellaneous APIs

Some APIs don't fit nicely into the main categories. They are listed here.

## Tag List

The webserver maintains a list of all tags known to the indexers. This list can be fetched with a GET request on `/api/tags`. This will return a list of tags:

```
["default", "gravwell", "pcap", "windows"]
```

## Search Module List

To get a list of all available search modules and some info about each one, do a GET on `/api/info/searchmodules/`. This will return a list of module info structures:

```
[
    {
        "Collapsing": true,
        "Examples": [
            "min by src",
            "min by someKey"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "min",
        "Sorting": true
    },
    {
        "Collapsing": true,
        "Examples": [
            "unique",
            "unique chuck",
            "unique chuck,testa"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "unique",
        "Sorting": false
    },
[...]
    {
        "Collapsing": false,
        "Examples": [
            "alias src dst"
        ],
        "FrontendOnly": false,
        "Info": "Alias enumerated values",
        "Name": "alias",
        "Sorting": false
    },
    {
        "Collapsing": true,
        "Examples": [
            "count",
            "count by chuck",
            "count by src",
            "count by someKey"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "count",
        "Sorting": true
    }
]
```

## Render Module List

To get a list of all available render modules and some info about each one, do a GET on `/api/info/rendermodules/`. This will return a list of module info structures:

```
[
    {
        "Description": "A raw entry storage system, it can store and handle anything.",
        "Examples": [
            "raw"
        ],
        "Name": "raw",
        "SortRequired": false
    },
    {
        "Description": "A chart storage system system.\n\t       Chart looks for numeric types, storing them.\n\t       Requested entries will be a set of types with column names.",
        "Examples": [
            "chart"
        ],
        "Name": "chart",
        "SortRequired": false
    },
[...]
    {
        "Description": "A point mapping system that supports condensing and geofencing",
        "Examples": [],
        "Name": "point2point",
        "SortRequired": false
    }
]
```

## GUI Settings

This API provides some basic information for the user interface. A GET on `/api/settings` will return a structure similar to the following:

```
{
    "DisableMapTileProxy": false,
    "DistributedWebservers": false,
    "MapTileUrl": "http://localhost:8080/api/maps",
    "MaxFileSize": 8388608,
    "MaxResourceSize": 134217728
}
```

* `DisableMapTileProxy`, if true, tells the UI that it should send map requests directly to OpenStreetMap servers, rather than using the Gravwell proxy.
* `MapTileUrl` is the URL which the UI should use to fetch map tiles.
* `DistributedWebservers` will be set to true if there are multiple webservers coordinating via a datastore.
* `MaxFileSize` is the maximum allowable file size (in bytes) which may be uploaded to the `/api/files` APIs.
* `MaxResourceSize` is the maximum allowable resource size, in bytes.
