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
	Icon              string    
	Dependencies      []KitDependency
	ConfigMacros	  []KitConfigMacro
	ScriptDeployRules map[int32]ScriptDeployConfig
}
```

Note that while the ID, Name, Description, and Version fields are required, the arrays of templates/pivots/dashboards etc. are optional. For example, here is a request to build a kit containing two dashboards, a pivot, a resource, and a scheduled search:

```
{
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG"
        }
    ],
    "Dashboards": [
        7,
        10
    ],
    "Description": "Test Gravwell kit",
    "ID": "io.gravwell.test",
    "Name": "test-gravwell",
    "Pivots": [
        "ae9f2598-598f-4859-a3d4-832a512b6104"
    ],
    "Resources": [
        "84270dbd-1905-418e-b756-834c15661a54"
    ],
    "ScheduledSearches": [
        1439174790
    ],
    "ScriptDeployRules": {
        "1439174790": {
            "Disabled": true,
            "RunImmediately": false
        }
    },
    "Version": 1
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

### Dependencies

A kit may depend on other kits. List these dependencies in the Dependencies array using the following sturcture:

```
{
	ID			string
	MinVersion	uint
}
```

The ID field specifies the dependency's ID, e.g. io.gravwell.testresource. The MinVersion field specifies the minimum version of that kit which must be installed, e.g. 3.

### Config Macros

A kit may define "config macros", which are special macros which will be created by Gravwell when the kit is installed. A config macro looks like this:

```
{
	"MacroName": "KIT_WINDOWS_TAG",
	"Description": "Tag or tags containing Windows event entries",
	"DefaultValue": "windows",
	"Value": "",
	"Type": "TAG"
}
```

The UI should prompt for the desired value of the macro at installation time and include the user's response in the KitConfig structure.

Config macro definitions can include a Type field, which give a hint about the sort of value that the macro expects. The following options are currently defined:

	* "TAG": the value should be a valid tag. This tag does not necessarily have to exist on the current system, but it may be useful to check and alert the user if they enter a non-existent tag.
	* "STRING": the value can be a free-form string.

If no Type is specified, assume "STRING" (free-form entry).

### Script Deploy Configs

By default, scripts included in a kit will be set enabled at installation. This behavior can be controlled through the script deploy config structures:

```
{
	"Disabled": false,
	"RunImmediately": true,
}
```

The structure contains two fields, "Disabled" and "RunImmediately". If Disabled is set to true, the script will be installed in a disabled state. If RunImmediately is set to true, the script will be executed as soon as possible after installation, *even if the script is otherwise disabled*.

Script deploy options can be set at *kit build time*, or at *kit deploy time* to override the kit's built-in options. 

When building a kit, the `ScriptDeployRules` field should contain mappings from scheduled script ID numbers (as listed in the `ScheduledSearches` field) to script deploy config structures.

When installing a kit, the `ScriptDeployRules` field should contain mappings from scheduled script *names* to configurations. Note that deployment options only need to be specified at installation time if you wish to override the defaults.

## Uploading a Kit

Before a kit can be installed, it must first be uploaded to the webserver. Kits are uploaded by a POST request to `/api/kits`. The request should contain a multipart form. To upload a file from the local system, add a file field to the form named `file` containing the kit file. To upload a file from a remote system such as an HTTP server, add a field named `remote` containing the URL of the kit.

You can also add a field named `metadata` to the request. The contents of this field are not parsed by the server; instead, it adds the contents to the Metadata field on the uploaded kit. This allows you to keep track of e.g. the URL from which the kit originated, the date on which the kit was uploaded, etc.

The server will respond with a description of the kit which has been uploaded, e.g.:

