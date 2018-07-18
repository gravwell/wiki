## Require

The require module is a simple filter that examines entries in the pipeline and drops any entries that do not have specific enumerated fields.  The example use case is second level filtering when an upstream module is extracting metadata but may not always populate all names desired.  Require allows a search to ensure that only entries that actually have all extracted feature sets make it through the module.

### Supported Options

* `-v `: The “-v” option inverts the requirement logic, essentially saying "drop all entries that have any of these enumerated values."  Inverting the requirement module can be useful when upstream modules may or may not extract some field, and you only want to see entries that did not have the field.

### Example Usage

The following search takes packet entries and eliminates any which do not have a "SrcPort" enumerated value set, then counts how many times each source port appeared. This has the effect of eliminating all packets which are not TCP traffic:

```
tag=pcap packet tcp.SrcPort | require SrcPort | count by SrcPort | table SrcPort count
```

The Following search looks for DNS requests which do not have a successful IPv4 resolution by dropping all entries where the A field is present in an enumerated value:

```
tag=dns json Question.Hdr.Name Question.A | require -v A | count by Name | table Name count
```
