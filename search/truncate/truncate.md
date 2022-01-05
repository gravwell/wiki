## truncate 

`truncate`モジュールは、列挙値の最初のN文字（バイナリモードを使用する場合はバイト）のみを保存します。例えば、EV "Message "の最初の20文字以外を切り捨てるには、次のようにします:

```
tag=data json IP Message | truncate -e Message 20 | table
```

`truncate`は、文字列とバイトスライスの列挙値に対してのみ動作し、デフォルトでは、データがUTF-8でエンコードされていることを前提としています。この動作をオーバーライドするには、`-binary` フラグを使用します。

### サポートされているオプション

* `-ellipsis`: オプション。省略記号（ピリオド3文字、"..."）を切り捨て後の文字列に追加します。つまり、切り捨てられたすべての値は、指定されたものより3文字長くなります。
* `-binary`: オプション。データをUTF-8の文字列ではなく、バイトスライスとして扱います。

### 例

この例では、DNSからのリクエストを抽出します:

```
tag=dns json Question.Hdr.Name | table
```

![Example 1](example1.png)

`Truncate`モジュールを使って、省略文字を切り詰めたり、戻したりすることができます: 

```
tag=dns json Question.Hdr.Name 
| truncate -ellipsis Name 10 
| table
```

![Example 2](example2.png)

