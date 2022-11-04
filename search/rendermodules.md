# Renderer Modules

Renderer modules are in charge of receiving data from the search module pipeline and organizing it for display to the user. When possible, the renderers provide for a second order temporal index. This allows for moving around and zeroing in on time spans within the original search. Renderers can optionally save search results, which can be reopened and viewed or even passed to another instance of Gravwell. This is useful for archiving a view of data or saving the results which survive well after stored data is expired or purposefully deleted.

Every search module has universal enumerated values for records.

* SRC -- the source of the data.
* TAG -- the Tag attached to the data.
* TIMESTAMP -- the timestamp associated with the entry

## Renderer Module list

### Charts, Graphs, and Gauges

```{toctree}
chart - Render data as line graphs, bar graphs, etc. <chart/chart.md>
fdg - Force-directed graphs. <fdg/fdg.md>
stackgraph - Stack graphs. <stackgraph/stackgraph.md>
gauge/numbercard - Gauges and numeric cards. <gauge/gauge.md>
wordcloud - Word clouds. <wordcloud/wordcloud.md>
```

### Tables and Text

```{toctree}
table - Display tables of enumerated values. <table/table.md>
text - Output the body of entries with minimal formatting. <text/text.md>
raw - Output data completely unformatted. <raw/raw.md>
pcap - Show an overview of the contents of network packets. <pcap/pcap.md>
```

### Maps

```{toctree}
pointmap / heatmap - Display heatmaps or individual points on a map. <map/map.md>
point2point - Display data entries which have both a source and a destination. <point2point/point2point.md>
```