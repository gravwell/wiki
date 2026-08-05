---
myst:
  substitutions:
    package: "gravwell-hosted-runner"
    standalone: "gravwell_hosted_runner"
    dockername: "hosted_runner"
---
# Jamf Ingester

The [Jamf](https://www.jamf.com/) ingester polls the Jamf Pro computer-inventory API v3 endpoint. By default, it ingests data from the GENERAL section of the computer-inventory endpoint, but other sections listed in the "section" query parameter list [here](https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory) can be added. Timestamps are preserved from the original events to maintain accuracy even across polling gaps or downtime. This ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration).

## Installation

```{include} installation_instructions_template  
```

If you already have the hosted runner installed, you can modify the config.

## Configuration

To configure the ingester you will need the following from Jamf:

* **Client ID**: The OAuth 2.0 client ID for your API 2.0 integration  
* **Client Secret**: The OAuth 2.0 client secret for your API 2.0 integration

See the [Jamf documentation](https://learn.jamf.com/r/en-US/jamf-pro-documentation-current/API_Roles_and_Clients) for instructions on creating an API 2.0 integration and obtaining these credentials.

The Jamf ingester is configured via `[Jamf "name"]` stanzas in the Hosted Runner configuration file, typically `/opt/gravwell/etc/hosted_runner.conf`. The `[Global]` and `[State]` blocks common to all Hosted Runner plugins are described in [Hosted Runner Configuration](hosted_runner_configuration).

### Jamf Stanza Parameters

Each `[Jamf "name"]` stanza configures an independent polling connection to the Jamf Pro API. Currently, only a single stanza is needed per Jamf Pro API instance and ingests to a single tag.

| Config Parameter | Type | Required | Default Value | Description |
| --- | --- | --- | --- | --- |
| Ingester-UUID | UUID | yes |     | A unique UUID for this ingester instance. Used for state tracking. |
| Client-Id | string | yes |     | OAuth 2.0 client ID from your Jamf Pro API 2.0 integration. |
| Client-Secret | string | yes |     | OAuth 2.0 client secret from your Jamf Pro API 2.0 integration. |
| Host | URL | yes |     | The Jamf Pro API base URL. |
| Lookback | integer | no  | 1 (hours) | How far back in time to fetch events on first run in hours. |
| Tag-Name | string | no  | jamf | Tag to assign ingested entries. Only valid when a single `Api` is configured. Cannot be used with `Tag-Prefix`. |
| Page-Size | integer | no | 100 | Maximum number of objects per page. |
| Requests-Per-Minute | integer | no  | 5   | Maximum number of API requests per minute. |
| Request-Interval | integer (seconds) | no  | 300 (seconds) | How often to poll the API for new events in seconds |

### Available Sections

The following values can be added with the `Sections` parameter:

| Sections | Description|
| --- | --- |
| `APPLICATIONS` | Information about installed applications on a computer e.g. title, version |
| `ATTACHMENTS` | Upload and delete attachments to the inventory record using this category |
| `CERTIFICATES` | List of certificates installed on a device e.g. issuer, name |
| `CONFIGURATION_PROFILES` | Information about the configuration profiles installed on a mobile device e.g. name, identifier |
| `CONTENT_CACHING` | Information collected by the `ContentCachingInformation` MDM command e.g. alerts, registrationStatus |
| `DISK_ENCRYPTION` | Disk encryption information for partitions on a computer e.g. name, fileVault2Enabled |
| `EXTENSION_ATTRIBUTES` | List of custom data fields collected using extension attributes |
| `GROUP_MEMBERSHIPS` | Information about local, managed groups and membership e.g. groupId, groupName |
| `HARDWARE` | Hardware details for a computer e.g. Make, Model |
| `IBEACONS` | Information for a computer or mobile device iBeacon (Apple’s iBeacon technology) region |
| `LICENSED_SOFTWARE` | Information about licensed software managed by Jamf Pro e.g. appName, version  |
| `LOCAL_USER_ACCOUNTS` | Information about managed local administrator accounts, as well as other local user accounts on a computer e.g. uid, username |
| `OPERATING_SYSTEM` | Operating system details for a computer e.g. operatingSystem, operatingSystemVersion |
| `PACKAGE_RECEIPTS` | Information about the packages installed on a computer e.g. cachedPackages, "Packages installed by Jamf Pro" |
| `PRINTERS` | Information about printer profiles present on a computer e.g. name, type |
| `PURCHASING` | Purchasing information from Apple’s Global Service Exchange (GSX) e.g. "P.O. Number", vendor |
| `SECURITY` | View information from security related categories e.g. "System Integrity Protection", "Gatekeeper" |
| `SERVICES` | Information about active services on a computer e.g. name |
| `SOFTWARE_UPDATES` | Information about available software updates e.g. name, version |
| `STORAGE` | View information from storage related categories e.g. "S.M.A.R.T. Status", serialNumber |
| `USER_AND_LOCATION` | Displays user/location inventory attributes; populated automatically by assigning a user to a computer e.g. "Full Name", "Email address" |

## Example Configuration

The following example shows a default Jamf stanza:

```  
[Jamf "yourserver"]
Ingester-UUID="99100000-0000-0000-0000-000000000000" #example UUID, remember to set
Host=https://yourserver.jamfcloud.com
Client-Id="api-client-id"
Client-Secret="api-client-secret"
```

To ingest additional `Sections`:

```  
[Jamf "yourserver"]
Ingester-UUID="99100000-0000-0000-0000-000000000000" #example UUID, remember to set
Host=https://yourserver.jamfcloud.com
Client-Id="api-client-id"
Client-Secret="api-client-secret"
Sections=OPERATING_SYSTEM
Sections=SERVICES
```

## Additional Resources

* [Jamf Pro API Overview](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-overview)
* [Jamf Pro API Reference](https://developer.jamf.com/jamf-pro/reference/jamf-pro-api)
* [Jamf Pro computer-inventory Reference](https://learn.jamf.com/r/en-US/jamf-pro-documentation-current/Computer_Inventory_and_Criteria_Reference)
* [Jamf Pro API computer-inventory Reference](https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory)
*
