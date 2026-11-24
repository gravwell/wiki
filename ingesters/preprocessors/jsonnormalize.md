# JSON Normalize Preprocessor

The JSON normalize preprocessor repairs heavily escaped JSON entries and re-emits a single, clean JSON document. It handles two distinct flavors of heavily escaped JSON:

1. **Whole-record escaping**, where an entire JSON document has been serialized as a string one or more times, e.g. an entry whose data is literally `"{\"a\":1}"` instead of `{"a":1}`, or, in the case where an upstream system stripped the enclosing quotes, the invalid fragment `{\"a\":1}`.
2. **Field-level escaping**, where an otherwise well-formed JSON document contains a field whose value is itself a JSON-encoded string, e.g. `{"user":"alice","payload":"{\"a\":1}"}`. This is common when a logging pipeline serializes a sub-object independently before embedding it in a parent document.

The JSON Normalize preprocessor Type is `jsonnormalize`.

## Supported Options

- `Max-Depth` (integer, optional): Bounds how many layers of string-escaping the processor will unwind, both when repairing a malformed, heavily escaped document and when recursively inlining JSON-encoded strings found as field values. Defaults to `8` if unset or set to `0`. This budgets escaping *layers*, not structural nesting -- walking down through an already-valid document's plain objects/arrays costs nothing against it, so a document nested many objects deep with a single escaped field at the bottom is handled the same as a shallow one, as long as that field is only escaped once.
- `Passthrough-Non-JSON` (boolean, optional): By default, entries which are not valid JSON and cannot be repaired into valid JSON by unescaping are dropped. Setting `Passthrough-Non-JSON` to true will instead pass such entries through unmodified. This also governs entries that decode as structurally valid JSON but contain invalid UTF-8 in a string value.
- `Pretty` (boolean, optional): By default, the normalized output is compact, single-line JSON. Setting `Pretty` to true indents the output for readability.

```{note}
Max-Depth counts escaping layers, not object/array nesting depth. CloudTrail, Kubernetes audit logs, and Windows event JSON routinely nest many objects deep; the default of 8 is almost always enough because it only needs to cover how many times a value was re-encoded as a string, not how deep the document's structure goes.
```

## Example: Repairing a Fully Escaped Document

To illustrate the use of this preprocessor, consider a situation where an upstream system serializes an entire JSON document as a string before handing it off, and in doing so strips the outer quotes that would normally make it valid JSON. Incoming logs look like this:

```
{\"action\":\"login\",\"user\":\"alice\",\"timestamp\":1234567890}
{\"action\":\"logout\",\"user\":\"bob\",\"timestamp\":1234567891}
```

We can apply a JSON normalize preprocessor to repair these entries:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=normalizer

[preprocessor "normalizer"]
        Type=jsonnormalize
```

With the above configuration, each entry is repaired back into valid JSON:

```
{"action":"login","user":"alice","timestamp":1234567890}
{"action":"logout","user":"bob","timestamp":1234567891}
```

## Example: Inlining a Field-Level Escaped JSON Field

The JSON normalize preprocessor also finds and inlines fields whose values are themselves JSON-encoded strings, rather than only repairing whole-record escaping. Consider logs where a sub-object has been serialized independently before being embedded in the parent document:

```
{"user":"alice","action":"login","payload":"{\"ip\":\"10.0.0.5\",\"success\":true}"}
{"user":"bob","action":"logout","payload":"{\"ip\":\"10.0.0.9\",\"success\":true}"}
```

Applying the preprocessor:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=normalizer

[preprocessor "normalizer"]
        Type=jsonnormalize
```

produces entries with `payload` inlined as a real nested object instead of an escaped string:

```
{"action":"login","payload":{"ip":"10.0.0.5","success":true},"user":"alice"}
{"action":"logout","payload":{"ip":"10.0.0.9","success":true},"user":"bob"}
```

## Example: Deeply Nested Documents

Because `Max-Depth` only budgets escaping layers and not structural nesting, documents from sources like CloudTrail or Kubernetes audit logs -- which are often a dozen or more objects deep -- are handled without needing to raise the default. Given a deeply nested entry with a single escaped field buried inside it:

```
{"eventSource":"s3.amazonaws.com","requestParameters":{"bucketName":"example","object":{"key":"file.txt","metadata":{"tags":{"custom":"{\"team\":\"platform\",\"owner\":\"alice\"}"}}}}}
```

We can apply a JSON normalize preprocessor to repair this entry:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=normalizer

[preprocessor "normalizer"]
        Type=jsonnormalize
```

The `custom` field is still found and inlined, and the surrounding structure is left untouched:

```
{"eventSource":"s3.amazonaws.com","requestParameters":{"bucketName":"example","object":{"key":"file.txt","metadata":{"tags":{"custom":{"owner":"alice","team":"platform"}}}}}}
```

## Example: Passing Through Unrecoverable Entries

By default, entries that cannot be repaired into valid JSON -- or that decode but contain invalid UTF-8 -- are dropped. Setting `Passthrough-Non-JSON=true` instead leaves such entries untouched, which is useful when a stream mixes JSON and non-JSON records:

```
{"action":"login","user":"alice"}
connection reset by peer
{"action":"logout","user":"bob"}
```

Preprocessor to keep non-JSON records:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=normalizer

[preprocessor "normalizer"]
        Type=jsonnormalize
        Passthrough-Non-JSON=true
```

With the above configuration, the two JSON entries are normalized as usual and the plain-text `connection reset by peer` line is passed through unmodified rather than being dropped. Without `Passthrough-Non-JSON` set, that line would simply be dropped.

## Example: Pretty-Printed Output

By default, output is compact. Setting `Pretty=true` indents the result, which can be useful when chaining `jsonnormalize` before a preprocessor or downstream tool that expects human-readable JSON:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=normalizer

[preprocessor "normalizer"]
        Type=jsonnormalize
        Pretty=true
```

Given the entry `{"user":"alice","payload":"{\"ip\":\"10.0.0.5\"}"}`, the output becomes:

```
{
  "payload": {
    "ip": "10.0.0.5"
  },
  "user": "alice"
}
```
