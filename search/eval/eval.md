# Eval

## Introduction

The `eval` module runs a small C-style program against each entry in the pipeline. A program can be a simple filter or a sequence of statements that read and create enumerated values.

For example, to keep only entries where the enumerated value `foo` is `bar` or `baz`:

```
tag=gravwell json foo | eval ( foo == "bar" || foo == "baz" )
```

Statements allow more involved work. For example, to create an EV only when a condition holds:

```
tag=gravwell json foo | eval if (foo == "bar") { other_variable = "hey! foo is bar!"; }
```

## Use cases

Eval is used mostly for filters that extraction modules can't express on their own. To filter a single field extracted by the `json` module, you would normally filter in the `json` module directly:

```
tag=gravwell json foo=="my value"
```

To keep entries where `foo` is either "my value" or "my other value", use `eval`:

```
tag=gravwell json foo | eval ( foo == "my value" || foo == "my other value" )
```

## Examples

### Example: Filtering with multiple possible values

This example uses the `in()` function to keep syslog entries whose Appname is "indexer" or "webserver".

```
tag=gravwell
  syslog Appname
| eval in(Appname, "indexer", "webserver")
| table
```

### Example: JSON encoding

This example creates a JSON object of several EVs and stores the object in another EV.

```
tag=gravwell
syslog Appname Hostname Message
| eval
  output = "{}"; // an empty JSON object
  output = json_set(output, "Appname", Appname);
  output = json_set(output, "Hostname", Hostname);
  output = json_set(output, "Message", Message);
| table
```

### Example: Loops

This example calculates e⁵ using the Taylor series for eᶻ. It ignores the entry's contents and exists only to show loops and flow control.

```
tag=gravwell limit 1

| eval

// Taylor series of eᶻ is ∑zⁿ/n! from 0 to ∞

z = 5.0; // exponent
ez = 0.0; // starting value

// An int64 cannot hold a value much larger than 20!, so we are forced 
// to stop at 20 iterations. That still gives us an accurate result to 
// about 5 decimal places. 
for (i = 0; i < 20; i++) {
    // zⁿ
    n = math_pow(z,i);

    // n!
    // eval doesn't have a factorial function built-in, 
    // so we implement it with a for loop here. There are 
    // much better ways to do this, but we want to 
    // illustrate flow control and nested loops.
    if (i == 0 || i == 1) {
        d = 1;
    } else {
        d = 1;
        for (j = 2; j <= i; j++) {
            d = d*j;
        }
    }

    // next value of z
    zn = n/d;

    ez = ez + zn;
}
| eval result = printf("e⁵ == %v", ez);
| table result
```

### Example: Calculating leg currents of a 3-phase PDU

This example calculates the leg currents of a 3-phase PDU that logs one phase per entry. It uses a persistent map, comments, and math functions.

```
tag=pdus 
  ax phase value
| eval

// THIS IS FOR 8.6kVA 3-phase PDUs. Other ratings need changes.

map linepairs;

Vnominal = 208.0;
I = value/Vnominal;

// save the phase currents in our map
linepairs[phase] = I;

// grab the last known currents for the other phases
if ( phase == "L1-L2" ) {
    l1l2 = I;
    l2l3 = linepairs["L2-L3"];
    l3l1 = linepairs["L3-L1"];
} else if ( phase == "L2-L3" ) {
    l1l2 = linepairs["L1-L2"];
    l2l3 = I;
    l3l1 = linepairs["L3-L1"];
} else if ( phase == "L3-L1" ) {
    l1l2 = linepairs["L1-L2"];
    l2l3 = linepairs["L2-L3"];
    l3l1 = I;
}

// calculate leg currents
ϕr1 = -0.5;   // cos(120)
ϕj1 = 0.866;  // sin(120)
ϕr2 = -0.5;   // cos(-120)
ϕj2 = -0.866; // sin(-120)

// leg current is the magnitude of the phasor difference of the leg pairs on this leg:
// leg1 = | L1L2∠120 - L3L1∠-120 |
Ileg1 = math_sqrt( math_pow( (ϕr1*l1l2 - ϕr2*l3l1), 2 ) + math_pow( (ϕj1*l1l2 - ϕj2*l3l1) , 2 ) );
Ileg2 = math_sqrt( math_pow( (ϕr1*l2l3 - ϕr2*l1l2), 2 ) + math_pow( (ϕj1*l2l3 - ϕj2*l1l2) , 2 ) );
Ileg3 = math_sqrt( math_pow( (ϕr1*l3l1 - ϕr2*l2l3), 2 ) + math_pow( (ϕj1*l3l1 - ϕj2*l2l3) , 2 ) );

Imax = 24.0;

Ptotal = (l1l2 + l2l3 + l3l1)*Vnominal; 

| chart Ileg1 Ileg2 Ileg3 Imax Ptotal
```

