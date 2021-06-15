# Dashboard Storage API

The dashboard api is essentially a generic CRUD api for managing json blobs used by the GUI to render dashboards. Searches present on a dashboard are launched by the GUI and the backend/frontend/webserver doesn't really have a concept of what they are.

## Data Structure

* ID: A 64-bit integer which uniquely identifies the dashboard.
* GUID: A UUID which designates the dashboard. Persists across kit installation (see below). Used when referring to dashboards from actionables.
* Name: A user-friendly name for the dashboard.
* Description: A more detailed description of the dashboard.
* UID: The numeric ID of the dashboard's owner.
* GIDs: An array of numeric group IDs with which this dashboard is shared.
* Global: A boolean, set to true if dashboard should be visible to all users (admin only).
* Created: The timestamp at which the dashboard was created.
* Updated: The timestamp at which the dashboard was last updated.
* Labels: An array of strings containing [labels](#!gui/labels/labels.md).
* Data: The actual definition of the dashboard contents (see below).

Note that every dashboard has both an `ID` field and a `GUID` field. This is because dashboards may be packed in kits along with actionables which *refer* to those dashboards. A dashboard packed into a kit includes its existing GUID, and that GUID is preserved when the kit is installed, so it is safe for actionables to refer to the dashboard by its GUID. The ID field, on the other hand, is randomly generated whenever a dashboard is created or installed. A given system may actually have several dashboards with the same GUID (installed by different users, typically) but each dashboard will have its own unique ID.

Although the webserver does not care what goes into the `Data` field (except that it should be valid JSON), there is a particular format which the **GUI** uses. The following is a complete Typescript definition of the Dashboard structure including the Data field as used by the GUI:

```
interface RawDashboard {
    ID: RawNumericID;
    GUID: RawUUID;

    UID: RawNumericID;
    GIDs: Array<RawNumericID> | null;

    Name: string;
    Description: string; // empty string is null
    Labels: Array<string> | null;

    Created: string; // Timestamp
    Updated: string; // Timestamp

    Data: {
        liveUpdateInterval?: number; // 0 is undefined
        linkZooming?: boolean;

        grid?: {
            gutter?: string | number | null; // string is a number
            margin?: string | number | null; // string is a number
        };

        searches: Array<{
            alias: string | null;
            timeframe?: {} | RawTimeframe;
            query?: string;
            searchID?: RawNumericID;
            reference?: {
                id: RawUUID;
                type: 'template' | 'savedQuery' | 'scheduledSearch';
                extras?: {
                    defaultValue: string | null;
                };
            };
        }>;
        tiles: Array<{
            id: RawNumericID;
            title: string;
            renderer: string;
            span: { col: number; row: number; x: number; y: number };
            searchesIndex: number;
            rendererOptions: RendererOptions;
        }>;
        timeframe: RawTimeframe;
        version?: 1 | 2;
        lastDataUpdate?: string; // Timestamp
    };
}

interface RawTimeframe {
    durationString: string | null;
    timeframe: string;
    start: string | null; // Timestamp
    end: string | null; // Timestamp
}

interface RendererOptions {
    XAxisSplitLine?: 'no';
    YAxisSplitLine?: 'no';
    IncludeOther?: 'yes';
    Stack?: 'grouped' | 'stacked';
    Smoothing?: 'normal' | 'smooth';
    Orientation?: 'v' | 'h';
    ConnectNulls?: 'no' | 'yes';
    Precision?: 'no';
    LogScale?: 'no';
    Range?: 'no';
    Rotate?: 'yes';
    Labels?: 'no';
    Background?: 'no';
    values?: {
        Smoothing?: 'smooth';
        Orientation?: 'h';
        columns?: Array<string>;
    };
}
```

Note: Throughout this document, we include valid `Data` structures, but for brevity we will tend to use structures describing an empty dashboard rather than one containing tiles.

## Creating a dashboard

To add a dashboard, issue a POST request to `/api/dashboards` with a payload in the following format:

```
{
        "Name": "test2",
        "Description": "test2 description",
		"UID": 2,
		"GIDs": [],
		"Global": false,
        "Data": {
			"tiles": []
        }
}
```

The `Data` property is the JSON used by the GUI to create the actual dashboard. It can be initialized or left blank for later populating; here we have included an empty "tiles" field for demonstration purposes.

If the `UID` parameter is omitted from the request, it will default to the UID of the requesting user. Only admin users can set the UID to any ID except their own.

The `GIDs` array should contain a list of group IDs with which the dashboard will be shared. If left empty, the dashboard is not shared with anyone.

Admin users may also set the boolean `Global` field to true; setting this will allow all users on the system to access the dashboard.

If a `GUID` field is included and is a valid UUID, it will be used for the dashboard instead of a randomly generated one. In most cases, the GUID field should be left blank so the webserver can generate a random one.

The webserver's response contains the numeric ID of the newly-created dashboard.

## Listing Dashboards

### Getting all dashboards for the current user

A user can get all dashboards to which they have access (via ownership or group) by issuing a GET request on `/api/dashboards`. The following shows a response containing two dashboards:

```
[
  {
    "ID": 203486809563715,
    "Name": "test",
    "UID": 1,
    "GIDs": [],
    "Description": "test dashboard",
    "Created": "2020-09-22T09:16:51.66798721-06:00",
    "Updated": "2020-09-22T09:17:06.241311128-06:00",
    "Data": {
      "searches": [
        {
          "alias": "Search 1",
          "timeframe": {},
          "query": "tag=* count\n",
          "searchID": 4780372388
        }
      ],
      "tiles": [
        {
          "title": "count",
          "renderer": "text",
          "span": {
            "col": 4,
            "row": 4,
            "x": 0,
            "y": 0
          },
          "searchesIndex": 0,
          "id": 16007878262310,
          "rendererOptions": {}
        }
      ],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      },
      "version": 2,
      "lastDataUpdate": "2020-09-22T09:17:06-06:00"
    },
    "Labels": null,
    "GUID": "9719d92a-df1f-4a05-885a-ad10915d8b42",
    "Synced": false
  },
  {
    "ID": 69148521436807,
    "Name": "Test 2",
    "UID": 1,
    "GIDs": [],
    "Description": "dashboard 2",
    "Created": "2020-09-22T09:17:13.809070187-06:00",
    "Updated": "2020-09-22T09:17:13.809070187-06:00",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    },
    "Labels": null,
    "GUID": "2c55cf84-bb63-40cf-bf54-3bff8c8d7fb6",
    "Synced": false
  }
]
```

### Getting all dashboards owned by a user
To get dashboards explicitly owned by a user, issue a GET request on `/api/users/{uid}/dashboards`, replacing {uid} with the desired user ID. The webserver will return ONLY those dashboards specifically owned by that UID.  This WILL NOT include dashboards the user has access to through group memberships.

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
    "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```

### Getting all of a group's dashboards
To get all dashboards that a specific group has access to, issue a GET request to `/api/groups/{gid}/dashboards`, replacing {gid} with the desired group ID. The server will return any dashboards shared with that group.  This deviates a little from normal Unix permissions in that a dashboard can be shared with multiple groups (so groups don't really OWN a dashboard, but rather dashboards are kind of members of a group).

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
    "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  },
  {
    "ID": 2,
    "Name": "test2",
    "UID": 3,
    "GIDs": [
      2
    ],
    "Description": "test2 description",
    "Created": "2016-12-18T23:28:08.250051418Z",
    "GUID": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```


### Admin-only: List ALL dashboards of ALL users
To get *all* dashboards on the system, an admin user may issue a GET request to `/api/dashboards/all`. If this request is issued by a non-admin user it should return all dashboards to which they have access (equivalent to a GET on `/api/dashboards`)

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
    "GUID": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  },
  {
    "ID": 2,
    "Name": "test2",
    "UID": 3,
    "GIDs": [],
    "Description": "test2 description",
    "Created": "2016-12-18T23:28:08.250051418Z",
    "GUID": "55bc7236-39e4-11e9-94e9-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```

## Getting a specific dashboard
To fetch a particular ID, issue a GET request to `/api/dashboards/{id}`, replacing `{id}` with the dashboard's ID, e.g. `/api/dashboards/69148521436807`.

It is also possible to fetch a specific dashboard by GUID rather than ID: `GET /api/dashboards/d28b6887-ad55-479e-8af3-0cbcbd5084b1`.


## Updating a Dashboard
Updating a dashboard (to change the data or alter the name, description, GID list, etc) is done by issuing a PUT request to `/api/dashboards/{id}` where {id} is the unique identifier for a given dashboard.

In this example a user (UID 3) wishes to add permission for group 1 to access a dashboard (ID 2). The user first fetches the dashboard:

```
GET /api/dashboards/2:
{
  "ID": 2,
  "Name": "test2",
  "UID": 3,
  "GIDs": [],
  "Description": "test2 description",
  "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
  "Created": "2016-12-18T23:28:08.250051418Z",
  "Data": {
    "searches": [
      {
        "alias": "Search 1",
        "timeframe": {},
        "query": "tag=* count\n",
        "searchID": 4780372388
      }
    ],
    "tiles": [
      {
        "title": "count",
        "renderer": "text",
        "span": {
          "col": 4,
          "row": 4,
          "x": 0,
          "y": 0
        },
        "searchesIndex": 0,
        "id": 16007878262310,
        "rendererOptions": {}
      }
    ],
    "timeframe": {
      "durationString": "PT1H",
      "timeframe": "PT1H",
      "end": null,
      "start": null
    },
    "version": 2,
    "lastDataUpdate": "2020-09-22T09:17:06-06:00"
  }
}
```

The user has now retrieved the target dashboard, makes the modifications, and posts a PUT request:
```
WEB PUT /api/dashboards/2:
{
  "ID": 2,
  "Name": "marketoverview",
  "UID": 3,
  "GIDs": [
    3
  ],
  "Description": "marketing group dashboard",
  "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
  "Created": "2016-12-18T23:28:08.250051418Z",
  "Data": {
    "searches": [
      {
        "alias": "Search 1",
        "timeframe": {},
        "query": "tag=* count\n",
        "searchID": 4780372388
      }
    ],
    "tiles": [
      {
        "title": "count",
        "renderer": "text",
        "span": {
          "col": 4,
          "row": 4,
          "x": 0,
          "y": 0
        },
        "searchesIndex": 0,
        "id": 16007878262310,
        "rendererOptions": {}
      }
    ],
    "timeframe": {
      "durationString": "PT1H",
      "timeframe": "PT1H",
      "end": null,
      "start": null
    },
    "version": 2,
    "lastDataUpdate": "2020-09-22T09:17:06-06:00"
  }
}

```

The server will respond to update requests with the updated dashboard structure.

To be safe, take care to send back all fields which were present in the original fetch, even if unchanged. For instance, although this update did not modify the `Description` field, we include it in the update request because the webserver cannot distinguish between an *un-set* field and an *empty* field.

Note: The GUID may be used in place of the dashboard ID if desired.

## Deleting a dashboard
To remove a dashboard issue a request with the DELETE method on the url `/api/dashboards/{id}` where {id} is the numeric ID of the dashboard.
