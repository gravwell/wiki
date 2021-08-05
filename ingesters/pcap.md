# ネットワークキャプチャインジェスター

Gravwellの主な強みは、バイナリデータを取り込む機能です。ネットワークインジェスターを使用すると、後で分析するためにネットワークから完全なパケットをキャプチャできます。 これにより、NetFlowやその他の凝縮されたトラフィック情報を単に保存するよりもはるかに優れた柔軟性が得られます。

## 基本構成

Network Capture インジェスターは、[ingester section](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバルコンフィギュレーションブロックを使用します。 他の多くのGravwellインジェスターと同様に、Network Captureインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## スニファーの例

```
[Sniffer "spy1"]
	Interface="p1p1" #sniffing from interface p1p1
	Tag-Name="pcap"  #assigning tag  of pcap
	Snap-Len=0xffff  #maximum capture size
	BPF-Filter="not port 4023" #do not sniff any traffic on our backend connection
	Promisc=true

[Sniffer "spy2"]
	Interface="p5p2"
	Source-Override=10.0.0.1
```

## インストール

GravwellのDebianリポジトリを使用している場合、インストールはaptコマンド1つで完了します。

```
apt-get install libpcap0.8 gravwell-network-capture
```

それ以外の場合は、[Downloads page](#!quickstart/downloads.md)からインストーラーをダウンロードしてください。ネットワークインジェスターをインストールするには、rootとしてインストーラーを実行するだけです（ファイル名が若干異なる場合があります）。

```
root@gravserver ~ # bash gravwell_network_capture_installer.sh
```

注：インジェスターを動作させるには、libpcapがインストールされている必要があります。

可能な限りネットワークインジェスターをインデクサーと同居させ、データの送信には `clear-conn` や `tls-conn` リンクではなく `pipe-conn` リンクを使用することを強くお勧めします。 ネットワーク・インジェスターがエントリーをプッシュするのに使用しているのと同じリンクからキャプチャーしている場合、リンクを急速に飽和させるフィードバック・ループが発生する可能性があります（例えば、eth0からキャプチャーしながら、eth0経由でインジェスターにエントリーを送信している場合など）。BPF-Filter`オプションを使用すると、この問題を軽減することができます。

Gravwellバックエンドがインストールされているマシンにインジェスターがある場合、インストーラは自動的に正しい`Ingest-Secrets`の値を拾い、設定ファイルに入力します。そうでない場合は、インデクサのIPアドレスとインジェスト・シークレットの入力を求められます。いずれにしても、実行する前に `/opt/gravwell/etc/network_capture.conf` にある設定ファイルを確認してください。eth0からのトラフィックをキャプチャする例は次のようになります。

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify = false
#Cleartext-Backend-target=127.1.0.1:4023 #example of adding a cleartext connection
#Encrypted-Backend-target=127.1.1.1:4023 #example of adding an encrypted connection
Pipe-Backend-target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when ingester is on the same machine as a backend
Log-Level=INFO #options are OFF INFO WARN ERROR
Ingest-Cache-Path=/opt/gravwell/cache/network_capture.cache

#basic default logger, all entries will go to the default tag
#no Tag-Name means use the default tag
[Sniffer "spy1"]
	Interface="eth0" #sniffing from interface eth0
	Tag-Name="pcap"  #assigning tag  of pcap
	Snap-Len=0xffff  #maximum capture size
	#BPF-Filter="not port 4023" #do not sniff any traffic on our backend connection
	#Promisc=true
```

異なるインターフェイスからキャプチャするために、任意の数の`Sniffer`エントリを構成することができます。

ディスク容量が気になる場合は、パケットのメタデータだけをキャプチャするために、`Snap-Len`パラメータを変更するとよいでしょう。通常、ヘッダのみをキャプチャするには96の値で十分です。

非常に高い帯域幅のリンクの可能性があるため、ネットワークキャプチャーデータを独自のウェルに割り当てることもお勧めします。これには、パケットキャプチャータグ用に別のウェルを定義するために、インデクサでの設定が必要です。

NetworkCaptureのインゲスターは、libpcapのシンタックスに準拠した`BPF-Filter`パラメータを使ったネイティブなBPFフィルタリングもサポートしています。 ポート22のすべてのトラフィックを無視するには、次のようにスニファーを設定します。

```
[Sniffer "no-ssh"]
	Interface="eth0"
	Tag-Name="pcap"
	Snap-Len=0xffff
	BPF-Filter="not port 22"
```

インジェスターがインデクサーとは別のシステムにあり、エントリーがインジェストされるためにネットワークを通過しなければならない場合は、`BPF-Filter`を「not port 4023」（クリアテキストを使用する場合）または「not port 4024」（TLSを使用する場合）に設定する必要があります。

## ネットワーク検索の例

次の検索は、RSTフラグが設定されたTCPパケットのうち、IP 10.0.0.0/24クラスCサブネットから発信されていないものを探し、IPごとにグラフ化します。 この検索は、ネットワークからのアウトバウンドポートスキャンを迅速に特定するために使用できます。

```
tag=pcap packet tcp.RST==TRUE ipv4.SrcIP !~ 10.0.0.0/24 | count by SrcIP | chart count by SrcIP limit 10
```

![](portscan.png)

以下の検索では、IPv6トラフィックを検索してFlowLabelを抽出し、これを数学演算に渡しています。これにより、パケットの長さを合計してチャート・レンダラに渡すことで、フローごとのトラフィック・アカウンティングが可能になります。

```
tag=pcap packet ipv6.Length ipv6.FlowLabel | sum Length by FlowLabel | chart sum by FlowLabel limit 10
```

TCPペイロードで使用されている言語を識別するために、ネットワークデータをフィルタリングして、langfindモジュールに渡します。 このクエリは、アウトバウンドのHTTPリクエストを探し、TCPペイロードデータをlangfindモジュールに渡し、識別された言語をカウント、そしてチャートに渡します。 これにより、アウトバウンドHTTPクエリで使用される人間の言語のチャートが作成されます。

```
tag=pcap packet ipv4.DstIP != 10.0.0.100 tcp.DstPort == 80 tcp.Payload | langfind -e Payload | count by lang | chart count by lang
```

![](langfind.png)

トラフィックアカウンティングは、レイヤー2でも行うことができます。これは、イーサネットヘッダーからパケットの長さを抽出し、その長さを宛先MACアドレスで合計し、トラフィックカウントでソートすることで実現されます。 これにより、特におしゃべりをしている可能性のあるイーサネットネットワーク上の物理デバイスを迅速に特定することができます。

```
tag=pcap packet eth.DstMAC eth.Length > 0 | sum Length by DstMAC | sort by sum desc | table DstMAC sum
```

同様の方法で、パケット数からおしゃべりな機器を特定することができます。例えば、あるデバイスが小さなイーサネット・パケットを積極的にブロードキャストしていて、スイッチに負担をかけているが、大量のトラフィックにはなっていない場合があります。

```
tag=pcap packet eth.DstMAC eth.Length > 0 | count by DstMAC | sort by count desc | table DstMAC count
```

非標準のHTTPポートで動作するHTTPトラフィックを識別することが望ましい場合があります。 このような場合は、フィルタリングオプションを使用して、ペイロードを他のモジュールに渡すことで実現できます。 例えば、特定のサブネットから発信されているTCPポート80以外のアウトバウンドトラフィックを探し、パケットのペイロードに含まれるHTTPリクエストを探すことで、異常なHTTPトラフィックを特定できます。

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP tcp.DstPort !=80 ipv4.SrcIP ~ 10.0.0.0/24 tcp.Payload | regex -e Payload "(?P<method>[A-Z]+)\s+(?P<url>[^\s]+)\s+HTTP/\d.\d" | table method url SrcIP DstIP DstPort
```

![](nonstandardhttp.png)