## Lexical elements

### Identifiers

Identifiers name variables. In `foo = "bar";`, `foo` is an identifier. Identifiers contain letters (including Unicode letters), digits, and underscores, and must begin with a letter or an underscore. The following are all valid identifiers:

```
foo
_foo
foo314
foo_bar
```

The identifiers `true` and `false` are predeclared and reserved.

To use an enumerated value whose name isn't a valid identifier, wrap the name in `$( )`:

```
tag=gravwell json "55 crazy enumerated value!" | eval $(55 crazy enumerated value!) == "value"
```

### Variables

Variables are enumerated values. A variable assigned in eval is available to later modules in the pipeline:

```
tag=gravwell eval foo = "bar"; | table foo
```

Enumerated values produced earlier in the pipeline are available as typed variables, as are the entry's `DATA`, `SRC`, `TAG`, `TIMESTAMP`, and `NOW` values:

```
tag=gravwell json foo | eval foo == "bar"
```

Reading a variable that no earlier module produces, and that the program hasn't assigned earlier in its text, is a parse error.

### Persistent Variables

Normally no state carries over from one entry to the next. Persistent variables keep their value for the whole query. For example, to number the entries in a query:

```
tag=gravwell
eval
	var count = 0;
	count++;
	output = count;
| table
```

This program initializes `count` to 0 once, and the value carries over from entry to entry.

Declare a persistent variable with `var <name>;` or `var <name> = <expression>;`. Declarations must come at the start of the program, and one `var` can declare several variables separated by commas. A variable declared without a value starts as an empty string. The initializer runs once, when the query starts, so it can't refer to enumerated values or other persistent variables.

Persistent variables are not attached to entries. To use a persistent variable's value outside of eval, assign it to a regular variable.

A program that declares persistent variables or maps collapses the query: eval and every module after it run on the webserver, so the state covers entries from all indexers.

### Persistent Maps

Persistent maps are key/value stores that also last for the whole query. Keys are strings, and values can be any type except a map. Declare a map with `map <name>;` at the start of the program, and read or write a key with `name[key]`.

For example, to keep a running count of each Appname in syslog entries:

```
tag=gravwell
  syslog Appname
| eval
    map appnames;
    if (has_key(appnames, Appname)) {
        appnames[Appname] = appnames[Appname] + 1;  // The key exists. Increment.
    } else {
        appnames[Appname] = 1;                      // The key doesn't exist. Create it.
    }
    seen = appnames[Appname];
| table Appname seen
```

Like persistent variables, maps are not attached to entries. Assign a value to a regular variable to use it outside of eval.

A map holds at most 1,000,000 keys. Once a map grows past that limit, any assignment into it stops the program and drops the entry.

### Arrays

Eval supports the `array` enumerated value type. Arrays currently can only be created in the eval and regex modules and appear to other modules as string representations of the array contents (for example, `[ apple orange banana ]`).

#### Declaring arrays

Declare an array with `{ }` when assigning to a variable:

```
foo = {1,2,3,4};
```

This creates a four-element array containing the integers 1, 2, 3, and 4.

Arrays can mix types:

```
foo = {1, 2, "my string!"};
```

Arrays can contain other arrays:

```
foo = {1, 2, {"foo", "bar"}};
```

#### Indexing arrays

Index an array with `[ ]`. Indexes start at 0:

```
foo = {"a","b","c","d"};
bar = foo[1];
```

Here `bar` contains the string "b".

Assign into an array the same way:

```
foo = {"a","b","c","d"};
foo[0] = "bar";
```

This replaces the first element, "a", with "bar".

#### Ranged sub-arrays

Take a sub-array with the `:` range operator:

```
foo = {"a","b","c","d"};
bar = foo[1:3];
```

