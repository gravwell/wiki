# Templates API

Templates are special objects which define a Gravwell query containing *variables*. Multiple templates using the same variable(s) can be included in a dashboard to create a powerful investigative tool--for instance, templates which expect an IP address as their variable can be used to create an IP address investigation dashboard.

## Data Structure

The template structure contains the following fields:

* GUID: A global reference for the template. Persists across kit installation. (see next section)
* ThingUUID: A unique ID for this particular template instance. (see next section)
* UID: The numeric ID of the template's owner.
* GIDs: An array of numeric group IDs with which this template is shared.
* Global: A boolean, set to true if template should be visible to all users (admin only).
* Name: The template's name.
* Description: A more detailed description of the template.
* Updated: A timestamp representing the last update time for the template.
* Labels: An array of strings containing [labels](#!gui/labels/labels.md).
* Contents: The actual definition of the template itself (see below).

Although the webserver does not care what goes into the `Contents` field (except that it must be valid JSON), there is a particular format which the **GUI** uses. The Contents field should conform to this structure in order to be usable for the GUI:

```
Contents: {
  query: string;
  variable: string;
  variableLabel: string;
  variableDescription: string | null;
  required: boolean;
  testValue: string | null;
};
```

The following is a complete Typescript definition of the template data type:

```
interface RawTemplate {
    GUID: RawUUID;
    ThingUUID: RawUUID;
    UID: RawNumericID;
    GIDs: null | Array<RawNumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Empty string is null
    Updated: string; // Timestamp
    Contents: {
        query: string;
        variable: string;
        variableLabel: string;
        variableDescription: string | null;
        required: boolean;
        testValue: string | null;
    };
    Labels: null | Array<string>;
}
```

## Naming: GUIDs and ThingUUIDs

Templates have two different IDs attached to them: a GUID, and a ThingUUID. They are both UUIDs, which can be confusing--why have two identifiers for one object? We will attempt to clarify in this section.

In Gravwell, a dashboard may refer to a particular template. This dashboard & corresponding template may also be packed into a kit for distribution to other users. The dashboard needs a way to refer to the template that will **persist** when packed in a kit and installed elsewhere, so we introduce the GUID as a "global" name for the template: wherever the kit gets installed, that template will have the same GUID. However, multiple users are allowed to install the same kit, so we also need a different identifier for *for each individual instantiation* of the template. This role is filled by the ThingUUID field.

Consider an example: I build a kit which includes a dashboard and a template. I create the template from scratch, so it gets assigned a random GUID, `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`, and a random ThingUUID, `c3b24e1e-5186-4828-82ee-82724a1d4c45`. I then create a tile in the dashboard which refers to the template by its GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`) and bundle both template & dashboard into a kit. Another user on the same system then installs this kit for themselves, which instantiates a template with the **same** GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`) but a **random** ThingUUID (`f07373a8-ea85-415f-8dfd-61f7b9204ae0`). When the user opens the dashboard, the dashboard will ask for a template with GUID == `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`. The webserver will return that user's instance of the template, with ThingUUID == `f07373a8-ea85-415f-8dfd-61f7b9204ae0`.

Note that if a user installs a template with the same GUID as a global template, theirs will transparently override the global one, but only for themselves. If multiple templates exist with the same GUID, they are prioritized in the following order:

* Owned by the user
* Shared with a group the user is a member of
* Global

This means that if a user is accessing a global dashboard, they can override a particular template referred to by the dashboard by creating their own copy of the template with the same GUID. In practice, this should be rare.

### Accessing Templates via GUID vs ThingUUID

Regular users must always access templates by GUID. Admin users may refer to a template by ThingUUID instead, but the `?admin=true` parameter must be set in the request URL.

## Create a template

To create a template, issue a POST to `/api/templates`. The body should be a JSON structure with a 'Contents' field containing any valid JSON, and optionally a GUID, Labels, Name, and Description. The following are all valid:

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  }
}
```

```
{
  "GUID": "ce95b152-d47f-443f-884b-e0b506a215be",
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  }
}
```

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate"
}
```

