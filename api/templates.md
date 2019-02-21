# Templates API

This API creates and manages templates

## Create templates

To create a template, do a POST to `/api/templates`. The body should contain valid JSON. The API will respond with the UUID of the newly-created template

## List templates

To list all templates available to a user, do a GET on `/api/templates`. The result will be an array of templates:

```
[{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":null,"Contents":null,"Updated":"2019-02-21T23:18:24.565968066Z","Synced":true}]
```

This example shows only one template. The 'Synced' field can be safely ignored. Note that when listing templates, the 'Contents' field is set to null to avoid shipping excessive data. To access the contents of a template, use the API to fetch a single template.

## Fetch a template's contents

To fetch the contents of a single template, do a GET on `/api/templates/<uuid>`.

## Update a template's contents

To update the contents of a template, do a PUT to `/api/templates/<uuid>`. The request body should contain valid JSON.

## Set group access to a template

To change which groups can access a template, do a PUT on `/api/templates/<uuid>/groups`. The request body should be a JSON-formatted array of positive integers, e.g.: `[1, 8]`.

## Get details about a template

To get information about a particular template, such as the owner's UID, do a GET on `/api/templates/<uuid>/details`. The result will look like this:

```
{"UUID":"f0b359d4-362e-11e9-8417-0242ac11000a","UID":1,"GIDs":null,"Contents":null,"Updated":"2019-02-21T23:18:24.565968066Z","Synced":true}
```

## Delete a template

To delete a template, send a DELETE request to `/api/templates/<uuid>`.
