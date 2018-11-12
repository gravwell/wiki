## Eval

Eval is most commonly used for performing AND and OR logic on searches and enumerated values. However, the eval module is a bit of a Swiss Army knife, providing access to a limited subset of the Anko programming language (a dynamically-typed Go-like language, see [https://github.com/mattn/anko/](https://github.com/mattn/anko/) and the [Gravwell documentation for the Anko language](#!scripting/scripting.md)) to allow flexible operations on data within Gravwell. The eval module will execute exactly one expression or statement. In order to keep this page relatively simple, this section provides only a brief overview of some example eval invocations; more details are available [in this article](#!scripting/eval.md)

### Syntax

`eval <expression>`

The <expression> must be a single Anko expression, as described in [the eval documentation](#!scripting/eval.md).

### Examples

A simple application of the eval module might be to separate out Reddit comments which are less than 20 characters long. We do this by using the json module to extract the `Body` field, then passing it to eval with an expression which evaluates to true whenever the length of the comment’s `Body` field is less than 20. Only entries for which the expression evaluate to true are allowed to continue down the pipeline. Finally, we simply send the result of the eval to the table module to display the comment bodies which are less than 20 characters long.

```
tag=reddit json Body | eval len(Body) < 20 | table Body
```

AND and OR logic can be done in a similar manner to the length example above. For instance, if you have a source port and a destination port and are interested in verifying ranges on each to filter queries, the syntax might look something like:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort | eval ( (DstPort < 5000 && DstPort > 2000) || (SrcPort > 9000 && SrcPort > 8000) ) | table SrcIP SrcPort DstIP DstPort
```

A more complex example along similar lines looks at the relative frequency of different comment lengths. It sets an enumerated value, `postlen`, to “short” if the comment is 10 characters or less, “medium” if it’s between 10 and 300, and “long” if it’s longer. We then use the count module to tally up each length, and the table module to display the counts for each length.

```
tag=reddit json Body | eval if len(Body) <= 10 { setEnum("postlen", "short") } else if len(Body) > 10 && len(Body) < 300 { setEnum("postlen", "medium") } else { setEnum("postlen", "long") } | count by postlen | table postlen count
```

Note the use of the `setEnum()` function. It allows you to set an enumerated value which will be passed downstream. The enumerated value doesn’t have to be a string; it can be a boolean, a number, a slice, or even an IP.

The eval module supports `if` and `switch` logic statements, as demonstrated below:

```
if len(Body) <= 10 { setEnum("postlen", "short"); setEnum(“anotherEnum”, “foo”) }
```

```
switch DstPort { case 80: setEnum(“protocol”, “http”); case 22: setEnum(“protocol”, “ssh”); default: setEnum(“protocol”, “unknown”) }
```

### Further reference

* [The Gravwell documentation for the Anko language](#!scripting/scripting.md) is a generic description of the Anko scripting language
* [The eval module article](#!scripting/eval.md) describes the eval module in more detail.
