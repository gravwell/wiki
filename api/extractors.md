# Autoextractor API

The extractor web API provides methods for accessing, modifying, adding, and deleting autoextractor definitions. For more information about autoextractors and their configuration, see the [Auto-Extractors](/#!configuration/autoextractors.md) section.

## Data Structure

Autoextractors contain the following fields:

* Tag: The tag which is being extracted.
* Name: A user-friendly name for the extractor.
* Desc: A more detailed description of the extractor.
* Module: The module to use ("csv", "fields", "regex", or "slice").
* Params: Extraction module parameters.
* Args: Extraction module arguments.
* Labels: An array of strings containing [labels](#!gui/labels/labels.md).
* UID: The numeric ID of the dashboard's owner.
* GIDs: An array of numeric group IDs with which this dashboard is shared.
* Global: A boolean, set to true if dashboard should be visible to all users (admin only).
* UUID: A unique ID for this particular extractor.
* LastUpdated: The time at which this extractor was most recently modified.

This is a Typescript description of the structure:

```
type RawAutoExtractorModule = 'csv' | 'fields' | 'regex' | 'slice';

interface RawAutoExtractor {
	UUID: RawUUID;

	UID: RawNumericID;
	GIDs: Array<RawNumericID> | null;

	Name: string;
	Desc: string;
	Labels: Array<string> | null;

	Global: boolean;
	LastUpdated: string; // Timestamp

	Tag: string;
	Module: RawAutoExtractorModule;
	Params: string;
	Args?: string;
}
```

## Listing

Issuing a GET on `/api/autoextractors` will return a list of JSON structures that represent the set of extractors to which the current user has access.  An example response is:

```
[
  {
    "Name": "Apache Combined Access Log",
    "Desc": "Apache Combined access logs using regex module.",
    "Module": "regex",
    "Params": "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<auth>\\S+) \\[(?P<date>[^\\]]+)\\] \\\"(?P<method>\\S+) (?P<url>.+) HTTP\\/(?P<version>\\S+)\\\" (?P<response>\\d+) (?P<bytes>\\d+) \\\"(?P<referrer>\\S+)\\\" \\\"(?P<useragent>.+)\\\"",
    "Tag": "apache",
    "Labels": [
      "apache"
    ],
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "UUID": "0e105901-92a7-4131-87bb-a00287d46f96",
    "LastUpdated": "2020-06-24T13:49:39.013266326-06:00"
  },
  {
    "Name": "vpcflow",
    "Desc": "VPC flow logs (TSV format)",
    "Module": "fields",
    "Params": "version, account_id, interface_id, srcaddr, dstaddr, srcport, dstport, protocol, packets, bytes, start, end, action, log_status",
    "Args": "-d \" \"",
    "Tag": "vpcflowraw",
    "Labels": null,
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "UUID": "7f80df6a-a2ce-42aa-b531-ac11c596f64a",
    "LastUpdated": "2020-05-29T15:00:41.883390284-06:00"
  }
]
```

Performing a GET request with the admin flag set (`/api/autoextractors?admin=true`) will return a list of *all* extractors on the system.

## Adding

Adding an autoextractor is performed by issuing a POST to `/api/autoextractors` with a valid definition in the request body.  The structure must be valid and the user cannot have an existing autoextractor defined for the same tag.  An example POST JSON structure that adds a new auto-extractor:

```
{
  "Tag": "foo",
  "Name": "my extractor",
  "Desc": "an extractor using the fields module",
  "Module": "fields",
  "Params": "version, account_id, interface_id, srcaddr, dstaddr, srcport, dstport, protocol, packets, bytes, start, end, action, log_status",
  "Args": "-d \" \"",
  "Labels": [
    "foo"
  ],
  "Global": false
}
```

If an error occurs when adding an auto-extractor the webserver will return a list of errors. If successful, the server will respond with the UUID of the new extractor.

Note: There is no need to set the `UUID`, `UID`, `GIDs`, or `LastUpdated` fields when creating an extractor--these are automatically filled in. Only an admin user may set the `Global` flag to true.

## Updating

Updating an autoextractor is performed by issuing a PUT request to `/api/autoextractors` with a valid definition JSON structure in the request body.  The structure must be valid and there must be an existing auto-extractor with the same UUID.  All non-modified fields should be included as originally returned by the server.  If the definition is invalid a non-200 response with an error message in the body is returned.  If the structure is valid but an error occurs in distributing the updated definition a list of errors is returned in the body.

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

## Downloading Files

You can download autoextractor definitions in TOML format by issuing a GET request to `/api/autoextractors/download`. For each definition you wish to download, add its UUID as a parameter to the URL; thus if you wish to download two extractors with UUIDs ad782c81-7a60-4d5f-acbf-83f70e68ecb0 and c7389f9b-ba52-4cbe-b883-621d577c6bcc, you would send a GET request to `/api/autoextractors/download?id=ad782c81-7a60-4d5f-acbf-83f70e68ecb0&id=c7389f9b-ba52-4cbe-b883-621d577c6bcc`.

If the current user has access to all the specified extractors, the server will respond with a downloadable file containing the definitions in TOML format. This file can be uploaded to another Gravwell system using the file upload API described above.

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
