# インジェスト・プリプロセッサー

取り込んだデータをインデクサーに送信する前に、さらに加工が必要な場合があります。例えば、syslogで送られてきたJSONデータを受信して、syslogのヘッダーを除去したい場合。Apache Kafkaのストリームからgzipで圧縮されたデータを取得している場合。エントリの内容に基づいて、エントリを異なるタグにルーティングしたい場合もあるでしょう。インジェストのプリプロセッサは、エントリーがインデクサに送られる前に1つ以上の処理ステップを挿入することで、これを可能にします。

## プリプロセッサのデータフロー

インゲスターは、何らかのソース（ファイル、ネットワーク接続、Amazon Kinesisストリームなど）から生データを読み込み、その入力データストリームを個々のエントリに分割します。これらのエントリがGravwellのインデクサに送信される前に、任意の数のプリプロセッサを通過させることができます。

![](arch.png)

それぞれのプリプロセッサは、エントリーを修正する機会があります。つまり、例えば、エントリーのデータを解凍してから、解凍されたデータに基づいてエントリーのタグを修正することができます。

## プリプロセッサの設定

プリプロセッサは、すべてのパッケージ化されたインジェスターでサポートされています。 単発のインジェスターやサポートされていないインジェスターは、プリプロセッサーをサポートしていない場合があります。

プリプロセッサは、インゲスターの設定ファイルで `preprocessor` 設定スタンザを使用して設定されます。 各プリプロセッサスタンザは、使用するプリプロセッサモジュールを `Type` 設定パラメータで宣言し、続いてプリプロセッサの特定の設定パラメータを宣言しなければなりません。Simple Relay インゲスターの例を以下に示します。

```
[Global]
Ingester-UUID="e985bc57-8da7-4bd9-aaeb-cc8c7d489b42"
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=true
Cleartext-Backend-target=127.0.0.1:4023 #example of adding a cleartext connection
Log-Level=INFO

[Listener "default"]
	Bind-String="0.0.0.0:7777" #we are binding to all interfaces, with TCP implied
	Tag-Name=default
	Preprocessor=timestamp

[Listener "syslog"]
	Bind-String="0.0.0.0:601" # TCP syslog
	Tag-Name=syslog

[preprocessor "timestamp"]
	Type = regextimestamp
	Regex ="(?P<badtimestamp>.+) MSG (?P<goodtimestamp>.+) END"
	TS-Match-Name=goodtimestamp
	Timezone-Override=US/Pacific
```

この設定では、"default"と "syslog"という2つのデータコンシューマー（Simple Relayでは "Listener"と呼びます）を定義しています。また、"timestamp"という名前のプリプロセッサも定義しています。default "リスナーには、`Preprocessor=timestamp`というオプションが付いていることに注目してください。これは、ポート7777でそのリスナーから送られてくるエントリーが「timestamp」プリプロセッサに送られることを指定しています。syslog "リスナーは`Preprocessor`オプションを設定していないので、ポート601で入ってくるエントリーはどのプリプロセッサも経由しません。

## gzip プリプロセッサ

gzipプリプロセッサは、GNU 'gzip' アルゴリズムで圧縮されたエントリを解凍することができます。

GZIP プリプロセッサの型は `gzip` です。

### サポートされるオプション

* `Passthrough-Non-Gzip` (boolean, optional): true に設定すると、gzip で解凍できない内容のエントリーはスルーされます。デフォルトでは、プリプロセッサは gzip で圧縮されていないすべてのエントリーを削除します。

### 一般的な使用例

多くのクラウドデータバスプロバイダーは、エントリーやパッケージを圧縮して出荷します。 このプリプロセッサは、クラウドのラムダ関数を経由するのではなく、インゲスターでデータストリームを解凍することができますが、コストがかかります。


### 例：圧縮されたエントリーの解凍

設定例:

```
[Preprocessor "gz"]
	Type=gzip
	Passthrough-Non-Gzip=true
```

## JSON抽出プリプロセッサ

JSON 抽出プリプロセッサは、エントリの内容を JSON としてパースし、その JSON から 1 つ以上のフィールドを抽出し、エントリの内容をそのフィールドで置き換えることができます。これは、複雑すぎるメッセージを、必要な情報だけを含むより簡潔なエントリーに簡略化するのに便利な方法です。

複数のフィールドが指定された場合、プリプロセッサはそれらのフィールドを含む有効なJSONを生成します。

JSON抽出のプリプロセッサの型は `jsonextract` です。

### サポートされるオプション

* `Extractions` (文字列、必須)。JSONから抽出するフィールドを指定します（コンマで区切ってください）。入力が `{"foo": "a", "bar":2, "baz":{"frog": "womble"}}` とすると、`Extractions=foo`, `Extractions=foo,bar`, `Extractions=baz.frog,foo` などと指定することができます。
* `Force-JSON-Object` (boolean, optional): デフォルトでは、単一の抽出が指定された場合、プリプロセッサはエントリーのコンテンツをその拡張子のコンテンツで置き換えます。したがって、`Extraction=foo`を選択すると、`{"foo": "a", "bar":2, "baz":{"frog": "womble"}}` を含むエントリを、単に `a` を含むように変更します。このオプションが設定されていると、プリプロセッサは常に完全なJSON構造を出力します。例えば、`{"foo": "a"}`です。
* `Passthrough-Misses` (boolean, optional): trueに設定すると、プリプロセッサは要求されたフィールドを抽出できなかったエントリーを通過させます。デフォルトでは、これらのエントリは削除されます。
* `Strict-Extraction` (boolean, optional): デフォルトでは、少なくとも1つの抽出が成功した場合、プリプロセッサはエントリを渡します。このパラメータがtrueに設定されている場合、すべての抽出が成功することが要求されます。

### 一般的な使用例

多くのデータソースは、実際のログストリームの一部ではない、輸送や保管に関連する追加のメタデータを提供することがあります。 jsonextract プリプロセッサは、ストレージコストを削減するためにフィールドをダウンセレクトすることができます。

