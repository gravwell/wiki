## Enrich

The `enrich` module can add enumerated values to each entry in a pipeline; these values can be specified in the module arguments or can come from a resource. `enrich` can be used to annotate data with constant enumerated values, such as "User=admin", in order to simplify visualization and reporting, pivoting within compound queries, and working with non-temporal data. 

### Supported Options

* `-r`: Optional. The `-r` option requires an argument and specifies the name or UUID of a resource to extract data from.
* `-o`: Optional. Overwrite any existing enumerated values that are specified. Usually combined with `-r`. 

### Enriching with string constants

The simplest use of `enrich` is to specify enumerated values and data within the query. For example, to add an enumerated value "foo" with the value "bar" to every entry, simply specify the name and content of the enumerated value, separated by whitespace:

```
tag=jsondata json val | enrich foo bar | table val foo
```

You can specify multiple enumerated value pairs as well. For example, to create enumerated values "foo" and "bar", each with different content:

```
tag=jsondata json val | enrich foo "my data" bar "my other data" | table val foo bar
```

### Enriching with resources 

The `enrich` module can extract columns from a CSV or lookup table resource. When using resources, only the first row will be used by `enrich` (excluding the column headers for CSV resources). Columns can be specified, and all columns will be used if none are specified. If a column conflicts with an existing enumerated value, it will only be overwritten if using the `-o` flag. 

For example, to use columns "foo" and "bar" from a resource "data":

```
tag=jsondata val | enrich -r data foo bar | table val foo bar
```

Additionally, the name of the enumerated value can be different from the column name by using the `as` keyword. For example, to enrich with the column "foo", but name the resulting enumerated value "bar":

```
tag=jsondata val | enrich -r data foo as bar | table val bar
```

NOTE: Only the first row of resources are used with the `enrich` module (excluding the column headers for CSV resources).
