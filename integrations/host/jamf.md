# Jamf

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Jamf Hosted Ingester](/ingesters/jamf)
:::

## Jamf Configuration

### Jamf Pro API Roles and Clients

To configure Jamf for ingestion with the API you will need the following:

- **Host:** Your Jamf Pro base URL (e.g. `https://yourserver.jamfcloud.com`).

- **Client ID:** The OAuth 2.0 client ID for a dedicated Jamf Pro API integration.

- **Client Secret:** The OAuth 2.0 client secret generated for that same API integration. This should be created for a dedicated, least-privilege API role and client, not credentials tied to an individual admin account.

See the [Jamf documentation](https://learn.jamf.com/r/en-US/jamf-pro-documentation-current/API_Roles_and_Clients) for instructions on creating an API role and API client.

#### Creating a Jamf API Role

Start by creating a dedicated API role scoped only to what the ingester needs (something like "Gravwell Ingest Read Only"). The Jamf ingester reads from the computer-inventory endpoint, so the role only needs read privileges for the inventory categories you plan to collect (e.g. General, Hardware, Operating System, and any additional [Sections](https://docs.gravwell.io/ingesters/jamf.html#available-sections) you enable).

```{attention}
Do not assign write, update, or delete privileges to the API role used by the ingester. This gives significantly more access than is needed for read-only log collection.
```

1. In Jamf Pro, go to **Settings > System > API Roles and Clients**.

2. On the **API Roles** tab, click **+ New**.

3. Name the role, then select only the read privileges required for the inventory sections you intend to ingest.

4. Click **Save**.

#### Creating a Jamf API Client

1. In **API Roles and Clients**, switch to the **API Clients** tab and click **+ New**.

2. Name the client and assign the API role created above.

3. Enable the API client, then click **Save**.

4. Click **Generate Client Secret** and copy the value immediately. The secret is only shown once.

5. Copy the **Client ID** shown on the API client details page for use in the ingester configuration.

```{attention}
Store the client secret securely; if it is lost, you must generate a new one from the API client's details page, which will invalidate the previous secret.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

#### Sample well config
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/jamf-well.conf`
```ini
[Storage-Well "jamf"]
    Location=/opt/gravwell/storage/jamf
    Tags=jamf*
```

### Gravwell Ingester Configuration

#### Sample Jamf config: Jamf Hosted Ingester
If the Hosted Runner is not installed, follow the [configuration guide for Jamf](/ingesters/jamf) to create your own configuration.  

Edit: `/opt/gravwell/etc/hosted_runner.conf`

```ini
[Jamf "yourserver"]
    Ingester-UUID="99d00000-0000-0000-0000-000000000000"
    Host=https://yourserver.jamfcloud.com
    Client-Id="api-client-id"
    Client-Secret="api-client-secret"
```

#### Sample Jamf config: Additional inventory sections
```ini
[Jamf "yourserver"]
    Ingester-UUID="99d00000-0000-0000-0000-000000000000"
    Host=https://yourserver.jamfcloud.com
    Client-Id="api-client-id"
    Client-Secret="api-client-secret"
    Sections=OPERATING_SYSTEM
    Sections=SERVICES
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_hosted_runner.service`
```
