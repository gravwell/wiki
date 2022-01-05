# ユーザーアカウントAPI

このページでは、ユーザーと対話するためのAPIについて説明します。

`/api/users` や `/api/users/` パス以下の他の URL に送られたリクエストは、ユーザアカウントに対して操作します。ウェブサーバは正常なリクエストに対してはStatusOK (200) を送信し、エラーに対しては400-500のステータスを送信します (エラーの内容に依存します)。

デフォルトの "admin" アカウントのユーザー名を変更したり削除したり、adminステータスから降格させたりすることはできません。他のユーザーアカウントにも管理者権限を割り当てることができます。

## データタイプ

このページのAPIは、主にユーザーとグループの詳細構造を扱います。 ユーザーの詳細構造は次のとおりです：

```
{
  UID: int,			// Numeric user ID
  User: string,		// Username e.g. "jdoe"
  Name: string,		// User's real name e.g. "John Doe"
  Email: string,	// User's email address e.g. "john@example.org"
  Admin: bool,		// Set to true if the user has admin rights
  Locked: bool,		// Set to true if the user's account has been locked
  DefaultGID: int,	// The user's default search group
  TS: string,		// Contains a timestamp representing when the user was last active
  Synced: bool,		// Tracks if user is synchronized with the datastore (can usually be ignored)
  Groups: Array<GroupDetails>	// An array of groups to which the user belongs
}
```

グループの詳細構造は次のとおりです：

```
{
  GID: int,		// Numeric group ID
  Name: string, // Group name
  Desc: string, // Group description
  Synced: bool	// Tracks if group is synchronized with the datastore (can usually be ignored)
}
```

## ログインとログアウト

ウェブサーバーで認証を行うためのAPIは[こちら](login.md)に記述されています。


## 現在のユーザー情報を取得

 /api/info/whoamiでGETすると、現在のアカウント情報が返されます

```
{
        "UID": 1,
        "User": "admin",
        "Name": "Sir changeme of change my password the third",
        "Email": "admin@admin.admin",
        "Admin": true,
        "Locked": false,
        "TS": "2016-12-05T17:05:33.121180268-07:00",
        "DefaultGID": 0,
        "Groups": [
                {
                        "GID": 1,
                        "Name": "foo",
                        "Desc": "This is the foo group"
                }
        ]
}

```

## 全アカウントの一覧表示

GETリクエストを `/api/users` に送信すると、すべてのアカウントに関する情報を含むJSONパケットが返されます。 返されるJSONは以下の通りです：


```
[
    {
        "Admin": true,
        "Email": "admin@admin.admin",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Admin John",
        "Synced": true,
        "TS": "2020-07-30T08:51:35.205998608-06:00",
        "UID": 1,
        "User": "admin"
    },
    {
        "Admin": false,
        "Email": "john@example.net",
        "Groups": [
            {
                "Desc": "bar",
                "GID": 1,
                "Name": "foo",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "John",
        "Synced": false,
        "TS": "2020-08-10T14:47:44.58356179-06:00",
        "UID": 14,
        "User": "john"
    },
    {
        "Admin": false,
        "Email": "joe@gravwell.io",
        "Groups": [
            {
                "Desc": "",
                "GID": 7,
                "Name": "gravwell-users",
                "Synced": false
            }
        ],
        "Locked": false,
        "Name": "Joe User",
        "Synced": false,
        "TS": "2020-08-10T14:49:30.72782227-06:00",
        "UID": 16,
        "User": "joe"
    }
]
```

## 単一のユーザー情報を取得する

あるユーザーのアカウント情報を取得するには、`/api/users/{id}/`に GET を送信してください。 管理者は任意のアカウントを取得することができますが、非管理者は自分のアカウント情報のみを取得することができます。 200のレスポンスには有効なJSONが含まれ、400-500のレスポンスには誰かがリクエストを失敗させたことを意味し、エラーメッセージが含まれます。 バックエンドは、このページの冒頭で説明したようなユーザー情報を含むJSONパケットで応答します。

