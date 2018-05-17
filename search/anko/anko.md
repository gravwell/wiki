## Anko

The anko module provides a more complete scripting environment as a supplement to eval. It allows more complex operations on search entries, but it also requires more work to develop, test, and deploy an anko script than a simple eval expression. Scripts are stored as resources in the [resource system](#!resources/resources.md).

The syntax of anko is identical to that of eval; both derive from [github.com/mattn/anko](https://github.com/mattn/anko), with some additional functions added for Gravwell-specific tasks.

We recommend using anko in situations where no other modules are capable enough. Typically this means situations where entries need to be compared against previous entries, entries need to be duplicated, complex operations are required to extract data from entries, or a combination of these.

This portion of the documentation only briefly describes the usage of the anko module; for a more detailed description, see the [full anko module documentation](#!scripting/anko.md) and the [Anko scripting language documentation](#!scripting/scripting.md).

### Syntax

`anko <script name> [script arguments]`

Anko scripts are stored as resources. The name of the resource must be specified as the first argument to the `anko` module. After the script name, any additional arguments are passed on to the script itself.

### Example script

The following script is a re-formatted version of an example from the eval module documentation. Note that it is far easier to read than the one-line eval example:

```
func Process() {
	if len(Body) <= 10 {
		setEnum("postlen", "short")
	} else if len(Body) > 10 && len(Body) < 300 {
		setEnum("postlen", "medium")
	} else {
		setEnum("postlen", "long")
	}
}
```

Assuming the script is uploaded to a resource named `CheckPostLen`, the script can be executed like this:

```
tag=reddit json Body | anko CheckPostLen | count by postlen | table postlen count
```

The `Process` function will be executed once for every search entry which reaches the anko module, checking the length of the enumerated value `Body` and setting a new enumerated value `postlen` based on the length of the body.

This example is quite simple; it implements only a `Process` function (not the optional `Parse` or `Finalize` functions). For more complex examples, refer to the [full anko module documentation](#!scripting/anko.md)