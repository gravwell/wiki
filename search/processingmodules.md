# 検索処理モジュール

検索モジュールは、パススルーモードでデータを操作するモジュールです。つまり、何らかのアクション（フィルタ、修正、ソートなど）を実行し、そのエントリーをパイプラインに渡します。検索モジュールは多数存在し、それぞれが独自の軽量スレッドで動作します。 つまり、1つの検索モジュールに10個のモジュールがある場合、パイプラインは10個のスレッドに分散して使用されます。 各モジュールのドキュメントには、そのモジュールが分散検索の崩壊やソートを引き起こすかどうかが示されています。 モジュールが崩壊すると、分散したパイプラインも崩壊し、そのモジュールと下流のモジュールがフロントエンドで実行されることになります。 検索を開始する際には、最初に崩壊するモジュールの上流にできるだけ多くの並列モジュールを置くことで、通信パイプの圧力を減らし、より大きな並列性を確保することができます。

モジュールの中には、`count` のようにデータを大幅に変換したり、折り畳んだりするものがあります。これらの折りたたみモジュールに続くパイプラインモジュールは、生データや以前に作成された列挙型の値を扱っていない可能性があります。つまり、`count by Src` のようなモジュールは、データを `10.0.0.1 3` のようなエントリで折り畳まれた結果にします。例を挙げると、検索で `tag=* limit 10 | count by TAG | raw` を実行すると、count モジュールからの生の出力を見ることができ、`tag=* limit 10 | count by TAG | table TAG count DATA` を実行すると、table モジュールから見て生のデータが凝縮されているのがわかります。

## ユニバーサルフラグ

フラグの中には、複数の異なる検索モジュールに表示され、同じ意味を持つものがあります。

* `-e <ソース名>` は、モジュールがエントリーのデータフィールドからではなく、与えられた列挙値から入力データを読み取ろうとすることを指定します。これは、[json](json/json.md) のようなモジュールで、JSON エンコードされたデータがより大きなデータレコードから抽出されている場合に便利です。例えば、以下の検索は、HTTP パケットのペイロードから JSON フィールドを読み取ろうとします： `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-r <リソース名>` は、[resources](#!resources/resources.md) システムのリソースを指定します。これは一般的に、[geoip](geoip/geoip.md) モジュールが使用する GeoIP マッピングテーブルのように、モジュールが使用する追加データを格納するために使用されます。
* `-v` は、通常のパス/ドロップのロジックを反転させることを意味します。例えば、[grep](grep/grep.md) モジュールは、通常、与えられたパターンにマッチしたエントリーを通過させ、マッチしないエントリーをドロップしますが、`-v` フラグを指定すると、マッチしたエントリーをドロップし、マッチしないエントリーを通過させます。
* `-s` は strict モードを示します。モジュールが通常、いくつかの条件のうちどれかひとつが満たされていれば、エントリーをパイプラインで進めることができる場合、strict  フラグを設定すると、すべての条件が満たされている場合にのみエントリーを進めることができます。たとえば、[require](require/require.md) モジュールは、通常は、必要な列挙型の値のいずれかが含まれていればエントリーを通過させますが、`-s` フラグが使用されると、指定された列挙型の値がすべて含まれているエントリーのみを通過させます。
* `-p` は permissive モードを示します。 パターンやフィルターがマッチしないときに、通常モジュールがエントリーをドロップする場合、permissive フラグはモジュールを通すように指示します。 [regex](regex/regex.md) モジュールや [grok](grok/grok.md) モジュールは、この permissive フラグが有効な例です。

## ユニバーサルな列挙値

以下の列挙値は、すべてのエントリーで利用可能です。これらは実際には、生のエントリー自体のプロパティの便利な名前ですが、列挙値の名前として扱うことができます。

* SRC -- エントリーデータのソース。
* TAG -- エントリーに付けられたタグ。
* TIMESTAMP -- エントリーのタイムスタンプ。
* DATA -- 実際のエントリーのデータ。
* NOW -- 現在の時刻。

これらは、ユーザー定義の列挙型の値と同じように使用できるので、`table foo bar DATA NOW` は有効です。これらは、どこかに明示的に抽出する必要はなく、常に利用可能です。

## 検索モジュールのドキュメント

* [abs](abs/abs.md)
* [alias](alias/alias.md)
* [anko](anko/anko.md)
* [base64](base64/base64.md)
* [count](math/math.md#Count)
* [diff](diff/diff.md)
* [enrich](enrich/enrich.md)
* [entropy](math/math.md#Entropy)
* [eval](eval/eval.md)
* [first/last](firstlast/firstlast.md)
* [geoip](geoip/geoip.md)
* [grep](grep/grep.md)
* [hexlify](hexlify/hexlify.md)
* [ip](ip/ip.md)
* [ipexist](ipexist/ipexist.md)
* [iplookup](iplookup/iplookup.md)
* [join](join/join.md)
* [langfind](langfind/langfind.md)
* [length](length/length.md)
* [limit](limit/limit.md)
* [lookup](lookup/lookup.md)
* [lower](upperlower/upperlower.md)
* [maclookup](maclookup/maclookup.md)
* [Math (list of math modules)](math/math.md)
* [max](math/math.md#Max)
* [mean](math/math.md#Mean)
* [min](math/math.md#Min)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [require](require/require.md)
* [slice](slice/slice.md)
* [sort](sort/sort.md)
* [split](split/split.md)
* [src](src/src.md)
* [stats](stats/stats.md)
* [stddev](math/math.md#Stddev)
* [strings](strings/strings.md)
* [subnet](subnet/subnet.md)
* [sum](math/math.md#Sum)
* [taint](taint/taint.md)
* [time](time/time.md)
* [transaction](transaction/transaction.md)
* [unique](math/math.md#Unique)
* [upper](upperlower/upperlower.md)
* [variance](math/math.md#Variance)
* [words](words/words.md)
