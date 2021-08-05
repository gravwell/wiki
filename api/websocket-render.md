# レンダーモジュール

レンダリングモジュールは、検索に関する情報、検索の実際のエントリ、検索に関する統計を取得するためのAPIを提供します。 APIコマンドは、検索の起動時に確立されたwebsocketサブプロトコルを介してJSONとして送信されます。 サブプロトコルと検索の起動については、[検索APIドキュメント]（＃！api / websocket-search.md）を参照してください。

これらの記事の内容を超えて、レンダリングモジュールコマンドの動作を確認する最も簡単な方法は、Webブラウザーのコンソール（ChromeのF12）を使用することです。 Websocketトラフィックを示すセクションを見つけ、送受信されたメッセージを参照します。

![](webconsole.png)

websocketトラフィックを監視する別の方法は、`-debug`フラグを指定して[Gravwell CLI client]（＃！cli / cli.md）を実行することです。 フラグは引数としてファイル名を取ります。 websocketとの間で送受信されるJSONメッセージは、そのファイルに書き込まれます。

## すべてのレンダーモジュールに共通の操作ID

すべてのレンダーモジュールは、次のリクエストに応答します。
| リクエスト名 | 16進値| 10進値 | 説明 |
|--------------|-----------|---------------|-------------|
| REQ_CLOSE				| 0x1			| 1			| Close the channel |
| REQ_ENTRY_COUNT		| 0x3			| 3			| Get the number of entries seen|
| REQ_SEARCH_DETAILS	| 0x4			| 4			| Get detailed info about the search |
| REQ_SEARCH_TAGS		| 0x5			| 5			| Get the tag map used in the search |
| REQ_GET_ENTRIES		| 0x10			| 16		| Request a block of search entries by index |
| REQ_STREAMING			| 0x11			| 17		| Request that search entries be sent as they come in |
| REQ_TS_RANGE			| 0x12			| 18		| Request a block of entries by time range |
| REQ_STATS_SIZE		| 0x7F000001	| 2130706433| Request the size of the statistics data |
| REQ_STATS_RANGE		| 0x7F000002	| 2130706434| Request the time range of available stats |
| REQ_STATS_GET			| 0x7F000003	| 2130706435| Request stats |
| REQ_STATS_GET_RANGE	| 0x7F000004	| 2130706436| Request stats from a particular time range |
| REQ_STATS_GET_SUMMARY	| 0x7F000005	| 2130706437| Request a summary of statistics |

応答値はリクエストと同じですが、特別な応答コード `RESP_ERROR`（0xFFFFFFFF）が追加されています。

| リクエスト名 | 16進数 | 10進数 | 説明 |
|--------------|-----------|---------------|-------------|
| RESP_ERROR				| 0xFFFFFFFF	| 4294967295| (error) |
| RESP_CLOSE				| 0x1			| 1			| Socket will be closed |
| RESP_ENTRY_COUNT			| 0x3			| 3			| Returning # of entries |
| RESP_SEARCH_DETAILS		| 0x4			| 4			| Returning info about search |
| RESP_SEARCH_TAGS			| 0x5			| 5			| Returning tag map for search |
| RESP_GET_ENTRIES			| 0x10			| 16		| Returning search entries|
| RESP_STREAMING			| 0x11			| 17		| Search entries will be streamed|
| RESP_TS_RANGE				| 0x12			| 18		| Returning a block of entries for a time range |
| RESP_STATS_SIZE			| 0x7F000001	| 2130706433| Returning size of stats data |
| RESP_STATS_RANGE			| 0x7F000002	| 2130706434| Returning the time range for stats |
| RESP_STATS_GET			| 0x7F000003	| 2130706435| Returning stats |
| RESP_STATS_GET_RANGE		| 0x7F000004	| 2130706436| Returning stats for a time range |
| RESP_STATS_GET_SUMMARY	| 0x7F000005	| 2130706437| Returning stats summary |


API要求は、検索の作成中に確立された検索サブプロトコルを介してJSONとして送信する必要があります。

## 応答形式

すべてのモジュールは同じコマンドに応答しますが、関連するデータ型の性質が異なるため、エントリを返す形式は異なります。 [Renderer API応答形式ドキュメント]（websocket-render-responses.md）は、各レンダリングモジュールで使用される応答形式について説明しています。 この記事の残りの部分で示す例は、テキストおよびテーブルレンダラーからの応答を示しています。

