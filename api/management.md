# System Management

This page documents admin-only APIs for managing the Gravwell processes and configuration.

## Restarting Gravwell

Two APIs are provided to restart Gravwell processes; one to restart the webserver, and one to restart the indexers. In both cases, "restarting" is accomplished by shutting down the process and allowing systemd (or whatever init system is in use) to restart it.

### Restarting the webserver

To restart the webserver, send a POST request with an empty body to `/api/restart/webserver`. This should trigger the webserver to shut down and restart immediately.

### Restarting the indexers

To restart all indexers to which the webserver is currently connected, send a POST request with an empty body to `/api/restart/indexers`. The webserver will signal to each indexer that it should shut itself down and restart. As the individual indexers come back up, the webserver will reconnect automatically.

### Checking for a Distributed Frontend and deployment info

To check whether the Gravwell cluster is operating in a distributed frontend mode, perform a GET on `/api/deployment`.  The webserver will responde with a JSON object indicating whether the frontend is configured in a distributed mode.

An example response when not in distributed mode:

```
{
	"Distributed": false,
	"DefaultLanguage": "en_US",
}
```


### Performing a system backup

Admin users may request a system backup which will provide a backup file containing all content related to the state of Gravwell.

A system backup can be used to save user and group accounts, dashboards, kits, query libraries, and even saved searches.  This is essentially everything but data and system configuration.

A backup is obtained by performing a `GET` request on `/api/backup` as an admin user, the API will then return a file download with the backup file.

By default a backup does not contain any saved searches; to include saved searches in the backup append the `savedsearch=true` URL parameter on the GET request.

The `/api/backup` API may be authenticated using either the JWT authorization token or a cookie.

NOTE: Only a single backup and/or restore may take place at any given time.

### Restoring from a system backup

Admin users may restore the system from a backup archive by performing a `POST` to the `/api/backup` API and uploading a backup file using a multipart form.

The API expects that the backup file be located in the `backup` form field and be uploaded using a multipart form.

The `/api/backup` API may be authenticated using either the JWT authorization token or a cookie.

NOTE: A restoration is a complete restoration, any changes to users, content, or saved searches will after the restoration point will be lost.

NOTE: Only a single backup and/or restore may take place at any given time.

NOTE: Once the restoration begins, all sessions will be terminated.  Upon successful completion of a restore, all users will need to log back in.
