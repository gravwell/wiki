# Actionables API

Actionables (previously called "pivots"), are objects stored in Gravwell which the web GUI uses to pivot from search result data. For instance, an actionable could define a set of queries which can be run on an IP address, along with a regular expression which *matches* IP addresses. When the user runs a query that includes IP addresses in the results, those addresses will be clickable, bringing up a menu to launch the pre-defined queries.

## Data Structure

The actionable structure contains the following fields:

* GUID: A global reference for the actionable. Persists across kit installation. (see next section)
* ThingUUID: A unique ID for this particular actionable instance. (see next section)
* UID: The numeric ID of the actionable's owner.
* GIDs: An array of numeric group IDs with which this actionable is shared.
* Global: A boolean, set to true if the actionable should be visible to all users (admin only).
* Name: The actionable's name.
* Description: A more detailed description of the actionable.
* Updated: A timestamp representing the last update time for the actionable.
* Labels: An array of strings containing [labels](/gui/labels/labels).
* Disabled: A boolean value indicating if the actionable has been disabled.
* Contents: The actual definition of the actionable itself (see below).
  * Contents.menuLabel: Optional. If not present, the first 20 characters of the name will be used in the dropdown menu.
  * Contents.actions: Array of actions that can be executed from this actionable.
    * Contents.actions\[n].name: Action name.
    * Contents.actions\[n].description: Optional action description.
    * Contents.actions\[n].placeholder: Placeholder will be replaced with the value of the trigger or cursor highlight. Defaults to "\_VALUE_".
    * Contents.actions\[n].start: Optional ActionableTimeVariable (see interface below) with definitions to handle the start date variable.
    * Contents.actions\[n].end: Optional ActionableTimeVariable (see interface below) with definitions to handle the end date variable.
    * Contents.actions\[n].command: ActionableCommand (see interface below) with definitions for the action execution.
  * Contents.triggers: Array of triggers for this actionable.
    * Contents.triggers\[n].pattern: Serialized regular expression pattern for the actionable to match on.
    * Contents.triggers\[n].hyperlink: True if the actionable can be activated with clicks and text selection. False if it can only be activated with text selection.

Although the webserver does not care what goes into the `Contents` field (except that it should be valid JSON), there is a particular format which the **GUI** uses. Below is a complete Typescript definition of the actionable structure, including the Contents field and descriptions for the various types used.

```
interface Actionable {
    GUID: UUID;
    ThingUUID: UUID;
    UID: NumericID;
    GIDs: null | Array<NumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Could be an empty string
    Updated: string; // Timestamp
    Contents: {
        menuLabel: null | string;
        actions: Array<ActionableAction>;
        triggers: Array<ActionableTrigger>;
    };
    Labels: null | Array<string>;
    Disabled: boolean;
}

type UUID = string;
type NumericID = number;

interface ActionableTrigger {
    pattern: string;
    hyperlink: boolean;
}

interface ActionableAction {
    name: string;
    description: string | null;
    placeholder: string | null;
    start?: ActionableTimeVariable;
    end?: ActionableTimeVariable;
    command: ActionableCommand;
}

type ActionableTimeVariable =
    | { type: 'timestamp'; format: null | string; placeholder: null | string }
    | { type: 'string'; format: null | string; placeholder: null | string };

type ActionableCommand =
    | { type: 'query'; reference: string; options?: {} }
    | { type: 'template'; reference: UUID; options?: { variable?: string } }
    | { type: 'savedQuery'; reference: UUID; options?: {} }
    | { type: 'dashboard'; reference: UUID; options?: { variable?: string } }
    | { type: 'url'; reference: string; options: { modal?: boolean; modalWidth?: string } };
```

## Naming: GUIDs and ThingUUIDs

Actionables have two different IDs attached to them: a GUID, and a ThingUUID. They are both UUIDs, which can be confusing--why have two identifiers for one object? We will attempt to clarify in this section.