## チャンネルを閉じる（リクエスト0x1）

検索への現在の接続を閉じるには、次の構造をwebsocketに送信します。

```
{
        "ID": 1
}
```

ソケットは応答するはずです：

```
{
        "ID": 1,
        "EntryCount": 0,
        "AdditionalEntries": false,
        "Finished": false,
        "Entries": []
}
```

## エントリ数を取得する（リクエスト0x3）

REQ_ENTRY_COUNTコマンドは、レンダーモジュールに到達したエントリの総数を要求します。

```
{
        "ID": 3
}
```

この例では、モジュールは41個のエントリがあると応答しました。

```
{
        "ID": 3,
        "EntryCount": 41,
        "AdditionalEntries": false,
        "Finished": true
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
}
```

重要：ここで報告される数は、**レンダラーが受け取ったエントリの総数**です。 REQ_GET_ENTRIESなどの他のコマンドを使用すると、返されるエントリが少なくなる場合があります。 これは、一部のレンダラー（テーブルやチャートなど）が*凝縮*レンダラーであるためです。

## 検索の詳細を取得する（リクエスト0x4）

REQ_SEARCH_DETAILSコマンド（0x4）は、検索自体に関する情報を要求します。

```
{
        "ID": 4
}
```

応答には、統計情報と検索自体に関する情報が含まれます。

```
{
	"ID": 4,
	"Stats": {
		"Size": 0,
		"Set": [
			{
				"TS": "2018-04-02T07:56:47.422345249-06:00",
				"Stats": []
			},
<some entries elided for brevity>
			{
				"TS": "2018-04-02T10:17:24.922345249-06:00",
				"Stats": [
					{
						"Name": "packet",
						"Args": "packet ipv4.SrcIP",
						"InputCount": 1619,
						"OutputCount": 1619,
						"InputBytes": 2034920,
						"OutputBytes": 2052729,
						"Duration": 54936015
					},
					{
						"Name": "count",
						"Args": "count by SrcIP",
						"InputCount": 1619,
						"OutputCount": 88,
						"InputBytes": 2052729,
						"OutputBytes": 28931,
						"Duration": 129787584
					}
				]
			},
<some entries elided for brevity>
			{
				"TS": "2018-04-02T12:47:24.922345249-06:00",
				"Stats": []
			},
			{
				"TS": "2018-04-02T12:52:06.172345249-06:00",
				"Stats": [
					{
						"Name": "packet",
						"Args": "packet ipv4.SrcIP",
						"InputCount": 1,
						"OutputCount": 1,
						"InputBytes": 99,
						"OutputBytes": 110,
						"Duration": 14480829
					},
					{
						"Name": "count",
						"Args": "count by SrcIP",
						"InputCount": 1,
						"OutputCount": 1,
						"InputBytes": 110,
						"OutputBytes": 110,
						"Duration": 52235061
					}
				]
			}
		],
		"RangeStart": "2018-04-02T07:56:47.422345249-06:00",
		"RangeEnd": "2018-04-02T12:56:47.422345249-06:00",
		"Current": "2018-04-02T07:56:47.422345249-06:00"
	},
	"SearchInfo": {
		"ID": "677124412",
		"UID": 1,
		"UserQuery": "tag=pcap packet ipv4.SrcIP | count by SrcIP | table SrcIP count",
		"EffectiveQuery": "tag=pcap packet ipv4.SrcIP | count by SrcIP | table SrcIP count",
		"StartRange": "2018-04-02T07:56:47.422345249-06:00",
		"EndRange": "2018-04-02T12:56:47.422345249-06:00",
		"Descending": true,
		"Started": "2018-04-02T12:56:47.430572676-06:00",
		"LastUpdate": "0001-01-01T00:00:00Z",
		"StoreSize": 75960,
		"IndexSize": 32,
		"ItemCount": 1575,
		"TimeZoomDisabled": false,
		"RenderDownloadFormats": [
			"json",
			"csv",
			"lookupdata"
		],
		"Duration": "0s"
	},
	"EntryCount": 1575,
	"AdditionalEntries": false,
	"Finished": true
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},

}
```

## タグマップの取得（リクエスト0x5）

タグマップを要求するには、リクエストID 0x5を標準のwebsocketサブプロトコルに送信します。

```
{
        "ID": 0x5,
}
```

応答は次のようになり、タグ名と数値タグインデックスのマップが表示されます。

