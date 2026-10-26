# Cisco ISE Preprocessor

The Cisco ISE preprocessor is designed to parse and accommodate the format and transport of Cisco ISE logs.  See the [Cisco Introduction to ISE Syslogs](https://www.cisco.com/c/en/us/td/docs/security/ise/syslog/Cisco_ISE_Syslogs/m_IntrotoSyslogs.pdf) for more information.

The Cisco ISE preprocessor is named `cisco_ise` and supports the ability to reassemble multipart messages, reformat the messages into a format more appropriate for Gravwell and modern syslog systems, filter unwanted message pairs, and remove redundant message headers.

## Attribute Filtering and Formatting

The Cisco ISE logging system is designed to split a single message across multiple syslog messages.  Gravwell will accept messages that far exceed the maximum message size of syslog, however if you are supporting multiple targets for Cisco ISE messages it may be necessary to enable multipart messages.  Disabling multipart messages in your Cisco device and letting Gravwell handle large payloads will be far more efficient.

## Supported Options

* `Drop-Misses` (boolean, optional): If set to true, the preprocessor will drop entries for which it was unable to extract a valid ISE message. By default, these entries are passed along.
* `Enable-Multipart-Reassembly` (boolean, optional): If set to true the preprocessor will attempt to reassemble messages that contain a Cisco remote message header.
* `Max-Multipart-Buffer` (uint64, optional): Specifies a maximum in-memory buffer to use when reassembling multipart messages. When the buffer is exceeded, the oldest partially-reassembled message will be sent to Gravwell.  The default buffer size is 8MB.
* `Max-Multipart-Latency` (string, optional): Specifies a maximum time duration that a partially reassembled multipart message will be held before it is sent.  Time spans should be specified in `ms` and `s` values.
* `Output-Format` (string, optional): Specifies the output format for sending ISE messages. The default format is `raw`, other options are `json` and `cef`.
* `Attribute-Drop-Filter` (string array, optional): Specifies globbing patterns that can be used to match against attributes within a message which will be removed from the output.  The arguments must be [Unix Glob patterns](https://en.wikipedia.org/wiki/Glob_(programming)); many patterns can be specified.  Attribute drop filters are not compatible with the `raw` output format.
* `Attribute-Strip-Header` (boolean, optional): Specifies that attributes with nested names and/or type information should have the header values stripped.  This is extremely useful for cleaning up poorly-formed attribute values.
* `Timestamp-Max-Past-Delta` (string, optional): The maximum age of a message header timestamp, as a duration such as `720h`. Timestamps older than this are ignored. Defaults to `8760h` (one year).
* `Timestamp-Max-Future-Delta` (string, optional): How far into the future a message header timestamp may be, as a duration such as `1h`. Timestamps further ahead than this are ignored. Defaults to `24h`.

## Timestamp Bounds

When the output format is `json` or `cef`, the preprocessor sets the entry timestamp from the ISE message header. It only applies that timestamp if it falls within a window around the current time, set by `Timestamp-Max-Past-Delta` and `Timestamp-Max-Future-Delta`. If the header timestamp falls outside the window, the message is still reformatted but the entry timestamp is left unchanged. Timestamps before 1970-01-01 are always ignored, whatever the window. Both durations must be non-zero; negative values are treated as positive.

The `raw` output format does not change the entry timestamp, so these options have no effect with it.

```{warning}
The default window rejects timestamps more than a year old. If you replay archives or backfill older ISE logs through this preprocessor, set `Timestamp-Max-Past-Delta` large enough to cover them, or those entries will keep their original timestamp.
```

## Example Configuration

The following `cisco_ise` preprocessor configuration is designed to re-assemble multipart messages, remove unwanted `Step` attributes, and reform the output messages in the CEF format.  It also strips the cisco attribute headers.

```
[preprocessor "iseCEF"]
    Type=cisco_ise
    Drop-Misses=true #if its malformed just drop it
    Enable-Multipart-Reassembly=true
    Attribute-Drop-Filter="Step*"
    Attribute-Strip-Header=true
    Output-Format=cef
```

An example output message using this configuration is:

```
CEF:0|CISCO|ISE_DEVICE|||Passed-Authentication|NOTICE| sequence=1 ode=5200 class=Passed-Authentication text=Authentication succeeded ConfigVersionId=44 DeviceIPAddress=8.8.8.8 DestinationIPAddress=1.2.3.4 DestinationPort=1645 UserName=user@company.com Protocol=Radius RequestLatency=10301 audit-session-id=0a700e191cff70005fbbf63f
```

The following `cisco_ise` preprocessor configuration achieves a similar result, but it contains two attribute filters and re-formats the output message into JSON:

```
[preprocessor "iseCEF"]
    Type=cisco_ise
    Drop-Misses=true #if its malformed just drop it
    Enable-Multipart-Reassembly=true
    Attribute-Drop-Filter="Step*"
    Attribute-Strip-Header=true
    Output-Format=json
```

An example output message using this configuration is:

```
{
  "TS":"2020-11-23T12:50:01.926-05:00",
  "Sequence":1,
  "ODE":"5200",
  "Severity":"NOTICE",
  "Class":"Passed-Authentication",
   "Text":"Authentication succeeded",
   "Attributes":{
     "AcsSessionID":"ISE_DEVICE/384429556/212087299",
     "AuthenticationIdentityStore":"AzureBackup",
     "AuthenticationMethod":"PAP_ASCII",
     "AuthenticationStatus":"AuthenticationPassed",
     "audit-session-id":"0a700e191cff70005fbbf63f",
     "device-mac":"00-0c-29-74-9d-e8",
     "device-platform":"win",
     "device-platform-version":"10.0.17134",
  }
}
```
