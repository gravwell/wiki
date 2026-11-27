# Base64

The `base64` module encodes or decodes base64 data. Web applications often use base64 to carry binary data in text, such as parameters in an HTTP request.

With no arguments, `base64` operates on the entry data. Give a single enumerated value name to operate on that value instead. The result replaces the entry data unless you name an output enumerated value with `-t`; the source enumerated value itself is never modified.

If a value cannot be decoded, the entry passes through unchanged and no output is written. There is no partial decode. Use `-s` to drop those entries instead.

## Supported Options

* `-d`: Decode rather than encode.
* `-raw`: Use unpadded base64, with no trailing `=` characters. When decoding, padded input fails.
* `-s`: Strict mode. Drop entries that cannot be decoded, and entries where the source value is missing or empty.
* `-t <name>`: Write the result to the named enumerated value instead of the entry data.
* `-u`: Use the URL-safe alphabet instead of the standard one. Can be combined with `-d` and `-raw`. See [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648#section-5) for more information.

## Example

This example uses regex to extract a base64-encoded `id` parameter from HTTP GET requests. The base64 module decodes `id` into `decodedid`, then grep keeps only entries where the decoded value contains "admin".

```gravwell
tag=apache regex "GET \S+\?id=(?P<id>[A-Za-z0-9+/=]+)" | base64 -d -t decodedid id | grep -e decodedid admin | table decodedid
```
