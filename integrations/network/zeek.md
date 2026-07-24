# Zeek

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [File Follower ingester](/ingesters/file_follow)
         Kit, [Zeek Kit](https://github.com/gravwell/kits/tree/main/zeek)
:::


## Zeek Configuration

The standard method to collect Zeek logs is using the gravwell-file-follower.

File Follower can be installed following instructions in the [file follower](/ingesters/file_follow.md). File follower will need to be configured to point to your current Gravwell environment.

## Gravwell Configuration

### Gravwell Storage Well Configuration

The Zeek data set contains many highly orthogonal data sources, including unique identifiers, GUIDs, floating point timestamps, and IPv6 addresses. To ensure that Gravwell performs well and minimizes memory usage when indexing Zeek data we highly recommend a separate well for the Zeek data with specific indexing options.

Gravwell supports two indexing engines designed to provide different capabilities and tradeoffs. Both engines can perform very well with the Zeek datasets. The bloom engine can provide a balance of good performance and minimal disk usage while the index engine provides precise indexing performance in exchange for greater disk and memory usage. Regardless of the chosen engine, Gravwell recommends that Zeek data be fulltext indexed with the "ignoreFloat" and "ignoreUUID" options. The following well configurations work well with Zeek data:

Create or edit: `/opt/gravwell/etc/gravwell.conf.d/zeek.conf`

**Sample Bloom Engine well config:**
```
[Storage-Well "zeek"]
    Location=/opt/gravwell/storage/zeek
    Tags=zeek*
    Accelerator-Name=fulltext
    Accelerator-Args="-ignoreFloat -ignoreUUID"
    Accelerator-Engine-Override=bloom
```
**Sample Index Engine well config:**
```
[Storage-Well "zeek"]
    Location=/opt/gravwell/storage/zeek
    Tags=zeek*
    Accelerator-Name=fulltext
    Accelerator-Args="-ignoreFloat -ignoreUUID"
    Accelerator-Engine-Override=index
```

### Gravwell Ingester Configuration: File Follower
**Sample File Follower config:**  
Create or edit: `/opt/gravwell/etc/file_follow.conf.d/zeek.conf`
```ini
[Follower "barnyard2"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="barnyard2.log"
    Tag-Name="zeekbarnyard2"

[Follower "conn"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="conn.log"
    Tag-Name="zeekconn"

[Follower "dce_rpc"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="dce_rpc.log"
    Tag-Name="zeekdce_rpc"

[Follower "dhcp"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="dhcp.log"
    Tag-Name="zeekdhcp"

[Follower "dnp3"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="dnp3.log"
    Tag-Name="zeekdnp3"

[Follower "dns"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="dns.log"
    Tag-Name="zeekdns"

[Follower "dpd"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="dpd.log"
    Tag-Name="zeekdpd"

[Follower "files"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="files.log"
    Tag-Name="zeekfiles"

[Follower "ftp"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="ftp.log"
    Tag-Name="zeekftp"

[Follower "http"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="http.log"
    Tag-Name="zeekhttp"

[Follower "intel"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="intel.log"
    Tag-Name="zeekintel"

[Follower "irc"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="irc.log"
    Tag-Name="zeekirc"

[Follower "kerberos"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="kerberos.log"
    Tag-Name="zeekkerberos"

[Follower "known_certs"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="known_certs.log"
    Tag-Name="zeekknown_certs"

[Follower "known_hosts"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="known_hosts.log"
    Tag-Name="zeekknown_hosts"

[Follower "known_modbus"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="known_modbus.log"
    Tag-Name="zeekknown_modbus"

[Follower "known_services"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="known_services.log"
    Tag-Name="zeekknown_services"

[Follower "modbus"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="modbus.log"
    Tag-Name="zeekmodbus"

[Follower "modbus_register_change"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="modbus_register_change.log"
    Tag-Name="zeekmodbus_register_change"

[Follower "mysql"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="mysql.log"
    Tag-Name="zeekmysql"

[Follower "notice_alarm"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="notice_alarm.log"
    Tag-Name="zeeknotice_alarm"

[Follower "notice"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="notice.log"
    Tag-Name="zeeknotice"

[Follower "ntlm"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="ntlm.log"
    Tag-Name="zeekntlm"

[Follower "ocsp"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="ocsp.log"
    Tag-Name="zeekocsp"

[Follower "openflow"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="openflow.log"
    Tag-Name="zeekopenflow"

[Follower "pe"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="pe.log"
    Tag-Name="zeekpe"

[Follower "radius"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="radius.log"
    Tag-Name="zeekradius"

[Follower "rdp"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="rdp.log"
    Tag-Name="zeekrdp"

[Follower "rfb"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="rfb.log"
    Tag-Name="zeekrfb"

[Follower "signatures"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="signatures.log"
    Tag-Name="zeeksignatures"

[Follower "sip"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="sip.log"
    Tag-Name="zeeksip"

[Follower "smb_cmd"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="smb_cmd.log"
    Tag-Name="zeeksmb_cmd"

[Follower "smb_files"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="smb_files.log"
    Tag-Name="zeeksmb_files"

[Follower "smb_mapping"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="smb_mapping.log"
    Tag-Name="zeeksmb_mapping"

[Follower "smtp"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="smtp.log"
    Tag-Name="zeeksmtp"

[Follower "snmp"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="snmp.log"
    Tag-Name="zeeksnmp"

[Follower "socks"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="socks.log"
    Tag-Name="zeeksocks"

[Follower "software"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="software.log"
    Tag-Name="zeeksoftware"

[Follower "ssh"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="ssh.log"
    Tag-Name="zeekssh"

[Follower "ssl"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="ssl.log"
    Tag-Name="zeekssl"

[Follower "sy"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="syslog.log"
    Tag-Name="zeeksyslog"

[Follower "tunnel"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="tunnel.log"
    Tag-Name="zeektunnel"

[Follower "weird"]
    Ignore-Line-Prefix="#"
    Timestamp-Format-Override="UnixMilli"
    Base-Directory="/logs/"
    File-Filter="weird.log"
    Tag-Name="zeekweird"

[Follower "x509"]
    Timestamp-Format-Override="UnixMilli"
    Ignore-Line-Prefix="#"
    Base-Directory="/logs/"
    File-Filter="x509.log"
    Tag-Name="zeekx509"
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_file_follow.service`
```