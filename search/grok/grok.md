# Grok

The grok module allows you to extract data from complicated text structures without specifying the whole regular expression every time. Instead, grok assigns names to regular expressions, allowing the user to specify the name instead. Grok patterns may contain additional patterns nested within them, making it easy to build up new definitions. It pre-defines a selection of useful patterns, but can also read your own customized set of patterns from a [resource](#!resources/resources.md).

By default, grok passes through any entry which matches the pattern and drops any which does not. This behavior can be inverted with the `-v` flag.

Grok is a filtering module; after specifying the desired pattern, you may also specify a list of filters to apply to the extracted fields.

Note: Because some filters (such as COMBINEDAPACHELOG) incorporate many nested regular expressions, they can be relatively slow when processing large numbers of entries. Use modules such as [grep](#!search/grep/grep.md), [regex](#!search/regex/regex.md), and [words](#!search/words/words.md) to pre-filter as much as possible.

## Supported Options

* `-e <arg>`: Operate on the specified enumerated value instead of the entire record.
* `-r <resource>`: load custom grok patterns from the resource with the specified name, rather than the default "grok" resource.
* `-v`: Operate in inverse mode; entries which do *not* match the pattern will be passed, and entries which *do* match will be dropped. You cannot specify any filters when using this flag.

### Parse Apache Logs

The following query finds all Apache logs for "PUT" requests and parses them out into their components:

```
tag=apache grep PUT | grok "%{COMBINEDAPACHELOG}" | table
```

![](apache.png)

Note that this query may take some time if you have millions of entries, since the COMBINEDAPACHELOG pattern is complex.

### Filtering

We can build on the previous query to return only those entries whose "clientip" field matches a particular IP:

```
tag=apache grep PUT | grok "%{COMBINEDAPACHELOG}" clientip=="128.10.247.36" | table clientip
```

![](apache-filter.png)

## Pre-defined Patterns
