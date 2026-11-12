# FleetDM

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http) <br /> [HTTP - HEC](http_splunk_hec_compatibility)
    Preprocessor, [Regex Timestamp Extraction](/ingesters/preprocessors/regextimestamp.md)
:::

[Fleet](https://fleetdm.com) is an open-source device management and osquery platform. This guide sends Fleet's **activity feed** to Gravwell. The activity feed is Fleet's admin audit trail, such as logins and failed logins, live queries, script runs, software installs, user and role changes, and policy edits. Fleet's [audit logs reference](https://fleetdm.com/docs/using-fleet/audit-logs) lists every activity type.

This guide does not cover osquery result logs, osquery status logs or Fleet's other webhooks (failing policies, host status, vulnerabilities). If you send those to Gravwell, use a different tag and listener URL; do not send them to the activity tag.

Fleet can deliver activities in two ways. Use **one** of them per tag:

| | Option 1: Audit log streaming | Option 2: Activities webhook |
|---|---|---|
| Fleet edition | Premium | Free and Premium |
| Fleet hosting | Self-hosted (needs Fleet server settings) | Self-hosted and Fleet-managed cloud |
| Fleet version | 4.90.0 or later | Any current version |
| Delivery | Every 5 minutes | Within seconds |
| If Gravwell is unreachable | Fleet resends on the next run | The activity is lost |
| Gravwell listener | `HEC-Compatible-Listener` | `Listener` |

Use Option 1 if your Fleet qualifies, because it does not lose activities. Use Option 2 otherwise.

## FleetDM Configuration

```{important}
Fleet 4.89.0 and later refuse to send to private network addresses (for example `10.0.0.0/8`, `172.16.0.0/12` and `192.168.0.0/16`). If the HTTP Ingester has a private address, set `FLEET_SERVER_ALLOW_PRIVATE_NETWORK_INTEGRATIONS=true` on the Fleet server and restart Fleet. Fleet always refuses loopback addresses (`127.0.0.1`), even with this setting, so use the ingester host's network address even when Fleet runs on the same host. Fleet-managed cloud customers cannot change this setting and need an ingester address that is reachable from the internet.
```

Fleet does not test the destination when you save the settings. A wrong address, token or certificate shows up only as missing data and an error in the Fleet server log (see Troubleshooting).

### [Option 1] Audit Log Streaming (Fleet Premium, Self-Hosted)

Fleet Premium can stream activities with its `splunk` audit log plugin, which speaks the Splunk HTTP Event Collector protocol. Set the following environment variables on the Fleet server and restart Fleet:

```bash
FLEET_ACTIVITY_ENABLE_AUDIT_LOG=true
FLEET_ACTIVITY_AUDIT_LOG_PLUGIN=splunk
FLEET_SPLUNK_URL=https://<HTTP-INGESTER-HOST>:<PORT>/fleetdm
FLEET_SPLUNK_TOKEN=<TOKEN>
```

* `FLEET_SPLUNK_URL` ends in `/fleetdm`. Fleet adds `/services/collector/event` to it, so it posts to `/fleetdm/services/collector/event`, which the listener below receives.
* `<TOKEN>` is the `TokenValue` of the Gravwell listener below.
* The same settings can be made in the Fleet configuration file (`activity.enable_audit_log`, `activity.audit_log_plugin`, `splunk.url`, `splunk.token`). See Fleet's [server configuration](https://fleetdm.com/docs/configuration/fleet-server-configuration).
* Fleet sends the activities recorded since its last run every 5 minutes. It marks an activity as sent only after Gravwell accepts it, so activities recorded while Gravwell is down arrive on a later run.
* To undo, remove the four variables and restart Fleet.

```{warning}
The `splunk` settings are shared with Fleet's osquery log plugins. If you also set `FLEET_OSQUERY_RESULT_LOG_PLUGIN` or `FLEET_OSQUERY_STATUS_LOG_PLUGIN` to `splunk`, osquery results go to the same listener and land on the activity tag. Use another plugin for osquery logs.
```

### [Option 2] Activities Webhook (Fleet Free and Premium)

Fleet's activities webhook posts each activity to a URL as it happens. It cannot send a header, so the Gravwell listener takes its token from the URL. Set it as a global admin. On Fleet Premium, select **All fleets** in the fleet selector first; Fleet Free has no selector.

1. Open the **Dashboard**.
2. On the **Activity** card, select **Manage automations**.
3. Turn the slider to **Enabled**.
4. In **Destination URL**, enter `https://<HTTP-INGESTER-HOST>:<PORT>/fleetdm?token=<TOKEN>`, where `<TOKEN>` is the `TokenValue` of the Gravwell listener below.
5. Select **Save**.

Or use the Fleet REST API with an API token for a global admin:

```bash
curl -X PATCH "https://<FLEET-HOST>/api/v1/fleet/config" \
  -H "Authorization: Bearer <FLEET-API-TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"webhook_settings":{"activities_webhook":{"enable_activities_webhook":true,"destination_url":"https://<HTTP-INGESTER-HOST>:<PORT>/fleetdm?token=<TOKEN>"}}}'
```

If Fleet's GitOps mode is on, set the same two keys under `org_settings.webhook_settings.activities_webhook` in your GitOps files instead.

To undo, turn the slider off in **Manage automations** and select **Save**, or send the same API request with `"enable_activities_webhook":false`.

```{warning}
* **Activities are lost when a delivery fails.** Fleet retries only when Gravwell answers HTTP 429 (busy). A connection error, a certificate error, a wrong token or a restart of the ingester loses the activities recorded at that moment. They remain in Fleet's own activity feed.
* **The token is visible in Fleet and in Gravwell.** Fleet records the whole destination URL, token included, in the `enabled_activity_automations` and `edited_activity_automations` activities, which appear in Fleet's activity feed and on the Gravwell tag. The token only allows writing to the listener. Restrict who can search the tag, limit the listener to the Fleet server's address where you can, and rotate the token if it is exposed.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/fleetdm-well.conf`
```ini
[Storage-Well "fleetdm"]
    Location=/opt/gravwell/storage/fleetdm
    Tags=fleetdm
```

### Gravwell Ingester Configuration: HTTP Ingester

This guide assumes the HTTP Ingester is already installed and configured as described in [HTTP Ingester Setup](/integrations/generic/http_ingester). Add the following to a new overlay file; you do not need a separate HTTP Ingester for this integration.

Serve the HTTP Ingester over HTTPS (see HTTP Ingester Setup). Both options send the token with every request. The certificate must be one the Fleet server trusts. Fleet rejects a self-signed certificate, and with Option 2 the activities are lost. For a certificate from a private CA, add the CA to the trust store of the Fleet server; for the Fleet container, mount the CA file and set `SSL_CERT_FILE` to its path.

Create a token for the listener:

```bash
openssl rand -hex 32
```

**Sample FleetDM config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/fleetdm.conf`

For Option 1:
```ini
[HEC-Compatible-Listener "fleetdm"]
    URL="/fleetdm/services/collector"
    TokenValue="<TOKEN>"
    Tag-Name=fleetdm
    Preprocessor=fleetdm-created-at

[Preprocessor "fleetdm-created-at"]
    Type=regextimestamp
    Regex=`created_at\W+(?P<ts>[0-9T:.Z+-]+)`
    TS-Match-Name=ts
    Timestamp-Format-Override=RFC3339Nano
```

* Fleet stamps each event with the time it sends it, up to 5 minutes after the activity. The preprocessor sets the entry time to the activity's own `created_at` instead.
* Keep the backticks around the `Regex` value. In double quotes, `\W` is an invalid escape and the ingester does not start.
* Each activity becomes one entry.

For Option 2:
```ini
[Listener "fleetdm"]
    URL="/fleetdm"
    Tag-Name=fleetdm
    AuthType="preshared-parameter"
    TokenName=token
    TokenValue="<TOKEN>"
```

* Each webhook post is one activity and becomes one entry. Gravwell uses the activity's `timestamp` as the entry time.
* `TokenName=token` matches the `?token=` in the destination URL. Fleet hides parameters with that name when it writes the URL to its own server log.

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_http_ingester.service`
```

## Verification

1. Cause a failed login in Fleet, by entering a wrong password on the Fleet login page or with:
   ```bash
   curl -s -X POST https://<FLEET-HOST>/api/v1/fleet/login \
     -H "Content-Type: application/json" \
     -d '{"email":"gravwell-test@example.invalid","password":"not-the-password"}'
   ```
   Fleet answers `Authentication failed` and records a `user_failed_login` activity.
2. Search for it. With Option 2 it arrives within seconds; with Option 1, within 5 minutes:
   ```text
   tag=fleetdm json type as atype details.email as email details.public_ip as ip
   | eval atype == "user_failed_login"
   | table TIMESTAMP atype email ip
   ```
   `TIMESTAMP` is the time of the failed login on both options, not the time the entry arrived.
3. See which activity types arrive:
   ```text
   tag=fleetdm json type as atype | stats count by atype
   ```

## Troubleshooting

**Nothing arrives.** Look in the Fleet server log. Option 2 errors start with `fire activity webhook to`; Option 1 errors are on the `activities_streaming` job and mention `splunk`.
* `connections to private network addresses are blocked`: set `FLEET_SERVER_ALLOW_PRIVATE_NETWORK_INTEGRATIONS=true` (see FleetDM Configuration). For a loopback address, use the host's network address.
* `x509: certificate signed by unknown authority`: the Fleet server does not trust the ingester's certificate.
* A `401` status: the token does not match. The HTTP Ingester also logs `access denied` to the `gravwell` tag, with `Bad token value`, `Missing token parameter` or `Unauthorized`.

**Activities are missing (Option 2).** Activities recorded while a delivery failed are lost (see the warning in Option 2). They remain in Fleet's activity feed. Switch to Option 1 if your Fleet qualifies.

**Duplicate activities.** Both options send to the same tag. Disable one.

**Entries that are not activities.** Activities always have a top-level `type`. Entries without one are osquery logs or other webhook payloads sent to the activity listener:
```text
tag=fleetdm json type as atype | eval len(atype) == 0 | count
```
Send those streams to another listener and tag.

## Additional Resources

* [Fleet audit logs reference](https://fleetdm.com/docs/using-fleet/audit-logs)
* [Fleet server configuration](https://fleetdm.com/docs/configuration/fleet-server-configuration)
* [Fleet log destinations](https://fleetdm.com/guides/log-destinations)
* [Fleet REST API: update configuration](https://fleetdm.com/docs/rest-api/rest-api)