```
{
    "AdminRequired": false,
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG",
            "Type": "TAG",
            "Value": "winlog"
        }
    ],
    "ConflictingItems": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        }
    ],
    "Description": "Test Gravwell kit",
    "GID": 0,
    "ID": "io.gravwell.test",
    "Installed": false,
    "Items": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        },
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        },
        {
            "AdditionalInfo": {
                "DefaultDeploymentRules": {
                    "Disabled": false,
                    "RunImmediately": true
                },
                "Description": "A script",
                "Name": "myScript",
                "Schedule": "* * * * *",
                "Script": "println(\"hi\")"
            },
            "Name": "5aacd602-e6ed-11ea-94d9-c771bfc07a39",
            "Type": "scheduled search"
        }
    ],
    "ModifiedItems": [
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        }
    ],
    "Name": "test-gravwell",
    "RequiredDependencies": [
        {
            "AdminRequired": false,
            "Assets": [
                {
                    "Featured": true,
                    "Legend": "Littering AAAAAAND",
                    "Source": "cover.jpg",
                    "Type": "image"
                },
                {
                    "Featured": false,
                    "Legend": "",
                    "Source": "readme.md",
                    "Type": "readme"
                }
            ],
            "Created": "2020-03-23T15:36:00.294625802-06:00",
            "Dependencies": null,
            "Description": "A simple test kit that just provides a resource",
            "ID": "io.gravwell.testresource",
            "Ingesters": [
                "simplerelay"
            ],
            "Items": [
                {
                    "AdditionalInfo": {
                        "Description": "hosts",
                        "ResourceName": "devlookup",
                        "Size": 610,
                        "VersionNumber": 1
                    },
                    "Name": "devlookup",
                    "Type": "resource"
                },
                {
                    "AdditionalInfo": "Testkit resource\n\nThis really has no restrictions, go nuts!\n",
                    "Name": "LICENSE",
                    "Type": "license"
                }
            ],
            "MaxVersion": {
                "Major": 0,
                "Minor": 0,
                "Point": 0
            },
            "MinVersion": {
                "Major": 0,
                "Minor": 0,
                "Point": 0
            },
            "Name": "Testing resource kit",
            "Signed": true,
            "Size": 10240,
            "Tags": [
                "syslog"
            ],
            "UUID": "d2a0cb10-ff25-4426-8b87-0dd0409cae48",
            "Version": 1
        }
    ],
    "Signed": false,
    "UID": 7,
    "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
    "Version": 2
}
```

Note the "ModifiedItems" field. If an earlier version of this kit is already installed, this field will contain a list of items which *the user has modified*. Installing the staged kit will overwrite these items, so users should be notified and given a chance to save their changes.

"ConflictingItems" lists items which appear to conflict with user-created objects. In this example, it appears that the user has previously created their own resource named "maxmind_asn". If an installation request is sent with `OverwriteExisting` set to true, that resource will be overwritten with the version in the kit; if set to false, the installation process will return an error.

The "RequiredDependencies" field contains a list of metadata structures for any currently-uninstalled dependencies of this kit, including an Items set which may contain licenses which should be displayed.

The ConfigMacros field contains a list of configuration macros (see previous section) which will be installed by this kit. If a previous version of this kit (or another kit altogether) has already installed a macro with the same name, the webserver will pre-populate the "Value" field with the current value in the macro. If a *user* has previously installed a macro with the same name, the webserver will return an error.

Take note of the scheduled search named "myScript", particularly the `DefaultDeploymentRules` field. This describes how the script will be installed: it will be marked enabled, and it will run as soon as possible.

## Listing Kits

A GET request on `/api/kits` will return a list of all known kits. Here is an example showing the result when the system has one kit uploaded but not yet installed:

```
[
    {
        "AdminRequired": false,
        "Description": "Test Gravwell kit",
        "GID": 0,
        "ID": "io.gravwell.test",
        "Installed": false,
        "Items": [
            {
                "AdditionalInfo": {
                    "Description": "ASN database",
                    "ResourceName": "maxmind_asn",
                    "Size": 6196221,
                    "VersionNumber": 1
                },
                "Name": "84270dbd-1905-418e-b756-834c15661a54",
                "Type": "resource"
            },
            {
                "AdditionalInfo": {
                    "DefaultDeploymentRules": {
                        "Disabled": false,
                        "RunImmediately": true
                    },
                    "Description": "count all entries",
                    "Duration": -3600,
                    "Name": "count",
                    "Schedule": "* * * * *",
                    "Script": "var time = import(\"time\")\n\naddSelfTargetedNotification(7, \"hello\", \"/#/search/486574780\", time.Now().Add(30 * time.Second))"
                },
                "Name": "55c81086",
                "Type": "scheduled search"
            },
            {
                "AdditionalInfo": {
                    "Description": "My dashboard",
                    "Name": "Foo",
                    "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
                },
                "Name": "a",
                "Type": "dashboard"
            },
            {
                "AdditionalInfo": {
                    "Description": "foobar",
                    "Name": "foo",
                    "UUID": "ae9f2598-598f-4859-a3d4-832a512b6104"
                },
                "Name": "ae9f2598-598f-4859-a3d4-832a512b6104",
                "Type": "pivot"
            }
        ],
        "Name": "test-gravwell",
        "Signed": false,
        "UID": 7,
        "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
        "Version": 1
    }
]
```

