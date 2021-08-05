# 名前付きフィールド

namedfields モジュールは、検索エントリからデータを抽出し、後で使用するために列挙値にフィルタリングするために使用されます。[fields モジュール](#!search/fields/fields.md) のように、バイト列で区切られたレコードからフィールドを抽出します。しかし、fields モジュールがインデックスを使って要素を参照するのに対して（例えば、`fields -d "\t" [5] as foo` は、「6番目の要素を抽出して、それを foo と呼ぶ」という意味です）、namedfields は特別なフォーマットの [resources](#!resources/resources.md) を使って、特定のデータフォーマットに人間が親しみやすい名前を付けます。これは、[Bro](https://www.bro.org) ログや CSV ファイルのようなものを解析しようとするときに特に有効です。

多くの人が Bro を使用しているので、Bro のフィールドをデコードするためのリソースファイルを [https://github.com/gravwell/resources](https://github.com/gravwell/resources) の `bro/namedfields` サブディレクトリに用意しました。Gravwell のリソースとして `namedfields.json` をアップロードするだけです。このドキュメントの例では、`brofields` という名前のリソースにアップロードされたと仮定しています。

## キーコンセプト

namedfields モジュールは、その中核となるもので、行内の数値インデックスをユーザーフレンドリーな名前にマッピングします。インデックスから名前へのマッピングのセットをグループと呼びます。ネットワークフローのテキスト表現を解析するグループは、次のようなマッピングを持っています。

| インデックス | 名称 |
|-------|------|
| 0 | start_time |
| 1 | duration |
| 2 | protocol |
| 3 | src_ip |
| 4 | src_port |
| 5 | dst_ip |
| 6 | dst_port |
| 7 | packets |
| 8 | bytes |

その後、1 つまたは複数のグループが、このドキュメントの別の場所で指定されたフォーマットで Gravwell リソースに集められます。namedfields が実行されるとき、ユーザーはどのリソースをロードするか、そしてそのリソース内のどのグループを使ってユーザーが指定した名前をインデックスにマッピングするかを指定します。

## サポートされているオプション

* `-r <arg>`: 「-r」オプションは必須です。インデックスと名前のマッピングを含むリソースの名前または GUID を指定します。
* `-g <arg>`: 「-g」オプションは必須で、指定したリソース内で使用するグループを指定します。
* `-e <arg>`: 「-e」オプションは、レコード全体ではなく、列挙値に対して操作を行います。
* `-s` : 「-s」オプションは、namedfields モジュールがストリクトモードで動作することを指定します。 ファイルされた仕様が満たされない場合、そのエントリーは削除されます。 例えば、0 番目、1 番目、2 番目のフィールドが必要なのに、エントリーには 2 つのフィールドしかない場合、strict フラグを指定すると、そのエントリーはドロップされます。

## フィルタリング演算子

namedfields モジュールでは、等値性に基づくフィルタリングが可能です。 等価性を指定するフィルタが有効な場合（「等しい」、「等しくない」、「含まれる」、「含まれない」）、フィルタの指定に失敗したエントリは完全に削除されます。 Not equal "!=" と指定されたフィールドが存在しない場合、そのフィールドは抽出されませんが、エントリが完全に削除されることはありません。

| 演算子 | 名称 | 意味 |
|----------|------|-------------|
| == | 等しい | フィールドは等しい
| != | 等しくない | フィールドは等しくない
| ~ | 含む | フィールドはその値を含む
| !~ | 含まない | フィールドはその値を含まない

## 例

Bro の conn.log ファイルが broconn フラグで取り込まれたと仮定して、以下は各エントリから service、dst、resp_bytes フィールドを抽出します。service フィールドが dns という文字列と一致しないエントリはすべて削除され、抽出された dst フィールドの名前は server に変更されます。そして、各サーバーの DNS レスポンスの平均長さを計算し、グラフ化します。brofields というリソースと、Conn というグループを指定していますが、これらは brofields リソース内で定義されていることに注意してください。

```
tag=broconn namedfields -r brofields -g Conn service==dns dst as server resp_bytes  | mean resp_bytes by server | chart mean by server
```

次の例では、intel.log という異なる Bro ファイルを解析しています。リソースは同じですが、異なるグループを指定していることに注意してください。

```
tag=brointel namedfields -r brofields -g Intel source | count source | table source count
```

## 名前付きフィールドのリソースフォーマット

namedfields モジュールを使用する前に、名前をフィールド内のインデックスにマッピングするためのリソースを作成する必要があります。リソースは JSON で構造化されています。各リソースには複数のグループを含めることができ、モジュールの実行時にそのうちの 1 つが選択されます。

 以下の例では、Bro の `intel.log` ファイルと、CSV 形式のカスタムアプリケーションログのエントリに名前を付けています。

```
{
	"Version": 2,
	"Set": [
		{
			"Delim": "\t",
			"Name": "Intel",
			"Subs": [
				{
					"Name": "source",
					"Index": 0
				},
				{
					"Name": "desc",
					"Index": 1
				},
				{
					"Name": "url",
					"Index": 2
				}
			]
		},{
			"Name": "App",
			"Engine": "csv",
			"Subs": [
				{
					"Name": "user",
					"Index": 0
				},
				{
					"Name": "host",
					"Index": 1
				},
				{
					"Name": "GUID",
					"Index": 2
				}
			]
		}

	]
}
```

必須のコンポーネントに注意してください。

* `Version` では、このファイルがどのバージョンの namedfields モジュールに対応しているかを指定します。1 のままにしておいてください。
* `Set` は、グループの配列を含みます。
* このファイルのセットには、「Intel」という名前の 1 つのグループが含まれています。区切り文字はタブ文字（\t）で、`Subs` のリストが用意されています。
* Subs メンバーは、このグループ内のサブフィールドを定義します。インデックス 0 のフィールドは source という名前で、インデックス 1 は desc、インデックス 2 は url という名前であることがわかります。
* Engine メンバーは、グループにどのエンジンを使用するかを宣言します（フィールド、csv など...）。

Gravwell が配布している Bro Log 用の [namedfields.json](https://github.com/gravwell/resources/blob/master/bro/namedfields/namedfields.json) ファイルには、多くのグループが含まれているので、その例を参照してください。

### namedfields のリソース生成

Gravwell は、namedfields リソースの生成を支援するシンプルな golang ライブラリを提供しています。 このライブラリを使って、プログラムでリソースを生成し、namedfields モジュールで使用することができます。 このライブラリは、github の tools Gravwell リポジトリの nfgen ディレクトリにあります。

名前付きフィールドを使って、1 つのリソースに 2 つのグループを生成する最も簡単な方法は、次のようになります。

```
package main

import (
	"github.com/gravwell/tools/nfgen"
	"log"
)

func main() {
	//create a new named fields resource using the CSV engine that knows how to deal with 2
	//data types, one for login events and one for password failed events
	nf := nfgen.NewGen()
	g, err := nfgen.NewGroup("logins", "csv", ``)
	if err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`username`, ``, 1); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`host`, ``, 2); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`srcip`, ``, 3); err != nil {
		log.Fatal(err)
	}
	if err = nf.AddGroup(g); err != nil {
		log.Fatal(err)
	}
	if g, err = nfgen.NewGroup("failedlogins", "csv", ``); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`srcip`, ``, 2); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`username`, ``, 3); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`password`, ``, 4); err != nil {
		log.Fatal(err)
	}
	if err = g.AddSub(`host`, ``, 5); err != nil {
		log.Fatal(err)
	}
	if err = nf.AddGroup(g); err != nil {
		log.Fatal(err)
	}
	if err = nf.Export("/tmp/lookups.json"); err != nil {
		log.Fatal(err)
	}
}
```

golang のビルドチェーンがインストールされていれば、前述のジェネレーターをビルドして実行するのは簡単です。

```
go get -u github.com/gravwell/tools/nfgen
go build main.go -o test
./test
```
