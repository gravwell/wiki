# Automation Scripts

Gravwell provides a robust scripting engine in which you can run searches, update resources, send alerts, or take action.  The engine can run searches and examine data automatically, taking action based on search results without the need to involve a human.  

Automation scripts can be run [on a schedule](scheduledsearch) or by hand from the [command line client](/cli/cli). 

## Building Scripts

The Gravwell user interface provides a built-in editor for creating and testing scripts. This interface allows rapid debugging and is the best way to build scripts. It is documented [in the scheduled search/script UI documentation](scheduledsearch).

(scripting_built-in_functions)=
## Built-in functions

Scripts can use built-in functions that mostly match those available for the [anko](scripting) module, with some additions for launching and managing searches. The functions are listed below in the format `functionName(<functionArgs>) <returnValues>`.  These functions are provided as convenience wrappers for specific functionality, however the complete [Gravwell client](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client) is available using the `getClient` wrapper.  The `getClient` wrapper will return a client object that is signed in as the user executing the script. 

## Controlling Versions

Gravwell is constantly adding new modules, methods, and functionality.  It is often desirable to be able to validate that a given script will work with the current version.  This is achieved through two built-in scripting functions which specify the minimum and maximum versions of Gravwell that they are compatible with.  If either assertion fails, the script will fail immediately with an error indicating that the version is incompatible.

* `MinVer(major, minor, point)` Ensures that Gravwell is at least a particular version.
* `MaxVer(major, minor, point)` Ensures that Gravwell is no newer than a particular version.

## Libraries and external functions

Version 3.3.1 of Gravwell now allows automation scripts to include external scripting libraries.  Two functions are provided for including additional libraries:

* `include(path, commitid, repo) error` Includes a library file. The repo and commitid arguments are optional.  If the include fails, the failure reason is returned.
* `require(path, commitid, repo)` Identical behavior to `include`, but if it fails the script is halted and the failure reason is attached to the script's results.

Both `include` and `require` can optionally specify an exact repository or commitid.  If the `repo` argument is omitted the Gravwell default library repo of `https://github.com/gravwell/libs` is used.  If the `commitid` is omitted then the `HEAD` commit is used.  Repos should be accessible by the Gravwell webserver via the schema defined (either `http://`, `https://`, or `git://`) in the repo path.  The scripting system will automatically go get repos as needed: if a commit id is requested that isn't currently known Gravwell will attempt to update the repo.

```
require("email/htmlEmail.ank")
em = htmlEmail
em.SetTitle("Infrastructure Monitoring Alert!")
em.AddSubTitle("Results from Gravwell infrastructure monitoring script")
from = "gravwell@example.com"
to = "recipient@example.com"
subject = "This is the subject line"
em.AddBodyData("This is the body of my message")
em.SendEmail(from, to, subject)
```

```
err = include(`email/htmlEmail.ank`, `df2c1a8792d12be066fec5ea7146ba5325bbaa1d`)
if err != nil {
	return err
}
...
```

If you are in an airgapped system, or otherwise do not want Gravwell to have access to GitHub, you can specify an internal mirror and/or default commit in the `gravwell.conf` file using the `Library-Repository` and `Library-Commit` configuration variables.  For example:

```
Library-Repository="https://github.com/foobar/baz" #override the default library
Library-Commit=da4467eb8fe22b90e5b2e052772832b7de464d63
```

The Library-Repository can also be a local folder that is readable by the Gravwell webserver process.  For example, if you are running Gravwell in a completely airgapped environment, you may still want access to the libs and the ability to update them.  Just unpack the git repository and set `Library-Repository` to that path.

```
Library-Repository="/opt/gitstuff/gravwell/libs"
```

The `include` and `require` can be disabled (thereby disallowing external code) by setting `Disable-Library-Repository` in the `gravwell.conf` file.

## Global configuration

* `loadConfig(resource) (map[string]interface, error)` loads the specified resource and attempts to parse it as a JSON structure, returning the results in a map. If the resource contains `{"foo":"bar","a":1}`, this function will return a map where "foo" → "bar" and "a" → 1.

If you have many automation scripts on the system, you may find it useful to keep a repository of configuration values somewhere for use across all the scripts. Suppose you wanted to use a particular revision of the email alerting library across all your scripts; you could put the following into a resource named "script-config":

```
{
"email_lib_revision":"14ceec90b69943992f4efae8fc9e24c3f4767944"
}
```

and then use the following code in your scripts:

