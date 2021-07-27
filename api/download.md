# Downloading search results

Each renderer supports downloading search results in a variety of formats.  For example, if a query is rendered via a table we can directly download the results as csv, json, or the native lookup format.  Performing a download may require an extra step if the download request is coming from a browser.  Clients must be authenticated against the webserver with a valid JWT and the user must have access to the search in order to download it.  Download links are guarded using the standard JWT token scheme as all other api URLs.

## Downloading Search Results
Downloading the search with id 150460229 in the form of a CSV

```
GET /api/searchctrl/150460229/download/csv
```
## Accessing a download via temporary JWT token

Because web browsers cannot perform file downloads via Javascript (and therefor cannot deliver the JWT via a request header) search result downloads can also authenticate using a temporary JWT stored in a cookie named "token".  Clients **SHOULD NOT** store the normal JWT token in a cookie, instead clients should ask for a temporary JWT that can **ONLY** be used to download search results.  These temporary JWT tokens only grant access to download API URLs and cannot be used on any other APIs (as a header or cookie).  Temporary JWT tokens are valid for 3 seconds.

Non-browser based clients that can deliver additional headers as part of a download request can ignore the temporary JWT token and cookie business and just send the normal JWT token in the appropriate request header.  If the server sees a valid JWT token in the standard header it will use that for authentication.


### Acquiring JWT temporary token

```
GET /api/login/tmptoken:
{
        "LoginStatus": true,
	"JWT": "7b22616c676f223a36323733382c22747970223a226a7774227d.7b22756964223a312c2265787069726573223a22323031382d30362d32305431333a32343a32382e393436393338312d30363a3030222c22696174223a5b3138312c36312c35382c3138352c3139322c3135392c3233302c3130372c37312c33322c3130382c33332c3134362c3138362c37392c35372c35382c33342c3135362c36342c33322c3234372c39352c35352c3138342c3235322c3135312c39382c34322c31382c35312c375d7d.cec81d84a3c96e8fd6961c1113a026eba08344d06d518e65f28bd6b92655fb6a22433fd6e42b51d62d45f8ed2a1665f3f951019b982251ebc614e8be5e4fdb6e"
}

```

Browsers should issue the GET request the JWT value in a cookie named "token" with an expiration of 3 seconds.  The cookie is not respected for any APIs but downloading search results.
