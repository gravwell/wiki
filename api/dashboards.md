# ダッシュボードストレージAPI

ダッシュボードAPIは、基本的に、GUIがダッシュボードをレンダリングするために使用するjsonblobを管理するための汎用CRUDAPIです。 ダッシュボードに表示される検索はGUIによって起動され、backend/frontend/webserverには実際にはそれらが何であるかという概念がありません。

## データ構造

* ID：ダッシュボードを一意に識別する64ビット整数。
* GUID：ダッシュボードを指定するUUID。 キットのインストール全体で持続します（以下を参照）。 アクション可能からダッシュボードを参照するときに使用されます。
* Name：ダッシュボードのわかりやすい名前。
* Description：ダッシュボードのより詳細な説明。
* UID：ダッシュボードの所有者の数値ID。
* GID：このダッシュボードが共有される数値グループIDの配列。
* Global：ブール値。ダッシュボードをすべてのユーザーに表示する必要がある場合はtrueに設定します（管理者のみ）。
* Created：ダッシュボードが作成されたタイムスタンプ。
* Updated：ダッシュボードが最後に更新されたタイムスタンプ。
* Labels： [labels](#!gui/labels/labels.md).を含む文字列の配列。
* Data：ダッシュボードコンテンツの実際の定義（以下を参照）。

すべてのダッシュボードには、`ID`フィールドと`GUID`フィールドの両方があることに注意してください。 これは、ダッシュボードが、それらのダッシュボードを*参照*するアクション可能なものと一緒にキットにパックされている可能性があるためです。 キットにパックされたダッシュボードには既存のGUIDが含まれており、キットのインストール時にそのGUIDが保持されるため、アクション担当者がGUIDでダッシュボードを参照しても安全です。 一方、IDフィールドは、ダッシュボードが作成またはインストールされるたびにランダムに生成されます。 特定のシステムには、実際には同じGUIDを持つ複数のダッシュボード（通常は異なるユーザーによってインストールされます）がありますが、各ダッシュボードには独自の一意のIDがあります。

Webサーバーは `Data`フィールドに何が入力されるかを気にしませんが（有効なJSONである必要があることを除いて）、**GUI**が使用する特定の形式があります。 以下は、GUIで使用されるデータフィールドを含むダッシュボード構造の完全なTypescript定義です。

```
interface RawDashboard {
    ID: RawNumericID;
    GUID: RawUUID;

    UID: RawNumericID;
    GIDs: Array<RawNumericID> | null;

    Name: string;
    Description: string; // empty string is null
    Labels: Array<string> | null;

    Created: string; // Timestamp
    Updated: string; // Timestamp

    Data: {
        liveUpdateInterval?: number; // 0 is undefined
        linkZooming?: boolean;

        grid?: {
            gutter?: string | number | null; // string is a number
            margin?: string | number | null; // string is a number
        };

        searches: Array<{
            alias: string | null;
            timeframe?: {} | RawTimeframe;
            query?: string;
            searchID?: RawNumericID;
            reference?: {
                id: RawUUID;
                type: 'template' | 'savedQuery' | 'scheduledSearch';
                extras?: {
                    defaultValue: string | null;
                };
            };
        }>;
        tiles: Array<{
            id: RawNumericID;
            title: string;
            renderer: string;
            span: { col: number; row: number; x: number; y: number };
            searchesIndex: number;
            rendererOptions: RendererOptions;
        }>;
        timeframe: RawTimeframe;
        version?: 1 | 2;
        lastDataUpdate?: string; // Timestamp
    };
}

interface RawTimeframe {
    durationString: string | null;
    timeframe: string;
    start: string | null; // Timestamp
    end: string | null; // Timestamp
}

interface RendererOptions {
    XAxisSplitLine?: 'no';
    YAxisSplitLine?: 'no';
    IncludeOther?: 'yes';
    Stack?: 'grouped' | 'stacked';
    Smoothing?: 'normal' | 'smooth';
    Orientation?: 'v' | 'h';
    ConnectNulls?: 'no' | 'yes';
    Precision?: 'no';
    LogScale?: 'no';
    Range?: 'no';
    Rotate?: 'yes';
    Labels?: 'no';
    Background?: 'no';
    values?: {
        Smoothing?: 'smooth';
        Orientation?: 'h';
        columns?: Array<string>;
    };
}
```

注：このドキュメント全体に有効な`データ`構造が含まれていますが、簡潔にするために、タイルを含むダッシュボードではなく、空のダッシュボードを記述する構造を使用する傾向があります。

## ダッシュボードの作成

ダッシュボードを追加するには、次の形式のペイロードを使用して`/api/dashboards`にPOSTリクエストを発行します。

```
{
        "Name": "test2",
        "Description": "test2 description",
		"UID": 2,
		"GIDs": [],
		"Global": false,
        "Data": {
			"tiles": []
        }
}
```

`Data`プロパティは、実際のダッシュボードを作成するためにGUIによって使用されるJSONです。 初期化することも、後で入力するために空白のままにすることもできます。 ここでは、デモンストレーション用に空の「タイル」フィールドを含めました。

`UID`パラメータがリクエストから省略されている場合、デフォルトでリクエストしているユーザーのUIDになります。 管理者ユーザーのみが、自分以外のIDにUIDを設定できます。

`GIDs`配列には、ダッシュボードを共有するグループIDのリストが含まれている必要があります。 空のままにすると、ダッシュボードは誰とも共有されません。

管理者ユーザーは、ブール値の「グローバル」フィールドをtrueに設定することもできます。 これを設定すると、システム上のすべてのユーザーがダッシュボードにアクセスできるようになります。

`GUID`フィールドが含まれていて、有効なUUIDである場合、ランダムに生成されたものではなく、ダッシュボードに使用されます。 ほとんどの場合、Webサーバーがランダムなものを生成できるように、GUIDフィールドは空白のままにする必要があります。

Webサーバーの応答には、新しく作成されたダッシュボードの数値IDが含まれます。

## ダッシュボードの一覧表示

### 現在のユーザーのすべてのダッシュボードを取得する

ユーザーは、`/api/dashboards`でGETリクエストを発行することにより、（所有権またはグループを介して）アクセスできるすべてのダッシュボードを取得できます。 以下に、2つのダッシュボードを含む応答を示します。
```
[
  {
    "ID": 203486809563715,
    "Name": "test",
    "UID": 1,
    "GIDs": [],
    "Description": "test dashboard",
    "Created": "2020-09-22T09:16:51.66798721-06:00",
    "Updated": "2020-09-22T09:17:06.241311128-06:00",
    "Data": {
      "searches": [
        {
          "alias": "Search 1",
          "timeframe": {},
          "query": "tag=* count\n",
          "searchID": 4780372388
        }
      ],
      "tiles": [
        {
          "title": "count",
          "renderer": "text",
          "span": {
            "col": 4,
            "row": 4,
            "x": 0,
            "y": 0
          },
          "searchesIndex": 0,
          "id": 16007878262310,
          "rendererOptions": {}
        }
      ],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      },
      "version": 2,
      "lastDataUpdate": "2020-09-22T09:17:06-06:00"
    },
    "Labels": null,
    "GUID": "9719d92a-df1f-4a05-885a-ad10915d8b42",
    "Synced": false
  },
  {
    "ID": 69148521436807,
    "Name": "Test 2",
    "UID": 1,
    "GIDs": [],
    "Description": "dashboard 2",
    "Created": "2020-09-22T09:17:13.809070187-06:00",
    "Updated": "2020-09-22T09:17:13.809070187-06:00",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    },
    "Labels": null,
    "GUID": "2c55cf84-bb63-40cf-bf54-3bff8c8d7fb6",
    "Synced": false
  }
]
```

### ユーザーが所有するすべてのダッシュボードを取得する
ユーザーが明示的に所有するダッシュボードを取得するには、`/api/users/{uid}/dashboards`でGETリクエストを発行し、{uid}を目的のユーザーIDに置き換えます。 Webサーバーは、そのUIDによって特別に所有されているダッシュボードのみを返します。 これには、ユーザーがグループメンバーシップを通じてアクセスできるダッシュボードは含まれません。

```
WEB GET /api/users/1/dashboards:
[
  {
    "ID": 4,
    "Name": "dashGroup2",
    "UID": 5,
    "GIDs": [
      3
    ],
    "Description": "dashGroup2",
    "Created": "2016-12-28T21:37:12.703358455Z",
    "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```

### グループのすべてのダッシュボードを取得する
特定のグループがアクセスできるすべてのダッシュボードを取得するには、GETリクエストを`/api/groups/{gid}/dashboards`に発行し、{gid}を目的のグループIDに置き換えます。 サーバーは、そのグループと共有されているダッシュボードを返します。 これは、ダッシュボードを複数のグループと共有できるという点で、通常のUnix権限とは少し異なります（したがって、グループは実際にはダッシュボードを所有していませんが、ダッシュボードは一種のグループのメンバーです）。

```
WEB GET /api/groups/2/dashboards:
[
  {
    "ID": 3,
    "Name": "dashGroup1",
    "UID": 5,
    "GIDs": [
      2
    ],
    "Description": "dashGroup1",
    "Created": "2016-12-28T21:37:12.696460531Z",
    "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  },
  {
    "ID": 2,
    "Name": "test2",
    "UID": 3,
    "GIDs": [
      2
    ],
    "Description": "test2 description",
    "Created": "2016-12-18T23:28:08.250051418Z",
    "GUID": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```


### 管理者のみ：すべてのユーザーのすべてのダッシュボードを一覧表示します
システム上の*すべて*のダッシュボードを取得するために、管理者ユーザーは`/api/Dashboards/all`にGETリクエストを発行できます。 このリクエストが管理者以外のユーザーによって発行された場合、アクセスできるすべてのダッシュボードを返す必要があります(`/api/dashboards`のGETに相当)

```
WEB GET /api/dashboards/all:
[
  {
    "ID": 1,
    "Name": "test1",
    "UID": 1,
    "GIDs": [],
    "Description": "test1 description",
    "Created": "2016-12-18T23:28:07.679322121Z",
    "GUID": "d28b6887-ad55-479e-8af3-0cbcbd5084b1",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  },
  {
    "ID": 2,
    "Name": "test2",
    "UID": 3,
    "GIDs": [],
    "Description": "test2 description",
    "Created": "2016-12-18T23:28:08.250051418Z",
    "GUID": "55bc7236-39e4-11e9-94e9-54e1ad7c66cf",
    "Data": {
      "searches": [],
      "tiles": [],
      "timeframe": {
        "durationString": "PT1H",
        "timeframe": "PT1H",
        "end": null,
        "start": null
      }
    }
  }
]

```

## 特定のダッシュボードを取得する
特定のIDを取得するには、`/api/dashboards/{id}`に対してGETリクエストを発行し、`{id}`をダッシュボードのIDに置き換えてください（例：`/api/dashboards/69148521436807`）。

IDではなくGUIDで特定のダッシュボードをフェッチすることもできます：`GET/api/dashboards/d28b6887-ad55-479e-8af3-0cbcbd5084b1`。


## ダッシュボードの更新
ダッシュボードの更新（データの変更や名前、説明、GIDリストの変更など）は、`/api/dashboards/{id}`にPUTリクエストを発行することで行われます。ここで、{id}は特定のダッシュボードの一意の識別子です。 。

この例では、ユーザー（UID 3）は、グループ1がダッシュボード（ID 2）にアクセスするためのアクセス許可を追加したいと考えています。 ユーザーは最初にダッシュボードを取得します。
```
GET /api/dashboards/2:
{
  "ID": 2,
  "Name": "test2",
  "UID": 3,
  "GIDs": [],
  "Description": "test2 description",
  "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
  "Created": "2016-12-18T23:28:08.250051418Z",
  "Data": {
    "searches": [
      {
        "alias": "Search 1",
        "timeframe": {},
        "query": "tag=* count\n",
        "searchID": 4780372388
      }
    ],
    "tiles": [
      {
        "title": "count",
        "renderer": "text",
        "span": {
          "col": 4,
          "row": 4,
          "x": 0,
          "y": 0
        },
        "searchesIndex": 0,
        "id": 16007878262310,
        "rendererOptions": {}
      }
    ],
    "timeframe": {
      "durationString": "PT1H",
      "timeframe": "PT1H",
      "end": null,
      "start": null
    },
    "version": 2,
    "lastDataUpdate": "2020-09-22T09:17:06-06:00"
  }
}
```

これで、ユーザーはターゲットダッシュボードを取得し、変更を加えて、PUTリクエストを投稿しました:
```
WEB PUT /api/dashboards/2:
{
  "ID": 2,
  "Name": "marketoverview",
  "UID": 3,
  "GIDs": [
    3
  ],
  "Description": "marketing group dashboard",
  "GUID": "5c6099dc-39e4-11e9-81a7-54e1ad7c66cf",
  "Created": "2016-12-18T23:28:08.250051418Z",
  "Data": {
    "searches": [
      {
        "alias": "Search 1",
        "timeframe": {},
        "query": "tag=* count\n",
        "searchID": 4780372388
      }
    ],
    "tiles": [
      {
        "title": "count",
        "renderer": "text",
        "span": {
          "col": 4,
          "row": 4,
          "x": 0,
          "y": 0
        },
        "searchesIndex": 0,
        "id": 16007878262310,
        "rendererOptions": {}
      }
    ],
    "timeframe": {
      "durationString": "PT1H",
      "timeframe": "PT1H",
      "end": null,
      "start": null
    },
    "version": 2,
    "lastDataUpdate": "2020-09-22T09:17:06-06:00"
  }
}

```

サーバーは、更新されたダッシュボード構造で更新要求に応答します。

安全のため、変更されていない場合でも、元のフェッチに存在していたすべてのフィールドを返送するように注意してください。 たとえば、この更新では `Description`フィールドは変更されませんでしたが、ウェブサーバーは*un-set*フィールドと*empty*フィールドを区別できないため、更新リクエストに含めます。

注：必要に応じて、ダッシュボードIDの代わりにGUIDを使用できます。

## ダッシュボードの削除
ダッシュボードを削除するには、`/api/dashboards/{id}`というURLにDELETEメソッドでリクエストを発行します。{id}はダッシュボードの数値IDです。
