# Canbus

The canbus module parses CAN bus frames (for example, vehicle data) and extracts their fields. With no arguments it extracts every field. Name one or more fields to extract only those, optionally with a filter. Entries that are not valid CAN frames are dropped.

| mod | Field | Operators | Example
|-----|-------|-----------|----------
| canbus | ID | == != < > <= >= | canbus ID==0x341
| canbus | EID | == != < > <= >= | canbus EID==0x123456
| canbus | IDE | == != | canbus IDE==true
| canbus | RTR | == != | canbus RTR==true
| canbus | data | == != ~ !~ | canbus data

`ID` is the standard identifier and `EID` the extended identifier. `ID` filters take values up to 11 bits and `EID` filters up to 29 bits. A query can use `ID` or `EID`, not both. `IDE` is true for frames with an extended identifier, and `RTR` is true for remote transmission requests. `data` is the frame payload.

## Example Search

The following search counts frames by ID and shows the most frequent IDs.

```gravwell
tag=vehicles canbus | count by ID | sort by count desc | table ID count
```

The following search extracts throttle messages and charts the mean throttle position. Each make and model may use different message IDs and data formats for throttle.

```gravwell
tag=vehicles canbus ID==0x123 data | slice uint16be(data[2:4]) as throttle | mean throttle | chart mean
```
