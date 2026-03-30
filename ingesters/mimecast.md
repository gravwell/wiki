---
myst:
  substitutions:
    package: "gravwell-hosted-runner"
    standalone: "gravwell_hosted_runner"
    dockername: "hosted_runner"
---
# Mimecast Ingester

The [Mimecast](https://www.mimecast.com/) ingester polls the Mimecast MTA SIEM and Audit APIs. It ingests email security events including delivery, receipt, spam, AV, URL protection, impersonation protection, attachment protection, and audit events. Timestamps are preserved from the original events to maintain accuracy even across polling gaps or downtime.

This ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration). Multiple Mimecast stanzas can coexist alongside other Hosted Runner plugins in a single configuration file.

## Installation

```{include} installation_instructions_template 
```

If you already have the hosted runner installed, you can modify the config.

## Configuration

To configure the ingester you will need the following from Mimecast:

* **Client ID**: The OAuth 2.0 client ID for your API 2.0 integration
* **Client Secret**: The OAuth 2.0 client secret for your API 2.0 integration

See the [Mimecast documentation](https://mimecastsupport.zendesk.com/hc/en-us/articles/34000360548755-API-Integrations-Managing-API-2-0-for-Cloud-Gateway#h_01KFBA474MS5X46Z6H5XRNKPJR) for instructions on creating an API 2.0 integration and obtaining these credentials.

The Mimecast ingester is configured via `[Mimecast "name"]` stanzas in the Hosted Runner configuration file, typically `/opt/gravwell/etc/hosted_runner.conf`. The `[Ingest]` and `[State]` blocks common to all Hosted Runner plugins are described in [Hosted Runner Configuration](hosted_runner_configuration).

### Mimecast Stanza Parameters

Each `[Mimecast "name"]` stanza configures an independent polling connection to the Mimecast API. You can define multiple stanzas to ingest from different API endpoints or with different tag configurations.

| Config Parameter    | Type              | Required | Default Value                     | Description                                                                                                     |
|---------------------|-------------------|----------|-----------------------------------|-----------------------------------------------------------------------------------------------------------------|
| Ingester-UUID       | UUID              | yes      |                                   | A unique UUID for this ingester instance. Used for state tracking.                                              |
| Client-Id           | string            | yes      |                                   | OAuth 2.0 client ID from your Mimecast API 2.0 integration.                                                     |
| Client-Secret       | string            | yes      |                                   | OAuth 2.0 client secret from your Mimecast API 2.0 integration.                                                 |
| Api                 | string            | yes      |                                   | The Mimecast API to poll. Can be specified multiple times. See [Available APIs](available-apis).                |
| Host                | URL               | no       | https://api.services.mimecast.com | The Mimecast API base URL. Override for regional endpoints or testing.                                          |
| Lookback            | integer           | no       | 24 (hours)                        | How far back in time to fetch events on first run in hours.                                                     |                                                        
| Tag-Name            | string            | no       | (derived from API name)           | Tag to assign ingested entries. Only valid when a single `Api` is configured. Cannot be used with `Tag-Prefix`. |
| Tag-Prefix          | string            | no       |                                   | Prefix for auto-generated tag names. Tags will be `<prefix>-<api>`. Cannot be used with `Tag-Name`.             |
| Requests-Per-Minute | integer           | no       | 5                                 | Maximum number of API requests per minute.                                                                      |
| Request-Interval    | integer (seconds) | no       | 300 (seconds)                     | How often to poll the API for new events in seconds                                                             |

(available-apis)=
### Available APIs

The following API values can be specified in the `Api` parameter:

| API Value           | Description                                                 |
|---------------------|-------------------------------------------------------------|
| `audit`             | Mimecast audit events (admin actions, policy changes, etc.) |
| `mta-delivery`      | MTA delivery events                                         |
| `mta-receipt`       | MTA receipt events                                          |
| `mta-process`       | MTA process events                                          |
| `mta-av`            | MTA antivirus scan events                                   |
| `mta-spam`          | MTA spam detection events                                   |
| `mta-internal`      | Internal Email Protect events                               |
| `mta-impersonation` | Impersonation Protect events                                |
| `mta-url`           | URL Protect events                                          |
| `mta-attachment`    | Attachment Protect events                                   |
| `mta-journal`       | MTA journal events                                          |

The `audit` API uses the [Mimecast Audit Events API](https://developer.services.mimecast.com/docs/auditevents/1/routes/api/audit/get-audit-events/post). All `mta-*` APIs use the [SIEM API](https://developer.services.mimecast.com/docs/threatssecurityeventsanddataforcg/1/routes/siem/v1/events/cg/get).

### Tag Naming

By default (without `Tag-Name` or `Tag-Prefix`), ingested entries are tagged using the API name directly (e.g., `mta-delivery`, `audit`).

Use `Tag-Prefix` to namespace tags. For example, `Tag-Prefix=mimecast` produces tags like `mimecast-audit`, `mimecast-mta-delivery`, etc. This is useful when ingesting data from a variety of data sources to keep it clear where data came from. 

Use `Tag-Name` to assign a fixed tag when ingesting from exactly one API.

```{note}
`Tag-Name` and `Tag-Prefix` are mutually exclusive. `Tag-Name` can only be used when a single `Api` value is configured.
```

## Example Configuration

The following example shows two Mimecast stanzas: one ingesting MTA delivery events with an explicit tag, and one ingesting audit events with a tag prefix applied. The `[Ingest]` and `[State]` blocks are omitted here — see [Hosted Runner Configuration](hosted_runner_configuration) for those common settings.

```
[Mimecast "mta"]
    Ingester-UUID="99000000-0000-0000-0000-000000000000"
    Client-Id="your-client-id"
    Client-Secret="your-client-secret"
    Api=mta-delivery
    Tag-Name=mimecast-delivery

[Mimecast "audit"]
    Ingester-UUID="99a00000-0000-0000-0000-000000000000"
    Client-Id="your-client-id"
    Client-Secret="your-client-secret"
    Api=audit
    Tag-Prefix="mimecast"
```

To ingest all MTA SIEM event types into individually tagged streams:

```
[Mimecast "mta"]
    Ingester-UUID="99b00000-0000-0000-0000-000000000000"
    Client-Id="your-client-id"
    Client-Secret="your-client-secret"
    Api=mta-delivery
    Api=mta-receipt
    Api=mta-process
    Api=mta-av
    Api=mta-spam
    Api=mta-internal
    Api=mta-impersonation
    Api=mta-url
    Api=mta-attachment
    Tag-Prefix="mimecast"
```

## Additional Resources

* [Mimecast API Overview](https://developer.services.mimecast.com/api-overview)
* [Mimecast SIEM Tutorial](https://developer.services.mimecast.com/siem-tutorial-cg)
* [Mimecast API 2.0 Setup Guide](https://mimecastsupport.zendesk.com/hc/en-us/articles/34000360548755-API-Integrations-Managing-API-2-0-for-Cloud-Gateway#h_01KFBA474MS5X46Z6H5XRNKPJR)
* [Mimecast Audit API Reference](https://developer.services.mimecast.com/docs/auditevents/1/routes/api/audit/get-audit-events/post)
* [Mimecast SIEM Events API Reference](https://developer.services.mimecast.com/docs/threatssecurityeventsanddataforcg/1/routes/siem/v1/events/cg/get)
* [Mimecast SIEM Batch API Reference](https://developer.services.mimecast.com/docs/threatssecurityeventsanddataforcg/1/routes/siem/v1/batch/events/cg/get)
