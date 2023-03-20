## Regex Router Preprocessor

The regex router preprocessor is a flexible tool for routing entries to different tags based on the contents of the entries. The configuration specifies a regular expression containing a [named capturing group](https://www.regular-expressions.info/named.html), the contents of which are then tested against user-defined routing rules.

The Regex Router preprocessor Type is `regexrouter`.

### Supported Options

* `Regex` (string, required): This parameter specifies the regular expression to be applied to the incoming entries. It must contain at least one [named capturing group](https://www.regular-expressions.info/named.html), e.g. `(?P<app>.+)` which will be used with the `Route-Extraction` parameter.
* `Route-Extraction` (string, required): This parameter specifies the name of the named capturing group from the `Regex` parameter which will contain the string used to compare against routes.
* `Route` (string, required): At least one `Route` definition is required. This consists of two strings separated by a colon, e.g. `Route=sshd:sshlogtag`. The first string ('sshd') is matched against the value extracted via regex, and the second string defines the name of the tag to which matching entries should be routed. If the second string is left blank, entries matching the first string *will be dropped*.
* `Drop-Misses` (boolean, optional): By default, entries which do not match the regular expression will be passed through unmodified. Setting `Drop-Misses` to true will make the ingester drop any entries which 1) do not match the regular expression, or 2) match the regular expression but do not match any of the specified routes.

### Example: Routing to Tag Based on App Field Value

To illustrate the use of this preprocessor, consider a situation where many systems are sending syslog entries to a Simple Relay ingester. We would like to separate out the sshd logs to a separate tag named `sshlog`. Incoming sshd logs are in old-style BSD syslog format (RFC3164):

```
<29>1 Nov 26 11:26:36 localhost sshd[11358]: Failed password for invalid user administrator from 202.198.122.184 port 49828 ssh2
```

By experimenting with regular expressions, we find that the following is a reasonable regular expression to extract the application name (e.g. sshd) from RFC3164 logs into a capturing group named "app":

```
^(<\d+>)?\d?\s?\S+ \d+ \S+ \S+ (?P<app>[^\s\[]+)(\[\d+\])?:
```

We can apply that regular expression to a preprocessor definition, as shown below:

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
        # Regex: <pri>version Month Day Time Host App[pid]
        Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog
```

Note that the preprocessor defines the regular expression, then calls out the capturing group "app" in the `Route-Extraction` parameter. It then uses the `Route=ssh:sshlog` definition to specify that those entries whose application name matches "sshd" should be routed to the tag "sshlog". We could define additional `Route` parameters as needed, e.g. `Route=apache:apachelog`.

With the above configuration, logs from sshd will be sent to the "sshlog" tag, while all other logs will go straight to the "syslog" tag. We could extract other applications from similarly-formatted syslog entries by adding additional `Route` specifications, but suppose we had some intermingled logs in RFC 5424 format, as shown below?

```
<101>1 2019-11-26T13:24:56.632535-07:00 web01.example.org webservice 21581 - [useragent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3191.0 Safari/537.36"] GET /
```

The regular expression we already have won't extract the application name ("webservice") properly, but we can define a *second* preprocessor and put it in the preprocessor chain after the existing one:

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter
        Preprocessor=rfc5424router

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
        # Regex: <pri>version Month Day Time Host App[pid]
        Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog

[preprocessor "rfc5424router"]
        Type=regexrouter
        Drop-Misses=false
        # Regex: <pri>version Date Host App
        Regex="^<\\d+>\\d? \\S+ \\S+ (?P<app>\\S+)"
        Route-Extraction=app
        Route=webservice:weblog
        Route=apache:weblog
        Route=postfix:		# drop
```

Note that this new preprocessor definition defines routes for the applications named "webservice" and "apache", sending both to the "weblog" tag. Note also that it specifies that logs from the "postfix" application should be *dropped*, perhaps because those logs are already being ingested from another source.


