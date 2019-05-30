# Controlling System Resource Usage

This section describes how you can tune the Gravwell system's consumption of system resources.

## Webserver Renderer Storage

Render modules such as chart and table store their results on-disk on the webserver. Depending on the search, the results may consume significant amounts of disk space. To prevent users from filling up the disk with over-large queries, use the `Render-Store-Limit` parameter in the webserver's gravwell.conf file. Setting `Render-Store-Limit=64`, for instance, would set a limit of 64 MB of on-disk storage per query.

## Limiting Gravwell Resource Size

User-created [resources](#!resources/resources.md) can take up a lot of space on disk, on both the webserver and the indexers. The `Resource-Max-Size` parameter in gravwell.conf specifies a limit, in bytes, for resource size. Thus, setting `Resource-Max-Size=20971520` will limit resources to no more than 20 megabytes.
