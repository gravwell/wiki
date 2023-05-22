# Capability Based Access Control

Capability Based Access Control (CBAC) is a feature permissions system that enables users and groups to be configured with fine-grained access to various Gravwell features. For example, using CBAC, a user can be configured to have access to search, but not resources or kits. Additionally, CBAC can be used to define which tags are available to users and groups.

CBAC is based around a deny-all default policy. Capabilities and tag access must be granted to each user (or group a user belongs to) in order to access those features. Admin users are not restricted by CBAC and always have full system access.

## Enabling CBAC

CBAC is enabled by adding the following clause to the global section of the webserver's `gravwell.conf` and restarting the webserver:

```
Enable-CBAC=true
```

Because CBAC has a deny-all default policy, if this is the first time enabling CBAC, _all_ non-admin users will begin with no capabilities or tag access. 

## Granting Capabilities to Users and Groups

Admin users can grant capability and tag access to both users and groups via the Users and Groups view under Administrator->Users and Administrator->Groups.

![CBAC User Editor](cbac_user.png)

When creating or editing a user or group, select the "Capabilities" tab and select the capabilities you wish to add. Users that are part of a group will also inherit capabilities from that group.

Tag access is configured by selecting the Tags tab, and selecting the tags the user or group has access to.

Users that don't have access to a particilar feature will see a menu system with those features disabled. For example, a user that does not have access to dashboard or data ingest will see a menu system like this:

![CBAC Menu](cbac_menu.png)

## Granting Capabilities in Practice

A typical use of user and group CBAC grants is to not provide any grants to individual users, and instead create groups with specific roles. For example, creating a group named "IT Users" that has access to IT related data (syslog, router logs, firewall logs, etc.), and a group named "Incident Response Users" that has access to IDS and other security realted data, allows the admin to grant access to users based on their role. Users that need access to both IT and Incident response data in this example can simply be added to both groups.

## List of CBAC Capabilities

| Capability Name | Desecription |
|--------|-------|
| Search | Search data and execute queries. |
| Download | Download search results. |
| SaveSearch | Save a search and add notes. |
| AttachSearch | Load a search by search ID. |
| BackgroundSearch | Execute a search in the background. |
| GetTags | View tags in the system. |
| SetSearchGroup | Assign a default group to searches. |
| SearchHistory | View search history of authenticated user |
| SearchAllHistory | View search history of items user has access to. |
| SearchGroupHistory | View a group search history |
| DashboardRead | View and launch dashboard |
| SearchHistory | View search history of the authenticated user. |
| SearchAllHistory | View search history of items the user has access to. |
| SearchGroupHistory | View group search history. |
| DashboardRead | View and launch a dashboard. |
| DashboardWrite | Create and edit a dashboard's searches and settings. |
| ResourceRead | View a resource and use it in a query. |
| ResourceWrite | Create and edit a resource. |
| TemplateRead | View and execute a template. |
| TemplateWrite | Create and edit a template. |
| PivotRead | View and click on an actionable. |
| PivotWrite | Create and edit an actionable. |
| MacroRead | View and use a macro in a query. |
| MacroRead | View a macro and use it in a query. |
| MacroWrite | Create and edit a macro. |
| LibraryRead | View and execute a saved query. |
| LibraryWrite | Create and edit a saved query. |
| ExtractorRead | View and use an extractor in a query. |
| ExtractorRead | View an extractor and use it in a query. |
| ExtractorWrite | Create and edit an extractor. |
| UserFileRead | View a file. |
| UserFileWrite | Create and edit a file. |
| KitRead | View a kit and its contents. |
| KitWrite | Create and edit a kit. |
| KitBuild | Build a kit. |
| KitDownload | Download a kit. |
| ScheduleRead | View a flow, script or scheduled search and its results. |
| ScheduleWrite | Create and edit a flow, script or scheduled search. |
| ScheduleRead | View a flow, script, or scheduled search and its results. |
| ScheduleWrite | Create and edit a flow, script, or scheduled search. |
| SOARLibs | Import an external library into a script. |
| SOAREmail | Send an email in a script. |
| PlaybookRead | View a playbook. |
| PlaybookWrite | Create and edit a playbook. |
| LicenseRead | View the license. |
| Stats | View health statistics. |
| Ingest | Ingest data. |
| ListUsers | View the list of users. |
| ListGroups | View the list of groups. |
| ListGroupMembers | View the members of a group. |
| NotificationRead | View notifications. |
| NotificationWrite | Create and edit notifications. |
| SystemInfoRead | View systems. |
| SystemInfoRead | View systems info. |
| TokenRead | View API tokens. |
| TokenWrite | Create and edit an API token. |
| SecretRead | User can read and access secrets. |
| SecretWrite | User can create, update, and delete secrets. |
| SecretRead | Read and access secrets. |
| SecretWrite | Create, update, and delete secrets. |

## Determining a CBAC Grant

A user is given a capability or tag grant from both the grants given to the user directly, and from any groups the user is a part of. 

For example, if user "Bob" has access to Search and Resources (but nothing else), and the "gravwell" tag, and is part of a group that grants access to "Dashboards" and the "default" tag, "Bob" has access to Search, Resources, and Dashboards, and both the "gravwell" and "default" tags.

## CBAC Restrictions in Search

CBAC capabilities also apply to search. A user that does not have access to resources will still be able to invoke the `lookup` module (and other resource-based modules), but that module will list no resources as being available. Similarly, macros, auto extractors, and related features in the `anko` and legacy `eval` modules will be restricted based on the user's CBAC grants.

## Caveats 

- Some capabilities require both read and write grants in order to function correctly, such as Resources and Playbooks.
- The GUI automatically selects read access for a given feature if any of the write capabilities for that feature are selected. If write-only access is needed, you must use the Gravwell command line tool to configure the capability.