```
cfg, err = loadConfig("script-config")
if err != nil {
	return nil
}
require(`alerts/email.ank`, cfg.email_lib_revision)
```

## Resources and persistent data

* `getResource(name) []byte, error` returns the slice of bytes is the content of the specified resource, while the error is any error encountered while fetching the resource.
* `setResource(name, value) error` creates (if necessary) and updates a resource named `name` with the contents of `value`, returning an error if one arises.
* `setPersistentMap(mapname, key, value)` stores a key-value pair in a map which will persist between executions of a scheduled script.
* `getPersistentMap(mapname, key) value` returns the value associated with the given key from the named persistent map.
* `delPersistentMap(mapname, key)` deletes the specified key/value pair from the given map.
* `persistentMap(mapname)` returns the entire named map, changes to the returned map will persist automatically on successful execution.
* `getMacro(name) string, error` returns the value of the given macro or an error if it does not exist. Note that this function does not perform macro expansion.
* `getSecret(name) string, error` returns the value of the given secret or an error if it does not exist or if the executing user does not have access.  Note that secret values are scrubbed from debug output.

## Search entry manipulation

* `setEntryEnum(ent, key, value)` sets an enumerated value on the specified entry.
* `getEntryEnum(ent, key) value, error` reads an enumerated value from the specified entry.
* `delEntryEnum(ent, key)` deletes the specified enumerated value from the given entry.

## General utilities

* `len(val) int` returns the length of val, which can be a string, slice, etc.
* `toIP(string) IP` converts string to an IP, suitable for comparing against IPs generated by e.g. the packet module.
* `toMAC(string) MAC` converts string to a MAC address.
* `toString(val) string` converts val to a string.
* `toInt(val) int64` converts val to an integer if possible. Returns 0 if no conversion is possible.
* `toFloat(val) float64` converts val to a floating point number if possible. Returns 0.0 if no conversion is possible.
* `toBool(val) bool` attempts to convert val to a boolean. Returns false if no conversion is possible. Non-zero numbers and the strings “y”, “yes”, and “true” will return true.
* `toHumanSize(val) string` attempts to convert val into an integer, then represent it as a human-readable byte count, e.g. `toHumanSize(15127)` will be converted to "14.77 KB".
* `toHumanCount(val) string` attempts to convert val into an integer, then represent it as a human-friendly number, e.g. `toHumanCount(15127)` will be converted to "15.13 K".
* `typeOf(val) type` returns the type of val as a string, e.g. “string”, “bool”.
* `hashItems(val...) (uint64, ok)` hashes one or more items into a uint64 using the siphash algorithm. 'ok' is true if at least one of the items could be hashed. Note that the hash function can really only hash scalars; passing slices or maps will typically not work.

(scripting_search_management)=
## Search management

Due to the way Gravwell's search system works, some of the functions in this section return Search structs (written as `search` in the parameters) while others return search IDs (written as `searchID` in the parameters). Each Search struct contains a search ID which can be accessed as `search.ID`.

Search structs are used to actively read entries from a search, while search IDs tend to refer to inactive searches to which we may attach or otherwise manage.

* `startBackgroundSearch(query, start, end) (search, err)` creates a backgrounded search with the given query string, executed over the time range specified by 'start' and 'end'. The return value is a Search struct. These time values should be specified using the time library; see the examples for a demonstration.
* `startSearch(query, start, end) (search, err)` acts exactly like `startBackgroundSearch`, but does not background the search.
* `detachSearch(search)` detaches the given search (a Search struct). This will allow non-backgrounded searches to be automatically cleaned up and should be called whenever you're done with a search.
* `waitForSearch(search) error` waits for the given search to complete execution and returns an error if there is one.
* `attachSearch(searchID) (search, error)` attaches to the search with the given ID and returns a Search struct which can be used to read entries etc.
* `getSearchStatus(searchID) (string, error)` returns the status of the specified search, e.g. "SAVED".
* `getAvailableEntryCount(search) (uint64, bool, error)` returns the number of entries that can be read from the given search, a boolean specifying if the search is complete, and an error if anything went wrong.
* `getEntries(search, start, end) ([]SearchEntry, error)` pulls the specified entries from the given search. The bounds for `start` and `end` can be found with the `getAvailableEntryCount` function.
* `isSearchFinished(search) (bool, error)` returns true if the given search is complete
* `executeSearch(query, start, end) ([]SearchEntry, error)` starts a search, waits for it to complete, retrieves up to ten thousand entries, detaches from search and returns the entries.
* `deleteSearch(searchID) error` deletes the search with the specified ID
* `backgroundSearch(searchID) error` sends the specified search to the background; this is useful for "keeping" a search for later manual inspection.
* `saveSearch(searchID) error` Marks a given search results for long term storage.  This call does not wait for the query to complete and only returns an error if the request to mark it as saved fails.
* `downloadSearch(searchID, format, start, end) ([]byte, error)` downloads the given search as if a user had clicked the 'Download' button in the web UI. `format` should be a string containing either "json", "csv", "text", "pcap", or "lookupdata" as appropriate. `start` and `end` are time values.
* `getDownloadHandle(searchID, format, start, end) (io.Reader, error)` returns a streaming handle to the results of the given search as if the user had clicked the 'Download' button in the web UI. The handle returned is suitable for use with the HTTP library functions shown later in this document.

