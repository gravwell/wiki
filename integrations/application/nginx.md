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

Nginx's default `combined` format is space-delimited. To better setup for Gravwell ingestion we recommend replacing it with a `log_format` directive that produces one JSON object per request, then apply that format to each vhost.

In `/etc/nginx/nginx.conf` (inside the `http {}` block):

```ini
log_format json_access escape=json
    '{'
    '"time":"$time_iso8601",'
    '"remote_addr":"$remote_addr",'
    '"method":"$request_method",'
    '"uri":"$uri",'
    '"status":$status,'
    '"bytes_sent":$bytes_sent,'
    '"request_time":$request_time,'
    '"upstream":"$upstream_addr",'
    '"user_agent":"$http_user_agent",'
    '"referer":"$http_referer"'
    '}';
```

Then in each vhost (or the default server block):

```nginx
access_log /var/log/nginx/access.log json_access;
error_log  /var/log/nginx/error.log warn;
```

### Key Parameters
* `escape=json` parameter is critical. Without it, special characters inside User Agents or URIs will break JSON parsing downstream. 
* `upstream` field is empty for directly served content and populated for proxied requests, which lets you distinguish traffic at query time.


### Proxy Configuration
If nginx is acting as a reverse proxy, add these to the proxy location block so the backend sees the real client IP:

```nginx
proxy_set_header X-Real-IP       $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
```

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