Here `bar` contains the elements from index 1 up to, but not including, index 3: `[ "b" "c" ]`.

#### Array operators on strings

Indexing also works on strings, byte by byte:

```
foo = "I am a potato";
foo[7] = "P";
```

This changes the "p" at index 7 to "P". Assigning into a string requires a single-character value.

#### Example: Iterating over an array

```
tag=data json foo
| eval
    k = json_to_gravwell_array(foo);
    for (i=0; i<len(k); i++) {
        k[i] = strings_tolower(k[i]);
    }
```

This converts the JSON array in `foo` to a Gravwell array, `k`, and replaces each element with its lowercase form.

### Keywords

The following keywords are reserved and may not be used as identifiers.

```
if
else
for
break
continue
return
case
default
switch
var
```

`var` and `map` start persistent variable and map declarations.

### Operators

Operators behave as in other C-style languages, with one exception: the bitwise OR operator `|` can only appear inside parentheses or braces, so that it isn't mistaken for the query pipe.

```
+ - * /         Arithmetic operators (+ also concatenates two strings)
%               Modulo
& | ^           Bitwise AND, OR, and XOR
<< >>           Shift operators
&& || !         Boolean operators
-- ++           Postfix increment and decrement
< > <= >=       Relational operators
== !=           Equality operators
~ !~            Contains and does-not-contain: a ~ b is true if a contains b
? :             Ternary conditional: cond ? a : b
```

### Punctuation

Eval uses C-style punctuation. The following character sequences are reserved:

```
( )             Expression and function grouping
{ }             Statement grouping and array literals
,               Function argument separator
;               Statement separator
$               Alternative enumerated value accessor
[ ]             Array, string, and map indexing
:               Array range operator, ternary operator, and case label
.               Scope (not currently supported)
```

### Numbers

Numbers can be decimal or hexadecimal integers, or decimal floating-point values. Hexadecimal values begin with "0x". The following are valid numbers:

```
5
-5
0x05
5.00
```

### String literals

Strings are any sequence of characters enclosed in double quotes `"`.

## Variables and types

Enumerated values from earlier in the pipeline are variables in eval, and variables created in eval are enumerated values for later modules.

An enumerated value can have a different type in each entry, so eval's behavior can depend on a variable's type at runtime. For example, `foo < 5` only works if `foo` can be read as a number. A different type isn't an error, but the result of the comparison is undefined. See [Failed Type Promotions](#eval-failed-type-promotions).

The `type()` cast returns the name of a value's type in the current entry: `bool`, `byte`, `int8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `float`, `mac`, `ip`, `ts`, `duration`, `location`, `array`, or `string`. If its argument fails, for example a cast that can't convert its input, `type()` returns `nil`.

Eval can attempt to convert a value to another type with the following casts:

```
int (equivalent to int64)
byte (equivalent to uint8)
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
location
array
```

Eval also "promotes" values to a common type before comparing or combining them. For example, if `foo` contains the string "56", `foo < 3.14` converts `foo` to a floating-point number before comparing.

## Scope

Eval has no variable scoping. All variables are global, whether or not they are assigned inside a block.

## Expressions

Expressions compute values, as in C. In `if (foo == "bar") ...`, the boolean expression `foo == "bar"` decides whether the body runs. The right-hand side of an assignment is also an expression: in `foo = type(bar);` it is `type(bar)`.

Expressions nest. For example, `( a == b || b < 3.14 || ( type(b) == "string" && error == "none" ))` combines several smaller expressions.

```{note}
When a program consists of only an expression, the program is a filter. If the expression returns false (or the "zero" value of its type), the entry is dropped.
```

```{note}
Expressions that operate on enumerated values that don't exist have undefined behavior. If an EV might be missing from some entries, check for it with `has()` first:

	if (has(myEV)) {
		if (myEV == "foo") {
			...
		}
	}
```

### Type Promotion in Expressions

When an operator's two operands have different types, eval converts them to a common type first. In `1 + 2.14`, the integer is promoted to a floating-point value and the result is a float.

Strings take part in promotion too. In `5 + "10"`, eval converts the string to the integer 10, and the result is 15. When both operands are strings, `+` concatenates them: `"5" + "10"` is `"510"`.

