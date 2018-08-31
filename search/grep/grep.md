## Grep

Grep is a very basic pipeline module that searches for a text string (not Unicode). Any record containing such text will match and be passed through the pipeline. Any record not containing the text is dropped from the pipeline. For example, `grep foo` will pass on any records containing the text “foo” and drop any records that do not have “foo” anywhere within. Grep is case sensitive so `grep foo` would match “foo” but drop “Foo”.  Grep also supports a standard set of escape codes similar to printf, allowing for binary filters as well.

Grep supports the standard GNU wildcards as well as fast string and binary matching.  To look for all entries that start that contain “foo” and “bar” separated by 0 or N bytes you can use `grep foo*bar`.  For more information on available wildcards see the TLDP miniguide[1].

Grep allows multiple patterns to be specified. If any pattern is matched, the entry is passed down the pipeline. If the `-v` flag is used to invert the search, the entry will be dropped if *any* pattern matches.

### Supported options

* `-v`: “Inverse” grep. For instance, `grep -v bar` would drop any records containing the text “bar” and pass on any records that do not contain “bar”.
* `-i`: Match case insensitive values. Thus, `grep -i foo` would match “Foo” and “foo”. Case insensitive search tends to be one of the slowest operations; put it later in your pipeline if possible to keep things fast.
* `-e <arg>`: Operate on an enumerated value instead of on the entire record. For example, a pipeline that showed packets that contain HTTP text but aren’t destined for port 80 would be `tag=pcap packet ipv4.DstPort!=80 tcp.Payload | grep -e Payload "GET / HTTP/1.1"`
* `-s`: Strict match.  All patterns must match, or in the case of a negated strict match, no pattern may match.`
* `-simple`: Simple match. With this flag, `grep` will match exactly the characters you specify, with no wildcard matching. This allows you to find asterisks and other normally-reserved characters: `grep -s *`

Attention: Case-insensitive search is significantly slower. If you must do case-insensitive grep, try to put it later in your search pipeline to improve speed.

### Parameter Structure
```
grep <argument list> <search parameter>
```
### Example Search

To find any Apache logs containing the exact string "Mozilla\*Firefox" (no wildcards):

```
tag=apache grep "Mozilla\*Firefox"
```

To find packets over port 80 whose payloads begin with the bytes 0, 1, 2, 3:

```
tag=pcap packet tcp.Port==80 tcp.Payload | grep -e Payload "\x01\x02\x03\x04"
```

Match any Reddit post which contains words ending in "ing" or "ed":

```
tag=reddit json Body | grep -e Body "*ing" "*ed"
```

Drop any Reddit posts on subreddits beginning with "Ask" or containing "foo":

```
tag=reddit json Subreddit | grep -v -e Subreddit "Ask*" "foo"
```

Grab only user agents that contain Mozilla and Windows

```
tag=apache grep -s Mozilla Apache
```
