# ankoモジュール

[検索モジュールのドキュメント](#!search/searchmodules.md#Anko)で紹介されているように、Gravwellのankoモジュールは検索パイプライン内の汎用スクリプトツールである。これにより、スクリプト作成者にとっては複雑さを犠牲にしても、非常に柔軟に検索エントリを操作することができます。しかし、一度作成したスクリプトは、他のユーザーと簡単に共有することができます。

この言語自体の詳細については、[Anko スクリプト言語ドキュメント](scripting.md) で使用されているスクリプト言語の一般的な説明を参照してください。

### anko スクリプトでのネットワーク機能の無効化

デフォルトでは、ankoスクリプトはhttpやnetライブラリ、sftp、sshなどのネットワークユーティリティの使用を許可されています。Gravwellユーザーにネットワークアクセスを与えたくない場合は、`/opt/gravwell/etc/gravwell.conf`のオプション`Disable-Network-Script-Functions=true'を設定することで、この機能を無効にすることができます。

## ankoスクリプトの管理

ankoスクリプトを検索で実行するには、スクリプトを含むテキストファイルをリソースとしてアップロードする必要があります。リソースの作成とアップロードの方法については、[resources section](#!resources/resources.md)を参照してください。

現時点では、スクリプトを変更するには、元のテキストファイル内のスクリプトを編集してから、リソースにアップロードし直す必要があります。Gravwellの将来のバージョンでは、スクリプトをより簡単に作成できるようにテキストエディタが統合される予定です。

## ankoスクリプトの実行

ankoスクリプトを実行するには、"anko "の後にスクリプト名を指定し、さらにそのスクリプトに必要な引数を指定します。例えば、2つの数字を引数に取る「foo」というスクリプトを実行するには、パイプラインで「anko foo 1 3」と入力します。

## ankoスクリプトの書き方

Anko スクリプトは eval コマンドと同じ構文を使用します。以下の例と eval モジュールのセクションを参照してください。

### 必須関数の定義

anko スクリプトは、`Process` という名前の関数か、`Main` という名前の関数のどちらかを **必ず** 含む必要があります。これらの関数は引数を取りません。これら2つの関数は、検索エントリを処理するための2つの異なるオプションを表します。

エントリー上の列挙された値はローカル変数として扱うことができ、`Process`関数の戻り値はエントリーがパイプラインを進むことを許可される(true)かどうか(false)を決定します。

`main`関数が定義されている場合、それは一度だけ呼ばれます。したがって、プログラマは各検索エントリを取得し、それを操作した後にラインに渡すために、`readEntry`と`writeEntry`関数を呼ばなければなりません。

可能な限り、`Main`の代わりに`Process`を使ってスクリプトを書くことを強くお勧めします。なぜならば、概念的にずっとシンプルだからです。

### オプション関数の定義

スクリプトには `Parse` と `Finalize` という名前の関数を含めることができます。

`Parse`関数は `Process` や `Main` の前に呼び出され、コマンドライン引数を引数の配列として与えられます。Parse`関数は引数が正常に処理されたことをnilを返すことで示します。nil以外の値が返された場合にはエラーとして扱われ、ユーザに提示されます。スクリプトの引数を解析する方法のサンプルは、以下のサンプルスクリプトをご覧ください。

注意 Parse関数は、明示的に値を返さなければなりません。nilを返すと解析が成功したことになり、それ以外の値を返すとエラーになります。エラーが発生した場合には、問題を説明する文字列を返すことをお勧めします。

`Finalize`関数は`Process`や`Main`が完了した後に呼び出されます。この関数はスクリプトの最後に実行されるコードで、必要に応じてリソースを作成するのに適した場所です。

### Process()関数を使用したサンプルスクリプト

このサンプルスクリプトでは、スクリプトの引数で指定した2つのモードで動作します。build」モードでは、パケットモジュールから抽出した「SrcIP」フィールドを用いて、現在の検索で見られるIPアドレスのリストを構築し、そのリストをリソースとして保存します。適用」モードでは、以前に構築したテーブルを使用し、以前に見た「SrcIP」フィールドを含むエントリをすべて削除します。このスクリプトは、ネットワーク上の新しいデバイスを探すために使用されていましたが、現在では `lookup` モジュールが同じ機能をより柔軟に提供しています。

```
table = make(map[string]interface)
task = "build"

var json = import("encoding/json")

# first arg = "build" or "apply"
func Parse(args) {
	errstr = "argument must be 'build' or 'apply'"
	if len(args) == 1 {
		task = args[0]
	} else {
		return errstr
	}
	switch task {
	case "apply":
		# load the table
		data, _ = getResource("lookuptable")
		json.Unmarshal(data, &table)
	case "build":
	default:
		return errstr
	}
	return nil
}

func Process() {
	if task == "build" {
		s = toString(SrcIP)
		table[s] = true
		return true
	} else if task == "apply" {
		s = toString(SrcIP)
		# create & set an enumerated value named "new" to true or false
		setEnum("new", !table[s])
		# if it's not in the table, return true
		return !table[s]
	}
}

func Finalize() {
	if task == "build" {
		data, err = json.Marshal(table)
		if err != nil {
			return err
		}
		return setResource("lookuptable", data)
	}
}
```

"SrcIP"列挙型の値は、`Process`関数内の他の変数と同様に**読み取ることができますが、列挙型の値を設定するためには、その動作を明示するために`setEnum`関数を使用しなければならないことに注意してください。

変数 `table` と `task` は、関数定義の外側で宣言されていることに注意してください。組み込みのjsonエンコーディングライブラリも、関数定義の前にインポートされています。利用可能なライブラリの一覧は以下のとおりです。


### Main()関数を使ったサンプルスクリプト

Main()関数を使ってスクリプトを書くのは難易度が高いですが、エントリを複製する必要がある場合には必要です。以下のスクリプトは、Modbusメッセージを含むエントリを読み込み、メッセージが0x10タイプのリクエスト（"Write multiple registers"）である場合、スクリプトは書き込まれるレジスタごとに元のエントリを一度だけ複製し、各エントリに1つのレジスタアドレス＋レジスタ値を含む "RegAddr "および "RegValue "列挙値を付加します。

注：このスクリプトは単独では機能しません。このスクリプトは、パイプラインの初期にある別のankoスクリプトの出力を消費するように意図されており、「Request」や「WriteAddr」などの列挙値を生成します。

```
func Main() {
	for i = 0; i != -1; i++ {
		ent, err = readEntry()
		if err != nil {
			return
		}

		# Check if this is a request or a response
		Request, err = getEntryEnum(ent, "Request")
		if err != nil {
			#Request isn't set, this isn't a modbus packet, skip
			continue
		}

		# read the function value
		Function, err = getEntryEnum(ent, "Function")
		if err != nil {
			continue
		}

		ReqResp, err = getEntryEnum(ent, "ReqResp")
		if err != nil {
			continue
		}

		if Request == true {
			if Function == 0x10 {
				# write multiple registers
				Addr, err = getEntryEnum(ent, "WriteAddr")
				if err != nil {
					writeEntry(ent)
					continue
				}
				Count, err = getEntryEnum(ent, "WriteCount")
				if err != nil {
					writeEntry(ent)
					continue
				}
				if Count == 0 || len(ReqResp) < 5 + 2*Count {
					writeEntry(ent)
					continue
				}
				for j = 0; j < Count; j++ {
					newEnt = cloneEntry(ent)
					# read the register value
					val = (toInt(ReqResp[5+(2*j)]) << 8) | toInt(ReqResp[5+(2*j)+1])
					setEntryEnum(newEnt, "RegAddr", Addr + j)
					setEntryEnum(newEnt, "RegValue", val)
					err = writeEntry(newEnt)
					if err != nil {
						continue
					}
				}
			} else {
				writeEntry(ent)
			}
		}
	}
}
```

`ReadEntry`、`cloneEntry`, `writeEntry` 関数の使用に注意してください。このような検索エントリの明示的な管理は、`Process` 関数を使用するスクリプトでは必要ありません。また、列挙された値を変数として扱うのではなく、読み取るための `getEntryEnum` および `setEntryEnum` 関数の使用にも注意してください。

## ユーティリティー関数

Ankoはビルトインのユーティリティー関数を提供しており、以下のように `functionName(<functionArgs>) <returnValues>` という形式でリストアップされています。以下の関数は任意のAnkoスクリプトで使用できます。

* `getResource(name) []byte, error` は、指定されたリソースのコンテンツであるバイトのスライスを返し、エラーはリソースのフェッチ中に発生したあらゆるエラーです。
* `setResource(name, value) error` は、`name` という名前のリソースを（必要に応じて）作成し、`value` の内容で更新し、エラーが発生した場合はエラーを返します。
* `getMacro(name) string, error` は、与えられたマクロの値を返し、マクロが存在しない場合はエラーを返します。この関数はマクロの展開を行いませんのでご注意ください。
* `len(val) int` は val の長さを返します。val は文字列やスライスなどです。
* `toIP(string) IP` 文字列をIPに変換します。これは、パケットモジュールなどで生成されたIPと比較するのに適しています。
* `toMAC(string) MAC` は、文字列をMACアドレスに変換します。
* `toString(val) string` は、valを文字列に変換します。
* `toInt(val) int64` 可能であれば，valを整数に変換します．変換できない場合は0を返します。
* `toFloat(val) float64` は、可能であればvalを浮動小数点数に変換します。変換できない場合は0.0を返します。
* `toBool(val) bool` は，valを真偽値に変換しようとします。変換できない場合は false を返します。0以外の数値や文字列 "y"，"yes"，"true "はtrueを返します。
* 例： `toHumanSize(15127)` は "14.77KB" に変換されます。
* 例えば、 `toHumanCount(15127)` は "15.13 K" に変換されます。
* `typeOf(val) type` は、valの型を文字列で返します。
* `producesEnum(val)` スクリプトがその名前のEnumerated Valueを生成する予定であることをパイプラインに通知します。 Parse()関数の中で呼び出されなければなりません。
* `consumesEnum(val)` スクリプトがその名前のEnumerated Valueを消費する予定であることをパイプラインに通知します。 Parse()関数の中で呼び出されるべきです。

以下の関数は、`Process`関数を実装したスクリプトでのみ利用可能です:

* `setEnum(key, value) error` は、`key` という名前の現在のエントリに、`value` を含む列挙型の値を作成します。
* `getEnum(key) value, error` は、`key` で指定された列挙型の値を返します。
* `delEnum(key)` は、現在のエントリから `key` という名前の列挙型の値を削除します。
* `hasEnum(key) bool` は、現在のエントリが列挙された値を持っているかどうかを返します。

以下の関数は、`Main`関数を実装したスクリプトでのみ利用可能です。

* `readEntry() entry, error` は、次のエントリとエラー（もしあれば）を返します。エントリが残っていない場合は、エラーを返します。
* `writeEntry(ent) error` 与えられたエントリをパイプラインに書き出し、エラーがあればそれを返します。
* `cloneEntry(ent) entry` 指定したエントリのコピーを返します。
* `dupEntryByReference(ent, names...) entry` は、与えられたエントリを複製し、`names` パラメータで指定された列挙型の値のコピーを作成します。
* `newEntry() entry` は全く新しいエントリを作成し、タイムスタンプは現在のクエリの終了時に設定されます。
* `setEntryEnum(ent, key, value)` 指定したエントリに列挙型の値を設定します。
* `getEntryEnum(ent, key) value, error` 指定したエントリから列挙型の値を読み込みます。
* `hasEntryEnum(ent, key) bool` エントリが列挙型の値を含んでいるかどうかを返します。
* `delEntryEnum(ent, key)` 指定したエントリから、指定した列挙型の値を削除します。
* `setEntryData(ent, value)` エントリのデータ部分を設定します。
* `setEntrySrc(ent, ip)` エントリのソース・フィールドを設定します。
* `setEntryTimestamp(ent, time)` エントリのタイムスタンプを設定します。

注: `Process` 関数と `Main` 関数を使用するスクリプトでは、`setEnum`, `hasEnum`, および `delEnum` 関数が異なります。

## 組み込み変数

以下の変数は anko スクリプト用にあらかじめ定義されています。:

* `START`: クエリの開始時刻です。
* `END`: クエリの終了時刻です。
* `TAGMAP`: 文字列のタグ名からentry.EntryTagのタグ番号へのマップ。これは現在のクエリで使用されているタグのみを含むので、`tag=default,foo`とすると、TAGMAPには'default'→0と'foo'→1が含まれます。これは `cloneEntry` や `newEntry` 関数と組み合わせて使用します。

## 利用可能なパッケージ

特定のGoライブラリのankoラッパーを以下のような構文でインポートすることができます。

```
var json = import("encoding/json")
```

セキュリティ上の理由から、ankoモジュールは、Ankoスクリプト言語に含まれるすべての*パッケージへのアクセスを許可していません。以下のパッケージが anko スクリプトで使用できます:

* [bytes](https://golang.org/pkg/bytes): バイトスライスを扱う。
* [crypto/md5](https://golang.org/pkg/crypto/md5), [crypto/sha1](https://golang.org/pkg/crypto/sha1), [crypto/sha256](https://golang.org/pkg/crypto/sha256), [crypto/sha512](https://golang.org/pkg/crypto/sha512): 暗号のハッシュ化
* [crypto/tls](https://golang.org/pkg/crypto/tls): 限定的な TLS 機能
* [encoding/base64](https://golang.org/pkg/encoding/base64): 制限付き base64 機能
* [encoding/csv](https://golang.org/pkg/encoding/csv): CSV データのエンコードとデコードを行います。
* [encoding/hex](https://golang.org/pkg/encoding/hex): 限定的な 16 進数エンコーディング機能です。
* [encoding/json](https://golang.org/pkg/encoding/json): json データのエンコードとデコードを行ないます。
* [encoding/xml](https://golang.org/pkg/encoding/xml): 限定的な XML エンコーディング機能を提供します。
* [errors](https://golang.org/pkg/errors): Go のエラーを処理します。
* [flag](https://golang.org/pkg/flag): 限られたフラグ解析機能
* [fmt](https://golang.org/pkg/fmt): 文字列の印刷と書式設定を行います。
* [github.com/google/uuid](https://github.com/google/uuid): UUIDの生成と検査
* [github.com/gravwell/ipexist](https://github.com/gravwell/ipexist): GravwellのIPヘルパー機能
* [github.com/RackSec/srslog](https://github.com/RackSec/srslog): golang の標準ライブラリに代わる syslog パッケージです。
* [io](https://golang.org/pkg/io): 基本的な I/O プリミティブ
* [io/util](https://golang.org/pkg/io/util): `ioutil.ReadAll` 関数のみ (後述)
* [math](https://golang.org/pkg/math): 数学関数
* [math/big](https://golang.org/pkg/math/big): ビグナム(bignum)
* [math/rand](https://golang.org/pkg/math/rand): 乱数
* [net](https://golang.org/pkg/net): 限られたネットワーク機能
* [net/https](https://golang.org/pkg/net/https): 限定的な HTTP 機能 (クライアントのみ)
* [net/url](https://golang.org/pkg/net/url): URL
* [path](https://golang.org/pkg/path): パス
* [path/filepath](https://golang.org/pkg/path/filepath): ファイルに特化したパス機能
* [regexp](https://golang.org/pkg/regexp): 正規表現
* [sort](https://golang.org/pkg/sort): 並べ替え
* [strings](https://golang.org/pkg/strings): 文字列処理関数
* [time](https://golang.org/pkg/time): 時間処理関数
* [github.com/ziutek/telnet](https://github.com/ziutek/telnet): telnetクライアント関数 

すべてのパッケージを網羅的に説明することはできません。各パッケージでエクスポートされる利用可能な機能は、[anko公式リポジトリ](https://github.com/mattn/anko/tree/master/packages)で見ることができます。いくつかのパッケージは、公式 anko リポジトリでエクスポートされる機能を完全には提供していないため、以下でさらに説明します。

## パッケージの制限

パッケージの中には、スクリプト言語でのエクスポートが危険な機能を持つものがあります。Gravwellではパッケージのエクスポートを制限していますが、これはフルパッケージで利用できる機能の一部です。以下は各パッケージの制限事項の一覧です。

### crypto/md5

`crypto/md5`は "New" と "Sum" 関数のみをエクスポートします:

- `md5.New`
- `md5.Sum`

### crypto/sha1

`crypto/sha1`は "New" および "Sum" 関数のみをエクスポートします:

- `sha1.New`
- `sha1.Sum`

### crypto/sha256

`crypto/sha256`では、さまざまな「New」と「Sum」の関数のみをエクスポートします:

- `sha256.New`
- `sha256.New224`
- `sha256.Sum224`
- `sha256.Sum256`

### crypto/sha512

`crypto/sha512`では、さまざまな「New」および「Sum」関数のみをエクスポートします:

- `sha512.New`
- `sha512.New384`
- `sha512.New512_224`
- `sha512.New512_256`
- `sha512.Sum384`
- `sha512.Sum512`
- `sha512.Sum512_224`
- `sha512.Sum512_256`

### crypto/tls

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ使用できます。`crypto/tls`は、`net/http`モジュールで使用するTLSコンフィグタイプのみをエクスポートするします:

- `tls.Config`

### encoding/csv

`encoding/csv` は CSV のイニシャライザのみをエクスポートします。:

- `csv.NewReader` (実際には、LazyQuotesオプションをtrueに設定して、`csv.NewReader`を呼び出します)
- `csv.NewWriter`
- `csv.NewBuilder`

### encoding/base64

`encoding/base64` は base64 のイニシャライザとエンコーディングタイプのみをエクスポートします:

- `base64.NewDecoder`
- `base64.NewEncoder`
- `base64.NewEncoding`
- `base64.RawStdEncoding`
- `base64.RawURLEncoding`
- `base64.StdEncoding`
- `base64.URLEncoding`

### encoding/hex

`encoding/hex`は、イニシャライザとラッパのサブセットをエクスポートします

- `hex.Decode`
- `hex.DecodeString`
- `hex.DecodedLen`
- `hex.Dump`
- `hex.Dumper`
- `hex.Encode`
- `hex.EncodeToString`
- `hex.EncodedLen`
- `hex.NewDecoder`
- `hex.NewEncoder`

### encoding/xml

`encoding/exml` は、イニシャライザ、ラッパー、およびエンコーディングオプションのサブセットをエクスポートします:

- `xml.Escape`
- `xml.EscapeText`
- `xml.Marshal`
- `xml.MarshalIndent`
- `xml.Unmarshal`
- `xml.NewDecoder`
- `xml.NewTokenDecoder`
- `xml.NewEncoder`
- `xml.HTMLAutoClose`
- `xml.HTMLEntity`
- `xml.Attr`
- `xml.CharData`
- `xml.Comment`
- `xml.Directive`
- `xml.EndElement`
- `xml.Name`
- `xml.ProcInst`
- `xml.StartElement`

### flag 

`flag`はパッケージ全体ではなく、タイプのサブセットのみをエクスポートします:

- `flag.NewFlagSet`
- `flag.PanicOnError`
- `flag.ContinueOnError`

### github.com/google/uuid

`github.com/google/uuid`では、"New "と "Parse "機能のみをエクスポートします:

- `uuid.New`
- `uuid.Parse`
- `uuid.ParseBytes`

### github.com/gravwell/ipexist

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ使用できます。github.com/gravwell/ipexist`では、"New "関連の関数のみをエクスポートします:

- `ipexist.New`
- `ipexist.NewIPBitMap`

### github.com/RackSec/srslog

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ利用できます。github.com/RackSec/srslog`ではsyslog関連の機能のみを公開しています:

- `srslog.Dial`
- `srslog.DefaultFormatter`
- `srslog.DefaultFramer`
- `srslog.RFC3164Formatter`
- `srslog.RFC5424Formatter`
- `srslog.RFC5425MessageLengthFramer`
- `srslog.UnixFormatter`
- `srslog.LOG_EMERG`
- `srslog.LOG_ALERT`
- `srslog.LOG_CRIT`
- `srslog.LOG_ERR`
- `srslog.LOG_WARNING`
- `srslog.LOG_NOTICE`
- `srslog.LOG_INFO`
- `srslog.LOG_DEBUG`
- `srslog.LOG_KERN`
- `srslog.LOG_USER`
- `srslog.LOG_MAIL`
- `srslog.LOG_DAEMON`
- `srslog.LOG_AUTH`
- `srslog.LOG_SYSLOG`
- `srslog.LOG_LPR`
- `srslog.LOG_NEWS`
- `srslog.LOG_UUCP`
- `srslog.LOG_CRON`
- `srslog.LOG_AUTHPRIV`
- `srslog.LOG_FTP`
- `srslog.LOG_LOCAL0`
- `srslog.LOG_LOCAL1`
- `srslog.LOG_LOCAL2`
- `srslog.LOG_LOCAL3`
- `srslog.LOG_LOCAL4`
- `srslog.LOG_LOCAL5`
- `srslog.LOG_LOCAL6`
- `srslog.LOG_LOCAL7`

### github.com/ziutek/telnet

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ使用できます。エクスポートされる関数や型は以下の通りです:

- `telnet.Dial`
- `telnet.DialTimeout`
- `telnet.NewConn`
- `telnet.Conn`

### io/ioutil

`io/ioutil`、`ioutil.ReadAll()`という1つの関数しかエクスポートしていません

### net

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ使用できます。エクスポートされる関数や型は以下の通りです:

- `net.CIDRMask`
- `net.Dial`
- `net.DialIP`
- `net.DialTCP`
- `net.DialTimeout`
- `net.DialUDP`
- `net.ErrWriteToConnected`
- `net.FlagBroadcast`
- `net.FlagLoopback`
- `net.FlagMulticast`
- `net.FlagPointToPoint`
- `net.FlagUp`
- `net.IPv4`
- `net.IPv4Mask`
- `net.IPv4allrouter`
- `net.IPv4allsys`
- `net.IPv4bcast`
- `net.IPv4len`
- `net.IPv4zero`
- `net.IPv6interfacelocalallnodes`
- `net.IPv6len`
- `net.IPv6linklocalallnodes`
- `net.IPv6linklocalallrouters`
- `net.IPv6loopback`
- `net.IPv6unspecified`
- `net.IPv6zero`
- `net.InterfaceAddrs`
- `net.InterfaceByIndex`
- `net.InterfaceByName`
- `net.Interfaces`
- `net.JoinHostPort`
- `net.LookupAddr`
- `net.LookupCNAME`
- `net.LookupHost`
- `net.LookupIP`
- `net.LookupMX`
- `net.LookupNS`
- `net.LookupPort`
- `net.LookupSRV`
- `net.LookupTXT`
- `net.ParseCIDR`
- `net.ParseIP`
- `net.ParseMAC`
- `net.ResolveIPAddr`
- `net.ResolveTCPAddr`
- `net.ResolveUDPAddr`
- `net.ResolveUnixAddr`
- `net.SplitHostPort`

### net/http

このモジュールは、Gravwellの設定で`Disable-Network-Script-Functions`が`false`に設定されている場合にのみ利用可能です。

`net/http` はHTTP *requests* を実行するための関数、型、変数のサブセットをエクスポートします。タイプとしては`Client`, `Cookie`, `Request`, `Response`がエクスポートされます。これらのタイプの説明は[Go documentation](https://golang.org/pkg/net/http/)を参照してください。

関数 `NewRequest(operation, url, body) (Request, error)` は、指定されたURL文字列に対して、指定された操作("PUT", "POST", "GET", "DELETE")を行う新しいhttp.Requestを準備します。bodyはPUTおよびPOSTリクエストで使用するオプションのパラメータで、'nil'またはio.Readerのいずれかに設定する必要があります。

ほとんどの場合、NewRequest関数でリクエストを作成し、http.DefaultClientを使用してそのリクエストを実行します。:

```
req, _ = http.NewRequest("GET", "http://example.org/foo", nil)
resp, _ = http.DefaultClient.Do(req)
resp.Body.Close()
```

送信前のリクエストに、追加のヘッダーやクッキーを設定することができます:

```
req, _ = http.NewRequest("GET", "http://example.org/foo", nil)
# Add a header
req.Header.Add("My-Header", "gravwell")
# Add a cookie
cookie = make(http.Cookie)
cookie.Name = "foo"
cookie.Value = "bar"
req.AddCookie(&cookie)
resp, _ = http.DefaultClient.Do(req)
resp.Body.Close()
```

警告: 上の図のように、終了時には http.Response の Body フィールドを *必ず* 閉じなければなりません。Bodyを開いたままにしておくと、ネットワーク接続が開いたままになり、最終的にサーチエージェントのソケットが足りなくなってしまいます。httpGet`と`httpPost`関数は自動的にBodyを閉じますので、可能な限りこれらの使用を検討してください。