```
{
        "ID": 0x5,
        "Tags": {
                "default": 0,
				"tagmctaggy": 1,
				"apache": 2,
				"syslog: 3,
				"gravwell": 4
			},
}
```

## エントリの取得（リクエスト0x10）

要求0x10（10進数の16）は、レンダラーにエントリのブロックを要求します。 「最初の」フィールドと「最後の」フィールドを使用して、必要なエントリを指定します。 レンダラーは、REQ_ENTRY_COUNT（0x3）リクエストへの応答として、エントリの数を報告します。

このコマンドは最初の1024エントリを要求します（指定しない場合、`First`はデフォルトで0になります）：

```
{
	"ID": 16,
	"EntryRange": {
		"Last": 1024
	}
}
```

サーバーは、エントリの配列と追加情報で応答します。

```
{
	"ID": 16,
	"EntryCount": 1575,
	"AdditionalEntries": false,
	"Finished": true,
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
	"Entries": {
		"Rows": [
			{
				"TS": "2018-04-02T10:30:29-06:00",
				"Row": [
					"10.144.162.236",
					"9410"
				]
			},
<67 similar entries elided>
			{
				"TS": "2018-04-02T10:28:51-06:00",
				"Row": [
					"192.168.1.1",
					"2"
				]
			}
		],
		"Columns": [
			"SrcIP",
			"count"
		]
	}
}
```

`"AdditionalEntries"：false`フィールドに注意してください。 これは、読み込むエントリがこれ以上ないことを意味します。 このフィールドがtrueに戻った場合、「First」を1024に、「Last」を2048に設定してコマンドを再発行することで、エントリがなくなるまで繰り返すことで、さらにエントリを読み取ることができます。

## ストリーミング結果のリクエスト（リクエスト0x11）

クライアントは、リクエストID 0x11を送信することで、レンダラーがwebsocketを介してできるだけ早くエントリを送信するようにリクエストできます。 通常、エントリがディスクまたは他の簡単な操作にすぐに書き込まれない限り、これは推奨されません。 レンダリングモジュールは、クライアントが処理できるよりも高速に結果を頻繁に出力できます。 REQ_GET_ENTRIESコマンドを使用して、ブロックごとに結果を取得することをお勧めします。

ストリーミングをリクエストするには：

```
{
	"ID": 17
}
```

レンダラーは、できる限り早くエントリの大きなブロックの送信を開始します。

```
{
	"ID": 17,
	"EntryCount": 1000,
	"AdditionalEntries": false,
	"Finished": false,
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
	"Entries": [
<1000 entries elided>
	]
}
<many other entry blocks elided>
{
	"ID": 17,
	"EntryCount": 861,
	"AdditionalEntries": false,
	"Finished": false,
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
	"Entries": [
<861 entries elided>
	]
}
{
	"ID": 17,
	"EntryCount": 0,
	"AdditionalEntries": false,
	"Finished": true,
	"Entries": []
}
```

この場合、レンダラーは1000エントリの多くのブロック、残りの861ブロックを含むブロック、および0エントリを含む最終ブロックを送信しました。 0エントリのこの最後のブロックは、それ以上エントリが送信されないことを示します。

注意：ストリーミング結果を有効にするときは細心の注意を払ってください。

## 特定の時間範囲のエントリを要求する（要求0x12）

リクエスト0x12（REQ_TS_RANGE）を使用して、検索範囲の特定の部分のエントリを取得します。 多数のエントリが存在する可能性があるため、REQ_GET_ENTRIESの場合と同様に `First`および` Last`フィールドを使用して、一度にブロックをフェッチします。この場合、最初の100エントリです。

```
{
	"ID":18,
	"EntryRange": {
		"First":0,
		"Last":100,
		"StartTS":"2018-04-02T16:19:51.579Z",
		"EndTS":"2018-04-02T16:42:28.649Z"
	}
}
```

サーバーは、要求された時間内に収まるエントリで応答します。

```
{
	"ID":18,
	"EntryCount":1575,
	"AdditionalEntries":false,
	"Finished":true,
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
	"Entries": {
		"Rows": [
			{
				"TS":"2018-04-02T10:30:29-06:00",
				"Row":["10.194.162.236","9410"]
			},
			{
				"TS":"2018-04-02T10:30:36-06:00",
				"Row":["192.168.1.101","8212"]
			},
<entries elided>
			{
				"TS":"2018-04-02T10:28:51-06:00",
				"Row":["192.168.1.1","2"]
			}
		],
		"Columns":["SrcIP","count"]
	}
}
```

