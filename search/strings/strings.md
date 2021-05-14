## Strings

The strings module finds strings of at least a minimum length in either the entry data or an enumerated value. It converts non-printable characters into periods and drops any entries which contain no strings of the required length. The syntax is:

```
strings [-t target] [-n length] [source]
```

The `source` parameter is an optional enumerated value name on which to operate; if no source is specified, strings will use the raw entry data. By default, the converted data is written back to the source. Specifying a `target` will instead write the converted data out to an enumerated value of that name. The `length` parameter allows the user to select the minimum string length; the default is 6 characters.

### Supported Options

* `-t <target>`: Write converted strings data to a specified enumerated value
* `-n <len>`: Select minimum string length (default 6 characters)

### Example Usage

| Command | Description |
|---------|-------------|
| strings | Operate directly on the entry data |
| strings -n 10 | Operate directly on the entry data, minimum string length 10 |
| strings -t foo | Convert the entry data and store the result in `foo`, rather than modifying the data |
| strings -t foo Payload | Convert the data in the enumerated value `Payload` and store it in `foo` |
