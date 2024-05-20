# Eval

## Introduction

The `eval` module is a general purpose filtering and programming environment that runs a "program" against each entry in the query pipeline. It uses a C-style syntax and supports creating simple filters, as well as arbitrary operations on each entry. 

For example, to filter entries to just those where the enumerated value `foo` is equal to `bar` or `baz`:

```
tag=gravwell json foo | eval ( foo == "bar" || foo == "baz" )
```

Additionally, C-style statements can be used for more advanced operations. For example, to conditionally create an EV:

```
tag=gravwell json foo | eval if (foo == "bar") { other_variable = "hey! foo is bar!"; }
```

The rest of this document describes the syntax and semantics of the eval language. 

## Use cases

Eval is used primarily for advanced filtering, especially multifiltering. For example, in order to filter a field extracted from the `json` module, you would normally use filtering in the `json` module directly:

```
tag=gravwell json foo=="my value"
```

However, if you want to filter "foo" in the example above to "my value" or "my other value", you must use the `eval` module:

```
tag=gravwell json foo | eval ( foo == "my value" || foo == "my other value" )
```

## Lexical elements

### Identifiers

Identifiers name variables and types. For example, in `foo = "bar";`, "foo" is an identifier. Identifiers can contain any number of letters, digits, and the underscore "_" character, and must begin with either a letter or an underscore. No other characters are allowed.

It is possible to create an enumerated value that doesn't follow the identifier syntax rules in eval. When using enumerated values like this, you can wrap the name of the enumerated value in `$( )`. For example: 

```
tag=gravwell json "55 crazy enumerated value!" | eval $(55 crazy enumerated value!) == "value"
```

For example, the following are all valid identifiers.

```
foo
_foo
foo314
foo_bar
```

The identifiers "true" and "false" are always predeclared, and are reserved words.

### Variables

Variables are enumerated values and when declared can be used by other modules in the query pipeline. For example:

```
tag=gravwell eval foo = "bar"; | table foo
```

Additionally, all previously produced enumerated values in the query pipeline are available as typed variables. For example:

```
tag=gravwell json foo | eval foo == "bar"
```

### Persistent Variables

Normally, eval executes against every entry in the query, with no state carried between entries. Eval has support for "persistent" variables, that persist throughout the execution of a query. For example, to use eval to count the number of entries in a query:

```
tag=gravwell
eval
	var count = 0;
	count++;
	output = count;
| table
```

This program will initialize a variable "count" to 0, and the value will persist across entries.

To use a persistent variable, it must be declared with `var <variable name>;`. Optionally, you can initialize the variable to a value with the syntax `var <variable name> = <expression>;` 

### Keywords

The following keywords are reserved and may not be used as identifiers.

```
if
else
for
break		
continue
return

(the remaining keywords are not currently supported, but are reserved)

case
default
switch
```

### Operators

Operators are used for filtering and boolean/arithmetic logic. Operators behave similarly to all C-style programming languages (C, C++, Go, ...) with one caveat -- the bitwise OR operator `|` can only appear inside parenthetical expressions in order to avoid a syntax conflict with the Gravwell query language.

The following operators are supported.

```
+ - * / 	Arithmetic operators
%		Modulo
& | 		Bitwise operators
<< >>		Shift operators
&& || !		Boolean operators
-- ++ 		Postfix operators
< > <= >= == != Equality operators
~ !~		String "contains" operators
```

### Punctuation

Eval uses punctuation similar to the C programming language. For example, `( )` are used to enclose expressions in statements as in `if (foo == "bar")...`.

The following character sequences are reserved punctuation.

```
( )		Expression and function grouping
{ }		Statement and block grouping
,		Function parameter separator
;		Statement separator
$		Alternative enumerated value accessor
[ ]		Array indexing (not currently supported)
. :		Scope and range (not currently supported)
```

### Numbers

Numbers can be expressed as decimal and hexadecimal integers, as well as decimal floating-point. Hexadecimal values always begin with "0x". 

The following examples are valid numbers.

```
5
-5
0x05
5.00
```

### String literals

Strings are any sequence of characters enclosed in double quotes `"`. 

## Variables and types

All enumerated values created in the query pipeline are available in eval as variables. Additionally, all variables created in eval are available as enumerated values to any modules after eval in the query pipeline. 

Since enumerated values can have a different type from entry to entry, eval may behave differently based on the type of a variable at runtime. For example, the filter `foo < 5` will only work if "foo" can be interpreted at runtime as a number. It is not an error for "foo" to be a different type (in this example the boolean expression "foo < 5" would be false if foo is not a number). 

