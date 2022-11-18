# Search Processing Modules

Search modules are modules that operate on data in a pass-through mode, meaning that they perform some action (filter, modify, sort, etc.) and pass the entries down the pipeline. There can be many search modules and each operates in its own lightweight thread.  This means that if there are 10 modules in a search, the pipeline will spread out and use 10 threads.  Documentation for each module will indicate if the module causes distributed searches to collapse and/or sort.  Modules that collapse force the distributed pipelines to collapse, meaning that the module as well as all downstream modules execute on the frontend.  When starting a search it's best to put as many parallel modules as possible upstream of the first collapsing module, decreasing pressure on the communication pipe and allowing for greater parallelism.

Some modules significantly transform or collapse data, such as `count`. Pipeline modules following these collapsing modules may not be dealing with Raw data or previously created Enumerated Values. In short, things like `count by Src` turn data into the collapsed results with entries such as `10.0.0.1  3`. To illustrate, run the search `tag=* limit 10 | count by TAG | raw` to see the raw output from the count module or `tag=* limit 10 | count by TAG | table TAG count DATA` to observe the raw data has been condensed as seen by the table module.

## Universal Flags

Some flags appear in several different search modules and have the same meaning throughout:

* `-e <source name>` specifies that the module should attempt to read its input data from the given enumerated value rather than from the entry's data field. This is useful in for modules like [json](json/json), where the JSON-encoded data may have been extracted from a larger data record, for example the following search will attempt to read JSON fields from the payloads of HTTP packets: `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-r <resource name>` specifies a resource in the [resources](/resources/resources) system. This is generally used to store additional data used by the module, such as a GeoIP mapping table used by the [geoip](geoip/geoip) module.
* `-v` indicates that the normal pass/drop logic should be inverted. For example the [grep](grep/grep) module normally passes entries which match a given pattern and drop those which do not match; specifying the `-v` flag will cause it to drop entries which match and pass those which do not.
* `-s` indicates a "strict" mode. If a module normally allows an entry to proceed down the pipeline if any one of several conditions are met, setting the strict flag means an entry will proceed only if *all* conditions are met. For example, the [require](require/require) module will normally pass an entry if it contains any one of the required enumerated values, but when the `-s` flag is used, it will only pass entries which contain *all* specified enumerated values.
* `-p` indicates "permissive" mode.  If a module normally drops entries when patterns and filters do not match, the permissive flag tells the module to let the module go through.  The [regex](regex/regex) and [grok](grok/grok) modules are good examples where the permissive flag can be valuable.

## Universal Enumerated Values

The following enumerated values are available for every entry. They're actually convenient names for properties of the raw entries themselves, but can be treated as enumerated value names.

* SRC -- the source of the entry data.
* TAG -- the tag attached to the entry.
* TIMESTAMP -- the timestamp of the entry.
* DATA -- the actual entry data.
* NOW -- the current time.

These can be used just like user-defined enumerated values, thus `table foo bar DATA NOW` is valid. They do not need to be explicitly *extracted* anywhere; they are always available.

(searchmodule_list)=
## Search module documentation

```{toctree}
---
maxdepth: 1
hidden: true
---
abs <abs/abs>
alias <alias/alias>
anko <anko/anko>
anonymize <anonymize/anonymize>
awk <awk/awk>
base64 <base64/base64>
communityid <communityid/communityid>
count <Count_module>
diff <diff/diff>
dns <dns/dns>
enrich <enrich/enrich>
entropy <entropy/entropy>
eval <eval/eval>
filetype <filetype/filetype>
first/last <firstlast/firstlast>
fuse <fuse/fuse>
geodist <geodist/geodist>
geoip <geoip/geoip>
grep <grep/grep>
hexlify <hexlify/hexlify>
ip <ip/ip>
ipexist <ipexist/ipexist>
iplookup <iplookup/iplookup>
join <join/join>
langfind <langfind/langfind>
length <length/length>
limit <limit/limit>
location <location/location>
lookup <lookup/lookup>
lower <upperlower/upperlower>
maclookup <maclookup/maclookup>
Math Modules (list) <math/math>
nosort <nosort/nosort>
packetlayer <packetlayer/packetlayer>
printf <printf/printf>
regex <regex/regex>
require <require/require>
slice <slice/slice>
sort <sort/sort>
split <split/split>
src <src/src>
stats <stats/stats>
strings <strings/strings>
subnet <subnet/subnet>
taint <taint/taint>
time <time/time>
transaction <transaction/transaction>
truncate <truncate/truncate>
unescape <unescape/unescape>
upper <upperlower/upperlower>
words <words/words>
```

