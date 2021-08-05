# ピボット/アクション可能API

*actionables*とも呼ばれるピボットは、Web GUIが検索結果データから「ピボット」するために使用するGravwellに格納されているオブジェクトです。たとえば、actionableは、IPアドレスで実行できるクエリのセットを、IPアドレスに*一致する*正規表現とともに定義できます。ユーザーが結果にIPアドレスを含むクエリを実行すると、それらのアドレスをクリックできるようになり、事前定義されたクエリを起動するためのメニューが表示されます。

## データ構造

ピボット構造には、以下の項目が含まれています。

* GUID：テンプレートのグローバルリファレンス。キットのインストール後も持続します。 （次のセクションを参照）
* ThingUUID：この特定のテンプレートインスタンスの一意のID。 （次のセクションを参照）
* UID：テンプレートの所有者の数値ID。
* GID：このテンプレートが共有される数値グループIDの配列。
* グローバル：ブール値。テンプレートをすべてのユーザーに表示する必要がある場合はtrueに設定します（管理者のみ）。
* Name：テンプレートの名前。
* Description：テンプレートのより詳細な説明。
* Updated：テンプレートの最終更新時刻を表すタイムスタンプ。
* Labels：[labels](#!gui/labels/labels.md)を含む文字列の配列。
* Disabled：ピボットが無効になっているかどうかを示すブール値。
* Contents：テンプレート自体の実際の定義（以下を参照）。

Webサーバーは`Contents`フィールドに何が入力されるかを気にしませんが（有効なJSONである必要があることを除いて）、**GUI**が使用する特定の形式があります。以下は、実行可能な構造の完全なTypescript定義であり、Contentsフィールドと、使用されるさまざまなタイプの説明が含まれています。

```
interface Actionable {
    GUID: UUID;
    ThingUUID: UUID;
    UID: NumericID;
    GIDs: null | Array<NumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Empty string is null
    Updated: string; // Timestamp
    Contents: {
        menuLabel: null | string;
        actions: Array<ActionableAction>;
        triggers: Array<ActionableTrigger>;
    };
    Labels: null | Array<string>;
    Disabled: boolean;
}

type UUID = string;
type NumericID = number;

interface ActionableTrigger {
    pattern: string;
    hyperlink: boolean;
}

interface ActionableAction {
    name: string;
    description: string | null;
    placeholder: string | null;
    start?: ActionableTimeVariable;
    end?: ActionableTimeVariable;
    command: ActionableCommand;
}

type ActionableTimeVariable =
    | { type: 'timestamp'; format: null | string; placeholder: null | string }
    | { type: 'string'; format: null | string; placeholder: null | string };

type ActionableCommand =
    | { type: 'query'; reference: string; options?: {} }
    | { type: 'template'; reference: UUID; options?: {} }
    | { type: 'savedQuery'; reference: UUID; options?: {} }
    | { type: 'dashboard'; reference: UUID; options?: { variable?: string } }
    | { type: 'url'; reference: string; options: { modal?: boolean; modalWidth?: string } };
```

## ネーミング：GUIDとThingUUID

ピボット/アクション可能オブジェクトには、GUIDとThingUUIDの2つの異なるIDが関連付けられています。これらは両方ともUUIDであり、混乱を招く可能性があります。1つのオブジェクトに2つの識別子があるのはなぜですか。このセクションで明確にしようとします。

例を考えてみましょう。ピボットを最初から作成するので、ランダムなGUID `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`とランダムなThingUUID`c3b24e1e-5186-4828-82ee-82724a1d4c45`が割り当てられます。次に、ピボットをキットにバンドルします。次に、同じシステム上の別のユーザーがこのキットを自分でインストールします。これにより、**同じ** GUID（ `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`）でピボットがインスタンス化されますが、**ランダム** ThingUUID（` f07373a8- ea85-415f-8dfd-61f7b9204ae0`）。

このシステムは、[templates](templates.md)で使用されているものと同じです。テンプレートはGUIDとThingUUIDを使用するため、ダッシュボードはGUIDでテンプレートを参照できますが、複数のユーザーが同じキット（サンプルテンプレートを使用）を競合することなく同時にインストールできます。ダッシュボードがテンプレートを参照するのと同じ方法でアクション可能オブジェクトを参照するGravwellコンポーネントはありませんが、将来を見据えた動作として含まれています。

### GUIDとThingUUIDを介したピボットへのアクセス

通常のユーザーは、常にGUIDによってピボットにアクセスする必要があります。管理ユーザーは代わりにThingUUIDでピボットを参照できますが、リクエストURLに`？admin=true`パラメータを設定する必要があります。

## ピボットを作成する

ピボットを作成するには、`/api/pivots`にPOSTを発行します。本文は、「コンテンツ」フィールドと、オプションでGUID、ラベル、名前、説明を含むJSON構造である必要があります。例えば：

```
{
  "Name": "IP actions",
  "Description": "Actions for an IP address",
  "Contents": {
    "actions": [
      {
        "name": "Whois",
        "description": null,
        "placeholder": null,
        "start": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "end": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "command": {
          "type": "url",
          "reference": "https://www.whois.com/whois/_VALUE_",
          "options": {}
        }
      }
    ],
    "menuLabel": null,
    "triggers": [
      {
        "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
        "hyperlink": true
      }
    ]
  }
}
```

APIは、新しく作成されたピボットのGUIDで応答します。リクエストでGUIDが指定されている場合、そのGUIDが使用されます。GUIDが指定されていない場合、ランダムなGUIDが生成されます。

注：現時点では、ピボットの作成中に`UID`、`GID`、および「グローバル」フィールドを設定することはできません。 代わりに、更新呼び出しを介して設定する必要があります（以下を参照）。

## リストピボット

ユーザーが利用できるすべてのピボットを一覧表示するには、`/api/pivots`でGETを実行します。結果は、ピボットの配列になります。

```
[
  {
    "GUID": "afba4f9b-f66a-4f9f-9c58-f45b3db6e474",
    "ThingUUID": "196a3cc3-ec9e-11ea-bfde-7085c2d881ce",
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "Name": "IP actions",
    "Description": "Actions for an IP address",
    "Updated": "2020-09-01T15:57:23.416537696-06:00",
    "Contents": {
      "actions": [
        {
          "name": "Whois",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "url",
            "reference": "https://www.whois.com/whois/_VALUE_",
            "options": {}
          }
        }
      ],
      "menuLabel": null,
      "triggers": [
        {
          "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
          "hyperlink": true
        }
      ]
    },
    "Labels": null,
    "Disabled": false
  },
  {
    "GUID": "34ba8372-0314-460a-9742-5a65c18d6241",
    "ThingUUID": "e1bdf35a-de7b-11ea-9709-7085c2d881ce",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Network Port",
    "Description": "Actions to take on a network port, e.g. 22",
    "Updated": "2020-08-14T16:17:03.790048874-06:00",
    "Contents": {
      "actions": [
        {
          "name": "Netflow - Most active hosts on this port",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Src Dst SrcPort DstPort Port==_VALUE_ Protocol Bytes | stats sum(Bytes) as ByteTotal by Port Src Dst | lookup -r network_services Protocol proto_number proto_name as Proto Port service_port service_name as Service | table Src Dst Port Service Proto ByteTotal",
            "options": {}
          }
        },
        {
          "name": "Netflow - Chart traffic",
          "description": "Traffic on this port over time",
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Src Dst SrcPort DstPort Port==_VALUE_ Protocol Bytes | lookup -r network_services Protocol proto_number proto_name as Proto Port service_port service_name as Service | stats sum(Bytes) by Service Port | chart sum by Service Port",
            "options": {}
          }
        },
        {
          "name": "Netflow - Internal IPs serving this port",
          "description": null,
          "placeholder": null,
          "start": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "end": {
            "type": "string",
            "format": null,
            "placeholder": null
          },
          "command": {
            "type": "query",
            "reference": "tag=netflow netflow Dst ~ PRIVATE DstPort==_VALUE_ Bytes Protocol | lookup -r ip_protocols Protocol Number Name as ProtocolName | stats sum(Bytes) as TotalTraffic by Dst | table Dst DstPort Protocol ProtocolName TotalTraffic",
            "options": {}
          }
        }
      ],
      "menuLabel": null,
      "triggers": []
    },
    "Labels": [
      "kit/io.gravwell.netflowv5"
    ],
    "Disabled": false
  }
]
```

## 単一のピボットをフェッチします

単一のピボットをフェッチするには、`/api/pivots/<guid>`にGETリクエストを発行します。 サーバーはそのピボットの内容で応答します。たとえば、`/api/ivots/afba4f9b-f66a-4f9f-9c58-f45b3db6e474`のGETは次を返します。

```
{
  "GUID": "afba4f9b-f66a-4f9f-9c58-f45b3db6e474",
  "ThingUUID": "196a3cc3-ec9e-11ea-bfde-7085c2d881ce",
  "UID": 1,
  "GIDs": null,
  "Global": false,
  "Name": "IP actions",
  "Description": "Actions for an IP address",
  "Updated": "2020-09-01T15:57:23.416537696-06:00",
  "Contents": {
    "actions": [
      {
        "name": "Whois",
        "description": null,
        "placeholder": null,
        "start": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "end": {
          "type": "string",
          "format": null,
          "placeholder": null
        },
        "command": {
          "type": "url",
          "reference": "https://www.whois.com/whois/_VALUE_",
          "options": {}
        }
      }
    ],
    "menuLabel": null,
    "triggers": [
      {
        "pattern": "/\\b(?:[0-9]{1,3}\\.){3}[0-9]{1,3}\\b/g",
        "hyperlink": true
      }
    ]
  },
  "Labels": null,
  "Disabled": false
}

```

管理者は、ThingUUIDとadminパラメータを使用して、この特定のピボットを明示的にフェッチできることに注意してください。`/api/pivots/196a3cc3-ec9e-11ea-bfde-7085c2d881ce?admin=true`。

## ピボットを更新する

ピボットを更新するには、`/api/pivots/<guid>`にPUTリクエストを発行します。 リクエストの本文は、同じパスでGETによって返されるものと同じである必要がありますが、必要な要素は変更されています。 GUIDとThingUUIDは変更できないことに注意してください。 次のフィールドのみを変更できます。

* Contents：ピボットの実際の本体/内容
* Name：ピボットの名前を変更します
* Description：ピボットの説明を変更します
* GID：32ビット整数グループIDの配列に設定できます。 `"GIDs":[1,4]`
* UID :(管理者のみ）32ビット整数に設定
* Global:(管理者のみ）ブール値のtrueまたはfalseに設定します。 グローバルピボットはすべてのユーザーに表示されます。

注：これらのフィールドのいずれかを空白のままにすると、ピボットがそのフィールドのnull値で更新されます。

##ピボットを削除する

ピボットを削除するには、`/api/pivots/<guid>`にDELETEリクエストを発行します。

## 管理者のアクション

管理ユーザーは、システム上のすべてのピボットを表示、変更、または削除する必要がある場合があります。 GUIDは必ずしも一意ではないため、管理APIは、代わりにGravwellがアイテムを格納するために内部で使用する一意のUUIDを参照する必要があります。 上記のピボットリストの例には、「ThingUUID」という名前のフィールドが含まれていることに注意してください。 これは、そのピボットの内部の一意の識別子です。

管理者ユーザーは、`/api/pivots?admin=true`のGETリクエストを使用して、システム内のすべてのピボットのグローバルリストを取得できます。

次に、管理者は、PUTを使用して特定のピボットを`/api/pivots/<ThingUUID>?admin=true`に更新し、目的のピボットをThingUUID値に置き換えることができます。 同じパターンが削除にも当てはまります。

管理者は、`/api/pivots/<ThingUUID>?admin=true`でGETまたはDELETEリクエストを使用して、特定のピボットにアクセスまたは削除できます。