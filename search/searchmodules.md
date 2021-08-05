# Search Modules

検索モジュールは、パススルーモードでデータを処理するモジュールです。  つまり、検索モジュールは何らかのアクション（フィルタ、変更、並べ替えなど）を実行し、エントリをパイプラインに渡します。  検索モジュールは多数存在する可能性があり、それぞれが独自の軽量スレッドで動作します。  つまり、検索に10個のモジュールがある場合、パイプラインは10個のスレッドを使用して分散します。  各モジュールのドキュメントは、そのモジュールによって分散検索が折りたたまれたりソートされたりするかどうかを示します。  モジュールが崩壊すると、分散パイプラインは強制的に崩壊します。  つまり、モジュールとすべてのダウンストリームモジュールはフロントエンドで実行されます。  検索を開始するときは、最初の折りたたみモジュールの上流にできるだけ多くの並列モジュールを配置して、通信パイプへの負担を軽減し、より大きな並列処理を可能にすることが最善です。

## ユニバーサルフラグ

下記のフラグはいくつかの異なる検索モジュール間で共通して使用することができます。:

* `-e <source name>` モジュールがエントリのデータフィールドからではなく、指定された列挙値から入力データを読み取ろうとすることを指定します。これは、JSONエンコードされたデータがより大きなデータレコードから抽出された可能性があるjsonのようなモジュールに役立ちます。たとえば、次の検索はHTTPパケットのペイロードからJSONフィールドを読み込もうとします: `tag=pcap packet tcp.Payload | json -e Payload user.email`
* `-t <target name>` モジュールがソースを上書きするのではなく、指定された名前の列挙値に出力を書き込むように指定します。たとえば、[hexlify](hexlify/hexlify.md)モジュールは通常、16進エンコードされた出力文字列をエントリのデータフィールドに書き戻しますが、`-t`フラグが指定されている場合は、代わりにソースをそのままにして名前付き列挙値に出力を書き込みます。
* `-r <resource name>` [resources](#!resources/resources.md)システム内のリソースを指定します。これは通常、[geoip](geoip/geoip.md)モジュールが使用するGeoIPマッピングテーブルなど、モジュールが使用する追加データを保存するために使用されます
* `-v` 通常のパス/ドロップロジックを反転する必要があることを示します。例えば、[grep](grep/grep.md)モジュールは通常、与えられたパターンにマッチするエントリを渡し、マッチしないものを削除します。`-v`フラグを指定すると、一致するエントリを削除し、一致しないエントリを渡します。
* `-s` "厳密な"モードを示します。いくつかの条件のうちのどれか1つでも満たされている場合、モジュールが通常エントリがパイプラインを進むことを許可する場合、strictフラグを設定することは、すべての条件が満たされた場合にのみエントリが進むことを意味します。たとえば、[require](require/require.md)モジュールは通常、必要な列挙値のいずれかが含まれている場合は`-s`エントリを渡しますが、フラグが使用されている場合は、指定された*すべて*の列挙値を含むエントリのみを渡します。
* `-p` "許容"モードを示します。 パターンとフィルターが一致しないときにモジュールが通常エントリーをドロップする場合、permissiveフラグはモジュールが通過できるようにモジュールに指示します。 [regex](regex/regex.md)および[grok](grok/grok.md)モジュールは、permissiveフラグが有用な良い例です。

## ユニバーサル列挙値

すべての検索モジュールには、レコードのユニバーサル列挙値があります

* SRC -- エントリデータのソース
* TAG -- エントリに添付されているタグ
* TIMESTAMP -- エントリのタイムスタンプ
* DATA -- 実際のエントリデータ

これらはユーザー定義の列挙値と同じように使用できます。

## 検索モジュールの詳細

* [abs](abs/abs.md)
* [alias](alias/alias.md)
* [anko](anko/anko.md)
* [ax](ax/ax.md)
* [base64](base64/base64.md)
* [canbus](canbus/canbus.md)
* [cef](cef/cef.md)
* [count](math/math.md#Count)
* [csv](csv/csv.md)
* [entropy](math/math.md#Entropy)
* [eval](eval/eval.md)
* [fields](fields/fields.md)
* [geoip](geoip/geoip.md)
* [grep](grep/grep.md)
* [grok](grok/grok.md)
* [hexlify](hexlify/hexlify.md)
* [ip](ip/ip.md)
* [ipexist](ipexist/ipexist.md)
* [ipfix](ipfix/ipfix.md)
* [j1939](j1939/j1939.md)
* [join](join/join.md)
* [json](json/json.md)
* [langfind](langfind/langfind.md)
* [length](length/length.md)
* [limit](limit/limit.md)
* [lookup](lookup/lookup.md)
* [lower](upperlower/upperlower.md)
* [Math (list of math modules)](math/math.md)
* [max](math/math.md#Max)
* [mean](math/math.md#Mean)
* [min](math/math.md#Min)
* [namedfields](namedfields/namedfields.md)
* [netflow](netflow/netflow.md)
* [packet](packet/packet.md)
* [packetlayer](packetlayer/packetlayer.md)
* [regex](regex/regex.md)
* [require](require/require.md)
* [slice](slice/slice.md)
* [sort](sort/sort.md)
* [src](src/src.md)
* [stats](stats/stats.md)
* [stddev](math/math.md#Stddev)
* [strings](strings/strings.md)
* [subnet](subnet/subnet.md)
* [sum](math/math.md#Sum)
* [syslog](syslog/syslog.md)
* [taint](taint/taint.md)
* [unique](math/math.md#Unique)
* [upper](upperlower/upperlower.md)
* [variance](math/math.md#Variance)
* [winlog](winlog/winlog.md)
* [words](words/words.md)
* [xml](xml/xml.md)
