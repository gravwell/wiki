## User Preferences
The user preferences API is used to store GUI preferences to persist across logins and between devices.

GET the /api/users/{id}/preferences and it will return a chunk of JSON.  Admins can request any users preferences, users can ONLY request their own sessions. If no preferences exist, return null.

GET on /api/users/preferences will return all users preferences

PUT the /api/users/{id}/preferences to update the user preferences. If no preferences exist, update with the provided JSON anyway. No POST will ever occur on this api. The payload of the PUT method will be the JSON blob.

GET and PUT are the only relevant methods. Each user should inherently have one and only one preferences JSON blob.

DELETE on /api/users/{id}/preferences will delete the preferences (if admin or canning your own)


Example returned JSON on a GET:
```json
{
     "foo": "bar",
	 "bar": "baz"
}
```

## Examples from the client
### Requesting preferences
```
WEB GET /api/users/5/preferences:
{
        "Name": "TestPref2",
        "Value": 57005,
        "DataData": "bW9yZSBpbXBvcnRhbnQgZGF0YQ=="
}
WEB GET /api/users/1/preferences:
{
        "Name": "TestPref1",
        "Val": 1234567890,
        "Data": "some important data",
        "Things": 3.1415
}
```
### Requesting ALL preferences (as admin)
```
WEB GET /api/users/preferences:
[]
```
### Pushing
```
WEB REQ PUT /api/users/1/preferences:
{
        "Name": "TestPref1",
        "Val": 1234567890,
        "Data": "some important data",
        "Things": 3.1415
}
```
```
WEB REQ PUT /api/users/5/preferences:
{
        "Name": "TestPref2",
        "Value": 57005,
        "DataData": "bW9yZSBpbXBvcnRhbnQgZGF0YQ=="
}
```
### Pushing to non existent user

Will get a 404 not found

### Pushing and pulling someone else's preferences as non-admin

Will get a 403 forbidden

### Deleting our preferences
```
WEB REQ DELETE /api/users/5/preferences:
```
### Attempting to delete someone else's preferences as non-admin

Will get a 403 forbidden
