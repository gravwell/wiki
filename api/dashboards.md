# Dashboard Storage API

The dashboard api is essentially a generic CRUD api for managing json blobs used by the GUI to render dashboards. Searches present on a dashboard are launched by the GUI and the backend/frontend/webserver doesn't really have a concept of what they are.

## Creating a dashboard

To add a dashboard a **POST** request is issued to **/api/dashboards** with a payload of the following format:

The 'Data' property is the JSON used by the GUI to create the actual dashboard.

```
{
        "Name": "test2",
        "Description": "test2 description",
		"UID": 2,
		"GIDs": [],
        "Data": {
                "A": "A2",
                "B": "B2",
                "C": 12610078956637388,
                "D": false
        }
}
```

If the "UID" parameter is omitted from the request, it should default to the UID of the requesting user.

The webserver's response contains the ID of the newly-created dashboard.

## Retrieving Dashboards

### Getting all dashboards for the current user

A user can get all dashboards to which they have access (via ownership or group) by issuing a **GET** request on **/api/dashboards**.

```
[
        {
                "ID": 2,
                "Name": "test2",
                "UID": 3,
                "GIDs": [],
                "Description": "test2 description",
                "Created": "2016-12-18T23:28:08.250051418Z",
				"Guid": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]

```
### Getting a specific dashboard
To fetch a particular ID, put its ID on the end of the dashboards URL:

```
GET /api/dashboards/2:
```

It is also possible to fetch a specific dashboard by GUID, although this is not recommended
### Admin getting ALL dashboards of ALL users
To get all dashboards the user MUST be an admin, and issue a **GET** request to **/api/dashboards/all**. If this request is issued by a non-admin user it should return all dashboards to which they have access (which would be equivocal to a GET on **/api/dashboards**)

```
WEB GET /api/dashboards/all:
[
        {
                "ID": 1,
                "Name": "test1",
                "UID": 1,
                "GIDs": [],
                "Description": "test1 description",
                "Created": "2016-12-18T23:28:07.679322121Z",
				"Guid": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
                "Data": {
                        "A": "A",
                        "B": "B",
                        "C": 57646075230342348,
                        "D": true
                }
        },
        {
                "ID": 2,
                "Name": "test2",
                "UID": 3,
                "GIDs": [],
                "Description": "test2 description",
                "Created": "2016-12-18T23:28:08.250051418Z",
				"Guid": "55bc7236-39e4-11e9-94e9-54e1ad7c66cf",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```

## Updating a Dashboard
Updating a dashboard (to change the data or alter the name, description, GID list, etc) is done by issuing a **PUT** request to **/api/dashboards/ID** where ID is the unique identifier for a given dashboard.

In this example a user (UID 3) wishes to add permission for group 3 to access a dashboard (ID 2).

```
GET /api/dashboards/2:
[
        {
                "ID": 2,
                "Name": "test2",
                "UID": 3,
                "GIDs": [],
                "Description": "test2 description",
				"Guid": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
                "Created": "2016-12-18T23:28:08.250051418Z",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```

The user has now retrieved the target dashboard, makes the modifications, and posts a PUT request:
```
WEB PUT /api/dashboards/2:
[
        {
                "ID": 2,
                "Name": "marketoverview",
                "UID": 3,
                "GIDs": [3],
                "Description": "marketing group dashboard",
				"Guid": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
                "Created": "2016-12-18T23:28:08.250051418Z",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```

The server will respond to update requests with the updated dashboard structure.

## Deleting a dashboard
To remove a dashboard issue a request with the **DELETE** method on the url **/api/dashboards/ID** where ID is the numeric ID of the dashboard.

## Getting all dashboards owned by a user
To get dashboards explictely owned by a user issue a **GET** request on **/api/users/UID/dashboards** which will hand back ONLY those dashboards specifically owned by that UID.  This WILL NOT include dashboards the user has access to through group memberships.
```
WEB GET /api/users/1/dashboards:
[
                {
                "ID": 4,
                "Name": "dashGroup2",
                "UID": 5,
                "GIDs": [
                        3
                ],
                "Description": "dashGroup2",
                "Created": "2016-12-28T21:37:12.703358455Z",
				"Guid": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```

## Getting all of a groups dashboards
To get all dashboards that a specific group has access to issue a **GET** request to **/api/groups/GID/dashboards** which will hand back any dashboards tagged to that group.  This deviates a little from normal Unix permissions in that a dashboard can be in multiple groups (so groups don't really OWN a dashboard, but rather dashboards are kind of members of a group).
```
WEB GET /api/groups/2/dashboards:
[
        {
                "ID": 3,
                "Name": "dashGroup1",
                "UID": 5,
                "GIDs": [
                        2
                ],
                "Description": "dashGroup1",
                "Created": "2016-12-28T21:37:12.696460531Z",
				"Guid": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        },
        {
                "ID": 2,
                "Name": "test2",
                "UID": 3,
                "GIDs": [],
                "Description": "test2 description",
                "Created": "2016-12-18T23:28:08.250051418Z",
				"Guid": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```
### Admin getting ALL dashboards of ALL users
To get all dashboards the user MUST be an admin, and issue a **GET** request to **/api/dashboards/all**. If this request is issued by a non-admin user it should return all dashboards to which they have access (which would be equivocal to a GET on **/api/dashboards**)

```
WEB GET /api/dashboards/all:
[
        {
                "ID": 1,
                "Name": "test1",
                "UID": 1,
                "GIDs": [],
                "Description": "test1 description",
                "Created": "2016-12-18T23:28:07.679322121Z",
				"Guid": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
                "Data": {
                        "A": "A",
                        "B": "B",
                        "C": 57646075230342348,
                        "D": true
                }
        },
        {
                "ID": 2,
                "Name": "test2",
                "UID": 3,
                "GIDs": [],
                "Description": "test2 description",
                "Created": "2016-12-18T23:28:08.250051418Z",
				"Guid": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
                "Data": {
                        "A": "A2",
                        "B": "B2",
                        "C": 12610078956637388,
                        "D": false
                }
        }
]
```
