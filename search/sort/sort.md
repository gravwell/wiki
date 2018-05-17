## Sort

By default, everything in the Gravwell search pipeline is temporally sorted. The web interface provides some additional sorting capabilities while maintaining the raw power of temporal sorting on the search backend.

The sort module, however, allows the user to sort on other values. This can be very useful for organizing information. For example, a query to display a table of the top domains requested from the dnsmasq daemon might look like:

```
tag=syslog grep dnsmasq | regex ".*query\[A\]\s(?P<dnsquery>[a-zA-Z0-9\.\-]+)" | count by dnsquery | sort by count desc | table dnsquery count
```

The syntax is `sort [by sortparam] [asc/desc]`. `sortparam` is the parameter to sort by, which can be `time`, `tag`, `src`, or any enumerated value. The sort parameter is optional; if not specified, it defaults to `time`. The other parameter selects the direction of sorting, either ascending (`asc`) or descending (`desc`). If not specified, it defaults to descending sort for time and ascending sort for all other parameters.

Some example sort invocations:

| Command | Description |
|---------|-------------|
| `sort` | Sort by time in descending order |
| `sort by tag` | Sort by entry tag in ascending order |
| `sort by count asc` | Sort by an enumerated value called "count" in ascending order |
| `sort asc` | Sort by time in ascending order |
| `sort by src desc` | Sort by entry source in descending order |

Note: The sort module collapses the pipeline  and can restrict second order temporal searching for all renderers.  This means that the overview graph and timeslice selection in the web interface will not affect a search that has been sorted non-temporally. Care must be taken to ensure that any pipeline modules following sort are expecting non-temporally ordered data.