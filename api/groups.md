# Group APIs

This page describes the APIs for interacting with user groups.

Requests sent to `/api/groups` and other URLs under the `/api/groups/` path operate on user accounts. The webserver will send StatusOK (200) on a good request, while a 400-500 status will be sent on error (depending on the error).

Note: Adding and removing users from groups is done through the [user account APIs](account.md), not the group APIs. These APIs are purely for creating, deleting, and describing groups.

## Data types

The APIs on this page primarily deal with group details structures. The group details structure is as follows:

```
{
  GID: int,		// Numeric group ID
  Name: string, // Group name
  Desc: string, // Group description
  Synced: bool	// Tracks if group is synchronized with the datastore (can usually be ignored)
}
```

## Admin only: Adding a group

To add a new group, send a POST request to `/api/groups`. The body of the request should contain a structure which defines the group's name and description, as below:

```
{
	"Name": "newgroup",
	"Desc": "This is the new group"
}
```

The server will attempt to parse the request and create the group; if successful, it will respond with a 200 status code and the GID of the new group in the body of the response.

## Admin only: Deleting a group

To delete a group, send a DELETE request to `/api/groups/{gid}`.

## Admin only: Updating group info

To modify a group's information, send a PUT request to `/api/groups/{gid}`. The body of the request should contain a group details structure, as below:

```
{
	"GID": 3,
	"Name": "newname",
	"Desc": "this is the new description",
}
```

If the Name or Desc fields are excluded, the current values will be kept.

## Listing all groups

To get a list of all groups on the system, send a GET request to `/api/groups`. The response will contain an array of group details structures:

```
[
    {
        "Desc": "bar",
        "GID": 1,
        "Name": "foo",
        "Synced": true
    },
    {
        "Desc": "",
        "GID": 7,
        "Name": "gravwell-users",
        "Synced": false
    },
    {
        "Desc": "",
        "GID": 8,
        "Name": "testgroup",
        "Synced": false
    }
]
```

## Getting info about a group

Administrators or members of a particular group may get information about a particular group. Issue a GET request to `/api/groups/{gid}`. The response will contain the group details:

```
{
    "Desc": "",
    "GID": 7,
    "Name": "gravwell-users",
    "Synced": false
}
```

## Listing users in a group

Administrators or members of a particular group may query the list of users in that group by issuing a GET request to `/api/groups/{gid}/members`. The response will be an array of user details structures:

```
[
    {
        "Admin": false,
        "Email": "joe@gravwell.io",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            },
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            },
            {
                "Desc": "",
                "GID": 8,
                "Name": "testgroup",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Joe User",
        "Synced": false,
        "TS": "2020-08-10T14:49:30.72782227-06:00",
        "UID": 16,
        "User": "joe"
    },
    {
        "Admin": false,
        "Email": "bkeaton@example.net",
        "Groups": [
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Buster Keaton",
        "Synced": false,
        "TS": "0001-01-01T00:00:00Z",
        "UID": 17,
        "User": "buster"
    }
]
```
