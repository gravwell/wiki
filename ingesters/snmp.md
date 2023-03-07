---
myst:
  substitutions:
    package: "gravwell-snmp"
    standalone: "gravwell_snmp_ingest"
    dockername: ""
---
# SNMP Trap Ingester

The SNMP ingester can receive SNMP traps for SNMP versions 2c and 3. The trap messages are ingested in a JSON structure for ease of use.

## Installation

```{include} installation_instructions_template.md 
```

## Basic Configuration

The SNMP ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters SNMP supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

To receive SNMP traps, you must define at least one `Listener` block in the configuration as well, which tells the ingester which port to listen on, which SNMP version it should expect, etc. Below is an example configuration with two Listeners:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO
Log-File=/opt/gravwell/log/snmp.log

[Listener "default"]
	Tag-Name=snmp
	Bind-String="0.0.0.0:162"
	Version=2c
	Community="public"

[Listener "v3"]
	Tag-Name=snmp3
	Bind-String="0.0.0.0:163"
	Version=3
	Username=myuser
	Auth-Passphrase=myauthpw
	Auth-Protocol=MD5
	Privacy-Passphrase=myprivpw
	Privacy-Protocol=DES
```

The Listener named "default" listens for SNMP version 2c traps on UDP port 162, ingesting them into the tag `snmp`. The Listener named "v3" listens on UDP port 163 for SNMP version 3 traps. It requires incoming messages to be authenticated with the passphrase "myauthpw" and the MD5 auth protocol, and encrypted using the shared key "myprivpw" and the DES protocol.

## Listener Configs

Listeners support the following configuration parameters:

| Parameter | Type | Description |
|-----------|------|-------------|
| Bind-String | string | (Required) An IP:port pair on which to listen for SNMP traps, e.g. 0.0.0.0:162. UDP is assumed. |
| Version | string | (Required) SNMP version, either "2c" or "3". |
| Username | string | Username for v3 authentication. |
| Auth-Passphrase | string | Passphrase for v3 authentication. |
| Auth-Protocol | string | Protocol for v3 authentication, "MD5" and "SHA" supported. |
| Privacy-Passphrase | string | Passphrase for v3 encryption. |
| Privacy-Protocol | string | Protocol for v3 encryption, currently only "DES" is supported. |
| Tag-Name  | string | The tag into which traps will be ingested. |
| Community | string | The SNMP community for the listener. If not set, any incoming community value is acceptable. |
| Source-Override | ip | Optional override for the SRC field, otherwise trap sender is used. |
| Preprocessor | string | Name of a preprocessor to apply to ingested data. Many Preprocessor parameters can be applied. |

### Bind-String

The Bind-String parameter controls which interface and port the listener will bind to.  All listeners are currently UDP only.  IPv4 and IPv6 are supported.

```
#bind to all interfaces on UDP port 7777
Bind-String=0.0.0.0:7777

#bind to IPv6 globally routable address on UDP port 901
Bind-String=[2600:1f18:63ef:e802:355f:aede:dbba:2c03]:901
```

### Version

The Version parameter specifies which SNMP protocol version this Listener should speak.

* If set to "2c", consider setting the `Community` parameter as well.
* If set to "3", evaluate if you need to also set the auth & encryption parameters as well (see later sections)

### Community

SNMP v2c provides a very basic security method: agents (clients) wishing to send a trap message must know the correct "community" string that the manager (server/ingester) is expecting. Incoming trap messages whose community string does not match the value in the Community parameter will be dropped (and logged).

Note that if you do not set Community on a version 2c Listener, it will accept traps with *any* Community string.

### Authentication and Privacy

SNMP v3 provides more advanced options for authentication and privacy compared to version 2. Messages may be authenticated with a password, and the contents of the messages may be encrypted for privacy. These two functions can be enabled separately; you can set up an SNMP v3 listener with no auth or privacy at all, with only authentication, or with both authentication and privacy.

To enable authentication, set `Username`, `Auth-Passphrase`, and `Auth-Protocol` (which can be "MD5" or "SHA"):

```
[Listener "v3-auth"]
	Tag-Name=snmp3
	Bind-String="0.0.0.0:9163"
	Version=3
	Username=myuser
	Auth-Passphrase=authpw
	Auth-Protocol=MD5
```

To enable encryption of the trap itself, add the `Privacy-Passphrase` and `Privacy-Protocol` (currently only "DES" is supported) parameters:

```
[Listener "v3-authpriv"]
	Tag-Name=snmp3
	Bind-String="0.0.0.0:9163"
	Version=3
	Username="myuser"
	Auth-Passphrase="authpw"
	Auth-Protocol=MD5
	Privacy-Passphrase="privpw"
	Privacy-Protocol=DES
```

Incoming messages whose authentication and privacy settings do not match those of the listener will be dropped. The ingester will log when it drops a message for this reason.
