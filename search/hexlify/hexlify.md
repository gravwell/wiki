# Hexlify

The `hexlify` module encodes data as ASCII hex, or decodes ASCII hex with `-d`. Viewing unknown binary data as hex can help reveal its structure. For example, hex-encoding CAN bus payloads makes it easier to spot bytes that change in predictable patterns, such as speed or throttle position.

With no arguments, `hexlify` converts the entry data in place. Given one or more enumerated value names, it converts each one in place, or writes the result to a new enumerated value when followed by `as <name>`. Enumerated values are encoded from their raw bytes, so non-string types such as integers produce the hex of their internal binary form. Results stored in enumerated values are strings.

Entries missing any of the listed enumerated values are dropped.

## Supported Options

* `-d`: Decode ASCII hex back into raw bytes rather than encoding. Entries whose value is not valid hex, including hex with an odd number of digits, are dropped.

## Examples

### Search to hexlify all data

```gravwell
tag=stuff hexlify
```

### Search to hexlify a single enumerated value

```gravwell
tag=CAN canbus ID data | hexlify data | table ID data
```

### Search to hexlify all data and assign to a new name

```gravwell
tag=stuff hexlify DATA as hexdata | table DATA hexdata
```

### Search to hexlify a few enumerated values with reassignment

```gravwell
tag=CAN canbus ID data | hexlify ID as hexid data as hexdata | table ID hexid data hexdata
```

### Decoding hex data

```gravwell
tag=apache json val | hexlify -d val as decodedval | table val decodedval
```
