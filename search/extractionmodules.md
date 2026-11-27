# Search Extraction Modules

Gravwell is a structure-on-read data lake, and extraction modules add that structure. You don't need to know the shape of your data when you collect it: ingest it raw and extract fields at search time.

Extraction modules are often the first module in a search. Netflow data, for example, sits on disk in its native binary format, so searches over it typically start with the `netflow` module to extract fields and do basic filtering. You *could* filter with a module like `grep` before `netflow`, but that is unlikely to be effective.

Extracted fields become enumerated values that "ride along" with the entry for later modules to use. For example, `netflow Src | subnet Src /16 as srcsub | grep -e srcsub "10.1.0.0"` extracts the source IP from the netflow record, extracts its /16 subnet as `srcsub`, and filters on that subnet with grep. The grep module has access to the raw data as well as the `Src` and `srcsub` enumerated values.

## Query Accelerators

Extraction modules can use [query accelerators](/configuration/accelerators) (full text indexing, JSON indexing, etc.) when you filter with them. For example, `netflow Src Dst Port==22` can use a properly configured accelerator to dramatically reduce search time, because not every record needs to be evaluated. Some filtering modules, such as [words](words/words), can also engage acceleration by passing hints to the accelerator.

Not every module works with every accelerator. If the `netflow` tag uses the netflow accelerator, the [words](words/words) module cannot engage it: the netflow accelerator indexes structure in binary data, while words is designed for fulltext acceleration.

Gravwell examines query parameters and engages acceleration whenever possible, with some caveats:

1. Acceleration that uses specific structure (such as netflow, ipfix, packet, fields, regex, etc.) requires that the query parameters match exactly.

For example, if a tag's accelerator uses a regular expression to apply structure, your query must use the exact same regular expression to benefit. If the accelerator's expression is `(?P<ip>\S+)\s(?P<description>(\S+\s+)+)` and your query is `regex "(?P<ip>\d+\.\d+\.\d+\.\d+)\s" ip=="1.2.3.4"`, acceleration is not engaged because Gravwell cannot guarantee the query's expression is a subset of the accelerator's.

2. Acceleration may use subsets of your query.

Suppose the tag `test` uses fulltext acceleration and you run `tag=test grep "this that the other"`. grep is a brute-force substring match, so the first and last words of the pattern may not fall on word boundaries in the data. The interior words do, and Gravwell uses them to accelerate the query, as if you had run `tag=test words that the | grep "this that the other"`.

3. Acceleration does not depend on module order.

Modules that can accelerate a query send hints no matter where they appear. In `tag=test grep foo* | regex "my name is (?P<name>\S+)" | words foobar`, the `words` module hints that it wants the word "foobar", and a fulltext accelerator uses that hint. Very few entries may enter the pipeline even though words comes last.

4. Acceleration is decided shard by shard.

Accelerator configuration doesn't have to be consistent across time. You can set up acceleration, ingest data, and later change the configuration without re-indexing. At query time, each shard checks the acceleration hints against its own configuration, so a query may be accelerated differently, or not at all, as it moves through historical data. A query may be fast over some time ranges and slower over others.

This also means Gravwell still returns accurate results if an index is lost. If a filesystem corruption or hardware failure destroys an index but the underlying data is intact, Gravwell detects that it cannot use the index and ignores it. Queries still complete, just more slowly.

5. Acceleration operates on positive matches.

Gravwell is not a bit field index, so it only accelerates direct matches. `tag=syslog syslog Hostname=="foobar"` can use the accelerator, but `tag=syslog syslog Severity < 4` cannot. Negative filters are not accelerated either: `tag=syslog syslog Hostname != "foobar"` does not use the accelerator.

## Universal Flags

Several modules share these flags. Most modules that accept them use them as described here, but a few reuse the same letter for something else (for example, `-e` sets the output name in `printf` and names the enumerated value holding the cutset in the trim modules, and `-s` sets the separator in `concat`). Check each module's page.

