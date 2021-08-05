# TLS 証明書の設定

GravwellはデフォルトでTLS証明書を使用せずに出荷されます。自動生成された自己署名証明書を使用すると、ブラウザの警告でユーザーを怖がらせ、誤った安全性を提供する傾向があるため、このようにしました。自己署名証明書を適切に検証することは困難であり、ユーザーに偽装の可能性のある証明書を受け入れるように教育するリスクがあります。 これは、Chromium ベースのブラウザの挙動が非常に気まぐれで、予測不可能な方法で証明書の例外をタイムアウトさせてしまうことにも起因しています（証明書を再受諾するために、Chromium/Chrome のザイゴートプロセスをすべて閉じなければならないことがよくあります）。

Gravwellシステムをインターネットに公開するつもりなら、信頼できるプロバイダから完全に検証された証明書を取得することを強くお勧めします。 [LetsEncrypt](https://letsencrypt.org)の人々は、適切な証明書の検証について学ぶための素晴らしいリソースであり、主要なブラウザで信頼されている無料の証明書を提供しています。

Gravwell 管理者には、証明書のための 3 つのオプションがあります：

* 暗号化されていない HTTP のみを使用し続けます。これは、信頼されたプライベートネットワーク上でのみアクセスするインストールや、GravwellがnginxのようなHTTPプロキシでフロントする場合に適しています。
* 適切に署名された TLS 証明書をインストールします。これは理想的な構成ですが、一般的にはGravwellインスタンスが公開されたホスト名を持つ必要があります。
* 自己署名証明書をインストールします。これは、Gravwellへのトラフィックを暗号化したいけれども、何らかの理由で適切に署名された証明書を取得できない場合に理にかなっています。

## HTTPのみの使用

これはGravwellのデフォルト設定であり、使用するための変更は必要ありません。ホームネットワークで Gravwell を実験している人や、仕事のために実験的なネットワークで Gravwell を評価している人に適しています。また、Gravwell ウェブサーバが nginx のようなロードバランサ/リバースプロキシを介してアクセスされる場合にも許容できる設定です。これにより、プロキシが HTTPS 暗号化/復号化を実行し、Gravwell システムの負荷を軽減することができます。

証明書がない場合、インジェスタはインデクサへのトラフィックを暗号化できないことに注意してください。インジェスタのトラフィックを暗号化したいが、ウェブサーバをHTTPのみのモードにしておきたい場合、他のセクションで説明したように証明書をインストールすることができますが、 gravwell.conf の `Certificate-File`, `Key-File`, `TLS-Ingest-Port` オプションのコメントを外すだけです。これにより、インデクサに対してはTLSが有効になりますが、ウェブサーバに対しては有効になりません。

注意: 分散型ウェブサーバとデータストアをHTTPSを無効にして設定する場合、データストアとウェブサーバの両方に gravwell.conf で `Datastore-Insecure-Disable-TLS` フラグを設定しなければなりません。

## 適切に署名されたTLS証明書をインストールする

適切に署名されたTLS証明書は、Gravwellにアクセスする最も安全な方法です。ブラウザは文句を言わずに証明書を自動的に受け入れます。

証明書の取得はこの文書の範囲外です。従来のプロバイダで証明書を購入するか、[LetsEncrypt](https://letsencrypt.org)を利用して無料の証明書を取得することを検討してください。

証明書を使用するためには、Gravwell に証明書と鍵のファイルがどこにあるかを教えなければなりません。ファイルが `/etc/certs/cert.pem` と `/etc/certs/key.pem` にあると仮定して、gravwell.conf を編集して `Certificat-File` と `Key-File` オプションをアンコメントして入力します：

```
Certificate-File=/etc/certs/cert.pem
Key-File=/etc/certs/key.pem
```

注意: これらのファイルは「gravwell」ユーザーが読めるようにしなければなりません。しかし、他のユーザーから鍵ファイルを保護するように注意してください。

ウェブサーバでHTTPSを有効にするには、`Web-Port` ディレクティブを80から443に変更し、`Insecure-Disable-HTTPS` ディレクティブをコメントアウトします。

TLSで暗号化されたインジェスター接続を有効にするには、`TLS-Ingest-Port=4024`という行を見つけてコメントを外します。

検索エージェントでHTTPSを有効にするには、/opt/gravwell/etc/searchagent.confを開き、`Insecure-Use-HTTP=true`行をコメントアウトし、`Webserver-Address`行のポートを80から443に変更します。

最後に、ウェブサーバ、インデクサ、検索エージェントを再起動します：

```
systemctl restart gravwell_webserver.service
systemctl restart gravwell_indexer.service
systemctl restart gravwell_searchagent.service
```

注意: データストアと複数のWebサーバを使用する場合、自己署名証明書を使用してWebサーバ同士が通信できるようにするには、`Search-Forwarding-Insecure-Skip-TLS-Verifyパラメータ`を`true`に設定しなければなりません。データストアが自己署名証明書を使用している場合は、`Datastore-Insecure-Skip-TLS-Verify`をウェブサーバに設定して、ウェブサーバがデータストアと通信できるようにします。

## 自己署名証明書のインストール

適切なTLS証明書ほど安全ではありませんが、自己署名証明書はユーザーとGravwell間の暗号化された通信を保証します。ブラウザに自己署名証明書を信頼するように指示することで、繰り返し発生する警告画面を回避することも可能です。

まず、`/opt/gravwell/etc`にある `gencert` というGravwellインストール時に同梱されているプログラムを使って、1年間の証明書を生成します：

```
cd /opt/gravwell/etc
sudo -u gravwell ../bin/gencert -h HOSTNAME
```

HOSTNAMEをGravwellシステムのホスト名またはIPアドレスに置き換えてください。カンマで区切ることで、複数のホスト名またはIPを指定することができます。例: `gencert -h gravwell.floren.lan,10.0.0.0.1,192.168.0.3`。

gravwell.confを開き、`Certificate-File` と `Key-File` ディレクティブのコメントを外してください。デフォルトでは、先ほど作成した2つのファイルが正しく指し示されているはずです。

ウェブサーバでHTTPSを有効にするには、`Web-Port` ディレクティブを80から443に変更し、`Insecure-Disable-HTTPS` ディレクティブをコメントアウトします。

TLSで暗号化されたインジェスター接続を有効にするには、`TLS-Ingest-Port=4024`行を見つけてコメントアウトします。

検索エージェントでHTTPSを有効にするには、/opt/gravwell/etc/searchagent.confを開き、`Insecure-Use-HTTP=true`行をコメントアウトし、`Webserver-Address`行のポートを80から443に変更します。

最後に、ウェブサーバ、インデクサ、検索エージェントを再起動します：

```
systemctl restart gravwell_webserver.service
systemctl restart gravwell_indexer.service
systemctl restart gravwell_searchagent.service
```

### 自己署名証明書をブラウザに信頼させる

証明書が認められたルートCAによって署名されていない場合、ブラウザは警告を表示します。しかし、証明書を手動でインストールすることで、ブラウザに証明書を信頼させることができます。

#### Firefox

Firefoxに証明書をインストールするのは非常に簡単です。まず、HTTPS経由でGravwellインスタンスに移動します。Firefoxはこのような画面を表示するはずです：

![](firefox-warning.png)

詳細設定ボタンをクリックしてください：

![](firefox-warning-advanced.png)

次に、「例外を追加...」をクリックします。

![](firefox-exception.png)

デフォルトは適切なものでなければなりませんが、「この例外を永続的に保存する」がチェックされていることを確認してください。「Confirm Security Exception」をクリックします。

これで Firefox は証明書の有効期限が切れるまで自己署名証明書を受け入れるようになりました。

#### Chrome

Chromeブラウザで証明書をインストールするのは少し複雑です。まず、HTTPS経由でGravwellインスタンスに移動します。Chromeは警告画面を表示します。

![](chrome-warning.png)

アドレスバーの「安全でない」ラベルをクリックします。

![](chrome-export1.png)

次に、「Certificate」の下にある「Invalid」リンクをクリックします。証明書ビューア」ウィンドウが開きますので、「詳細」タブをクリックします。

![](chrome-export2.png)

「エクスポート」ボタンを選択します。Chromeは証明書を保存するためのファイルダイアログを表示するので、どこかに保存して場所を覚えておきます。

アドレスバーに[chrome://settings](chrome://settings)と入力するか、Chromeブラウザのメニューから設定を開きます。一番下までスクロールして、「詳細設定」ボタンをクリックします。

![](chrome-advanced.png)

「プライバシーとセキュリティ」セクション内で、「証明書の管理」を見つけてクリックします。

![](chrome-advanced2.png)

ここで「Authorities」タブを選択し、「Import」をクリックします。

![](chrome-authorities.png)

ファイルダイアログが開きますので、先ほど保存した証明書ファイルを選択します。次のダイアログで「この証明書をウェブサイトを識別するために信頼する」にチェックを入れ、「OK」をクリックします。

![](chrome-import.png)

これで、SSLの警告なしでGravwellタブを更新できるようになりました。
