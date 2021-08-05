# パケットフリートインジェスター

パケットフリートインジェスター、Google Stenographerのインスタンスに問い合わせを行い、その結果をパケット単位でGravwellに取り込む仕組みを提供します。

各Stenographerインジェスターは、指定されたポート（``Listen-Address``）でリッスンし、HTTP POSTでStenographerのクエリ（以下のクエリ構文を参照）を受け付けます。クエリを受信すると、インゲスターは整数のジョブIDを返し、非同期的にステノグラファー・インスタンスにクエリを行い、返されたPCAPの取り込みを開始します。複数のインフライト・クエリを同時に実行することができます。ジョブのステータスは、"/status "に対してHTTP GETを発行することで見ることができ、インフライト・ジョブIDのJSONエンコードされた配列が返されます。

また、指定されたインゲスターポートにアクセスすることで、ジョブの送信やステータスの確認ができるシンプルなWebインターフェースも利用できます。

## 基本設定

パケットフリートは、[ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters)に記載されている統一されたグローバルコンフィギュレーションブロックを使用しています。 他の多くのGravwellインゲスターと同様に、パケットフリートは複数のアップストリームインデクサー、TLS接続、クリアテキスト接続、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## ステノグラファーの例

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	Assume-Local-Timezone=false #Default for assume localtime is false
	Source-Override="DEAD::BEEF" #override the source for just this Queue 

[Stenographer "Region 2"]
	URL="https://my.url:1234"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
```

## 設定オプション 

パケットフリートでは、いくつかのグローバル設定とステノグラファーごとの設定オプションが必要です。グローバル設定では、TLSの設定（該当する場合）や、Webインターフェースのリスンアドレスの設定などがあります。

```
Use-TLS=true
Listen-Address=":9002"
Server-Cert="server.cert"
Server-Key="server.key"
```

ステノグラファーの各インスタンスには、以下のスタンザが必要です。ここでの例題名「Region 1」は、ウェブインターフェイスがステノグラファーインスタンスをリストアップする際に使用されます。

```
[Stenographer "Region 1"]
	URL="https://127.0.0.1:9001"
	CA-Cert="ca_cert.pem"
	Client-Cert="client_cert.pem"
	Client-Key="client_key.pem"
	Tag-Name=steno
	#Assume-Local-Timezone=false #Default for assume localtime is false
	#Source-Override="DEAD::BEEF" #override the source for just this Queue 
```

## 検索言語 

ユーザーは、非常にシンプルな検索言語でパケットを指定して、速記者にパケットを要求します。検索言語で指定してステノグラファーにパケットを要求します。 この言語はBPFの単純なサブセットであり、以下のプリミティブを含んでいます。

    host 8.8.8.8          # Single IP address (hostnames not allowed)
    net 1.0.0.0/8         # Network with CIDR
    net 1.0.0.0 mask 255.255.255.0  # Network with mask
    port 80               # Port number (UDP or TCP)
    ip proto 6            # IP protocol number 6
    icmp                  # equivalent to 'ip proto 1'
    tcp                   # equivalent to 'ip proto 6'
    udp                   # equivalent to 'ip proto 17'

    # Stenographer-specific time additions:
    before 2012-11-03T11:05:00Z      # Packets before a specific time (UTC)
    after 2012-11-03T11:05:00-07:00  # Packets after a specific time (with TZ)
    before 45m ago        # Packets before a relative time
    before 3h ago         # Packets after a relative time

**注釈**:相対的な時間は、上記のように時間または分の整数値で測定する必要があります。


プリミティブは、and/&&やor/||と組み合わせることができます。
これらは同じ優先順位を持ち、左から右へと評価されます。 パレンはグループ化にも使用できます。

    (udp and port 514) or (tcp and port 8080)

**注**:この部分は[Google Stenographer](https://github.com/google/stenographer/blob/master/README.md)からの引用です。
