## Eval

evalは、検索や列挙値に対してANDやORの論理を実行するために最もよく使用されます。しかし、evalモジュールはちょっとしたスイスアーミーナイフのようなもので、Ankoプログラミング言語（動的に型付けされた囲碁のような言語で、[https://github.com/mattn/anko/](https://github.com/mattn/anko/)や[Anko言語用のSolitonNKのドキュメント](#!scripting/scripting.md)を参照）の限られたサブセットへのアクセスを提供し、SolitonNK内のデータに対する柔軟な操作を可能にする。evalモジュールは正確に1つの式または文を実行する。このページを比較的シンプルなものにするために、このセクションではevalの実行例について簡単に説明します。詳細は[こちら](#!scripting/eval.md)を参照してください。

### 構文

`eval <expression>`

<expression>は、[evalのドキュメント](#!scripting/eval.md)で説明されているように、単一のAnko式でなければなりません。

### 例

evalモジュールの簡単な応用例としては、長さが20文字未満のRedditコメントを分離することが挙げられます。そのためには、jsonモジュールを使って`Body`フィールドを抽出し、コメントの`Body`フィールドの長さが20未満のときにtrueと評価される式をevalに渡します。この式がtrueと評価されたエントリーだけがパイプラインを進むことができます。最後に、evalの結果をtableモジュールに送って、20文字未満のコメント本文を表示します。

```
tag=reddit json Body | eval len(Body) < 20 | table Body
```

ANDおよびORロジックは、上記の長さの例と同様の方法で行うことができます。例えば、ソースポートとデスティネーションポートがあり、クエリをフィルタリングするためにそれぞれのポートの範囲を確認したい場合、以下のような構文になります:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort | eval ( (DstPort < 5000 && DstPort > 2000) || (SrcPort > 9000 && SrcPort > 8000) ) | table SrcIP SrcPort DstIP DstPort
```

もっと複雑な例では、コメントの長さの違いによる相対的な頻度を調べています。この例では、列挙値 `postlen` を、コメントが10文字以下の場合は"short"、10～300文字の場合は"medium"、それ以上の場合は"long"を設定しています。そして、countモジュールを使って各長さを集計し、tableモジュールを使って各長さのカウントを表示します。

```
tag=reddit json Body | eval if len(Body) <= 10 { setEnum("postlen", "short") } else if len(Body) > 10 && len(Body) < 300 { setEnum("postlen", "medium") } else { setEnum("postlen", "long") } | count by postlen | table postlen count
```

`setEnum()`関数の使用に注意してください。これにより、下流に渡される列挙値を設定することができます。列挙値は、文字列である必要はなく、ブーリアン、数値、スライス、あるいはIPであっても構いません。

evalモジュールは、以下に示すように、`if`と`switch`の論理文をサポートしています:

```
if len(Body) <= 10 { setEnum("postlen", "short"); setEnum(“anotherEnum”, “foo”) }
```

```
switch DstPort { case 80: setEnum(“protocol”, “http”); case 22: setEnum(“protocol”, “ssh”); default: setEnum(“protocol”, “unknown”) }
```

### 参考

* [Anko言語用のSolitonNKのドキュメント](#!scripting/scripting.md) は、Ankoスクリプト言語の説明です。
* [evalモジュール](#!scripting/eval.md) では、evalモジュールについて詳しく説明しています。
