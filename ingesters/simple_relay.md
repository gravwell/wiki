---
myst:
  substitutions:
    package: "gravwell-simple-relay"
    standalone: "gravwell_simple_relay"
    dockername: "simple_relay"
---
# Simple Relay

Simple Relay is the go-to ingester for text based data sources that can be delivered over plaintext TCP, encrypted TCP, or plaintext UDP network connections via either IPv4 or IPv6.

Some common use cases for Simple Relay are:

* Remote syslog collection
* Devops log collection over a network
* Bro sensor log collection
* Simple integration with any text source capable of delivering over a network

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

The Simple Relay ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters Simple Relay supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/simple_relay.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/simple_relay.conf.d`).

An example configuration for the Simple Relay ingester, configured to listen on several ports and apply a unique tag to each is as follows:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Cleartext-Backend-target=127.0.0.1:4023 #example of a cleartext connection
Cleartext-Backend-target=127.1.0.1:4023 #example of second cleartext connection
Encrypted-Backend-target=127.1.1.1:4024 #example of encrypted connection
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection
Ingest-Cache-Path=/opt/gravwell/cache/simple_relay.cache #local storage cache when uplinks fail
Max-Ingest-Cache=1024 #Number of MB to store, localcache will only store 1GB before stopping
Log-Level=INFO
Log-File=/opt/gravwell/log/simple_relay.log

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Listener "default"]
	Bind-String="0.0.0.0:7777" #bind to all interfaces, with TCP implied
	#Lack of "Reader-Type" implies line break delimited logs
	#Lack of "Tag-Name" implies the "default" tag
	#Assume-Local-Timezone=false #Default for assume localtime is false
	#Source-Override="DEAD::BEEF" #override the source for just this listener

[Listener "syslogtcp"]
	Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
	Reader-Type=rfc5424
	Tag-Name=syslog
	Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
	Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry

[Listener "syslogudp"]
	Bind-String="udp://0.0.0.0:514" #standard UDP based RFC5424 syslog
	Reader-Type=rfc5424
	Tag-Name=syslog
	Timezone-Override="America/Chicago"
	Keep-Priority=true	# leave the <nnn> priority tag at the start of each syslog entry
```

```{note}
The `Keep-Priority` field is necessary if you plan to analyze syslog entries with the [syslog search module](/search/syslog/syslog).
```

## Listeners

Each listener contains a set of universal configuration values, regardless of whether the listener is a line, RFC5424, RFC6587, regex, or JSON reader.

Listeners support several configuration parameters for specifying protocols, listening interfaces and ports, and fine tuning ingest behavior.

| Parameter     | Type    | Default Value     | Description |
|---------------|---------|-------------------|-------------|
| Tag-Name      | string  |                   | Tag to be assigned to data ingested on this listener |
| Bind-String   | string  |                   | Define optional interface and port to bind to |
| Ignore-Timestamps | Boolean | false   | Do not attempt to resolve timestamps from entries, instead use the time of collection |
| Assume-Local-Timezone | Boolean | false | Assume timestamps that do not contain a timezone are located in the local timezone |
| Timezone-Override | string |  | Override the timezone for timestamps that do not contain timezone data |
| Source-Override |   string |  | Override the source IP assigned to entries ingested on the listener |
| Timestamp-Format-Override | string | | Override the timestamp format used to derive timestamps from ingested data |
| Cert-File | string | | Path to an X509 public certificate file for use in TLS listeners |
| Key-File  | string | | Path to an X509 private key file for use in TLS listeners |
| Preprocessor | string | | Name of a preprocessor to apply to ingested data, many Preprocessor parameters can be applied |


### Bind-String

The Bind-String parameter controls which interface and port the listener will bind to.  The listener can bind to plaintext/encrypted TCP or plaintext UDP ports, specific addresses, and specific ports.  IPv4 and IPv6 are supported.

```
#bind to all interfaces on TCP port 7777
Bind-String=0.0.0.0:7777

#bind to all interfaces on UDP port 514
Bind-String=udp://0.0.0.0:514

#bind to port 1234 on link local IPv6 address on interface p1p
Bind-String=[fe80::4ecc:6aff:fef9:48a3%p1p1]:1234

#bind to IPv6 globally routable address on TCP port 901
Bind-String=[2600:1f18:63ef:e802:355f:aede:dbba:2c03]:901

#listen for TLS connections on port 9999
Bind-String=tls://0.0.0.0:9999
```

