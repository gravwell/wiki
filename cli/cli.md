# Gravwell CLI

The Gravwell command line client can be used to remotely manage Gravwell and perform searches (with limited renderer support).  Administrators can manage users and monitor cluster health without the need for a full web browser.  Users can perform searches and easily export results to files for additional analysis with other tools.

The command line client is slightly limited in that it cannot render some search results (e.g. the CLI cannot draw a chart in a terminal, so it will refuse to render a search that uses the chart module).  However, the CLI does have access to all renderer modules when issuing backgrounded searches, which may be useful if an advanced user wanted to login remotely and start a few very large searches that will be ready for viewing on a full browser once they get on-site.

On a typical installation, the CLI tool will be installed as `/usr/sbin/gravwell`; passing the `-h` flag will give you an idea on where to start.  By default the Gravwell client expects the webserver to be listening on the local machine, specify the `-s` flag to point it at other webservers or a remote Gravwell instance.

```
gravwell options
  -b	Background the search
  -debug string
    	Enable JSON output debugging to a file
  -f string
    	Output format: "simple" or "grid" (default "grid")
  -insecure
    	Do NOT enforce webserver certificates, TLS operates in insecure mode
  -insecure-no-https
    	Use insecure HTTP connection, passwords are shipped plaintext
  -o string
    	Output to file rather than stdio
  -query string
    	Query string
  -r	Raw output, no pretty print
  -s string
    	Address and port of Gravwell webserver
  -si
    	Enable additional search information output
  -t	Disable sessions, always require logins
  -time string
    	Query time range
  -v	Output client version number and exit
  -watch-interval int
    	Watch update interval
OPTIONS:
	shell: enter the interactive shell
	state: Show the state of configured indexers
	desc: Show the description of each connected indexer
	storage: Show the current state of each connected indexer
	systems: Show performance metrics of each addr
	indexes: Show size of each index
	ingesters: Show activity and performance metrics of each ingester
	sessions: Show your other active sessions
	notifications: Show your active notifications
	search: Perform search
	attach: Reattaches to existing search
	download: Download search results in a packaged format
	admin: Perform admin actions
	user: Perform user actions
	dashboards: Perform dashboard actions
	logout: Logout of the current session
	logoutall: Logout all sessions using your UID
	list_searches: list_searches
	search_ctrl: Issue search control command
	resource: Create and manage resources
	macro: Manage search macros
	kits: Manage and upload kits
	userfiles: Manage user files
	templates: Manage templates
	pivots: Manage pivots
	ingest: Ingest entries directly
	scheduled_search: Manage scheduled searches
	script: Run a script
	help: Display available commands
MODIFIERS:
	watch: Continually show results of stats commands

EXAMPLE: gravwell -s=localhost state
```

The Gravwell client is also a great way to perform searches on Gravwell and feed the output to other tools.  For instance if you have a custom program for processing security data, but prefer to store your log entries in Gravwell, you can run a background query using the CLI client to extract the entries, then save the results to a file for the custom program to read.

## Using the CLI interactively.

The Gravwell CLI client provides an interactive shell similar to those found on commercial switches. It has different "menu" levels; for example, from the top level menu one might select the 'dashboards' sub-menu, which contains commands for managing dashboards. This section will describe the basics of using the client interactively.

### Connecting & Logging In

By default, the client will assume the Gravwell webserver is listening on `localhost:443`. If this is correct, you can connect by simply running the command `gravwell`. The client will prompt for your username and password, then display a prompt:

```
$ gravwell
Username:  admin
Password:  changeme
#> 
```

If your webserver is on a different host, use the `-s` flag to specify the hostname and port, e.g. `gravwell -s webserver.example.com:4443`.

If your webserver has self-signed TLS certificates installed, you will need to add the `-insecure` flag to disable TLS verification but still use HTTPS.

If your webserver does not have TLS certificates installed, add the `-insecure-no-https` flag to use HTTP-only mode. Note that this is insecure: your password will be sent to the server in plain text.

