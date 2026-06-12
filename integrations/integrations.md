# Integrations 

This page covers currently supported and developed integrations. If your specific log source is not listed, Gravwell can likely still ingest it with a custom configuration.

To get started with unlisted sources, please refer to our [Generic Integrations](generic_integrations) below to see some examples integrations. Whether you are working with a new source or have suggestions for improving an existing integration, you can connect with us through the following channels:
* **Community:** Join the conversation on [Discord server](https://discord.com/invite/gravwell).
* **Requests & Feedback:** Visit our [GitHub Wiki page](https://github.com/gravwell/wiki) to report issues or suggest improvements.
* **Enterprise:** Contact the *Gravwell Customer Success* team to discuss custom plugin development or professional support.

If you successfully integrate a custom log source, please reach out to us. We would love to review your configuration and potentially incorporate it into our list of supported integrations.


```{toctree}
---
hidden: true
---
Apache <application/apache>
Auditd <host/auditd>
Auth0 <application/auth0>
AWS - CloudTrail <cloud/aws/cloudtrail>
AWS - GuardDuty <cloud/aws/guardduty>
AWS - Kinesis <cloud/aws/kinesis>
Azure <cloud/azure>
Bitwarden <application/bitwarden>
Cisco ASA <network/ciscoasa>
Cisco FTD <network/ciscoftd>
CoreDNS <network/coredns>
Corelight <network/corelight>
Duo <application/duo>
Fluentd <host/fluentd>
Fortinet <network/fortinet>
Github <application/github>
IPFIX <network/ipfix>
IPMI <generic/ipmi>
Juniper <network/juniper>
Netflow <network/netflow>
Nginx <application/nginx>
Office 365 <cloud/office365>
Okta <application/okta>
OpenTelemetry <generic/opentelemetry>
OpenWeatherMap <application/openweathermap>
Palo Alto <network/paloalto>
pfSense <network/pfsense>
PiHole <network/pihole>
Syslog <generic/syslog>
Sysmon <host/sysmon>
Thinkst <network/thinkst>
Windows Event <host/windowsevent>
Zeek <network/zeek>
```

## Cloud
::::{grid} 4
:::{grid-item-card}
:link: cloud/aws/cloudtrail
:link-type: doc
**AWS - CloudTrail**
:::

:::{grid-item-card}
:link: cloud/aws/guardduty
:link-type: doc
**AWS - GuardDuty**
:::

:::{grid-item-card}
:link: cloud/aws/kinesis
:link-type: doc
**AWS - Kinesis**
:::

:::{grid-item-card}
:link: cloud/azure
:link-type: doc
**Azure**
:::
::::

::::{grid} 4
:::{grid-item-card}
:link: cloud/office365
:link-type: doc
**Office 365**
:::
::::

## Network
::::{grid} 4
:::{grid-item-card}
:link: network/ciscoasa
:link-type: doc
**Cisco ASA**
:::

:::{grid-item-card}
:link: network/ciscoftd
:link-type: doc
**Cisco FTD**
:::

:::{grid-item-card}
:link: network/coredns
:link-type: doc
**CoreDNS**
:::

:::{grid-item-card}
:link: network/corelight
:link-type: doc
**Corelight** 
:::
::::

::::{grid} 4
:::{grid-item-card}
:link: network/fortinet
:link-type: doc
**Fortinet** 
:::

:::{grid-item-card}
:link: network/ipfix
:link-type: doc
**IPFIX**
:::

:::{grid-item-card}
:link: network/juniper
:link-type: doc
**Juniper** 
:::

:::{grid-item-card}
:link: network/netflow
:link-type: doc
**Netflow** 
:::
::::

::::{grid} 4
:::{grid-item-card}
:link: network/paloalto
:link-type: doc
**Palo Alto**
:::

:::{grid-item-card}
:link: network/pfsense
:link-type: doc
**pfSense**
:::

:::{grid-item-card}
:link: network/pihole
:link-type: doc
**PiHole**
:::

:::{grid-item-card}
:link: network/thinkst
:link-type: doc
**Thinkst**
:::
::::

::::{grid} 4
:::{grid-item-card}
:link: network/zeek
:link-type: doc
**Zeek**
:::
::::

## Host
::::{grid} 4
:::{grid-item-card}
:link: host/auditd
:link-type: doc
**Auditd**
:::

:::{grid-item-card}
:link: host/fluentd
:link-type: doc
**Fluentd**
:::

:::{grid-item-card}
:link: host/sysmon
:link-type: doc
**Sysmon**
:::

:::{grid-item-card}
:link: host/windowsevent
:link-type: doc
**Windows Event**
:::
::::

## Application
::::{grid} 4
:::{grid-item-card}
:link: application/apache
:link-type: doc
**Apache**
:::

:::{grid-item-card}
:link: application/auth0
:link-type: doc
**Auth0**
:::

:::{grid-item-card}
:link: application/bitwarden
:link-type: doc
**Bitwarden**
:::

:::{grid-item-card}
:link: application/duo
:link-type: doc
**Duo**
:::
::::

::::{grid} 4
:::{grid-item-card}
:link: application/github
:link-type: doc
**Github**
:::

:::{grid-item-card}
:link: application/nginx
:link-type: doc
**Nginx**
:::

:::{grid-item-card}
:link: application/okta
:link-type: doc
**Okta** 
:::

:::{grid-item-card}
:link: application/openweathermap
:link-type: doc
**Open Weather Map** 
:::
::::

(generic_integrations)=
## Generic Integrations
::::{grid} 4
:::{grid-item-card}
:link: generic/ipmi
:link-type: doc
**IPMI**
:::

:::{grid-item-card}
:link: generic/opentelemetry
:link-type: doc
**OpenTelemetry**
:::

:::{grid-item-card}
:link: generic/syslog
:link-type: doc
**Syslog**
:::
::::
