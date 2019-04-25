# Scheduling Searches with the Search Agent

It is often advantageous to perform searches automatically, for instance running a search every morning to detect malicious behavior from the previous night. Using Gravwell's search agent, searches can be run on customized scheduled.

The scheduling feature allows the user to schedule both regular searches and [search scripts](scriptingsearch.md).

## Setting up the Search Agent

The Gravwell Search Agent is now included in the main Gravwell install packages and will be installed by default. Disabling the webserver component with the `--no-webserver` flag or setting the `--no-searchagent` flag will disable installation of the Search Agent. The Search Agent is installed automatically by the Gravwell Debian package.

Verify the Search Agent is running with the following command:

```
$ ps aux | grep gravwell_searchagent
```

## Disabling the Search Agent

The Search Agent is installed by default but can be disabled if desired by running the following:

```
systemctl stop gravwell_searchagent.service
systemctl disable gravwell_searchagent.service
```

## Disabling network functions in search agent scripts

By default, scheduled scripts run by the search agent are allowed to use network utilities such as the http library, sftp, and ssh. Setting the option `Disable-Network-Script-Functions=true' in `/opt/gravwell/etc/searchagent.conf` will disable this.