### Listing Available Commands

The `help` command will list the commands available at the current menu level. Immediately after launching the client, it will be at the top level:

```
#>  help
shell                enter the interactive shell
state                Show the state of configured indexers
desc                 Show the description of each connected indexer
storage              Show the current state of each connected indexer
systems              Show performance metrics of each addr
indexes              Show size of each index
ingesters            Show activity and performance metrics of each ingester
sessions             Show your other active sessions
notifications        Show your active notifications
search               Perform search
attach               Reattaches to existing search
download             Download search results in a packaged format
admin                Perform admin actions
user                 Perform user actions
dashboards           Perform dashboard actions
logout               Logout of the current session
logoutall            Logout all sessions using your UID
list_searches        list_searches
search_ctrl          Issue search control command
resource             Create and manage resources
macro                Manage search macros
kits                 Manage and upload kits
userfiles            Manage user files
templates            Manage templates
pivots               Manage pivots
ingest               Ingest entries directly
scheduled_search     Manage scheduled searches
script               Run a script
help                 Display available commands
```

Some of the items listed are commands:

```
#>  state
+----------------------+----------+
|               System |    State |
+======================+==========+
|    192.168.2.60:9404 |       OK |
+----------------------+----------+
|            webserver |       OK |
+----------------------+----------+
```

Others are menus which will contain their own commands. In the example below, we select the 'dashboards' menu, list the available commands, and run the 'list' command:

```
#>  dashboards
dashboards>  help
list                	List available user dashboards
mine                	List dashboards owned by you
del                 	Delete a dashboard available user dashboards
clone               	Clone a dashboard to enable ownership and editing
dashboards>  list
+---------+-------+-----------------+------------------------------+----------+-----------+
|    Name |    ID |     Description |                      Created |    Owner |    Groups |
+=========+=======+=================+==============================+==========+===========+
|     Foo |    10 |    My dashboard |    2019-04-15T12:19:49-06:00 |    admin |           |
+---------+-------+-----------------+------------------------------+----------+-----------+
dashboards>  
```

## Searching via CLI

The `search` command runs a search in the foreground:

```
#>  search
query>  tag=* count
time range> -1h
count 100
1/1
Press q[Enter] to quit, [Entry] to continue

Total Items: 1
101 stats records from Apr 15 12:39:37.000 <-> Apr 15 13:39:38.000
count 100.00/1.00 61.66 KB/616 B 8.109585ms
```

If you wish to save the results of a search, we can run the client with the '-b' flag, which specifies that searches should be run in the background, then use the `search` and `download` commands to run a search and save the results:

```
$ gravwell -b
#>  search
query>  tag=* json state=="NM"
time range> -1h
Background search with ID 065015787 launched
#>  download
+--------------+---------+----------+------------+---------------------+------------+
|    Search ID |    User |    Group |      State |    Attached Clients |    Storage |
+==============+=========+==========+============+=====================+============+
|    065015787 |       1 |        0 |    DORMANT |                   0 |    1.56 KB |
+--------------+---------+----------+------------+---------------------+------------+
search ID>  065015787
Available formats:
json
text
csv
format>  text
Save Location (default: /tmp)>  /tmp/nm.txt
Saving to  /tmp/nm.txt
#>  
```


## Admin

The Gravwell client implements many commands for managing the system in the admin sub-menu:

