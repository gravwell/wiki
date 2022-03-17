# Flow Storage Read

This node, in conjunction with the [Flow Storage Write node](storagewrite.md) enables users to persist data across different *runs* of a flow. Every flow can have multiple key-value maps, which persist from execution to execution. This can be used to record the last time an event occurred, or track how many times a particular hostname has been seen, etc.

The Flow Storage Read node pulls elements back out of a key-value map. The Flow Storage Write node stores items into a key-value map. 

## Configuration

* `Map Name`, required: The name of the map to access. This can be almost anything, such as "state" or "previously seen systems".
* `Items`: a collection of keys to extract from the specified map. They will be placed into the payload with the same name.

## Output

The node will insert an item into the output payload for each key extracted from the map. Thus if the value "foo" was extracted from a map, the outgoing payload will contain an item named "foo".

## Example

Refer to the [Flow Storage Write](storagewrite.md) node documentation for an example showing the use of both nodes.
