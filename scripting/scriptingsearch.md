# Scripting Searches

Gravwell provides a scripting environment in which you can run searches, even launching new searches based on the results of the previous ones, automatically. These scripts can either be run [on a schedule](scheduledsearch.md) or by hand from the [command line client](#!cli/cli.md). Because the CLI allows the script to be re-executed interactively, we recommend developing and testing scripts in the CLI first before creating a scheduled search.

## An example script

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
