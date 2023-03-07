---
myst:
  substitutions:
    package: "gravwell-ipmi"
    standalone: "gravwell_ipmi"
    dockername: "ipmi"
---
# IPMI Ingester

The IPMI Ingester collects Sensor Data Record (SDR) and System Event Log (SEL) records from any number of IPMI devices. 

The configuration file provides a simple host/port, username, and password field for connecting to each IPMI device. SEL and SDR records are ingested in a JSON-encoded schema. For example:

```
{
    "Type": "SDR",
    "Target": "10.10.10.10:623",
    "Data": {
        "+3.3VSB": {
            "Type": "Voltage",
            "Reading": "3.26",
            "Units": "Volts",
            "Status": "ok"
        },
        "+5VSB": {...},
        "12V": {...}
    }
}

{
    "Target": "10.10.10.10:623",
    "Type": "SEL",
    "Data": {
        "RecordID": 25,
        "RecordType": 2,
        "Timestamp": {
            "Value": 1506550240
        },
        "GeneratorID": 32,
        "EvMRev": 4,
        "SensorType": 5,
        "SensorNumber": 81,
        "EventType": 111,
        "EventDir": 0,
        "EventData1": 240,
        "EventData2": 255,
        "EventData3": 255
    }
}
```

## Installation

```{include} installation_instructions_template.md 
```

## Basic Configuration

The IPMI ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, the IPMI ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

## Configuration Options

IPMI uses the default set of Global configuration options. IPMI devices are configured with an "IPMI" stanza and each stanza can support multiple IPMI devices that share the same credentials. For example:

```
[IPMI "Server 1"]
	Target="127.0.0.1:623"
	Target="1.2.3.4:623"
	Username="user"
	Password="pass"
	Tag-Name=ipmi
	Rate=60
	Source-Override="DEAD::BEEF" 
```

The IPMI stanza is simple, only taking one or more Targets (the IP:PORT of the IPMI device), username, password, tag, and a poll rate, in seconds. The default poll rate is 60 seconds. Optionally, you can set a source override to force the SRC field on all ingested entries to another IP. By default, the SRC field is set to the IP of the IPMI device. 

Additionally, all IPMI stanzas can use the "Preprocessor" options, as described [here](/ingesters/preprocessors/preprocessors).
