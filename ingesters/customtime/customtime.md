# インジェスターカスタム時刻フォーマット

多くのインジェスターは、Gravwell TimeGrinder時間解決システムの機能を拡張できるカスタム時間フォーマットの組み込みをサポートしています。[TimeGrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder)には、自動的に識別して解決できるタイムスタンプフォーマットが豊富に用意されています。 しかし、実際の開発者がいる現実世界では、システムがどのような時間フォーマットを使用するかはわかりません。そのため、TimeGrinderシステムでは、ユーザーがカスタムタイムフォーマットを指定できるようになっています。

## サポートされるインジェスター

すべてのインジェスターがカスタムタイムフォーマットをサポートしているわけではありません。[シングルファイル](https://github.com/gravwell/gravwell/blob/v3.7.0/ingesters/singleFile/main.go)のようなワンオフまたはスタンドアロンのインジェスターは、手動で起動することを目的としたアプリケーションであり、設定ファイルはありません。[netflow](#!ingesters/ingesters.md#Netflow_Ingester)のような専用のインジェスターは、タイムスタンプを解決する必要がないため、カスタムフォーマットの必要はありません。

以下のインジェスターは、カスタムタイムフォーマットの組み込みをサポートしています:

* [シンプルリレー](#!ingesters/ingesters.md#Simple_Relay)
* [ファイルフォロワー](#!ingesters/ingesters.md#File_Follower)
* [HTTPインジェスター](#!ingesters/ingesters.md#HTTP)
* [Amazon Kinesis](#!ingesters/ingesters.md#Kinesis_Ingester)
* [Microsoft Graph API](#!ingesters/ingesters.md#Microsoft_Graph_API_Ingester)
* [Office 365](#!ingesters/ingesters.md#Office_365_Log_Ingester)
* [Kafka](#!ingesters/ingesters.md#Kafka)

## カスタムフォーマットの定義

カスタムフォーマットには3つのアイテムが必要です:

* 名前
* 正規表現
* フォーマット

カスタム時刻フォーマットの名前は、他のカスタム時刻フォーマットや含まれるtimegrinderフォーマットで一意でなければなりません。 含まれる時間フォーマットとその名前の完全な最新リストについては、[TimeGrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder#pkg-constants)をご覧ください。

カスタムの時間フォーマットは、サポートされているインジェスターの設定ファイルで、`TimeFormat`という名前のINIブロックを指定して宣言します。ここでは、アンダースコアで区切られたタイムスタンプを扱う"foo"というフォーマットの例を紹介します:

```
[TimeFormat "foo"]
	Format="2006_01_02_15_04_05"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}`
```

このフォーマットは、以下のログのタイムスタンプを適切に処理します:

```
2021_02_05_09_00_00 and my id is 1
2021_02_05_09_00_00 and my id is 2
2021_02_05_09_00_00 and my id is 3
2021_02_05_09_00_00 and my id is 4
2021_02_05_09_00_00 and my id is 5
2021_02_05_09_00_00 and my id is 6
```

ここでは、タイムスタンプのみのログを扱う別のフォーマットを紹介します:

```
[TimeFormat "foo2"]
	Format="15^04^05"
	Regex=`\d{1,2}\^\d{1,2}\^\d{1,2}`
```

このフォーマットでは、抽出された各タイムスタンプに適切に現在の日付を適用して、以下のログを処理します:

```
09^00^00 and my id is 1
09^00^00 and my id is 2
09^00^00 and my id is 3
09^00^00 and my id is 4
09^00^00 and my id is 5
09^00^00 and my id is 6
```

注：カスタムのタイムスタンプフォーマット名は、[Timestamp-Format-Override](#!ingesters/ingesters.md#Time_Parsing_Overrides)の値で使用できます。 例えば、`Timestamp-Format-Override="foo"`を使用して、タイムスタンプフォーマットをカスタムフォーマットに強制することができます。

### 時刻フォーマット

`Format`コンポーネントは[Go standard time format specification](https://golang.org/pkg/time/#pkg-constants)を使用しています。 簡単に説明すると、どのようなフォーマットであっても、日付 `Mon Jan 2 15:04:05 MST 2006` を記述しなければなりません。

時間フォーマットでは、日付の要素を省略することができます。 カスタムフォーマットシステムは、カスタム時間フォーマットに日付コンポーネントが含まれていないことを識別すると、抽出されたタイムスタンプの日付を自動的に `today` に更新します。

### タイムゾーン

すべてのカスタム時刻フォーマットは、`Format` ディレクティブで指定されていない限り、 UTC で動作するようになっています。 つまり、日付要素のない時刻フォーマットを使用する場合は、タイムゾーンに特別な注意を払わなければなりません。 アプリケーションがMSTの`12:00:00`というタイムスタンプを発行し、タイムゾーンコンポーネントやタイムゾーンオーバーライドがない場合、timegrinderはタイムスタンプをUTCと解釈し、抽出された日付は7時間前のものになります。

タイムスタンプにタイムゾーンが含まれている場合には、タイムグラインダーシステムが正しいタイムゾーンでタイムスタンプを解釈できるように、`Format` ディレクティブにタイムゾーンを含める必要があります。 例えば、前述の "foo "カスタムフォーマットにタイムゾーンの要素を加えたものは以下の通りです:

```
[TimeFormat "foo"]
	Format="2006_01_02_15_04_05_MST"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\S+`
```

この例では、それぞれのタイムゾーンのタイムスタンプを適切に処理し、抽出時に正しいタイムスタンプを適用します。

## 例

ここでは、2つのカスタムタイムフォーマットを追加する[ファイルフォロワー](#!ingesters/file_follow.md)の設定例を紹介します:

```
[Global]
Ingester-UUID="463c1889-2954-40a0-a3b4-705ea66459f6"
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
State-Store-Location=/opt/gravwell/etc/file_follow.state
Log-Level=INFO #options are OFF INFO WARN ERROR
Log-File=/opt/gravwell/log/file_follow.log
Max-Files-Watched=64 # Maximum number of files to watch before rotating out old ones, this can be bumped but will need sysctl flags adjusted

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Follower "auth"]
	Base-Directory="/tmp/logs/"
	File-Filter="*.log" #we are looking for all authorization log files
	Tag-Name=test
	Assume-Local-Timezone=true #Default for assume localtime is false

[TimeFormat "foo"]
	Format="2006_01_02_15_04_05"
	Regex=`\d{4}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}_\d{1,2}`

[TimeFormat "foo2"]
	Format="15!04!05"
	Regex=`\d{1,2}!\d{1,2}!\d{1,2}`

```

`2021_02_14_12_33_52`や`15!05!22`と指定されたタイムスタンプは、カスタムタイムフォーマットの追加により、ファイルフォロワーで適切に処理されます。
