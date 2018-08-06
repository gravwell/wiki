# Require

The require module is a simple filter that examines entries in the pipeline and drops any entries that do not have specific enumerated fields.  The example use case is second level filtering when an upstream module is extracting metadata but may not always populate all names desired.  Require allows a search to ensure that only entries that actually have the desired feature sets make it through the module.

By default, given a list of enumerated values the require module will pass an entry down the pipeline if it contains **at least one** of those enumerated value names. This behavior can be modified using flags as shown below.

## Supported Options

* `-s`: The `-s` option specifies strict operation: *all* listed enumerated values must exist, not just one. Essentially, changes the module from a logical OR operation to a logical AND.
* `-v `: The `-v` option inverts the requirement logic, essentially saying "drop all entries that have any of these enumerated values." This flag implies the `-s` flag. Inverting the requirement module can be useful when upstream modules may or may not extract some field, and you only want to see entries that did not have the field.

## Example Usage

The following search takes packet entries and eliminates any which do not have a "SrcPort" enumerated value set, then counts how many times each source port appeared. This has the effect of eliminating all packets which are not TCP traffic:

```
tag=pcap packet tcp.SrcPort | require SrcPort | count by SrcPort | table SrcPort count
```

The following search looks for DNS requests which do not have a successful IPv4 resolution by dropping all entries where the A field is present in an enumerated value:

```
tag=dns json Question.Hdr.Name Question.A | require -v A | count by Name | table Name count
```

This search passes through any packet which has *either* a TCP source port or a UDP source port specified:

```
tag=pcap packet tcp.SrcPort as tsp udp.SrcPort as usp | require tsp usp | table tsp usp
```

This search drops any packet which does not have *both* an IPv4 source IP *and* a TCP source port:

```
tag=pcap packet tcp.SrcPort ipv4.SrcIP | require -s SrcPort SrcIP | table SrcPort SrcIP
```
