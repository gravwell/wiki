# Search Processing Modules

Search modules are modules that operate on data in a passthrough mode, meaning that they perform some action (filter, modify, sort, etc.) and pass the entries down the pipeline. There can be many search modules and each operates in its own lightweight thread.  This means that if there are 10 modules in a search, the pipeline will spread out and use 10 threads.  Documentation for each module will indicate if the module causes distributed searches to collapse and/or sort.  Modules that collapse force the distributed pipelines to collapse, meaning that the module as well as all downstream modules execute on the frontend.  When starting a search it's best to put as many parallel modules as possible upstream of the first collapsing module, decreasing pressure on the communication pipe and allowing for greater parallelism.

Some modules significantly transform or collapse data, such as `count`. Pipeline modules following these collapsing modules may not be dealing with Raw data or previously created Enumerated Values. In short, things like `count by Src` turn data into the collapsed results with entries such as `10.0.0.1  3`. To illustrate, run the search `tag=* limit 10 | count by TAG | raw` to see the raw output from the count module or `tag=* limit 10 | count by TAG | table TAG count DATA` to observe the raw data has been condensed as seen by the table module.

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

* [abs](abs/abs.md)
* [alias](alias/alias.md)
* [anko](anko/anko.md)
* [base64](base64/base64.md)
* [count](math/math.md#Count)
* [diff](diff/diff.md)
* [entropy](math/math.md#Entropy)
* [eval](eval/eval.md)
* [first/last](firstlast/firstlast.md)
* [geoip](geoip/geoip.md)
* [grep](grep/grep.md)
* [hexlify](hexlify/hexlify.md)
* [ip](ip/ip.md)
* [ipexist](ipexist/ipexist.md)
* [iplookup](iplookup/iplookup.md)
* [join](join/join.md)
* [langfind](langfind/langfind.md)
* [length](length/length.md)
* [limit](limit/limit.md)
* [lookup](lookup/lookup.md)
* [lower](upperlower/upperlower.md)
* [maclookup](maclookup/maclookup.md)
* [Math (list of math modules)](math/math.md)
* [max](math/math.md#Max)
* [mean](math/math.md#Mean)
* [min](math/math.md#Min)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [require](require/require.md)
* [slice](slice/slice.md)
* [sort](sort/sort.md)
* [split](split/split.md)
* [src](src/src.md)
* [stats](stats/stats.md)
* [stddev](math/math.md#Stddev)
* [strings](strings/strings.md)
* [subnet](subnet/subnet.md)
* [sum](math/math.md#Sum)
* [taint](taint/taint.md)
* [time](time/time.md)
* [unique](math/math.md#Unique)
* [upper](upperlower/upperlower.md)
* [variance](math/math.md#Variance)
* [words](words/words.md)
