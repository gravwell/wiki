# ログインサブシステム

## ログイン

ログインするには、以下のような構造のJSONを `/api/login` にPOSTします。

```
{
    "User": "username",
    "Pass": "password"
}
```

と表示され、サーバーはログインが成功したかどうかを示す次のような応答をします:

```
{
  "LoginStatus":true,
  "JWT":"reallylongjsonwebtokenstringishere"
}

```

ログインに失敗した場合、サーバーは "reason "プロパティを持つ構造体を返します:
```
{
  "LoginStatus":false,
  "Reason":"Invalid username or password"
}
```

JSONを送信する代わりに、ログインPOSTリクエストにフォームフィールド「User」と「Pass」を設定することもできます。

## ログアウト

* PUT /api/logout - 現在のインスタンスをログアウトします。
* DELETE /api/logout - すべてのユーザーのインスタンスをログアウトします。

## JWT保護は、ファイルのダウンロード操作に使用されていないすべてのリクエストに適用されます。
ログインAPIから受け取ったJWTは、他のすべてのAPIリクエストのAuthorization Bearerヘッダーとして含める必要があります。

```Authorization: Bearer reallylongjsonwebtokenstringishere```

### ウェブソケット認証

便利なことに、ウェブソケット API のエンドポイントは、`Sec-Websocket-Protocol` ヘッダー値で JWT トークンを探します。 多くのウェブソケットの実装は、ヘッダ値の受け渡しを適切にサポートしていないため、ウェブソケットのサブプロトコルネゴシエーションヘッダをオーバーロードします。 APIエンドポイントは、標準的な `Authentication` ヘッダー値も引き続き探します。

## アクティブなセッションの表示

GET を `/api/users/{id}/sessions` に送ると、JSON のチャンクが返されます。 管理者はすべてのユーザーのセッションをリクエストできますが、ユーザーは自分のセッションしかリクエストできません。

```
{
    "Sessions": [
        {
            "LastHit": "2020-08-04T15:28:12.601899275-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        },
        {
            "LastHit": "2020-08-03T23:59:53.807610997-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        },
        {
            "LastHit": "2020-08-04T09:45:48.291770859-06:00",
            "Origin": "127.0.0.1",
            "Synced": false,
            "TempSession": false
        }
    ],
    "UID": 1,
    "User": "admin"
}
```
