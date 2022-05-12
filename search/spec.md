# The grav.y Specification

## Introduction

This is the reference specification the Gravwell query language syntax (grav.y). A query is made up of indexer and webserver constraints, modules, a pipeline, and a renderer. This document provides documentation for how input text is interpreted and tokenized. Some lexical meaning of input is also defined here. Modules have context-specific semantics that differ between modules (such as numbers being implied to be strings). The user should read the [search module](#!search/complete-module-list.md) documentation for more information on module-specific considerations.

## Text Encoding 

**All** Gravwell input is Unicode text encoded in UTF-8. 

### Characters and digits 

A "character" is any of the Unicode points in the "General Category" of the Unicode specification, ([The Unicode Standard 8.0, Section 4.5](https://www.unicode.org/versions/Unicode8.0.0/)). This includes letters, numbers, marks, symbols, and punctuation. 

## Lexical grammar

This section defines the syntax of a Gravwell query. Token semantics are module-specific, so the user should read the [search module](#!search/complete-module-list.md) documentation for more information on module-specific considerations.

Note: The grammar is specified in Extended Backus-Naur Form (EBNF).

### Quotes

Many characters and keywords have special meaning in the Gravwell syntax. When using a special character or keyword as a literal string, you must use double quoted strings `"`. For example

```
json foo bar | table
```

extracts `foo` and `bar` from some JSON input, then passes the extractions to the `table` renderer. To use whitespace and the `|` character, both of which have special meaning, in the input to the `json` module, use a quoted string:

```
{
	"TS":"2022-05-06T09:48:41.371259747-06:00",
	"foo bar | table":"data!"
}
```

```
json "foo bar | table"
```

This example will extract a field named `foo bar | table` from the above JSON. 

All other lexical definitions below are implied to be interpreted outside of a quoted string, except for escaped productions, which are always interpreted.

A quoted string is defined as 

```
quoted_string 	= '"' { unicode_print | whitespace } '"' 
whitespace 	= Characters from Unicode's whitespace category and Unicode category Z
unicode_print 	= Characters from Unicode categories L, M, N, P, and S
```

### Escaped input

Escaped text follows the same rules as described in the [Rune Literals](https://go.dev/ref/spec#Rune_literals) section of the Go programming language specification. See that document for more information.
 
### Whitespace

Whitespace is defined as all Unicode whitespace characters, which includes whitespace characters outside of the `Latin-1` category. Whitespace characters are folded, meaning multiple, continuous whitespace characters are interpreted as a single whitespace character. Whitespace delimits tokens when used outside of a quoted string.

For example, the input

```
json foo bar
```

is made up of three tokens `json`, `foo`, and `bar`. In the `json` module in Gravwell, this would extract two enumerated values `foo` and `bar`.

Quoted whitespace is treated as part of a single string. For example, the input

```
json "foo bar"
```

is made up of two tokens `json`, and `foo bar`. The `json` module in this case would extract a single enumerated value named `foo bar`.

Whitespace is defined as

```
whitespace = Characters from Unicode's whitespace category and Unicode category Z
```

### Comments

A comment is any input between ANSI-C style comment specifiers `/* */`, and is not considered part of the input to Gravwell. 

For example,

```
kv myKey /* this is my key extraction */ | table
```

is implicitly reduced to

```
kv myKey | table
```

A comments is defined as

```
comment 	= '/' '*' { unicode_print | whitespace } '*' '/' 
whitespace 	= Characters from Unicode's whitespace category and Unicode category Z
unicode_print 	= Characters from Unicode categories L, M, N, P, and S
```

### Module tokens

Tokens make up the "vocabulary" of the Gravwell query language. Tokens are groups of characters separated by whitespace (as defined above) and reserved characters (such as `|`), unless grouped in a quoted string. The semantic meaning of a token depends on the position the token occurs in the _token stream_. For example,

```
tag=default json tag
```

extracts the enumerated value `tag`, using the `json` module, all from the default Gravwell tag. While the token `tag` shows up twice, the meaning is different based on the position in the token stream. The first occurrence tells Gravwell to pull data from the default tag. The second occurrence tells the `json` module to extract a value named `tag`.

Tokens cannot contain the following reserved characters, unless quoted:

| Character | Description |
|-----------|-------------|
| &#124; |  Pipe: separates modules in the module pipeline |
| @ | Compound query reference |
| {} | Compound query block |
| ; | Compound query delimiter |
| = | Assignment operator |
| ==, <, >, <=, >=, ~, !=, !~ | Comparison operators |
| ., !, #, $, %, ^, &, *, (, ), [, ], ? | Other reserved characters |

### Tokenizing in the R-value of a filter

When filtering, tokenizing in the R-value (the value of the filter) of the filter behaves differently. All reserved characters except `|[](){}` will be considered part of the token until the next whitespace character. This means that while `uint16(Data[2:5])` is split into 9 tokens (all reserved characters cause token splitting), the filter value in `foo == ::!!!.50` is a single token.

The grammar of a module token is

```
token 		= string | r_string (after an operator) | special 
special 	= "@" | "|" | "{" | "}" | ";" | "=" | "==" | "<" | ">" | "<=" |
		  ">=" | "~" | "!=" | "!~" | "!" | "%" | "^" | "&" | "*" | "(" |
		  ")" | "," | "." | "#" | "$" | "?" 
string 		= normal_print { normal_print } | quoted_string 
r_string	= quoted_string | r_print { r_print } 
normal_print	= unicode_print // except special and " "
r_print 	= unicode_print // except |{}()[]
quoted_string 	= '"' { unicode_print | whitespace } '"' 
whitespace 	= Characters from Unicode's whitespace category and Unicode category Z
unicode_print 	= Characters from Unicode categories L, M, N, P, and S
```


### Tokenizing in `eval` and code fragments

Gravwell syntax supports inline code fragments for filtering and other operations. This is accomplished with either the `eval` module, followed by the code fragment, or a module stage wrapped in parentheses. For example,

```
tag=default json foo-bar baz | eval baz > 10 | table
```

has the code fragment `baz > 10`. This is easily parsed using the tokenizing rules described above. This same query can be written as

```
tag=default json foo-bar baz | (baz>10) | table
```

However, the code fragment syntax supports C-style notation for bitwise and logic operations, so Gravwell parses these fragments differently than the regular token stream. For example,

```
tag=default json foo-bar foo bar | ( foo-bar > 10 ) | table
```

has a code fragment `foo-bar > 10`, but it is unclear if the user meant "foo minus bar is greater than 10" or "the enumerated value 'foo-bar' is greater than 10". This is because hyphens are allowed in the tokenizing in the `json` module preceding the code fragment. 

Another example,

```
tag=default json foo bar | ( baz = foo | bar ) | table
```

has an interior `|` character, which would otherwise cause a module split, but the intended use here is to perform a bitwise-or of the two enumerated values "foo" and "bar".

To reconcile this behavior, `eval` and implied code fragments tokenize in a different way:

- Enumerated value names (identifiers in the grammar below) are limited to unquoted strings beginning with a non-number, and not containing any of the special characters in the grammar below.
- String literals must be quoted.
- Numeric literals are all forms of numbers, floating point numbers, hexadecimal syntax (eg 0xfa), and binary (eg 0b0010).
- `|` and `||` are treated as bitwise and logical OR operations, respectively.

NOTE: Enumerated values containing reserved characters or whitespace cannot be used in code fragments. These variables must be renamed or aliased.

This form of tokenizing occurs until the outermost parenthetical group in the code fragment is closed.

The grammar of a token in a code fragment is:

```
token 			= identifier | string | int_lit | float_lit | special .
special 		= "|" | "{" | "}" | ";" | "=" | "==" | "<" | ">" | "<=" | ">=" |
		  	  "~" | "!=" | "!~" | "!" | "%" | "^" | "&" | "*" | "(" | ")" |
		  	  "&&" | "||" | "/" | "<<" | ">>" | "+=" | "-=" | "++" | "--" |
		  	  "," | "+" | "-" 
string			= quoted_string 
identifier 		= code_print_not_number { code_print } 
code_print_not_number 	= code_print // except unicode_numbers
code_print		= unicode_print // except special and " "
int_lit        		= decimal_digits | binary_lit | hex_lit 
decimal_digits 		= decimal_digit { decimal_digit } 
binary_lit     		= "0" ( "b" | "B" ) binary_digits 
hex_lit        		= "0" ( "x" | "X" ) hex_digits 
binary_digits  		= binary_digit { binary_digit } 
hex_digits     		= hex_digit { hex_digit } 
binary_digit		= "0" | "1" 
hex_digit		= "0…9" | "a…f" | "A…F" 
decimal_digit		= "0…9" 
float_lit 		= decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    	  decimal_digits decimal_exponent |
                    	  "." decimal_digits [ decimal_exponent ] 
exponent	  	= ( "e" | "E" ) [ "+" | "-" ] decimal_digits 
quoted_string 		= '"' { unicode_print | whitespace } '"' 
whitespace 		= Characters from Unicode's whitespace category and Unicode category Z
unicode_print 		= Characters from Unicode categories L, M, N, P, and S
```

### Operators and filters

Operators are reserved characters that are used when applying filters in certain modules. Filters, and their operators, are always in the form of

```
<identifier> <operator> <value>
```

For example,

```
foo <= 1.5
```

applies a filter on the identifier `foo`, requiring that it is less than or equal to the floating point value 1.5. 

Individual modules specify the rules for what types and operators are allowed on given identifiers. See the module documentation for more information.

## Type inference

Modules individually specify how values are typed. For example, the `json` module extracts all values as strings, and using a filter `json foo == 1.5` will perform a string comparison on the value of foo with the _string_ `1.5`. 

## Query structure

![](query.png)

A query is made up of any number of query constraints (such as which tags to pull data from), search modules (arranged in a pipeline), and a render module. Additionally, multiple queries can be packed into a _compound query_.

### Query constraints 

All input before the first module in a query represents the query constraints. Usually this is just the tag(s) to extract. Query constraints are a whitespace-delimited list of key=value pairs. The following constraints are supported.

| Constraint | Description | Example |
|------------|-------------|---------|
| tag | The tag(s) to extract. Supports comma separated lists and wildcards. Defaults to "tag=default" if omitted. | tag=dns,zeek* |

A query constraint uses the following grammar:

```
constraint 		= token "=" token 
token 			= constraint_graphic { constraint_graphic } 
constraint_print 	= unicode_print // except ", =, and " "
unicode_print		= Characters from Unicode categories L, M, N, P, and S
```

### Modules

Please see the [list of modules](#!search/complete-module-list.md) for module specific documentation.

Modules are pipelined functions that extract, transform, and render data. Conceptually, data flows left-to-right in the module pipeline, and modules can drop, pass, modify, or inject data into the pipeline. The last module in the pipeline is the render module (such as `table` or `chart`); note that if no render module is explicitly defined, Gravwell will add one automatically. The module pipeline is split by the `|` character. A module invocation is made up of the module name, optional flags, and optional arguments. 

#### Module name

The first token of a module invocation is the module name. See the [list of modules](#!search/complete-module-list.md) for the list of available modules.

#### Module flags

Flags are whitespace-delimited lists of module-specific configuration options. Some flags take arguments, and some modules support multiple invocations of the same flag. See the module documentation for more information on the specific flags for that module.

Flags take two forms. First is a flag with no arguments, such as

```
kv -q
```

The second form takes an argument, which must be quoted if it contains any reserved characters or whitespace.

```
kv -e "my enumerated value"
```

The first token that does not begin with a hyphen, and is not the argument to a flag, represents the first argument to the module.

#### Module arguments and semantics

Module arguments are module-specific lists of tokens containing identifiers, filters, and other keywords. Every token up to the end of the input or the `|` character is considered an argument to that module. Modules apply their own semantics to their argument tokens. For example, the `fields` module uses the special characters `[` and `]` to denote a column offset, while the `lookup` module uses `(` and `)` for groups of extractions. Refer to the module documentation for more information.

### Renderer

The last module in a query is always the renderer. Render modules take flags and arguments just like a search module, but must appear at the end of the pipeline.

If no renderer is specified, the text renderer is implied.

### Compound queries

Multiple modules can be grouped into a single _compound query_ using the [compound query notation](#!search/complete-module-list.md). A compound query takes the form of

```
@foo{tag=default ...}; @bar{tag=default ...}; tag=default lookup -r @foo | ...
```

Where `@foo`, `@bar` represent the names of "inner" queries. The `@` is required. Any query can be specified in the inner query body, enclosed in `{}`, but the renderer must be the `table` renderer. Any number of inner queries can be specified. Inner queries and the main query (the final query in the list of queries, which is not wrapped in the `@{}` notation) are separated by semicolons. 

Queries are executed in order, and any later query (including other inner queries) can use the output of an earlier inner query anywhere that a tabular resource can be used (such as the `lookup` module), by referencing the query by name with the `@` symbol. 
