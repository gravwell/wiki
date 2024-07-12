# Top / Bottom Modules

The `top` and `bottom` modules show the top or bottom N values of a given enumerated value (or values) or entry. For example, to show the top 10 values of the enumerated value `foo`:

```gravwell
tag=foo json foo | top -n 10 foo
```

The output is sorted in descending order for the `top` module and in ascending order for the `bottom` module.

The modules use the same syntax. Each optionally takes one or more enumerated values. The modules will emit the top/bottom values of the given value. If no enumerated value is provided, the `DATA` field is used.
If multiple enumerated values are provided, the modules will sort by the top/bottom values of the enumerated values in the provided order. If two entries have the same value for a given enumerated value, the second enumerated value will be used.

For example, given:

| foo | bar |
|-----|-----|
| 1   | 1   |
| 10  | 9   |
| 10  | 200 |
| 10  | 100 |

and the query: 

```gravwell
tag=foo json foo bar | top -n 2 foo bar
```

The output will be:

| foo | bar |
|-----|-----|
| 10  | 200 | 
| 10  | 100 |

In the example above, `top` took the top two values of `foo` first, and where `foo` had the same value, took the top value of `bar` instead.

Values must be numeric or able to be casted to a number. Non-numeric values are ignored.

The top module is functionally equivalent to sorting with a limit. For example:

```gravwell
tag=foo json foo | top -n 10 foo
```

will produce the same result as

```gravwell
tag=foo json foo | sort by foo desc | limit 10
```

The `top` and `bottom` modules are however far more performant than using sort/limit.

## Flags

- `-n <number>`: The "-n" flag specifies the number of entries to track. The default is 10. 

## Examples

To get the top 10 values of the DATA field:

```gravwell
tag=gravwell top
```

To get the top 10 values of foo:

```gravwell
tag=gravwell json foo | top foo
```

To get the bottom 300 values of foo, and where foo is equal, bar:

```gravwell
tag=gravwell json foo bar | bottom -n 300 foo bar
```
