# Search Library

REST API located at /api/library

The search library API is used to store and retrieve saved search queries.  The search library is a useful system to build up valuable queries with names, notes, tags, and anything else that might be valuable in day-to-day operations.

The search library is a permissioned API, which means each user owns their search library and can optionally share them with groups.  Each library entry also has a global flag, which means any user can read the library entry.  ONLY admins can set the global flag.

Only owners can delete library entries, even if another user has access to a library entry through group membership, they cannot delete the entry.

## Admin Operations

Administrators interact with the search library in the same way as all other users.  If an admin wants administrative access to the library API, whether to list, delete, or modify library entries outside of their ownership they must append the `admin` flag to requests.

For example, performing a GET request on `/api/library` will return only the calling users library entries (admin or not).  Performing the same GET request on `/api/library?admin=true` will return all library entries for all users.  The admin flag is ignored for non-admin users.

## Basic API Overview

The Library API is rooted at `/api/library` and responds to the following request methods:

| method | description | Supports Admin Calls |
| ------ | ----------- | -------------------- |
| GET    | Retrieve a list of available library entries | TRUE |
| POST   | Add a new library entry | FALSE |
| PUT    | Update an existing library entry | TRUE |
| DELETE | Delete an existing library entry | TRUE |

Every search library entry has both a "ThingUUID" and a "GUID" associated with it. The "ThingUUID" is always unique: there will only ever be one search library entry with a given ThingUUID on the system. The "GUID", on the other hand, is an ID which is used to refer to a search library entry from other things like dashboards or actionables; when you install a kit, all the search library entries will have the same GUID as on the kit creator's system, allowing cross-linking. Each user could potentially have a search library entry with the same GUID, but each of those entries will have a unique ThingUUID.

The `DELETE` and `PUT` method require that the "ThingUUID" or "GUID" of a specific library entry be appended to the URL. The ThingUUID and the GUID can be used interchangeably in the API, but be aware that in "admin mode" there may be multiple accessible items with the same GUID. For example, to update the entry with the GUID of `5f72d51e-d641-11e9-9f54-efea47f6014a` a `PUT` request would be issued against `/api/library/5f72d51e-d641-11e9-9f54-efea47f6014a` with a complete entry structure encoded into the body of the request.

The `GET` method can optionally append a GUID to request a specific library entry, if not GUID is present the GET method returns a list of all available entries.  If a users does not have access to a specific entry specified by the GUID the webserver will return a response of 403.

The structure of a library entry is as follows:

```
struct {
	ThingUUID   uuid.UUID
	GUID        uuid.UUID
	UID         int32
	GIDS        []int32
	Global      boolean
	Name        string
	Description string
	Query       string
	Labels      []string
	Metadata    RawObject
}
```

The structure members are:

| Member      | Description                     | Omitted if Empty |
| ----------- | ------------------------------- | ---------------- |
| ThingUUID   | Unique identifier *on the local system* |                  |
| GUID        | Global "name"; persists across kit installation. Multiple users may have search library entries with the same GUID. | |
| UID         | Owners system user id           |                  |
| GIDs        | List of group ids the entry is shared with | X |
| Global      | Boolean indicating whether the entry is globally readable | |
| Name        | A human readable name for the query | |
| Description | A human readable description of the query | |
| Query       | The query string | |
| Labels      | A list of human readable labels used for categorizing queries | X |
| Metadata    | An opaque JSON blob used for arbitrary parameter storage, just be valid JSON but does not have any particular structure | X |

Here is an example entry that is owned by the user with UID 1 and shared with 3 groups.  The UID and GID values must be mapped back to user and group names using the user API:

```
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"GIDs": [1, 3, 5],
	"Global": false,
	"Name": "syslog counting query",
	"Description": "A simple chart that shows total syslog over time",
	"Query": "tag=syslog syslog Appname | stats count by Appname | chart count by Appname",
	"Labels": [
		"syslog",
		"chart",
		"trending"
	],
	"Metadata": {}
}

```


## Examples API Interaction

This section contains example interactions with the search library API endpoint.  These examples were generated using the Gravwell CLI using the `-debug` flag.

### Creating a New Entry

Request:
```
POST /api/library
{
	"GIDs": [1, 2],
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	]
}
```

Note: Because the "GUID" field is not set here, the system will assign one. You may also include the GUID in the request.

Response:
```
{
	"ThingUUID": "c9169d15-d643-11e9-99d3-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"GIDs": [1, 2],
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	]
}

```
### Retrieving Entries

Request:

```
GET http://172.19.0.5:80/api/library
```

Response:
```
[
	{
		"ThingUUID": "0b5a66cb-d642-11e9-931c-0242ac130005",
		"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
		"UID": 1,
		"Global": false,
		"Name": "netflow agg",
		"Description": "Total traffic using netflow data",
		"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
		"Labels": [
			"netflow",
			"traffic",
			"aggs"
		],
		"Metadata": {
			"value": 1,
			"extra": "some extra field value"
		}
	},
	{
		"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
		"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
		"UID": 1,
		"Global": false,
		"Name": "test2",
		"Description": "testing second",
		"Query": "tag=foo grep bar",
		"Labels": [
			"foo",
			"bar",
			"baz"
		],
	}
]
```

### Requesting Specific Entry

Request:

```
GET http://172.19.0.5:80/api/library/ae132ecc-88dd-11ea-a6aa-373f4c2439d4
```

Response:
```
{
	"ThingUUID": "0b5a66cb-d642-11e9-931c-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	],
	"Metadata": {
		"value": 1,
		"extra": "some extra field value"
	}
}
```

Note that you would get the same response from `api/library/0b5a66cb-d642-11e9-931c-0242ac130005` too.

### Updating an Entry

Request:
```
PUT /api/library/69755a85-d5b1-11e9-89c2-0242ac130005
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
	"Global": false,
	"Name": "SyslogAgg",
	"Description": "Updated Syslog aggregate",
	"Query": "tag=syslog length | stats sum(length) | chart sum",
	"Labels": [
		"syslog",
		"agg",
		"totaldata"
	],
	"Metadata": {}
}
```

Response:
```
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
	"UID": 1,
	"Global": false,
	"Name": "SyslogAgg",
	"Description": "Updated Syslog aggregate",
	"Query": "tag=syslog length | stats sum(length) | chart sum",
	"Labels": [
		"syslog",
		"agg",
		"totaldata"
	]
}
```

### Deleting an Entry

Request:
```
DELETE /api/library/69755a85-d5b1-11e9-89c2-0242ac130005
```

### Admin Deleting an Entry

If an non-admin appends the admin flag the webserver will ignore the flag, if the non-admin is the owner of the specified entry the action (DELETE in this case) still works.  If the non-admin user does NOT own the entry the webserver responds with a 403 StatusForbidden.

When performing an admin deletion, always use the ThingUUID in the URL parameter or else the wrong item may be deleted.

Request:
```
DELETE /api/library/69755a85-d5b1-11e9-89c2-0242ac130005?admin=true
```
