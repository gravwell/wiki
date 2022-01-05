# Actionables API

Actionables (以前は "pivots" と呼ばれていました) は、Gravwell に保存されているオブジェクトで、Web GUI が検索結果データからピボットするために使用します。例えば、actionableはIPアドレスに実行できるクエリのセットを定義し、IPアドレスに*マッチする*正規表現と一緒にすることができます。ユーザーがIPアドレスを含むクエリを実行すると、それらのアドレスがクリック可能になり、定義済みのクエリを起動するためのメニューが表示されます。

## データ構造

アクショナブル構造体は、以下のフィールドを含みます：

* GUID: アクショナブルのグローバルリファレンスです。キットのインストールをまたいで永続化します。(次のセクションを参照)
* ThingUUID: この特定のアクション可能なインスタンスのための一意のID。(次のセクションを参照)
* UID: アクション可能な所有者の数値ID。
* GIDs: このアクショナブルが共有されるグループIDの数値の配列
* Global: ブール値で、操作対象がすべてのユーザーに見えるようにする場合はtrueに設定されます（管理者のみ）。
* Name: アクショナブルの名前です。
* Description: アクショナブルをより詳しく説明したものです。
* Updated: アクション可能なものの最終更新時刻を表すタイムスタンプ
* Labels:[labels](#!gui/labels/labels.md)を含む文字列の配列です。
* Disabled: 操作可能が無効になっているかどうかを示すブーリアン値
* Contents:アクショナブルの実際の定義自体（下記参照）
  * Contents.menuLabel: オプションです。存在しない場合、名前の最初の20文字がドロップダウンメニューに使用されます。
  * Contents.actions: このアクショナブルから実行可能なアクションの配列です。
    * Contents.actions\[n].name: アクション名
    * Contents.actions\[n].description: オプションでアクションの説明を記述します。
    * Contents.actions\[n].placeholder: プレースホルダーは、トリガーまたはカーソルのハイライトの値に置き換えられます。 デフォルトは「\ _VALUE_」です。
    * Contents.actions\[n].start: 開始日変数を処理するための定義を持つオプションのActionableTimeVariable（以下のインターフェースを参照）。
    * Contents.actions\[n].end: オプションの ActionableTimeVariable (以下のインターフェイスを参照) は、終了日変数を処理するための定義があります。
    * Contents.actions\[n].command: アクション実行のための定義を持つ ActionableCommand（以下のインターフェイスを参照）。
  * Contents.triggers: このアクション可能なトリガーの配列です。
    * Contents.triggers\[n].pattern: アクションがマッチするためのシリアル化された正規表現パターン
    * Contents.triggers\[n].hyperlink: クリックとテキスト選択で操作可能な場合はTrue。テキスト選択のみで有効な場合はFalse。

ウェブサーバは `Contents` フィールドに何が入るかを気にしませんが (有効な JSON であることを除いて)、 **GUI** が使用する特定のフォーマットが存在します。以下は、Contents フィールドと、使用される様々な型の説明を含む、アクション可能な構造体の完全な Typescript の定義です。

```
interface Actionable {
    GUID: UUID;
    ThingUUID: UUID;
    UID: NumericID;
    GIDs: null | Array<NumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Could be an empty string
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
    | { type: 'template'; reference: UUID; options?: { variable?: string } }
    | { type: 'savedQuery'; reference: UUID; options?: {} }
    | { type: 'dashboard'; reference: UUID; options?: { variable?: string } }
    | { type: 'url'; reference: string; options: { modal?: boolean; modalWidth?: string } };
```

## ネーミング GUIDとThingUUID

アクショナブルには、GUIDとThingUUIDという2種類のIDが付与されています。どちらもUUIDですが、これは混乱を招く可能性があります。このセクションでは、その理由を説明します。

例を考えてみましょう。アクションをゼロから作成し、ランダムなGUID、 `e80293f0-5732-4c7e-a3d1-2fb779b91bf7` とランダムなThingUUID、 `c3b24e1e-5186-4828-82ee-82724a1d4c45` が割り当てられました。そして、実行可能なものをキットにバンドルします。同じシステムの別のユーザーがこのキットを自分用にインストールすると、**same**  GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`) で **random** ThingUUID (`f07373a8-ea85-415f-8dfd-61f7b9204ae0`) の actionable がインスタンス化されるのですが、このキットには **same** という名前のついた GUID はありません。

このシステムは [templates](templates.md) で使われているものと同じです。テンプレートはGUIDとThingUUIDを使用しているので、ダッシュボードはGUIDでテンプレートを参照できますが、複数のユーザーが同時に同じキット（サンプルテンプレート付き）をインストールしても衝突することはありません。ダッシュボードがテンプレートを参照するのと同じように、Gravwellコンポーネントがactionableを参照することはありませんが、将来の対策としてこの挙動を含めました。

### GUID と ThingUUID を使用した Actionables へのアクセス

一般ユーザーは常にGUIDでactionableにアクセスする必要があります。管理者ユーザーは、代わりにThingUUIDでactionableを参照することができますが、`?admin=true`パラメータをリクエストURLに設定する必要があります。

## アクショナブルを作成する

アクションを作成するには、 `/api/pivots` に POST してください。ボディはJSON構造で、'Contents'フィールドと、オプションでGUID、Labels、Name、Descriptionが必要です。例えば：

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

APIは新しく作成されたactionableのGUIDで応答します。リクエストでGUIDが指定された場合、そのGUIDが使用されます。GUIDが指定されない場合、ランダムなGUIDが生成されます。

注意: 現時点では、`UID`、`GIDs`、`Global` フィールドは、アクション可能の作成中に設定することはできません。代わりに、updateコールで設定する必要があります(下記参照)。

## 実行可能なリスト

ユーザーが利用できるすべてのアクションをリストアップするには、 `/api/pivots` を GET してください。その結果、actionablesの配列が生成されます：

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

## 実行可能な単一のフェッチ

一つのアクションを取得するには、 `/api/pivots/<guid>` にGETリクエストを発行してください。例えば、 `/api/pivots/afba4f9b-f66a-4f9f-9c58-f45b3db6e474` にGETすると、サーバーはそのアクションの内容で応答するでしょう。

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

管理者は、ThingUUIDとadminパラメータを使用して、明示的にこの特定のactionableを取得できることに注意してください、例`/api/pivots/196a3cc3-ec9e-11ea-bfde-7085c2d881ce?admin=true`。

## アクションの更新

アクションを更新するには、 `/api/pivots/<guid>` にPUTリクエストを発行してください。リクエストのボディは、同じパスのGETで返されるものと同じでなければなりませんが、必要な要素は変更されます。GUIDとThingUUIDは変更できないことに注意してください。変更できるのは以下のフィールドのみです：

* コンテンツ: アクショナブルの実際のボディ/コンテンツ
* Name: アクションの名前を変更します。
* Description: アクションの説明を変更します。
* GIDs:  32ビット整数のグループIDの配列を設定することができます。
* UID: （管理者のみ）32ビット整数に設定
* Global:  (管理者のみ) true または false のブール値に設定します; Global actionables はすべてのユーザーから見えます。

注：これらのフィールドのいずれかを空白にすると、そのフィールドのNULL値で実行可能ファイルが更新されます。

## アクションの削除

アクションを削除するには、`/api/pivots/<guid>` に DELETE リクエストを発行してください。

## 管理者の行動

管理者ユーザーは、システム上のすべてのactionableを表示、変更、または削除する必要がある場合があります。GUIDは必ずしも一意ではないので、管理者APIは代わりにGravwellがアイテムを保存するために内部で使用する一意のUUIDを参照する必要があります。上記のアクション可能なリストの例では、"ThingUUID "というフィールドがあることに注意してください。これは、そのアクション可能な内部でユニークな識別子です。

管理者ユーザーは、`/api/pivots?admin=true`のGETリクエストで、システム内のすべてのactionableのグローバルリストを取得することができます。

次に管理者は、`/api/pivots/<ThingUUID>?admin=true`へのPUTで特定のactionableを更新し、希望のactionableのThingUUID値に代入することができます。同じパターンが削除にも適用されます。

管理者は `/api/pivots/<ThingUUID>?admin=true` に対して GET と DELETE リクエスト（それぞれ）で特定のアクショナブルにアクセスしたり削除したりすることができます。