# Kits Web API

This API implements the creation, installation, and deletion of Gravwell kits. Kits contain other components which are installed on the local system to provide a ready-to-go solution to a particular problem. Kits can contain:

* Resources
* Scheduled searches
* Dashboards
* Auto-extractor definitions
* Templates
* Pivots
* User files

A given kit will also have the following attributes, specified at build time:

* ID: A unique identifier for this kit. We recommend following Android naming practice, e.g. "com.example.my-kit".
* Name: A human-friendly name for the kit, e.g. "My Kit".
* Description: A description of the kit.
* Version: An integer version of the kit.

## Building a kit

Kits are built by sending a POST request to `/api/kit/build` containing a KitBuildRequest structure, as defined below:

```
type KitBuildRequest struct {
	ID                string
	Name              string
	Description       string
	Version           uint
	Dashboards        []uint64    
	Templates         []uuid.UUID 
	Pivots            []uuid.UUID 
	Files             []uuid.UUID 
	Resources         []string    
	ScheduledSearches []int32     
	Macros            []uint64    
	Extractors        []string    
}
```

Note that while the ID, Name, Description, and Version fields are required, the arrays of templates/pivots/dashboards etc. are optional. For example, here is a request to build a kit containing two dashboards, a pivot, a resource, and a scheduled search:

```
{
	"ID": "io.gravwell.test",
	"Name": "test-gravwell",
	"Description": "Test Gravwell kit",
	"Version": 1,
	"Dashboards": [
		7,
		10
	],
	"Pivots": [
		"ae9f2598-598f-4859-a3d4-832a512b6104"
	],
	"Resources": [
		"84270dbd-1905-418e-b756-834c15661a54"
	],
	"ScheduledSearches": [
		1439174790
	]
}
```

Attention: The UUIDs specified for templates, pivots, and userfiles should be the *GUIDs* associated with those structures, not the *ThingUUID* field which is also reported in a listing of items.

The system will respond with a structure describing the newly-built kit:

```
{
	"UUID": "2f5e485a-2739-475b-810d-de4f80ae5f52",
	"Size": 8268288,
	"UID": 1
}
```

This kit can be downloaded by doing a GET on `/api/kit/build/<uuid>`; given the above response, one would fetch the kit from `/api/kit/build/2f5e485a-2739-475b-810d-de4f80ae5f52`

## Uploading a Kit

Before a kit can be installed, it must first be uploaded to the webserver. Kits are uploaded by a POST request to `/api/kit`. The request should contain a multipart form. To upload a file from the local system, add a file field to the form named `file` containing the kit file. To upload a file from a remote system such as an HTTP server, add a field named `remote` containing the URL of the kit.

## Listing Kits

A GET request on `/api/kit` will return a list of all known kits. Here is an example showing the result when the system has one kit uploaded but not yet installed:

```
[
	{
		"UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
		"UID": 7,
		"GID": 0,
		"ID": "io.gravwell.test",
		"Name": "test-gravwell",
		"Description": "Test Gravwell kit",
		"Version": 1,
		"Installed": false,
		"Signed": false,
		"AdminRequired": false,
		"Items": [
			{
				"Name": "84270dbd-1905-418e-b756-834c15661a54",
				"Type": "resource",
				"AdditionalInfo": {
					"VersionNumber": 1,
					"ResourceName": "maxmind_asn",
					"Description": "ASN database",
					"Size": 6196221
				}
			},
			{
				"Name": "55c81086",
				"Type": "scheduled search",
				"AdditionalInfo": {
					"Name": "count",
					"Description": "count all entries",
					"Schedule": "* * * * *",
					"Duration": -3600,
					"Script": "var time = import(\"time\")\n\naddSelfTargetedNotification(7, \"hello\", \"/#/search/486574780\", time.Now().Add(30 * time.Second))"
				}
			},
			{
				"Name": "a",
				"Type": "dashboard",
				"AdditionalInfo": {
					"UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32",
					"Name": "Foo",
					"Description": "My dashboard"
				}
			},
			{
				"Name": "ae9f2598-598f-4859-a3d4-832a512b6104",
				"Type": "pivot",
				"AdditionalInfo": {
					"UUID": "ae9f2598-598f-4859-a3d4-832a512b6104",
					"Name": "foo",
					"Description": "foobar"
				}
			}
		]
	}
]
```

## Installing a Kit

To install a kit once it has been uploaded, send a PUT request to `/api/kit/<uui>`, where the UUID is the UUID field from the list of kits. The server will return a 200 status code upon successful installation.

## Uninstalling a kit

To remove a kit, issue a DELETE request on `/api/kit/<uuid>`.