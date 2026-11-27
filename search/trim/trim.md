# Trim modules

The six trim modules remove characters from the start or end of an enumerated value. For example, `trimleft` removes a set of characters from the beginning of an enumerated value.

- `trim`: Remove leading and trailing characters found in a cutset.
- `trimleft`: Remove leading characters found in a cutset.
- `trimprefix`: Remove a prefix string.
- `trimright`: Remove trailing characters found in a cutset.
- `trimspace`: Remove leading and trailing whitespace.
- `trimsuffix`: Remove a suffix string.

A cutset is a set of individual Unicode characters, not a substring: `trim foo "abcd"` removes any run of `a`, `b`, `c`, or `d` from both ends of `foo`. Use `trimprefix` or `trimsuffix` to remove an exact string.

Each module operates on one enumerated value. Entries without that enumerated value pass through unchanged. String and byte slice values keep their type; other types, such as numbers or IP addresses, are converted to strings before trimming and stored as strings.

## Supported Options

* `-e <enumerated value>`: Optional. Read the cutset, prefix, or suffix from another enumerated value instead of a string literal. Give either `-e` or a literal, not both. Entries that lack the `-e` enumerated value pass through unchanged. `trimspace` does not support this flag.

## trim

### Syntax

	trim <enumerated value> <cutset>

### Example

	trim foo "abcd"

### Example output

	"amy has queried" -> "my has querie"
	"cab fare"        -> " fare"
	"bad dad"         -> " "

This removes any of the characters `a`, `b`, `c`, and `d` from both ends of `foo`.

### Example using an EV instead of a string literal

	trim -e bar foo

### Example output

If the enumerated value `bar` contains "abcd", the output matches the previous example:

	"amy has queried" -> "my has querie"
	"cab fare"        -> " fare"
	"bad dad"         -> " "

## trimleft

### Syntax

	trimleft <enumerated value> <cutset>

### Example

	trimleft foo "abcd"

### Example output

	"amy has queried" -> "my has queried"
	"cab fare"        -> " fare"
	"bad dad"         -> " dad"

This removes any of the characters `a`, `b`, `c`, and `d` from the start of `foo`.

## trimprefix

### Syntax

	trimprefix <enumerated value> <prefix>

### Example

	trimprefix foo "Watch out!"

### Example output

	"Watch out! A rabid logbot!" -> " A rabid logbot!"
	"Watch logbot dance!"        -> "Watch logbot dance!"

This removes the prefix "Watch out!" from `foo`. Values that do not start with the prefix are unchanged.

## trimright

### Syntax

	trimright <enumerated value> <cutset>

### Example

	trimright foo "abcd"

### Example output

	"amy has queried" -> "amy has querie"
	"cab fare"        -> "cab fare"
	"bad dad"         -> "bad "

This removes any of the characters `a`, `b`, `c`, and `d` from the end of `foo`.

## trimspace

### Syntax

	trimspace <enumerated value>

### Example

	trimspace foo

### Example output

	"\n\n   This EV has much whitespace    " -> "This EV has much whitespace"

This removes all leading and trailing whitespace, including spaces, tabs, and newlines, from `foo`.

## trimsuffix

### Syntax

	trimsuffix <enumerated value> <suffix>

### Example

	trimsuffix foo "logbot!"

### Example output

	"Watch out! A rabid logbot!" -> "Watch out! A rabid "
	"Watch logbot dance!"        -> "Watch logbot dance!"

This removes the suffix "logbot!" from `foo`. Values that do not end with the suffix are unchanged.