### Search Datatype

When executing a search via the startSearch or startBackgroundSearch functions the `search` datatype is returned.  The `search` datatype contains the following members:

* `ID` - A string containing the search ID.  Use this member for other functions like getSearchStatus and attachSearch
* `RenderMod` - A string indicating the renderer attached to the search.  It may be something like raw, text, table, chart, or fdg.
* `SearchString` - A string containing the search string passed in during the request
* `SearchStart` - A string containing the start timestamp for the search
* `SearchEnd` - A string containing the end timestamp for the search
* `Background` - A boolean indicating whether the search was started as a background search
* `Name` - An optional string with a search name.

## Script Information

The following functions let you get information about scheduled scripts/searches and the current script:

* `scheduledSearchInfo() ([]ScheduledSearch, error)` returns all scheduled searches or scripts visible to the user, including the currently-running script. See below for a description of the ScheduledSearch type.
* `thisScriptID() int32` returns the ID number of the currently-running script.

The ScheduledSearch structure contains the following fields:

* `ID` - A 32-bit integer containing the ID number of this scheduled search.
* `Owner` - A 32-bit integer representing the scheduled search's owner. 
* `Groups` - An array of 32-bit group IDs which may view this scheduled search.
* `Name` - A string containing the scheduled search's name.
* `Description` - A string describing the search.
* `Labels` - An array of strings giving further labels to the search.
* `Schedule` - A cron-format string defining the run schedule.
* `Updated` - A timestamp set when the scheduled search was last modified.
* `Disabled` - Boolean, set to true if the search is disabled.
* `SearchString` - If this is a scheduled search (not a script), contains the search string to run.
* `Duration` - Number of seconds into the past to search (only if SearchString is set).
* `Script` - The contents of the script to be run.

* `LastRun` - The time at which the most recent run occurred.
* `LastRunDuration` - The number of nanoseconds the most recent run took.
* `LastSearchIDs` - An array of strings listing the IDs of any searches created during the last run.
* `LastError` - A string containing any error results from the previous run of the script/search.

## Querying Infrastructure Information

The automation scripting system can also be used to monitor the state of the Gravwell installation using API calls.  This allows you to monitor ingester status, system loads, and indexer connectivity within the platform.  The following calls can provide information about the physical deployment:

* `ingesters` - Returns a map containing an ingester status block for each indexer.
* `indexers` - Returns a map containing a well status for each indexer.
* `indexerStates` - Returns a map with a boolean indicating whether the indexer is healthy.
* `systemStates` - Returns a map with disk, CPU, and memory loads for each system.
* `systems` - Returns a map with physical system information such as CPU, memory, disk, and software versions.

## Sending results

The scripting system provides several methods for transmitting script results to external systems.

The following functions provide basic HTTP functionality:

* `httpGet(url) (string, error)` performs an HTTP GET request on the given URL, returning the response body as a string.
* `httpPost(url, contentType, data) (response, error)` performs an HTTP POST request to the given URL with the specified content type (e.g. "application/json") and the given data as the POST body.

More elaborate HTTP operations are possible with the "net/http" library. See the package documentation in the [anko document](scripting) for a description of what is available, or see below for an example.

If the user has configured their personal email settings within Gravwell, the `email` function is a very simple way to send an email:

* `email(from, to, subject, message, attachments...) error` sends an email via SMTP. The `from` field is simply a string, while `to` should be a slice of strings containing email addresses or a single string containing one email address. The `subject` and `message` fields are also strings which should contain the subject line and body of the email. The attachments parameter is optional.
  * Attachments can be sent as a byte array, and they will be given an automatic file name
  * If the attachment parameter is a map, the key is the file name and the value is the attachment
