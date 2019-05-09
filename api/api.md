# API

This section documents the web API used between the GUI and the "frontend" webserver.

The bulk of the API is RESTful. The exception to this rule is the searching API which uses websockets due to the nature of data exchange and transfer involved in launching and observing data from a search.

## Primary APIs

* [Login](login.md)
* [User Preferences](userprefs.md)
* [Account controls](account.md)
* [Dashboards](dashboards.md)
* [Notifications](notifications.md)
* [Search Controls](searchctrl.md)
* [Downloading Search Results](download.md)
* [Search History](searchhistory.md)
* [Logging](loglevel.md)
* [Resources](resources.md)
* [Scheduled Searches](scheduledsearches.md)
* [Ingesting Entries](ingest.md)
* [Macros](macros.md)
* [Auto-extractors](extractors.md)
* [Miscellaneous APIs](misc.md)
* [Templates and Pivots](templates.md)
* [User Files](userfiles.md)
* [System Management](management.md)

## Searching and Search Stats

[Search Websocket](websocket-search.md)

[Reattaching to Searches](websocket-search-attach.md)

[Interacting with Renderers](websocket-render.md)

## System Stats

The system stats also use a websocket for communication. This contains all information necessary for monitoring general cluster health.

[System Stats Websocket](websocket-stats.md)

Some other stats may be accessed via REST calls.

[REST Stats API](stats-json.md)

## Test API

The System contains a test API located at _/api/test_ which can be used to test if the webserver is alive and functioning.  The test API is entirely unauthenticated and always responds with a StatusOK 200 and an empty body.
