# Winlog

The winlog processor is a special-purpose extractor for XML-formatted Windows log data. It provides "shortcuts" for extracting many common fields from Windows log entries rather than requiring the more general [xml module](/#!search/xml/xml.md)

## Supported Options

* `-e`: The “-e” option specifies that the winlog module should operate on an enumerated value.  Operating on enumerated values can be useful when you have extracted log entries using upstream modules.
* `-or`: The "-or" flag specifies that the winlog module should allow an entry to continue down the pipeline if ANY of the extractions or filters are successful.

## Processing Operators

Each winlog field supports a set of operators that can act as fast filters. In the case of the winlog module, all fields will be extracted as strings, so only string filters are available.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field must be a member of
| !~ | Not subset | Field must not be a member of

## Data Fields

Given a log entry in this format:

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

The following fields can be extracted

| Field | XML spec | Type | Filter Options |
|-------|----------|------|----------------|
| System | Event.System | bytes | == != ~ !~ |
| EventData | Event.EventData | bytes | == != ~ !~ |
| UserData | Event.UserData | bytes | == != ~ !~ |
| Provider | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderName | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderGUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| GUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| EventID | Event.System.EventID | uint | == != < <= > >= |
| Version | Event.System.Version | uint | == != < <= > >= |
| Level | Event.System.Level | uint | == != < <= > >= |
| Task | Event.System.Task | uint | == != < <= > >= |
| Opcode | Event.System.Opcode | bytes | == != ~ !~ |
| Keywords | Event.System.Keywords | bytes | == != ~ !~ |
| TimeCreated | Event.System.TimeCreated[SystemTime] | bytes | == != ~ !~ |
| EventRecordID | Event.System.EventRecordID | uint | == != < <= > >= |
| ProcessID | Event.System.Execution[ProcessID] | uint | == != < <= > >= |
| ThreadID | Event.System.Execution[ThreadID] | uint | == != < <= > >= |
| Channel | Event.System.Channel | bytes | == != ~ !~ |
| Computer | Event.System.Computer | bytes | == != ~ !~ |
| Correlation | Event.System.Correlation | bytes | == != ~ !~ |
| UserID | Event.System.Security[UserID] | uint | == != < <= > >= |

Specifying a field not listed above will cause the winlog module to attempt to extract `Event.Data[Name]==<field>`. For example, the SubjectLogonId in the example above (0x3e3) can be extracted by simply specifying `SubjectLogonId` to the winlog module.

### Data Field Filtering

Every extractable field in a windows log supports inline filtering with various comparison operations depending on the type of the field.  Some of the Event System fields are integers and can be compared as an integer, such as the `EventID`, `Version`, and `Level` fields.  The other System fields and all data fields are treated as byte arrays.  As with most Gravwell search modules it is almost always faster to perform filtering right when an enumerated value is created.


## Examples

The following examples refer to the sample log shown above.

To extract the process ID (4) and the user name (GRAVUSER$):

```
winlog ProcessID SubjectUserName
```

To extract the process name from only those events on the Security channel with EventID == 4689:

```
winlog EventID==4689 Channel==Security ProcessName
```