### Cert-File

If using a TLS `Bind-String`, you must also specify `Cert-File`. The value should be the path to a file containing a valid TLS certificate:

```
Cert-File=/opt/gravwell/etc/cert.pem
``` 

### Key-File

If using a TLS `Bind-String`, you must also specify `Key-File`. The value should be the path to a file containing a valid TLS key:

```
Key-File=/opt/gravwell/etc/key.pem
``` 

### Listener Reader Types and configurations

Simple relay supports the following types of basic readers which are useful in different contexts

* line reader
* RFC5424
* RFC6587
* json
* regex

The basic listeners are specified via the "Listener" block and support line delimited and RFC5424 reader types.  Each Listener must have a unique name and a unique bind port (two different listeners cannot bind to the same protocol, address, and port).  The reader type for a basic listener is controlled by the "Reader-Type" parameter.  If no "Reader-Type" is specified, the line reader type is assumed.

Basic Listeners also require that each listener designate the tag the listener will apply to all incoming data via the "Tag-Name" parameter.  If the "Tag-Name" parameter is omitted, the "default" tag is applied.  The most basic listener named "test" which expects line broken data on TCP port 5555 and applies the tag "testing" would have the following configuration specification:

```
[Listener "test"]
	Bind-String=0.0.0.0:5555
	Tag-Name=testing
```

## Line Reader Listener

The line reader listener is designed to read newline broken data streams from either a TCP or UDP stream.  Applications which can deliver simple line broken data over a network can utilize this type of reader to very simply and easily integrate with Gravwell.  The Line Reader listener can also be used for simple log file delivery by simply sending log files to the listening port.

For example, an existing log file can be imported into Gravwell using netcat and Simple Relay:
```
nc -q 1 10.0.0.1 7777 < /var/log/syslog
```

### Example Line Reader Listener

The most basic Listener requires only one the "Bind-String" argument which tells the listener what port to listen on:

```
[Listener "default"]
	Bind-String="0.0.0.0:7777" #bind to all interfaces, with TCP implied
```

## RFC5424 Listener

