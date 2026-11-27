# Anko

The anko module runs a script against search entries. It can do more than eval, but a script takes more work to write, test, and deploy than an eval expression. Scripts are stored as resources in the [resource system](/resources/resources).

The syntax of anko is identical to that of eval; both derive from [github.com/mattn/anko](https://github.com/mattn/anko), with some additional functions added for Gravwell-specific tasks.

Use anko when no other module can do the job: for example, when entries must be compared against earlier entries, when entries must be duplicated, or when extracting data takes complex logic.

Anko runs on the webserver, not the indexers, so every entry that reaches it is sent from the indexers first.

You may also refer to the [scripting overview](/scripting/scripting) for additional details or the [Automation scripting documentation](/scripting/scriptingsearch) for more examples.

## Syntax

`anko <script name> [script arguments]`

The first argument is the name of the resource holding the script. Any further arguments are passed to the script's `Parse` function.

## Example script

This script is a reformatted version of an example from the eval documentation, and is easier to read than the one-line eval version:

```
func Process() {
	if len(Body) <= 10 {
		setEnum("postlen", "short")
	} else if len(Body) > 10 && len(Body) < 300 {
		setEnum("postlen", "medium")
	} else {
		setEnum("postlen", "long")
	}
	return true
}
```

Assuming the script is uploaded to a resource named `CheckPostLen`, the script can be executed like this:

```gravwell
tag=reddit json Body | anko CheckPostLen | count by postlen | table postlen count
```

The `Process` function runs once for every entry that reaches the anko module. It checks the length of the `Body` enumerated value and sets a new enumerated value, `postlen`, based on that length.

```{note}
The `return true` at the end of `Process` is required. `Process` must return a boolean: true passes the entry down the pipeline and false drops it. Returning anything else, or nothing, stops the search with an error.
```

## Essential function definitions

An anko script **must** define either a `Process` function or a `Main` function, but not both. Neither takes arguments.

`Process` is called once per entry. Enumerated values on the entry can be read like variables, as can `DATA`, `TAG` (the tag name), `SRC`, `TIMESTAMP`, and `NOW`. Reading an enumerated value the entry doesn't have stops the search with an error, so check with `hasEnum` first if it might be missing. The return value decides whether the entry continues down the pipeline (true) or is dropped (false).

`Main` is called only once. It must call `readEntry` to fetch each entry and `writeEntry` to pass entries down the pipeline.

Use `Process` instead of `Main` whenever possible; it is much simpler.

### Sample script using Process() function

This script has two modes, chosen by its argument. In `build` mode, it collects every `SrcIP` value (extracted by the packet module) seen in the search and stores the list as a resource. In `apply` mode, it loads that list and drops entries whose `SrcIP` is already on it. It was used to find new devices on a network; the `lookup` module now does the same job with more flexibility.

```
table = make(map[string]interface)
task = "build"

var json = import("encoding/json")

# first arg = "build" or "apply"
func Parse(args) {
	errstr = "argument must be 'build' or 'apply'"
	if len(args) == 1 {
		task = args[0]
	} else {
		return errstr
	}
	switch task {
	case "apply":
		# load the table
		data, _ = getResource("lookuptable")
		json.Unmarshal(data, &table)
	case "build":
	default:
		return errstr
	}
	return nil
}

func Process() {
	if task == "build" {
		s = toString(SrcIP)
		table[s] = true
		return true
	} else if task == "apply" {
		s = toString(SrcIP)
		# create & set an enumerated value named "new" to true or false
		setEnum("new", !table[s])
		# if it's not in the table, return true
		return !table[s]
	}
}

func Finalize() {
	if task == "build" {
		data, err = json.Marshal(table)
		if err != nil {
			return err
		}
		return setResource("lookuptable", data)
	}
}
```

Inside `Process`, the `SrcIP` enumerated value can be **read** like any other variable, but setting an enumerated value requires `setEnum`.

The `table` and `task` variables and the `json` import are declared outside the functions. The available packages are listed below.


### Sample script using Main() function

A `Main` function is harder to write, but it is required to duplicate entries. This script reads entries containing Modbus messages. When a message is a request of type 0x10 ("Write multiple registers"), the script makes one copy of the entry for each register written, and sets `RegAddr` and `RegValue` enumerated values on each copy to that register's address and value.

```{note}
This script does not work on its own. It expects an earlier anko script in the pipeline to set enumerated values such as `Request` and `WriteAddr`.
```

```
func Main() {
	for i = 0; i != -1; i++ {
		ent, err = readEntry()
		if err != nil {
			return
		}

		# Check if this is a request or a response
		Request, err = getEntryEnum(ent, "Request")
		if err != nil {
			#Request isn't set, this isn't a modbus packet, skip
			continue
		}

		# read the function value
		Function, err = getEntryEnum(ent, "Function")
		if err != nil {
			continue
		}

		ReqResp, err = getEntryEnum(ent, "ReqResp")
		if err != nil {
			continue
		}

		if Request == true {
			if Function == 0x10 {
				# write multiple registers
				Addr, err = getEntryEnum(ent, "WriteAddr")
				if err != nil {
					writeEntry(ent)
					continue
				}
				Count, err = getEntryEnum(ent, "WriteCount")
				if err != nil {
					writeEntry(ent)
					continue
				}
				if Count == 0 || len(ReqResp) < 5 + 2*Count {
					writeEntry(ent)
					continue
				}
				for j = 0; j < Count; j++ {
					newEnt = cloneEntry(ent)
					# read the register value
					val = (toInt(ReqResp[5+(2*j)]) << 8) | toInt(ReqResp[5+(2*j)+1])
					setEntryEnum(newEnt, "RegAddr", Addr + j)
					setEntryEnum(newEnt, "RegValue", val)
					err = writeEntry(newEnt)
					if err != nil {
						continue
					}
				}
			} else {
				writeEntry(ent)
			}
		}
	}
}
```

The script manages entries explicitly with `readEntry`, `cloneEntry`, and `writeEntry`, which `Process` scripts don't need. It reads and sets enumerated values with `getEntryEnum` and `setEntryEnum` instead of treating them as variables.

## Optional function definitions

The script may also contain functions named `Parse` and `Finalize`.

`Parse` is called before `Process` or `Main` and receives the script arguments as an array. It returns nil on success; any other return value is shown to the user as an error. The `Process` sample script above shows how to parse arguments.

```{attention}
The Parse function MUST explicitly return a value. Returning nil signals a successful parse; returning anything else indicates an error. In case of an error, we recommend returning a string describing the problem.
```

`Finalize` is called after `Process` or `Main` has finished and is the last code the script runs. It is a good place to create resources. If it returns an error, the search reports it.

### Time limits

Scripts that run too long are stopped with an error:

* Top-level script code and `Parse` must each finish within 1 second.
* A single `Process` call is stopped if it runs for more than a few seconds.
* `Main` must return within 10 seconds after its input ends. `writeEntry` returns an error if it can't hand off the entry within 1 second.
* `Finalize` must finish within 10 seconds.

## Utility functions

Anko provides built-in utility functions, listed below in the format `functionName(<functionArgs>) <returnValues>`. These functions can be used in any anko script:

* `getResource(name) []byte, error` returns the contents of the named resource.
* `setResource(name, value) error` writes `value` (a byte slice) to the named resource, creating it if needed.
* `getMacro(text) string, error` expands any macros referenced in `text` and returns the result, so pass the macro with its `$`: `getMacro("$MYMACRO")`. Text with no known macros is returned unchanged, without an error.
* `len(val) int` returns the length of val, which can be a string, slice, etc.
* `toIP(string) IP` converts string to an IP, suitable for comparing against IPs generated by e.g. the packet module.
* `toMAC(string) MAC` converts string to a MAC address.
* `toString(val) string` converts val to a string.
* `toInt(val) int64` converts val to an integer if possible. Returns 0 if no conversion is possible.
* `toFloat(val) float64` converts val to a floating point number if possible. Returns 0.0 if no conversion is possible.
* `toBool(val) bool` attempts to convert val to a boolean. Returns false if no conversion is possible. Positive numbers and the strings `y`, `yes`, `true`, `t`, and `1` (case-insensitive) return true.
* `toHumanSize(val) string` attempts to convert val into an integer, then represent it as a human-readable byte count, e.g. `toHumanSize(15127)` will be converted to "14.77 KB".
* `toHumanCount(val) string` attempts to convert val into an integer, then represent it as a human-friendly number, e.g. `toHumanCount(15127)` will be converted to "15.13 K".
* `typeOf(val) type` returns the type of val as a string, e.g. "string", "bool".
* `toDuration(val) time.Duration` converts a number (nanoseconds) or duration string to a duration.
* `sprintf(format, args...) string` formats a string, like Go's `fmt.Sprintf`.
* `keys(map) []string` returns the keys of a map.
* `defined(name) bool` returns whether a variable or function with that name is defined.
* `producesEnum(name)` tells the pipeline that the script creates an enumerated value with that name. Call it in `Parse`.
* `consumesEnum(name)` tells the pipeline that the script reads an enumerated value with that name. Call it in `Parse`.

The `print`, `println`, and `printf` functions are accepted but print nothing.

The following functions are only available in scripts implementing the `Process` function:

* `setEnum(key, value) error` creates an enumerated value on the current entry named `key` containing `value`.
* `getEnum(key) value, error` returns the enumerated value specified by `key`, or an error if the entry doesn't have it.
* `delEnum(key)` deletes an enumerated value named `key` from the current entry.
* `hasEnum(key) bool` returns whether the current entry has the enumerated value.

The following functions are only available in scripts implementing the `Main` function:

* `readEntry() entry, error` returns the next entry and an error (if any). It will return an error when no entries remain.
* `writeEntry(ent) error` writes out the given entry down the pipeline, returning an error if any.
* `cloneEntry(ent) entry` returns a copy of the specified entry.
* `dupEntryByReference(ent, names...) entry` duplicates the given entry and makes copies of the enumerated values specified by the `names` parameter(s).
* `newEntry() entry` creates an entirely new entry, with the timestamp set to the end of the current query.
* `setEntryEnum(ent, key, value)` sets an enumerated value on the specified entry.
* `getEntryEnum(ent, key) value, error` reads an enumerated value from the specified entry.
* `hasEntryEnum(ent, key) bool` returns whether the entry contains the enumerated value.
* `delEnum(ent, key)` deletes the specified enumerated value from the given entry. (In `Main` scripts this function is named `delEnum`, not `delEntryEnum`.)
* `setEntryData(ent, value)` sets the data portion of an entry.
* `setEntrySrc(ent, ip)` sets the source field of an entry.
* `setEntryTimestamp(ent, time)` sets the timestamp of an entry.

```{note}
`delEnum` takes different arguments in `Process` and `Main` scripts: `delEnum(key)` acts on the current entry, while `delEnum(ent, key)` acts on the given entry.
```

## Built-in variables

The following variables are pre-defined for anko scripts:

* `START`: the start time of the query.
* `END`: the end time of the query.
* `TAGMAP`: a map from tag names to tag numbers. It only contains tags used in the current query; for `tag=default,foo` it holds `default` and `foo`. Use it to set the tag on entries made with `cloneEntry` or `newEntry`.

## Available packages

It is possible to import anko wrappers for certain Go libraries with syntax similar to the following:

```
var json = import("encoding/json")
```

For security reasons, anko scripts can't use every package from the full Anko language. These packages are available:

* [archive/tar](https://golang.org/pkg/archive/tar): limited tar archive functionality
* [bytes](https://golang.org/pkg/bytes): work with byte slices
* [compress/bzip2](https://golang.org/pkg/compress/bzip2), [compress/flate](https://golang.org/pkg/compress/flate), [compress/gzip](https://golang.org/pkg/compress/gzip), [compress/lzw](https://golang.org/pkg/compress/lzw), [compress/zlib](https://golang.org/pkg/compress/zlib): limited compression functionality
* [crypto/md5](https://golang.org/pkg/crypto/md5), [crypto/sha1](https://golang.org/pkg/crypto/sha1), [crypto/sha256](https://golang.org/pkg/crypto/sha256), [crypto/sha512](https://golang.org/pkg/crypto/sha512): cryptographic hashing
* [crypto/tls](https://golang.org/pkg/crypto/tls): limited TLS functionality
* [database/sql](https://golang.org/pkg/database/sql): limited SQL database functionality
* [encoding/base64](https://golang.org/pkg/encoding/base64): limited base64 functionality
* [encoding/csv](https://golang.org/pkg/encoding/csv): encode and decode CSV data
* [encoding/hex](https://golang.org/pkg/encoding/hex): limited hex encoding functionality
* [encoding/json](https://golang.org/pkg/encoding/json): encode and decode json data
* [encoding/xml](https://golang.org/pkg/encoding/xml): limited XML encoding functionality
* [errors](https://golang.org/pkg/errors): handle Go errors
* [flag](https://golang.org/pkg/flag): limited flag parsing functionality
* [fmt](https://golang.org/pkg/fmt): formatting strings
* [google/uuid](https://github.com/google/uuid): generate and inspect UUIDs. Import it as `google/uuid`.
* [github.com/gravwell/ipexist](https://github.com/gravwell/ipexist): Gravwell IP helper functions
* [github.com/RackSec/srslog](https://github.com/RackSec/srslog): alternate syslog package to Go's standard library. Also importable as `github.com/gravwell/srslog`.
* [io](https://golang.org/pkg/io): basic I/O primitives
* [io/ioutil](https://golang.org/pkg/io/ioutil): just the `ioutil.ReadAll` function (see below)
* [math](https://golang.org/pkg/math): mathematical functions
* [math/big](https://golang.org/pkg/math/big): bignums
* [math/rand](https://golang.org/pkg/math/rand): random numbers
* [net](https://golang.org/pkg/net): limited network functionality
* [net/http](https://golang.org/pkg/net/http): limited HTTP functionality (client only)
* [net/url](https://golang.org/pkg/net/url): URLs
* [path](https://golang.org/pkg/path): paths
* [path/filepath](https://golang.org/pkg/path/filepath): path functions specific to files
* [regexp](https://golang.org/pkg/regexp): regular expressions
* [sort](https://golang.org/pkg/sort): sorting
* [strings](https://golang.org/pkg/strings): string processing functions
* [time](https://golang.org/pkg/time): time processing functions
* [github.com/ziutek/telnet](https://github.com/ziutek/telnet): telnet client functions 

An exhaustive description of every package is not possible in this document; you can view the available functions exported for each package at [the official anko repository](https://github.com/mattn/anko/tree/master/packages). Some specific packages are described further below, as they do not offer the complete functionality exported by the official anko repository.

## Package restrictions

Some package functions are too dangerous to expose to scripts, so Gravwell exports only a subset of certain packages. The restrictions are listed below.

### archive/tar

`archive/tar` exports the `tar.Type*` header type constants, the `tar.Err*` error values, and these functions:

- `tar.FileInfoHeader`
- `tar.NewReader`
- `tar.NewWriter`

### compress packages

The compression packages export only their readers, writers, and error values:

- `bzip2.NewReader`
- `flate.NewReader`, `flate.NewReaderDict`, `flate.NewWriter`, `flate.NewWriterDict`
- `gzip.NewReader`, `gzip.NewWriter`, `gzip.ErrChecksum`, `gzip.ErrHeader`
- `lzw.NewReader`, `lzw.NewWriter`
- `zlib.NewReader`, `zlib.NewReaderDict`, `zlib.NewWriter`, `zlib.NewWriterLevelDict`, `zlib.ErrChecksum`, `zlib.ErrDictionary`, `zlib.ErrHeader`

### crypto/md5

`crypto/md5` only exports the "New" and "Sum" functions:

- `md5.New`
- `md5.Sum`

### crypto/sha1

`crypto/sha1` only exports the "New" and "Sum" functions:

- `sha1.New`
- `sha1.Sum`

### crypto/sha256

`crypto/sha256` only exports the various "New" and "Sum" functions:

- `sha256.New`
- `sha256.New224`
- `sha256.Sum224`
- `sha256.Sum256`

### crypto/sha512

`crypto/sha512` only exports the various "New" and "Sum" functions:

- `sha512.New`
- `sha512.New384`
- `sha512.New512_224`
- `sha512.New512_256`
- `sha512.Sum384`
- `sha512.Sum512`
- `sha512.Sum512_224`
- `sha512.Sum512_256`

### crypto/tls

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. `crypto/tls` only exports the TLS config type, for use with the `net/http` package:

- `tls.Config`

### database/sql

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. It exports `sql.Open`, `sql.Drivers`, the `sql.Err*` error values, the `sql.Level*` isolation levels, and the `DB`, `Conn`, `Tx`, `Stmt`, `Row`, `Rows`, `Null*` and related types. It also adds `sql.ExtractRows(rows, n) []string, error`, which scans the current row (after `rows.Next()`) into `n` strings.

### encoding/csv

`encoding/csv` only exports CSV initializers:

- `csv.NewReader` (actually calls `csv.NewReader` with the LazyQuotes option set to true)
- `csv.NewWriter`
- `csv.NewBuilder`

### encoding/base64

`encoding/base64` only exports base64 initializers and encoding types:

- `base64.NewDecoder`
- `base64.NewEncoder`
- `base64.NewEncoding`
- `base64.RawStdEncoding`
- `base64.RawURLEncoding`
- `base64.StdEncoding`
- `base64.URLEncoding`

### encoding/hex

`encoding/hex` exports a subset of initializers and wrappers:

- `hex.Decode`
- `hex.DecodeString`
- `hex.DecodedLen`
- `hex.Dump`
- `hex.Dumper`
- `hex.Encode`
- `hex.EncodeToString`
- `hex.EncodedLen`
- `hex.NewDecoder`
- `hex.NewEncoder`

### encoding/xml

`encoding/xml` exports a subset of initializers, wrappers, and encoding options:

- `xml.Escape`
- `xml.EscapeText`
- `xml.Marshal`
- `xml.MarshalIndent`
- `xml.Unmarshal`
- `xml.NewDecoder`
- `xml.NewTokenDecoder`
- `xml.NewEncoder`
- `xml.HTMLAutoClose`
- `xml.HTMLEntity`
- `xml.Attr`
- `xml.CharData`
- `xml.Comment`
- `xml.Directive`
- `xml.EndElement`
- `xml.Name`
- `xml.ProcInst`
- `xml.StartElement`

### flag

`flag` only exports a subset of types, instead of the entire `flag` package:

- `flag.NewFlagSet`
- `flag.PanicOnError`
- `flag.ContinueOnError`

### fmt

`fmt` exports the string formatting and scanning functions: `fmt.Errorf`, `fmt.Sprint`, `fmt.Sprintf`, `fmt.Sprintln`, `fmt.Sscan`, `fmt.Sscanf`, `fmt.Sscanln`, `fmt.Fprint`, `fmt.Fprintf`, `fmt.Fprintln`, `fmt.Fscan`, `fmt.Fscanf`, and `fmt.Fscanln`. `fmt.Print` and `fmt.Println` exist but print nothing.

### google/uuid

`google/uuid` only exports the "New" and "Parse" functions:

- `uuid.New`
- `uuid.Parse`
- `uuid.ParseBytes`

### github.com/gravwell/ipexist

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. `github.com/gravwell/ipexist` only exports the "New" related functions:

- `ipexist.New`
- `ipexist.NewIPBitMap`

### github.com/RackSec/srslog

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. `github.com/RackSec/srslog` only exposes the syslog related functionality:

- `srslog.Dial`
- `srslog.DefaultFormatter`
- `srslog.DefaultFramer`
- `srslog.RFC3164Formatter`
- `srslog.RFC5424Formatter`
- `srslog.RFC5424CustomFormatter`
- `srslog.RFC5425MessageLengthFramer`
- `srslog.UnixFormatter`
- `srslog.LOG_EMERG`
- `srslog.LOG_ALERT`
- `srslog.LOG_CRIT`
- `srslog.LOG_ERR`
- `srslog.LOG_WARNING`
- `srslog.LOG_NOTICE`
- `srslog.LOG_INFO`
- `srslog.LOG_DEBUG`
- `srslog.LOG_KERN`
- `srslog.LOG_USER`
- `srslog.LOG_MAIL`
- `srslog.LOG_DAEMON`
- `srslog.LOG_AUTH`
- `srslog.LOG_SYSLOG`
- `srslog.LOG_LPR`
- `srslog.LOG_NEWS`
- `srslog.LOG_UUCP`
- `srslog.LOG_CRON`
- `srslog.LOG_AUTHPRIV`
- `srslog.LOG_FTP`
- `srslog.LOG_LOCAL0`
- `srslog.LOG_LOCAL1`
- `srslog.LOG_LOCAL2`
- `srslog.LOG_LOCAL3`
- `srslog.LOG_LOCAL4`
- `srslog.LOG_LOCAL5`
- `srslog.LOG_LOCAL6`
- `srslog.LOG_LOCAL7`

### github.com/ziutek/telnet

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. Exported functions and types include:

- `telnet.Dial`
- `telnet.DialTimeout`
- `telnet.NewConn`
- `telnet.Conn`

### io/ioutil

`io/ioutil` only exports one function, `ioutil.ReadAll()`.

### net

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config. Exported functions and types include:

- `net.CIDRMask`
- `net.Dial`
- `net.DialIP`
- `net.DialTCP`
- `net.DialTimeout`
- `net.DialUDP`
- `net.ErrWriteToConnected`
- `net.FlagBroadcast`
- `net.FlagLoopback`
- `net.FlagMulticast`
- `net.FlagPointToPoint`
- `net.FlagUp`
- `net.IPv4`
- `net.IPv4Mask`
- `net.IPv4allrouter`
- `net.IPv4allsys`
- `net.IPv4bcast`
- `net.IPv4len`
- `net.IPv4zero`
- `net.IPv6interfacelocalallnodes`
- `net.IPv6len`
- `net.IPv6linklocalallnodes`
- `net.IPv6linklocalallrouters`
- `net.IPv6loopback`
- `net.IPv6unspecified`
- `net.IPv6zero`
- `net.InterfaceAddrs`
- `net.InterfaceByIndex`
- `net.InterfaceByName`
- `net.Interfaces`
- `net.JoinHostPort`
- `net.LookupAddr`
- `net.LookupCNAME`
- `net.LookupHost`
- `net.LookupIP`
- `net.LookupMX`
- `net.LookupNS`
- `net.LookupPort`
- `net.LookupSRV`
- `net.LookupTXT`
- `net.ParseCIDR`
- `net.ParseIP`
- `net.ParseMAC`
- `net.ResolveIPAddr`
- `net.ResolveTCPAddr`
- `net.ResolveUDPAddr`
- `net.ResolveUnixAddr`
- `net.SplitHostPort`

### net/http

This package is only available if `Disable-Network-Script-Functions` is set to `false` in the Gravwell config.

`net/http` exports a subset of functions, types, and variables for making HTTP *requests*: the `NewRequest` function, the `DefaultClient` and `DefaultTransport` variables, and the `Client`, `Cookie`, `Request`, `Response`, and `Transport` types. See [the Go documentation](https://golang.org/pkg/net/http/) for details on these types.

`NewRequest(method, url, body) (Request, error)` prepares a request with the given method ("PUT", "POST", "GET", "DELETE") for the URL. `body` is used with PUT and POST requests; set it to `nil` or an `io.Reader`.

Usually you create a request with `NewRequest` and send it with `http.DefaultClient`:

```
req, _ = http.NewRequest("GET", "http://example.org/foo", nil)
resp, _ = http.DefaultClient.Do(req)
resp.Body.Close()
```

Additional headers or cookies can be set on the request before sending:

```
req, _ = http.NewRequest("GET", "http://example.org/foo", nil)
# Add a header
req.Header.Add("My-Header", "gravwell")
# Add a cookie
cookie = make(http.Cookie)
cookie.Name = "foo"
cookie.Value = "bar"
req.AddCookie(&cookie)
resp, _ = http.DefaultClient.Do(req)
resp.Body.Close()
```

```{warning}
You *must* close the response's `Body` when you are finished, as shown above. An open `Body` holds a network connection open, and enough of them will exhaust the available sockets.
```
