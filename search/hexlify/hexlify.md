## Hexlify

hexlifyモジュールは、データをASCIIの16進数表現にエンコードするために使用します。 このモジュールは、未知のデータタイプに取り組むときや、バイナリデータの処理方法を学ぶときに役立ちます。 例えば、canbusデータから抽出した未知の列挙値をエンコードする場合があります。 ほとんどのメーカーはcanbusの仕様を公開していませんが、IDから抽出して16進数でエンコードすることで、予測可能なパターンで変化している値を特定し、パラメータを特定するのに役立ちます。 Gravwellチームは、Fiat Chrysler of Americaのcanbus IDにアクセスすることなく、RAM 1500トラックのガスレベル、速度、スロットルポジションのPDUを導き出したのです。

### サポートされているオプション

* `-d`: intをASCIIの16進数でエンコードするのではなく、ASCIIの16進数を整数にデコードします。


### すべてのデータを16進数にするための検索例

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

### 複数の列挙値を再割り当てで16進数にするための検索例

```
tag=CAN canbus ID Data | hexlify ID as hexid Data as hexdata | table ID hexid DATA hexdata
```

### 16進数データのデコード例

```
tag=apache json val | hexlify -d val as decodedval | table val decodedval
```
