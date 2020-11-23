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


## Shard Storage and Replication Stats

Indexers maintain a list of all shards and wells and can produce a shard level view into the total stored data within Gravwell.  This view can provide a very quick order of magnitude observation of wells and data volumes over long periods of time.

When in high availability mode, the indexers also maintain a mapping of replicated data and can resolve where data is replicated too, enabling a quick overview of which indexers are replicating for other indexers.

The shard level view is accessed via a `GET` request to `/api/indexers/info` and will return a JSON map of each indexer. The returned data set has extensive information about the configuration of the well, what tags have been assigned to it, and the shards populated within the well.

### Example JSON Response
Here is an example response from a cluster of 4 Gravwell indexers with 2 wells each and replication enabled.  Only a single shard is populated in the syslog well.

<details><summary>Expand JSON Response</summary>
<pre>
```
{
  "172.19.0.4:9404": {
    "UUID": "f71ae8ea-5659-4ed2-8e4e-d7ebad4853c6",
    "Wells": [
      {
        "Name": "default",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "default",
          "gravwell"
        ],
        "Shards": []
      },
      {
        "Name": "syslog",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "syslog",
          "kernel",
          "dmesg"
        ],
        "Shards": [
          {
            "Name": "76ba7",
            "Start": "2020-11-23T19:09:52Z",
            "End": "2020-11-25T07:34:24Z",
            "Entries": 25794,
            "Size": 13163774,
            "RemoteState": {
              "UUID": "f71ae8ea-5659-4ed2-8e4e-d7ebad4853c6",
              "Entries": 25794,
              "Size": 13163262
            }
          }
        ]
      }
    ],
    "Replicated": {
      "5e20794b-0b73-4eb0-b49b-ece17089bf28": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "9a779454-95d8-457b-9841-aab9b93661fe": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "bc5ff11f-34a1-460c-adeb-4adc8c031777": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": [
            {
              "Name": "76ba7",
              "Start": "2020-11-23T19:09:52Z",
              "End": "2020-11-25T07:34:24Z",
              "Entries": 24092,
              "Size": 12261276,
              "RemoteState": {
                "UUID": "00000000-0000-0000-0000-000000000000",
                "Entries": 0,
                "Size": 0
              }
            }
          ]
        }
      ]
    }
  },
  "172.19.0.5:9404": {
    "UUID": "5e20794b-0b73-4eb0-b49b-ece17089bf28",
    "Wells": [
      {
        "Name": "default",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "default",
          "gravwell"
        ],
        "Shards": []
      },
      {
        "Name": "syslog",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "syslog",
          "kernel",
          "dmesg"
        ],
        "Shards": [
          {
            "Name": "76ba7",
            "Start": "2020-11-23T19:09:52Z",
            "End": "2020-11-25T07:34:24Z",
            "Entries": 25861,
            "Size": 13182056,
            "RemoteState": {
              "UUID": "5e20794b-0b73-4eb0-b49b-ece17089bf28",
              "Entries": 25861,
              "Size": 13181768
            }
          }
        ]
      }
    ],
    "Replicated": {
      "9a779454-95d8-457b-9841-aab9b93661fe": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "bc5ff11f-34a1-460c-adeb-4adc8c031777": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "f71ae8ea-5659-4ed2-8e4e-d7ebad4853c6": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ]
    }
  },
  "172.19.0.6:9404": {
    "UUID": "bc5ff11f-34a1-460c-adeb-4adc8c031777",
    "Wells": [
      {
        "Name": "default",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "default",
          "gravwell"
        ],
        "Shards": []
      },
      {
        "Name": "syslog",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "syslog",
          "kernel",
          "dmesg"
        ],
        "Shards": [
          {
            "Name": "76ba7",
            "Start": "2020-11-23T19:09:52Z",
            "End": "2020-11-25T07:34:24Z",
            "Entries": 24092,
            "Size": 12261596,
            "RemoteState": {
              "UUID": "bc5ff11f-34a1-460c-adeb-4adc8c031777",
              "Entries": 24092,
              "Size": 12261276
            }
          }
        ]
      }
    ],
    "Replicated": {
      "5e20794b-0b73-4eb0-b49b-ece17089bf28": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "9a779454-95d8-457b-9841-aab9b93661fe": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": [
            {
              "Name": "76ba7",
              "Start": "2020-11-23T19:09:52Z",
              "End": "2020-11-25T07:34:24Z",
              "Entries": 24253,
              "Size": 12359413,
              "RemoteState": {
                "UUID": "00000000-0000-0000-0000-000000000000",
                "Entries": 0,
                "Size": 0
              }
            }
          ]
        }
      ],
      "f71ae8ea-5659-4ed2-8e4e-d7ebad4853c6": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "raw",
          "Tags": [
            "raw"
          ],
          "Shards": [
            {
              "Name": "76ba7",
              "Start": "2020-11-23T19:09:52Z",
              "End": "2020-11-25T07:34:24Z",
              "Entries": 0,
              "Size": 4112,
              "RemoteState": {
                "UUID": "00000000-0000-0000-0000-000000000000",
                "Entries": 0,
                "Size": 0
              }
            }
          ]
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": [
            {
              "Name": "76ba7",
              "Start": "2020-11-23T19:09:52Z",
              "End": "2020-11-25T07:34:24Z",
              "Entries": 25794,
              "Size": 13163262,
              "RemoteState": {
                "UUID": "00000000-0000-0000-0000-000000000000",
                "Entries": 0,
                "Size": 0
              }
            }
          ]
        }
      ]
    }
  },
  "172.19.0.7:9404": {
    "UUID": "9a779454-95d8-457b-9841-aab9b93661fe",
    "Wells": [
      {
        "Name": "default",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "default",
          "gravwell"
        ],
        "Shards": [
          {
            "Name": "76ba7",
            "Start": "2020-11-23T19:09:52Z",
            "End": "2020-11-25T07:34:24Z",
            "Entries": 0,
            "Size": 4112,
            "RemoteState": {
              "UUID": "9a779454-95d8-457b-9841-aab9b93661fe",
              "Entries": 0,
              "Size": 4112
            }
          }
        ]
      },
      {
        "Name": "syslog",
        "Accelerator": "fulltext",
        "Engine": "index",
        "Tags": [
          "syslog",
          "kernel",
          "dmesg"
        ],
        "Shards": [
          {
            "Name": "76ba7",
            "Start": "2020-11-23T19:09:52Z",
            "End": "2020-11-25T07:34:24Z",
            "Entries": 24253,
            "Size": 12359637,
            "RemoteState": {
              "UUID": "9a779454-95d8-457b-9841-aab9b93661fe",
              "Entries": 24253,
              "Size": 12359413
            }
          }
        ]
      }
    ],
    "Replicated": {
      "5e20794b-0b73-4eb0-b49b-ece17089bf28": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": [
            {
              "Name": "76ba7",
              "Start": "2020-11-23T19:09:52Z",
              "End": "2020-11-25T07:34:24Z",
              "Entries": 25861,
              "Size": 13181768,
              "RemoteState": {
                "UUID": "00000000-0000-0000-0000-000000000000",
                "Entries": 0,
                "Size": 0
              }
            }
          ]
        }
      ],
      "bc5ff11f-34a1-460c-adeb-4adc8c031777": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
        {
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ],
      "f71ae8ea-5659-4ed2-8e4e-d7ebad4853c6": [
        {
          "Name": "default",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "default",
            "gravwell"
          ],
          "Shards": []
        },
          "Name": "syslog",
          "Accelerator": "fulltext",
          "Engine": "index",
          "Tags": [
            "syslog",
            "kernel",
            "dmesg"
          ],
          "Shards": []
        }
      ]
    }
  }
}
```
</pre>
</details>
