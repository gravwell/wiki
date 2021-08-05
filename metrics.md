# Gravwellのメトリクスとクラッシュレポート

Gravwellのユーザーは、自分のネットワークで何が起こっているかを気にかけています。それは、人々が最初に私たちをチェックアウトする大きな理由です。Gravwellに組み込まれている自動クラッシュレポートとメトリックスシステムについて、すべてのユーザーが理解し、快適に使用できるようにしたいと考えています。このドキュメントでは、Gravwellのサーバーに送信される内容の完全な例とともに、両方のシステムについて説明します。

## クラッシュレポート

Gravwellのコンポーネントがクラッシュすると、自動クラッシュレポートがGravwellに送信されます。これは対象コンポーネントのコンソール出力で構成されており、通常はライセンスに関する簡単な情報（誰のシステムがクラッシュしたかを判断するため）とスタックトレースが含まれている。**Gravwellのすべてのコンポーネント（ウェブサーバ、インデクサ、インジェスタ、サーチエージェント）は、クラッシュレポートを送信するように設定されています。
 
注：クラッシュレポートは常にTLSで検証されたHTTPSでupdate.gravwell.ioに送信されます。リモートの証明書を完全に検証できない場合は、レポートは*出ません*。

以下は、Gravwell社員のテストシステムからのクラッシュレポートの例です。

```
Component:      webserver
Version:        3.3.5
Host:           X.X.X.X
Domain: c-X-X-X-X.hsd1.nm.comcast.net
Full Log Location:      /opt/gravwellCustomers/uploads/crash/webserver_3.3.5_X.X.X.X_2020-01-31T14:39:42


Log Snippet:
Version         3.3.10
API Version     0.1
Build Date      2020-Apr-30
Build ID        745dc6ca
Cmdline         /opt/gravwell/bin/gravwell_webserver -stderr gravwell_webserver.service
Executing user  gravwell
Parent PID      1
Parent cmdline  /sbin/init
Parent user     root
Total memory    4147781632
Memory used     5.781707651865122%
Process heap allocation 2005776
Process sys reserved    72112017
CPU count       4
Host hash       4224be94ae35247ed32013d9021f64bc40986c9fbbafac97787ab58b400f1666
Virtualization role     guest
Virtualization system   kvm
max_map_count   65530
RLIMIT_AS (address space)       18446744073709551615 / 18446744073709551615
RLIMIT_DATA (data seg)  18446744073709551615 / 18446744073709551615
RLIMIT_FSIZE (file size)        18446744073709551615 / 18446744073709551615
RLIMIT_NOFILE (file count)      1048576 / 1048576
RLIMIT_STACK (stack size)       8388608 / 18446744073709551615
SKU             2UX
Customer NUM    00000000
Customer GUID   ffffffff-ffff-ffff-ffff-ffffffffffff

panic: send on closed channel

goroutine 90 [running]:
gravwell/pkg/search.(*SearchManager).clientSystemStatsRoutine(0xc01414edc0,
0xc00017fd20, 0x16, 0xc000c300c0, 0xc000c1dc80)
        gravwell@/pkg/search/manager_handlers.go:284 +0x106
created by gravwell/pkg/search.(*SearchManager).GetSystemStats.func1
        gravwell@/pkg/search/manager_handlers.go:301 +0x65
```

メッセージは、クラッシュした特定のコンポーネント（この場合はウェブサーバ）から始まります。そして、Gravwellのバージョン、クラッシュしたシステムのIPとホスト名、Gravwellスタッフがクラッシュログの完全なコピーを見つけられるパスが記載されています（バックトレースが特に長い場合は、最初の100行ほどだけがメールで送られてきます）。

メッセージの残りの部分は、クラッシュしたプログラムのコンソール出力です。クラッシュレポーターは、`/dev/shm`にあるコンポーネントの出力ファイルから直接これを取り出します。例えば`/dev/shm/gravwell_webserver`を見れば、あなたのシステムが何を送信しているかがわかります。また、`/opt/gravwell/log/crash` で過去のクラッシュレポートを見ることができますが、クラッシュレポーターを *disable* した場合、クラッシュログはそのディレクトリに保存されなくなることに注意してください。

