# Upper / Lower

The `upper` and `lower` modules convert text to upper or lower case. With no arguments, they convert the entry's raw data. Given one or more enumerated value names, they convert those enumerated values instead. Only string and byte slice enumerated values are converted; other types, and enumerated values missing from an entry, are left alone.

Converting case is useful for normalizing data before `unique`, `count`, or other modules that compare values.

## Example Usage

This example uses the `upper` module to normalize Shodan data prior to counting.

```gravwell
tag=shodan json location.region_code | upper region_code | count by region_code | table region_code count
```
