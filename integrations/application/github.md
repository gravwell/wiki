# GitHub

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, • [HTTP - HEC](/ingesters/http.md#splunk-hec-compatibility) <br /> • [Simple Relay](/ingesters/simple_relay.md)
         Kit, [GitHub Kit](https://github.com/gravwell/kits/tree/main/)
        test, [test]({ref}splunk-hec-compatibility)
:::

## GitHub Configuration

### [Option 1] Streaming Github Logs
* [Streaming the audit log for your enterprise](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-splunk)

Follow the instructions for setting up streaming to Splunk. For the configuration page point to your Gravwell HTTP ingester.

### [Option 2] Using WebHooks to export Logs
* [Creating a repository webhook](https://docs.github.com/en/enterprise-cloud@latest/webhooks/using-webhooks/creating-webhooks)

Github provides webhooks for exporting logs depending on what you want to export for example for monitoring single repository, app, enterprise, global, etc.. You can follow the instructions posted point the *Payload URL* to a Gravwell Simple Relay Ingester. Change *Content Type* to *application/json*


## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/github-well.conf`
```ini
[Storage-Well "github"]
    Location=/opt/gravwell/storage/github
    Tags=github*
```
### Gravwell HTTP HEC Ingester Configuration

Setup the HTTP HEC configuration file.

**Sample GitHub config:**  
Create or edit: `/opt/gravwell/etc/gravwell_http_ingester.conf.d/github.conf`
```ini
[HEC-Compatible-Listener "github"]
    URL="/services/collector"
    TokenValue="thisisyourtoken"
    Tag-Match=github:github
    Tag-Match=github-audit:github_audit
```
