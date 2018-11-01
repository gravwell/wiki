# Named Fields

The namedfields module is used to extract and filter data from search entries into enumerated values for later use. Much like the [fields module](#!search/fields/fields.md), it extracts fields from records delimited by sequences of bytes. However, where the fields module uses indexes to refer to elements (e.g. `fields -d "\t" [5] as foo` means "extract the 6th element and call it foo"), namedfields uses specially-formatted [resources](#!resources/resources.md) to give human-friendly names to particular data formats. This is especially useful when attempting to parse things like [Bro](https://www.bro.org) logs or CSV files.

Because so many people use Bro, we have provided a resource file for decoding Bro fields at [https://github.com/gravwell/resources](https://github.com/gravwell/resources), in the `bro/namedfields` subdirectory. Simply upload `namedfields.json` as a resource in Gravwell; the examples in this document assume it was uploaded to a resource named `brofields`

## Key concepts

The namedfields module, at its core, maps numeric indexes within a line to user-friendly names. A set of index-to-name mappings is called a group; a group to parse a textual representation of a network flow might have a mapping like this:

| Index | Name |
|-------|------|
| 0 | start_time |
| 1 | duration |
| 2 | protocol |
| 3 | src_ip |
| 4 | src_port |
| 5 | dst_ip |
| 6 | dst_port |
| 7 | packets |
| 8 | bytes |

One or more groups are then gathered into a Gravwell resource in a format specified elsewhere in this document. When namedfields is run, the user specified which resource to load and which group within that resource should be used to map user-specified names to indexes.

## Supported Options

* `-r <arg>`: The "-r" option is required; it specified the name or GUID of a resource which contains index-to-name mappings.
* `-g <arg>`: The "-g" option is required; it specifies which group to use within the specified resource.
* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.
* `-s` : The “-s” option speciies that the namedfields module operate in a strict mode.  If any filed specification cannot be met, the entry is dropped.  For example if you want the 0th, 1st, and 2nd field but an entry only has 2 fields, the strict flag will cause the entry to be dropped.

## Filtering Operators

The namedfields module allows for a filtering based on equality.  If a filter is enabled that specifies equality ("equal", "not equal", "contains", "not contains") any entry that fails the filter specification will be dropped entirely.  If a field is specified as not equal "!=" and the field does not exist, the field is not extracted but the entry won't be dropped entirely.

| Operator | Name | Description |
|----------|------|-------------|
| == | Equal | Field must be equal
| != | Not equal | Field must not be equal
| ~ | Subset | Field contains the value
| !~ | Not Subset | Field does NOT contain the value

## Examples

Assuming Bro's conn.log file is ingested with the "broconn" flag, the following will extract the "service", "dst", and "resp_bytes" fields from each entry. It will drop all entries whose "service" field does not match the string "dns" and it will rename the extracted "dst" field to "server". It then calculates and graphs the average length of DNS responses for each server. Note that we specify a resource named "brofields" and a group named "Conn", which is defined within the "brofields" resource.

```
tag=broconn namedfields -r brofields -g Conn service==dns dst as server resp_bytes  | mean resp_bytes by server | chart mean by server
```

The following example parses a different Bro file, intel.log. Note that while the resource is the same, we specify a different group:

```
tag=brointel namedfields -r brofields -g Intel source | count source | table source count
```

## Named fields resource format

Before the namedfields module can be used, a resource must be created to map names to indexes within a field. The resource is structured with JSON. Each resource can contain multiple groups, one of which is selected when running the module.

 The example below gives names to entries in Bro's `intel.log` file as well as a custom application log formated as a CSV:

```
{
	"Version": 2,
	"Set": [
		{
			"Delim": "\t",
			"Name": "Intel",
			"Subs": [
				{
					"Name": "source",
					"Index": 0
				},
				{
					"Name": "desc",
					"Index": 1
				},
				{
					"Name": "url",
					"Index": 2
				}
			]
		},{
			"Name": "App",
			"Engine": "csv",
			"Subs": [
				{
					"Name": "user",
					"Index": 0
				},
				{
					"Name": "host",
					"Index": 1
				},
				{
					"Name": "GUID",
					"Index": 2
				}
			]
		}

	]
}
```

Note the essential components:

* `Version` specifies which version of the namedfields module this file is meant for. Leave it as 1.
* `Set` contains an *array* of groups
* This file's Set contains one group, named "Intel". The delimiter is specified as a tab character ("\t"), and a list of `Subs` are provided.
* The "Subs" member defines sub-fields within this group. We see that the field at index 0 is named "source", while index 1 is named "desc" and index 2 is named "url".
* The "Engine" member declares which engine should be used on the group (fields, csv, etc...).

The Gravwell-distributed [namedfields.json](https://github.com/gravwell/resources/blob/master/bro/namedfields/namedfields.json) file for Bro logs contains many groups; refer to it for more examples.

### Namedfields Resource Generation

Gravwell has provided a simple golang library to aid in the generation of namedfields resources.  The library can be used to programatically generate a resource which can the be used by the namedfields module.  The library is available in the tools Gravwell repository on github within the "nfgen" directory.

The simplest usage of the named fields to generate two groups within a single resource would appear as:

```
package main

import (
	"github.com/gravwell/tools/nfgen"
	"log"
)

func main() {
	//create a new named fields resource using the CSV engine that knows how to deal with 2
	//data types, one for login events and one for password failed events
	nf := nfgen.NewGen()
	g, err := nfgen.NewGroup("logins", "csv", ``)
	if err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`username`, ``, 1); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`host`, ``, 2); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`srcip`, ``, 3); err != nil {
		log.Fatal(err)
	}
	if err = nf.AddGroup(g); err != nil {
		log.Fatal(err)
	}
	if g, err = nfgen.NewGroup("failedlogins", "csv", ``); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`srcip`, ``, 2); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`username`, ``, 3); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`password`, ``, 4); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`host`, ``, 5); err != nil {
		log.Fatal(err)
	}
	if err = nf.AddGroup(g); err != nil {
		log.Fatal(err)
	}
	if err = nf.Export("/tmp/lookups.json"); err != nil {
		log.Fatal(err)
	}
}
```

Building and executing the afformentioned generator is simple if you have the golang buildchain installed:

```
go get -u github.com/gravwell/tools/nfgen
go build main.go -o test
./test
```
