# データエクスプローラーAPI

データエクスプローラーシステムは、ユーザーが独自のクエリを作成する必要がなく、クエリ内のデータの自動抽出とフィルタリングを提供します。 これは主にWebSocketAPIのメッセージを介してアクセスされますが、RESTエンドポイントもあります。

データエクスプローラシステムは、[autoextractor](extractors.md)定義を使用して、与えられたタグ内のデータがどのように解析されるべきかを決定します。タグに定義が存在しない場合、抽出生成REST APIを使用して自動抽出器の候補定義を作成することができます。定義が配置されると、クライアントは、検索ウェブソケット上の特別なコマンドを使用して、生の検索レンダラーから「濃縮された」エントリを要求することができます。

## データ構造

### 要素

エントリから抽出されたデータは、要素構造の配列として表されます。 各要素は、データの単一の「フィールド」、たとえばNetflowレコードの送信元アドレスを表します。 一部のモジュールは*ネストされた*要素を放出する可能性があることに注意してください。SubElementフィールドを参照してください。

*名前：要素のわかりやすい名前。
*パス：要素の完全なパス指定。例： jsonモジュールの場合は `foo.bar。[0]`。
*値：データから抽出されたもの（文字列、数値、IPアドレスなど）の値。
* SubElements：自然なツリーのような構造を持つデータ型の場合、この要素の「下」にある要素の（オプションの）配列。
*フィルター：要素に適用できる可能性のあるフィルターのリスト。 "！="、 "〜"、 ">"。

```
interface Element {
	Name:        string;
	Path:        string;
	Value:       string | boolean | number;
	SubElements: Array<Element> | null;
	Filters:     Array<string> | null;
}
```

### エクスペリエンス

ExploreResult型は、特定のエントリから引き出された要素のセットを返すために使用されます。これには、抽出を生成したモジュールの名前 (例: "json") とエントリのタグの文字列名 (例: "syslog") も含まれます。

```
interface ExploreResult {
	Elements:	Array<Element>;
	Module:	string;
	Tag:		string;
}
```


### RESTエンドポイントのリクエスト/レスポンス
次の構造がRESTエンドポイントに使用されます：

```
interface GenerateAXRequest {
	Tag:     string;
	Entries: Array<SearchEntry>;
}
```

```
interface GenerateAXResponse {
	Extractor:	AXDefinition;
	Entries:	Array<SearchEntry>;
	Explore:	Array<ExploreResult>;
}
```

AXDefinitionタイプの説明については、[自動抽出器](autoextractors.md)のドキュメントを参照してください。

### フィルタリクエスト

フィルタリクエストタイプは、検索クエリにフィルターを追加するために使用されます。フィルタの配列は、ParseSearchRequestまたはStartSearchRequestオブジェクトにアタッチできます。フィルタは自動的にクエリに挿入されます。

* タグ：フィルタリングするタグ名（通常はExploreResultオブジェクトから取得）。
* モジュール：フィルタリングするモジュール名（通常はExploreResultオブジェクトから取得）。
* パス：フィルタリングする要素のパス（要素オブジェクトのパスフィールドにあります）。
* Op：使用するオプションのフィルター操作（ElementオブジェクトのFilters配列にあります）。
* 値：フィルタリングするオプションの値（ElementオブジェクトのValueフィールドにあります）。

OpとValueが指定されていない場合、「フィルター」は、指定された要素をフィルタリングするのではなく、明示的に抽出するだけであることに注意してください。

```
interface FilterRequest {
	Tag:    string;
	Module: string;
	Path:   string;
	Op:     string;
	Value:  string;
}
```

## 抽出生成RESTエンドポイント

エントリをデータエクスプローラで解析する前に、タグに自動抽出定義をインストールする必要があります。 抽出生成エンドポイントは、タグ名と1つ以上のエントリを受け取り、可能な抽出のコレクションを返します。 データ探索モジュールごとに1つの可能な抽出を返します。 ユーザーは、最も適切な抽出を選択し、[対応する自動抽出定義をインストールする](autoextractors.md)必要があります。

エンドポイントは`/api/explore/generate`にあります。 GenerateAXRequestを含む本文でPOSTリクエストを実行します。 サーバーは、（文字列）モジュール名をGenerateAXResponseオブジェクトにマッピングして応答します。各オブジェクトは、その特定のモジュールによって生成された抽出を表します。 各GenerateAXResponseオブジェクトには、Entries配列のSearchEntryごとにExplore配列に1つのExploreResultオブジェクトが含まれています。

次のリクエストには、1つのエントリが含まれています：

