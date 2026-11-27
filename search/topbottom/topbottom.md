# Top / Bottom Modules

The `top` and `bottom` modules return the N entries with the highest or lowest values of one or more enumerated values. For example, to get the 10 entries with the highest values of `foo`:

```gravwell
tag=foo json foo | top -n 10 foo
```

`top` outputs entries in descending order and `bottom` in ascending order.

Both modules use the same syntax and take optional enumerated value names. If none are given, the `DATA` field is used. With several names, entries are ranked by the first, and ties are broken by the next one, in the order given.

For example, given:

| foo | bar |
|-----|-----|
| 1   | 1   |
| 10  | 9   |
| 10  | 200 |
| 10  | 100 |

and the query:

```gravwell
tag=foo json foo bar | eval foo = int(foo); bar = int(bar); | top -n 2 foo bar
```

The output will be:

| foo | bar |
|-----|-----|
| 10  | 200 |
| 10  | 100 |

Three of the four entries have `foo` equal to 10, so `top` uses `bar` to break the tie and keeps the entries where `bar` is 200 and 100.

Values are compared according to their type: numbers numerically, strings lexically. Extraction modules such as `json` produce strings, and as strings `9` ranks above `200`. The `eval` in the query above casts the values to integers so they are ranked numerically.

`top` and `bottom` are equivalent to sorting with a limit. For example:

```gravwell
tag=foo json foo | top -n 10 foo
```

returns the same entries as

```gravwell
tag=foo json foo | sort by foo desc | limit 10
```

`top` and `bottom` are much faster, because they only hold N entries in memory instead of sorting everything.

## Flags

- `-n <number>`: the number of entries to return. Must be greater than 0. Defaults to 10.

## Examples

To get the 10 entries with the highest `DATA` field:

```gravwell
tag=gravwell top
```

To get the 10 entries with the highest values of `foo`:

```gravwell
tag=gravwell json foo | top foo
```

To get the 300 entries with the lowest values of `foo`, breaking ties with `bar`:

```gravwell
tag=gravwell json foo bar | bottom -n 300 foo bar
```
