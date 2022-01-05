## Dump

`dump`モジュールは、リソースをクエリ可能なデータとして扱うように設計されています。つまり、CSVやルックアップテーブルをクエリパイプラインに注入して、あたかも本格的な検索エントリであるかのように行に対して操作を実行することができます。 dumpは、他の静的なソースからインポートされたデータや、スケジュールされたクエリを使用して集約されたデータを操作する必要のあるダッシュボードやビューを構築する際に特に役立ちます。

SolitonNKの検索システムは、クエリから、インデクサや保存データを使用する必要があるかどうかを推測します。 もし `dump` モジュールが列挙値の要件をすべて満たし、他の抽出モジュールが存在しない場合、クエリはインデクサ上ではまったく実行されません。 しかし、dumpモジュールは、以下で説明する `-p` 実行フラグによって、他のデータと一緒に値を注入することをサポートしています。

### サポートされているオプション

* `-r`: `-r` オプションには引数が必要で、データを抽出するリソースの名前または UUID を指定します。 リソースを指定するには、現在のユーザーがそのリソースに対して読み取り権限を持っている必要があります。
* `-p`: `-p`は、dumpモジュールが他のデータと一緒に値を注入していることを示し、パイプラインがインデクサに問い合わせなければならないことを示します。このフラグは `-t` と一緒に使うことはできません。
* `-t`: 注入されたエントリーのタイムスタンプとして使用するリソースのカラムを選択します。このフラグは `-tz` や `-f` と組み合わせることができます。このフラグは `-p` と併用することはできません。
* `-tz`: `-t`を使って処理されるタイムスタンプのタイムゾーンを、[tzデータベースフォーマット](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)で設定します。例えば、"America/Denver"、"UTC"、"Atlantic/Reykjavik"などです。 
* `-f`: `-t`を使ってタイムスタンプを解析するときに使うフォーマットを指定します。このフォーマットは、[Go timeライブラリ](https://golang.org/pkg/time/#pkg-constants)で使われているように、"Mon Jan 2 15:04:05 MST 2006"という特定の時間の文字列表現で構成されています。詳しい例はリンク先のドキュメントを参照してください。

### フィルタリングとカラムインクルージョン

`dump`モジュールは、CSVやルックアップテーブルのリソースから、カラムが指定されていない場合や、指定されたすべてのカラムにフィルタ演算がある場合に、すべてのカラムを抽出します。 例えば、`dump -r hosts hostname` は、ホスト名のカラムのみを抽出することを意味しますが、`dump -r hosts hostname == "ad.example.com"` は、指定されたすべてのカラムにフィルタ操作が付加されているため、すべてのカラムを抽出することを意味します。 また、`dump -r hosts hostname == "ad.example.com" IP`というクエリは、`hostname`と`IP`のカラムのみを抽出します。

#### サポートされているフィルタ演算子

`dump`モジュールでは、等値性に基づいたフィルタリングが可能です。 等価性を指定したフィルタ（"equal"、"not equal"、"contains"、"not contains"）が有効な場合、フィルタの指定に失敗した行のすべてのカラムはパイプラインに注入されません。

| 演算子 | 名前 | 説明 |
|----------|------|-------------|
| == | 等しい | フィールドは等しくなければならない
| != | 等しくない | フィールドは等しくてはいけない
| ~ | 含む | フィールドには値が含まれている
| !~ | 含まない | フィールドには値が含まれていない


#### フィルタリング例

`hosts`リソースから、`hostname`カラムが "ad.example.com"ではないすべてのカラムを取得します:
```
dump -r hosts hostname != "ad.example.com"
```

ホスト名に "example.com" が含まれる `hosts` リソースから、`hostname`, `IP`, `MAC` カラムのみを取得します:
```
dump -r hosts hostname ~ "example.com" IP MAC
```

ホスト名が空でなく、所有する組織が "finance" である `hosts` リソースからすべてのカラムを取得します:
```
dump -r hosts hostname != "" org=="finance"
```

### クエリ例

リソース全体をパイプラインにダンプし、テーブルモジュールに反映させる:

```
dump -r devlookup | table
```

![Table produced from resource](dump_table.png)


ホスト列に "Chrome"が含まれるリソース全体をダンプする:

```
dump -r devlookup Host ~ Chrome | table
```

![Table produced from resource with filters](dump_filter_table.png)

リソースをダンプし、他のモジュールと一緒にエントリを操作する:

```
dump -r devlookup Host ~ Chrome | maclookup -r mac_prefixes MAC.Manufacturer MAC.Country | table
```

![Table produced dump and maclookup](dump_filter_lookup_table.png)
