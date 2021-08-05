# IPMIインジェスター

IPMIインジェスターは、任意の数のIPMIデバイスからセンサーデータレコード（SDR）とシステムイベントログ（SEL）のレコードを収集します。

設定ファイルには、各IPMIデバイスに接続するためのシンプルなホスト/ポート、ユーザー名、パスワードのフィールドが用意されています。SELとSDRのレコードは、JSONにエンコードされたスキーマで取り込まれます。例えば、以下のようになります:

```
{
    "Type": "SDR",
    "Target": "10.10.10.10:623",
    "Data": {
        "+3.3VSB": {
            "Type": "Voltage",
            "Reading": "3.26",
            "Units": "Volts",
            "Status": "ok"
        },
        "+5VSB": {...},
        "12V": {...}
    }
}

{
    "Target": "10.10.10.10:623",
    "Type": "SEL",
    "Data": {
        "RecordID": 25,
        "RecordType": 2,
        "Timestamp": {
            "Value": 1506550240
        },
        "GeneratorID": 32,
        "EvMRev": 4,
        "SensorType": 5,
        "SensorNumber": 81,
        "EventType": 111,
        "EventDir": 0,
        "EventData1": 240,
        "EventData2": 255,
        "EventData3": 255
    }
}
```

## 基本設定

IPMIインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、IPMIインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## 設定オプション

IPMIでは、グローバル設定オプションのデフォルトセットを使用します。IPMIデバイスは「IPMI」スタンザで構成され、各スタンザは同じ認証情報を共有する複数のIPMIデバイスをサポートできます。例えば、以下のようになります:

```
[IPMI "Server 1"]
	Target="127.0.0.1:623"
	Target="1.2.3.4:623"
	Username="user"
	Password="pass"
	Tag-Name=ipmi
	Rate=60
	Source-Override="DEAD::BEEF" 
```

IPMIスタンザはシンプルで、1つまたは複数のターゲット（IPMIデバイスのIP:PORT）、ユーザー名、パスワード、タグ、ポールレート（秒）を指定するだけです。デフォルトのポーリングレートは60秒です。オプションで、ソースオーバーライドを設定して、取り込まれたすべてのエントリーのSRCフィールドを別のIPに強制的に割り当てることができます。既定では、SRCフィールドはIPMIデバイスのIPに設定されます。

さらに、すべてのIPMIスタンザは、[ここ](https://docs.gravwell.io/#!ingesters/preprocessors/preprocessors.md)で説明されているように、「プリプロセッサ」オプションを使用することができます。
