## maclookup 

maclookupモジュールは、カスタムMACプレフィックスデータベースを使用して、MACアドレスのブロックの所有者に関するメーカー、住所、国の情報を抽出します。 

### データベースの設定

maclookupモジュールを使用する前に、macdbデータベースを含む[resource](#!resources/resources.md)をインストールする必要があります。 

デフォルトでは、maclookupモジュールは、macdbデータベースが "macdb"という名前のリソースにあることを想定しています。これにより、リソース名を明示的に指定することなく抽出を行うことができるようになります。

![](maclookup.png)

### サポートされているオプション

* `-r <arg>`: "-r"オプションは、macdb データベースを含むリソース名または UUID を指定します。 "-r"が指定されない場合、geoipモジュールはデフォルトの "macdb"リソース名を使用します。
* `-s`: "-s"オプションは、maclookup モジュールを厳密モードで動作するように指定します。 厳密モードでは、指定された演算子のいずれかがMACを解決できない場合、そのエントリーは削除されます。

### 処理演算子

maclookup抽出器は、モジュール内で非常に高速なフィルタリングを可能にする直接演算子をサポートしています。これらのフィルタを使用すると、メーカー、住所、または国に基づいてエントリを高速にフィルタリングできます。抽出フィルタは、「等しい(==)」、「等しくない(!=)」、「含む(~)」の各演算子をサポートしています。maclookupモジュールの1回の起動で複数の演算子を指定することができ、出力される列挙値の名前は "as"ディレクティブを使って変更することができます。 

| 演算子 | 名前 | 説明 |
|----------|------|-------------|
| == | 等しい | フィールドは等しくなければならない
| != | 等しくない | フィールドは等しくてはいけない
| ~ | 含む | フィールドはそのメンバーでなければならない
| !~ | 含まない | フィールドはそのメンバーであってはならない

### データ項目抽出器

maclookupデータベースでは、以下のような抽出が可能です:

| 抽出器 | 演算子 | 説明 | 例 
|-----------|-----------|-------------|----------
| Manufacturer | == != ~ !~ | メーカー名（"Apple Inc"、"IBM"など） | mac.Manufacturer~Apple
| Address | == != ~ !~ | 住所（"One Apple Park Way, Cupertino, CA"など） | mac.Address~Cupertino as addr
| Country | == != ~ !~ | 国コード | mac.Country == "US"

### 例

### 米国のPCAPから全メーカーをリストアップ

```
tag=pcap packet eth.MAC | maclookup MAC.Manufacturer MAC.Country == US | table
```

![US Manufacturers](tableByUS.png)

### PCAPのメーカーの円グラフ

```
tag=pcap packet eth.MAC | maclookup mac.Manufacturer | count by Manufacturer | chart count by Manufacturer
```

![Pie chart by manufacturer](chartByManufacturer.png)