最初の数行（Version, API Version, Build Date, Build ID）は、Gravwellのどのバージョンが実行されていたかを正確に判断するのに役立ちます。「Cmdline"、"Executing user"、"Parent PID"、"Parent cmdline"、"Parent user "は、Gravwellのプロセスがどのように実行されているかを把握し、潜在的な問題を特定するのに役立ちます。この例では、親プロセスのPIDが1で、名前が "manager "であることから、GravwellはDockerコンテナで実行されていると推測できます。この例では、親プロセスのPIDが1で、名前が "manager "であることから、GravwellはDockerコンテナで実行されていると推測できます。

また、システム上のメモリ量と設定されているrlimitsの情報も含まれています。これは、特定のクラスのクラッシュを追跡するのに役立つからです。例えば、512MBのRAMを持つシステムでメモリ不足のエラーが発生しても、特に驚くことではありません。Host hash」フィールドは、プロセスを実行しているホストのユニークな識別子ですが、ハッシュなので「これは他のクラッシュレポートと同じマシンだ」と言うためにしか使えず、他の情報は含まれていないことに注意してください。

SKU」、「Customer NUM」、「Customer GUID」フィールドは、使用中のライセンスを示します。SKUは、ユーザーのライセンスで許可されている機能を表します。このケースでは、Gravwellの従業員は無制限（"UX"）ライセンスを使用しています。顧客番号と顧客GUIDフィールドは、顧客データベースを参照し、誰が問題を抱えているかを確認するためのものです。

これらの情報の下には、Gravwellのプロセスからのバックトレースがあります。このケースでは、アルファビルドのバグにより、GUIのハードウェア統計ページで使用するためにウェブサーバがインデクサのCPU/メモリ情報をチェックするために使用するルーチンがクラッシュしたことがわかります。このスタックトレースにはユーザーデータは含まれておらず、ソースコードの行番号のみが含まれていることを明確にしておきます。

### クラッシュレポートの無効化

何らかの理由でクラッシュレポートを送信したくないと思った場合、レポートシステムを無効にする方法が複数あります。

* スタンドアロンのシェルインストーラを使用している場合は、インストール時に `--no-crash-report` フラグを付けて無効にすることができます。
* DebianのリポジトリからGravwellをインストールした場合は、`systemctl disable gravwell_crash_report`で無効にすることができます。
* GravwellのDockerイメージを使用している場合は、Dockerコマンドに`-e DISABLE_ERROR_HANDLING=true`を渡すことで、クラッシュレポーターを無効にすることができます。

しかし、クラッシュレポートを有効にしたままにしておいていただけるとありがたいです。クラッシュレポートのおかげで、ユーザーが気づかないような問題を発見し、修正することができます。これは我々のソフトウェアを改善するための最高のフィードバックメカニズムの一つです。

過去のクラッシュレポートの削除をご希望の場合は、support@gravwell.io までご連絡ください。

## メトリクスレポート

Gravwellのウェブサーバコンポーネント（ウェブサーバのみ）は、一般的な使用状況を示すHTTPS POSTリクエストをGravwellコーポレートサーバに送信することがあります。この情報は、どの機能が最も利用されているか、どの機能がもっと利用されているかを把握するのに役立ちます。GravwellがどのくらいのRAMを消費しているのか、ガベージコレクションを最適化する必要があるのか、あるいはデフォルトの設定をより保守的にする必要があるのか、といった統計情報を生成することができます。また、有料のライセンスが不適切に導入されていないかどうかも確認できます。

これらのメトリクスを収集する際の最も重要な目的は、お客様のデータの匿名性を守ることです。これらのメトリクスレポートには、Gravwellに保存されているデータの実際の内容が含まれることはなく、実際の検索クエリやシステム上のタグのリストが送信されることもありません。

