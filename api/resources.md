# Resources Web API

The web API provides access to resources. Resources must be referenced by GUID, to prevent ambiguity.

## Resource metadata structure
The resources system keeps a metadata structure for each resource. The web API uses a JSON-encoded version of this struct to communicate. The fields are largely self-explanatory, but are explained here for the sake of precision:

* UID: the UID of the resource owner
* GUID: the resource's unique identifier
* LastModified: the time when the resource was most recently updated
* VersionNumber: incremented every time the resource contents are changed
* GroupACL: a list of integer group IDs which are allowed to access the resource
* Global: if true, the resource is readable by all users on the system. Only administrators can create global resources
* ResourceName: the resource's name
* Description: a verbose description of the resource
* Size: the size, in bytes, of the resource contents
* Hash: a sha1 hash of the resource contents
* Synced: (internal use only)

## Listing resources

To get a list of all resources, do a GET on `/api/resources`. The result will look like this:

```
[{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:19:10.945117816-07:00","VersionNumber":0,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":0,"Hash":"","Synced":true},{"UID":1,"GUID":"66f7be7d-893b-4dc4-b0ad-3609b348385d","LastModified":"2018-02-12T11:06:44.215431364-07:00","VersionNumber":1,"GroupACL":[1],"Global":false,"ResourceName":"test","Description":"test resource","Size":543,"Hash":"zkTmUEV+AR6JZdqhobIeYw==","Synced":true}]
```

This example shows two resources, "newresource" (GUID 2332866c-9b8d-469f-bf40-de9fad828362) and "test" (GUID 66f7be7d-893b-4dc4-b0ad-3609b348385d)

## Creating resources

To create a resource, perform a POST request on `/api/resources`, sending a JSON structure in the following format:

```
{
	"GroupACL": [3,7],
	"Global": false,
	"ResourceName": "newresource",
	"Description": "Description of the resource"
}
```

Note: the structure is a subset of the metadata structure, containing fields which can be set by the user.

The server will respond with a resource metadata structure for the newly-created resource:

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:19:10.945117816-07:00","VersionNumber":0,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":0,"Hash":"","Synced":false}
```

## Setting resource contents

A newly-created resource contains no data. To modify the contents of a resource, issue a PUT request to `/api/resources/{guid}/raw`, replacing `{guid}` with the appropriate GUID of the resource. Thus, to set the contents of the resource created above, perform a PUT on `/api/resources/2332866c-9b8d-469f-bf40-de9fad828362/raw`. The server will respond with an updated metastructure showing the new modification time, size, and hash.

## Reading resource contents

To read the contents of a resource, simply perform a GET request on `/api/resources/{guid}/raw`, replacing `{guid}` with the appropriate GUID of the resource.

## Reading & updating resource metadata

To read the metadata for a single resource, do a GET request on `/api/resources/{guid}`. For example, a GET on `/api/resources/2332866c-9b8d-469f-bf40-de9fad828362` might yield:

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:29:10.557490321-07:00","VersionNumber":1,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":6,"Hash":"QInZ92Blt3TopFBeBTD0Cw==","Synced":true}
```

Metadata can be modified by performing a PUT request on `/api/resources/{guid}`. The contents of the request should be the structure as read with a GET request, with the desired fields modified. For example, to change the description of "newresource", perform a PUT with the following contents:

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:29:10.557490321-07:00","VersionNumber":1,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"A new description for the resource!","Size":6,"Hash":"QInZ92Blt3TopFBeBTD0Cw==","Synced":true}
```

Note: Only the GroupACL, ResourceName, and Description fields can be modified by regular users. Admin users can also modify the Global field. Any other modified fields will be ignored by the server.

## Deleting resources

To delete a resource, simply issue a DELETE request on `/api/resources/{guid}`, replacing `{guid}` with the appropriate GUID of the resource as usual.

## Admin actions

Admin users may occasionally need to view all resources on the system. An administrator user may obtain a global listing of all resources in the system with a GET request on `/api/resources?admin=true`.

Because resource GUIDs are unique across the system, the administrator may then modify/delete/retrieve any resource without the need to specify `?admin=true`, although adding the parameter unecessarily will not cause an error.