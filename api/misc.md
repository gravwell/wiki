# Miscellaneous APIs

Some APIs don't fit nicely into the main categories. They are listed here.

## Tag List

The webserver maintains a list of all tags known to the indexers. This list can be fetched with a GET request on `/api/tags`. This will return a list of tags:

```
["default", "gravwell", "pcap", "windows"]
```
