# キットウェブAPI

このAPIでは、Gravwell Kitの作成、インストール、削除を行います。キットとは、ローカルシステムにインストールされる他のコンポーネントを含み、特定の問題に対する解決策をすぐに提供するものです。キットには以下のものが含まれます。

* Resources
* Scheduled searches
* Dashboards
* Auto-extractor definitions
* Templates
* Pivots
* User files
* Macros
* Search library entries
また、1つのキットは、ビルド時に指定された以下の属性を持ちます。

* ID: ID：このキットを一意に識別するための識別子。Androidの命名規則に従い、"com.example.my-kit "のようにすることをお勧めします。
* Name: 例："My Kit"。
* Description: キットの説明です。
* Version: キットの整数バージョンです。

## キットの構築

キットは、以下に定義するKitBuildRequest構造体を含むPOSTリクエストを`/api/kits/build`に送信することで構築されます。

```
type KitBuildRequest struct {
	ID                string
	Name              string
	Description       string
	Version           uint
	MinVersion        CanonicalVersion 
	MaxVersion        CanonicalVersion 
	Dashboards        []uint64 
	Templates         []uuid.UUID 
	Pivots            []uuid.UUID 
	Resources         []string 
	ScheduledSearches []int32 
	Macros            []uint64 
	Extractors        []uuid.UUID 
	Files             []uuid.UUID 
	SearchLibraries   []uuid.UUID 
	Playbooks         []uuid.UUID 
	EmbeddedItems     []KitEmbeddedItem 
	Icon              string 
	Banner            string 
	Cover             string 
	Dependencies      []KitDependency 
	ConfigMacros      []KitConfigMacro
	ScriptDeployRules map[int32]ScriptDeployConfig
}
```

ID、Name、Description、Version の各フィールドは必須ですが、Template/Pivot/Dashboards などの配列はオプションであることに注意してください。例えば、2つのDashboards、pivot、resource、およびscheduled searchを含むキットを構築するリクエストを以下に示します：

```
{
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG"
        }
    ],
    "Dashboards": [
        7,
        10
    ],
    "Description": "Test Gravwell kit",
    "ID": "io.gravwell.test",
    "Name": "test-gravwell",
    "Description":"testing\n\n## TESTING",
    "Pivots": [
        "ae9f2598-598f-4859-a3d4-832a512b6104"
    ],
    "Resources": [
        "84270dbd-1905-418e-b756-834c15661a54"
    ],
    "ScheduledSearches": [
        1439174790
    ],
    "EmbeddedItems":[
        {
           "Name":"TEST",
           "Type":"license",
           "Content":"VGVzdCBsaWNlbnNlIHRoYXQgYWxsb3dzIEdyYXZ3ZWxsIHRvIGdpdmUgeW91ciBmaXJzdCBib3JuIHNvbiBhIHN0ZXJuIHRhbGtpbmcgdG8h"
        }
    ],
    "Files":[
        "810a014d-1373-4d57-95b6-0638a7a01442",
        "09a26a2e-e449-4857-88d1-56cede1b8d95",
        "92bcfe5e-2c9a-4f39-9083-dd3f7a6f9738"
    ],
    "MinVersion":{"Major":4,"Minor":0,"Point":0},
    "MaxVersion":{"Major":4,"Minor":2,"Point":0},
    "Icon":"810a014d-1373-4d57-95b6-0638a7a01442",
    "Banner":"09a26a2e-e449-4857-88d1-56cede1b8d95",
    "Cover":"92bcfe5e-2c9a-4f39-9083-dd3f7a6f9738",
    "ScriptDeployRules": {
        "1439174790": {
            "Disabled": true,
            "RunImmediately": false
        }
    },
    "Version": 1
}
```

注意 テンプレート、ピボット、ユーザーファイルに指定されるUUIDは、それらの構造に関連付けられた*GUID*でなければならず、アイテムのリストでも報告される*ThingUUID*フィールドではありません。

注意してください。Banner、Cover、およびIconに指定されたUUIDは、ビルド要求のFilesリストに含まれていなければなりません。 ビルドリクエストに、メインのファイルリクエストに含まれていないファイルUUIDへの参照が含まれている場合、APIサーバーはリクエストを拒否します。

