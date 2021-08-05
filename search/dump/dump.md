## Dump

`dump`は、ダッシュボードやビューを構築する際に、他の静的なソースから既にインポートされているデータや、スケジュールされたクエリを集約したデータを操作する必要がある場合に特に便利です。つまり、CSVやルックアップテーブルをクエリパイプラインに注入して、あたかも本格的な検索エントリであるかのように行に対して操作を行うことができます。

Gravwellの検索システムは、クエリからインデクサーや保存されたデータを利用する必要があるかどうかを推論します。 `dump`モジュールが列挙された値の要件をすべて満たすことができ、他の抽出モジュールが存在しない場合、クエリはインデクサー上では全く実行されません。 しかし、ダンプモジュールは `-p` 実行フラグを用いて他のデータと一緒に値を挿入することをサポートしています。

### サポートされているオプション

* `-r`: "-r"オプションは引数を必要とし、データを抽出するリソースの名前またはUUIDを指定する。 リソースを指定し、現在のユーザがそのリソースへの読み込みアクセス権を持っていなければならない
* `-p`: "-p "は、ダンプモジュールが他のデータと一緒に値を注入していることを示し、パイプラインはインデクサーに問い合わせを行う必要がある。
* `-t`: インジェクションされたエントリのタイムスタンプとして使用するカラムをリソースから選択する。`tz`や`-f`と組み合わせることができる。このフラグは`-p`とは併用できない。
* `-tz`: [tz database format](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)で、`-t`で処理されるタイムスタンプのタイムゾーンを設定する。例："America/Denver", "UTC", "Atlantic/Reykjavik"
* `-f`: `t`を使用してタイムスタンプを解析する際に使用するフォーマットを指定する。フォーマットは、[Go time library](https://golang.org/pkg/time/#pkg-constantsで使用されている "Mon Jan 2 15:04:05 MST 2006 "という特定の時刻を文字列で表現したもの。詳細な例については、リンク先のドキュメントを参照。

### Filtering and Column Inclusion

`dump`モジュールは、カラムが指定されていない場合、あるいは指定されたすべてのカラムがフィルタ操作を行っている場合に、CSVまたはルックアップテーブルリソースからすべてのカラムを抽出します。 例えば、`dump -r hosts hostname` はホスト名の列のみを抽出することを意味し、`dump -r hosts hostname == "ad.example.com"` は指定されたすべての列にフィルタ操作が付加されているので、すべての列を抽出することを意味します。 クエリ `dump -r hosts hostname == "ad.example.com" IP` は `hostname` と `IP` のカラムのみを抽出します。

#### サポートされているフィルタ演算子

`dump`モジュールは、等しさに基づくフィルタリングを可能にします。 等値を指定するフィルタが有効な場合（"equal"、"not equal"、"contains"、"not contains"）、フィルタの指定に失敗した行のカラムはパイプラインに注入されません。

| 演算子 | 名前 | 説明 |
|----------|------|-------------|
| == | 等しい | フィールドは等しくなければならない
| != | 等しくない | フィールドは同じであってはならない
| ~ | サブセット｜フィールドには値が含まれている
| !~ | サブセットではない | フィールドに値が含まれていない


#### フィルタリング例

`hostname`カラムが "ad.example.com" と等しくない場合、`hosts` リソースからすべてのカラムを取得します:
```
dump -r hosts hostname != "ad.example.com"
```

ホスト名に "example.com" が含まれる `hosts` リソースから `hostname`, `IP`, `MAC` カラムのみを取得します:
```
dump -r hosts hostname ~ "example.com" IP MAC
```

ホスト名が空ではなく、所有する組織が "finance" である `hosts` リソースからすべてのカラムを取得します。:
```
dump -r hosts hostname != "" org=="finance"
```

### クエリ例

パイプラインにリソース全体をダンプし、テーブルモジュールを生成します:

```
dump -r devlookup | table
```

![Table produced from resource](dump_table.png)


ホストカラムに "Chrome "が含まれているリソース全体をダンプする:

```
dump -r devlookup Host ~ Chrome | table
```

![Table produced from resource with filters](dump_filter_table.png)

リソースをダンプし、他のモジュールでエントリを操作する:

```
dump -r devlookup Host ~ Chrome | maclookup -r mac_prefixes MAC.Manufacturer MAC.Country | table
```

![Table produced dump and maclookup](dump_filter_lookup_table.png)
