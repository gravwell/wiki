# Playbooks Web API

This API is used to manipulate playbooks withing Gravwell. A playbook is a user-friendly way to bring together notes and search queries into a user-formatted document.

The playbook structure contains the following components:

* `UUID`: The unique identifier of the playbook, set at installation time.
* `GUID`: The global name of the playbook; this is set by the original creator of the playbook and will remain the same even if the playbook is bundled into a kit and installed on another system. Each user may only have one playbook with a given GUID, but multiple users could each have a copy of a playbook with the same GUID.
* `UID`: The user ID of the playbook owner.
* `GIDs`: A list of group IDs which are allowed to access this playbook.
* `Global`: A boolean flag; if set, all users on the system may view the playbook.
* `Name`: The user-friendly name of the playbook.
* `Desc`: A description of the playbook.
* `Body`: A byte array which stores the playbook contents.
* `Metadata`: A byte array which stores playbook metadata, for use by the client.
* `Labels`: An array of strings containing optional labels to apply to the playbook.
* `LastUpdated`: A timestamp indicating when the playbook was last modified.
* `Author`: A structure containing information about the author of the playbook (see below).
* `Synced`: Used internally by Gravwell.

Note that the UUID and GUID fields may be used interchangeably in all API calls. This is so playbooks included in kits may link to each other, by using links containing GUIDs which will persist across kit installation.

The author information structure contains the following fields, any of which may be left blank:

* `Name`: The author's name.
* `Email`: The author's email address.
* `Company`: The author's company.
* `URL`: A web address for more information about the author.

## Listing Playbooks

To list playbooks, send a GET request to `/api/playbooks`. The server will respond with an array of playbook structures which the user has permission to view:

```
[
  {
    "UUID": "2cbc8500-5fc5-453f-b292-8386fe412f5b",
    "GUID": "c9da126b-1608-4740-a7cd-45495e8341a3",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Netflow V5 Playbook",
    "Desc": "A top-level playbook for netflow, with background and starting points.",
    "Body": "",
    "Metadata": "eyJkYXNoYm9hcmRzIjpbXSwiYXR0YWNobWVudHMiOlt7ImNvbnRleHQiOiJjb3ZlciIsImZpbGVHVUlEIjoiNDhjNmIwZWYtNmU3Ni00MjA4LWJjYTctMGI5NWU0NzAwYmRkIiwidHlwZSI6ImltYWdlIn1dfQ==",
    "Labels": [
      "netflow",
      "netflow-v5",
      "kit/io.gravwell.netflowv5"
    ],
    "LastUpdated": "2020-08-14T16:17:03.778971838-06:00",
    "Author": {
      "Name": "John Floren",
      "Email": "john@example.org",
      "Company": "Gravwell",
      "URL": "http://grawell.io"
    },
    "Synced": false
  },
  {
    "UUID": "973fcc22-1964-4efa-848c-7196ac67094e",
    "GUID": "dbd84b95-11b7-450d-9111-9bb33d63741b",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Network Enrichment Kit Overview",
    "Desc": "",
    "Body": "",
    "Metadata": "eyJkYXNoYm9hcmRzIjpbXSwiYXR0YWNobWVudHMiOlt7ImNvbnRleHQiOiJjb3ZlciIsImZpbGVHVUlEIjoiOGIwZjQzMjItOTY1My00OTQyLWJkODctY2Y4ZWM5NjZmNmFmIiwidHlwZSI6ImltYWdlIn1dfQ==",
    "Labels": [
      "kit/io.gravwell.networkenrichment"
    ],
    "LastUpdated": "2020-08-05T12:14:48.739069332-06:00",
    "Author": {
      "Name": "John Floren",
      "Email": "john@example.org",
      "Company": "Gravwell",
      "URL": "http://grawell.io"
    },
    "Synced": false
  }
]
```

Note that the Body parameter is empty; because playbooks can be quite large, the body is left out when listing all playbooks.

Appending the `?admin=true` parameter to the URL will return a list of *all* playbooks on the system, provided the user is an Administrator.

## Fetching a Playbook

To get a specific playbook, including the Body, send a GET request to `/api/playbooks/<uuid>`. The web server will attempt to find a playbook with a matching UUID field; if that is not successful, it will look for a playbook that the user can read with the following precedence:

* Top precedence: playbooks owned by the user.
* Next: playbooks shared with one of the user's groups.
* Finally: playbooks with the Global flag set.

## Creating a Playbook

Playbooks are created by sending a POST request to `/api/playbooks`. The body of the request should contain those fields the user wishes to set; note that the server will ignore the UUID, UID, LastUpdated, and Synced fields if set.

```
{
    "Body": <contents of the playbook>,
	"Metadata": <any desired metadata>,
    "Name": "ssh syslog",
    "Desc": "A playbook for monitoring syslog entries for ssh sessions",
    "GIDs": null,
    "Global": true,
	"Author": {
		"Name": "Dean Martin"
	},
    "Labels": [
        "syslog"
    ]
}
```

The server will respond with the UUID of the newly-created playbook. If the `GUID` field is set in the request, the server will use it, otherwise it will generate a new one.

## Modifying a Playbook

To update the contents of a playbook, send a PUT request to `/api/playbooks/<uuid>`, where the UUID matches the desired playbook. The body of the request should contain the playbook structure to be updated. Note that changes to the UUID, GUID, LastUpdated, and Synced fields will be ignored. Administrators are allowed to modify the UID field, but regular users cannot.

Note: If you do not intend to update the contents of a field, you should send the original value in the request. The server has no way to know if e.g. an un-set "Desc" field means you wish to preserve the original value, or you wish to clear the field.

## Deleting a Playbook

To delete a playbook, send a DELETE request to `/api/playbooks/<uuid>`, where the UUID matches the desired playbook.