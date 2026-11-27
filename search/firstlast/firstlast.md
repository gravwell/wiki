# First / Last Modules

The `first` and `last` modules return the earliest or most recent entry by timestamp. For instance, given syslog messages from many sources, you can get the earliest or most recent entry from each host, or from each daemon.

Both modules use the same syntax. With no arguments they return a single entry. Given one or more enumerated value names, they return the first or last entry for each combination of those values, similar to the [unique](/search/unique/unique) module. Entries that have none of the named values are dropped.

```
first [enumerated value]... [over <duration>]
```

## Supported Options

* `-maxtracked <arg>`: sets the maximum number of unique keys to track, e.g. `first -maxtracked 5000 DstIP`. This guards against memory exhaustion, for example when keying on millions of IPv6 addresses. If the limit is exceeded, the search fails with an error. Defaults to 1000000. See the [stats module documentation](/search/stats/stats) for more about maxtracked.
* `-maxsize <arg>`: sets the maximum amount of memory, in megabytes, to hold when tracking entries. If the limit is exceeded, the search fails with an error. Defaults to 1024, or 15% of system memory if that is larger.

## Examples

To get the earliest entry in the search, invoke `first` with no arguments:

```gravwell
tag=gravwell first
```

To get the earliest entry for each unique value of the enumerated value `foo`:

```gravwell
tag=gravwell json foo | first foo
```

Add `over <duration>` to get one result per key in each time window instead of one for the whole search. The duration must be a whole number of seconds, minimum one second. This returns the most recent entry for each value of `foo` in each 10-minute window:

```gravwell
tag=gravwell json foo | last foo over 10m
```

## Caveats

`first` and `last` compare entry timestamps, not the order in which entries reach the module. This query still returns the entry with the earliest timestamp, even though the data was sorted by `foo`:

```gravwell
tag=gravwell json foo | sort by foo | first
```

If several entries share the earliest (or latest) timestamp, the one that reaches the module first is kept.
