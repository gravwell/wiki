# Simple Relay

Simple Relay is the go-to ingester for text based data sources that can be delivered over plaintext TCP and/or UDP network connections via either IPv4 or IPv6.

Some common use cases for Simple Relay are:

* Remote syslog collection
* Devop log collection over a network
* Bro sensor log collection
* Simple integration with any text source capable of delivering over a network

## Basic Configuration

The Simple Relay ingester uses the unified global configuration block described in the [ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters).  Like most other Gravwell ingesters Simple Relay supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

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
	#Lack of "Reader-Type" implines line break delimited logs
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

Note: The `Keep-Priority` field is necessary if you plan to analyze syslog entries with the [syslog search module](#!search/syslog/syslog.md).

## Listeners

Each listener contains a set of universal configuration values, regardless of whether the listener is a line reader, RFC5424 reader, or JSON Listener.

### Universal Listener Configuration Parameters

Listeners support several configuration parameters which allow for specifying protocols, listening interfaces and ports, and fine tuning ingest behavior.

#### Bind-String

The Bind-String parameter controls which interface and port the listener will bind to.  The listener can bind to TCP or UDP ports, specific addresses, and specific ports.  IPv4 and IPv6 are supported.

```
#bind to all interfaces on TCP port 7777
Bind-String=0.0.0.0:7777

#bind to all interfaces on UDP port 514
Bind-String=udp://0.0.0.0:514

#bind to port 1234 on link local IPv6 address on interface p1p
Bind-String=[fe80::4ecc:6aff:fef9:48a3%p1p1]:1234

#bind to IPv6 globally routable address on TCP port 901
Bind-String=[2600:1f18:63ef:e802:355f:aede:dbba:2c03]:901
```

#### Ignore-Timestamps

The "Ignore-Timestamps" parameter instructs the listener to not attempt to derive a timestamp from the read values, instead apply the current timestamp.  This parameter is useful for reading data where there may not be a timestamp present, or the timestamp is wrong on the originating system due to unreliable system clocks.  "Ignore-Timestamps" is false by default, to enable specify ```Ignore-Timestamps=true```

#### Assume-Local-Timezone and Timezone-Override

Most timestamp formats have a timezone attached which indicates an offset to Universal Cordinated Time (UTC).  However, some systems do not specify the timezone leaving it up to the receiver to determine what timezone a log entry may be in.  Assume-Local-Timezone causes the reader to assume that the timestamp is in the same timezone as the Simple Relay reader when the timzeone is omitted. Timezone-Override takes a string in the IANA timezone database format (e.g. "America/Chicago") and applies that timezone to timestamps which do not specify a timezone.

Assume-Local-Timezone and Timezone-Override are mutually exclusive.

#### Source-Override

The "Source-Override" parameter instructs the listener to ignore the source of the data and apply a hard coded value.  It may be desirable to hard code source values for incoming data as a method to organize and/or group data sources.  "Source-Override" values can be IPv4 or IPv6 values.

```
Source-Override=192.168.1.1
Source-Override=127.0.0.1
Source-Override=[fe80::899:b3ff:feb7:2dc6]
```

#### Timestamp-Format-Override

Data values may contain multiple timestamps which can cause some confusion when attempting to derive timestamps out of the data.  Normally, the Listeners will grab the left most timestamp that can be derived, but it may be desirable to only look for a timestamp in a very specific format.  "Timestamp-Format-Override" tells the listener to only respect timestamps in a specific format.  The following timstamp formats are available:

* AnsiC
* Unix
* Ruby
* RFC822
* RFC822Z
* RFC850
* RFC1123
* RFC1123Z
* RFC3339
* RFC3339Nano
* Apache
* ApacheNoTz
* Syslog
* SyslogFile
* SyslogFileTZ
* DPKG
* Custom1Milli
* NGINX
* UnixMilli
* ZonelessRFC3339
* SyslogVariant
* UnpaddedDateTime

To force the Listener to only look for timestamps that match the RFC3339 specification add ```Timestamp-Format-Override=RFC3339``` to the Listener.

### Listener Reader Types and configurations

Simple relay supports the following types of basic readers which are useful in different contexts

* line reader
* RFC5424

The basic listeners are specified via the "Listener" block and support line delimited and RFC5424 reader types.  Each Listener must have a unique name and a unique bind port (two different listeners cannot bind to the same protocol, address, and port).  The reader type for a basic listener is controlled by the "Reader-Type" parameter.  Currently there are two types of listeners (line and RFC5424).  If no Reader-Type is specified the line reader type is assumed.

Basic Listeners also require that each listener designate the tag the listener will apply to all incoming data via the "Tag-Name" parameter.  If the "Tag-Name" parameter is omitted, the "default" tag is applied.  The most basic listener named "test" which expects line broken data on TCP port 5555 and applies the tag "testing" would have the following configuration specification:

```
[Listener "test"]
	Bind-String=0.0.0.0:5555
	Tag-Name=testing
```

#### Line Reader Listener

The line reader listener is designed to read newline broken data streams from either a TCP or UDP stream.  Applications which can deliver simple line broken data over a network can utilize this type of reader to very simply and easily integrate with Gravwell.  The Line Reader listener can also be used for simple log file delivery by simply sending log files to the listening port.

For example, an existing log file can be imported into Gravwell using netcat and Simple Relay
```
nc -q 1 10.0.0.1 7777 < /var/log/syslog
```

##### Example Line Reader Listener

The most basic Listener requires only one the "Bind-String" argument which tells the listener what port to listen on. 

```
[Listener "default"]
	Bind-String="0.0.0.0:7777" #bind to all interfaces, with TCP implied
```

#### RFC5424 Listener

A listener designed to accept structured syslog messages based on either RFC5424 or RFC3164 enables Simple Relay to act as a syslog aggregation point.  To enable a listener that expects syslog messages using a reliable TCP connection on port 601 set the "Reader-Type" to "RFC5424.

```
[Listener "syslog"]
	Bind-String=0.0.0.0:601
	Reader-Type=RFC5424
```

To accept syslog messages over stateless UDP via port 514 the listener would look like the following:

```
[Listener "syslog"]
	Bind-String=udp://0.0.0.0:514
	Reader-Type=RFC524
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
	Reader-Type=RFC524
	Keep-Priority=false
```

Note: The priority portion of a syslog message is codified in the RFC specification.  Removing the priority means that the Gravwell [syslog](#!search/syslog/syslog.md) search module will be unable to properly parse the values.  Paid Gravwell licenses are all unlimited and we reccomend that the priority field is left in syslog messages.  The syslog search module is also dramatically faster than attempting to hand parse syslog messages with regular expressions.

### JSON Listeners

The JSON Listener type enables some mild JSON processing at the time of ingest.  The purpose of a JSON reader would be to apply a unique tag to an entry based on the value of a field in a JSON entry.   Many applications export JSON data with a field that indicates the format of the JSON, from a processing efficiency standpoint it can be beneficial to tag the different formats with specific tags.

A great example use case is the JSON over TCP data export functionality found in many Bro sensor appliances.  The appliances export all Bro log data over a single TCP stream, however there are multiple data types within the stream built by different modules.  Using the JSON Listener we can derive the data type from the module field and apply a unique tag.  this allows us to do things like keep the Bro conn logs in one well, the Bro DNS logs in another, and all other Bro logs in yet another.  As a result, we can differentiate the data types with different tags and take advantage of Gravwell Wells when multiple JSON data types are coming in via a single stream.

#### JSON Listener Configuration Parameters

The JSON Listener blocks implement the universion listener types as documented above.  Additional parameters allow for specifying which field we wish to pivot on to define a tag.

##### Extractor Parameter

The "Extractor" parameter specifies a JSON extraction string which is used to pull a field from a JSON entry.  The Extraction string follows the same syntax as the Gravwell [json](#!search/json/json.md) search module minus any inline filtering.

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

##### Tag-Match

Each JSONListener supports multiple field value to tag match specifications.  The value to tag assignment is specified as an argument to the "Tag-Match" paramter in the form <field value>:<tag name>.

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

##### Default-Tag

When extracting fields and applying tags, the JSON Listener will apply a default tag if there is no matching Tag-Match specified.

#### Example JSONListener behaviors

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

##### Matched field

```
{ "field1": "test1", "field2": "test2" }
```

The entry gets the tag "tag1" because the field "field1" matched the "Tag-Match=test1:tag1"

##### Unmatched field

```
{ "field1": "foobar", "field2": "test2" }
```

The entry gets the tag "json" because the field "field1" did not match any "Tag-Match" parameters.

##### Extraction field not found

```
{ "fieldfoo": "test1", "fieldbar": "test2" }
```

The entry gets the tag "json" because the extractor could not find the field "field1".
