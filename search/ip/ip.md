# IP

The ip module can convert enumerated values to the IP type and optionally perform filtering. This allows the user to, for instance, extract a string containing an IP address from a JSON structure, then use the ip module to convert that string to an IP address and check if it is in a certain subnet.

## Supported Options

* `-or`: The "-or" flag specifies that the ip module should allow an entry to continue down the pipeline if ANY of the filters are successful.
* `-categorize`: If the "-categorize" flag is set, the module will attempt to categorize each enumerated value as "PRIVATE", "PUBLIC", or "MULTICAST". Running `ip -categorize srcIP` will create an enumerated value named "srcIP_category" containing the category string.

## Processing Operators

Enumerated values passed to the ip module can be compared against IP addresses or subnets using the following operators.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | IP must be equal to the given address
| != | Not equal | IP must not be equal to the given address
| ~ | Subset | IP must be a member of the given subnet
| !~ | Not subset | IP must not be a member of the given subnet

The ip module defines the keyword PRIVATE to match any of the standard private networks:

* 10.0.0.0/8
* 172.16.0.0/12
* 192.168.0.0/16
* 127.0.0.0/8
* 224.0.0.0/24
* 169.254.0.0/16
* fd00::/8
* fe80::/10

## Examples

### Convert a string to an IP

Assuming JSON-formatted entries containing an 'ipaddr' field, extract that field and convert it to an IP address for later use:

```
tag=json json ipaddr | ip ipaddr
```

The resulting IP enumerated value can also be assigned to a different enumerated value name rather than overwriting the original:

```
tag=json json ipaddr | ip ipaddr as IP
```

### Filter by address or subnet

Assuming CSV-formatted data in which the 3rd field describes the source IP address of a connection, we can drop all connections not originating from 192.168.1.5:

```
tag=csv csv [2] as srcip | ip srcip==192.168.1.5
```

We can also eliminate any connections which originated in the local subnet:

```
tag=csv csv [2] as srcip | ip srcip !~ 192.168.0.0/16
```

### Use the PRIVATE keyword

Assuming CSV-formatted data in which the 3rd field describes the source IP address of a connection, we can use the ip module to keep only those entries originating from private networks:

```
tag=csv csv [2] as srcip | ip srcip ~ PRIVATE
```

### Categorize IP addresses

Assuming CSV-formatted data in which the 3rd field describes the source IP address of a connection, we can use the ip module to assign a network category for each IP address:

```
tag=csv csv [2] as srcip | ip -categorize srcip | table srcip srcip_category
```

The resulting table will contain two columns: `srcip`, and `srcip_category` which will contain one of "PRIVATE", "PUBLIC", or "MULTICAST".