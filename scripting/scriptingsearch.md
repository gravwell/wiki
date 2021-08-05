# 自動化スクリプト

Gravwellは、検索の実行、リソースの更新、アラートの送信、アクションの実行が可能な堅牢なスクリプトエンジンを提供します。 このエンジンは、検索を実行して自動的にデータを調べ、人間を介さずに検索結果に基づいてアクションを起こすことができます。 

オートメーションスクリプトは、[スケジュール]（schedulesearch.md）または[コマンドラインクライアント]（#!cli/cli.md）から手動で実行することができます。 CLIではスクリプトを対話的に再実行することができるため、スケジュール検索を作成する前にCLIでスクリプトを開発およびテストすることをお勧めします。 下記の[Client](#!scripting/scriptingsearch.md#Gravwell_Client_Usage)のサンプルスクリプトをご覧ください。

## 組込み機能

スクリプトは、[anko](#!scripting/anko.md)モジュールで利用可能なものとほぼ同じ組み込み関数を使用することができますが、検索の起動や管理のためにいくつか追加されています。 関数は以下のように、`functionName(<functionArgs>) <returnValues>`というフォーマットで表示されます。  これらの関数は特定の機能のための便利なラッパーとして提供されていますが、完全な[Gravwell client](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client)は`getClient`ラッパーを使って入手できます。  ラッパーの`getClient`は、スクリプトを実行するユーザーとしてサインインしているクライアントオブジェクトを返します。

## バージョン管理

Gravwellは常に新しいモジュール、メソッド、機能を追加しています。 与えられたスクリプトが現在のバージョンで動作するかどうかを検証できることが望ましい場合が多々あります。 これは、互換性のあるGravwellの最小バージョンと最大バージョンを指定する2つのビルトインスクリプト関数によって達成されます。 どちらかのアサーションが失敗すると、スクリプトはバージョンが互換性がないことを示すエラーで直ちに失敗します。

* `MinVer(major, minor, point)` Gravwellが少なくとも特定のバージョンであることを確認します。
* `MaxVer(major, minor, point)` Gravwellが特定のバージョンより新しくないことを保証します。

## ライブラリー外部機能

Gravwellのバージョン3.3.1では、オートメーションスクリプトに外部スクリプトライブラリをインクルードできるようになりました。 追加のライブラリをインクルードするために2つの機能が提供されています。

* `include(path, commitid, repo) error`ライブラリファイルをインクルードします。repoとcommitid引数はオプションです。インクルードに失敗した場合、失敗の理由が返されます。
* `require(path, commitid, repo)include`と同様の動作をしますが、失敗した場合はスクリプトが停止され、失敗の理由がスクリプトの結果に添付されます。

`include`と`require`はどちらもオプションで正確なリポジトリやcommitidを指定することができる。`repo`引数が省略されている場合、`https://github.com/gravwell/libs`の Gravwell デフォルトのライブラリレポが使われます。`commitid`が省略されている場合は、`HEAD`のコミットが使われる。 レポは、レポパスに定義されたスキーマ(`http://`, `https://`, `git://`)を介してGravwellウェブサーバからアクセス可能でなければなりません。 スクリプトシステムは必要に応じて自動的にレポを取得します: もし現在知られていないコミットIDが要求された場合、Gravwellはレポの更新を試みます。

エアギャップシステムを利用している場合や、Gravwellがgithubにアクセスできないようにしたい場合は、`gravwell.conf`ファイルで`Library-Repository`と`Library-Commit`設定変数を使って、内部ミラーやデフォルトコミットを指定することができます。 例えば、以下のようになります。

```
Library-Repository="https://github.com/foobar/baz" #override the default library
Library-Commit=da4467eb8fe22b90e5b2e052772832b7de464d63
```

Library-Repository は、Gravwell ウェブサーバープロセスが読めるローカルフォルダにすることもできます。 たとえば、完全にエアギャップされた環境で Gravwell を実行している場合でも、lib にアクセスして更新できるようにしたいと思うかもしれません。 git リポジトリを解凍して `Library-Repository`をそのパスに設定するだけです。

```
Library-Repository="/opt/gitstuff/gravwell/libs"
```

`gravwell.conf`ファイルで`Disable-Library-Repository`を設定することにより、`include`と`require`を無効にすることができます（それにより外部コードを禁止します）。

## グローバル設定

* `loadConfig(resource) (map[string]interface, error)` は指定されたリソースを読み込み、それをJSON構造体として解析し、結果をマップとして返す。リソースに `{"foo": "bar", "a":1}` が含まれている場合、この関数は "foo" → "bar" と "a" → 1 のマップを返します。

システムに多数の自動化スクリプトがある場合は、すべてのスクリプトで使用できるように、構成値のリポジトリをどこかに保持しておくと便利な場合があります。 すべてのスクリプトで電子メールアラートライブラリの特定のリビジョンを使用したいとします。 「script-config」という名前のリソースに次のものを入れることができます。

```
{
"email_lib_revision":"14ceec90b69943992f4efae8fc9e24c3f4767944"
}
```

で、以下のコードをスクリプトで使用します：

```
cfg, err = loadConfig("script-config")
if err != nil {
	return nil
}
require(`alerts/email.ank`, cfg.email_lib_revision)
```

## リソースと持続的データ

* `getResource(name) []byte, error`は指定されたリソースの内容をバイト数で返し、エラーはリソースを取得する際に発生したエラーを返します。
* `setResource(name, value) error`は`name`という名前のリソースを作成(必要に応じて)して`value`の内容で更新し、エラーが発生した場合にはエラーを返します。
* `setPersistentMap(mapname, key, value)` は、スケジュールされたスクリプトを実行する間に持続するキーと値のペアをマップに格納します。
* `getPersistentMap(mapname, key)value` は、指定されたキーに関連付けられた値を、指定された永続マップから返します。
* `delPersistentMap(mapname, key)` は指定されたキーと値のペアを指定されたマップから削除します。
* `persistentMap(mapname)`は返されたマップの変更は実行に成功したときに自動的に持続します。
* `getMacro(name) string, error`与えられたマクロの値を返し、存在しない場合はエラーを返します。この関数はマクロの展開を行わないことに注意してください。

## 検索エントリー操作

* `setEntryEnum(ent, key, value)`は指定されたエントリに列挙された値を設定します。
* `getEntryEnum(ent, key)value, error`指定されたエントリから列挙された値を読み込みます。
* 指定されたエントリから指定された列挙された値を削除します。

## 一般ユーティリティー

* `len(val) int`はvalの長さを返します。
* これはパケットモジュールなどで生成されたIPと比較するのに適しています。
* `toMAC(string) MAC`は文字列をMACアドレスに変換します。
* `toString(val) string`はvalを文字列に変換します。
* `toInt(val) int64`はvalを可能なら整数に変換します。変換できない場合は0を返します。
* `toFloat(val) float64`は、可能であればVALを浮動小数点数に変換します。変換できない場合は0.0を返します。
* `toBool(val) bool`はvalをブール値に変換しようとします。変換できない場合は false を返す。0以外の数値と文字列 "y", "yes", "true" は真を返します。
* `toHumanSize(val) string`はvalを整数に変換し、人間が読めるバイト数で表現しようとします。
* 例えば、`toHumanSize(15127)`は"14.77KB"に変換されます。
* `typeOf(val) type` はvalの型を文字列で返します。
* `hashItems(val...) (uint64,ok)` 1つ以上の項目をサイファッシュアルゴリズムを用いてuint64にハッシュします。'ok' は、少なくとも1つの項目がハッシュ化できた場合に真であります。ハッシュ関数は実際にはスカラのみをハッシュすることができることに注意してください。

## 検索管理

Gravwellの検索システムの動作方法により、このセクションの一部の関数は検索構造体（パラメーターでは「search」と記述）を返し、その他の関数は検索ID（パラメーターでは「searchID」と記述）を返します。 各Search構造体には、`search.ID`としてアクセスできる検索IDが含まれています。

検索構造体は、検索からエントリをアクティブに読み取るために使用されますが、検索IDは、アタッチまたはその他の方法で管理する可能性のある非アクティブな検索を参照する傾向があります。

* `startBackgroundSearch（query、start、end）（search、err）`は、指定されたクエリ文字列を使用してバックグラウンド検索を作成し、「start」と「end」で指定された時間範囲で実行されます。 戻り値は検索構造体です。 これらの時間値は、時間ライブラリを使用して指定する必要があります。 デモの例を参照してください。
* `startSearch(query, start, end) (search, err)` は `startBackgroundSearch` と全く同じように動作しますが、検索のバックグラウンド化は行いません。
* `detachSearch(search)` は与えられた検索(Search構造体)を切り離します。これにより、バックグラウンド化されていない検索を自動的にクリーンアップできるようになります。
* `waitForSearch(search) error`は、与えられた検索が完全に実行されるのを待ち、エラーがあればそれを返します。
* `attachSearch（searchID）（search、error）`は、指定されたIDの検索にアタッチし、エントリなどの読み取りに使用できるSearch構造体を返します。
* `getSearchStatus(searchID) (string, error)` は、指定された検索の状態、例えば "SAVED "を返します。
* `getAvailableEntryCount（search）（uint64、bool、error）`は、指定された検索から読み取ることができるエントリの数、検索が完了したかどうかを指定するブール値、および問題が発生した場合のエラーを返します。
* `getEntries（search、start、end）（[] SearchEntry、error）`は、指定された検索から指定されたエントリをプルします。 `start`と` end`の境界は、`getAvailableEntryCount`関数で見つけることができます。
* `isSearchFinished（search）（bool、error）`は、指定された検索が完了した場合にtrueを返します。
* `executeSearch（query、start、end）（[] SearchEntry、error）`は検索を開始し、検索が完了するのを待ち、最大1万のエントリを取得し、検索から切り離してエントリを返します。
* `deleteSearch（searchID）error`は、指定されたIDの検索を削除します。
* `backgroundSearch（searchID）error`は、指定された検索をバックグラウンドに送信します。これは、後で手動で検査するために検索を「維持」するのに役立ちます。
* `saveSearch(searchID) error`は与えられた検索結果を長期保存用にマークします。 この呼び出しはクエリが完了するのを待たず、保存したものとしてマークするリクエストが失敗した場合にのみエラーを返します。
* `downloadSearch（searchID、format、start、end）（[] byte、error）`は、ユーザーがWebUIの[ダウンロード]ボタンをクリックしたかのように、指定された検索をダウンロードします。 `format`は、必要に応じて「json」、「csv」、「text」、「pcap」、または「lookupdata」のいずれかを含む文字列である必要があります。 `start`と` end`は時間の値です。
* `getDownloadHandle（searchID、format、start、end）（io.Reader、error）`は、ユーザーがWeb UIの[ダウンロード]ボタンをクリックしたかのように、指定された検索結果へのストリーミングハンドルを返します。返されるハンドルは、このドキュメントで後述するHTTPライブラリ関数での使用に適しています。

### 検索データタイプ

startSearchまたはstartBackgroundSearch関数を介して検索を実行すると、`search`データ型が返されます。`search`データ型には、次のメンバーが含まれています。

* `ID`-検索IDを含む文字列。 getSearchStatusやattachSearchなどの他の関数にこのメンバーを使用します
* `RenderMod`-検索にアタッチされたレンダラーを示す文字列。 raw、text、table、chart、fdgなどの場合があります。
* `SearchString`-リクエスト中に渡された検索文字列を含む文字列
* `SearchStart`-検索の開始タイムスタンプを含む文字列
* `SearchEnd`-検索の終了タイムスタンプを含む文字列
* `Background`-検索がバックグラウンド検索として開始されたかどうかを示すブール値
* `Name`-検索名を含むオプションの文字列。

## スクリプト情報

次の関数を使用すると、スケジュールされたスクリプト/検索と現在のスクリプトに関する情報を取得できます。

* `scheduledSearchInfo（）（[] ScheduledSearch、error）`は、現在実行中のスクリプトを含む、ユーザーに表示されるすべてのスケジュールされた検索またはスクリプトを返します。 ScheduledSearchタイプの説明については、以下を参照してください。
* `thisScriptID（）int32`は、現在実行中のスクリプトのID番号を返します。

ScheduledSearch構造には、次のフィールドが含まれています：

* `ID`-このスケジュールされた検索のID番号を含む32ビット整数
* `Owner`-スケジュールされた検索の所有者を表す32ビット整数
* `Groups`-このスケジュールされた検索を表示できる32ビットのグループIDの配列
* `Name`-スケジュールされた検索の名前を含む文字列
* `Description`-検索を説明する文字列
* `Labels`-検索にさらにラベルを付ける文字列の配列
* `Schedule`-実行スケジュールを定義するcron形式の文字列
* `Updated`-スケジュールされた検索が最後に変更されたときに設定されたタイムスタンプ
* `Disabled`-ブール値。検索が無効になっている場合はtrueに設定されます。
* `SearchString`-これがスケジュールされた検索（スクリプトではない）の場合、実行する検索文字列が含まれます。
* `Duration`-検索する過去の秒数（SearchStringが設定されている場合のみ）。
* `Script`-実行するスクリプトの内容
* `LastRun`-最新の実行が発生した時刻
* `LastRunDuration`-最新の実行にかかったナノ秒数
* `LastSearchIDs`-前回の実行中に作成された検索のIDを一覧表示する文字列の配列
* `LastError`-スクリプト/検索の前回の実行からのエラー結果を含む文字列

## インフラストラクチャ情報のクエリ

自動化スクリプトシステムは、APIコールを使用してGravwellインストールの状態を監視するために使用することもできます。これにより、プラットフォーム内のインゲスターの状態、システム負荷、インデクサーの接続性を監視することができます。以下の呼び出しにより、物理的なデプロイメントに関する情報を提供することができます。

* `ingesters`-各インデクサーのingesterステータスブロックを含むマップを返します。
* `indexers`-各インデクサーのウェルステータスを含むマップを返します。
* `indexerStates`-インデクサーが正常であるかどうかを示すブール値を含むマップを返します。
* `systemStates`-各システムのディスク、CPU、およびメモリの負荷を含むマップを返します。
* `systems`-CPU、メモリ、ディスク、ソフトウェアバージョンなどの物理システム情報を含むマップを返します。

## 結果の送信

スクリプトシステムは、スクリプトの結果を外部システムに送信するためのいくつかの方法を提供します。

次の関数は、基本的なHTTP機能を提供します：

* `httpGet（url）（string、error）`は、指定されたURLに対してHTTP GETリクエストを実行し、レスポンスの本文を文字列として返します。
* `httpPost（url、contentType、data）（response、error）`は、指定されたコンテンツタイプ（「application / json」など）と指定されたデータをPOST本文として指定されたURLに対してHTTPPOSTリクエストを実行します。

[net/http]ライブラリを使用すると、より複雑なHTTP操作が可能になります。 利用可能なものの説明については、 [anko document](anko.md) のパッケージドキュメントを参照するか、例については以下を参照してください。

ユーザーがGravwell内で個人の電子メール設定を構成している場合、`email`機能は電子メールを送信する非常に簡単な方法です。

* `email(from, to, subject, message, attachments...) error`はSMTP経由でメールを送信します。`from`フィールドは単なる文字列であり`to`はメールアドレスを含む文字列のスライスか、1つのメールアドレスを含む単一の文字列でなければなりません。フィールド`subject`と`message`もまた文字列で、メールの件名と本文を含む必要があります。attachmentsパラメータはオプションです。
  * 添付ファイルはバイト配列で送信することができ、自動ファイル名が付与されます。
  * 添付ファイルパラメータがマップの場合、キーはファイル名、値は添付ファイルです。
  * `emailWithCC(from, to, cc, bcc, subject, message, attachments...) error`はSMTP経由で電子メールを送信する。これは `email` 関数と全く同じように動作しますが、CCとBCCの受信者を指定することもできます。これらは単一の文字列(`"foo@example.com"`)または文字列の配列(`["foo@example.com", "bar@example.com"]`)のいずれかです。 

添付ファイル付きの電子メールの送信例：
```
#名前の付いた添付ファイルにテキストをポストだけの簡単な方法"attachment1"
email(`user@example.com`, `bob@accounting.org`, "Hey bob", "We need to talk", "A random attachement")

#Adding a list of attached files with specific names
mp = map[interface]interface{}
mp["stuff.txt"] = "this is some stuff"
mp["things.csv"] = CsvData

subj="Forgot attachments"
body="Forgot to send the stuff and things files"
email(`user@example.com`, `bob@accounting.org`, subj, body, mp)
```

次の機能は廃止されましたが、引き続き使用できるため、ユーザーの電子メールオプションを構成せずに電子メールを送信できます：

* `sendMail（hostname、port、username、password、from、to、subject、message）error`はSMTP経由でメールを送信します。`hostname`と`port`は使用するSMTPサーバーを指定します。`username`と`password`はサーバーへの認証用です。`from`フィールドは単なる文字列ですが、`to`フィールドはメールアドレスを含む文字列のスライスである必要があります。 `subject`フィールドと`message`フィールドも文字列であり、メールの件名と本文を含める必要があります。
* `sendMailTLS（hostname、port、username、password、from、to、subject、message、disableValidation）error`は、TLSを使用してSMTP経由で電子メールを送信します。`hostname`と`port`は使用するSMTPサーバーを指定します。`username`と`password`はサーバーへの認証用です。 `from`フィールドは単なる文字列ですが、`to`フィールドはメールアドレスを含む文字列のスライスである必要があります。`subject`フィールドと`message`フィールドも文字列であり、メールの件名と本文を含める必要があります。 disableValidation引数は、TLS証明書の検証を無効にするブール値です。 disableValidationをtrueに設定することは安全ではなく、電子メールクライアントをman-in-the-middle攻撃にさらす可能性があります。

## 通知の作成

スクリプトは、スクリプトの所有者を対象とした通知を作成する場合があります。通知は、整数ID、文字列メッセージ、オプションのHTTPリンク、および有効期限で構成されます。有効期限が過去、または24時間以上先の場合、Gravwellは代わりに有効期限を12時間に設定します。

	addSelfTargetedNotification(7, "これは私の通知です", "https://gravwell.io", time.Now().Add(3 * time.Hour)

通知IDは、通知を一意に識別します。 これにより、ユーザーは同じ通知IDで関数を再度呼び出すことで既存の通知を更新できますが、異なるIDを指定して複数の通知を同時に追加することもできます。

## エントリの作成と取り込み

次の関数を使用して、スクリプト内からインデクサーに新しいエントリを取り込むことができます。

* `newEntry（time.Time、data）Entry`は、指定されたタイムスタンプ（time.Now（）からのtime.Time）とデータ（多くの場合文字列）を使用して新しいエントリを返します。
* `ingestEntries（[] Entry、tag）error`は、指定されたタグ文字列を持つエントリの指定されたスライスを取り込みます。

`getEntries`関数によって返されるエントリは、必要に応じて変更し、`ingestEntries`を介して再取り込みするか、新しいエントリを一括で作成することができます。たとえば、以前の検索の一部のエントリをタグ「newtag」に再取り込みするには、次のようにします。

```
# Get the first 100 entries from the search
ents, _ = getEntries(mySearch, 0, 100)
ingestEntries(ents, "newtag")
```

他の条件に基づいて新しいエントリを取り込むには：

```
if condition == true {
	ents = make([]Entry)
	ent += newEntry(time.Now(), "Script condition triggered")
	ingestEntries(ents, "results")
}
```

## その他のネットワーク機能

一連のラッパー関数は、SSH および SFTP クライアントへのアクセスを提供します。これらの関数が返す構造体から呼び出すことができるメソッドについては、[ssh ライブラリのドキュメント](https://godoc.org/golang.org/x/crypto/ssh) および [sftp ライブラリのドキュメント](https://godoc.org/github.com/pkg/sftp) を参照してください。

重要：これらの関数によって返されるクライアントは、使用が終了したら、Close（）メソッドを介して閉じる必要があります。 以下の例を参照してください。

* `sftpConnectPassword（hostname、username、password、hostkey）（* sftp.Client、error）`は、指定されたユーザー名とパスワードを使用して、指定されたsshサーバー上でSFTPセッションを確立します。hostkeyパラメーターがnil以外の場合、ホストから期待される公開鍵として使用され、ホスト鍵の検証が実行されます。hostkeyパラメーターがnilの場合、ホストキーの検証はスキップされます。
* `sftpConnectKey（hostname、username、privkey、hostkey）（*sftp.Client、error）`は、指定された秘密鍵（文字列または[]バイト）を使用して認証することにより、指定されたユーザー名で指定されたsshサーバー上にSFTPセッションを確立します。
* `sshConnectPassword（hostname、username、password、hostkey）（* ssh.Client、error）`は、指定されたホスト名のSSHクライアントを返し、パスワードで認証します。クライアントを確立したら、通常、client.NewSession（）を呼び出して使用可能なセッションを確立する必要があることに注意してください。goのドキュメントまたは以下の例を参照してください。
* `sshConnectKey（hostname、username、privkey、hostkey）（* sftp.Client、error）`は、指定された秘密鍵（文字列または[]バイト）を使用して、指定されたユーザー名で指定されたSSHサーバーに接続します。

注：hostkeyパラメーターは、known_hosts/authorized_keys形式である必要があります。"ecdsa-sha2-nistp256AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOcrwoHMonZ/l3OJOGrKYLky2FHKItAmAMPzZUhZEgEb86NNaqfdAj4qmiBDqM04 / o7B45mc 〜/.ssh/known_hostsから適切なキーを抽出するには、`ssh-keygen -H -F <hostname>`を実行します。

telnetライブラリも利用できます。 直接ラッパーは提供されていませんが、スクリプトに `github.com/ziutek/telnet`をインポートし、telnet.Dialなどを呼び出すことで使用できます。以下の例は、telnetライブラリの簡単な使用法を示しています。

また、 [github.com/RackSec/srslog](https://github.com/RackSec/srslog) syslogパッケージへのアクセスも提供します。これにより、スクリプトはsyslogを介して通知を送信できます。 以下に例を示します。

最後に、低レベルのGo [ネットライブラリ](https://golang.org/pkg/net) が利用可能です。 リスナー機能は無効になっていますが、スクリプトはIP解析機能と、Dial、DialIP、DialTCPなどのダイヤル機能を使用できます。例については、以下を参照してください。

### SFTPの例

このスクリプトは、パスワード認証を使用し、ホストキーチェックを使用せずにSFTPサーバーに接続します。 ユーザー「sshtest」としてログインし、そのユーザーのホームディレクトリの内容を出力して、「hello.txt」という名前の新しいファイルを作成します。

```
conn, err = sftpConnectPassword("example.com:22", "sshtest", "foobar", nil)
if err != nil {
	println(err)
	return
}

w = conn.Walk("/home/sshtest")
for w.Step() {
	if w.Err() != nil {
		continue
	}
	println(w.Path())
}

f, err = conn.Create("/home/sshtest/hello.txt")
if err != nil {
	conn.Close()
	println(err)
	return
}
_, err = f.Write("Hello world!")
if err != nil {
	conn.Close()
	println(err)
	return
}

// check it's there
fi, err = conn.Lstat("hello.txt")
if err != nil {
	conn.Close()
	println(err)
	return
}
println(fi)
conn.Close()
```

### SSHの例

このスクリプトは、公開鍵認証を使用してサーバーに接続します。 ここでは、読みやすくするために秘密鍵ブロックが短縮されていることに注意してください。 また、ホストキーの検証も行います。 次に、`/bin/ps aux`を実行し、結果を出力します。

```
var bytes = import("bytes")

# `ssh-keygen -H -F <hostname`を介してこれを取得します
pubkey = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOcrwoHMonZ/l3OJOGrKYLky2FHKItAmAMPzZUhZEgEb86NNaqfdAj4qmiBDqM04/o7B45mcbjnkTYRuaIUwkno="

privkey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAr1vpoiftxU7Jj7P0bJIvgCQLTpM0tMrPmuuwvGMba/YyUO+A
[...]
5iMqZFaUncYZyOFE9hhHqY1xhgwxyjgCTeaI/J/KfbsaCSvrkeBq
-----END RSA PRIVATE KEY-----`

# 秘密鍵でログインします
conn, err = sshConnectKey("example.com:22", "sshtest", privkey, pubkey)
if err != nil {
	println(err)
	return
}

session, err = conn.NewSession()
if err != nil {
	println("Failed to create session: ", err)
	return err
}

// Once a Session is created, you can execute a single command on
// the remote side using the Run method.
var b = make(bytes.Buffer)
session.Stdout = &b
err = session.Run("/bin/ps aux")
if err != nil {
	println("Failed to run: " + err.Error())
	return err
}
println(b.String())

session.Close()
conn.Close()
```

### Telnetの例

このスクリプトはTelnetサーバーに接続し、パスワード「testing」を使用してrootとしてログインし、プロンプト(`$`)まで受信したすべてのものを出力します。

```
var telnet = import("github.com/ziutek/telnet")
t, err = telnet.Dial("tcp", "example.org:23")
if err != nil {
	println(err)
	return
}
t.SetUnixWriteMode(true)
b, err = t.ReadUntil(":")
println(toString(b))
t.Write("root\n")
b, err = t.ReadUntil(":")
println(toString(b))
t.Write("testing\n")
for {
	r, err = t.ReadUntil("$ ")
	print(toString(r))
}
```

### Syslogの例

Dial関数を使用してsyslogサーバーに接続し、Alertおよびその他の関数を呼び出してメッセージを送信します。 使用可能な関数のリストについては、 [the godoc](https://pkg.go.dev/github.com/RackSec/srslog?tab=doc)を参照してください。`NewLogger`および`New`関数は、ローカルシステムにのみメッセージを書き込むため、有効になっていないことに注意してください。これは通常、役に立ちません。

```
var syslog = import("github.com/RackSec/srslog")

c, err = syslog.Dial("tcp", "localhost:601", syslog.LOG_ALERT|syslog.LOG_DAEMON, "gravalerts")
if err != nil {
    println(err)
    return err
}
c.Alert("Detected something bad")
c.Close()
```

### ネットの例

このスクリプトは、TCPポート7778でローカルホストに接続し、接続に「foo」を書き込みます。`nc -l 7778`を実行することで、netcatに対してこれをテストできます。

```
var net = import("net")

c, err = net.Dial("tcp", "localhost:7778")
if err != nil {
	println(err)
	return err
}

c.Write("foo")
c.Close()
```

## 管理機能

スクリプトシステムは、様々なGravwell APIと自動的に対話するために使用できる管理機能にアクセスすることができます。

### システムのバックアップを実行する

システムのバックアップや`io.ReadCloser`の操作を容易にするために、ヘルパー関数`backup`が提供されます。関数`backup`の定義は以下の通りです:
```
backup(io.Writer, bool) error
```
backup関数は、提供されたライターにバックアップパッケージ全体を書き込み、失敗した場合はエラーを返します。 backup関数の第2引数は、保存された検索をバックアップに含めるかどうかを示します。 保存された検索は非常に大きくなる可能性があるので、バックアップパッケージが非常に大きくなる可能性があることに注意してください。

#### バックアップスクリプトの例

このサンプルバックアップスクリプトは、バックアップを実行し、結果のパッケージをTCPネットワークソケットに送信します。基本的に、バックアップをリモートTCPリスナーにストリーミングします。

```
var net = import("net")
c, err = net.Dial("tcp", "127.0.0.1:9876")
if err != nil {
	return err
}
err = backup(c, true)
if err != nil {
	c.Close()
	return err
}
return c.Close()
```

## 検索スクリプトの例

このスクリプトは、過去1日間にCloudflareの1.1.1.1DNSサービスと通信したIPを見つけるバックグラウンド検索を作成します。結果が見つからない場合、検索は削除されますが、結果がある場合、検索はGUIの[検索の管理]画面でユーザーが後で閲覧できるように残ります。

```
# タイムライブラリをインポートします。
var time = import("time")
# Define start and end times for the search
start = time.Now().Add(-24 * time.Hour)
end = time.Now()
# 検索を開始します。
s, err = startSearch("tag=netflow netflow Dst==1.1.1.1 Src | unique Src | table Src", start, end)
if err != nil {
	return err
}
printf("s.ID = %v\n", s.ID)
# 検索が終了するまで待ちます。
for {
	f, err = isSearchFinished(s)
	if err != nil {
		return err
	}
	if f {
		break
	}
	time.Sleep(1 * time.Second)
}
# Find out how many entries were returned
c, _, err = getAvailableEntryCount(s)
if err != nil {
	return err
}
printf("%v entries\n", c)
# エントリが返されない場合は、検索を削除します。
# そうでなければ、それをバックグラウンドにします。
if c == 0 {
	deleteSearch(s.ID)
} else {
	err = backgroundSearch(s.ID)
	if err != nil {
		return err
	}
}
# 実行終了時に常に検索から切り離します。
detachSearch(s)
```

スクリプトをテストするには、その内容を`/tmp/script`というファイルに貼り付けます。次に、Gravwell CLIツールを使ってスクリプトを実行します。これにより、スクリプトの変更を簡単にテストすることができます。

```
$ gravwell -s gravwell.example.org watch script
Username: <user>
Password: <password>
script file path> /tmp/script
0 entries
deleting 910782920
Hit [enter] to re-run, or [q][enter] to cancel

s.ID = 285338682
1 entries
Hit [enter] to re-run, or [q][enter] to cancel
```

上記の例は、スクリプトが2回実行されたことを示しています。最初の実行では、結果が見つからず、検索が削除されました。2番目では、1つのエントリが返されたため、検索はそのまま残されました。

CLIでスクリプトをテストするときは、古い検索を削除するように注意してください。スケジュールされた検索システムは、CLIではなく、古い検索を自動的に削除します。

## スクリプトでのHTTPの使用

最近の多くのコンピューターシステムは、HTTP要求を使用してアクションをトリガーします。Gravwellスクリプトは、次の基本的なHTTP操作を提供します。

* httpGet(url)
* httpPost(url, contentType, data)

これらの関数とAnkoに含まれているJSONライブラリを使用して、前のサンプルスクリプトを変更できます。スクリプトは、後で閲覧するために検索をバックグラウンドで実行するのではなく、1.1.1.1と通信したIPのリストを作成し、それらをJSONとしてエンコードし、HTTPPOSTを実行してサーバーに送信します。

```
var time = import("time")
var json = import("encoding/json")
var fmt = import("fmt")

# 検索を開始
start = time.Now().Add(-24 * time.Hour)
end = time.Now()
s, err = startSearch("tag=netflow netflow Dst==1.1.1.1 Src | unique Src", start, end)
if err != nil {
	return err
}
# 検索が終了するのを待ちます
for {
	f, err = isSearchFinished(s)
	if err != nil {
		return err
	}
	if f {
		break
	}
	time.Sleep(1*time.Second)
}
# 返されたエントリの数を取得します
c, _, err = getAvailableEntryCount(s)
if err != nil {
	return err
}
# エントリがない場合は戻ります
if c == 0 {
	return nil
}
# エントリをフェッチします
ents, err = getEntries(s, 0, c)
if err != nil {
	return err
}
# IPのリストを作成します
ips = []
for i = 0; i < len(ents); i++ {
	src, err = getEntryEnum(ents[i], "Src")
	if err != nil {
		continue
	}
	ips = ips + fmt.Sprintf("%v", src)
}
# IPリストをJSONとしてエンコードします
encoded, err = json.Marshal(ips)
if err != nil {
	return err
}
# HTTPサーバーに投稿します
httpPost("http://example.org:3002/", "application/json", encoded)
detachSearch(s)
```

検索結果が非常に大きく、メモリに保持するには大きすぎる場合があります。 「net/http」ライブラリを「getDownloadHandle」関数と組み合わせると、Gravwell検索からHTTP POST/PUTリクエストに結果を直接ストリーミングできます。また、Cookieまたは追加のヘッダーを設定することもできます。

```
var http = import("net/http")
var time = import("time")
var bytes = import("bytes")

start = time.Now().Add(-72 * time.Hour)
end = time.Now()
s, err = startSearch("tag=gravwell", start, end)
if err != nil {
		return err
}
for {
		f, err = isSearchFinished(s)
		if err != nil {
				return err
		}
		if f {
				break
		}
		time.Sleep(1*time.Second)
}

# 検索結果の処理します
rhandle, err = getDownloadHandle(s.ID, "text", start, end)
if err != nil {
		return err
}
# リクエストを作成します
req, err = http.NewRequest("POST", "http://example.org:3002/", rhandle)
if err != nil {
		return err
}
# ヘッダーを追加します
req.Header.Add("My-Header", "gravwell")
# クッキーを追加します
cookie = make(http.Cookie)
cookie.Name = "foo"
cookie.Value = "bar"
req.AddCookie(&cookie)

# 実際のリクエストを実行します
resp, err = http.DefaultClient.Do(req)
detachSearch(s)
return err
```

## CSVヘルパー

CSVは、リソースの非常に一般的なエクスポート形式であり、Gravwellからデータを一般的に取得します。`encoding/csv`によって提供されるCSVライブラリは堅牢で柔軟性がありますが、少し冗長です。 Gravwellスクリプトシステム内で使用するためのよりシンプルなインターフェイスを提供するために、CSVライターをラップしました。 簡略化されたCSVビルダーを作成するには、`encoding/csv`パッケージをインポートし、`NewWriter`を呼び出す代わりに、引数なしで`NewBuilder`を呼び出します。

CSVビルダーは、独自の内部バッファーを管理し、`Flush`の実行時にバイト配列を返します。これにより、エクスポートまたは保存するためのCSVを構築するプロセスを簡素化できます。簡略化されたcsvBuilderを使用して、2つのテーブル列で構成されるリソースを作成するスクリプトの例を次に示します。
```
csv = import("encoding/csv")
time = import("time")

query = `tag=pcap packet ipv4.SrcIP ipv4.DstIP ipv4.Length | sum Length by SrcIP DstIP | table SrcIP DstIP sum`
end = time.Now()
start = end.Add(-1 * time.Hour)

ents, err = executeSearch(query, start, end)
if err != nil {
	return err
}

bldr = csv.NewBuilder()
err = bldr.WriteHeaders([`src`, `dst`, `total`])
if err != nil {
	return err
}

for ent in ents {
	src, err = getEntryEnum(ent, "SrcIP")
	if err != nil {
		return err
	}
	dst, err = getEntryEnum(ent, "DstIP")
	if err != nil {
		return err
	}
	sum, err = getEntryEnum(ent, "sum")
	if err != nil {
		return err
	}
	err = bldr.Write([src, dst, sum])
	if err != nil {
		return err
	}
}

buff, err = bldr.Flush()
if err != nil {
	return err
}
return setResource("csv", buff)
```

## SQL Usage

Gravwellスクリプトシステムでは、自動化スクリプトが外部のSQLデータベースと対話できるように、SQLデータベースパッケージを公開しています。 SQLライブラリを使用するには、Gravwellバージョン4.1.6以降が必要です。

現在、スクリプトシステムは以下のデータベースドライバーをサポートしています:

* MySQL/MariaDB
* Postgresql
* MSSQL
* OracleDB

SQLインターフェースの使用は、Go [database/sql](https://golang.org/pkg/database/sql/)パッケージを直接インポートした`database/sql`パッケージによって行われます。

Gravwellの`database/sql`パッケージには、`ExtractRows`というGo sqlパッケージには含まれないヘルパー関数も含まれています。 `ExtractRows`ヘルパー関数はSQLの結果の行を後の操作のために文字列の一部に変換するのを簡単にします。 `ExtractRows`関数のインターフェイスは以下のとおりです:

`ExtractRows(*sql.Row, columnCount) ([]string, error)`

検索スクリプトでSQLリソースを使用すると、強力なツールになります。 しかし、SQLインターフェースは冗長で、適切に使用するには注意が必要です。

特定のAPIの使用方法についての詳細なドキュメントは、Goの公式ドキュメントを参照してください。

### SQL Example Script

次の例は、リモートのMariaDB SQLデータベースに問い合わせ、その結果をCSVリソースにするスクリプトです。:

```
var sql = import("database/sql")
var csv = import("encoding/csv")

MinVer(4, 1, 6)

csvbldr = csv.NewBuilder()

// connect to the DB
db, err = sql.Open("mysql", "root:password@tcp(172.19.0.2)/foo")
if err != nil {
	return err
}

// Query example
rows, err = db.Query("SELECT * FROM bar WHERE name!=?", "stuff")
if err != nil {
	return err
}

// Get Headers
headers, err = rows.Columns()
if err != nil {
	return err
}
csvbldr.WriteHeaders(headers)

for rows.Next() {
	vals, err = sql.ExtractRows(rows, len(headers)) //pass in the number of columns
	if err != nil {
		return err
	}
	err = csvbldr.Write(vals)
	if err != nil {
		return err
	}
}

err = rows.Err()
if err != nil {
	return err
}
err = rows.Close()
if err != nil {
	return err
}

err = db.Close()
if err != nil {
	return err
}

data, err = csvbldr.Flush()
if err != nil {
	return err
}

return setResource("foobar", data)
```


## IPExistデータセット

[ipexist](#!search/ipexist/ipexist.md)検索モジュールは、IPv4アドレスがセットに存在するかどうかをテストするように設計されています。このモジュールは、速度という1つのことだけを目的として設計された単純なフィルタリングモジュールです。 内部的には、 `ipexist`は高度に最適化されたビットマップシステムを使用しているため、適度なマシンがそのフィルターシステム内のIPv4アドレス空間全体を表すことができます。 IPExistは、[iplookup](#!search/iplookup/iplookup.md)モジュールを使用してより高価なルックアップを実行する前に、脅威リストを保持し、非常に大きなデータセットに対して初期フィルタリング操作を実行するための優れたツールです。

Gravwellスクリプトシステムはipexistビルダー関数にアクセスできるため、既存のデータから高速IPメンバーシップテーブルを生成できます。 ipexistビルダー関数はオープンソースであり、[github](https://github.com/gravwell/ipexist)で利用できます。 以下は、クエリを使用してIPメンバーシップリソースを生成する基本的なスクリプトです。

```
ipexist = import("github.com/gravwell/ipexist")
bytes = import("bytes")
time = import("time")

query = `tag=ipfix ipfix port==22 src dst | stats count by src dst | table`
end = time.Now()
start = end.Add(-1 * time.Hour)

ipe = ipexist.New()

ents, err = executeSearch(query, start, end)
if err != nil {
	return err
}

for ent in ents {
	ip, err = getEntryEnum(ent, "src")
	if err != nil {
		return err
	}
	err = ipe.AddIP(toIP(ip))
	if err != nil {
		return err
	}
	
	ip, err = getEntryEnum(ent, "dst")
	if err != nil {
		 return err
	}
	err = ipe.AddIP(toIP(ip))
	if err != nil {
		return err
	}
}

bb = bytes.NewBuffer(nil)
err = ipe.Encode(bb)
if err != nil {
	return err
}
ipe.Close()
buff = bb.Bytes()
println("buffer", len(buff))
return setResource("sshusers", buff)
```
## Gravwell Client Usage

関数 `getClient` は、現在のユーザとしてログインして同期している新しい [Client](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client) オブジェクトへのポインタを返します。通常の動作状態であれば、この新しいクライアントはすぐに使用することができます。  ただし、ネットワーク障害やシステムアップグレードなどにより、スクリプト運用中にGravwellのウェブサーバーが利用できなくなる可能性があります。  そのため、スクリプトでは、[TestLogin()](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client.TestLogin)メソッドを使用して、クライアントの接続状態をテストすることをお勧めします。

このスクリプト例では、クライアントを取得し、リモートサーバーへのTCP接続を行い、Gravwellシステムの[バックアップ](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client#Client.Backup)を実行し、リモートTCP接続を介してバックアップファイルを送信します:

```
net = import("net")
time = import("time")

BACKUP_SERVER=`10.0.0.1:5555`

cli = getClient()
if cli == nil {
	return "Failed to get client"
}
err = cli.TestLogin()
if err != nil {
	return err
}
// Backup requests can take some time, increase the client request timeout
err = cli.SetRequestTimeout(10*time.Minute)
if err != nil {
	return err
}

conn, err = net.Dial("tcp", BACKUP_SERVER)
if err != nil {
	return err
}

err = cli.Backup(conn, false)
if err != nil {
	return err
}
err = conn.Close()
if err != nil {
	return err
}

return cli.Close()
``` 

注：Gravwellクライアントのデフォルトのリクエストタイムアウトは5秒である。システムバックアップのような長時間のリクエストの場合は、このタイムアウトを増やすべきだが、長時間のリクエストが完了したら元のタイムアウトに戻すのがベストプラクティスであることに注意してほしい。