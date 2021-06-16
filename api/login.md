# Login subsystem

## Login

To login, POST JSON to `/api/login` with the following structure:

```
{
    "User": "username",
    "Pass": "password"
}
```

and the server will respond with the following to indicate whether login was successful:

```
{
  "LoginStatus":true,
  "JWT":"reallylongjsonwebtokenstringishere"
}

```

If the login failed, the server will return a structure with a "reason" property:
```
{
  "LoginStatus":false,
  "Reason":"Invalid username or password"
}
```

Instead of sending JSON, you may also set form fields "User" and "Pass" in the login POST request.

## Logout

* PUT /api/logout - logs your current instance out
* DELETE /api/logout - logs out ALL your user's instances

## JWT protections are enforced on all requests that are not used for file download operations.
The JWT received from the login API must be included as an Authorization Bearer header on all other API requests.

```Authorization: Bearer reallylongjsonwebtokenstringishere```

### Websocket Authentication

As a convenience, the websocket API endpoints will also look for the JWT token in the `Sec-Websocket-Protocol` header value.  Many websocket implementations do not properly support passing header values, so we overload the websocket subprotocol negotiation header.  The API endpoints will still look for the standard `Authentication` header values as well.

## View active sessions
Send a GET to `/api/users/{id}/sessions` and it will return a chunk of JSON.  Admins can request any users sessions, users can ONLY request their own sessions.

```
{
    "Sessions": [
        {
            "LastHit": "2020-08-04T15:28:12.601899275-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        },
        {
            "LastHit": "2020-08-03T23:59:53.807610997-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        },
        {
            "LastHit": "2020-08-04T09:45:48.291770859-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        }
    ],
    "UID": 1,
    "User": "admin"
}
```
