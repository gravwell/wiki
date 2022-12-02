# Plugin Preprocessor

## Supported Global Options

* `Plugin-Path` (string, required, multiple allowed): Specify the path to the plugin application file.  May specify multiple paths for some engine types.
* `Plugin-Engine` (string, optional): Override the default plugin engine ("scriggo"), currently only `scriggo` is supported.
* `Debug` (Boolean, optional): Enable debug mode, which allows STDOUT debugging by the plugin.  Defaults to false.

Additional plugin specific options can be specified using a key/value pattern with a single name and a string value.  If the same name is specified twice, the plugin will receive the second value.

Plugins are responsible for interpreting additional configuration data.

## Common Use Cases

The variety of desired preprocessors and data manipulators is nearly infinite. Sometimes applying a Turing Complete machine against the data is all that is necessary to make it fit a schema or enrich using some external system.  The plugin system is designed to allow complex operations to be performed against data in a reasonably performant manner.  The current plugin is designed to support multiple engines and each engine may have varying performance, completeness, and/or stability.  The Scriggo engine is the most thoroughly tested, but it does have some known issues.  See the [Scriggo Issue List](https://github.com/open2b/scriggo/issues) for more information.


## Scriggo Engine Plugins

The Scriggo engine is currently the only supported plugin engine, but we have high hopes for the likes of WebAssembly and a few other interpreters so we made sure to scaffold the system so that we can add additional engines at a later date.

### Structure of a Scriggo Plugin

A Scriggo plugin is a fully valid Go Program that uses an injected "gravwell" package which controls execution and provides some needed types.  A plugin must define and register a few functions then turn over execution control to the `gravwell.Execute` function, essentially making your Go program a callback handler, with the meat of the execution happening in the `gravwell` package.  Execution begins in the `Main()` function just like a real Go program.

The `gravwell` package is a sort of virtual package, in that it is not a real package you can go lookup on [pkg.go.dev](https://pkg.go.dev) but instead provides some interface definitions and some scaffolding.

The types provided by the `gravwell` package are:

```
type ConfigMap interface {
        Names() []string
        GetBool(string) (bool, error)
        GetInt(string) (int64, error)
        GetUint(string) (uint64, error)
        GetFloat(string) (float64, error)
        GetString(string) (string, error)
        GetStringSlice(string) ([]string, error)
}

type Tagger interface {
        NegotiateTag(name string) (entry.EntryTag, error)
        LookupTag(entry.EntryTag) (string, bool)
        KnownTags() []string
}

type StartFunc func() error
type CloseFunc func() error
type ConfigFunc func(ConfigMap, Tagger) error
type FlushFunc func() []*entry.Entry
type ProcessFunc func([]*entry.Entry) ([]*entry.Entry, error)

```

The function definition for the `gravwell.Execute` function is:

```
func Execute(string, ConfigFunc, StartFunc, CloseFunc, ProcessFunc, FlushFunc) error
```

The first parameter to the `Execute` function is the name you wish to provide for your plugin. This name will be reported upstream into the Systems & Health page and in logs.  All parameters must be valid functions that fit the type definitions described above.

The call order of each function is as follows:

1. `Config`
2. `Start`
3. `Process`
4. `Flush`
5. `Close`

The `Config` function is used to provide the plugin the opportunity to parse configuration options and provide feedback to the user when starting a service.  If a configuration is invalid, the plugin should return an error indicating why.  During application startup and configuration validation, a plugin will be initialized and the `Config` function will be called so that the plugin can indicate to the user that its configuration is invalid.  A non-nil error returned by the `Config` function is considered fatal and will prevent the ingester from starting up.

The `Start` and `Close` functions provide the plugin the opportunity to do any startup and/or shutdown work.  These functions might be used to establish network connections, open files, close network connections, or clean up any temporary resources the plugin may have created.  The `Start` function will be called prior to any calls to `Process`. `Close` indicates that the ingester is shutting down and the plugin will not receive any more data.

A non-nil error returned by the `Start` function is fatal and will cause the plugin to shutdown and reload.  Plugin writers are encouraged to only return errors on `Start` for truly fatal errors which indicate the plugin could never run; do not return an error on temporary errors like network connectivity problems.

The `Flush` function may be called periodically when the system is under pressure or when the ingester is shutting down.  `Flush` will always be called immediately prior to the `Close` call.


### Skeleton Scriggo Plugin

The following Scriggo plugin program is the bare minimum plugin required to implement all required interfaces.  This example skeleton plugin performs no operations and essentially drops all entries:

```
package main

import (
	"gravwell" //virtual package to expose the builtin plugin funcs

	"github.com/gravwell/gravwell/v3/ingest/entry" //needed for types
)

const (
	PluginName = "example"
)

func Start() error {
	return nil
}

func Close() error {
	return nil
}

func Config(cm gravwell.ConfigMap, tgr gravwell.Tagger) error {
	return nil
}

func Flush() []*entry.Entry {
	return nil
}

func Process(ents []*entry.Entry) ([]*entry.Entry, error) {
	return nil, nil
}

func main() {
	if err := gravwell.Execute(PluginName, Config, Start, Close, Process, Flush); err != nil {
		panic(err) //panic on failure, generally not needed
	}
}
```


### Caveats

The Scriggo engine is **NOT** a complete implementation of the Golang spec, there are limititations and missing features.  Some notable missing features is its lack of method declarations.  While you can execute methods on native types you cannot define methods for your own types.  For a complete list of limitations see the [Scriggo limitations page](https://scriggo.com/limitations).

The plugin preprocessor incurs overhead and may not be as performant as a native preprocessor, in most cases the Gravwell ingest system is fast enough that simple plugins will not adversely affect ingest performance.  However, if you are performing complex operations or attempting to operate on a very high speed ingest pipeline we advise that you enable `Cache-Mode=always` on the ingester.

```{warning}
The Scriggo plugin engine allows for creating goroutines in a plugin. More often than not, this will decrease performance due to nature of the Scriggo interpreter.  Concurrency and synchronization primitives may also behave unexpectedly due to the abstracted runtime.  Be forewarned, a Scriggo plugin is not well suited to fan out and crunch heavy data.
```

## Examples

We have a set of examples used for testing in our open source [Github](https://github.com/gravwell/gravwell/tree/master/ingest/processors/test_data/plugins) repository.

### Example: Forcing all data records to lower case

The following plugin is an example which takes a single configuration parameter and either forces all data records to upper case or lower case based on the configuration parameter.  This plugin is an example of an atomic operation plugin which does not require extensive startup or shutdown logic and does not buffer data at all.  The example configuration snippet is:

```
[Preprocessor "case_adjust"]
	Type=plugin
	Plugin-Path=case_adjust.go
	Upper=true
```

The complete plugin is:

```
package main

import (
	"bytes"
	"errors"
	"fmt"
	"gravwell" //virtual package expose the builtin plugin funcs

	"github.com/gravwell/gravwell/v3/ingest/entry"
)

const (
	PluginName = "recase"
)

var (
	cfg   CaseConfig
	tg    gravwell.Tagger
	ready bool

	ErrNotReady = errors.New("not ready")
)

type CaseConfig struct {
	Upper bool
	Lower bool
}

func nop() error {
	return nil //this is a synchronous plugin, so no "start" or "close"
}

func Config(cm gravwell.ConfigMap, tgr gravwell.Tagger) (err error) {
	if cm == nil || tgr == nil {
		err = errors.New("bad parameters")
	}
	cfg.Upper, _ = cm.GetBool("upper")
	cfg.Lower, _ = cm.GetBool("lower")

	if cfg.Upper && cfg.Lower {
		err = errors.New("upper and lower case are exclusive")
	} else if !cfg.Upper && !cfg.Lower {
		err = errors.New("at least one upper/lower config must be set")
	} else {
		tg = tgr
		ready = true
	}
	return
}

func Flush() []*entry.Entry {
	return nil //we don't hold on to anything
}

func Process(ents []*entry.Entry) ([]*entry.Entry, error) {
	if !ready {
		return nil, ErrNotReady
	}
	if cfg.Upper {
		for i := range ents {
			ents[i].Data = bytes.ToUpper(ents[i].Data)
		}
	} else if cfg.Lower {
		for i := range ents {
			ents[i].Data = bytes.ToLower(ents[i].Data)
		}
	}
	return ents, nil
}

func main() {
	if err := gravwell.Execute(PluginName, Config, nop, nop, Process, Flush); err != nil {
		panic(fmt.Sprintf("Failed to execute dynamic plugin %s - %v\n", PluginName, err))
	}
}
```

### Example: Negotiating tags and Routing Data

This example uses a user provided regex to extract a field from the data and then combine it with the `SRC` value attached to each entry to route the tags.

An example config is:
```
[Preprocessor "tagroute"]
	Type=plugin
	Plugin-Path=tag_route.go
	Regex=`TAG[a-zA-Z0-9]+`
```

The complete plugin is:

```
package main

import (
	"errors"
	"fmt"
	"regexp"
	"gravwell" //virtual package to expose the builtin plugin funcs

	"github.com/gravwell/gravwell/v3/ingest/entry" //needed for types
)

const (
	PluginName = "example"
)

var (
	rx *regexp.Regexp
	tg tagger.Tagger
)

func Start() error {
	return nil
}

func Close() error {
	return nil
}

func Config(cm gravwell.ConfigMap, tgr gravwell.Tagger) error {
	var err error
	rxstr, ok := cm.GetBool("Regex")
	if !ok || rxstr == `` {
		return errors.New("missing Regex parameter")
	} else if rx, err = regexp.Compile(rxstr); err != nil {
		return fmt.Errorf("invalid regex %v", err)
	}
	tg = tgr
	return nil
}

func Flush() []*entry.Entry {
	return nil
}

func Process(ents []*entry.Entry) ([]*entry.Entry, error) {
	for i := range ents {
		//try to match the regex
		if v := rx.Find(ents[i].Data); v != nil {
			//create a combined tag and negotiate it
			ntagstr := fmt.Sprintf("%s_%v", string(v), ents[i].SRC)
			if ntag, err := tg.NegotiateTag(ntagstr); err == nil {
				ents[i].Tag = ntag //assign the tag
			}
		}
	}
	return ents, nil
}

func main() {
	if err := gravwell.Execute(PluginName, Config, Start, Close, Process, Flush); err != nil {
		panic(err) //panic on failure, generally not needed
	}
}
```

## Available Libraries

Plugins may only make use of code that is available in the parent ingester application or is fully self contained in the plugin.  This is due to the way preprocessor plugins are run in an interpreted version of Go, supported by the [Scriggo](https://scriggo.com/) [library](https://github.com/open2b/scriggo/).

Therefore the set of libraries available for import are limited to the following standard library packages:

- archive/tar
- archive/zip
- bufio
- bytes
- compress/bzip2
- compress/flate
- compress/gzip
- compress/lzw
- compress/zlib
- container/heap
- container/list
- container/ring
- context
- crypto
- crypto/aes
- crypto/cipher
- crypto/des
- crypto/dsa
- crypto/ecdsa
- crypto/elliptic
- crypto/hmac
- crypto/md5
- crypto/rand
- crypto/rc4
- crypto/rsa
- crypto/sha1
- crypto/sha256
- crypto/sha512
- crypto/subtle
- crypto/tls
- crypto/x509
- crypto/x509/pkix
- encoding
- encoding/ascii85
- encoding/asn1
- encoding/base32
- encoding/base64
- encoding/binary
- encoding/csv
- encoding/gob
- encoding/hex
- encoding/json
- encoding/pem
- encoding/xml
- errors
- expvar
- flag
- fmt
- go/ast
- go/build
- go/constant
- go/doc
- go/format
- go/importer
- go/parser
- go/printer
- go/scanner
- go/token
- go/types
- hash
- hash/adler32
- hash/crc32
- hash/crc64
- hash/fnv
- hash/maphash
- html
- html/template
- image
- image/color
- image/color/palette
- image/draw
- image/gif
- image/jpeg
- image/png
- index/suffixarray
- io
- io/fs
- io/ioutil
- log
- log/syslog
- math
- math/big
- math/bits
- math/cmplx
- math/rand
- mime
- mime/multipart
- mime/quotedprintable
- net
- net/http
- net/http/cgi
- net/http/cookiejar
- net/http/fcgi
- net/http/httptest
- net/http/httptrace
- net/http/httputil
- net/http/pprof
- net/mail
- net/rpc
- net/rpc/jsonrpc
- net/smtp
- net/textproto
- net/url
- os
- os/exec
- os/user
- path
- path/filepath
- reflect
- regexp
- regexp/syntax
- runtime/debug
- sort
- strconv
- strings
- sync
- sync/atomic
- text/scanner
- text/tabwriter
- text/template
- text/template/parse
- time
- time/tzdata
- unicode
- unicode/utf16
- unicode/utf8

And the following external packages:

- github.com/gravwell/gravwell/v3/ingest
- github.com/gravwell/gravwell/v3/ingest/config
- github.com/gravwell/gravwell/v3/ingest/entry
- github.com/crewjam/rfc5424
- github.com/dchest/safefile
- github.com/gobwas/glob
- github.com/gofrs/flock
- github.com/google/gopacket
- github.com/google/renameio
- github.com/google/uuid
- github.com/gravwell/ipfix
- github.com/h2non/filetype
- github.com/k-sone/ipmigo
- github.com/klauspost/compress
- github.com/open-networks/go-msgraph
- github.com/tealeg/xlsx
- github.com/miekg/dns
- github.com/buger/jsonparser

```{note}
If you want an additional library, file a Feature Request issue on our [Github Repo](https://github.com/gravwell/gravwell/issues).
```
