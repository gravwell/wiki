# Microsoft Graph API Ingester

Gravwell provides an ingester which can pull security information from Microsoft's Graph API. In order to configure the ingester, you will need to register a new *application* within the Azure Active Directory management portal; this will generate a set of keys which can be used to access the logs. You will need the following information:

* Client ID: A UUID generated for your application via the Azure management console
* Client secret: A secret token generated for your application via the Azure console
* Tenant Domain: The domain of your Azure domain, e.g. "mycorp.onmicrosoft.com"

## Basic Configuration

The MS Graph ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters, the MS Graph ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## ContentType Examples

```
[ContentType "alerts"]
	Content-Type="alerts"
	Tag-Name="graph-alerts"

[ContentType "scores"]
	Content-Type="secureScores"
	Tag-Name="graph-scores"
	Ignore-Timestamps=true

[ContentType "profiles"]
	Content-Type="controlProfiles"
	Tag-Name="graph-profiles"
```

## Installation and configuration

First, download the installer from the [Downloads page](#!quickstart/downloads.md), then install the ingester:

```
root@gravserver ~# bash gravwell_msgraph_installer.sh
```

If the Gravwell services are present on the same machine, the installation script should automatically extract and configure the `Ingest-Auth` parameter and set it appropriately. You will now need to open the `/opt/gravwell/etc/msgraph_ingest.conf` configuration file and set it up for your application, replacing the placeholder fields and modifying tags as desired. Once you have modified the configuration as described below, start the service with the command `systemctl start gravwell_msgraph_ingest.service`.

By default, the ingester will ingest security alerts as they arrive. It will also periodically query for new security score results (typically issued daily), and ingest the associated control profiles which are used to build those security score results. These three data sources are by default ingested to the tags `graph-alerts`, `graph-scores`, and `graph-profiles`, respectively.

The example below shows a sample configuration which connects to an indexer on the local machine (note the `Pipe-Backend-target` setting) and feeds it logs from all supported types:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=ERROR #options are OFF INFO WARN ERROR
State-Store-Location=/opt/gravwell/etc/o365_ingest.state

Client-ID=79fb8690-109f-11ea-a253-2b12a0d35073
Client-Secret="<secret>"
Tenant-Domain=mycorp.onmicrosoft.com

[ContentType "alerts"]
	Content-Type="alerts"
	Tag-Name="graph-alerts"

[ContentType "scores"]
	Content-Type="secureScores"
	Tag-Name="graph-scores"

[ContentType "profiles"]
	Content-Type="controlProfiles"
	Tag-Name="graph-profiles"
```
