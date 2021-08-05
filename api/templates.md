# テンプレート API

テンプレートとは、Gravwellのクエリを定義する特殊なオブジェクトです。同じ変数を使った複数のテンプレートをダッシュボードに組み込むことで、強力な調査ツールを作ることができます。例えば、IPアドレスを変数とするテンプレートを使って、IPアドレス調査ダッシュボードを作ることができます。

## データ構造

テンプレート構造には、以下のフィールドがあります。

* GUID： テンプレートのグローバルリファレンス。キットをインストールしても永続します。(次のセクションを参照)
* ThingUUID：この特定のテンプレート・インスタンスのためのユニークなID。(次のセクションを参照)
* UID：テンプレートの所有者の数字によるIDです。
* GIDs： このテンプレートが共有されているグループIDの配列
* Global： ブーリアン値で、テンプレートをすべてのユーザー（adminのみ）に表示する場合はtrueに設定されます。
* Name： テンプレートの名前です。
* Description： テンプレートの詳細な説明です。
* Updated： テンプレートの最終更新時刻を表すタイムスタンプ
* Labels： [ラベル](#!gui/labels/labels.md)を含む文字列の配列です。
* コンテンツ。コンテンツ： テンプレートの実際の定義 (下記参照)


ウェブサーバは、`Contents`フィールドに何が入るかは気にしませんが(有効なJSONでなければならないことを除いて)、**GUI**が使用する特定のフォーマットがあります。GUIで使用するためには、Contentsフィールドがこの構造に準拠している必要があります。

```
Contents: {
  query: string;
  variable: string;
  variableLabel: string;
  variableDescription: string | null;
  required: boolean;
  testValue: string | null;
};
```

以下は、テンプレートデータタイプの完全なTypescriptの定義です：

```
interface RawTemplate {
    GUID: RawUUID;
    ThingUUID: RawUUID;
    UID: RawNumericID;
    GIDs: null | Array<RawNumericID>;
    Global: boolean;
    Name: string;
    Description: string; // Empty string is null
    Updated: string; // Timestamp
    Contents: {
        query: string;
        variable: string;
        variableLabel: string;
        variableDescription: string | null;
        required: boolean;
        testValue: string | null;
    };
    Labels: null | Array<string>;
}
```

## ネーミング GUIDとThingUUID

テンプレートには、GUIDとThingUUIDという2つの異なるIDが付けられています。これらはどちらもUUIDですが、これは混乱を招く恐れがあります。このセクションではそれを明らかにしていきます。

Gravwellでは、ダッシュボードは特定のテンプレートを指すことがあります。このダッシュボードと対応するテンプレートは、他のユーザーに配布するためにキットにまとめられることもあります。ダッシュボードは、キットにまとめられて他の場所にインストールされたときにも、テンプレートを参照する方法が必要です。そこで、テンプレートの「グローバル」な名前としてGUIDを導入しています。しかし、複数のユーザーが同じキットをインストールすることが許されているので、テンプレートの個々のインスタンス化*のために異なる識別子も必要です。この役割を果たすのが、ThingUUIDフィールドです。

例を挙げてみましょう。ダッシュボードとテンプレートを含むキットを作ります。テンプレートをゼロから作成するので、ランダムなGUID、`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`と、ランダムなThingUUID、`c3b24e1e-5186-4828-82ee-82724a1d4c45`が割り当てられます。そして、ダッシュボードに、テンプレートをGUID（`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`）で参照するタイルを作成し、テンプレートとダッシュボードの両方をキットにバンドルします。同じシステムの別のユーザーが自分用にこのキットをインストールすると、**同じ** GUID (`e80293f0-5732-4c7e-a3d1-2fb779b91bf7`)で**ランダム**なThingUUID (`f07373a8-ea85-415f-8dfd-61f7b9204ae0`)を持つテンプレートがインスタンス化されます。ユーザーがダッシュボードを開くと、ダッシュボードはGUID == `e80293f0-5732-4c7e-a3d1-2fb779b91bf7`のテンプレートを要求します。ウェブサーバは、ThingUUID == `f07373a8-ea85-415f-8dfd-61f7b9204ae0`のテンプレートのそのユーザのインスタンスを返します。

ユーザーがグローバルテンプレートと同じGUIDのテンプレートをインストールした場合、そのユーザーのテンプレートは透過的にグローバルテンプレートをオーバーライドしますが、それは自分自身に対してのみであることに注意してください。同じGUIDで複数のテンプレートが存在する場合は、以下の順序で優先されます。

* ユーザーが所有しているもの
* ユーザーが属しているグループで共有されているもの
* グローバル

これは、ユーザーがグローバルなダッシュボードにアクセスしている場合、同じGUIDでテンプレートの独自のコピーを作成することで、ダッシュボードが参照する特定のテンプレートをオーバーライドできることを意味します。実際には、このようなことはほとんどないはずです。

### GUIDとThingUUIDの違いによるテンプレートへのアクセス

一般ユーザーは常にGUIDでテンプレートにアクセスする必要があります。管理者ユーザーは、ThingUUIDでテンプレートを参照することができますが、リクエストURLに`?admin=true`パラメータを設定する必要があります。

## テンプレートの作成

テンプレートを作成するには、`/api/templates`にPOSTを発行します。ボディは、任意の有効なJSONを含む「Contents」フィールドと、オプションでGUID、ラベル、名前、説明を含むJSON構造でなければなりません。以下のものはすべて有効です。

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  }
}
```

```
{
  "GUID": "ce95b152-d47f-443f-884b-e0b506a215be",
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  }
}
```

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate"
}
```

