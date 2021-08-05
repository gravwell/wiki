# Gravwell CLI

Gravwellコマンドラインクライアントを使用すると、Gravwellをリモートで管理し、検索を実行できます（レンダラーのサポートは限られています）。管理者は、フルWebブラウザを使用せずにユーザーを管理し、クラスタの状態を監視できます。ユーザーは検索を実行し、他のツールを使用して追加の分析のために結果をファイルに簡単にエクスポートすることができます。

コマンドラインクライアントは、検索結果をレンダリングできないという点でわずかに制限があります（たとえば、CLIは端末にチャートを描画できないため、chartモジュールを使用する検索のレンダリングを拒否します）。ただし、バックグラウンド検索を発行する場合、CLIはすべてのレンダラーモジュールにアクセスできます。上級ユーザーがリモートログインして、大規模な検索を開始してフルブラウザで表示できるようにする場合に便利です。サイト。

通常のインストールでは、CLIツールは/usr/sbin/gravwellとしてインストールされます。 -hフラグを渡すと、どこから始めればよいかがわかります。デフォルトでは、GravwellクライアントはWebサーバーがローカルマシン上で待機していると想定し、-sフラグを指定して他のWebサーバーまたはリモートGravwellインスタンスを指すようにします。


```
gravwell options
  -b	Background the search
  -debug string
    	Enable JSON output debugging to a file
  -f string
    	Output format: "simple" or "grid" (default "grid")
  -insecure
    	Do NOT enforce webserver certificates, TLS operates in insecure mode
  -insecure-no-https
    	Use insecure HTTP connection, passwords are shipped plaintext
  -o string
    	Output to file rather than stdio
  -query string
    	Query string
  -r	Raw output, no pretty print
  -s string
    	Address and port of Gravwell webserver
  -si
    	Enable additional search information output
  -t	Disable sessions, always require logins
  -time string
    	Query time range
  -v	Output client version number and exit
  -watch-interval int
    	Watch update interval
OPTIONS:
	shell: enter the interactive shell
	state: Show the state of configured indexers
	desc: Show the description of each connected indexer
	storage: Show the current state of each connected indexer
	systems: Show performance metrics of each addr
	indexes: Show size of each index
	ingesters: Show activity and performance metrics of each ingester
	sessions: Show your other active sessions
	notifications: Show your active notifications
	search: Perform search
	attach: Reattaches to existing search
	download: Download search results in a packaged format
	admin: Perform admin actions
	user: Perform user actions
	dashboards: Perform dashboard actions
	logout: Logout of the current session
	logoutall: Logout all sessions using your UID
	list_searches: list_searches
	search_ctrl: Issue search control command
	resource: Create and manage resources
	macro: Manage search macros
	kits: Manage and upload kits
	userfiles: Manage user files
	templates: Manage templates
	pivots: Manage pivots
	ingest: Ingest entries directly
	scheduled_search: Manage scheduled searches
	script: Run a script
	help: Display available commands
MODIFIERS:
	watch: Continually show results of stats commands

EXAMPLE: gravwell -s=localhost state
```

GravwellクライアントはGravwellで検索を実行し、その出力を他のツールにフィードするための優れた方法でもあります。たとえば、セキュリティデータを処理するためのカスタムプログラムがあるが、ログエントリをGravwellに格納したい場合は、CLIクライアントを使用してバックグラウンドクエリを実行してエントリを抽出し、その結果をカスタムプログラム用のファイルに保存します。読む。

## 対話的にCLIを使う

Gravwell CLIクライアントは、商用スイッチに見られるものと同様の対話式シェルを提供します。それは異なる「メニュー」レベルを持っています。たとえば、トップレベルのメニューから、ダッシュボードを管理するためのコマンドを含む「ダッシュボード」サブメニューを選択できます。このセクションでは、クライアントを対話的に使用するための基本について説明します。

### 接続とログイン

デフォルトでは、クライアントはGravwell Webサーバーがlocalhost：443をリッスンしていると想定します。これが正しい場合は、単にgravwellコマンドを実行して接続できます。クライアントはあなたのユーザ名とパスワードの入力を促し、それからプロンプトを表示します。

```
$ gravwell
Username:  admin
Password:  changeme
#> 
```

Webサーバが別のホストにある場合は、-sフラグを使用してホスト名とポートを指定します。 gravwell -s webserver.example.com:4443