```
{
  "GUID": "ce95b152-d47f-443f-884b-e0b506a215be",
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate"
}
```

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate",
  "Labels": [
    "suits",
    "ladders"
  ]
}
```

The API will respond with the GUID of the newly-created template. If a GUID is specified in the request, that GUID will be used. If no GUID is specified, a random GUID will be generated.

Note: At this time, the `UID`, `GIDs`, and `Global` fields cannot be set during template creation. They must instead be set via an update call (see below).

## List templates

To list all templates available to a user, do a GET on `/api/templates`. The result will be an array of templates:

```
[
  {
    "ThingUUID": "1b36a1d7-a5ac-11ea-b07e-7085c2d881ce",
    "UID": 1,
    "GIDs": [
      6,
      8
    ],
    "Global": false,
    "GUID": "780b1d31-e46b-4460-ad83-2fc11c34a162",
    "Name": "json ip",
    "Description": "JSON tag, filter by IP",
    "Contents": {
      "variable": "%%IP%%",
      "query": "tag=json* json ip==%%IP%% | table",
      "variableLabel": "IP address",
      "variableDescription": "the IP to investigate!",
      "required": true,
      "testValue": "\"10.0.0.1\""
    },
    "Updated": "2020-09-01T15:01:18.354750806-06:00",
    "Labels": [
      "test"
    ]
  }
]

```

## Fetch a single template

To fetch a single template, issue a GET request to `/api/templates/<guid>`. The server will respond with the contents of that template, for instance a GET on `/api/templates/780b1d31-e46b-4460-ad83-2fc11c34a162` might return:

```
{
  "ThingUUID": "1b36a1d7-a5ac-11ea-b07e-7085c2d881ce",
  "UID": 1,
  "GIDs": [
    6,
    8
  ],
  "Global": false,
  "GUID": "780b1d31-e46b-4460-ad83-2fc11c34a162",
  "Name": "json ip",
  "Description": "JSON tag, filter by IP",
  "Contents": {
    "variable": "%%IP%%",
    "query": "tag=json* json ip==%%IP%% | table",
    "variableLabel": "IP address",
    "variableDescription": "the IP to investigate!",
    "required": true,
    "testValue": "\"10.0.0.1\""
  },
  "Updated": "2020-09-01T15:01:18.354750806-06:00",
  "Labels": [
    "test"
  ]
}
```

Note that an administrator can fetch this particular template explicitly by using the ThingUUID and the admin parameter, e.g. `/api/templates/1b36a1d7-a5ac-11ea-b07e-7085c2d881ce?admin=true`.

## Update a template

To update a template, issue a PUT request to `/api/templates/<guid>`. The request body should be identical to that returned by a GET on the same path, with any desired elements changed. Note that the GUID and ThingUUID cannot be changed; only the following fields may be modified:

* Contents: The actual body/contents of the template
* Name: Change the name of the template
* Description: Change the template's description
* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global templates are visible to all users.

Note: Leaving any of these field blank will result in the template being updated with a null value for that field!

## Delete a template

To delete a template, issue a DELETE request to `/api/templates/<guid>`.

## Admin actions

Admin users may occasionally need to view all templates on the system, modify them, or delete them. Because GUIDs are not necessarily unique, the admin API must refer instead to the unique UUID Gravwell uses internally to store the items. Note that the example template listings above include a field named "ThingUUID". This is the internal, unique identifier for that template.

An administrator user may obtain a global listing of all templates in the system with a GET request on `/api/templates?admin=true`.

The administrator may then update a particular template with a PUT to `/api/templates/<ThingUUID>?admin=true`, substituting in the ThingUUID value for the desired template. The same pattern applies to deletion.

An administrator may access or delete a particular template with a GET or DELETE request (respectively) on `/api/templates/<ThingUUID>?admin=true`.