システムは、新しく構築されたキットを説明する構造体で応答します：

```
{
	"UUID": "2f5e485a-2739-475b-810d-de4f80ae5f52",
	"Size": 8268288,
	"UID": 1
}
```

このキットは、`/api/kits/build/<uuid>`をGETすることでダウンロードできます。上記のレスポンスがあれば、`/api/kits/build/2f5e485a-2739-475b-810d-de4f80ae5f52`からキットを取得することになります。

### 依存関係

キットは他のキットに依存している場合があります。これらの依存関係を以下のような構造でDependencies配列に列挙します:

```
{
	ID			string
	MinVersion	uint
}
```

IDフィールドは依存関係のIDを指定します。MinVersionフィールドには、インストールしなければならないそのキットの最小バージョンを指定します。

### コンフィグマクロ

キットでは、インストール時にGravwellが作成する特別なマクロである"コンフィグマクロ"を定義することができます。コンフィグマクロは次のようなものです:

```
{
	"MacroName": "KIT_WINDOWS_TAG",
	"Description": "Tag or tags containing Windows event entries",
	"DefaultValue": "windows",
	"Value": "",
	"Type": "TAG"
}
```

UI は、インストール時にマクロの希望値を尋ね、ユーザーの応答を KitConfig 構造体に含める必要があります。

コンフィグマクロの定義には、マクロが期待する値の種類についてのヒントとなるTypeフィールドを含めることができます。現在、以下のオプションが定義されています。

	* "TAG": 値は有効なタグでなければならない。このタグは、必ずしも現在のシステム上に存在する必要はありませんが、存在しないタグを入力した場合には、チェックしてユーザーに警告すると便利でしょう。
	* "STRING": 値は自由形式の文字列である必要があります。
   

Typeが指定されていない場合は、「STRING」（自由形式の入力）とします。

### スクリプトのデプロイ設定

デフォルトでは、キットに含まれるスクリプトはインストール時に有効に設定されます。この動作は script deploy config 構造体で制御できます。

```
{
	"Disabled": false,
	"RunImmediately": true,
}
```

この構造体には、「Disabled」と「RunImmediately」という2つのフィールドがあります。Disabledがtrueに設定されている場合、スクリプトは無効な状態でインストールされます。RunImmediatelyをtrueに設定すると、スクリプトが無効化されている場合でも、インストール後できるだけ早くスクリプトが実行されます。

スクリプトのデプロイオプションは、*キットのビルド時*に設定することも、*キットのデプロイ時*に設定して、キットの組み込みオプションを上書きすることもできます。

キットの構築時には、`ScriptDeployRules`フィールドに、スケジュールされたスクリプトのID番号（`ScheduledSearches`フィールドにリストされている）からスクリプトのデプロイ設定構造へのマッピングを含める必要があります。

キットをインストールする場合、`ScriptDeployRules` フィールドには、スケジュールされたスクリプトの *名前* から設定へのマッピングが含まれている必要があります。デプロイメントオプションは、デフォルトをオーバーライドしたい場合にのみ、インストール時に指定する必要があることに注意してください。

## キットのアップロード

キットをインストールするには、まずウェブサーバにアップロードする必要があります。キットは `/api/kits` へのPOSTリクエストによってアップロードされます。リクエストには、マルチパートのフォームを含める必要があります。ローカルシステムからファイルをアップロードするには、キットのファイルを含む `file` という名前のファイルフィールドをフォームに追加します。HTTPサーバなどのリモートシステムからファイルをアップロードするには，キットのURLを含む`remote`というフィールドを追加します。

また，`metadata`という名前のフィールドをリクエストに追加することもできます。このフィールドの内容はサーバでは解析されず，アップロードされたキットのMetadataフィールドに追加されます。これにより、キットの発信元のURLや、キットがアップロードされた日付などを記録することができます。

サーバーはアップロードされたキットの説明を応答します。