Webサーバーに自己署名TLS証明書がインストールされている場合は、TLS検証を無効にしてもHTTPSを使用するために-insecureフラグを追加する必要があります。

WebサーバーにTLS証明書がインストールされていない場合は、HTTP専用モードを使用するために-insecure-no-httpsフラグを追加してください。 これは安全ではありません。パスワードはプレーンテキストでサーバーに送信されます。

### 利用可能なコマンドの一覧表示

helpコマンドは現在のメニューレベルで利用可能なコマンドをリストします。 クライアントを起動した直後は、最上位レベルになります。

```
#>  help
shell                enter the interactive shell
state                Show the state of configured indexers
desc                 Show the description of each connected indexer
storage              Show the current state of each connected indexer
systems              Show performance metrics of each addr
indexes              Show size of each index
ingesters            Show activity and performance metrics of each ingester
sessions             Show your other active sessions
notifications        Show your active notifications
search               Perform search
attach               Reattaches to existing search
download             Download search results in a packaged format
admin                Perform admin actions
user                 Perform user actions
dashboards           Perform dashboard actions
logout               Logout of the current session
logoutall            Logout all sessions using your UID
list_searches        list_searches
search_ctrl          Issue search control command
resource             Create and manage resources
macro                Manage search macros
kits                 Manage and upload kits
userfiles            Manage user files
templates            Manage templates
pivots               Manage pivots
ingest               Ingest entries directly
scheduled_search     Manage scheduled searches
script               Run a script
help                 Display available commands
```

リストされている項目のいくつかはコマンドです。

```
#>  state
+----------------------+----------+
|               System |    State |
+======================+==========+
|    192.168.2.60:9404 |       OK |
+----------------------+----------+
|            webserver |       OK |
+----------------------+----------+
```

他のものはそれら自身のコマンドを含むメニューです。 以下の例では、 'ダッシュボード'メニューを選択し、利用可能なコマンドを一覧表示して、 'list'コマンドを実行します。

```
#>  dashboards
dashboards>  help
list                	List available user dashboards
mine                	List dashboards owned by you
del                 	Delete a dashboard available user dashboards
clone               	Clone a dashboard to enable ownership and editing
dashboards>  list
+---------+-------+-----------------+------------------------------+----------+-----------+
|    Name |    ID |     Description |                      Created |    Owner |    Groups |
+=========+=======+=================+==============================+==========+===========+
|     Foo |    10 |    My dashboard |    2019-04-15T12:19:49-06:00 |    admin |           |
+---------+-------+-----------------+------------------------------+----------+-----------+
dashboards>  
```
## キットの管理

キットの管理は、サブメニューの「キット」から行います:

```
#>  kits
kits>  help
list                	List installed kits
get                 	Get kit details
uninstall           	Remove an installed kit
install             	Install an uploaded kit
upload              	Upload a new kit
pull                	Pull a kit from a remote kitserver
build               	Build a kit from existing queries, resources, and dashboards
rebuild             	Rebuild a kit from a previous build request, incrementing the version
repack              	Repack a currently-installed kit, optionally changing attributes.
remote              	List available kits from the remote kitserver
```

### キットのインストール

キットをインストールするには、キットサーバーからインストールする方法と、ローカルファイルをアップロードしてインストールする方法があります。いずれの場合も、インストール作業は、ウェブサーバー上にキットをステージングしてからインストールするという2つのステップで構成されています。

キットサーバからキットをインストールするには、まず `remote` コマンドを使ってサーバ上のキットをリストアップします。目的のキットのUUIDをコピーしてから、`pull`コマンドを実行し、プロンプトが表示されたらUUIDをペースト操作で貼り付けます。これでキットがダウンロードされ、ステージングされます。キットがステージングされたら、`install`コマンドを実行し、ステージングされたキットを選択すると、インストールプロセスが開始されます。

ローカルファイルからキットをインストールするには、`upload`コマンドを実行し、プロンプトが表示されたらファイルのパスを入力します。これでキットがステージングされます。次に、`install`コマンドを実行して、ステージングされたキットを選択すると、インストールプロセスが開始されます。

The `install` command will ask the user if the kit should be installed with default options; if the user answers "no", they must select each option individually:
`install`コマンドを入れると、キットをデフォルトのオプションでインストールするかどうかが尋ねられます。「no」と答えた場合は、インストールするオプションを個別に選択する必要があります:

