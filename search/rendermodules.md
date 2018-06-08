# Renderer Modules

Renderer modules are in charge of receiving data from the search module pipeline and organizing it for display to the user. When possible, the renderers provide for a second order temporal index. This allows for moving around and zeroing in on time spans within the original search. Renderers can optionally save search results, which can be reopened and viewed or even passed to another instance of Gravwell. This is useful for archiving a view of data or saving the results which survive well after stored data is expired or purposefully deleted.

Every search module has universal enumerated values for records.

* SRC -- the source of the data.
* TAG -- the Tag attached to the data.
* TIMESTAMP -- the timestamp associated with the entry

## Renderer Module list

* [chart](chart/chart.md)
* [fdg](fdg/fdg.md)
* [raw](raw/raw.md)
* [stackgraph](stackgraph/stackgraph.md)
* [text](text/text.md)
* [table](table/table.md)
* [pointmap / heatmap](map/map.md)