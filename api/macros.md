# Macros API

The web API provides methods for accessing and creating search macros, which are mappings of a short string to a longer string which is expanded during the parse phase of a search.

## The SearchMacro structure

The web server returns macros in a JSON struct which is also used to update fields. Note that when sending a struct to update a macro, not all fields are updated (it is not possible to change the LastUpdated field manually, for instance). The fields are largely self-explanatory but are explained here for precision:

* ID: a unique integer representing this macro
* UID: the macro owner's integer UID
* GIDs: a list of integer group IDs which are allowed to access the macro
* Name: the short name of the macro as typed in a search query
* Expansion: the string to which the macro expands
* LastUpdated: the time at which this macro was most recently modified
* Synced: (internal use only)

## Listing macros

To get a list of all macros belonging to the current user, do a GET on `/api/macros`. The result will look like this:

```
[{"ID":1,"UID":1,"GIDs":null,"Name":"FOO","Expansion":"grep foo","LastUpdated":"2018-10-31T20:56:24.629561628Z","Synced":true}]
```

In this example, the user with UID 1 has one macro named "FOO" which expands to the string "grep foo". The macro ID is 1, which can be used in other APIs.

### Getting macros by user ID

Admin users can retrieve a list of a specific user's macros by performing a GET on `/api/users/{uid}/macros`, replacing `{uid}` with the desired user ID. Non-admin users can use this API to retrieve their own macros.

### Getting macros by group ID

Admin users or members of a group can retrieve a list of macros to which a specified group has access by performing a GET on `/api/groups/{gid}/macros`.

### Getting all macros

Admin users can retrieve a list of all macros on the system by performing a GET on `/api/macros/all`.

## Retrieving a specific macro

The structure for a specific macro may be retrieved by doing a GET on `/api/macros/{id}`, replacing `{id}` with the macro ID. For instance, a GET on `/api/macros/1` returns the following:

```
{"ID":1,"UID":1,"GIDs":null,"Name":"FOO","Expansion":"grep foo","LastUpdated":"2018-10-31T20:56:24.629561628Z","Synced":true}
```

## Creating a new macro

New macros may be created via POST to `/api/macros`. The body of the request should contain the Name and Expansion fields and (optionally) the GIDs field, as shown below:

```
{"GIDs": [1, 2], "Name": "TEST", "Expansion": "grep test | count"}
```

A successful creation will return the ID of the new macro; in this example the new ID was 2, so a GET on `/api/macros/2` yields the full body of the new macro:

```
{"ID":2,"UID":1,"GIDs":null,"Name":"TEST","Expansion":"grep test | count","LastUpdated":"2018-10-31T21:05:34.231426316Z","Synced":true}
```

## Updating a macro

Updating a macro may be done by PUT to `/api/macros/{id}` using the same body format as a creation; note that for safety, it is best to populate the GIDs, Name, and Expansion fields every time even if no changes are made.

A successful update returns HTTP 200 and the updated macro in the body.

## Deleting a macro

Macro deletion is performed via a DELETE on `/api/macros/{id}`. This returns HTTP 200 upon success.