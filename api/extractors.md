# Autoextractor API

The extractor web API provides methods for accessing, modifying, adding, and deleting autoextractor definitions. For more information about autoextractors and their configuration, see the [Auto-Extractors](/#!configuration/autoextractors.md) section.

## Definition Structure

Auto-Extractors are defined using the following JSON structure:

```
{
	"tag": string,
	"name": string,
	"desc": string,
	"module": string,
	"params": string,
	"args": string,
	"uid": int32,
	"gids": []int32,
	"global": bool,
	"uuid": string,
	"synced": bool,
	"lastupdated": string
}
```

## Listing

Listing autoextractors is done by performing a GET on `/api/autoextractors`.  The webserver will return a list of JSON structures that represent the set of installed to which the current user has access.  An example response is:

```
[
    {
        "accelerated": "",
        "desc": "csv extractor",
        "gids": null,
        "global": false,
        "lastupdated": "2020-01-08T09:45:29.617936398-07:00",
        "module": "csv",
        "name": "csvextract",
        "params": "col1, col2, col3",
        "synced": false,
        "tag": "csv",
        "uid": 1,
        "uuid": "3674b59e-6064-4cf0-8023-4bf444e84625"
    },
    {
        "accelerated": "",
        "args": "-d \"\\t\"",
        "desc": "Bro conn logs",
        "gids": null,
        "global": true,
        "lastupdated": "2020-01-08T14:45:28.514723502-07:00",
        "module": "fields",
        "name": "bro-conn",
        "params": "ts, uid, orig, orig_port, resp, resp_port, proto, service, duration, orig_bytes, dest_bytes, conn_state, local_orig, local_resp, missed_bytes, history, orig_pkts, orig_ip_pkts, resp_pkts, resp_ip_bytes, tunnel_parents",
        "synced": false,
        "tag": "bro-conn",
        "uid": 1,
        "uuid": "bdb2c2f6-5d50-4222-8558-ecbe3a0822aa"
    }
]
```

Performing a GET request with the admin flag set (`/api/autoextractors?admin=true`) will return a list of *all* extractors on the system.

## Adding

Adding an autoextractor is performed by issuing a POST to `/api/autoextractors` with a valid definition JSON structure in the request body.  The structure must be valid and there cannot be an existing auto-extractor that is assigned to the tag.  An example POST JSON structure that adds a new auto-extractor:

```
{
	"name": "testCSVExt",
	"desc": "testing extractor",
	"module": "csv",
	"params": "a, b, c, src, dst, extra",
	"tag": "test4"
}
```

If an error occurs when adding an auto-extractor the webserver will return a list of errors. If successful, the server will respond with the UUID of the new extractor.

Note: There is no need to set the `uuid`, `uid`, `gids`, `synced`, or `lastupdated` fields when creating an extractor--these are automatically filled in. Only an admin user may set the `global` flag to true.

## Updating

Updating an autoextractor is performed by issuing a PUT request to `/api/autoextractors` with a valid definition JSON structure in the request body.  The structure must be valid and there must be an existing auto-extractor with the same UUID.  All non-modified fields should be included as originall returned by the server.  If the definition is invalid a non-200 response with an error message in the body is returned.  If the structure is valid but an error occurs in distributing the updated definition a list of errors is returned in the body.

## Testing Extractor Syntax

Before adding or updating an autoextractor, it may be useful to validate the syntax. Doing a POST request to `/api/autoextractors/test` will validate the request.If there is a problem with the definition, an error will be returned:

```
{"Error":"asdf is not a supported engine"}
```

When adding a new auto-extractor, it is important that the new extractor does not conflict with an existing extraction on the same tag. When updating an existing extraction, this is not a concern. If an extraction already exists for the specified tag, the test API will set the 'TagExists' field in the returned structure:

```
{"TagExists":true,"Error":""}
```

If `TagExists` is true, it should be treated as an error if you intend to create a new extractor, and ignored if updating an existing extractor.

## Uploading Files

Autoextractor definitions can be represented in a TOML format. This format is human-readable and can be a convenient way to distribute extractor definitions. An example is shown below:

```
[[extraction]]
	tag="bro-conn"
	name="bro-conn"
	desc="Bro conn logs"
	module="fields"
	args='-d "\t"'
	params="ts, uid, orig, orig_port, resp, resp_port, proto, service, duration, orig_bytes, dest_bytes, conn_state, local_orig, local_resp, missed_bytes, history, orig_pkts, orig_ip_pkts, resp_pkts, resp_ip_bytes, tunnel_parents"
```

Rather than parsing out this file to populate a JSON structure, this type of definition can be uploaded directly to the webserver via a multipart form sent in a POST request to `/api/autoextractors/upload`. The form should contain a file field named `extraction` which holds the contents of the extractor definition. The server will respond with a 200 response if the definition is valid and was successfully installed.

## Deleting

Deleting an existing auto-extractor is performed by issuing a DELETE request to `/api/autoextractors/{uuid}` where `uuid` is the UUID associated with the auto-extractor. If the auto-extractor does not exist or there is an error removing it, the webserver will respond with a non-200 response and error in the response body.

## Listing Modules

Auto-extractor definitions must specify a valid module.  The API to get a list of supported modules is performed by issuing a GET request to `/api/autoextractors/engines`.  The resulting set is a list of strings:

```
[
	"fields",
	"csv",
	"slice",
	"regex"
]
```
