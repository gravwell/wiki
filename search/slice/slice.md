## Slice

The slice module is a powerful but very low-level tool for extracting bytes from entries or enumerated values by simply specifying offsets within the entry/enumerated value and optionally casting those bytes to specific type.  Slice can reference bytes via relative indexing, including negative numbers.

The following search extracts the RPM from canbus messages on a Toyota vehicle by reading the first two bytes of the "data" enumerated value and parsing it as a 16-bit big-endian integer.

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | mean RPM | chart mean
```

Slice can extract from raw entry contents, or it can operate on an enumerated value. The extracted bytes can optionally be parsed into a different type, such as an integer or a string. Some examples:

| Command | Description |
|---------|-------------|
| `slice [0:4] as foo` | Extract the first 5 bytes directly from the entry's data and place them into an enumerated value "foo" |
| `slice Payload[9] as tenth` | Extract the tenth byte from the enumerated value "Payload" |
| `slice uint16le(Payload[9:10]) as value` | Pull two bytes from the enumerated value "Payload", parse them as an unsigned 16-bit little-endian integer, and store it as "value" |
| `slice uint16be(Payload[-2:]) as value` | Pull the last two bytes from "Payload", parse them as an unsigned 16-bit big-endian integer, and store it as "value" |
| `slice uint16be(Payload[-4:-2]) as value2` | Pull the two bytes preceding the last two bytes of "Payload", parse them as an unsigned 16-bit big-endian integer, and store as "value2"

Supported Types

* byte
* int16
* int16le
* int16be
* uint16
* uint16le
* uint16be
* int32
* int32le
* int32be
* uint32
* uint32le
* uint32be
* int64
* int64le
* int64be
* uint64
* uint64le
* uint64be
* float32
* float32le
* float32be
* float64
* float64le
* float64be
* array
* string
