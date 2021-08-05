## Limit

limitモジュールは、指定した数のエントリを通過させ、それ以上は通過させないようにします。これは、例えば正規表現をテストしているときなど、クエリを構築している間に特に便利かもしれません。

構文は単純です: `limit <n>`, ここで `n` は通過させるエントリの最大数、`limit <n> <m>` は N 番目から M 番目までのエントリを通過させるものです。

具体的には、`limit X Y` はエントリ `[X,Y)` を渡す。つまり、limitは最初の項を含み、2番目の項を含まない。つまり、集合 `[a,b,c,d,e,f]` が与えられると、`limit 2 5` は `[c d e]` を返す。

例えば、10個のパケットのペイロードを見るには:

```
tag=pcap packet tcp.Payload | limit 10 | table Payload
```

パケット5から10を見るには:

```
tag=pcap packet tcp.Payload | limit 5 10 | table Payload
```

