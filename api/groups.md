# グループAPI

このページでは、ユーザーグループと対話するためのAPIについて説明します。

`/api/groups` や `/api/groups/` パス以下の他の URL に送られたリクエストは、ユーザアカウントで動作します。ウェブサーバは正常なリクエストにはStatusOK(200)を、エラーには(エラーに応じて)400-500のステータスを送信します。

注意: グループへのユーザーの追加や削除は、グループAPIではなく、[ユーザーアカウントAPI](account.md)を使用して行います。これらのAPIは純粋にグループの作成、削除、記述のためのものです。

## データタイプ

このページのAPIは、主にグループ詳細構造を扱います。グループ詳細構造は以下の通りです。

```
{
  GID: int,		// Numeric group ID
  Name: string, // Group name
  Desc: string, // Group description
  Synced: bool	// Tracks if group is synchronized with the datastore (can usually be ignored)
}
```

## 管理者のみグループの追加

新しいグループを追加するには、`/api/groups`にPOSTリクエストを送信します。リクエストの本文には、以下のようにグループの名前と説明を定義する構造体が含まれていなければなりません。

```
{
	"Name": "newgroup",
	"Desc": "This is the new group"
}
```

サーバーはリクエストを解析してグループを作成しようとします。成功した場合、200のステータスコードと新しいグループのGIDをレスポンスの本文に記載して応答します。

## 管理者のみグループの削除

グループを削除するには、`/api/groups/{gid}`にDELETEリクエストを送ります。

## 管理者のみグループ情報の更新

グループの情報を変更するには、`/api/groups/{gid}`にPUTリクエストを送る。リクエストの本文には、以下のようにグループの詳細構造が含まれている必要があります。

```
{
	"GID": 3,
	"Name": "newname",
	"Desc": "this is the new description",
}
```

Name または Desc フィールドが除外されている場合は、現在の値が保持されます。

## すべてのグループをリストアップ

システム上のすべてのグループのリストを取得するには、`/api/groups`にGETリクエストを送る。レスポンスには、グループの詳細構造の配列が含まれます。

```
[
    {
        "Desc": "bar",
        "GID": 1,
        "Name": "foo",
        "Synced": true
    },
    {
        "Desc": "",
        "GID": 7,
        "Name": "gravwell-users",
        "Synced": false
    },
    {
        "Desc": "",
        "GID": 8,
        "Name": "testgroup",
        "Synced": false
    }
]
```

## グループの情報を取得する

管理者や特定のグループのメンバーは、特定のグループに関する情報を取得することができます。GETリクエストを `/api/groups/{gid}` に発行します。レスポンスには、グループの詳細が含まれます。

```
{
    "Desc": "",
    "GID": 7,
    "Name": "gravwell-users",
    "Synced": false
}
```

## グループ内のユーザーをリストアップする

管理者や特定のグループのメンバーは、`/api/groups/{gid}/members`にGETリクエストを発行することで、そのグループのユーザーのリストを問い合わせることができます。応答は、ユーザ詳細構造体の配列となります。

```
[
    {
        "Admin": false,
        "Email": "joe@gravwell.io",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            },
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            },
            {
                "Desc": "",
                "GID": 8,
                "Name": "testgroup",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Joe User",
        "Synced": false,
        "TS": "2020-08-10T14:49:30.72782227-06:00",
        "UID": 16,
        "User": "joe"
    },
    {
        "Admin": false,
        "Email": "bkeaton@example.net",
        "Groups": [
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Buster Keaton",
        "Synced": false,
        "TS": "0001-01-01T00:00:00Z",
        "UID": 17,
        "User": "buster"
    }
]
```
