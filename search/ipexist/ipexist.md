# IPexist

ipexistモジュールは、IPアドレスの簡単な存在チェックを可能な限り高速に実行するように設計されています。SolitonNK	の[ipexistライブラリ](https://github.com/gravwell/ipexist)を使用して、IPアドレスのセットを管理し、そのセット内の指定されたIPの存在を素早く照会します。デフォルトでは、すべての列挙値がセット内のアドレスに一致した場合、そのエントリはパスされます。

## サポートされているオプション

* `-r <resource>`: "-r"フラグは、ipexist形式のルックアップ・セットを含むリソースの名前を指定します。このフラグを複数回指定することで、複数のリソースからの検索を試みることができます。これらのセットの作成に関する詳細は以下を参照してください。
* `-v`: "-v"フラグは、ipexist モジュールに逆のモードで動作するように指示します。したがって、`ipexist -r ips SrcIP` というクエリは、通常、SrcIP がリソースの ip に一致するエントリをすべて通過させますが、`ipexist -v -r ips SrcIP` では、代わりにそれらのエントリを削除し、他のすべてのエントリを通過させます。
* `-or`: "-or"フラグは、すべてのフィルタが成功した場合に、ipexistモジュールがエントリをパイプラインに沿って続行することを許可することを指定します。

## IPセットを作成

ipexistモジュールは、IPv4アドレスのセットを格納するために特定のフォーマットを使用しています。このフォーマットは、比較的スペースを有効に利用しつつ、高速な検索ができるように設計されています。このフォーマットは[ipexistライブラリ](https://github.com/gravwell/ipexist)に実装されており、コマンドラインでセットを生成するツールも含まれています。

まず、ツールを取得します:

	go get github.com/gravwell/ipexist/textinput

次に、セットに入れたいIPアドレスのリストをテキストファイルに入力し、1行に1つのIPを入力します。順序は関係ありません:

	10.0.0.2
	192.168.3.77
	10.3.2.1
	8.8.8.8

次にtextinputツールを実行し、入力ファイルのパスと出力ファイルのパスを指定します:

	$GOPATH/bin/textinput -i /path/to/inputfile -o /path/to/outputfile

これにより、適切にフォーマットされた出力ファイルが作成され、ipexistモジュールで使用するリソースとしてアップロードできるようになります。

## 使用例

パケットが `pcap` タグでキャプチャされているとすると、以下のクエリは、ソースIPアドレスが `ips` リソースのIPに一致するパケットのみを通過させます:

```
tag=pcap packet ipv4.SrcIP | ipexist -r ips SrcIP | table SrcIP
```

![](ipexist1.png)

このクエリは、 **SrcIPとDstIP** がリソースで見つかったすべてのエントリを渡します:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP | ipexist -r ips SrcIP DstIP | table SrcIP DstIP
```

`-or`フラグを追加すると、照会が緩和されます。 **SrcIPまたはDstIP** がリソースに見つかったエントリをすべて渡します:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP | ipexist -or -r ips SrcIP DstIP | table SrcIP DstIP
```

## 反転クエリ

`-v`フラグはクエリを反転させます。クエリに`-v`を追加すると、通常は削除されるエントリがパスされ、その逆も同様です。

このクエリは、ソースIPアドレスがリソースに見つかったエントリを*削除*します:

```
tag=pcap packet ipv4.SrcIP | ipexist -v -r ips SrcIP | table SrcIP
```

![](ipexist2.png)

次のクエリでは、 **SrcIPおよびDstIP** がリソースに存在するエントリはすべて*削除*されます。このクエリは基本的に「送信元または宛先が既知のリストにないすべてのパケットを表示する」というものです。

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP | ipexist -v -r ips SrcIP DstIP | table SrcIP DstIP
```

`-or`フラグと組み合わせると、モジュールは、与えられた列挙値のうち1つでもリソース内で見つかったエントリをすべて削除します。以下の例では、**送信元IPと宛先IP**がリソースに*見つからない*エントリだけがパイプラインを通過します。

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP | ipexist -or -r ips SrcIP DstIP | table SrcIP DstIP
```

### 複数のリソース

`-r` フラグを繰り返すことで、複数の固有IPセットを指定できます。ipexists モジュールは、基本的にこれらをひとつの大きなセットとして扱います。

```
tag=pcap packet ipv4.SrcIP | ipexist -r ips -r externalips SrcIP | table SrcIP
```
