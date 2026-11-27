# Slice

Slice is a low-level tool that extracts bytes from the entry data or an enumerated value by offset, and optionally casts them to a type.

This search extracts the RPM from canbus messages on a Toyota vehicle by reading the first two bytes of the "data" enumerated value as a 16-bit big-endian integer:

```gravwell
tag=CAN canbus ID==0x2C4 data | slice uint16be(data[0:2]) as RPM | mean RPM | chart mean
```

Each extraction has the form `[start:end]`, `ev[start:end]`, `type([start:end])`, or `type(ev[start:end])`, optionally followed by a filter and `as <name>`. You can give several extractions in one slice invocation.

Ranges work like Go slices:

* `start` is inclusive and `end` is exclusive, so `[0:4]` extracts 4 bytes. Offsets can be decimal or hex (`0x10`).
* An omitted `start` means 0; an omitted `end` (or an `end` of 0) means the end of the data.
* Negative offsets count back from the end of the data, so `[-2:]` is the last two bytes.

Without `as`, an extraction from an enumerated value replaces that enumerated value, and a typed extraction from the entry data is named after its type (for example, `uint16be`).

An entry is dropped if any extraction fails or doesn't pass its filter. An extraction fails when the source enumerated value is missing, `start` is past the end of the data, or there aren't enough bytes for a fixed-size type. Byte array and string extractions that run past the end of the data are cut short instead.

Some examples:

| Command | Description |
|---------|-------------|
| `slice [0:4] as foo` | Extract the first 4 bytes of the entry data into an enumerated value "foo" |
| `slice uint16le(Payload[9:11]) as value` | Pull bytes 9 and 10 from the enumerated value "Payload", parse them as an unsigned 16-bit little-endian integer, and store it as "value" |
| `slice uint16be(Payload[-2:]) as value` | Pull the last two bytes from "Payload", parse them as an unsigned 16-bit big-endian integer, and store it as "value" |
| `slice uint16be(Payload[-4:-2]) as value2` | Pull the two bytes preceding the last two bytes of "Payload", parse them as an unsigned 16-bit big-endian integer, and store as "value2"

## Supported Types

By default, slice extracts a byte array. A type cast converts the bytes to another type. Types with a "be" suffix use [big-endian](https://en.wikipedia.org/wiki/Endianness) byte order; all others are little-endian. The "le" and "be" suffixes can also be written in uppercase (`uint16BE`). For fixed-size types, the range must be exactly the size of the type when both offsets are known at parse time, for example `[0:2]` for `uint16`.

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

## Inline filtering

Slice supports inline filtering for fast processing of binary data. Not every type supports every operator: a subset test on a float makes no sense, nor does "less than" on a byte array. For example, `slice byte(DATA[-1:]) == 0x0a` keeps entries whose last byte is a newline.

### Filter Operators

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

### Supported Operators by Type

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

```{note}
The `IPv4` and `IPv6` types expect 4- and 16-byte network-encoded addresses. They can't parse IP addresses written as text.
```
