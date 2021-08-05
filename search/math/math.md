# 数学モジュール

数学モジュールは、パイプライン上で動作し、統計的な分析を行います。また、時間軸に沿って情報を凝縮する際にも重要な役割を果たしています。例えば、温度を1秒間に10回計測しているが、ユーザーからは1秒単位で表示してほしいという要望があった場合、そのデータを凝縮するために数学モジュールが使われます。

## 和

和のモジュールは、レコードの値の和をとります。これはデフォルトの動作であり、直接呼び出されることはないでしょう。

ある MAC アドレスがネットワーク上で送信したデータを合計する検索例がこちらです。

```
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

## 平均

平均モジュールは、レコードの平均値を返します。
車両の回転数をチャート化した検索例がこちらです。

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | mean RPM | chart mean
```

## カウント

カウントモジュールは、レコードのインスタンスをカウントします。レコード内のデータに対する演算処理は行いません。

Linux マシンから sudo コマンドをカウントする検索例がこちらです。

```
grep sudo | regex "COMMAND=(?P<command>\S+)" | count by command | chart count by command
```

これは、ある MAC アドレスがネットワーク上で何個のパケットを送信したかを示す検索例です（和のモジュールの例で示した各パケットのサイズとは異なります）。

```
tag=pcap eth.SrcMAC eth.Length | sum Length by SrcMAC | chart sum by SrcMAC
```

## 最大値

最大値モジュールは、見られる最大値を返します。

フリート全体の各車両の最大回転数を表にした検索例がこちらです。

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | max RPM by SRC | table SRC max
```

## 最小値

最小値モジュールは、見られる最小値を返します。

フリート全体の各車両の最低回転数の表を表示した検索例がこちらです。

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | min RPM by SRC | table SRC min
```

## 分散

分散モジュールは、レコードの分散情報を返します。これは、変化の割合を強調するのに便利です。

トヨタ車のスロットルデータの分散をグラフ化した検索例がこちらです。

```
tag=CAN canbus ID=0x2C1 | slice byte(data[6:7]) as throttle | variance throttle | chart variance
```

## stddev

標準偏差です。

stddev モジュールは、レコードの標準偏差情報を返します。これは変則的なイベントを強調するのに便利です。

外れ値であるRPM信号をグラフ化する検索の例がこちらです。

```
tag=CAN canbus ID=0x2C4 | slice uint16BE(data[0:2]) as RPM | stddev RPM | chart stddev
```

## unique

unique モジュールは、クエリデータの中の重複したエントリを排除します。単に `unique` を指定すると、各エントリのデータのハッシュに基づいて重複エントリをチェックします。一方、1つ以上の列挙型の値の名前を指定すると、unique は列挙型の値だけでフィルタリングを行います。この違いを説明すると、次のようになります。

```
tag=pcap packet tcp.DstPort | unique
```

```
tag=pcap packet tcp.DstPort | unique DstPort
```

最初のクエリは、パケットの内容全体を見て、重複するパケットをフィルタリングします。パケットには通常、シーケンス番号などが含まれているため、これではあまり効果がありません。2つ目のクエリは、抽出された DstPort 列挙値を使用して一意性をテストします。これは、例えば、TCP ポート 80 宛ての最初のパケットは通過しますが、ポート 80 宛てのそれ以降のパケットはすべてドロップされることを意味します。

複数の引数を指定すると、unique モジュールはそれらの引数のユニークな組み合わせをそれぞれ探します。

```
tag=pcap packet tcp.DstPort tcp.DstIP | eval DstPort < 1024 | unique DstPort DstIP | table DstIP DstPort
```

上記の検索では、ポートが 1024 以下であれば、IP とポートのユニークな組み合わせがすべて出力されます。これは、例えば、ネットワーク上のサーバを探すのに便利な方法です。

## エントロピー

エントロピーモジュールはフィールド値の時間的なエントロピーを計算します。 引数なしで `entropy` を指定すると、検索範囲内のすべてのエントリのデータセクションのエントロピーを生成します。 エントロピーモジュールは時間検索モードをサポートしており、時系列でエントロピーのチャートを作成することができます。 エントロピーは、他の数学モジュールと同様に、列挙値や複数のキーを使用したグループ化についても操作できます。 エントロピーの出力値は 0 から 1 の間です。

こちらが、ポートに基づいて TCP パケットのペイロードのエントロピーを計算し、グラフ化するクエリの例です。

```
tag=pcap packet tcp.Port tcp.Payload | entropy Payload by Port | chart entropy by Port
```

こちらは、ホスト別に URL のエントロピーを計算し、エントロピー値の大きい順にリストをソートするクエリの例です。

```
tag=pcap packet tcp.Port==80 ipv4.IP !~ 10.0.0.0/8 tcp.Payload | grep -e Payload GET PUT HEAD POST | regex -e Payload "[A-Z]+\s(?P<url>\S+)\sHTTP\/" | entropy url by IP | table IP entropy
```

エントロピーモジュールは、`-p` フラグを取ることができます。これは、エントロピーを通常のようにウィンドウ上で計算するのではなく、エントリごとに計算するように指示します。以下は、ウィンドウの各ログエントリのエントロピーを計算し、そのエントロピーとデータを表示します。

```
tag=winlog entropy -p | table DATA entropy
```