Not all types are compatible. Eval picks a common type using the following rules, in order:

* If either is a Location, try to convert the other to a Location.
* If either is an IP, try to convert the other to an IP.
* If either is a MAC, try to convert the other to a MAC.
* If either is a timestamp, try to convert the other to a timestamp.
* If either is a float, try to convert the other to a float.
* If either is an integer, try to convert the other to an integer.

If none of these cases apply, or if either value is a duration, the values keep their original types.

Timestamps and durations have three special cases:

* A timestamp plus a duration results in a timestamp.
* A timestamp minus a duration results in a timestamp.
* A timestamp minus a timestamp results in a duration.

When type promotion results in two integers, or two integers are provided in an expression, additional promotion logic is applied, conforming to a modified form of the [ISO/IEC 9899:2011 C standard](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1548.pdf) rank rules:

* No two signed integer types shall have the same rank, even if they have the same representation.
* The rank of a signed integer type shall be greater than the rank of any signed integer type with less precision.
* An int is equivalent to an int64.
* The rank of int64 shall be greater than the rank of int32, which shall be greater than the rank of int16, which shall be greater than the rank of int8.
* The rank of any unsigned integer type shall equal the rank of the corresponding signed integer type, if any.
* The rank of byte shall equal the rank of uint8.
* The rank of bool shall be less than the rank of all other integer types.
* For all integer types T1, T2, and T3, if T1 has greater rank than T2 and T2 has greater rank than T3, then T1 has greater rank than T3.

After integer promotion, the "usual arithmetic conversions" apply, in order:

* If both operands have the same type, no further conversion is needed.
* If both operands are of the same integer type (signed or unsigned), the operand with the type of lesser integer conversion rank is converted to the type of the operand with greater rank.
* If the operand that has unsigned integer type has rank greater than or equal to the rank of the type of the other operand, the operand with signed integer type is converted to the type of the operand with unsigned integer type.
* If the type of the operand with signed integer type can represent all of the values of the type of the operand with unsigned integer type, the operand with unsigned integer type is converted to the type of the operand with signed integer type.
* Otherwise, both operands are converted to the unsigned integer type corresponding to the type of the operand with signed integer type.

(eval-failed-type-promotions)=
#### Failed Type Promotions

When type promotion fails, the result of the operation is undefined. This matters most in conditions. Check types with `type()` and casts before comparing. For example, given this program:

```
x = "foo";
y = 5;

if ( x < y ) {
    output = "x is less than y";
}
```

The comparison result is undefined, and the `output =` statement may run. To guard against this, check the cast first:

```
x = "foo";
y = 5;

if ( type(int(x)) == "int64" ) { // The cast succeeded, we can do the comparison
    if ( x < y ) {
        output = "x is less than y";
    }
} else {
    output = "x isn't a number";
}
```

## Statements

Eval supports `if`, `switch`, `for`, `break`, `continue`, and `return` statements, assignments, function calls, and `{ }` blocks. Assignments, function calls, `return`, `break`, and `continue` end with a semicolon `;`.

### if statements

An `if` statement runs its body when a boolean expression is true, and its optional `else` branch when it is false. The expression is always in parentheses `( )`. The body is a single statement or a block in braces `{ }`:

```
if ( foo == "bar" ) {
	output = "foo is bar!";
} else {
	output = "foo is not bar!";
}
```

Chain conditions with `else if`:

```
if ( foo == "bar" ) {
	output = "foo is bar!";
} else if ( foo == "baz" ) {
    output = "foo is baz!";
} else {
	output = "foo is not bar!";
}
```

### switch statements

A `switch` statement compares an expression in `( )` against a list of `case` expressions, each followed by a `:` and any number of statements. An optional `default` case comes last. For example:

```
switch (int(foo)) {
    case 1:
        output = "foo is one";
        break;
    case 2:
        output = "foo is two";
        break;
    default:
        output = "foo is neither one nor two";
}
```

Execution jumps to the first case, from top to bottom, whose value equals the switch value. If no case matches, the `default` case runs; without one, execution continues after the switch.

As in C, cases fall through: execution continues from the matching case until a `break` or the end of the switch, including into `default`. For example:

```
switch (int(foo)) {
    case 1:
        output = "started at the first case";
    case 2:
        output = output + " and continued into the second case";
}
```

