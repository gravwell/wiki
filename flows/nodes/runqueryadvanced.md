# Run Query - Advanced Node

This node executes a Gravwell query, much like the [Run a Query](runquery) node, but provides more advanced options for specifying the search duration. It outputs a structure into the payload (named `search` by default) which contains information about the search and allows other nodes to access the results.

## Configuration

* `Query String`, required: the Gravwell query to run.
* `Start Time`, required: the starting point for the query timeframe. This can be either a literal timestamp ("2023-01-01T00:00:00Z07:00") or a [Go-style duration string](https://pkg.go.dev/time#ParseDuration) such as "-10h" or "-20m".
* `End Time`, required: the end point for the query timeframe. This can be either a literal timestamp ("2023-01-31T11:59:59Z07:00") or a [Go-style duration string](https://pkg.go.dev/time#ParseDuration) such as "0h" or "-5m".
* `Output Variable Name`: the name to use for results in the payload, default "search".

```{note}
Durations should usually be negative, e.g. set Start Time to "-10h" to start the query ten hours ago. A positive duration indicates a time in the future; setting Start Time to "-1h" and End Time to "1h" will run a search over data spanning from 1 hour in the past to 1 hour in the future (potentially useful in the case of time zone issues or clock skew).

Durations are always taken relative to the *scheduled* execution time of the flow
```

Obviously it is rarely useful to run a flow over the exact same timeframe again and again, so users will rarely enter a literal timestamp by hand. However, reading start/end timestamps from *variables* means you can programmatically decide (using other nodes) a specific timeframe to search over.

## Output

The node inserts an object (named `search` by default) into the payload containing information about the search. The structure is identical to that output by the [Run a Query](runquery) node.

## Example

This node behaves exactly like the [Run a Query](runquery) node, except for the difference in timeframe specification. This section will therefore show a few examples of how the timeframe could be configured:

| Start | End | Explanation |
|-------|-----|-------------|
| 2023-01-01T00:00:00Z07:00 | 2023-01-31T11:59:59Z07:00 | Search over the entire month of January 2023 |
| -24h | 0h | Search over the last day (equivalent to using the normal Run a Query node with a duration of 24 hours) |
| -36h | -12h | If scheduled at noon, searches over the *previous day's* data. This can help if, for whatever reason, it takes time for the entirety of a day's logs to be collected |
| -24h | 24h | Searches data over a 2-day span centered on the scheduled time, extending 24h into the *future*. This is not usually a good idea, but sometimes data may come in with timestamps in the future |
