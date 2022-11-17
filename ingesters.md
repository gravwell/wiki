# Ingesters

This section contains more detailed instruction for configuring and running Gravwell ingesters, which gather incoming data, package it into Gravwell entries, and ship it to Gravwell indexers for storage. The ingesters described in these pages are primarily designed to capture *live* data as it is generated; if you have existing data you want to import, check out the [migration documents](migrate/migrate.md).

The Gravwell-created ingesters are released under the BSD open source license and can be found on [Github](https://github.com/gravwell/gravwell/tree/master/ingesters). The ingest API is also open source, so you can create your own ingesters for unique data sources, performing additional normalization or pre-processing, or any other manner of things. The ingest API code [is located here](https://github.com/gravwell/gravwell/tree/master/ingest).

In general, for an ingester to send data to Gravwell, the ingester will need to know the “Ingest Secret” of the Gravwell instance, for authentication. This can be found by viewing the `/opt/gravwell/etc/gravwell.conf` file on the Gravwell server and finding the entry for `Ingest-Auth`. If the ingester is running on the same system as Gravwell itself, the installer will usually be able to detect this value and set it automatically.

The Gravwell GUI has an Ingesters page (under the System menu category) which can be used to easily identify which remote ingesters are actively connected, for how long they have been connected, and how much data they have pushed.

![](remote-ingesters.png)

Attention: The [replication system](/configuration/replication) does not replicate entries larger than 999MB. Larger entries can still be ingested and searched as usual, but they are omitted from replication. This is not a concern for 99.9% of use cases, as all the ingesters detailed in this page tend to create entries no larger than a few kilobytes.

```{toctree}
---
maxdepth: 1
caption: Ingester Configuration
---
Ingesters List & Configuration <ingesters/ingesters.md>
Ingester Preprocessors <ingesters/preprocessors/preprocessors.md>
Custom Time Formats <ingesters/customtime/customtime.md>
Service Integrations <ingesters/integrations.md>
Federators <ingesters/federators/federators.md>
```