この場合、エントリは100未満でした。 さらにある場合、 `AdditionalEntries`フィールドは` true`に設定されます。 同じタイムスタンプで別のリクエストを送信し、 `First`フィールドを100に、` Last`フィールドを200に変更して次の100エントリのブロックを取得することで、より多くのエントリを取得できます。

## 現在の統計エントリカウントを取得します（リクエスト0x7F000001）

検索を実行すると、統計エントリが生成されます。 エントリの数は、0x7F000001（REQ_STATS_SIZE）要求を使用して取得できます。

```
{
        "ID": 2130706433,
}
```

サーバーは、統計エントリの数で応答します。

```
{
        "ID": 2130706433,
        "Stats": {
			"Size": 466
		}
}
```

## 統計でカバーされる現在の時間範囲を取得します（リクエスト0x7F000002）

このコマンドは、検索統計が利用可能な時間範囲を返します。 コマンドを送信します。

```
{
        "ID": 2130706434,
}
```

サーバーが応答します。

```
{
        "ID": 2130706434,
        "Stats": {
                "RangeStart": "2016-09-02T08:59:37.943271552-06:00",
                "RangeEnd": "2016-09-02T08:59:37.943271552-06:00",
                "Size": 2
        },
        "EntryCount": 510000
}
```

戻されたSizeパラメーターは、使用可能な最大粒度を示します。 この例では、発行された検索は2秒しかカバーしなかったため、クライアントは最大2の粒度を要求できます（または、データの1秒ごとに1つの統計エントリ）。 検索が「メモリとストレージで許可されているように」無制限の粒度を可能にする「フル解像度の統計」で発行されない限り、ウェブサーバーは最大65kの統計エントリを保持します。

## 統計セットの要求（リクエスト0x7F000003）

統計セットは、必要な「チャンク」の数を指定することで要求されます。 たとえば、検索が1か月分のデータにわたって実行された場合、65kの統計セットがありますが、表示のために、10、100、または1000の増分でそれらの統計に異なる粒度を表示することができます。

さまざまな粒度を取得するには、StatsRequestでSetCountを送信します。 この例では、サイズ1のセットを要求します（すべてのモジュールを要約します）：

```
{
        "ID": 2130706435,
        "Stats": {
                "SetCount": 1
        }
}
```

応答には、単一の統計エントリのみが含まれます。

```
{
        "ID": 2130706435,
        "Stats": {
                "RangeStart": "0001-01-01T00:00:00Z",
                "RangeEnd": "0001-01-01T00:00:00Z",
                "Set": [
                        {
                                "Stats": [
                                        {
                                                "Name": "grep",
                                                "Args": "grep HEE",
                                                "InputCount": 510000,
                                                "OutputCount": 510000,
                                                "InputBytes": 52363340,
                                                "OutputBytes": 52363340,
                                                "Duration": 0
                                        },
                                        {
                                                "Name": "sort",
                                                "Args": "sort by time",
                                                "InputCount": 510000,
                                                "OutputCount": 500000,
                                                "InputBytes": 52363340,
                                                "OutputBytes": 51344450,
                                                "Duration": 0
                                        }
                                ],
                                "TS": "2016-09-02T08:59:37.943271552-06:00"
                        }
                ],
                "Size": 2
        },
        "OverLimit":false,
        "LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
        "EntryCount": 510000
}
```

## 特定の時間範囲にわたる統計セットを要求する（リクエスト0x7F000006）

この例では、2016-09-09T06：02：14Zと2016-09-09T06：02：16Zの間にサイズ1のセット（すべてのモジュールを要約する）を要求します。 SetCount番号は、要求された範囲全体で均一な「ChunkSize」を生成するために使用されることに注意することが重要です。 たとえば、検索に2016-01-09T06：02：16Zから2016-12-09T06：02：16Z（11か月）のデータがあるが、1901-01-01T06：02：16Zから2016-01の範囲を指定した場合 -01T06：02：16ZおよびSetSizeが100の場合、その範囲とサイズの「ChunkSize」は実際にデータがある期間よりもはるかに大きいため、1つの統計サイズのみを取得します。

```
{
        "ID": 2130706436,
        "Stats": {
                "SetCount": 1,
				"SetStart": "2016-09-09T06:02:14Z",
				"SetEnd": "2016-09-09T06:02:16Z",
        }
}
```

