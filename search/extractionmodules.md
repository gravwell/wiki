# Search Extraction Modules

Gravwell is a structure-on-read data lake, and these are the modules that add that structure. This is where the power and flexibility of the platform can be seen, as information about data doesn't need to be known before collection. Instead, we can ingest the raw data and perform the extractions at search time which gives tremendous flexibility to search operations.

Extraction modules are often the first module in a search. Netflow data, as an example, sits in it's native binary format on disk and any searches looking to operate on that data will be using the netflow extraction module as a first module in the analysis pipeline to do extraction and basic filtering. While you *could* do filtering before the `netflow` extraction module using modules like `grep`, that is unlikely to be effective.

Search extraction modules extract fields from data which "ride along" with the raw data through the rest of the search pipeline. Any extracted data/fields/properties become Enumerated Values which exist alongside an entry for use by following modules in a pipeline. For example, a search using the module sequence `netflow Src | subnet Src /16 as srcsub | grep -e srcsub "10.1.0.0"` will extract the Src IP address from the raw netflow record, extract the /16 subnet from the Src IP address, and then filter based on the extracted subnet using grep. In this example, the raw data is available to the grep module as well as the "Src" and "srcsub" Enumerated Values.

## Query Accelerators

Extraction modules can make use of [query accelerators](/configuration/accelerators) (like full text indexing, JSON indexing, etc) when filtering is used with a given module. For example, using the module `netflow Src Dst Port==22` can use a properly configured accelerator to dramatically reduce search time because not all records need to be evaluated.  Some filtering modules (such as [words](words/words.md)) can also invoke query acceleration by passing hints to the underlying accelerator.

Not all query modules are compatible with all query accelerators.  For example, the `netflow` tag is configured to be accelerated using the netflow accelerator, but the [words](words/words.md) module will not be able to invoke netflow query acceleration. This is because the netflow accelerator is expecting to operate on binary data and apply a specific structure to data during indexing where the words module is designed to operate on a fulltext acceleration.

Gravwell will intelligently examine query parameters and invoke the acceleration system whenever possible, but there are some caveats to be aware of:

1. Query acceleration that uses specific structure (such as netflow, ipfix, packet, fields, regex, etc) requires that the query parameters match exactly.

For example, if you customize your accelerator for a given tag so that it uses regular expressions to apply structure, then you must use the exact regular expression in your query in order to benefit from the accelerator.  This means that if your accelerator regular expression is `(?P<ip>\S+)\s(?P<description>(\S+\s+)+` and your query is `regex "(?P<ip>\d+\.\d+\.\d+\.\d+)\s" ip=="1.2.3.4"` the system will not engage query acceleration because it cannot guarantee that the query regular expression is a direct subset of the accelerator regular expression.

2. Query acceleration may accelerate on subsets of your query.

For example, let's assume the tag `test` is using fulltext acceleration and we are executing the query `tag=test grep "this that the other"`.  The query is using grep as a brute force sub-string match which means that the boundaries of the match string are not necessarily word boundaries.  However, the query system is smart enough to know that some internal parts of the matched string are words and it will use them to accelerate the query.  The query will accelerate the search as if you had run `tag=test words this the | grep "this that the other"`.

3. Query acceleration is not order of operation dependent.

Modules which can accelerate queries will hint about their ability to accelerate no matter where they are in the query string.  Consider the following query: `tag=test grep foo* | regex "my name is (?P<name>\S+)" | words foobar`.  The `words` module is capable of invoking query acceleration on a fulltext accelerator and will hint to the acceleration system that it wants the word "foobar" and the accelerator will use that hint.  You may notice that very few entries actually enter the pipeline even though words came much later in the query.

4. Query acceleration is on a shard-by-shard basis.

Gravwell does not require tag accelerators to be consistent across all time. You can setup acceleration, ingest some data, and then change that acceleration configuration without re-indexing data.  When you issue a query, the acceleration hints are handed to each shard of data across time and the compatibility of acceleration is checked on each shard.  Gravwell will automatically invoke acceleration wherever possible; this means that as your query moves over historical data, it may be engaging acceleration in different ways transparently.  You may notice that a query is fast on some sections of data and slower on others.  That is just the system engaging acceleration where it can.