See the listing at the end of this page for a list of what "AdditionalInfo" fields are available for each type of kit item.

## Kit Info

A GET request on `/api/kits/<GUID>` where `<GUID>` is a guid of a specifically installed or staged kit will provide info about that specific kit.

For example, a GET request on `/api/kits/549c0805-a693-40bd-abb5-bfb29fc98ef1` will yield:

```
{
    "AdminRequired": false,
    "Description": "Test Gravwell kit",
    "GID": 0,
    "ID": "io.gravwell.test",
    "Installed": false,
    "Items": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        },
        {
            "AdditionalInfo": {
                "DefaultDeploymentRules": {
                    "Disabled": false,
                    "RunImmediately": true
                },
                "Description": "count all entries",
                "Duration": -3600,
                "Name": "count",
                "Schedule": "* * * * *",
                "Script": "var time = import(\"time\")\n\naddSelfTargetedNotification(7, \"hello\", \"/#/search/486574780\", time.Now().Add(30 * time.Second))"
            },
            "Name": "55c81086",
            "Type": "scheduled search"
        },
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        },
        {
            "AdditionalInfo": {
                "Description": "foobar",
                "Name": "foo",
                "UUID": "ae9f2598-598f-4859-a3d4-832a512b6104"
            },
            "Name": "ae9f2598-598f-4859-a3d4-832a512b6104",
            "Type": "pivot"
        }
    ],
    "Name": "test-gravwell",
    "Signed": false,
    "UID": 7,
    "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
    "Version": 1
}

```

If the kit does not exist a 404 is returned, if the user does not have access to the specific kit requested a 400 is returned.

## Installing a Kit

To install a kit once it has been uploaded, send a PUT request to `/api/kits/<uuid>`, where the UUID is the UUID field from the list of kits. The server will perform some preliminary checks and return an integer, which can be used to query the progress of the installation using the installation status API (see below).

During installation, all required dependencies (as listed in the RequiredDepdencies field of the staging response) will be staged and installed automatically before the final installation of the kit itself.

Additional kit installation options may be specified by passing a configuration structure in the body of the request, e.g.:

```
{
    "AllowUnsigned": false,
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG",
            "Value": "winlog"
        }
    ],
    "Global": true,
    "InstallationGroup": 3,
    "Labels": [
        "foo",
        "bar"
    ],
    "OverwriteExisting": true,
    "ScriptDeployRules": {
        "myScript": {
            "Disabled": true,
            "RunImmediately": false
        }
    }
}
```

Note: All of the following are optional. Any or all of them may be omitted; to take the default options, simply omit the body from the request.

If set, `OverwriteExisting` tells the installer to simply replace any existing items which have the name unique identifier as the kit's version.

The `Global` flag may only be set by the administrator. If set, all items will be marked as Global, meaning all users will have access.

Regular users can only install properly-signed kits from Gravwell. If `AllowUnsigned` is set, *administrators* can install unsigned kits.

`InstallationGroup` allows the installing user to share the contents of the kit with one of the groups to which he belongs.

`Labels` is a list of additional labels which should be applied to all label-able items in the kit upon installation. Note that Gravwell automatically labels kit-installed items with "kit" and the ID of the kit (e.g. "io.gravwell.coredns").

`ConfigMacros` is the list of ConfigMacros found in the kit information structure, with the "Value" fields optionally set to whatever the user wishes. If the "Value" field is blank, the webserver will use the "DefaultValue".

`ScriptDeployRules` should contain overrides for any scheduled scripts in the kit whose deployment rules you wish to override. In this example, a script named "myScript" will be installed in a disabled state. If the default deployment options are acceptable, this field can be left empty.

### Installation Status API

When an installation request is sent, the server places the request into a queue for processing, since installation of a large package with many dependencies may take some time. The server responds to the installation request with an integer, e.g. `2019727887`. This can be used with the installation status API to query the progress of the installation by sending a GET request to `/api/kits/status/<id>`, e.g. `/api/kits/status/2019727887` might return this:

```
{
    "CurrentStep": "Done",
    "Done": true,
    "Error": "",
    "InstallID": 2019727887,
    "Log": "\nQueued installation of kit io.gravwell.testresource, with 0 dependencies also to be installed\nBeginning installation of io.gravwell.testresource (9b701e75-76ee-40fc-b9b5-4c7e1706339d) for user Admin John (1)\nInstalling requested kit io.gravwell.testresource\nDone",
    "Owner": 1,
    "Percentage": 1,
    "Updated": "2020-03-25T15:39:37.184221203-06:00"
}
```