応答：

```
{
        "ID": 2130706436,
        "Stats": {
                "RangeStart": "0001-01-01T00:00:00Z",
                "RangeEnd": "0001-01-01T00:00:00Z",
                "Set": [
                        {
                                "Stats": [
                                        {
                                                "Name": "grep",
                                                "Args": "grep HEE",
                                                "InputCount": 510000,
                                                "OutputCount": 510000,
                                                "InputBytes": 52363340,
                                                "OutputBytes": 52363340,
                                                "Duration": 0
                                        },
                                        {
                                                "Name": "sort",
                                                "Args": "sort by time",
                                                "InputCount": 510000,
                                                "OutputCount": 500000,
                                                "InputBytes": 52363340,
                                                "OutputBytes": 51344450,
                                                "Duration": 0
                                        }
                                ],
                                "TS": "2016-09-09T06:02:14.943271552-06:00"
                        }
                ],
                "Size": 100
        },
        "OverLimit":false,
        "LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
        "EntryCount": 500000
}
```

## 統計情報の要約を要求する（リクエスト0x7F000005）

統計セット概要のリクエストは、サイズ1の統計セットのリクエストと同等です：

```
{
	"ID":2130706437
}
```

応答：

```
{
    "AdditionalEntries": false,
    "EntryCount": 1575,
    "Finished": true,
    "OverLimit":false,
    "LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
    "ID": 2130706437,
    "Stats": {
        "Set": [
            {
                "Stats": [
                    {
                        "Args": "packet ipv4.SrcIP",
                        "Duration": 61425123,
                        "InputBytes": 27364970,
                        "InputCount": 24886,
                        "Name": "packet",
                        "OutputBytes": 27636389,
                        "OutputCount": 24861
                    },
                    {
                        "Args": "count by SrcIP",
                        "Duration": 215555286,
                        "InputBytes": 27636389,
                        "InputCount": 24861,
                        "Name": "count",
                        "OutputBytes": 541861,
                        "OutputCount": 1575
                    }
                ],
                "TS": "2018-04-02T09:06:41.441-06:00"
            }
        ],
        "Size": 0
    }
}
```
## 検索用のメタデータを要求する（リクエスト 0x10001)

このメッセージは、パイプラインを通過した列挙型の値のハイレベルな調査を要求しています。以下に、サンプルのクエリと、生成されるメタデータの例を示します。

```
tag=syslog syslog Hostname Appname |
     length |
     stats sum(length) count by Hostname Appname |
     table Hostname Appname sum count
```


```
{
	"ID": 65537
}
```

Response:

```
{
	"ID": 65537,
	"EntryCount": 1575,
	"AdditionalEntries": false,
	"Finished": true,
	"OverLimit":false,
	"LimitDroppedRange":{"StartTS":"0000-12-31T16:07:02-07:52","EndTS":"0000-12-31T16:07:02-07:52"},
	"Metadata": {
		"ValueStats": [
			{
				"Name": "count",
				"Type": "number",
				"Number": {
					"Count": 963,
					"Min": 1,
					"Max": 3
				},
				"Raw": {
					"Map": null,
					"Other": 0
				}
			},
			{
				"Name": "Hostname",
				"Type": "raw",
				"Number": {
					"Count": 0,
					"Min": 0,
					"Max": 0
				},
				"Raw": {
					"Map": {
						"ant": 25,
						"tracker": 25,
						"voice": 22,
						"warrior": 31,
						"whale": 29
					},
					"Other": 0
				}
			},
			{
				"Name": "Appname",
				"Type": "raw",
				"Number": {
					"Count": 0,
					"Min": 0,
					"Max": 0
				},
				"Raw": {
					"Map": {
						"alpine": 4,
						"time": 2,
						"zenith": 5
					},
					"Other": 865
				}
			},
			{
				"Name": "length",
				"Type": "number",
				"Number": {
					"Count": 963,
					"Min": 314,
					"Max": 812
				},
				"Raw": {
					"Map": null,
					"Other": 0
				}
			},
			{
				"Name": "sum",
				"Type": "number",
				"Number": {
					"Count": 963,
					"Min": 314,
					"Max": 1397
				},
				"Raw": {
					"Map": null,
					"Other": 0
				}
			}
		],
		"SourceStats": [
			{
				"IP": "192.168.1.1",
				"Count": 963
			}
		],
		"TagStats": {
			"rawsyslog": 963
		}
	}
}
```