```
#>  admin
admin>  help
add_user            Add a new user
impersonate_user    Impersonate an existing users
del_user            Delete an existing user
get_user            Get an existing users details
update_user         Update an existing user
list_users          List all users
lock_user           Lock a user account
user_activity       Show a specific users activity
user_sessions       Show all open sessions
change_user_pwd     Change a users password
change_admin        Set a users admin status
add_group           Create a new group
del_group           Delete an existing group
list_groups         Lists all existing groups
list_group_users    Lists all members of an existing group
update_group        Update an existing group
add_users_group     Add users to an existing group
del_users_group     Delete users from an existing group
add_user_groups     Add user to existing groups
del_user_groups     Delete a user from groups
get_log_level       Get the webservers current logging level
set_log_level       Set the webservers current logging level
all_dashboards      Get all dashboards for all users
del_dashboard       Delete a dashboard owned by another user
license_info        Display license information
license_sku         Display license SKU
license_serial      Display license Serial Number
license_update      Upload a new license
list_queries        List all queries (active and saved) for all users
delete_queries      Delete any query (active or saved) for any user
list_users_storage  List all users current storage usage
add_indexer         Add another indexer to the configuration
list_kits           List all kits across all users
uninstall_kit       Uninstall a kit owned by any user
list_extractions    List installed autoextractors
add_extraction      Add a new autoextractor
delete_extraction   Delete an installed autoextractor
update_extraction   Update an installed autoextractor
sync_extractions    Force a sync of installed autoextractors to indexers
```

In addition to user/group management, the admin menu also provides tools to manage dashboards, kits, and other objects belonging to other users on the system.


## CLI examples

### Check On Indexer Health

```
$ gravwell state
+----------------------+----------+
|               System |    State |
+======================+==========+
|    10.0.0.2:9404     |       OK |
+----------------------+----------+
|    10.0.0.3:9404     |       OK |
+----------------------+----------+
|    webserver         |       OK |
+----------------------+----------+
```

Output of every command can be set to “raw” with no tables or formatting.  The raw output can be easier to digest if you are passing Gravwell data to other tools or scripts.

```
$ gravwell -r state
10.0.0.3:9404 OK
webserver     OK
10.0.0.2:9404 OK
```

### View Indexer Wells and Storage Size

```
$ gravwell -r indexes
10.0.0.2:9404 default /mnt/storage/gravwell/default 14.8 MB 93.76 K
10.0.0.2:9404 pcap /mnt/storage/gravwell/pcap 3.6 MB 29.63 K
10.0.0.2:9404 bench /mnt/storage/gravwell/bench 142.5 GB 686.66 M
10.0.0.2:9404 reddit /mnt/storage/gravwell/reddit 34.3 GB 73.72 M
10.0.0.2:9404 fcc /mnt/storage/gravwell/fcc 21.7 GB 11.09 M
10.0.0.2:9404 raw /mnt/storage/gravwell/raw 76.3 KB 0
10.0.0.2:9404 syslog /mnt/storage/gravwell/syslog 60.2 MB 406.62 K
10.0.0.3:9404 default /mnt/storage/gravwell/default 12.2 MB 77.56 K
10.0.0.3:9404 reddit /mnt/storage/gravwell/reddit 34.3 GB 73.66 M
10.0.0.3:9404 fcc /mnt/storage/gravwell/fcc 21.6 GB 11.06 M
10.0.0.3:9404 pcap /mnt/storage/gravwell/pcap 5.1 MB 44.85 K
10.0.0.3:9404 syslog /mnt/storage/gravwell/syslog 79.6 MB 536.86 K
10.0.0.3:9404 raw /mnt/storage/gravwell/raw 76.3 KB 0
10.0.0.3:9404 bench /mnt/storage/gravwell/bench 136.5 GB 658.69 M
```

### View Remote Ingesters

```
$ gravwell -r ingesters
10.0.0.2:9404
        tcp://10.0.0.1:49544 111h27m33.8s [reddit] 5.16 M 2.68 GB
        tcp://192.210.192.202:45578 34m52.1s [pcap] 62.00 3.69 KB
        tcp://192.210.192.202:43369 34m51.9s [kernel] 1.1 K 121.43 KB
10.0.0.3:9404
        tcp://10.0.0.1:49770 111h27m0.01s [reddit] 5.24 M 2.72 GB
        tcp://192.210.192.202:43364 34m52.6s [pcap] 119.00 6.93 KB
        tcp://192.210.192.202:43368 34m51.9s [kernel] 1.33 K 141.57 KB
```

### Run a script

```
$ gravwell script
script file path>  /tmp/myscript.ank
```