The `type()` built-in function can be used at runtime to determine the type of a variable in a given entry.

`eval` can attempt to cast a variable to another type using the following cast operators.

```
int (equivalent to int64)
uint8
int8
uint16
int16
uint32
int32
uint64
int64
float
string
ip
mac
duration
time
type
bool
```

Additionally, eval will attempt to "promote" types that can be automatically cast. For example, if "foo" contains the string "56", the expression `foo < 3.14` will cause eval to attempt to promote foo to a floating-point number before the expression is evaluated.

## Scope

Eval has no variable scoping. All variables are "global", regardless of if they are declared in a statement block or not.

## Expressions

Expressions are simply computations of values. The behavior of an expression depends on where in a program it is used. 

Expressions follow the C-style form of expressions. For example, `if (foo == "bar")...` contains the boolean expression `foo == "bar"` which is used to compute a true/false value for determining if the body of the if statement should be executed or not. 

The right-hand-side of an assignment is also an expression. For example, `foo = type(bar);` contains the expression `type(bar)`, which is a built-in function that returns a value. 

Expressions can be complex, and contain other expressions, just like in C-style programs. For example, `( a == b || b < 3.14 || ( type(b) == "string" && error == "none" ))` is an expression made up of several, smaller expressions.

```{note}
When a program consists of only an expression, the program as a whole is treated as a filter. If the expression returns false (or returns the "zero" value of the type returned), the entry will be dropped.
```

```{note}
Expressions that operate on enumerated values that don't exist have undefined behavior. If you aren't sure that an EV will exist in each entry being processed, be sure to use the has() built-in to ensure the behavior you intend. For example:

	if (has(myEV)) {
		if (myEV == "foo") {
			...
		}
	}
```

## Statements

Statements control program execution. For example, `if (foo == "bar") { ... }` contains an "if" statement, which will determine how the program is to proceed.

Statements are separated by the `{ }` punctuation in the case of "if" and "for", and semicolons ";" for assignments. Currently, only "if", "for", and assignment statements are supported in eval.

Statements other than "if" and "for" must be ended with a semicolon `;`.

### if statements

"if" statements specify conditional execution of code, according to a boolean expression. If the expression evaluates to "true", the code will be executed. "if" statements can contain an "else" code path as well, for executing when the expression evaluates to "false". 

The boolean expression is always contained in parentheses `( )`, and code blocks are contained in braces `{ }`. For example:

```
if ( foo == "bar" ) {
	output = "foo is bar!";
} else {
	output = "foo is not bar!";
}
```

### for statements

"for" statements specify the repeated execution of a block. "for" statements use the C-style syntax of an initializer, condition, and post statement, and are contained in parentheses `( )` and separated by semicolons `;`. Code blocks are contained in braces `{ }`. For example, to iterate 10 times over a code block:

```
for ( i = 0; i < 10; i++ ) {
	count = i;
}
```

### break statements

A "break" statement stops the execution of the innermost "for" statement it is contained in. "break" can only be used inside a "for" statement code block. Execution continues at the statement immediately after the "for" statement. For example:

```
/* for loop up to 10, but break after 5 iterations */
for ( i = 0; i < 10; i++ ) {
	if ( i == 5 ) {
		break;
	}
}
```

### continue statements

A "continue" statement begins the next iteration of the innermost "for" statement it is contained in. "continue" can only be used inside a "for" statement code block. Execution continues at the beginning of the next iteration of the "for" statement. For example:

```
/* for loop up to 10, but only count the last 8 */
for ( i = 0; i < 10; i++ ) {
	if ( i < 2 ) {
		continue;
	}
	count++;
}
```

### return statements

The `return` keyword ends execution of the eval program and determines if the entry being executed on is dropped or passed in the query pipeline. `return` takes a single expression. Returning `true` will pass the entry, `false`, will drop it. For example:

```
if ( foo == "bar" && bar == "baz" ) {
	return true;
}
return false;
```

## Built-in functions

### Convenience

#### delete

	function delete(key string)

Deletes the given enumerated value.

#### has

	function has(ev string) bool

Returns true if the given EV exists.

#### in

	function in(ev string, string...) bool

Returns true if any of the given strings match the given EV.

Example

```
vegetable = "potato";
hasPotato = in(vegetable, "turnip", "potato", "cabbage");
```

#### unix

	function unix(ev) int

Returns a UNIX time of the given enumerated value.

#### len

	function len(<expression>) int

Return the length of the given expression or enumerated value.

