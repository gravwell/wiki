## CEF

The CEF or Common Event Format is used by several data providers and many data enrichment systems.  The Gravwell CEF parser is designed to allow for very fast extraction of the common CEF header variables, as well as an arbitrary set of key value pairs.  The loose specifcation that is CEF technically defines a set of known key names, but we have yet to see a data generating product that produces CEF and strictly holds to those set of keys, which is why the Gravwell module can handle any key name.

### Standard CEF Header Key Names

CEF contains a set of standardized header values that should be present in every CEF record.  The header record names are:

* Version
* DeviceVendor
* DeviceProduct
* Signature
* Name
* Severity

The Gravwell CEF parser allows additional flexibility in extracting submember keys that collide with the header names.  By default, if a search specifies the key Version, we will extract the header value Version.  However, if a misbehaving data source is providing CEF records with a key named Version you can still access it by prepending "EXT" to the name.

Extracted header and key values are extracted into enumerated values with the same name as the key or header.  However, using the "as" syntax the extracted values can be renamed to any value.  The gravwell CEF parser is designed with flexibility in mind, and can deal with poorly formed CEF records and records that technically violate the loosely defined CEF spec.

## Supported Options

* `-e`: The “-e” option specifies that the CEF module should operate on an enumerated value.  Operating on enumerated values can be useful when you have extracted a CEF record using upstream modules.  You could e.g. extract CEF records from raw PCAP and pass the records into the CEF module.

## Processing Operators

Each CEF field supports a set of operators that can act as fast filters.  The filters supported by each operator are determined by the data type of the field.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~  | Contains | Field must contain the sub sequence
| !~ | Not contains | Field must NOT contain the sub sequence

When possible, use the CEF inline filters so that the CEF module can zero into the types of records you want immediately, rather than relying on downstream modules.  Using inline filters is not only faster during worst case operation, it also enables you to take advantage of field accelerators when they are enabled.

### Examples

If we wanted to extract the device vendor, product, severity, and msg from the following CEF record and draw a table with only records where the severity was > 7.

An example CEF record:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

The Query:

```
tag=firewall cef DeviceVendor DeviceProduct Severity==7 msg | table DeviceVendor DeviceProduct msg
```

However, if we had the technically invalid CEF record which contained a key value named Version and we wanted that instead of the header Version we can still access that value using the Ext designator.

The poorly formed CEF record:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 Version=11.0 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

The Query:

```
tag=firewall cef DeviceVendor==Citrix DeviceProduct==NetScaler Severity Ext.Version msg ~ Banned | table DeviceVendor DeviceProduct msg Version
```

The query would extract the value "11.0" for version rather than "0" but if you ALSO wanted the header value we can make use of the "as" syntax to pull both the header Version and the key Version.  The "~" inline filter states that we only want records containing the word "Banned" in the message field.

```
tag=firewall cef DeviceVendor DeviceProduct Severity Version as hdrversion Ext.Version msg | eval Severity > 7 | table DeviceVendor DeviceProduct msg Version hdrversion
```
