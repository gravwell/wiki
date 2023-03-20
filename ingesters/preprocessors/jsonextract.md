## JSON Extraction Preprocessor

The JSON extraction preprocessor can parse the contents of an entry as JSON, extract one or more fields from the JSON, and replace the entry contents with those fields. This is a useful way to simplify overly-complex messages into more concise entries containing only the information of interest.

If only a single field extraction is specified, the result will contain purely the contents of that field; if multiple fields are specified, the preprocessor will generate valid JSON containing those fields.

The JSON Extraction preprocessor Type is `jsonextract`.

### Supported Options

* `Extractions` (string, required): This specifies the field or fields (comma-separated) to be extracted from the JSON. Given an input of `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}`, you could specify `Extractions=foo`, `Extractions=foo,bar`, `Extractions=baz.frog,foo`, etc.
* `Force-JSON-Object` (boolean, optional): By default, if a single extraction is specified the preprocessor will replace the entry contents with the contents of that extension; thus selecting `Extraction=foo` will change an entry containing `{"foo":"a", "bar":2, "baz":{"frog": "womble"}}` to simply contain `a`. If this option is set, the preprocessor will always output a full JSON structure, e.g. `{"foo":"a"}`.
* `Drop-Misses` (boolean, optional): If set to true, the preprocessor will drop entries for which it was unable to extract the requested fields. By default, these entries are passed.
* `Strict-Extraction` (boolean, optional): By default, the preprocessor will pass along an entry if at least one of the extractions succeeds. If this parameter is set to true, it will require that all extractions succeed.

### Common Use Cases

Many data sources may provide additional metadata related to transport and/or storage that are not part of the actual log stream.  The jsonextract preprocessor can down-select fields to reduce storage costs.

### Example: Condensing JSON Data Records

```
[Preprocessor "json"]
	Type=jsonextract
	Extractions=IP,Alert.ID,Message
	Drop-Misses=false
```


