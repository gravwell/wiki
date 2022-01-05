# ユーザー作成オブジェクトの管理

ユーザーは、Gravwellシステム内でさまざまなオブジェクトを作成できます:

* リソース
* 保存された検索/バックグラウンド検索
* スケジュールされた検索/スクリプト
* ダッシュボード
* テンプレート
* ユーザーファイル

現時点では、これらのオブジェクトを管理者として管理するためのGUIユーティリティはありません。 ただし、 [Gravwellコマンドラインクライアント](#!cli/cli.md)は、**admin**サブメニューのオプションを使用して、これらすべてのオブジェクトタイプを一覧表示、削除、および変更できます。

これらの管理オプションにアクセスするには、クライアントを実行し、管理者ユーザーとしてログインして、管理メニューに入ります:

```
$ ./client -s gravwell.example.org
Username:  admin
Password:  
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
list_extractions    List installed autoextractors
add_extraction      Add a new autoextractor
delete_extraction   Delete an installed autoextractor
update_extraction   Update an installed autoextractor
sync_extractions    Force a sync of installed autoextractors to indexers
resource            Create and manage resources
scheduled_search    Manage scheduled searches
templates           Manage templates
pivots              Manage actionables
userfiles           Manage user files
kits                Manage and upload kits
admin>
```

このセクションの残りの部分では、各オブジェクトタイプの管理オプションについて簡単に説明します。

## ダッシュボードの管理

システム上のすべてのダッシュボードを一覧表示するには、**admin**メニューから`all_dashboards`コマンドを実行します。

ダッシュボードを削除するには、**admin**メニューから`del_dashboard`コマンドを実行します。

## 検索の管理

システム上のすべての検索（保存済み、バックグラウンド、またはアクティブ）を一覧表示するには、**admin**メニューから`list_queries`コマンドを実行します。

クエリを削除するには、`delete_queries`コマンドを実行します。

## リソースの管理

adminサブメニューには、通常のリソースメニューで使用可能なコマンドを反映したコマンドでリソースを管理するための独自のサブメニューが含まれています:

```
admin>  resource
resource>  help
list                	List available resources
create              	Create a new resource
update              	Upload new data to a resource
delete              	Delete a resource
updatemeta          	Update resource metadata
resource>  
```

このメニューから、管理者はシステム上の*すべての*リソースを一覧表示したり、リソースの内容を変更したり、名前/説明/所有権を変更したり、削除したりできます。

## スケジュールされた検索の管理

adminサブメニューには、スケジュールされた検索を管理するための独自のサブメニューが含まれています:

```
admin>  scheduled_search
scheduled search>  help
list                	List saved searches
listall             	List all saved searches
create              	Create a new scheduled search
createscript        	Create a new scheduled search w/ script
delete              	Delete a scheduled search
```

管理者は、このメニューから、自分の検索だけでなく、システム上の*すべての*スケジュールされた検索を管理できます。

## テンプレート/アクショナブルの管理

テンプレートとアクションテーブル（ここでは「ピボット」と呼びます）は、それぞれ管理メニュー内にサブメニュー（`templates`と`pivots`）があり、管理者向けの同じコマンド群が用意されています。

```
admin>  templates
template>  help
list                	List templates
create              	Create a new template
update              	Upload new contents to a template
delete              	Delete a template
print               	Print template contents
updatemeta          	Update template metadata
template>  quit
admin>  pivots
pivot>  help
list                	List actionables
create              	Create a new actionable
update              	Upload new contents to an actionable
delete              	Delete an actionable
print               	Print actionable contents
updatemeta          	Update actionabl metadata
pivot>
```

これらのコマンドを使用して、システム上の任意のテンプレートまたはピボットに影響を与えることができます。

## ユーザーファイルの管理

テンプレートやリソースなどと同様に、ユーザーファイルにも管理メニュー内に管理管理用のサブメニューがあります。管理メニュー内で実行されるコマンドは、システム全体の任意のユーザーファイルで操作できます。

```
admin>  userfiles
userfile>  help
list                	List available userfiles
add                 	Add a new userfile
update              	Update an existing userfile
del                 	Delete a userfile
get                 	Download a userfile
userfile> 
```
