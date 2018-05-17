# Getting Status

Websocket located at /ws/stats

The server will immediately start throwing stats at you via the RoutingWebsocket protocol.  It expects the following types to be registered and will produce data on each:

* ping
* idxStats
* sysStats
* sysDesc

Upon connecting the websocket will throw a single JSON package on the sysDesc type first (and ALWAYS FIRST).  This can reliably be used to enumerate the number of backend systems you will continue to get stats about.  Use that first packet to build out tables and what not, to be continually filled via the idxStats and sysStats types.  The ping type is a keep alive, and should you not get an update from the ping type after 30 seconds, the backend is dead, let the user know.

# type examples

## system description subproto sysDesc

```
{
   "HostNameA": {
        CPUCount:      int
        CPUModel:      string
        CPUMhz:        string
        CPUCache:      string
        CPUMIPS:       string
        TotalMemoryMB: uint64
        SystemVersion: string
    },
    "HostNameB": {
        CPUCount:      int
        CPUModel:      string
        CPUMhz:        string
        CPUCache:      string
        CPUMIPS:       string
        TotalMemoryMB: uint64
        SystemVersion: string
    }
}
```

## example ping JSON
```
{
	"Error": "",
	"States": [
		{
			"Addr": "localhost:9404",
			"State": "OK"
		},
		{
			"Addr": "10.0.0.1:9404",
			"State": "OK"
		}
	]
}
```

## example sysStats JSON packet
```json
{
	"Error": "",
	"Stats": [
		{
			"Name": "webserver",
			"Error": "",
			"Host": {
				"Uptime": 1237164,
				"TotalMemory": 8219250688,
				"FreeMemory": 501071872,
				"CachedMemory": 5131911168,
				"Disks": [
					{
						"Mount": "/",
						"Partition": "/dev/mapper/mint--vg-root",
						"Total": 243347922944,
						"Used": 30278324224
					},
					{
						"Mount": "/boot",
						"Partition": "/dev/sda1",
						"Total": 246755328,
						"Used": 47550464
					}
				],
				"CPUUsage": 9.66123,
				"Net": {
					"Up": 0,
					"Down": 0
				}
			}
		},
		{
			"Name": "localhost:9404",
			"Error": "",
			"Host": {
				"Uptime": 3946,
				"TotalMemory": 8248610816,
				"FreeMemory": 7495507968,
				"CachedMemory": 317022208,
				"Disks": [
					{
						"Mount": "/",
						"Partition": "/dev/disk/by-uuid/b3b698db-b6c7-4490-a34f-e60e57c9b8e0",
						"Total": 9707950080,
						"Used": 3922296832
					},
					{
						"Mount": "/",
						"Partition": "/dev/disk/by-uuid/b3b698db-b6c7-4490-a34f-e60e57c9b8e0",
						"Total": 9707950080,
						"Used": 3922296832
					},
					{
						"Mount": "/home",
						"Partition": "/dev/sda3",
						"Total": 19549782016,
						"Used": 12020940800
					},
					{
						"Mount": "/mnt/storage1",
						"Partition": "/dev/sda6",
						"Total": 284401197056,
						"Used": 214626893824
					}
				],
				"CPUUsage": 2.5773196,
				"Net": {
					"Up": 1709294,
					"Down": 1710926
				}
			}
		},
		{
			"Name": "10.0.0.1:9404",
			"Error": "",
			"Host": {
				"Uptime": 3946,
				"TotalMemory": 8248610816,
				"FreeMemory": 7495507968,
				"CachedMemory": 317022208,
				"Disks": [
					{
						"Mount": "/",
						"Partition": "/dev/disk/by-uuid/b3b698db-b6c7-4490-a34f-e60e57c9b8e0",
						"Total": 9707950080,
						"Used": 3922296832
					},
					{
						"Mount": "/",
						"Partition": "/dev/disk/by-uuid/b3b698db-b6c7-4490-a34f-e60e57c9b8e0",
						"Total": 9707950080,
						"Used": 3922296832
					},
					{
						"Mount": "/home",
						"Partition": "/dev/sda3",
						"Total": 19549782016,
						"Used": 12020940800
					},
					{
						"Mount": "/mnt/storage1",
						"Partition": "/dev/sda6",
						"Total": 284401197056,
						"Used": 214626893824
					}
				],
				"CPUUsage": 2.5773196,
				"Net": {
					"Up": 1709294,
					"Down": 1710926
				}
			}
		}
	]
}
```


## example indexer stats JSON packet

```json
{
	"type": "idxStats",
	"data": {
		"Stats": {
			"127.0.0.1:9404": {
				"IndexStats": [{
					"Name": "default",
					"Stats": [{
						"Data": 39859808,
						"Entries": 504549,
						"Path": "/opt/sinkhole/storage/default"
					}]
				}, {
					"Name": "testingB",
					"Stats": [{
						"Data": 24,
						"Entries": 0,
						"Path": "/opt/sinkhole/storage/testB"
					}]
				}, {
					"Name": "testingA",
					"Stats": [{
						"Data": 3100139,
						"Entries": 30000,
						"Path": "/opt/sinkhole/storage/testA"
					}]
				}, {
					"Name": "testingC",
					"Stats": [{
						"Data": 24,
						"Entries": 0,
						"Path": "/opt/sinkhole/storage/testC"
					}]
				}]
			}
		}
	}
}
```