A side affect of this per-shard acceleration calculation is that Gravwell can still query accurately in the event of catastrophic index corruption.  For example, if you suffer a filesystem corruption or hardware failure where an index is entirely corrupted but the underlying data is still intact, Gravwell will determine that it cannot use the index and ignore it.  Your queries will still complete albeit slower.

5. Query acceleration operates on positive matches.

Gravwell is not a bit field index which means that it will only accelerate on direct matches.  For example `tag=syslog syslog Hostname=="foobar"` can invoke the accelerator, but `tag=syslog ProcID < 20` will not.  The accelerators also do not accelerate on the negative, meaning that `tag=syslog syslog Hostname != "foobar"` will not invoke the accelerator engine.

## Universal Flags

Some flags appear in several different search modules and have the same meaning throughout:

* `-e <source name>` specifies that the module should attempt to read its input data from the given enumerated value rather than from the entry's data field. This is useful in for modules like [json](json/json.md), where the JSON-encoded data may have been extracted from a larger data record, for example the following search will attempt to read JSON fields from the payloads of HTTP packets: `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-r <resource name>` specifies a resource in the [resources](/resources/resources) system. This is generally used to store additional data used by the module, such as a GeoIP mapping table used by the [geoip](geoip/geoip.md) module.
* `-v` indicates that the normal pass/drop logic should be inverted. For example the [grep](grep/grep.md) module normally passes entries which match a given pattern and drop those which do not match; specifying the `-v` flag will cause it to drop entries which match and pass those which do not.
* `-s` indicates a "strict" mode. If a module normally allows an entry to proceed down the pipeline if any one of several conditions are met, setting the strict flag means an entry will proceed only if *all* conditions are met. For example, the [require](require/require.md) module will normally pass an entry if it contains any one of the required enumerated values, but when the `-s` flag is used, it will only pass entries which contain *all* specified enumerated values.
* `-p` indicates "permissive" mode.  If a module normally drops entries when patterns and filters do not match, the permissive flag tells the module to let the module go through.  The [regex](regex/regex.md) and [grok](grok/grok.md) modules are good examples where the permissive flag can be valuable.

## Universal Enumerated Values

The following enumerated values are available for every entry. They're actually convenient names for properties of the raw entries themselves, but can be treated as enumerated value names.

* SRC -- the source of the entry data.
* TAG -- the tag attached to the entry.
* TIMESTAMP -- the timestamp of the entry.
* DATA -- the actual entry data.
* NOW -- the current time.

These can be used just like user-defined enumerated values, thus `table foo bar DATA NOW` is valid. They do not need to be explicitly *extracted* anywhere; they are always available.

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

* [ax](ax/ax.md) - automatically extract fields from entries.
* [canbus](canbus/canbus.md) - decode CANBUS data.
* [cef](cef/cef.md) - decode CEF data.
* [csv](csv/csv.md) - extract fields from CSV data.
* [dump](dump/dump.md) - dump entries from a resource into the pipeline.
* [fields](fields/fields.md) - extract data from entries using arbitrary field separators.
* [grok](grok/grok.md) - extract data from complicated text structures using pre-defined regular expressions.
* [ip](ip/ip.md) - convert & filter IP addresses.
* [ipfix](ipfix/ipfix.md) - extract data from IPFIX records.
* [j1939](j1939/j1939.md) - parse J1939 data.
* [json](json/json.md) - extract elements from JSON data.
* [kv](kv/kv.md) - parse key-value data.
* [netflow](netflow/netflow.md) - parse Netflow records.
* [packet](packet/packet.md) - parse raw packets.
* [packetlayer](packetlayer/packetlayer.md) - parse portions of a packet.
* [path](path/path.md) - extract portions of pathnames.
* [regex](regex/regex.md) - match and extract data using regular expressions.
* [slice](slice/slice.md) - low-level binary parsing & extraction.
* [strings](strings/strings.md) - find strings from binary data.
* [subnet](subnet/subnet.md) - extract & filter based on IP subnets.
* [syslog](syslog/syslog.md) - parse and extract syslog entries.
* [winlog](winlog/winlog.md) - parse Windows logs.
* [xml](xml/xml.md) - parse XML data.
