## Nosort

By default, everything in the Gravwell search pipeline is temporally sorted (when appropriate).  This means that if you execute the query `tag=gravwell` Gravwell will automatically insert a `sort by time desc` so that the data you see is strictly sorted.

However, there may be times where the extra overhead from the sort may not be required or explicitly not wanted; this is where `nosort` comes into play.  The `nosort` module does nothing but inform Gravwell that you explicitly do not want the data sorted at any stage, it basically turns off the `sort by time` injection.

The `nosort` module is purely for query optimization and is never required, don't use it unless you really know what you are doing.  The `nosort` module has a single optional flag `-asc` that tells Gravwell that you don not care about explicit time sorting but you would like the data read from oldest to newest (roughly).  The `-asc` flag is useful on big aggregate queries because it means that we will likely pull data in the order it was ingested which means the disks are probably moving in a more or less linear pattern.

### Examples

`tag=syslog nosort`

Show syslog entries in a rough ordering, do not strictly sort them by time.

`tag=syslog nosort -asc`

Show syslog entries in a rough ordering from oldest to newest, do not strictly sort them by time.

Note: The `nosort` module collapses the pipeline, this means that if you are running on a cluster environment it may actually reduce the query performance if you place it in the wrong location.
