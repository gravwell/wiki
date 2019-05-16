# Search Control

REST API located at /api/searchctrl

The searchctrl group is used to query information about active searches and optionally invoke some action.  We currently provide the ability to query all active searches, get information about a specific search, background a an active search, archive an active search, and delete/terminate a search.

## Basic API overview

The basic action here is to perform a GET, DELETE, or PATCH on a REST url

## Getting a list of searches
When requesting the list of searches the web server will return all searches the current user is authorized to view.  No json package is posted, but rather an empty GET is performed against '/api/searchctrl/'.  

```
[
        {
                "ID": "181382061",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "010985768",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "795927171",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        }
]
```

## Getting the info of a specific search
Getting the status of a specific search is performed by performing a GET the REST url /api/searchctrl/:ID

```
WEB GET /api/searchctrl/795927171:
{
        "ID": "795927171",
        "UID": 1,
        "UserQuery": "grep paco | grep chico",
        "EffectiveQuery": "grep paco | grep chico | text",
        "StartRange": "2016-12-22T12:41:27.011080417-07:00",
        "EndRange": "2016-12-22T13:01:27.011080417-07:00",
        "Started": "2016-12-22T13:01:27.01227455-07:00",
        "Finished": "0001-01-01T00:00:00Z",
        "StoreSize": 0,
        "IndexSize": 0
}
```

## Backgrounding a search

Backgrounding a search is used to inform the web server that if the last client lets go, its ok to continue the search.  To background a search perform a PATCH on the url /api/searchctrl/:ID/background correct ID.   The Search MUST be active or already backgrounded for the command to succeed and the user must be an admin or have access to the search.

```
WEB PATCH /api/searchctrl/795927171/background:
null
```

## Saving a search

Saving a search is used to inform the webserver that we wish to keep the results of this search.  A backgrounded search will stay resident (even if no one is connected to it) as long as the webserver doesn't need the disk space (or it isn't explicitly deleted).  Saving moves the results to the saved location, and the results will not be deleted unless someone (with the proper authority) explicitly requests it.  To Save a search perform a PATCH on the url /api/searchctrl/:ID/save correct ID.   The Search can be in in any state, but will only begin transferring to the persistent storage once it hits the dormant state.  The transfer to persistent storage is either instantaneous (if the persistent storage is on the same drive) or requires a full copy.  This is done in the background in its own goroutine, so nothing is blocked while it happens.

```
WEB PATCH /api/searchctrl/010985768/save:
null
```

## Deleting/terminating a search

Deleting a search terminates the search (and kicks off any active users) and immediately removes any storage associated with the search results.  A search may be deleted while in any state.  To delete a search peroform a DELETE request to /api/searchctrl/:ID with the correct ID.  The server will return 200 on success, 5XX on error, and 403 if the user is not authorized to modify the search.

```
WEB DELETE /api/searchctrl/010985768:
null
```

## Admin APIs

Admin users can get information about any search, delete any search, load any search, send any search to the background, etc. using the API endpoints documented above.

### List all searches

In order to get a list of all searches that exist on the system, an admin user may do a GET on `/api/searchctrl/all`. The format is identical to that returned from `/api/searchctrl`, but includes all searches on the system.

```
[
    {
        "AttachedClients": 0,
        "GID": 0,
        "ID": "486574780",
        "State": "DORMANT",
        "StoredData": 9355,
        "UID": 1
    },
    {
        "AttachedClients": 0,
        "GID": 0,
        "ID": "815623546",
        "State": "DORMANT",
        "StoredData": 3536,
        "UID": 4
    },
    {
        "AttachedClients": 0,
        "GID": 0,
        "ID": "525125903",
        "State": "DORMANT",
        "StoredData": 0,
        "UID": 7
    },
    {
        "AttachedClients": 0,
        "GID": 0,
        "ID": "274379984",
        "State": "DORMANT",
        "StoredData": 319,
        "UID": 4
    }
]

```