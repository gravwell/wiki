# Gravwell Notification Node

This node sends a Gravwell notification to the owner of the flow. Notifications consist of a message, an optional URL, and a notification ID. The notification ID is used to differentiate and deduplicate notifications; only one notification with a given ID will exist at one time, to prevent the same flow from creating hundreds of notifications when run repeatedly.

## Configuration

* `Message`, required: the text of the notification.
* `URL`: an optional URL. Setting this field will make the notification text clickable, opening the URL.
* `Notification ID`, required: an integer ID. This can be set to anything, but keep in mind that if multiple flows try to create notifications with the same ID, they will overwrite each other.

## Output

The node does not modify the payload.

## Example

Refer to the [Run Query](runquery.md) node documentation for a basic example of using the Notification node.
