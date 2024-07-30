# First / Last Modules

The `first` and `last` modules are a convenient way to see specific entries over a time period. For instance, given a collection of syslog messages from a number of sources, one might wish to see the earliest or most recent entry from each individual host, or from each daemon.

The modules use the same syntax. Each optionally takes one or more enumerated value names; if specified, the modules will emit the first/last entry for each combination of values of the enumerated values, similar to the behavior of the unique module.

```
first [enumerated value]...
```

## Supported Options

* `-maxtracked <arg>`: Set the maximum number of entries to track when using keys.
* `-maxsize <arg>`: Set the maximum number of MB to hold in memory when tracking using keys.

## Examples

To get just the first entry in a query by time, simply invoke the `first` module with no arguments:

```gravwell
tag=gravwell first
```

To get the first entry for each unique value of the enumerated value "foo", invoke `first` with the argument "foo":

```gravwell
tag=gravwell json foo | first foo
```

## Caveats

The first and last module operate on the timestamp of entries for most queries. However, if you write a query that isn't ordered temporally, first/last will operate simply on the first or last entry seen in the search pipeline. For example:

```gravwell
tag=gravwell json foo | sort by foo | first
```

This query will give you the first entry seen, after sorting by foo.