#### log

	function log(<expression>)

Log the given expression to the "gravwell" tag.

#### match

	function match(input string, pattern string) bool

Returns true if the regular expression in pattern matches the given input. Regular expressions use the [RE2 syntax](https://github.com/google/re2/wiki/Syntax).

#### printf

	function printf(format string, <expression>...) string

Formats a string according to the given format and any number of expressions. printf uses the formatting rules and verbs defined in [Golang's fmt.Printf()](https://pkg.go.dev/fmt#hdr-Printing) function.

#### set_data

	function set_data(data <expression>)

Overwrites the entry's DATA field. This is different than creating an enumerated value named DATA.

### String manipulation

#### strings_count

	function strings_count(input string, pattern string) int

Counts the number of non-overlapping instances of pattern in the given input. If pattern is an empty string, strings_count returns 1 + the number of Unicode code points in the input.

#### strings_hasprefix

	function strings_hasprefix(input string, prefix string) bool

Returns true if the string input begins with prefix.

#### strings_hassuffix

	function strings_hassuffix(input string, suffix string) bool

Returns true if the string input ends with suffix.

#### strings_index

	function strings_index(input string, pattern string) int

Returns the index of the first instance of pattern in the input, or -1 if pattern is not in input.

#### strings_replace

	function strings_replace(input string, old string, new sting, n int) string

Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If n is -1, there is no limit on the number of replacements.

#### strings_tolower

	function strings_tolower(input string) string

Returns the input with all letters mapped to their lower case.

#### strings_totitle

	function strings_totitle(input string) string

Returns the input with all letters mapped to their Unicode title case.

#### strings_toupper

	function strings_toupper(input string) string

Returns the input with all letters mapped to their upper case.

#### strings_trim

	function strings_trim(input string, cutset string) string

Returns the input with all leading and trailing code points in cutset removed.

#### strings_trimleft

	function strings_trimleft(input string, cutset string) string

Returns the input with all leading code points in the cutset removed.

#### strings_trimprefix

	function strings_trimprefix(input string, prefix string) string

Returns the input with the leading prefix removed.

#### strings_trimright

	function strings_trimright(input string, cutset string) string

Returns the input with all trailing code points in the cutset removed.

#### strings_trimspace

	function strings_trimspace(input string) string

Returns the input with all leading and trailing whitespace removed.

#### strings_trimsuffix

	function strings_trimsuffix(input string, suffix string) string

Returns the input with the trailing suffix removed.

#### pretty_size

    function pretty_size(input string) string

Converts a number to an abreviated pretty printed size, 1234567 becomes "1.18 MB".

#### pretty_count

    function pretty_count(input string) string

Converts a number to an abreviated pretty printed magnitude, 1234567 becomes "1.24 M".

#### pretty_count

    function pretty_count(input string) string

Converts a number to an abreviated pretty printed magnitude, 1234567 becomes "1.24 M".

#### pretty_rate

    function pretty_rate(number, duration) string

Converts a number to an abreviated pretty printed rate in bytes, kilobytes, or megabytes per second given a magnitude and duration; "pretty_rate(1234567, "2s")" becomes "588.87 KB/s".

#### pretty_line_rate

    function pretty_line_rate(number, duration) string

Converts a number to an abreviated pretty printed line rate in bits, kilobits, and megabits per second given a magnitude and duration; "pretty_line_rate(1234567, "2s")" becomes "4.71 Mb/s".

### Hash

#### hash_md5

	function hash_md5(input <expression>) string

Returns the MD5 sum of the given input.

#### hash_sha1

	function hash_sha1(input <expression>) string

Returns the SHA1 sum of the given input.

#### hash_sha256

	function hash_sha256(input <expression>) string

Returns the SHA256 sum of the given input.

#### hash_sha512

	function hash_sha512(input <expression>) string

Returns the SHA512 sum of the given input.


### Path

#### path_base

	function path_base(path string) string

Returns the last elements of the given path. Trailing slashes are removed before extracting the last element.

#### path_clean

	function path_clean(path string) string

Returns the shortest path name equivalent to the given path. For example, "/../a/b/../././/c" would return as "/a/c".

#### path_dir

	function path_dir(path string) string

Returns all but the last element of path.

#### path_ext

	function path_ext(path string) string

Returns the file name extension used by path.

#### path_isabs

	function path_isabs(path string) bool

Returns true if the path is absolute (no relative path elements).

#### path_join

	function path_join(string...) string

Joins any number of path elements into a single path, adding slashes between elements.

#### path_match

	function path_match(pattern string, name string) bool

Returns true if the name matches the given pattern. 

The pattern syntax is identical to [Golang's path.Match()](https://pkg.go.dev/path#Match) function.


### Random

#### rand_float

	function rand_float() float

Returns a random 64-bit float in the interval [0.0,1.0)

#### rand_int

	function rand_int() int

Returns a randomm, non-negative, 64-bit integer.

#### rand_intn

	function rand_intn(n int) int

Returns a random, non-negative, 64-bit integer in the range [0,n).

#### rand_normal

	function rand_normal() float

Returns a normally distributed 64-bit float with standard normal distribution (mean = 0, stddev = 1). To produce a different normal distribution, you can adjust the output using:

```
sample = rand_normal() * desiredStdDev + desiredMean;
```

### Encoding

#### base64_decode

	function base64_decode(input string) string

Returns the decoded form of the given base64 input.

#### base64_encode

	function base64_encode(input string) string

Returns the encoded base64 form of the given input.

#### hex_decode

	function hex_decode(input string) string

Returns the decoded form of the given hexadecimal string.

#### hex_encode

	function hex_encode(input string) string

Returns the encoded hexadecimal form of the given input.

#### html_escape

	function html_escape(input string) string

Returns the input with special characters like "<" changed to their escaped HTML form. It escapes only five characters: <, >, &, ', and ".

#### html_unescape

	function html_unescape(input string) string

Returns the input with escaped HTML entities such as "\&lt;" in their unescaped form.

#### json

	function json(input string) bool

Returns true if the given input is valid JSON.

#### json_append

	function json_append(array string, value <expression>) string

Appends the value to the JSON array given in input.

#### json_array

	function json_array(value <expression>) string

Returns a JSON array of the given value. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.


#### json_get

	function json_get(object string, key string) boo/float/string

Returns a typed item from the given object with the given key.

#### json_index

	function json_index(array string, index int) bool/float/string

Returns a typed item from the given array at the given index.

#### json_len

	function json_len(array string) int

Returns the length of the given JSON array.

#### json_object

	function json_object(key string, value <expression>) string

Returns a JSON object of the given key/value pair. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.

#### json_pretty

	function json_pretty(input string) string

Pretty prints the given JSON input.

#### json_set

	function json_set(object string, key string, value <expression>) string

Sets a key/value pair in the given object. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.

### Math

```{note}
Some math functions retain their legacy function names for backwards compatability.
```

#### ceil

	function ceil(x float) float

Returns the least integer value greater than or equal to x.

#### floor

	function floor(x float) float

Returns the greatest integer value less than or equal to x.

#### math_abs

	function math_abs(x float) float

Returns the absolut value of x.

#### math_ceil

	function math_ceil(x float) float

Same as ceil(). Returns the least integer value greater than or equal to x.

#### math_floor

	function math_floor(x float) float

Same as floor(). Returns the greatest integer value less than or equal to x.

#### math_log

	function math_log(x float) float

Returns the natural logarithm of x.

#### math_log10

	function math_log10(x float) float

Returns the decimal logarithm of x.

#### math_log2

	function math_log2(x float) float

Returns the binary logarithm of x.

#### math_max

	function math_max(x float, y float) float

Returns the larger of x or y.

#### math_min
### Math

```{note}
Some math functions retain their legacy function names for backwards compatability.
```

#### ceil

	function ceil(x float) float

Returns the least integer value greater than or equal to x.

#### floor

	function floor(x float) float

Returns the greatest integer value less than or equal to x.

#### math_abs

	function math_abs(x float) float

Returns the absolut value of x.

#### math_ceil

	function math_ceil(x float) float

Same as ceil(). Returns the least integer value greater than or equal to x.

#### math_floor

	function math_floor(x float) float

Same as floor(). Returns the greatest integer value less than or equal to x.

#### math_log

	function math_log(x float) float

Returns the natural logarithm of x.

#### math_log10

	function math_log10(x float) float

Returns the decimal logarithm of x.

#### math_log2

	function math_log2(x float) float

Returns the binary logarithm of x.

#### math_max

	function math_max(x float, y float) float

Returns the larger of x or y.

#### math_min

	function math_min(x float, y float) float

Returns the smaller of x or y.

#### math_mod

	function math_mod(x float, y float) float

Returns the floating-point remainder of x/y.

#### math_pow

	function math_pow(x float, y float) float

Returns the base-x exponential of y.

#### math_remainder

	function math_remainder(x float, y float) float

Returns the IEEE 754 floating-point remainder of x/y.

#### math_round

	function math_round(x float) float

Same as round(). Returns the nearest integer, rounding half away from zero.

#### math_roundtoeven

	function math_roundtoeven(x float) float

Returns the nearest integer, rounding ties to even.

#### math_sqrt

	function math_sqrt(x float) float

Returns the square root of x.

#### math_trunc

	function math_trunc(x float) float

Returns the integer value of x.

#### round

	function round(x float) float

Returns the nearest integer, rounding half away from zero.


## Acceleration and eval

Eval is capable of producing acceleration hints for the fulltext acceleration engine when used as a filter. For example:

```
tag=gravwell syslog Appname | eval ( Appname == "webserver" || Appname == "indexer" )
```

This query will provide the acceleration hints of "webserver" OR "indexer" to the fulltext engine. By using eval, complex acceleration hints can be composed that aren't possible using extraction modules alone.

## Syntax

The eval syntax is expressed using a [variant](https://github.com/gravwell/pbpg) of Extended Backus-Naur Form (EBNF):

```
Program                  = ( "(" Expression ")" EOF ) | ( "(" Vars StatementList ")" EOF ) | ( "(" StatementList ")" EOF ) | ( "(" Assignment ")" EOF ) | ( Expression EOF ) | ( Vars StatementList EOF ) | ( StatementList EOF ) | ( Assignment EOF )
Vars                     = VarSpec { VarSpec }
VarSpec                  = "var" VarSpecAssignment { "," VarSpecAssignment } ";"
VarSpecAssignment        = AssignmentIdentifier [ "=" Expression ]
StatementList            = Statement { Statement }
Statement                = ( "if" "(" Expression ")" Statement "else" Statement ) | ( "if" "(" Expression ")" Statement ) | ( "for" "(" Assignment ";" Expression ";" Assignment ")" "{" StatementList "}" ) | "{" StatementList "}" | Function ";" | Assignment ";" | "return" Expression ";" | "break" ";" | "continue" ";" | ";"
Assignment               = ( AssignmentIdentifier "=" Expression ) | Expression
Expression               = ( LogicalOrExpression "?" Expression ":" LogicalOrExpression ) | LogicalOrExpression
LogicalOrExpression      = LogicalAndExpression { LogicalOrOp LogicalAndExpression }
LogicalAndExpression     = InclusiveOrExpression { LogicalAndOp InclusiveOrExpression }
InclusiveOrExpression    = ExclusiveOrExpression { BitwiseOrOp ExclusiveOrExpression }
ExclusiveOrExpression    = AndExpression { ExclusiveOrOp AndExpression }
AndExpression            = EqualityExpression { BitwiseAndOp EqualityExpression }
EqualityExpression       = RelationalExpression { EqualityOp RelationalExpression }
RelationalExpression     = ShiftExpression { RelationalOp ShiftExpression }
ShiftExpression          = AdditiveExpression { ShiftOp AdditiveExpression }
AdditiveExpression       = MultiplicativeExpression { AdditiveOp MultiplicativeExpression }
MultiplicativeExpression = UnaryExpression { MultiplicativeOp UnaryExpression }
UnaryExpression          = UnaryOp PostfixExpression | PostfixExpression
PostfixExpression        = PrimaryExpression [ PostfixOp ]
PrimaryExpression        = NestedExpression | Identifier | Literal
NestedExpression         = ( Function ) | ( Cast "(" Expression ")" ) | ( "(" Expression ")" )
Literal                  = DecimalLiteral | FloatLiteral | StringLiteral | "true" | "false"
Function                 = FunctionName "(" [ Expression { "," Expression } ] ")"
LogicalOrOp              = "||"
LogicalAndOp             = "&&"
ExclusiveOrOp            = "^"
BitwiseOrOp              = "|"
BitwiseAndOp             = "&"
UnaryOp                  = "-" | "!"
PostfixOp                = "++" | "--"
EqualityOp               = "==" | "!=" | "~" | "!~"
RelationalOp             = "<" | ">" | "<=" | ">="
ShiftOp                  = "<<" | ">>"
AdditiveOp               = "+" | "-"
MultiplicativeOp         = "*" | "/" | "%"
Cast                     = "int" | "float" | "string" | "mac" | "ip" | "time" | "duration" | "type" | "bool" | "location" | "byte" | "int8" | "uint8" | "int16" | "uint16" | "int32" | "uint32" | "int64" | "uint64"
```

## Legacy Eval

There is a legacy version of eval that you may still see in older queries. For more details, see the [Legacy eval page](legacy-eval) for reference.
