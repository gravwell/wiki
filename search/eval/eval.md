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
int
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

`eval` has several built-in functions:

```
round     Round floating-point number to nearest even 
len	  Returns an integer length of the variable
floor     Returns the integer part of a floating-point number
ceil 	  Returns the integer part of a floating-point number after rounding to nearest even
delete    Deletes the enumerated value in the entry
rand      Returns a random 64-bit integer
log	  Logs a message according to the logging configuration in the deployment's gravwell.conf
has       Returns true if the specified EV exists
```

## Acceleration and eval

Eval is capable of producing acceleration hints for the fulltext acceleration engine when used as a filter. For example:

```
tag=gravwell syslog Appname | eval ( Appname == "webserver" || Appname == "indexer" )
```

This query will provide the acceleration hints of "webserver" OR "indexer" to the fulltext engine. By using eval, complex acceleration hints can be composed that aren't possible using extraction modules alone.

## Syntax

The eval syntax is expressed using a [variant](https://github.com/gravwell/pbpg) of Extended Backus-Naur Form (EBNF):

```
Program                  = ( "(" Expression ")" EOF ) | ( "(" Vars StatementList ")" EOF ) | ( "(" StatementList ")" EOF ) | ( Expression EOF ) | ( Vars StatementList EOF ) | ( StatementList EOF )
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
Cast                     = "int" | "float" | "string" | "mac" | "ip" | "time" | "duration" | "type" | "bool"
```
