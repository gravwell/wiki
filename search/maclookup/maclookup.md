## maclookup

maclookupモジュールは、カスタムMACプレフィックスデータベースを使用して、MACアドレスのブロックの所有者に関するManufacturer、Address、Countryの情報を抽出します。

### データベースのセットアップ

maclookupモジュールを使用する前に、macdbデータベースを含む[リソース](#!resources/resources.md)をインストールする必要があります。

デフォルトでは、maclookupモジュールは、macdbデータベースが "macdb" という名前のリソースにあることを想定しています。これにより、リソース名を明示的に指定することなく抽出を行うことができるようになります。

![](maclookup.png)

### サポートされているオプション

* `-r <arg>`: 「-r」オプションは、macdb データベースを含むリソース名または UUID を指定します。「-r」が指定されない場合、geoip モジュールはデフォルトの「macdb」リソース名を使用します。
* `-s`: 「-s」オプションは、maclookup モジュールをストリクトモードで動作させることを指定します。 ストリクトモードでは、指定された演算子のいずれかがMACを解決できない場合、そのエントリーはドロップされます。

### 処理演算子

maclookup 抽出器は、モジュール内で非常に高速なフィルタリングを可能にする直接演算子をサポートしています。これらのフィルターを使用すると、製造者、住所、または国に基づいてエントリーを高速にフィルタリングできます。抽出フィルターは、等しい（==）、等しくない（!=）、含まれる（~）の各演算子をサポートしています。maclookup モジュールの1回の起動で複数の演算子を指定することができ、出力される列挙値の名前は as ディレクティブを使って変更することができます。

| 演算子 | 名称 | 意味
|----------|------|-------------
| == | 等しい | フィールドは等しい
| != | 等しくない | フィールドは等しくない
| ~ | 含まれる | フィールドはそれに含まれる
| !~ | 含まれない | フィールドはそれに含まれない

### データ項目の抽出器

maclookupデータベースでは、以下のような抽出が可能です。

| 抽出器 | 演算子 | 意味 | 例
|-----------|-----------|-------------|----------
| Manufacturer | == != ~ !~ | メーカー名（"Apple Inc"、"IBM" など） | mac.Manufacturer~Apple
| Address | == != ~ !~ | 住所（"One Apple Park Way, Cupertino, CA" など） | mac.Address~Cupertino as addr
| Country | == != ~ !~ | 国別コード | mac.Country == "US"

### 例

### 米国の PCAP から全メーカーをリストアップ

```
tag=pcap packet eth.MAC | maclookup MAC.Manufacturer MAC.Country == US | table
```

![US Manufacturers](tableByUS.png)

### PCAP のメーカーの円グラフ

```
tag=pcap packet eth.MAC | maclookup mac.Manufacturer | count by Manufacturer | chart count by Manufacturer
```

![Pie chart by manufacturer](chartByManufacturer.png)
