# Sort

By default, everything in the Gravwell search pipeline is temporally sorted. The web interface provides some additional sorting capabilities while maintaining the raw power of temporal sorting on the search backend.

The sort module allows the user to sort on one or more enumerated values. This can be very useful for organizing information. For example, a query to display a table of the top domains requested from the dnsmasq daemon might look like:

```gravwell
tag=syslog grep dnsmasq | regex ".*query\[A\]\s(?P<dnsquery>[a-zA-Z0-9\.\-]+)" | count by dnsquery | sort (count desc) | table dnsquery count
```

## Syntax

The sort module requires at least one column specification. Each column can be specified in one of two forms:

**Simple form:** `sort <enumerated value>`

**Extended form:** `sort (<enumerated value> <asc|desc> [type])...`

Multiple columns can be specified for multi-column sorting. When sorting by multiple columns, entries are first sorted by the first column; if values are equal, the second column is used as a tiebreaker, and so on.

### Sort Direction

Each column can specify a sort direction:

- `asc` - Ascending order (smallest to largest)
- `desc` - Descending order (largest to smallest)

When using the simple form (just an enumerated value name), the default direction is descending.

### Type Casting

The extended form allows specifying a type for comparison. Available types are:

| Type | Description |
|------|-------------|
| `time` | Compare as timestamps |
| `number` | Compare as numeric values |
| `ip` | Compare as IP addresses |
| `string` | Compare as strings (default) |

If no type is specified, the default is `string`, except for:
- `TIMESTAMP` - defaults to `time`
- `SRC` - defaults to `ip`

```{note}
When sorting on `TIMESTAMP`, the module sorts by the entry's timestamp, not an extracted field. If you have extracted a field named `TIMESTAMP`, it will be ignored in favor of the entry timestamp.
```

## Examples

| Command | Description |
|---------|-------------|
| `sort TIMESTAMP` | Sort by entry timestamp in descending order |
| `sort (TIMESTAMP asc)` | Sort by entry timestamp in ascending order |
| `sort (count desc)` | Sort by an enumerated value called "count" in descending order |
| `sort (SRC asc)` | Sort by entry source IP in ascending order |
| `sort (count desc number)` | Sort by "count" as a number in descending order |
| `sort (timestamp desc time)` | Sort by an extracted "timestamp" field as a time value |
| `sort (name asc) (count desc)` | Multi-column sort: first by "name" ascending, then by "count" descending |
| `sort foo bar baz` | Multi-column sort by foo, bar, and baz (all descending) |
| `sort (priority asc number) (timestamp desc time)` | Sort by priority (numeric, ascending), then by timestamp (time, descending) |

## Multi-Column Sorting

Multi-column sorting allows you to specify multiple sort keys. Entries are compared by the first column first; if two entries have equal values for that column, they are compared by the second column, and so on.

For example, to sort log entries first by severity level (highest first), then by name (alphabetically), you could use:

```gravwell
tag=syslog json level name | sort (level desc number) (name desc)
```

To sort a list of users by last name, then by first name:

```gravwell
tag=users json lastname firstname | sort (lastname asc) (firstname asc)
```

```{note}
The sort module can restrict second order temporal searching for all renderers when sorting non-temporally. This means that the overview graph and timeslice selection in the web interface will not affect a search that has been sorted non-temporally. Care must be taken to ensure that any pipeline modules following sort are expecting non-temporally ordered data.
```
