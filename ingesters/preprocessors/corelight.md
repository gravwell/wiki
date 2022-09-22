## Corelight JSON to TSV Preprocessor

Corelight is the commercial version of the highly popular network monitoring tool Zeek, Gravwell provides a robust Zeek kit which is designed to provide fast and flexible access to the TSV formated Zeek data.  Rather than replicate the entire Zeek kit for Corelight JSON data which would require additional storage and perform worse than the standard TSV formatted data which comes from Zeek.  This preprocessor is designed to recieve entries from Corelight via their JSON over TCP export option and translate them into Zeek standard TSV data.  As a result you can use the Gravwell Zeek kit and integrations with a commercial Corelight deployment.

The Corelight JSON to TSV preprocessor type is "corelight".

### Supported Options

* Prefix (string, optional): This parameter specifies the tag prefix attached to entries.  The default is "zeek".
* Custom-Format (string, optional): This parameter allows for specifying custom path and header values.  A Custom-Format can be a wholly new format or override an existing format.

Multiple `Custom-Format` options can be specified in a single preprocessor block, this allows for translating output from Corelight plugins or overriding an existing default handler.  Formats are specified as: `<tag suffix>:<TSV headers>`.

For example, here is a format specification which looks for the S7Comm Zeek plugin:

```
s7comm:ts,uid,id,pdu_type,rosctr,parameter,item_count,data_info
```

This specification would output a TSV log on the `zeeks7comm` tag with 8 fields.

### Example Data Translations
Example input conn.log JSON data:
```
{
  "_path": "conn",
  "_system_name": "ds61",
  "_write_ts": "2020-08-16T06:26:04.077276Z",
  "_node": "worker-01",
  "ts": "2020-08-16T06:26:03.553287Z",
  "uid": "CMdzit1AMNsmfAIiQc",
  "id.orig_h": "192.168.4.76",
  "id.orig_p": 36844,
  "id.resp_h": "192.168.4.1",
  "id.resp_p": 53,
  "proto": "udp",
  "service": "dns",
  "duration": 0.06685185432434082,
  "orig_bytes": 62,
  "resp_bytes": 141,
  "conn_state": "SF",
  "missed_bytes": 0,
  "history": "Dd",
  "orig_pkts": 2,
  "orig_ip_bytes": 118,
  "resp_pkts": 2,
  "resp_ip_bytes": 197
}
```

Resulting TSV output:
```
1597559163.553287    CMdzit1AMNsmfAIiQc      192.168.4.76    36844   192.168.4.1     53      udp     dns     0.06685 62      141     SF      -       -       0       Dd      2       118     2       197     -       -
```

#### Example: Default Handlers

Here is a basic [Simple Relay](/#!ingesters/simple_relay.md) example which listens for JSON over TCP on port 7890 and translates the JSON payloads to TSV data with no custom handlers:

```
[Listener "corelight"]
	Bind-String="0.0.0.0:7890"
	Preprocessor=corelight

[Preprocessor "corelight"]
	Type=corelight

```

#### Example: Custom Handlers and Prefix

This example shows how to set a custom prefix and inject a custom handler, using the Prefix override we are setting all the routed tags to prepend `core` instead of `zeek`:

```
[Listener "corelight"]
	Bind-String="0.0.0.0:7890"
	Preprocessor=corelight

[Preprocessor "corelight"]
	Type=corelight
	Prefix="core"
	Custom-Format="s7comm:ts,uid,id,pdu_type,rosctr,parameter,item_count,data_info"

```

#### Example: Override Existing Handler

This example shows that we can inject a custom handler as well as override an existing default handler:

```
[Listener "corelight"]
	Bind-String="0.0.0.0:7890"
	Preprocessor=corelight

[Preprocessor "corelight"]
	Type=corelight
	Custom-Format="s7comm:ts,uid,id,pdu_type,rosctr,parameter,item_count,data_info"
	Custom-Format="conn:ts,uid,id.orig_h,id.orig_p,id.resp_h,id.resp_p,proto,service,duration"

```

### Default Handlers

The following handlers are included by default:

- conn
- dns
- dhcp
- ssh
- ftp
- http
- files
- ssl
- x509
- smtp
- pe
- ntp
- notice
- weird
- dpd
- irc
- rdp
- sip
- snmp
- tunnel
- socks
- software
- syslog
- rfb
- radius
- intel
- kerberos
- mysql
- modbus
- signature
- smb_mapping
- smb_files
- zeekdnp3

