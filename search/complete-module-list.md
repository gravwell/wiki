# Complete List of Search Pipeline Modules

* [abs](abs/abs.md) - calculate the absolute value of an enumerated value.
* [alias](alias/alias.md) - create copies of enumerated values with new names.
* [anko](anko/anko.md) - run arbitrary code in the pipeline.
* [anonymize](anonymize/anonymize.md) - anonymize IP addresses.
* [ax](ax/ax.md) - automatically extract fields from entries.
* [awk](awk/awk.md) - execute AWK code.
* [base64](base64/base64.md) - encode/decode base64.
* [canbus](canbus/canbus.md) - decode CANBUS data.
* [cef](cef/cef.md) - decode CEF data.
* [chart](chart/chart.md) - render charts.
* [communityid](communityid/communityid.md) - calculate Zeek community ID values.
* [count](math/math.md#Count) - count entries.
* [csv](csv/csv.md) - extract fields from CSV data.
* [diff](diff/diff.md) - compare fields between entries.
* [dns](dns/dns.md) - do DNS and reverse DNS lookups.
* [dump](dump/dump.md) - dump entries from a resource into the pipeline.
* [enrich](enrich/enrich.md) - manually attach enumerated values to entries.
* [entropy](entropy/entropy.md) - calculate entropy of enumerated values.
* [eval](eval/eval.md) - evaluate arbitrary logic expressions.
* [filetype](filetype/filetype.md) - detect filetypes of binary data.
* [fdg](fdg/fdg.md) - generate Force Directed Graphs.
* [fields](fields/fields.md) - extract data from entries using arbitrary field separators.
* [first/last](firstlast/firstlast.md) - take the first or last entry.
* [fuse](fuse/fuse.md) - join data from disparate data sources.
* [gauge](gauge/gauge.md) - render gauges.
* [geodist](geodist/geodist.md) - compute distance between locations.
* [geoip](geoip/geoip.md) - look up GeoIP locations.
* [grep](grep/grep.md) - search for strings in entries.
* [grok](grok/grok.md) - extract data from complicated text structures using pre-defined regular expressions.
* [hexlify](hexlify/hexlify.md) - encode data into ASCII hex representation, or vice versa.
* [ip](ip/ip.md) - convert & filter IP addresses.
* [ipexist](ipexist/ipexist.md) - check if IP address exists in a lookup table.
* [ipfix](ipfix/ipfix.md) - extract data from IPFIX records.
* [iplookup](iplookup/iplookup.md) - enrich entries by looking up IP addresses in a table which can contain CIDR subnets rather that individual IPs.
* [j1939](j1939/j1939.md) - parse J1939 data.
* [join](join/join.md) - join two or more enumerated values into a single enumerated value.
* [json](json/json.md) - extract elements from JSON data.
* [kv](kv/kv.md) - parse key-value data.
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
* [netflow](netflow/netflow.md) - parse Netflow records.
* [nosort](nosort/nosort.md) - disable sorting in the pipeline.
* [numbercard](gauge/gauge.md) - display number cards.
* [packet](packet/packet.md) - parse raw packets.
* [packetlayer](packetlayer/packetlayer.md) - parse portions of a packet.
* [path](path/path.md) - extract portions of pathnames.
* [pcap](pcap/pcap.md) - render human-friendly representations of entire packets.
* [pointmap / heatmap](map/map.md) - display point maps or heatmaps.
* [point2point](point2point/point2point.md) - display a point-to-point map.
* [printf](printf/printf.md) - format text in the pipeline.
* [raw](raw/raw.md) - render entries with zero processing applied.
* [regex](regex/regex.md) - match and extract data using regular expressions.
* [require](require/require.md) - drop any entries which lack a given enumerated value.
* [slice](slice/slice.md) - low-level binary parsing & extraction.
* [sort](sort/sort.md) - sort entries by a given key.
* [split](split/split.md) - split a single entry into multiple entries.
* [src](src/src.md) - filter based on the SRC field of entries.
* [stackgraph](stackgraph/stackgraph.md) - render a stacked graph.
* [stats](stats/stats.md) - perform math operations.
* [stddev](math/math.md#Stddev) - calculate standard deviation.
* [strings](strings/strings.md) - find strings from binary data.
* [subnet](subnet/subnet.md) - extract & filter based on IP subnets.
* [sum](math/math.md#Sum) - sum up enumerated values.
* [syslog](syslog/syslog.md) - parse and extract syslog entries.
* [table](table/table.md) - render a table from entries.
* [taint](taint/taint.md) - taint tracking.
* [text](text/text.md) - render entries as text.
* [time](time/time.md) - convert strings to time enumerated values, and vice versa.
* [transaction](transaction/transaction.md) - group multiple entries into single-entry "transactions" based on keys.
* [truncate](truncate/truncate.md) - truncate entries or enumerated values to a specified number of characters.
* [unescape](unescape/unescape.md) - convert escaped text into an unescaped representation.
* [unique](math/math.md#Unique) - eliminate duplicate entries.
* [upper](upperlower/upperlower.md) - convert text to upper-case.
* [variance](math/math.md#Variance) - find variance of enumerated values.
* [winlog](winlog/winlog.md) - parse Windows logs.
* [wordcloud](wordcloud/wordcloud.md) - render word clouds.
* [words](words/words.md) - highly optimized search for individual words.
* [xml](xml/xml.md) - parse XML data.
