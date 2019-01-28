# Logging APIs

The API provides utilities for admin users to manage logging and inject log entries into Gravwell's on-disk log files.

## Show/set webserver logging level

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

## Inject logs

Admin users can insert log entries by sending POST requests to the appropriate URLs. These logs will be written out to the Gravwell webserver's on-disk log files.

The URLs are:

```
/api/logging/access
/api/logging/info
/api/logging/warn
/api/logging/error
```

The POST request body should contain a JSON structure with a field named 'Body' containing the desired log message:

```
{
	"Body": "This is my error message"
}
```

The server will return boolean 'true' on success.