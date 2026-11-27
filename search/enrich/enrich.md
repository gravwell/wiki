# Enrich

The `enrich` module adds enumerated values to every entry in the pipeline. The values come either from the module arguments or from a resource. Use it to annotate data with constant values, such as "User=admin", to simplify visualization and reporting, pivoting within compound queries, and working with non-temporal data.

## Supported Options

* `-r <arg>`: Optional. The name or UUID of a resource to extract data from. This can also be an inner query of a compound query, such as `@foo`.
* `-o`: Optional. Overwrite existing enumerated values with the same name. Without `-o`, an entry that already has an enumerated value of that name keeps its existing value.

## Enriching with string constants

The simplest use of `enrich` is to give enumerated value names and contents in the query, as name/value pairs separated by whitespace. For example, to add an enumerated value "foo" with the value "bar" to every entry:

```gravwell
tag=jsondata json val | enrich foo bar | table val foo
```

You can specify multiple pairs. For example, to create enumerated values "foo" and "bar", each with different content:

```gravwell
tag=jsondata json val | enrich foo "my data" bar "my other data" | table val foo bar
```

## Enriching with resources 

The `enrich` module can extract columns from a CSV or lookup table resource. List the columns to use; if none are listed, all columns are used. If the resource has no data rows, entries pass through unchanged. If a column conflicts with an existing enumerated value, it is only overwritten with the `-o` flag.

For example, to use columns "foo" and "bar" from a resource "data":

```gravwell
tag=jsondata json val | enrich -r data foo bar | table val foo bar
```

Use `as` to give the enumerated value a different name from the column. For example, to enrich with the column "foo" but name the resulting enumerated value "bar":

```gravwell
tag=jsondata json val | enrich -r data foo as bar | table val bar
```

```{note}
Only the first row of a resource is used (excluding the column headers for CSV resources).
```
