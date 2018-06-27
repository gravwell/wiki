# Syslog

The syslog processor extracts fields from [RFC 5424-formatted](https://tools.ietf.org/html/rfc5424) syslog messages as ingested with the [Simple Relay ingester](#!ingesters/ingesters.md) (be sure to set the Keep-Priority flag on your listener, or it won't work).

## Supported Options

* `-e`: The “-e” option specifies that the syslog module should operate on an enumerated value.  Operating on enumerated values can be useful when you have extracted a syslog record using upstream modules.  You could e.g. extract syslog records from raw PCAP and pass the records into the syslog module.

## Processing Operators

Each syslog field supports a set of operators that can act as fast filters.  The filters supported by each operator are determined by the data type of the field.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| < | Less than | Field must be less than
| > | Greater than | Field must be greater than
| <= | Less than or equal | Field must be less than or equal to
| >= | Greater than or equal | Field must be greater than or equal to

## Data Fields

The syslog module extracts individual fields from an RFC 5424-formatted syslog record. It makes a best-effort attempt to parse from left to right, meaning that if a field is missing, only those fields to the right of it will be available for a given record.

| Field | Description | Supported Operators | Example |
|-------|-------------|---------------------|---------|
| Facility | Numeric code indicating the facility from which the message originates | > < <= >= == != | Facility == 0
| Severity | Numeric code indicating the severity of the message, with 0 being the most severe and 7 the least | > < <= >= == != | Severity < 3
| Priority | The message priority, defined as (20*Facility)+Severity | > < <= >= == != | Priority >= 100
| Version | The version of the syslog protocol in use | > < <= >= == != | Version != 1
| Timestamp | A string representation of the timestamp provided in the log message | == != | |
| Hostname | The hostname of the machine which originally sent the syslog message | == != | Hostname != "myhost"
| Appname | The application which originally sent the syslog message, e.g. `systemd` | == != | Appname != "dhclient"
| ProcID | A string representing the process which sent the message, often a PID | == != | ProcID != "7053"
| MsgID | A string representing the type of message | == != | MsgID == "TCPIN"
| StructuredData | A collection of key-values containing additional data. See below for further discussion. | == != | |
| Message | The log message itself | == != | Message == "Critical error!" |

Consider the following syslog record (sourced from [https://github.com/influxdata/go-syslog](https://github.com/influxdata/go-syslog)):

```
<165>4 2018-10-11T22:14:15.003Z mymach.it e - 1 [ex@32473 iut="3"] An application event log entry...
```

The syslog module would extract the following fields:

* Facility: 20
* Severity: 5
* Priority: 165
* Version: 4
* Timestamp: "2018-10-11T22:14:15.003Z"
* Hostname: "mymach.it"
* Appname: "e"
* ProcID: <nil> (not set)
* MsgID: "1"
* Message: "An application event log entry..."

The portion `[ex@32473 iut="3"]` is the *Structured Data* section. Structured Data sections contain key-value pairs; to access a value using the syslog module, specify a path to it: executing `syslog StructuredData.ex@32473.iut` will set an enumerated value named `iut` containing the value "3".

Structured Data can consist of multiple elements contained in square brackets. A more complex example might look like `[foo@1234 name="Gravwell" year="2018"][bar@432 cake="lie"]`. Specifying `syslog StructuredData.foo@1234.year StructuredData.bar@432.cake` will extract two enumerated values, `year` with a value of "2018" and `cake` with a value of "lie".

## Examples

### Number of events by severity

```
tag=syslog syslog Severity | count by Severity | chart count by Severity
```

![Number of events by severity](severity.png)

### Number of events at each severity level by application

```
tag=syslog syslog Appname Severity | count by Appname,Severity | table Appname Severity count
```

![Number of events at each severity by application](severity2.png)