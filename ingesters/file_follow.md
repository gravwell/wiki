# ファイルフォロワー

ファイルフォロワー・インジェスターは、ローカルファイルシステム上のファイルがその場で更新される可能性のある状況でファイルを取り込むための最適な方法です。ファイルの各行を1つのエントリとして取り込みます。

ファイルフォロワーの最も一般的な使用例は、/var/logのような、活発に更新されているログファイルを含むディレクトリを監視することです。ファイルフォロワーはログのローテーションをインテリジェントに処理し、`logfile` が `logfile.1` に移動したことなどを検出します。ディレクトリ内の特定のパターンに一致するファイルを取り込むように設定することができ、オプションでトップレベルのディレクトリのサブディレクトリに再帰的に降りていくことができます。

注: RHEL/Centosでは、`/var/log`は "adm"ではなく "root"グループに属しています。ファイルフォロワーはデフォルトではadmグループで動作しますので、`/var/log`を読みたい場合は`chgrp -R adm /var/log`とするか、systemdユニットファイルでグループを変更する必要があります。


## 基本的な設定

ファイルフォロワーの設定ファイルは、デフォルトでLinuxでは`/opt/gravwell/etc/file_follow.conf`に、Windowsでは`C:GRAVWEL\file_follow.cfg`にあります。

