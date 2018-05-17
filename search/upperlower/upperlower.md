## Upper / Lower

The upper and lower modules convert text to upper or lower case. Simply invoking `upper` (or `lower`) will cause the entry's raw data to be converted to upper case. Invoking the module with a list of one or more enumerated value names will cause those enumerated values to be converted to all upper or lower case. This can be useful to normalize data before passing it to `unique` or other modules.

### Example Usage

This example uses the `upper` module to normalize Shodan data prior to counting.

```
tag=shodan json location.region_code | upper region_code | count by region_code | table region_code count
```