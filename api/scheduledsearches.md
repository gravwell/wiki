# Scheduled Searches API

This API allows for the creation and management of scheduled searches. Searches are referred to by a randomly-generated ID.

## Scheduled search structure

A scheduled search contains the following fields of interest:

* ID: the ID of the scheduled search
* GUID: a unique ID for this particular search. If left blank at creation, a random GUID will be assigned (this should be the standard use case)
* Owner: the uid of the search's owner
* Groups: a list of groups which are allowed to see the results of this search
* Name: the name of this scheduled search
* Description: a textual description of the scheduled search
* Schedule: a cron-compatible string specifying when to run
* Permissions: a 64-bit integer used to store permission bits
* LastRun: the time at which this scheduled search last ran
* LastRunDuration: how long the last run took
* LastSearchIDs: an array of strings containing the search IDs of the most recently performed searches from this scheduled search
* LastError: any error resulting from the last run of this search

If the search is a 'standard' scheduled search, it will also set these fields:

* SearchString: the Gravwell query to execute
* Duration: a value in seconds specifying how far back to run the search. This must be a negative value.
* SearchSinceLastRun: a boolean. If set, the Duration field will be ignored and the search will instead run from the LastRun time to the present.

If the search is on the other hand a script, it will set the following field:

* Script: a string containing an anko script

## User commands

The API commands in this section can be executed by any user.

### Listing scheduled searches

To get a list of all scheduled searches visible to the user (either owned by the user or marked accessible to one of the user's groups), perform a GET on `/api/scheduledsearches`. The result will look like this:

```
[{"ID":1824856041,"GUID":"126108d3-0159-4b1f-8f9e-26a93bb84433","Groups":null,"Name":"count","Description":"count all entries","Owner":1,"Schedule":"* * * * *","Permissions":0,"Updated":"2019-03-11T15:50:01.0327611-06:00","Synced":false,"SearchString":"tag=* count","Duration":0,"SearchSinceLastRun":true,"Script":"","PersistentMaps":{},"LastRun":"2019-03-11T15:50:00.011889037-06:00","LastRunDuration":1017076531,"LastSearchIDs":["892579845"],"LastError":""}]
```

This example shows a single scheduled search named "count", owned by UID 1 (admin). It runs every minute and executes the search `tag=* count` over the last hour hours.

### Creating a scheduled search

To create a new scheduled search, perform a POST request on `/api/scheduledsearches` with a JSON structure containing information about the scheduled search. To create a standard search, be sure to populate the SearchString and Duration fields, as in this example which runs a search over the last 24 hours every day at 8 a.m.:

```
{
	"Name": "myscheduledsearch",
	"Description": "a scheduled search",
	"Groups": [2],
	"Schedule": "0 8 * * *",
	"SearchString": "tag=default grep foo",
	"Duration": -86400,
	"SearchSinceLastRun": false
}
```

Alternately, if the SearchSinceLastRun field is set to true, the search agent will ignore the Duration (except for the first run of this new search) and instead perform the search over the time of the last run to the present time.

To create a scheduled search using a script, populate the "Script" field instead of the "SearchString" and "Duration" fields. If both are populated, the script will take precedence.

The server will respond with the ID of the new scheduled search.

### Fetching a specific scheduled search

Information about a single scheduled search may be accessed with a GET on `/api/scheduledsearches/{id}`. For example, given a scheduled search ID of 1353491046, we would query `/api/scheduledsearches/1353491046` and receive the following:

```
{"ID":1353491046,"GUID":"cdf011ae-7e60-46ec-827e-9d9fcb0ae66d","Groups":[2],"Name":"myscheduledsearch","Description":"a scheduled search","Owner":1,"Schedule":"0 8 * * *","Permissions":0,"SearchString":"tag=default grep foo","Duration":-86400,"Script":"","LastSearchIDs":null}
```

A scheduled search can also be fetched by GUID. Note that this requires more work for the webserver and should only be used when necessary. To fetch the scheduled search shown above, do a GET on `/api/scheduledsearches/cdf011ae-7e60-46ec-827e-9d9fcb0ae66d`.

### Updating an existing search

To modify a scheduled search, do an HTTP PUT to `/api/scheduledsearches/{id}` containing an updated structure with the desired changes. Take care to push the unchanged fields too, or they will be overwritten with empty values.

The following fields can be updated:

* Name
* Description
* Schedule
* SearchString
* Duration
* SearchSinceLastRun
* Script
* Groups

A script scheduled search can be changed to a standard scheduled search by pushing a SearchString and a Duration with the Script field empty. Likewise, a standard scheduled search can be converted to a script scheduled search by pushing a Script field and setting the SearchString empty.

### Clearing a scheduled search error

The LastError field in the scheduled search structure will be set if an error is encountered and will not be cleared by subsequent successful executions. It can be cleared manually by a DELETE on `/api/scheduledsearches/{id}/error`

### Deleting a scheduled search

An existing scheduled search can be removed by performing a DELETE on `/api/scheduledsearches/{id}`.

## Admin commands

The following commands are only available to admin users.

### Listing all searches

Admin users may occasionally need to view all scheduled searches on the system. An administrator user may obtain a global listing of all scheduled searches in the system with a GET request on `/api/scheduledsearches?admin=true`.

Because scheduled search IDs are unique across the system, the administrator may then modify/delete/retrieve any search without the need to specify `?admin=true`, although adding the parameter unecessarily will not cause an error.

### Fetching a specific user's searches

Performing a GET on `/api/scheduledsearches/user/{uid}`, where `uid` is a numeric userid, will fetch an array of all searches belonging to that user.

### Deleting all of a specific user's searches

Performing a DELETE on `/api/scheduledsearches/user/{uid}` will delete all scheduled searches belonging to the specified user.
