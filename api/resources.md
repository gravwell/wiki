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

A newly-created resource contains no data. To modify the contents of a resource, issue a multipart PUT request to `/api/resources/{guid}/raw`, replacing `{guid}` with the appropriate GUID of the resource. The request only needs one part, named `file`, containing the data which should be stored in the resource. Thus, to set the contents of the resource created above, perform a multipart PUT on `/api/resources/2332866c-9b8d-469f-bf40-de9fad828362/raw`. The server will respond with an updated metastructure showing the new modification time, size, and hash. An example curl invocation is shown below, uploading the file named "maxmind.db" to the resource (note that the Bearer token will need to be set appropriately for your user session, this is simply an example):

```
curl 'http://gravwell.example.com/api/resources/2332866c-9b8d-469f-bf40-de9fad828362/raw' -X PUT -H 'Authorization: Bearer 7b22616c676f223a35323733382c22747970223a226a3774227d.7b22756964223a312c2265787069726573223a22323031392d31302d30395431333a33333a32352e3231343632203131352d30363a3030222c22696174223a5b33392c32323c2c35382c36362c3231372c32362c3131392c33362c3234312c33352c39302c312c39312c3138312c3234322c33362c3137342c3139342c3130382c37342c3133382c32362c3133392c3234362c37362c3132352c3136342c38382c39322c39302c3231312c36365d7d.ef9ca1e0ac7f012adcd796d8cca0746a6fabecd7e787c025d754e54a072be5c89dc7bac5f648ae26b422f0bbe6b69a806e8de4a0fe2b7d06d3293ed4c1323daf' -H 'Content-Type: multipart/form-data' -H 'Accept: */*' --form file=@maxmind.db
```

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

## Getting resource content-type

A GET request to `/api/resources/{guid}/contenttype` will return a structure containing the detected content-type of the resource and the first 512 bytes of the resource body:

```
{"ContentType":"text/plain; charset=utf-8","Body":"IyBEdW1wcyB0aGUgcm93cyBvZiB0aGUgc3BlY2lmaWVkIENTViByZXNvdXJjZSBhcyBlbnRyaWVzLCB3aXRoCiMgZW51bWVyYXRlZCB2YWx1ZXMgY29udGFpbmluZyB0aGUgY29sdW1ucy4KIyBlLmcuIGdpdmVuIGEgcmVzb3VyY2UgbmFtZWQgImZvbyIgY29udGFpbmluZyB0aGUgZm9sbG93aW5nOgojCWhvc3RuYW1lLGRlcHQKIwl3czEsc2FsZXMKIwl3czIsbWFya2V0aW5nCiMJbWFpbHNlcnZlcjEsSVQKIyBydW5uaW5nIHRoZSBmb2xsb3dpbmcgcXVlcnk6CiMJdGFnPWRlZmF1bHQgYW5rbyBkdW1wIGZvbwojIHdpbGwgb3V0cHV0IDQgZW50cmllcyB3aXRoIHRoZSB0YWcgImRlZmF1bHQiLCBjb250YWluaW5nIGVudW1lcmF0ZWQKIyB2YWx1ZXMgbmFtZWQgImhvc3RuYW1lIiBhbmQgImRlcHQiIHdob3NlIGNvbnRlbnRzIG1hdGNoIHRoZSByb3dzCiMgb2YgdGhlIHJlc291cmNlLgojIEZsYWdzOgojICAtZDogc3BlY2lmaWVzIHRoYXQgaW5jb21pbmcgZW50cmllcyBzaG91bGQgYmUgZHJvcHBlZCA="}
```

Adding the bytes parameter, e.g. `/api/resources/{guid}/contenttype?bytes=1024` will change the number of bytes returned. Note that for successful content-type detection, the API will always read at least 128 bytes; if fewer than 128 bytes are specified, the API will default to reading 512 bytes.

## Deleting resources

To delete a resource, simply issue a DELETE request on `/api/resources/{guid}`, replacing `{guid}` with the appropriate GUID of the resource as usual.

## Admin actions

Admin users may occasionally need to view all resources on the system. An administrator user may obtain a global listing of all resources in the system with a GET request on `/api/resources?admin=true`.

Because resource GUIDs are unique across the system, the administrator may then modify/delete/retrieve any resource without the need to specify `?admin=true`, although adding the parameter unecessarily will not cause an error.