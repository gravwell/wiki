# Renderer Modules

Renderer modules are in charge of receiving data from the search module pipeline and organizing it for display to the user. When possible, the renderers provide for a second order temporal index. This allows for moving around and zeroing in on time spans within the original search. Renderers can optionally save search results, which can be reopened and viewed or even passed to another instance of Gravwell. This is useful for archiving a view of data or saving the results which survive well after stored data is expired or purposefully deleted.

Every search module has universal enumerated values for records.

* SRC -- the source of the data.
* TAG -- the Tag attached to the data.
* TIMESTAMP -- the timestamp associated with the entry

## Renderer Module list

### Charts, Graphs, and Gauges

```{toctree}
---
maxdepth: 1
hidden: true
---
chart <chart/chart>
fdg <fdg/fdg>
stackgraph <stackgraph/stackgraph>
gauge/numbercard <gauge/gauge>
wordcloud <wordcloud/wordcloud>
```

* [chart](chart/chart) - Render data as line graphs, bar graphs, etc.
* [fdg](fdg/fdg) - Force-directed graphs.
* [stackgraph](stackgraph/stackgraph) - Stack graphs.
* [gauge/numbercard](gauge/gauge) - Gauges and numeric cards.
* [wordcloud](wordcloud/wordcloud) - Word clouds.

### Tables and Text

```{toctree}
---
maxdepth: 1
hidden: true
---
table <table/table>
text <text/text>
raw <raw/raw>
pcap <pcap/pcap>
```

* [table](table/table) - Display tables of enumerated values.
* [text](text/text) - Output the body of entries with minimal formatting.
* [raw](raw/raw) - Output data completely unformatted.
* [pcap](pcap/pcap) - Show an overview of the contents of network packets.

### Maps

```{toctree}
---
maxdepth: 1
hidden: true
---
pointmap/heatmap <map/map>
point2point <point2point/point2point>
```

* [pointmap/heatmap](map/map) - Display heatmaps or individual points on a map.
* [point2point](point2point/point2point) - Display data entries which have both a source and a destination.
