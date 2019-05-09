# REST Stats API

## Ping

The ping stats API returns the status of the webserver and indexers. This is retrieved via a GET to `/api/stats/ping`. If all systems are up, they will be reported as "OK":

```
{"192.168.2.60:9404":"OK","webserver":"OK"}
```

If an indexer is down, it will be marked "Disconnected":

```
{"192.168.2.60:9404":"Disconnected","webserver":"OK"}
```

## Index Stats

The indexer stats API provides information about the indexes on each indexer. It is accessed via a GET to `/api/stats/idxStats`.

```
{
    "192.168.2.60:9404": {
        "IndexStats": [
            {
                "Name": "default",
                "Stats": [
                    {
                        "Cold": false,
                        "Data": 461610307,
                        "Entries": 4327438,
                        "Path": "/opt/gravwell/storage/default"
                    },
                    {
                        "Cold": true,
                        "Data": 3637724726,
                        "Entries": 33315409,
                        "Path": "/opt/gravwell/cold-storage/default"
                    }
                ]
            },
            {
                "Name": "csv",
                "Stats": [
                    {
                        "Accelerator": "index",
                        "Cold": false,
                        "Data": 69904,
                        "Entries": 0,
                        "Extractor": "csvAcceleratorV1",
                        "Path": "/opt/gravwell/storage/csv"
                    }
                ]
            },
[...]
            {
                "Name": "test",
                "Stats": [
                    {
                        "Accelerator": "index",
                        "Cold": false,
                        "Data": 775980546,
                        "Entries": 2000000,
                        "Extractor": "jsonAcceleratorV1",
                        "Path": "/opt/gravwell/storage/test2"
                    }
                ]
            }
        ]
    }
}

```

## Ingester Stats

Sending a GET request to `/api/stats/igstStats` will return a structure describing the ingesters attached to each indexer. The example below shows a single indexer (192.168.2.60) with two attached ingesters: the Simple Relay ingester and the Network Capture ingester.

```
{
    "192.168.2.60:9404": {
        "Ingesters": [
            {
                "Count": 6,
                "RemoteAddress": "unix://@",
                "Size": 659,
                "Tags": [
                    "bro",
                    "default",
                    "syslog"
                ],
                "Uptime": 5639681950
            },
            {
                "Count": 3,
                "RemoteAddress": "tcp://192.168.2.60:43684",
                "Size": 229,
                "Tags": [
                    "pcap"
                ],
                "Uptime": 2846761051
            }
        ],
        "QuotaMax": 0,
        "QuotaUsed": 0,
        "TotalCount": 9,
        "TotalSize": 888
    }
}
```

In each ingester description, the "Count" field gives the number of entries ingested and the "Size" field gives the number of bytes ingested. "Uptime" is how long, in nanoseconds, the ingester has been connected.

Note the "QuotaMax" and "QuotaUsed" fields. Community licenses can only ingest a certain amount each day. "QuotaMax" specifies how many bytes the given indexer is allowed to ingest per day. "QuotaUsed" shows how many bytes have been ingested so far today.

## System Stats

Sending a GET request to `/api/stats/sysStats` will return a structure containing information about the indexer and webserver systems, such as number of CPUs, CPU utilization, memory and network utilization, and more.

```
{
    "192.168.2.60:9404": {
        "Stats": {
            "CPUCount": 4,
            "CPUUsage": 28.717948741951247,
            "Disks": [
                {
                    "Mount": "/",
                    "Partition": "/dev/mapper/foo--vg-root",
                    "Total": 233377820672,
                    "Used": 170719322112
                }
            ],
            "HostHash": "bef3ac74c4bd31fc15d37bbbd927ea7213d7ea0d922126ed07c34e2c41a9ca12",
            "IO": [
                {
                    "Device": "nvme0n1p2",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "foo--vg-root",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "nvme0n1p1",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "sda1",
                    "Read": 0,
                    "Write": 0
                }
            ],
            "MemoryUsedPercent": 39.42591390055748,
            "Net": {
                "Down": 1160,
                "Up": 310
            },
            "TotalMemory": 16721588224,
            "Uptime": 15178980
        }
    },
    "webserver": {
        "Stats": {
            "CPUCount": 4,
            "CPUUsage": 28.589743582518338,
            "Disks": [
                {
                    "Mount": "/boot",
                    "Partition": "/dev/nvme0n1p2",
                    "Total": 247772160,
                    "Used": 108133376
                },
                {
                    "Mount": "/",
                    "Partition": "/dev/mapper/foo--vg-root",
                    "Total": 233377820672,
                    "Used": 170719322112
                }
            ],
            "HostHash": "bef3ac74c4bd31fc15d37bbbd927ea7213d7ea0d922126ed07c34e2c41a9ca12",
            "IO": [
                {
                    "Device": "nvme0n1p1",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "foo--vg-root",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "sda1",
                    "Read": 0,
                    "Write": 0
                },
                {
                    "Device": "nvme0n1p2",
                    "Read": 0,
                    "Write": 0
                }
            ],
            "MemoryUsedPercent": 39.42591390055748,
            "Net": {
                "Down": 747,
                "Up": 255
            },
            "TotalMemory": 16721588224,
            "Uptime": 15178980
        }
    }
}
```

Most of the fields are self-explanatory. Note that the "IO" array gives information about disks, with "Read" and "Write" fields specifying read and write *rates* in bytes per second. Similarly, the "Net" component describes network utilization in bytes per second.

## System Description

Sending a GET request to `/api/stats/sysDesc` will return a structure giving additional information about the webserver and indexer host systems:

```
{
    "192.168.2.60:9404": {
        "CPUCache": "4096",
        "CPUCount": 4,
        "CPUMhz": "3500",
        "CPUModel": "Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz",
        "SystemVersion": "4.9.0-8-amd64",
        "TotalMemoryMB": 15946
    },
    "webserver": {
        "CPUCache": "4096",
        "CPUCount": 4,
        "CPUMhz": "3500",
        "CPUModel": "Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz",
        "SystemVersion": "4.9.0-8-amd64",
        "TotalMemoryMB": 15946
    }
}
```