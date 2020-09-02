# API

This section documents the web API used between the GUI and the "frontend" webserver.

The bulk of the API is RESTful. The exception to this rule is the searching API which uses websockets due to the nature of data exchange and transfer involved in launching and observing data from a search.

## Basic APIs

* [Login](login.md)
* [User Preferences](userprefs.md)
* [User account controls](account.md)
* [User group controls](groups.md)
* [Notifications](notifications.md)
* [Search Controls](searchctrl.md)
* [Downloading Search Results](download.md)
* [Search History](searchhistory.md)
* [Logging](loglevel.md)
* [Ingesting Entries](ingest.md)
* [Miscellaneous APIs](misc.md)
* [System Management](management.md)

## Objects within Gravwell

There are a variety of "things" which users may create and modify. Their APIs are listed in this section.

* [Auto-extractors](extractors.md)
* [Dashboards](dashboards.md)
* [Kits](kits.md)
* [Macros](macros.md)
* [Playbooks](playbooks.md)
* [Resources](resources.md)
* [Scheduled Searches](scheduledsearches.md)
* [Search Library](searchlibrary.md)
* [Templates](templates.md)
* [Pivots (actionables)](pivots.md)
* [User Files](userfiles.md)

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
