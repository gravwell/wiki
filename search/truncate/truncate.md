## truncate 

`truncate`モジュールは、列挙された値の最初のN文字（バイナリモードを使用する場合はバイト）のみを保存します。例えば、EV "Message"の最初の20文字以外を切り捨てるには、次のようにします：

```
tag=data json IP Message | truncate -e Message 20 | table
```

`truncate`は、文字列とバイトスライスの列挙された値に対してのみ動作し、デフォルトではデータがUTF-8でエンコードされていると仮定します。この動作をオーバーライドするには、`-binary`フラグを使用します。

### 対応オプション

* `-ellipsis`: オプションです。つまり、すべての切り捨てられた値は、指定されたものよりも3文字長くなります。
* `-binary`: オプションです。オプションです。データを UTF-8 文字列ではなく、バイトスライスとして扱います。

### 例

この例では、DNSからのリクエストを抽出します：

```
tag=dns json Question.Hdr.Name | table
```

![例 1](example1.png)

また、`truncate`モジュールを使って、省略文字を切り捨てたり、戻したりすることができます：

```
tag=dns json Question.Hdr.Name 
| truncate -ellipsis Name 10 
| table
```

![例 2](example2.png)

