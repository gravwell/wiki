## anonymize

`anonymize` モジュールは、指定された列挙値の内容を匿名化された値に置き換えるために使用します。例えば、データセットに含まれる IP アドレスを匿名化されたアドレスに置き換えることができます:

```
tag=data json IP Message | anonymize IP | table
```

このクエリは、各エントリーのIP列挙値をランダムなIPアドレスにマッピングします。こうすることで、同じ IP アドレスが再び現れたときに、 同じ匿名化された値を受け取ることができます。

`anonymize` モジュールは、文字列、バイト配列、IP アドレス、MAC アドレス、整数、浮動小数点数、およびロケーションをサポートしています。 

### サポートされているオプション

`anonymize` モジュールはフラグを使用しません。匿名化するために列挙された値を指定するだけでよいのです。

### 例

この例では、DNSからリクエスト/レスポンスペアを抽出します:

```
tag=dns json Question.Hdr.Name~google.com Question.A | require A | table
```

![Example 1](example1.png)

`anonymize` モジュールを追加するだけで、IP アドレスを匿名化することができます:

```
tag=dns json Question.Hdr.Name~google.com Question.A | require A | anonymize IP | table
```

![Example 2](example2.png)

`anonymize` モジュールは、テキストを匿名化することもできます。この例では、ウェブサーバのログからフィールドを抽出し、ユーザエージェントフィールドを匿名化します。匿名化されたテキストは、元の値と同じ長さになります。 

```
tag=apache grok "%{COMBINEDAPACHELOG}" | anonymize agent | table timestamp verb response agent
```

![Example 3](example3.png)
