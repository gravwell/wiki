# Transaction

```{note}
The `transaction` module can consume a large amount of memory. Use caution when using this module on memory constrained systems.
```

The `transaction` module groups entries that share the same key values into a single entry, called a transaction. Use it to collect the activity of a given user, IP, and so on across many entries.

(transaction-options)=
## Supported Options

* `-e <ev>`: an enumerated value to include in the transaction. Repeat the flag to include several, e.g. `-e action -e message`. By default, all enumerated values except the keys are included.
* `-rsep <string>`: the string inserted between records (one record per grouped entry). Defaults to `\n`.
* `-fsep <string>`: the string inserted between enumerated values within a record. Defaults to a single space.
* `-o <name>`: the name of the output enumerated value. Defaults to `transaction`.
* `-c <name>`: adds an enumerated value with this name holding the number of entries in each transaction. No count is added unless this flag is set.
* `-maxsize <n>`: the maximum size of a transaction, in kilobytes, before it is evicted from the tracking table (see "Memory considerations" below). Defaults to 500.
* `-maxstate <n>`: the maximum number of transactions to track. When exceeded, the least recently updated transaction is evicted (see "Memory considerations" below). Defaults to 200.
* `-maxcount <n>`: the maximum number of entries per transaction. When a transaction reaches this count, it is evicted. Unlimited by default.

All flags are optional.

## Overview

Entries are grouped by the keys given as arguments. For example, given a dataset with enumerated values `host`, `message`, and `action`, the query:

```gravwell
tag=data kv host action message | transaction -fsep " -- " host | table
```

collapses all entries with the same `host` into a single entry. By default, `transaction` includes every enumerated value that is *not* a key. Here, `action` and `message` from each entry are joined with `-fsep` to form a record, and the records for each `host` are joined with `-rsep`. Given the following entries:

```
Entry 1: host="foo" message="Host foo login" action="login"
Entry 2: host="foo" message="Host foo delete file X" action="delete"
Entry 3: host="bar" message="Host bar login" action="login"
Entry 4: host="foo" message="Host foo logout" action="logout"
```

the output is two entries, one for `foo` and one for `bar`:

```
Entry 1: transaction="login -- Host foo login
                      delete -- Host foo delete file X
                      logout -- Host foo logout"
Entry 2: transaction="login -- Host bar login"
```

To choose which enumerated values to include, use one or more `-e` flags. They are included in the order given. For example:

```gravwell
tag=data kv host action message user group | transaction -e action -e message host | table
```

includes only `action` and `message`, ignoring `user` and `group`.

With multiple keys, entries are grouped by the combination of all keys. For example:

```gravwell
tag=data kv host action message user group | transaction host action user | table
```

groups entries with the same `host`, `action`, and `user`. An entry that is missing a key is grouped using the keys it does have.

## Memory considerations

The `transaction` module must buffer entries to build transactions. Queries over large amounts of data can quickly exhaust available memory. The `-maxsize`, `-maxstate`, and `-maxcount` flags control how much data is held before it is passed down the pipeline.

The module keeps a table with one transaction for each unique set of key values. Each entry is added to the transaction for its keys, or starts a new one. Every time an entry is added, these checks run:

* If the transaction is larger than `-maxsize`, it is evicted immediately: it is sent down the pipeline and removed from the table.
* If the transaction has reached `-maxcount` entries, it is evicted.
* If the table holds more than `-maxstate` transactions, the _least recently updated_ transaction is evicted.

If a later entry has the same keys as an evicted transaction, a new transaction is started. If your output looks fragmented, check the `-maxsize`, `-maxstate`, and `-maxcount` flags.

To limit memory use:

* Put the `transaction` module as late in the query as possible.
* Search the smallest time window that works for your query.
* Start with small `-maxsize` and `-maxstate` values, and increase them only if needed.
* Name only the enumerated values you need with `-e` instead of including all of them.
