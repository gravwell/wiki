# Gravwell Accelerators

Gravwellは、フィールド抽出を実行するために、エントリが*取り込まれた*ときにエントリを処理できます。抽出されたフィールドは、各シャードに付随するアクセラレーションブロックに記録されます。アクセラレータを使用すると、最小限のストレージオーバーヘッドでスループットを大幅に高速化できます。加速器はウェルごとに指定され、できるだけ目立たず、柔軟になるように設計されています。データがアクセラレーションディレクティブと一致しないウェルに入力された場合、または指定されたフィールドが欠落している場合、Gravwellは他のエントリと同じようにデータを処理します。可能な場合は加速が行われます。

2つの理由から、「インデクサー」と「インデックス作成」ではなく、「アクセラレーター」と「アクセラレーション」を参照します。まず、Gravwellにはすでに「インデクサー」と呼ばれる非常に重要なコンポーネントがあります。第2に、加速はブルームフィルターを使用して**または**直接インデックスを作成することで実行できるため、「インデックス」について説明することは必ずしも正確ではありません。

## 加速の基本

Gravwellアクセラレータは、データが比較的ユニークな場合に最も効果的なフィルタリング技術を使用しています。 フィールド値が非常に一般的であるか、ほとんどすべてのエントリに存在する場合、アクセラレータ仕様に含めることはあまり意味がありません。 また、複数のフィールドを指定してフィルタ処理すると、精度が向上し、クエリーの処理速度が向上します。 高速化のための候補となるフィールドは、ユーザが直接問い合わせるフィールドである。 例えば、プロセス名、ユーザ名、IPアドレス、モジュール名、あるいはスタック内の型の問い合わせに使用される他のフィールドなどがあります。

タグは、使用中の抽出モジュールに関係なく、常にアクセラレーションに含まれます。クエリがインラインフィルタを指定していない場合でも、1つのウェルに複数のタグがある場合には、アクセラレーションシステムがクエリを絞り込み、加速化するのに役立ちます。

ほとんどのアクセラレーションモジュールは、ブルームエンジンを使用した場合、約1～1.5%のストレージオーバーヘッドが発生しますが、非常にスループットの低いウェルでは、より多くのストレージを消費する可能性があります。一般的に1秒間に約1～10件のエントリーがあるウェルの場合、アクセラレーションには5～10％のストレージペナルティが発生しますが、1秒間に10～15,000件のエントリーがあるウェルでは0.5％のストレージオーバーヘッドが発生することもあります。グラブウェルアクセラレータでは、ユーザーが指定した衝突速度の調整も可能です。ストレージに余裕がある場合は、コリジョンレートを低くすることで精度が向上し、クエリが高速化される一方で、ストレージのオーバーヘッドが増加することがあります。精度を下げるとストレージのペナルティは減りますが、精度は低下し、アクセラレータの有効性は低下します。インデックスエンジンは、抽出されるフィールドの数や抽出されるデータの変動性に応じて、大幅に多くのスペースを消費することになります。例えば、フルテキストのインデックスを作成すると、アクセラレータのファイルが保存されているデータファイルと同じくらいの容量を消費することがあります。

## 加速エンジン

エンジンとは、抽出した加速度データを実際に格納するシステムのことです。Gravwellは2つのアクセラレーションエンジンをサポートしています。各エンジンは、希望するインジェスト レート、ディスク オーバーヘッド、検索パフォーマンス、およびデータ量に応じて異なる利点を提供します。加速エンジンは、アクセラレータ抽出器自体から完全に独立しています（アクセラレータ名設定オプションで指定されます）。

デフォルトのエンジンは「index」エンジンです。「index」エンジンは、すべてのクエリタイプで高速に動作するように設計された完全なインデックス作成システムです。インデックスエンジンは通常、bloomエンジンよりもかなり多くのディスク容量を消費しますが、非常に大きなデータ量や、総データのかなりの部分に触れる可能性のあるクエリで動作する場合には、非常に高速です。インデックスエンジンは、重度のインデックスを持つシステムでは圧縮データと同じくらいのスペースを消費することも珍しくありません。

