# Sort

By default, search results are sorted by time. The sort module sorts on other values instead. For example, this query shows a table of the domains most often requested from the dnsmasq daemon:

```gravwell
tag=syslog grep dnsmasq | regex ".*query\[A\]\s(?P<dnsquery>[a-zA-Z0-9\.\-]+)" | count by dnsquery | sort by count desc | table dnsquery count
```

The syntax is `sort [by sortparam] [asc/desc]`. `sortparam` is `time`, `tag`, `src`, or the name of a single enumerated value, and defaults to `time`. The direction is ascending (`asc`) or descending (`desc`), and defaults to descending for time and ascending for everything else.

Enumerated values are compared according to their type: numbers numerically, strings lexically. `sort by tag` orders entries by internal tag ID, not alphabetically by tag name.

Some example sort invocations:

| Command | Description |
|---------|-------------|
| `sort` | Sort by time in descending order |
| `sort by tag` | Sort by entry tag in ascending order |
| `sort by count asc` | Sort by an enumerated value called `count` in ascending order |
| `sort asc` | Sort by time in ascending order |
| `sort by src desc` | Sort by entry source in descending order |

```{note}
The sort module collapses the pipeline and can restrict second-order temporal searching for all renderers. The overview graph and timeslice selection in the web interface will not affect a search that has been sorted by anything other than time. Make sure any modules after `sort` can handle data that is not in time order.
```
