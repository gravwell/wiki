## Lookup

The Lookup module is used to do data enrichment and translation off of a static data source stored in a Resource.

```
lookup -r <resource name> <enumerated value> <column to match> <column to extract> as <valuename>
```

Note: If you do NOT provide an ```as <valuename>``` addition to the syntax, lookup will REPLACE the enumerated value with the extracted value from the lookup.

Multiple lookup operations can be specified in a single invocation of the lookup module by stringing together additional operations.

### Supported Options
* `-r <arg>`: The "-r" option informs the lookup module which lookup resource should be used to enrich data.
* `-s`: The "-s" option specifies that the lookup modules should require that all specified operations succeeed.
* `-v`: The "-v" flag inverts the flow logic in the lookup module, meaning that successful matches are suppressed and missed matches are passed on.  The "-v" and "-s" flags can be combined to provide basic whiltelisting, passing only values which do not exist in the specified lookup table.

### Setting up a lookupdata resource

Lookup data can be downloaded from compatible render modules (e.g. the table module) and downloaded or stored in a resource for sharing and utilization. Using the menu on the search results page, we can opt to download a table of search results in this format by selecting "LOOKUPDATA".

![Lookup Download](lookup-download.png)

The [table renderer](#!search/table/table.md) also provides a `-save` option, which will automatically save the search result table as a resource for later use by lookup:

```
tag=syslog regex "DHCPACK on (?P<ip>\S+) to (?P<mac>\S+)" | unique ip mac | table -save ip2mac ip mac
```

In the above example, the table renderer automatically creates a resource named 'ip2mac' which contains a mapping of IP addresses to MAC addresses as derived from DHCP logs.

#### CSVs

CSV data can also be used for the lookup module. In order to use a csv file as a resource in the Gravwell lookup search module the CSV must contain unique headers for the columns.

### Example Search

In this example, we have a resource called "macresolution" which was created from the following csv:
```
mac,hostname
mobile-device-1,40:b0:fa:d7:af:01
desktop-1,64:bc:0c:87:bc:71
mobile-device-2,40:b0:fa:d7:ae:02
desktop-2,64:bc:0c:87:9a:11
```

Then we issue a search off of packet data and use the lookup module to enrich our data stream to include hostnames, which in this case we are assigning to the "devicename" enumerated value.

```
tag=pcap packet eth.SrcMAC | count by SrcMAC | lookup -r macresolution SrcMAC mac hostname as devicename |  table SrcMAC devicename count
```

This results in a table containing the enriched data of
```
64:bc:0c:87:bc:71	|	desktop-1       	|	52183
40:b0:fa:d7:ae:02	|	mobile-device-2 	|	21278
64:bc:0c:87:9a:11	|	desktop-2       	|	 2901
40:b0:fa:d7:af:01	|	mobile-device-1 	|	  927
```

#### Example whitelist operation using the lookup table
```
tag=pcap packet eth.SrcMAC | count by SrcMAC | lookup -v -s -r macresolution SrcMAC mac hostname |  table SrcMAC count
```

This results in a table containing any mac addresses which were NOT in the look up list.  System administrators can use the "-v" and "-s" flag to provide basic white listing and identification of new devices on a network or new logs in an event stream.
```
64:bc:0c:87:bc:60	|	24382
40:b0:fa:d7:ae:13	|	93485
64:bc:0c:87:9a:02	|	11239
40:b0:fa:d7:af:fe	|	   21
```
