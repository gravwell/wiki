# Notifications REST API
A notification is basically a Message and potentially a Type.  The type is used so that something can continually send notifications without duplicating it.  E.g. The webserver can keep adding the notification (Indexer X is down) with Type 0xDEADBEEF.  The notification engine will just keep replacing the old notification that has the same Type.

Notifications can come in two forms, targeted and broadcast.  Broadcast go to everyone, targeted only go to users with the appropriate UID or GID.  Broadcast notificatons inherently have no UID/GID (they will always be zero).

Notifications will expire (deleted) at their expires date.

Notifications with a IgnoreUntil date in the future won't come back on requests (hidden)

If a notification is added with blank dates (Sent, Expires, IgnoreUntil) they are populated with default values.

All notifications are ephemerial.  If the webserver/frontend reboots, they are lost.

All notifications have an ID, and each ID monotonically increases and is always represented as a base10 uint64.

The basic REST APIs URLs are:
```
/api/notifications/all/{id:[0-9]+}
/api/notifications/{id:[0-9]+}
/api/notifications/broadcast
/api/notifications/targeted
/api/notifications/targeted/user/{id:[0-9]+}
/api/notifications/targeted/group/{id:[0-9]+}
```

Methods for requests
GET - pull back notifications
PUT - update a specific notification
POST - add a new notification
DELETE - remove a specific notification

### Notification rules
* Only admins can add new notifications
* Users can only update notifications that they explicitely own (UIDs match).  GID match isn't enough
* Users cannot change the UID or GID associated with a notification
* Users cannot delete notifications they don't own (UID does not match)

## Adding a new notification

## Getting notifications
Users can get all their notifications by hitting /api/notifications/{id}  with no id (e.g. /api/notifications/).  To get all notifications after a specific ID issue a GET on /api/notifications/{id}.  For example if you pull back notifications and the largest ID is 10, you can make a request for /api/notifications/10 and you will only get NEW notifications that the user has access to.

### Admins requesting all notifications
Admins can get all notifications (irrespective of UID or GID) by issuing a GET on /api/notifications/all/ with no id.