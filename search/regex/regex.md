## Regex

Regex is a pipeline module that uses regular expressions to match text data. It is an extremely powerful way of matching complex patterns and extracting enumerable fields from text. For those unfamiliar with regular expressions, a decent starting point is [the Wikipedia article](https://en.wikipedia.org/wiki/Regular_expression).

Building regular expressions is well outside the scope of this document, but one important thing to note is that the `(?P<foo> .*)` style syntax which will assign any matched group into an enumerated field (named “foo” in this case). These enumerated fields are useful for performing analysis, charting, etc.

For example, the following search will enumerate the method, user, and ip from an sshd Accepted log entry.

```
".*sshd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"
```

Because regular expressions can get very long, the regex module takes the `-r` flag, which specifies a resource containing a regular expression. When populating the resource, do not include "wrapping quotes" around the whole expression as you would when typing directly into a search: e.g. `".*ssh.*Accepted"` becomes `.*ssh.*Accepted`. This is because the quotes are normally stripped out by the search parser prior to being handed to the regex module.

### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record. For example, a pipeline that showed packets not headed for port 80 but that have HTTP text would be `tag=pcap packet ipv4.DstPort!=80 tcp.Payload | regex -e Payload ".*GET \/ HTTP\/1.1.*"`
* `-r <arg>`: The “-r” option specifies that the regular expression statement is located in a resource file. 

Note: Storing especially large regular expressions in resource files can clean up queries, and allows for easy reuse.  If `-r` is specified, do not specify a regular expression in the query -- instead the contents of the resource will be used. Handy!

### Parameter Structure
```
regex <argument list> <regular expression>
```
### Example Search
```
tag=syslog grep sshd | regex *shd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"
```