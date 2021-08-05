# プレイブック Web API

この API は、Gravwell でプレイブックを操作するために使用されます。プレイブックは、メモや検索クエリをユーザーフォーマットのドキュメントにまとめるためのユーザーフレンドリーな方法です。

プレイブックの構造には、以下のコンポーネントが含まれています:

* `UUID`: インストール時に設定されるプレイブックの一意の識別子。
* `GUID`: プレイブックのグローバル名。これはプレイブックのオリジナル作成者によって設定され、プレイブックがキットにバンドルされて別のシステムにインストールされた場合でも同じままです。各ユーザーは与えられた GUID を持つプレイブックを 1 つだけ持つことができますが、複数のユーザーがそれぞれ同じ GUID を持つプレイブックのコピーを持つことができます。
* `UID`: プレイブックの所有者のユーザーID。
* `GIDs`: このプレイブックへのアクセスを許可されたグループIDのリスト。
* `Global`: ブール値のフラグ。設定されている場合、システム上のすべてのユーザがそのプレイブックを閲覧することができます。
* `Name`: わかりやすいプレイブックの名前。
* `Desc`: プレイブックの説明。
* `Body`: プレイブックの内容を格納するバイト配列。
* `Metadata`: クライアントが使用するプレイブックのメタデータを格納するバイト配列。
* `Labels`: プレイブックに適用するオプションのラベルを含む文字列の配列。
* `LastUpdated`: プレイブックが最後に修正された時刻を示すタイムスタンプ。
* `Author`: プレイブックの作者に関する情報を含む構造体（後述）。
* `Synced`: Gravwellで内部的に使用されています。

UUID と GUID フィールドは、すべての API 呼び出しで互換性を持って使用できます。これは、キットに含まれるプレイブックが、キットのインストール間で同一のGUIDを含むリンクを使用して、互いにリンクすることができるようにするためです。

作者情報の構造には以下のフィールドが含まれており、いずれかのフィールドは空白のままにしておくことができます:

* `Name`: 作者の名前。
* `Email`: 作者のメールアドレス。
* `Company`: 作者の会社。
* `URL`: 作者の詳細についてのウェブアドレス。

## プレイブックのリストアップ

プレイブックをリストアップするには、`/api/playbooks`にGETリクエストを送ります。サーバは、ユーザが閲覧権限を持つプレイブックの構造体の配列を返します:

```
[
  {
    "UUID": "2cbc8500-5fc5-453f-b292-8386fe412f5b",
    "GUID": "c9da126b-1608-4740-a7cd-45495e8341a3",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Netflow V5 Playbook",
    "Desc": "A top-level playbook for netflow, with background and starting points.",
    "Body": "",
    "Metadata": "eyJkYXNoYm9hcmRzIjpbXSwiYXR0YWNobWVudHMiOlt7ImNvbnRleHQiOiJjb3ZlciIsImZpbGVHVUlEIjoiNDhjNmIwZWYtNmU3Ni00MjA4LWJjYTctMGI5NWU0NzAwYmRkIiwidHlwZSI6ImltYWdlIn1dfQ==",
    "Labels": [
      "netflow",
      "netflow-v5",
      "kit/io.gravwell.netflowv5"
    ],
    "LastUpdated": "2020-08-14T16:17:03.778971838-06:00",
    "Author": {
      "Name": "John Floren",
      "Email": "john@example.org",
      "Company": "Gravwell",
      "URL": "http://grawell.io"
    },
    "Synced": false
  },
  {
    "UUID": "973fcc22-1964-4efa-848c-7196ac67094e",
    "GUID": "dbd84b95-11b7-450d-9111-9bb33d63741b",
    "UID": 1,
    "GIDs": [
      0
    ],
    "Global": false,
    "Name": "Network Enrichment Kit Overview",
    "Desc": "",
    "Body": "",
    "Metadata": "eyJkYXNoYm9hcmRzIjpbXSwiYXR0YWNobWVudHMiOlt7ImNvbnRleHQiOiJjb3ZlciIsImZpbGVHVUlEIjoiOGIwZjQzMjItOTY1My00OTQyLWJkODctY2Y4ZWM5NjZmNmFmIiwidHlwZSI6ImltYWdlIn1dfQ==",
    "Labels": [
      "kit/io.gravwell.networkenrichment"
    ],
    "LastUpdated": "2020-08-05T12:14:48.739069332-06:00",
    "Author": {
      "Name": "John Floren",
      "Email": "john@example.org",
      "Company": "Gravwell",
      "URL": "http://grawell.io"
    },
    "Synced": false
  }
]
```

bodyが空であることに注意してください。プレイブックは非常に大きくなることがあるので、すべてのプレイブックをリストアップする際には、bodyが省略されます。

URLに`?admin=true`パラメータを追加すると、ユーザーが管理者であれば、システム上の*すべての*プレイブックのリストを返します。

## プレイブックの取得

Bodyを含む特定のプレイブックを取得するには、`/api/playbooks/<uuid>`にGETリクエストを送信します。ウェブサーバは、一致するUUIDフィールドを持つプレイブックを見つけようとします; それが成功しなかった場合は、ユーザが読めるプレイブックを以下の優先順位で探します:

* 1：ユーザーが所有しているプレイブック。
* 2：ユーザーのグループで共有しているプレイブックです。
* 3：グローバルフラグが設定されているプレイブック。

## プレイブックの作成

プレイブックは、`/api/playbooks`にPOSTリクエストを送ることで作成されます。リクエストの本文には、ユーザーが設定したいフィールドが含まれていなければなりません; UUID、UID、LastUpdated、Syncedフィールドが設定されている場合、サーバーはそれらを無視することに注意してください。

```
{
    "Body": <contents of the playbook>,
	"Metadata": <any desired metadata>,
    "Name": "ssh syslog",
    "Desc": "A playbook for monitoring syslog entries for ssh sessions",
    "GIDs": null,
    "Global": true,
	"Author": {
		"Name": "Dean Martin"
	},
    "Labels": [
        "syslog"
    ]
}
```

サーバは新しく作成されたプレイブックのUUIDで応答します。リクエストに `GUID` フィールドが設定されていれば、サーバはそれを使用し、そうでなければ新しいものを生成します。

## プレイブックの修正

プレイブックの内容を更新するには、`/api/playbooks/<uuid>`にPUTリクエストを送信します。リクエストの本文には、更新するプレイブックの構造が含まれていなければなりません。UUID、GUID、LastUpdated、Syncedフィールドの変更は無視されることに注意してください。管理者はUIDフィールドを変更することができますが、通常のユーザーは変更できません。

注意: フィールドの内容を更新するつもりがない場合は、リクエストで元の値を送るべきです。サーバは、例えば「Desc」フィールドが未設定の場合、元の値を保持したいのか、フィールドをクリアしたいのかを知る方法がありません。

## プレイブックの削除

プレイブックを削除するには、`/api/playbooks/<uuid>`にDELETEリクエストを送ります。