"Owner" is the UID of the user who submitted the installation request. "Done" is set to true when the kit is fully installed. "Percentage" is a value between 0 and 1 which indicates how much of the installation has been completed. "CurrentStep" is the current status of the installation, while "Log" maintains a complete record of statuses from the entire installation. "Error" will be empty unless something has gone wrong with the installation process. "Updated" is the time at which the status was last modified.

One may also request a list of *all* kit installation statuses by doing a GET on `/api/kits/status`, which returns an array of the sort of objects seen above. Note that by default this will only return statuses for the current user; administrators may append `?admin=true` to the URL to get *all* statuses on the system.

## Uninstalling a kit

To remove a kit, issue a DELETE request on `/api/kits/<uuid>`. If any of the items in the kit have been modified by the user since installation, the response will have a 400 status code and contain a structure detailing what has changed:

```
{
    "Error": "Kit items have been modified since installation, set ?force=true to override",
    "ModifiedItems": [
        {
            "AdditionalInfo": {
                "Description": "Network services (protocol + port) database",
                "ResourceName": "network_services",
                "Size": 531213,
                "VersionNumber": 1
            },
            "ID": "2e4c8f31-92a4-48b5-a040-d2c895caf0b2",
            "KitID": "io.gravwell.networkenrichment",
            "KitName": "Network enrichment",
            "KitVersion": 1,
            "Name": "network_services",
            "Type": "resource"
        }
    ]
}
```

The UI should prompt the user at this point; to force removal of the kit, add the `?force=true` parameter to the request.

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

## Pull Single Kit Information

The Remote kit API also supports pulling back information about a specific kit by issuing a `GET` on `/api/kits/remote/<guid>`, which will return a single `KitMetadata` structure.

For example if we issue a `GET` on `/api/kits/remote/c2870b48-ff31-4550-bd58-7b2c1c10eeb3` the webserver will return:

```
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
```

### Pulling kit assets from the remote kitserver

Kits also contain assets that can be used to display images, markdown, licenses, and additional files that help explore the purpose of the kit prior to actually downloading/installing a kit.  These assets can be retrieved from the remote system by executing GET requests on `api/kits/remote/<guid>/<asset>`.  For example, if we wanted to pull back the asset of Type "image" and Legend "TEAM RAMROD!" for the kit with the guid `c2870b48-ff31-4550-bd58-7b2c1c10eeb3` you would issue a GET on `/api/kits/remote/c2870b48-ff31-4550-bd58-7b2c1c10eeb3/cover.jpg`.


## Kit item "Additional Info" fields

When listing kits (GET on `/api/kits`), each kit will include a list of items, which contain AddditionalInfo fields. These fields give more information about the items within the kit; the contents vary based on the item type and are enumerated below:

```
Resources:
		VersionNumber int
		ResourceName  string
		Description   string
		Size          uint64

Scheduled Search:
		Name                    string
		Description             string
		Schedule                string
		SearchString            string 
		Duration                int64  
		Script                  string 
		DefaultDeploymentRules  ScriptDeployConfig

Dashboard:
		UUID        string
		Name        string
		Description string

Extractor:
		Name   string 
		Desc   string 
		Module string 
		Tag    string 

Template:
		UUID        string
		Name        string
		Description string

Pivot:
		UUID        string
		Name        string
		Description string

File:
		UUID        string
		Name        string
		Description string
		Size        int64
		ContentType string

Macro:
		Name      string
		Expansion string

Search Library:
		Name        string
		Description string
		Query       string

Playbook:
		UUID        string
		Name        string
		Description string

License:
		(contents of license file itself)
```

## Kit Build Request History

Successful kit build requests are stored by the webserver. You can get a list of build requests for the current user by sending a GET request to `/api/kits/build/history`. The response will be an array of build requests:

```
[{"ID":"io.gravwell.test","Name":"test","Description":"","Version":1,"MinVersion":{"Major":0,"Minor":0,"Point":0},"MaxVersion":{"Major":0,"Minor":0,"Point":0},"Macros":[4,41],"ConfigMacros":null}]
```

Note: This store is keyed on UID + kit ID; if I build a kit named "io.gravwell.test" again, it will overwrite the version in the store.

You can delete a particular item by sending a DELETE request to `/api/kits/build/history/<id>`, e.g. `/api/kits/build/history/io.gravwell.test`.
