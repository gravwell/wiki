# スケジュール検索API

このAPIは、スケジュールされた検索の作成と管理を可能にします。検索はランダムに生成されたIDで参照されます。

## スケジュール検索の構造

スケジュール検索には、以下の重要なフィールドが含まれます:

* ID： スケジュール検索のID
* GUID： この検索に固有のID、作成時に空白にしておくと、ランダムなGUIDが割り当てられます（これは標準的なユースケースとなります）。
* Owner： 検索の所有者のUIDです。
* Groups： この検索の結果を見ることを許可されたグループのリスト
* Global： 検索結果をすべてのユーザーに表示することを示すブール値（管理者のみがこのフィールドを設定できます）
* Name： このスケジュールされた検索の名前
* Description： スケジュールされた検索のテキストによる説明です。
* Schedule： cronと互換性のある文字列で、いつ実行するかを指定します。
* Permissions： パーミッション・ビットを格納するための64ビット整数
* Disabled： 真を設定すると、スケジュールされた検索が実行されないようにするブール値です。
* OneShot： 真に設定された場合、無効化されていない限り、スケジュールされた検索が可能な限り一度だけ実行されるようにするブーリアンです。
* LastRun： このスケジュール検索が最後に実行された時間です。
* LastRunDuration： 最後の実行にかかった時間
* LastSearchIDs： このスケジュール検索から最も最近実行された検索の検索IDを含む文字列の配列
* LastError：この検索の最後の実行から生じたすべてのエラー

検索が「標準」のスケジュール検索の場合は、以下のフィールドも設定されます:

* SearchString：実行するGravwell検索文
* Duration（期間）：検索をどれだけ遡って実行するかを秒単位で指定します。これは負の値でなければならない。
* SearchSinceLastRun: ブーリアン値、設定されている場合、Durationフィールドは無視され、代わりにLastRun時間から現在までの検索が実行されます。

一方、検索がスクリプトの場合は、次のフィールドが設定されます:

* Script: ankoのスクリプトを含む文字列

## ユーザーコマンド

このセクションのAPIコマンドは、どのユーザーでも実行できます。

### 予定されている検索の一覧表示

ユーザーが表示している（ユーザーが所有しているか、ユーザーのグループがアクセス可能とマークされている）すべてのスケジュールされた検索のリストを取得するには、`/api/scheduledsearches`に対してGETを実行します。結果は以下のようになります：

```
[
  {
    "ID": 1439174790,
    "GUID": "efd1813d-283f-447a-a056-729768326e7b",
    "Groups": null,
	"Global": false,
    "Name": "count",
    "Description": "count all entries",
    "Owner": 1,
    "Schedule": "* * * * *",
    "Permissions": 0,
    "Updated": "2019-05-21T16:01:01.036703243-06:00",
    "Disabled": false,
    "OneShot": false,
    "Synced": true,
    "SearchString": "tag=* count",
    "Duration": -3600,
    "SearchSinceLastRun": false,
    "Script": "",
    "PersistentMaps": {},
    "LastRun": "2019-05-21T16:01:00.013062447-06:00",
    "LastRunDuration": 1015958622,
    "LastSearchIDs": [
      "672586805"
    ],
    "LastError": ""
  }
]

```

この例では、UID 1（admin）が所有する "count"という名前の単一のスケジュールされた検索を示しています。1分ごとに実行され、過去1時間の間に`tag=*count`という検索を実行します。

### スケジュール検索の作成

新しいスケジュール検索を作成するには、スケジュール検索に関する情報を含む JSON 構造体を `/api/scheduledsearches` に POST リクエストします。標準的な検索を作成するには、SearchString と Duration フィールドに必ず入力してください。

```
{
  "Name": "myscheduledsearch",
  "Description": "a scheduled search",
  "Groups": [
    2
  ],
  "Global": false,
  "Schedule": "0 8 * * *",
  "SearchString": "tag=default grep foo",
  "Duration": -86400,
  "SearchSinceLastRun": false
}

```

また、SearchSinceLastRunフィールドがtrueに設定されている場合、検索エージェントはDurationを無視し（この新しい検索の最初の実行を除く）、代わりに最後の実行の時間から現在の時間までの検索を実行します。

スクリプトを使用してスケジュール検索を作成するには、"SearchString"および"Duration"フィールドの代わりに"Script"フィールドを入力します。両方が入力されている場合は、スクリプトが優先されます。

スケジュール検索は、Disabledフラグをtrueに設定して作成し、ユーザーの準備が整うまで実行しないようにすることができます。また、OneShotフラグをtrueに設定して作成すると、作成後すぐに検索が実行されるようになります。

サーバーは新しいスケジュール検索のIDを応答します。

### 特定のスケジュール検索を取得します

単一のスケジュール検索に関する情報は、`/api/scheduledsearches/{id}`のGETでアクセスできます。例えば、スケジュール検索のIDが1439174790の場合、`/api/scheduledsearches/1439174790`に問い合わせると、以下のような結果が得られます。

