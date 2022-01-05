# スクリプトについて

スクリプトはGravwell内で2つの方法で使用されます：検索パイプラインの一部として、および検索の起動を自動化する方法として。 スクリプト言語（[Anko]（https://github.com/mattn/anko））は両方のケースで同じですが、ユースケースの違いを考慮して若干の違いがあります。 この記事では、両方のユースケースを紹介し、Anko言語の概要を提供します。

* [`anko` モジュールのドキュメント](anko.md)
* [`eval` モジュールのドキュメント](eval.md)
* [検索数リプとの資料](scriptingsearch.md)（CLIクライアントを使用したスクリプト作成およびスケジュール検索で使用可能な機能の詳細な説明が含まれています）
* [検索スケジュールの資料](scheduledsearch.md)

## スクリプトモジュール

Gravwellにはankoとeval2つのモジュールがあり、これらは[Ankoスクリプト言語](https://github.com/mattn/anko) を使用して、検索パイプラインでチューリング完全なスクリプト機能を提供します。ankoモジュールは、完全な[Turing-Complete](https://en.wikipedia.org/wiki/Turing_completeness)とランタイムを提供するためにankoの全機能セットを有効にします。Evalはankoランタイムを使用して、検索クエリに直接入力された単一のステートメントを実行します。

ankoは何でもできますが、evalにはいくつかの重要な制限があります。

* 1つのステートメントしか定義できません。(x==y || j < 2)これsetEnum(foo, "1"); setEnum(bar, "2")は受け入れ可能ですが、2つのステートメントであり、機能しません。
* 関数を定義またはインポートすることはできません.
* ループは許可されていません.
* リソースシステムにアクセスできない.

この文書はAnkoプログラミング言語そのものについて説明しています。2つの検索モジュールに関する文書は別々のページで管理されています。

* [ankoドキュメンテーション](anko.md) （ankoは、検索モジュールのドキュメンテーションでも簡単に説明されています）
* [evalドキュメンテーション](eval.md) （evalは、検索モジュールのドキュメンテーションにも簡単に説明されています）

## 検索スクリプト

`anko`と` eval`モジュールが検索パイプラインの*内部*でスクリプトを実行する場合、Gravwellは独自の検索を*起動*し、結果を操作するスクリプトもサポートします。 これは、自動クエリに役立ちます。 特定の不審なネットワーク動作を探すために毎朝午前6時に実行されるスクリプト。

これらのスクリプトは、スケジュールに従って実行するか（[スケジュール検索](#!scripting/scheduledsearch.md)を参照）、[コマンドラインクライアント](#!cli/cli.md)を使用して手動で実行できます。 スクリプト言語はどちらの場合も同じですが、スクリプトはスケジュール検索として実行され、出力を表示するために `print`関数を使用できません。

[スクリプト検索](scriptingsearch.md) のドキュメントには、例を含むこのタイプのスクリプトの記述方法に関する詳細が記載されています。

## Anko概要

簡単に言うと、Ankoは動的に型付けされたスクリプト言語で、構文はGoに似ていますが、コンパイルされるのではなく解釈されます。Ankoのgithubページは、Ankoがどのように見えて動作するかのアイデアを与えるためにこの例を提供します：

```
# declare function
func plus(n){
  return n + 1
}

# declare variables
x = 1
y = x + 1

# print values
println(x * (y + 2 * x + plus(x) / 2))

# if/else condition
if plus(y) > 1 {
  println("こんにちわ世界")
} else {
  println("Hello, World")
}

# array type
a = [1,2,3]
println(a[2])
println(len(a))

# map type
m = {"foo": "bar", "far": "boo"}
m.foo = "baz"
for k in keys(m) {
  println(m[k])
}
```

[Anko Playground](http://play-anko.appspot.com/)はAnkoコードを実験するための便利な方法です。Ankoを感じるために、上記の例やドキュメント内の他の例を試してみることをお勧めします。（setEnumなどのGravwell固有の関数を使用した例は、もちろんPlaygroundでは機能しません。）

## データ型

Goとは異なり、Ankoは明示的な型宣言を使用することはめったにありません。型は通常推測され、可能な限り自動的に変換されます。

Ankoは以下の基本型をサポートしています。

* 整数：int、int32、int64、uint、uint32、uint64
* 浮動小数点：float32、float64
* ブール値：bool
* ストリング：string
* 文字：byte、rune
* インタフェース

`interface` タイプは特別です。これは一般的なオブジェクトを表すので、 `interface` 型の配列は文字列、整数、浮動小数点などを保持できます。

Ankoは2種類の複合型を提供します。

* 配列: データの多次元配列
* マップ: 

スカラを宣言するには、変数の型を指定する必要はありません（または不可能です）。単に変数を代入すると、型が推測されます:

```
a = 1		# integer
b = 2.5		# float
c = true	# bool
d = "hi"	# string
```

暗黙の型変換は可能な限り行われ、正確性を保つために適切な型が選択されます。

```
x = a + b	# x == 1 + 2.5 == "3.5"

y = b - c	# boolean true converts to 1, so y == 2.5 - true == 2.5 - 1.0 == 1.5
```

さまざまな型に対するさまざまな操作の正確な意味は、後のセクションで説明します。

<span style="color: red; ">重要：Ankoで16進定数を使用する場合は、必ずAFを大文字にしてください。`0x1E`は数値14の有効な表現ですが、そうで`0x1e` はありません。</span>

### 配列

Ankoの配列は一般的なもの（さまざまな型の要素を保持する）または型指定されたものです。総称配列を作成するには、単純に初期値を代入します。

```
myArray = ["hi", 1]
printf("%v\n", myArray)    # prints ["hi" 1]
myArray[0] = 3.5
printf("%v\n", myArray)    # prints [3.5 1]
```

型付き配列を作成するには、`make` 関数を使用してください。`make` 引数として配列型と初期長を取ります。配列タイプは、前のセクションにリストされているスカラータイプのいずれでも構いません。

```
myArray = make([]int, 5)	# make an array of 5 ints
myArray[1] = 7
printf("%v\n", myArray)		# prints [0 7 0 0 0]
```

注意：ジェネリック配列はmake関数を使っても構築できます。make([]interface, 10)

多次元配列が可能です。次の例は、同じ多次元配列を実現するための2つの異なる方法を示しています。

```
a = [[1, "foo"][3.2, 4]]

b = make([][]interface, 2)
b[0] = [1, "foo"]
b[1] = make([]interface, 2)
b[1][0] = 3.2
b[1][1] = 4
# at this point, a and b are equivalent
```

#### 配列に追加する

`+=` 演算子を使用して配列を追加することができます。
```
a = [1, 2]
a += 3	# a is now [1 2 3]
```

ある配列を別の配列に追加することができます。

```
a = [1, 2]
b = [3, 4]
a += b	# a is now [1 2 3 4]
```

Ankoは暗黙の追加も許可します。指定されたインデックスが配列の現在の末尾のインデックスよりちょうど1大きい場合は、配列が展開されて項目が追加されます。

```
foo = ["a", "b"]
foo[2] = "c"	# foo is now ["a" "b" "c"]

# This will fail!
foo[5] = "bar"
```

#### 配列のスライス

Goと同様に、境界を指定することで配列の一部を抽出することができます。配列 `a` が与えられた場合、 `a[low:high]` 式は、 `low <= index < high` 式を満たすインデックスを持つ `a` の要素からなるサブスライスを抽出します。

したがって、 `a = [1, 2, 3, 4]` が与えられた場合、 `a[1:3]` 式は `[2 3]` に評価されます。

下限または上限を省略すると、それぞれ配列の先頭または末尾に設定されます。つまり、`a = [1, 2, 3, 4]` が与えられると、式 `a[:2]` は `[1 2]` と評価され、`a[1:]` は `[2 3 4]` と評価されるわけです。

この`len`関数を使用すると、より複雑なスライスを実行することができます。

```
a = [1, 2, 3, 4, 5, 6]
b = a[:len(a)-3]			# b == [1 2 3]
c = a[len(a)-4:len(a)-2]	# c == [3 4]
```

<span style="color: red; ">注意：同じ方法で文字列をスライスすることも可能です。このように与えられて`a = "hello"`、に`a[1:4]`評価され`"ell"`ます。</span>

### マップ

Ankoは限定された形式のGoのマップを提供しています（Pythonプログラマーはそれらを「辞書」として認識するでしょう）。Ankoのマップでは、キーは文字列で、値は任意の型です。例えば：

```
a = {}					# define an empty map
a["foo"] = 2			# key = "foo", value is int
a["bar"] = [1, 2, 3]	# key = foo, value is array
```

マップは事前に設定できます。

```
a = {"foo": 2, "bar": [1, 2, 3]}
```

マップから要素を削除するには、次の `delete` 関数を使います:

```
a = {"foo": 2, "bar": [1, 2, 3]}
delete(a, "bar")
```

### チャンネル

AnkoはGoチャンネルを使用するためのインターフェースを提供します。 チャネルは、同時実行に役立つデータの先入れ先出しのパイプラインです。 チャネルは任意のタイプで作成できますが、Ankoの暗黙的な入力のため、 `interface`、`int64`、`float64`、`string`、 `bool`以外のタイプのチャネルでは注意が必要です。 一般に、ほとんどのタスクには`interface`のチャネルで十分です。

チャネルには「サイズ」があります。これは、チャネルブロックへの書き込み前にチャネルが保持できる要素の数です。 サイズ0のチャネルはバッファリングされません。つまり、チャネルへの書き込みは読み取りが実行されるまでブロックされ、その逆も同様です。 サイズ1のチャネルには、書き込みブロックの前に1つのアイテムを書き込むことができます。 チャネルの詳細については、[「効果的なGo」でのチャネルの説明](https://golang.org/doc/effective_go.html#channels)を参照してください。

チャンネルは引数としてチャンネルタイプとオプションのチャンネルサイズをとる `make`関数を使用して作成されます：

```
unbuf = make(chan interface)	# an unbuffered channel
buf = make(chan bool, 1)		# a buffered channel
```

作成されたチャネルは、Goのように読み書きできます。

```
c = make(chan interface, 2)
c <- "foo"
c <- "bar"
a = <- c	# the variable 'a' will contain the string "foo" read from the channel
b = <- c	# variable 'b' contains "bar"
```

## Ankoコードを書く

### 変数の作成、割り当て、スコープ
Ankoは動的スコープを使用します。動的スコープは、Cや他の多くの言語で実装されているレキシカルスコープに慣れているプログラマには馴染みがないかもしれません。 変数は `=`演算子を使用して割り当てられます。 変数が存在しない場合、現在のスコープで作成されます。 名前付き変数が現在のスコープまたはその上の任意のスコープに既に存在する場合、変数の値は新しい変数に設定されます。

次の例は、2番目の割り当て（ `a = 2`）が新しいスコープを作成するのではなく、外部スコープの「a」変数を変更する方法を示しています。

```
func foo() {
	a = 2
}
a = 1
println(a)		# prints "1"
foo()
println(a)		# prints "2"
```
内側のスコープに明示的に新しい変数を作成するには、`var` キーワードを使用します。
```
func foo() {
	var a = 2
}
a = 1
println(a)		# prints "1"
foo()
println(a)		# prints "1"
```
上記のコードは、関数定義の内部スコープ内に 'a'という名前の別の変数を作成します。

### 数学および論理演算

Ankoは、GoやCなどの言語に見られるような基本的な数学的および論理的操作の標準セットを提供します。Ankoは必要に応じて暗黙的に値を変換することに注意してください。

Ankoは以下の数学演算をサポートしています。

| 式の構文| 説明
|-------------------|--------------
| lhs + rhs | lhsとrhsの合計を返します
| lhs-rhs | lhsとrhsの差を返します
| lhs * rhs | rhsを掛けたlhsを返します
| lhs / rhs | rhsをrhsで割った値を返します
| lhs％rhs | rhsを法とするlhsを返します
| lhs == rhs | lhsがrhsと同じ場合にtrueを返します
| lhs！= rhs | lhsがrhsと同じでない場合にtrueを返します
| lhs> rhs | lhsがrhsより大きい場合にtrueを返します
| lhs> = rhs | lhsがrhs以上の場合にtrueを返します
| lhs <rhs | lhsがrhsより小さい場合にtrueを返します
| lhs <= rhs | lhsがrhs以下の場合にtrueを返します
| lhs && rhs | lhsおよびrhsがtrueの場合、trueを返します
| lhs＆＃124;＆＃124; rhs | lhsまたはrhsがtrueの場合、trueを返します
| lhs＆rhs | lhsとrhsのビット単位のANDを返します
| lhs＆＃124; rhs | lhsとrhsのビット単位のORを返します
| lhs << rhs |左にlhsビットシフトされたrhsビットを返します
| lhs >> rhs | lhsビットシフトされたrhsビットを右に返します
| val ++ |ポストインクリメント：valを返し、valをインクリメントします。
| val-- |デクリメント後：valを返し、デクリメントします。
| ^ val | valのビット単位の補数を返します
| ！val | valの否定を返します

標準的なGo / Cの振る舞いに従わないオペレータは以下に文書化されています。

AnkoはC演算子の優先順位規則に従います。

#### +演算子

`+` 予想されるように、オペレータは、必要に応じて浮動小数点数に整数を変換し、番号を追加します:

```
1 + 1		== 2
1.5 + 1		== 2.5
1 + true	== 2	# boolean 'true' evaluates to 1
1 + false 	== 1	# boolean 'false' evaluates to 0
```

また、文字列を連結し、可能な場合は型を変換します。

```
"hello " + "world"	== "hello world"
"anko is #" + 1		== "anko is #1"
2.5 + "apples"		== "2.5apples"
"result is " + true	== "result is true"
```

配列も結合します。

```
[1, 2] + [3, 4]			== [1 2 3 4]
["hi"] + ["there", 7]	== ["hi" "there" 7]
```

#### * 演算子

The `*` オペレータは、標準の乗算を実行します。

```
5 * 2		== 10
3.5 * 3		== 10.5
false * 2	== 0
```

Pythonと同様に文字列の乗算も行います。

```
"hi" * 3	== "hihihi"
```

#### ** 演算子

**オペレータは、べき乗を実行します。

```
2**3	== 8		# 2 to the third power
10**4	== 10000	# 10 to the fourth power
```

### ステートメント

`if`AnkoのステートメントはGoと同じように動作します。

```
if myBoolVar {
	println("myBoolVar is true")
}
if foo == 3 && !bar {
	println("foo is 3, bar is false")
} else if bar {
	println("bar is true")
} else {
	println("neither case")
}
```

<span style="color: red; ">注：Goはフォームを許可しますが、if result := foo(); result true { ... }Ankoでは受け入れられません。</span>

#### 三項演算子

必要に応じて、Ankoでは3項演算子を使用できます。

```
result == 3 ? return true : return false
```

### ループ用

`for` AnkoのループはGoと同じように動作します。

単一の条件を持つループ：

```
for a < b {
	a *= 2
}
```

初期化ステートメント、条件、および事後ステートメントを含むループ：

```
for i = 0; i < max; i++ {
	# do things
}
```

フィボナッチ数列を実装するrange節を持つループ（Ankoリポジトリの例から）:

```
func fib(n) {
  a, b = 1, 1
  f = []
  for i in range(n) {
    f += a
    b += a
    a = b - a
  }
  return f
}
```

### スイッチステートメント
Ankoのswitch文は、デフォルトではフォールスルーしません。`fallthrough`また、Goのようにケースを強制的にフォールスルーさせる文はありません。

例:

```
switch foo {
case "a":
	# do things
case "b":
	# do other things
default:
	# base case
}
```

### 関数宣言

Ankoは動的に型指定されているため、関数定義は戻り型や引数の型を指定しません。暗黙的キャストのため、これは通常は問題ありませんが、関数が引数として配列またはマップを期待する場合は、これを明確にするために関数宣言の上にコメントを挿入することをお勧めします。

引数なしの関数:

```
func incrementCount() {
	counter++
}
```

与えられた引数の2倍を返す関数:

```
func double(x) {
	return 2*x
}
```

関数は可変個引数を取ることができます。引数は配列として関数に渡されます。次の関数は、渡された2番目の引数を出力し、渡された引数の総数を返します。

```
func bar(x ...) {
	println(x[1])
	return len(x)
}

bar(10, 20, 30)		# prints "20", function returns 3.
```

### 並行性

Ankoは`go`ステートメントを使ってゴルーチンを作成できます。ゴルーチン間の通信は通常、チャネルを介して行われます。

この例では、バッファなしのチャンネルを作成してから、そのチャンネルに3つの値を書き込む新しいゴルーチンを起動します。その後、元のゴルーチンは同じチャネルから3つの値を読み取ります。チャネルはバッファされていないので、他のゴルーチンが読み取りを発行するまで、各書き込みはブロックされます。

```
c = make(chan int64)

go func() {
  c <- 1
  c <- 2
  c <- 3
}()

println(<-c)
println(<-c)
println(<-c)
```