### 例：JSONデータレコードの凝縮

```
[Preprocessor "json"]
	Type=jsonextract
	Extractions=IP,Alert.ID,Message
	Passthrough-Misses=true
```

## JSON配列分割プリプロセッサ

このプリプロセッサは、JSONオブジェクトの配列を個々のエントリに分割できます。 たとえば、名前の配列を含むエントリが与えられた場合、プリプロセッサは代わりに名前ごとに1つのエントリを発行します。 したがって、これは：

```
{"IP": "10.10.4.2", "Users": ["bob", "alice"]}
```

"bob"を含むエントリーと "alice"を含むエントリーの2つになります。

JSON Array Splitのプリプロセッサタイプは、`jsonarraysplit`です。

### サポートされるオプション

* `Extraction` (string): 分割すべき構造体を含む JSON フィールドを指定します。例: `Extraction=Users`, `Extraction=foo.bar`. 例えば、`Extraction=Users`, `Extraction=foo.bar` などです。`Extraction` を設定しない場合、プリプロセッサはオブジェクト全体を分割するための配列として扱おうとします。
* `Passthrough-Misses` (boolean, optional): trueに設定された場合、プリプロセッサは要求されたフィールドを抽出できなかったエントリーを通過させます。デフォルトでは、これらのエントリーはドロップされます。
* `Force-JSON-Object` (boolean, optional): デフォルトでは、プリプロセッサはそれぞれがリストの1つのアイテムを含み、それ以外は何も含まないエントリーを出力します。["a", "b"]}` から `foo` を抽出すると、それぞれ "a" と "b" を含む 2 つのエントリになります。このオプションが設定されていると、同じエントリーでも `{"foo": "a"}` と `{"foo": "b"}`.
* `Additional-Fields` (string, optional): 例えば、`Additional-Fields="foo,bar, foo.bar.baz"`のように、分割される配列の外側にある追加フィールドをカンマ区切りで指定します。

### 一般的な使用例

多くのデータプロバイダーは、複数のイベントを1つのエントリにまとめることがありますが、これはイベントの原子性を低下させ、分析の複雑さを増大させます。 複数のイベントを含む1つのメッセージを個々のエントリに分割することで、イベントの扱いが簡単になります。


### 例：複数のメッセージを1つのレコードに分割する

「アラート」という名前の配列を持つJSONレコードで構成されるエントリを分割します。

```
[preprocessor "json"]
	Type=jsonarraysplit
	Extraction=Alerts
	Force-JSON-Object=true
```

インプットデータ:

```
{ "Alerts": [ "alert1", "alert2" ] }
```

アウトプット:

```
{ "Alerts": "alert1" }
```

```
{ "Alerts": "alert2" }
```

### 例：トップレベルの配列の分割

エントリー全体が配列になっていることがあります：

```
[ {"foo": "bar"}, {"x": "y"} ]
```

これを分割するには、次のような定義を用います：

```
[preprocessor "json"]
	Type=jsonarraysplit
```

Extractionパラメータを設定しないでおくと、モジュールはエントリ全体を配列として扱い、次の2つの出力エントリが得られます：

```
{"foo": "bar"}
```

```
{"x": "y"}
```

## JSONフィールドフィルタリングプリプロセッサ

このプリプロセッサは、入力データをJSONオブジェクトとして解析し、指定されたフィールドを抽出して、許容値のリストと比較します。許容される値のリストは、ディスク上のファイルで、1行に1つの値を指定します。

リストに一致するフィールドを持つエントリのみを*通過*させるか、リストに一致するエントリを*ドロップ*するように設定できます（ホワイトリストまたはブラックリスト）。また、複数のフィールドに対してフィルタリングするように設定することもできます。その場合、*すべての*フィールドが一致しなければならない（論理的AND）か、*少なくとも1つの*フィールドが一致しなければならない（論理的OR）かのいずれかを要求します。

このプリプロセッサは、一般的なデータのファイアーホースを遅いネットワークリンクで送信する前に絞り込むのに特に役立ちます。

JSONフィールドフィルタリングプリプロセッサのタイプは `jsonfilter` です。

### サポートされるオプション

* `Field-Filter` (文字列、必須)。対象となるJSONフィールドの名前と、照合する値を含むファイルのパスの2つを指定します。たとえば、`Field-Filter=ComputerName,/opt/gravwell/etc/computernames.txt` と指定すれば、「ComputerName」というフィールドを抽出して、`/opt/gravwell/etc/computernames.txt` の中の値と比較することができます。Field-Filter`オプションは複数回指定することができ、複数のフィールドに対してフィルタリングを行うことができます。
* `Match-Logic` (文字列、オプション)。このパラメータは、複数のフィールドに対してフィルタリングする際に使用する論理演算を指定します。and "に設定すると、指定されたリストに対して指定されたフィールドがすべて一致した場合にのみ、エントリーが一致したとみなされます。or "に設定すると、*任意の*フィールドがマッチしたときにエントリーがマッチしたとみなされます。
* `Match-Action` (文字列、オプション): フィールドが与えられたリストにマッチしたエントリーに対して取るべきオプションを指定します。省略した場合のデフォルトは "pass "です。pass "に設定すると、マッチしたエントリーはインデクサへの通過が許可されます（ホワイトリスト）。drop "に設定すると、マッチしたエントリーはドロップされます（ブラックリスト化）。

Match-Logic "パラメータは、複数の "Field-Filter "が指定されている場合にのみ必要です。

注意: 設定でフィールドが指定されていても、エントリーに存在しない場合、プリプロセッサは、そのフィールドが存在するが何にもマッチしないかのように*エントリーを扱います。したがって、ホワイトリストにマッチするフィールドを持つエントリのみを通過させるようにプリプロセッサを設定した場合、フィールドの1つを欠くエントリはドロップされます。

### 一般的な使用例

