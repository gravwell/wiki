# Abs

absモジュールは、指定された列挙値の絶対値を取ります。 引数が1つしか指定されていない場合、列挙値はその絶対値で上書きされます:

```
tag=default json offset | abs offset
```

元の名前をそのままにしておく 'destination'列挙値の名前を指定することもできます:

```
tag=default json offset | abs offset as offsetAbsolute | chart offsetAbsolute
```
