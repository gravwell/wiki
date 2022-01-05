# evalモジュール

[検索モジュールのドキュメント](#!search/searchmodules.md#Eval)で紹介されているように、Gravwellのevalモジュールは、他のモジュールでは不十分な場合に検索エントリを操作するための一般的なツールです。このモジュールは[Ankoスクリプト言語](scripting.md)を使用して、パイプライン内での汎用的なスクリプト機能を提供します。

evalモジュールにはいくつかの重要な制限があります:

* 1つのステートメントのみ定義することができます: `(x==y || j < 2)` は `if SrcPort==80 { setEnum("http", true) }`と同様に許容されますが、 `setEnum(foo, "1"); setEnum(bar, "2")` は2つのステートメントであり、動作しません。詳細は以下のセクションを参照してください。
* 関数の定義やインポートはできません。
* ループは使用できません
* リソースシステムへのアクセス不可

注：eval式の構造をより明確にするために、クエリの入力中にCtrl-Enterを押すと、必要に応じて改行が挿入されます。

この言語自体の詳細については、[the Anko scripting language documentation](scripting.md)で使用されているスクリプト言語の一般的な説明を参照してください。

## フィルタリング: 表現方法とステートメントの比較

evalモジュールは、引数が式の場合にはエントリーをフィルタリングしますが、引数が文の場合にはフィルタリングしません。次のような例を考えてみましょう:

```
tag=reddit json Body | eval len(Body) < 20 | table Body
```

`len(Body) < 20`は式なので、その式にマッチしたエントリーだけがパイプラインを進むことができます。これに対して、次のように考えてみましょう:

```
tag=reddit json Body | eval if len(Body) <= 10 { setEnum("postlen", "short"); setEnum(“anotherEnum”, “foo”) } | table Body
```

`if`形式はステートメントであり、ifステートメントの結果に関わらず、すべてのエントリーはパイプラインを進みます。

式とは，文字列 `“foo”`や数値 `1.5`のように，ある値に評価されるものです．これには、`DoStuff(15)`のような関数呼び出しや、`myVariable == 42`のようなブーリアン式も含まれます。基本的には、変数に代入したり、関数の引数として渡すことができるものはすべて式と考えることができます。

ステートメントは、`if`や`switch`などのステートメント、変数の作成や割り当て、リターンなど、スクリプトの流れや構造を制御します。例えば、`if`文には、式と複数の文のリストが含まれます。例えば、`if`文には、式と複数の文のリストが含まれており、式の評価値がtrueの場合、1つの文のリストが実行されます。式が評価されて真であれば、ステートメントの1つのリストが実行され、式が評価されて偽であれば、異なるリストが実行されます。

## 列挙された値

eval文の中では、既存の列挙された値を通常の変数のように参照することができます:

```
tag=reddit json Body | eval len(Body) < 20 | table Body
```

しかし、列挙型の値を設定するには、より明確な記述が必要です。`setEnum`関数は引数として名前と値を取ります。これは、与えられた名前の列挙型を作成または更新し、与えられた値に対して正しい列挙型を推論します。:

```
tag=reddit json Body | eval if len(Body) <= 10 { setEnum("postlen", "short") } else if len(Body) > 10 && len(Body) < 300 { setEnum("postlen", "medium") } else { setEnum("postlen", "long") } | count by postlen | table postlen count
```

`delEnum`関数は、必要に応じて指定された列挙型の値を削除しますが、これが必要になることはほとんどありません。

## ユーティリティー関数

Evalにはビルトインのユーティリティー関数が用意されており、以下のように`functionName(<functionArgs>) <returnValues>`という形式で表示されます:

* `setEnum(key, value)` は、任意の有効な列挙型の値を含む key という名前の列挙型の値を作成します。
* `delEnum(key)` は key という名前の列挙型の値を削除します。
* `hasEnum(key) bool` は、エントリが列挙型の値を持っているかどうかを示すブール値を返します。
* `setPersistentMap(mapname, key, value)` 検索全体で永続するキーと値のペアをマップに保存します。
* `getPersistentMap(mapname, key) value` は、指定されたキーに関連する値を、指定された永続的なマップから返します。
* `len(val) int` はvalの長さを返します。valは文字列やスライスなどです。
* `toIP(string) IP` 文字列をIPに変換し、パケットモジュールなどで生成されたIPと比較します。
* `toMAC(string) MAC` は、文字列をMACアドレスに変換します。
* `toString(val) string` は、valを文字列に変換します。
* `toInt(val) int64` 可能であれば，valを整数に変換します．変換できない場合は0を返します。
* `toFloat(val) float64` は、可能であればvalを浮動小数点数に変換します。変換できない場合は0.0を返します。
* `toBool(val) bool` は，valを真偽値に変換しようとします。変換できない場合は false を返します。0以外の数値や文字列 "y"，"yes"，"true "はtrueを返します。
* 例： `toHumanSize(15127)` は "14.77KB" に変換されます。
* 例えば、 `toHumanCount(15127)` は "15.13 K" に変換されます。
例： `toHumanCount(15127)` は "15.13 K" に変換されます。 * `typeOf(val) type` は val の型を文字列で返します。

evalが暗黙の変換を思い通りにできないこともあるので、変換関数は特に重要です。たとえば，パケットモジュールでは，IPを単なる文字列ではなく特別な型で抽出します。IPを含む列挙型の値と適切に比較するためには、`IP`関数を使わなければなりません:

```
tag=pcap packet ipv4.SrcIP | eval SrcIP != toIP("192.168.0.1") | count by SrcIP | table SrcIP count
```

想定外の結果が出た場合に備えて、`typeOf`関数を使って列挙型の値の型を確認することができます:

```
tag=pcap packet ipv4.SrcIP | eval setEnum("type", typeOf(SrcIP)) | table type
```

この場合、結果は`SrcIP`がnet.IPであることを示しています。

#### パーシステントマップ

Evalはデータをマップに格納するメソッドを提供します。つまり、あるエントリのデータを格納して、後で別のエントリと比較することができます。SetPersistentMap(mapname, key, value)`関数は，与えられた文字列 `key` と `value` を対応付けるマップ（`mapname`パラメータで指定される）のエントリを作成または更新する。その後、`GetPersistentMap(mapname, key)`関数を使ってその情報を取得することができます。

パーシステントマップの機能の例として、次のようなJson形式のHacker Newsのコメントを考えてみましょう（実際のコメントではありません）:

```
{"body": "The eval function has persistent map capabilities", "author": "Gravwell", "article-id": 1234000, "parent-id": 1234111, "article-title": "Gravwell for data analysis", "date-string": "0 minutes ago", "type": "comment", "id": 1234222}
```

このコメントには、現在のコメントのID（1234222）、著者名（"Gravwell"）、親コメントのID（1234111）が含まれていますが、親コメントの著者名は含まれていないことに注意してください。

次のコマンドは、親のIDと著者名の照合を試みます。すべてのコメントに対して、コメントIDと著者名のマッピングを "id_to_name "という名前のマップに保存します。次に、親IDがそのマップにエントリを持っているかどうかをチェックし、エントリがあれば、`parentauthor`という列挙型の値をその名前に設定し、そうでなければ列挙型の値を「unknown」に設定します。

```
tag=hackernews json author id "parent-id" as parentid | sort asc | eval if true { setPersistentMap("id_to_name", id, author);  name = getPersistentMap("id_to_name", parentid); if name != nil { setEnum("parentauthor", name) } else { setEnum("parentauthor", "unknown") } } | table author id parentid parentauthor
```

この方法の限界は、最初に処理された最も古いコメントが、検索に表示されないさらに古いコメントに返信してしまうことです。そのため、最初の検索結果はすべて親著作者が「unknown」となっており、結果をスクロールするにつれて、有効な親著作者を見つけるコメントが増えていきます。より良い解決策は、おそらく1週間に渡って検索を行い、各コメントの著者名とコメントIDを抽出して、結果をルックアップテーブルに保存することです。そして、そのルックアップテーブルをリソースとして保存しておけば、`lookup`モジュールを使って、より短い期間の検索を行うことができます。

この例では、evalモジュールの現状におけるもうひとつの限界を示しています。それは、ひとつのステートメントや式しか実行されないため、ロジックを `if true {}` ステートメントの中に収めなければならないことです。
