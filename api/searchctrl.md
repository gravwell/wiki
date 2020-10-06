# Search Control

REST API located at /api/searchctrl

The searchctrl group is used to query information about persistent searches and optionally invoke some action.  We currently provide the ability to query all active searches, get information about a specific search, background a an active search, archive an active search, and delete/terminate a search.

## Basic API overview

The basic action here is to perform a GET, DELETE, or PATCH on a REST url.

## Search States

Searches can be in any of the following states, and can be in more than one state at a time (such as ACTIVE/BACKGROUNDED, SAVED/ATTACHED, or DORMANT/SAVED). 

- Active: The search is running and/or finished and there is a session attached to it.
- Backgrounded: The search is actively running but marked in a way that it will persist without an attached session.
- Saving: The search has been marked as saved, but is still waiting for completion to move its contents to a persistent location.
- Saved: The search is marked as saved and moved to the appropriate persistent location.
- Dormant: The search is being kept (background or saved) and no sessions are attached.
- Attached: The search is saved and a session has re-attached to it.

## Getting a list of searches
When requesting the list of searches the web server will return all searches the current user is authorized to view.  No json package is posted, but rather an empty GET is performed against '/api/searchctrl'.  

```
[
        {
                "ID": "181382061",
                "UID": 1,
                "GID": 0,
                "State": "ACTIVE",
                "AttachedClients": 1
        },
        {
                "ID": "010985768",
                "UID": 1,
                "GID": 0,
                "State": "BACKGROUNDED",
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
[
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
]
```

## Backgrounding a search

Backgrounding a search is used to inform the web server that if the last client lets go, its ok to continue the search.  To background a search perform a PATCH on the url /api/searchctrl/:ID/background correct ID.   The Search MUST be active or already backgrounded for the command to succeed and the user must be an admin or have access to the search.

```
WEB PATCH /api/searchctrl/795927171/background:
null
```

## Saving a search

Saving a search is used to inform the webserver that we wish to keep the results of this search.  A backgrounded search will stay resident (even if no one is connected to it) as long as the webserver doesn't need the disk space (or it isn't explicitly deleted).  Saving moves the results to the saved location, and the results will not be deleted unless someone (with the proper authority) explicitly requests it.  To Save a search perform a PATCH on the url /api/searchctrl/:ID/save correct ID.   The Search can be in in any state, but will only begin transferring to the persistent storage once it hits the dormant state.  The transfer to persistent storage is either instantaneous (if the persistent storage is on the same drive) or requires a full copy.  This is done in the background in its own goroutine, so nothing is blocked while it happens.

An optional set of metadata can be attached to a saved search, the metadata must be a valid JSON structure.  This metadata can be used to attach notes, info, or pretty much anything that can be encoded as a JSON blob.  Metadata should be encoded into the body of the `PATCH` request to save the search.

```
WEB PATCH /api/searchctrl/10985768/save:
null
```

## Deleting/terminating a search

Deleting a search terminates the search (and kicks off any active users) and immediately removes any storage associated with the search results.  A search may be deleted while in any state.  To delete a search peroform a DELETE request to /api/searchctrl/:ID with the correct ID.  The server will return 200 on success, 5XX on error, and 403 if the user is not authorized to modify the search.

```
WEB DELETE /api/searchctrl/010985768:
null
```

## Importing a saved search archive

An optional download format for a search is an `archive`.  An archive represents a fully self-contained search that can be imported into another Gravwell instance.  The import API accepts the saved search archives as an upload and unpacks the search into the saved search system.  Users can then attach to the search as if it were saved on the local system.

When a search archive is reimported, the imported search is owned by the user that *imported* it, regardless of which user *downloaded* it.  An optional `GID` form field may be supplied in the import request to share the imported search with a group.

Searches are imported by performing a multipart form `POST` to the `/api/searchctrl/import` URL.

The API expects that the file upload be in the form field called `file`.

The import API can be authenticated using either the JWT authorization token or a cookie.

NOTE: An admin can specify groups it is not a member of in the `GID` form field, but non-admin users must be a member of the group specified in `GID`.

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
