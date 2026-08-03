## \---

myst:  
substitutions:  
package: "gravwell-hosted-runner"  
standalone: "gravwell_hosted_runner"  
dockername: "hosted_runner"

# Jamf Ingester

The \[Jamf\](https://www.jamf.com/) ingester polls the Jamf Pro computer-inventory API v3 endpoint. By default, it ingests data from the GENERAL section of the computer-inventory endpoint, but other sections listed in the "section" query parameter list \[here\](https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory) can be added. Timestamps are preserved from the original events to maintain accuracy even across polling gaps or downtime. This ingester runs as a plugin inside the \[Gravwell Hosted Runner\](hosted_runner_configuration).

## Installation

\`\`\`{include} installation_instructions_template  
\`\`\`

If you already have the hosted runner installed, you can modify the config.

## Configuration

To configure the ingester you will need the following from Jamf:

\* \*\*Client ID\*\*: The OAuth 2.0 client ID for your API 2.0 integration  
\* \*\*Client Secret\*\*: The OAuth 2.0 client secret for your API 2.0 integration

See the \[Jamf documentation\](https://learn.jamf.com/r/en-US/jamf-pro-documentation-current/API_Roles_and_Clients) for instructions on creating an API 2.0 integration and obtaining these credentials.

The Jamf ingester is configured via \`\[Jamf "name"\]\` stanzas in the Hosted Runner configuration file, typically \`/opt/gravwell/etc/hosted_runner.conf\`. The \`\[Global\]\` and \`\[State\]\` blocks common to all Hosted Runner plugins are described in \[Hosted Runner Configuration\](hosted_runner_configuration).

### Jamf Stanza Parameters

Each \`\[Jamf "name"\]\` stanza configures an independent polling connection to the Jamf Pro API. Currently, only a single stanza is needed per Jamf Pro API instance and ingests to a single tag.

| Config Parameter | Type | Required | Default Value | Description |
| --- | --- | --- | --- | --- |
| Ingester-UUID | UUID | yes |     | A unique UUID for this ingester instance. Used for state tracking. |
| Client-Id | string | yes |     | OAuth 2.0 client ID from your Jamf Pro API 2.0 integration. |
| Client-Secret | string | yes |     | OAuth 2.0 client secret from your Jamf Pro API 2.0 integration. |
| Host | URL | yes |     | The Jamf Pro API base URL. |
| Lookback | integer | no  | 1 (hours) | How far back in time to fetch events on first run in hours. |
| Tag-Name | string | no  | jamf | Tag to assign ingested entries. Only valid when a single \`Api\` is configured. Cannot be used with \`Tag-Prefix\`. |
| Page-Size | integer | no | 100 | Maximum number of objects per page. |
| Requests-Per-Minute | integer | no  | 5   | Maximum number of API requests per minute. |
| Request-Interval | integer (seconds) | no  | 300 (seconds) | How often to poll the API for new events in seconds |

### Available Sections

The following values can be added with the \`Sections\` parameter:

| Sections |
| --- |
| \`DISK_ENCRYPTION\` |
| \`PURCHASING\` |
| \`APPLICATIONS\` |
| \`STORAGE\` |
| \`USER_AND_LOCATION\` |
| \`CONFIGURATION_PROFILES\` |
| \`PRINTERS\` |
| \`SERVICES\` |
| \`HARDWARE\` |
| \`LOCAL_USER_ACCOUNTS\` |
| \`CERTIFICATES\` |
| \`ATTACHMENTS\` |
| \`PACKAGE_RECEIPTS\` |
| \`SECURITY\` |
| \`OPERATING_SYSTEM\` |
| \`LICENSED_SOFTWARE\` |
| \`IBEACONS\` |
| \`SOFTWARE_UPDATES\` |
| \`EXTENSION_ATTRIBUTES\` |
| \`CONTENT_CACHING\` |
| \`GROUP_MEMBERSHIPS\` |

## Example Configuration

The following example shows a default Jamf stanzas:

\`\`\`  
\[Jamf "yourserver"\]
Ingester-UUID="99100000-0000-0000-0000-000000000000" #example UUID, remember to set
Host=https://yourserver.jamfcloud.com
Client-Id="api-client-id"
Client-Secret="api-client-secret"
Sections=GENERAL
\`\`\`

To ingest addition \`Sections\` add them:

\`\`\`  
\[Jamf "yourserver"\]
Ingester-UUID="99100000-0000-0000-0000-000000000000" #example UUID, remember to set
Host=https://yourserver.jamfcloud.com
Client-Id="api-client-id"
Client-Secret="api-client-secret"
Sections=GENERAL
Sections=OPERATING_SYSTEM
Sections=SERVICES
\`\`\`

## Additional Resources

\* \[Jamf Pro API Overview\](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-overview)  
\* \[Jamf Pro API Reference\](https://developer.jamf.com/jamf-pro/reference/jamf-pro-api)  
\* \[Jamf Pro API computer-inventory Reference\](https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory)
