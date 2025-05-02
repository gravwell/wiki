# Keep

The keep module is a simple module that removes all unspecified enumerated values from entries. Keep does not drop entries.

The keep module is useful for scenarios where extracting certain enumerated values is needed for filtering or enrichment, but not wanted in the final dataset.

## Arguments

Keep requires one or more enumerated value names, such as:

```gravwell
keep foo bar
```

## Supported Options

The keep module has no flags.

## Example Usage

The following example removes all enumerated values except "bananas" and "potatoes" from all entries:

```gravwwell
tag=data json fruit vegetables
| fields -e fruit [0] as bananas [1] as apples
| fields -e vegetables [0] as broccoli potatoes
| keep bananas potatoes
```
