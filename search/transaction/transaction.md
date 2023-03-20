# Transaction

```{note}
The `transaction` module can consume a large amount of memory. Use caution when using this module on memory constrained systems.
```

The `transaction` module transforms and groups entries in the pipeline into single-entry "transactions" - groupings of entries - based on any number of keys. It is a powerful tool for capturing the activity of a given user, IP, etc., across multiple entries in a datastream. 

## Supported Options

* `-e`: The `-e` option operates on an enumerated value instead of on the entire record. Multiple EVs are supported by providing additional `-e` flags.
* `-rsep`: The `-rsep` option sets the string to insert between transaction records. The default is "\n".
* `-fsep`: The `-fsep` option sets the string to insert between enumerated values within a given record. The default is " ".
* `-o`: The `-o` option sets the output EV to produce. The default is "transaction".
* `-c`: The `-c` option enables a count of the number of entries that make up a given transaction in the provided name. The default is "count".
* `-maxsize`: The `-maxsize` flag sets the maximum size, in kilobytes, of a given transaction before it is evicted from the tracking table (see "Memory considerations" below). The default is 500kb.
* `-maxstate`: The `-maxstate` flag sets the maximum number of transactions to track. Once exceeded, the oldest transaction will be evicted (see "Memory considerations" below). The default is 200.

All flags are optional.

## Overview

The `transaction` module groups entries into single entries based on a provided set of keys. For example, given a dataset with enumerated values "host", "message", and "action", the query:

```gravwell
tag=data kv host action message | transaction -fsep " -- " host | table
```

Will collapse all entries with the same value for the EV "host" into a single entry. By default, `transaction` will group all EVs that are *not* part of the key into the output. In the example above, the EVs "host" and "message" will be grouped, using `-fsep` as a separator, and all entries that match this key will be further grouped by `-rsep`. To illustrate the example above, given the following entries:

```
Entry 1: host="foo" message="Host foo login" action="login"
Entry 2: host="foo" message="Host foo delete file X" action="delete"
Entry 3: host="bar" message="Host bar login" action="login"
Entry 4: host="foo" message="Host foo logout" action="logout"
```

Will be collapsed into two entries, one for "foo", and another for "bar":

```
Entry 1: transaction="login -- Host foo login
                      delete -- Host foo delete file X
                      logout -- Host foo logout"
Entry 2: transaction="login -- Host bar login"
```

To specify exactly which EVs to group, you can use one or more `-e` flags in the query. EVs will be grouped in the order provided. For example:

```gravwell
tag=data kv host action message user group | transaction -e action -e message host | table
```

Will only group EVs "action" and "message", ignoring "user" and "group". 

Multiple keys can be provided, and records will be created based on the grouping of all provided keys. For example:

```gravwell
tag=data kv host action message user group | transaction host action user | table
```

Will group records with the same host, action, and user. 

## Memory considerations

The `transaction` module must buffer all entries in the datastream in order to create transactions. For queries that produce large amounts of data, this can quickly exhaust the available memory on a system. In order to prevent this, the `transaction` module provides two flags, `-maxsize`, and `-maxstate`, to control how much and how long to retain data before passing it downstream in the pipeline. 

When running, the `transaction` module keeps a table of records, with one record for every unique set of provided keys. When an entry matches the provided keys, it is added to other entries with the same match in the table (or creates a new record if it's the first one encountered). Two checks are asserted every time an entry is added to the table:

* If the size of a given record exceeds the `-maxsize` argument, the record is immediately "evicted" - meaning it is sent down the query pipeline and is removed from the table. 
* If the number of records exceeds the `-maxstate` argument, the _least recently updated_ record is evicted. 

If a record is evicted, and later an entry with a key matching that of the evicted record is encountered, a new record is created. If you notice "fragmentation" in your output, check the `-maxsize` and `-maxstate` flags. 

Because the `transaction` module can easily exhaust all available memory on your Gravwell system, follow these general guidelines when writing queries with `transaction`:

* Put the `transaction` module as late in the query as possible. 
* Work on the smallest time window possible for your query. 
* Start with small `-maxsize` and `-maxstate` values, and increase only if needed.
* Instead of grouping all enumerated values, only group those of concern for your query by explicitly naming them with `-e`.
