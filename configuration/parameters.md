# Configuration Parameters

The `gravwell.conf` configuration file is used by indexers, webservers, and the datastore for configuration. The configuration contains *sections*, which are defined inside square brackets (e.g. `[Global]`); each section contains parameters. The following example includes a Global section, a Replication section, a Default-Well section, and a Storage-Well section:

```
[global]
### Authentication tokens
Ingest-Auth=IngestSecrets
Control-Auth=ControlSecrets
Search-Agent-Auth=SearchAgentSecrets

### Web server HTTP/HTTPS settings
Web-Port=80
Insecure-Disable-HTTPS=true

### Other web server settings
Remote-Indexers=net:127.0.0.1:9404
Persist-Web-Logins=True
Session-Timeout-Minutes=1440
Login-Fail-Lock-Count=4
Login-Fail-Lock-Duration=5

### Ingester settings
Ingest-Port=4023
Control-Port=9404
Search-Pipeline-Buffer-Size=4

### Other settings
Log-Level=WARN

### Paths
Pipe-Ingest-Path=/opt/gravwell/comms/pipe
Log-Location=/opt/gravwell/log
Web-Log-Location=/opt/gravwell/log/web
Render-Store=/opt/gravwell/render
Saved-Store=/opt/gravwell/saved
Search-Scratch=/opt/gravwell/scratch
Web-Files-Path=/opt/gravwell/www
License-Location=/opt/gravwell/etc/license
User-DB-Path=/opt/gravwell/etc/users.db
Web-Store-Path=/opt/gravwell/etc/webstore.db

[Replication]
	Peer=10.0.0.1
	Storage-Location=/opt/gravwell/replication_storage
	Insecure-Skip-TLS-Verify=true
	Disable-TLS=true
	Connect-Wait-Timeout=20

[Default-Well]
	Location=/opt/gravwell/storage/default/
	Accelerator-Name=fulltext #fulltext is the most resilent to varying data types
	Accelerator-Engine-Override=bloom #The bloom engine is effective and fast with minimal disk overhead
	Disable-Compression=true

[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Tags=syslog
	Tags=kernel
	Tags=dmesg
	Accelerator-Name=fulltext #fulltext is the most resilent to varying data types
	Accelerator-Args="-ignoreTS" #tell the fulltext accelerator to not index timestamps, syslog entries are easy to ID
```

Each parameter has a default value which applies if the parameters is empty, or not specified in the configuration file.

## Global Parameters

The parameters in this section go under the `[Global]` heading of the config file, which is typically at the top of the file.

**Indexer-UUID**
Applies to:			Indexer
Default Value:		[randomly generated if not set]
Example:			`Indexer-UUID="ecdeeff8-8382-48f1-a24f-79f83af00e97"`
Description:		Sets a unique identifier for a particular indexer. No two indexers should have the same UUID. If this parameter is not set, the indexer will generate a UUID and write it into the config file; this is usually the best choice, unless an indexer has catastrophically failed and is to be rebuilt from replication (see [the replication docs](configuration/replication.md)). Think twice before modifying this parameter.

**Webserver-UUID**
Applies to:			Webserver
Default Value:		[randomly generated if not set]
Example:			`Webserver-UUID="ecdeeff8-8382-48f1-a24f-79f83af00e97"`
Description:		Sets a unique identifier for a particular webserver. No two webserver should have the same UUID. If this parameter is not set, the webserver will generate a UUID and *write it back into the config file*; this is usually the best choice. Think twice before modifying this parameter.

**License-Location**
Applies to:        Indexer and Webserver
Default Value:        `/opt/gravwell/etc/license`
Example:        `License-Location=/opt/gravwell/etc/my_license`
Description:        Sets the path to the gravwell license file, the path must be readable by the gravwell user and group.

**Config-Location**
Applies to:        Indexer and Webserver
Default Value:        `/opt/gravwell/etc`
Example:        `Config-Location=/tmp/path/to/etc`
Description:        The config location allows for specifying an alternate location for housing all other configuration parameters.  Specifying an alternate Config-Location allows for setting a single parameter without requiring that all other parameters be specified with the alternate path.

**Web-Port**
Applies to:        Webserver
Default Value:        `443`
Example:        `Web-Port=80`
Description:        Specifies the listening port for the webserver. Note that setting the Web-Port parameter to 80 will not switch the webserver to HTTP-only mode; that requires the Insecure-Disable-HTTPS setting.

**Disable-HTTP-Redirector**
Applies to:        Webserver
Default Value:        `False`
Example:        `Disable-HTTP-Redirector=true`
Description:        By default Gravwell starts an HTTP redirector which redirects clients requesting the cleartext HTTP portal to the encrypted HTTPS portal.

**Insecure-Disable-HTTPS**
Applies to:		Webserver
Default Value:	`false`
Example:		`Insecure-Disable-HTTPS=true`
Description:	By default Gravwell operates in HTTPS mode. Setting `Insecure-Disable-HTTPS=true` instructs Gravwell to instead use plaintext HTTP, listening on `Web-Port`.

