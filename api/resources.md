# リソース Web API

Web APIは、リソースへのアクセスを提供します。曖昧さを避けるために、リソースはGUIDで参照する必要があります。

## リソースのメタデータ構造
リソースシステムは、各リソースのメタデータ構造を保持します。 Web APIは、この構造体のJSONエンコードバージョンを使用して通信します。 フィールドは主に自明ですが、正確さのためにここで説明します。

* UID：リソースのオーナーのUID
* GUID：リソースのユニークな識別子
* LastModified：リソースが最後に更新された時刻
* VersionNumber：リソースの内容が変更されるたびに増加します。
* GroupACL：リソースへのアクセスが許可されている整数のグループIDのリスト
* Global：trueの場合、リソースはシステム上のすべてのユーザーが読めるようになります。グローバルリソースを作成できるのは管理者のみです。
* ResourceName：リソースの名前
* Description：リソースの詳細な説明です。
* Size：リソースのコンテンツのサイズをバイト単位で指定します。
* Hash：リソースの内容を表すSha1ハッシュです。
* Synced：(内部使用のみ)

## リソースの一覧表示

すべてのリソースのリストを取得するには、`/api/resources` を GET してください。結果は以下のようになります。

```
[{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:19:10.945117816-07:00","VersionNumber":0,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":0,"Hash":"","Synced":true},{"UID":1,"GUID":"66f7be7d-893b-4dc4-b0ad-3609b348385d","LastModified":"2018-02-12T11:06:44.215431364-07:00","VersionNumber":1,"GroupACL":[1],"Global":false,"ResourceName":"test","Description":"test resource","Size":543,"Hash":"zkTmUEV+AR6JZdqhobIeYw==","Synced":true}]
```

この例は、「newresource」（GUID 2332866c-9b8d-469f-bf40-de9fad828362）と「test」（GUID 66f7be7d-893b-4dc4-b0ad-3609b348385d）の2つのリソースを示しています。

## リソースの作成

リソースを作成するには、`/api/resources`にPOSTリクエストを行い、以下のフォーマットのJSON構造を送信します。

```
{
	"GroupACL": [3,7],
	"Global": false,
	"ResourceName": "newresource",
	"Description": "Description of the resource"
}
```

注：この構造体は、メタデータ構造体のサブセットであり、ユーザーが設定できるフィールドを含んでいます。

サーバは、新たに作成されたリソースのリソース・メタデータ構造体を応答します:

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:19:10.945117816-07:00","VersionNumber":0,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":0,"Hash":"","Synced":false}
```

## リソースコンテンツの設定

新しく作成されたリソースにはデータが含まれていません。リソースの内容を変更するには、`/api/resources/{guid}/raw`に対してマルチパートのPUTリクエストを発行し、`{guid}`をリソースの適切なGUIDに置き換えます。このリクエストに必要なのは、リソースに格納されるべきデータを含む、`file`という名前の1つのパートだけです。したがって、上記で作成したリソースのコンテンツを設定するには、`/api/resources/2332866c-9b8d-469f-bf40-de9fad828362/raw`に対してマルチパートのPUTを実行します。サーバは、新しい更新時刻、サイズ、およびハッシュを示す更新されたメタストラクチャで応答します。以下にcurl呼び出しの例を示します。「maxmind.db」という名前のファイルをリソースにアップロードしています（ベアラートークンはユーザーセッションに合わせて適切に設定する必要があることに注意してください。

```
curl 'http://gravwell.example.com/api/resources/2332866c-9b8d-469f-bf40-de9fad828362/raw' -X PUT -H 'Authorization: Bearer 7b22616c676f223a35323733382c22747970223a226a3774227d.7b22756964223a312c2265787069726573223a22323031392d31302d30395431333a33333a32352e3231343632203131352d30363a3030222c22696174223a5b33392c32323c2c35382c36362c3231372c32362c3131392c33362c3234312c33352c39302c312c39312c3138312c3234322c33362c3137342c3139342c3130382c37342c3133382c32362c3133392c3234362c37362c3132352c3136342c38382c39322c39302c3231312c36365d7d.ef9ca1e0ac7f012adcd796d8cca0746a6fabecd7e787c025d754e54a072be5c89dc7bac5f648ae26b422f0bbe6b69a806e8de4a0fe2b7d06d3293ed4c1323daf' -H 'Content-Type: multipart/form-data' -H 'Accept: */*' --form file=@maxmind.db
```

## リソースコンテンツの読み込み

リソースのコンテンツを読み取るには、`/api/resources/{guid}/raw`に対してGETリクエストを行います。`{guid}`はリソースの適切なGUIDに置き換えてください。

## リソースのメタデータの読み込みと更新

1つのリソースのメタデータを読み込むには、`/api/resources/{guid}`に対してGETリクエストを行います。例えば、`/api/resources/2332866c-9b8d-469f-bf40-de9fad828362`をGETすると、以下のようになります。

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:29:10.557490321-07:00","VersionNumber":1,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"Description of the resource","Size":6,"Hash":"QInZ92Blt3TopFBeBTD0Cw==","Synced":true}
```

