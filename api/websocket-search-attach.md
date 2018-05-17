## Reattaching to an existing search
Attaching to a search that already exists is performed through the existing websocket using another subprotocol named "attach" which must be negotiated along with "search", "parse", and "PONG" when connecting.

#### Permissions
Users must either be a member of the group the search is assigned, the owner, or an admin in order to attach to searches.  If a user is not allowed to attach to a search, a permission denied message will be sent back over the socket.

### Requesting to attach to a search
Basically you just pass in the ID in and the server will respond with either an Error, or a new SubProto, RenderMod, RenderCmd, and SearchInfo.

The new Subproto is the name of the new subprotocol which will be created to service the newly attached search.  This means that you can fire up the websocket and then attach to lots of searches at the same time.

### Example transfers
Below is an example of the client asking for a list of searches, then attaching to one.

Asking for list of searches
```
WEB GET /api/searchctrl:
[
        {
                "ID": "004081950",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "560752652",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "608274427",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        }
]
```

Transaction on the "attach" subprotocol
```
SUBPROTO PUT attach:
{
        "ID": "560752652"
}
SUBPROTO GET attach:
{
        "Subproto": "attach5",
        "RendererMod": "text",
        "Info": {
                "ID": "560752652",
                "UID": 1,
                "UserQuery": "grep paco",
                "EffectiveQuery": "grep paco | text",
                "StartRange": "2017-01-14T06:08:32.024425042-07:00",
                "EndRange": "2017-01-14T16:08:32.024425042-07:00",
                "Started": "2017-01-14T16:08:32.025987218-07:00",
                "Finished": "2017-01-14T16:08:32.746482323-07:00",
                "StoreSize": 0,
                "IndexSize": 0
        }
}
```
After negotiating a new attached search, continue on using the same old APIs as you would when normally taking to a renderer.
```
SUBPROTO PUT attach5:
{
        "ID": 3
}
SUBPROTO GET attach5:
{
        "ID": 3,
        "EntryCount": 20000,
        "Finished": true
}
SUBPROTO PUT attach5:
{
        "ID": 16777218,
        "EntryRange": {
                "First": 0,
                "Last": 1024,
                "StartTS": "0001-01-01T00:00:00Z",
                "EndTS": "0001-01-01T00:00:00Z"
        }
}
```