* [abs](abs/abs) - calculate the absolute value of an enumerated value.
* [alias](alias/alias) - create copies of enumerated values with new names.
* [anko](anko/anko) - run arbitrary code in the pipeline.
* [anonymize](anonymize/anonymize) - anonymize IP addresses.
* [awk](awk/awk) - execute AWK code.
* [base64](base64/base64) - encodes or decodes base64 strings.
* [communityid](communityid/communityid) - calculate Zeek community ID values.
* [diff](diff/diff) - compare fields between entries.
* [dns](dns/dns) - do DNS and reverse DNS lookups.
* [enrich](enrich/enrich) - manually attach enumerated values to entries.
* [entropy](entropy/entropy) - calculate entropy of enumerated values.
* [eval](eval/eval) - evaluate arbitrary logic expressions.
* [filetype](filetype/filetype) - detect filetypes of binary data.
* [first/last](firstlast/firstlast) - take the first or last entry.
* [fuse](fuse/fuse) - join data from disparate data sources.
* [geodist](geodist/geodist) - compute distance between locations.
* [geoip](geoip/geoip) - look up GeoIP locations.
* [grep](grep/grep) - search for strings in entries.
* [hexlify](hexlify/hexlify) - encode data into ASCII hex representation, or vice versa.
* [ip](ip/ip) - convert & filter IP addresses.
* [ipexist](ipexist/ipexist) - check if IP address exists in a lookup table.
* [iplookup](iplookup/iplookup) - enrich entries by looking up IP addresses in a table which can contain CIDR subnets rather that individual IPs.
* [join](join/join) - join two or more enumerated values into a single enumerated value.
* [langfind](langfind/langfind) - classify the language of text.
* [length](length/length) - compute the length of entries or enumerated values.
* [limit](limit/limit) - limit the number of entries which will pass further down the pipeline.
* [location](location/location) - convert individual lat/lon enumerated values into a single Gravwell Location enumerated value.
* [lookup](lookup/lookup) - enrich entries by looking up keys in a table.
* [lower](upperlower/upperlower) - convert text to lower-case.
* [maclookup](maclookup/maclookup) - look up manufacturer, address, and country information based on a MAC address.
* [Math Modules (list)](math/math.md) - perform math operations.
  * count - count entries.
  * max - find a maximum value.
  * mean - find a mean value.
  * min - find a minimum value.
  * stddev - calculate standard deviation.
  * sum - sum up enumerated values.
  * unique - eliminate duplicate entries.
  * variance - find variance of enumerated values.
* [nosort](nosort/nosort) - disable sorting in the pipeline.
* [packetlayer](packetlayer/packetlayer) - parse portions of a packet.
* [printf](printf/printf) - format text in the pipeline.
* [regex](regex/regex) - match and extract data using regular expressions.
* [require](require/require) - drop any entries which lack a given enumerated value.
* [slice](slice/slice) - low-level binary parsing & extraction.
* [sort](sort/sort) - sort entries by a given key.
* [split](split/split) - split a single entry into multiple entries.
* [src](src/src) - filter based on the SRC field of entries.
* [stats](stats/stats) - perform math operations.
* [strings](strings/strings) - find strings from binary data.
* [subnet](subnet/subnet) - extract & filter based on IP subnets.
* [taint](taint/taint) - taint tracking.
* [time](time/time) - convert strings to time enumerated values, and vice versa.
* [transaction](transaction/transaction) - group multiple entries into single-entry "transactions" based on keys.
* [truncate](truncate/truncate) - truncate entries or enumerated values to a specified number of characters.
* [unescape](unescape/unescape) - convert escaped text into an unescaped representation.
* [upper](upperlower/upperlower) - convert text to upper-case.
* [words](words/words) - highly optimized search for individual words.