メタデータを修正するには、`/api/resources/{guid}`に対してPUTリクエストを実行します。リクエストの内容は、GETリクエストで読み込んだ構造体に、必要なフィールドを変更したものである必要があります。例えば、"newresource"の記述を変更するには、次のような内容のPUTを実行します。

```
{"UID":1,"GUID":"2332866c-9b8d-469f-bf40-de9fad828362","LastModified":"2018-03-07T15:29:10.557490321-07:00","VersionNumber":1,"GroupACL":[3,7],"Global":false,"ResourceName":"newresource","Description":"A new description for the resource!","Size":6,"Hash":"QInZ92Blt3TopFBeBTD0Cw==","Synced":true}
```

注：通常のユーザーが変更できるのは、GroupACL、ResourceName、およびDescriptionの各フィールドのみです。管理者ユーザーは、Globalフィールドを変更することもできます。それ以外のフィールドを変更した場合、サーバーはそれを無視します。

## リソースのコンテントタイプの取得

`api/resources/{guid}/contenttype` へのGETリクエストは、検出されたリソースのコンテントタイプと、リソースボディの最初の512バイトを含む構造体を返します:

```
{"ContentType":"text/plain; charset=utf-8","Body":"IyBEdW1wcyB0aGUgcm93cyBvZiB0aGUgc3BlY2lmaWVkIENTViByZXNvdXJjZSBhcyBlbnRyaWVzLCB3aXRoCiMgZW51bWVyYXRlZCB2YWx1ZXMgY29udGFpbmluZyB0aGUgY29sdW1ucy4KIyBlLmcuIGdpdmVuIGEgcmVzb3VyY2UgbmFtZWQgImZvbyIgY29udGFpbmluZyB0aGUgZm9sbG93aW5nOgojCWhvc3RuYW1lLGRlcHQKIwl3czEsc2FsZXMKIwl3czIsbWFya2V0aW5nCiMJbWFpbHNlcnZlcjEsSVQKIyBydW5uaW5nIHRoZSBmb2xsb3dpbmcgcXVlcnk6CiMJdGFnPWRlZmF1bHQgYW5rbyBkdW1wIGZvbwojIHdpbGwgb3V0cHV0IDQgZW50cmllcyB3aXRoIHRoZSB0YWcgImRlZmF1bHQiLCBjb250YWluaW5nIGVudW1lcmF0ZWQKIyB2YWx1ZXMgbmFtZWQgImhvc3RuYW1lIiBhbmQgImRlcHQiIHdob3NlIGNvbnRlbnRzIG1hdGNoIHRoZSByb3dzCiMgb2YgdGhlIHJlc291cmNlLgojIEZsYWdzOgojICAtZDogc3BlY2lmaWVzIHRoYXQgaW5jb21pbmcgZW50cmllcyBzaG91bGQgYmUgZHJvcHBlZCA="}
```

bytesパラメータを追加すると、例えば、`/api/resources/{guid}/contenttype?bytes=1024`のように、返されるバイト数が変わります。なお、コンテントタイプの検出に成功した場合、APIは常に最低128バイトを読み込みます。128バイト未満が指定された場合、APIはデフォルトで512バイトを読み込みます。

## リソースの削除

リソースを削除するには、`/api/resources/{guid}`に対してDELETEリクエストを発行するだけです。ただし、`{guid}`は通常通りリソースの適切なGUIDに置き換えてください。

## リソースのクローン作成

既存のリソースをクローンするには、`/api/resources/{guid}/clone`にPOSTリクエストを発行し、`{guid}`を元のリソースのGUIDに置き換えます。リクエストのボディには、新しく作成されるクローンの名前を含むJSON構造を指定します。

```
{
	"Name": "Copy of Foo"
}
```

サーバーは、新たにクローン化されたリソースのメタデータを応答します:

```
{
  "UID": 1,
  "GUID": "bbf682f8-363f-4245-8182-d7f6286022ff",
  "Domain": 0,
  "LastModified": "2020-08-19T20:31:11.258257937Z",
  "VersionNumber": 1,
  "GroupACL": null,
  "Global": false,
  "ResourceName": "Copy of Foo",
  "Description": "foobar",
  "Size": 207,
  "Hash": "xfV5dQG4eRe75ULzdb2e2A==",
  "Synced": false,
  "Labels": [
    "blah"
  ]
}
```

## 管理者権限

管理者ユーザーは、システム上のすべてのリソースを見る必要がある場合があります。管理者ユーザーは、`/api/resources?admin=true`のGETリクエストで、システム内のすべてのリソースのグローバルリストを取得することができます。

リソースのGUIDはシステム全体で一意であるため、管理者は`?admin=true`を指定することなく、リソースの修正、削除、取得を行うことができますが、不必要にパラメータを追加してもエラーにはなりません。