注）メトリクスレポートは、常にTLSで検証されたHTTPSでupdate.gravwell.ioに送信されます。リモート証明書を完全に検証できない場合、レポートは*出ません*。

メトリクスレポートが送信されると、サーバーはGravwellの最新バージョンで応答します。これにより、新しいバージョンが利用可能になったときに、Gravwell UIに通知を表示することができます（これらの通知は、gravwell.confの`Disable-Update-Notification`パラメータで無効にすることができます）。

以下は、Gravwell社員のホームシステムから送られてきた例です:

```
{
    "ApiVer": {
        "Major": 0,
        "Minor": 1
    },
    "AutomatedSearchCount": 1,
    "BuildVer": {
        "BuildDate": "2020-04-02T00:00:00Z",
        "BuildID": "e755ee13",
        "GUIBuildID": "87e5e523",
        "Major": 3,
        "Minor": 3,
        "Point": 8
    },
    "CustomerNumber": 000000000,
    "CustomerUUID": "ffffffff-ffff-ffff-ffff-ffffffffffff",
    "DashboardCount": 5,
    "DashboardLoadCount": 13,
    "DistributedFrontends": false,
    "ForeignDashboardLoadCount": 0,
    "ForeignSearchLoadCount": 0,
    "Groups": 2,
    "IndexerCount": 4,
    "IndexerNodeInfo": [
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 47899944,
            "ProcessSysReserved": 282423040,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 66157568,
            "ProcessSysReserved": 282554112,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 58577296,
            "ProcessSysReserved": 351827712,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 58304584,
            "ProcessSysReserved": 282226432,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        }
    ],
    "IndexerStats": [
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 658757162,
                    "Entries": 2770447
                },
                {
                    "Cold": false,
                    "Data": 12681258,
                    "Entries": 9882
                },
                {
                    "Cold": false,
                    "Data": 325462303,
                    "Entries": 1344586
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 45312907669,
                    "Entries": 119150365
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 50161444,
                    "Entries": 297743
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 669469662,
                    "Entries": 2931573
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 325986097,
                    "Entries": 1348645
                },
                {
                    "Cold": false,
                    "Data": 50301788,
                    "Entries": 298556
                },
                {
                    "Cold": false,
                    "Data": 45316008062,
                    "Entries": 119174395
                },
                {
                    "Cold": false,
                    "Data": 12341038,
                    "Entries": 9559
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 663669955,
                    "Entries": 2782081
                },
                {
                    "Cold": false,
                    "Data": 326449600,
                    "Entries": 1350525
                },
                {
                    "Cold": false,
                    "Data": 50427080,
                    "Entries": 299538
                },
                {
                    "Cold": false,
                    "Data": 12552734,
                    "Entries": 9759
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 45445347364,
                    "Entries": 119473828
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 660249138,
                    "Entries": 2794164
                },
                {
                    "Cold": false,
                    "Data": 45332590720,
                    "Entries": 119204014
                },
                {
                    "Cold": false,
                    "Data": 50572152,
                    "Entries": 300251
                },
                {
                    "Cold": false,
                    "Data": 12608944,
                    "Entries": 9730
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 325899751,
                    "Entries": 1347670
                }
            ]
        }
    ],
    "IngesterCount": 8,
    "LicenseHash": "kH3R+R4AdTCnXFYDi3L4nZ==",
    "LicenseTimeLeft": 23550517079782204,
    "ManualSearchCount": 330,
    "ResourceUpdates": 13356,
    "ResourcesCount": 4,
    "SKU": "2UX",
    "ScheduledSearchCount": 4,
    "SearchCount": 331,
    "Source": "X.X.X.X",
    "SystemMemory": 67479150592,
    "SystemProcs": 3,
    "SystemUptime": 1920449,
    "TimeStamp": "2020-04-02T22:11:23Z",
    "TotalData": 185614443921,
    "TotalEntries": 494907311,
    "Uptime": 300,
    "UserLoginCount": 27,
    "Users": 2,
    "WebserverNodeInfo": {
        "CPUCount": 12,
        "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
        "ProcessHeapAllocation": 311618224,
        "ProcessSysReserved": 420052881,
        "TotalMemory": 67479150592,
        "VirtRole": "guest",
        "VirtSystem": "docker"
    },
    "WebserverUUID": "17405830-3ac4-4b75-a639-6a265e6718a4",
    "WellCount": 28
}
```

