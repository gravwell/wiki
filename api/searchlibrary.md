# 検索ライブラリ

REST API は/api/libraryにあります。

保存された検索クエリの保存・取得には、検索ライブラリAPIを使用します。検索ライブラリは、名前、メモ、タグなど、日々の業務で価値のあるクエリを蓄積するのに便利なシステムです。

検索ライブラリはパーミッション付きのAPIで、各ユーザーが検索ライブラリを所有し、オプションでグループと共有することができます。各ライブラリエントリにはグローバルフラグが設定されており、どのユーザーもそのライブラリエントリを読むことができます。グローバルフラグを設定できるのは管理者のみです。

所有者のみがライブラリエントリを削除することができ、他のユーザがグループメンバーシップを通じてライブラリエントリにアクセスできたとしても、そのエントリを削除することはできません。

## 管理者の業務

管理者は、他のすべてのユーザと同じ方法で検索ライブラリと対話します。管理者がライブラリAPIへの管理者アクセスを希望する場合、所有者以外のライブラリエントリをリストアップ、削除、変更するかどうかに関わらず、リクエストに `admin` フラグを追加しなければなりません。

例えば、`/api/library`に対してGETリクエストを実行すると、呼び出したユーザのライブラリエントリ(adminかどうか)のみが返ってくる。同じGETリクエストを`/api/library?admin=true`で実行すると、すべてのユーザのすべてのライブラリエントリが返される。adminフラグは管理者でないユーザには無視されます。

## 基本APIの概要

LibraryAPIは`/api/library`をルートとし、次のリクエストメソッドに応答します。

| メソッド| 説明| 管理者の呼び出しをサポートします。|
| ------ | ----------- | -------------------- |
| GET | 利用可能なライブラリエントリのリストを取得する| TRUE |
| POST | 新しいライブラリエントリを追加する| FALSE |
| PUT| 既存のライブラリエントリを更新する| TRUE |
| DELETE| 既存のライブラリエントリを削除する| TRUE |

すべての検索ライブラリ エントリには、それに関連付けられた"ThingUUID" と"GUID"の両方があります。ThingUUID」は常に一意であり、システム上には、指定されたThingUUIDを持つ検索ライブラリ項目は1つしか存在しません。一方、「GUID」は、ダッシュボードやアクショナブルなどの他のものから検索ライブラリ エントリを参照するために使用される ID です。キットをインストールすると、すべての検索ライブラリエントリがキット作成者のシステム上のものと同じ GUID を持つようになり、相互リンクが可能になります。各ユーザーが同じ GUID を持つ検索ライブラリエントリを持つ可能性がありますが、それらのエントリはそれぞれ固有の ThingUUID を持ちます。

`DELETE`および`PUT`メソッドでは、特定のライブラリエントリの「ThingUUID」または「GUID」をURLに追加する必要があります。ThingUUIDとGUIDはAPIで交換可能に使用できますが、「管理モード」では、同じGUIDを持つ複数のアクセス可能なアイテムが存在する可能性があることに注意してください。たとえば、`5f72d51e-d641-11e9-9f54-efea47f6014a`のGUIDでエントリを更新するには、`/api/library/5f72d51e-d641-11e9-9f54-efea47f6014a`に対して`PUT`リクエストが発行されます。リクエストの本文にエンコードされたエントリ構造です。

`GET`メソッドは、オプションでGUIDを追加して、特定のライブラリエントリを要求できます。GUIDが存在しない場合、GETメソッドは使用可能なすべてのエントリのリストを返します。 ユーザーがGUIDで指定された特定のエントリにアクセスできない場合、Webサーバーは403の応答を返します。

ライブラリエントリの構造は次のとおりです：

```
struct {
	ThingUUID   uuid.UUID
	GUID        uuid.UUID
	UID         int32
	GIDS        []int32
	Global      boolean
	Name        string
	Description string
	Query       string
	Labels      []string
	Metadata    RawObject
}
```

構造体のメンバーは次のとおりです：

| メンバー| 説明| 空の場合は省略|
| ----------- | ------------------------------- | ---------------- |
| ThingUUID | *ローカルシステム上の一意の識別子* | |
| GUID | グローバル「名前」; キットのインストール後も持続します。 複数のユーザーが同じGUIDの検索ライブラリエントリを持っている場合があります。 | |
| UID | 所有者システムのユーザーID | |
| GID | エントリが共有されるグループIDのリスト| X |
| Global| エントリがグローバルに読み取り可能かどうかを示すブール値| |
| Name| クエリの人間が読める名前| |
| Description| 人間が読める形式のクエリの説明| |
| Query| クエリ文字列| |
| Labels | クエリの分類に使用される人間が読めるラベルのリスト| X |
| Metadata | 任意のパラメーターストレージに使用される不透明なJSONブロブ。有効なJSONであるだけで、特定の構造はありません。 X |

