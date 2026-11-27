# Strings

The strings module finds runs of printable ASCII characters (0x20 through 0x7E) of at least a minimum length in the entry data or an enumerated value. Each non-printable byte, including the bytes of multi-byte UTF-8 characters, becomes a period. Entries with no string of the required length are dropped.

```
strings [-t target] [-n length] [source]
```

`source` is an optional enumerated value name. Without it, strings uses the entry data. Entries that don't have the `source` enumerated value, or whose data is empty, are dropped. By default the result replaces the source; `-t` writes it to a different enumerated value instead.

## Supported Options

* `-t <target>`: Write the result to the named enumerated value instead of replacing the source.
* `-n <len>`: Minimum string length. The default is 6.

## Example Usage

| Command | Description |
|---------|-------------|
| strings | Operate directly on the entry data |
| strings -n 10 | Operate directly on the entry data, minimum string length 10 |
| strings -t foo | Convert the entry data and store the result in `foo`, rather than modifying the data |
| strings -t foo Payload | Convert the data in the enumerated value `Payload` and store it in `foo` |
