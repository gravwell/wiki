# Scheduled Searches API

This API allows for the creation and management of scheduled searches. Searches are referred to by a randomly-generated ID.

## Scheduled search structure

A scheduled search contains the following fields:

* ID: the ID of the scheduled search
* Owner: the uid of the search's owner
* Groups: a list of groups which are allowed to see the results of this search
* Name: the name of this scheduled search
* Description: a textual description of the scheduled search
* Schedule: a cron-compatible string specifying when to run
* Permissions: a 64-bit integer used to store permission bits
* LastSearchIDs: an array of strings containing the search IDs of the most recently performed searches from this scheduled search

If the search is a 'standard' scheduled search, it will also set these fields:

* SearchString: the Gravwell query to execute
* Duration: a value in nanoseconds specifying how far back to run the search. This must be a negative value.

If the search is on the other hand a script, it will set the following field:

* Script: a string containing an anko script

## User commands

The API commands in this section can be executed by any user.

### Listing scheduled searches

To get a list of all scheduled searches visible to the user (either owned by the user or marked accessible to one of the user's groups), perform a GET on `/api/scheduledsearches`. The result will look like this:

```
[{"ID":2007987335,"Groups":null,"Name":"foo","Description":"test search","Owner":1,"Schedule":"* * * * *","Permissions":0,"SearchString":"tag=gravwell","Duration":-86400000000000,"Script":"","LastSearchIDs":null}]
```

This example shows a single scheduled search named "foo", owned by UID 1 (admin). It runs every minute and executes the search `tag=gravwell` over the last 24 hours.

### Creating a scheduled search

To create a new scheduled search, perform a POST request on `/api/scheduledsearches` with a JSON structure containing information about the scheduled search. To create a standard search, be sure to populate the SearchString and Duration fields, as in this example which runs a search over the last 24 hours every day at 8 a.m.:

```
{
	"Name": "myscheduledsearch",
	"Description": "a scheduled search",
	"Groups": [2],
	"Schedule": "0 8 * * *",
	"SearchString": "tag=default grep foo",
	"Duration": -86400000000000
}
```

To create a scheduled search using a script, populate the "Script" field instead of the "SearchString" and "Duration" fields. If both are populated, the script will take precedence.

The server will respond with the ID of the new scheduled search.

### Fetching a specific scheduled search

Information about a single scheduled search may be accessed with a GET on `/api/scheduledsearches/{id}`. For example, given a scheduled search ID of 1353491046, we would query `/api/scheduledsearches/1353491046` and receive the following:

```
{"ID":1353491046,"Groups":[2],"Name":"myscheduledsearch","Description":"a scheduled search","Owner":1,"Schedule":"0 8 * * *","Permissions":0,"SearchString":"tag=default grep foo","Duration":-86400000000000,"Script":"","LastSearchIDs":null}
```

### Updating an existing search

To modify a scheduled search, do an HTTP PUT to `/api/scheduledsearches/{id}` containing an updated structure with the desired changes. It is only necessary to push those fields which are to be changed, thus to change the description one would push:

```
{
	"Description": "my new description"
}
```

The following fields can be updated:

* Name
* Description
* Schedule
* SearchString
* Duration
* Script

A script scheduled search can be changed to a standard scheduled search by pushing a SearchString and a Duration. Likewise a standard scheduled search can be converted to a script scheduled search by pushing a Script field.

### Deleting a scheduled search

An existing scheduled search can be removed by performing a DELETE on `/api/scheduledsearches/{id}`.

## Admin commands

The following commands are only available to admin users.

### Listing all searches

A GET on `/api/scheduledsearches/all` will return an array containing all scheduled searches on the system.

### Fetching a specific user's searches

Performing a GET on `/api/scheduledsearches/user/{uid}`, where `uid` is a numeric userid, will fetch an array of all searches belonging to that user.

### Deleting all of a specific user's searches

Performing a DELETE on `/api/scheduledsearches/user/{uid}` will delete all scheduled searches belonging to the specified user.