* Global: (管理者専用) trueの場合、キットアイテムはすべてのユーザーに表示されます（デフォルト：false）
* Overwrite existing: trueの場合、インストール時に、キットの内容と競合する既存のオブジェクトを上書きします（デフォルト：false）
* Allow unsigned: 署名されていないキットをインストールする時には、このオプションｈをtrueに設定する必要があります（デフォルト：false）
* Item labels: キット内のアイテムに適用される追加ラベルのオプションリスト（デフォルト：なし）
* Kit labels: キット自体に適用される追加ラベルのオプションリスト（デフォルト：なし）
* Group: キットの内容を見ることができるメンバーのグループの選択（デフォルト：なし）

### キットのビルド

`build`コマンドを使うと、キットの構築手順が案内されます。キットのID、名前、説明、およびバージョンを入力するように求められます。IDは、衝突を避けるために、"io.gravwell.networkenrichment "のような "名前付き "のIDでなければならないことに注意してください。名前と説明のフィールドには何を入力してもかまいません。また、バージョンは整数でなければなりません。

基本的な情報を収集した後、CLIはキットに含まれるべきオブジェクトを選択するようユーザに促します。ダッシュボード、テンプレート、アクショナブルなどのプロンプトが次々と表示されます。プロンプトが出た時に、そのタイプのオブジェクトに何も含めたくないこと時には、何も入力しないでEnterキーを押してください。

すべてのオブジェクトが選択されると、CLIは生成されたキットをダウンロードする場所を尋ねるプロンプトを表示します。ファイル名を入力することも、ディレクトリのみを入力することもできます。完了すると、新しいキットの場所が表示されます。

### キットの再ビルド

`rebuild`コマンドは、以前にビルドしたキットのアップデート版をビルドするのに使います。実行すると、既にビルド済みのキットがリストアップされます。キットを選択すると、必要に応じて、キットに含まれるアイテムのリストを変更するオプションがUIに表示されます。その後、キットのバージョン番号が自動的に増加し、新しいキットファイルの出力が生成され、（`install`コマンドの時と同様に）生成されたファイルをどこに保存する場所を尋ねるプロンプトが表示されます。

### キットの再パッケージ

`repack`コマンドは、`rebuild`コマンドとよく似た動作をしますが、*以前にビルドされた*キットを再構築するのではなく、*現在インストールされている*キットの一つを再パッケージ化します。これは、Gravwellや他のユーザーから入手した既存のキットを修正したい場合に便利です。キットをインストールして、変更が必要な項目を修正してから、そのキットに対してrepackコマンドを実行します。


## CLIによる検索

searchコマンドはフォアグラウンドで検索を実行します。

```
#>  search
query>  tag=* count
time range> -1h
count 100
1/1
Press q[Enter] to quit, [Entry] to continue

Total Items: 1
101 stats records from Apr 15 12:39:37.000 <-> Apr 15 13:39:38.000
count 100.00/1.00 61.66 KB/616 B 8.109585ms
```

検索の結果を保存したい場合は、検索をバックグラウンドで実行するように指定する '-b'フラグを指定してクライアントを実行し、次にsearchおよびdownloadコマンドを使用して検索を実行し、結果を保存します：

```
$ gravwell -b
#>  search
query>  tag=* json state=="NM"
time range> -1h
Background search with ID 065015787 launched
#>  download
+--------------+---------+----------+------------+---------------------+------------+
|    Search ID |    User |    Group |      State |    Attached Clients |    Storage |
+==============+=========+==========+============+=====================+============+
|    065015787 |       1 |        0 |    DORMANT |                   0 |    1.56 KB |
+--------------+---------+----------+------------+---------------------+------------+
search ID>  065015787
Available formats:
json
text
csv
format>  text
Save Location (default: /tmp)>  /tmp/nm.txt
Saving to  /tmp/nm.txt
#>  
```


## 管理者

Gravwellクライアントはadminサブメニューでシステムを管理するための多くのコマンドを実装しています。

