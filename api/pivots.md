# Pivots/Actionables API

Pivots, also called *actionables*, are objects stored in Gravwell which the web GUI uses to "pivot" from search result data. For instance, an actionable could define a set of queries which can be run on an IP address, along with a regular expression which *matches* IP addresses. When the user runs a query that includes IP addresses in the results, those addresses will be clickable, bringing up a menu to launch the pre-defined queries.

## Data Structure

The pivot structure contains the following fields:

* GUID: A global reference for the template. Persists across kit installation. (see next section)
* ThingUUID: A unique ID for this particular template instance. (see next section)
* UID: The numeric ID of the template's owner.
* GIDs: An array of numeric group IDs with which this template is shared.
* Global: A boolean, set to true if template should be visible to all users (admin only).
* Name: The template's name.
* Description: A more detailed description of the template.
* Updated: A timestamp representing the last update time for the template.
* Labels: An array of strings containing [labels](#!gui/labels/labels.md).
* Disabled: A boolean value indicating if the pivot has been disabled.
* Contents: The actual definition of the template itself (see below).

Although the webserver does not care what goes into the `Contents` field (except that it should be valid JSON), there is a particular format which the **GUI** uses. Below is a complete Typescript definition of the actionable structure, including the Contents field and descriptions for the various types used.

```
interface Actionable {
    GUID: UUID;
    ThingUUID: UUID;
    UID: NumericID;
    GIDs: null | Array<NumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Empty string is null
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
    | { type: 'template'; reference: UUID; options?: {} }
    | { type: 'savedQuery'; reference: UUID; options?: {} }
    | { type: 'dashboard'; reference: UUID; options?: { variable?: string } }
    | { type: 'url'; reference: string; options: { modal?: boolean; modalWidth?: string } };
```

## Naming: GUIDs and ThingUUIDs

Pivots/actionables have two different IDs attached to them: a GUID, and a ThingUUID. They are both UUIDs, which can be confusing--why have two identifiers for one object? We will attempt to clarify in this section.

Consider an example: I create the pivot from scratch, so it gets assigned a random GUID, `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`, and a random ThingUUID, `c3b24e1e-5186-4828-82ee-82724a1d4c45`. I then bundle the pivot into a kit. Another user on the same system then installs this kit for themselves, which instantiates a pivot with the **same** GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`) but a **random** ThingUUID (`f07373a8-ea85-415f-8dfd-61f7b9204ae0`).

This system is identical to the one used in [templates](templates.md). Templates use GUIDs and ThingUUIDs so that dashboards can refer to templates by GUID, but multiple users can still install the same kit (with the sample template) at the same time without conflict. Although no Gravwell components reference actionables in the same way dashboards reference templates, we have included the behavior as future-proofing.

### Accessing Pivots via GUID vs ThingUUID

Regular users must always access pivots by GUID. Admin users may refer to a pivot by ThingUUID instead, but the `?admin=true` parameter must be set in the request URL.

## Create a pivot

To create a pivot, issue a POST to `/api/pivots`. The body should be a JSON structure with a 'Contents' field, and optionally a GUID, Labels, Name, and Description. For example:

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

The API will respond with the GUID of the newly-created pivot. If a GUID is specified in the request, that GUID will be used. If no GUID is specified, a random GUID will be generated.

Note: At this time, the `UID`, `GIDs`, and `Global` fields cannot be set during pivot creation. They must instead be set via an update call (see below).

## List pivots

To list all pivots available to a user, do a GET on `/api/pivots`. The result will be an array of pivots:

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

## Fetch a single pivot

To fetch a single pivot, issue a GET request to `/api/pivots/<guid>`. The server will respond with the contents of that pivot, for instance a GET on `/api/pivots/afba4f9b-f66a-4f9f-9c58-f45b3db6e474` might return:

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

Note that an administrator can fetch this particular pivot explicitly by using the ThingUUID and the admin parameter, e.g. `/api/pivots/196a3cc3-ec9e-11ea-bfde-7085c2d881ce?admin=true`.

## Update a pivot

To update a pivot, issue a PUT request to `/api/pivots/<guid>`. The request body should be identical to that returned by a GET on the same path, with any desired elements changed. Note that the GUID and ThingUUID cannot be changed; only the following fields may be modified:

* Contents: The actual body/contents of the pivot
* Name: Change the name of the pivot
* Description: Change the pivot's description
* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global pivots are visible to all users.

Note: Leaving any of these field blank will result in the pivot being updated with a null value for that field!

## Delete a pivot

To delete a pivot, issue a DELETE request to `/api/pivots/<guid>`.

## Admin actions

Admin users may occasionally need to view all pivots on the system, modify them, or delete them. Because GUIDs are not necessarily unique, the admin API must refer instead to the unique UUID Gravwell uses internally to store the items. Note that the example pivot listings above include a field named "ThingUUID". This is the internal, unique identifier for that pivot.

An administrator user may obtain a global listing of all pivots in the system with a GET request on `/api/pivots?admin=true`.

The administrator may then update a particular pivot with a PUT to `/api/pivots/<ThingUUID>?admin=true`, substituting in the ThingUUID value for the desired pivot. The same pattern applies to deletion.

An administrator may access or delete a particular pivot with a GET or DELETE request (respectively) on `/api/pivots/<ThingUUID>?admin=true`.