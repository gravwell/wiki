# Web API to show webserver logging level and optionally set it

This API allows admins to show current logging level as well as available logging levels.
An admin can change the logging level at will via this API.

To get the current logging level as well as available logging levels perform a GET request to /api/logging
The request will return the following

```
{
     Levels: ["Error", "Warn", "Off"],
     Current: "Error"
}
```

To set the log level perform a PUT to /api/logging/ with the following

```
{
     Level: "Error"
}
```