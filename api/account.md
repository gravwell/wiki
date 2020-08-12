# User Account APIs

This page describes the APIs for interacting with users.

Requests sent to `/api/users` and other URLs under the `/api/users/` path operate on user accounts. The webserver will send StatusOK (200) on a good request, while a 400-500 status will be sent on error (depending on the error).

The default "admin" account username cannot be changed or deleted, nor can the account be demoted out of admin status. Other user accounts may be also assigned administrator privileges.

## Data types

The APIs on this page primarily deal with user and group details structures. The user details structure is as follows:

```
{
  UID: int,			// Numeric user ID
  User: string,		// Username e.g. "jdoe"
  Name: string,		// User's real name e.g. "John Doe"
  Email: string,	// User's email address e.g. "john@example.org"
  Admin: bool,		// Set to true if the user has admin rights
  Locked: bool,		// Set to true if the user's account has been locked
  DefaultGID: int,	// The user's default search group
  TS: string,		// Contains a timestamp representing when the user was last active
  Synced: bool,		// Tracks if user is synchronized with the datastore (can usually be ignored)
  Groups: Array<GroupDetails>	// An array of groups to which the user belongs
}
```

The group details structure is as follows:

```
{
  GID: int,		// Numeric group ID
  Name: string, // Group name
  Desc: string, // Group description
  Synced: bool	// Tracks if group is synchronized with the datastore (can usually be ignored)
}
```

## Log in and log out

The API to authenticate with the webserver is [documented here](login.md).

## Fetch current user's info
GET on /api/info/whoami returns current account info

```
{
        "UID": 1,
        "User": "admin",
        "Name": "Sir changeme of change my password the third",
        "Email": "admin@admin.admin",
        "Admin": true,
        "Locked": false,
        "TS": "2016-12-05T17:05:33.121180268-07:00",
        "DefaultGID": 0,
        "Groups": [
                {
                        "GID": 1,
                        "Name": "foo",
                        "Desc": "This is the foo group"
                }
        ]
}

```

## List all accounts

Sending a GET request to `/api/users` will return a JSON packet containing information about every account.  The JSON returned is as follows:

```
[
    {
        "Admin": true,
        "Email": "admin@admin.admin",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Admin John",
        "Synced": true,
        "TS": "2020-07-30T08:51:35.205998608-06:00",
        "UID": 1,
        "User": "admin"
    },
    {
        "Admin": false,
        "Email": "john@example.net",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "John",
        "Synced": false,
        "TS": "2020-08-10T14:47:44.58356179-06:00",
        "UID": 14,
        "User": "john"
    },
    {
        "Admin": false,
        "Email": "joe@gravwell.io",
        "Groups": [
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Joe User",
        "Synced": false,
        "TS": "2020-08-10T14:49:30.72782227-06:00",
        "UID": 16,
        "User": "joe"
    }
]
```

## Get a single user's information
To get a single user's account information, send a GET to `/api/users/{id}/`.  Admins can retrieve any account, while non-admins can ONLY retrieve their own account information.  Response of 200 will contain valid JSON in the body, 400-500 means someone boned the request and the body will contain an error message.  The backend will respond with a JSON packet containing the user details as described at the top of this page; an example response follows:

```
{
    "Admin": true,
    "DefaultGID": 1,
    "Email": "admin@admin.admin",
    "Groups": [
        {
            "Desc": "bar",
            "GID": 1,
            "Name": "foo",
            "Synced": false
        }
    ],
    "Locked": false,
    "Name": "Admin",
    "Synced": false,
    "TS": "2020-07-30T08:51:35.205998608-06:00",
    "UID": 1,
    "User": "admin"
}
```

## Add a new user
To add a new user, an admin account can POST to `/api/users` with a request containing these fields:

```
{
     User: "buster",
     Pass: "gr4vwellRulez",
     Name: "Buster Keaton",
     Email: "bkeaton@example.net"
     Admin: false,
}
```

All fields must be populated.  The back end will respond with the new user's UID, e.g. `17`.

## Lock a user account
To lock an account, send an empty PUT to `/api/users/{id}/lock`, where {id} is the UID for the user that will be locked.  Locking a user account will immediately log out the user's active sessions and prevent any new logins.

## Unlock a user account
To unlock an account, send an empty DELETE to `/api/users/{id}/lock` where {id} is the UID for the user that will be unlocked.  The webserver will respond with success regardless of whether the account was actually unlocked if the action is allowed.  We do this because locking a locked account ends in the state that the account is locked.  So its all good.  Same with unlocking.

## Changing user info
To change user info the client should send a PUT `/api/users/{id}` with the desired fields to update:

