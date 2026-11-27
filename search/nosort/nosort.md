# Nosort

When the renderer needs ordered data and no module in the query sorts it, Gravwell automatically adds a `sort by time` (newest first) to the end of the pipeline. For example, `tag=gravwell` becomes `tag=gravwell sort by time`.

The `nosort` module turns off that automatic sort. It passes entries through unchanged; use it when you don't need strict time ordering and want to avoid the cost of the sort.

`nosort` is purely an optimization and is never required. Don't use it unless you know you need it. Its one optional flag, `-asc`, asks Gravwell to read data roughly from oldest to newest instead of newest to oldest. This can help large aggregate queries, since data is then read in roughly the order it was ingested and disk access is more sequential.

## Examples

Show syslog entries in rough time order, without strictly sorting them:

```gravwell
tag=syslog nosort
```

Show syslog entries roughly from oldest to newest, without strictly sorting them:

```gravwell
tag=syslog nosort -asc
```

```{note}
The `nosort` module collapses the pipeline. In a clustered deployment, placing it in the wrong spot can reduce query performance.
```
