# time

timeモジュールは、タイムスタンプ列挙値をフォーマットされた文字列に変換するために使用されます。例えば、以下の例では、指定されたフォーマットを使用して各エントリの組み込みタイムスタンプ値を表示し、その結果を「formattedTS」という名前の新しい列挙値に文字列として保存します。

```
time -f "Mon Jan _2 15:04:05 2006 MST" TIMESTAMP formattedTS
```

列挙値の最初の引数がタイムスタンプではなく文字列の場合、モジュールは代わりに文字列を時間として解析し、出力をタイムスタンプ列挙値に保存しようとします。特定のフォーマットが与えられない場合、モジュールは `timegrinder` ライブラリを使用して多くの可能性を試すことに注意してください。以下は、「tsString」という名前の文字列列挙値を見て、それをタイムスタンプに変換しようとし、成功した場合には結果を「extractedTS」に格納します。

```
time tsString extractedTS
```

## サポートされているオプション

* `-f <format>`: タイムスタンプを表示するときや、オプションで文字列を解析するときに使用するフォーマットを指定します。フォーマットは、[Go time library](https://golang.org/pkg/time/#pkg-constants)で使用されている "Mon Jan 2 15:04:05 MST 2006" という特定の時刻を文字列で表現したものです。例えば、`-f "Mon 3:04PM"` とすると、非常に短いタイムスタンプ形式が得られます。より多くの例については、リンク先のドキュメントを参照してください。
* `-tz <timezone>`: タイムゾーンを [tz データベース形式](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) で指定します。このタイムゾーンは、タイムスタンプ(タイムゾーンに関連付けられていないもの)を表示する際や、タイムゾーン指定を含まない文字列を解析する際に使用されます。

## 例

エントリーのタイムスタンプを特定のフォーマットとタイムゾーンで表示する。

```
tag=json time -f "Mon Jan _2 15:04:05 2006 MST" -tz "America/Chicago" TIMESTAMP foo | table TIMESTAMP foo
```

![](time1.png)

前のモジュール呼び出しの出力をタイムスタンプに変換するために、タイムモジュールにフィードバックすることができます。

```
tag=json time -f "Mon Jan _2 15:04:05 2006 MST" -tz "America/Chicago" TIMESTAMP foo | time -f "Mon Jan _2 15:04:05 2006 MST" -tz "America/Chicago" foo bar | table TIMESTAMP foo bar
```

![](time2.png)

中間時点フォーマットには端数秒が含まれていないため、最終的な出力では端数秒が切り捨てられていることに注意してください。
