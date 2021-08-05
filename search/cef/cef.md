## CEF

CEFまたはCommon Event Formatは、いくつかのデータプロバイダおよび多くのデータ強化システムによって使用されています。  Gravwell CEFパーサーは、共通のCEFヘッダー変数、および任意のキー値ペアのセットを非常に高速に抽出できるように設計されています。  厳密な仕様であるCEFは、一連の既知のキー名を技術的に定義していますが、CEFを生成し、それらのキーのセットを厳密に保持するデータ生成製品はまだ見ていません。

### 標準CEFヘッダキー名

CEFには、すべてのCEFレコードに含まれるべき標準化されたヘッダー値のセットが含まれています。  ヘッダーレコード名は次のとおりです:

* Version
* DeviceVendor
* DeviceProduct
* Signature
* Name
* Severity

Gravwell CEFパーサーを使用すると、ヘッダー名と衝突するサブメンバーキーを柔軟に抽出できます。  デフォルトでは、検索でキーのVersionが指定されている場合、ヘッダー値Versionが抽出されます。  ただし、不正なデータソースがVersionという名前のキーを持つCEFレコードを提供している場合でも、名前の前に "EXT"を追加することでアクセスできます。

抽出されたヘッダーとキーの値は、キーまたはヘッダーと同じ名前の列挙値に抽出されます。  ただし、"as"構文を使用すると、抽出された値は任意の値に名前を変更できます。  gravwell CEFパーサーは、柔軟性を念頭に置いて設計されており、不完全に形成されたCEFレコードおよび厳密に定義されていないCEF仕様に技術的に違反するレコードを処理できます。

## サポートされているオプション

* `-e`: "-e"オプションは、CEFモジュールが列挙値で動作するように指定します。  列挙値を操作すると、アップストリームモジュールを使用してCEFレコードを抽出した場合に役立ちます。  たとえば、生のPCAPからCEFレコードを抽出し、そのレコードをCEFモジュールに渡すことができます。

## 処理オペレータ

各CEFフィールドは、高速フィルターとして機能できる一連の演算子をサポートします。  各演算子でサポートされているフィルタは、フィールドのデータ型によって決まります。

| オペレーター | 名 | 説明 |
|----------|------|-------------|
| == | 等しい | フィールドは等しくなければなりません
| != | 等しくない | フィールドは等しくてはいけません
| ~  | 含む | フィールドはサブシーケンスを含まなければなりません
| !~ | 含まない | フィールドはサブシーケンスを含んではいけません

可能であれば、CEFインラインフィルタを使用して、CEFモジュールがダウンストリームモジュールに頼るのではなく、すぐに必要な種類のレコードに集中できるようにします。 インラインフィルタを使用すると、最悪の場合の操作時の速度が向上するだけでなく、フィールドアクセラレータを有効にしたときにそれを利用することもできます。

### 例

次のCEFレコードからデバイスの製造元、製品、重大度、およびmsgを抽出し、重大度が> 7のレコードだけを含むテーブルを作成したい場合:

CEFレコードの例:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

クエリ:

```
tag=firewall cef DeviceVendor DeviceProduct Severity==7 msg | table DeviceVendor DeviceProduct msg
```

ただし、Versionというキー値を含む技術的に無効なCEFレコードがあり、ヘッダVersionの代わりにExt指定子を使用してその値にアクセスできるようにしたい場合は、

形成が不十分なCEFレコード:

```
CEF:0|Citrix|NetScaler|NS11.0|APPFW|APPFW_STARTURL|6|src=192.168.1.1 Version=11.0 method=GET request=http://totallynotmalware.safedomain.cn/stuff.html msg=Banned domain request attempt cs1=FIREWALL cs2=APP cs3=deadbeef1009 cs4=WARN cs5=2018 act=blocked
```

クエリ:

```
tag=firewall cef DeviceVendor==Citrix DeviceProduct==NetScaler Severity Ext.Version msg ~ Banned | table DeviceVendor DeviceProduct msg Version
```

クエリは "0"ではなくversionの値 "11.0"を抽出しますが、ヘッダ値が必要な場合は、ヘッダバージョンとキーVersionの両方を取得するために "as"構文を使用することもできます。  "~"インラインフィルタは、メッセージフィールドに "Banned"という単語を含むレコードのみが必要であることを示しています。

```
tag=firewall cef DeviceVendor DeviceProduct Severity Version as hdrversion Ext.Version msg | eval Severity > 7 | table DeviceVendor DeviceProduct msg Version hdrversion
```
