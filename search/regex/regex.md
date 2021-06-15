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
* `-v`: The "-v" option tells regex to operate in inverse mode, dropping any entries which match the regex and passing entries which do not match.
* `-p`: The "-p" option tells regex to allow entries through if the regular expression does not match at all.  The permissive flag does not change the operation of filters.

Note: Storing especially large regular expressions in resource files can clean up queries, and allows for easy reuse.  If `-r` is specified, do not specify a regular expression in the query -- instead the contents of the resource will be used. Handy!

### Inline Filtering

The regex module supports inline filtering to allow for down-selecting data directly within the regex module.  The inline filtering also enables regex to engage accelerators to dramatically reduce the amount of data that needs to be processed.  Inline filtering is achieved in the same manner as other modules by using comparison operators.  If a filter is enabled that specifies equality ("equal", "not equal", "contains", "not contains") any entry that fails the filter specification will be dropped entirely.  If a field is specified as not equal "!=" and the field does not exist, the field is not extracted but the entry won't be dropped entirely.


| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

#### Filtering Examples

```
tag=syslog regex "shd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)" user==root ip ~ "192.168"
```

### Parameter Structure
```
regex <argument list> <regular expression> <filter arguments>
```
### Example Search
```
tag=syslog grep sshd | regex *shd.*Accepted (?P<method>\S*) for (?P<user>\S*) from (?P<ip>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"
```
