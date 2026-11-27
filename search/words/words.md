# Words

Words matches whole words, as delimited by split characters. By default it behaves like `grep -s -w`, and it is designed to use the fulltext accelerator. Words handles UTF-8 and usually behaves well with binary data, so you can look for the word "foo" in a pcap stream. Watch the word boundaries, though: if `foo` is next to the byte 0x44 (`D`) in binary data, the word is `Dfoo` and won't match `foo`. Words is a good first-level filter on unknown data.

Words does not support wildcards.

You can give multiple words. By default every word must match for the entry to pass. With `-or`, the entry passes if any word matches. With `-v`, the logic is inverted.

## Supported options

* `-v`: Invert the match. `words -v bar` drops entries containing the word "bar" and passes the rest. With multiple words, `-v` drops only entries that contain *all* of them; combine it with `-or` to drop entries that contain *any* of them.
* `-e <arg>`: Match against an enumerated value instead of the entire entry. Entries without that enumerated value are dropped, even with `-v`. For example, to find HTTP requests in packets not destined for port 80: `tag=pcap packet tcp.DstPort!=80 tcp.Payload | words -e Payload GET HTTP 1.1`
* `-or`: Pass the entry if any word matches. With `-v`, drop the entry if any word matches.

## Parameter Structure
```
words <argument list> <word> [<word> ...]
```

## Example Search

To find any Apache logs containing the words `Mozilla` and `Firefox`:

```gravwell
tag=apache words Mozilla Firefox
```

To find port 80 packets whose payloads contain neither the word HTTP nor 1.1:

```gravwell
tag=pcap packet tcp.Port==80 tcp.Payload | words -e Payload -v -or HTTP 1.1
```

Match any Reddit post containing the word `Gravwell` or `gravwell`:

```gravwell
tag=reddit json Body | words -e Body -or Gravwell gravwell
```

Keep only entries containing both Mozilla and Windows:

```gravwell
tag=apache words Mozilla Windows
```

(working-with-word-matches-target)=
## Working With Word Matches

The word match system is designed to match complete words, using the exact same word-breaking rules as the [fulltext accelerator](#fulltext-word-extraction-target). A query word matches only when it appears in the data bounded on both sides by a split character (or the start/end of the data). Most punctuation is a split character, but `.` (period), `:` (colon), `-` (hyphen), `_` (underscore), and `@` (at sign) are treated as part of a word. This is what allows IP addresses, decimals, hostnames, and email addresses to be matched as single words. Leading and trailing `.`, `:`, `;`, and `-` are trimmed from a word, so a sentence-ending period does not prevent a match.

The words module cannot match a *substring* of a word. Because `.` and `:` are not split characters, `192.168.1.100:8080` is a single word; a search for `192.168.1.100` will NOT match it. For substring matching, see the [grep module](/search/grep/grep). For a complete list of split and trim characters, see the [fulltext word extraction](#fulltext-word-extraction-target) documentation.

Consider this entry:

```
16.246.30.72 - - [08/May/2017:15:20:35 -0600] "DELETE /search/tag/list HTTP/1.0" 200 5032 "http://nguyen.biz/category/tags/tag/home.htm" "Opera/8.74.(Windows 98; Win 9x 4.90; it-IT) Presto/2.9.173 Version/11.00"
```

| Words Invocation | MATCHES | Explanation |
|-----------------|---------|-------------|
| words Ver       |   NO    | The words module will NOT match the `Version/11.00` text because `Ver` is not a complete word; the words module cannot match a subset of a word |
| words 16.246.30.72   |   YES    | The words module will match IPs, because the `.` character is not a split character |
| words 8.74   |   YES    | The words module will match the `8.74` value in `Opera/8.74.(Windows` even though it is followed by a `.`. The `/` and `(` are split characters and the trailing `.` is a trim character, so the word `8.74` is isolated. |
| words Version   |   YES   | The words module WILL match because `Version` is a full word, the `/` character is a split character |
| words 11.00     |   YES     | The word will match.  The `.` character is not a split character, so `11.00` is a complete word bounded by the `/` before it and the `"` after it |
| words Windows   |   YES   | `Windows` is a complete word; the `(` before it and the space after it are split characters |
| words "8.74.(Windows"   |  ERROR  | The words module will throw an error because the search term contains the split character `(`; you cannot have a split character inside a single match word |
| words "Version/11.00"   |  ERROR  | The words module will throw an error, you cannot have a split character (here, `/`) in a match word | 
| words "Ver*"   |  ERROR  | The words module does not support wildcards, and `*` is itself a split character, so the search term is rejected. |
