# ゲージとナンバーカード

ゲージレンダラーは、エントリを"ゲージ"として表示するのに適した1つまたは複数の最終値に変換するために使用される凝縮レンダラーです。例えば、過去1時間のブルートフォース攻撃の総数を求め、ダッシュボードに表示したいとします。これはテーブルレンダラーでも実現できますが、ゲージレンダラーの方が一目で読みやすい魅力的な結果となります。

![](gauge-example.png)

ナンバーカードレンダラーは、ゲージの"エイリアス"です。ゲージとまったく同じ構文を持ちますが、デフォルトではゲージの代わりにシンプルな数字のタイルを表示します:

![](numbercard-example.png)

## 基本的な使用方法

ゲージレンダラーを使用する最もシンプルな方法は、単一の列挙値を引数として渡すことです:

```
tag=json json class | stats mean(class) | gauge mean
```

![](gauge1.png)

歯車のアイコンを選択すると、ゲージのいくつかのオプションを変更することができます。"Half"をクリックすると、ゲージの表示スタイルが変わります:

![](gauge2.png)

チャートの種類のドロップダウンで`ナンバーカード`を選択すると、他の種類のゲージに表示が変わります:

![](gauge3.png)

もし、`gauge`の代わりに`numbercard`を指定すると、デフォルトで`ナンバーカード`ビューになります。:


```
tag=json json class | stats mean(class) | numbercard mean
```

![](numbercard-basic.png)

## ラベルの指定

特にダッシュボードで使用するゲージを作成する場合、デフォルトのラベルは必ずしも理想的ではありません。より分かりやすいラベルが必要な場合は、以下のようにマグニチュードの列挙値と必要なラベルを括弧で囲みます:

```
tag=json json class | mean class | gauge (mean "Avg Class")
```

![](gauge-label.png)

## 最大と最小制限の指定

マグニチュードの列挙値と希望する最小値/最大値を括弧で括ることで、ゲージの最小値/最大値を指定することができます:

```
tag=json json class | stats mean(class) | gauge (mean 1 100000)
```

![](gauge-minmax1.png)

また、最小値と最大値を列挙値で指定することもできます:

```
tag=json json class | stats mean(class) min(class) max(class) | gauge (mean min max)
```

![](gauge-minmax2.png)

または、定数と列挙値の組み合わせを使用します:

```
tag=json json class | stats mean(class) max(class) | gauge (mean 1 max)
```

## 最小/最大値とラベルの組み合わせ

もちろん、min/max値とラベルを指定してゲージを指定することもできます:

```
tag=json json class | mean class | gauge (mean 0 100000 "Avg Class")
```

![](gauge-label2.png)

## 複数のゲージ

複数の列挙値をリストして、ゲージに複数の針を配置できます:

```
tag=json json class | stats mean(class) stddev(class) | gauge mean stddev
```

![](gauge-multi1.png)

各針の最小値／最大値を別々に指定することもできますが、デフォルトのシングルゲージのレンダラーでは、最小値と最大値のうち最も低いものが選択され、他のものは無視されることに注意してください。そのため、設定メニューで"multiple gauges"を選択することをお勧めします:

```
tag=json json class | stats mean(class) stddev(class) min(class) max(class) | gauge (mean min max) (stddev 1 35000)
```

![](gauge-multi2.png)

また、複数のアイテムがある"number card"モードでも、レンダラーは適切に動作します:

![](gauge-multi3.png)

## キー付きマルチゲージ

*キー*付きのマグニチュードを指定すると、ゲージはキーの組み合わせごとに値を出力します。例えば、複数の都市の気象データを取得して、*都市ごとの*平均値を求め、その平均値を `gauge` や `numbercard` に渡すことができます:

```
tag=weather json main.temp name | stats mean(temp) by name | numbercard mean
```

![](keyed1.png)

ラベルを指定すると適切に使用されます:

```
tag=weather json main.temp name | stats mean(temp) by name | gauge (mean "Fahrenheit temp")
```

![](keyed2.png)
