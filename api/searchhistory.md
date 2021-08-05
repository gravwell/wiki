# 検索履歴

`/api/searchhistory`にあるRESTAPI

検索履歴APIは、ユーザー、グループ、またはそれらの組み合わせが起動した検索のリストをプルバックするために使用されます。 結果には、検索を開始したユーザー、所有しているグループ、開始された日時、検索を表す2つの文字列（ユーザーが実際に入力した内容、バックエンドが処理した内容）などの基本情報が表示されます。

## 基本的なAPIの概要

ここでの基本的なアクションは、`/api/searchhistory/{cmd}/{id}`に対してGETを実行することです。ここで、`cmd`は必要な検索のセットを表し、`id`はその検索に関連するIDを表します。 たとえば、UID 1のユーザーが所有するすべての検索が必要な場合は、`/api/searchhistory/user/1`でGETを実行し、GID 4のグループが所有するすべての検索が必要な場合は、`/api/searchhistory/group/4`でGETを実行します。`/api/searchhistory`でGETを実行するだけで、現在のユーザーの履歴が返されます。

「all」コマンドを使用して、特定のUIDがアクセスできるすべての検索を要求できます。 つまり、UIDが所有するすべての検索と、彼がメンバーになっているグループが所有するすべての検索を取得できます。 たとえば、`/api/searchhistory/all/1`のGETは、UID 1のユーザーがこれらのグループのメンバーである場合、GID 1、2、3、および4のグループが所有する検索を返す場合があります。 返される結果は、JSON形式のSearchLog構造のリストになります。

JSONの例：
```
[
        {
                "UID": 1,
                "GID": 2,
                "UserQuery": "grep stuff",
                "EffectiveQuery": "grep stuff | text",
                "Launched": "2015-12-30T23:30:23.298945825-07:00"
        },
        {
                "UID": 1,
                "GID": 2,
                "UserQuery": "grep stuff | grep things | grep that | grep this | sort by time",
                "EffectiveQuery": "grep stuff | grep things | grep that | grep this | sort by time | text",
                "Launched": "2015-12-30T23:31:08.237520376-07:00"
        }
]
```

### 管理者クエリ

管理者ユーザーとして`/api/searchhistory？admin = true`へのGETリクエストは、すべてのユーザーによるすべての検索を返します。

### 洗練された検索履歴

`/api/searchhistory`のGETメソッドの特殊なケースとして、「refine」値を設定することで単純な部分文字列検索用語を提供できます。 たとえば、`/api/searchhistory？refine = foo`へのGETリクエストは、検索のどこかに「foo」という用語を含む現在のユーザーのすべての検索を返します。