bloomエンジンは、bloomフィルタを使用して、データの一部が所定のブロックに存在するかどうかを示します。bloom エンジンは通常、ディスクのオーバーヘッドがほとんどなく、特定のIPが表示されているログを見つけるなど、針の筵のようなスタイルのクエリでうまく動作します。bloomエンジンは、フィルタリングされたエントリが定期的に発生するようなフィルタリングではパフォーマンスが低い。また、bloomエンジンは、フルテキストアクセラレータと組み合わせた場合には、あまり良い選択ではありません。

### インデックスエンジンの最適化

「インデックス」は、ファイルに裏打ちされたデータ構造を使用して、キーデータを格納およびクエリします。ファイルのバッキングはメモリマップを使用して実行されます。これは、カーネルがダーティページを書き戻すことに熱心すぎる場合に非常に悪用される可能性があります。カーネルのダーティページパラメータを調整して、カーネルがダーティページを書き戻す頻度を減らすことを強くお勧めします。これは「/proc」インターフェースを介して行われ、「/etc/sysctl.conf」構成ファイルを使用して永続的にすることができます。次のスクリプトは、いくつかの効率的なパラメーターを設定し、それらが再起動後も維持されるようにします。

```
#!/bin/bash
user=$(whoami)
if [ "$user" != "root" ]; then
	echo "must run as root"
fi

echo 70 > /proc/sys/vm/dirty_ratio
echo 60 > /proc/sys/vm/dirty_background_ratio
echo 2000 > /proc/sys/vm/dirty_writeback_centisecs
echo 3000 > /proc/sys/vm/dirty_expire_centisecs

echo "vm.dirty_ratio = 70" >> /etc/sysctl.conf
echo "vm.dirty_background_ratio = 60" >> /etc/sysctl.conf
echo "vm.dirty_writeback_centisecs = 2000" >> /etc/sysctl.conf
echo "vm.dirty_expire_centisecs = 3000" >> /etc/sysctl.conf

```

## アクセラレーションの構成

加速器はウェルごとに構成されます。各ウェルは、加速モジュール、抽出用のフィールド、衝突率、および入力ソースフィールドを含めるオプションを指定できます。通常、特定のソースでフィルタリングする場合(たとえば、特定のデバイスからのsyslogエントリのみを確認する場合)、ソースフィールドを含めると、抽出されるフィールドに関係なく、アクセラレータの精度を高める効果的な方法が提供されます。

| 加速パラメータ| 説明| 例|
| ---------- | ------ | ------------- |
| アクセラレータ-名前| 取り込み時に使用するフィールド抽出モジュールを指定します| アクセラレータ-Name = "json" |
| アクセラレータ-引数| アクセラレーションモジュールの引数、通常は抽出するフィールドを指定します| アクセラレータ-Args = "username hostname appname" |
| 衝突率| ブルームエンジンを使用して加速モジュールの精度を制御します。 0.1から0.000001の間でなければなりません。 デフォルトは0.001です。 | 衝突率= 0.01
| Accelerate-On-Source | 各モジュールのSRCフィールドを含める必要があることを指定します。 これにより、CEFなどのモジュールをSRCと組み合わせることができます。 | Accelerate-On-Source = true
| アクセラレータ-エンジン-オーバーライド| インデックス作成に使用するエンジンを指定します。 デフォルトでは、インデックスエンジンが使用されます。 | Accelerator-Engine-Override = index

### サポートされている抽出モジュール

