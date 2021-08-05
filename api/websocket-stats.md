# ステータスを取得する

/ws/statsにあるWebsocket

サーバーは、RoutingWebsocketプロトコルを介してすぐに統計情報を投げ始めます。 以下のタイプが登録されることを期待し、それぞれのデータを生成します：

* ping
* idxStats
* sysStats
* sysDesc

接続すると、websocketは最初にsysDescタイプで単一のJSONパッケージをスローします（常に最初に）。 これは、統計情報を取得し続けるバックエンドシステムの数を列挙するために確実に使用できます。 その最初のパケットを使用して、テーブルを構築しますが、idxStatsおよびsysStatsタイプを介して継続的に入力します。 pingタイプはキープアライブであり、30秒後にpingタイプから更新を取得しない場合、バックエンドは停止しています。ユーザーに知らせてください。

# タイプの例

## システム記述サブプロトコルsysDesc

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

## JSONのpingの例
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

## sysStats JSONパケットの例
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


## インデクサーの統計JSONパケットの例

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