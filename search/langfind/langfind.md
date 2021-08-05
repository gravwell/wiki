## Langfind

langfindモジュールは、テキストのデータを検索し、そのテキストを人間の言語として分類しようとする基本的な人間言語分析モジュールです。

### 検索例

以下の検索では、Redditのコメントで最も一般的に使用されている言語を、人気のあるものから少ないものへと降順に並べた表を作成します。

```
tag=reddit json Body | langfind -e Body | count by lang | sort by count desc | table lang count
```

### サポートされているオプション

* `-e <arg>`: "e"オプションはレコード全体ではなく、列挙値に対して操作する。例えば、HTTPペイロードの言語解析を行うパイプラインは、`tag=pcap ipv4.DstPort==80 tcp.Payload | langfind -e Payload`のようになります。
* デフォルトでは、列挙値 "lang "で出力されます。オプションで、最後の引数に列挙値名を指定することができます。例えば、列挙値 "foo "で出力を生成するには:

```
tag=reddit json Body | langfind -e Body foo | table foo
```
