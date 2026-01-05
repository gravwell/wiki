# Attach Preprocessor

The attach preprocessor allows for attaching intrinsic enumerated values to data streams. This is useful for adding metadata to entries without modifying the underlying data payload. Values can be static strings or dynamic values derived from the environment or entry metadata.

The Attach preprocessor Type is `attach`.

## Supported Options

* Any Key-Value Pair (string, optional): The key specifies the name of the attached value, and the value specifies what to attach. 
* Supported dynamic values:
    * `$HOSTNAME`: The hostname of the machine running the ingester.
    * `$NOW`: The current timestamp when the entry is processed.
    * `$ENV:VAR_NAME`: The value of the environment variable `VAR_NAME`.

```{note}
The `$UUID` dynamic value is NOT supported in the preprocessor version of attach, as it is only available in the global Attach configuration.
```

## Common Use Cases

The attach preprocessor is commonly used to tag data with environmental context like the hostname or deployment environment, which can then be used for filtering or aggregation in search.

### Example: Attaching Hostname and a Static Tag

Given a stream of logs where you want to identify the source host and label the environment as "production" without altering the log message itself:

```
[Listener "syslog"]
	Bind-String="0.0.0.0:514"
	Preprocessor=add_meta

[Preprocessor "add_meta"]
	Type=attach
	host=$HOSTNAME
	environment="production"
```

In this example, every entry processed by the `add_meta` preprocessor will have two enumerated values attached:
1. `host`: Contains the hostname of the system.
2. `environment`: Contains the static string "production".

### Example: Attaching Timestamps

You can also attach the processing timestamp to entries, which can be useful for debugging latency or improved time tracking:

```
[Preprocessor "add_time"]
	Type=attach
	ingress_time=$NOW
```

This will attach a value named `ingress_time` with the current timestamp to every entry.
