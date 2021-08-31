## Slice

The slice module is a powerful but very low-level tool for extracting bytes from entries or enumerated values by simply specifying offsets within the entry/enumerated value and optionally casting those bytes to specific type.  Slice can reference bytes via relative indexing, including negative numbers.

The following search extracts the RPM from canbus messages on a Toyota vehicle by reading the first two bytes of the "data" enumerated value and parsing it as a 16-bit big-endian integer.

```
tag=CAN canbus ID=0x2C4 | slice uint16be(data[0:2]) as RPM | mean RPM | chart mean
```

Slice can extract from raw entry contents, or it can operate on an enumerated value. The extracted bytes can optionally be parsed into a different type, such as an integer or a string. Some examples:

| Command | Description |
|---------|-------------|
| `slice [0:4] as foo` | Extract the first 5 bytes directly from the entry's data and place them into an enumerated value "foo" |
| `slice Payload[9] as tenth` | Extract the tenth byte from the enumerated value "Payload" |
| `slice uint16le(Payload[9:10]) as value` | Pull two bytes from the enumerated value "Payload", parse them as an unsigned 16-bit little-endian integer, and store it as "value" |
| `slice uint16be(Payload[-2:]) as value` | Pull the last two bytes from "Payload", parse them as an unsigned 16-bit big-endian integer, and store it as "value" |
| `slice uint16be(Payload[-4:-2]) as value2` | Pull the two bytes preceding the last two bytes of "Payload", parse them as an unsigned 16-bit big-endian integer, and store as "value2"

### Supported Types

An integral function of the slice module is casting the data to the appropriate type.  By default, data is extracted as a byte slice, but the option cast allows us to transform it into a type.  Types that have a suffix of "be" indicate a [Big Endian](https://en.wikipedia.org/wiki/Endianness) bit order, those without a "be" suffix use a Little Endian bit order.

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
* IPv4
* IPv6

### Inline filtering

The slice module supports inline filtering which allows for very fast processing of binary data.  Every type does not support every filter operation.  For example attempting to find a subset in a floating point number does not make any sense, nor does applying "less than" to a byte slice.  Below is the complete list of filter operators and a table showing which operators can be applied to which types:

#### Filter Operators

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value
| < | Less Than | Numeric value of field is less than
| <= | Less Than or Equal to | Numeric value of field is less than or equal to
| > | Greater Than | Numeric value of field is greater than
| >= | Greater Than or Equal to | Numeric value of field is greater than or equal to

#### Supported Operators by Type

Type     | == | != | ~ | !~ | < | <= | > | >=
----------|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:
byte     | X | X |   |   | X | X | X | X 
int16    | X | X |   |   | X | X | X | X
int16le  | X | X |   |   | X | X | X | X
int16be  | X | X |   |   | X | X | X | X
uint16   | X | X |   |   | X | X | X | X
uint16le | X | X |   |   | X | X | X | X
uint16be | X | X |   |   | X | X | X | X 
int32    | X | X |   |   | X | X | X | X
int32le  | X | X |   |   | X | X | X | X
int32be  | X | X |   |   | X | X | X | X
uint32   | X | X |   |   | X | X | X | X
uint32le | X | X |   |   | X | X | X | X
uint32be | X | X |   |   | X | X | X | X
int64    | X | X |   |   | X | X | X | X
int64le  | X | X |   |   | X | X | X | X
int64be  | X | X |   |   | X | X | X | X
uint64   | X | X |   |   | X | X | X | X
uint64le | X | X |   |   | X | X | X | X
uint64be | X | X |   |   | X | X | X | X
float32  | X | X |   |   | X | X | X | X
float32le| X | X |   |   | X | X | X | X
float32be| X | X |   |   | X | X | X | X
float64  | X | X |   |   | X | X | X | X
float64le| X | X |   |   | X | X | X | X
float64be| X | X |   |   | X | X | X | X
array    | X | X | X | X |   |   |   |
string   | X | X | X | X |   |   |   |
IPv4     | X | X | X | X |   |   |   |
IPv6     | X | X | X | X |   |   |   |

Note: The `IPv4` and `IPv6` operators expect 4 and 16 byte network encoded values, text encoding of IP addresses will not extract appropriately.