```
{
    "Admin": true,
    "DefaultGID": 1,
    "Email": "admin@admin.admin",
    "Groups": [
        {
            "Desc": "bar",
            "GID": 1,
            "Name": "foo",
            "Synced": false
        }
    ],
    "Locked": false,
    "Name": "Admin",
    "Synced": false,
    "TS": "2020-07-30T08:51:35.205998608-06:00",
    "UID": 1,
    "User": "admin"
}
```

## 新しいユーザーの追加

新しいユーザーを追加するには、管理者アカウントはこれらのフィールドを含むリクエストを `/api/users` にPOSTすることができます。

```
{
     User: "buster",
     Pass: "gr4vwellRulez",
     Name: "Buster Keaton",
     Email: "bkeaton@example.net"
     Admin: false,
}
```

すべてのフィールドに入力する必要があります。 バックエンドは新しいユーザーの UID (例: `17`) で応答します。

## ユーザーアカウントをロックする
アカウントをロックするには、`/api/users/{id}/lock` に空のPUTを送信します。ここで {id} はロックされるユーザーのUIDです。 ユーザーアカウントをロックすると、そのユーザーのアクティブなセッションは直ちにログアウトし、新たなログインを防ぐことができます。

## ユーザーアカウントのロック解除
アカウントのロックを解除するには、`/api/users/{id}/lock`宛に空のDELETEを送信します（{id}はロックを解除されるユーザーのUID）。 このアクションが許可されている場合、ウェブサーバーはアカウントが実際にロック解除されたかどうかに関係なく、成功として応答します。 ロックされたアカウントをロックすると、そのアカウントがロックされている状態で終了するからです。 そのため、すべてうまくいきます。 ロック解除と同じです。

## ユーザー情報の変更
ユーザー情報を変更するには、クライアントは更新したいフィールドを指定して `/api/users/{id}` というPUTを送信する必要があります。

```
{
     User: "chuck",
     Name: "Chuck Testa",
     Email: "chuck@testa.net"
}
```

入力されていないフィールドは無視されます。 バックエンドは上記のように標準的なレスポンスJSONで応答します。 現在のユーザーが管理者ではなく、自分のアカウントを変更していない場合、リクエストは拒否されます。 管理者は任意のアカウントの情報を変更することができます。 プライマリ管理者アカウント(UID 0)は管理者ステータスを変更することができません。 バックエンドは成功時には200、エラー時には400～500で応答します。 エラーメッセージはレスポンスの本文に返され、表示できるようになっています。


## ユーザーのパスワードを変更する

パスワードを変更するには、`/api/users/{id}/pwd`というURLにPUTしてください。

```
{
     OrigPass: "my old password was bad",
     NewPass: "thisis mynewpassword",
}
```

注：現在のユーザが管理者で、自分が所有していないアカウントのパスワードを変更する場合、OrigPassフィールドは必須ではありません。

## ユーザーの削除
ユーザーを削除するには、クライアントは `/api/users/{id}/` に DELETE を送信し、`{id}` には削除する UID を設定する必要があります。 この機能は管理者のみが使用することができ、ユーザが自分自身のアカウントを削除することはできません。 プライマリ管理者(UID 1)は削除することができません。

## ユーザーの管理者状態の取得/設定
ユーザの管理者権限を変更したり、管理者権限を問い合わせたりするには、`/api/users/{id}/admin` に空の本文でリクエストを送信してください。メソッドには、希望するアクションを指定します。 ウェブサーバは成功した場合は200、失敗した場合はその種類に応じて400-500で応答します。

* GET - 現在の管理者のステータスを返します。
* PUT - ユーザーを管理者として設定し、新しいステータスを返します。
* DELETE - ユーザーの管理者ステータスを削除します。

成功時のJSONの例：

```
{
     UID: 1,
     Admin: true
}
```

## ユーザーセッションの取得

ユーザは、`/api/users/{id}/sessions`にGETリクエストを発行することで、自分のセッションのリストを取得することができます。管理者は、任意のUIDに対してリクエストを発行することができます。レスポンスは、以下に示すように、セッションオブジェクトの配列になります。