```
{
    "AdminRequired": false,
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG",
            "Type": "TAG",
            "Value": "winlog"
        }
    ],
    "ConflictingItems": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        }
    ],
    "Description": "Test Gravwell kit",
    "GID": 0,
    "ID": "io.gravwell.test",
    "Installed": false,
    "Items": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        },
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        },
        {
            "AdditionalInfo": {
                "DefaultDeploymentRules": {
                    "Disabled": false,
                    "RunImmediately": true
                },
                "Description": "A script",
                "Name": "myScript",
                "Schedule": "* * * * *",
                "Script": "println(\"hi\")"
            },
            "Name": "5aacd602-e6ed-11ea-94d9-c771bfc07a39",
            "Type": "scheduled search"
        }
    ],
    "ModifiedItems": [
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        }
    ],
    "Name": "test-gravwell",
    "RequiredDependencies": [
        {
            "AdminRequired": false,
            "Assets": [
                {
                    "Featured": true,
                    "Legend": "Littering AAAAAAND",
                    "Source": "cover.jpg",
                    "Type": "image"
                },
                {
                    "Featured": false,
                    "Legend": "",
                    "Source": "readme.md",
                    "Type": "readme"
                }
            ],
            "Created": "2020-03-23T15:36:00.294625802-06:00",
            "Dependencies": null,
            "Description": "A simple test kit that just provides a resource",
            "ID": "io.gravwell.testresource",
            "Ingesters": [
                "simplerelay"
            ],
            "Items": [
                {
                    "AdditionalInfo": {
                        "Description": "hosts",
                        "ResourceName": "devlookup",
                        "Size": 610,
                        "VersionNumber": 1
                    },
                    "Name": "devlookup",
                    "Type": "resource"
                },
                {
                    "AdditionalInfo": "Testkit resource\n\nThis really has no restrictions, go nuts!\n",
                    "Name": "LICENSE",
                    "Type": "license"
                }
            ],
            "MaxVersion": {
                "Major": 0,
                "Minor": 0,
                "Point": 0
            },
            "MinVersion": {
                "Major": 0,
                "Minor": 0,
                "Point": 0
            },
            "Name": "Testing resource kit",
            "Signed": true,
            "Size": 10240,
            "Tags": [
                "syslog"
            ],
            "UUID": "d2a0cb10-ff25-4426-8b87-0dd0409cae48",
            "Version": 1
        }
    ],
    "Signed": false,
    "UID": 7,
    "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
    "Version": 2
}
```

ModifiedItems "フィールドに注目してください。このキットの旧バージョンがすでにインストールされている場合、このフィールドには、ユーザーが変更した*アイテムのリストが含まれています。ステージングされたキットをインストールすると、これらの項目が上書きされるため、ユーザーに通知し、変更を保存する機会を与える必要があります。