```
#>  admin
admin>  help
add_user            Add a new user
impersonate_user    Impersonate an existing users
del_user            Delete an existing user
get_user            Get an existing users details
update_user         Update an existing user
list_users          List all users
lock_user           Lock a user account
user_activity       Show a specific users activity
user_sessions       Show all open sessions
change_user_pwd     Change a users password
change_admin        Set a users admin status
add_group           Create a new group
del_group           Delete an existing group
list_groups         Lists all existing groups
list_group_users    Lists all members of an existing group
update_group        Update an existing group
add_users_group     Add users to an existing group
del_users_group     Delete users from an existing group
add_user_groups     Add user to existing groups
del_user_groups     Delete a user from groups
get_log_level       Get the webservers current logging level
set_log_level       Set the webservers current logging level
all_dashboards      Get all dashboards for all users
del_dashboard       Delete a dashboard owned by another user
license_info        Display license information
license_sku         Display license SKU
license_serial      Display license Serial Number
license_update      Upload a new license
list_queries        List all queries (active and saved) for all users
delete_queries      Delete any query (active or saved) for any user
list_users_storage  List all users current storage usage
add_indexer         Add another indexer to the configuration
list_kits           List all kits across all users
uninstall_kit       Uninstall a kit owned by any user
list_extractions    List installed autoextractors
add_extraction      Add a new autoextractor
delete_extraction   Delete an installed autoextractor
update_extraction   Update an installed autoextractor
sync_extractions    Force a sync of installed autoextractors to indexers
```

ユーザー/グループ管理に加えて、管理メニューにはダッシュボード、キット、およびシステム上の他のユーザーに属する他のオブジェクトを管理するためのツールもあります。


## CLIの例

### インデクサーの健全性チェック

```
$ gravwell state
+----------------------+----------+
|               System |    State |
+======================+==========+
|    10.0.0.2:9404     |       OK |
+----------------------+----------+
|    10.0.0.3:9404     |       OK |
+----------------------+----------+
|    webserver         |       OK |
+----------------------+----------+
```

すべてのコマンドの出力は、テーブルやフォーマットなしで「raw」に設定できます。 Gravwellのデータを他のツールやスクリプトに渡すと、生の出力を要約しやすくなります。

```
$ gravwell -r state
10.0.0.3:9404 OK
webserver     OK
10.0.0.2:9404 OK
```

### インデクサーウェルとストレージサイズの表示

```
$ gravwell -r indexes
10.0.0.2:9404 default /mnt/storage/gravwell/default 14.8 MB 93.76 K
10.0.0.2:9404 pcap /mnt/storage/gravwell/pcap 3.6 MB 29.63 K
10.0.0.2:9404 bench /mnt/storage/gravwell/bench 142.5 GB 686.66 M
10.0.0.2:9404 reddit /mnt/storage/gravwell/reddit 34.3 GB 73.72 M
10.0.0.2:9404 fcc /mnt/storage/gravwell/fcc 21.7 GB 11.09 M
10.0.0.2:9404 raw /mnt/storage/gravwell/raw 76.3 KB 0
10.0.0.2:9404 syslog /mnt/storage/gravwell/syslog 60.2 MB 406.62 K
10.0.0.3:9404 default /mnt/storage/gravwell/default 12.2 MB 77.56 K
10.0.0.3:9404 reddit /mnt/storage/gravwell/reddit 34.3 GB 73.66 M
10.0.0.3:9404 fcc /mnt/storage/gravwell/fcc 21.6 GB 11.06 M
10.0.0.3:9404 pcap /mnt/storage/gravwell/pcap 5.1 MB 44.85 K
10.0.0.3:9404 syslog /mnt/storage/gravwell/syslog 79.6 MB 536.86 K
10.0.0.3:9404 raw /mnt/storage/gravwell/raw 76.3 KB 0
10.0.0.3:9404 bench /mnt/storage/gravwell/bench 136.5 GB 658.69 M
```

### リモートインジェスターを表示する

```
$ gravwell -r ingesters
10.0.0.2:9404
        tcp://10.0.0.1:49544 111h27m33.8s [reddit] 5.16 M 2.68 GB
        tcp://192.210.192.202:45578 34m52.1s [pcap] 62.00 3.69 KB
        tcp://192.210.192.202:43369 34m51.9s [kernel] 1.1 K 121.43 KB
10.0.0.3:9404
        tcp://10.0.0.1:49770 111h27m0.01s [reddit] 5.24 M 2.72 GB
        tcp://192.210.192.202:43364 34m52.6s [pcap] 119.00 6.93 KB
        tcp://192.210.192.202:43368 34m51.9s [kernel] 1.33 K 141.57 KB
```

### スクリプトを実行する

```
$ gravwell script
script file path>  /tmp/myscript.ank
```
