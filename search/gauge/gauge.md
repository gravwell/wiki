# ゲージ＆ナンバーカード

ゲージレンダラは、エントリを「ゲージ」として表示するのに適した1つまたは複数の最終値に変換するために使用される凝縮レンダラーです。例えば、過去1時間のブルートフォース試行回数の合計を求め、ダッシュボードに表示したいとします。これはテーブルレンダラーでも実現できますが、ゲージレンダラーの方が一目で読み取れる魅力的な結果となります。

![](gauge-example.png)

ナンバーカードレンダラーは、ゲージの「エイリアス」です。ゲージとまったく同じシンタックスを持ちますが、デフォルトではゲージの代わりにシンプルな数字のタイルを表示します。

![](numbercard-example.png)

## 基本的な使い方

ゲージレンダラーを使用する最も簡単な方法は、列挙された値の引数を1つ渡すことです。

```
tag=json json class | stats mean(class) | gauge mean
```

![](gauge1.png)

歯車のアイコンを選択すると、ゲージのいくつかのオプションを変更することができます。半分」をクリックすると、ゲージの表示スタイルが変わります:

![](gauge2.png)

チャートタイプのドロップダウンで「ナンバーカード」を選択すると、他の種類のゲージに表示が変わります:

![](gauge3.png)

もし、`gauge`の代わりに`numbercard`を指定すると、デフォルトで「ナンバーカード」ビューになります:

```
tag=json json class | stats mean(class) | numbercard mean
```

![](numbercard-basic.png)

## ラベルの指定

デフォルトのラベルは、特にダッシュボードで使用するゲージを作成する場合、必ずしも理想的ではありません。より分かりやすいラベルが必要な場合は、以下のようにマグニチュード列挙値と希望するラベルを括弧で囲みます:

```
tag=json json class | mean class | gauge (mean "Avg Class")
```

![](gauge-label.png)

## 最大値と最小値の指定

マグニチュードの列挙された値と、希望する最小・最大値を括弧で囲むことで、ゲージの最小・最大値を指定することができます。

```
tag=json json class | stats mean(class) | gauge (mean 1 100000)
```

![](gauge-minmax1.png)

また、最小値と最大値を列挙型の値で指定することもできます:

```
tag=json json class | stats mean(class) min(class) max(class) | gauge (mean min max)
```

![](gauge-minmax2.png)

また、定数と列挙型の値を組み合わせて使用することもできます:

```
tag=json json class | stats mean(class) max(class) | gauge (mean 1 max)
```

## Min/Maxとラベルの組み合わせ

もちろん、最小/最大値とラベルの両方を持つゲージを指定することもできます:

```
tag=json json class | mean class | gauge (mean 0 100000 "Avg Class")
```

![](gauge-label2.png)

## 複数のゲージ

複数の列挙型の値を列挙して、ゲージに複数の針を配置することができます:

```
tag=json json class | stats mean(class) stddev(class) | gauge mean stddev
```

![](gauge-multi1.png)

各針の最小値／最大値を別々に指定することもできますが、デフォルトのシングルゲージのレンダラーでは、最小値と最大値のうち最も低いものが選択され、他のものは無視されることに注意してください。そのため、設定メニューで「複数ゲージ」を選択するとよいでしょう。

```
tag=json json class | stats mean(class) stddev(class) min(class) max(class) | gauge (mean min max) (stddev 1 35000)
```

![](gauge-multi2.png)

また、複数のアイテムを持つ「ナンバーカード」モードでも、レンダラーは適切に動作します:

![](gauge-multi3.png)

## キー付きマルチゲージ

*キー付き*の大きさを指定すると、ゲージはキーの組み合わせごとに値を出力します。 したがって、たとえば、複数の都市から気象データを取得し、*都市ごとの*平均を見つけて、結果の平均を `gauge`または` numbercard`に渡すことができます。

```
tag=weather json main.temp name | stats mean(temp) by name | numbercard mean
```

![](keyed1.png)

ラベルを指定した場合は、そのラベルが適切に使用されます。

```
tag=weather json main.temp name | stats mean(temp) by name | gauge (mean "Fahrenheit temp")
```

![](keyed2.png)
