# Amazon SQS インジェスター

Amazon SQS Ingester（sqsIngester）は、標準的なSQSキューとFIFO SQSキューの両方をサブスクライブしてインジェストすることができるシンプルなインジェスターです。 Amazon SQSは、メッセージ配信保証、メッセージの"soft"オーダー、メッセージの"at-least-once"配信をサポートする大容量のメッセージキューサービスです。 

Gravwellの場合、"at-least-once"配信には重要な注意点があります。SQSインジェスターは、同一のタイムスタンプを持つ重複したメッセージを受信する可能性があります（設定によります）。 また、SQSのワークフローが他の接続されたサービスとどのように展開されているかによって、SQSインジェスターに一部のメッセージが表示されない可能性もあります。 詳しくは[Amazon SQS](https://aws.amazon.com/sqs/)をご覧ください。

## 基本設定

SQSインジェスターは、[インジェスターセクション]（#!ingesters/ingesters.md#Global_Configuration_Parameters）に記載されている統一されたグローバル設定ブロックを使用します。  他の多くのGravwellインジェスターと同様に、SQSは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## キューの例

```
[Queue "default"]
	Region="us-east-2"
	Queue-URL="https://us-east-2.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
	Assume-Local-Timezone=false #Default for assume localtime is false
	Source-Override="DEAD::BEEF" #override the source for just this Queue 

[Queue "default"]
	Region="us-west-1"
	Queue-URL="https://us-west-1.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
```

## インストール
GravwellのDebianリポジトリを使用している場合、インストールはaptコマンド1つで完了します:

```
apt-get install gravwell-sqs
```

それ以外の場合は、[ダウンロードページ](#!quickstart/downloads.md)からインストーラーをダウンロードしてください。 Netflow インジェスターをインストールするには、root 権限でインストーラーを実行します（実際のファイル名には通常、バージョン番号が含まれます）:

```
root@gravserver ~ # bash gravwell_sqs.sh
```

ローカルマシンにGravwellインデクサーがない場合、インストーラはIngest-Secretの値とインデクサー（またはフェデレーター）のIPアドレスの入力を求める。 そうでない場合は、既存のGravwellの設定から適切な値を引き出します。いずれにしても、インストール後は `/opt/gravwell/etc/sqs.conf` にある設定ファイルを確認してください。典型的な設定は以下のようになります: 

```
[Global]
Ingest-Secret = IngestSecrets
Connection-Timeout = 0
Insecure-Skip-TLS-Verify=false
Pipe-Backend-Target=/opt/gravwell/comms/pipe 
Log-Level=INFO
Log-File=/opt/gravwell/log/sqs.log

# A Queue pulls from a specific SQS queue with a given AKID and Secret. See
# https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
# for information about obtaining an AKID/Secret for your user.
[Queue "default"]
	Region="us-east-2"
	Queue-URL="https://us-east-2.amazon..."
	Tag-Name="sqs"
	AKID="AKID..."
	Secret="..."
```

この設定では、`/opt/gravwell/comms/pipe`を介してローカルのインデクサにエントリーを送信することに注意してください。エントリーには'sqs'というタグが付けられます。

SQSキューごとに1つずつ、任意の数の`Queue`エントリーを構成し、それぞれに固有の認証やタグ名などを提供することができます。
