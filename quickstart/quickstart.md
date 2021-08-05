# クイックスタート手順

このセクションは、Gravwellを起動して実行するための基本的なクイックスタート手順です。この説明は、最も一般的なユースケースをサポートし、Gravwell入門としても機能します。 クイックスタートの手順では、Cluster Editionで利用可能な、分散検索とストレージに関する高度なGravwellの機能は利用できないことに注意してください。より高度な設定が必要な場合は、本ガイドの「高度なトピック」のセクションを参照してください。

このガイドはCommunity Editionユーザーおよび、有料のシングルノードGravwellサブスクリプションをお持ちのユーザーに適しています。

注：Community Editionユーザーは、インストールを開始する前に[https://www.gravwell.io/download](https://www.gravwell.io/download)から独自のライセンスを取得する必要があります。 購入済みユーザーは、すでに電子メールでライセンスファイルを受け取っているはずです。

## インストール
Gravwellを1台のマシンにインストールするのは非常に簡単です。このセクションの指示に従ってください。 複数のシステムを含むより高度な環境については、「高度なトピック」セクションを確認してください。

Gravwellは4つの方法で配布されています。 Docker コンテナ経由、ディストリビューションに依存しない自己解凍インストーラー経由、Debian パッケージリポジトリ経由、Redhat パッケージリポジトリ経由です。システムが Debian や Ubuntu を動かしている場合は Debian リポジトリ、システムが RHEL、CentOS、SuSE を動かしている場合は Redhat パッケージ、それ以外の場合は自己解凍インストーラーを使うことをお勧めします。Docker ディストリビューションは、Docker に慣れている人にも便利です。Gravwellは主要なLinuxディストリビューションのすべてでテストされており、問題なく動作しますが、より望ましいのはUbuntu Server LTSです。Ubuntuのインストールに関するヘルプは、https://tutorials.ubuntu.com/tutorial/tutorial-install-ubuntu-server を参照してください。

### Debianリポジトリ

Debianリポジトリからのインストールは非常に簡単です。 最初にいくつかの手順を実行して、GravwellのPGP署名キーおよびDebianパッケージリポジトリを追加する必要がありますが、その後は単に`gravwell`パッケージをインストールするだけです。

```
sudo apt install apt-transport-https gnupg curl
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debian/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get update
sudo apt-get install gravwell
```

インストールプロセスでは、Gravwellのコンポーネントが使用するいくつかの共有シークレット値を設定するよう求められます。 セキュリティのために、インストーラーがランダムな値（デフォルト）を生成できるようにすることを強くお勧めします。

![Read the EULA](eula.png)

![Accept the EULA](eula2.png)

![Generate secrets](secret-prompt.png)

### Redhat/CentOS リポジトリ

Gravwell は、Redhat と CentOS Linux ディストリビューションの両方で `yum` リポジトリとして利用できます。Gravwell の yum リポジトリを使うには、以下のスタンザを `yum.conf` (`/etc/yum.conf` にあります) に追加してください。

```
[gravwell]
name=gravwell
baseurl=https://update.gravwell.io/rhel
gpgkey=https://update.gravwell.io/rhel/gpg.key
```

次に以下を実行します。

```
yum update
yum install -y gravwell
```

インストールしたら、次のように実行して、webports用のcentOSファイアウォールをバンプします。

```
sudo firewall-cmd --zone=public --add-service=http
sudo firewall-cmd --zone=public --add-service=https
```

これで、centOS/RHELシステムに割り当てられたIP上のGravwell Webインターフェースにアクセスできるようになります。

### Dockerコンテナ

Gravwellは、ウェブサーバーとインデクサーの両方を含む単一のコンテナとしてDockerhubで利用できます。 GravwellをDockerにインストールする詳細な手順については、[Dockerのインストール手順](#!configuration/docker.md) を参照してください。

### 自己完結型インストーラー

Debian以外のシステムの場合、[ダウンロードページ](#!quickstart/downloads.md)から自己完結型のインストーラーをダウンロードします。

次に、インストーラーを実行します。

```
sudo bash gravwell_X.X.X.sh
```

プロンプトに従い、完了後、実行中のGravwellインスタンスが必要です。

注：ディストリビューションがSystemDを使用していない場合、インストール後にGravwellプロセスを手動で開始する必要があります。 サポートが必要な場合は、support@gravwell.ioまでご連絡ください。

## ライセンスの設定

Gravwellがインストールされたら、Webブラウザを開き、サーバーに移動します(例: [http://localhost/](http://localhost/))。ライセンスファイルをアップロードするよう求められます。

![ライセンスをアップロード](upload-license.png)

アップロードしたライセンスが検証されると、Gravwell ログイン画面が表示されます。ユーザー名「admin」とパスワード「changeme」を入力しログインします。

重要：Gravwellのデフォルトのユーザー名/パスワードの組み合わせはadmin/changemeです。 できるだけ早く管理者パスワードを変更することを強くお勧めします！ これを行うには、ナビゲーションサイドバーから[アカウント設定]を選択するか、右上の[ユーザー]アイコンをクリックします。

![](login.png)

## インジェスターの設定

インストールされたばかりのGravwellインスタンスには、検索できるデータがありません。データを取り込むためにはインジェスターが必要です。リポジトリからパッケージをインストールするか、[ダウンロード](downloads.md)にアクセスして、各インジェスターの自己解凍インストーラーを取得することができます。

Debianリポジトリで利用可能なインジェスターは、`apt-cache search gravwell`を実行することで表示できます：

```
root@debian:~# apt-cache search gravwell
gravwell - Gravwell data analytics platform (gravwell.io)
gravwell-collectd - Gravwell collectd ingester
gravwell-crash-reporter - Gravwell crash reporter service
gravwell-datastore - Gravwell datastore service
gravwell-federator - Gravwell ingest federator
gravwell-file-follow - Gravwell file follow ingester
gravwell-http-ingester - Gravwell HTTP ingester
gravwell-ipmi - Gravwell IPMI ingester
gravwell-kafka - Gravwell Kafka ingester
gravwell-kafka-federator - Gravwell Kafka federator
gravwell-kinesis - Gravwell Kinesis ingester
gravwell-loadbalancer - Gravwell load balancing service
gravwell-netflow-capture - Gravwell netflow ingester
gravwell-network-capture - Gravwell packet ingester
gravwell-o365 - Gravwell Office 365 log ingester
gravwell-offline-replication - Gravwell offline replication service
gravwell-packet-fleet - Gravwell Packet Fleet ingester
gravwell-pubsub - Gravwell ingester for Google Pub/Sub streams
gravwell-shodan - Gravwell Shodan ingester
gravwell-simple-relay - Gravwell simple relay ingester
gravwell-sqs - Gravwell SQS ingester
```

メインGravwellインスタンスと同じノードにインストールする場合、インデクサーに接続するように自動的に構成する必要がありますが、ほとんどの場合、データソースを設定する必要があります。 その手順については、[インジェスター設定](#!ingesters/ingesters.md) を参照してください。

最初の実験として、File Follow インジェスター（gravwell-file-follow）をインストールすることを強くお勧めします。これはLinuxログファイルを取り込むように事前に設定されているため、`tag=auth`などで検索することにより、いくつかのエントリをすぐに見ることができるはずです。

![Auth entries](auth.png)

### File Follower インジェスター

File Follower インジェスターは、標準のLinuxログファイルを取り込むように事前に構成されているため、ログをGravwellに取り込む最も簡単な方法の1つです。

Gravwell Debianリポジトリを使用している場合、インストールはaptコマンド1つで済みます。

```
apt-get install gravwell-file-follow
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードします。 Gravwellサーバー上のターミナルを使用して、次のコマンドをスーパーユーザーとして（たとえば`sudo`コマンド経由で）発行し、インジェスターをインストールします。

```
root@gravserver ~ # bash gravwell_file_follow_installer.sh
```

Gravwellサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメーターを抽出して設定し、適切に設定します。 ただし、ご使用のコンピューターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラーは認証トークンとGravwellインデクサーのIPアドレスの入力を求めます。 インストール時にこれらの値を設定するか、空白のままにして`/opt/gravwell/etc/file_follow.conf`の設定ファイルを手動で変更します。 インジェスターの構成の詳細については、[インジェスター設定](#!ingesters/ingesters.md)を参照してください。

### Simple Relay インジェスター

GravwellのSimple Relay グループは、ネットワーク経由で行区切りまたはsyslog形式のメッセージを取り込むことができます。 既存のデータソースからデータをGravwellに取り込むもう1つの良い方法です。

Gravwell Debianリポジトリを使用している場合、インストールはaptコマンド1つで済みます。

```
apt-get install gravwell-simple-relay
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードします。 Gravwellサーバー上のターミナルを使用して、次のコマンドをスーパーユーザーとして（たとえば`sudo`コマンド経由で）発行して、インジェスターをインストールします。

```
root@gravserver ~ # bash gravwell_simple_relay_installer.sh
```

Gravwellサービスが同じマシンに存在する場合、インストールスクリプトは自動的に`Ingest-Auth`パラメーターを抽出して設定し、適切に設定します。ただし、ご使用のコンピューターが既存のGravwellバックエンドと同じマシンに常駐していない場合、インストーラーは認証トークンとGravwellインデクサーのIPアドレスの入力を求めます。 インストール時にこれらの値を設定するか、空白のままにして`/opt/gravwell/etc/simple_relay.conf`の設定ファイルを手動で変更します。 インジェスターの構成の詳細については、[インジェスター設定](#!ingesters/ingesters.md)を参照してください。

### インジェスターのメモ
これらのクイックスタート手順にあるように、インストールが1台のマシンに完全に含まれている場合、インジェスターインストーラーは構成オプションを抽出し、適切に構成します。すべてのGravwellコンポーネントが単一のシステムで実行されているわけではない高度なセットアップを使用している場合は、[インジェスター設定](#!ingesters/ingesters.md)を確認してください。

これで、Gravwellサーバー上でFile FollowとSimple Relayサービスが動作するようになりました。File Followは、`/var/log/`のいくつかのファイルからログエントリを自動的に取り込みます。 デフォルトでは、「auth」タグ付きの/var/log/auth.log、「dpkg」タグ付きの/var/log/dpkg.logおよび「kernel」タグ付き/var/log/dmesgおよび/var/log/kern.logを取り込みます。

Simple Relayは、TCPポート601またはUDPポート514で送信されたsyslogエントリを取り込みます。これらは「syslog」タグでタグ付けされます。Simple Relayの設定ファイルには、ポート7777で行区切りデータをListenするためのエントリも含まれています。syslogのみを使用する場合、これを無効にできます。設定ファイルの`[Listener "default"]`セクションをコメントアウトして、Simple Relayサービスを再起動してください。 このサービスの設定ファイルは`/opt/gravwell/etc/simple_relay.conf`にあります。高度な設定オプションについては、[インジェスター設定](#!ingesters/ingesters.md)のSimple Relayセクションを参照してください。

## Gravwellへのデータの供給
このセクションでは、データをGravwellに送信するための基本的な手順を説明します。他のデータインジェスターの設定手順については、[インジェスター設定](#!ingesters/ingesters.md)セクションを確認してください。

Gravwellの「システム統計」ページは、Gravwellサーバーがデータを受信しているかどうかを確認するのに役立ちます。 データが表示されず、それがエラーによると思われる場合、インジェスターが実行されていること(`ps aux | grep gravwell`コマンドを実行し、`gravwell_webserver`、`gravwell_indexer`、`gravwell_simple_relay`、および`gravwell_file_follow`が表示される必要があります)と、それらの設定ファイルが正しいことを再確認してください。

![](stats.png)

### Syslogの取り込み
Gravwellサーバーがインストールされ、Simple Relayテキストインジェスターサービスが実行されたら、syslogプロトコルを介してログやテキストデータのGravwellへの供給を開始することができます。デフォルトでは、Simple Relayインジェスターはポート601のTCP syslogとポート514のUDP syslogをリッスンします。

rsyslogを実行しているLinuxサーバからGravwellにsyslogエントリを送信するには、サーバ上に`/etc/rsyslog.d/90-gravwell.conf`という名前の新しいファイルを作成し、以下の行をUDP syslog用に貼り付けます。ホスト名`gravwell.example.com`を自分のGravwellインスタンスに変更するよう注意してください。

```
*.* @gravwell.example.com;RSYSLOG_SyslogProtocol23Format
```

または、代わりにTCP syslogにこれを使用します。

```
*.* @@gravwell.example.com;RSYSLOG_SyslogProtocol23Format
```

（UDP構成での`@`の使用とTCPでの`@@`の使用に注意してください）

次に、rsyslogデーモンを再起動します。

```
sudo systemctl restart rsyslog.service
```

多くのLinuxサービス（DNS、Apache、sshなど）は、syslogを介してイベントデータを送信するように構成できます。 これらのサービスとGravwellの「間」としてsyslogを使用することは、多くの場合、イベントをリモートで送信するようにこれらのサービスを構成する最も簡単な方法です。

たとえば、Apacheの構成エントリにこの行を追加すると、すべてのApacheログがrsyslogに送信され、それらがGravwellに転送されます。

```
CustomLog "|/usr/bin/logger -t apache2.access -p local6.info" combined
```

### アーカイブされたログ
Simple Relay インジェスターは、ファイルシステム上にある古いログ（Apache、syslogなど）を取り込むためにも使用できます。 netcatのような基本的なネットワーク通信ツールを利用することにより、どんなデータでもSimple Relay インジェスターの行区切りリスナーに送り込むことができます。 デフォルトでは、Simple RelayはTCPポート7777で行区切りエントリをリッスンします。

たとえば、Gravwellで分析したい古いApacheログファイルがある場合、次のようなコマンドを実行して取り込むことができます。

```
user@webserver ~# cat /tmp/apache-oct2017.log | nc -q gravwell.server.address 7777
```

注：複数のファイルからなる非常に大きなログのセットを取り込む場合、Simple Relay インジェスターではなく、Mass File インジェスターユーティリティを使用して、事前に最適化してまとめて取り込むことをお勧めします。

### ネットワークパケットインジェスター

Gravwellの第一の強みは、バイナリデータを取り込めることです。ネットワークインジェスターを使用すると、分析用にネットワークから完全なパケットをキャプチャすることができます。ストレージの使用量は増えますが、Netflowまたはその他の凝縮されたトラフィック情報を保存するより、はるかに優れた柔軟性が提供されます。

Gravwell Debianリポジトリを使用している場合、インストールはaptコマンド1つで済みます。

```
apt-get install libpcap0.8 gravwell-network-capture
```

それ以外の場合は、[ダウンロード](#!quickstart/downloads.md)からインストーラーをダウンロードします。 ネットワークインジェスターをインストールするには、rootとしてインストーラーを実行するだけです（ファイル名は若干異なる場合があります）。

```
root@gravserver ~ # bash gravwell_network_capture_installer.sh
```

ネットワークインジェスターには、libpcap共有ライブラリが必要です。スタンドアロンインストーラーを使用する場合は、ライブラリもインストールされていることを確認する必要があります。パッケージはDebianでは `libpcap0.8`です。

もしインジェスターが、すでにGravwellバックエンドがインストールされているマシン上にある場合、インストーラーは自動的に正しい`Ingest-Secrets`値を取得し、それを設定ファイルに追加してくれるでしょう。いずれにしても、実行する前に`/opt/gravwell/etc/network_capture.conf`の設定ファイルを確認してください。Interfaceフィールドがシステムのネットワークインターフェースの1つに設定され、かつ少なくとも1つの"Sniffer"セクションがコメントアウトされていないことを確認してください。 詳細については、[インジェスター設定](#!ingesters/ingesters.md)を参照してください

注：Debianパッケージとスタンドアロンインストーラーば、どちらもキャプチャ元のデバイスを要求してくるはずです。設定を変更したい場合は、`/opt/gravwell/etc/network_capture.conf`を開き、希望するインターフェイスを設定し、`service gravwell_network_capture restart`を実行して、インジェスターを再起動します。

## 検索
Gravwellサーバーが稼働し、データを受信できるようになると、検索パイプラインの能力が利用可能になります。

ここでは、このクイックスタートでの設定で取り込まれたデータのタイプに基づいた検索の例を紹介します。これらの例では、Linuxサーバーによって生成されたsyslogデータが、Simple Relayインジェスターを介して取り込まれ、前のセクションで説明したようにネットワークからパケットがキャプチャされていると想定しています。

### syslogの例
Syslogは、Unixのロギングおよび監査操作の中核的コンポーネントです。 UNIXインフラストラクチャのデバッグと防御を行いながら、ログイン、クラッシュ、セッション、またはその他のサービスアクションを完全に可視化することが重要です。 Gravwellを使用すると、多くのリモートマシンからsyslogデータを中央に簡単に集約して検索できます。この例では、いくつかのSSHログを追跡し、管理者やセキュリティの専門家がSSHアクティビティを監視する方法を検討します。

この例は、GravwellインスタンスにSSHログインデータを送信するためのサーバです。すべてのSSH関連エントリの一覧を見たい場合は、以下のように検索を実行します。

```
tag=syslog grep ssh
```

検索コマンドの内訳は次のとおりです。

* `tag=syslog`: "syslog"とタグ付けされたデータを検索します。 Simple Relay インジェスターは、TCPポート601またはUDPポート514を介して入ってくるデータに"syslog"タグを付けるように設定されています。
* `grep ssh`: "grep"モジュール（同様のlinuxコマンドにちなんで命名）は、特定のテキストを検索します。 この場合、検索は"ssh"を含むエントリを探します。

検索結果は、概要グラフと一連のログエントリとして表示されます。概要グラフには、パイプラインを通過したマッチングレコードの頻度プロットが表示されます。このグラフを使用して、ログエントリの頻度を特定したり、検索のタイムウィンドウをナビゲートしたりして、検索を再実行せずに表示を絞り込むことができます。ほぼすべての検索には、タイムウィンドウをリフォーカスして調整する機能があり、時間の順序を変更する検索のみ、概要グラフが表示されません。

"ssh"を含むすべてのエントリの結果を以下に示します。

![概要グラフ](overview.png)

これらの結果も非常に大まかな洞察を与えてくれるかもしれませんが、本当に有用な情報を抽出するには、検索を絞り込む必要があります。この例では、成功したSSHログインを抽出します。また、ログレコードからいくつかの特定のフィールドを抽出して、結果を見やすくします。

```
tag=syslog syslog Appname==sshd Message~Accepted | regex -e Message "Accepted\s(?P<method>\S+)\sfor\s(?P<user>\S+)\sfrom\s(?P<ip>\S+)"
```

検索の内訳は以下の通りです：

* ```tag=syslog```："syslog"とタグ付けされたデータに検索を制限する
* ```syslog Appname==sshd Message~Accepted```：syslogモジュールを呼び出し、"sshd"アプリケーションによって生成され、かつメッセージ本文に"Accepted"という文字列を含むsyslogメッセージのみにフィルタリングします。
* ```regex -e Message "Accepted\s(?P<method>\S+)\sfor\s(?P<user>\S+)\sfrom\s(?P<ip>\S+)"```：syslogモジュールを用いて抽出したメッセージ本文のみを正規表現で処理しています。ユーザ、IP、ログインに成功した方法を抽出します。

これにより、結果はログインのみに絞り込まれます：

![ログインに絞り込まれた検索](logins-only.png)

結果の下部にある「列挙値」ボタンをクリックすると、この検索で抽出された使用可能な列挙値がすべて表示されます。

![ログインに絞り込まれた検索](logins-only-enums.png)

検索文の最後に*レンダリングモジュール*を指定して、結果の表示方法を変更できます。 ログインしたすべてのユーザー名のグラフが必要な場合は、次の検索を実行します。

```
tag=syslog syslog Appname==sshd Message~Accepted | regex -e Message "Accepted\s(?P<method>\S+)\sfor\s(?P<user>\S+)\sfrom\s(?P<ip>\S+)" | count by user | chart count by user
```

新しい検索クエリアイテムの内訳は次のとおりです。

* ```count by user```: `count`モジュールは（regexによって抽出された）各`user`値が出現する回数をカウントします。
* ```chart count by user```: カウントモジュールの出力をチャート描画レンダラーにパイプし、カウントモジュールの結果によって決定される大きさでユーザーごとに個別の線を描画します。

結果には、検索期間中にシステムにログインしたすべてのユーザーの素晴らしいグラフが表示されます。グラフの種類を変更してデータにさまざまなビューを表示したり、概要チャートを使用して結果のより短い時間枠を選択したりできます。予想通り、最近これらのシステムにログインしているのはIT管理者の「クリス」だけのようです。

![ユーザーによる検索カウント](users-chart.png)

チャートアイコン（ジグザグ線）をクリックして、チャートのタイプを変更することもできます。 これは、棒グラフに表示されるまったく同じデータです。

![ユーザーによる検索カウント](users-chart-bar.png)

### ネットワークの例
ユーザーがLinuxルーターでパケットキャプチャインジェスターを設定し、インターネットとの間で送受信されるすべてのパケットをキャプチャするホームネットワークの例を考えてみましょう。このデータを使用して、特定のゲームがいつ遊ばれたかといった利用パターンを分析できます。サンプルは10.0.0.0/24ネットワークサブネットを使用し、Blizzard Entertainmentゲームがゲームトラフィックにポート1119を使用しています。次の検索では、どのPCがいつBlizzardゲームをプレイしているかが表示されます。

```
tag=pcap packet ipv4.DstIP !~ 10.0.0.0/24 tcp.DstPort==1119 ipv4.SrcIP | count by SrcIP | chart count by SrcIP
```

検索コマンドの内訳は次のとおりです。

* ```tag=pcap```: Gravwellに、"pcap"というタグの付いたアイテムのみを検索するように指示します。
* ```packet```: パケット解析検索パイプラインモジュールを呼び出し、このコマンドの残りのオプションを有効にします。
  * ```ipv4.DstIP !~ 10.0.0.0/24```: Gravwellパケットパーサーは、パケットをさまざまなフィールドに分割します。 この場合、検索は宛先IPを比較し、10.0.0.xクラスCサブネットにないものを探しています。
  * ```tcp.DstPort == 1119```: 宛先ポートを指定します。 これにより、ほとんどのBlizzard Entertainmentゲームで使用されるポート1119宛てのパケットのみがフィルターされます。
  * ```ipv4.SrcIP```: 比較演算子なしでこのフィールドを指定すると、パケットパーサーはソースIPを抽出してパイプラインに配置します。
* ```count by SrcIP```: フィルタリングされた結果をパケットパーサーからカウントモジュールにパイプし、各ソースIPが表示される回数をカウントするように指示します。
* ```chart count by SrcIP```: カウント結果をグラフ表示レンダラーにパイプして表示し、ソースIP値ごとに個別の線を描画します。

結果：2つのシステムがポート1119にトラフィックを送信しています。黄色（10.0.0.6）で表されるIPはパッシブトラフィックであるように見えますが、青色の10.0.0.183はBlizzardゲームサービスとアクティブに通信しています。

![ゲームトラフィック](games.png)

パケット解析検索モジュールの使用法の詳細については、[パケット検索モジュールのドキュメント](#!search/packet/packet.md)を参照してください。

## ダッシュボード
ダッシュボードは、データの複数の側面を一度に表示する検索の集約ビューです。

「ダッシュボード」ページに移動し（左上のメニューを使用）、「+追加」ボタンをクリックして新しいダッシュボードを作成します。これを「SSH auth monitoring」と呼びます。次に、検索を追加します。この例では、先ほどのSSH認証検索を使用します。その検索を再実行し、結果画面から、右上の3ドットメニューを使用して[ダッシュボードに追加]を選択し、新しいダッシュボードを選択します。右下のポップアップで検索がダッシュボードに追加されたことが通知され、そのダッシュボードに移動するためのリンクが表示されますので、リンクをクリックします。

ダッシュボードには検索用のタイルが自動的に作成されているはずですが、そのサイズを変更することもできます。タイルのメニューから[タイルの編集]を選択すると、タイルの表示方法を変更できます。

### 動作中のダッシュボード
Gravwellの一般的な使用例の1つは、ネットワークアクティビティの追跡です。ここでは、アウトバウンドおよびインバウンドの帯域幅レート、wifi上のアクティブなMAC、Windowsネットワーキングイベント、および一般的なパケット頻度をレポートするダッシュボードを見てみます。このデータはすべて、pcap、netflow、およびWindowsイベントから抽出されています。

![ネットワークダッシュボード](network-dashboard.png)

アウトバウンドトラフィックチャートは、午前10時34分ごろ静かなシステムでかなり大きなスパイクを示しているため、概要チャートで短い時間枠を「ブラッシング」してズームインします。 デフォルトでは、1つの概要をズームすると、このダッシュボードに関連付けられている他の検索もズームされます。そのため、成功したログインをズームすると、その短い時間範囲が反映されて残りのグラフも更新されます。ズームインすると、アドレス10.0.0.57のスパイクを確認できます。さらに調査するには、事前に構築されたネットワーク調査ダッシュボードを使用できますが、これはこのクイックスタートの範囲外です。

![ネットワークダッシュボード、ズームイン](network-dashboard-zoomed.png)

## Kitインストール

Gravwellキットは、特定のデータソースを分析するための事前にパッケージ化されたツールセットです。キットは、Netflow v5、IPFIX、CoreDNSなどを分析するために存在します。これらのキットは、データの解析を始めるのに最適な方法であり、独自の解析を構築するためのジャンプオフの場となります。

ほとんどのキットはインジェスターの設定に依存していますが（例えば Netflow v5 キットは Netflowレコードを収集するために Netflowインジェスターを実行していることを想定しています）、*Weather* キットは完全に自己完結しています。このキットには、毎分ごとに実行され、指定した場所の天気データを取得するスクリプトが含まれています。

注：ウェザーキットを使用するには、[openweathermap.org](https://openweathermap.org)のAPIキーが必要です。APIキーの取得方法は[こちら](https://openweathermap.org/appid)をご覧ください。

メインメニューの "Kits"項目をクリックするとキットを見つけることができます。キットが何もインストールされていない場合、GUIは自動的に *利用可能な* キットのリストを表示します。

![](available-kits.png)

Weatherキットをインストールするには、Weatherキットのタイル上のデプロイアイコン（箱から出ている矢印）をクリックします。これにより、インストールウィザードが表示されます。最初のページでは、キットに含まれるアイテムが一覧表示され、内容を確認することができます。

![](kit-wizard1.png)

2ページ目には設定マクロが含まれています。これらはキットを設定するために使用します。最初のマクロに OpenWeatherMap API キーを入力し、2番目に監視する場所のリストを設定する必要があります。3番目のマクロでは、使用する単位を制御し、デフォルトのままにするか、「メートル法」に変更することができます。設定マクロのフィールドに値を入力するときは、まず「カスタム値を入力」リンクをクリックして、特定の検証ルールをオフにします。

注：場所のリストは、[ここ](https://openweathermap.org/current#one)に記載されているように、コロンで区切られた場所のリストで構成されている必要があります。複数の国が米国と同じ郵便番号フォーマットを使用しているため、"87110,us"と指定する方が通常は"87110"よりも良いことに注意してください。

![](kit-wizard2.1.png)

![](kit-wizard2.2.png)

![](kit-wizard2.3.png)

設定マクロの設定が終わったら、ウィザードの最終ページの "Next"をクリックします。これでキットのインストールに関連した最終的なオプションがいくつか出てきます。

![](kit-wizard3.png)

キットがインストールされると、インストールされたキットのリストが表示され、新しくインストールされたWeatherキットが表示されます。

![](kit-list.png)

キットに含まれているスクリプトは、すぐに天気データの取り込みを開始します。1、2分後には、いくつかのデータが取得されるはずなので、メインメニューをクリックしてダッシュボードページを開き、"Weather Overview" ダッシュボードをクリックしてください。気温チャートはまだ少ししか表示されないはずですが、少なくとも左下の「現在の状況」の表を見ることができるはずです。

![](current-conditions.png)

1日ほどすると、このような素敵なチャートが見られるほどのデータが集まってきます。

![](weather.png)

Gravwellキットの詳細は[キット](#!kits/kits.md)を参照してください。

## Gravwellの更新

Gravwellを問題なくアップグレードするために、私たちはインストールとアップグレードのプロセスが迅速かつ容易になるよう細心の注意を払っています。アップグレードのプロセスは、元々のインストール方法によって異なります。 Debian のようなパッケージリポジトリを使用している場合、Gravwell は他のアプリケーションと同様にアップグレードできます。


```
apt update
apt upgrade
```


元のインストール方法が自己完結型シェルインストーラの場合は、最新バージョンのインストーラをダウンロードして実行するだけです。 自己完結型インストーラは、インストールとアップグレードの両方のシステムとして機能し、既存のインストールを検出し、アップグレードに適用されないステップをスキップします。


### アップグレードのヒント

いくつかのインストールで役立つアップグレードのヒントがあります。

* クラスタ構成ではインデクサーのアップグレードを数珠つなぎに行い、アップグレード中にインジェスターが通常の操作を継続できるようにすべきです。
  * 同じことが分散型ウェブサーバーの構成に当てはまり、ロードバランサーは、必要に応じてユーザーをシフトします
* 可能であれば、大規模な自動化スクリプトジョブが実行されていないときに、検索エージェントのアップグレードを行います。
* ディストリビューションパッケージマネージャが設定ファイルの変更を要求してくることがありますが、 *既存の*設定を維持するようにしてください。
  * 何が変わったかを確認するのはいいのですが、通常は新機能のための設定を追加しているだけです。
  * 新たな設定ファイルにすると、設定を上書きしてコンポーネントの障害を引き起こす可能性があります。

アップグレード後、すべてのインデクサーが表示されていること、インジェスターが再接続され、期待されたバージョン番号が表示されていることを確認し、Gravwellの状態をチェックしてください。

![インジェスターのステータス](ingesters.png)


## 高度なトピック

### クラッシュレポートとメトリクス

Gravwellソフトウェアには、自動化されたクラッシュレポートとメトリクスレポートが組み込まれています。Gravwellに送られてくるものの詳細については、[クラッシュレポートとメトリクス](#!metrics.md)を参照してください。


### クラスター構成

マルチノードライセンスを持つユーザーは、複数のインデクサーとウェブサーバーのインスタンスを展開し、ネットワーク上でそれらを調整できます。 このドキュメントで基本的な手順の概要を説明しますが、このようなセットアップを展開する前に、Gravwellのサポートチームと調整することを強くお勧めします。

ほとんどのユースケースでは、単一のウェブサーバーと複数のインデクサーノードが望ましいでしょう。 簡単にするために、ウェブサーバーがインデクサーの1つと同じノードに存在する環境について説明します。

最初に、ヘッドノードとなるシステムで、上記のシングルノードGravwellインストールを実行します。 これにより、ウェブサーバーとインデクサーがインストールされ、認証シークレットが生成されます。

```
root@headnode# bash gravwell_installer.sh
```

次に、`/opt/gravwell/etc/gravwell.conf`のコピーを別の場所に作成し、"Indexer-UUID"で始まる行を削除します。 このgravwell.confファイルとインストーラーを各インデクサーノードにコピーします。 インデクサーノードで、追加の引数をインストーラーに渡して、ウェブサーバーのインストールを無効にし、新しいファイルを生成するのではなく、既存のgravwell.confファイルを使用するように指定します。

```
root@indexer0# bash gravwell_installer.sh --no-webserver --no-searchagent --use-config /root/gravwell.conf
```

インデクサーノードごとにこのプロセスを繰り返します。

インストールの最後のステップは、これらすべてのインデクサーをウェブサーバーに通知することです。 *ヘッドノード*で、`/opt/gravwell/etc/gravwell.conf`を開き、 'Remote-Indexers'行を見つけます。`Remote-Indexers=net:127.0.0.1:9404`のようになります。 次に、その行を複製し、他のインデクサーを指すようにIPを変更します（IPアドレスまたはホスト名を指定できます）。 たとえば、ローカルマシンとindexer0.example.net、indexer1.example.net、indexer2.example.netという3つの他のマシンにインデクサーがある場合、設定ファイルには次の行が含まれている必要があります。

```
Remote-Indexers=net:127.0.0.1:9404
Remote-Indexers=net:indexer0.example.net:9404
Remote-Indexers=net:indexer1.example.net:9404
Remote-Indexers=net:indexer2.example.net:9404
```

コマンド`systemctl restart gravwell_webserver`でウェブサーバーを再起動します。「システム統計」ページを表示して「ハードウェア」タブをクリックすると、4つのインデクサープロセスのそれぞれのエントリが表示されるはずです。
