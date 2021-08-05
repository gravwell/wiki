# Join

joinモジュールを使用すると、2つ以上の列挙値を1つの列挙値に簡単に結合することができます。すべての列挙値型は文字列に変換されて連結されますが、バイトスライスにのみ結合でき、バイトスライスのままになるバイトスライスを除いては、バイトスライスのままになります。

以下の検索は、ネットフローレコードから宛先の IP とポートを抽出し、セミコロンを区切りとしてそれらを結合し、その結果を `dialstring` という名前の列挙値に配置します。

```
tag=netflow netflow Dst DstPort | join -s : Dst DstPort as dialstring | table Dst DstPort dialstring
```

列挙値は何個でも指定できます。列挙値の出力は `as` 引数で指定します。

## サポートされているオプション

* `-s <separator>`：結果として得られる文字列の列挙値の間に、指定された区切り文字列を配置します。指定しない場合、セパレータは使用されません。バイトスライスの場合は無視されます。

## 例

```
tag=pcap packet ipv4.SrcIP ~ 192.168.0.0/16 tcp.SrcPort | join -s : SrcIP SrcPort as dialstring | unique SrcIP SrcPort | table SrcIP SrcPort dialstring
```
![](join.png)
