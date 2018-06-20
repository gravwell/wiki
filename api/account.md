# Account control and info page for users and groups
This page describes the API used for interacting with users and groups.

## Connectivity test
#### allows for checking connectivity with the backend
Basically a GET on /api/test returns a 200 with no content.  Pretty simple.

## Version API
#### Path is /api/version/
Perform a GET and you get version info.  Any, no authentication required.
```
{"Version":0.1,"Date":"0001-01-01T00:00:00Z"}
```
## Logging in and logging out

[documented here](login.md)

# List accounts and activity
The /api/user/ subdirectory is for admins to be able to get account info and also see account activity

## List all accounts

An admin can perform a GET request against /api/user and the backend will return a JSON packet containing a summery for every account.  The JSON returned is as follows:

```
{[
    {
        UID: 1
        User: "Administrator"
        Name: "I am the law"
        Email: "admin@example.com"
        Admin: true
        Locked: false
        TS: "2009-11-01 17:45:02.1000 +0000 UTC"
    },
    {
        UID: 2
        User: "paco"
        Name: "Paco Chingato"
        Email: "paco@gmail.com"
        Admin: false
        Locked: true
        TS: "2014-05-01 01:31:33.1234 +0000 UTC"
    }
]}
```

# account control
An admin can operate on /api/user/ in order to modify accounts.  The backend will always respond with a JSON packet which indicates whether the command succeeded or not and a possible informational error message.  On failure the frontend should display the error message to the user (it is formatted for a human).  The appropriate error message will be dumped in the body of the response.  StatusOK (200) will be sent on a good request and a 400-500 status will be sent on error (depending on what error and why).

The Admin account username cannot be changed or deleted, nor can the Admin account be demoted out of admin status.

## Adding a new user
To add a new user the front end should POST the following struct/JSON to /api/user/

```
{
     User: "chuck",
     Pass: "chuckTesta4Eva",
     Name: "Chuck Testa",
     Email: "chuck@testa.net"
     Admin: true,
}
```

All fields must be populated.  The back end will respond with the standard response JSON as mentioned above.

## Locking a user account
To Lock an account the front end should send an empty POST to /api/user/{id}/lock with the {id} of the UID for the user that will be locked.  Locking a user account will immediately stop any new connections to the backend, even if a session is active.

## UnLocking a user account
To UnLock an account the front end should send an empty DELETE to /api/user/{id}/lock with the {id} for the UID of the user that will be unlocked.  The backend will respond with success regardless of whether the account was actually unlocked if the action is allowed.  We do this because locking a locked account ends in the state that the account is locked.  So its all good.  Same with unlocking.

## Changing user info
To change user info the front end should send a PUT /api/user/{id}/ with replaced account JSON:

```
{
     User: "chuck",
     Name: "Chuck Testa",
     Email: "chuck@testa.net"
}
```

Any field that is NOT populated will be ignored.  The backend will respond with the standard response JSON as mentioned above.  If the current user is NOT an admin and not changing their own account, the request will be rejected.  Admins can change information for any account.  The primary admin account (UID zero) cannot change its admin status.  The backend responds with a 200 on success and 400-500 on error depending on who caused the error and why.  Error messages will be returned in the body of the response and are human displayable.

## Changing a users password
To change a password the frontend should PUT JSON to the url /api/user/{id}/pwd
If the User is an admin and changing the password for an account they DO NOT OWN, the OrigPass field is NOT required.
```
{
     OrigPass: "my old password was bad",
     NewPass: "thisis mynewpassword",
}
```

## Deleting a user
To Delete a user the front end should send a DELETE to /api/user/{id}/ with the id of the UID to delete.  Only Admins can use this facility, and a user cannot delete their own account.  The Primary admin (UID zero) cannot be deleted.

## Changing admin status
To change the admin status of a user or query the admin status the frontend should hit /api/user/{id}/admin
with an empty body where the method specifies the action.  The backend will respond with 200 on success and 400-500 on error depending on who munged the request and why.
```
GET - returns current admin status
PUT - sets the user as an admin and returns the new status
DELETE - removes admin status for the user


//example JSON on success
{
     UID: 1,
     Admin: true,
}
```

## Getting a single users account information
To get a single users account information the front end should send a GET to /api/user/{id}/.  Admins can retrieve any account, non-admins can ONLY retrieve their own account information.  Response of 200 will contain valid JSON in the body, 400-500 means someone boned the request and the body will contain an error message.  The backend will respond with a JSON packet containing the account info as follows:
```

//the JSON returned on success
{
    Status: bool,
    Reason: string
    Info: {
        UID: int,
        User: string,
        Name: string,
        Email: string,
        Admin: bool,
        Locked: bool,
        TS: string
    }
}
//an example successful response
{
    Status: true,
    Reason: "",
    Info: {
        UID: 5,
        User: "ChuckTesta",
        Name: "Chuck Testa",
        Email: "ChuckTesta@gmail.com",
        Admin: false,
        Locked: false,
        TS: "05-12-2014 13:05:01.2242 +0000 UTC"
    }
}

```

## account and group info
The info URLs can be hit by any authenticated user.  Hitting the urls gets the current users info.
### Base URL for account user info /api/info/whoami
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
                        "Name": "TheNinos",
                        "Desc": "This is the Ninos group"
                }
        ]
}

```
