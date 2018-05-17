## Hexlify

The hexlify module is used to encode a data into ASCII hex representations.  The module can be useful when tackling unknown data types and learning how to process binary data.  For example, one might encode an unknown enumerated value extracted from canbus data.  Most manufacturers do not publish canbus specs, but by extracting from IDs and encoding it in hex it can assist in identifying values that are changing in predictable patterns, helping to identify parameters.  This is exactly how the Gravwell team derived the PDUs for gas level, speed, and throttle position of a RAM 1500 truck without having access to canbus IDs from Fiat Chrysler of America.

### Supported Options

* `-d`: Decode ASCII hex into an integer, rather than encoding an int as ASCII hex.
* `-t <name>`: Assign encoded or decoded results into an enumerated value with the given name.

### Example Search

```
tag=CAN canbus ID Data | hexlify Data | table ID Data
```