# Length

The `length` module calculates the length in bytes of the entry data or of an enumerated value.

```
length [-t target] [source]
```

`source` is an optional enumerated value name. Without it, length measures the entry data. The result goes into an enumerated value named `length`, or the name given with `-t`. If an entry doesn't have the `source` enumerated value, it passes through unchanged.

## Supported Options

* `-t <target>`: Write the length to the named enumerated value instead of `length`.

## Example Usage

| Command | Description |
|---------|-------------|
| length | Store the length of the entry data in `length` |
| length -t foo | Store the length of the entry data in `foo` |
| length Payload | Store the length of the enumerated value `Payload` in `length` |
| length -t foo Payload | Store the length of the enumerated value `Payload` in `foo` |
