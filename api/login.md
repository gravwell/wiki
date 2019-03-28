# Login subsystem

## Login

A GET on this page redirects to /api/login/login.html

To login you POST JSON to /api/login/ with the following structure:
```
{
    "User": "username",
    "Pass": "password"
}
```

and the server will respond with the following to indicate whether login was successful.

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

## Logout

* PUT /api/logout - logs your current instance out
* DELETE /api/logout - logs out ALL your users instances

## JWT protections are enforced on all POSTs
The JWT received from the login API must be included as an Authorization Bearer header on all other API requests.

```Authorization: Bearer reallylongjsonwebtokenstringishere```

## View active sessions
GET the /api/account/{id}/sessions and it will return a chunk of JSON.  Admins can request any users sessions, users can ONLY request their own sessions.

```
{
     User:  "user",
     [
          {
              Origin: "192.168.4.2",
              LastHit: "2014-12-1 12:32:02.2332422Z07:00"
          },
          {
              Origin: "1.35.2.2",
              LastHit: "2014-12-1 12:05:02.2332422Z07:00"
          },
          {
              Origin: "10.0.0.22",
              LastHit: "2014-11-1 12:05:02.2332422Z07:00"
          }
     ]
}
```
