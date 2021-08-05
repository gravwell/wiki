## Eval

Evalは、検索および列挙値に対してANDおよびOR論理を実行するために最も一般的に使用されています。 しかし、evalモジュールはスイスアーミーナイフのようなもので、Ankoプログラミング言語（動的に型付けされたGoのような言語）の限られたサブセットへのアクセスを提供します。(Gravwell内のデータに対する柔軟な操作を可能にする[https://github.com/mattn/anko/](https://github.com/mattn/anko/) および[Anko言語用のGravwellのドキュメント](#!scripting/scripting.md))  evalモジュールはちょうど1つの式または文を実行します。のページを比較的単純にするために、このセクションではいくつかのeval呼び出し例の簡単な概要だけを説明します; [詳細](#!scripting/eval.md)

### 構文

`eval <expression>`

の [evalのドキュメント](#!scripting/eval.md)で説明されているように、単一のAnko式でなければなりません。

### 例

evalモジュールの簡単な用途は、20文字以下のRedditコメントを切り離すことです。これを行うには、jsonモジュールを使用して`Body`フィールドを抽出してから、コメントの`Body`フィールドの長さが20未満になるたびにtrueと評価される式を使用してそれをevalに渡します。式がtrueと評価されるエントリだけを続行できます。最後に、単純にevalの結果をテーブルモジュールに送信して、20文字未満のコメント本文を表示します。

```
tag=reddit json Body | eval len(Body) < 20 | table Body
```

ANDおよびORロジックは、上記の長さの例と同じ方法で実行できます。たとえば、送信元ポートと宛先ポートがあり、クエリをフィルタ処理するためにそれぞれの範囲を確認することに関心がある場合、構文は次のようになります:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP tcp.SrcPort tcp.DstPort | eval ( (DstPort < 5000 && DstPort > 2000) || (SrcPort > 9000 && SrcPort > 8000) ) | table SrcIP SrcPort DstIP DstPort
```

似たような線に沿ったより複雑な例では、コメントの長さが異なる場合の相対頻度を調べます。postlenコメントが10文字以下の場合は「short」、10〜300文字の場合は「medium」、長い場合は「long」に列挙値を設定します。次に、countモジュールを使用して各長さを集計し、tableモジュールを使用して各長さの数を表示します。

```
tag=reddit json Body | eval if len(Body) <= 10 { setEnum("postlen", "short") } else if len(Body) > 10 && len(Body) < 300 { setEnum("postlen", "medium") } else { setEnum("postlen", "long") } | count by postlen | table postlen count
```

`setEnum()`関数の使用に注意してください。それはあなたが下流に渡される列挙値を設定することを可能にします。列挙値は文字列である必要はありません;それはブール値、数値、スライス、あるいはIPでも構いません。

評価モジュールがサポートする`if`と`switch`以下に示すような論理文:

```
if len(Body) <= 10 { setEnum("postlen", "short"); setEnum(“anotherEnum”, “foo”) }
```

```
switch DstPort { case 80: setEnum(“protocol”, “http”); case 22: setEnum(“protocol”, “ssh”); default: setEnum(“protocol”, “unknown”) }
```

### 更に詳しい説明

* [Anko言語のGravwellドキュメント](#!scripting/scripting.md)は、Ankoスクリプト言語の一般的な説明です
* [evalモジュール](#!scripting/eval.md)では、evalモジュールについて詳しく説明しています。