json フィールドフィルタリングプリプロセッサは、エントリ内のフィールドに基づいてエントリをダウンセレクトすることができます。 これにより、データフローにブラックリストやホワイトリストを作成して、データがストレージに入るか入らないかを確認することができます。

### 例：単純なホワイトリスト化

例えば、企業内で発生している事象の詳細を示す毎秒数千のイベントを送信しているエンドポイント監視ソリューションがあるとします。イベントの量が多いため、特定の深刻度のイベントのみをインデックス化したいと考えるかもしれません。幸い、イベントにはSeverityフィールドが含まれています。

```
{ "EventID": 1337, "Severity": 8, "System": "email-server-01.example.org", [...] }
```

Severityフィールドは0から9までありますが、Severityが6以上のイベントのみを通過させたいと考えています。そこで、インゲスターの設定ファイルに次のように追加します：

```
[preprocessor "severity"]
	Type=jsonfilter
	Match-Action=pass
	Field-Filter=Severity,/opt/gravwell/etc/severity-list.txt
```

を追加し、適切なデータ入力に対して `Preprocessor=severity` を設定します（例：Simple Relay を使用している場合）：

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
```

最後に、`/opt/gravwell/etc/severity-list.txt`を作成し、許容できるSeverity値のリストを1行ごとに入力します：

```
6
7
8
9
```

インジェスターの再起動後、インジェスターは各エントリから `Severity` フィールドを抽出し、その結果得られた値をファイルに記載されている値と比較します。値がファイル内の行と一致した場合、そのエントリはインデクサに送られます。それ以外の場合は、ドロップされます。

### 例：ブラックリスト化

前述の例に基づいて、エンドポイント監視システムが特定のシステムから *ロット* の高感度の誤検出を生成していることに気づくかもしれません。例えば、`EventID` フィールドが 219, 220, 1338 に設定され、`System` フィールドが "webserver-prod.example.org" と "webserver-dev.example.org" に設定されたイベントは、常に誤検出であると判断できます。別のプリプロセッサを定義して、インデクサに送信される前にこれらのエントリを取り除くことができます。

```
[preprocessor "falsepositives"]
	Type=jsonfilter
	Match-Action=drop
	Match-Logic=and
	Field-Filter=EventID,/opt/gravwell/etc/eventID-blacklist.txt
	Field-Filter=System,/opt/gravwell/etc/system-blacklist.txt
```

このプリプロセッサをデータ入力構成に既存のプリプロセッサの後に追加すると、インゲスターは2つのフィルタを順番に適用します：

```
[Listener "endpoint_monitoring"]
	Bind-String="0.0.0.0:7700
	Tag-Name=endpoint
	Preprocessor=severity
	Preprocessor=falsepositives
