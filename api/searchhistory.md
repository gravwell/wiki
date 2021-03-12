# Search History

REST API located at `/api/searchhistory`

The search history API is used to pull back a list of searches that a user, group, or combination thereof have launched. The results provide basic info like who launched the search, what group owns it, when it was launched and the two strings representing the search (what the user actually typed, and what the backend processed).

## Basic API Overview

The basic action here is to perform a GET on `/api/searchhistory/{cmd}/{id}` with `cmd` representing what set of searches you want and `id` representing the ID related to that search. For example if you wanted all searches owned by the user with the UID 1 you would perform a GET on `/api/searchhistory/user/1` while if you wanted all searches owned by group with GID 4 you would perform a GET on `/api/searchhistory/group/4`. Simply performing a GET on `/api/searchhistory` will return the history for the current user.

You can request ALL searches a specific UID has access to using the "all" cmd. This means give me all searches owned by the UID as well as all searches that are owned by groups he is a member of. For example a GET on `/api/searchhistory/all/1` may return searches owned by group with GIDs 1, 2, 3, and 4 if the user with UID 1 is a member of those groups. The returned results will be a list of SearchLog structures in JSON format.

Example JSON:
```
[
        {
                "UID": 1,
                "GID": 2,
                "UserQuery": "grep stuff",
                "EffectiveQuery": "grep stuff | text",
                "Launched": "2015-12-30T23:30:23.298945825-07:00"
        },
        {
                "UID": 1,
                "GID": 2,
                "UserQuery": "grep stuff | grep things | grep that | grep this | sort by time",
                "EffectiveQuery": "grep stuff | grep things | grep that | grep this | sort by time | text",
                "Launched": "2015-12-30T23:31:08.237520376-07:00"
        }
]
```

### Admin Queries

A GET request to `/api/searchhistory?admin=true` as an admin user will return all searches by all users.

### Refined Search History

As a special case to the GET method on `/api/searchhistory`, you can provide a simple substring search term by setting a "refine" value. For example, a GET request to `/api/searchhistory?refine=foo` will return all searches for the current user containing the term "foo" anywhere in the search.