If `foo` equals 1, `output` becomes "started at the first case and continued into the second case", because no `break` stops execution after the first case.

### for statements

A `for` statement repeats a block. It uses the C-style initializer, condition, and post statement, in parentheses `( )` and separated by semicolons `;`. The body must be a block in braces `{ }`. For example, to run a block 10 times:

```
for ( i = 0; i < 10; i++ ) {
	count = i;
}
```

A loop runs at most 1,000,000 iterations. After that, execution continues with the statement after the loop.

### break statements

A `break` statement ends the innermost `for` loop or `switch` statement that contains it. Execution continues with the statement after that loop or switch. For example:

```
/* for loop up to 10, but break after 5 iterations */
for ( i = 0; i < 10; i++ ) {
	if ( i == 5 ) {
		break;
	}
}
```

### continue statements

A `continue` statement starts the next iteration of the innermost `for` loop that contains it. It can only be used inside a `for` loop. For example:

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

A `return` statement ends the program and decides whether the entry continues down the pipeline. It takes a single expression: `true` passes the entry, and `false` drops it. For example:

```
if ( foo == "bar" && bar == "baz" ) {
	return true;
}
return false;
```

## Built-in functions

### Convenience

#### append

    function append(x array, y <expression>) array

Returns a copy of array x with y added as a new last element. Assign the result to keep it, for example `x = append(x, y);`.

#### delete

	function delete(key string)

Deletes the given enumerated value.

#### has

	function has(ev string) bool

Returns true if the given EV exists.

#### has_key

	function has_key(m map, key string) bool

Returns true if the key exists in the persistent map m. Pass the map's name as an identifier, for example `has_key(appnames, Appname)`.

#### in

	function in(x <expression>, <expression>...) bool

Returns true if x equals any of the following arguments. Values are promoted to a common type before comparison, as with `==`.

Example

```
vegetable = "potato";
hasPotato = in(vegetable, "turnip", "potato", "cabbage");
```

#### unix

	function unix(t time) int

Returns the Unix time, in seconds, of the given timestamp. Strings are parsed as timestamps, and integers are treated as Unix seconds.

#### len

	function len(<expression>) int

Returns the number of elements in an array, the number of keys in a persistent map, or the length in bytes of a string.

#### log

	function log(<expression>)

Log the given expression to the "gravwell" tag.

#### match

	function match(input string, pattern string) bool

