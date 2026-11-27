# Intrinsic

The `intrinsic` module extracts intrinsic enumerated values, which are attached to entries at ingest time, so later modules in the search can use them. The text and raw renderers, and the table renderer when given no column arguments, show all intrinsic enumerated values automatically. Everywhere else, extract them with `intrinsic` first.

`intrinsic` can also be used as the module for an [autoextractor](/configuration/autoextractors).

## Supported Options

The intrinsic module has no flags.

## Arguments and syntax

The intrinsic module takes a list of enumerated values to extract, each optionally followed by a filter.

For example, to extract the enumerated values "foo" and "bar":

```gravwell
tag=data intrinsic foo bar | table
```

To also keep only entries where "foo" is equal to "potato":

```gravwell
tag=data intrinsic foo == "potato" bar | table
```

The supported filter operators are:

| Operator | Name |
|----------|------|
| == | Equal |
| != | Not equal |
| ~ | Subset |
| !~ | Not subset |

With the `==` and `~` operators, entries that lack the enumerated value are dropped.