```
{
  "ID": 1439174790,
  "GUID": "efd1813d-283f-447a-a056-729768326e7b",
  "Groups": null,
  "Global": false,
  "Name": "count",
  "Description": "count all entries",
  "Owner": 1,
  "Schedule": "* * * * *",
  "Permissions": 0,
  "Updated": "2019-05-21T16:01:01.036703243-06:00",
  "Disabled": false,
  "OneShot": false,
  "Synced": true,
  "SearchString": "tag=* count",
  "Duration": -3600,
  "SearchSinceLastRun": false,
  "Script": "",
  "PersistentMaps": {},
  "LastRun": "2019-05-21T16:01:00.013062447-06:00",
  "LastRunDuration": 1015958622,
  "LastSearchIDs": [
    "672586805"
  ],
  "LastError": ""
}
```

スケジュールされた検索は、GUIDによってもフェッチできます。ただし、これはウェブサーバにとってより多くの作業を必要とするため、必要な場合にのみ使用すべきです。上記のスケジュール検索を取得するには、`/api/scheduledsearches/cdf011ae-7e60-46ec-827e-9d9fcb0ae66d`をGETしてください。

### 既存の検索の更新

スケジュールされた検索を修正するには、必要な変更を含む更新された構造を含む `/api/scheduledsearches/{id}` への HTTP PUT を行います。変更されていないフィールドもプッシュしないと、空の値で上書きされてしまうので注意が必要です。

更新できるフィールドは以下の通りです。

* Name
* Description
* Schedule
* SearchString
* Duration
* SearchSinceLastRun
* Script
* Groups
* Global (admin only)
* Disabled
* OneShot

スクリプトスケジュール検索は、Scriptフィールドを空にしてSearchStringとDurationをプッシュすることで、標準のスケジュール検索に変更することができます。同様に、標準のスケジュール検索は、スクリプトフィールドをプッシュしてSearchStringを空にすることで、スクリプトスケジュール検索に変換することができます。

### スケジュール検索のエラーを解除します

スケジュール検索構造体の LastError フィールドは、エラーが発生した場合に設定され、その後の実行が成功してもクリアされません。このフィールドは、`/api/scheduledsearches/{id}/error`をDELETEすることで手動でクリアできます。

### スケジュール検索の永続的な状態の消去

`api/scheduledsearches/{id}/state`をDELETEすると、LastErrorフィールドとスケジュール検索のパーシステントマップの両方がクリアされます。これにより、不正なスクリプトによって状態が破損した場合に、スケジュール検索をリセットすることができます。

### スケジュール検索の削除

既存のスケジュール検索を削除するには、`/api/scheduledsearches/{id}`に対してDELETEを実行します。

## 管理者コマンド

以下のコマンドは admin ユーザーのみが使用できます。

### すべての検索結果の表示

管理者ユーザーは、システム上のすべてのスケジュール検索を表示する必要がある場合があります。管理者ユーザーは、`/api/scheduledsearches?admin=true`のGETリクエストで、システム内のすべてのスケジュールされた検索のグローバルリストを取得できます。

スケジュールされた検索IDはシステム全体で一意であるため、管理者は`?admin=true`を指定することなく、検索の修正・削除・取得を行うことができますが、不必要にパラメータを追加してもエラーにはなりません。

### 特定のユーザーの検索結果を取得します

`API/scheduledsearches/user/{uid}`（`uid`は数字のユーザーID）に対してGETを実行すると、そのユーザーに属するすべての検索結果の配列が取得されます。

### 特定のユーザーの検索をすべて削除します

`api/scheduledsearches/user/{uid}`に対してDELETEを実行すると、指定したユーザーに属するすべてのスケジュール検索が削除されます。

### スケジュール検索のテストパーズの実行

scheduledsearches APIは、スケジュールされた検索を保存する前にテストするためのAPIを提供しています。 Parse APIは `/api/scheduledsearches/parse` にあり、PUTリクエストでアクセスします。 認証されたユーザは、既存のスケジュールされたスクリプトを保存したり変更したりすることなく、スケジュールされたスクリプトを送信して解析とチェックを行うことができます。

解析を実行するには、PUTリクエスト `/api/scheduledsearches/parse` のボディに、以下のJSON構造を送信します。

```
{
	Script string
}
```

APIは以下のJSON構造で応答します。
```
{
	OK bool
	Error string
	ErrorLine int
	ErrorColumn int
}
```

スクリプトが解析テストに合格した場合、レスポンスのOKフィールドには`true`が入ります。 Errorフィールドは省略され、ErrorLineとErrorColumnフィールドは両方とも`-1`になります。 提供されたスクリプトが正しく解析できなかった場合は、OKフィールドは`false`となり、Errorフィールドは失敗の理由を示し、ErrorLineとErrorColumnはスクリプトのどこでエラーが発生したかを示します。

ErrorLine と ErrorColumn フィールドは常に入力されるとは限りません。1の値は、スクリプト解析システムがスクリプト内のどこでエラーが発生したのかわからないことを示します。

以下は、リクエストとレスポンスの例です。

#### 有効なスクリプト
リクエスト
```
{
	"Script":"fmt = import(\"fmt\")\nfmt.Println(\"Hello\")\nfmt.Sstuff(\"Goodbye\")\n"
}
```

レスポンス
```
{
	"OK":true,
	"ErrorLine":-1,
	"ErrorColumn":-1
}
```

#### 無効なスクリプト
リクエスト
```
{
	"Script":"fmt = import(\"fmt\")\nfmt.Println(\"Hello\")\nfmt.Sstuff(\"Goodbye)\n"
}
```

レスポンス
```
{
	"OK":false,
	"Error":"syntax error",
	"ErrorLine":3,
	"ErrorColumn":21
}
```
