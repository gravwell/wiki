## Length

The `length` module calculates the length (in bytes) of either the entry data or an enumerated value. The syntax is:

```
length [-t target] [source]
```

The `source` parameter is an optional enumerated value name on which to operate; if no source is specified, length will use the raw entry data. By default, the length of the data or enumerated value is written to an enumerated value named "length". Specifying a `target` will instead write the length out to an enumerated value of that name.

### Supported Options

* `-t <target>`: Write the computed length to a specified enumerated value instead of the default named "length".

### Example Usage

| Command | Description |
|---------|-------------|
| length | Get the length in bytes of the entry data and store the result in the default enumerated value `length` |
| length -t foo | Calculate the length of the entry data and store the result in an enumerated value named `foo` instead of `length` |
| length Payload | Find the length of the enumerated value `Payload` and store it in an enumerated value named `length` |
| length -t foo Payload | Find the length of the enumerated value `Payload` and store it in `foo` |
