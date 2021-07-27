# Search websocket

Websocket URL: /api/ws/search

This page documents the websocket protocol for searches. A full example of JSON transferred between client and server when initiating a "grep foo" search, complete with entry data retrieval, can be found at the [Websocket Search Example](websocket-search-example.md) page.

## Ping/Pong keepalive

The search websocket is used for checking queries, sending searches, and receiving search results and search stats.  The search websocket expects to use the RoutingWebsocket system to handle message "types".  `/api/ws/search` expects the following message "subtypes" or "types" to be registered at startup: PONG, parse, search, attach.

Note: Message "types" are sometimes referred to as "SubProto" due to legacy naming. This is likely to change in the future but if developing against this API, be aware that "SubProto" refers to the "type" value sent with a message and not the RFC Websocket subprotocol spec.

The PONG type is a keepalive system, and the client should periodically send PING/PONG requests.

This could be used so that if a user is sitting at a search prompt or whatever that you can tell the user if the conn to the back is healthy. This may not be necessary at all as we can just probe to see if the websocket itself is alive.

## Parsing searches

The "parse" websocket type is used for rapidly testing the validity of queries without invoking any of the search backend.

An example request with a valid query and response would result in the following JSON:

request from frontend:
```json
{
        SearchString: "tag=apache grep firefox | regex "Firefox(<version>[0-9]+) .+" | count by version""
}
```

response from backend:
```
{
        GoodQuery: true,
        ParseQuery: "tag=apache grep firefox | regex "Firefox(<version>[0-9]+) .+" | count by version"",
        ModuleIndex: 0,
}
```

An example request with an invalid query and response would result in the following JSON:

request from frontend:
```
{
        SearchString: "tag=apache grep firefox | MakeRainbows",
}
```

response from backend:
```
{
        GoodQuery: false,
        ParseError: "ModuleError: MakeRainbows is not a valid module",
        ModuleIndex: 1,
}
```

## Initiating searches
All searches are initiated through websockets and require that the "parse", "PONG", "search", and "attach" subtypes are requested at start.  

This is done by sending the following JSON upon websocket establishment:
```
{"Subs":["PONG","parse","search","attach"]}
```


The SearchString member should contain the actual query which will invoke the search.

SearchStart and SearchEnd should be the time ranges that the query will operate over.  The time ranges should be formatted in the RFC3339Nano format which looks like "2006-01-02T15:04:05.999999999Z07:00"

An example search request with a good query would contain the following JSON:
```
{
       SearchString: "tag=apache grep firefox | nosort",
       SearchStart:  "2015-01-01T12:01:00.0Z07:00",
       SearchEnd:    "2015-01-01T12:01:30.0Z07:00",
       Background:   false,
}
```

//server responds yay/nay plus new subtypes if the search is cool
//searchStart and searchEnd should be strings in RFC3339Nano format

The response to a good query would contain the following JSON:
```
{
        SearchString: "tag=apache grep firefox | nosort",
        RenderModule: "text",
        RenderCmd:    "text",
        OutputSearchSubproto:  "searchSDF8973",
        OutputStatsSubproto:   "statsSDF8973",
        SearchID:              "skdlfjs9098",
		SearchStartRange:      "2015-01-01T12:01:00.0Z07:00",
        SearchEndRange:        "2015-01-01T12:01:30.0Z07:00",
        Background:            false,
}
```

On error the JSON response would be:
```
{
        Error: "Search error: The parameter "ChuckTesta" is invalid",
}
```

On a good search request response the client must response with a search ACK. The Ack must respond with the either a true or false.  A false response may be used when the backend requests a render module that the front end doesn't understand, which may happen when there is a version mismatch between the frontend and backend.

The following JSON would represent an affirmative ACK to the previous response example:
```
{
       Ok: True,
       OutputSearchSubproto: "searchSDF8973"
}
```

After the ACK is sent the backend will fire up the search and begin providing search results on the new subtypes.  The original search, parse, and PONG subtypes stay active and can be used by the frontend to check new queries, or kick off additional searches.  All interaction with active queries needs to occur via the newly negotiated search specific subtypes though.

## Notes
All searches are fully asynchronous, however, if a client disconnects or the connection crashes without requesting that a search be placed in a background state, the active search will terminate and the data will be garbage collected.  This is to prevent resource exhaustion.  A user must EXPLICITLY request a background search.

Searches can have multiple consumers.  For example Bob can kick off a search and Janet may attach to it and see the results.  A non-backgrounded search will only terminate and cleanup if ALL consumers disconnect.  So if Bob kicks off a search and Janet attaches, but Bob then navigates away or closes his browser the search will not terminate.  Janet can continue to interact with it.  However, if Janet also navigates away or closes her browser the search will then terminate and garbage collect.

## Stats output during an active search

Stats are requested via the stats IDs

## Request/Response ID reference

The list of request and response ID codes is:
```
{
    req: {
        REQ_CLOSE: 0x1,
        REQ_ENTRY_COUNT: 0x3,
        REQ_DETAILS: 0x4,
        REQ_TAGS: 0x5,
        REQ_STATS_SIZE: 0x7F000001, //gets backend "size" value of stats chunks. never used
        REQ_STATS_RANGE: 0x7F000002, //gets current time range covered by stats. rarely used
        REQ_STATS_GET: 0x7F000003, //gets stats sets over all time. may be used initially
        REQ_STATS_GET_RANGE: 0x7F000004, //gets stats in a specific range
        REQ_STATS_GET_SUMMARY: 0x7F000005, //gets stats summary for entire results
        REQ_STATS_GET_LOCATION: 0x7F000006, //get current timestamp for search progress
        REQ_GET_ENTRIES: 0x10, //1048578
        REQ_STREAMING: 0x11,
        REQ_TS_RANGE: 0x12,
		REQ_GET_EXPLORE_ENTRIES: 0xf010,
		REQ_EXPLORE_TS_RANGE: 0xf012,
        SEARCH_CTRL_CMD_DELETE: 'delete',
        SEARCH_CTRL_CMD_ARCHIVE: 'archive',
        SEARCH_CTRL_CMD_BACKGROUND: 'background',
        SEARCH_CTRL_CMD_STATUS: 'status'
    },
    rep: {
        RESP_CLOSE: 0x1,
        RESP_ENTRY_COUNT: 0x3,
        RESP_DETAILS: 0x4,
        RESP_TAGS: 0x5,
        RESP_STATS_SIZE: 0x7F000001, //2130706433
        RESP_STATS_RANGE: 0x7F000002, //2130706434
        RESP_STATS_GET: 0x7F000003, //2130706435
        RESP_STATS_GET_RANGE: 0x7F000004, //2130706436
        RESP_STATS_GET_SUMMARY: 0x7F000005,
        RESP_STATS_GET_LOCATION: 0x7F000006, //2130706438
        RESP_GET_ENTRIES: 0x10,
        RESP_STREAMING: 0x11,
        RESP_TS_RANGE: 0x12,
		RESP_GET_EXPLORE_ENTRIES: 0xf010,
		RESP_EXPLORE_TS_RANGE: 0xf012,
        RESP_ERROR: 0xFFFFFFFF,
        SEARCH_CTRL_CMD_DELETE: 'delete',
        SEARCH_CTRL_CMD_ARCHIVE: 'archive',
        SEARCH_CTRL_CMD_BACKGROUND: 'background',
        SEARCH_CTRL_CMD_STATUS: 'status'
    }
}
```
