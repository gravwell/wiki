# Grep

Grep filters entries by text. Entries containing the pattern pass down the pipeline; entries that don't are dropped. For example, `grep foo` passes any entry containing "foo" anywhere. Matching is case sensitive, so `grep foo` matches "foo" but not "Foo".

Patterns support Go-style escape sequences such as `\t`, `\n`, and `\x01`, so you can match binary data.

Patterns can contain glob wildcards:

* `*` matches any sequence of bytes, including none. `grep foo*bar` matches entries containing "foo" followed by "bar".
* `?` matches any single character.
* `[abc]` and `[a-z]` match one character from a set or range; `[!abc]` matches one character not in the set.
* `{foo,bar}` matches any of the comma-separated alternatives.

You can give multiple patterns. By default an entry passes if *any* pattern matches. With `-s`, all patterns must match. With `-v`, the entry is dropped if any pattern matches.

## Supported options

* `-v`: Invert the match. `grep -v bar` drops entries containing "bar" and passes everything else.
* `-i`: Case-insensitive match. `grep -i foo` matches "Foo" and "foo".
* `-e <arg>`: Match against an enumerated value instead of the entire entry. Entries that don't have the enumerated value are dropped, even with `-v`.
* `-s`: Strict match. All patterns must match. Combined with `-v`, an entry is dropped if any pattern matches.
* `-simple`: Match the pattern literally, with no wildcards. Use this to find asterisks and other wildcard characters: `grep -simple "*"`.
* `-w`: Word match. The pattern must match a complete word, as split by the fulltext indexer. See "Working With Word Matches" below.

```{attention}
Case-insensitive search is significantly slower. If you need `-i`, put grep later in the pipeline so it sees fewer entries.
```

```{attention}
A backslash in a double-quoted or unquoted pattern is consumed as an escape character, so `"foo\*bar"` is the same as `"foo*bar"` and `*` is still a wildcard. Use `-simple` to match wildcard characters literally.
```

## Parameter Structure

```
grep <argument list> <pattern> [<pattern> ...]
```

## Example Search

To find Apache logs containing the literal string "Mozilla*Firefox" (no wildcards):

```gravwell
tag=apache grep -simple "Mozilla*Firefox"
```

To find HTTP requests in packets not destined for port 80:

```gravwell
tag=pcap packet tcp.DstPort!=80 tcp.Payload | grep -e Payload "GET / HTTP/1.1"
```

To find port 80 packets whose payloads contain the bytes 0x01, 0x02, 0x03, 0x04:

```gravwell
tag=pcap packet tcp.Port==80 tcp.Payload | grep -e Payload "\x01\x02\x03\x04"
```

Match any Reddit post containing a word ending in "ing" or "ed":

```gravwell
tag=reddit json Body | grep -e Body "*ing" "*ed"
```

Drop any Reddit posts on subreddits beginning with "Ask" or containing "foo":

```gravwell
tag=reddit json Subreddit | grep -v -e Subreddit "Ask*" "foo"
```

Keep only entries that contain both "Mozilla" and "Windows":

```gravwell
tag=apache grep -s Mozilla Windows
```

## Working With Word Matches

The `-w` flag matches complete words, using the same word boundaries as the fulltext indexer. Word patterns can't contain word-separator characters (spaces, `/`, `*`, `?`, brackets, and similar) and can't begin or end with `.`, `:`, `;`, or `-`; grep returns an error if they do. Because of this, wildcards can't be used with `-w`.

Grep can use the fulltext accelerator when it searches the entire entry (no `-e`) without `-v` or `-i`, and either has one pattern or uses `-s`. With `-w`, each whole pattern is looked up in the index. Without `-w`, only complete interior words of a pattern are used (for `"foo to the bar"`, "to" and "the"), so single-word and wildcard patterns aren't accelerated.

Consider this entry:

```
16.246.30.72 - - [08/May/2017:15:20:35 -0600] "DELETE /search/tag/list HTTP/1.0" 200 5032 "http://nguyen.biz/category/tags/tag/home.htm" "Opera/8.74.(Windows 98; Win 9x 4.90; it-IT) Presto/2.9.173 Version/11.00"
```

| Grep Invocation | MATCHES | Explanation |
|-----------------|---------|-------------|
| grep Ver        |   YES   | A plain grep matches the bytes `Ver` in `Version/11.00` |
| grep -w Ver     |   NO    | `Ver` is not a complete word |
| grep -w Ver*    |  ERROR  | `*` is a word separator, so it can't appear in a word match |
| grep -w Version |   YES   | `Version` is a complete word; `/` is a separator |
| grep -w "11.00" |   YES   | `.` is not a separator inside a word, only trimmed from its start or end, so IP addresses and version numbers stay whole |
| grep -w "Version/11.00" |  ERROR  | `/` is a word separator, so it can't appear in a word match |
