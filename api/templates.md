# Templates and Pivots API

Templates and pivots have identical APIs. This document describes the template API; to manage pivots, simply replace `templates` in the URLs with `pivots`.

Templates and pivots are referred to by GUID. Note that if a user installs a template or pivot with the same GUID as a global template, theirs will transparently override the global one, but only for themselves. If multiple templates or pivots exist with the same GUID, they are prioritized in the following order:

* Owned by the user
* Shared with a group the user is a member of
* Global

## Create a template

To create a template, do a POST to `/api/templates`. The body should be a JSON structure with a 'Contents' field containing any valid JSON, and optionally a GUID, Name, and Descrption. The following are all valid:

```
{"Contents": "foo"}
```

```
{"GUID":"ce95b152-d47f-443f-884b-e0b506a215be","Contents":{"a":1, "b":2}}
```

```
{"Contents": "foo","Name":"mytemplate"}
```

```
{"GUID":"ce95b152-d47f-443f-884b-e0b506a215be","Contents": "foo","Name":"mytemplate"}
```
The API will respond with the GUID of the newly-created template

## List templates

To list all templates available to a user, do a GET on `/api/templates`. The result will be an array of templates:

[{"GUID":"a8fbdee6-9d92-4d5e-80ab-540532babd54","ThingUUID":"7a4de770-6c31-11e9-b1ef-54e1ad7c66cf","UID":1,"GIDs":[2,3],"Global":false,"Name":"blah","Description":"","Updated":"2019-03-29T10:55:40.127258032-06:00","Contents":"bar"}]

## Fetch a single template

To fetch a single template, do a GET on `/api/templates/<guid>`:

{"GUID":"a8fbdee6-9d92-4d5e-80ab-540532babd54","ThingUUID":"7a4de770-6c31-11e9-b1ef-54e1ad7c66cf","UID":1,"GIDs":[2,3],"Global":false,"Name":"blah","Description":"","Updated":"2019-03-29T10:55:40.127258032-06:00","Contents":"bar"}

## Update a template

To update a template, do a PUT to `/api/templates/<guid>`. The request body should be identical to that returned by a GET on the same path, with any desired elements changed. Note that the GUID cannot be changed; only the following fields may be modified:

* Contents: The actual body/contents of the template
* Name: Change the name of the template
* Description: Change the template's description
* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global templates are visible to all users.

Note: Leaving any of these field blank will result in the template being updated with a null value for that field!

## Delete a template

To delete a template, send a DELETE request to `/api/templates/<guid>`.

## Admin actions

Admin users may occasionally need to view all templates/pivots on the system, modify them, or delete them. Because GUIDs are not necessarily unique, the admin API must refer instead to the unique UUID Gravwell uses internally to store the items. Note that the example template listings above include a field named "ThingUUID". This is the internal, unique identifier for that template.

An administrator user may obtain a global listing of all templates in the system with a GET request on `/api/templates?admin=true`.

The administrator may then update a particular template with a PUT to `/api/templates/<ThingUUID>?admin=true`, substituting in the ThingUUID value for the desired template. The same pattern applies to deletion.