```

最後に、`/opt/gravwell/etc/eventID-blacklist.txt`を作成します：

```
219
220
1338
```

と、`/opt/gravwell/etc/system-blacklist.txt`の2つを用意しました：

```
webserver-prod.example.org
webserver-dev.example.org
```

この新しいプリプロセッサは、すべてのエントリから `EventID`フィールドと` System`フィールドを抽出し、最初のフィルタを通過させます。 次に、それらをファイル内の値と比較します。 `Match-Logic = and`を設定したため、エントリと見なされます` Match-Action = drop`を設定したため、両方のフィールドに一致するエントリはすべて削除されます。 したがって、EventID = 220およびSystem = webserver-devのエントリ。 .example.orgは削除されますが、EventID = 220およびSystem = email-server-01.example.orgの1つは削除されません。

## 正規表現ルータープリプロセッサ

regexルータプリプロセッサは、エントリの内容に基づいて異なるタグにエントリをルーティングするための柔軟なツールです。設定では、[named capturing group](https://www.regular-expressions.info/named.html)を含む正規表現を指定し、その内容をユーザー定義のルーティングルールに対してテストします。

Regex Routerのプリプロセッサタイプは `regexrouter` です。

### サポートされるオプション

* `Regex` (文字列、必須)。このパラメータでは、入力されるエントリーに適用する正規表現を指定します。例えば、`(?P<app>.+)` は `Route-Extraction` パラメータと一緒に使用されます。
* `Route-Extraction` (文字列, 必須): Route-Extraction` (string, required): このパラメータは、`Regex`パラメータで指定したキャプチャグループの名前を指定します。
* `Route` (文字列、必須)。少なくとも1つの `Route` 定義が必要です。例えば、`Route=sshd:sshlogtag`のように、コロンで区切られた2つの文字列で構成されます。最初の文字列('sshd')は、正規表現で抽出された値と照合され、2番目の文字列は、一致したエントリがルーティングされるべきタグの名前を定義します。2番目の文字列を空白にすると、最初の文字列にマッチしたエントリは *ドロップ* されます。
* `Drop-Misses` (boolean, optional): デフォルトでは、正規表現にマッチしないエントリーは、そのまま通過します。Drop-Misses`をtrueに設定すると、インゲスターは、1) 正規表現にマッチしない、または、2) 正規表現にマッチするが、指定されたルートのいずれにもマッチしないエントリをドロップします。

### 例：Appフィールドの値に基づくタグへのルーティング

このプリプロセッサの使い方を説明するために、多くのシステムがSimple Relayインゲスターにsyslogエントリを送信している状況を考えてみましょう。sshd のログを `sshlog` という名前の別のタグに分離したいとします。受信する sshd のログは、古いスタイルの BSD syslog フォーマット (RFC3164) になっています。

```
<29>1 Nov 26 11:26:36 localhost sshd[11358]: Failed password for invalid user administrator from 202.198.122.184 port 49828 ssh2
```

正規表現を試してみたところ、RFC3164のログからアプリケーション名（例：sshd）を「app」という名前のキャプチャーグループに抽出するには、以下のような正規表現が妥当であることがわかりました。

```
^(<\d+>)?\d?\s?\S+ \d+ \S+ \S+ (?P<app>[^\s\[]+)(\[\d+\])?:
```

この正規表現をプリプロセッサの定義に適用すると、以下のようになります：

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
	# Regex: <pri>version Month Day Time Host App[pid]
	Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog
```

プリプロセッサは、正規表現を定義してから、`Route-Extraction`パラメータでキャプチャグループの「app」を呼び出していることに注意してください。そして、`Route=ssh:sshlog`の定義を使用して、アプリケーション名が「sshd」にマッチするエントリを「sshlog」というタグにルーティングするように指定しています。必要に応じて、追加の `Route` パラメータを定義することもできます。例えば、`Route=apache:apachelog` などです。

上記の構成では、sshd からのログは "sshlog" タグに送られ、他のすべてのログは "syslog" タグに直接送られます。Route`の指定を追加することで、同じようなフォーマットのsyslogエントリから他のアプリケーションを抽出することができますが、以下のようにRFC 5424フォーマットのログが混ざっていたとしたらどうでしょうか。

```
<101>1 2019-11-26T13:24:56.632535-07:00 web01.example.org webservice 21581 - [useragent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3191.0 Safari/537.36"] GET /
```

すでにある正規表現では、アプリケーション名（"webservice"）を正しく抽出できませんが、*2つ目の*プリプロセッサを定義し、既存のプリプロセッサの後にプリプロセッサチェーンに入れることができます。

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=bsdrouter
	Preprocessor=rfc5424router

[preprocessor "bsdrouter"]
        Type = regexrouter
        Drop-Misses=false
	# Regex: <pri>version Month Day Time Host App[pid]
	Regex="^(<\\d+>)?\\d?\\s?\\S+ \\d+ \\S+ \\S+ (?P<app>[^\\s\\[]+)(\\[\\d+\\])?:"
        Route-Extraction=app
        Route=sshd:sshlog

[preprocessor "rfc5424router"]
	Type=regexrouter
	Drop-Misses=false
	# Regex: <pri>version Date Host App
	Regex="^<\\d+>\\d? \\S+ \\S+ (?P<app>\\S+)"
	Route-Extraction=app
	Route=webservice:weblog
	Route=apache:weblog
	Route=postfix:		# drop
```

この新しいプリプロセッサの定義では、"webservice "と "apache "というアプリケーションのためのルートを定義し、両方とも "weblog "タグに送っていることに注意してください。また、"postfix "アプリケーションからのログを *ドロップ* するように指定していることにも注意してください。これはおそらく、これらのログがすでに他のソースから取り込まれているためです。

## ソースルータープリプロセッサ

ソースルータプリプロセッサは、エントリーのSRCフィールドに基づいて、エントリーを異なるタグにルーティングすることができます。通常、SRCフィールドはエントリーの起点のIPアドレスになります。例えば、Simple Relayに送信されたsyslogメッセージを作成したシステムなどです。

ソースルータのプリプロセッサタイプは `srcrouter` です。

### サポートされるオプション

* `Route` (string, optional): `Route` は、SRC フィールドの値とタグのマッピングをコロンで区切って定義します。たとえば、`Route=192.168.0.1:server-logs` とすると、SRC=192.168.0.1 のすべてのエントリを「server-logs」タグに送信します。複数の`Route`パラメータを指定することができます。タグを空白にすると (`Route=192.168.0.1:`)、プリプロセッサは代わりにマッチするすべてのエントリをドロップします。
* `Route-File` (string, optional): `Route-File` には、`192.168.0.1:server-logs` のように、改行で区切られたルート指定を含むファイルへのパスを指定します。
* `Drop-Misses` (boolean, optional): デフォルトでは、定義されたルートのどれにもマッチしないエントリーは、修正されずに通過します。Drop-Misses` を true に設定すると、ルート定義に明示的にマッチしないエントリはすべてドロップされます。

Route-File` を使用しない限り、少なくとも 1 つの `Route` 定義が必要です。

ルートは単一のIPアドレスか、適切に形成されたCIDR仕様のいずれかで、IPv4とIPv6の両方がサポートされています。

### 例：インラインでのルート定義

以下のスニペットは、ソースルータのプリプロセッサを使用し、インラインでルートを定義したSimple Relayインゲスターの設定の一部を示しています。10.0.0.1からのエントリーには "internal-syslog "というタグが付けられ、7.82.33.4からのエントリーには "external-syslog "というタグが付けられ、それ以外のエントリーにはデフォルトのタグ "syslog "が付けられます。SRC=3.3.3.3のエントリはすべて削除されます。

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=srcroute

[preprocessor "srcroute"]
        Type = srcrouter
        Route=10.0.0.0/24:internal-syslog
        Route=7.82.33.4:external-syslog
        Route=3.3.3.3:
        Route=DEAD::BEEF:external-syslog
        Route=FEED:FEBE::0/64:external-syslog
```

### 例：ファイルベースの定義

以下のスニペットは、ファイルで定義されたルートを持つソースルータのプリプロセッサを使用するSimple Relayインゲスターの設定の一部です。

```
[Listener "syslog"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=syslog
        Preprocessor=srcroute

[preprocessor "srcroute"]
        Type = srcrouter
        Route-File=/opt/gravwell/etc/syslog-routes
```

以下の内容が `/opt/gravwell/etc/syslog-routes` に書き込まれます：

```
10.0.0.0/24:internal-syslog
7.82.33.4:external-syslog
3.3.3.3:
```

## 正規表現タイムスタンプ抽出プリプロセッサ

摂取者は通常、有効なタイムスタンプであると思われる最初のものを探してそれを解析することにより、エントリからタイムスタンプを抽出しようとします。 タイムスタンプを解析するための追加のingester構成ルール（検索する特定のタイムスタンプ形式の指定など）と組み合わせると、通常、適切なタイムスタンプを適切に抽出するにはこれで十分ですが、一部のデータソースはこれらの単純な方法に反する場合があります。 ネットワークデバイスがsyslogでラップされたCSV形式のイベントログを送信する可能性がある状況を考えてみましょう。これはGravwellで見られた状況です。

正規表現のタイムスタンプ抽出のプリプロセッサタイプは、 `regextimestamp`です。

### サポートされているオプション

* `Regex` (文字列、必須)。このパラメータは、入力されたエントリーに適用される正規表現を指定します。例えば、`(?P<timestamp>.+)` は、`TS-Match-Name` パラメータと一緒に使用されます。
* `TS-Match-Name` (文字列、必須)。このパラメータは、抽出されたタイムスタンプを含む、`Regex`パラメータから名前が付けられたキャプチャグループの名前を与えます。
* `Timestamp-Format-Override` (string, optional): これを使って、別のタイムスタンプの解析フォーマットを指定することができます。利用可能なタイムフォーマットは
	- AnsiC
	- Unix
	- Ruby
	- RFC822
	- RFC822Z
	- RFC850
	- RFC1123
	- RFC1123Z
	- RFC3339
	- RFC3339Nano
	- Apache
	- ApacheNoTz
	- Syslog
	- SyslogFile
	- SyslogFileTZ
	- DPKG
	- NGINX
	- UnixMilli
	- ZonelessRFC3339
	- SyslogVariant
	- UnpaddedDateTime
	- UnpaddedMilliDateTime
	- UK
	- Gravwell
	- LDAP
	- UnixSeconds
	- UnixMs
	- UnixNano

	タイムスタンプのフォーマットの中には、重複する値を持つものがあります（例えば、LDAPとUnixNanoは同じ桁数のタイムスタンプを生成できます）。`Timestamp-Format-Override`が使用されていない場合、プリプロセッサは上記の順序でタイムスタンプを導き出そうとします。このリストの他のものと衝突する可能性のあるタイムスタンプフォーマットを使用する場合は、常に `Timestamp-Format-Override` を使用してください。

* `Timezone-Override` (文字列、オプション): 抽出されたタイムスタンプにタイムゾーンが含まれていない場合は、ここで指定されたタイムゾーンが適用されます。例: `US/Pacific`、`Europe/Rome`、 `Cuba`。
* `Assume-Local-Timezone` (boolean, optional): このオプションは、タイムゾーンが含まれていない場合、タイムスタンプがローカルのタイムゾーンであると仮定するようにプリプロセッサに指示します。これは `Timezone-Override` パラメータとは相互に排他的です。


### 一般的な使用例

多くのデータストリームには、複数のタイムスタンプや、タイムスタンプとして解釈されやすい値が含まれていることがあります。 regextimestamp プリプロセッサを使用すると、 timegrinder にログストリーム内の特定のタイムスタンプを検査させることができます。 良い例は、それ自身のタイムスタンプを含むが、そのタイムスタンプをsyslog APIに中継しないアプリケーションを使用して、syslog経由で転送されるログストリームです。 syslogのラッパーは、よくできたタイムスタンプを持っていますが、実際のデータストリームでは、正確なタイムスタンプのためにいくつかの内部フィールドを使用する必要があるかもしれません。

### 例：ラップされたSyslogデータ

```
Nov 25 15:09:17 webserver alerts[1923]: Nov 25 14:55:34,GET,10.1.3.4,/info.php
```

インジェストされたエントリーのTSフィールドの内側のタイムスタンプ、"Nov 25 14:55:34 "を抽出したいと思います。これは行頭のsyslogタイムスタンプと同じフォーマットを使用しているので、巧妙なタイムスタンプフォーマットルールでは抽出できません。しかし、正規表現タイムスタンププリプロセッサを使って抽出することができます。名前付きのサブマッチで必要なタイムスタンプを捕らえる正規表現を指定することで、エントリーのどこからでもタイムスタンプを抽出することができます。このエントリーでは、正規表現 `\S+\s+\S+[\d+\]: (?<timestamp>.+),`という正規表現を使えば、目的のタイムスタンプを適切に抽出することができます。

この設定を利用して、上の例で示したタイムスタンプを抽出することができます：

```
[Preprocessor "ts"]
	Type=regextimestamp
	Regex="\S+\s+\S+\[\d+\]: (?P<timestamp>.+),"
	TS-Match-Name=timestamp
	Timezone-Override=US/Pacific
```

## 正規表現抽出プリプロセッサ

トランスポートバスでは、データストリームに実際のイベントには関係のない追加のメタデータを付加することがよくあります。Syslogはその好例で、Syslogヘッダーが基礎となるデータに価値を与えない、あるいはデータの分析を単に複雑にする可能性があります。regexextractorプリプロセッサは、複数のフィールドを抽出し、新しい構造のフォーマットに変換する正規表現を宣言することができます。

正規表現抽出プリプロセッサは、名前付きの正規表現抽出フィールドとテンプレートを使ってデータを抽出し、それを出力レコードに再構成します。 出力テンプレートには静的な値を含めることができ、必要に応じて出力データを完全に再構築します。

テンプレートは、bashと同様にフィールド定義を使って抽出された値を名前で参照します。 たとえば、正規表現で `foo` という名前のフィールドを抽出した場合、その抽出値を `${foo}` としてテンプレートに挿入することができます。また、テンプレートは以下の特殊キーをサポートしています。

* `${_SRC_}`, これは現在のエントリのSRCフィールドに置き換えられます。
* `${_DATA_}`, これは現在のエントリの文字列形式のDataフィールドで置き換えられます。
* `${_TS_}`, これは現在のエントリの文字列形式の TS (timestamp) フィールドに置き換えられます。

正規表現抽出プリプロセッサのTypeは `regexextract` です。

### サポートされるオプション

* Passthrough-Misses (boolean, optional)：このパラメータは、正規表現がマッチしない場合に、プリプロセッサがレコードを変更せずに通過させるかどうかを指定します。
* Regex (文字列、必須)：このパラメータは、抽出のための正規表現を定義します。
* Template (string, required): テンプレートを定義します。このパラメータは、レコードの出力形式を定義します。

### 一般的な使用例

正規表現プリプロセッサは、データストリームから必要のないヘッダーを取り除くために最も一般的に使用されますが、データを処理しやすい形式に変換するためにも使用できます。

#### 例：Syslog ヘッダーの除去

次のようなレコードがあった場合、syslogヘッダーを除去してJSON blobだけを出荷したいとします。

```
<30>1 2020-03-20T15:35:20Z webserver.example.com loginapp 4961 - - {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}
```

syslogメッセージはよく構造化されたJSON blobを含んでいますが、syslogトランスポートは、必ずしもレコードを強化しない追加のメタデータを追加します。 Regexエクストラクタを使って必要なデータを取り出し、使いやすいレコードに整形することができます。

ここでは、正規表現抽出器を使用してデータフィールドとホスト名を抽出し、テンプレートを使用してホストが挿入された新しいJSON blobを構築します。


```
[Listener "logins"]
	Bind-String="0.0.0.0:7777"
	Preprocessor=loginapp

[Preprocessor "loginapp"]
	Type=regexextract
	Regex="\\S+ (?P<host>\\S+) \\d+ \\S+ \\S+ (?P<data>\\{.+\\})$"
	Template="{\"host\": \"${host}\", \"data\": ${data}}"
```

注意：正規表現では、文字セットを記述するためにバックスラッシュを使用することがよくありますが、これらのバックスラッシュはエスケープする必要があります。

結果のデータは：

```
{"host": "loginapp", "data": {"user": "bob", "action": "login", "result": "success", "UID": 123, "ts": "2020-03-20T15:35:20Z"}}
```

注：テンプレートでは、複数のフィールドの定数値を指定できます。 抽出されたフィールドは複数回挿入することができます。

## フォワーディングプリプロセッサ

フォワーディングプリプロセッサは、ログストリームを分割して別のエンドポイントに転送するために使用します。 これは、追加のロギングツールを立ち上げるときや、外部のアーカイブプロセッサにデータを供給するときに非常に便利です。 フォワーディング」プリプロセッサは、フォーク型プリプロセッサです。 これは、データストリームを変更せず、データを追加のエンドポイントに転送するだけであることを意味します。

デフォルトでは、フォワーディングプリプロセッサはブロッキングです。つまり、TCPやTLSなどのステートフルな接続を使用してフォワーディングエンドポイントを指定しても、そのエンドポイントが稼働していなかったり、データを受け入れることができなかったりすると、取り込みをブロックします。 この動作は、`Non-Blocking` フラグを使用するか、UDP プロトコルを使用することで変更できます。

フォワーディングプリプロセッサは、転送されるデータストリームを削減したり、正確に指定するためのいくつかのフィルタメカニズムもサポートしています。 ストリームは、エントリータグ、ソース、または実際のログデータ上で動作する正規表現を使用してフィルタリングできます。 各フィルター仕様は、複数回指定してORパターンを作ることができます。

複数のフォワーディングプリプロセッサを指定することができ、特定のログストリームを複数のエンドポイントにフォワードすることができます。

フォワーディングプリプロセッサのタイプは `forwarder` です。

### サポートされるオプション

* `Target` (string, required): 転送されるデータのエンドポイントです。 unix` プロトコルを使用する場合を除き、ホストとポートのペアでなければなりません。
* `Protocol` (string, 必須): データを転送する際に使用するプロトコルです。Protocol` (string, required): データを転送するときに使用するプロトコルです。 オプションは `tcp`, `udp`, `unix`, および `tls` です。
* `Delimiter` (文字列、オプション)：`Raw` 出力フォーマットでデータを送信するときに使用するオプションのデリミタです。
* `Format` (string, optional): データを送信する出力フォーマットです。 オプションは `raw`, `json`, `syslog` です。
* `Tag` (string, optional, 複数可): イベントをフィルタリングするためのタグ。 複数指定は OR を意味する。
* `Regex` (文字列、オプション、複数可): イベントのフィルタリングに使用される正規表現。 複数の指定は OR を意味します。
* `Source` (文字列、オプション、複数可): イベントのフィルタリングに使用されるIPまたはCIDRの仕様。 複数の指定は OR を意味します。
* `Timeout` (unsigned int, オプション, 秒単位で指定): フォワーダへの接続と書き込みの試行のタイムアウト
* `Buffer` (unsigned int, オプション): フォワーダがデータ送信を試みる際にバッファリングできるイベント数を指定します。
* `Non-Blocking` (boolean, optional): Trueを指定すると、フォワーダーはデータを転送するために最善の努力をしますが、取り込みをブロックしません。
* `Insecure-Skip-TLS-Verify` (boolean, optional): TLS ベースの接続で、無効な TLS 証明書を無視できることを指定します。

### 例 特定のホストセットからのsyslogの転送

この例では、SimpleRelayのインゲスターを使ってsyslogメッセージを取り込み、生のまま別のエンドポイントに転送します。 `forward`プリプロセッサの`Source`フィルタを使用して、`192.168.0.1`のIPまたは`192.168.1.0/24`のサブネットのいずれかからのsourceタグを持つログのみを転送しています。 ログはそれぞれの間に改行を入れた元のフォーマットで転送されます。

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=forwardprivnet

[Preprocessor "forwardprivnet"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false
```

### 例：特定のWindowsイベントログの転送

この例では、潜在的に多くのダウンストリーム・インゲスターからのデータ・ストリームを転送するためにフェデレーターを使用しています。 この例では、Federatorを使用して、多くのダウンストリームインゲスターからのデータストリームを転送しています。 ここでは`syslog`フォーマットを使用しており、RFC5424ヘッダーとsyslogメッセージのボディにデータを入れてエンドポイントにデータを送信していることに注意してください。 syslog形式で転送されたデータは、Hostnameに`gravwell`、AppnameにエントリーTAGを指定しています。

`Tag`フィルターは、`windows`または`sysmon`タグを使用しているエントリーのみを転送したいことを指定します。

`Regex`フィルターは、特定のチャンネルとイベントIDの組み合わせからのイベントデータのみを取得するために使用されます。 例えば、セキュリティプロバイダからのログインイベントや、sysmonプロバイダからの実行イベントなどです。
```
[IngestListener "enclaveA"]
	Ingest-Secret = IngestSuperSecrets
	Cleartext-Bind = 0.0.0.0:4023
	Tags=win*
	Preprocessor=forwardloginsysmon
	Preprocessor=forwardprivnet

[Preprocessor "forwardloginsysmon"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="syslog"
	Buffer=128
	Tag=windows
	Tag=sysmon
	Regex="Microsoft-Windows-Sysmon.+>(1|22)</EventID>"
	Regex="Microsoft-Windows-Security-Auditing.+>(4624|4625|4626)</EventID>"
	Non-Blocking=false

[Preprocessor "forwardsyslog"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Tag=syslog
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false

```


### 例 複数のホストにログを転送する

この例では、Gravwell Federatorを使って、ログのサブセットを異なるフォーマットで異なるエンドポイントに転送します。 フォワーダープリプロセッサは他のプリプロセッサと同じようにスタックできるので、独自のフィルタ、エンドポイント、フォーマットを持つ複数のフォワーダープリプロセッサを指定することができます。

```
[IngestListener "enclaveB"]
	Ingest-Secret = IngestSuperSecrets
	Cleartext-Bind = 0.0.0.0:4123
	Tags=win*
	Preprocessor=forwardloginsysmon

[Preprocessor "forwardloginsysmon"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="syslog"
	Buffer=128
	Tag=windows
	Tag=sysmon
	Regex="Microsoft-Windows-Sysmon.+>(1|22)</EventID>"
	Regex="Microsoft-Windows-Security-Auditing.+>(4624|4625|4626)</EventID>"
	Non-Blocking=false


```

## Gravwell転送プリプロセッサ

Gravwell転送プロセッサは、複数のGravwellインスタンスにエントリを複製できる完全なGravwell muxerの作成を可能にします。 このプリプロセッサは、テストや、特定のGravwellデータストリームを別のインデクサに複製する必要がある場合に役立ちます。 Gravwell転送プリプロセッサは、パッケージ化されたインジェスターと同じ構成構造を利用して、インデクサー、インジェスト・シークレット、およびキャッシュ・コントロールを指定します。 Gravwellフォワーディングプリプロセッサはブロッキングプリプロセッサであり、これはローカルキャッシュを有効にしていない場合、プリプロセッサが指定されたインデクサにエントリを転送できない場合、インジェストパイプラインをブロックすることができることを意味します。

Gravwell Forwarding プリプロセッサの Type は `gravwellforwarder` です。

### サポートされているオプション

すべてのGravwellインジェスターオプションの詳細については、[グローバルコンフィギュレーションパラメータ](#!ingesters/ingesters.md#Global_Configuration_Parameters)セクションを参照してください。 ほとんどのインゲスターのグローバル設定オプションはGravwell Forwarderのプリプロセッサでサポートされています。

### 例：フェデレーターでデータを複製する

この例では、すべてのエントリを2つ目のクラスタに複製する完全なFederatorの構成を指定します。

注意: フォワーディングプリプロセッサで`always`キャッシュを有効にして、通常のインジェストパスをブロックしないようにしています。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Verify-Remote-Certificates = true
Cleartext-Backend-Target=172.19.0.2:4023 #example of adding a cleartext connection
Log-Level=INFO

[IngestListener "enclaveA"]
	Ingest-Secret = CustomSecrets
	Cleartext-Bind = 0.0.0.0:4423
	Tags=windows
	Tags=syslog-*
	Preprocessor=dup

[Preprocessor "dup"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets
	Connection-Timeout = 0
	Cleartext-Backend-Target=172.19.0.4:4023 #indexer1
	Cleartext-Backend-Target=172.19.0.5:4023 #indexer2 (cluster config)
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup.cache # must be a unique path
	Max-Ingest-Cache=1024 #Limit forwarder disk usage
```

### 例：重複フォワーダのスタック

この例では、完全なFederator構成と複数のGravwellプリプロセッサを指定して、1つのストリームのエントリを複数のGravwellクラスタに複製できるようにします。

注意：プリプロセッサの制御ロジックは、同じクラスタに複数回転送していないかどうかはチェックしません。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Verify-Remote-Certificates = true
Cleartext-Backend-Target=172.19.0.2:4023 #example of adding a cleartext connection
Log-Level=INFO

[IngestListener "enclaveA"]
	Ingest-Secret = CustomSecrets
	Cleartext-Bind = 0.0.0.0:4423
	Tags=windows
	Tags=syslog-*
	Preprocessor=dup1
	Preprocessor=dup2
	Preprocessor=dup3

[Preprocessor "dup1"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets1
	Cleartext-Backend-Target=172.19.0.101:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup1.cache
	Max-Ingest-Cache=1024

[Preprocessor "dup2"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets2
	Cleartext-Backend-Target=172.19.0.102:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup2.cache
	Max-Ingest-Cache=1024

[Preprocessor "dup3"]
	Type=GravwellForwarder
	Ingest-Secret = IngestSecrets3
	Cleartext-Backend-Target=172.19.0.103:4023
	Ingest-Cache-Path=/opt/gravwell/cache/federator_dup3.cache
	Max-Ingest-Cache=1024
```

## ドロッププリプロセッサ

ドロッププリプロセッサは、その名が示す通り、インジェストパイプラインからエントリを単純にドロップし、効果的に破棄します。

このプロセッサは、インジェストストリームが主に他のプリプロセッサのセットによって処理される場合に便利です。 例えば、フォワーダープリプロセッサを使ってリモートシステムにデータを送信したいが、Gravwellのインデクサにはアップストリームでインジェストしたくない場合、最終的に`drop`プリプロセッサを追加することで、見たエントリをすべて単純に廃棄することができます。

### サポートされているオプション

なし。

### 例：すべてを廃棄

この例では、シンプルなRelayリスナーのすべてのエントリを廃棄するだけの単一のプリプロセッサ `drop` を使用しています：

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=dropit

[Preprocessor "dropit"]
	Type=Drop               
```

### 例：エントリの転送とドロップ

この例では、TCPフォワーダーを使ってエントリーを転送し、ドロップします。

```
[Listener "default"]              
	Bind-String="0.0.0.0:601"
	Reader-Type=rfc5424
	Tag-Name=syslog
	Preprocessor=forwardprivnet
	Preprocessor=dropit

[Preprocessor "forwardprivnet"]
	Type=Forwarder               
	Protocol=tcp
	Target="172.17.0.3:601"
	Format="raw"
	Delimiter="\n"
	Buffer=128
	Source=192.168.0.1
	Source=192.168.1.0/24
	Non-Blocking=false

[Preprocessor "dropit"]
	Type=Drop               
```

## CiscoISEプリプロセッサ

Cisco ISEプリプロセッサは、Cisco ISE ログのフォーマットやトランスポートを解析し、適応させるために設計されています。 詳細については、[Cisco Introduction to ISE Syslogs](https://www.cisco.com/c/en/us/td/docs/security/ise/syslog/Cisco_ISE_Syslogs/m_IntrotoSyslogs.pdf)を参照してください。

Cisco ISE プリプロセッサは `cisco_ise` という名前で、マルチパート メッセージの再構築、Gravwell や最新の syslog システムに適した形式へのメッセージの再フォーマット、不要なメッセージ ペアのフィルタリング、冗長なメッセージ ヘッダの削除などの機能をサポートしています。

### 属性のフィルタリングとフォーマット

Cisco ISEのログシステムは、1つのメッセージを複数のsyslogメッセージに分割するように設計されています。 Gravwellはsyslogの最大メッセージサイズをはるかに超えるメッセージを受け入れますが、Cisco ISEメッセージの複数のターゲットをサポートしている場合、マルチパートメッセージを有効にする必要があります。 Ciscoデバイスでマルチパートメッセージを無効にして、Gravwellに大きなペイロードを処理させる方がはるかに効率的です。

### サポートされているオプション

* `Passthrough-Misses` (boolean, optional)：trueに設定すると、プリプロセッサは有効なISEメッセージを抽出することができなかったエントリを通過させます。デフォルトでは、これらのエントリーはドロップされます。
* `Enable-Multipart-Reassembly` (boolean, optional): trueに設定すると、プリプロセッサはCiscoリモートメッセージヘッダを含むメッセージの再組み立てを試みます。
* `Max-Multipart-Buffer` (uint64, オプション): このバッファを超えると、最も古い部分的に再構成されたメッセージがGravwellに送信されます。 デフォルトのバッファサイズは8MBです。
* `Max-Multipart-Latency` (string, optional): 部分的に再構成されたマルチパートメッセージが送信される前に保持される最大時間帯を指定します。 時間間隔は `ms` や `s` の値で指定します。
* `Output-Format` (string, optional): デフォルトのフォーマットは `raw` で、他のオプションは `json` と `cef` です。
* `Attribute-Drop-Filter` (string array, optional): 出力から削除されるメッセージ内の属性に対するマッチングに使用できるグロブイングパターンを指定します。 引数は [Unix Glob patterns](https://en.wikipedia.org/wiki/Glob_(programming)) でなければならず、多くのパターンを指定することができます。 属性ドロップフィルターは `raw` 出力フォーマットとは互換性がありません。
* `Attribute-Strip-Header` (boolean, optional): ネストされた名前やタイプ情報を持つ属性のヘッダ値を除去するように指定します。 これは、不適切に形成された属性値を整理するのに非常に便利です。


### 構成例

以下の `cisco_ise` プリプロセッサの構成は、マルチパートのメッセージを再構成し、不要な `Step` 属性を削除して、出力メッセージを CEF 形式に変換するように設計されています。 また、cisco属性のヘッダを除去します。

```
[preprocessor "iseCEF"]
    Type=cisco_ise
    Passthrough-Misses=false #if its malformed just drop it
    Enable-Multipart-Reassembly=true
    Attribute-Drop-Filters="Step*"
    Attribute-Strip-Header=true
    Output-Format=cef
```

この設定を使った出力メッセージの例は以下の通りです：

```
CEF:0|CISCO|ISE_DEVICE|||Passed-Authentication|NOTICE| sequence=1 ode=5200 class=Passed-Authentication text=Authentication succeeded ConfigVersionId=44 DeviceIPAddress=8.8.8.8 DestinationIPAddress=1.2.3.4 DestinationPort=1645 UserName=user@company.com Protocol=Radius RequestLatency=10301 audit-session-id=0a700e191cff70005fbbf63f
```

次の `cisco_ise` プリプロセッサの設定は、同様の結果をもたらしますが、2つの属性フィルタを含み、出力メッセージをJSONに再フォーマットします。

```
[preprocessor "iseCEF"]
    Type=cisco_ise
    Passthrough-Misses=false #if its malformed just drop it
    Enable-Multipart-Reassembly=true
    Attribute-Drop-Filters="Step*"
    Attribute-Strip-Header=true
    Output-Format=json
```

この設定を使った出力メッセージの例は以下の通りです：

```
{
  "TS":"2020-11-23T12:50:01.926-05:00",
  "Sequence":1,
  "ODE":"5200",
  "Severity":"NOTICE",
  "Class":"Passed-Authentication",
   "Text":"Authentication succeeded",
   "Attributes":{
     "AcsSessionID":"ISE_DEVICE/384429556/212087299",
     "AuthenticationIdentityStore":"AzureBackup",
     "AuthenticationMethod":"PAP_ASCII",
     "AuthenticationStatus":"AuthenticationPassed",
     "audit-session-id":"0a700e191cff70005fbbf63f",
     "device-mac":"00-0c-29-74-9d-e8",
     "device-platform":"win",
     "device-platform-version":"10.0.17134",
  }
}
```
