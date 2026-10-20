# macOS Unified Log

The macOS Unified Log ingester streams events from the macOS [unified logging system](https://developer.apple.com/documentation/os/logging) into Gravwell. It runs `/usr/bin/log stream --style json` and ingests each event it emits as a single-line JSON entry.

The ingester only runs on macOS; there is no Linux or Windows build. Running as a non-root user works and still ingests events, but macOS redacts any field a format string marks `%{private}` down to the literal string `<private>` unless the reader is root, and refuses to emit `Level=debug` output to a non-root reader at all. The installer therefore runs the ingester as root via a `launchd` daemon rather than a per-user agent, so output is never silently redacted. See [macOS Configuration](macos_unified_log_macos_configuration) below.

```{attention}
If a stream's `Type` parameter is left unset entirely, `log stream` streams **all three** event types: ordinary `log` messages, `activity` tracing spans, and `trace` signposts. The latter two are high-volume OS instrumentation events rather than human-readable log messages, and leaving `Type` unset is usually far noisier than intended. Set `Type=log` unless you specifically want the other two.
```

(macos_unified_log_installation)=
## Installation

Like the macOS [File Follower](file_follow), the macOS Unified Log ingester is distributed as a signed, notarized `.pkg` installer, available from the [downloads page](/quickstart/downloads). The package installs the ingester binary and a default configuration file under `/Library/Gravwell/MacOSUnifiedLog/` and registers a LaunchDaemon so the ingester starts automatically on boot and runs as root.

Unlike the Windows File Follower installer, the macOS package does not prompt for connection details during install. It lays down a default configuration with no indexer target configured; you must edit it by hand (see [File Locations](macos_unified_log_file_locations)) and restart the service before it will connect to anything.

```{note}
The installer requires macOS 12.0 (Monterey) or later. The unified logging system and `log stream --style json` have been available since macOS 10.12 (Sierra), but this ingester is built with `CGO_ENABLED=0`, and under that configuration Go's internal linker stamps the binary's minimum macOS version as 12.0, independent of any `MACOSX_DEPLOYMENT_TARGET` setting.
```

(macos_unified_log_file_locations)=
## File Locations

The installer places the ingester binary and its configuration file together in the same directory, rather than splitting them the way the Windows installer does, matching the existing macOS File Follower installer.

| File | Location |
| :--- | :------- |
| Ingester binary | `/Library/Gravwell/MacOSUnifiedLog/macosUnifiedLog` |
| Configuration file | `/Library/Gravwell/MacOSUnifiedLog/macos_unified_log.conf` |
| LaunchDaemon | `/Library/LaunchDaemons/io.gravwell.macosunifiedlog.plist` |

```{note}
If you run the `macosUnifiedLog` binary directly without passing `-config-file`, it falls back to looking for a configuration file at `/opt/gravwell/etc/macos_unified_log.conf`, with overlays in `/opt/gravwell/etc/macos_unified_log.conf.d`. This only matters if you are invoking the binary yourself; the installed LaunchDaemon always passes `-config-file` pointing at the path above.
```

## Basic Configuration

The macOS Unified Log ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters). Like most other Gravwell ingesters, it supports multiple upstream indexers, TLS, cleartext, and named pipe connections, an ingest cache, and local logging.

At least one `[Stream "name"]` stanza is required. Each stanza runs its own `log stream` process, so different predicates and filters can be routed to different tags. Here is a minimal configuration with a single stream that ingests ordinary log messages, excluding the high-volume `activity` and `trace` event types, into the `macos` tag:

```
[Global]
	Ingest-Secret = IngestSecrets
	Connection-Timeout = 0
	Insecure-Skip-TLS-Verify = false
	Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when the ingester is on the same machine as a backend
	Log-Level = INFO
	Log-File = /Library/Gravwell/MacOSUnifiedLog/macos_unified_log.log

[Stream "system"]
	Tag-Name = macos
	Type = log
```

A stream can also narrow events down with a predicate and route them to a dedicated tag. This example captures only authorization-related events, including both ordinary log messages and activity-tracing spans:

```
[Stream "auth"]
	Tag-Name = macos-auth
	Predicate = "subsystem == 'com.apple.Authorization' OR process == 'sudo'"
	Type = log
	Type = activity
```

```{note}
`Type` can be repeated to select more than one event type, as in the example above. It cannot be set as a comma-separated list on one line; each occurrence is a separate value.
```

### Additional Global Configuration Parameters

