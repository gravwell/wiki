# JSON Field Filtering Preprocessor

This preprocessor will parse entry data as a JSON object, then extract specified fields and compare them against lists of acceptable values. The lists of acceptable values are specified in files on the disk, one value per line.

It can be configured to either *pass* only those entries whose fields match the lists, or to *drop* those entries which match the lists--whitelisting, or blacklisting. It can be set up to filter against multiple fields, requiring either that *all* fields must match (logical AND) or that *at least one* field must match (logical OR).

This preprocessor is particularly useful to narrow down a firehose of general data before sending it across a slow network link.

The JSON Field Filtering preprocessor Type is `jsonfilter`.

## Supported Options

* `Field-Filter` (string, required): This specifies two things: the name of the JSON field of interest, and the path to a file which contains values to match against. For example, one might specify `Field-Filter=ComputerName,/opt/gravwell/etc/computernames.txt` in order to extract a field named "ComputerName" and compare it against values in `/opt/gravwell/etc/computernames.txt`. The `Field-Filter` option may be specified multiple times in order to filter against multiple fields.
* `Match-Logic` (string, optional): This parameter specifies the logic operation to use when filtering against multiple fields. If set to "and", an entry is only considered a match when *all* specified fields match against the given lists. If set to "or", an entry is considered a match when *any* field matches.
* `Match-Action` (string, optional): This specifies the option which should be take for entries whose fields match the provided lists. It may be set to "pass" or "drop"; if omitted, the default is "pass". If set to "pass", entries which match will be allowed to pass to the indexer (whitelisting). If set to "drop", entries which match will be dropped (blacklisting).

The `Match-Logic` parameter is only necessary when more than one `Field-Filter` has been specified.

```{note}
If a field is specified in the configuration but is not present on an entry, the preprocessor will treat the entry *as if the field existed but did not match anything*. Thus, if you have configured the preprocessor to only pass those entries whose fields match your whitelist, an entry which lacks one of the fields will be dropped.
```

## Common Use Cases

The json field filtering preprocessor can down select entries based on fields within the entries.  This can be used to build blacklists and whitelists on data flows to ensure that specific data either does or does not make it to storage.

## Example: Simple Whitelisting

Suppose we have an endpoint monitoring solution which is sending thousands of events per second detailing things which are occurring across the enterprise. Due to the high event volume, we may decide we only want to index events with a certain severity. Luckily, the events include a Severity field:

```
{ "EventID": 1337, "Severity": 8, "System": "email-server-01.example.org", [...] }
```

We know the Severity field goes from 0 to 9, and we decide we want to only pass events with a severity of 6 or higher. We would therefore add the following to our ingester configuration file:

```
[preprocessor "severity"]
	Type=jsonfilter
	Match-Action=pass
	Field-Filter=Severity,/opt/gravwell/etc/severity-list.txt
```

and set `Preprocessor=severity` on the appropriate data input, for instance if we were using Simple Relay:

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
```

Finally, we create `/opt/gravwell/etc/severity-list.txt` and populate it with a list of acceptable Severity values, one per line:

```
6
7
8
9
```

After restarting the ingester, it will extract the `Severity` field from each entry and compare the resulting value against those listed in the file. If the value matches a line in the file, the entry will be sent to the indexer. Otherwise, it will be dropped.

## Example: Blacklisting

Building on the previous example, we may find that that our endpoint monitoring system is generating a *lot* of high-severity false positives from certain systems. We may determine that events with the `EventID` field set to 219, 220, or 1338 and the `System` field set to "webserver-prod.example.org" and "webserver-dev.example.org" are always false positives. We can define another preprocessor to get rid of these entries before they are sent to the indexer:

```
[preprocessor "falsepositives"]
	Type=jsonfilter
	Match-Action=drop
	Match-Logic=and
	Field-Filter=EventID,/opt/gravwell/etc/eventID-blacklist.txt
	Field-Filter=System,/opt/gravwell/etc/system-blacklist.txt
```

If we now add this preprocessor to the data input configuration *after* the existing one, the ingester will apply the two filters in order:

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
	Preprocessor=falsepositives
```

Last, we create `/opt/gravwell/etc/eventID-blacklist.txt`:

```
219
220
1338
```

and `/opt/gravwell/etc/system-blacklist.txt`:

```
webserver-prod.example.org
webserver-dev.example.org
```

This new preprocessor extracts the `EventID` and `System` fields from every entry which makes it past the first filter. It then compares them against the values in the files. Because we set `Match-Logic=and`, it considers an entry a match if *both* field values are found in the files. Because we set `Match-Action=drop`, any entry which matches on both fields will be dropped. Thus, an entry with EventID=220 and System=webserver-dev.example.org is dropped, while one with EventID=220 and System=email-server-01.example.org will *not* be dropped.