"ConflictingItems "には、ユーザーが作成したオブジェクトと衝突すると思われるアイテムがリストアップされます。この例では、ユーザが以前に "maxmind_asn "という名前の独自のリソースを作成しているようです。OverwriteExisting` を true に設定してインストールリクエストを送信すると、そのリソースはキット内のバージョンで上書きされます。false に設定すると、インストールプロセスはエラーを返します。

"RequiredDependencies"フィールドには、このキットの現在インストールされていない依存関係のメタデータ構造のリストが含まれており、これには表示すべきライセンスを含むアイテムセットも含まれています。

"ConfigMacros"フィールドには、このキットによってインストールされる構成マクロ（前のセクションを参照）のリストが含まれます。このキットの前のバージョン（または別のキット）が同じ名前のマクロをすでにインストールしていた場合、ウェブサーバは「Value」フィールドにマクロの現在の値をあらかじめ入力します。ユーザー*が同じ名前のマクロをインストールしていた場合、ウェブサーバはエラーを返します。

"myScript"という名前のスケジュール検索、特に`DefaultDeploymentRules`フィールドに注目してください。これは、スクリプトがどのようにインストールされるかを説明しています。有効とマークされ、できるだけ早く実行されます。

## キットの一覧表示

`api/kits`に GET リクエストをすると、すべての既知のキットのリストが返されます。以下は、システムにアップロードされたものの、まだインストールされていないキットがある場合の結果を示す例です。

```
[
    {
        "AdminRequired": false,
        "Description": "Test Gravwell kit",
        "GID": 0,
        "ID": "io.gravwell.test",
        "Installed": false,
        "Items": [
            {
                "AdditionalInfo": {
                    "Description": "ASN database",
                    "ResourceName": "maxmind_asn",
                    "Size": 6196221,
                    "VersionNumber": 1
                },
                "Name": "84270dbd-1905-418e-b756-834c15661a54",
                "Type": "resource"
            },
            {
                "AdditionalInfo": {
                    "DefaultDeploymentRules": {
                        "Disabled": false,
                        "RunImmediately": true
                    },
                    "Description": "count all entries",
                    "Duration": -3600,
                    "Name": "count",
                    "Schedule": "* * * * *",
                    "Script": "var time = import(\"time\")\n\naddSelfTargetedNotification(7, \"hello\", \"/#/search/486574780\", time.Now().Add(30 * time.Second))"
                },
                "Name": "55c81086",
                "Type": "scheduled search"
            },
            {
                "AdditionalInfo": {
                    "Description": "My dashboard",
                    "Name": "Foo",
                    "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
                },
                "Name": "a",
                "Type": "dashboard"
            },
            {
                "AdditionalInfo": {
                    "Description": "foobar",
                    "Name": "foo",
                    "UUID": "ae9f2598-598f-4859-a3d4-832a512b6104"
                },
                "Name": "ae9f2598-598f-4859-a3d4-832a512b6104",
                "Type": "pivot"
            }
        ],
        "Name": "test-gravwell",
        "Signed": false,
        "UID": 7,
        "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
        "Version": 1
    }
]
```

キットアイテムの種類ごとにどのような"AdditionalInfo"フィールドが利用できるかについては、このページの最後にあるリストを参照してください。

## キット情報

`/api/kits/<GUID>` でのGETリクエスト(`<GUID>`は特別にインストールまたはステージングされたキットのGUID)は、その特定のキットに関する情報を提供します。

たとえば、`/api/kits/549c0805-a693-40bd-abb5-bfb29fc98ef1`に対するGETリクエストは、次のようになります：

```
{
    "AdminRequired": false,
    "Description": "Test Gravwell kit",
    "GID": 0,
    "ID": "io.gravwell.test",
    "Installed": false,
    "Items": [
        {
            "AdditionalInfo": {
                "Description": "ASN database",
                "ResourceName": "maxmind_asn",
                "Size": 6196221,
                "VersionNumber": 1
            },
            "Name": "84270dbd-1905-418e-b756-834c15661a54",
            "Type": "resource"
        },
        {
            "AdditionalInfo": {
                "DefaultDeploymentRules": {
                    "Disabled": false,
                    "RunImmediately": true
                },
                "Description": "count all entries",
                "Duration": -3600,
                "Name": "count",
                "Schedule": "* * * * *",
                "Script": "var time = import(\"time\")\n\naddSelfTargetedNotification(7, \"hello\", \"/#/search/486574780\", time.Now().Add(30 * time.Second))"
            },
            "Name": "55c81086",
            "Type": "scheduled search"
        },
        {
            "AdditionalInfo": {
                "Description": "My dashboard",
                "Name": "Foo",
                "UUID": "5567707c-8508-4250-9121-0d1a9d5ebe32"
            },
            "Name": "a",
            "Type": "dashboard"
        },
        {
            "AdditionalInfo": {
                "Description": "foobar",
                "Name": "foo",
                "UUID": "ae9f2598-598f-4859-a3d4-832a512b6104"
            },
            "Name": "ae9f2598-598f-4859-a3d4-832a512b6104",
            "Type": "pivot"
        }
    ],
    "Name": "test-gravwell",
    "Signed": false,
    "UID": 7,
    "UUID": "549c0805-a693-40bd-abb5-bfb29fc98ef1",
    "Version": 1
}

