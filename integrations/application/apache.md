# Apache

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower](/ingesters/file_follow)
         Kit, [Apache Kit](https://github.com/gravwell/kits/tree/main/apache)
:::

## Apache Configuration
Apache features a highly configurable logging framework. Logging can be applied globally (on Debian/Ubuntu hosts, this is typically stored in `/etc/apache2.conf`) or for each virtual host (typically stored in `/etc/apache2/sites-available/{VHOST_NAME}.conf`). 

Follow the steps below to configure Apache to output clean data for Gravwell ingestion.

### Define LogFormat
Apache uses the Common Log Format by default, so we will define a custom structure called: `json_combined`. This will write the logs cleanly as JSON before ingestion into Gravwell.

You can define this globally in `/etc/apache2.conf` or inside a specific Virtual Host in `/etc/apache2/sites-available/{VHOST_NAME}.conf`.

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
\"referer\":\"%{Referer}i\"}" json_combined
```

#### Key Field Mechanics:
* `%a` uses the client IP address ( or the proxy-decode address if `mod_remoteip` is enabled).
* `%D` reports the request duration in microseconds. 
* `%q` includes the query string, complete with leading `?`. If no query string exists, outputs an empty string.

```{Note}
If you're using this in a template or config management tool, the `%{%Y-%m-%dT%H:%M:%S}t` strftime pattern contains `{%` which many templating engines will try to interpret — escape it appropriately for your tooling.
```

#### Apply custom logging
Once the format is defined, update the configuration file to apply it:

```apache
CustomLog /var/log/apache2/access.log json_combined
ErrorLog  /var/log/apache2/error.log
```

### Common Gotchas & Advance Tweaks

#### RewriteRule placement (Extesionless URLS)
If you're using `mod_rewrite` to handle extensionless URLs (e.g. routing `/status` to `/status.php`), your rules must must be placed inside a  `<Directory>` block rather than directly at the global VirtualHost level.

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
If Apache sits behind a proxy, the standard remote host field will log the proxy's internal IP address instead of the real clients. We can fix this using `mod_remoteip`.

1. **Enable the required modules:** Run the following commands to enable both the proxy headers and rewrite engines:

```bash
a2enmod remoteip
a2enmod rewrite   # if you need extensionless URL rewrites
```

2. **Configure mod_remoteip:** Create a dedicated configuration file at `/etc/apache2/conf-available/remoteip.conf`:

```apache
RemoteIPHeader X-Real-IP
RemoteIPTrustedProxy 172.18.100.60   # nginx's IP
```

3. **Activate the Configuration**

```
a2enconf remoteip
systemctl restart apache2
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