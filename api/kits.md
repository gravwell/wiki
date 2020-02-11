# Kits Web API

This API implements the creation, installation, and deletion of Gravwell kits. Kits contain other components which are installed on the local system to provide a ready-to-go solution to a particular problem. Kits can contain:

* Resources
* Scheduled searches
* Dashboards
* Auto-extractor definitions
* Templates
* Pivots
* User files
* Macros
* Search library entries

A given kit will also have the following attributes, specified at build time:

* ID: A unique identifier for this kit. We recommend following Android naming practice, e.g. "com.example.my-kit".
* Name: A human-friendly name for the kit, e.g. "My Kit".
* Description: A description of the kit.
* Version: An integer version of the kit.

## Building a kit

Kits are built by sending a POST request to `/api/kits/build` containing a KitBuildRequest structure, as defined below:

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
	SearchLibraries   []uuid.UUID    
	Extractors        []uuid.UUID    
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

This kit can be downloaded by doing a GET on `/api/kits/build/<uuid>`; given the above response, one would fetch the kit from `/api/kits/build/2f5e485a-2739-475b-810d-de4f80ae5f52`

## Uploading a Kit

Before a kit can be installed, it must first be uploaded to the webserver. Kits are uploaded by a POST request to `/api/kits`. The request should contain a multipart form. To upload a file from the local system, add a file field to the form named `file` containing the kit file. To upload a file from a remote system such as an HTTP server, add a field named `remote` containing the URL of the kit.

## Listing Kits

A GET request on `/api/kits` will return a list of all known kits. Here is an example showing the result when the system has one kit uploaded but not yet installed:

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

To install a kit once it has been uploaded, send a PUT request to `/api/kits/<uuid>`, where the UUID is the UUID field from the list of kits. The server will return a 200 status code upon successful installation.

Additional kit installation options may be specified by passing a configuration structure in the body of the request, e.g.:

```
{
	"OverwriteExisting": true,
	"Global": true,
	"AllowUnsigned": false,
	"InstallationGroup": 3,
	"Labels": ["foo", "bar"]
}
```

Note: All of the following are optional. Any or all of them may be omitted; to take the default options, simply omit the body from the request.

If set, `OverwriteExisting` tells the installer to simply replace any existing items which have the name unique identifier as the kit's version.

The `Global` flag may only be set by the administrator. If set, all items will be marked as Global, meaning all users will have access.

Regular users can only install properly-signed kits from Gravwell. If `AllowUnsigned` is set, *administrators* can install unsigned kits.

`InstallationGroup` allows the installing user to share the contents of the kit with one of the groups to which he belongs.

`Labels` is a list of additional labels which should be applied to all label-able items in the kit upon installation. Note that Gravwell automatically labels kit-installed items with "kit" and the ID of the kit (e.g. "io.gravwell.coredns").

## Uninstalling a kit

To remove a kit, issue a DELETE request on `/api/kits/<uuid>`.

## Querying Remote Kit Server

To get a list of remote kits from the Gravwell Kit Server, issue a GET on `/api/kits/remote/list`.  This will return a JSON encoded list of kit metadata structures that represents the latest versions for all available kits.  The API path `/api/kits/remote/list/all` will provide all kits for all versions.

The Metadata structure is as follows:

```
type KitMetadata struct {
	ID            string
	Name          string
	GUID          string
	Version       uint
	Description   string
	Signed        bool
	AdminRequired bool
	MinVersion    CanonicalVersion
	MaxVersion    CanonicalVersion
	Size          int64
	Created       time.Time
	Ingesters     []string //ingesters associated with the kit
	Tags          []string //tags associated with the kit
	Assets        []KitMetadataAsset
}

type KitMetadataAsset struct {
	Type     string
	Source   string //URL
	Legend   string //some description about the asset
	Featured bool
}

type CanonicalVersion struct {
	Major uint32
	Minor uint32
	Point uint32
}
```

Here is an example:

```
WEB GET http://172.19.0.2:80/api/kits/remote/list:
[
	{
		"ID": "io.gravwell.test",
		"Name": "testkit",
		"GUID": "c2870b48-ff31-4550-bd58-7b2c1c10eeb3",
		"Version": 1,
		"Description": "Testing a kit with a license in it",
		"Signed": true,
		"AdminRequired": false,
		"MinVersion": {
			"Major": 0,
			"Minor": 0,
			"Point": 0
		},
		"MaxVersion": {
			"Major": 0,
			"Minor": 0,
			"Point": 0
		},
		"Size": 0,
		"Created": "2020-02-10T16:31:23.03192303Z",
		"Ingesters": [
			"SimpleRelay",
			"FileFollower"
		],
		"Tags": [
			"syslog",
			"auth"
		],
		"Assets": [
			{
				"Type": "image",
				"Source": "cover.jpg",
				"Legend": "TEAM RAMROD!",
				"Featured": true
			},
			{
				"Type": "readme",
				"Source": "readme.md",
				"Legend": "",
				"Featured": false
			},
			{
				"Type": "image",
				"Source": "testkit.jpg",
				"Legend": "",
				"Featured": false
			}
		]
	}
]
```

### Pulling kit assets from the remote kitserver

Kits also contain assets that can be used to display images, markdown, licenses, and additional files that help explore the purpose of the kit prior to actually downloading/installing a kit.  These assets can be retrieved from the remote system by executing GET requests on `api/kits/remote/<guid>/<asset>`.  For example, if we wanted to pull back the asset of Type "image" and Legend "TEAM RAMROD!" for the kit with the guid `c2870b48-ff31-4550-bd58-7b2c1c10eeb3` you would issue a GET on `/api/kits/remote/c2870b48-ff31-4550-bd58-7b2c1c10eeb3/cover.jpg`.