* [CSV](#!search/csv/csv.md)
* [Fields](#!search/fields/fields.md)
* [Syslog](#!search/syslog/syslog.md)
* [JSON](#!search/json/json.md)
* [CEF](#!search/cef/cef.md)
* [Regex](#!search/regex/regex.md)
* [Winlog](#!search/winlog/winlog.md)
* [Slice](#!search/slice/slice.md)
* [Netflow](#!search/netflow/netflow.md)
* [IPFIX](#!search/ipfix/ipfix.md)
* [Packet](#!search/packet/packet.md)
* Fulltext

### 構成例

以下は、タブ区切りのエントリの2番目、4番目、および5番目のフィールド(たとえば、broログファイルからの行)を抽出する構成の例です。この例では、各bro接続ログから送信元IP、宛先IP、および宛先ポートを抽出して加速しています。「bro」ウェルに入るすべてのエントリ（この例では「bro」タグのみを含む）は、取り込み中に抽出モジュールを通過します。データがアクセラレーション仕様に準拠していない場合、データは保存されますが、アクセラレーションは行われません。クエリに含まれますが、多くの不適合エントリがウェルにある場合、クエリははるかに遅くなります。

```
[Storage-Well "bro"]
	Location=/opt/gravwell/storage/bro
	Tags=bro
	Accelerator-Name="fields"
	Accelerator-Args="-d \"\t\" [2] [4] [5]"
	Accelerate-On-Source=true
	Collision-Rate=0.0001
```

## 加速の基本

各アクセラレーションモジュールは、基本的なフィールド抽出のために、コンパニオンサーチモジュールと同じ構文を使用します。アクセラレータは、名前の変更、フィルタリング、または列挙値の操作をサポートしていません。それらは第1レベルのフィルターです。アクセラレーションモジュールは、対応する検索モジュールが動作し、等式フィルターを実行するたびに透過的に呼び出されます。

たとえば、JSONアクセラレータを使用する次のウェル構成について考えます。

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname app.field1 app.field2"
```

次のクエリを発行する場合：

```
tag=app json username==admin app.field1=="login event" app.field2 != "failure" | count by hostname | table hostname count
```

json検索モジュールは、アクセラレーションフレームワークを透過的に呼び出し、「username」と「app.field1」で抽出された値に第1レベルのフィルターを提供します。「app.field2」フィールドは、直接等式フィルターを使用しないため、このクエリでは高速化されません。サブセットを除外、比較、またはチェックするフィルターは、アクセラレーションの対象にはなりません。

## 全文

フルテキストアクセラレータは、テキストログ内の単語にインデックスを付けるように設計されており、最も柔軟なアクセラレーションオプションと見なされています。他の検索モジュールの多くは、クエリの実行時に全文アクセラレータの呼び出しをサポートしています。ただし、全文アクセラレータを使用するための主要な検索モジュールは、`-w`フラグが設定された[grep](/search/grep/grep.md)モジュールです。UNIXのgrepユーティリティと同様に、 `grep -w`は、提供されたフィルタがバイトのサブセットではなく単語に期待されることを指定します。`words foo bar baz`で検索を実行すると、単語foo、bar、bazが検索され、全文アクセラレータが使用されます。

フルテキストアクセラレータは最も柔軟性がありますが、最もコストがかかります。フルテキストアクセラレータを使用すると、Gravwellの取り込みパフォーマンスが大幅に低下し、大量のストレージスペースを消費する可能性があります。これは、フルテキストアクセラレータがすべてのエントリのほぼすべてのコンポーネントでインデックスを作成しているためです。

### フルテキスト引数

フルテキストアクセラレータは、インデックスが作成されるデータの種類を絞り込み、ストレージのオーバーヘッドが大きくなるがクエリ時にはあまり役に立たないフィールドを削除できるいくつかのオプションをサポートしています。

| 引数 | 説明 | 例 | デフォルト状態 |
|----------|-------------|---------|---------------|
| -acceptTS |デフォルトでは、フルテキストアクセラレータは、データ内のタイムスタンプを識別して無視しようとします。このフラグはその動作を無効にし、タイムスタンプにインデックスを付けることを許可します。 | `-acceptTS` |無効|
| -acceptFloat |デフォルトでは、フルテキストアクセラレータは浮動小数点数を識別して無視しようとします。これは、浮動小数点数は通常大きく異なり、明示的に照会されないためです。このフラグを設定すると、その動作が無効になり、浮動小数点数にインデックスを付けることができます。 | `-acceptFloat` |無効|
| -min |抽出されたトークンは少なくともXバイトの長さである必要があります。これにより、「is」や「I」などの非常に小さな単語のインデックス作成を防ぐことができます。 | `-min 3` |無効|
| -max |抽出されたトークンの長さがXバイト未満である必要があります。これにより、実行可能なクエリが実行されないログ内の非常に大きな「blob」のインデックス作成を防ぐことができます。 | `-最大256` |無効|
| -ignoreUUID |フィルタを有効にして、UUID / GUID値を無視します。一部のログでは、エントリごとにUUIDが生成されます。これにより、インデックス作成のオーバーヘッドが大幅に増加し、価値がほとんど提供されません。 | `-ignoreUUID` |無効|
| -ignoreTS |加速中のタイムスタンプを特定して無視します。タイムスタンプは頻繁に変更されるため、肥大化の重大な原因となる可能性があります。フルテキストアクセラレータはデフォルトでタイムスタンプを無視します| `-ignoreTS` |有効|
| -ignoreFloat |浮動小数点数は無視してください。フィルタに浮動小数点数が使用されているログでは、 `-accptFloat`を使用できます。 | `-acceptFloat` |有効|
| -maxInt |特定のサイズ未満の整数のみにインデックスを付けるフィルターを有効にします。これは、HTTPアクセスログなどのデータにインデックスを付けるときに役立ちます。戻りコードにインデックスを付けたいが、応答時間とデータサイズにはインデックスを付けたくない。 | `-maxInt 1000` |無効|

注：`-acceptTS`フラグと` -acceptFloat`フラグを有効にする前に、データを理解していることを確認してください。これらは、インデックスエンジンを使用すると、インデックスを大幅に肥大化させる可能性があります。 Bloomエンジンは、タイムスタンプや浮動小数点数などの直交データによる影響が少なくなります。

### ウェル構成の例

次の適切な構成は、`index`エンジンを使用して全文アクセラレーションを実行します。タイムスタンプ、UUIDを識別して無視しようとしていますが、すべてのトークンの長さが2バイト以上である必要があります。

```
[Default-Well]
	Location=/opt/gravwell/storage/default
	Accelerator-Name="fulltext"
	Accelerator-Args="-ignoreTS -ignoreUUID -min 2"
```

## JSON

JSONアクセラレータモジュールは、アクセラレータ名「json」を使用して指定され、JSONモジュールとまったく同じ構文を使用してフィールドを選択します。フィールド抽出の詳細については、[JSON検索モジュール](#!search/json/json.md)セクションを参照してください。

### ウェル構成の例

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="json"
	Accelerator-Args="username hostname \"strange-field.with.specials\".subfield"
```

## Syslog

Syslogアクセラレータは、準拠したRFC5424Syslogメッセージで動作するように設計されています。フィールド抽出の詳細については、[syslog検索モジュール](#!search/syslog/syslog.md)セクションを参照してください。

### ウェル構成の例

```
[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Tags=syslog
	Accelerator-Name="syslog"
	Accelerator-Args="Hostname Appname MsgID valueA valueB"
```

## CEF

CEFアクセラレータは、CEFログメッセージを操作するように設計されており、検索モジュールと同じように柔軟性があります。フィールド抽出の詳細については、[CEF検索モジュール](#!search/cef/cef.md) セクションを参照してください。

### ウェル構成の例

```
[Storage-Well "ceflogs"]
	Location=/opt/gravwell/storage/cef
	Tags=app1
	Accelerator-Name="cef"
	Accelerator-Args="DeviceVendor DeviceProduct Version Ext.Version"
```

## フィールド

フィールドアクセラレータは、CSV、TSV、またはその他の区切り文字であるかどうかに関係なく、任意の区切りデータ形式で動作できます。フィールドアクセラレータを使用すると、検索モジュールと同じ方法で区切り文字を指定できます。フィールド抽出の詳細については、[フィールド検索モジュール](#!search/fields/fields.md)セクションを参照してください。

###　ウェル構成の例

この構成では、コンマ区切りのエントリから4つのフィールドが抽出されます。区切り文字を指定するために`-d`フラグを使用していることに注意してください。

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
```

## CSV

CSVアクセラレータは、カンマ区切りの値データを操作するように設計されており、データから周囲の空白と二重引用符を自動的に削除します。列抽出の詳細については、[CSV検索モジュール](#!search/csv/csv.md)セクションを参照してください。

### ウェル構成の例

```
[Storage-Well "security"]
	Location=/opt/gravwell/storage/seclogs
	Tags=secapp
	Accelerator-Name="csv"
	Accelerator-Args="[1] [2] [5] [3]"
```

## regex

regexアクセラレータを使用すると、非標準のデータ形式を処理するために、取り込み時に複雑な抽出を行うことができます。regexは抽出速度が遅い形式の1つであるため、特定のフィールドで高速化すると、クエリのパフォーマンスが大幅に向上します。

### ウェル構成の例

```
[Storage-Well "webapp"]
	Location=/opt/gravwell/storage/webapp
	Tags=webapp
	Accelerator-Name="regex"
	Accelerator-Args="^\\S+\\s\\[(?P<app>\\w+)\\]\\s<(?P<uuid>[\\dabcdef\\-]+)>\\s(?P<src>\\S+)\\s(?P<srcport>\\d+)\\s(?P<dst>\\S+)\\s(?P<dstport>\\d+)\\s(?P<path>\\S+)\\s"
```

重要：gravwell.confファイルで正規表現を指定するときは、円記号 '\\'をエスケープすることを忘れないでください。 正規表現引数 '\\ w'は '\\\\ w'になります

## Winlog

winlogモジュールは、おそらく最も遅い検索モジュールです。Windowsログスキーマと組み合わされたXMLデータの複雑さは、モジュールが非常に冗長である必要があることを意味し、パフォーマンスがかなり低下します。これは、winlogモジュールを使用した数百万または数十億の高速化されていないエントリの処理が非常に遅くなるため、Windowsログデータの高速化が最も重要なパフォーマンスの最適化である可能性があることを意味します。アクセラレータは、すべてのデータでwinlog検索モジュールを呼び出すことなく、必要な特定のログエントリを絞り込むのに役立ちます。ただし、winlogデータを高速化すると、処理が検索時間から取り込み時間にシフトするだけです。つまり、高速化を有効にすると、Windowsログの取り込みが遅くなるため、Gravwellの通常の取り込み速度である1秒あたり数十万エントリの取り込みを期待しないでください。winlog-よく加速されました。

### ウェル構成の例

```
[Storage-Well "windows"]
	Location=/opt/gravwell/storage/windows
	Tags=windows
	Accelerator-Name="winlog"
	Accelerator-Args="EventID Provider Computer TargetUserName SubjectUserName"
```

重要：winlogアクセラレーターは許容的です（「-or」フラグが暗示されます）。したがって、2つのフィールドが同じエントリに存在しない場合でも、検索のフィルタリングに使用する予定のフィールドを指定します。

## Netflow

[netflow](#!search/netflow/netflow.md)モジュールを使用すると、netflow V5フィールドを高速化し、大量のnetflowデータに対するクエリを高速化できます。NetFlowモジュールは非常に高速で、データは非常にコンパクトですが、NetFlowデータの量が非常に多い場合は、アクセラレーションを使用すると効果的です。NetFlowモジュールは、任意の直接NetFlowフィールドを使用できますが、ピボットヘルパーフィールドは使用できません。これは、 `IP`ではなく`Src`または`Dst`を指定する必要があることを意味します。`IP`および`Port`フィールドはacceleration引数で指定できません。

注：ヘルパー抽出の`Timestamp`と` Duration`はアクセラレーターでは使用できません。

### ウェル構成の例

この設定例では、「bloom」エンジンを使用しており、プロトコルだけでなく、送信元と宛先のIP /ポートペアで高速化しています。

```
[Storage-Well "netflow"]
	Location=/opt/gravwell/storage/netflow
	Tags=netflow
	Accelerator-Name="netflow"
	Accelerator-Args="Src Dst SrcPort DstPort Protocol"
	Accelerator-Engine-Override="bloom"
```

## IPFIX

[ipfix](#!search/ipfix/ipfix.md)モジュールは、IPFIX形式のレコードに対するクエリを高速化できます。このモジュールは「通常の」IPFIXフィールドのいずれかで加速できますが、ヘルパーフィールドをピボットすることはできません。つまり、`port`ではなく`sourceTransportPort`または `destinationTransportPort`を指定するか、`ip`ではなく`src`/`dst`を指定する必要があります。

### ウェル構成の例

この設定例では、`index`エンジンを使用して、送信元/宛先のIP/ポートペアとフローのIPプロトコルを高速化します。これは、netflowセクションに示されている例に相当します。

```
[Storage-Well "ipfix"]
	Location=/opt/gravwell/storage/ipfix
	Tags=ipfix
	Accelerator-Name="ipfix"
	Accelerator-Args="src dst sourceTransportPort destinationTransportPort protocolIdentifier"
	Accelerator-Engine-Override=index
```

## Packet

[packet](#!search/packet/packet.md) モジュールは、同じ名前の検索モジュールと同じ構文を使用して、生のパケットキャプチャを高速化できます。検索モジュールと比較して、パケットアクセラレータの適用方法には微妙ですが重要な違いがあります。アクセラレータは重複するレイヤーを使用できます。 これは、UDPアイテムとTCPアイテムの両方を指定し、処理中のパケットに応じて適切なフィールドを抽出できることを意味します。

ウェル構成は、IPv4、IPv6、TCP、UDP、ICMPなどをすべて同時に高速化するように構成できます。パケットアクセラレータは、指定されたフィールドを暗黙のフィルタとして扱いません。

パケットアクセラレータには直接フィールドも必要です。つまり、`IP`や` Port`などの便利なエクストラクタは使用できません。 加速したいものを正確に指定する必要があります。

### ウェル構成の例

```
[Storage-Well "packets"]
	Location=/opt/gravwell/storage/pcap
	Tags=pcap
	Accelerator-Name="packet"
	Accelerator-Args="ipv4.SrcIP ipv4.DstIP ipv6.SrcIP ipv6.DstIP tcp.SrcPort tcp.DstPort udp.SrcPort udp.DstPort"
```

## SRC

srcアクセラレータは、エントリのソースフィールドのみを高速化する必要がある場合に使用できます。ただし、「Accelerate-On-Source」フラグを有効にし、クエリでsrc検索モジュールを使用することで、srcアクセラレータを他のアクセラレータと組み合わせることが基本的に可能です。 ィルタリングの詳細については、[src検索モジュール](#!search/src/src.md)を参照してください。

### ウェル構成の例

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="src"
```

### SRCを組み合わせたウェル構成とクエリの例

```
[Storage-Well "applogs"]
	Location=/opt/gravwell/storage/app
	Tags=app
	Accelerator-Name="fields"
	Accelerator-Args="-d \",\" [1] [2] [5] [3]"
	Accelerate-On-Source=true
```

次のクエリは、フィールドアクセラレータとsrcアクセラレータの両方を呼び出して、特定のソースからの特定のログタイプを指定します。

```
tag=app src dead::beef | fields -d "," [1]=="security" [2]="process" [5]="domain" [3] as processname | count by processname | table processname count
```

## 加速性能とベンチマーク

アクセラレーションのメリットとデメリットを理解するには、アクセラレーションがストレージの使用、取り込みのパフォーマンス、クエリのパフォーマンスにどのように影響するかを確認するのが最善です。[github](https://github.com/kiritbasu/Fake-Apache-Log-Generator)で利用可能なジェネレーターを使用して生成されたいくつかのApache結合アクセスログを使用します。私たちのデータセットは、約24時間にわたって分散された1,000万のappache結合アクセスログです。合計データは2.1GBです。4つの異なる構成で4つのウェルを定義します。返されるバイト数など、インデックスを作成する意味があまりないパラメータが多数あるため、このデータのインデックス作成にはかなり単純なアプローチを採用します。


| ウェル| 抽出器| エンジン| 説明|
| ------ | ----------- | -------- | ------------- |
| raw| なし| なし| 加速のない完全に生のウェル|
| 全文| 全文| インデックス| インデックスエンジンを使用し、すべての単語に対してフルテキストアクセラレーションを実行するフルテキストアクセラレーションウェル|
| regexindex | 正規表現| インデックス| 正規表現抽出機能とインデックスエンジンを使用すると、十分に加速されます。 各パラメーターが抽出され、インデックスが付けられます|
| regexbloom | 正規表現| bloom| regexindexウェルと同じエクストラクターを備えていますが、bloomエンジンを備えたウェル。 各パラメーターが抽出され、bloomフィルターに追加されます|

ウェルの構成は次のとおりです:

```
[Storage-Well "raw"]
	Location=/opt/gravwell/storage/raw
	Tags=raw
	Enable-Transparent-Compression=true

[Storage-Well "fulltext"]
	Location=/opt/gravwell/storage/fulltext
	Tags=fulltext
	Enable-Transparent-Compression=true
	Accelerator-Name=fulltext
	Accelerator-Args="-ignoreTS -min 2"

[Storage-Well "regexindex"]
	Location=/opt/gravwell/storage/regexindex
	Tags=regexindex
	Enable-Transparent-Compression=true
	Accelerator-Name=regex
	Accelerator-Engine-Override=index
	Accelerator-Args="^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"

[Storage-Well "regexbloom"]
	Location=/opt/gravwell/storage/regexbloom
	Tags=regexbloom
	Enable-Transparent-Compression=true
	Accelerator-Name=regex
	Accelerator-Engine-Override=bloom
	Accelerator-Args="^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
```

### 検証機

クエリ、取り込み、およびストレージのパフォーマンス特性は、データセットとハードウェアプラットフォームごとに異なりますが、このテストでは、次のハードウェアを使用しています。

| コンポーネント| 説明|
| ----------- | ------------- |
| CPU | AMD Ryzen 1700 |
| メモリ| 16GB DDR4-2400 |
| ディスク| サムスン960EVO NVME |
| ファイルシステム| zstd透過圧縮を使用したBTRFS |

これらのテストは、Gravwellバージョン `3.1.5`を使用して実施されました。

### 取り込みパフォーマンス

取り込みには、singleFileingesterを使用します。singleFile ingesterは、改行で区切られた単一のファイルを取り込むように設計されており、タイムスタンプを取得します。ingesterはタイムスタンプを取得しているため、CPUリソースが必要です。singleFile ingesterは、[githubページ](https://github.com/gravwell/ingesters/)で入手できます。 singleFileingesterの正確な呼び出しは次のとおりです。

```
./singleFile -i apache_fake_access_10M.log -clear-conns 172.17.0.3 -block-size 1024 -timeout=10 -tag-name=fulltext
```

| ウェル| 1秒あたりのエントリ数| データレート|
| ------------ | -------------------- | ------------ |
| raw| 313.54 KE / s | 65.94 MB / s |
| regexbloom | 112.91 KE / s | 23.75 MB / s |
| regexindex | 57.58 KE / s | 12.11 MB / s |
| 全文| 26.37 KE / s | 5.55MB /秒|

取り込みパフォーマンスから、フルテキストアクセラレーションシステムが取り込みパフォーマンスを大幅に低下させることがわかります。5.55MB/sは取り込みパフォーマンスが低いように見えますが、これは依然として約480GBのデータと1日あたり23億のエントリであることに言及する価値があります。

### ストレージの使用量

取り込みパフォーマンスといくつかの追加メモリ要件以外では、アクセラレーションを有効にすることの主なペナルティは使用量です。各抽出方法のインデックスエンジンは50％以上多くのストレージを消費し、ブルームエンジンはさらに4％を消費したことがわかります。ストレージの使用量は消費されるデータに大きく依存しますが、平均すると、インデックス作成システムはかなり多くのストレージを消費します。

| well| 使用済みストレージ| 生との違い|
| ------------ | -------------- | --------------- |
| raw| 2.49GB | 0％|
| 全文| 3.83GB | 53％|
| regexindex | 3.76GB | 51％|
| regexbloom | 2.60GB | 4％|

### クエリのパフォーマンス

クエリのパフォーマンスの違いを示すために、スパースとデンスに分類できる2つのクエリを実行します。スパースクエリは、データセット内の特定のIPを検索し、ほんの一握りのエントリを返します。密なクエリは、データセットでかなり一般的な特定のURLを探します。クエリを簡略化するために、アクセラレーションシステムに一致するregexindexタグとregexbloomタグのaxモジュールをインストールします。密なクエリはデータセット内のエントリの約12％を取得しますが、疎なクエリは約0.01％を取得します

スパースクエリとデンスクエリは次のとおりです。

```
ax ip=="106.218.21.57"
ax url=="/list" method | count by method | chart count by method
```

各クエリを実行する前に、rootとして次のコマンドを実行してシステムキャッシュを削除します。

```
echo 3 > /proc/sys/vm/drop_caches
```

|  Well      | クエリタイプ | クエリタイム| 処理済みエントリの割合| 加速 |
|------------|------------|------------|------------------------------|---------|
| raw        | sparse     | 71.5s      |  100%                        | 0X      |
| regexbloom | sparse     | 397ms      |  0.00389%                    | 180X    |
| regexindex | sparse     | 190ms      |  0.000001%                   | 386X    |
| fulltext   | sparse     | 195ms      |  0.000001%                   | 376X    |
| raw        | dense      | 73.5s      |  100%                        | 0X      |
| regexbloom | dense      | 71.5s      |  100%                        | 1.02X   |
| regexindex | dense      | 14.2s      |  13%                         | 5.17X   |
| fulltext   | dense      | 24.6s      |  30%                         | 2.98X   |

注：正規表現検索モジュール/自動抽出機能は全文アクセラレーターと完全には互換性がないため、アクセラレーターを使用するにはクエリを少し変更する必要があります。それらは `` `grep -w" 106.218.21.57 "` ``と `` `grep -w list | ax url == "/ list"メソッド| 方法によるカウント| メソッド別のチャート数 `` `

#### 全文

上記のベンチマークは、フルテキストアクセラレータには重大な取り込みとストレージのペナルティがあり、クエリ例はそれらの費用を正当化するようには見えないことを非常に明確にしています。データがタブ区切り、csv、jsonなどの完全にトークンベースであり、すべてのトークンが完全に控えめなもの（単一の単語、数値、IP、値など）である場合、フルテキストアクセラレータはあまり意味がありません。ただし、データに複雑なコンポーネントが含まれている場合、フルテキストアクセラレータは他のアクセラレータでは実行できないことを実行できます。Apacheを組み合わせたアクセスログを使用してきました。フルテキストアクセラレータを実際に輝かせるクエリを見てみましょう。

URLのサブコンポーネントを調べて、PowerPCMacintoshコンピューターを使用して`/apps`サブディレクトリを閲覧しているユーザーのグラフを取得します。上記の例の正規表現は、Apacheログ内の完全なフィールドにインデックスを付けます。それらはドリルダウンしてそれらのフィールドの一部を加速に使用することはできません、フルテキストはできます。

フルテキストインデクサーとその他の両方のクエリを最適化して公平を期しますが、どちらのクエリもどちらのデータセットでも機能します。

フルテキストアクセラレータ最適化クエリ：
```
grep -s -w apps Macintosh PPC | ax url~"/apps" useragent~"Macintosh; U; PPC" | count | chart count
```

全文以外に最適化されたクエリ：
```
ax url~"/apps" useragent~"Macintosh; U; PPC" | count | chart count
```

結果は、フルテキストがストレージと取り込みのペナルティに値することが多い理由を示しています。

| well| クエリ時間| 加速|
| ------------ | ------------ | --------- |
| raw| 71.7秒| 0X |
| regexbloom | 72.6秒| 〜0X |
| regexindex | 72.6 | 〜0X |
| 全文| 5.73秒| 12.49X |


#### AXモジュールのクエリ

4つのタグすべてのAX定義ファイルは以下のとおりです。詳細については、[AX]()のドキュメントを参照してください。

```
[[extraction]]
  tag = 'regexindex'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apacheindex'
  desc = 'apache index'

[[extraction]]
  tag = 'regexbloom'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apachebloom'
  desc = 'apache bloom'

[[extraction]]
  tag = 'fulltext'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apachefulltext'
  desc = 'apache fulltext'

[[extraction]]
  tag = 'raw'
  module = 'regex'
  params = "^(?P<ip>\\S+) (?P<ident>\\S+) (?P<username>\\S+) \\[([\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<method>\\S+)\\s?(?P<url>\\S+)?\\s?(?P<proto>\\S+)?\" (?P<resp>\\d{3}|-) (?P<bytes>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
  name = 'apacheraw'
  desc = 'apache raw'
```
