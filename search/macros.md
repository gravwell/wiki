# 検索マクロ

検索マクロは、SolitonNKをより効果的に使うための強力な機能です。マクロは、長くて反復的な検索クエリを、覚えやすいショートカットに変えることができます。

## マクロの基礎

マクロは、基本的には検索クエリの文字列に適用される文字列置換ルールです。マクロは、短い名前($MYMACRO のような)を長い文字列にマッピングします。SolitonNKが検索クエリを解析するとき、マクロ名(ドル記号の後に少なくとも1つの大文字または数字が続くように定義されている)を探し、検索を開始する前に置換を行います。マクロを使用すると、GUIにはクエリの拡張バージョンが表示されます。

例えば、`$DHCPACK`というマクロを定義すると、`regex "DHCPACK on (?P<ip>\S+) to (?P<mac>\S+) via (?P<iface>S+)`に展開されます。例えば、`tag=syslog $DHCPACK | unique ip mac | table ip mac`のように、そのマクロをregex呼び出しの代わりに使うことができます。

マクロは、タグ指定、検索モジュール、レンダリングモジュールなど、通常の SolitonNK クエリの任意の部分を含むことができます。マクロはクエリ全体を含むこともできますが、クエリ全体を格納するには検索ライブラリの方が便利なツールです。

### マクロ引数

Macros can be defined with arguments. To define a macro that takes arguments, put replacement directives in the format `%%1%%`, `%%2%%`, etc. in the query string. Those directives will be replaced by the arguments given at run-time. For example, we might define a macro named `HTTPUSER` which expands to `tag=bro-http ax user_agent~"%%1%%"`. Later, we can pass arguments to the macro as though it were a C or Python function:

![](macro-args.png)

### ネストされたマクロ

マクロは別のマクロを含むことができます。マクロ `$FOO` を定義して `tag=foo json timestamp $BAR` に展開することができます; マクロを使用すると、SolitonNK はその展開に別のマクロが含まれていることを認識し、順番に $BAR マクロも展開します。

SokutionNL マクロを数回繰り返して展開し続けますが、マクロのループがある場合は無限再帰を起こします。

## マクロの定義

マクロ管理のページは、SOlitonNK のメインメニューにあります:

![](macro-page.png)

新しいマクロを追加するには、右上の Add ボタンをクリックします。マクロ名とクエリ文字列の入力を求めるウィンドウが表示されます。以下のスクリーンショットは、前述のDHCPACKマクロの定義を示しています:

![](macro-dhcpack.png)
