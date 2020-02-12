# Playbooks Web API

This API is used to manipulate playbooks withing Gravwell. A playbook is a user-friendly way to bring together notes and search queries into a user-formatted document.

The playbook structure contains the following components:

* `UUID`: The unique identifier of the playbook, set at installation time.
* `UID`: The user ID of the playbook owner.
* `GIDs`: A list of group IDs which are allowed to access this playbook.
* `Global`: A boolean flag; if set, all users on the system may view the playbook.
* `Name`: The user-friendly name of the playbook.
* `Desc`: A description of the playbook.
* `Body`: A byte array which stores the playbook contents.
* `Labels`: An array of strings containing optional labels to apply to the playbook.
* `LastUpdated`: A timestamp indicating when the playbook was last modified.
* `Synced`: Used internally by Gravwell


## Listing Playbooks

To list playbooks, send a GET request to `/api/playbooks`. The server will respond with an array of playbook structures which the user has permission to view:

```
[{"UUID":"0bbff773-9ee2-4874-89a4-bf85a4b800df","UID":1,"GIDs":null,"Global":false,"Name":"foo","Desc":"bar","Body":"blah","Labels":["test"],"LastUpdated":"2020-02-12T10:13:02.864666484-07:00","Synced":false}]
```

Appending the `?admin=true` parameter to the URL will return a list of *all* playbooks on the system, provided the user is an Administrator.

## Creating a Playbook

Playbooks are created by sending a POST request to `/api/playbooks`. The body of the request should contain those fields the user wishes to set; note that the server will ignore the UUID, UID, LastUpdated, and Synced fields if set.

```
{
    "Body": <contents of the playbook>,
    "Name": "ssh syslog",
    "Desc": "A playbook for monitoring syslog entries for ssh sessions",
    "GIDs": null,
    "Global": true,
    "Labels": [
        "syslog"
    ]
}
```

The server will respond with the UUID of the newly-created playbook.

## Modifying a Playbook

To update the contents of a playbook, send a PUT request to `/api/playbooks/<uuid>`, where the UUID matches the desired playbook. The body of the request should contain the playbook structure to be updated. Note that changes to the UUID, LastUpdated, and Synced fields will be ignored. Administrators are allowed to modify the UID field, but regular users cannot.

Note: If you do not intend to update the contents of a field, you should send the original value in the request. The server has no way to know if e.g. an un-set "Desc" field means you wish to preserve the original value, or you wish to clear the field.

## Deleting a Playbook

To delete a playbook, send a DELETE request to `/api/playbooks/<uuid>`, where the UUID matches the desired playbook.