# Templates and Pivots API

Templates and pivots have identical APIs. This document describes the template API; to manage pivots, simply replace `templates` in the URLs with `pivots`.

Templates and pivots are referred to by GUID. Note that if a user installs a template or pivot with the same GUID as a global template, theirs will transparently override the global one, but only for themselves. If multiple templates or pivots exist with the same GUID, they are prioritized in the following order:

* Owned by the user
* Shared with a group the user is a member of
* Global

## Create templates

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

```
[{"GUID":"218ea16b-d831-48c0-bd1c-50c7b1b6079e","Name":"mytemplate","Description":"a template","Contents":{"test":"blah"}},{"GUID":"a0fe90f6-94ea-4f34-86a0-3fdeaaa11c80","Name":"template2","Description":"","Contents":"bar"}]
```

This example shows two templates.

## Fetch a template's contents

To fetch the contents of a single template, do a GET on `/api/templates/<guid>`.

## Update a template's contents

To update the contents of a template, do a PUT to `/api/templates/<guid>`. The request body should resemble the structure sent to create a new template. Sending the GUID field is not necessary (it will be ignored) but the Name and Description fields should be sent even if not changed:

```
{"Name":"mytemplate","Description":"a template","Contents": "foo"}
```

```
{"GUID":"ce95b152-d47f-443f-884b-e0b506a215be","Name":"mytemplate","Description":"a template","Contents":{"a":1, "b":2}}
```

## Get details about a template

To get information about a particular template, such as the owner's UID, do a GET on `/api/templates/<guid>/details`. The result will look like this:

```
{"GUID":"ce95b152-d47f-443f-884b-e0b506a215bf","UID":1,"GIDs":null,"Global":true,"Name":"mytemplate","Description":"a template","Updated":"2019-02-26T16:46:47.571887406-07:00"}
```

## Set group access, global flag, and UID of a template

To change access rules on a template, do a PUT on `/api/templates/<uuid>/details`. The request body should be the same structure returned by a GET on that path, with the desired fields modified. The following fields can be set; all others will be ignored:

* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global templates are visible to all users.

If the original details are the following:

```
{"GUID":"ce95b152-d47f-443f-884b-e0b506a215bf","UID":1,"GIDs":null,"Global":true,"Name":"mytemplate","Description":"a template","Updated":"2019-02-26T16:46:47.571887406-07:00"}
```

We can set the Global flag true by sending the following; note that the UID and GID fields must remain the same, although other fields can be omitted:

```
{"UID":1,"GIDs":[2],"Global":true,"}
```

## Delete a template

To delete a template, send a DELETE request to `/api/templates/<guid>`.
