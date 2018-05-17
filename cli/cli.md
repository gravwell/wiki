# Gravwell CLI

The Gravwell command line client can be used to remotely manage Gravwell as well as perform searches (with limited renderer support).  Administrators can manage users and monitor cluster health without the need for a full web browser.  Users can perform searches and easily export results to files for additional analysis with other tools.

The command line client is slightly limited in that it cannot render some search results (e.g. the CLI cannot draw a chart in a terminal, so it will refuse to render a search that uses the render module).  However, the CLI does have access to all renderer modules when issuing backgrounded searches, which may be useful if an advanced user wanted to login remotely and start a few very large searches that will be ready for viewing on a full browser once they get onsite.

The CLI has an abundance of options and features, including watching cluster status.  Passing the `-h` flag will give you an idea on where to start.  By default the Gravwell client expects the webserver to be listening on the local machine, specify the `-s` flag to point it at other webservers or a remote Gravwell instance.

```
gravwell options
  -b        Background the search
  -debug string
            Enable JSON output debugging to a file
  -e        Enforce webserver certificates
  -f string
            Output format: "simple" or "grid" (default "grid")
  -o string
            Output to file rather than stdio
  -query string
            Query string
  -r        Raw output, no pretty print
  -s string
            Address and port of Gravwell webserver (default "127.0.0.1:443")
  -si
            Enable additional search information output
  -t        Disable sessions, always require logins
  -time string
            Query time range
  -v        Output client version number and exit
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
        admin: Perform admin actions
        user: Perform user actions
        logout: Logout of the current session
        logoutall: Logout all sessions using your UID
        list_searches: list_searches
        search_ctrl: Issue search control command
        help: Display available commands
MODIFIERS:
        watch: Continually show results of stats commands
EXAMPLE: gravwell -s=localhost state
```

The Gravwell client is also a great way to perform searches on Gravwell and feed the output to other tools.  For instance if you have a super secret program for processing security data, but don’t want to burden it with all your log entries you can perform a search in Gravwell and pass long only the data that your super secret system cares about.  Redirect output of searches with the `-o` flag.
Example CLI Actions

Below are just some of the operations you can perform with the CLI.

## Check On Indexer Health

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

## View Indexer Wells and Storage Size

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

## View Remote Ingesters

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
