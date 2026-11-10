# Mimecast

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Mimecast Hosted Ingester](/ingesters/mimecast)
:::

## Mimecast Configuration

To configure the ingester you will need the following from Mimecast:

* **Client ID**: The OAuth 2.0 client ID for your API 2.0 integration
* **Client Secret**: The OAuth 2.0 client secret for your API 2.0 integration

See the [Mimecast documentation](https://mimecastsupport.zendesk.com/hc/en-us/articles/34000360548755-API-Integrations-Managing-API-2-0-for-Cloud-Gateway) for instructions on creating an API 2.0 integration and obtaining these credentials.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/mimecast-well.conf`
```ini
[Storage-Well "mimecast"]
    Location=/opt/gravwell/storage/mimecast
    Tags=mimecast*
```

### Gravwell Ingester Configuration: Mimecast

If the Hosted Runner is not installed, follow the [configuration guide for Mimecast](/ingesters/mimecast) to create your own configuration.

Edit: `/opt/gravwell/etc/hosted_runner.conf`
```ini
[Mimecast "mimecast"]
    Ingester-UUID="9a000000-0000-0000-0000-000000000000"
    Client-Id="your-client-id"
    Client-Secret="your-client-secret"
    Api=mta-delivery
    Api=mta-receipt
    Api=mta-av
    Api=mta-spam
    Api=audit
    Tag-Prefix="mimecast"
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_hosted_runner.service`
```

## Additional Resources
* [Mimecast API Overview](https://developer.services.mimecast.com/api-overview)
* [Mimecast SIEM Tutorial](https://developer.services.mimecast.com/siem-tutorial-cg)
* [Mimecast API 2.0 Setup Guide](https://mimecastsupport.zendesk.com/hc/en-us/articles/34000360548755-API-Integrations-Managing-API-2-0-for-Cloud-Gateway)