ファイルフォロワーインジェスターは[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。他の多くのGravwellインジェスターと同様に、ファイルフォロワーは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルロギングをサポートしています。

注: ファイルフォロワーインジェスターではファイルキャッシュを使用しないことを強くお勧めします。なぜなら、ソースファイル内の位置をすでに追跡しているからです。

以下は、/var/logにある複数の異なるタイプのログファイルを監視し、/tmp/incomingにあるファイルを再帰的に追跡するように構成されたファイルフォロワーインジェスターの設定例です:

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
Cleartext-Backend-target=172.20.0.1:4023 #example of adding a cleartext connection
Cleartext-Backend-target=172.20.0.2:4023 #example of adding another cleartext connection
State-Store-Location=/opt/gravwell/etc/file_follow.state
Log-Level=ERROR #options are OFF INFO WARN ERROR
Max-Files-Watched=64

[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true
[Follower "auth"]
        Base-Directory="/var/log/"
        File-Filter="auth.log,auth.log.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
        Assume-Local-Timezone=true #Default for assume localtime is false
[Follower "packages"]
        Base-Directory="/var/log"
        File-Filter="dpkg.log,dpkg.log.[0-9]" #we are looking for all dpkg files
        Tag-Name=dpkg
        Ignore-Timestamps=true
[Follower "external"]
        Base-Directory="/tmp/incoming"
		Recursive=true
        File-Filter="*.log"
        Tag-Name=external
		Timezone-Override="America/Los_Angeles"
```

この例では、"syslog "フォロワーが `/var/log/syslog` とそのローテーションを読み、行を syslog タグに取り込み、日付はローカルのタイムゾーンであると仮定しています。同様に、"auth "フォロワーもsyslogタグを使って、`/var/log/auth.log`を取り込みます。"packages"フォロワーは、Debian のパッケージ管理ログを dpkg タグで取り込みます。説明のために、タイムスタンプを無視して、各エントリに読み込まれた時間をマークします。

最後に、"external"フォロワーは、ディレクトリ `/tmp/incoming` から `.log` で終わるすべてのファイルを読み込み、再帰的にディレクトリをさかのぼります。タイムスタンプはタイムゾーンがPacific時間であるかのように解析されます。このフォロワーでは、例えばアメリカ西海岸にある複数のサーバーが定期的にログファイルをこのシステムにアップロードしている場合に便利な設定を紹介しています。

上記で使用した設定パラメータの詳細は、以下のセクションで説明します。

## 追加のグローバルパラメーター

### Max-Files-Watched

Max-Files-Watched パラメーターは、ファイルフォロワーが開いているファイルの数が増えすぎることを防ぎます。Max-Files-Watched=64`が指定された場合、ファイルフォロワーは最大64個のログファイルをアクティブに監視します。新しいファイルが作成されると、ファイルフォロワーは新しいファイルを監視するために、既存の最も古いファイルの監視を停止します。しかし、古いファイルが後に更新された場合は、キューの先頭に戻ります。

ほとんどの場合、この設定を64のままにしておくことをお勧めします。上限を高く設定しすぎると、カーネルが設定した制限に抵触する可能性があります。

## フォロワー設定

ファイルフォロワー設定ファイルには、1つ以上の"フォロワー"ディレクティブが含まれています:

```
[Follower "syslog"]
        Base-Directory="/var/log/"
        File-Filter="syslog,syslog.[0-9]" #we are looking for all authorization log files
        Tag-Name=syslog
```

各フォロワーは、最低でもベースディレクトリとファイル名のフィルタリングパターンを指定します。このセクションでは、フォロワーごとに設定可能な設定パラメーターについて説明します。

###	Base-Directory

Base-Directoryパラメータは、取り込み対象となるファイルを格納するディレクトリを指定します。絶対パスでなければならず、ワイルドカードは使用できません。

### File-Filter

File-Filterパラメーターは、取り込むべきファイル名を定義します。それは、単一のファイル名のように単純なものである可能性があります:

```
File-Filter="foo.log"
```

または複数のパターンを含むことができます:

```
File-Filter="kern*.log,kern*.log.[0-9]"
```

これは、"kern "で始まり".log "で終わるファイル名、または "kern "で始まり".log.0 "から".log.9 "で終わるファイル名にマッチします。

完全なマッチング構文は、[https://golang.org/pkg/path/filepath/#Match](https://golang.org/pkg/path/filepath/#Match)で定義されています:

```
pattern:
	{ term }
term:
	'*'         matches any sequence of non-Separator characters
	'?'         matches any single non-Separator character
	'[' [ '^' ] { character-range } ']'
	            character class (must be non-empty)
	c           matches character c (c != '*', '?', '\\', '[')
	'\\' c      matches character c

character-range:
	c           matches character c (c != '\\', '-', ']')
	'\\' c      matches character c
	lo '-' hi   matches character c for lo <= c <= hi
```

### Recursive

recursive パラメータは、File Follower がBase-Directory下のFile-Filterにマッチするファイルを再帰的に取り込むことを指示します。

デフォルトでは、インジェスターは Base-Directory のトップレベルにある File-Filter にマッチするファイルのみを取り込みます。以下の例では、`/tmp/incoming/foo.log`は取り込みますが、`/tmp/incoming/system1/foo.log`は取り込みません:

```
Base-Directory="/tmp/incoming"
File-Filter="foo.log"
Recursive=false
```

Recusive=trueを設定すると、`/tmp/incoming`以下の任意のディレクトリの深さにある、**foo.logという名前のファイルを取り込むことができます。

### Tag-Name

Tag-Nameパラメータは、このフォロワーが取り込んだエントリーに適用するタグを指定します。

### Ignore-Line-Prefix

インジェスターは、Ignore-Line-Prefixに渡された文字列で始まるすべての行をドロップします(インジェストしません)。これは、Broログのようなコメントを含むログファイルをインジェストするときに便利です。Ignore-Line-Prefixパラメータは複数回指定できます。

以下は、`#`または`//`で始まる行がインジェストされないことを示しています。:

```
Ignore-Line-Prefix="#"
Ignore-Line-Prefix="//"
```

### Regex-Delimiter

`Regex-Delimiter`オプションでは、エントリーを分割する際に、改行ではなく、正規表現を指定することができます。たとえば、入力ファイルが次のようなものだとすると:

```
####This is the first entry
additional data
####This is the second entry
```

フォロワーの定義に次のような行を追加することができます:

```
Regex-Delimiter="####"
```

これは、前のファイルを解析し2つのエントリーにします:

```
####This is the first entry
additional data
```

と

```
####This is the second entry
```

注: `Timestamp-Delimited` は `Regex-Delimiter` よりも優先されるので、どちらか一方を設定してください。

### Timestamp-Delimited

Timestamp-Delimitedパラメータは、タイムスタンプが出現するたびに、新しいエントリの開始とみなすことを指定するブール値です。これは、ログエントリが複数の行にまたがる場合に便利です。Timestamp-Delimitedを指定する場合は、Timestamp-Format-Overrideパラメータも設定する必要があります。

ログファイルが以下のような場合:

```
2012-11-01T22:08:41+00:00 Line 1 of the first entry
Line 2 of the first entry
2012-11-01T22:08:43+00:00 Line 1 of the second entry
Line 2 of the second entry
Line 3 of the second entry
```

フォロワーが`Timestamp-Delimited=true`と`Timestamp-Format-Override=RFC3339`で設定されている場合、以下の2つのエントリーが生成されます:

```
2012-11-01T22:08:41+00:00 Line 1 of the first entry
Line 2 of the first entry
```
```
2012-11-01T22:08:43+00:00 Line 1 of the second entry
Line 2 of the second entry
Line 3 of the second entry
```

注: `Timestamp-Delimited` は `Regex-Delimiter` よりも優先されるので、どちらか一方を設定してください。

### Ignore-Timestamps

Ignore-Timestampsパラメータは、フォロワーがファイルの各行からタイムスタンプを抽出せず、各行に現在時刻をタグ付けすることを示します。

### Assume-Local-Timezone

Assume-Local-Timezoneは、タイムゾーンの指定がないタイムスタンプを、デフォルトのUTCではなくローカルのタイムゾーンで解析するよう、インジェスターに指示するブール値の設定です。

Assume-Local-TimezoneとTimezone-Overrideは相互に排他的です。

### Timezone-Override

Timezone-Overrideパラメータは、タイムゾーンの指定がないタイムスタンプを、デフォルトのUTCではなく指定されたタイムゾーンで解析するようにインジェスターに指示します。タイムゾーンは、[https://en.wikipedia.org/wiki/List_of_tz_database_time_zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)に示すように、IANAデータベースの文字列形式で指定する必要があります:

```
Timezone-Override="America/Chicago"
```

Assume-Local-TimezoneとTimezone-Overrideは相互に排他的です。`Timezone-Override="Local"`は、機能的には`Assume-Local-Timezone=true`と同じです。

### Timestamp-Format-Override

データ値に複数のタイムスタンプが含まれている場合、データからタイムスタンプを導き出そうとすると、混乱を招くことがあります。通常、フォロワーは導出可能な一番左のタイムスタンプを取得しますが、エントリ内に複数のタイムスタンプがある場合、最初に試すフォーマットを指定すると便利です。"Timestamp-Format-Override"は、フォロワーが特定のフォーマットを最初に試すように指示します。 以下のようなタイムスタンプフォーマットがあります:

* AnsiC
* Unix
* Ruby
* RFC822
* RFC822Z
* RFC850
* RFC1123
* RFC1123Z
* RFC3339
* RFC3339Nano
* Apache
* ApacheNoTz
* Syslog
* SyslogFile
* SyslogFileTZ
* DPKG
* Custom1Milli
* NGINX
* UnixMilli
* ZonelessRFC3339
* SyslogVariant
* UnpaddedDateTime

可能なオーバーライドの全リストとその例については、[TimeGrinder](https://pkg.go.dev/github.com/gravwell/gravwell/v3/timegrinder)を参照してください。

RFC3339の仕様にマッチするタイムスタンプを最初に探すようにフォロワーに強制するには、フォロワーに `Timestamp-Format-Override=RFC3339`を追加します。なお、RFC3339のタイムスタンプが見つからない場合は、他のフォーマットにもマッチさせようとします。

### Timestamp-Regex and Timestamp-Format-String

`Timestamp-Regex`と`Timestamp-Format-String`オプションは、このフォロワーのタイムスタンプを解析する際に使用する追加のタイムスタンプフォーマットを指定するために、同時に使用することができます。例えば、Oracle WebLogicのタイムスタンプ（例："Sep 18, 2020 12:26:48,992 AM EDT"）を含むログを取り込む場合は、以下を設定に追加します:

```
	Timestamp-Regex=`[JFMASOND][anebriyunlgpctov]+\s+\S{1,2},\s+\d{4}\s+\d{1,2}:\d\d:\d\d,\d+\s+\S{2}\s+\S+`
	Timestamp-Format-String="Jan _2, 2006 3:04:05,999 PM MST"
```

`Timestamp-Format-String`パラメータは、[このドキュメント](https://golang.org/pkg/time/)で定義されているGoスタイルのタイムスタンプフォーマットでなければなりません。`Timestamp-Regex`パラメータは、抽出したいタイムスタンプにマッチする正規表現でなければなりません。これは、`Timestamp-Format-String`にもマッチしなければならず、マッチしない場合はエラーを返します。

これらのオプションを使って定義されたフォーマットは、タイムグラインダーが使用するフォーマットのリストの一番上に挿入され、最初にチェックされますが、ユーザー定義のフォーマットで有効なタイムスタンプが見つからない場合は、タイムグラインダーの他のフォーマットも試されることになります。
