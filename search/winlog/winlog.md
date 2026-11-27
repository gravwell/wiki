# Winlog

The winlog module extracts fields from XML-formatted Windows event logs. It provides short names for common fields, so you don't need full paths as with the more general [xml module](/search/xml/xml).

## Supported Options

* `-e <arg>`: Operate on an enumerated value instead of the entire entry, for example when an upstream module extracted the event XML. Entries without that enumerated value are dropped.
* `-s`: Strict mode. Drop the entry if any extraction fails.

## Processing Operators

Fields can be filtered inline. Which operators are available depends on the field's type: string fields support `==`, `!=`, `~`, and `!~`; numeric fields support `==`, `!=`, `<`, `<=`, `>`, and `>=`. The table under Data Fields lists each field's type.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| < | Less than | Field must be less than (numeric fields)
| <= | Less than or equal | Field must be less than or equal to (numeric fields)
| > | Greater than | Field must be greater than (numeric fields)
| >= | Greater than or equal | Field must be greater than or equal to (numeric fields)
| ~ | Subset | Field contains the value (string fields)
| !~ | Not subset | Field does NOT contain the value (string fields)

Entries that fail a filter are dropped. A filtered field that is missing also drops the entry, unless the operator is `!=` or `!~`.

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

The following fields can be extracted:

| Field | XML spec | Type | Filter Options |
|-------|----------|------|----------------|
| System | Event.System | bytes | == != ~ !~ |
| EventData | Event.EventData | bytes | == != ~ !~ |
| UserData | Event.UserData | bytes | == != ~ !~ |
| Provider | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderName | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderGUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| ProviderGuid | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| GUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| EventSourceName | Event.System.Provider[EventSourceName] | bytes | == != ~ !~ |
| EventID | Event.System.EventID | uint | == != < <= > >= |
| Version | Event.System.Version | uint | == != < <= > >= |
| Level | Event.System.Level | uint | == != < <= > >= |
| Task | Event.System.Task | uint | == != < <= > >= |
| Opcode | Event.System.Opcode | uint | == != < <= > >= |
| Keywords | Event.System.Keywords | bytes | == != ~ !~ |
| TimeCreated | Event.System.TimeCreated[SystemTime] | bytes | == != ~ !~ |
| EventRecordID | Event.System.EventRecordID | uint | == != < <= > >= |
| ProcessID | Event.System.Execution[ProcessID] | uint | == != < <= > >= |
| ThreadID | Event.System.Execution[ThreadID] | uint | == != < <= > >= |
| Channel | Event.System.Channel | bytes | == != ~ !~ |
| Computer | Event.System.Computer | bytes | == != ~ !~ |
| Correlation | Event.System.Correlation | bytes | == != ~ !~ |
| ActivityID | Event.System.Correlation[ActivityID] | bytes | == != ~ !~ |
| UserID | Event.System.Security[UserID] | bytes | == != ~ !~ |
| Binary | Event.EventData.Binary | bytes | == != ~ !~ |
| Data | First Event.EventData.Data element | bytes | == != ~ !~ |

`ProviderGUID` and `ActivityID` are extracted without their surrounding braces; `GUID` and `ProviderGuid` keep them.

Any other field name extracts the `Event.EventData.Data` element whose `Name` attribute matches it. For example, `winlog SubjectLogonId` extracts `0x3e3` from the sample above.

### Data Field Filtering

Fields listed as `uint` above are compared as integers. A `uint` field whose value isn't a valid unsigned integer is treated as missing. All other fields, including every `EventData` field looked up by name, are compared as strings. Filtering in the winlog module is usually faster than filtering later in the pipeline.


## Examples

The following examples refer to the sample log shown above.

To extract the process ID (`4`) and the user name (`GRAVUSER$`):

```
winlog ProcessID SubjectUserName
```

To extract the process name only from events on the `Security` channel with `EventID` 4689:

```
winlog EventID==4689 Channel==Security ProcessName
```
