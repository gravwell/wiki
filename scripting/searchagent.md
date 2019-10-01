# The Search Agent

The aearch agent is the component which runs [automated searches](scheduledsearch.md). The aearch agent is included in the main Gravwell install packages and will be installed by default. Disabling the webserver component with the `--no-webserver` flag or setting the `--no-searchagent` flag will disable installation of the aearch agent. The aearch agent is installed automatically by the Gravwell Debian package.

You can verify the aearch agent is running with the following command:

```
$ ps aux | grep gravwell_searchagent
```

## Disabling the search agent

The aearch agent is installed by default but can be disabled if desired by running the following:

```
systemctl stop gravwell_searchagent.service
systemctl disable gravwell_searchagent.service
```

## Configuring the search agent

The aearch agent is configured in `/opt/gravwell/etc/searchagent.conf`. An example configuration is shown below:

```
[global]
Webserver-Address=127.0.0.1:80
Insecure-Skip-TLS-Verify=true
Insecure-Use-HTTP=true
Search-Agent-Auth=SearchAgentSecrets
Max-Script-Run-Time=10	# Minutes, set to 0 for unlimited
Log-File=/opt/gravwell/log/searchagent.log
Log-Level=INFO
```

This configuration is suitable when running the search agent on the same node as the webserver, provided the webserver is configured to use HTTP rather than HTTPS. Note that the webserver is located on the loopback interface (127.0.0.1) and that HTTP is explicitly enabled.

The individual configuration options available for the Search Agent configuration file are described below.

**Webserver-Address**

The `Webserver-Address` option gives an IP address or hostname, plus a port, which the search agent should use to connect to a webserver. This option can be specified multiple times; if multiple webservers are defined (as shown below), the search agent will load-balance its searches across them.

```
Webserver-Address=gravwell1.example.org:443
Webserver-Address=gravwell2.example.org:443
```

Attention: Do not specify multiple webservers unless they are all synchronized using the [datastore](#!distributed/frontend.md)

**Search-Agent-Auth**

The `Search-Agent-Auth` parameter sets the authentication token used to authenticate with the webserver. This should be set automatically during the installation process. It *must* match the `Search-Agent-Auth` value found in `/opt/gravwell/etc/gravwell.conf` on the target webserver!

**Insecure-Skip-TLS-Verify**

If `Insecure-Skip-TLS-Verify` is set to true, the search agent will *not* verify the validity of TLS certificates when connecting to an HTTPS-enabled Gravwell webserver. Use this option with care and see [the certificates documentation](#!configuration/certificates.md) for more information.

**Insecure-Use-HTTP**

If `Insecure-Use-HTTP` is set to true, the search agent will attempt to communicate with the Gravwell webserver using plaintext HTTP rather than the default HTTPS. This option is set to true in the default configuration file because [Gravwell requires manual configuration to enable HTTPS](#!configuration/certificates.md)

**Disable-Network-Script-Functions**

By default, scheduled scripts run by the search agent are allowed to use network utilities such as the http library, sftp, and ssh. Setting the option `Disable-Network-Script-Functions=true' will disable this.

**HTTP-Proxy**

The `HTTP-Proxy` parameter allows you to define an HTTP proxy which will be used *by scheduled scripts*. Thus if you set `HTTP-Proxy=https://proxy.example.com:3128`, any HTTP requests originating in scheduled scripts will be routed through this proxy.

**Max-Script-Run-Time**

The `Max-Script-Run-Time` parameter sets, in minutes, the maximum length of wall-clock time a scheduled script may run. If a script goes over the limit, it is immediately terminated. Setting this parameter to 0 gives scripts unlimited time, but we recommend setting *some* maximum time. The default configuration file sets a maximum time of 10 minutes, which is suitable for many purposes.

**Log-File**

The `Log-File` parameter tells the search agent where it should output its logs.

**Log-Level**

The `Log-Level` parameter tells the search agent the minimum level of severity which should be logged. The options are INFO, WARN, ERROR, or OFF. Selecting WARN means that logs of severity WARN or ERROR will be logged. Selecting INFO logs everything. Selecting OFF logs nothing.
