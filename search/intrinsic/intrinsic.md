# Intrinsic

The `intrinsic` module extracts enumerated values that were created at ingest time to a given search. By default, when using the text or raw renderers, all intrinsic enumerated values are added to the search. When using any other render module, such as table, the intrinsic module must be used. 

## Supported Options

The intrinsic module has no flags.

## Arguments and syntax

The intrinsic module simply takes a list of enumerated values to extract, and optionally a filter for each enumerated value. 

For example, to extract the enumerated values "foo" and "bar":

```gravwell
tag=data intrinsic foo bar | table
```

Additionally, to filter "foo" to just entries where foo is equal to "potato":

```gravwell
tag=data intrinsic foo == "potato" bar | table
```
