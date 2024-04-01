# Automation

Gravwell provides several utilities to enable automated operations. At the most basic level, users can schedule searches to be executed at specific times, for example every morning at 1:00. They can also schedule flows and scripts, which can execute multiple searches, sift and parse search results, and send notifications via email or HTTP. Flows are designed using a drag-and-drop graphical interface and are generally preferred over scripts, which should be reserved for legacy situations or very particular use cases that flows cannot cover.

```{attention}
Automations run as the user who owns the automation. This includes flows, scheduled searches, and scripts. When you share an automation with another user or group, you are granting them access to modify and execute the automation as your user. Consider using a machine user with the least privileges necessary as the owner of shared automations. 
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
The Anko Module <scripting/anko>
The Eval Module <scripting/eval>
```

```{toctree}
---
maxdepth: 1
caption: Alerts
---
Alerts <alerts/alerts>
```
