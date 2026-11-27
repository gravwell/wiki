# Src

The src module filters entries by their source, a metadata field every entry has. Use it to look at entries from a specific host or network.

```
src [<op>] <value>
```

The operator is one of `==`, `!=`, `~` (in subnet), or `!~` (not in subnet). The value can be:

* An IPv4 or IPv6 address, such as `192.168.1.1`.
* A subnet in CIDR notation, such as `192.168.0.0/16`. A subnet with `==` or `!=` is treated as `~` or `!~`. `~` and `!~` require a subnet.
* A decimal integer, or a hex integer with a `0x` prefix, such as `1` or `0xfeadbeef`.
* A hex string of up to 32 digits (16 bytes) with an even number of digits, such as `1234567890ABCDEF0011223344556677`.

Without an operator, `src <value>` is an equality match, or a subnet match if the value is a subnet.

```{note}
The accelerator can use the source field to speed up queries, but only for equality matches. Subnet matches and negations are not accelerated.
```

## Example Usage

Eliminate entries coming from a specific source:

```gravwell
tag=syslog,apache,pcap src != 192.168.1.1 | count by TAG | chart count by TAG
```

Select only those entries coming from a specific subnet:

```gravwell
tag=syslog,apache,pcap src ~ 192.168.1.0/24 | count by SRC | chart count by SRC
```

Eliminate entries from a specific subnet:

```gravwell
tag=syslog src !~ 192.168.0.0/16
```

Select only entries whose source is the integer ID 1:

```gravwell
tag=syslog src == 1
```

Eliminate entries whose source is a hex-encoded integer ID:

```gravwell
tag=syslog src != 0xfeadbeef
```

Select entries whose source is a 16-byte hex string:

```gravwell
tag=syslog src == 1234567890ABCDEF0011223344556677
```
