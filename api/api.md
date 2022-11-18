# API

This section documents the web API used between the GUI and the "frontend" webserver.

The bulk of the API is RESTful. The exception to this rule is the searching API which uses websockets due to the nature of data exchange and transfer involved in launching and observing data from a search.

## APIs & Objects

```{toctree}
---
maxdepth: 1
caption: Basic APIs
---
Login <login>
User Preferences <userprefs>
User account controls <account>
User group controls <groups>
Notifications <notifications>
Search Controls <searchctrl>
Downloading Search Results <download>
Search History <searchhistory>
Logging <loglevel>
Ingesting Entries <ingest>
Miscellaneous APIs <misc>
System Management <management>
```

```{toctree}
---
maxdepth: 1
caption: API Objects
---
Auto-extractors <extractors>
Dashboards <dashboards>
Kits <kits>
Macros <macros>
Playbooks <playbooks>
Resources <resources>
Scheduled Searches <scheduledsearches>
Search Library <searchlibrary>
Templates <templates>
Actionables <actionables>
User Files <userfiles>
```

```{toctree}
---
maxdepth: 1
caption: Search and Stats
---
Gravwell Direct Query API <../search/directquery/directquery>
Search Websocket <websocket-search>
Reattaching to Searches <websocket-search-attach>
Interacting with Renderers <websocket-render>
```
## System Stats

The system stats also use a websocket for communication. This contains all information necessary for monitoring general cluster health.

```{toctree}
---
maxdepth: 1
caption: System Stats (WebSocket)
---
System Stats Websocket <websocket-stats>
```

Some other stats may be accessed via REST calls.

```{toctree}
---
maxdepth: 1
caption: System Stats (REST)
---
REST Stats API <stats-json>
```

## Test API

The System contains a test API located at _/api/test_ which can be used to test if the webserver is alive and functioning.  The test API is entirely unauthenticated and always responds with a StatusOK 200 and an empty body.

## Tokens
```{toctree}
---
maxdepth: 1
caption: API Tokens System
---
API Tokens System <../tokens/tokens>
```