```
{
  "GUID": "ce95b152-d47f-443f-884b-e0b506a215be",
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate"
}
```

```
{
  "Contents": {
    "query": "tag=json json ip==%%IP%% | table",
    "required": true,
    "testValue": "\"10.0.0.1\"",
    "variable": "%%IP%%",
    "variableDescription": "the IP to investigate",
    "variableLabel": "IP address"
  },
  "Name": "mytemplate",
  "Labels": [
    "suits",
    "ladders"
  ]
}
```

APIは、新しく作成されたテンプレートのGUIDを応答します。リクエストにGUIDが指定されている場合は、そのGUIDが使用されます。GUIDが指定されていない場合は、ランダムなGUIDが生成されます。

注：現時点では、`UID`、`GIDs`、`Global`の各フィールドは、テンプレートの作成時には設定できません。これらのフィールドは、アップデートコールで設定する必要があります。

## テンプレートの一覧表示

あるユーザーが利用できるすべてのテンプレートをリストアップするには、`/api/templates`を GET します。結果として、テンプレートの配列が得られます。

```
[
  {
    "ThingUUID": "1b36a1d7-a5ac-11ea-b07e-7085c2d881ce",
    "UID": 1,
    "GIDs": [
      6,
      8
    ],
    "Global": false,
    "GUID": "780b1d31-e46b-4460-ad83-2fc11c34a162",
    "Name": "json ip",
    "Description": "JSON tag, filter by IP",
    "Contents": {
      "variable": "%%IP%%",
      "query": "tag=json* json ip==%%IP%% | table",
      "variableLabel": "IP address",
      "variableDescription": "the IP to investigate!",
      "required": true,
      "testValue": "\"10.0.0.1\""
    },
    "Updated": "2020-09-01T15:01:18.354750806-06:00",
    "Labels": [
      "test"
    ]
  }
]

```

## 一つのテンプレートを取得

1つのテンプレートを取得するには、`/api/templates/<guid>`に対してGETリクエストを発行します。例えば、`/api/templates/780b1d31-e46b-4460-ad83-2fc11c34a162`をGETすると、そのテンプレートのコンテンツが返ってきます。

```
{
  "ThingUUID": "1b36a1d7-a5ac-11ea-b07e-7085c2d881ce",
  "UID": 1,
  "GIDs": [
    6,
    8
  ],
  "Global": false,
  "GUID": "780b1d31-e46b-4460-ad83-2fc11c34a162",
  "Name": "json ip",
  "Description": "JSON tag, filter by IP",
  "Contents": {
    "variable": "%%IP%%",
    "query": "tag=json* json ip==%%IP%% | table",
    "variableLabel": "IP address",
    "variableDescription": "the IP to investigate!",
    "required": true,
    "testValue": "\"10.0.0.1\""
  },
  "Updated": "2020-09-01T15:01:18.354750806-06:00",
  "Labels": [
    "test"
  ]
}
```

なお、管理者は、ThingUUIDとadminパラメータを使用して、この特定のテンプレートを明示的に取得することができます。

## テンプレートの更新

テンプレートを更新するには、`/api/templates/<guid>`にPUTリクエストを発行します。 リクエストの本文は、必要な要素を変更して、同じパスのGETによって返されるものと同じである必要があります。 GUIDとThingUUIDは変更できないことに注意してください。 次のフィールドのみを変更できます：

* Contents： テンプレートの実際のボディ／コンテンツ
* Name： テンプレートの名前を変更
* Description： テンプレートの説明を変更
* GIDs： 32ビット整数のグループIDの配列を設定することができます。
* UID： (管理者のみ) 32ビット整数で設定可能
* Global： (Admin only) 真偽値を設定します。グローバルテンプレートはすべてのユーザーに表示されます。

注：これらのフィールドを空白にすると、そのフィールドのNULL値でテンプレートが更新されます。

## テンプレートの削除

テンプレートを削除するには、`/api/templates/<guid>`にDELETEリクエストを発行します。

## 管理者権限

管理者ユーザーは、システム上のすべてのテンプレートを表示したり、修正したり、削除したりする必要がある場合があります。GUIDは必ずしも一意ではないので、管理者APIはGravwellがアイテムを保存するために内部的に使用している一意のUUIDを参照する必要があります。上記のテンプレートリストの例では、"ThingUUID "というフィールドが含まれていることに注意してください。これは、そのテンプレートの内部的な一意の識別子です。

管理者ユーザーは、`/api/templates?admin=true`のGETリクエストで、システム内の全テンプレートのグローバルリストを取得することができます。

その後、管理者は、`/api/templates/<ThingUUID>?admin=true`へのPUTで、希望するテンプレートのThingUUID値に置き換えて、特定のテンプレートを更新することができます。削除の場合も同じパターンです。

管理者は、`/api/templates/<ThingUUID>?admin=true`へのGETまたはDELETEリクエスト（それぞれ）で、特定のテンプレートにアクセスまたは削除することができます。