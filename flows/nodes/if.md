# If Node

This node executes a boolean logic expression specified by the user. If the expression evaluates to `true`, execution of downstream nodes will proceed. If the expression evaluates to `false`, none of the nodes directly downstream of this node will be executed; the node is said to be "blocking".

## Configuration

The configuration of the node defines the parts of a boolean logic expression, e.g. in `foo != 0`, "foo" is the left-hand side, "!=" is the operator, and "0" is the right-hand side.

* `Left-hand Side`, required: the item on the left side of the operator, frequently a variable.
* `Operator`, required: the logic operator to use, e.g. "!=", "==", ">".
* `Right-hand Side`, required: the item on the right side of the operator, frequently a constant.

## Output

The node does not modify the payload. However, if the expression evaluates to "false", the node will block execution of any downstream nodes.

## Example

Refer to the [Run Query](runquery.md) node's documentation for an example demonstrating the use of the If node.
