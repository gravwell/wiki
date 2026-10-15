# Apache

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
    Kit, [Apache Kit](https://github.com/gravwell/kits/tree/main/apachehttp)
:::

## Apache Configuration

Apache defines its log locations via the `ErrorLog` and `CustomLog` directives, which can be configured globally (typically `/etc/apache2/apache2.conf`) or per virtual host (typically `/etc/apache2/sites-available/{VHOST_NAME}.conf`). By default, these write to `error.log` and `access.log` in `/var/log/apache2/`. Follow the steps below to configure Apache to output clean data for Gravwell ingestion, then install File Follower on your Apache host by following the instructions in [File Follower](/ingesters/file_follow). Then add the configuration from the File Follower section below.

### Define LogFormat
Apache uses the Common Log Format by default. Define a custom format named `json_combined` that writes the logs cleanly as JSON before ingestion into Gravwell.

You can define this globally in `/etc/apache2/apache2.conf` or inside a specific Virtual Host in `/etc/apache2/sites-available/{VHOST_NAME}.conf`.

```apache
LogFormat \
  "{\"time\":\"%{%Y-%m-%dT%H:%M:%S}t\",\
  \"remote_addr\":\"%a\",\
  \"method\":\"%m\",\
  \"uri\":\"%U\",\
  \"query\":\"%q\",\
  \"status\":%>s,\
  \"bytes\":%B,\
  \"duration\":%D,\
  \"user_agent\":\"%{User-Agent}i\",\
  \"referer\":\"%{Referer}i\",\
  \"vhost\":\"%v\"}" json_combined
```

#### Key Field Mechanics
* `%a` uses the client IP address (or the address decoded by `mod_remoteip`, if it is enabled).
* `%D` reports the request duration in microseconds. 
* `%q` includes the query string, including the leading `?`. If no query string exists, outputs an empty string.
* `%v` logs the `ServerName` of the virtual host that served the request, so sites on the same server can be told apart. Requests that don't match a named virtual host (for example, a request to the server's bare IP) are logged under the default virtual host's `ServerName`. Give every `<VirtualHost>` its own `ServerName`: one without it logs the server's global name, so several sites end up sharing one value.

```{note}
If you're using this in a template or config management tool, the `%{%Y-%m-%dT%H:%M:%S}t` strftime pattern contains `{%` which many templating engines will try to interpret — escape it appropriately for your tooling.
```

#### Apply Custom Logging
Once the format is defined, update the configuration file to apply it:

```apache
CustomLog /var/log/apache2/access.log json_combined
ErrorLog  /var/log/apache2/error.log
```

### Common Gotchas and Advanced Tweaks

#### RewriteRule Placement (Extensionless URLs)
If you're using `mod_rewrite` to handle extensionless URLs (e.g. routing `/status` to `/status.php`), your rules must be placed inside a `<Directory>` block rather than directly at the global VirtualHost level.

At the global VirtualHost level, `%{REQUEST_FILENAME}` treats the target as a plain URI string instead of a filesystem path. This causes `-f` (file) and `-d` (directory) checks to always evaluate false:

```apache
<Directory /var/www/html>
    RewriteEngine On
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteRule ^(.+)$ $1.php [L]
</Directory>
```

#### Proxy Configuration
If Apache sits behind a proxy, `remote_addr` (`%a`) will log the proxy's IP for every request unless you configure `mod_remoteip` first.

1. **Enable the required modules:** Run the following commands to enable both the proxy headers and rewrite engines:

```bash
a2enmod remoteip
a2enmod rewrite   # if you need extensionless URL rewrites
```

2. **Configure mod_remoteip:** Create a dedicated configuration file at `/etc/apache2/conf-available/remoteip.conf`:

```apache
# Replace 192.0.2.10 with your reverse proxy's IP address or CIDR range
RemoteIPHeader X-Real-IP
RemoteIPInternalProxy 192.0.2.10
```

Use `RemoteIPInternalProxy` when your proxy is on your own network. `RemoteIPTrustedProxy` won't accept private client addresses (10/8, 172.16/12, 192.168/16, 127/8) from the header, so internal clients would be logged with the proxy's IP instead of their own.

3. **Activate the Configuration**

```bash
a2enconf remoteip
systemctl restart apache2
```

4. **Log the forwarding header:** If Apache is behind a proxy, add a `forwarded_for` field to the LogFormat, between `referer` and `vhost`:

```text
\"forwarded_for\":\"%{X-Forwarded-For}i\",\
```

The full LogFormat for a proxied server:

```apache
LogFormat \
  "{\"time\":\"%{%Y-%m-%dT%H:%M:%S}t\",\
\"remote_addr\":\"%a\",\
\"method\":\"%m\",\
\"uri\":\"%U\",\
\"query\":\"%q\",\
\"status\":%>s,\
\"bytes\":%B,\
\"duration\":%D,\
\"user_agent\":\"%{User-Agent}i\",\
\"referer\":\"%{Referer}i\",\
\"forwarded_for\":\"%{X-Forwarded-For}i\",\
\"vhost\":\"%v\"}" json_combined
```

`forwarded_for` records the proxy chain a request came through. Requests that reach Apache without going through the proxy log `-`, which lets the Apache kit spot clients bypassing the proxy. Only add this field if Apache is behind a proxy. On a server with no proxy every request logs `-`, and the bypass check would flag every client.

```{note}
Whoever sends the request sets this header, so a client can put anything in it. Treat it as context, not proof of where a request came from. `remote_addr` stays the reliable client address, because `mod_remoteip` only accepts it from your configured proxy.
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/apache-well.conf`
```ini
[Storage-Well "apache"]
    Location=/opt/gravwell/storage/apache
    Tags=apache*
```
### Gravwell Ingester Configuration: File Follower
**Sample Apache config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/apache.conf`
```ini
[Follower "apache-access"]
    Base-Directory = /var/log/apache2
    File-Filter    = access.log
    Tag-Name       = apache

[Follower "apache-error"]
    Base-Directory = /var/log/apache2
    File-Filter    = error.log
    Tag-Name       = apache-err
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```
