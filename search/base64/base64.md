# Base64

The base64 module is useful for encoding and decoding base64 values.  Most web transport systems use base64 in order to ship binary data via HTTP (a text based protocol). For example, one might use `base64` to extract encoded values from a HTTP PUT request.  The base64 module makes a best effort when decoding and will decode as much data as it can.  The strict flag allows for enforcing a clean decode on all values, dropping any entries that cannot be entirely decoded.

## Supported Options

* `-d`: Decode rather than encode
* `-raw`: Assume RAW base64 encoding/decoding, excluding any padding = characters
* `-s`: Enforce strict mode, drop entry if entire field cannot be decoded.
* `-t <arg>`: Assign encoded or decoded values to an enumerated value
* `-u`: Use URL-safe encoding variant instead of the default. May be used in combination with the `-d` and `-raw` flags. See [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648#section-3.5) for more information.

## Example

This example first uses regex to look for HTTP GET requests with a base64-encoded "id" parameter and extract the parameter value. The base64 module then decodes the value in the enumerated value "id" and passes it along to grep, which checks if the decoded value matches "admin".

```
grep gravwell | regex "(?P<ts>[^\.]+)([^\s]+)\sgravwell\sapache2\.access:\s(?P<ip>[0-9\.]+)\s.+\s\"GET\s.+\?id=(?P<id>[0-9a-f]+)" | base64 -t decodedid -d id | grep -e decodedid “admin” | table decodedid
```
