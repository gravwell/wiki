# Reingest

The reingest module copies entries in the search pipeline into another tag. Along with the entry DATA, it attaches all enumerated values to the copy as intrinsic enumerated values. A common use is building data aggregates and summary indexes.

```{warning}
The `reingest` module copies all data in the search pipeline to the destination tag. Queries over large datasets can produce large copies. Use with caution!
```

```{note}
CBAC controls apply to the reingest module. The user must have the ingest permission and permission to ingest into the destination tag.
```

The reingest module has these restrictions:

* It must be the last module in the query, so it can be used only once.
* The query must use the text or raw renderer. If no renderer is given, text is used.
* It cannot be used in the inner query of a compound query.
* The destination tag cannot be one of the tags the query reads from.

Every entry that reaches the reingest module is copied to the destination tag, and also passed on to the renderer. Use the earlier modules in the query to filter out unwanted entries, add enumerated values, and otherwise transform the data before reingesting.

To control which enumerated values are copied, restrict extractions or use the [keep](/search/keep/keep) module. To replace the DATA field, use the [set_data](#set-data-target) function in [eval](/search/eval/eval).

## Arguments

Reingest takes a single argument: the destination tag. For example, to reingest into tag "foo":

```gravwell
tag=gravwell reingest foo
```

## Supported Options

* `-nodata`: Do not include the DATA field of the entry. Enumerated values are always included.
* `-now`: Use the current time as the timestamp of the copied entries instead of the original timestamp.

## Example Usage

This example builds a simple aggregation dataset. It counts the entries in each tag and ingests those counts into the tag "aggs":

```gravwell
tag=* count by TAG
| reingest -now -nodata aggs
```

This example uses eval to replace the DATA field of entries before reingesting them:

```gravwell
tag=gravwell syslog Hostname Appname
| eval
    set_data( printf("hostname is %v, appname is %v", Hostname, Appname) );
| reingest foo
```
