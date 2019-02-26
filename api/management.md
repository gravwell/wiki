# System Management

This page documents admin-only APIs for managing the Gravwell processes and configuration.

## Restarting Gravwell

Two APIs are provided to restart Gravwell processes; one to restart the webserver, and one to restart the indexers. In both cases, "restarting" is accomplished by shutting down the process and allowing systemd (or whatever init system is in use) to restart it.

### Restarting the webserver

To restart the webserver, send a POST request with an empty body to `/api/restart/webserver`. This should trigger the webserver to shut down and restart immediately.

### Restarting the indexers

To restart all indexers to which the webserver is currently connected, send a POST request with an empty body to `/api/restart/indexers`. The webserver will signal to each indexer that it should shut itself down and restart. As the individual indexers come back up, the webserver will reconnect automatically.

### Checking for a Distributed Frontend

To check whether the Gravwell cluster is operating in a distributed frontend mode, perform a GET on `/api/distributed`.  The webserver will responde with a JSON object indicating whether the frontend is configured in a distributed mode.

An example response when not in distributed mode:

```
{
	"Distributed": false
}
```
