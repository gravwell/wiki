## Require

The require module is a simple filter that examines entries in the pipeline and drops any entries that do not have specific enumerated fields.  The example use case is second level filtering when an upstream module is extracting metadata but may not always populate all names desired.  Require allows a search to ensure that only entries that actually have all extracted feature sets make it through the module.

### Example Usage

The following search takes packet entries and eliminates any which do not have a "SrcPort" enumerated value set, then counts how many times each source port appeared. This has the effect of eliminating all packets which are not TCP traffic:

```
tag=pcap packet tcp.SrcPort | require SrcPort | count by SrcPort | table SrcPort count
```