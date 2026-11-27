# Search Processing Modules

Processing modules operate in pass-through mode: each one performs some action (filter, modify, sort, etc.) and passes entries down the pipeline. Each module runs in its own lightweight thread, so a search with 10 modules spreads across 10 threads. Each module's documentation notes whether it causes distributed searches to collapse and/or sort. A collapsing module forces the distributed pipelines to collapse, so that module and everything downstream of it runs on the frontend. Put as many parallel modules as possible upstream of the first collapsing module; this reduces traffic between indexers and the frontend and allows more parallelism.

Some modules, such as `count`, significantly transform or collapse data. Modules after them may no longer see the original raw data or previously extracted enumerated values. For example, `count by Src` turns the data into collapsed results like `10.0.0.1  3`. Run `tag=* limit 10 | count by TAG | raw` to see the raw output of the count module, or `tag=* limit 10 | count by TAG | table TAG count DATA` to see how the table module sees the condensed data.

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

(searchmodule_list)=
## Search module documentation

```{toctree}
---
maxdepth: 2
hidden: true
---
abs <abs/abs>
alias <alias/alias>
anko <anko/anko>
anonymize <anonymize/anonymize>
awk <awk/awk>
base64 <base64/base64>
communityid <communityid/communityid>
concat <concat/concat>
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
keep <keep/keep>
kmeans <kmeans/kmeans>
langfind <langfind/langfind>
length <length/length>
limit <limit/limit>
location <location/location>
lookup <lookup/lookup>
maclookup <maclookup/maclookup>
Math Modules (list) <math/math>
nosort <nosort/nosort>
packetlayer <packetlayer/packetlayer>
printf <printf/printf>
regex <regex/regex>
reingest <reingest/reingest>
require <require/require>
semantic <semantic/semantic>
slice <slice/slice>
sort <sort/sort>
split <split/split>
src <src/src>
stats <stats/stats>
strings <strings/strings>
subnet <subnet/subnet>
taint <taint/taint>
time <time/time>
top/bottom <topbottom/topbottom>
transaction <transaction/transaction>
Trim Modules (list) <trim/trim>
truncate <truncate/truncate>
unescape <unescape/unescape>
unique <unique/unique>
upper/lower <upperlower/upperlower>
words <words/words>
```

* [abs](abs/abs) - calculate the absolute value of an enumerated value.
* [alias](alias/alias) - create copies of enumerated values with new names.
* [anko](anko/anko) - run arbitrary code in the pipeline.
* [anonymize](anonymize/anonymize) - replace enumerated values with consistent anonymized values.
* [awk](awk/awk) - execute AWK code.
* [base64](base64/base64) - encodes or decodes base64 strings.
* [communityid](communityid/communityid) - calculate Zeek community ID values.
* [concat](concat/concat) - concatenate two or more enumerated values into a single enumerated value. `join` is an alias.
* [diff](diff/diff) - compare fields between entries.
* [dns](dns/dns) - do DNS and reverse DNS lookups.
* [enrich](enrich/enrich) - attach enumerated values to entries from strings or a resource.
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
* [iplookup](iplookup/iplookup) - enrich entries by looking up IP addresses in a table of CIDR subnets.
* [keep](keep/keep) - remove all enumerated values except those specified.
* [kmeans](kmeans/kmeans) - perform k-means clustering on numeric enumerated values.
* [langfind](langfind/langfind) - classify the language of text.
* [length](length/length) - compute the length of entries or enumerated values.
* [limit](limit/limit) - limit the number of entries which will pass further down the pipeline.
* [location](location/location) - convert individual lat/lon enumerated values into a single Gravwell Location enumerated value.
* [lookup](lookup/lookup) - enrich entries by looking up keys in a table.
* [lower](upperlower/upperlower) - convert text to lower-case.
* [maclookup](maclookup/maclookup) - look up manufacturer, address, and country information based on a MAC address.
* [Math Modules (list)](math/math) - shorthand for common [stats](stats/stats) operations.
  * count - count entries.
  * max - find a maximum value.
  * mean - find a mean value.
  * min - find a minimum value.
  * stddev - calculate standard deviation.
  * sum - sum up enumerated values.
  * variance - find variance of enumerated values.
* [nosort](nosort/nosort) - disable sorting in the pipeline.
* [packetlayer](packetlayer/packetlayer) - parse portions of a packet.
* [printf](printf/printf) - format text in the pipeline.
* [regex](regex/regex) - match and extract data using regular expressions.
* [reingest](reingest/reingest) - reingest data into another tag.
* [require](require/require) - drop any entries which lack a given enumerated value.
* [semantic](semantic/semantic) - search by meaning using vector embeddings.
* [slice](slice/slice) - low-level binary parsing & extraction.
* [sort](sort/sort) - sort entries by a given key.
* [split](split/split) - split a single entry into multiple entries.
* [src](src/src) - filter based on the SRC field of entries.
* [stats](stats/stats) - compute statistics such as count, sum, mean, min, and max.
* [strings](strings/strings) - find strings from binary data.
* [subnet](subnet/subnet) - extract & filter based on IP subnets.
* [taint](taint/taint) - track tainted values transitively across enumerated values.
* [time](time/time) - convert strings to time enumerated values, and vice versa.
* [top/bottom](topbottom/topbottom) - display the top/bottom N values.
* [transaction](transaction/transaction) - group multiple entries into single-entry "transactions" based on keys.
* [Trim Modules (list)](trim/trim) - various string trim functions.
  * trim - remove leading and trailing characters in a cutset.
  * trimleft - remove leading characters in a cutset.
  * trimprefix - remove the provided prefix.
  * trimright - remove trailing characters in a cutset.
  * trimspace - remove all leading and trailing whitespace.
  * trimsuffix - remove the provided suffix.
* [truncate](truncate/truncate) - truncate enumerated values to a specified number of characters.
* [unescape](unescape/unescape) - convert escaped text into an unescaped representation.
* [unique](unique/unique) - eliminate duplicate entries.
* [upper](upperlower/upperlower) - convert text to upper-case.
* [words](words/words) - highly optimized search for individual words.
