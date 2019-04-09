# Autoextractor API

The extractor web API provides methods for accessing, modifying, adding, and deleting autoextractor definitions.  All users can retrieve the list of auto extractors and their definitions.  However, only admins may add, modify, delete, or sync auto extractor definitions.  As of version 3.0.2 the auto extractor API is available only when operating in non-distributed mode.  If you are operating a Gravwell cluster in a distributed frontend configuration the autoextractors must be pre-installed and managed manually.  For more information about autoextractors and their configuration, see the [Auto-Extractors](/#!configuration/autoextractors.md) section.

## Definition Structure

Auto-Extractors are defined using the following JSON structure (the `module` and `tag` parameters must be populated):

```
{
	"tag": string,
	"name": string,
	"desc": string,
	"module": string,
	"params": string,
	"arts": string
}
```

## Listing

Listing autoextractors is done by performing a GET on `/api/autoextractors`.  The webserver will return a list of JSON structures that represent the set of installed auto-extractors on the webserver.  An example response is:

```
[
	{
		"name": "testCSVExt",
		"module": "csv",
		"params": "ts, app, id, uuid, src, srcport, dst, dstport, data, name, country, city, hash",
		"tag": "test1"
	},
	{
		"name": "testFieldsExt",
		"desc": "testing extraction",
		"module": "fields",
		"params": "src, dst, extra",
		"tag": "test2"
	}
]
```

## Syncing

Any operation that changes the installed set of auto-extractors will also invoke a sync operation.  A sync operation causes the webserver to push changes to each of the attached indexers.  For example, if you are running a Gravwell cluster with 10 indexers, after adding a new auto-extractor the webserver will push the configured auto-extractor set to each of the indexers.  However, failures can happen whether it be due to network connectivity issues, or because an indexer is down when the change occurs.  The Sync API allows for manually invoking a sync operation so that all attached indexers are forced into the same state.  When managing a large Gravwell cluster the Sync operation can be used to intialize auto-extractor definitions after a new Indexer is brought online or restored.

A sync is performed by issuing a PUT request to `/api/autoextractors/sync`.  The webserver will ruturn a 200 response on success and non-200 on failure.  In the event of a partial success (not all indexers were successfully synced) a warning structure is returned in the body of the response.  The warning structure is a list of indexer names and errors, here is an example response when an indexer is down:

```
[
	{
		"Name": "net:172.17.0.4:9404",
		"Err": "is disconnected"
	},
	{
		"Name": "net:172.17.0.5:9404",
		"Err": "is disconnected"
	}
]
```

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

If an error occurs when adding an auto-extractor the webserver will return a list of errors.

## Upating

Updating an autoextractor is performed by issuing a PUT request to `/api/autoextractors` with a valid definition JSON structure in the request boyd.  The structure must be valid and there must be an existing auto-extractor that is assigned to the same tag.  The tag associated with an updated auto-extractor cannot be changed via the update API.  To change the tag associated with an existing auto-extractor, the definition must be deleted then added again.  The data structure is identical to the add API.  If the definition is invalid a non-200 response with an error message in the body is returned.  If the structure is valid but an error occurs in distributing the updated definition a list of errors is returned in the body.

## Testing Extractor Syntax

Before adding or updating an autoextractor, it may be useful to validate the syntax. Doing a POST request to `/api/autoextractors/test` will validate the request.If there is a problem with the definition, an error will be returned:

```
{"Error":"asdf is not a supported engine"}
```

When adding a new auto-extractor, it is important that the new extractor does not conflict with an existing extraction on the same tag. When updating an existing extraction, this is not a concern. If an extraction already exists for the specified tag, the test API will set the 'TagExists' field in the returned structure:

```
{"TagExists":true,"FileExists":true,"Error":""}
```

If `TagExists` is true, it should be treated as an error if you intend to create a new extractor, and ignored if updating an existing extractor.

The `FileExists` flag indicates that the proposed extraction would overwrite an existing extraction on disk; typically it will only be set when 'TagExists' is set. It should be treated as an error when creating a new extractor and ignored when updating.

## Deleting

Deleting an existing auto-extractor is performed by issuing a DELETE request to `/api/autoextractors/{id}` where id is the tag associated with the auto-extractor.  For example, to delete the auto-extractor associated with the tag "syslog" the request would go to `/api/autoextractors/syslog`.  If the auto-extractor does not exist or there is an error removing it, the webserver will respond with a non-200 response and error in the response body.  If an error occurs when distributing the deletion to indexers there will be a 200 response and a list of warnings.

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
