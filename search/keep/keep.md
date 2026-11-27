# Keep

Keep removes every enumerated value from each entry except the ones you list. It never drops entries.

Use it when you extract enumerated values for filtering or enrichment but don't want them in the final results.

## Arguments

Keep takes one or more enumerated value names. Each name must be produced by an earlier module in the pipeline.

```gravwell
keep foo bar
```

## Supported Options

Keep has no flags.

## Example Usage

This search removes all enumerated values except "bananas" and "potatoes":

```gravwell
tag=data json fruit vegetables
| fields -e fruit [0] as bananas [1] as apples
| fields -e vegetables [0] as broccoli [1] as potatoes
| keep bananas potatoes
```
