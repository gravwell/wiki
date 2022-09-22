## JSON Array Split Preprocessor

This preprocessor can split an array in a JSON object into individual entries. For example, given an entry which contains an array of names, the preprocessor will instead emit one entry for each name. Thus this:

```
{"IP": "10.10.4.2", "Users": ["bob", "alice"]}
```

Becomes two entries, one containing "bob" and one containing "alice".

The JSON Array Split preprocessor Type is `jsonarraysplit`.

### Supported Options

* `Extraction` (string): specifies the JSON field containing a struct which should be split, e.g. `Extraction=Users`, `Extraction=foo.bar`. If you do not set `Extraction`, the preprocessor will attempt to treat the entire object as an array to split.
* `Drop-Misses` (boolean, optional): If set to true, the preprocessor will drop entries for which it was unable to extract the requested field. By default, these entries are passed along.
* `Force-JSON-Object` (boolean, optional): By default, the preprocessor will emit entries with each containing one item in the list and nothing else; thus extracting `foo` from `{"foo": ["a", "b"]}` would result in two entries containing "a" and "b" respectively. If this option is set, that same entry would result in two entries containing `{"foo": "a"}` and `{"foo": "b"}`.
* `Additional-Fields` (string, optional): A comma delimited list of additional fields outside the array to be split that will be extracted and included in each entry, e.g. `Additional-Fields="foo,bar,foo.bar.baz"`.

### Common Use Cases

Many data providers may pack multiple events into a single entry, which can degrade the atomic nature of an event and increase the complexity of analysis.  Splitting a single message that contains multiple events into individual entries can simplify working with the events.


### Example: Splitting Multiple Messages In a Single Record

To split entries which consist of JSON records with an array named "Alerts":

```
[preprocessor "json"]
	Type=jsonarraysplit
	Extraction=Alerts
	Force-JSON-Object=true
```

Input data:

```
{ "Alerts": [ "alert1", "alert2" ] }
```

Output:

```
{ "Alerts": "alert1" }
```

```
{ "Alerts": "alert2" }
```

### Example: Splitting a Top-Level Array

Sometimes the entire entry is an array:

```
[ {"foo": "bar"}, {"x": "y"} ]
```

To split this, use the following definition:

```
[preprocessor "json"]
	Type=jsonarraysplit
```

Leaving the Extraction parameter un-set tells the module to treat the entire entry as an array, giving the following two output entries:

```
{"foo": "bar"}
```

```
{"x": "y"}
```


