# Tag Router Preprocessor

The tag router preprocessor can route entries to different tags based on the TAG field of the entry.

The tag router preprocessor Type is `plugin`.

## Supported Options

* `Route` (string, required): `Route` defines a mapping of TAG field value to tag, separated by a colon. For instance, `Route=tag1:tag2` will send all entries with TAG=tag1 to the “tag2” tag. Multiple Route parameters can be specified. Any unmatched tags will pass through unmodified. At least one Route definition is required.

* IPFilter (IP or CIDR specification, optional): Optionally a filter can be included to restrict tag changes to SRC. Either a single IP address or a properly formed CIDR specification. Only IPv4 is supported.

## Example: Inline Route Definitions

The snippet below shows part of a federator configuration that uses the tag router preprocessor.

```
[Listener "tagManipulator"]
    Cleartext-Bind = 0.0.0.0:4423
    Preprocessor="changetags"
    Tags=*

[Preprocessor "changetags"]
    Type=tagrouter
    Route=tag1:tag2
    Route=tag3:tag4:10.1.1.1
    Route=tag5:tag6:10.2.2.1/24
```

In the example, the following routes are given:

* `Route=tag1:tag2` will convert any logs passed through the Listener with tag1 to tag2
* `Route=tag3:tag4:10.1.1.1` will convert any logs passed through the Listener with tag3 and SRC 10.1.1.1 to tag2
* `Route=tag5:tag6:10.2.2.2/24` will convert any logs passed through the Listener with tag5 and SRC in the ip range of 10.2.2.1-10.2.2.254 to tag6
