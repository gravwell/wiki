# その他のAPI

いくつかのAPIは、主要なカテゴリにうまく収まらないものがあります。ここではそれらを紹介します。

## 接続性テスト

このAPIは、バックエンドがHTTPリクエストに応答しているかどうかを検証するためのものです。`api/test` への GET は、200 ステータスを返し、ボディコンテンツはありません。認証は必要ありません。

## バージョン認証

`api/version` を GET して、バージョン情報を取得します。 認証は必要ありません。

```
{
    "API": {
        "Major": 0,
        "Minor": 1
    },
    "Build": {
        "BuildDate": "2020-05-04T00:00:00Z",
        "BuildID": "6c48dd4c",
        "GUIBuildID": "b5c8cd58",
        "Major": 4,
        "Minor": 0,
        "Point": 0
    }
}
```

## タグリスト

ウェブサーバは、インデクサが知っているすべてのタグのリストを管理しています。このリストは `/api/tags` への GET リクエストで取得することができます。これはタグのリストを返します。

```
["default", "gravwell", "pcap", "windows"]
```

## 検索モジュール一覧

利用可能なすべての検索モジュールのリストと、それぞれのモジュールに関する情報を取得するには、`/api/info/searchmodules`をGETしてください。これは、モジュール情報構造のリストを返します。

```
[
    {
        "Collapsing": true,
        "Examples": [
            "min by src",
            "min by someKey"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "min",
        "Sorting": true
    },
    {
        "Collapsing": true,
        "Examples": [
            "unique",
            "unique chuck",
            "unique chuck,testa"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "unique",
        "Sorting": false
    },
[...]
    {
        "Collapsing": false,
        "Examples": [
            "alias src dst"
        ],
        "FrontendOnly": false,
        "Info": "Alias enumerated values",
        "Name": "alias",
        "Sorting": false
    },
    {
        "Collapsing": true,
        "Examples": [
            "count",
            "count by chuck",
            "count by src",
            "count by someKey"
        ],
        "FrontendOnly": false,
        "Info": "No information available",
        "Name": "count",
        "Sorting": true
    }
]
```

## レンダーモジュールリスト

利用可能なすべてのレンダリングモジュールのリストと、各モジュールに関する情報を取得するには、`/api/info/rendermodules`をGETします。これは、モジュール情報構造のリストを返します。

```
[
    {
        "Description": "A raw entry storage system, it can store and handle anything.",
        "Examples": [
            "raw"
        ],
        "Name": "raw",
        "SortRequired": false
    },
    {
        "Description": "A chart storage system system.\n\t       Chart looks for numeric types, storing them.\n\t       Requested entries will be a set of types with column names.",
        "Examples": [
            "chart"
        ],
        "Name": "chart",
        "SortRequired": false
    },
[...]
    {
        "Description": "A point mapping system that supports condensing and geofencing",
        "Examples": [],
        "Name": "point2point",
        "SortRequired": false
    }
]
```

## GUIの設定

このAPIは、ユーザーインターフェイスの基本的な情報を提供します。`/api/settings` を GET すると、以下のような構造が返されます。

```
{
  "DisableMapTileProxy": false,
  "DistributedWebservers": false,
  "MapTileUrl": "http://localhost:8080/api/maps",
  "MaxFileSize": 8388608,
  "MaxResourceSize": 134217728,
  "ServerTime": "2020-11-30T11:50:29.478092519-08:00",
  "ServerTimezone": "PST",
  "ServerTimezoneOffset": -28800
}

```

* `DisableMapTileProxy`, trueの場合、UIにマップリクエストをGravwellプロキシを使わずにOpenStreetMapサーバーに直接送信するように指示します。
* `MapTileUrl` は UI がマップタイルを取得する際に使用する URL です。
* `DistributedWebservers` は、データストアを介して連携している複数のウェブサーバがある場合、true に設定されます。
* `MaxFileSize` は `/api/files` API にアップロード可能な最大許容ファイルサイズ (バイト単位) です。
* `MaxResourceSize` は、許容可能なリソースの最大サイズ（バイト）です。
* `ServerTime` はウェブサーバの現在の時刻です。
* `ServerTimezone` はウェブサーバのタイムゾーンである。
* `ServerTimezoneOffset` はウェブサーバーのタイムゾーンのオフセットで、UTCからの秒数です。

## スクリプティングライブラリ

このAPIでは、自動化スクリプトが `require` 関数を使って github リポジトリからライブラリをインポートすることができます。また、ユーザーのすべてのリポジトリに対して git pull を実行するエンドポイントもあります。

### ライブラリの取得

このエンドポイントは、サーチエージェントがライブラリ機能を利用する場合にのみ有用だと思われますが、念のため記載します。指定したリポジトリからファイルを取得するには、URLにパラメータを指定してGETを行います。

```
/api/libs?repo=github.com/gravwell/libs&commit=40e98d216bb6e69642df392b255e8edc0f57eb06&path=utils/links.ank
```

"repo"と"commit"の値は省略可能です。"repo"が省略された場合、デフォルトではgithub.com/gravwell/libsとなります。"commit"を省略した場合、デフォルトでは master ブランチの先端になります。

### ライブラリの更新

リポジトリはユーザーごとに管理されています。ユーザーは `/api/libs/pull` に GET リクエストを送信することで、自分のリポジトリセットに `git pull` を強制的に適用することができます。これには時間がかかりますのでご注意ください。