* `-e <source name>` reads input from the named enumerated value instead of the entry's DATA field. This is useful in modules like [json](json/json) when the JSON was extracted from a larger record. For example, this search reads JSON fields from the payloads of HTTP packets: `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-r <resource name>` names a [resource](/resources/resources) that holds extra data for the module, such as the GeoIP database used by [geoip](geoip/geoip).
* `-v` inverts the module's pass/drop logic. For example, [grep](grep/grep) normally passes entries that match a pattern and drops the rest; with `-v` it drops matching entries and passes the rest.
* `-s` enables "strict" mode. Where a module normally passes an entry if any one of several conditions holds, `-s` passes it only if *all* hold. For example, [require](require/require) normally passes an entry that has any of the listed enumerated values; with `-s` the entry must have all of them. In extraction modules such as [json](json/json) or [kv](kv/kv), `-s` means every extraction must succeed.
* `-p` enables "permissive" mode. Entries that don't match the module's pattern pass through instead of being dropped; filters still apply to entries that do match. [regex](regex/regex) and [grok](grok/grok) support it.
* `-or` applies to modules that normally require every condition to match: with `-or`, the entry passes if *any* one matches. [ip](ip/ip), [ipexist](ipexist/ipexist), and [words](words/words) support it.

## Universal Enumerated Values

These enumerated values are available on every entry. They are names for properties of the entry itself, so they never need to be extracted:

* `SRC`: the source of the entry data.
* `TAG`: the tag attached to the entry.
* `TIMESTAMP`: the timestamp of the entry.
* `DATA`: the entry's data.
* `NOW`: the current time.

Use them like any other enumerated value, e.g. `table foo bar DATA NOW`.

## Search module documentation

```{note}
The modules listed here have a primary function of extraction. They may also perform filtering and/or processing.
```

```{toctree}
---
maxdepth: 1
hidden: true
---
ax <ax/ax>
canbus <canbus/canbus>
cef <cef/cef>
csv <csv/csv>
dump <dump/dump>
fields <fields/fields>
grok <grok/grok>
intrinsic <intrinsic/intrinsic>
ip <ip/ip>
ipfix <ipfix/ipfix>
j1939 <j1939/j1939>
json <json/json>
kv <kv/kv>
netflow <netflow/netflow>
packet <packet/packet>
packetlayer <packetlayer/packetlayer>
path <path/path>
regex <regex/regex>
slice <slice/slice>
strings <strings/strings>
subnet <subnet/subnet>
syslog <syslog/syslog>
winlog <winlog/winlog>
xml <xml/xml>
```

* [ax](ax/ax) - automatically extract fields from entries.
* [canbus](canbus/canbus) - decode CANBUS data.
* [cef](cef/cef) - decode CEF data.
* [csv](csv/csv) - extract fields from CSV data.
* [dump](dump/dump) - dump entries from a resource into the pipeline.
* [fields](fields/fields) - extract data from entries using arbitrary field separators.
* [namedfields](#namedfields-module) - like fields, with field indexes looked up by name from a resource.
* [grok](grok/grok) - extract data from complicated text structures using pre-defined regular expressions.
* [intrinsic](intrinsic/intrinsic) - populate entries with enumerated values that were created at ingest time.
* [ip](ip/ip) - convert & filter IP addresses.
* [ipfix](ipfix/ipfix) - extract data from IPFIX records.
* [j1939](j1939/j1939) - parse J1939 data.
* [json](json/json) - extract elements from JSON data.
* [kv](kv/kv) - parse key-value data.
* [netflow](netflow/netflow) - parse Netflow records.
* [packet](packet/packet) - parse raw packets.
* [packetlayer](packetlayer/packetlayer) - parse portions of a packet.
* [path](path/path) - extract portions of pathnames.
* [regex](regex/regex) - match and extract data using regular expressions.
* [slice](slice/slice) - low-level binary parsing & extraction.
* [strings](strings/strings) - find strings from binary data.
* [subnet](subnet/subnet) - extract & filter based on IP subnets.
* [syslog](syslog/syslog) - parse and extract syslog entries.
* [winlog](winlog/winlog) - parse Windows logs.
* [xml](xml/xml) - parse XML data.
