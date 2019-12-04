# First / Last Modules

The `first` and `last` modules are a convenient way to see specific entries over a time period. For instance, given a collection of syslog messages from a number of sources, one might wish to see the earliest or most recent entry from each individual host, or from each daemon.

The modules use the same syntax. Each optionally takes one or more enumerated value names; if specified, the modules will emit the first/last entry for each combination of values of the enumerated values, similar to the behavior of the unique module.

```
first [enumerated value]...
```
