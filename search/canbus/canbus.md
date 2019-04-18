## Canbus

The canbus module extracts fields from CAN messages (i.e. vehicle data).

| mod | Field | Operators | Example
|-----|-------|-----------|----------
| canbus | ID | == != < > <= >= | canbus ID==0x341
| canbus | EID | == != < > <= >= | canbus EID==0x123456
| canbus | RTR | == != | canbus RTR==true
| canbus | Data | ~ !~ | canbus Data

### Example Search

The following search will count by canbus packet IDs and display a table with the most frequent IDs.

```
tag=vehicles canbus ID | count by ID | sort by count desc | table ID count
```

The following search extracts packets specifying throttle data and plots the mean position of the throttle.
```
tag=vehicles canbus ID==0x123 Data | slice uint16be(Data[2:4]) as throttle | mean throttle | chart mean
```
