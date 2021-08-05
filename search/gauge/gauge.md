# ゲージ

ゲージレンダラーは、エントリを「ゲージ」として表示するのに適した1つ以上の最終値に変換するために使用される凝縮レンダラーです。 たとえば、過去1時間のブルートフォース攻撃の総数を調べて、ダッシュボードに表示したい場合があります。 これはテーブルレンダラーで実現できますが、ゲージレンダラーは一目で読みやすく、より魅力的な結果をもたらします。

## 基本的な利用方法

ゲージレンダラを使用する最もシンプルな方法は、列挙された値の引数を1つ渡すことです。

```
tag=json json class | stats mean(class) | gauge mean
```

![](gauge1.png)

歯車アイコンを選択すると、ゲージのいくつかのオプションを変更できます。 「半分」をクリックすると、ゲージ表示のスタイルが変更されます。

![](gauge2.png)

チャートタイプのドロップダウンで[ナンバーカード]を選択すると、表示が他の種類のゲージに変わります。

![](gauge3.png)

## ラベルの指定

特にダッシュボードで使用するゲージを作成する場合、デフォルトのラベルが常に理想的であるとは限りません。 より有益なラベルが必要な場合は、以下のように、マグニチュード列挙値と目的のラベルを括弧で囲みます。

```
tag=json json class | mean class | gauge (mean "Avg Class")
```

![](gauge-label.png)

## 最大および最小制限の指定

マグニチュードの列挙値と希望する最小値/最大値を括弧で括ることで、ゲージの最小値/最大値を指定することができます。

```
tag=json json class | stats mean(class) | gauge (mean 1 100000)
```

![](gauge-minmax1.png)

また、最小値と最大値を列挙値で指定することもできます。

```
tag=json json class | stats mean(class) min(class) max(class) | gauge (mean min max)
```

![](gauge-minmax2.png)

または、定数と列挙された値の組み合わせを使用します。

```
tag=json json class | stats mean(class) max(class) | gauge (mean 1 max)
```

## 最小/最大値とラベルの組み合わせ

もちろん、min/max値とラベルを指定してゲージを指定することもできます。

```
tag=json json class | mean class | gauge (mean 0 100000 "Avg Class")
```

![](gauge-label2.png)

## 複数のゲージ

複数の列挙値をリストして、ゲージに複数の針を配置できます。

```
tag=json json class | stats mean(class) stddev(class) | gauge mean stddev
```

![](gauge-multi1.png)

必要に応じて、各針の最小/最大値を個別に指定できますが、デフォルトのシングルゲージレンダラーは、他の値を無視して、表示に最小の最小値と最大の最大値を選択することに注意してください。 そのため、構成メニューで「複数のゲージ」オプションを選択することをお勧めします。

```
tag=json json class | stats mean(class) stddev(class) min(class) max(class) | gauge (mean min max) (stddev 1 35000)
```

![](gauge-multi2.png)

レンダラーは、複数のアイテムがある「ナンバーカード」モードでも適切に動作します。

![](gauge-multi3.png)

## キー付きマルチゲージ

マグニチュードを指定した場合、ゲージはキーの組み合わせごとに値を出力します。例えば、複数の都市から天気データを取得して、都市ごとの平均値を求め、その平均値をゲージに渡すことができます。

```
tag=weather json main.temp name | stats mean(temp) by name | gauge mean
```

![](keyed1.png)

ラベルを指定すると適切に使用されます。

```
tag=weather json main.temp name | stats mean(temp) by name | gauge (mean "Fahrenheit temp")
```

![](keyed2.png)
