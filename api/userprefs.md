## ユーザープリファレンス
ユーザープリファレスAPIは、ログイン後やデバイス間で保持されるGUIプリファレンスを保存するために使用されます。

/api/users/{id}/preferencesをGETすると、JSONのチャンクが返されます。 管理者はすべてのユーザーのプリファレンスをリクエストできますが、ユーザーは自分のセッションしかリクエストできません。プリファレンスが存在しない場合はnullを返します。

/api/users/preferencesをGETすると、すべてのユーザーのプリファレンスが返されます。

ユーザープリファレンスを更新するためには、/api/users/{id}/preferencesをPUTしてください。プリファレンスが存在しない場合は、提供されたJSONで更新します。このAPIではPOSTは行われません。PUTメソッドのペイロードはJSON blobです。

関連するメソッドはGETとPUTのみです。各ユーザーは、本来、ただ1つだけのプリファレンスJSON blobを持つべきです。

/api/users/{id}/preferencesをDELETEすると、プリファレンスは削除されます（管理者や自分で作成した場合）。


GETで返されたJSONの例です:
```json
{
     "foo": "bar",
	 "bar": "baz"
}
```

## クライアントからの例
### プリファレンスを要求する
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
### すべてのプリファレンスを要求する(管理者として)
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

404 not foundが得られます。

## 非管理者として他の誰かのプリファレンスをプッシュおよびプルする

403 forbiddenが得られます。

### プリファレンスの削除
```
WEB REQ DELETE /api/users/5/preferences:
```
## 非管理者として他の誰かのプリファレンスを削除しようとする

403 forbiddenが得られます。
