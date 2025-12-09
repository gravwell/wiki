# Regex Replace Preprocessor

The Regex Replace preprocessor performs regular expression-based find and replace operations on entry data. This is useful for sanitizing sensitive data, normalizing log formats, or transforming data before ingestion.

The Regex Replace preprocessor Type is `regexreplace`.

## Supported Options

* `Regex` (string, required): The regular expression pattern to match against entry data. Supports standard Go regular expression syntax including named capture groups.
* `Replacement` (string, required): The replacement string. Can reference capture groups using `$1`, `$2`, etc. for numbered groups or `${name}` for named groups.
* `Case-Sensitive` (boolean, optional): When set to `true`, the regex matching is case-sensitive. When `false` (the default), matching is case-insensitive.

## Common Use Cases

The regexreplace preprocessor is commonly used for:

* Sanitizing sensitive data 
* Normalizing log formats across different sources
* Stripping or replacing unwanted characters or patterns
* Transforming data for easier downstream processing

### Example: Redacting Numbers

To redact all numbers from log entries (e.g., for removing phone numbers, IDs, or other sensitive numeric data):

```
Phone: 123-456-7890, Age: 25, ID: 987654321
```

Use the following configuration:

```
[Preprocessor "redact-numbers"]
	Type=regexreplace
	Regex=`\d+`
	Replacement=`REDACTED`
```

The result is:

```
Phone: REDACTED-REDACTED-REDACTED, Age: REDACTED, ID: REDACTED
```

### Example: Case-Insensitive Replacement

To replace all occurrences of a word regardless of case:

```
This is a TEST string with Test words
```

Use the following configuration:

```
[Preprocessor "normalize-test"]
	Type=regexreplace
	Regex=`test`
	Replacement=`example`
	Case-Sensitive=false
```

The result is:

```
This is a example string with example words
```

### Example: Modifying JSON Fields

Given JSON log data where you want to modify a specific field:

```
{"name":"john","age":30,"city":"new york"}
```

Use named capture groups to extract and modify the value:

```
[Preprocessor "modify-json"]
	Type=regexreplace
	Regex=`"name":"(?P<n>[^"]*)"`
	Replacement=`"name":"${n}_modified"`
	Case-Sensitive=true
```

The result is:

```
{"name":"john_modified","age":30,"city":"new york"}
```