```

キットが存在しない場合は404が返され、要求された特定のキットにユーザーがアクセスできない場合は400が返されます。

## キットのインストール

アップロードされたキットをインストールするには、`/api/kits/<uuid>`にPUTリクエストを送ります。UUIDはキットのリストにあるUUIDフィールドです。サーバーはいくつかの予備チェックを行い、整数値を返します。この整数値は、installation status API (下記参照) を使ってインストールの進捗状況を問い合わせるのに使用できます。

インストール中、すべての必要な依存関係（ステージング応答のRequiredDepdenciesフィールドにリストされているもの）はステージングされ、キット自体の最終インストールの前に自動的にインストールされます。

追加のキットインストールオプションは、リクエストのボディに設定構造を渡すことで指定できます。

```
{
    "AllowUnsigned": false,
    "ConfigMacros": [
        {
            "DefaultValue": "windows",
            "Description": "Tag or tags containing Windows event entries",
            "MacroName": "KIT_WINDOWS_TAG",
            "Value": "winlog"
        }
    ],
    "Global": true,
    "InstallationGroup": 3,
    "Labels": [
        "foo",
        "bar"
    ],
    "OverwriteExisting": true,
    "ScriptDeployRules": {
        "myScript": {
            "Disabled": true,
            "RunImmediately": false
        }
    }
}
```

注：以下の項目はすべてオプションです。デフォルトのオプションを使用するには、単にリクエストからボディを省略してください。

`OverwriteExisting` が設定されている場合、インストーラはキットのバージョンと同じ名前のユニークな識別子を持つ既存のアイテムを単純に置き換えることができます。

`Global` フラグは管理者のみが設定できます。設定された場合、すべてのアイテムは Global としてマークされ、すべてのユーザーがアクセスできることになります。

一般ユーザはGravwellからの適切に署名されたキットのみをインストールできます。AllowUnsigned`が設定されている場合、*管理者*は署名のないキットをインストールすることができます。

`InstallationGroup` を設定すると、インストールするユーザが所属するグループの一つとキットの内容を共有することができます。

`Labels`は、インストール時にキット内のラベル付け可能なアイテムに適用する追加ラベルのリストです。Gravwellはキットにインストールされたアイテムに "kit "とキットのID(例："io.gravwell.coredns")を自動的にラベル付けするので注意が必要です。

`ConfigMacros` はキット情報構造の中にあるConfigMacrosのリストで、"Value "フィールドにはユーザーが任意に設定することができます。Value" フィールドが空白の場合、ウェブサーバは "DefaultValue" を使用します。

`ScriptDeployRules` には、キット内のスケジュールされたスクリプトのうち、オーバーライドしたいデプロイメントルールが含まれている必要があります。この例では、"myScript "という名前のスクリプトが無効な状態でインストールされます。デフォルトのデプロイメントオプションが許容できる場合は、このフィールドは空にしておくことができます。

### インストールステータス API

依存関係の多い大型パッケージのインストールには時間がかかることがあるため、インストールリクエストが送信されると、サーバーはリクエストをキューに入れて処理します。サーバーはインストール要求に対して、`2019727887`のような整数値で応答します。これをインストールステータスAPIで使用して、`/api/kits/status/<id>`にGETリクエストを送信することで、インストールの進捗状況を照会することができます。

```
{
    "CurrentStep": "Done",
    "Done": true,
    "Error": "",
    "InstallID": 2019727887,
    "Log": "\nQueued installation of kit io.gravwell.testresource, with 0 dependencies also to be installed\nBeginning installation of io.gravwell.testresource (9b701e75-76ee-40fc-b9b5-4c7e1706339d) for user Admin John (1)\nInstalling requested kit io.gravwell.testresource\nDone",
    "Owner": 1,
    "Percentage": 1,
    "Updated": "2020-03-25T15:39:37.184221203-06:00"
}
```

"Owner"は、インストールリクエストを送信したユーザーのUIDです。"Done"は、キットが完全にインストールされたときにtrueに設定されます。"Percentage"は0から1の間の値で、インストールがどの程度完了したかを示します。"CurrentStep"はインストールの現在の状態を表し、"Log"はインストール全体の状態を完全に記録しています。"Error"は、インストールプロセスで何か問題が発生していない限り、空になります。"Updated"は、ステータスが最後に変更された時間です。

