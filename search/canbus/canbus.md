## Canbus

The canbus module extracts fields from CAN messages (i.e. vehicle data). These fields are automatically extracted with the invocation of the canbus module.

| mod | Field | Operators | Example
|-----|-------|-----------|----------
| canbus | ID | == != < > <= >= | canbus ID==0x341
| canbus | EID | == != < > <= >= | canbus EID==0x123456
| canbus | RTR | == != | canbus RTR==true
| canbus | Data | ~ !~ | canbus Data

### Example Search

The following search will count by canbus packet IDs and display a table with the most frequent IDs.

```
tag=vehicles canbus | count by ID | sort by count desc | table ID count
```

The following search extracts messages specifying throttle data and plots the mean position of the throttle. Note that each make/model may use different message IDs and data formats for throttle.
```
tag=vehicles canbus ID==0x123 Data | slice uint16be(Data[2:4]) as throttle | mean throttle | chart mean
```
