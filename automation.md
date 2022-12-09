# Automation

Gravwell provides several utilities to enable automated operations. At the most basic level, users can schedule searches to be executed at specific times, for example every morning at 1:00. They can also schedule flows and scripts, which can execute multiple searches, sift and parse search results, and send notifications via email or HTTP. Flows are designed using a drag-and-drop graphical interface and are generally preferred over scripts, which should be reserved for legacy situations or very particular use cases that flows cannot cover.

## Gravwell Automation

```{toctree}
---
maxdepth: 1
caption: Flows
---
Flows <flows/flows>
The Flow Editor <flows/editor>
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
