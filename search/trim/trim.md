# Trim modules

The trim modules are six modules that modify given string-based enumerated values in various ways. For example, you can use `trimleft` to remove a given list of characters from the beginning of a given enumerated value.

The trim module can use either a string literal as the cutset/prefix/suffix, or an EV can be specified using the `-e` flag.

The trim modules consist of:

- `trim`: Remove leading and trailing Unicode code points from an enumerated value.
- `trimleft`: Remove leading Unicode code points from an enumerated value.
- `trimprefix`: Remove a given prefix from an enumerated value.
- `trimright`: Remove trailing Unicode code points from an enumerated value.
- `trimspace`: Remove all leading and trailing whitespace from an enumerated value.
- `trimsuffix`: Remove a given suffix from an enumerated value.

## Supported Options

* `-e`: Optional. Use the contents of an EV, specified with `-e`, as the cutset, prefix, or suffix. `trimspace` does not support this flag.

## trim

### Syntax

	trim <enumerated value> <code points>

### Example

	trim foo "abcd"

### Example output

	"amy has queried" -> "my has querie"
	"cab fare"        -> " fare"
	"bad dad"	  -> " "

This example will remove any of the leading or trailing characters "abcd" from the EV. 

### Example using an EV instead of a string literal

	trim -e bar foo 

### Example output

This assumes the EV "bar", has the contents "abcd"

	"amy has queried" -> "my has querie"
	"cab fare"        -> " fare"
	"bad dad"	  -> " "

## trimleft

### Syntax

	trimleft <enumerated value> <code points>

### Example

	trimleft foo "abcd"

### Example output

	"amy has queried" -> "my has queried"
	"cab fare"        -> " fare"
	"bad dad"	  -> " dad"

This example will remove any of the leading characters "abcd" from the EV. 

## trimprefix

### Syntax

	trimprefix <enumerated value> <prefix>

### Example

	trimprefix foo "Watch out!"

### Example output

	"Watch out! A rabid logbot!" -> " A rabid logbot!"
	"Watch logbot dance!"        -> "Watch logbot dance!"

This example will remove the prefix "Watch out!" from the EV. 

## trimright

### Syntax

	trimright <enumerated value> <code points>

### Example

	trimright foo "abcd"

### Example output

	"amy has queried" -> "amy has querie"
	"cab fare"        -> "cab fare"
	"bad dad"	  -> "bad "

This example will remove any of the trailing characters "abcd" from the EV. 

## trimspace

### Syntax

	trimspace <enumerated value>

### Example

	trimspace foo

### Example output

	"
         
         This EV has much whitespace    " -> "This EV has much whitesapce"

This example will remove all leading and trailing whitespace from the EV.

## trimsuffix

### Syntax

	trimsuffix <enumerated value> <suffix>

### Example

	trimpsuffix foo "logbot!"

### Example output

	"Watch out! A rabid logbot!" -> "Watch out! A rabid"
	"Watch logbot dance!"        -> "Watch logbot dance!"

This example will remove the suffix "logbot!" from the EV. 

