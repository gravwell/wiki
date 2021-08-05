## lookup

lookup モジュールは、リソースに保存されている静的な lookup テーブルから、データの濃縮やトランスレーションを行うために使用されます。1つまたは複数の*列挙値*の内容は、一致するものが見つかるまで *match 列*の値と比較され、その行の *extract 列*の値が別の列挙値に抽出されます。

```
lookup -r <resource name> <enumerated value> <column to match> <column to extract> as <valuename>
```

注意：シンタックスに ```as <valuename>``` を追加しない場合、lookup は抽出する列の名前を持つ列挙値を作成します。

lookup モジュールの1回の呼び出しで、複数の lookup の操作を、追加の操作を連結して指定することができます。

```
lookup -r mytable A B C as foo D E F
```

一致するごとに複数の列を抽出することもできます。次の例では、列挙値Aの内容をB列の値と照合します。一致した場合、C列とD列の両方を抽出します。

```
lookup -r mytable A B (C as foo D as bar)
```

Lookup はベクトル化された一致もサポートしています。つまり、列のセットに対して列挙された値のセットをマッチさせることができます。 ベクトル化された一致は、マッチリストと抽出リストを指定して実行されます。 ベクトル化された一致を実行する場合、抽出された列挙値の数は、一致する列の数と一致する必要があります。

```
lookup -r mytable [A B] [A B] (C as foo D as bar)
```

### サポートされているオプション
* `-r <arg>`: 「-r」オプションは、どの lookup リソースを使用してデータを充実させるかを lookup モジュールに通知します。
* `-s`: 「-s」オプションは、lookup モジュールが、すべての抽出が成功しないとエントリーが削除されることを要求するように指定します。
* `-v`: 「-v」フラグは、lookup モジュールのフローロジックを反転させます。つまり、マッチしたものは抑制され、マッチしなかったものは引き継がれます。 「-v」フラグと「-s」フラグを組み合わせることで、指定した lookup テーブルに存在しない値のみを渡す、基本的なホワイトリストを提供することができます。

注意： `-s` または `-v` フラグを使用する際に、抽出を行わないように指定することは合法です。 この操作は、ホワイトリストやブラックリストを実行するときに便利です。

列 `X` と `Y` に列挙値 `A` と `B` が存在することを保証するものの、データをリッチ化しない例を示します。

```
lookup -v -r mytable [A B] [X Y] ()
```

### lookupdata リソースの設定

検索結果のデータは、互換性のあるレンダリングモジュール（table モジュールなど）からダウンロードしてリソースに保存し、共有・活用することができます。検索結果ページのメニューで「LOOKUPDATA」を選択すると、この形式で検索結果の表をダウンロードすることができます。

![Lookup Download](lookup-download.png)

[テーブルレンダラー](#!search/table/table.md)には、`-save`オプションもあり、検索結果のテーブルを自動的にリソースとして保存し、後で lookup で利用できるようにします。

```
tag=syslog regex "DHCPACK on (?P<ip>\S+) to (?P<mac>\S+)" | unique ip mac | table -save ip2mac ip mac
```

上記の例では、テーブルレンダラーが自動的に「ip2mac」というリソースを作成し、DHCPログから得られたIPアドレスとMACアドレスのマッピングを含んでいます。

#### CSV テーブル

CSV データは検索モジュールにも使用できます。Gravwell の検索モジュールで CSV ファイルをリソースとして使用するためには、CSV には列に対応するユニークなヘッダーが含まれていなければなりません。

### 例

#### 基本的な抽出

この例では、以下の csv から作成された「macresolution」というリソースがあります。
```
mac,hostname
mobile-device-1,40:b0:fa:d7:af:01
desktop-1,64:bc:0c:87:bc:71
mobile-device-2,40:b0:fa:d7:ae:02
desktop-2,64:bc:0c:87:9a:11
```

次に、パケットデータから検索を行い、lookup モジュールを使用して、データストリームにホスト名を追加します。ここでは、「devicename」という列挙値に割り当てています。

```
tag=pcap packet eth.SrcMAC | count by SrcMAC | lookup -r macresolution SrcMAC mac hostname as devicename |  table SrcMAC devicename count
```

この結果、次のような内容を含むテーブルができあがります。

```
64:bc:0c:87:bc:71	|	desktop-1       	|	52183
40:b0:fa:d7:ae:02	|	mobile-device-2 	|	21278
64:bc:0c:87:9a:11	|	desktop-2       	|	 2901
40:b0:fa:d7:af:01	|	mobile-device-1 	|	  927
```

#### ホワイトリスティング

先ほどの「macresolution」の表と同じものを使用します。

```
tag=pcap packet eth.SrcMAC | count by SrcMAC | lookup -v -s -r macresolution SrcMAC mac hostname |  table SrcMAC count
```

この結果、lookup リストに含まれていない MAC アドレスがテーブルに表示されます。 システム管理者は、「-v」および「-s」フラグを使用して、ネットワーク上の新しいデバイスや、イベントストリームの新しいログの基本的なホワイトリストとの識別を行うことができます。

```
64:bc:0c:87:bc:60	|	24382
40:b0:fa:d7:ae:13	|	93485
64:bc:0c:87:9a:02	|	11239
40:b0:fa:d7:af:fe	|	   21
```

#### 多重抽出

次のような「places」という名前の lookup テーブルを考えてみましょう。

```
name,lat,long
Albuquerque,35.122,-106.553
Santa Fe,35.6682,-105.96
Sacramento,38.527,-121.347
```

都市名を指定すると、緯度と経度の両方の値を抽出できるようにしたいと思います。City フィールドを含む JSON エントリーを扱う場合、以下のクエリはまさにそれを実行します。

```
tag=default json City | lookup -r places City name (lat long) | table
```

列挙値 City をテーブルの name 列と照合し、一致したものがあれば、以下のように緯度と経度の両方を抽出します。

![](city.png)
