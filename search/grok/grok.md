# Grok

grokモジュールを使用すると、正規表現全体を毎回指定することなく、複雑なテキスト構造からデータを抽出できます。代わりに、grokは正規表現に名前を割り当て、ユーザーが代わりに名前を指定できるようにします。 Grokパターンには、ネストされた追加のパターンが含まれている場合があり、新しい定義を簡単に作成できます。有用なパターンの選択を事前定義しますが、[リソース](#!resources/resources.md)から独自のカスタマイズされたパターンのセットを読み取ることもできます。

デフォルトでは、grokはパターンに一致するエントリをすべて通過し、一致しないエントリはすべてドロップします。この動作は、`-v`フラグで反転できます。

Grokはフィルタリングモジュールです。目的のパターンを指定した後、抽出されたフィールドに適用するフィルターのリストを指定することもできます。

注：一部のフィルターには非常に厳密で複雑なパターンが組み込まれているため、多数のエントリーを処理する場合は比較的遅くなる可能性があります。 [grep]（＃！search / grep / grep.md）、[regex](#!search/regex/regex.md)、[words](#!search/words/words.md)などのモジュールを使用して可能な限り事前にフィルタします。

## サポートされているオプション 

* `-e <arg>`: レコード全体ではなく、指定された列挙値を操作します。
* `-r <resource>`: デフォルトの `grok`リソースではなく、指定された名前のリソースからカスタムgrokパターンをロードします。
* `-v`: 反転モードで動作します;パターンに*一致しない*エントリは渡され、*一致する*エントリはドロップされます。このフラグを使用する場合、フィルターを指定できません。
* `-p`: "-p"オプションは、式がまったく一致しない場合にエントリを許可するようにgrokに指示します。許容フラグは、フィルターの動作を変更しません。

### Apacheログを解析する

次のクエリは、リソースを利用して複雑なパターンを実装し、Apache2.0結合アクセスログを厳密に処理します。 `COMBINEDAPACHELOG`パターンは、[github](https://raw.githubusercontent.com/gravwell/resources/master/grok/all.grok)でGravwellが提供する非常に大きなパターンセットの一部です。パターンセットをダウンロードし、リソース名`grok`としてアップロードして、定義済みのgrokパターンの大規模なスイートにアクセスします。

次のクエリは、"PUT"リクエストのすべてのApacheログを検索し、それらをコンポーネントに解析します:

```
tag=apache words PUT | grok "%{COMBINEDAPACHELOG}" | table
```

![](apache.png)

注：COMBINEDAPACHELOGパターンは複雑で非常に厳密であるため、数百万のエントリがある場合、このクエリには時間がかかることがあります。

### フィルタリング

前のクエリに基づいて、"clientip"フィールドが特定のIPと一致し、PUTメソッドを使用するエントリのみを返すように構築できます:

```
tag=apache words PUT 128.10.247.36 | grok "%{COMBINEDAPACHELOG}" clientip=="128.10.247.36" verb==PUT | table clientip
```

![](apache-filter.png)

注：単語モジュールを使用してPUTとIPをフィルター処理し、インデックス作成を行い、高価な COMBINEDAPACHELOG` grokパターンによって処理されるエントリの数を減らします。

## パフォーマンス

Grokは複雑な正規表現を劇的に単純化し、単なる人間が大きなログフラグメントを分解できるようにします。ただし、ログ内のすべてのフィールドを抽出して検証するように設計されたgrokパターンは、複雑で遅くなります。すべてのフィールドが必要でない場合は、フラグメントとプリミティブの使用を検討してください。少数のアイテムのみを抽出する小さなgrokパターンは、すべてを抽出する完全なパターンよりも劇的に高速です。

たとえば、各HTTPメソッドの応答コードカウントをコンパイルしてスタックグラフに表示する2つのクエリを見てみましょう。最初のクエリではgrokを使用し、`COMBINEDAPACHELOG`パターンで非常に単純なクエリを実行します:

```
tag=apache grok "%{COMBINEDAPACHELOG}" | stats count by verb response | stackgraph verb response count
```

2番目のクエリはgrokプリミティブを使用して、明示的に必要なフィールドのみを抽出します:

```
tag=apache grok "] \"%{WORD:verb}\s\S+\s\S+\s%{POSINT:response}" | stats count by verb response | stackgraph verb response count
```

両方のクエリは同じ結果を生成します:

![](apachestackgraph.png)

ただし、10MのApacheアクセスログを処理するには、最初のクエリに`2m 39s`が必要でした。 2番目のクエリは、`3.46s`だけで済み、45倍のスピードアップになりました。そのため、単純なクエリは見栄えがよくなりますが、大規模なデータセットを操作するときにプリミティブパターンを操作するのに時間をかける価値があります。

## 事前定義されたパターン

Grokモジュールは、すぐに使用できる定義済みパターンの基本セットを提供します。これらの基本パターンは基本的なデータ型をカバーし、一般に、受け入れられるものと受け入れられないものについて非常に厳格です。ログセット全体を処理するように設計されたパターンの大きなセットについては、[github](https://raw.githubusercontent.com/gravwell/resources/master/grok/all.grok)

| パターン名 | 置換パターン                                                                               |
| ------------ | -------------------------------------------------------------------------------------------------- |
| USERNAME | `[a-zA-Z0-9._-]+` |
| USER | `%{USERNAME}` |
| EMAILLOCALPART | `[a-zA-Z][a-zA-Z0-9_.+-=:]+` |
| HOSTNAME | `\b[0-9A-Za-z][0-9A-Za-z-]{0,62}(?:\.[0-9A-Za-z][0-9A-Za-z-]{0,62})*(\.?&#124;\b)` |
| EMAILADDRESS | `%{EMAILLOCALPART}@%{HOSTNAME}` |
| HTTPDUSER | `%{EMAILADDRESS}&#124;%{USER}` |
| INT | `[+-]?(?:[0-9]+)` |
| BASE10NUM | `[+-]?(?:(?:[0-9]+(?:\.[0-9]+)?)&#124;(?:\.[0-9]+))` |
| NUMBER | `%{BASE10NUM}` |
| BASE16NUM | `[+-]?(?:0x)?(?:[0-9A-Fa-f]+)` |
| BASE16FLOAT | `\b[+-]?(?:0x)?(?:(?:[0-9A-Fa-f]+(?:\.[0-9A-Fa-f]*)?)&#124;(?:\.[0-9A-Fa-f]+))\b` |
| POSINT | `\b[1-9][0-9]*\b` |
| NONNEGINT | `\b[0-9]+\b` |
| WORD | `\b\w+\b` |
| NOTSPACE | `\S+` |
| SPACE | `\s*` |
| DATA | `.*?` |
| GREEDYDATA | `.*` |
| QUOTEDSTRING | ``("(\\.&#124;[^\\"]+)+")&#124;""&#124;('(\\.&#124;[^\\']+)+')&#124;''&#124;(`(\\.|[^\\`]+)+`)|`` |
| UUID | `[A-Fa-f0-9]{8}-(?:[A-Fa-f0-9]{4}-){3}[A-Fa-f0-9]{12}` |
| UNIXPATH | `(/([\w_%!$@:.,~-]+&#124;\\.)*)+` |
| TTY | `/dev/(pts&#124;tty([pq])?)(\w+)?/?(?:[0-9]+)` |
| WINPATH | `(?:[A-Za-z]+:&#124;\\)(?:\\[^\\?*]*)+` |
| PATH | `%{UNIXPATH}&#124;%{WINPATH}` |
| URIPROTO | `[A-Za-z]+(\+[A-Za-z+]+)?` |
| URIHOST | `%{IPORHOST}(?::%{POSINT:port})?` |
| URIPATH | `(?:/[A-Za-z0-9$.+!*'(){},~:;=@#%_\-]*)+` |
| URIPARAM | `\?[A-Za-z0-9$.+!*'&#124;(){},~@#%&/=:;_?\-\[\]<>]*` |
| URIPATHPARAM | `%{URIPATH}(?:%{URIPARAM})?` |
| URI | `%{URIPROTO}://(?:%{USER}(?::[^@]*)?@)?(?:%{URIHOST})?(?:%{URIPATHPARAM})?` |
| MONTH | `\bJan(?:uary&#124;uar)?&#124;Feb(?:ruary&#124;ruar)?&#124;M(?:a&#124;ä)?r(?:ch&#124;z)?&#124;Apr(?:il)?&#124;Ma(?:y&#124;i)?&#124;Jun(?:e&#124;i)?&#124;Jul(?:y)?&#124;Aug(?:ust)?&#124;Sep(?:tember)?&#124;O(?:c&#124;k)?t(?:ober)?&#124;Nov(?:ember)?&#124;De(?:c&#124;z)(?:ember)?\b` |
| MONTHNUM | `0?[1-9]&#124;1[0-2]` |
| MONTHNUM2 | `0[1-9]&#124;1[0-2]` |
| MONTHDAY | `(?:0[1-9])&#124;(?:[12][0-9])&#124;(?:3[01])&#124;[1-9]` |
| DAY | `Mon(?:day)?&#124;Tue(?:sday)?&#124;Wed(?:nesday)?&#124;Thu(?:rsday)?&#124;Fri(?:day)?&#124;Sat(?:urday)?&#124;Sun(?:day)?` |
| YEAR | `(?:\d\d){1,2}` |
| HOUR | `2[0123]&#124;[01]?[0-9]` |
| MINUTE | `[0-5][0-9]` |
| SECOND | `(?:[0-5]?[0-9]&#124;60)(?:[:.,][0-9]+)?` |
| TIME | `%{HOUR}:%{MINUTE}:%{SECOND}` |
| DATE_US | `%{MONTHNUM}[/-]%{MONTHDAY}[/-]%{YEAR}` |
| DATE_EU | `%{MONTHDAY}[./-]%{MONTHNUM}[./-]%{YEAR}` |
| DATE_X | `%{YEAR}/%{MONTHNUM2}/%{MONTHDAY}` |
| ISO8601_TIMEZONE | `Z&#124;[+-]%{HOUR}(?::?%{MINUTE})` |
| ISO8601_SECOND | `%{SECOND}&#124;60` |
| TIMESTAMP_ISO8601 | `%{YEAR}-%{MONTHNUM}-%{MONTHDAY}[T ]%{HOUR}:?%{MINUTE}(?::?%{SECOND})?%{ISO8601_TIMEZONE}?` |
| DATE | `%{DATE_US}&#124;%{DATE_EU}&#124;%{DATE_X}` |
| DATESTAMP | `%{DATE}[- ]%{TIME}` |
| TZ | `[A-Z]{3}` |
| NUMTZ | `[+-]\d{4}` |
| DATESTAMP_RFC822 | `%{DAY} %{MONTH} %{MONTHDAY} %{YEAR} %{TIME} %{TZ}` |
| DATESTAMP_RFC2822 | `%{DAY}, %{MONTHDAY} %{MONTH} %{YEAR} %{TIME} %{ISO8601_TIMEZONE}` |
| DATESTAMP_OTHER | `%{DAY} %{MONTH} %{MONTHDAY} %{TIME} %{TZ} %{YEAR}` |
| DATESTAMP_EVENTLOG | `%{YEAR}%{MONTHNUM2}%{MONTHDAY}%{HOUR}%{MINUTE}%{SECOND}` |
| HTTPDERROR_DATE | `%{DAY} %{MONTH} %{MONTHDAY} %{TIME} %{YEAR}` |
| ANSIC | `%{DAY} %{MONTH} [_123]\d %{TIME} %{YEAR}"` |
| UNIXDATE | `%{DAY} %{MONTH} [_123]\d %{TIME} %{TZ} %{YEAR}` |
| RUBYDATE | `%{DAY} %{MONTH} [0-3]\d %{TIME} %{NUMTZ} %{YEAR}` |
| RFC822Z | `[0-3]\d %{MONTH} %{YEAR} %{TIME} %{NUMTZ}` |
| RFC850 | `%{DAY}, [0-3]\d-%{MONTH}-%{YEAR} %{TIME} %{TZ}` |
| RFC1123 | `%{DAY}, [0-3]\d %{MONTH} %{YEAR} %{TIME} %{TZ}` |
| RFC1123Z | `%{DAY}, [0-3]\d %{MONTH} %{YEAR} %{TIME} %{NUMTZ}` |
| RFC3339 | `%{YEAR}-[01]\d-[0-3]\dT%{TIME}%{ISO8601_TIMEZONE}` |
| RFC3339NANO | `%{YEAR}-[01]\d-[0-3]\dT%{TIME}\.\d{9}%{ISO8601_TIMEZONE}` |
| KITCHEN | `\d{1,2}:\d{2}(AM&#124;PM&#124;am&#124;pm)` |
| SYSLOGTIMESTAMP | `%{MONTH} +%{MONTHDAY} %{TIME}` |
| LOGLEVEL | `[Aa]lert&#124;ALERT&#124;[Tt]race&#124;TRACE&#124;[Dd]ebug&#124;DEBUG&#124;[Nn]otice&#124;NOTICE&#124;[Ii]nfo&#124;INFO&#124;[Ww]arn?(?:ing)?&#124;WARN?(?:ING)?&#124;[Ee]rr?(?:or)?&#124;ERR?(?:OR)?&#124;[Cc]rit?(?:ical)?&#124;CRIT?(?:ICAL)?&#124;[Ff]atal&#124;FATAL&#124;[Ss]evere&#124;SEVERE&#124;EMERG(?:ENCY)?&#124;[Ee]merg(?:ency)?` |
| QS | `%{QUOTEDSTRING}` |
| NQS | `[^"]*` |
| PROG | `[\x21-\x5a\x5c\x5e-\x7e]+` |
| CISCOMAC | `(?:[A-Fa-f0-9]{4}\.){2}[A-Fa-f0-9]{4}` |
| WINDOWSMAC | `(?:[A-Fa-f0-9]{2}-){5}[A-Fa-f0-9]{2}` |
| COMMONMAC | `(?:[A-Fa-f0-9]{2}:){5}[A-Fa-f0-9]{2}` |
| MAC | `%{CISCOMAC}&#124;%{WINDOWSMAC}&#124;%{COMMONMAC}` |
| IPV4 | `(?:(?:25[0-5]&#124;2[0-4][0-9]&#124;[01]?[0-9][0-9]?)\.){3}(?:25[0-5]&#124;2[0-4][0-9]&#124;[01]?[0-9][0-9]?)` |
| IPV6 | `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}&#124;((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3})&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})&#124;:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3})&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})&#124;((:[0-9A-Fa-f]{1,4})?:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3}))&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})&#124;((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3}))&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})&#124;((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3}))&#124;:))&#124;(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})&#124;((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3}))&#124;:))&#124;(:(((:[0-9A-Fa-f]{1,4}){1,7})&#124;((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)(\.(25[0-5]&#124;2[0-4]\d&#124;1\d\d&#124;[1-9]?\d)){3}))&#124;:)))(%.+)?` |
| IP | `%{IPV6}&#124;%{IPV4}` |
| IPORHOST | `%{IP}&#124;%{HOSTNAME}` |
| HOSTPORT | `%{IPORHOST}:%{POSINT}` |
| HTTPDATE | `%{MONTHDAY}/%{MONTH}/%{YEAR}:%{TIME} %{INT}` |
