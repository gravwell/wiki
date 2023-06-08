---
myst:
  substitutions:
    package: "gravwell-packet-fleet"
    standalone: "gravwell_packet_fleet"
    dockername: ""
---
# Packet Fleet Ingester

The Packet Fleet Ingester provides a mechanism to query Google Stenographer instances and have results ingested per-packet into Gravwell. 

Each Stenographer ingester listens on a given port (```Listen-Address```) and accepts Stenographer queries (see query syntax below) as an HTTP POST. On receiving a query, the ingester returns an integer job ID, and asynchronously queries the Stenographer instance and begins to ingest the returned PCAP. Multiple in-flight queries can be ran concurrently. Job status can be viewed by issuing an HTTP GET on "/status", which returns a JSON-encoded array of in-flight job IDs. 

A simple web interface to submit and view job status is also available by browsing to the specified ingester port.

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

Packet Fleet uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters).  Like most other Gravwell ingesters, Packet Fleet supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/packet_fleet.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/packet_fleet.conf.d`).

## Stenographer Examples

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	Assume-Local-Timezone=false #Default for assume localtime is false
	Source-Override="DEAD::BEEF" #override the source for just this Queue 

[Stenographer "Region 2"]
	URL="https://my.url:1234"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
```

## Configuration Options 

Packet Fleet requires several Global and per-stenographer configuration options. Global settings include setting up TLS (if applicable) and the listen address for the web interface, as shown below:

```
Use-TLS=true
Listen-Address=":9002"
Server-Cert="server.cert"
Server-Key="server.key"
```

For each Stenographer instance, the following stanza is required. The example name `Region 1` here is used by the web interface to list Stenographer instances. 

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	#Assume-Local-Timezone=false #Default for assume localtime is false
	#Source-Override="DEAD::BEEF" #override the source for just this Queue 
```

## Query Language 

A user requests packets from stenographer by specifying them with a very simple
query language.  This language is a simple subset of BPF, and includes the
primitives:

    host 8.8.8.8          # Single IP address (hostnames not allowed)
    net 1.0.0.0/8         # Network with CIDR
    net 1.0.0.0 mask 255.255.255.0  # Network with mask
    port 80               # Port number (UDP or TCP)
    ip proto 6            # IP protocol number 6
    icmp                  # equivalent to 'ip proto 1'
    tcp                   # equivalent to 'ip proto 6'
    udp                   # equivalent to 'ip proto 17'

    # Stenographer-specific time additions:
    before 2012-11-03T11:05:00Z      # Packets before a specific time (UTC)
    after 2012-11-03T11:05:00-07:00  # Packets after a specific time (with TZ)
    before 45m ago        # Packets before a relative time
    before 3h ago         # Packets after a relative time

```{note}
Relative times must be measured in integer values of hours or minutes
as demonstrated above.
```

Primitives can be combined with and/&& and with or/||, which have equal
precedence and evaluate left-to-right.  Parens can also be used to group.

    (udp and port 514) or (tcp and port 8080)

```{note}
This section sourced from [Google Stenographer](https://github.com/google/stenographer/blob/master/README.md)
```