Returns true if the regular expression in pattern matches the given input. Regular expressions use the [RE2 syntax](https://github.com/google/re2/wiki/Syntax).

#### printf

	function printf(format string, <expression>...) string

Formats a string according to the given format and any number of expressions. printf uses the formatting rules and verbs defined in [Golang's fmt.Printf()](https://pkg.go.dev/fmt#hdr-Printing) function.

(set-data-target)=
#### set_data

	function set_data(data <expression>)

Overwrites the entry's DATA field. This is different from creating an enumerated value named DATA.

#### evs

    function evs() array

Returns an array of all enumerated value names attached to the entry.

#### indirect

    function indirect(<expression>)

Returns the value of the EV whose name is the given expression. For example, `indirect("f"+"oo")` returns the value of the EV named `foo`. If no such EV exists, it returns the string "foo".

#### keys

    function keys(m map) array

Returns an array of all keys in the persistent map m, in no particular order.

#### ip_in_cidr

    function ip_in_cidr(ip ip, cidr string) bool

Returns true if the given IP is in the given CIDR-formatted string. For example, `ip_in_cidr(src_ip, "10.2.254.0/24")`.

### String manipulation

#### strings_contains

	function strings_contains(input string, pattern string) bool

Returns true if pattern appears anywhere in input.

#### strings_count

	function strings_count(input string, pattern string) int

Returns the number of non-overlapping instances of pattern in input. If pattern is empty, it returns 1 plus the number of Unicode code points in input.

#### strings_hasprefix

	function strings_hasprefix(input string, prefix string) bool

Returns true if input begins with prefix.

#### strings_hassuffix

	function strings_hassuffix(input string, suffix string) bool

Returns true if input ends with suffix.

#### strings_index

	function strings_index(input string, pattern string) int

Returns the byte index of the first instance of pattern in input, or -1 if pattern is not in input.

#### strings_replace

	function strings_replace(input string, old string, new string, n int) string

Returns input with the first n non-overlapping instances of old replaced by new. If n is -1, all instances are replaced.

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

    function pretty_size(n int) string

Formats a byte count as an abbreviated size. For example, 1234567 becomes "1.18 MB".

#### pretty_count

    function pretty_count(n int) string

Formats a number as an abbreviated magnitude. For example, 1234567 becomes "1.23 M".

#### pretty_rate

    function pretty_rate(n int, d duration) string

Formats n bytes over duration d as a rate in bytes, kilobytes, or megabytes per second. For example, `pretty_rate(1234567, "2s")` becomes "602.82 KB/s".

#### pretty_line_rate

    function pretty_line_rate(n int, d duration) string

Formats n bytes over duration d as a line rate in bits, kilobits, or megabits per second. For example, `pretty_line_rate(1234567, "2s")` becomes "4.71 Mb/s".

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

Returns the last element of the given path. Trailing slashes are removed before extracting the last element.

#### path_clean

	function path_clean(path string) string

Returns the shortest path name equivalent to the given path. For example, "/../a/b/../././/c" becomes "/a/c".

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

Returns true if name matches the given pattern. The pattern syntax is identical to [Golang's path.Match()](https://pkg.go.dev/path#Match) function.


### Random

#### rand_float

	function rand_float() float

Returns a random 64-bit float in the interval [0.0,1.0)

#### rand_int

	function rand_int() int

Returns a random, non-negative, 64-bit integer.

#### rand_intn

	function rand_intn(n int) int

Returns a random, non-negative, 64-bit integer in the range [0,n). n must be greater than 0.

#### rand_normal

	function rand_normal() float

Returns a 64-bit float from the standard normal distribution (mean = 0, stddev = 1). To produce a different normal distribution, adjust the output:

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

Appends the value to the given JSON array and returns the updated array. If array isn't valid JSON, it returns a new JSON array containing only the value.

#### json_array

	function json_array(value <expression>) string

Returns a JSON array of the given value. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.

#### json_array_set

	function json_array_set(array string, index <expression>, value <expression>) string

Sets the value of the given JSON array at the given index and returns the updated array. If the index is out of range, the array is returned unchanged. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.

#### json_get

	function json_get(object string, key string) bool/float/string

Returns the value of key in the given JSON object. JSON numbers are returned as floats, and nested objects and arrays as JSON strings. If the key doesn't exist, it returns the string "null".

#### json_index

	function json_index(array string, index int) bool/float/string

Returns the item at the given index of a JSON array. JSON numbers are returned as floats, and nested objects and arrays as JSON strings.

#### json_keys

    function json_keys(object string) array

Returns an array of the keys of the given JSON object, in no particular order.

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

Sets a key/value pair in the given JSON object and returns the updated object. If object isn't valid JSON, it returns a new object containing only this key. The value's type is evaluated at runtime and will map to the corresponding JSON type (object, array, bool, number, string), or a string if the type doesn't map to a JSON type.

#### json_to_gravwell_array

    function json_to_gravwell_array(array string) array

Converts a JSON array to a Gravwell array. Numbers become floats, strings and booleans keep their types, and nested objects and arrays become JSON strings.

(eval-evs-to-json)=
#### evs_to_json

    function evs_to_json() string

Returns a JSON object containing all enumerated values attached to the current entry. Each EV name becomes a key, and each value is serialized to its JSON type (number, string, boolean, array, or object). Durations, MAC addresses, and timestamps (in UTC) are serialized as strings.

Example

```
tag=gravwell syslog Appname Hostname Message
| eval output = evs_to_json();
| table output
```

This would produce output like:

```
{"Appname":"webserver","Hostname":"potato","Message":"connection established"}
```


(eval-math)=
### Math

```{note}
`ceil`, `floor`, and `round` are legacy names for `math_ceil`, `math_floor`, and `math_round`.
```

#### ceil

	function ceil(x float) float

Returns the least integer value greater than or equal to x.

#### floor

	function floor(x float) float

Returns the greatest integer value less than or equal to x.

#### math_abs

	function math_abs(x float) float

Returns the absolute value of x.

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

Returns x raised to the power y.

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

Returns the integer value of x by dropping decimal digits. For example, `3.1` and `3.9` will both return `3`. 

#### round

	function round(x float) float

Returns the nearest integer, rounding half away from zero.

#### math_acos

    function math_acos(x float) float

Returns the arccosine, in radians, of x.

#### math_acosh

    function math_acosh(x float) float

Returns the inverse hyperbolic cosine of x.

#### math_asin

    function math_asin(x float) float

Returns the arcsine, in radians, of x.

#### math_asinh

    function math_asinh(x float) float

Returns the inverse hyperbolic sine of x.

#### math_atan

    function math_atan(x float) float

Returns the arctangent, in radians, of x.

#### math_atanh

    function math_atanh(x float) float

Returns the inverse hyperbolic tangent of x.

#### math_atan2

    function math_atan2(y, x float) float

Returns the arctangent of y/x, using the signs of the two to determine the quadrant of the return value.

#### math_cos

    function math_cos(x float) float

Returns the cosine of the radian argument x.

#### math_cosh

    function math_cosh(x float) float

Returns the hyperbolic cosine of x.

#### math_sin

    function math_sin(x float) float

Returns the sine of the radian argument x.

#### math_sinh

    function math_sinh(x float) float

Returns the hyperbolic sine of x.

#### math_tan

    function math_tan(x float) float

Returns the tangent of the radian argument x.

#### math_tanh

    function math_tanh(x float) float

Returns the hyperbolic tangent of x.

## Acceleration and eval

Eval is capable of producing acceleration hints for the fulltext acceleration engine when used as a filter. For example:

```
tag=gravwell syslog Appname | eval ( Appname == "webserver" || Appname == "indexer" )
```

This query will provide the acceleration hints of "webserver" OR "indexer" to the fulltext engine. By using eval, complex acceleration hints can be composed that aren't possible using extraction modules alone.

## Syntax

The eval syntax is expressed using a [variant](https://github.com/gravwell/pbpg) of Extended Backus-Naur Form (EBNF):

```
Program                  = ( "(" Expression ")" EOF ) | ( "(" [ Vars ] StatementList ")" EOF ) | ( "(" Assignment ")" EOF ) | ( Expression EOF ) | ( [ Vars ] StatementList EOF ) | ( Assignment EOF )
Vars                     = VarSpec { VarSpec }
VarSpec                  = ( "var" VarSpecAssignment { "," VarSpecAssignment } ";" ) | ( "map" AssignmentIdentifier ";" )
VarSpecAssignment        = AssignmentIdentifier [ "=" Expression ]
StatementList            = Statement { Statement }
Statement                = ( "if" "(" Expression ")" Statement "else" Statement ) | ( "if" "(" Expression ")" Statement ) | ( "for" "(" Assignment ";" Expression ";" Assignment ")" "{" StatementList "}" ) | ( "switch" "(" Expression ")" "{" { ExpressionCaseClause } [ DefaultCaseClause ] "}" ) | "{" StatementList "}" | Function ";" | Assignment ";" | "return" Expression ";" | "break" ";" | "continue" ";" | ";"
ExpressionCaseClause     = "case" Expression ":" [ StatementList ]
DefaultCaseClause        = "default" ":" [ StatementList ]
Assignment               = ( AssignmentIdentifier { "[" Expression "]" } "=" Expression ) | Expression
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
PostfixExpression        = ( PrimaryExpression Index { Index } ) | ( PrimaryExpression [ PostfixOp ] )
Index                    = ( "[" Expression ":" Expression "]" ) | ( "[" Expression "]" )
PrimaryExpression        = NestedExpression | Identifier | Literal
NestedExpression         = ( Function ) | ( Cast "(" Expression ")" ) | ( "(" Expression ")" )
Literal                  = ArrayLiteral | DecimalLiteral | FloatLiteral | StringLiteral | "true" | "false"
ArrayLiteral             = "{" [ Expression { "," Expression } ] "}"
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
Cast                     = "int" | "float" | "string" | "mac" | "ip" | "time" | "duration" | "type" | "bool" | "location" | "byte" | "int8" | "uint8" | "int16" | "uint16" | "int32" | "uint32" | "int64" | "uint64" | "array"
```

## Legacy Eval

There is a legacy version of eval that you may still see in older queries. For more details, see the [Legacy eval page](legacy-eval) for reference.