Consider an example: I create the actionable from scratch, so it gets assigned a random GUID, `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`, and a random ThingUUID, `c3b24e1e-5186-4828-82ee-82724a1d4c45`. I then bundle the actionable into a kit. Another user on the same system then installs this kit for themselves, which instantiates an actionable with the **same** GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`) but a **random** ThingUUID (`f07373a8-ea85-415f-8dfd-61f7b9204ae0`).

This system is identical to the one used in [templates](templates.md). Templates use GUIDs and ThingUUIDs so that dashboards can refer to templates by GUID, but multiple users can still install the same kit (with the sample template) at the same time without conflict. Although no Gravwell components reference actionables in the same way dashboards reference templates, we have included the behavior as future-proofing.

### Accessing Actionables via GUID vs ThingUUID

Regular users must always access actionables by GUID. Admin users may refer to an actionable by ThingUUID instead, but the `?admin=true` parameter must be set in the request URL.

## Create an actionable

To create an actionable, issue a POST to `/api/pivots`. The body should be a JSON structure with a 'Contents' field, and optionally a GUID, Labels, Name, and Description. For example:

```
{
  "Name": "IP actions",
  "Description": "Actions for an IP address",
  "Contents": {
    "actions": [
      {
        "name": "Whois",
        "description": null,
        "placeholder": null,
        "start": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "end": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "command": {
          "type": "url",
          "reference": "https://www.whois.com/whois/_VALUE_",
          "options": {}
        }
      }
    ],
    "menuLabel": null,
    "triggers": [
      {
        "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
        "hyperlink": true
      }
    ]
  }
}
```

The API will respond with the GUID of the newly-created actionable. If a GUID is specified in the request, that GUID will be used. If no GUID is specified, a random GUID will be generated.

Note: At this time, the `UID`, `GIDs`, and `Global` fields cannot be set during actionable creation. They must instead be set via an update call (see below).

## List actionables

To list all actionables available to a user, do a GET on `/api/pivots`. The result will be an array of actionables:

```
[
  {
    "GUID": "afba4f9b-f66a-4f9f-9c58-f45b3db6e474",
    "ThingUUID": "196a3cc3-ec9e-11ea-bfde-7085c2d881ce",
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "Name": "IP actions",
    "Description": "Actions for an IP address",
    "Updated": "2020-09-01T15:57:23.416537696-06:00",
    "Contents": {
      "actions": [
        {
          "name": "Whois",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "url",
            "reference": "https://www.whois.com/whois/_VALUE_",
            "options": {}
          }
        }
      ],
      "menuLabel": null,
      "triggers": [
        {
          "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
          "hyperlink": true
        }
      ]
    },
    "Labels": null,
    "Disabled": false
  },
  {
    "GUID": "34ba8372-0314-460a-9742-5a65c18d6241",
    "ThingUUID": "e1bdf35a-de7b-11ea-9709-7085c2d881ce",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Network Port",
    "Description": "Actions to take on a network port, e.g. 22",
    "Updated": "2020-08-14T16:17:03.790048874-06:00",
    "Contents": {
      "actions": [
        {
          "name": "Netflow - Most active hosts on this port",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Src Dst SrcPort DstPort Port==_VALUE_ Protocol Bytes | stats sum(Bytes) as ByteTotal by Port Src Dst | lookup -r network_services Protocol proto_number proto_name as Proto Port service_port service_name as Service | table Src Dst Port Service Proto ByteTotal",
            "options": {}
          }
        },
        {
          "name": "Netflow - Chart traffic",
          "description": "Traffic on this port over time",
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Src Dst SrcPort DstPort Port==_VALUE_ Protocol Bytes | lookup -r network_services Protocol proto_number proto_name as Proto Port service_port service_name as Service | stats sum(Bytes) by Service Port | chart sum by Service Port",
            "options": {}
          }
        },
        {
          "name": "Netflow - Internal IPs serving this port",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Dst ~ PRIVATE DstPort==_VALUE_ Bytes Protocol | lookup -r ip_protocols Protocol Number Name as ProtocolName | stats sum(Bytes) as TotalTraffic by Dst | table Dst DstPort Protocol ProtocolName TotalTraffic",
            "options": {}
          }
        }
      ],
      "menuLabel": null,
      "triggers": []
    },
    "Labels": [
      "kit/io.gravwell.netflowv5"
    ],
    "Disabled": false
  }
]
```

## Fetch a single actionable

To fetch a single actionable, issue a GET request to `/api/pivots/<guid>`. The server will respond with the contents of that actionable, for instance a GET on `/api/pivots/afba4f9b-f66a-4f9f-9c58-f45b3db6e474` might return:

```
{
  "GUID": "afba4f9b-f66a-4f9f-9c58-f45b3db6e474",
  "ThingUUID": "196a3cc3-ec9e-11ea-bfde-7085c2d881ce",
  "UID": 1,
  "GIDs": null,
  "Global": false,
  "Name": "IP actions",
  "Description": "Actions for an IP address",
  "Updated": "2020-09-01T15:57:23.416537696-06:00",
  "Contents": {
    "actions": [
      {
        "name": "Whois",
        "description": null,
        "placeholder": null,
        "start": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "end": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "command": {
          "type": "url",
          "reference": "https://www.whois.com/whois/_VALUE_",
          "options": {}
        }
      }
    ],
    "menuLabel": null,
    "triggers": [
      {
        "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
        "hyperlink": true
      }
    ]
  },
  "Labels": null,
  "Disabled": false
}

```

Note that an administrator can fetch this particular actionable explicitly by using the ThingUUID and the admin parameter, e.g. `/api/pivots/196a3cc3-ec9e-11ea-bfde-7085c2d881ce?admin=true`.

## Update an actionable

To update an actionable, issue a PUT request to `/api/pivots/<guid>`. The request body should be identical to that returned by a GET on the same path, with any desired elements changed. Note that the GUID and ThingUUID cannot be changed; only the following fields may be modified:

* Contents: The actual body/contents of the actionable
* Name: Change the name of the actionable
* Description: Change the actionable's description
* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global actionables are visible to all users.

Note: Leaving any of these field blank will result in the actionable being updated with a null value for that field!

## Delete an actionable

To delete an actionable, issue a DELETE request to `/api/pivots/<guid>`.

## Admin actions

Admin users may occasionally need to view all actionables on the system, modify them, or delete them. Because GUIDs are not necessarily unique, the admin API must refer instead to the unique UUID Gravwell uses internally to store the items. Note that the example actionable listings above include a field named "ThingUUID". This is the internal, unique identifier for that actionable.

An administrator user may obtain a global listing of all actionables in the system with a GET request on `/api/pivots?admin=true`.

The administrator may then update a particular actionable with a PUT to `/api/pivots/<ThingUUID>?admin=true`, substituting in the ThingUUID value for the desired actionable. The same pattern applies to deletion.

An administrator may access or delete a particular actionable with a GET or DELETE request (respectively) on `/api/pivots/<ThingUUID>?admin=true`.