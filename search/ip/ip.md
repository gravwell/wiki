# IP

The ip module converts enumerated values to the IP type and can filter on them. For example, you can extract a string containing an IP address from a JSON structure, convert it to an IP, and check whether it is in a given subnet.

By default, every filter must pass for an entry to continue. Entries where a listed enumerated value is missing or does not hold a valid IP address are dropped.

(search_ip_supported_options)=
## Supported Options

* `-or`: pass the entry if any filter succeeds, instead of requiring all of them.
* `-categorize`: label each IP as `PRIVATE`, `PUBLIC`, or `MULTICAST`, using the `PRIVATE` and `MULTICAST` lists below. The label goes in an enumerated value named after the output with `_category` appended: `ip -categorize srcIP` creates `srcIP_category`. An address in both lists, such as `224.0.0.1`, is labeled `PRIVATE`.

## Processing Operators

Enumerated values passed to the ip module can be compared against IP addresses or subnets using the following operators.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | IP must be equal to the given address
| != | Not equal | IP must not be equal to the given address
| ~ | Subset | IP must be a member of the given subnet
| !~ | Not subset | IP must not be a member of the given subnet

You can apply multiple filters to the same EV, even without the `-or` flag. For example, to keep only entries whose `ipaddr` is in neither `10.0.0.0/24` nor `192.168.0.0/24`:

```gravwell
tag=json json ipaddr | ip ipaddr !~ 10.0.0.0/24 ipaddr !~ 192.168.0.0/24
```

The `~` and `!~` operators accept two keywords in place of a subnet. `PRIVATE` matches the standard private, loopback, and link-local networks:

* 10.0.0.0/8
* 172.16.0.0/12
* 192.168.0.0/16
* 127.0.0.0/8
* 224.0.0.0/24
* 169.254.0.0/16
* fd00::/8
* fe80::/10
* ::1/128

`MULTICAST` matches the multicast networks:

* 224.0.0.0/4
* ff00::/8

## Examples

### Convert a string to an IP

Assuming JSON-formatted entries containing an `ipaddr` field, extract that field and convert it to an IP address:

```gravwell
tag=json json ipaddr | ip ipaddr
```

To keep the original string and write the IP to a new enumerated value, use `as`:

```gravwell
tag=json json ipaddr | ip ipaddr as IP
```

### Filter by address or subnet

Assuming CSV data where the third field is the source IP of a connection, keep only connections from 192.168.1.5:

```gravwell
tag=csv csv [2] as srcip | ip srcip==192.168.1.5
```

Drop connections that originate in 192.168.0.0/16:

```gravwell
tag=csv csv [2] as srcip | ip srcip !~ 192.168.0.0/16
```

### Use the PRIVATE keyword

Keep only connections that originate from private networks:

```gravwell
tag=csv csv [2] as srcip | ip srcip ~ PRIVATE
```

### Categorize IP addresses

Label each source IP with its network category:

```gravwell
tag=csv csv [2] as srcip | ip -categorize srcip | table srcip srcip_category
```

The table has two columns: `srcip` and `srcip_category`, which holds `PRIVATE`, `PUBLIC`, or `MULTICAST`.
