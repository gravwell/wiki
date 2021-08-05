# syslog

syslog プロセッサは、[Simple Relay インジェスター](#!ingesters/ingesters.md) でインジェストされた [RFC 5424-format](https://tools.ietf.org/html/rfc5424) の syslog メッセージからフィールドを抽出します（リスナーに Keep-Priority フラグを設定しておかないと動作しません）。

## サポートされているオプション

* `-e`: 「-e」オプションは、syslog モジュールが列挙値に対して操作することを指定します。 列挙値に対する操作は、上流のモジュールを使って syslog レコードを抽出した場合に便利です。 例えば、生の PCAP から syslog レコードを抽出し、そのレコードを syslog モジュールに渡すことができます。

## 処理演算子

各 syslog フィールドは、高速フィルタとして機能する演算子のセットをサポートしています。 各演算子がサポートするフィルタは、フィールドのデータタイプによって決まります。

| 演算子 | 名称 | 意味 |
|----------|------|-------------|
| == | 等しい | フィールドは等しい
| != | 等しくない | フィールドは等しくない
| < | 小なり | フィールドはその値より小さい
| > | 大なり | フィールドはその値より大きい
| <= | 小なりイコール | フィールドはその値以下である
| >= | 大なりイコール | フィールドはその値以上である

## データフィールド

syslog モジュールは、RFC 5424 形式の syslog レコードから個々のフィールドを抽出します。これは、左から右に向かって解析しようとする最善の努力をします。つまり、あるフィールドが欠けている場合には、その右にあるフィールドだけが与えられたレコードで利用可能になります。

| フィールド | 説明 | サポートされている演算子 | 例 |
|-------|-------------|---------------------|---------|
| Facility | メッセージの発信元となるファシリティを示す数値コード | > < <= >= == != | Facility == 0
| Severity | メッセージの深刻度を示す数値コード。0が最も深刻で、7が最も深刻でないことを示します。 | > < <= >= == != | Severity < 3
| Priority | メッセージの優先度。(20 * Facility)+Severity で定義されます。 | > < <= >= == != | Priority >= 100
| Version | 使用されている syslog プロトコルのバージョン | > < <= >= == != | Version != 1
| Timestamp | ログメッセージに含まれるタイムスタンプの文字列表現 | == != | |
| Hostname | syslog メッセージを最初に送信したマシンのホスト名 | == != | Hostname != "myhost"
| Appname | syslog メッセージを最初に送信したアプリケーション（例：`systemd`） | == != | Appname != "dhclient"
| ProcID | メッセージを送信したプロセスを表す文字列（多くはPID） | == != | ProcID != "7053"
| MsgID | メッセージの種類を表す文字列 | == != | MsgID == "TCPIN"
| Message | ログメッセージそのもの | == != | Message == "Critical error!" |
| StructuredID | 最初の構造化データ要素の構造化データ ID を含む文字列（以下を参照） | == != | StructuredID == "ourSDID@32473"

次のような syslog レコードを考えてみましょう（出典：[https://github.com/influxdata/go-syslog](https://github.com/influxdata/go-syslog)）。

```
<165>4 2018-10-11T22:14:15.003Z mymach.it e - 1 [ex@32473 iut="3" foo="bar"] An application event log entry...
```

syslog モジュールは、以下のフィールドを抽出します。

* Facility: 20
* Severity: 5
* Priority: 165
* Version: 4
* Timestamp: "2018-10-11T22:14:15.003Z"
* Hostname: "mymach.it"
* Appname: "e"
* ProcID: <nil> (not set)
* MsgID: "1"
* Message: "An application event log entry..."

### 構造化されたデータ

上のレコード例では、`[ex@32473 iut="3" foo="bar"]` の部分が構造化データセクションです。構造化データセクションには、構造化された値のID（"ex@32473"、`StructuredID` キーワードで抽出）と、任意の数のキーと値のペアが含まれます。syslog モジュールを使って値にアクセスするには、`Structured.key` を指定します。`syslog Structured.iut` を指定すると、"3" という値を含む `iut` という名前の列挙型の値が抽出されます。同様に、`syslog StructuredID Structured.foo` を指定すると、"ex@32473" を含む `StructuredID` と、"bar" を含む `foo` が抽出されます。

1 つの syslog メッセージには、それぞれが ID を持つ複数の構造化データセクションが含まれている可能性があることに注意してください。両方のセクションで同じキーが定義されている場合は、どのセクションから抽出するかを明示的に指定することができます。これは、`syslog Structured[ex@32473].foo` というように構造化 ID を抽出に挿入することで可能です。ID を指定しない場合、モジュールはいずれかのセクションから結果を抽出しますが、どのセクションから抽出するかは保証されません。

注意：複数の構造化データセクションが存在する場合、StructuredID フィールドを抽出すると、最初のセクションの ID が返されます。フィルターはすべてのセクションの ID と照合されます。例えば、`[foo@bar a=b][baz@quux x=y]` を含むエントリに対して、`StructuredID!="baz@quux"` または `StructuredID!="foo@bar"` を指定すると、そのエントリは削除されます。`StructuredID=="baz@quux"` を指定すると、いずれかのセクションにマッチするため、例のエントリーを通過させることができます。ただし、いずれかのセクションのIDに "baz@quux" が含まれていないエントリーのみをドロップします。

## 例

### 深刻度別イベント数

```
tag=syslog syslog Severity | count by Severity | chart count by Severity
```

![Number of events by severity](severity.png)

### アプリケーション別の各深刻度レベルでのイベント数

```
tag=syslog syslog Appname Severity | count by Appname,Severity | table Appname Severity count
```

![Number of events at each severity by application](severity2.png)
