# Slack

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, Slack Hosted Ingester
:::

## Slack Configuration

To configure Slack for ingestion with the API you will need the following:

- **Slack Enterprise organization:** The Audit Logs API is available only to Slack Enterprise organizations. It is not available on Slack Pro, Business+, or a single standalone workspace.
- **Organization-scoped OAuth token:** A token from a Slack app installed at the Enterprise organization level, granted only the `auditlogs:read` scope.
- **Enterprise organization ID:** Begins with the letter `E`; used as the `Scope-Identity` config value.

See the [Slack Audit Logs API documentation](https://docs.slack.dev/admins/audit-logs-api/) for instructions and background. This API only covers organization and workspace-management events (membership, settings, app installs, and similar administrative actions).

### Creating an Organization-Scoped Slack Token

Start by creating a dedicated Slack app for this integration rather than reusing one built for another purpose.

```{attention}
Grant only the `auditlogs:read` scope. No delete or write scope is ever required. The Slack plugin only issues `GET` requests against the Audit Logs API.
```

1. In the app's **OAuth & Permissions** settings, add the `auditlogs:read` scope.
2. Install the app **at the Enterprise organization level**, not an individual workspace. As the organization Owner, confirm the install target is the Enterprise organization itself.
3. Confirm the token is organization-scoped by calling Slack's [`auth.test`](https://docs.slack.dev/reference/methods/auth.test/) method with it. If the response includes an `enterprise_id` field, the token is correctly organization-scoped; if it is absent, reinstall the app at the organization level.
4. Record the Enterprise organization ID from that same `auth.test` response's `enterprise_id` field. This is the `Scope-Identity` value the ingester configuration needs.
5. Save the token to a file that contains only the token on a single line, for example `/opt/gravwell/secrets/slack-token`. Restrict access to the file (for example, `chmod 0600`) so that only the user running the Hosted Runner can read it. This path is the `Credential-File` value in the ingester configuration.

```{note}
A workspace-scoped token cannot call the Audit Logs API and will fail authentication.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

#### Sample Well Config
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/slack-well.conf`
```ini
[Storage-Well "slack"]
    Location=/opt/gravwell/storage/slack
    Tags=slack*
```

### Gravwell Ingester Configuration: Slack Hosted Ingester

This ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration). If the Hosted Runner is not installed, install it first.

```{note}
The Slack plugin is still in development (see [gravwell/gravwell#2786](https://github.com/gravwell/gravwell/pull/2786)) and is not yet available in a released Hosted Runner. The options below come from that development version and may change before release.
```

Edit: `/opt/gravwell/etc/hosted_runner.conf`

```ini
[Slack "primary"]
    Ingester-UUID="99e00000-0000-0000-0000-000000000000"
    Base-URL="https://api.slack.com"
    Credential-File="/opt/gravwell/secrets/slack-token"
    Scope-Identity="E00000000"
    Tag-Name=slack-audit
```

To use a shorter initial lookback and tighter rate limiting, add `Lookback` and `Requests-Per-Minute`. `Lookback` accepts durations such as `1d` or `24h`:

```ini
[Slack "primary"]
    Ingester-UUID="99e00000-0000-0000-0000-000000000000"
    Base-URL="https://api.slack.com"
    Credential-File="/opt/gravwell/secrets/slack-token"
    Scope-Identity="E00000000"
    Tag-Name=slack-audit
    Lookback=1d
    Requests-Per-Minute=30
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_hosted_runner.service`
```

## Additional Resources
* [Slack Audit Logs API](https://docs.slack.dev/admins/audit-logs-api/)
* [Slack `auth.test` method](https://docs.slack.dev/reference/methods/auth.test/)
