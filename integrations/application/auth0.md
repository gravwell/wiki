# Auth0

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
:::

## Auth0 Configuration

Follow the auth0 docs for setting up a [Custom Log Stream using Webhooks](https://auth0.com/docs/customize/log-streams/custom-log-streams).

```{note}
Auth0 does not support using self-signed HTTP certificates.
```

1. Go to `Dashboard > Monitoring > Streams > Create Stream > Custom Webhook`
2. Configure the settings:
   * **Name:** Enter a unique name for your new stream.
      * Example: `Gravwell Webhook`
   * **Payload URL:** Sets where the event payloads are sent as HTTP Post Requests.
      * Example: `https://path.to.gravwell:port/auth0`
   * **Authorization Token:** (Optional) The value in the Authorization header of the request.
      * Example: `AuthenticationToken`
   * **Content Type:** The media type of the payload that will be delivered to the webhook. 
      * Example: `application/json`
   * **Content Format:** Receive data in JSON lines, arrays, or objects.
      * Example: `JSON Lines`
   * **Filter by Event Category:** List of log stream filters.
      * Example: `Filter: All`
   * **Starting Cursor:** (Optional) Specific day and time to start the stream from.
3. Click `Save`.

Verify the `Stream Status` is active in the `Health` view.


## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/auth0-well.conf`
```ini
[Storage-Well "auth0"]
    Location=/opt/gravwell/storage/auth0
    Tags=auth0*
```
### Gravwell Ingester Configuration: HTTP
**Sample Auth0 HTTP config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/auth0.conf`
```ini
[Listener "auth0"]
    URL="/auth0"
    #TokenValue= "AuthenticationToken"
    Tag-Name=auth0
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```