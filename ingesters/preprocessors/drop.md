## Drop Preprocessor

The drop preprocessor does exactly what the name implies: it drops entries from the ingest pipeline, effectively throwing them away.

This processor can be useful if an ingest stream is to be primarily handled by another set of preprocessors.  For example, if you want to send data to a remote system using the forwarder preprocessor but *not* ingest it upstream into a Gravwell indexer, you could add a final `drop` preprocessor which will simply discard all entries that it sees.

### Supported Options

None.

### Example: Just Drop Everything

This example has a single preprocessor `drop` which just discards all entries on a Simple Relay listener.

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=dropit

[Preprocessor "dropit"]
	Type=Drop               
```

### Example: Forward Entries and Drop

This example forwards entries to another system via a TCP forwarder, then drops them before they can be ingested into Gravwell.

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=forwardprivnet
	Preprocessor=dropit

[Preprocessor "forwardprivnet"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false

[Preprocessor "dropit"]
	Type=Drop               
```