```
{
     User: "chuck",
     Name: "Chuck Testa",
     Email: "chuck@testa.net"
}
```

Any field that is NOT populated will be ignored.  The backend will respond with the standard response JSON as mentioned above.  If the current user is NOT an admin and not changing their own account, the request will be rejected.  Admins can change information for any account.  The primary admin account (UID zero) cannot change its admin status.  The backend responds with a 200 on success and 400-500 on error depending on who caused the error and why.  Error messages will be returned in the body of the response and are human displayable.

## Changing a users password

To change a password, issue a PUT to the url `/api/users/{id}/pwd`.

```
{
     OrigPass: "my old password was bad",
     NewPass: "thisis mynewpassword",
}
```

Note: If the current user is an admin and changing the password for an account they DO NOT OWN, the OrigPass field is NOT required.

## Delete a user
To delete a user the client should send a DELETE to `/api/users/{id}/` with `{id}` set to the UID to delete.  Only admins can use this facility, and a user cannot delete their own account.  The primary admin (UID 1) cannot be deleted.

## Get/Set user admin status
To change the admin status of a user or query the admin status, send a request to `/api/users/{id}/admin` with an empty body. The method specifies the desired action.  The webserver will respond with 200 on success and 400-500 on error depending on the type of failure.

* GET - returns current admin status
* PUT - sets the user as an admin and returns the new status
* DELETE - removes admin status for the user

Example JSON on success:

```
{
     UID: 1,
     Admin: true
}
```

## Get user sessions

A user may get a list of his or her own sessions by issuing a GET request to `/api/users/{id}/sessions`. An admin may issue the request for any UID. The response will be an array of session objects, as shown below:

```
{
    "Sessions": [
        {
            "LastHit": "2020-08-11T14:41:23.829366-06:00",
            "Origin": "::1",
            "Synced": false,
            "TempSession": false
        }
    ],
    "UID": 1,
    "User": "admin"
}
```

## Set default search group

Each user can set a default search group. If set, all queries run by the user will be shared with that group by default. To set the default search group, issue a PUT request to `/api/users/{id}/searchgroup`, with the desired group specified in the request body as shown below:

```
{
	"GID": 3
}
```

## Get default search group

Each user can set a default search group. If set, all queries run by the user will be shared with that group by default. To fetch the user's search group, issue a GET request to `/api/users/{id}/searchgroup`. The server's response will contain an integer GID, e.g. `3`.

## User preferences

"User preferences" consist of an arbitrary JSON blob set by the frontend for the user. This allows the frontend to store the user's UI settings (color scheme, etc.) across multiple computers. User preferences can only be set by 1) the user in question, or 2) an admin user.

### Set user preferences

To set the user's preferences, issue a PUT request to `/api/users/{id}/preferences`. The request body should contain valid JSON; the actual contents of the JSON are ignored by the webserver.

### Get user preferences

To fetch the user's preferences, issue a GET request on `/api/users/{id}/preferences`. The body of the response will contain any previously-set preferences.

### Clear user preferences

To clear out anything in the user's preferences, issue a DELETE on `/api/users/{id}/preferences`.

### Admin only: Fetch all user preferences

The administrator can issue a GET request on `/api/users/preferences` to get a complete listing of all users' preferences. The response will include both user preferences (with Name == "prefs") and user *email* configurations (Name == "emailSettings"). The interface to set and modify email settings is described elsewhere:

```
[
    {
        "Data": "ImJhciBiYXogcXV1eCBhc2Rmc2FkZmRzZiI=",
        "Name": "prefs",
        "Synced": false,
        "UID": 1,
        "Updated": "2020-07-13T13:50:08.286015047-06:00"
    },
    {
        "Data": "ImZvbyI=",
        "Name": "emailSettings",
        "Synced": false,
        "UID": 1,
        "Updated": "2019-12-18T09:45:38.683780752-07:00"
    }
]
```

## Add user to groups

Administrators can add users to groups by sending a POST request to `/api/users/{id}/group`. The body of the request should contain a list of GIDs which should be *added* to the user's group memberships. Thus, to add a user to a group with GID 8, send the following, regardless of any other group memberships:

```
{
	"GIDs": [8]
}
```

## Remove user from group

To remove a user from a specific group, an admin can send a DELETE request to `/api/users/{id}/group/{gid}`, where {id} is the user's UID and {gid} is the group's GID.

## Get user groups

To get a list of groups to which a user belongs, the user (or an admin) should issue a GET request to `/api/users/{id}/group`. The response will be an array of group details structures, as below:

```
[
    {
        "Desc": "foo group",
        "GID": 1,
        "Name": "foo",
        "Synced": false
    }
]
```
