# Scheduling Searches with the Search Agent

It is often advantageous to perform searches automatically, for instance running a search every morning to detect malicious behavior from the previous night. Using Gravwell's search agent, searches can be run on customized scheduled.

The scheduling feature allows the user to schedule both regular searches and [search scripts](scriptingsearch.md).

## Setting up the Search Agent

The Gravwell Search Agent is distributed as a separate installer from the main image, allowing administrators to omit it if desired. We recommend installing it on the same node as the webserver for ease of use.

The Search Agent authenticates with the webserver with a special key set in `gravwell.conf`. Before installing, ensure that `gravwell.conf` contains a `Search-Agent-Auth` line:

```
Search-Agent-Auth=SearchAgentSecret
```

Installing the Search Agent is a simple matter of [downloading the installer](#!quickstart/downloads.md) and running it:

```
$ bash gravwell_searchagent_installer_2.0.6.sh
```

If the webserver is not on the local machine, the installer will prompt for a webserver IP address and a Search-Agent-Auth key.

Once the installer has run, verify that the `gravwell_searchagent` process is running:

```
$ ps aux | grep gravwell_searchagent
```

