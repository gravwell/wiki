# Nginx

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
    Kit, [Nginx Kit](https://github.com/gravwell/kits/tree/main/nginx)
:::

## Nginx Configuration

Nginx's default `combined` log format is space-delimited. To better prepare logs for Gravwell ingestion, replace it with a `log_format` directive that produces one JSON object per request. Apply that format to each virtual host.

In `/etc/nginx/nginx.conf` (inside the `http {}` block):

```text
log_format json_access escape=json
    '{'
    '"time":"$time_iso8601",'
    '"remote_addr":"$remote_addr",'
    '"method":"$request_method",'
    '"uri":"$uri",'
    '"query":"$query_string",'
    '"status":$status,'
    '"bytes_sent":$bytes_sent,'
    '"request_time":$request_time,'
    '"upstream":"$upstream_addr",'
    '"user_agent":"$http_user_agent",'
    '"referer":"$http_referer",'
    '"vhost":"$server_name"'
    '}';
```

Then, in each virtual host (or the default server block):

```nginx
access_log /var/log/nginx/access.log json_access;
error_log  /var/log/nginx/error.log warn;
```

### Key Parameters
* The `escape=json` parameter is critical. Without it, special characters inside user agents or URIs will break JSON parsing downstream.
* The `upstream` field is empty for directly served content and populated for proxied requests, which lets you distinguish traffic at query time.

The `query` field logs the query string separately from `uri`, because Nginx's `$uri` is the path only. Attack payloads such as SQL injection, XSS and open redirects usually arrive in the query string, and the Nginx kit's detections read them from this field.

The `vhost` field logs the name of the `server` block that handled the request, so sites on the same server can be told apart. It's the **first** name listed in that block's `server_name`, so list each site's main hostname first. A catch-all block (`server_name _;`) logs `_`, which marks requests that didn't match any named site, such as scanners hitting the server's IP. If `server_name` is a regular expression, the expression itself is logged, not the hostname that matched.


### Proxy Configuration
If Nginx is acting as a reverse proxy, add these to the proxy location block so the backend sees the real client IP:

```nginx
proxy_set_header X-Real-IP       $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
```

Nginx writes logs to `access.log` and `error.log` in `/var/log/nginx/` by default. Install File Follower on your Nginx host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/nginx-well.conf`
```ini
[Storage-Well "nginx"]
    Location=/opt/gravwell/storage/nginx
    Tags=nginx*
```
### Gravwell Ingester Configuration: File Follower
**Sample Nginx config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/nginx.conf`
```ini
[Follower "nginx-access"]
    Base-Directory = /var/log/nginx
    File-Filter    = access.log
    Tag-Name       = nginx
    Assume-Local-Timezone = false
    Ignore-Timestamps = false

[Follower "nginx-error"]
    Base-Directory = /var/log/nginx
    File-Filter    = error.log
    Tag-Name       = nginx-err
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```