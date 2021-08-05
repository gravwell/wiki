## Hexlify

hexlifyモジュールは、データをASCIIの16進表現にエンコードするために使用されます。  このモジュールは、未知のデータ型に取り組み、バイナリデータを処理する方法を学ぶときに役立ちます。  たとえば、canbusデータから抽出された未知の列挙値をエンコードすることがあります。  ほとんどの製造元はcanbusの仕様を公開していませんが、IDから抽出して16進数でエンコードすることで、予測可能なパターンで変化している値を識別し、パラメータを識別するのに役立ちます。  これは、Gravwellチームが、Fiat Chrysler of Americaのcanbus IDにアクセスせずに、RAM 1500トラックのガスレベル、速度、およびスロットル位置のPDUを導き出した方法です。

### サポートされているオプション

* `-d`: intをASCII 16進数としてエンコードするのではなく、ASCII 16進数を整数にデコードします。


### すべてのデータを16進化するための検索例

```
tag=stuff hexlify
```

### 1つの列挙値を16進数にするための検索例

```
tag=CAN canbus ID Data | hexlify Data | table ID Data
```

### すべてのデータを16進数にして新しい名前に割り当てるための検索例

```
tag=stuff hexlify DATA as hexdata | table DATA hexdata
```

### いくつかの列挙値を再割り当てで16進数化するための検索例

```
tag=CAN canbus ID Data | hexlify ID as hexid Data as hexdata | table ID hexid DATA hexdata
```

### 16進データのデコード例

```
tag=apache json val | hexlify -d val as decodedval | table val decodedval
```
