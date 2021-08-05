## Upper / Lower

upperモジュールとlowerモジュールは、テキストを大文字または小文字に変換します。  単に`upper`（または`lower`）を呼び出すと、エントリの生データが大文字に変換されます。  1つ以上の列挙値の名前のリストを使用してモジュールを呼び出すと、それらの列挙値はすべて大文字または小文字に変換されます。  これは`unique`などのモジュールに渡す前にデータを正規化するのに役立ちます。

### 使用例

この例では、`upper`カウント前にShodanデータを正規化するためにモジュールを使用しています。

```
tag=shodan json location.region_code | upper region_code | count by region_code | table region_code count
```