**Webserver-Domain**
Applies to:		Webserver
Default Value:	0
Example:		`Webserver-Domain=17`
Description: The `Webserver-Domain` parameter controls the [resources](#!resources/resources.md) domain on the webserver. If two webservers are configured with the same domain, have differing resource sets, are connected to the same indexer(s), and are *not* synchronized via the datastore, the indexer(s) will thrash between the two sets of resources. Putting the webservers in different domains allows them both to use the same indexer without resource conflicts.

**Control-Listen-Address**
Applies to:        Indexer
Default Value:
Example:        `Control-Listen-Address=”10.0.0.1”`
Description:        The Control-Listen-Address parameter can bind the indexer's control listener to a specific address.  Gravwell installations on dual-home machines, or machines with high speed data networks and low speed control networks, may wish to bind indexers to specific addresses to ensure the traffic is routed appropriately.

**Control-Port**
Applies to:        Indexer and Webserver
Default Value:        `9404`
Example:        `Control-Port=12345`
Description:        The Control-Port parameter selects the port on which the indexer should listen for control commands from the webserver. This setting must be the same on indexers and webservers in order for them to communicate. The installer does not set the bind capability on the indexers by default, so ports must be set to a value greater than 1024.  Adjusting the control port may be necessary in environments where multiple indexers are run on a single machine, or where another application is binding to port 9404.

**Datastore-Listen-Address**
Applies to:			Datastore
Default Value:
Example:			`Datastore-Listen-Address=10.0.0.1`
Description:		The Datastore-Listen-Address parameter instructs the datastore to listen only on a particular address. By default, the datastore listens on all addresses on the system.

**Datastore-Port**
Applies to:			Datastore
Default Value:		`9405`
Example:			`Datastore-Port=7888`
Description:		The Datastore-Port parameter selects the port on which the datastore will communicate. The port should be greater than 1024. The default value of 9405 is typically suitable for most installations.

**Datastore**
Applies to:			Webserver
Default Value:
Example:			`Datastore=10.0.0.1:9405`
Description:		The Datastore parameter specifies that the webserver should connect to a datastore to synchronize its dashboards, resources, user preferences, and search history. This allows for [distributed webservers](distributed/frontend.md) but should only be set if needed. By default, webservers do not connect to a datastore.

**Datastore-Update-Interval**
Applies to:			Webserver
Default Value:		`10`
Example:			`Datastore-Update-Interval=5`
Description:		The Datastore-Update-Interval parameter determines how long (in seconds) the webserver should wait before checking the datastore for updates. The default value of 10 seconds is typically suitable.

**Datastore-Insecure-Disable-TLS**
Applies to:		Webserver and Datastore
Default Value:	`false`
Example:		`Datastore-Insecure-Disable-TLS=true`
Description:	The Datastore-Insecure-Disable-TLS parameter is used by both the webserver and the datastore. By default, the datastore listens for incoming HTTPS connections from webservers; setting this parameter to false makes the datastore expect plaintext HTTP and instructs the webservers to use HTTP.

**Datastore-Insecure-Skip-TLS-Verify**
Applies to:		Webserver
Default Value:	`false`
Example:		`Datastore-Insecure-Skip-TLS-Verify=true`
Description:	The Datastore-Insecure-Skip-TLS-Verify parameter instructs the webserver to ignore invalid TLS certificates when connecting to the datastore. This is necessary when using self-signed certificates but should be avoided when possible.

**External-Addr**
Applies to:		Webserver
Default Value:
Example:		`External-Addr=10.0.0.1:443`
Description:	The External-Addr parameter specifies the address other webservers should use to contact this webserver. This parameter is **required** when using a datastore, as it allows a user on one webserver to load the results of a search performed on another webserver.

**Search-Forwarding-Insecure-Skip-TLS-Verify**
Applies to:		Webserver
Default Value:	`false`
Example:		`Search-Forwarding-Insecure-Skip-TLS-Verify=true`
Description:	This parameter is only useful when operating multiple webservers in distributed mode using a datastore. If the webservers have self-signed certificates, users will be unable to access searches from remote webservers *unless* this parameter is set to true.

**Ingest-Port**
Applies to:        Indexer
Default Value:        `4023`
Example:        `Ingest-Port=14023`
Description:        The Ingest-Port parameter controls the port the indexers listen on for ingester connections.  Altering the Ingest-Port parameter can be useful when running multiple indexers on a single machine or another application is already bound to the default port of 4023.

**TLS-Ingest-Port**
Applies to:        Indexer
Default Value:        `4024`
Example:        `TLS-Ingest-Port=14024`
Description:        The TLS-Ingest-Port parameter controls the port the indexers listen on for ingester connections.  Altering the TLS-Ingest-Port parameter can be useful when running multiple indexers on a single machine or another application is already bound to the default port of 4024.  By default, all ingesters using the TLS transport will validate remote certificates.  If a deployment is using the auto generated certificates, ingesters either need to have the certificates installed as trusted, or they must disable certificate validation (this effectively destroys the protections provided by a TLS transport).

**Pipe-Ingest-Path**
Applies to:			Indexer
Default Value:		`/opt/gravwell/comms/pipe`
Example:			`Pipe-Ingest-Path=/tmp/path/to/pipe`
Description:		The Pipe-Ingest-Path specifies the full path to a Unix named pipe.  The indexer will create the named pipe and co-resident ingesters can attach to the pipe and use it as a very high speed and low latency transport.  Named pipes are excellent for ingesters that require extremely high performance, such as a network packet ingester operating above 1 gigabit.  Named pipes can also be used to facilitate transport over unusual network transports or very high speed non-IP based interconnects.

**Log-Location**
Applies to:        Indexer
Default Value:        `/opt/gravwell/log`
Example:        `Log-Location=/tmp/path/to/logs`
Description:        The Log-Location parameter controls the location that Gravwell infrastructure places its own logs.  Gravwell does not feed its own logs directly into indexers, and instead writes them to files (use the file follower ingester if you want to ingest Gravwell logs too).  This parameter specifies where those logs go.

**Web-Log-Location**
Applies to:        Webserver
Default Value:        `/opt/gravwell/log/web`
Example:        `Web-Log-Location=/tmp/path/to/logs/web`
Description:        The Web-Log-Location parameter controls where webserver logs are stored.  Gravwell does not feed its own logs directly into indexers, and instead writes them to files (use the file follower ingester if you want to ingest Gravwell logs too).  This parameter specifies where those logs go.

**Datastore-Log-Location**
Applies to:		Datastore
Default Value:	`/opt/gravwell/log/datastore`
Example:		`Datastore-Log-Location=/tmp/path/to/logs/datastore`
Description:	The Datastore-Log-Location parameter controls where datastore logs are stored.

**Log-Level**
Applies to:        Indexer, Datastore, and Webserver
Default Value:        `INFO`
Example:        `Log-Level=ERROR`
Description:        The Log-Level parameter controls the verbosity of logs from gravwell infrastructure.  There are three available arguments to the Log-Level: INFO, WARN, and ERROR.  INFO is the most verbose, and ERROR is the least.  The logging system will generate a file for each level of logging and rotate them in a similar manner to the syslog daemon.

**Disable-Access-Log**
Applies to:        Webserver
Default Value:        `false`
Example:        `Disable-Access-Log=true`
Description:        The Disable-Access-Log parameter is used to disable the access log generated by the webserver.  The access logging infrastructure logs individual page accesses; while is typically valuable to have these access logs to audit Gravwell access and to debug potential problems, the access logs can become large in environments with a lot of users, so it may be desirable to disable them.

**Persist-Web-Logins**
Applies to:        Webserver
Default Value:        `true`
Example:        `Persist-Web-Logins=false`
Description:        The Persist-Web-Logins parameter is used to inform the webserver that it should save user sessions on shutdown to non-volatile storage.  By default, if the webserver is shutdown or restarted, it will persist client sessions.  Setting the Persist-Web-Logins to false means sessions will be invalidated whenever the webserver is restarted.

**Session-Timeout-Minutes**
Applies to:        Webserver
Default Value:        `60`
Example:        `Session-Timeout-Minutes=1440`
Description:        The Session-Timeout-Minutes parameter controls how long a client can be idle before the webserver destroys the session.  For example, if a client closes a browser without logging out, the system will wait for the specified time period before invalidating the session.  The installers set this value to 1 day by default.

**Key-File**
Applies to:        Indexer, Datastore, and Webserver
Default Value:        `/opt/gravwell/etc/key.pem`
Example:        `Key-File=/opt/gravwell/etc/privkey.pem`
Description:        The Key-File parameter controls which file is used as a private key for the webserver, datastore, and indexer.  The private/public keys must be encoded in the PEM format.  The private key must be protected, and should be destroyed and reissued upon compromise.  For more information see http://www.tldp.org/HOWTO/SSL-Certificates-HOWTO/x64.html.

**Certificate-File**
Applies to:        Indexer, Datastore, and Webserver
Default Value:        `/opt/gravwell/etc/cert.pem`
Example:        `Certificate-File=/opt/gravwell/etc/cert.pem`
Description:        The Certificate-File parameter specifies the public key component of the public/private key pair used for TLS transport.  The public key will be delivered to every ingester and web client and is not considered sensitive.  Gravwell expects the public key to be encoded in the PEM format, and to only contain the public key portion.

**Ingest-Auth**
Applies to:        Indexer
Default Value:        `IngestSecrets`
Example:        `Ingest-Auth=abcdefghijklmnopqrstuvwxyzABCD`
Description:        The Ingest-Auth parameter specifies the shared secret token that is used to authenticate ingesters to indexers.  This token can be of arbitrary length; Gravwell recommends a high entropy token of at least 24 characters.  By default the installers will generate a random token.

**Control-Auth**
Applies to:        Indexer and Webserver
Default Value:        `ControlSecrets`
Example:        `Control-Auth=abcdefghijklmnopqrstuvwxyzABCD`
Description:        The Control-Auth parameter specifies the shared secret token that is used to authenticate ingesters to webservers and vice versa.  This token can be of arbitrary length; Gravwell recommends a high entropy token of at least 24 characters.  By default the installers will generate a random token.

**Search-Agent-Auth**
Applies to:		Webserver
Default Value:	
Example:		`Search-Agent-Auth=abcdefghijklmnopqrstuvwxyzABCD`
Description:	The Search-Agent-Auth parameter specifies the shared secret token that is used to authenticate the search agent to the webserver. The installers default to generating a random search agent token.

**Web-Files-Path**
Applies to:        Webserver
Default Value:        `/opt/gravwell/www`
Example:        `Web-Files-Path=/tmp/path/to/www`
Description:        The Web-Files-Path specifies the path containing the frontend GUI files to be served by the webserver.  The web files contain all Gravwell code responsible for displaying the webpage and interacting with the Gravwell system via a web browser.

**Tag-DB-Path**
Applies to:		Indexer
Default Value:	`/opt/gravwell/etc/tags.db`
Example:		`Tag-DB-Path=/tmp/path/to/tags.db`
Description:	The Tag-DB-Path parameter specifies the location of the tag database. This file maps the indexer's numeric tag IDs to tag name strings.

**User-DB-Path**
Applies to:        Webserver
Default Value:        `/opt/gravwell/etc/users.db`
Example:        `User-DB-Path=/tmp/path/to/users.db`
Description:        The User-DB-Path parameter specifies the location of the user database file.  The user database file contains user and group configurations.  The user database uses the bcrypt hash algorithm to store and validate passwords, which is considered very robust, but the users.db file should still be protected.  By default the installers set the filesystem permissions on the user database file to only be readable by the Gravwell user and group.

**Datastore-User-DB-Path**
Applies to:		Datastore
Default Value:	`/opt/gravwell/etc/datastore-users.db`
Example:		`Datastore-User-DB-Path=/tmp/path/to/datastore-users.db`
Description:	The Datastore-User-DB-Path parameter specifies the location of the user database file as managed by the datastore component. This **must not** be the same path as specified by the User-DB-Path parameter!

**Web-Store-Path**
Applies to:        Webserver
Default Value:        `/opt/gravwell/etc/webstore.db`
Example:        `Web-Store-Path=/tmp/path/to/webstore.db`
Description:        The Web-Store-Path points to the database file used to store search history, dashboards, user settings, user sessions, and any other miscellaneous user data.  The webstore database file does not contain any user credentials, but *does* contain user session cookies and CSRF tokens.  Gravwell ties cookies and CSRF tokens to origins, so while the risk of an attacker reusing as stolen cookie or token is low the datastore should be protected.  Installers set the filesystem permissions to only allow read/write by the Gravwell user.

**Datastore-Web-Store-Path**
Applies to:		Datastore
Default Value:	`/opt/gravwell/etc/datastore-webstore.db`
Example:		`Datastore-Web-Store-Path=/tmp/path/to/datastore-webstore.db`
Description:	The Datastore-Web-Store-Path parameter points to the database file used by the datastore to store search history, dashboards, and user preferences. This **must not** be the same path as specified by the Web-Store-Path parameter!

**Web-Listen-Address**
Applies to:        Webserver
Default Value:
Example:        `Web-Listen-Address=10.0.0.1`
Description:        The Web-Listen-Address parameter specifies the address the webserver should bind to and serve from.  By default the parameter is empty, meaning the webserver binds to all interfaces and addresses.

**Login-Fail-Lock-Count**
Applies to:        Webserver
Default Value:        `5`
Example:        `Login-Fail-Lock-Count=10`
Description:        The Login-Fail-Lock-Count parameter specifies the number of sequential failed logins against a user account can occur before brute-force protection is enabled on the account.  For example, if the value is set to 4 and a user provides a bad password 4 times in a row, additional login attempts will take longer to complete, slowing down an attacker. Note: Gravwell previously locked an account after a specific number of failures; it now engages a less aggressive brute-force protection, but for legacy reasons the configuration parameter retains the 'Lock' name.

**Login-Fail-Lock-Duration**
Applies to:        Webserver
Default Value:        `5`
Example:        `Login-Fail-Lock-Duration=10`
Description:        The Login-Fail-Lock-Duration parameter specifies the window (in minutes) used when calculating if the Login-Fail-Lock-Count has been exceeded. Note: Gravwell previously locked an account after a specific number of failures; it now engages a less aggressive brute-force protection, but for legacy reasons the configuration parameter retains the 'Lock' name.

**Remote-Indexers**
Applies to:        Webserver
Default Value:        `net:10.0.0.1:9404`
Example:        `Remote-Indexers=net:10.0.0.1:9404`
Description:        The Remote-Indexers parameter specifies the address and port of remote indexers that the webserver should connect to and control.  Remote-Indexers is a list parameter, meaning that it can be specified many times to provide multiple remote indexers. Gravwell Cluster editions will need to specify each indexer in the cluster.  The “net:” prefix indicates that the remote indexer is accessible via a network transport; special editions of Gravwell can use alternate transports, but most commercial customers should expect to use “net:”.

**Search-Scratch**
Applies to:        Indexer and Webserver
Default Value:        `/opt/gravwell/scratch`
Example:        `Search-Scratch=/tmp/path/to/scratch`
Description:        The Search-Scratch parameter specifies a storage location that search modules can use for temporary storage during an active search.  Some search modules may need to use temporary storage due to memory constraints.  For example, the sort module may need to sort 5GB of data but the physical machine may only have 4GB of physical RAM.  The module can intelligently use the scratch space to sort the large dataset without invoking the host's swap (which would penalize all modules, not just sort).  At the end of each search, scratch space is destroyed.

**Render-Store**
Applies to:        Webserver
Default Value:        `/opt/gravwell/render`
Example:        `Render-Store=/tmp/path/to/render`
Description:        The Render-Store parameter specifies where renderer modules store the results of a search.  Render-Store locations are temporary storage locations and typically represent reasonably small data sets.  When a search is actively running or dormant and interacting with a client, the Render-Store is where the renderer will store and retrieve its data set.  Render-Store should be on high speed storage such as flash-based or XPoint SSDs.  When a search is abandoned the Render-Store is deleted (unless the search is saved).

**Saved-Store**
Applies to:        Webserver
Default Value:        `/opt/gravwell/saved`
Example:        `Saved-Store=/path/to/saved/searches`
Description:        The Saved-Store parameter specifies where saved searches will be stored.  Saved searches represent the output state of a search and can be useful for auditing and situations where users want to be able to consult search results again later without relaunching the search.  Saved searches must be explicitly deleted and the data is not subject to shard age out policies.  Saved searches are entirely atomic, which means that the underlying data for a saved search can be completely aged out and even deleted and users can still re-open and examine the saved search.  Saved searches can also be shared, meaning users can pack up and share saved searches with other instances of Gravwell.

**Search-Pipeline-Buffer-Size**
Applies to:        Indexer and Webserver
Default Value:        `2`
Example:        `Search-Pipeline-Buffer-Size=8`
Description:        The Search-Pipeline-Buffer-Size specifies how many blocks can be in transit between each module during a search.  Larger sizes allow for better buffering and potentially higher throughput searches at the expense of resident memory usage.  Indexers are more sensitive to the pipeline size, but also use a shared memory technique whereby the system can evict and reinstantiate memory at will; the webserver typically keeps all entries resident when moving through the pipeline and relies on condensing modules to reduce the memory load.  If your system uses higher latency storage systems like spinning disks, it can be advantageous to increase this buffer size.
Increasing this parameter may make searches perform better, but it will directly impact the number of running searches the system can handle at once!  If you know you are storing extremely large entries like video frames, PE executables, or audio files you may need to reduce the buffer size to limit resident memory usage. If you see your host kernel invoking the Out Of Memory (OOM) firing and killing the Gravwell process, this is the first knob to turn.

**Search-Relay-Buffer-Size**
Applies to:		Webserver
Default Value:	`4`
Example:		`Search-Relay-Buffer-Size=8`
Description:	The Search-Relay-Buffer-Size parameter controls how many entry blocks the webserver will accept from each indexer while still waiting for outstanding blocks from another indexer. As search entries flow in temporally, it is possible that one indexer may still be processing older entries while another has moved ahead to more recent entries. Because the webserver must process entries in temporal order, it will buffer entries from the indexer which is "ahead" while waiting for the slower indexer to catch up. In general, the default value will help prevent memory problems while still providing acceptable performance. On systems with large amounts of memory, it may be useful to increase this value.

**Max-Search-History**
Applies to:        Webserver
Default Value:        `100`
Example:        `Max-Search-History=256`
Description:        The Max-Search-History parameter controls how many searches are kept for a user.  Search history is useful to be able to go back and examine old searches, or see what other users in your group are searching.  A larger history allows for a greater tail of old search strings, but if too many searches are kept in the history it can cause some slowdowns when interacting with the GUI.

**Prebuff-Block-Hint**
Applies to:        Indexer
Default Value:        `32`
Example:        `Prebuff-Block-Hint=8`
Description:        The Prebuff-Block-Hint specifies in megabytes a soft target that the indexer should shoot for when storing blocks of data.  Very high-throughput systems may want to push this value a little higher, where memory constrained systems may want to push this value lower.  This value is a soft target, and indexers will typically only engage it when ingest is occurring at high rates.

**Prebuff-Max-Size**
Applies to:        Indexer
Default Value:        `32`
Example:        `Prebuff-Max-Size=128`
Description:        The Prebuff-Max-Size parameter controls the maximum data size in megabytes the prebuffer will hold before forcing entries to disk.  The prebuffer is used to help optimize storage of entries when source clocks may not be very well synchronized.  A larger prebuffer means that the indexer can better optimize ingesters that are providing wildly out of order values.  Each well has its own prebuffer, so if your installation has 4 wells defined and a Prebuff-Max-Size of 256, the indexer can consume up to 1GB of memory holding data.  The prebuffer max size will typically only engage in high-throughput systems, as the prebuffer is periodically evicting entries and pushing them to the storage media all the time.  This is the second knob to turn (after Search-Pipeline-Buffer-Size) if your host system's OOM killer is terminating the Gravwell processes.

**Prebuff-Max-Set**
Applies to:        Webserver
Default Value:        `256`
Example:        `Prebuff-Max-Set=256`
Description:        The Prebuff-Max-Set specifies how many one-second blocks are allowed to be held in the prebuffer for optimization.  The more out of sync the timestamps are on entries provided by ingesters the larger this set should be.  For example, if you are consuming from sources that might have as much as a 2 hour swing in timestamps you might want to set this value to 7200, but if your data typically arrives with very tight timestamp tolerances you can shrink this value down as low as 10.  The Prebuff-Max-Size controls will still engage and force prebuffer evictions, so setting this value too high hurts less than setting it too low.

**Prebuff-Tick-Interval**
Applies to:        Webserver
Default Value:        `3`
Example:        `Prebuff-Tick-Interval=4`
Description:        The Prebuff-Tick-Interval parameter specifies in seconds how often the prebuffer should engage an artificial eviction of entries located in the prebuffer.  The prebuffer is always evicting values to persistent storage when there is active ingestion, but in very low-throughput systems this value can be used to ensure that entries are forcibly pushed to persistent storage.  Gravwell will never allow data to be lost when it can help it; when gracefully shutting down indexers the prebuffer ensures all entries make it to the persistent storage.  However, if you don’t have a lot of faith in the stability of your hosts you may want to set this interval closer to 2 to ensure that system failures, or angry admins, can’t pull the rug out from under the indexers.

**Prebuff-Sort-On-Consume**
Applies to:        Indexer
Default Value:        `false`
Example:        `Prebuff-Sort-On-Consume=true`
Description:        The Prebuff-Sort-On-Consume parameter tells the prebuffer to sort locks of data prior to pushing them to disk.  The sorting process is only applied to the individual block, and does NOT guarantee that data is sorted when entering the pipeline.  Sorting blocks prior to storage also incurs a significant performance penalty in ingestion.  Almost all installations should leave this value as false.

**Max-Block-Size**
Applies to:        Indexer
Default Value:        `4`
Example:        `Max-Block-Size=8`
Description:        The Max-Block-Size specifies a value in megabytes and is used as a hint to tell indexers the maximum block size they can generate when pushing entries into the pipeline.  Larger blocks reduce pressure on the pipeline, but increase memory pressure.  Large memory and high throughput systems can increase this value to increase throughput, smaller memory systems can decrease this size to reduce memory pressure.  The Prebuff-Block-Hint and Max-Block-Size parameters intersect to provide two knobs that tune ingest and search throughput.  At Gravwell, on the 128GB nodes, the following is achieved: a clean 1GB/s of search throughput; a 1.25 million entry per second ingest with a Max-Block-Size of 16; and a Prebuff-Block-Hint of 8 is achieved

**Render-Store-Limit**
Applies to:		Webserver
Default Value:	1024
Example:		`Render-Store-Limit=512`
Description:	The Render-Store-Limit parameter specifies how many megabytes a search renderer can store.

**Search-Control-Script**
Applies to:		Webserver
Default Value:
Example:		`Search-Control-Script=/opt/gravwell/etc/authscripts/limits.grv`
Description:	The Search-Control-Script parameter is a list parameter which can specify scripts to be applied at search time. Being a list parameter, it can be specified multiple times to specify multiple scripts. These scripts can apply additional restrictions to searches executed by users. All scripts are executed for every search. Contact Gravwell for more information about search control scripts.

**Webserver-Resource-Store**
Applies to:		Webserver
Default Value:	`/opt/gravwell/resources/webserver`
Example:		`Webserver-Resource-Store=/tmp/path/to/resources/webserver`
Description:	The Webserver-Resource-Store parameter specifies where the webserver should store its resources. This directory **must** be unused by any other process and cannot be specified as the resource location for the indexer or datastore.

**Indexer-Resource-Store**
Applies to:		Indexer
Default Value:	`/opt/gravwell/resources/indexer`
Example:		`Indexer-Resource-Store=/tmp/path/to/resources/indexer`
Description:	The Indexer-Resource-Store parameter specifies where the indexer should store its resources. This directory **must** be unused by any other process and cannot be specified as the resource location for the webserver or datastore.

**Datastore-Resource-Store**
Applies to:		Datastore
Default Value:	`/opt/gravwell/resources/datastore`
Example:		`Datastore-Resource-Store=/tmp/path/to/resources/datastore`
Description:	The Datastore-Resource-Store parameter specifies where the datastore should store its resources. This directory **must** be unused by any other process and cannot be specified as the resource location for the indexer or webserver.

**Resource-Max-Size**
Applies to:		Webserver, Datastore, and Indexer
Default Value:	`134217728`
Example:		`Resource-Max-Size=1000000000`
Description:	The Resource-Max-Size parameter specifies the maximum size of resources in bytes.

**Docker-Secrets**
Applies to:		Webserver, Datastore, and Indexer
Default Value:	`false`
Example:		`Docker-Secrets=true`
Description:	The Docker-Secrets parameter tells Gravwell that it should attempt to read the ingest, control, and search agent secrets from [Docker secrets](https://docs.docker.com/engine/swarm/secrets/). It expects the secrets to be named `ingest_secret`, `control_secret`, and `search_agent_secret`, respectively, and they should be accessible from within the VM in the `/run/secrets/` directory.

**HTTP-Proxy**
Applies to:		Webserver
Default Value:
Example:		`HTTP-Proxy=wwwproxy.example.com:8080`
Description:	The HTTP-Proxy parameter configures a proxy to be used for HTTP and HTTP requests by the webserver. It is effectively equivalent to setting the environment variable $http_proxy and allows the same syntax.  The specified proxy value will be used for both `HTTP` and `HTTPS` requests.

**Webserver-Ingest-Groups**
Applies to:		Webserver
Default Value:
Example:		`Webserver-Ingest-Groups=ingestUsers`
Description:	The Webserver-Ingest-Groups parameter is a list parameter which specifies groups whose users are allowed to ingest entries directly via the Gravwell web API. As a list parameter, it can be specified multiple times to enable multiple groups to ingest via web API.

**Disable-Update-Notification**
Applies to:		Webserver
Default Value:	`false`
Example:		`Disable-Update-Notification=false`
Description:	If Disable-Update-Notification is set to true, the web UI will not present a notification when a new version of Gravwell is available.

**Disable-Stats-Report**
Applies to: Webserver
Default Value: false
Example: `Disable-Stats-Report=true`
Description:	Setting this parameter to true will tell the webserver's [metrics reporting routine](#!metrics.md) to send only minimal information about the license, omitting the broader system statistics.

**Temp-Dir**
Applies to:		Webserver
Default Value:	`/opt/gravwell/tmp`
Example:		`Temp-Dir=/tmp/gravtmp`
Description:	The Temp-Dir parameter specifies a directory which can be used for temporary Gravwell files without risk of interference from other processes. It is used to store uploaded kits before installation, among other uses.

**Insecure-User-Unsigned-Kits-Allowed**
Applies to:		Webserver
Default Value:	`false`
Example:		`Insecure-User-Unsigned-Kits-Allowed=true`
Description:	This parameter, if set, allows all users to install unsigned kits. We strong recommend against enabling this option.

**Disable-Search-Agent-Notifications**
Applies to:		Webserver
Default Value:	`false`
Example:		`Disable-Search-Agent-Notifications=true`
Description:	If set to true, this parameter prevents the web UI from displaying a notification if the search agent fails to check in. This is useful if you have disabled the search agent and do not want to see the notification.

**Indexer-Storage-Notification-Threshold**
Applies to:		Indexer
Default Value:		`90`
Example:		`Indexer-Storage-Notification-Threshold=98`
Description:		A percentage value which determines when to warn about storage usage.  If the value is above 0, a notification will be thrown whenever a storage device that is used by the Indexer uses more than the specified storage percentage.  The value MUST be between 0 and 99.

**Disable-Network-Script-Functions**
Applies to:		Webserver
Default Value:	`false`
Example:		`Disable-Network-Script-Functions=true`
Description:	By default, anko scripts in the pipeline are allowed to use network functions such as the net/http library and the ssh/sftp utilities. Setting this to 'true' will disable those functions.

**Webserver-Enable-Frame-Embedding**
Applies to:		Webserver
Default Value:	`false`
Example:		`Webserver-Enable-Frame-Embedding=true`
Description:	By default, the webserver disallows Gravwell pages from being rendered within frames by setting the header X-Frame-Options: deny. Setting this configuration parameter to 'true' will eliminate that header, allowing the pages to be embedded within frames.

**Webserver-Content-Security-Policy**
Applies to:		Webserver
Default Value:	``
Example:		`Webserver-Content-Security-Policy="default-src https:"`
Description:	This parameter allows the administrator to defined a Content-Security-Policy header which will be sent with all Gravwell pages. This is an important security option and should be set for your organization based on your deployment requirements, such as requiring https-only.

**Default-Language**
Applies to:		Webserver
Default Value:		`en-US`
Example:		`Default-Language=en-US`
Description:		Setting the Default-Language parameter controls what is provided on the unauthenticated API at /api/language and is used by the GUI to determine which language should be default in deployments with multiple languages. This is the fallback if the user has not chosen a language and their browser is not providing a preferred language via `window.navigator.language`.

**Disable-Map-Tile-Server-Proxy**
Applies to:		Webserver
Default Value:	`false`
Example:		`Disable-Map-Tile-Server-Proxy=true`
Description:	This parameter controls Gravwell's built-in maps proxy. To avoid placing undue load on map servers, the Gravwell webserver caches map tiles. However, use of the proxy means that the requests sent to the actual map servers originate from the Gravwell webserver rather than the user's web browser; if the Gravwell installation is on a locked-down network, this may fail due to outgoing HTTP being disabled. Setting `Disable-Map-Tile-Server-Proxy` to true will disable the built-in proxy and cause the GUI to make map requests directly. If the proxy is disabled and the `Map-Tile-Server` parameter is also set, the GUI will make its requests to that server.

**Map-Tile-Server**
Applies to:		Webserver
Default Value:	``
Example:		`Map-Tile-Server=https://maps.example.com/osm/`
Description:	The Map-Tile-Server parameter allows the administrator to define a different source for map tiles. By default, Gravwell will fetch tiles from a Gravwell map server, falling back to OpenStreetMap servers if necessary. Setting this parameter forces Gravwell to use *only* the specified server. The URL specified should be a prefix to the standard OpenStreetMap tile server format as defined [here](https://wiki.openstreetmap.org/wiki/Tile_servers), leaving out the z/x/y coordinate parameters. For example, if tiles may be accessed at `https://maps.wikimedia.org/osm-intl/${z}/${x}/${y}.png`, e.g. https://maps.wikimedia.org/osm-intl/0/1/2.png, you could set `Map-Tile-Server=https://maps.wikimedia.org/osm-intl/`.

**Gravwell-Tile-Server-Cooldown-Minutes**
Applies to:		Webserver
Default Value:	5
Example:		`Gravwell-Tile-Server-Cooldown-Minutes=1`
Description:	When the Gravwell tile proxy is operating in normal mode (not disabled, `Map-Tile-Server` parameter not set), it will attempt to fetch map tiles from a Gravwell-operated server. If a request send to that server fails, the proxy will instead fall back to openstreetmap.org servers for the duration of the cooldown. Setting this parameter to 0 disables the cooldown.

**Gravwell-Tile-Server-Cache-MB**
Applies to:		Webserver
Default Value:	4
Example:		`Gravwell-Tile-Server-Cache-MB=32`
Description:	The Gravwell tile proxy maintains a cache of recently-accessed tiles to speed up map rendering. This parameter controls how many megabytes of storage the cache may use.

**Gravwell-Tile-Server-Cache-Timeout-Days**
Applies to:		Webserver
Default Value:	7
Example:		`Gravwell-Tile-Server-Cache-Timeout-Days=2`
Description:	The Gravwell tile proxy maintains a cache of recently-accessed tiles to speed up map rendering. This parameter controls the maximum number of days a cached tile should be considered valid; after that time has elapsed, the tile will be purged and re-fetched from the upstream server.

**Disable-Single-Indexer-Optimization**
Applies to:		Webserver
Default Value:	false
Example:		`Disable-Single-Indexer-Optimization=true`
Description:	When Gravwell is used with a single indexer, it will by default run all modules (except for the render module) on the *indexer* to reduce the amount of data transferred from the indexer to the webserver. This option disables that optimization. We strongly recommend leaving this option set to `false` unless instructed by Gravwell support.

**Library-Dir**
Applies to:		Webserver
Default Value:	`/opt/gravwell/libs`
Example:		`Library-Dir=/scratch/libs`
Description:	Scheduled scripts may import additional libraries using the `include` function. These libraries are fetched from an external repository and cached locally; this configuration option sets the directory in which the cached libraries are stored.

**Library-Repository**
Applies to:		Webserver
Default Value:	`https://github.com/gravwell/libs`
Example:		`Library-Repository=https://github.com/example/gravwell-libs`
Description:	Scheduled scripts may import additional libraries using the `include` function. These libraries are loaded from files found in the repository specified by this parameter. By default, it points to a Gravwell-maintained repository of convenient libraries. If you wish to provide your own set of libraries, set this parameter to point at a git repository you control.

**Library-Commit**
Applies to:		Webserver
Default Value:
Example:		`Library-Commit=19b13a3a8eb877259a06760e1ee35fae2669db73`
Description:	Scheduled scripts may import additional libraries using the `include` function. These libraries are loaded from files found in the repository specified by the `Library-Repository` option. By default, Gravwell uses the latest version. If a git commit string is specified, Gravwell will attempt to use the specified version of the repository instead.

**Disable-Library-Repository**
Applies to:		Webserver
Default Value:	false
Example:		`Disable-Library-Repository=true`
Description:	Scheduled scripts may import additional libraries using the `include` function. Setting `Disable-Library-Repository` to true disables this functionality.

**Gravwell-Kit-Server**
Applies to:	Webserver
Default Value:	https://kits.gravwell.io/kits
Example:	`Gravwell-Kit-Server=http://internal.mycompany.io/gravwell/kits`
Description:	Allows for overriding the Gravwell kitserver host, this can be useful in airgapped or segmented deployments where you host a mirror of the Gravwell kitserver.  Set this value to an empty string to completely disable access to the remote kitserver.
Example:
```
Gravwell-Kit-Server="" #disable remote access to gravwell kitserver
Gravwell-Kit-Server="http://gravwell.mycompany.com/kits" #override to use internal mirror
```

**Kit-Verification-Key**
Applies to: Webserver
Default Value:
Example: `Kit-Verification-Key=/opt/gravwell/etc/kits-pub.pem`
Description:	Specifies a file containing a public key to use when verifying kits from the kitserver. Set this value if you have specified an alternate Gravwell-Kit-Server; it is not necessary when using Gravwell's official kit server. Keys suitable for signing kits can be generated with the [gencert](https://github.com/gravwell/gencert) utility.

## Password Control

The `[Password-Control]` configuration section can be used to enforce password complexity rules when users are created or passwords are changed. Options set in this block apply only to webservers. These complexity configuration rules do not apply when using Single Sign On.

Note: The `Password-Control` section should not be declared more than once.

```
[Password-Control]
	Min-Length=8
	Require-Uppercase=true
	Require-Lowercase=true
	Require-Special=true
	Require-Special=true
```

**MinLength**
Default Value:	0
Example:		`MinLength=8`
Description:	`MinLength` specifies the minimum character length of passwords.

**Require-Uppercase**
Default Value:	false
Example:		`Require-Uppercase=true`
Description:	If `Require-Uppercase` is set, passwords must contain at least one upper-case character.

**Require-Lowercase**
Default Value:	false
Example:		`Require-Lowercase=true`
Description:	If `Require-Lowercase` is set, passwords must contain at least one lower-case character.

**Require-Number**
Default Value:	false
Example:		`Require-Number=true`
Description:	If `Require-Number` is set, passwords must contain at least one numeric digit.

**Require-Special**
Default Value:	false
Example:		`Require-Special=true`
Description:	If `Require-Special` is set, passwords must contain at least one "special" character. The set of special characters includes all Unicode characters which are not numbers or letters.

## Well Configuration

The parameters in this section apply to well specifications, including the `Default-Well` specification. Well configurations only apply to *indexers*; they are ignored by webservers. Below is a sample of two well configurations, `Default-Well` and another well named "pcap".

```
[Default-Well]
	Location=/opt/gravwell/storage/default/
	Cold-Location=/opt/gravwell/cold-storage/default
	Accelerator-Name=fulltext
	Accelerator-Engine-Override=bloom
	Max-Hot-Storage-GB=20

[Storage-Well "pcap"]
	Location=/opt/gravwell/storage/pcap
	Cold-Location=/opt/gravwell/cold-storage/pcap
	Hot-Duration=1D
	Cold-Duration=12W
	Delete-Frozen-Data=true
	Max-Hot-Storage-GB=20
	Disable-Compression=true
	Tags=pcap
```

The configuration file should contain exactly one `Default-Well` section. It may also include one or more `Storage-Well` sections.

Refer to the [ageout documentation](ageout.md) for more information on how wells move entries between hot, cold, and archive storage.

Note: `Default-Well` cannot include `Tags=` specifications; instead, the default well contains all tags *not contained in other wells*

**Location**
Default Value:	`/opt/gravwell/storage/default` for `Default-Well`, none for `Storage-Well`
Example:		`Location=/opt/gravwell/storage/foo`
Description:	This parameter controls where the well stores "hot" data. No two wells should be allowed to point at the same directory!

### Ageout Options

**Hot-Duration**
Default Value:
Example:		`Hot-Duration=1w`
Description:	This parameter determines how long data remains in the "hot" storage location. The value should consist of a number followed by a suffix, either "d" for "day" or "w" for "week"; thus `Hot-Duration=30d` indicates that data should be kept for 30 days. Note that data will not actually be moved out of hot storage unless the `Cold-Location` parameter is specified or `Delete-Cold-Data` is set to true.

**Cold-Location**
Default Value:
Example:		`Cold-Location=/opt/gravwell/cold_storage/foo`
Description:	This parameter sets the storage location for cold data, data which has been moved out of the hot store specified by `Location`.

**Cold-Duration**
Default Value:
Example:		`Cold-Duration=365d`
Description:	This parameter determines how long data remains in the "cold" storage location. Data will not actually be moved out of cold storage unless `Delete-Frozen-Data` is set to true.

**Max-Hot-Storage-GB**
Default Value:
Example:		`Max-Hot-Storage-GB=100`
Description:	This parameter sets a maximum disk consumption for a given well's hot storage, in gigabytes. If this number is exceeded, the oldest data will be migrated to cold storage (if possible), sent to cloud storage (if configured), or deleted (if allowed).

**Max-Cold-Storage-GB**
Default Value:
Example:		`Max-Cold-Storage-GB=100`
Description:	This parameter sets a maximum disk consumption for a given well's cold storage, in gigabytes. If this number is exceeded, the oldest data will be sent to cloud storage (if configured) and deleted (if allowed).

**Hot-Storage-Reserve**
Default Value:
Example:		`Hot-Storage-Reserve=10`
Description:	This parameter tells the well that it must leave at least a particular percentage free on the disk. Thus, if `Hot-Storage-Reserve=10` is set, the well will attempt to age-out data from hot storage when the disk reaches 90% utilization.

**Cold-Storage-Reserve**
Default Value:
Example:		`Cold-Storage-Reserve=10`
Description:	This parameter tells the well that it must leave at least a particular percentage free on the disk. Thus, if `Cold-Storage-Reserve=10` is set, the well will attempt to age-out data from cold storage when the disk reaches 90% utilization.

**Delete-Cold-Data**
Default Value:	false
Example:		`Delete-Cold-Data=true`
Description:	Setting this parameter true means that data can be deleted from the hot storage location when one of the ageout criteria has been met.

**Delete-Frozen-Data**
Default Value:	false
Example:		`Delete-Frozen-Data=true`
Description:	Setting this parameter true means that data can be deleted from the cold storage location when one of the ageout criteria has been met.

**Archive-Deleted-Shards**
Default Value:	false
Example:		`Archive-Deleted-Shards=true`
Description:	If this option is set, the well will attempt to upload shards to an external archive server before deleting them. Note that this will only work if the `[Cloud-Archive]` section is configured!

**Disable-Compression, Disable-Hot-Compression, Disable-Cold-Compression**
Default Value:	false
Example:		`Disable-Compression=true`
Description:	These parameters control user-mode compression of data in the wells. By default, Gravwell will compress data in the well. Setting `Disable-Hot-Compression` or `Disable-Cold-Compression` will disable it for the hot or cold storage, respectively; setting `Disable-Compression` disables it for both.

**Enable-Transparent-Compression, Enable-Hot-Transparent-Compression, Enable-Cold-Transparent-Compression**
Default Value:	false
Example:		`Enable-Transparent-Compression=true`
Description:	These parameters control kernel-level, transparent compression of data in the wells. If enabled, Gravwell can instruct the `btrfs` filesystem to transparently compress data. This is more efficient than user-mode compression. Setting `Enable-Transparent-Compression` true automatically turns off user-mode compression. Note that setting `Disable-Compression=true` will **disable** transparent compression.

**Ageout-Time-Override**
Default Value:
Example:		`Ageout-Time-Override="3:00AM"`
Description:	This parameter allows you to specify a particular time at which the ageout routine should run. This is typically not needed.

### Acceleration Options

**Accelerator-Name**
Default Value:	
Example:		`Accelerator-Name=json`
Description:	Setting the `Accelerator-Name` parameter (and the `Accelerator-Args` parameter) enables acceleration on the well. See [the acceleration documentation](#!configuration/accelerators.md) for more information.

**Accelerator-Args**
Default Value:	
Example:		`Accelerator-Args="username hostname \"strange-field.with.specials\".subfield"`
Description:	Setting the `Accelerator-Args` parameter (and the `Accelerator-Name` parameter) enables acceleration on the well. See [the acceleration documentation](#!configuration/accelerators.md) for more information.

**Accelerate-On-Source**
Default Value:	false
Example:		`Accelerate-On-Source=true`
Description:	Specifies that the SRC field of each module should be included. This allows combining a module like CEF with SRC.

**Accelerator-Engine-Override**
Default Value:	"index"
Example:		`Accelerator-Engine-Override=bloom`
Description:	Selects the acceleration engine to be used. By default, the indexing accelerator will be used. Setting this parameter to "bloom" will instead select the bloom filter.

**Collision-Rate**
Default Value:	0.001
Example:		`Collision-Rate=0.01`
Description:	Sets the accuracy of the bloom filter acceleration engine. Must be a value between 0.1 and 0.000001.

### General Options

**Disable-Replication**
Default Value:	false
Example:		`Disable-Replication=true`
Description:	If set, the contents of this well will not be replicated.

**Enable-Quarantine-Corrupted-Shards**
Default Value:	false
Example:		`Enable-Quarantine-Corrupted-Shards=true`
Description:	If set, corrupted shards which cannot be recovered will be copied to a quarantine location for later analysis. By default, badly corrupted shards may be deleted.

## Replication Configuration

The `[Replication]` section configures [Gravwell's replication capability](#!configuration/replication.md). An example configuration might look like this:

```
[Replication]
	Disable-Server=true
	Peer=10.0.01
	Storage-Location=/opt/gravwell/replication_storage
```

The replication configuration block only applies to indexers.

**Peer**
Default Value:
Example:		`Peer=10.0.0.1:9406`
Description:	The `Peer` parameter specifies a replication peer. It takes an IP or hostname, with an optional port at the end. If no port is specified, the default port (9406) is used. `Peer` may be specified multiple times.

**Listen-Address**
Default Value:	":9406"
Example:		`Listen-Address=192.168.1.1:9406`
Description:	This parameter defines the IP and port on which Gravwell should listen for *incoming* replication connections. By default, it listens on all interfaces on port 9406.

**Storage-Location**
Default Value:	
Example:		`Storage-Location=/opt/gravwell/replication`
Description:	Sets the storage location for data replicated from other Gravwell indexers.

**Max-Replicated-Data-GB**
Default Value:
Example:		`Max-Replicated-Data-GB=100`
Description:	Sets, in gigabytes, the maximum amount of replicated data to store. When this is exceeded, the indexer will begin walking the replicated data to clean up; it will first remove any shards which have been deleted on the original indexer, then it will begin deleting the oldest shards. Once the storage size is below the limit, deletion will stop.

**Replication-Secret-Override**
Default Value:
Example:		`Replication-Secret-Override=MyReplicationSecret`
Description:	By default, Gravwell uses the `Control-Auth` token to authenticate for replication. Setting this parameter will instead define a custom replication authentication token.

**Disable-TLS**
Default Value:	false
Example:		`Disable-TLS=true`
Description:	Setting this parameter to true will disable TLS for replication. The indexer will listen for unencrypted incoming connections, and will use unencrypted connections to communicate with peers.

**Key-File**
Default Value: (value of `Key-File` in `[Global]` section)
Example:		`Key-File=/opt/gravwell/etc/replication-key.pem`
Description:	This parameter allows you to use a separate key for TLS connections, rather than the globally-defined one.

**Certificate-File**
Default Value: (value of `Certificate-File` in `[Global]` section)
Example:		`Certificate-File=/opt/gravwell/etc/replication-cert.pem`
Description:	This parameter allows you to use a separate certificate for TLS connections, rather than the globally-defined one.

**Insecure-Skip-TLS-Verify**
Default Value:	false
Example:		`Insecure-Skip-TLS-Verify=false`
Description:	Setting this parameter to true will disable validation of TLS certificates when connecting to replication peers.

**Connect-Wait-Timeout**
Default Value:	30
Example:		`Connect-Wait-Timeout=60`
Description:	Configures the timeout, in seconds, to be used when connecting to replication peers.

**Disable-Server**
Default Value:	false
Example:		`Disable-Server=true`
Description:	Disables the replication *server* functionality. If set, the indexer will push its own data out to replication peers, but will not allow any other indexers to push to it.

**Disable-Compression**
Default Value:	false
Example:		`Disable-Compression=true`
Description:	Controls compression of replicated data. By default, replicated data is compressed on disk.

**Enable-Transparent-Compression**
Default Value:	false
Example:		`Enable-Transparent-Compression=true`
Description:	If this parameter is set to true, Gravwell will attempt to use btrfs transparent compression on the replicated data. Setting `Disable-Compression=true` will disable this!

## Single Sign-On Configuration

The `[SSO]` configuration section controls single sign-on options for the Gravwell webserver. A sample section could be as simple as this:

```
[SSO]
	Gravwell-Server-URL=https://10.10.254.1:8080
	Provider-Metadata-URL=https://sso.gravwell.io/FederationMetadata/2007-06/FederationMetadata.xml
```

But it will more frequently require additional configuration:

```
[SSO]
	Gravwell-Server-URL=https://10.10.254.1:8080
	Provider-Metadata-URL=https://sso.gravwell.io/FederationMetadata/2007-06/FederationMetadata.xml
	Groups-Attribute=http://schemas.xmlsoap.org/claims/Group
	Group-Mapping=Gravwell:gravwell-users
	Group-Mapping=TestGroup:testgroup
	Username-Attribute = "uid"
	Common-Name-Attribute = "cn"
	Given-Name-Attribute  = "givenName"
	Surname-Attribute = "sn"
	Email-Attribute = "mail"
```

Refer to the [SSO configuration documentation](sso.md) for more information.

**Gravwell-Server-URL**
Default Value:
Example:		`Gravwell-Server-URL=https://gravwell.example.org/`
Description:	Specifies the URL to which users will be redirected once the SSO server has authenticated them. This should be the user-facing hostname or IP address of your Gravwell server. This parameter is required.

**Provider-Metadata-URL**
Default Value:
Example:		 `Provider-Metadata-URL=https://sso.example.org/FederationMetadata/2007-06/FederationMetadata.xml`
Description:	Specifies the URL of the SSO server's XML metadata. The path shown above (`/FederationMetadata/2007-06/FederationMetadata.xml`) should work for AD FS servers, but may need to be adjusted for other SSO providers. This parameter is required.

**Insecure-Skip-TLS-Verify**
Default Value:	false
Example:		`Insecure-Skip-TLS-Verify=true`
Description:	If set to true, this parameter instructs Gravwell to ignore invalid TLS certificates when communicating with the SSO server. Set this option with care!

**Username-Attribute**
Default Value:	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"
Example: 		`Username-Attribute = "uid"`
Description:	Defines the SAML attribute which will contain the username. On a Shibboleth server this should be set to "uid" instead.

**Common-Name-Attribute**
Default Value: "http://schemas.xmlsoap.org/claims/CommonName"
Example:		`Common-Name-Attribute="cn"`
Description:	Defines the SAML attribute which will contain the user's "common name". On a Shibboleth server this should be set to "cn" instead.

**Given-Name-Attribute**
Default Value:	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname"
Example:		`Given-Name-Attribute="givenName"`
Description:	Defines the SAML attribute which will contain the user's given name. On a Shibboleth server this should be set to "givenName" instead.

**Surname-Attribute**
Default Value:	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"
Example:		`Surname-Attribute=sn`
Description:	Defines the SAML attribute which will contain the user's surname. On a Shibboleth server this should be set to "sn" instead.

**Email-Attribute**
Default Value:	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
Example:		`Email-Attribute="mail"`
Description:	Defines the SAML attribute which will contain the user's email address. On a Shibboleth server this should be set to "mail" instead.

**Groups-Attribute**
Default Value:	"http://schemas.microsoft.com/ws/2008/06/identity/claims/groups"
Example:		`Groups-Attribute="groups"`
Description:	Defines the SAML attribute which contain's the list of groups to which the user belongs. You will typically have to explicitly configure the SSO provider to send the group list.

**Group-Mapping**
Default Value:	
Example:		`Group-Mapping=Gravwell:gravwell-users`
Description:	Defines one of the groups which may be automatically created if listed in the user's group memberships. This may be specified multiple times to allow multiple groups. The argument should consist of two names separated by a colon; the first is the SSO server-side name for the group (typically a name for AD FS, a UUID for Azure, etc.) and the second is the name Gravwell should use. Thus, if we define `Group-Mapping=Gravwell Users:gravwell-users`, if we receive a login token for a user who is a member of the group "Gravwell Users", we will create a local group named "gravwell-users" and add the user to it.

## Cloud Archive Configuration

Gravwell can be configured to archive data shards to a remote Cloud Archive server before deleting them, via the `Archive-Deleted-Shards` parameter on individual wells. The `[Cloud-Archive]` configuration section defines information about the Cloud Archive server to enable this.

```
[Cloud-Archive]
	Archive-Server=10.0.0.2:443
	Archive-Shared-Secret=MyArchiveSecret
```

Cloud archive configurations only apply to indexers.

**Archive-Server**
Default Value:
Example:		`Archive-Server=cloudarchive.example.org:443`
Description:	This parameter specifies an IP/hostname and optionally a port for the Cloud Archive server. If no port is specified, default (443) is used.

**Archive-Shared-Secret**
Default Value:
Example:		`Archive-Shared-Secret=MyArchiveSecret`
Description:	Sets the shared secret to use when authenticating to the Cloud Archive server. The indexer will use the license's customer ID number as the other half of the authentication process.

**Insecure-Skip-TLS-Verify**
Default Value:	false
Example:		`Insecure-Skip-TLS-Verify=true`
Description:	If set to true, the indexer will not verify the Cloud Archive server's TLS certificate when connecting.
