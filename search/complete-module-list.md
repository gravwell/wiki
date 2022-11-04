# Complete List of Search Pipeline Modules

```{toctree}
---
maxdepth: 1
caption: Search Modules
---
abs - calculate the absolute value of an enumerated value. <abs/abs>
alias - create copies of enumerated values with new names. <alias/alias>
anko - run arbitrary code in the pipeline. <anko/anko>
anonymize - anonymize IP addresses. <anonymize/anonymize>
ax - automatically extract fields from entries. <ax/ax>
awk - execute AWK code. <awk/awk>
base64 - encode/decode base64. <base64/base64>
canbus - decode CANBUS data. <canbus/canbus>
cef - decode CEF data. <cef/cef>
chart - render charts. <chart/chart>
communityid - calculate Zeek community ID values. <communityid/communityid>
count - count entries. <math/math>
csv - extract fields from CSV data. <csv/csv>
diff - compare fields between entries. <diff/diff>
dns - do DNS and reverse DNS lookups. <dns/dns>
dump - dump entries from a resource into the pipeline. <dump/dump>
enrich - manually attach enumerated values to entries. <enrich/enrich>
entropy - calculate entropy of enumerated values. <entropy/entropy>
eval - evaluate arbitrary logic expressions. <eval/eval>
filetype - detect filetypes of binary data. <filetype/filetype>
fdg - generate Force Directed Graphs. <fdg/fdg>
fields - extract data from entries using arbitrary field separators. <fields/fields>
first/last - take the first or last entry. <firstlast/firstlast>
fuse - join data from disparate data sources. <fuse/fuse>
gauge - render gauges. <gauge/gauge>
geodist - compute distance between locations. <geodist/geodist>
geoip - look up GeoIP locations. <geoip/geoip>
grep - search for strings in entries. <grep/grep>
grok - extract data from complicated text structures using pre-defined regular expressions. <grok/grok>
hexlify - encode data into ASCII hex representation, or vice versa. <hexlify/hexlify>
ip - convert & filter IP addresses. <ip/ip>
ipexist - check if IP address exists in a lookup table. <ipexist/ipexist>
ipfix - extract data from IPFIX records. <ipfix/ipfix>
iplookup - enrich entries by looking up IP addresses in a table which can contain CIDR subnets rather that individual IPs. <iplookup/iplookup>
j1939 - parse J1939 data. <j1939/j1939>
join - join two or more enumerated values into a single enumerated value. <join/join>
json - extract elements from JSON data. <json/json>
kv - parse key-value data. <kv/kv>
langfind - classify the language of text. <langfind/langfind>
length - compute the length of entries or enumerated values. <length/length>
limit - limit the number of entries which will pass further down the pipeline. <limit/limit>
location - convert individual lat/lon enumerated values into a single Gravwell Location enumerated value. <location/location>
lookup - enrich entries by looking up keys in a table. <lookup/lookup>
lower - convert text to lower-case. <upperlower/upperlower>
maclookup - look up manufacturer, address, and country information based on a MAC address. <maclookup/maclookup>
Math (list of math modules) - perform math operations. <math/math>
max - find a maximum value. <math/math>
mean - find a mean value. <math/math>
min - find a minimum value. <math/math>
netflow - parse Netflow records. <netflow/netflow>
nosort - disable sorting in the pipeline. <nosort/nosort>
numbercard - display number cards. <gauge/gauge>
packet - parse raw packets. <packet/packet>
packetlayer - parse portions of a packet. <packetlayer/packetlayer>
path - extract portions of pathnames. <path/path>
pcap - render human-friendly representations of entire packets. <pcap/pcap>
pointmap / heatmap - display point maps or heatmaps. <map/map>
point2point - display a point-to-point map. <point2point/point2point>
printf - format text in the pipeline. <printf/printf>
raw - render entries with zero processing applied. <raw/raw>
regex - match and extract data using regular expressions. <regex/regex>
require - drop any entries which lack a given enumerated value. <require/require>
slice - low-level binary parsing & extraction. <slice/slice>
sort - sort entries by a given key. <sort/sort>
split - split a single entry into multiple entries. <split/split>
src - filter based on the SRC field of entries. <src/src>
stackgraph - render a stacked graph. <stackgraph/stackgraph>
stats - perform math operations. <stats/stats>
stddev - calculate standard deviation. <math/math>
strings - find strings from binary data. <strings/strings>
subnet - extract & filter based on IP subnets. <subnet/subnet>
sum - sum up enumerated values. <math/math>
syslog - parse and extract syslog entries. <syslog/syslog>
table - render a table from entries. <table/table>
taint - taint tracking. <taint/taint>
text - render entries as text. <text/text>
time - convert strings to time enumerated values, and vice versa. <time/time>
transaction - group multiple entries into single-entry "transactions" based on keys. <transaction/transaction>
truncate - truncate entries or enumerated values to a specified number of characters. <truncate/truncate>
unescape - convert escaped text into an unescaped representation. <unescape/unescape>
unique - eliminate duplicate entries. <math/math>
upper - convert text to upper-case. <upperlower/upperlower>
variance - find variance of enumerated values. <math/math>
winlog - parse Windows logs. <winlog/winlog>
wordcloud - render word clouds. <wordcloud/wordcloud>
words - highly optimized search for individual words. <words/words>
xml - parse XML data. <xml/xml>
```
