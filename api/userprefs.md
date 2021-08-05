## ユーザー設定
ユーザー設定APIは、ログイン間およびデバイス間で保持されるGUI設定を保存するために使用されます。

/api/users/{id}/preferencesをGETすると、JSONのチャンクが返されます。 管理者はすべてのユーザーのプリファレンスをリクエストできますが、ユーザーは自分のセッションしかリクエストできません。プリファレンスが存在しない場合は、nullを返します。

api/users/preferencesをGETすると、すべてのユーザーのプリファレンスが返されます。

ユーザープリファレンスを更新するために、/api/users/{id}/preferencesをPUTしてください。プリファレンスが存在しない場合は、提供されたJSONで更新します。このAPIではPOSTは行われません。PUTメソッドのペイロードは、JSON blobです。

関連するメソッドは、GET と PUT のみです。各ユーザーは、本来、1つだけの好みのJSON blobを持つべきです。

/api/users/{id}/preferencesをDELETEすると、設定が削除されます（管理者や自分で作成した場合）。


GETでJSONを返した例です:
```json
{
     "foo": "bar",
	 "bar": "baz"
}
```

## クライアントからの例
### プリファレンスのリクエスト
```
WEB GET /api/users/5/preferences:
{
        "Name": "TestPref2",
        "Value": 57005,
        "DataData": "bW9yZSBpbXBvcnRhbnQgZGF0YQ=="
}
WEB GET /api/users/1/preferences:
{
        "Name": "TestPref1",
        "Val": 1234567890,
        "Data": "some important data",
        "Things": 3.1415
}
```
### すべての設定を要求する(管理者として)
```
WEB GET /api/users/preferences:
[]
```
### プッシュ
```
WEB REQ PUT /api/users/1/preferences:
{
        "Name": "TestPref1",
        "Val": 1234567890,
        "Data": "some important data",
        "Things": 3.1415
}
```
```
WEB REQ PUT /api/users/5/preferences:
{
        "Name": "TestPref2",
        "Value": 57005,
        "DataData": "bW9yZSBpbXBvcnRhbnQgZGF0YQ=="
}
```
### 存在しないユーザへのプッシュ

404 not foundが表示されます。

### 非管理者として他の誰かをプッシュおよびプルします

403 forbidden が表示されます。

### プリファレンスの削除
```
WEB REQ DELETE /api/users/5/preferences:
```
## 他の人を削除しようとすること

403 forbiddenが表示されます。