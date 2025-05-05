# Reingest

The reingest module allows "reingesting" (copying) entries in the search pipeline to another tag. In addition to the raw entry data, the reingest module attaches all enumerated values as intrinsic enumerated values in copied entry.

```{note}
The `reingest` module will copy all data in the search pipeline to the destination tag. Queries over large datasets can result in large data copies. Use with caution!
```

```{note}
CBAC controls apply to the reingest module. The user invoking the module must have the ingest permission, as well as permissions to ingest into the destination tag.
```

The reingest module can only be used once in a query, and it must be the final module in the query pipeline. All entries that reach the reingest module will be copied to the destination tag. The query pipeline can be used to filter out unwanted entries, enrich with enumerated values, and otherwise transform the dataset before reingesting.

## Arguments

Reingest takes the destination tag as a single argument. For example, to reingest into tag "foo":

```gravwell
tag=gravwell reingest foo
```

## Supported Options

* `-nodata`: Do not include the DATA field of the entry. Enumerated Values are always included.
* `-now`: Reingest using the current time as the entry timestamp, overriding the original timestamp.

## Example Usage

The following example uses the reingest module to create a simple aggregation dataset. It counts the number of entries in all tags, and ingests those counts into a tag named "aggs".

```gravwwell
tag=* count by TAG
| reingest -now -nodata aggs
```