A listener designed to accept structured syslog messages based on either [RFC5424](https://www.rfc-editor.org/rfc/rfc5424) or [RFC3164](https://www.rfc-editor.org/rfc/rfc3164) enables Simple Relay to act as a syslog aggregation point.  To enable a listener that expects syslog messages using a reliable TCP connection on port 601 set the "Reader-Type" to "RFC5424".

Additional RFC5424 listener configuration parameters:

| Parameter     | Type    | Default Value     | Description |
|---------------|---------|-------------------|-------------|
| Drop-Priority | Boolean | false             | Removes the `<nnn>` priority header in RFC5424 messages |

### RFC5424 Examples

```
[Listener "syslog"]
	Bind-String=0.0.0.0:601
	Reader-Type=RFC5424
```

To accept syslog messages over stateless UDP via port 514 the listener would look like the following:

```
[Listener "syslog"]
	Bind-String=udp://0.0.0.0:514
	Reader-Type=RFC5424
```

RFC5424 reader types also support a parameter named "Keep-Priority" which is set to true by default.  A typical syslog message is prepended by a priority identifier, however some users may wish to discard the priority from stored messages.  This is accomplished by added "Keep-Priority=false" to an RFC5424 based listener.  Line based listeners ignore the "Keep-Priority" parameter.

An example syslog message with a priority attached:

```
<30>Sep 11 17:04:14 router dhcpd[9987]: DHCPREQUEST for 10.10.10.82 from e8:c7:4f:04:e1:af (Chromecast) via insecure
```

An example listener specification which removes the priority tag from entries:

```
[Listener "syslog"]
	Bind-String=udp://0.0.0.0:514
	Reader-Type=RFC5424
	Drop-Priority=true
```

```{note}
The priority portion of a syslog message is codified in the RFC specification.  Removing the priority means that the Gravwell [syslog](/search/syslog/syslog) search module will be unable to properly parse the values, the syslog search module is dramatically faster than attempting to hand parse syslog messages with regular expressions.  However, some systems use an RFC5424-like header to send non-syslog data (such as the Fortinet products), in which case the data is not syslog at all and dropping the priority is appropriate.
```

## RFC6587 Listener

A listener designed to perform octet counting on TCP streams of data with optional framing, the [RFC6587](https://www.rfc-editor.org/rfc/rfc6587) listener is designed to parse incoming messages that contain an octet (byte) count header which defines the length of the message.  Properly formatted RFC6587 messages contain a header with an ASCII base 10 number indicating the length of the message, each message is then terminated with an optional newline (0xa), null (0x0), or newline and carriage return (0xa 0xd).  The RFC6587 reader will process the framed messages and remove the octet count header and any optional framing.

Additional RFC6587 listener configuration parameters:

| Parameter     | Type    | Default Value     | Description |
|---------------|---------|-------------------|-------------|
| Drop-Priority | Boolean | false             | Removes the `<nnn>` priority header in RFC5424-like messages |


A common source of RFC6587 data is the Fortinet series of firewalls and switches when transmitting logs using the "reliable" mode, an example RFC6587 message is shown below:

```
464 <185>date=2022-09-28 time=08:49:59 devname="fortigate" devid="FGT60EABCDEF012" eventtime=1664380199448569680 tz="-0700" logid="0100032002" type="event" subtype="system" level="alert" vd="root" logdesc="Admin login failed" sn="0" user="admin" ui="ssh(192.168.1.100)" method="ssh" srcip=192.168.1.100 dstip=192.168.1.99 action="login" status="failed" reason="ssh_key_invalid" msg="Administrator admin login failed from ssh(192.168.1.100) because of invalid ssh key"
```

```{note}
RFC6587 listeners are not compatible with UDP transports.
```

### RFC6587 Examples

A basic listener which expects RFC6587 data on TCP port 601.

```
[Listener "switch logs"]
	Bind-String=0.0.0.0:601
	Reader-Type=rfc6587
```

An example listener designed to accept "reliable" mode data from a Fortinet device and strip the priority header to create clean key/value data.

```
[Listener "fortinet"]
	Bind-String=tcp://0.0.0.0:601
	Reader-Type=rfc6587
	Drop-Priority=true
```

The previously shown example would be ingested as:

```
date=2022-09-28 time=08:49:59 devname="fortigate" devid="FGT60EABCDEF012" eventtime=1664380199448569680 tz="-0700" logid="0100032002" type="event" subtype="system" level="alert" vd="root" logdesc="Admin login failed" sn="0" user="admin" ui="ssh(192.168.1.100)" method="ssh" srcip=192.168.1.100 dstip=192.168.1.99 action="login" status="failed" reason="ssh_key_invalid" msg="Administrator admin login failed from ssh(192.168.1.100) because of invalid ssh key"
```

## Regex Listeners

The regex listener type is a very flexible listener which can split entries based on arbitrary regular expressions. This can be useful when log sources do not adhere to standard formats like RFC5424 or RFC6587.


Additional Regex listener configuration parameters:

| Parameter     | Type    | Default Value     | Description |
|---------------|---------|-------------------|-------------|
| Regex | string | | Regular expression used to detect the start of an entry |
| Trim-Whitespace | Boolean | false | remove extra whitespace around entries after an extraction takes place |
| Max-Buffer | integer | 8MB | The maximum amount of data to be read before an entry is forced out |


For instance, there may be existing infrastructure which forwards Windows XML event logs over a plain TCP connection:

```
<Event xmlns="http://schemas.microsoft.com/win/2004/08/events/event">
  <System>
    <Provider Name="Microsoft-Windows-Security-Auditing" Guid="{543496D5-5478-49A4-A5BA-3E3B0428E31D}"/>
    <EventID>4689</EventID>
    <Version>0</Version>
    <Level>0</Level>
    <Task>13313</Task>
    <Opcode>0</Opcode>
    <Keywords>0x8020000000000000</Keywords>
    <TimeCreated SystemTime="2018-11-26T20:42:07.323695200Z"/>
    <EventRecordID>1624709</EventRecordID>
    <Correlation/>
    <Execution ProcessID="4" ThreadID="4392"/>
    <Channel>Security</Channel>
    <Computer>MY-PC</Computer>
    <Security/>
  </System>
  <EventData>
    <Data Name="SubjectUserSid">S-1-2-14</Data>
    <Data Name="SubjectUserName">GRAVUSER$</Data>
    <Data Name="SubjectDomainName">WORKGROUP</Data>
    <Data Name="SubjectLogonId">0x3e3</Data>
    <Data Name="Status">0x0</Data>
    <Data Name="ProcessId">0x1384</Data>
    <Data Name="ProcessName">C:\Windows\servicing\TrustedInstaller.exe</Data>
  </EventData>
</Event>
```

Because the event logs can span multiple lines, it is not safe to use the basic listener. However, we can create a *regex* listener which splits entries apart using the trailing `</Event>` tag. The following is a regex listener definition which can split on Windows event logs like the one above:

```
[RegexListener "windows"]
	Bind-String="0.0.0.0:6666"
	Tag-Name="winevent"
	Regex=`(?P<suffix></Event>)`
```

The behavior of the Regex parameter is discussed below.

### Regex Parameter Details

The `Regex` parameter specifies a regular expression to be used as the delimiter. The ingester will read incoming data and buffer it until it finds a match for the regular expression. It then takes all the buffered data *up to but not including* the delimiter and ingests it as an entry. The delimiter itself is *discarded*. Thus, if we define `Regex="X"` and send the following:

```
fooXbarXbaz
```

We will see three entries: "foo", "bar", and "baz".

Note that sometimes, you'll wish to keep portions of the delimiter. In the Windows event log example given earlier, we match on the trailing `</Event>` tag on the log, but we need to include that in the outgoing entry! The regex listener will look for two special [named capture groups](https://www.regular-expressions.info/refcapture.html) in the `Regex` parameter: "prefix" and "suffix". The contents of the "suffix" group will be attached at the end of the current entry; in the example above, that means that `</Event>` will be included at the end. The contents of the "prefix" group are stored and attached at the *start* of the *next* entry; this is useful if you're matching on the start of the next entry, for instance for multi-line syslogs:

```
Regex=`(?P<prefix><\d+>1 \d{4}-\d{1,2}-\d{1,2}T)`
```

### Trim-Whitespace Parameter

If `Trim-Whitespace` is set to true, any preceding or trailing whitespace on the outgoing entry data will be removed.

### Max-Buffer Parameter

The Max-Buffer parameter specifies, in bytes, how much data the regex listener should buffer as it looks for a matching regular expression. The default is 8 MB. If the listener reads more data than that without finding a match for the regex delimiter, it will ingest an entry containing the first `Max-Buffer` stored bytes; although the entries may end up malformed, we consider it more appropriate to ingest questionable data than to throw it away.

## JSON Listeners

The JSON Listener type enables some mild JSON processing at the time of ingest.  The purpose of a JSON reader would be to apply a unique tag to an entry based on the value of a field in a JSON entry.   Many applications export JSON data with a field that indicates the format of the JSON, from a processing efficiency standpoint it can be beneficial to tag the different formats with specific tags.

A great example use case is the JSON over TCP data export functionality found in many Bro sensor appliances.  The appliances export all Bro log data over a single TCP stream, however there are multiple data types within the stream built by different modules.  Using the JSON Listener we can derive the data type from the module field and apply a unique tag.  this allows us to do things like keep the Bro conn logs in one well, the Bro DNS logs in another, and all other Bro logs in yet another.  As a result, we can differentiate the data types with different tags and take advantage of Gravwell Wells when multiple JSON data types are coming in via a single stream.

Additional JSON listener configuration parameters:

| Parameter     | Type    | Default Value     | Description |
|---------------|---------|-------------------|-------------|
| Extractor     | string  |  | A JSON path which specifies the object to extract from the JSON object for use in tag assignment |
| Tag-Match     | string list | | A set of key/value specifications used to match against the Extractor parameter and lookup a tag for assignment, many Tag-Match parameters can be defined |
| Default-Tag   | string  |  | The tag to assign to the entry if no match can be made using the set of Tag-Match parameters |
| Max-Object-Size | unsigned int | 1048576 | Default maximum object size for a JSON object.  Defaults to 1MB and will drop values that are over this size.  This is a safety valve for unauthenticated listeners. |
| Disable-Compact | bool | false | By default the JSON listener will attempt to clean and compact JSON objects.  Set to true to keep JSON objects in their original form. |

### JSON Parameter Details

The JSON Listener blocks implement the universal listener types as documented above.  Additional parameters allow for specifying which field we wish to pivot on to define a tag.

#### Extractor Parameter Details

The "Extractor" parameter specifies a JSON extraction string which is used to pull a field from a JSON entry.  The Extraction string follows the same syntax as the Gravwell [json](/search/json/json) search module minus any inline filtering.

Given the following JSON:

```
{
  "time": "2018-09-12T12:25:33.503294982-06:00",
  "class": 5.1041415140005e+18,
  "data": "Have I come to Utopia to hear this sort of thing?",
  "identity": {
    "user": "alexanderdavis605",
    "name": "Noah White",
    "email": "alexanderdavis605@test.org",
    "phone": "+52 27 83 68 75069 2"
  },
  "location": {
    "address": "43 Wilson Pkwy,\nBury, AL, 66232",
    "state": "PW",
    "country": "Pakistan"
  },
  "group": "carp",
  "useragent": "Mozilla\/5.0 (X11; Fedora; Linux x86_64) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/52.0.2743.116 Safari\/537.36",
  "ip": "8.83.94.200"
}
```

We could extract the location and state value and apply a tag based on which state abbreviation we find using the following Extraction parameter:

```
Extractor=location.state
```

**Tag-Match**

Each JSONListener supports multiple field value to tag match specifications.  The value to tag assignment is specified as an argument to the "Tag-Match" parameter in the form <field value>:<tag name>.

For example, if we extracted a field with the value "foo" and wanted to assign it to the tag "bar" we would add the following to the JSONListener configuration block:

```
Tag-Match=foo:bar
```

The field extraction values can contain ":" characters, to specify a field value with a ":" character in it encapsulate the value with double quotes.

For example, if we wanted to assign the tag "baz" to the extracted value "foo:bar" the "Tag-Match" parameter would be as follows:

```
Tag-Match="foo:bar":baz
```

Extraction value to tag mappings can be many to one, meaning that multiple extraction values can be mapped to the same tag.  For example the following parameters will map both "foo" and "bar" extracted values to the tag "baz":

```
Tag-Match=foo:baz
Tag-Match=bar:baz
```

However, a single extraction value CANNOT be mapped to multiple tags.  The following is invalid:

```
Tag-Match=foo:baz
Tag-Match=foo:bar
```

**Default-Tag**

When extracting fields and applying tags, the JSON Listener will apply a default tag if there is no matching Tag-Match specified.

### Example JSONListener behaviors

Assume the following configured JSONListener:

```
[JSONListener "testing"]
	Bind-String=0.0.0.0:7777
	Extractor="field1"
	Default-Tag=json
	Tag-Match=test1:tag1
	Tag-Match=test2:tag2
	Tag-Match=test3:tag3
```

Some example JSON data and resulting tag:

#### Matched field

```
{ "field1": "test1", "field2": "test2" }
```

The entry gets the tag "tag1" because the field "field1" matched the "Tag-Match=test1:tag1"

#### Unmatched field

```
{ "field1": "foobar", "field2": "test2" }
```

The entry gets the tag "json" because the field "field1" did not match any "Tag-Match" parameters.

##### Extraction field not found

```
{ "fieldfoo": "test1", "fieldbar": "test2" }
```

The entry gets the tag "json" because the extractor could not find the field "field1".

## TLS Configuration

All listener types (line, syslog, regex, and JSON) support TLS connections. To enable TLS, use `tls://` in the `Bind-String` field and set the `Cert-File` and `Key-File` parameters to point at a valid TLS certificate and private key:

```
[JSONListener "testing"]
	Bind-String=tls://0.0.0.0:7777
	Cert-File=/opt/gravwell/etc/cert.pem
	Key-File=/opt/gravwell/etc/key.pem
	Extractor="field1"
	Default-Tag=json
	Tag-Match=test1:tag1
	Tag-Match=test2:tag2
	Tag-Match=test3:tag3
```
