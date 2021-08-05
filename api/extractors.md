# 自動抽出API

エクストラクターWebAPIは、自動エクストラクター定義にアクセス、変更、追加、および削除するためのメソッドを提供します。 自動抽出機能とその構成の詳細については、[自動抽出機能](/#!configuration/autoextractors.md) セクションを参照してください。

## データ構造

自動抽出機能には、次のフィールドが含まれています。

* Tag：抽出されているタグ。
* Name：エクストラクターのわかりやすい名前。
* Desc：抽出機能のより詳細な説明。
* Module：使用するモジュール（「csv」、「fields」、「regex」、または「slice」）。
* Params：抽出モジュールのパラメータ。
* Args：抽出モジュールの引数。
* Labels：[labels](#!gui/labels/labels.md)を含む文字列の配列。
* UID：ダッシュボードの所有者の数値ID。
* GID：このダッシュボードが共有される数値グループIDの配列。
* Global：ブール値で、ダッシュボードをすべてのユーザー（管理者のみ）に表示する場合はtrueに設定されます。
* UUID：この特定のエクストラクターの一意のID。
* LastUpdated：このエクストラクタが最後に変更された時刻。

これは、構造に関するTypescriptの記述です:
```
type RawAutoExtractorModule = 'csv' | 'fields' | 'regex' | 'slice';

interface RawAutoExtractor {
	UUID: RawUUID;

	UID: RawNumericID;
	GIDs: Array<RawNumericID> | null;

	Name: string;
	Desc: string;
	Labels: Array<string> | null;

	Global: boolean;
	LastUpdated: string; // Timestamp

	Tag: string;
	Module: RawAutoExtractorModule;
	Params: string;
	Args?: string;
}
```

## リスト

`/api/autoextractors`でGETを発行すると、現在のユーザーがアクセスできるエクストラクターのセットを表すJSON構造のリストが返されます。 応答の例は次のとおりです:

```
[
  {
    "Name": "Apache Combined Access Log",
    "Desc": "Apache Combined access logs using regex module.",
    "Module": "regex",
    "Params": "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<auth>\\S+) \\[(?P<date>[^\\]]+)\\] \\\"(?P<method>\\S+) (?P<url>.+) HTTP\\/(?P<version>\\S+)\\\" (?P<response>\\d+) (?P<bytes>\\d+) \\\"(?P<referrer>\\S+)\\\" \\\"(?P<useragent>.+)\\\"",
    "Tag": "apache",
    "Labels": [
      "apache"
    ],
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "UUID": "0e105901-92a7-4131-87bb-a00287d46f96",
    "LastUpdated": "2020-06-24T13:49:39.013266326-06:00"
  },
  {
    "Name": "vpcflow",
    "Desc": "VPC flow logs (TSV format)",
    "Module": "fields",
    "Params": "version, account_id, interface_id, srcaddr, dstaddr, srcport, dstport, protocol, packets, bytes, start, end, action, log_status",
    "Args": "-d \" \"",
    "Tag": "vpcflowraw",
    "Labels": null,
    "UID": 1,
    "GIDs": null,
    "Global": false,
    "UUID": "7f80df6a-a2ce-42aa-b531-ac11c596f64a",
    "LastUpdated": "2020-05-29T15:00:41.883390284-06:00"
  }
]
```

adminフラグを設定してGETリクエストを実行すると(`/api/autoextractors?admin=true`)、システム上の*すべての*エクストラクターのリストが返されます。

## タグのエクストラクタを見つけます

システムが特定のタグに使用するエクストラクタを確認するには、`/api/autoextractors/find/{tag}`でGETリクエストを発行します。ここで、"{tag}"は問題のタグに置き換えられます。 たとえば、タグ"syslog"を確認するには、`/api/autoextractors/find/syslog`でGETを発行します。 サーバーは、そのタグの単一の自動抽出定義、または一致する定義が存在しない場合は404で応答します。

## 追加

自動抽出機能の追加は、リクエスト本文に有効な定義を指定して`/api/autoextractors` にPOSTを発行することで実行されます。 構造は有効である必要があり、ユーザーは同じタグに対して既存の自動抽出機能を定義することはできません。 新しい自動抽出機能を追加するPOSTJSON構造の例：

```
{
  "Tag": "foo",
  "Name": "my extractor",
  "Desc": "an extractor using the fields module",
  "Module": "fields",
  "Params": "version, account_id, interface_id, srcaddr, dstaddr, srcport, dstport, protocol, packets, bytes, start, end, action, log_status",
  "Args": "-d \" \"",
  "Labels": [
    "foo"
  ],
  "Global": false
}
```

自動抽出器の追加時にエラーが発生した場合、ウェブサーバはエラーのリストを返します。成功すれば、サーバは新しい抽出器のUUIDを応答します。

注：抽出器を作成する際には、`UUID`, `UID`, `GIDs`, `LastUpdated` の各フィールドを設定する必要はありません。管理者のみが `Global` フラグを true に設定できます。

## 更新

自動抽出器の更新は、`/api/autoextractors`にPUTリクエストを発行して、リクエストボディに有効な定義のJSON構造を入れることで実行されます。 構造体は有効でなければならず、同じUUIDを持つ既存の自動抽出器が存在しなければなりません。 変更されていないすべてのフィールドは、サーバーから最初に返されたとおりに含まれていなければなりません。 定義が無効な場合は、ボディにエラーメッセージを含むnon-200レスポンスが返されます。 構造は有効だが、更新された定義の配布にエラーが発生した場合は、エラーのリストが本文に返される。

## 抽出器の構文のテスト

自動抽出器を追加したり更新したりする前に、その構文を検証しておくと便利です。`/api/autoextractors/test` に POST リクエストを行うと、リクエストが検証されます。定義に問題がある場合は、エラーが返されます。

```
{"Error":"asdf is not a supported engine"}
```

新しい自動抽出器を追加する際には、新しい抽出器が同じタグの既存の抽出器と衝突しないことが重要です。既存の抽出物を更新する際には、これは問題になりません。指定されたタグに対して既に抽出物が存在する場合、テスト API は返された構造体の`TagExists` フィールドを設定します。

```
{"TagExists":true,"Error":""}
```

`TagExists`がtrueの場合、新しい抽出器を作成しようとする場合はエラーとして扱い、既存の抽出器を更新する場合は無視する必要があります。

## ファイルのアップロード

自動抽出器の定義は、TOML形式で表現することができます。このフォーマットは人間が読むことができ、抽出器の定義を配布するのに便利な方法です。その例を以下に示します。

```
[[extraction]]
	tag="bro-conn"
	name="bro-conn"
	desc="Bro conn logs"
	module="fields"
	args='-d "\t"'
	params="ts, uid, orig, orig_port, resp, resp_port, proto, service, duration, orig_bytes, dest_bytes, conn_state, local_orig, local_resp, missed_bytes, history, orig_pkts, orig_ip_pkts, resp_pkts, resp_ip_bytes, tunnel_parents"
```

このファイルを解析して JSON 構造体を作成するのではなく、このタイプの定義は、`/api/autoextractors/upload`への POST リクエストで送信されるマルチパートフォームを介して、ウェブサーバに直接アップロードすることができます。このフォームには、抽出器定義のコンテンツを格納する `extraction`という名前のファイルフィールドを含める必要があります。定義が有効であり，インストールに成功した場合には，サーバは 200 レスポンスを返します。

## ファイルのダウンロード

自動抽出器の定義をTOML形式でダウンロードするには`/api/autoextractors/download`にGETリクエストを発行します。UUIDが`ad782c81-7a60-4d5f-acbf-83f70e68ecb0`と`c7389f9b-ba52-4cbe-b883-621d577c6bcc`の2つのエクストラクタをダウンロードしたい場合は、`/api/autoextractors/download?`にGETリクエストを送ります。 `id=ad782c81-7a60-4d5f-acbf-83f70e68ecb0&id=c7389f9b-ba52-4cbe-b883-621d577c6bcc`.

現在のユーザが指定されたすべての抽出器にアクセスできる場合、サーバはTOML形式の定義を含むダウンロード可能なファイルで応答します。このファイルは、上述のファイルアップロードAPIを使って他のGravwellシステムにアップロードすることができます。

## 削除

既存の自動抽出器を削除するには、`/api/autoextractors/{uuid}` に対して DELETE リクエストを発行します。ここで `uuid` は自動抽出器に関連付けられた UUID です。自動抽出器が存在しない場合や、削除にエラーがあった場合、ウェブサーバーは non-200 レスポンスを返し、レスポンスボディにエラーを記載します。

## モジュールのリストアップ

自動抽出器の定義には、有効なモジュールを指定する必要があります。 サポートされているモジュールのリストを取得するAPIは、`/api/autoextractors/engines`にGETリクエストを発行することで実行されます。 結果として、文字列のリストが得られます。

```
[
	"fields",
	"csv",
	"slice",
	"regex"
]
```
