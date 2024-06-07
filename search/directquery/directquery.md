# Gravwell Direct Query API

The Gravwell Direct Query API is designed to provide atomic, REST-powered access to the Gravwell query system.  This API provides an option for simple integrations with external tools and systems that do not normally know how to interact with Gravwell.  The API is designed to be as flexible as possible and support tools that know how to interact with an HTTP API.

The Direct Query API is authenticated and requires a valid Gravwell account with access to the Gravwell query system.  Most users will want to generate a [Gravwell Token](/tokens/tokens) and use that to access the query API.

Issuing a query via the Direct Query API requires the same set of parameters as issuing a query via the Gravwell web GUI.  You will need a query string, a time range, and an optional output format.  The Direct Query API has some limitations on which output formats can be provided.  For example, the [pointmap](/search/map/map) and [heatmap](/search/map/map) renderers cannot output rendered maps via this API, nor can this API draw a chart and deliver it as an image.  This API is primarily used for retrieving raw results and delivering them to other systems for direct integration.

```{note}
The Direct Query API is atomic, one request will execute and entire search and deliver the completed results.  Queries that cover large time durations or require significant time to execute may require that HTTP clients adjust their respective client timeouts.
```

## Query Endpoints

The Direct Query API consists of two REST endpoints which can parse a search and execute a search.  The parse API is useful for testing whether a query is valid and could execute while the search API will actually execute a search and deliver the results.  Both the query and parse APIs require the user and/or token to have the `Search` permission.

## Parse API

The parse API is accessed via the POST HTTP method and is located at `/api/parse`.  The parse API requires the following parameters delivered by header values, URL parameters, or the [ParseSearchRequest](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client/types#ParseSearchRequest) JSON object.  The [ParseSearchResponse](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client/types#ParseSearchResponse) object and a 200 code will be returned if the query is valid.

| Parameter | Description |
| --------- | ----------- |
| `query` | A complete Gravwell query string |

The `query` can be delivered as a header value, as query parameter, or as a `ParseSearchRequest` JSON object in the body of the request.

The following [curl](https://curl.se/) commands are all functionally equivalent:

### Using Headers

```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -H "query: tag=gravwell limit 10" \
   http://10.0.0.1/api/parse
```

### Using URL Parameters

```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   http://10.0.0.1/api/parse?query=tag%3Dgravwell%20limit%2010
```

### JSON Object

```
curl -X POST -d '{"SearchString":"tag=gravwell limit 10"}' \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   http://10.0.0.1/api/parse
```

### Example Response

```
{
  "Sequence": 0,
  "GoodQuery": false,
  "ParsedQuery": "tag=gravwell limit 10",
  "RawQuery": "tag=gravwell limit 10",
  "ModuleIndex": 0,
  "CollapsingIndex": 1,
  "RenderModule": "text",
  "TimeZoomDisabled": false,
  "Tags": [
    "gravwell"
  ],
  "ModuleHints": [
    {
      "Name": "limit",
      "ProducedEVs": null,
      "ConsumedEVs": null,
      "ResourcesNeeded": null,
      "Condensing": true
    },
    {
      "Name": "limitCollapser",
      "ProducedEVs": null,
      "ConsumedEVs": null,
      "ResourcesNeeded": null,
      "Condensing": true
    },
    {
      "Name": "sort",
      "ProducedEVs": null,
      "ConsumedEVs": null,
      "ResourcesNeeded": null,
      "Condensing": false
    }
  ]
}
```

## Query API

The query API is accessed via the POST HTTP method and is located at `/api/search/direct`.  The search API requires the following parameters delivered by header values, URL parameters, or a JSON object:

| Parameter | Description | Optional |
| --------- | ----------- | -------- |
| query     | A complete Gravwell query string | |
| format    | Query output format | |
| preview   | Boolean indicating that the query should execute as a preview query | X |
| start     | RFC3339 start timestamp for the query | X |
| end       | RFC3339 end timestamp for the query | X |
| duration  | Golang encoded duration | X |

```{note}
While the `start`, `end`, and `duration` parameters are optional at last one complete set must be provided, either `start` and `end` or `duration`.
```

```{note}
Each query renderer will support a different set of output formats, if the specified output format is not supported a 400 BadRequest response will be returned.
```

### Examples

The following examples assume a valid token with the `search` permission and a modern version of curl.  Any HTTP compatible tool or programming language will work too.

#### Pure Headers
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -H "query: tag=gravwell limit 10" \
   -H "duration: 1h" \
   -H "format: text" \
   http://10.0.0.1/api/search/direct
```

#### Pure Query Parameters
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   'http://10.0.0.1/api/search/direct?query=tag%3Dgravwell%20limit%2010&format=text&duration=1h'
```

#### Pure JSON Object
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -d '{"SearchString":"tag=gravwell limit 10","SearchStart":"2022-03-01T12:00:00Z","SearchEnd":"2022-03-01T13:00:00Z","Format":"text"}' \
   http://10.0.0.1/api/search/direct
```

#### Mixed Mode
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -d '{"SearchString":"tag=gravwell limit 10"}'
   -H 'duration: 1h' \
   http://10.0.0.1/api/search/direct?format=text
```

#### Downloading a CSV
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -H "query: tag=gravwell syslog Appname Hostname | stats count by Appname Hostname | table Appname Hostname count" \
   -H "duration: 1h" \
   -H "format: csv" \
   http://10.0.0.1/api/search/direct
```

#### Retrieving a PCAP
```
curl -X POST \
   -H "Gravwell-Token: aFOa_YbO7Pe0MAqK08PSD-oTrEZxopc5JBf0hu0W5_Vo-FxWsjHp" \
   -H "query: tag=packets packet tcp.Port==80 ipv4.IP==192.168.1.1 | pcap" \
   -H "duration: 1h" \
   -H "format: pcap" \
   --output /tmp/port80.pcap \
   http://10.0.0.1/api/search/direct
```
