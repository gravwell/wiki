# Collectdインジェスター

Collectdインジェスターは完全にスタンドアローンの[collectd](https://collectd.org/)コレクションエージェントで、collectdのサンプルを直接Gravwellに送ることができます。 インジェスターは複数のコレクターをサポートしており、それぞれに異なるタグ、セキュリティコントロール、プラグインからタグへのオーバーライドを設定できます。

## 基本設定

Collectdインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、Collectdインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## コレクター例

```
[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	Security-Level=encrypt
	User=user
	Password=secret
	Encoder=json

[Collector "example"]
	Bind-String=10.0.0.1:9999 #default is "0.0.0.0:25826
	Tag-Name=collectdext
	Tag-Plugin-Override=cpu:collectdcpu
	Tag-Plugin-Override=swap:collectdswap
```

## インストール
GravwellのDebianリポジトリを使用している場合、インストールはaptコマンド1つで完了します:

```
apt-get install gravwell-collectd
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードしてください。Gravwellサーバのターミナルを使って、スーパーユーザーとして（`sudo`コマンドなどで）以下のコマンドを実行し、インジェスターをインストールします:

```
root@gravserver ~ # bash gravwell_collectd_installer.sh
```

Gravwellのサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメータを抽出し、適切に設定します。しかし、インジェスターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラは認証トークンとGravwellインデクサーのIPアドレスを要求します。これらの値はインストール時に設定するか、あるいは空欄にして、`/opt/gravwell/etc/collectd.conf`の設定ファイルを手動で修正することができます。

## 設定

Collectdインジェスターは、他のすべてのインジェスターと同じグローバル設定システムに依存しています。グローバルセクションは、インデクサの接続、認証、およびローカルキャッシュコントロールを定義するために使用されます。

コレクター設定ブロックは、collectdのサンプルを受け入れることができるリスニングコレクターを定義するために使用されます。 各コレクター設定は、固有のセキュリティレベル、認証、タグ、ソースオーバーライド、ネットワークバインド、タグオーバーライドを持つことができます。複数のコレクター設定を使用すると、単一のcollectdインジェスターが複数のインターフェイスをリッスンし、複数のネットワークエンクレーブから来るcollectdサンプルに固有のタグを適用できます。

デフォルトでは、collectdインジェスターは _/opt/gravwell/etc/collectd.conf_ にある設定ファイルを読み込みます。

### 設定例

```
[Global]
	Ingest-Secret = SuperSecretKey
	Connection-Timeout = 0
	Cleartext-Backend-target=192.168.122.100:4023
	Log-Level=INFO

[Collector "default"]
	Bind-String=0.0.0.0:25826
	Tag-Name=collectd
	User=user
	Password=secret

[Collector "localhost"]
	Bind-String=[fe80::1]:25827
	Tag-Name=collectdlocal
	Security-Level=none
	Source-Override=[fe80::beef:1000]
	Tag-Plugin-Override=cpu:collectdcpu
```

### コレクター設定オプション

各コレクターブロックには、固有の名前と重複しないバインド文字列が必要です。 同じポートの同じインターフェイスにバインドされた複数のコレクターを持つことはできません。

#### Bind-String

Bind-Stringは、コレクターが受信するcollectdサンプルのリッスンに使用するアドレスとポートを制御します。有効な Bind-String には、IPv4 または IPv6 のアドレスとポートのいずれかが含まれていなければなりません。 すべてのインターフェースをリッスンするには、「0.0.0.0」のワイルドカードアドレスを使用します。

##### Bind-Stringの例
```
Bind-String=0.0.0.0:25826
Bind-String=127.0.0.1:25826
Bind-String=127.0.0.1:12345
Bind-String=[fe80::1]:25826
```

#### Tag-Name

Tag-Nameは、Tag-Plugin-Override が適用されない限り、collectd のサンプルに割り当てられるタグを定義します。

#### Source-Override

Source-Overrideディレクティブは、Gravwellに送信される際にエントリに適用されるSource値をオーバーライドするために使用します。 デフォルトでは、インジェスターはインジェスターのSourceを適用しますが、検索時にセグメンテーションやフィルタリングを適用するために、コレクターブロックに特定のSource値を適用することが望ましい場合があります。Source-Overrideは任意の有効なIPv4またはIPv6アドレスです。

#### Source-Overrideの例
```
Source-Override=192.168.1.1
Source-Override=[DEAD::BEEF]
Source-Override=[fe80::1:1]
```

#### Security-Level

Security-Level ディレクティブは、コレクターが collectd パケットを認証する方法を制御します。 利用可能なオプションは、encrypt、sign、none です。 デフォルトでは、コレクターは "encrypt "セキュリティレベルを使用し、ユーザとパスワードの両方を指定する必要があります。 "none"の場合は、ユーザもパスワードも必要ありません。

#### Security-Levelの例
```
Security-Level=none
Security-Level=encrypt
Security-Level = sign
Security-Level = SIGN
```

#### User と Password

Security-Levelが "sign"または "encrypt"に設定されている場合、エンドポイントで設定されている値と一致するユーザー名とパスワードを提供する必要があります。 デフォルト値は、collectdに同梱されているデフォルト値に合わせて、「user」と「secret」になっています。 collectd のデータに機密情報が含まれている可能性がある場合は、これらの値を変更する必要があります。

##### User と Password の例
```
User=username
Password=password
User = "username with spaces in it"
Password = "Password with spaces and other characters @$@#@()*$#W)("
```

#### Encoder

collectdのデフォルトのエンコーダーはJSONですが、シンプルなテキストエンコーダーも用意されています。 オプションは "JSON" または "text" です。

JSONエンコーダーを使ったエントリーの例です。:

```
{"host":"build","plugin":"memory","type":"memory","type_instance":"used","value":727789568,"dsname":"value","time":"2018-07-10T16:37:47.034562831-06:00","interval":10000000000}
```

## Tag Plugin Overrides

各 Collector ブロックは N 個の Tag-Plugin-Override 宣言をサポートしています。この宣言は、Collectd サンプルを生成したプラグインに基づいて一意のタグを適用するために使用されます。 Tag-Plugin-Overrides は、異なるプラグインからのデータを異なるウェルに保存し、異なるエージアウトルールを適用したい場合に役立ちます。 例えば、ディスク使用量に関するcollectdレコードを9ヶ月間保存することは価値があるかもしれませんが、CPU使用量のレコードは14日で期限切れになります。 Tag-Plugin-Overrideシステムはこれを容易にします。

Tag-Plugin-Overrideのフォーマットは、":"文字で区切られた2つの文字列で構成されています。 左側の文字列はプラグインの名前を表し、右側の文字列は希望するタグの名前を表します。 タグに関する通常のルールがすべて適用されます。 1つのプラグインを複数のタグに対応させることはできませんが、複数のプラグインを同じタグに対応させることはできます。

### Tag Plugin Overridesの例
```
Tag-Plugin-Override=cpu:collectdcpu # Map CPU plugin data to the "collectdcpu" tag.
Tag-Plugin-Override=memory:memstats # Map the memory plugin data to the "memstats" tag.
Tag-Plugin-Override= df : diskdata  # Map the df plugin data to the "diskdata" tag.
Tag-Plugin-Override = disk : diskdata  # Map the disk plugin data to the "diskdata" tag.
```
