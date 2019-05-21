# Global Configuration Parameters

Each parameter has a default value which applies if the parameters is empty, or not specified in the configuration file.

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

**HTTP-Proxy***
Applies to:		Webserver
Default Value:
Example:		`HTTP-Proxy=http://wwwproxy.example.com:8080/`
Description:	The HTTP-Proxy parameter configures a proxy to be used for HTTP and HTTP requests by the webserver. It is effectively equivalent to setting the environment variable $http_proxy and allows the same syntax.

**Webserver-Ingest-Groups**
Applies to:		Webserver
Default Value:
Example:		`Webserver-Ingest-Groups=ingestUsers`
Description:	The Webserver-Ingest-Groups parameter is a list parameter which specifies groups whose users are allowed to ingest entries directly via the Gravwell web API. As a list parameter, it can be specified multiple times to enable multiple groups to ingest via web API.

**Autoextract-Definition-Path**
Applies to:		Webserver and Indexer
Default Value:	`/opt/gravwell/extractions` 
Example:		`Autoextract-Definition-Path=/tmp/extractions`
Description:	The Autoextract-Definition-Path parameter specifies a directory which will contain autoextractor definitions.

**Disable-Update-Notification**
Applies to:		Webserver
Default Value:	`false`
Example:		`Disable-Update-Notification=false`
Description:	If Disable-Update-Notification is set to true, the web UI will not present a notification when a new version of Gravwell is available.

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
