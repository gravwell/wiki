# REST Stats API

## Ping

ping stats APIは、Webサーバーとインデクサーのステータスを返します。 これは、GETを介して`/api/stats/ping`に取得されます。 すべてのシステムが稼働している場合、それらは「OK」として報告されます。

```
{"192.168.2.60:9404"： "OK"、 "webserver"： "OK"}
```

インデクサがダウンしている場合は、「切断済み」とマークされます。

```
{"192.168.2.60:9404"： "Disconnected"、 "webserver"： "OK"}
```

## インデックス統計

インデクサー統計APIは、各インデクサーのインデックスに関する情報を提供します。 GETを介して `/api/stats/idxStats`にアクセスします。

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

## インジェスター統計

GETリクエストを`/api/stats/igstStats`に送信すると、各インデクサーに接続されているインジェスターを説明する構造が返されます。 以下の例は、2つのインジェスターが接続された単一のインデクサー（192.168.2.60）を示しています。SimpleRelayingesterとNetworkCaptureingesterです。

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

各ingesterの説明で、「Count」フィールドは取り込まれたエントリの数を示し、「Size」フィールドは取り込まれたバイト数を示します。 「稼働時間」とは、ナノ秒単位で、インジェスターが接続されている時間のことです。

「QuotaMax」フィールドと「QuotaUsed」フィールドに注意してください。 コミュニティライセンスは、毎日特定の量しか摂取できません。 「QuotaMax」は、特定のインデクサーが1日に取り込むことができるバイト数を指定します。 「QuotaUsed」は、今日までに取り込まれたバイト数を示します。

## シ ステム統計

GETリクエストを`/api/stats/sysStats`に送信すると、CPUの数、CPU使用率、メモリとネットワークの使用率など、インデクサーとウェブサーバーシステムに関する情報を含む構造が返されます。

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

ほとんどのフィールドは自明です。 「IO」アレイはディスクに関する情報を提供し、「読み取り」フィールドと「書き込み」フィールドは読み取りと書き込みの*レート*をバイト/秒で指定することに注意してください。 同様に、「ネット」コンポーネントは、ネットワーク使用率を1秒あたりのバイト数で表します。

## システムの説明

GETリクエストを`/api/stats/sysDesc`に送信すると、ウェブサーバーとインデクサーのホストシステムに関する追加情報を提供する構造が返されます。

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


## シャードストレージとレプリケーションの統計

インデクサーは、すべてのシャードとウェルのリストを維持し、Gravwell内に保存されているデータ全体のシャードレベルのビューを生成できます。 このビューは、長期間にわたる井戸とデータ量の非常に迅速な規模の観察を提供できます。

高可用性モードの場合、インデクサーは複製されたデータのマッピングも維持し、データが複製された場所も解決できるため、他のインデクサーに複製しているインデクサーの概要をすばやく確認できます。

シャードレベルのビューは、`/api/indexers/info`への`GET`リクエストを介してアクセスされ、各インデクサーのJSONマップを返します。 返されるデータセットには、ウェルの構成、それに割り当てられているタグ、およびウェル内に配置されているシャードに関する広範な情報が含まれています。

### JSON応答の例
これは、それぞれ2つのウェルがあり、レプリケーションが有効になっている4つのGravwellインデクサーのクラスターからの応答の例です。 Syslogウェルには単一のシャードのみが入力されます。

<details><summary>JSON応答を展開します</summary>
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