また、`/api/kits/status` を GET して、*all* kit のインストールステータスのリストを要求することもできます。これは、上記のようなオブジェクトの配列を返します。デフォルトでは、これは現在のユーザのステータスのみを返すことに注意してください。管理者は、システム上の *all* ステータスを取得するために、URL に `?admin=true` を追加することができます。

## キットのアンインストール

キットを削除するには、`/api/kits/<uuid>`に対してDELETEリクエストを発行します。キット内のアイテムがインストール後にユーザーによって変更されていた場合、レスポンスは400のステータスコードを持ち、何が変更されたかを詳細に説明する構造を含みます。

```
{
    "Error": "Kit items have been modified since installation, set ?force=true to override",
    "ModifiedItems": [
        {
            "AdditionalInfo": {
                "Description": "Network services (protocol + port) database",
                "ResourceName": "network_services",
                "Size": 531213,
                "VersionNumber": 1
            },
            "ID": "2e4c8f31-92a4-48b5-a040-d2c895caf0b2",
            "KitID": "io.gravwell.networkenrichment",
            "KitName": "Network enrichment",
            "KitVersion": 1,
            "Name": "network_services",
            "Type": "resource"
        }
    ]
}
```

強制的にキットを取り外すには、リクエストに `?force=true` パラメータを追加してください。

## リモートキットサーバへの問い合わせ

Gravwell Kit Serverからリモートキットのリストを取得するには、`/api/kits/remote/list`にGETを発行します。 これは、利用可能なすべてのキットの最新バージョンを表す、JSONエンコードされたキットメタデータ構造のリストを返します。 APIパスの `/api/kits/remote/list/all` は、すべてのバージョンのすべてのキットを提供します。

メタデータ構造は以下の通りです。

```
type KitMetadata struct {
	ID            string
	Name          string
	GUID          string
	Version       uint
	Description   string
	Signed        bool
	AdminRequired bool
	MinVersion    CanonicalVersion
	MaxVersion    CanonicalVersion
	Size          int64
	Created       time.Time
	Ingesters     []string //ingesters associated with the kit
	Tags          []string //tags associated with the kit
	Assets        []KitMetadataAsset
}

type KitMetadataAsset struct {
	Type     string
	Source   string //URL
	Legend   string //some description about the asset
	Featured bool
}

type CanonicalVersion struct {
	Major uint32
	Minor uint32
	Point uint32
}
```

ここではその一例をご紹介します:

```
WEB GET http://172.19.0.2:80/api/kits/remote/list:
[
	{
		"ID": "io.gravwell.test",
		"Name": "testkit",
		"GUID": "c2870b48-ff31-4550-bd58-7b2c1c10eeb3",
		"Version": 1,
		"Description": "Testing a kit with a license in it",
		"Signed": true,
		"AdminRequired": false,
		"MinVersion": {
			"Major": 0,
			"Minor": 0,
			"Point": 0
		},
		"MaxVersion": {
			"Major": 0,
			"Minor": 0,
			"Point": 0
		},
		"Size": 0,
		"Created": "2020-02-10T16:31:23.03192303Z",
		"Ingesters": [
			"SimpleRelay",
			"FileFollower"
		],
		"Tags": [
			"syslog",
			"auth"
		],
		"Assets": [
			{
				"Type": "image",
				"Source": "cover.jpg",
				"Legend": "TEAM RAMROD!",
				"Featured": true
			},
			{
				"Type": "readme",
				"Source": "readme.md",
				"Legend": "",
				"Featured": false
			},
			{
				"Type": "image",
				"Source": "testkit.jpg",
				"Legend": "",
				"Featured": false
			}
		]
	}
]
```

## 単一のキット情報を引き出す

リモートキットAPIは、`/api/kits/remote/<guid>`に対して、`GET`を発行することで、特定のキットに関する情報を引き出すこともサポートしています。これは、単一の`KitMetadata`構造体を返します。