```
{
  "Tag": "foo",
  "Entries": [
    {
      "TS": "2020-11-02T16:58:56.717034109-07:00",
      "SRC": "",
      "Tag": 0,
      "Data": "ewogICJUUyI6ICIyMDIwLTEwLTE0VDEwOjM1OjQxLjEzNjUyMjg4NC0wNjowMCIsCiAgIlByb3RvIjogInVkcCIsCiAgIkxvY2FsIjogIls6Ol06NTMiLAogICJSZW1vdGUiOiAiNzMuNDIuMTA3LjE4MTo0Nzc0MiIsCiAgIlF1ZXN0aW9uIjogewogICAgIkhkciI6IHsKICAgICAgIk5hbWUiOiAicG9ydGVyLmdyYXZ3ZWxsLmlvLiIsCiAgICAgICJScnR5cGUiOiAxLAogICAgICAiQ2xhc3MiOiAxLAogICAgICAiVHRsIjogNjUsCiAgICAgICJSZGxlbmd0aCI6IDQKICAgIH0sCiAgICAiQSI6ICIyMDguNzEuMTQxLjM0IiwKCSJOb25zZW5zZSI6IFsgMTAwLCAyMDAsIDMwMCBdLAoJIk1vcmVOb25zZW5zZSI6IFsgeyJmb28iOiAiYmFyIn0gXQogIH0KfQ==",
      "Enumerated": null
    }
  ]
}

```

以下は、上記のリクエストに対する応答の例です。 簡潔にするために、1つのモジュール（"json"）からの結果のみが含まれ、Elements配列は短縮されています。

```
{
	"json": [
		{
			"Extractor": {
				"Name": "foo",
				"Desc": "Auto-generated JSON extraction for tag foo",
				"Module": "json",
				"Params": "TS Proto Local Remote Question",
				"Tag": "foo",
				"Labels": null,
				"UID": 0,
				"GIDs": null,
				"Global": false,
				"UUID": "00000000-0000-0000-0000-000000000000",
				"LastUpdated": "0001-01-01T00:00:00Z"
			},
			"Entries": [
				{
					"TS": "2020-11-02T16:58:56.717034109-07:00",
					"SRC": "",
					"Tag": 0,
					"Data": "ewogICJUUyI6ICIyMDIwLTEwLTE0VDEwOjM1OjQxLjEzNjUyMjg4NC0wNjowMCIsCiAgIlByb3RvIjogInVkcCIsCiAgIkxvY2FsIjogIls6Ol06NTMiLAogICJSZW1vdGUiOiAiNzMuNDIuMTA3LjE4MTo0Nzc0MiIsCiAgIlF1ZXN0aW9uIjogewogICAgIkhkciI6IHsKICAgICAgIk5hbWUiOiAicG9ydGVyLmdyYXZ3ZWxsLmlvLiIsCiAgICAgICJScnR5cGUiOiAxLAogICAgICAiQ2xhc3MiOiAxLAogICAgICAiVHRsIjogNjUsCiAgICAgICJSZGxlbmd0aCI6IDQKICAgIH0sCiAgICAiQSI6ICIyMDguNzEuMTQxLjM0IiwKCSJOb25zZW5zZSI6IFsgMTAwLCAyMDAsIDMwMCBdLAoJIk1vcmVOb25zZW5zZSI6IFsgeyJmb28iOiAiYmFyIn0gXQogIH0KfQ==",
					"Enumerated": null
				}
			],
			"Explore": [
				{
					"Elements": [
						{
							"Name": "TS",
							"Path": "TS",
							"Value": "2020-10-14T10:35:41.136522884-06:00",
							"Filters": [
								"==",
								"!="
							]
						},
						{
							"Name": "Proto",
							"Path": "Proto",
							"Value": "udp",
							"Filters": [
								"==",
								"!="
							]
						},
						{
							"Name": "Local",
							"Path": "Local",
							"Value": "[::]:53",
							"Filters": [
								"==",
								"!="
							]
						},
						{
							"Name": "Remote",
							"Path": "Remote",
							"Value": "73.42.107.181:47742",
							"Filters": [
								"==",
								"!="
							]
						},
						{
							"Name": "Question",
							"Path": "Question",
							"Value": "{\n    \"Hdr\": {\n      \"Name\": \"porter.gravwell.io.\",\n      \"Rrtype\": 1,\n      \"Class\": 1,\n      \"Ttl\": 65,\n      \"Rdlength\": 4\n    },\n    \"A\": \"208.71.141.34\",\n\t\"Nonsense\": [ 100, 200, 300 ],\n\t\"MoreNonsense\": [ {\"foo\": \"bar\"} ]\n  }",
							"SubElements": [
								{
									"Name": "Question.A",
									"Path": "Question.A",
									"Value": "208.71.141.34",
									"Filters": [
										"==",
										"!="
									]
								},

							],
							"Filters": [
								"==",
								"!="
							]
						}
					],
					"Module": "json",
					"Tag": "foo"
				}
			]
		}
	]
}
```

ネストの例については、`Question`要素の`SubElements`フィールドに注意してください。

## ソケットコマンドの検索

`raw`および`text`レンダラーは、2つの追加のWebSocketコマンド、 `REQ_GET_EXPLORE_ENTRIES`および`REQ_EXPLORE_TS_RANGE`を実装します。 これらのコマンドは、それぞれREQ_GET_ENTRIESコマンドとREQ_TS_RANGEコマンドを反映していますが、データエクスプローラーコマンドへの応答には、エントリごとに1つずつ、ExploreResultオブジェクト（上記で定義）の配列を含む `Explore`という名前のフィールドが含まれます。
たとえば、このコマンドは最初の10エントリを要求します：

