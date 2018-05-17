## Subnet

Subnet is designed to allow for filtering and/or extracting based on subnet.  This query is useful for looking at values that only fall within a specific subnet, or for classifying attackers based on points of origin.  By default the subnet module assumes IPv4, but fully supports IPv6 via the `-6` flag.  If no target flag is specified, the subnet module populates the `subnet` enumerated value.

### Supported Options

* `-t <arg>`: Write the extracted subnet to an alternate enumerated value name.
* `-4`: Look for IPv4 subnets and IPs
* `-6`: Look for IPv6 subnets and IPs

### Example Usage

Charting failed SSH login attempts by origin subnet

```
tag=syslog grep sshd | grep "Failed password for" | regex "\sfrom\s(?P<ip>\S+)\s" | subnet ip /16 | count by subnet | chart count by subnet limit 64
```

Filtering failed SSH login attempts to only those from a specific subnet

```
tag=syslog grep sshd | grep "Failed password for" | regex "\sfrom\s(?P<ip>\S+)\s" | subnet -t atackersub ip /16 | grep -e attackersub “34.22.1.0” | count by ip | sort by count desc | table ip count
```