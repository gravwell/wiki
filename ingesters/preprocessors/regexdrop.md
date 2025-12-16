# Regex Drop Preprocessor

The Regex Drop preprocessor filters entries based on regular expression matching. Entries that match the regex pattern are dropped from ingestion, allowing you to selectively exclude unwanted data before it reaches your Gravwell instance.

The Regex Drop preprocessor Type is `regexdrop`.

## Supported Options

* `Regex` (string, required): The regular expression pattern to match against entry data. Supports standard Go regular expression syntax.
* `Invert` (boolean, optional): When set to `true`, the matching behavior is inverted—entries that **match** the regex are kept, and entries that **don't match** are dropped. Default is `false`.

## Common Use Cases

The regexdrop preprocessor is commonly used for:

* Filtering out noisy or irrelevant log entries
* Excluding debug or verbose logging levels
* Dropping entries that contain specific patterns (e.g., health checks, heartbeats)
* Keeping only entries that match a specific pattern (using Invert mode)

### Example: Dropping Debug Logs

To drop all log entries that contain "DEBUG" level messages:

```
[2024-01-15 10:30:45] DEBUG User session started
[2024-01-15 10:30:46] INFO Login successful
[2024-01-15 10:30:47] DEBUG Cache refreshed
[2024-01-15 10:30:48] ERROR Connection failed
```

Use the following configuration:

```
[Preprocessor "drop-debug"]
	Type=regexdrop
	Regex=`\bDEBUG\b`
```

The result is that only non-DEBUG entries are ingested:

```
[2024-01-15 10:30:46] INFO Login successful
[2024-01-15 10:30:48] ERROR Connection failed
```

### Example: Dropping Health Check Requests

To filter out health check endpoints from web server logs:

```
192.168.1.10 - - "GET /health HTTP/1.1" 200
192.168.1.11 - - "GET /api/users HTTP/1.1" 200
192.168.1.10 - - "GET /healthz HTTP/1.1" 200
192.168.1.12 - - "POST /api/login HTTP/1.1" 201
```

Use the following configuration:

```
[Preprocessor "drop-healthchecks"]
	Type=regexdrop
	Regex=`GET /health[z]? HTTP`
```

The result is:

```
192.168.1.11 - - "GET /api/users HTTP/1.1" 200
192.168.1.12 - - "POST /api/login HTTP/1.1" 201
```

### Example: Keeping Only Error Logs (Invert Mode)

To keep only entries that contain error-level messages and drop everything else:

```
[2024-01-15 10:30:45] INFO Application started
[2024-01-15 10:30:46] ERROR Database connection failed
[2024-01-15 10:30:47] WARN Disk space low
[2024-01-15 10:30:48] ERROR Authentication timeout
```

Use the following configuration with `Invert=true`:

```
[Preprocessor "keep-errors-only"]
	Type=regexdrop
	Regex=`\bERROR\b`
	Invert=true
```

The result is that only ERROR entries are ingested:

```
[2024-01-15 10:30:46] ERROR Database connection failed
[2024-01-15 10:30:48] ERROR Authentication timeout
```

### Example: Filtering JSON Logs by Field Value

To drop JSON log entries where a specific field has a certain value:

```
{"level":"debug","message":"Starting process"}
{"level":"info","message":"User logged in"}
{"level":"debug","message":"Cache hit"}
{"level":"error","message":"Connection refused"}
```

Use the following configuration:

```
[Preprocessor "drop-json-debug"]
	Type=regexdrop
	Regex=`"level":"debug"`
```

The result is:

```
{"level":"info","message":"User logged in"}
{"level":"error","message":"Connection refused"}
```