```
{
	"ID": 61456,
	"EntryRange": {
		"First": 0,
		"Last": 10
	}
}
```

エントリを1つだけ返した検索の応答例を次に示します：

```
{
  "ID": 61456,
  "EntryCount": 1,
  "AdditionalEntries": false,
  "Finished": true,
  "Entries": [
    {
      "TS": "2020-11-02T16:58:56.717034109-07:00",
      "SRC": "",
      "Tag": 0,
      "Data": "ewogICJUUyI6ICIyMDIwLTEwLTE0VDEwOjM1OjQxLjEzNjUyMjg4NC0wNjowMCIsCiAgIlByb3RvIjogInVkcCIsCiAgIkxvY2FsIjogIls6Ol06NTMiLAogICJSZW1vdGUiOiAiNzMuNDIuMTA3LjE4MTo0Nzc0MiIsCiAgIlF1ZXN0aW9uIjogewogICAgIkhkciI6IHsKICAgICAgIk5hbWUiOiAicG9ydGVyLmdyYXZ3ZWxsLmlvLiIsCiAgICAgICJScnR5cGUiOiAxLAogICAgICAiQ2xhc3MiOiAxLAogICAgICAiVHRsIjogNjUsCiAgICAgICJSZGxlbmd0aCI6IDQKICAgIH0sCiAgICAiQSI6ICIyMDguNzEuMTQxLjM0IiwKCSJOb25zZW5zZSI6IFsgMTAwLCAyMDAsIDMwMCBdLAoJIk1vcmVOb25zZW5zZSI6IFsgeyJmb28iOiAiYmFyIn0gXQogIH0KfQ==",
      "Enumerated": null
    }
  ],
  "Explore": [
    {
      "Elements": [
        {
          "Name": "TS",
          "Path": "TS",
          "Value": "2020-10-14T10:35:41.136522884-06:00",
          "Filters": [
            "==",
            "!="
          ]
        },
        {
          "Name": "Proto",
          "Path": "Proto",
          "Value": "udp",
          "Filters": [
            "==",
            "!="
          ]
        },
        {
          "Name": "Local",
          "Path": "Local",
          "Value": "[::]:53",
          "Filters": [
            "==",
            "!="
          ]
        },
        {
          "Name": "Remote",
          "Path": "Remote",
          "Value": "73.42.107.181:47742",
          "Filters": [
            "==",
            "!="
          ]
        },
        {
          "Name": "Question",
          "Path": "Question",
          "Value": "{\n    \"Hdr\": {\n      \"Name\": \"porter.gravwell.io.\",\n      \"Rrtype\": 1,\n      \"Class\": 1,\n      \"Ttl\": 65,\n      \"Rdlength\": 4\n    },\n    \"A\": \"208.71.141.34\",\n\t\"Nonsense\": [ 100, 200, 300 ],\n\t\"MoreNonsense\": [ {\"foo\": \"bar\"} ]\n  }",
          "SubElements": [
            {
              "Name": "Question.A",
              "Path": "Question.A",
              "Value": "208.71.141.34",
              "Filters": [
                "==",
                "!="
              ]
            }
          ],
          "Filters": [
            "==",
            "!="
          ]
        }
      ],
      "Module": "json",
      "Tag": "foo"
    }
  ]
}

```

## Adding Filters to Queries

データエクスプローラレンダラーコマンドから返された情報を使用して、*フィルターリクエスト*を作成できます。フィルタリクエストは、データ内の値に基づいてクエリの結果を絞り込みます。 たとえば、上記の応答例では、「Proto」フィールドが「udp」であるすべてのエントリを除外したい場合があります。 次のFilterRequestは、そのフィルターを実装します。

```
{
  "Tag": "foo",
  "Module": "json",
  "Path": "Proto",
  "Op": "!=",
  "Value": "udp"
}
```

以下は、フィルターを含むStartSearchRequestです：

```
{
  "SearchString": "tag=foo",
  "SearchStart": "2020-01-01T12:01:00.0Z07:00",
  "SearchEnd": "2020-01-01T12:01:30.0Z07:00",
  "Filters": [
    {
      "Tag": "foo",
      "Module": "json",
      "Path": "Proto",
      "Op": "!=",
      "Value": "udp"
    }
  ]
}

```

サーバーの応答には、必要に応じて検索ライブラリに保存するのに適した、書き直されたSearchStringフィールドが含まれます：

```
{
  "SearchString": "tag=foo json Proto!=udp",
  "RawQuery": "tag=foo",
  "RenderModule": "text",
  "RenderCmd": "text",
  "OutputSearchSubproto": "search484",
  "OutputStatsSubproto": "stats484",
  "SearchID": "7947106109",
  "SearchStartRange": "2015-01-01T12:01:00.0Z07:00",
  "SearchEndRange": "2015-01-01T12:01:30.0Z07:00",
  "Background": false
}
```

フィルタリクエストをParseSearchRequestメッセージに添付して、フィルタを検証することもできます。