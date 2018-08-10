# Search Modules

Search modules are modules that operate on data in a passthrough mode, meaning that they perform some action (filter, modify, sort, etc.) and pass the entries down the pipeline. There can be many search modules and each operates in its own lightweight thread.  This means that if there are 10 modules in a search, the pipeline will spread out and use 10 threads.  Documentation for each module will indicate if the module causes distributed searches to collapse and/or sort.  Modules that collapse force the distributed pipelines to collapse, meaning that the module as well as all downstream modules execute on the frontend.  When starting a search it's best to put as many parallel modules as possible upstream of the first collapsing module, decreasing pressure on the communication pipe and allowing for greater parallelism.

## Universal Enumerated Values

Every search module has universal enumerated values for records.

* SRC -- the source of the entry data.
* TAG -- the tag attached to the entry.
* TIMESTAMP -- the timestamp of the entry.

These can be used just like user-defined enumerated values.

## Search module documentation

* [anko](anko/anko.md)
* [base64](base64/base64.md)
* [canbus](canbus/canbus.md)
* [cef](cef/cef.md)
* [count](math/math.md)
* [eval](eval/eval.md)
* [fields](fields/fields.md)
* [geoip](geoip/geoip.md)
* [grep](grep/grep.md)
* [hexlify](hexlify/hexlify.md)
* [ipfix](ipfix/ipfix.md)
* [j1939](j1939/j1939.md)
* [join](join/join.md)
* [json](json/json.md)
* [langfind](langfind/langfind.md)
* [limit](limit/limit.md)
* [lookup](lookup/lookup.md)
* [lower](upperlower/upperlower.md)
* [Math (list of math modules)](math/math.md)
* [max](math/math.md)
* [mean](math/math.md)
* [min](math/math.md)
* [namedfields](namedfields/namedfields.md)
* [netflow](netflow/netflow.md)
* [packet](packet/packet.md)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [require](require/require.md)
* [slice](slice/slice.md)
* [sort](sort/sort.md)
* [src](src/src.md)
* [stddev](math/math.md)
* [strings](strings/strings.md)
* [subnet](subnet/subnet.md)
* [sum](math/math.md)
* [syslog](syslog/syslog.md)
* [taint](taint/taint.md)
* [unique](math/math.md)
* [upper](upperlower/upperlower.md)
* [variance](math/math.md)
* [xml](xml/xml.md)
