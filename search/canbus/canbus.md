## Canbus

canbusモジュールはCANメッセージからフィールドを抽出します（すなわち車両データ）。 これらのフィールドは、canbusモジュールの呼び出しで自動的に抽出されます。

| mod | フィールド | オペレータ | 例
|-----|-------|-----------|----------
| canbus | ID | == != < > <= >= | canbus ID==0x341
| canbus | EID | == != < > <= >= | canbus EID==0x123456
| canbus | RTR | == != | canbus RTR==true
| canbus | Data | ~ !~ | canbus Data

### 検索例

次の検索では、canbusパケットIDでカウントし、最も頻度の高いIDを含む表を表示します。

```
tag=vehicles canbus | count by ID | sort by count desc | table ID count
```

次の検索では、スロットルデータを指定しているパケットを抽出し、スロットルの平均位置をプロットします。
```
tag=vehicles canbus ID==0x123 Data | slice uint16be(Data[2:4]) as throttle | mean throttle | chart mean
```
