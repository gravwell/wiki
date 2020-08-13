# Search Extraction Modules

Gravwell is a structure-on-read data lake, and these are the modules that add that structure. This is where the power and flexibility of the platform can be seen, as information about data doesn't need to be known before collection. Instead, we can ingest the raw data and perform the extractions at search time which gives tremendous flexibility to search operations.

Extraction modules are often the first module in a search. Netflow data, as an example, sits in it's native binary format on disk and any searches looking to operate on that data will be using the netflow extraction module as a first module in the analysis pipeline to do extraction and basic filtering. While you *could* do filtering before the `netflow` extraction module using modules like `grep`, that is unlikely to be effective.

Search extraction modules extract fields from data which "ride along" with the raw data through the rest of the search pipeline. Any extracted data/fields/properties become Enumerated Values which exist alongside an entry for use by following modules in a pipeline. For example, a search using the module sequence `netflow Src | subnet Src /16 as srcsub | grep -e srcsub "10.1.0.0"` will extract the Src IP address from the raw netflow record, extract the /16 subnet from the Src IP address, and then filter based on the extracted subnet using grep. In this example, the raw data is available to the grep module as well as the "Src" and "srcsub" Enumerated Values.

## Query Accelerators

Extraction modules can make use of [query accelerators](configuration/accelerators.md) (like full text indexing, JSON indexing, etc) when filtering is used with a given module. For example, using the module `netflow Src Dst Port==22` can use a properly configured accelerator to dramatically reduce search time because not all records need to be evaluated.

Some processing modules (such as [words](words/words.md)) directly perform filtering against the accelerated indexes.

## Universal Flags

Some flags appear in several different search modules and have the same meaning throughout:

* `-e <source name>` specifies that the module should attempt to read its input data from the given enumerated value rather than from the entry's data field. This is useful in for modules like [json](json/json.md), where the JSON-encoded data may have been extracted from a larger data record, for example the following search will attempt to read JSON fields from the payloads of HTTP packets: `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-r <resource name>` specifies a resource in the [resources](#!resources/resources.md) system. This is generally used to store additional data used by the module, such as a GeoIP mapping table used by the [geoip](geoip/geoip.md) module.
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

Note: The modules listed here have a primary function of extraction. They may also perform filtering and/or processing.

* [ax](ax/ax.md)
* [canbus](canbus/canbus.md)
* [cef](cef/cef.md)
* [csv](csv/csv.md)
* [dump](dump/dump.md)
* [fields](fields/fields.md)
* [grok](grok/grok.md)
* [ip](ip/ip.md)
* [ipfix](ipfix/ipfix.md)
* [j1939](j1939/j1939.md)
* [json](json/json.md)
* [kv](kv/kv.md)
* [namedfields](namedfields/namedfields.md)
* [netflow](netflow/netflow.md)
* [packet](packet/packet.md)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [slice](slice/slice.md)
* [strings](strings/strings.md)
* [subnet](subnet/subnet.md)
* [syslog](syslog/syslog.md)
* [winlog](winlog/winlog.md)
* [xml](xml/xml.md)
