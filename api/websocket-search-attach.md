## Reattaching to an existing search
### 既存の検索に再接続する
既存の検索へのアタッチは、接続時に "search"、 "parse"、および "PONG"とネゴシエートする必要がある "attach"という名前の別のサブプロトコルを使用して既存のWebソケットを介して実行されます。

#### 許可
検索に参加するには、ユーザーは検索が割り当てられているグループのメンバー、所有者、または管理者である必要があります。ユーザーが検索にアタッチすることを許可されていない場合は、許可拒否メッセージがソケットを介して送り返されます。

### 検索への添付を要求する
基本的には単にIDを渡すだけで、サーバーはError、または新しいSubProto、RenderMod、RenderCmd、およびSearchInfoのいずれかで応答します。

新しいサブプロトコルは、新しく接続された検索を処理するために作成される新しいサブプロトコルの名前です。つまり、WebSocketを起動してから、同時に多数の検索にアタッチすることができます。

### 転送例
以下は、検索リストを要求してからそれに添付するクライアントの例です。

検索リストを尋ねる
```
WEB GET /api/searchctrl:
[
        {
                "ID": "004081950",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "560752652",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        },
        {
                "ID": "608274427",
                "UID": 1,
                "GID": 0,
                "State": "DORMANT",
                "AttachedClients": 0
        }
]
```

「付加」サブプロトコル上のトランザクション
```
SUBPROTO PUT attach:
{
        "ID": "560752652"
}
SUBPROTO GET attach:
{
        "Subproto": "attach5",
        "RendererMod": "text",
        "Info": {
                "ID": "560752652",
                "UID": 1,
                "UserQuery": "grep paco",
                "EffectiveQuery": "grep paco | text",
                "StartRange": "2017-01-14T06:08:32.024425042-07:00",
                "EndRange": "2017-01-14T16:08:32.024425042-07:00",
                "Started": "2017-01-14T16:08:32.025987218-07:00",
                "Finished": "2017-01-14T16:08:32.746482323-07:00",
                "StoreSize": 0,
                "IndexSize": 0
        }
}
```
新しい添付検索のネゴシエーションを行った後は、通常レンダラーを利用するときと同じ古いAPIを引き続き使用してください。
```
SUBPROTO PUT attach5:
{
        "ID": 3
}
SUBPROTO GET attach5:
{
        "ID": 3,
        "EntryCount": 20000,
        "Finished": true
}
SUBPROTO PUT attach5:
{
        "ID": 16777218,
        "EntryRange": {
                "First": 0,
                "Last": 1024,
                "StartTS": "0001-01-01T00:00:00Z",
                "EndTS": "0001-01-01T00:00:00Z"
        }
}
```