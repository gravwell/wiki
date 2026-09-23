---
myst:
  substitutions:
    package: "gravwell-hosted-runner"
    standalone: "gravwell_hosted_runner"
    dockername: "hosted_runner"
---
# Thinkst Canary Ingester

The [Thinkst Canary](https://canary.tools/) ingester polls the Canary Console API for incidents and audit trail events.

* **Incidents** (`Api=incident`) incidents raised by Canary birds and Canarytokens: connections, HTTP requests, file access, and other bait interactions. Events are ingested in order with timestamps preserved from the original records.
* **Audit Trail** (`Api=audit`) console-level activity: logins, configuration changes, and bird/token operations performed through the Console.

Each stanza polls exactly one Api, so a deployment that wants both feeds needs two stanzas, each writing to its own tag.

This ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration). Multiple Thinkst stanzas can coexist alongside other Hosted Runner plugins in a single configuration file.

## Installation

```{include} installation_instructions_template 
```

## Configuration

To configure the ingester you will need the following from your Canary Console:

* **Domain**: Your Canary Console's hostname, e.g. `example.canary.tools`. Do not include the protocol (`https://`) or a trailing slash.
* **Auth Token**: An API auth_token generated from the Console's API settings. This should be a dedicated Read-Only key, not an Admin or Analyst key, and not a token shared with any other integration.

See the [Canary API documentation](https://docs.canary.tools/console-settings/api.html) for background on API keys.

### Creating a Thinkst Auth Token

Log in to your Canary Console and go to Settings > API (under Global Settings). If the Console API hasn't been enabled yet, it must be enabled from the Console UI first; it cannot be enabled via the API itself.

From the API settings page, add a new Global API key:

* Give it a Note like "Gravwell Hosted Runner" so it's identifiable later in the key list and the audit trail.
* Set its role to Read-Only. A Read-Only key can fetch incidents and the audit trail but can't modify the Console, birds, or tokens.
* Optionally restrict it to the IP address(es) the Hosted Runner connects from, via Allowed IP Ranges.

```{attention}
Do **not** use an Admin or Analyst key for the ingester. Those grant significantly more access than is needed for monitoring. 
```

The auth_token value is only shown once so make sure to copy and store it for later use in the `Token` config parameter.

### Request Limits

The `Requests-Per-Minute` parameter controls how often the ingester calls the Console API (default 60/minute). The Canary API doesn't publish a fixed per-console rate limit the way some other vendors do, but polling faster than necessary provides no benefit and adds needless load to the Console; the default is a reasonable starting point for most deployments.

### Thinkst Stanza Parameters

The Thinkst ingester is configured via `[Thinkst "name"]` stanzas in the Hosted Runner configuration file, typically `/opt/gravwell/etc/hosted_runner.conf`. The `[Global]` and `[State]` blocks common to all Hosted Runner plugins are described in [Hosted Runner Configuration](hosted_runner_configuration).

| Config Parameter    | Type    | Required | Default | Description                                                                                                                      |
|----------------------|---------|----------|---------|------------------------------------------------------------------------------------------------------------------------------------|
| Ingester-UUID        | UUID    | yes      |         | A unique UUID for this ingester instance. Used for state tracking.                                                                 |
| Domain               | string  | yes      |         | Your Canary Console hostname, e.g. `example.canary.tools`. No scheme or trailing slash.                                            |
| Token                | string  | yes      |         | Canary Console API auth_token.                                                                                                     |
| Api                  | string  | yes      |         | Which feed this stanza polls: `incident` or `audit`.                                                                               |
| Tag-Name             | string  | no       | thinkst | Tag to write events to. If unset, defaults to Tag-Prefix (or `thinkst` if that's unset too) plus the Api, e.g. `thinkst-incident`. |
| Tag-Prefix           | string  | no       | thinkst | Prefix used to build the default tag when Tag-Name isn't set. Cannot be set together with Tag-Name.                                |
| Lookback             | integer | no       | 24      | Hours of history to consider when an instance first starts with no prior state. Only affects the `audit` Api.                      |
| Requests-Per-Minute  | integer | no       | 60      | Maximum number of Console API requests per minute.                                                                                 |
| Request-Interval     | integer | no       | 60      | Seconds to wait between polls once a feed has caught up (no further pages pending).                                                |

## Example Configuration

A minimal setup polling both feeds from a single Console, using the default tag names (`thinkst-incident` and `thinkst-audit`):

```
[Thinkst "incidents"]
    Ingester-UUID="99b00000-0000-0000-0000-000000000000"
    Domain="example.canary.tools"
    Token="your-thinkst-auth-token"
    Api=incident

[Thinkst "audit"]
    Ingester-UUID="99b00000-0000-0000-0000-000000000001"
    Domain="example.canary.tools"
    Token="your-thinkst-auth-token"
    Api=audit
```

With explicit tag names and a longer initial lookback window for the audit feed:

```
[Thinkst "incidents"]
    Ingester-UUID="99b00000-0000-0000-0000-000000000000"
    Domain="example.canary.tools"
    Token="your-thinkst-auth-token"
    Api=incident
    Tag-Name="thinkst-incidents"

[Thinkst "audit"]
    Ingester-UUID="99b00000-0000-0000-0000-000000000001"
    Domain="example.canary.tools"
    Token="your-thinkst-auth-token"
    Api=audit
    Tag-Name="thinkst-audit"
    Lookback=72
```

## Additional Resources

* [Canary API — Getting Started](https://docs.canary.tools/guide/getting-started.html)
* [Canary API — Console API Settings](https://docs.canary.tools/console-settings/api.html)
* [Canary API — Incidents Queries](https://docs.canary.tools/incidents/queries.html)
* [Canary API — Audit Trail](https://docs.canary.tools/console/audit-trail.html)
