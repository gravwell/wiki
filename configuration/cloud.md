# Cloud Customers

Gravwell offers a hosted cloud solution for customers who do not want to run Gravwell locally. This page documents common questions, configurations, and problems often encountered in the hosted solution.

## Email

Gravwell's customer success team can configure your hosted instance to send email using your email provider. We'll need the following information:

* Email server address & port.
* A username & password to authenticate.

There are some parameters which can be tweaked to customize your email configuration for your needs:

* Email sending can be limited to specific users or groups.
* The "From" field on outgoing emails can be overridden.
* Destination addresses can be restricted via regular expressions.

If you want to use any of these options, please let customer success know during the setup process.

```{note}
Gravwell works with most email servers, but please coordinate with your email server operator to ensure the Gravwell cloud instance is whitelisted etc.
```

## User Interface Differences

When using a hosted cloud deployment, a few things will look different in the web UI:

* The "Storage, Indexers & Wells" section of the Systems & Health UI will only show a single indexer named "indexers". This virtual indexer represents all the physical indexers in the system.
* The "Ingesters & Federators" section will coalesce ingester information across all indexers into a single virtual indexer, as above.
* The "Topology" section will not show anything.

## Network Considerations

For better network availability, Gravwell uses load balancers and multiple ISPs with failover. When your deployment is prepared, Gravwell CS will provide you a list of DNS names for the various components of the new system. Always use the DNS names when configuring integrations, never hard-code IP addresses -- this can interfere with redundancy and make your system *less* reliable.

Hosted cloud deployments are always IPv6-enabled, and we recommend taking advantage of it when possible.

## Hosted Ingesters

Gravwell will deploy two types of ingesters in the cloud environment for your use:

* The [HTTP Ingester](/ingesters/http), with authentication enabled.
* The [Federator](/ingesters/federator).

Many devices and software packages can forward logs to a Splunk-compatible HTTP endpoint, meaning you can integrate directly with the hosted HTTP ingester.

For other data sources, such as syslog, you'll need to deploy an appropriate [ingester](/ingesters/ingesters) locally and configure it to send its data to the hosted [Federator](/ingesters/federator).
