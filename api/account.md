# ユーザーアカウントAPI

このページでは、ユーザーと対話するためのAPIについて説明します。

`/api/users` や `/api/users/` パス以下の他の URL に送られたリクエストは、ユーザアカウントで動作します。ウェブサーバは正常なリクエストにはStatusOK (200)を、エラーには(エラーに応じて)400-500のステータスを送信します。

デフォルトの「admin」アカウントのユーザー名を変更または削除したり、アカウントを管理者ステータスから降格したりすることはできません。 他のユーザーアカウントにも管理者権限が割り当てられている場合があります。

## データタイプ

このページのAPIは、主にユーザーとグループの詳細構造を扱います。 ユーザーの詳細構造は次のとおりです。

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

グループの詳細構造は次のとおりです。

```
{
  GID: int,		// Numeric group ID
  Name: string, // Group name
  Desc: string, // Group description
  Synced: bool	// Tracks if group is synchronized with the datastore (can usually be ignored)
}
```

## ログインとログアウト

ウェブサーバで認証するためのAPIは[このドキュメント](login.md)です。

## 現在のユーザー情報を取得

GET on /api/info/whoami は現在のアカウント情報を返します。

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

## すべてのアカウントをリストアップ

GETリクエストを `/api/users` に送信すると、すべてのアカウントに関する情報を含むJSONパケットが返されます。 返されるJSONは以下の通りです。

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

単一のユーザのアカウント情報を取得するには、`/api/users/{id}/`にGETを送ってください。 管理者は任意のアカウントを取得することができますが、非管理者は自分のアカウント情報のみを取得することができます。 200のレスポンスは本文に有効なJSONが含まれていますが、400～500のレスポンスは誰かがリクエストを破棄したことを意味し、本文にはエラーメッセージが含まれています。 バックエンドは、このページのトップで説明したようにユーザーの詳細を含む JSON パケットで応答します。

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

すべてのフィールドを入力しなければなりません。 バックエンドは新しいユーザーのUIDで応答します。

## ユーザーアカウントをロックする
アカウントをロックするには、空のPUTを `/api/users/{id}/lock` に送ります。 ユーザアカウントをロックすると、そのユーザのアクティブなセッションが直ちにログアウトされ、新規ログインができなくなります。

## ユーザーアカウントのロックを解除する
アカウントのロックを解除するには、空のDELETEを `/api/users/{id}/lock` に送信します。 このアクションが許可されていれば、アカウントが実際にロックされていたかどうかに関係なく、ウェブサーバは成功して応答します。 ロックされたアカウントをロックすると、そのアカウントがロックされた状態で終了するため、このようにしています。 これで問題ありません。 ロック解除も同じです。

## ユーザー情報変更
ユーザー情報を変更するには、クライアントは更新したいフィールドを指定してPUT `/api/users/{id}`を送信しなければなりません。

```
{
     User: "chuck",
     Name: "Chuck Testa",
     Email: "chuck@testa.net"
}
```

入力されていないフィールドは無視されます。 バックエンドは上記のように標準的なレスポンスJSONで応答します。 現在のユーザーが管理者ではなく、自分のアカウントを変更していない場合、リクエストは拒否されます。 管理者は任意のアカウントの情報を変更することができます。 プライマリ管理者アカウント(UIDゼロ)は管理者ステータスを変更することができません。 バックエンドは成功時には200、エラー時には400～500で応答します。 エラーメッセージはレスポンスの本文に返され、表示できるようになっています。


## ユーザーのパスワードを変更する

パスワードを変更するには、URL `/api/users/{id}/pwd` にPUTを発行する。

```
{
     OrigPass: "my old password was bad",
     NewPass: "thisis mynewpassword",
}
```

注: 現在のユーザーが管理者であり、彼らが所有していないアカウントのパスワードを変更する場合、OrigPassフィールドは必要ありません。

## ユーザーを削除する
ユーザーを削除するには、クライアントは `/api/users/{id}/` に削除するUIDを `{id}` に設定してDELETEを送信しなければなりません。 この機能を使用できるのは管理者のみであり、ユーザは自分のアカウントを削除することはできません。 プライマリ管理者 (UID 1) は削除できません。

## ユーザーの管理者ステータスの取得、設定
ユーザの管理者の状態を変更したり、管理者の状態を問い合わせたりするには、`/api/users/{id}/admin`に空のボディでリクエストを送ります。メソッドは希望するアクションを指定します。 ウェブサーバは失敗の種類に応じて、成功の場合は200、エラーの場合は400から500の値で応答します。

* GET - 現在の管理者の状態を返します。
* PUT - ユーザーを管理者に設定し、新しいステータスを返します。
* DELETE - ユーザの管理者ステータスを削除します。

成功時のJSONの例。

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

各ユーザーはデフォルトの検索グループを設定することができます。設定された場合、そのユーザーが実行したすべてのクエリは、デフォルトでそのグループと共有されます。デフォルトの検索グループを設定するには、`/api/users/{id}/searchgroup`にPUTリクエストを発行する。

```
{
	"GID": 3
}
```

## デフォルトの検索グループを取得

各ユーザーはデフォルトの検索グループを設定することができます。設定されている場合、そのユーザーが実行したすべてのクエリは、デフォルトでそのグループと共有されます。ユーザの検索グループを取得するには、`/api/users/{id}/searchgroup`にGETリクエストを発行する。サーバの応答は整数のGID、例えば `3` を含む。

## ユーザー設定

"ユーザー設定 "は、フロントエンドがユーザーのために設定した任意のJSONブロブで構成されています。これにより、フロントエンドはユーザーの UI 設定 (配色など) を複数のコンピュータに保存することができます。ユーザー設定は、1) 問題のユーザーか、2) 管理者ユーザーのみが設定できます。

### ユーザー設定の設定

ユーザー設定を設定するには、`/api/users/{id}/preferences`にPUTリクエストを発行します。リクエスト本文には有効なJSONが含まれていなければなりません; 実際のJSONの内容はウェブサーバによって無視されます。

### ユーザー設定の取得

ユーザ設定を取得するには、`/api/users/{id}/preferences`にGETリクエストを発行します。レスポンスの本文には、以前に設定された設定が含まれます。

### ユーザー設定をクリア

ユーザーの環境設定をクリアするには、`/api/users/{id}/preferences`でDELETEを発行します。

### 管理者のみすべてのユーザー設定を取得

管理者は `/api/users/preferences` にGETリクエストを発行して、すべてのユーザの設定の完全なリストを取得することができます。レスポンスには、ユーザの設定(Name == "prefs")とユーザの*email*設定(Name == "emailSettings")の両方が含まれます。電子メール設定を設定・変更するためのインターフェイスは別の場所で説明されています。

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

管理者は、`/api/users/{id}/group`にPOSTリクエストを送信することで、ユーザーをグループに追加することができます。リクエストの本文には、ユーザのグループメンバシップに*追加されるべきGIDのリストが含まれていなければなりません。したがって、GID 8のユーザーをグループに追加するには、他のグループメンバーシップに関係なく、以下のように送信します。

```
{
	"GIDs": [8]
}
```

## グループからユーザーを削除する

特定のグループからユーザを削除するには、管理者は `/api/users/{id}/group/{gid}` にDELETEリクエストを送ることができます。

## ユーザーグループの取得

ユーザが所属するグループのリストを取得するには、ユーザ(または管理者)は `/api/users/{id}/group` にGETリクエストを発行しなければなりません。レスポンスは、以下のようなグループ詳細構造の配列になります。
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