* `emailWithCC(from, to, cc, bcc, subject, message, attachments...) error` sends an email via SMTP. It behaves exactly like the `email` function, but it lets you specify CC and BCC recipients as well. These can be either single strings (`"foo@example.com"`) or arrays of strings (`["foo@example.com", "bar@example.com"]`).

Example sending an email with attachments:
```
#easy way just throwing text into an attachment named "attachment1"
email(`user@example.com`, `bob@accounting.org`, "Hey bob", "We need to talk", "A random attachment")

#Adding a list of attached files with specific names
mp = map[interface]interface{}
mp["stuff.txt"] = "this is some stuff"
mp["things.csv"] = CsvData

subj="Forgot attachments"
body="Forgot to send the stuff and things files"
email(`user@example.com`, `bob@accounting.org`, subj, body, mp)
```

The following functions are deprecated but still available, allowing emails to be sent without configuring the user's email options:

* `sendMail(hostname, port, username, password, from, to, subject, message) error` sends an email via SMTP. `hostname` and `port` specify the SMTP server to use; `username` and `password` are for authentication to the server. The `from` field is simply a string, while the `to` field should be a slice of strings containing email addresses. The `subject` and `message` fields are also strings which should contain the subject line and body of the email.
* `sendMailTLS(hostname, port, username, password, from, to, subject, message, disableValidation) error` sends an email via SMTP using TLS. `hostname` and `port` specify the SMTP server to use; `username` and `password` are for authentication to the server. The `from` field is simply a string, while the `to` field should be a slice of strings containing email addresses. The `subject` and `message` fields are also strings which should contain the subject line and body of the email.  The disableValidation argument is a boolean which disables TLS certificate validation.  Setting disableValidation to true is insecure and may expose the email client to man-in-the-middle attacks.

## Creating Notifications

Scripts may create notifications targeted at the script owner. A notification consists of an integer ID, a string message, an optional HTTP link, and an expiration.

* `addSelfTargetedNotification(uint32, string, string, time.Time) error`

If the expiration is in the past, or more than 24 hours in the future, Gravwell will instead set the expiration to be 12 hours. The notification ID uniquely identifies the notification. This allows the user to update existing notifications by calling the function again with the same notification ID, but it also allows the user to add multiple simultaneous notifications by specifying different IDs.

### Example Notification Creation Script

This script will create a notification that is targeted at the current user, it contains a link and expires 12 hours after the notification is created.

```
var time = import("time")
MSG=`This is my notification`
ID=0x7
LINK="https://gravwell.io"
EXPIRES=time.Now().Add(3*time.Hour)
return addSelfTargetedNotification(ID, MSG, LINK, EXPIRES)
```

## Creating and Ingesting Entries

It is possible to ingest new entries into the indexers from within a script using the following functions:

* `newEntry(time.Time, data) Entry` hands back a new entry with the given timestamp (a time.Time, as from time.Now()) and data (frequently a string).
* `ingestEntries([]Entry, tag) error` ingests the given slice of entries with the specified tag string.

The entries returned by the `getEntries` function can be modified if desired and re-ingested via `ingestEntries`, or new entries can be created wholesale. For example, to re-ingest some entries from a previous search into the tag "newtag":

```
# Get the first 100 entries from the search
ents, _ = getEntries(mySearch, 0, 100)
ingestEntries(ents, "newtag")
```

To ingest new entries based on some other condition:

```
if condition == true {
	ents = make([]Entry)
	ent += newEntry(time.Now(), "Script condition triggered")
	ingestEntries(ents, "results")
}
```

## Other Network Functions

