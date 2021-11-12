## Entropy

`entropy`モジュールはフィールド値の時間的なエントロピーを計算します。`entropy`を引数なしで指定すると、検索範囲内の全てのエントリDATAフィールドのエントロピーを生成します。`entropy`モジュールは一時的な検索モードをサポートしており、時間経過に伴うエントロピーのチャートを作成することができます。また、`entropy`は列挙型の値を操作したり、列挙型の値でグループ化したりすることもできます。出力値は0から1の間です。

構文：

```
entropy [enumerated value] [by ...] [over <duration>]
```

`entropy`モジュールの構文では、エントロピーを計算する列挙型の値を指定することができます。指定しない場合、`entropy`はDATAフィールド全体のエントロピーを計算します。また、このモジュールでは、`by`キーワードを使って、グループ化するための1つ以上の引数を指定することができます。例えば、列挙された値 `foo` に対してエントロピーを計算し、 `bar` と `baz` でグループ化することができます。

```
tag=gravwell entropy foo by bar baz
```

キーワード `over` を使って、クエリを任意の時間範囲でグループ化することができます。

```
tag=gravwell entropy over 10m
```

すべての引数は任意です。

### サポートされるオプション

`entropy` には利用できるフラグがありません。

### 例

このクエリは、ポートに基づいてTCPパケット・ペイロードのエントロピーを計算し、グラフ化します。

```
tag=pcap packet tcp.Port tcp.Payload | entropy Payload by Port | chart entropy by Port
```

ホスト別にURLSのエントロピーを計算し、エントロピー値の大きい順にリストをソートするクエリの例です。

```
tag=pcap packet tcp.Port==80 ipv4.IP !~ 10.0.0.0/8 tcp.Payload | grep -e Payload GET PUT HEAD POST | regex -e Payload "[A-Z]+\s(?P<url>\S+)\sHTTP\/" | entropy url by IP | table IP entropy
```
