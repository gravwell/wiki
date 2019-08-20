## Grep

Grep is a very basic pipeline module that searches for a text string (not Unicode). Any record containing such text will match and be passed through the pipeline. Any record not containing the text is dropped from the pipeline. For example, `grep foo` will pass on any records containing the text “foo” and drop any records that do not have “foo” anywhere within. Grep is case sensitive so `grep foo` would match “foo” but drop “Foo”.  Grep also supports a standard set of escape codes similar to printf, allowing for binary filters as well.

Grep supports the standard GNU wildcards as well as fast string and binary matching.  To look for all entries that start that contain “foo” and “bar” separated by 0 or N bytes you can use `grep foo*bar`.  For more information on available wildcards see the TLDP miniguide[1].

Grep allows multiple patterns to be specified. If any pattern is matched, the entry is passed down the pipeline. If the `-v` flag is used to invert the search, the entry will be dropped if *any* pattern matches.

### Supported options

* `-v`: “Inverse” grep. For instance, `grep -v bar` would drop any records containing the text “bar” and pass on any records that do not contain “bar”.
* `-i`: Match case insensitive values. Thus, `grep -i foo` would match “Foo” and “foo”. Case insensitive search tends to be one of the slowest operations; put it later in your pipeline if possible to keep things fast.
* `-e <arg>`: Operate on an enumerated value instead of on the entire record. For example, a pipeline that showed packets that contain HTTP text but aren’t destined for port 80 would be `tag=pcap packet ipv4.DstPort!=80 tcp.Payload | grep -e Payload "GET / HTTP/1.1"`
* `-s`: Strict match.  All patterns must match, or in the case of a negated strict match, no pattern may match.`
* `-simple`: Simple match. With this flag, `grep` will match exactly the characters you specify, with no wildcard matching. This allows you to find asterisks and other normally-reserved characters: `grep -s * `
* `-w`: A word match.  The entire match pattern must a word as would be matched by the fulltext extractors.

Attention: Case-insensitive search is significantly slower. If you must do case-insensitive grep, try to put it later in your search pipeline to improve speed.

Attention: The `-w` word match implies a simple match as the wildcards allow for crossing word boundaries.

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
tag=apache grep -s Mozilla Windows
```

### Working With Word Matches

The word match system is designed to match complete words.  Grep with the -w flag is one of the primary methods to interacting with the fulltext indexing system.

The -w flag is designed to create some additional specificity when selecting values, lets look at some example data to see what will and will not match.

```
16.246.30.72 - - [08/May/2017:15:20:35 -0600] "DELETE /search/tag/list HTTP/1.0" 200 5032 "http://nguyen.biz/category/tags/tag/home.htm" "Opera/8.74.(Windows 98; Win 9x 4.90; it-IT) Presto/2.9.173 Version/11.00"
```

Lets look at a few invocations of grep to see what would and would not match:

| Grep Invocation | MATCHES | Explanation |
|-----------------|---------|-------------|
| grep Ver        |   YES   | A simple grep WILL match due to the `Version/11.00` byte pattern |
| grep -w Ver     |   NO    | The word match will NOT match `Version/11.00` pattern because Ver is not a complete word |
| grep -w Ver*    |   NO    | The word match will NOT match because `-w` implies a simple match and literal bytes `Ver*` are not in the log` |
| grep -w Version |   YES   | The word match WILL match because Version is a full word, the `/` character is a split character |
| grep -w "11.00" |   YES   | The word will match, the `.` character is only a separator if it is followed by a space, this allows matching IP addresses |
| grep -w "Version/11.00" |  ERROR  | The grep module will throw an error, you cannot have word boundary characters in a match |
