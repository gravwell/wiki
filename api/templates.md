# Templates and Pivots API

Templates and pivots have identical APIs. This document describes the template API; to manage pivots, simply replace `templates` in the URLs with `pivots`.

## Create templates

To create a template, do a POST to `/api/templates`. The body should contain valid JSON. The API will respond with the UUID of the newly-created template

## List templates

To list all templates available to a user, do a GET on `/api/templates`. The result will be an array of templates:

```
[{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":null,"Global":false,"Contents":null,"Updated":"2019-02-21T23:18:24.565968066Z","Synced":true}]
```

This example shows only one template. The 'Synced' field can be safely ignored. Note that when listing templates, the 'Contents' field is set to null to avoid shipping excessive data. To access the contents of a template, use the API to fetch a single template.

## Fetch a template's contents

To fetch the contents of a single template, do a GET on `/api/templates/<uuid>`.

## Update a template's contents

To update the contents of a template, do a PUT to `/api/templates/<uuid>`. The request body should contain valid JSON.

## Get details about a template

To get information about a particular template, such as the owner's UID, do a GET on `/api/templates/<uuid>/details`. The result will look like this:

```
{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":null,"Global":false,"Contents":null,"Updated":"2019-02-21T23:18:24.565968066Z","Synced":true}
```

## Set group access, global flag, and UID of a template

To change access rules on a template, do a PUT on `/api/templates/<uuid>/details`. The request body should be the same structure returned by a GET on that path, with the desired fields modified. The following fields can be set; all others will be ignored:

* GIDs: May be set to an array of 32-bit integer group IDs, e.g. `"GIDs":[1,4]`
* UID: (Admin only) Set to a 32-bit integer
* Global: (Admin only) Set to a boolean true or false; Global templates are visible to all users.

If the original details are the following:

```
{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":[2],"Global":false,"Contents":null,"Updated":"2019-02-21T23:18:24.565968066Z","Synced":true}
```

We can set the Global flag true by sending the following; note that the UID and GID fields must remain the same, although other fields can be omitted:

```
{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":[2],"Global":true,"}
```

## Delete a template

To delete a template, send a DELETE request to `/api/templates/<uuid>`.
