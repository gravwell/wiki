# Search Processing Modules

Search modules are modules that operate on data in a pass-through mode, meaning that they perform some action (filter, modify, sort, etc.) and pass the entries down the pipeline. There can be many search modules and each operates in its own lightweight thread.  This means that if there are 10 modules in a search, the pipeline will spread out and use 10 threads.  Documentation for each module will indicate if the module causes distributed searches to collapse and/or sort.  Modules that collapse force the distributed pipelines to collapse, meaning that the module as well as all downstream modules execute on the frontend.  When starting a search it's best to put as many parallel modules as possible upstream of the first collapsing module, decreasing pressure on the communication pipe and allowing for greater parallelism.

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

* [abs](abs/abs.md) - calculate the absolute value of an enumerated value.
* [alias](alias/alias.md) - create copies of enumerated values with new names.
* [anko](anko/anko.md) - run arbitrary code in the pipeline.
* [anonymize](anonymize/anonymize.md) - anonymize IP addresses.
* [awk](awk/awk.md) - execute AWK code.
* [base64](base64/base64.md) - encodes or decodes base64 strings.
* [communityid](communityid/communityid.md) - calculate Zeek community ID values.
* [count](math/math.md#Count) - count entries.
* [diff](diff/diff.md) - compare fields between entries.
* [dns](dns/dns.md) - do DNS and reverse DNS lookups.
* [enrich](enrich/enrich.md) - manually attach enumerated values to entries.
* [entropy](entropy/entropy.md) - calculate entropy of enumerated values.
* [eval](eval/eval.md) - evaluate arbitrary logic expressions.
* [filetype](filetype/filetype.md) - detect filetypes of binary data.
* [first/last](firstlast/firstlast.md) - take the first or last entry.
* [fuse](fuse/fuse.md) - join data from disparate data sources.
* [geodist](geodist/geodist.md) - compute distance between locations.
* [geoip](geoip/geoip.md) - look up GeoIP locations.
* [grep](grep/grep.md) - search for strings in entries.
* [hexlify](hexlify/hexlify.md) - encode data into ASCII hex representation, or vice versa.
* [ip](ip/ip.md) - convert & filter IP addresses.
* [ipexist](ipexist/ipexist.md) - check if IP address exists in a lookup table.
* [iplookup](iplookup/iplookup.md) - enrich entries by looking up IP addresses in a table which can contain CIDR subnets rather that individual IPs.
* [join](join/join.md) - join two or more enumerated values into a single enumerated value.
* [langfind](langfind/langfind.md) - classify the language of text.
* [length](length/length.md) - compute the length of entries or enumerated values.
* [limit](limit/limit.md) - limit the number of entries which will pass further down the pipeline.
* [location](location/location.md) - convert individual lat/lon enumerated values into a single Gravwell Location enumerated value.
* [lookup](lookup/lookup.md) - enrich entries by looking up keys in a table.
* [lower](upperlower/upperlower.md) - convert text to lower-case.
* [maclookup](maclookup/maclookup.md) - look up manufacturer, address, and country information based on a MAC address.
* [Math (list of math modules)](math/math.md) - perform math operations.
* [max](math/math.md#Max) - find a maximum value.
* [mean](math/math.md#Mean) - find a mean value.
* [min](math/math.md#Min) - find a minimum value.
* [nosort](nosort/nosort.md) - disable sorting in the pipeline.
* [packetlayer](packetlayer/packetlayer.md) - parse portions of a packet.
* [printf](printf/printf.md) - format text in the pipeline.
* [regex](regex/regex.md) - match and extract data using regular expressions.
* [require](require/require.md) - drop any entries which lack a given enumerated value.
* [slice](slice/slice.md) - low-level binary parsing & extraction.
* [sort](sort/sort.md) - sort entries by a given key.
* [split](split/split.md) - split a single entry into multiple entries.
* [src](src/src.md) - filter based on the SRC field of entries.
* [stats](stats/stats.md) - perform math operations.
* [stddev](math/math.md#Stddev) - calculate standard deviation.
* [strings](strings/strings.md) - find strings from binary data.
* [subnet](subnet/subnet.md) - extract & filter based on IP subnets.
* [sum](math/math.md#Sum) - sum up enumerated values.
* [taint](taint/taint.md) - taint tracking.
* [time](time/time.md) - convert strings to time enumerated values, and vice versa.
* [transaction](transaction/transaction.md) - group multiple entries into single-entry "transactions" based on keys.
* [truncate](truncate/truncate.md) - truncate entries or enumerated values to a specified number of characters.
* [unescape](unescape/unescape.md) - convert escaped text into an unescaped representation.
* [unique](math/math.md#Unique) - eliminate duplicate entries.
* [upper](upperlower/upperlower.md) - convert text to upper-case.
* [variance](math/math.md#Variance) - find variance of enumerated values.
* [words](words/words.md) - highly optimized search for individual words.
