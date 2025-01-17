# Automation

Gravwell provides several utilities to enable automated operations. At the most basic level, users can schedule searches to be executed at specific times, for example every morning at 1:00. They can also schedule flows and scripts, which can execute multiple searches, sift and parse search results, and send notifications via email or HTTP. Flows are designed using a drag-and-drop graphical interface and are generally preferred over scripts, which should be reserved for legacy situations or very particular use cases that flows cannot cover.

```{attention}
An automation (Flow, Scheduled Search, or Script) runs as the user who owns the automation, except when triggered by an alert.

Granting write access to an automation has important implications. When you grant a group write access to a Flow or Scheduled Search, you are granting them the ability to modify and execute it as **your user**. Consider using a machine user with the least privileges necessary as the owner of shared automations.

When a flow is run in response to an [alert](/alerts/alerts), it runs as the owner of the alert. This also has implications: the alert owner is executing code defined by the flow owner. Don't use flows owned by untrusted users as consumers on your alerts.
```

## Gravwell Automation

```{toctree}
---
maxdepth: 1
caption: Flows
---
Flows <flows/flows>
The Flow Editor <flows/editor>
Common Flow Patterns <flows/patterns/patterns>
```

```{toctree}
---
maxdepth: 1
caption: Scheduled Searches & Scripts
---
Scheduled Searches & Scripts <scripting/scheduledsearch>
Scripting Overview <scripting/scripting>
```

```{toctree}
---
maxdepth: 1
caption: Alerts
---
Alerts <alerts/alerts>
```
