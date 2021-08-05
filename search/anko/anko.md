## Anko

ankoモジュールはevalの補足としてより完全なスクリプト環境を提供します。検索エントリに対してより複雑な操作を可能にしますが、単純なeval式よりもankoスクリプトの開発、テスト、および展開に多くの作業が必要になります。スクリプトは[リソースシステム](#!resources/resources.md)のリソースとして保存されます。

ankoの構文はevalの構文と同じです。どちらも[github.com/mattn/anko](https://github.com/mattn/anko)から派生しており、Gravwell固有のタスク用にいくつかの追加機能が追加されています。

他のモジュールが十分に機能しない状況ではankoを使用することをお勧めします。通常、これは、エントリを以前のエントリと比較する必要がある、エントリを複製する必要がある、エントリからデータを抽出するのに複雑な操作が必要な、またはこれらの組み合わせの状況を意味します。

ドキュメントのこの部分では、ankoモジュールの使い方について簡単に説明しています。より詳細な説明については、[完全なankoモジュールのドキュメント](#!scripting/anko.md)と[Ankoスクリプト言語のドキュメント](#!scripting/scripting.md)を参照してください。

### 構文

`anko <script name> [script arguments]`

Ankoスクリプトはリソースとして保存されています。リソースの名前は、ankoモジュールへの最初の引数として指定する必要があります。スクリプト名の後に、追加の引数がスクリプト自体に渡されます。

### スクリプト例

次のスクリプトは、evalモジュールのドキュメントの例を再フォーマットしたものです。1行のevalの例よりもはるかに読みやすいことに注意してください:

```
func Process() {
	if len(Body) <= 10 {
		setEnum("postlen", "short")
	} else if len(Body) > 10 && len(Body) < 300 {
		setEnum("postlen", "medium")
	} else {
		setEnum("postlen", "long")
	}
	return true
}
```

スクリプトが`CheckPostLen`という名前のリソースにアップロードされていると仮定すると、スクリプトは次のように実行できます:

```
tag=reddit json Body | anko CheckPostLen | count by postlen | table postlen count
```
プロセス`関数`は、ankoモジュールに到達する検索項目ごとに1回実行され、列挙された値 `Body` の長さをチェックし、ボディの長さに基づいて新しい列挙値 `postlen` を設定します。

注意: プロセス関数の最後の `return true` は重要です。プロセス関数は、エントリを通過させるかフィルタリングするかを示すブール値を返します。 falseを返すと、エントリを削除します。true を返すと、エントリがパイプラインを通過することを意味します。

この例は非常に単純で、`Process` 関数のみを実装しています (オプションの `Parse` や `Finalize` 関数は実装していません)。より複雑な例については、[完全なankoモジュールのドキュメント](#!scripting/anko.md)を参照してください。