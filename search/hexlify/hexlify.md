## Hexlify

The hexlify module is used to encode a data into ASCII hex representations.  The module can be useful when tackling unknown data types and learning how to process binary data.  For example, one might encode an unknown enumerated value extracted from canbus data.  Most manufacturers do not publish canbus specs, but by extracting from IDs and encoding it in hex it can assist in identifying values that are changing in predictable patterns, helping to identify parameters.  This is exactly how the Gravwell team derived the PDUs for gas level, speed, and throttle position of a RAM 1500 truck without having access to canbus IDs from Fiat Chrysler of America.

### Supported Options

* `-d`: Decode ASCII hex into an integer, rather than encoding an int as ASCII hex.


### Example Search to hexlify all data

```
tag=stuff hexlify
```

### Example Search to hexlify a single enumerated value	

```
tag=CAN canbus ID Data | hexlify Data | table ID Data
```

### Example Search to hexlify all data and assign to a new name

```
tag=stuff hexlify DATA as hexdata | table DATA hexdata
```

### Example Search to hexlify a few enumerated values with reassignment

```
tag=CAN canbus ID Data | hexlify ID as hexid Data as hexdata | table ID hexid DATA hexdata
```

### Example decoding hex data

```
tag=apache json val | hexlify -d val as decodedval | table val decodedval
```