| Config Parameter        | Type    | Required | Default Value | Description |
|--------------------------|---------|----------|----------------|-------------|
| Disable-Hostname-Attach | boolean | NO       | false          | When false, every entry is stamped with a `host` [intrinsic enumerated value](intrinsic_enumerated_values) containing this machine's hostname, which makes it possible to tell entries from different Macs apart when several report into the same Gravwell instance without configuring a separate `attach` preprocessor. Set to `true` to disable this. |

## Stream Configuration

| Parameter          | Type         | Required | Default Value | Description |
|--------------------|--------------|----------|----------------|-------------|
| Tag-Name           | string       | NO       | `default`      | Tag entries from this stream are ingested into. |
| Level              | string       | NO       |                | `default`, `info`, or `debug`, passed to `log stream` as `--level`. Each level includes everything the ones before it capture; `info` and `debug` are substantially more verbose. |
| Predicate          | string       | NO       |                | An [NSPredicate](https://developer.apple.com/documentation/foundation/nspredicate) filter expression, passed as `--predicate`. It is passed directly to the `log` process rather than through a shell, so it needs no additional quoting. |
| Process            | string array | NO       |                | A process name or PID, passed as `--process`. Can be repeated to filter on multiple processes. |
| Type               | string array | NO       | all three types | `activity`, `log`, or `trace`, passed as `--type`. Can be repeated, with each occurrence as its own line, to select more than one type. If left unset, no `--type` flag is passed and `log stream` streams all three types. See the warning at the <a href="#macos-unified-log">top of this page</a>. |
| Include-Source     | boolean      | NO       | false          | Include symbol names and source line numbers where available (`--source`). |
| Ignore-Timestamps  | boolean      | NO       | false          | Stamp entries with the time they were read instead of parsing the event's own `timestamp` field. |
| Source-Override    | string       | NO       |                | Overrides the entry's SRC field for this stream only. See [Source-Override](ingesters_source-override). |
| Preprocessor       | string array | NO       |                | [Preprocessors](/ingesters/preprocessors/preprocessors) to apply to this stream's entries. Can be repeated. |

## Behavior

* Entry timestamps are parsed from each event's own `timestamp` field. If the ingester cannot parse it, it falls back to the current time and logs one warning per stream so a persistently unparseable format isn't logged on every event.
* Anything `log` writes to stderr, such as an invalid predicate or a permissions error, is relayed into the ingester's own log.
* If the `log` process exits, the ingester restarts it, with a restart delay that starts at 1 second and backs off to 30 seconds while the process keeps failing without running long enough to be considered healthy. A process that runs long enough resets the backoff, so a crash-looping process backs off instead of hot-looping, while a stream that occasionally restarts cleanly isn't penalized by a long delay it doesn't need.
* An event larger than 16MB, or output the ingester cannot parse, causes the `log` process for that stream to restart.
* Events emitted while a stream's `log` process is restarting, or while the ingester itself is stopped, are not backfilled.

(macos_unified_log_macos_configuration)=
## macOS Configuration

The macOS Unified Log ingester installs to `/Library/Gravwell/MacOSUnifiedLog` and runs under a `launchd` LaunchDaemon so that it starts automatically at boot and runs as root. To change its configuration, edit `/Library/Gravwell/MacOSUnifiedLog/macos_unified_log.conf` and restart the service with:

```
launchctl kickstart -k system/io.gravwell.macosunifiedlog
```

To disable the service, use:

```
launchctl disable system/io.gravwell.macosunifiedlog
```

To remove the macOS Unified Log ingester from macOS, remove the following:

1. `/Library/Gravwell/MacOSUnifiedLog`
2. `/Library/LaunchDaemons/io.gravwell.macosunifiedlog.plist`

## Example Queries

Show every entry from one Mac in a fleet reporting into the same Gravwell instance:

```gravwell
tag=macos intrinsic host == "alices-macbook-pro.local" | table TIMESTAMP DATA
```

Find error-level messages from a specific subsystem:

```gravwell
tag=macos json subsystem == "com.apple.Authorization" | json eventMessage messageType == "Error" | table TIMESTAMP eventMessage
```

Count log volume by subsystem to find the noisiest ones:

```gravwell
tag=macos json subsystem | count by subsystem | sort by count desc | table subsystem count
```

Find every message a specific process emitted:

```gravwell
tag=macos json eventMessage processImagePath == "\\/kernel" | table TIMESTAMP eventMessage processImagePath
```

Chart message volume over time for a dedicated, predicate-filtered stream such as the `macos-auth` example above:

```gravwell
tag=macos-auth | count | chart count
```