例えば、`/api/kits/remote/c2870b48-ff31-4550-bd58-7b2c1c10eeb3`に対して、`GET`を発行すると、ウェブサーバは次のように返します。

```
{
	"ID": "io.gravwell.test",
	"Name": "testkit",
	"GUID": "c2870b48-ff31-4550-bd58-7b2c1c10eeb3",
	"Version": 1,
	"Description": "Testing a kit with a license in it",
	"Signed": true,
	"AdminRequired": false,
	"MinVersion": {
		"Major": 0,
		"Minor": 0,
		"Point": 0
	},
	"MaxVersion": {
		"Major": 0,
		"Minor": 0,
		"Point": 0
	},
	"Size": 0,
	"Created": "2020-02-10T16:31:23.03192303Z",
	"Ingesters": [
		"SimpleRelay",
		"FileFollower"
	],
	"Tags": [
		"syslog",
		"auth"
	],
	"Assets": [
		{
			"Type": "image",
			"Source": "cover.jpg",
			"Legend": "TEAM RAMROD!",
			"Featured": true
		},
		{
			"Type": "readme",
			"Source": "readme.md",
			"Legend": "",
			"Featured": false
		},
		{
			"Type": "image",
			"Source": "testkit.jpg",
			"Legend": "",
			"Featured": false
		}
	]
}
```

### リモートのキットサーバーからキットのアセットを引き出します

キットには、実際にキットをダウンロード/インストールする前に、画像やマークダウン、ライセンス、キットの目的を探るための追加ファイルなどを表示するためのアセットも含まれています。 これらのアセットは、`api/kits/remote/<guid>/<asset>`に対してGETリクエストを実行することで、リモートシステムから取得することができます。 例えば、guidが `c2870b48-ff31-4550-bd58-7b2c1c10eeb3` のキットの Type "image" and Legend "TEAM RAMROD!" のアセットを取り出したい場合、`/api/kits/remote/c2870b48-ff31-4550-bd58-7b2c1c10eeb3/cover.jpg` に GET を発行します。


## キットアイテムの"追加情報"フィールド

キットをリストアップするとき（`/api/kits`でGET）、各キットはAddditionalInfoフィールドを含むアイテムのリストを含みます。これらのフィールドは、キット内のアイテムに関する詳細な情報を提供します。内容はアイテムの種類によって異なり、以下のように列挙されます。

```
Resources:
		VersionNumber int
		ResourceName  string
		Description   string
		Size          uint64

Scheduled Search:
		Name                    string
		Description             string
		Schedule                string
		SearchString            string 
		Duration                int64  
		Script                  string 
		DefaultDeploymentRules  ScriptDeployConfig

Dashboard:
		UUID        string
		Name        string
		Description string

Extractor:
		Name   string 
		Desc   string 
		Module string 
		Tag    string 

Template:
		UUID        string
		Name        string
		Description string

Pivot:
		UUID        string
		Name        string
		Description string

File:
		UUID        string
		Name        string
		Description string
		Size        int64
		ContentType string

Macro:
		Name      string
		Expansion string

Search Library:
		Name        string
		Description string
		Query       string

Playbook:
		UUID        string
		Name        string
		Description string

License:
		(contents of license file itself)
```

## キットビルドリクエスト履歴

成功したキットのビルドリクエストはウェブサーバに保存されます。GETリクエストを `/api/kits/build/history` に送ることで、現在のユーザーのビルドリクエストのリストを取得することができます。レスポンスは、ビルドリクエストの配列になります。

```
[{"ID":"io.gravwell.test","Name":"test","Description":"","Version":1,"MinVersion":{"Major":0,"Minor":0,"Point":0},"MaxVersion":{"Major":0,"Minor":0,"Point":0},"Macros":[4,41],"ConfigMacros":null}]
```

注：このストアはUID + キットIDをキーにしています。"io.gravwell.test "という名前のキットを再度作成すると、ストア内のバージョンは上書きされます。

例えば`/api/kits/build/history/io.gravwell.test`のように、`/api/kits/build/history/io.gravwell.test`にDELETEリクエストを送ることで、特定のアイテムを削除することができます。
