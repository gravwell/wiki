# Scripting in Gravwell

Scripting is used in two ways within Gravwell: as part of a search pipeline, and as a method to automate search launching. The scripting language ([Anko](https://github.com/mattn/anko)) is the same in both cases, with some slight differences to account for the differing use cases. This article introduces both use cases and provides a high-level overview of the Anko language.

* [`anko` module documentation](anko.md)
* [`eval` module documentation](eval.md)
* [Automation scripting documentation](scriptingsearch.md) (contains detailed descriptions of functions available for automation scripts)
* [Scheduling scripts & queries](scheduledsearch.md)

## Scripting modules

Gravwell includes two modules, `anko` and `eval`, which use the [Anko scripting language](https://github.com/mattn/anko) to provide Turing-complete scriptability in the search pipeline.  The anko module enables the full feature set of anko to provide a full [Turing-Complete](https://en.wikipedia.org/wiki/Turing_completeness) language and runtime.  Eval uses the anko runtime to execute a single statement entered directly in the search query.

While anko can do anything, eval has several important restrictions:

* Only a single statement can be defined: `(x==y || j < 2)` is acceptable, but `setEnum(foo, "1"); setEnum(bar, "2")` is two statements and will not work.
* Functions cannot be defined or imported
* Loops are not allowed
* No access to the resource system

This document describes the Anko programming language itself. Documentation for the two search modules is maintained on separate pages:

* [`anko` documentation](anko.md) (anko is also briefly described in [the search modules documentation](#!search/searchmodules.md#Anko))
* [`eval` documentation](eval.md) (eval is also briefly described in [the search modules documentation](#!search/searchmodules.md#Eval))

## Search Scripts

Where the `anko` and `eval` modules run scripts *inside* search pipelines, Gravwell also supports scripts which *launch* searches of their own and operate on the results. This is useful for automated queries, e.g. a script which runs every morning at 6 a.m. to look for particular suspicious network behavior.

These scripts can be either run on a schedule (see [automation scripts](#!scripting/scheduledsearch.md)) or run by hand using the [command line client](#!cli/cli.md). The scripting language is the same in both cases, although scripts run on a schedule cannot use `print` functions to display output.

The [automation script](scriptingsearch.md) documentation provides more information on how to write this type of script, including examples.

## Anko overview

Briefly, Anko is a dynamically-typed scripting language which largely resembles Go in syntax but is interpreted rather than compiled. The Anko github page provides this example to give an idea of how Anko looks and operates:

```
# declare function
func plus(n){
  return n + 1
}

# declare variables
x = 1
y = x + 1

# print values
println(x * (y + 2 * x + plus(x) / 2))

# if/else condition
if plus(y) > 1 {
  println("こんにちわ世界")
} else {
  println("Hello, World")
}

# array type
a = [1,2,3]
println(a[2])
println(len(a))

# map type
m = {"foo": "bar", "far": "boo"}
m.foo = "baz"
for k in keys(m) {
  println(m[k])
}
```

The [Anko playground](http://play-anko.appspot.com/) is a convenient way to experiment with Anko code; we recommend playing with the example above and other examples in the documentation in order to get a feel for Anko. (Examples using Gravwell-specific functions, such as setEnum, will of course not work in the playground.)

## Data types

Unlike Go, Anko rarely uses explicit type declarations. Type is usually inferred and automatically converted whenever possible.

Anko supports the following basic types:

* integer: int, int32, int64, uint, uint32, uint64
* floating point: float32, float64
* boolean: bool
* string: string
* character: byte, rune
* interface

The `interface` type is special; it represents a generic object, so an array of `interface` types can hold strings, integers, floating points, etc.

Anko provides two kinds of composite types:

* arrays: multi-dimensional arrays of data
* maps: 

To declare a scalar, it is not necessary (or possible) to specify the variable type; simply assign the variable and the type will be inferred:

```
a = 1		# integer
b = 2.5		# float
c = true	# bool
d = "hi"	# string
```

Implicit conversion is done whenever possible, selecting the appropriate type to preserve accuracy:

```
x = a + b	# x == 1 + 2.5 == "3.5"

y = b - c	# boolean true converts to 1, so y == 2.5 - true == 2.5 - 1.0 == 1.5
```

The exact semantics of various operations on the different types are explained in later sections.

Attention: If you use hex constants in Anko, be sure to capitalize A-F. `0x1E` is a valid representation of the number 14, but `0x1e` is not.

### Arrays

Arrays in Anko may be generic (holding elements of varying types) or typed. To create a generic array, simply assign an initial value:

```
myArray = ["hi", 1]
printf("%v\n", myArray)    # prints ["hi" 1]
myArray[0] = 3.5
printf("%v\n", myArray)    # prints [3.5 1]
```

To create a typed array, use the `make` function. `make` takes the array type and an initial length as arguments. The array type can be any of the scalar types listed in the preceding section.

```
myArray = make([]int, 5)	# make an array of 5 ints
myArray[1] = 7
printf("%v\n", myArray)		# prints [0 7 0 0 0]
```

Note: generic arrays can also be constructed using the `make` function: `make([]interface, 10)`

Multi-dimensional arrays are possible; the following example shows two different ways to achieve the same multi-dimensional array:

```
a = [[1, "foo"][3.2, 4]]

b = make([][]interface, 2)
b[0] = [1, "foo"]
b[1] = make([]interface, 2)
b[1][0] = 3.2
b[1][1] = 4
# at this point, a and b are equivalent
```

#### Appending to arrays

An array can be appended to using the `+=` operator:

```
a = [1, 2]
a += 3	# a is now [1 2 3]
```

One array can be appended to another:

```
a = [1, 2]
b = [3, 4]
a += b	# a is now [1 2 3 4]
```

Anko also allows implicit appending. If the specified index is exactly one greater than index of the current end of the array, the array is expanded and the item is appended:

```
foo = ["a", "b"]
foo[2] = "c"	# foo is now ["a" "b" "c"]

# This will fail!
foo[5] = "bar"
```

#### Slicing arrays

As in Go, it is possible to extract a portion of an array by specifying bounds. Given an array `a`, the expression `a[low:high]` extracts a sub-slice consisting of the elements of `a` whose indexes satisfy the expression `low <= index < high`.

Thus, given `a = [1, 2, 3, 4]`, the expression `a[1:3]` evaluates to `[2 3]`.

Omitting lower or upper bound sets it to the start or end of the array, respectively. So given `a = [1, 2, 3, 4]`, the expression `a[:2]` evaluates to `[1 2]`, while `a[1:]` evaluates to `[2 3 4]`.

Using the `len` function, it is possible to perform more complex slicing:

```
a = [1, 2, 3, 4, 5, 6]
b = a[:len(a)-3]			# b == [1 2 3]
c = a[len(a)-4:len(a)-2]	# c == [3 4]
```

Note: It is also possible to slice strings in the same fashion. Thus given `a = "hello"`, `a[1:4]` evaluates to `"ell"`.

### Maps

Anko provides a limited form of Go's maps (Python programmers will recognize them as 'dictionaries'). In Anko's maps, keys are strings and values are any type. For example:


```
a = {}					# define an empty map
a["foo"] = 2			# key = "foo", value is int
a["bar"] = [1, 2, 3]	# key = foo, value is array
```

Maps can be pre-populated:

```
a = {"foo": 2, "bar": [1, 2, 3]}
```

To remove an element from a map, use the `delete` function:

```
a = {"foo": 2, "bar": [1, 2, 3]}
delete(a, "bar")
```

### Channels

Anko provides an interface to use Go channels. Channels are first-in, first-out pipelines of data which are useful for concurrency. Channels can be created of any type, but due to the implicit typing of Anko you should use caution with channels of types other than `interface`, `int64`, `float64`, `string`, and `bool`. Generally, a channel of `interface` should be sufficient for most tasks.

Channels have a "size", which is the number of elements the channel can hold before writes to the channel block. A channel of size 0 is unbuffered, meaning a write to the channel will block until a read is performed, and vice versa. A channel of size 1 can have one item written to it before writes block. See the [discussion of channels in 'Effective Go'](https://golang.org/doc/effective_go.html#channels) for more information on channels.

Channels are created using the `make` function, which takes the channel type and an optional channel size as the argument:

```
unbuf = make(chan interface)	# an unbuffered channel
buf = make(chan bool, 1)		# a buffered channel
```

Once a channel has been created, it may be written to and read from as in Go:

```
c = make(chan interface, 2)
c <- "foo"
c <- "bar"
a = <- c	# the variable 'a' will contain the string "foo" read from the channel
b = <- c	# variable 'b' contains "bar"
```

## Writing Anko code

### Variable creation, assignment and scoping

Anko uses dynamic scoping, which may be unfamiliar to programmers used to lexical scoping as implemented in C and many other languages. Variables are assigned using the `=` operator. If the variable does not exist, it will be created in the current scope. If the named variable already exists in the current scope or any scope above it, the value of the variable will be set to the new variable.

The following example demonstrates how the second assignment (`a = 2`) modifies the outer scope's 'a' variable rather than creating a new one:

```
func foo() {
	a = 2
}
a = 1
println(a)		# prints "1"
foo()
println(a)		# prints "2"
```

In order to explicitly create a new variable in an inner scope, use the `var` keyword:

```
func foo() {
	var a = 2
}
a = 1
println(a)		# prints "1"
foo()
println(a)		# prints "1"
```

The above code creates another variable named 'a' in the inner scope of the function definition.

### Math and logic operations

Anko provides a standard set of basic mathematical and logical operations as found in languages such as Go and C. Take care to remember that Anko will implicitly convert values when needed.

Anko supports the following mathematical operations:

| Expression syntax | Description
|-------------------|--------------
| lhs + rhs | Returns the sum of lhs and rhs 
| lhs - rhs | Returns the difference of lhs and rhs 
| lhs * rhs | Returns lhs multiplied by rhs 
| lhs / rhs | Returns lhs divided by rhs 
| lhs % rhs | Returns lhs modulo rhs 
| lhs == rhs | Return true if lhs is the same as rhs 
| lhs != rhs | Return true if lhs is not the same as rhs 
| lhs > rhs | Returns true if lhs is greater than rhs 
| lhs >= rhs | Returns true if lhs is greater than or equal to rhs 
| lhs < rhs | Returns true if lhs is less than rhs 
| lhs <= rhs | Returns true if lhs is less than or equal to rhs 
| lhs && rhs | Returns true if lhs and rhs are true 
| lhs &#124;&#124; rhs | Returns true if lhs or rhs are true 
| lhs & rhs | Returns the bitwise AND of lhs and rhs 
| lhs &#124; rhs | Returns the bitwise OR of lhs and rhs 
| lhs << rhs | Returns lhs bit-shifted rhs bits to the left 
| lhs >> rhs | Returns lhs bit-shifted rhs bits to the right
| val++ | Post-increment: returns val, then increments val.
| val-- | Post-decrement: returns val, then decrements it.
| ^val | Returns the bitwise complement of val
| !val | Returns the negation of val

Operators not following standard Go/C behavior are documented below.

Anko follows C operator precedence rules.

#### + operator

The `+` operator will add numbers as expected, converting integers to floating point numbers if needed:

```
1 + 1		== 2
1.5 + 1		== 2.5
1 + true	== 2	# boolean 'true' evaluates to 1
1 + false 	== 1	# boolean 'false' evaluates to 0
```

It will also concatenate strings, converting types when it can:

```
"hello " + "world"	== "hello world"
"anko is #" + 1		== "anko is #1"
2.5 + "apples"		== "2.5apples"
"result is " + true	== "result is true"
```

It also joins arrays:

```
[1, 2] + [3, 4]			== [1 2 3 4]
["hi"] + ["there", 7]	== ["hi" "there" 7]
```

#### * operator

The `*` operator performs standard multiplication:

```
5 * 2		== 10
3.5 * 3		== 10.5
false * 2	== 0
```

It also performs string multiplication as in Python:

```
"hi" * 3	== "hihihi"
```

#### ** operator

The `**` operator performs exponentiation:

```
2**3	== 8		# 2 to the third power
10**4	== 10000	# 10 to the fourth power
```

### If statements

`if` statements in Anko behave as in Go:

```
if myBoolVar {
	println("myBoolVar is true")
}
if foo == 3 && !bar {
	println("foo is 3, bar is false")
} else if bar {
	println("bar is true")
} else {
	println("neither case")
}
```

Note: While Go allows the form `if result := foo(); result == true { ... }`, it is not acceptable in Anko.

#### Ternary operator

Anko allows the use of the ternary operator if desired:

```
result == 3 ? return true : return false
```

### For loops

`for` loops in Anko behave much as in Go.

A loop with a single condition:

```
for a < b {
	a *= 2
}
```

A loop with an initialization statement, a condition, and a post statement:

```
for i = 0; i < max; i++ {
	# do things
}
```

A loop with a range clause implementing the Fibonacci sequence (from the Anko repository examples):

```
func fib(n) {
  a, b = 1, 1
  f = []
  for i in range(n) {
    f += a
    b += a
    a = b - a
  }
  return f
}
```

### Switch statements

Anko's switch statements do **not** fall through by default, and there is no `fallthrough` statement as in Go to force a case to fall through.

Example:

```
switch foo {
case "a":
	# do things
case "b":
	# do other things
default:
	# base case
}
```

### Function declarations

Because Anko is dynamically typed, function definitions do not specify their return type or the types of the arguments. Due to implicit casting, this is typically fine, but if a function expects (e.g.) an array or map as an argument, we recommend inserting a comment above the function declaration to make this clear.

A function with no arguments:

```
func incrementCount() {
	counter++
}
```

A function which returns twice the argument given:

```
func double(x) {
	return 2*x
}
```

Functions can take variadic arguments. The arguments are presented to the function as an array. The following function prints the second argument passed to it and returns the total number of arguments that were passed:

```
func bar(x ...) {
	println(x[1])
	return len(x)
}

bar(10, 20, 30)		# prints "20", function returns 3.
```

### Concurrency

Anko can create goroutines using `go` statements. Communication between goroutines is typically accomplished via channels.

This example creates an unbuffered channel, then launches a new goroutine which will write three values to the channel. The original goroutine then goes on to read three values from the same channel. Because the channel is unbuffered, each write will block until the other goroutine issues a read.

```
c = make(chan int64)

go func() {
  c <- 1
  c <- 2
  c <- 3
}()

println(<-c)
println(<-c)
println(<-c)
```