A set of wrapper functions provide access to SSH and SFTP clients. See [the ssh library documentation](https://godoc.org/golang.org/x/crypto/ssh) and [the sftp library documentation](https://godoc.org/github.com/pkg/sftp) for information about the method which can be called on the structures these return.

```{attention}
The clients returned by these functions *must* be closed via their Close() method when you're done using them. See examples below.
```

* `sftpConnectPassword(hostname, username, password, hostkey) (*sftp.Client, error)`establishes an SFTP session on the given ssh server with the specified username and password. If the hostkey parameter is non-nil, it will be used as the expected public key from the host to perform host-key verification. If the hostkey parameter is nil, host key verification will be skipped.
* `sftpConnectKey(hostname, username, privkey, hostkey) (*sftp.Client, error)` establishes an SFTP session on the specified ssh server with the given username, using the provided private key (a string or []byte) to authenticate.
* `sshConnectPassword(hostname, username, password, hostkey) (*ssh.Client, error)` returns an SSH client for the given hostname, authenticating via password. Note that having established a Client, you will typically want to call client.NewSession() to establish a usable session; see the go documentation or the examples below.
* `sshConnectKey(hostname, username, privkey, hostkey) (*sftp.Client, error)` connects to the specified SSH server with the given username, using the provided private key (a string or []byte) to authenticate.

```{note}
The hostkey parameter should be in the known_hosts/authorized_keys format, e.g. "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOcrwoHMonZ/l3OJOGrKYLky2FHKItAmAMPzZUhZEgEb86NNaqfdAj4qmiBDqM04/o7B45mcbjnkTYRuaIUwkno=". To extract the appropriate key from your ~/.ssh/known_hosts, run `ssh-keygen -H -F <hostname>`.
```

A telnet library is also available; no direct wrappers are provided, but it can be used by importing `github.com/ziutek/telnet` in the script and calling telnet.Dial, etc. An example below demonstrates a simple use of the telnet library.

We also provide access to the [github.com/RackSec/srslog](https://github.com/RackSec/srslog) syslog package, which lets scripts send notifications via syslog. An example is shown below.

Finally, the low-level Go [net library](https://golang.org/pkg/net) is available. The listener functions are disabled, but scripts may use the IP parsing functions as well as dial functions such as Dial, DialIP, DialTCP, etc. See below for an example.

### SFTP example

This script connects to an SFTP server using password authentication and no host-key checking. It logs in as the user "sshtest", prints the contents of that user's home directory, and creates a new file named "hello.txt".

```
conn, err = sftpConnectPassword("example.com:22", "sshtest", "foobar", nil)
if err != nil {
	println(err)
	return
}

w = conn.Walk("/home/sshtest")
for w.Step() {
	if w.Err() != nil {
		continue
	}
	println(w.Path())
}

f, err = conn.Create("/home/sshtest/hello.txt")
if err != nil {
	conn.Close()
	println(err)
	return
}
_, err = f.Write("Hello world!")
if err != nil {
	conn.Close()
	println(err)
	return
}

// check it's there
fi, err = conn.Lstat("hello.txt")
if err != nil {
	conn.Close()
	println(err)
	return
}
println(fi)
conn.Close()
```

### SSH example

This script connects to a server using public-key authentication; note that the private key block is shortened for readability here. It also does host-key verification. It then runs `/bin/ps aux` and prints the results.

```
var bytes = import("bytes")

# Get this via `ssh-keygen -H  -F <hostname`
pubkey = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOcrwoHMonZ/l3OJOGrKYLky2FHKItAmAMPzZUhZEgEb86NNaqfdAj4qmiBDqM04/o7B45mcbjnkTYRuaIUwkno="

privkey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAr1vpoiftxU7Jj7P0bJIvgCQLTpM0tMrPmuuwvGMba/YyUO+A
[...]
5iMqZFaUncYZyOFE9hhHqY1xhgwxyjgCTeaI/J/KfbsaCSvrkeBq
-----END RSA PRIVATE KEY-----`

# Log in with a private key
conn, err = sshConnectKey("example.com:22", "sshtest", privkey, pubkey)
if err != nil {
	println(err)
	return
}

session, err = conn.NewSession()
if err != nil {
	println("Failed to create session: ", err)
	return err
}

// Once a Session is created, you can execute a single command on
// the remote side using the Run method.
var b = make(bytes.Buffer)
session.Stdout = &b
err = session.Run("/bin/ps aux")
if err != nil {
	println("Failed to run: " + err.Error())
	return err
}
println(b.String())

session.Close()
conn.Close()
```

### Telnet example

This script connects to a telnet server, logs in as root with the password "testing", and then prints everything it receives, up until a prompt (`$ `).

```
var telnet = import("github.com/ziutek/telnet")
t, err = telnet.Dial("tcp", "example.org:23")
if err != nil {
	println(err)
	return
}
t.SetUnixWriteMode(true)
b, err = t.ReadUntil(":")
println(toString(b))
t.Write("root\n")
b, err = t.ReadUntil(":")
println(toString(b))
t.Write("testing\n")
for {
	r, err = t.ReadUntil("$ ")
	print(toString(r))
}
```

### Syslog example

Use the Dial function to get a connection to a syslog server, then call Alert and other functions to send messages. See [the godoc](https://pkg.go.dev/github.com/RackSec/srslog?tab=doc) for a list of available functions; note that the `NewLogger` and `New` functions are not enabled, because they would write messages to the local system only, which is typically not useful.

```
var syslog = import("github.com/RackSec/srslog")

c, err = syslog.Dial("tcp", "localhost:601", syslog.LOG_ALERT|syslog.LOG_DAEMON, "gravalerts")
if err != nil {
    println(err)
    return err
}
c.Alert("Detected something bad")
c.Close()
```

### Net example

This script connects to localhost at TCP port 7778 and writes "foo" to the connection. You could test this against netcat by running `nc -l 7778`.

```
var net = import("net")

c, err = net.Dial("tcp", "localhost:7778")
if err != nil {
	println(err)
	return err
}

c.Write("foo")
c.Close()
```

(scripting_system_management_functions)=
## Management Functions

The scripting system also has access to management functions that can be used to automatically interact with various Gravwell APIs.

### Performing a system backup

The helper function `backup` is provided to make it easy to perform a system backup and gain a handle on a `io.ReadCloser`.  The `backup` function has the following definition:
```
backup(io.Writer, bool) error
```
The backup function will write the entire backup package to the provided writer and return and error on failure.  The second argument to the backup function indicates whether you want to include saved searches in the backup.  Be aware that saved searches can be very large, which can make backup packages very large.

#### An example backup script

This example backup script will execute a backup and send the resulting package to an TCP network socket, essentially streaming a backup to a remote TCP listener:

```
var net = import("net")
c, err = net.Dial("tcp", "127.0.0.1:9876")
if err != nil {
	return err
}
err = backup(c, true)
if err != nil {
	c.Close()
	return err
}
return c.Close()
```

## An example search script

This script creates a backgrounded search that finds which IPs have communicated with Cloudflare's 1.1.1.1 DNS service over the last day. If no results are found, the search is deleted, but if there are results the search will remain for later perusal by the user in the 'Manage Searches' screen of the GUI.

```
# Import the time library
var time = import("time")
# Define start and end times for the search
start = time.Now().Add(-24 * time.Hour)
end = time.Now()
# Launch the search
s, err = startSearch("tag=netflow netflow Dst==1.1.1.1 Src | unique Src | table Src", start, end)
if err != nil {
	return err
}
printf("s.ID = %v\n", s.ID)
# Wait until the search is finished
for {
	f, err = isSearchFinished(s)
	if err != nil {
		return err
	}
	if f {
		break
	}
	time.Sleep(1 * time.Second)
}
# Find out how many entries were returned
c, _, err = getAvailableEntryCount(s)
if err != nil {
	return err
}
printf("%v entries\n", c)
# If no entries returned, delete the search
# Otherwise, background it
if c == 0 {
	deleteSearch(s.ID)
} else {
	err = backgroundSearch(s.ID)
	if err != nil {
		return err
	}
}
# Always detach from the search at the end of execution
detachSearch(s)
```

To test the script, we can paste the contents into a file, say `/tmp/script`. We then use the Gravwell CLI tool to run the script. It will prompt to re-run as many times as desired and re-read the file for each run; this makes it easy to test changes to the script.

```
$ gravwell -s gravwell.example.org watch script
Username: <user>
Password: <password>
script file path> /tmp/script
0 entries
deleting 910782920
Hit [enter] to re-run, or [q][enter] to cancel

s.ID = 285338682
1 entries
Hit [enter] to re-run, or [q][enter] to cancel
```

The example above shows that the script was run twice; in the first run, no results were found and the search was deleted. In the second, 1 entry was returned, so the search was left intact.

Take care to delete old searches when testing scripts at the CLI; the scheduled search system will automatically delete old searches for you, but not the CLI.

## Using HTTP in scripts

Many modern computer systems use HTTP requests to trigger actions. Gravwell scripts offer the following basic HTTP operations:

* httpGet(url)
* httpPost(url, contentType, data)

We can use these functions and the JSON library included with Anko to modify the previous example script. Rather than backgrounding the search for later perusal, the script will build a list of IPs which have communicated with 1.1.1.1, encode them as JSON, and perform an HTTP POST to submit them to a server:

```
var time = import("time")
var json = import("encoding/json")
var fmt = import("fmt")

# Launch search
start = time.Now().Add(-24 * time.Hour)
end = time.Now()
s, err = startSearch("tag=netflow netflow Dst==1.1.1.1 Src | unique Src", start, end)
if err != nil {
	return err
}
# wait for search to finish
for {
	f, err = isSearchFinished(s)
	if err != nil {
		return err
	}
	if f {
		break
	}
	time.Sleep(1*time.Second)
}
# Get number of entries returned
c, _, err = getAvailableEntryCount(s)
if err != nil {
	return err
}
# Return if no entries
if c == 0 {
	return nil
}
# fetch the entries
ents, err = getEntries(s, 0, c)
if err != nil {
	return err
}
# Build up a list of IPs
ips = []
for i = 0; i < len(ents); i++ {
	src, err = getEntryEnum(ents[i], "Src")
	if err != nil {
		continue
	}
	ips = ips + fmt.Sprintf("%v", src)
}
# Encode IP list as JSON
encoded, err = json.Marshal(ips)
if err != nil {
	return err
}
# Post to an HTTP server
httpPost("http://example.org:3002/", "application/json", encoded)
detachSearch(s)
```

Sometimes, the results of a search may be very large, too large to hold in memory. The "net/http" library, combined with the `getDownloadHandle` function, allows you to stream results directly from the Gravwell search into an HTTP POST/PUT request. It also allows cookies or additional headers to be set:

```
var http = import("net/http")
var time = import("time")
var bytes = import("bytes")

start = time.Now().Add(-72 * time.Hour)
end = time.Now()
s, err = startSearch("tag=gravwell", start, end)
if err != nil {
		return err
}
for {
		f, err = isSearchFinished(s)
		if err != nil {
				return err
		}
		if f {
				break
		}
		time.Sleep(1*time.Second)
}

# Get a handle on the search results
rhandle, err = getDownloadHandle(s.ID, "text", start, end)
if err != nil {
		return err
}
# Build the request
req, err = http.NewRequest("POST", "http://example.org:3002/", rhandle)
if err != nil {
		return err
}
# Add a header
req.Header.Add("My-Header", "gravwell")
# Add a cookie
cookie = make(http.Cookie)
cookie.Name = "foo"
cookie.Value = "bar"
req.AddCookie(&cookie)

# Perform the actual request
resp, err = http.DefaultClient.Do(req)
detachSearch(s)
return err
```

## CSV Helpers

CSV is a pretty common export format for resources and just generally getting data out of Gravwell.  The CSV library provided by `encoding/csv` is robust and flexible but a little verbose.  We have wrapped the CSV writer to provide a simpler interface for use within the Gravwell scripting system.  To create a simplified CSV builder, import the `encoding/csv` package and instead of invoking `NewWriter` call `NewBuilder` without any arguments.

The CSV builder manages its own internal buffers and returns a byte array upon executing `Flush`.  This can simplify the process of building up CSVs for exporting or saving.  Here is an example script that uses the simplified csv Builder to create a resource comprised of two table columns:

```
csv = import("encoding/csv")
time = import("time")

query = `tag=pcap packet ipv4.SrcIP ipv4.DstIP ipv4.Length | sum Length by SrcIP DstIP | table SrcIP DstIP sum`
end = time.Now()
start = end.Add(-1 * time.Hour)

ents, err = executeSearch(query, start, end)
if err != nil {
	return err
}

bldr = csv.NewBuilder()
err = bldr.WriteHeaders([`src`, `dst`, `total`])
if err != nil {
	return err
}

for ent in ents {
	src, err = getEntryEnum(ent, "SrcIP")
	if err != nil {
		return err
	}
	dst, err = getEntryEnum(ent, "DstIP")
	if err != nil {
		return err
	}
	sum, err = getEntryEnum(ent, "sum")
	if err != nil {
		return err
	}
	err = bldr.Write([src, dst, sum])
	if err != nil {
		return err
	}
}

buff, err = bldr.Flush()
if err != nil {
	return err
}
return setResource("csv", buff)
```
(scripting_sql_usage)=
## SQL Usage

The Gravwell scripting system exposes SQL database packages so that automation scripts can interact with external SQL databases.  The SQL library requires Gravwell version 4.1.6 or newer.

The scripting system currently supports the following database drivers:

* MySQL/MariaDB
* PostgreSQL
* MSSQL
* OracleDB

Using the SQL interfaces is done through the `database/sql` package which is a direct import of the Go [database/sql](https://golang.org/pkg/database/sql/) package.

The Gravwell `database/sql` package also includes a helper function that is not part of the Go sql package called `ExtractRows`.  The `ExtractRows` helper function makes it easier to transform the an SQL result row into a slice of strings for later manipulation.  The `ExtractRows` function has the following interface:

`ExtractRows(*sql.Row, columnCount) ([]string, error)`

Using SQL resources in search scripts can be a powerful tool.  However, the SQL interface is verbose and requires some care to use properly.

For in-depth documentation on specific API usage see the official Go documentation.

### SQL Example Script

The following example is a script which queries a remote MariaDB SQL database and turns the result into a CSV resource:

```
var sql = import("database/sql")
var csv = import("encoding/csv")

MinVer(4, 1, 6)

csvbldr = csv.NewBuilder()

// connect to the DB
db, err = sql.Open("mysql", "root:password@tcp(172.19.0.2)/foo")
if err != nil {
	return err
}

// Query example
rows, err = db.Query("SELECT * FROM bar WHERE name!=?", "stuff")
if err != nil {
	return err
}

// Get Headers
headers, err = rows.Columns()
if err != nil {
	return err
}
csvbldr.WriteHeaders(headers)

for rows.Next() {
	vals, err = sql.ExtractRows(rows, len(headers)) //pass in the number of columns
	if err != nil {
		return err
	}
	err = csvbldr.Write(vals)
	if err != nil {
		return err
	}
}

err = rows.Err()
if err != nil {
	return err
}
err = rows.Close()
if err != nil {
	return err
}

err = db.Close()
if err != nil {
	return err
}

data, err = csvbldr.Flush()
if err != nil {
	return err
}

return setResource("foobar", data)
```


## IPExist Datasets

The [ipexist](/search/ipexist/ipexist) search module is designed to test whether an IPv4 address exists in a set, this module is a simple filtering module that is designed for one thing and one thing only: speed.  Under the hood, `ipexist` uses a highly optimized bitmap system so that its possible for a modest machine to represent the entirety of the IPv4 address space in it's filter system.  IPExist is a great tool for holding threat lists and performing initial filtering operations on very large sets of data before performing more expensive lookups using the [iplookup](/search/iplookup/iplookup) module.

The Gravwell scripting system has access to the ipexist builder functions, enabling you to generate high speed ip membership tables from existing data.  The ipexist builder functions are open source and available on [github](https://github.com/gravwell/ipexist).  Below is a basic script which generates an ip membership resource using a query:

```
ipexist = import("github.com/gravwell/ipexist")
bytes = import("bytes")
time = import("time")

query = `tag=ipfix ipfix port==22 src dst | stats count by src dst | table`
end = time.Now()
start = end.Add(-1 * time.Hour)

ipe = ipexist.New()

ents, err = executeSearch(query, start, end)
if err != nil {
	return err
}

for ent in ents {
	ip, err = getEntryEnum(ent, "src")
	if err != nil {
		return err
	}
	err = ipe.AddIP(toIP(ip))
	if err != nil {
		return err
	}
	
	ip, err = getEntryEnum(ent, "dst")
	if err != nil {
		 return err
	}
	err = ipe.AddIP(toIP(ip))
	if err != nil {
		return err
	}
}

bb = bytes.NewBuffer(nil)
err = ipe.Encode(bb)
if err != nil {
	return err
}
ipe.Close()
buff = bb.Bytes()
println("buffer", len(buff))
return setResource("sshusers", buff)
```

(scripting_gravwell_client_usage)=
## Gravwell Client Usage

The `getClient` function will hand back a pointer to a new [Client](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client) object that is logged in and synchronized as the current user. Under normal operating conditions, the new client should be ready for immediate use.  However, it is possible for Gravwell webservers to become unavailable during script operations due to network failures or system upgrades.  We therefore recommend that scripts test the status of the client connection using the [TestLogin()](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client.TestLogin) method.

This example script gets a client, makes a TCP connection to a remote server, and performs a [backup](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client.Backup) of the Gravwell system, sending the backup file out over the remote TCP connection:

```
net = import("net")
time = import("time")

BACKUP_SERVER=`10.0.0.1:5555`

cli = getClient()
if cli == nil {
	return "Failed to get client"
}
err = cli.TestLogin()
if err != nil {
	return err
}
// Backup requests can take some time, increase the client request timeout
err = cli.SetRequestTimeout(10*time.Minute)
if err != nil {
	return err
}

conn, err = net.Dial("tcp", BACKUP_SERVER)
if err != nil {
	return err
}

err = cli.Backup(conn, false)
if err != nil {
	return err
}
err = conn.Close()
if err != nil {
	return err
}

return cli.Close()
``` 

```{note}
The Gravwell client has a default request timeout of 5 seconds. For long running requests like system backups you should increase that timeout, but note that it is best practice to restore the original timeout when you've completed the long-running request; we have omitted that above for brevity.
```