```
{
    "Sessions": [
        {
            "LastHit": "2020-08-11T14:41:23.829366-06:00",
            "Origin": "::1",
            "Synced": false,
            "TempSession": false
        }
    ],
    "UID": 1,
    "User": "admin"
}
```

## デフォルトの検索グループを設定します。

各ユーザーは、デフォルトの検索グループを設定することができます。設定された場合、そのユーザが実行するすべてのクエリは、デフォルトでそのグループと共有されます。デフォルトの検索グループを設定するには、 `/api/users/{id}/searchgroup` にPUTリクエストを発行し、リクエスト本文に以下のように希望するグループを指定します：

```
{
	"GID": 3
}
```

## デフォルトの検索グループの取得

各ユーザーはデフォルトの検索グループを設定することができます。設定された場合、そのユーザが実行するすべてのクエリは、デフォルトでそのグループと共有されます。ユーザーの検索グループを取得するには、`/api/users/{id}/searchgroup`に対してGETリクエストを発行してください。サーバのレスポンスには、例えば `3` のような整数のGIDが含まれます。

## ユーザープリファレンス

「ユーザープリファレンス "は、フロントエンドがユーザに設定する任意のJSON blobから構成されます。これにより、フロントエンドはユーザーのUI設定（配色など）を複数のコンピュータに保存することができます。ユーザープリファレンスは、1) 当該ユーザ、または 2) 管理者ユーザのみが設定することができます。

### ユーザープリファレンスの設定

ユーザーのプリファレンスを設定するには、`/api/users/{id}/preferences`に対してPUTリクエストを発行してください。リクエストの本文は有効なJSONでなければなりません。JSONの実際のコンテンツはウェブサーバによって無視されます。

### ユーザープリファレンスの取得

ユーザーのプリファレンスを取得するには、`/api/users/{id}/preferences`に対してGETリクエストを発行してください。レスポンスの本文には、以前に設定されたプリファレンスが含まれます。

### ユーザープリファレンスのクリア

ユーザーのプリファレンスをすべて消去するには、`/api/users/{id}/preferences`に対してDELETEを実行してください。

### 管理者のみ 全ユーザープリファレンスの取得

管理者は `/api/users/preferences` に対して GET リクエストを発行して、すべてのユーザーのプリファレンスの完全なリストを取得することができます。レスポンスには、ユーザーのプリファレンス（Name == "prefs"）とユーザーの *email* 設定（Name == "emailSettings" ）の両方が含まれます。電子メール設定を設定・変更するためのインターフェースは、別の場所で説明されています。

```
[
    {
        "Data": "ImJhciBiYXogcXV1eCBhc2Rmc2FkZmRzZiI=",
        "Name": "prefs",
        "Synced": false,
        "UID": 1,
        "Updated": "2020-07-13T13:50:08.286015047-06:00"
    },
    {
        "Data": "ImZvbyI=",
        "Name": "emailSettings",
        "Synced": false,
        "UID": 1,
        "Updated": "2019-12-18T09:45:38.683780752-07:00"
    }
]
```

## ユーザーをグループに追加

管理者は、`/api/users/{id}/group`にPOSTリクエストを送信することで、ユーザーをグループに追加することができます。リクエストの本文には、ユーザのグループメンバシップに*追加*されるべきGIDのリストが含まれていなければなりません。したがって、GID 8のユーザーをグループに追加するには、他のグループメンバーシップに関係なく、以下のように送信します：

```
{
	"GIDs": [8]
}
```

## グループからユーザーを削除する

特定のグループからユーザを削除するには、管理者は `/api/users/{id}/group/{gid}` にDELETEリクエストを送ることができます。ここで {id} はユーザーのUID、{gid} はグループのGIDです。

## ユーザーグループの取得

ユーザが所属するグループのリストを取得するには、ユーザ(または管理者)は `/api/users/{id}/group` にGETリクエストを発行しなければなりません。レスポンスは、以下のようなグループ詳細構造の配列になります：
```
[
    {
        "Desc": "foo group",
        "GID": 1,
        "Name": "foo",
        "Synced": false
    }
]
```
