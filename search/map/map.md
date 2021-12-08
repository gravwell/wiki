# Mapモジュール

`pointmap` と `heatmap` のレンダラモジュールは、検索結果を地図上に変換します。どちらも、列挙値の位置に基づいて、地図上にエントリを配置します。デフォルトでは、[geoip](#!search/geoip/geoip.md) 検索モジュールで設定された 'Location' と呼ばれる列挙値を検索します。ロケーションは、以下のようにして明示的に指定することもできます:

* `-loc <enumerated value>` は、デフォルトの `Location` ではなく、指定された列挙値の中からロケーションを探すようにモジュールに指示します。
* `-lat <enumerated value> -long <enumerated value>` は、緯度と経度の値を別々に探すようにモジュールに指示します。これらの値は、浮動小数点数（`geoip`モジュールによって提供されるもの）または他のソースからの文字列です。

マップには最大 1000 地点が表示されます。地図の一部を拡大すると、その範囲内で最大 1000 地点が表示されます。

# pointmap

pointmap は、エントリを地図上の明確なマーカーに変換します。追加の列挙値名が指定された場合、ポイントがクリックされたときにその内容が表示されます。

以下の検索では、netflowレコードに取り込まれたすべてのIPアドレスのマップが表示されます:

```
tag=netflow netflow IP | geoip IP.Location | pointmap IP
```

![](map1.png)

各IPからのバイト数を集計し、IPとBytesの列挙値をpointmapの引数に追加すると、ポイントをクリックしたときに表示されます(誰と話しているのかわかるように、ASN Organizationもついかしています):

```
tag=netflow netflow IP Bytes | sum Bytes by IP | geoip IP.Location | geoip -r maxmindASN IP.ASNOrg | pointmap IP Bytes ASNOrg
```

![](map2.png)

# heatmap

Heatmapはpointmapと同様の動作をしますが、引数として0または1の追加の列挙値を受け取ります。列挙値の引数が与えられない場合は、各場所のエントリ数を'heat'としてヒートマップを生成する。netflow レコードを使用したこの例では、'heatは、ある場所からの接続数を表しています:

```
tag=netflow netflow IP | geoip IP.Lat IP.Long | heatmap -lat Lat -long Long
```

![](map3.png)

総バイト数を引数に加えれば、'heat'は接続数ではなく、接続で送信されたバイト数から導き出されます:

```
tag=netflow netflow IP Bytes | sum Bytes by IP | geoip IP.Location | heatmap sum
```

![](map4.png)

## 3Dについて

heatmap と pointmap には 3D レンダリング機能もあり、マップの右上にある"Globe" セレクターをクリックすると、マップが再描画されます。

![](selector.png)

全く同じ heatmap クエリを実行しますが、Globe システムを使ってレンダリングすると以下のようになります:

![](map5.png)

しかし、もう1つ、非常に気の利いた仕掛けがあります。ローテーションを追加することで、あなたが世界的な脅威についてリアルタイムに考えていることを、上司全員に知らせることができます。

![](rotation.gif)
