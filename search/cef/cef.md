# CEF

The `cef` module extracts header fields and extension key/value pairs from Common Event Format (CEF) records. The CEF specification defines a set of standard extension keys, but in practice few products stick to them, so the module accepts any key name. It also tolerates poorly formed records that bend the spec.

The record does not need to start at the beginning of the entry: the module looks for the `CEF:` marker, so records with a syslog header in front of them parse normally.

## Standard CEF Header Key Names

Every CEF record starts with these header fields:

* Version
* DeviceVendor
* DeviceProduct
* DeviceVersion
* Signature
* Name
* Severity

Header names are case-sensitive. Any other name is treated as an extension key.

If an extension key has the same name as a header field, the header field wins. To get the extension instead, prefix the key with `Ext.`: `Ext.Version` extracts the extension key `Version`, not the header `Version`.

Each value is extracted into an enumerated value with the same name as the header field or key (without any `Ext.` prefix). Use `as` to rename it, for example `src as source`.

## Supported Options

* `-e <arg>`: Operate on the enumerated value `<arg>` instead of the entry data. This is useful when an upstream module has already extracted a CEF record, for example from raw PCAP. Entries without the enumerated value are dropped.
* `-s`: Strict mode. Drop the entry if it is not a valid CEF record or if any requested extension key is missing.

## Processing Operators

Each field supports inline filters. `Version` is an integer and supports `==`, `!=`, `<`, `>`, `<=`, and `>=`. All other header fields and extension keys are strings and support the following:

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~  | Contains | Field must contain the sub sequence
| !~ | Not contains | Field must NOT contain the sub sequence

An entry that fails a filter is dropped. If an extension key is missing, `==` and `~` filters drop the entry, while `!=` and `!~` filters pass it.

Use inline filters where you can rather than filtering in downstream modules. They are faster, and `==` filters can take advantage of accelerators when they are enabled.

## Examples

The following example extracts the device vendor, product, severity, and `msg` from a CEF record and draws a table of only the records with severity 6.

An example CEF record:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

The Query:

```gravwell
tag=firewall cef DeviceVendor DeviceProduct Severity==6 msg | table DeviceVendor DeviceProduct msg
```

The following technically invalid record has an extension key named `Version`. To extract it instead of the header `Version`, use the `Ext.` prefix.

The poorly formed CEF record:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 Version=11.0 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

The Query:

```gravwell
tag=firewall cef DeviceVendor==Citrix DeviceProduct==NetScaler Severity Ext.Version msg ~ Banned | table DeviceVendor DeviceProduct msg Version
```

This query extracts "11.0" into `Version` rather than the header value "0". The `~` filter keeps only records whose `msg` contains "Banned".

To extract both, rename the header value with `as`:

```gravwell
tag=firewall cef DeviceVendor DeviceProduct Severity Version as hdrversion Ext.Version msg | table DeviceVendor DeviceProduct msg Version hdrversion
```
