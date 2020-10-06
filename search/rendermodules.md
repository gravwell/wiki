# Renderer Modules

Renderer modules are in charge of receiving data from the search module pipeline and organizing it for display to the user. When possible, the renderers provide for a second order temporal index. This allows for moving around and zeroing in on time spans within the original search. Renderers can optionally save search results, which can be reopened and viewed or even passed to another instance of Gravwell. This is useful for archiving a view of data or saving the results which survive well after stored data is expired or purposefully deleted.

Every search module has universal enumerated values for records.

* SRC -- the source of the data.
* TAG -- the Tag attached to the data.
* TIMESTAMP -- the timestamp associated with the entry

## Renderer Module list

### Charts, Graphs, and Gauges
* [chart](chart/chart.md) - Render data as line graphs, bar graphs, etc.
* [fdg](fdg/fdg.md) - Force-directed graphs.
* [stackgraph](stackgraph/stackgraph.md) - Stack graphs.
* [gauge/numbercard](gauge/gauge.md) - Gauges and numeric cards.

### Tables and Text
* [table](table/table.md) - Display tables of enumerated values.
* [text](text/text.md) - Output the body of entries with minimal formatting.
* [raw](raw/raw.md) - Output data completely unformatted.
* [pcap](pcap/pcap.md) - Show an overview of the contents of network packets.

### Maps
* [pointmap / heatmap](map/map.md) - Display heatmaps or individual points on a map.
* [point2point](point2point/point2point.md) - Display data entries which have both a source and a destination.