これは、UID 1のユーザーが所有し、3つのグループと共有されるエントリの例です。 UIDとGIDの値は、ユーザーAPIを使用してユーザー名とグループ名にマッピングし直す必要があります。

```
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"GIDs": [1, 3, 5],
	"Global": false,
	"Name": "syslog counting query",
	"Description": "A simple chart that shows total syslog over time",
	"Query": "tag=syslog syslog Appname | stats count by Appname | chart count by Appname",
	"Labels": [
		"syslog",
		"chart",
		"trending"
	],
	"Metadata": {}
}

```


## APIインタラクションの例

このセクションには、検索ライブラリAPIエンドポイントとの相互作用の例が含まれています。これらの例は、GravwellCLIを使用して`-debug`フラグを使用して生成されました。

### 新しいエントリの作成

リクエスト：
```
POST /api/library
{
	"GIDs": [1, 2],
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	]
}
```

注：ここでは「GUID」フィールドが設定されていないため、システムによって割り当てられます。リクエストにGUIDを含めることもできます。

リクエスト:
```
{
	"ThingUUID": "c9169d15-d643-11e9-99d3-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"GIDs": [1, 2],
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	]
}

```
### エントリの取得

リクエスト:

```
GET http://172.19.0.5:80/api/library
```

リクエスト:
```
[
	{
		"ThingUUID": "0b5a66cb-d642-11e9-931c-0242ac130005",
		"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
		"UID": 1,
		"Global": false,
		"Name": "netflow agg",
		"Description": "Total traffic using netflow data",
		"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
		"Labels": [
			"netflow",
			"traffic",
			"aggs"
		],
		"Metadata": {
			"value": 1,
			"extra": "some extra field value"
		}
	},
	{
		"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
		"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
		"UID": 1,
		"Global": false,
		"Name": "test2",
		"Description": "testing second",
		"Query": "tag=foo grep bar",
		"Labels": [
			"foo",
			"bar",
			"baz"
		],
	}
]
```

### 特定のエントリを要求します

リクエスト:

```
GET http://172.19.0.5:80/api/library/ae132ecc-88dd-11ea-a6aa-373f4c2439d4
```

リクエスト:
```
{
	"ThingUUID": "0b5a66cb-d642-11e9-931c-0242ac130005",
	"GUID": "ae132ecc-88dd-11ea-a6aa-373f4c2439d4",
	"UID": 1,
	"Global": false,
	"Name": "netflow agg",
	"Description": "Total traffic using netflow data",
	"Query": "tag=netflow netflow Bytes | stats sum(Bytes) as TotalTraffic | chart TotalTraffic",
	"Labels": [
		"netflow",
		"traffic",
		"aggs"
	],
	"Metadata": {
		"value": 1,
		"extra": "some extra field value"
	}
}
```

`api/library/0b5a66cb-d642-11e9-931c-0242ac130005`からも同じ応答が返されることに注意してください。

### エントリの更新

リクエスト:
```
PUT /api/library/69755a85-d5b1-11e9-89c2-0242ac130005
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
	"Global": false,
	"Name": "SyslogAgg",
	"Description": "Updated Syslog aggregate",
	"Query": "tag=syslog length | stats sum(length) | chart sum",
	"Labels": [
		"syslog",
		"agg",
		"totaldata"
	],
	"Metadata": {}
}
```

リクエスト:
```
{
	"ThingUUID": "69755a85-d5b1-11e9-89c2-0242ac130005",
	"GUID": "d57611be-88dd-11ea-a94d-df6bfb56a8a8",
	"UID": 1,
	"Global": false,
	"Name": "SyslogAgg",
	"Description": "Updated Syslog aggregate",
	"Query": "tag=syslog length | stats sum(length) | chart sum",
	"Labels": [
		"syslog",
		"agg",
		"totaldata"
	]
}
```

### エントリの削除

リクエスト:
```
DELETE /api/library/69755a85-d5b1-11e9-89c2-0242ac130005
```

### 管理者がエントリを削除します

非管理者が管理者フラグを追加した場合、Webサーバーはフラグを無視します。非管理者が指定されたエントリの所有者である場合、アクション（この場合はDELETE）は引き続き機能します。管理者以外のユーザーがエントリを所有していない場合、Webサーバーは403StatusForbiddenで応答します。

管理者による削除を実行するときは、常にURLパラメーターでThingUUIDを使用してください。そうしないと、間違ったアイテムが削除される可能性があります。

リクエスト:
```
DELETE /api/library/69755a85-d5b1-11e9-89c2-0242ac130005?admin=true
```
