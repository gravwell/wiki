## Eval

### Introduction

The `eval` module is a general purpose filtering and programming environment that runs a "program" against each entry in the query pipeline. It uses a C-style syntax and supports creating simple filters, as well as arbitrary operations on each entry. 

For example, to filter entries to just those where the enumerated value `foo` is equal to `bar` or `baz`:

```
tag=gravwell json foo | eval ( foo == "bar" || foo == "baz" )
```

Additionally, C-style statements can be used for more advanced operations. For example, to conditionally create an EV:

```
tag=gravwell json foo | eval if (foo == "bar") { other_variable = "hey! foo is bar!" }
```

The rest of this document describes the syntax and semantics of the eval language. 

### Lexical elements

#### Identifiers

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

#### Keywords

The following keywords are reserved and may not be used as identifiers.

```
if
else
for

(the remaining keywords are not currently supported, but are reserved)

break		
case
continue
default
return
switch
```

#### Operators

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

#### Punctuation

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

#### Numbers

Numbers can be expressed as decimal and hexadecimal integers, as well as decimal floating-point. Hexadecimal values always begin with "0x". 

The following examples are valid numbers.

```
5
-5
0x05
5.00
```

#### String literals

Strings are any sequence of characters enclosed in double quotes `"`. 

### Numeric literals

Numeric literals consist of any decimal values, hex values (for example, 0xffff), and floating point values.

### Variables and types

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
```

Additionally, eval will attempt to "promote" types that can be automatically cast. For example, if "foo" contains the string "56", the expression `foo < 3.14` will cause eval to attempt to promote foo to a floating-point number before the expression is evaluated.

### Scope

Eval has no variable scoping. All variables are "global", regardless of if they are declared in a statement block or not.

### Expressions

Expressions are simply computations of values. The behavior of an expression depends on where in a program it is used. 

Expressions follow the C-style form of expressions. For example, `if (foo == "bar")...` contains the boolean expression `foo == "bar"` which is used to compute a true/false value for determining if the body of the if statement should be executed or not. 

The right-hand-side of an assignment is also an expression. For example, `foo = type(bar);` contains the expression `type(bar)`, which is a built-in function that returns a value. 

Expressions can be complex, and contain other expressions, just like in C-style programs. For example, `( a == b || b < 3.14 || ( type(b) == "string" && error == "none" ))` is an expression made up of several, smaller expressions.

```{note}
When a program consists of only an expression, the program as a whole is treated as a filter. If the expression returns false (or returns the "zero" value of the type returned), the entry will be dropped.
```

### Statements

Statements control program execution. For example, `if (foo == "bar") { ... }` contains an "if" statement, which will determine how the program is to proceed.

Statements are separated by the `{ }` punctuation in the case of "if" and "for", and semicolons ";" for assignments. Currently, only "if", "for", and assignment statements are supported in eval.

### Built-in functions

`eval` has several built-in functions:

```
round     Round floating-point number to nearest even 
len	  Returns an integer length of the variable
floor     Returns the integer part of a floating-point number
ceil 	  Returns the integer part of a floating-point number after rounding to nearest even
delete    Deletes the enumerated value in the entry
rand      Returns a random 64-bit integer
log	  Logs a message according to the logging configuration in the deployment's gravwell.conf
```

### Syntax

The eval syntax is expressed using a [variant](https://github.com/gravwell/pbpg) of Extended Backus-Naur Form (EBNF):

```
Program                  = ( "(" Expression ")" EOF ) | ( "(" StatementList ")" EOF ) | ( Expression EOF ) | ( StatementList EOF )
StatementList            = Statement { Statement }
Statement                = ( "if" "(" Expression ")" Statement [ "else" Statement ] ) | ( "for" "(" Assignment ";" Expression ";" Assignment ")" "{" StatementList "}" ) | "{" StatementList "}" | Function "(" Expression { "," Expression } ")" | Assignment | ";"
Assignment               = ( Literal "=" Expression ) | Expression
Expression               = ( LogicalOrExpression "?" Expression ":" LogicalOrExpression ) | LogicalOrExpression
LogicalOrExpression      = LogicalAndExpression { "||" LogicalAndExpression }
LogicalAndExpression     = InclusiveOrExpression { "&&" InclusiveOrExpression }
InclusiveOrExpression    = ExclusiveOrExpression { "|" ExclusiveOrExpression }
ExclusiveOrExpression    = AndExpression { "^" AndExpression }
AndExpression            = EqualityExpression { "&" EqualityExpression }
EqualityExpression       = RelationalExpression { EqualityOp RelationalExpression }
RelationalExpression     = ShiftExpression { RelationalOp ShiftExpression }
ShiftExpression          = AdditiveExpression { ShiftOp AdditiveExpression }
AdditiveExpression       = MultiplicativeExpression { AdditiveOp MultiplicativeExpression }
MultiplicativeExpression = UnaryExpression { MultiplicativeOp UnaryExpression }
UnaryExpression          = [ "-" | "!" ] PostfixExpression
PostfixExpression        = PrimaryExpression [ "++" | "--" ]
PrimaryExpression        = ( Function "(" Expression ")" ) | ( [ Cast ] "(" Expression ")" ) | Literal
EqualityOp               = "==" | "!=" | "~" | "!~"
RelationalOp             = "<" | ">" | "<=" | ">="
ShiftOp                  = "<<" | ">>"
AdditiveOp               = "+" | "-"
MultiplicativeOp         = "*" | "/" | "%"
Cast                     = "int" | "float" | "string" | "mac" | "ip" | "time" | "duration" | "type"
Function                 = "round" | "len" | "floor" | "ceil" | "delete" | "rand" | "log"
```