このウェブサーバが4つのインデクサに接続されており、それぞれが独自の情報を取得していることもあって、構造が大きくなっています。以下にフィールドの詳細を説明します。

* `ApiVer`: Gravwell APIの内部バージョニング番号。
* `AutomatedSearchCount`: 検索エージェントやダッシュボードの読み込みによって）"自動的に "実行された検索の数です。
* `BuildVer`: * `BuildVer`: このシステム上のGravwellの特定のビルドを記述する構造体です。
* `CustomerNumber`: * `CustomerNumber`: このシステム上のライセンスに関連する顧客番号。
* `CustomerUUID`: このシステム上のライセンスに関連する顧客番号です。このシステム上のライセンスのUUIDです。
* `DashboardCount`: ダッシュボードの数です。存在するダッシュボードの数です。
* `DashboardLoadCount`: ダッシュボードの数です。* `DashboardLoadCount`: ユーザーによって開かれたダッシュボードの種類の数。
* `DistributedFrontends`: 分散型フロントエンドです。Distributed Webservers](#!distributed/frontend.md)が有効な場合、trueに設定されます。
* `ForeignDashboardLoadCount`: ユーザーがダッシュボードを閲覧した回数です。ForeignDashboardLoadCount`: ユーザーが他のユーザーが所有するダッシュボードを閲覧した回数です (ダッシュボードの共有オプションが十分な柔軟性を持っているかどうかを判断するのに役立ちます)。
* `ForeignSearchLoadCount`: 他のユーザーが所有する検索をユーザーが閲覧した回数（検索の共有オプションが十分な柔軟性を持っているかどうかの判断に役立ちます。
* `Groups`: システム上のユーザーグループの数です。
* `IndexerCount`: * `IndexerCount`: このウェブサーバが接続されているインデクサの数です。
* `IndexerNodeInfo`: 各インデクサの統計情報を簡潔に記述した、インデクサごとに1つの構造体の配列。
	- `CPUCount`: CPUCount`: インデクサに搭載されているCPUコアの数。
	- `HostHash`: 非可逆的なハッシュ（[github.com/denisbrodbeck/machineid](https://github.com/denisbrodbeck/machineid)を参照）で、インデクサーを実行しているホストマシンを一意に識別します。なお、この例では、インデクサはすべて1つのDockerホスト上で動作しているため、すべて同じHostHashを持っています。
	- `ProcessHeapAllocation`: インデクサプロセスによって割り当てられるヒープメモリの量です。
	- `ProcessSysReserved`: IndexerプロセスがOSから予約したメモリの合計量です。
	- `TotalMemory`: システムのメインメモリのサイズです。システムのメインメモリのサイズです。
	- `VirtRole`: VirtRole`: "host "または "guest "で、インデクサーが仮想マシンで動作しているかどうかによります。
	- `VirtSystem`: 仮想化システム。xen」、「kvm」、「vbox」、「vmware」、「docker」、「openvz」、「lxc」などがあります。
* `IndexerStats`: 各インデクサの統計構造体の配列です。
	* `WellStats`: インデクサ上の各ウェルに関する匿名化された情報の配列です。
		* `Cold`: Cold`: "コールド "ウェルであるかどうか。
		* `Data`: このウェルに含まれるデータのバイト数です。
		* `Entries`: エントリ数。このウェルに含まれるエントリーの数です。
* `IngesterCount`: インジェスター数です。システムに接続されているユニークなインゲスターの数です。
* `LicenseHash`: ライセンスの MD5 合計です。使用されているライセンスの MD5 合計です。
* `LicenseTimeLeft`: ライセンスの残り時間です。ライセンスの残りの秒数です。
* `ManualSearchCount`: 手動で実行された検索回数です。手動で実行された検索の数です。
* `ResourceUpdates`: リソースの更新回数です。リソースが変更された回数です。
* `ResourcesCount`: リソースの数。システム上のリソースの数です。
* `SKU`: 使用中のライセンスのSKUです。使用されているライセンスのSKUです。
* `ScheduledSearchCount`: システムにインストールされているスケジュール検索の数です。
* `SearchCount`: 非推奨のフィールドで、`ManualSearchCount` + `AutomatedSearchCount` の合計値です。
* `Source`: このレポートが発信されたIPです。このレポートが発信されたIPです。
* `SystemMemory`: * `SystemMemory`: ウェブサーバのホストシステムにインストールされているメモリのバイト数。
* `SystemProcs`: * `SystemProcs`: ホストシステム上で動作しているプロセスの数。
* `SystemUptime`: * `SystemUptime`: ホストシステムが稼働している秒数。
* `TimeStamp`: このレポートが生成された時間。
* `TotalData`: データの総量です。* `TotalData`: すべてのインデクサー上のすべてのウェルのバイト数です。
* `TotalEntries`: 全てのインデクサー上の全てのウェルにおけるエントリー数です。
* `Uptime`: `Uptime`: ウェブサーバプロセスが起動してからの秒数。
* `UserLoginCount`: ユーザーがログインした回数です。ユーザーがログインした回数です。
* `Users`: ユーザーがログインした回数。システムに登録されているユーザー数。
* `WebserverNodeInfo`: Webサーバプロセスを実行しているシステムの簡単な説明です。
	- `CPUCount`: `CPUCount`: ウェブサーバに搭載されているCPUコアの数。
	- `HostHash`: 非可逆的なハッシュ (github.com/denisbrodbeck/machineid](https://github.com/denisbrodbeck/machineid) を参照) で、ウェブサーバを実行しているホストマシンを一意に識別します。
	- `ProcessHeapAllocation` です。ウェブサーバプロセスによって割り当てられたヒープメモリの量です。
	- `ProcessSysReserved`: Web サーバプロセスが OS から予約したメモリの総量。
	- `TotalMemory`: システムのメインメモリのサイズ。システムのメインメモリのサイズです。
	- `VirtRole`: `VirtRole`: "host "または "guest "で、ウェブサーバが仮想マシンの中で動作しているかどうかによる。
	- `VirtSystem`: 仮想化システム。xen", "kvm", "vbox", "vmware", "docker", "openvz", "lxc" などの仮想化システムです。
* `WebserverUUID`: Gravwell の Web サーバはインストール時に UUID を生成するが、このフィールドにはその UUID が格納される。
* `WellCount`: * `WellCount`: 全てのインデクサーにおけるウェルの合計数です。

私たちは報告する情報を慎重に検討し、あなたがGravwellに持っているデータの種類や内容に関する情報を得ることができないようにするために努力しました。ご質問があれば、support@gravwell.io までお問い合わせください。

## メトリクスレポートの制限

お客様は gravwell.conf で `Disable-Stats-Report=true` を設定することで、メトリクスレポートを最小限にし、CustomerUUID、CustomerNumber、BuildVer、ApiVer、LicenseTimeLeft、LicenseHash フィールドを含む、正しいライセンスがインストールされ、システムが稼働していることを確認するのに十分な情報を提供することができます。

正しいライセンスがインストールされ、システムが稼働しているかどうかを確認するのに十分な情報です。ただし、完全な統計レポートを有効にしておいていただけると幸いです。前述したように、これらの統計レポートは、どの機能が最も使用されているか、Gravwellがどのようなシステム上で動作しているか、どのくらいのRAMを使用しているかなどを把握するのに役立ちます。