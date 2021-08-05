# Gravwellのネットワークに関する考察

Gravwellは、分散配置されたコンポーネント間の通信にいくつかのネットワークポートを使用します。この記事では、どのポートがどの目的で使用されているかを説明します。

## インデクサー制御用ポート: TCP 9404

このポートは gravwell.conf の `Control-Port` オプションで設定され、ウェブサーバーがインデクサーと通信するために使用します。全ての*インデクサー*上のファイアウォールが *ウェブサーバー* からのこのポートでの着信接続を許可していること、ウェブサーバーと全てのインデクサー間のネットワークインフラストラクチャがこのポートをブロックしていないことを確認してください。

## ウェブサーバー用ポート: TCP 80/443

このポートは、GravwellユーザーがGravwellウェブサーバーにアクセスするためのものです。デフォルトの設定では、gravwell.confの`Web-Port`オプションで指定された80番ポートで暗号化されていないHTTPを使用します。これは必要に応じて別の値、例えば8080に変更することができます。[TLS 証明書をインストール](#!configuration/certificates.md)する場合は、`Web-Port`で指定するウェブサーバー用ポートを 443 に変更することを推奨します。

## 平文通信のインジェスト用ポート: TCP 4023

このポートはインジェスターがインデクサーに接続し、非暗号化通信によってエントリをアップロードするために使用されます。デフォルトのポートはTCP 4023ですが、gravwell.confの`Ingest-Port`オプションを使って変更することができます。インジェスターとインデクサーは完全に異なるネットワーク上にあることが多いので、*インジェスター*からの*インデクサー*上のこのポートへの接続を許可されるようにファイアウォールが設定されていることが不可欠です。

## TLS通信のインジェスト用ポート: TCP 4024

このポートはインジェスターがインデクサーに接続し、TLS暗号通信によってエントリをアップロードするために使用されます。デフォルトのポートはTCP 4024ですが、gravwell.confの`TLS-Ingest-Port`オプションを使って変更することができます。インジェスターとインデクサーは完全に異なるネットワーク上にあることが多いので、*インジェスター*からの*インデクサー*上のこのポートへの接続を許可されるようにファイアウォールが設定されていることが不可欠です。

## インデクサーのレプリケーション用ポート: TCP 9606

このポートはインデクサーが[レプリケーション](#!configuration/replication.md)の際にお互いに通信するために使用されます。gravwell.confのレプリケーション部分の `Peer` と `Listen-Address` オプションで指定がなされていなければ、デフォルトのポートは9606です。インデクサーのみがこのポートを使用します。

## データストア用ポート: TCP 9405

このポートは、Gravwellクラスタで[複数のWebサーバー](#!distributed/frontend.md)を設定している場合に使用されます。*データストア*コンポーネントはこのポート(`Datastore-Port`オプションで指定)で*ウェブサーバー*から来る接続をリッスンします。

## RHEL (Redhat Enterprise Linux) や CentOS でのファイアウォールコマンド

RHEL/CentOS は独自のファイアウォールコマンドを使用します。利便性のために、ウェブサーバーとインデクサーコンポーネントとSimple Relayインジェスターのためにポートを開くために必要なコマンドを掲載します。ネットワークポートをリッスンするインジェスターは、これと同じ方法でポートを開く必要があることに注意してください。

注意: ここに示したコマンドは、*一時的に*ポートを開くだけです。システムを再起動するとルールがリセットされます。ルールの変更を恒久的なものにするには、`sudo firewall-cmd --runtime-to-permanent` を実行してください。

### インデクサーのポート

```
sudo firewall-cmd --zone=public --add-port=9404/tcp 
sudo firewall-cmd --zone=public --add-port=9405/tcp
sudo firewall-cmd --zone=public --add-port=4023/tcp
sudo firewall-cmd --zone=public --add-port=4024/tcp
```

### ウェブサーバーのポート

```
sudo firewall-cmd --zone=public --add-service=http
sudo firewall-cmd --zone=public --add-service=https
```

### Simple Relayインジェスターのポート

```
sudo firewall-cmd --zone=public --add-port=7777/tcp
sudo firewall-cmd --zone=public --add-port=601/tcp
sudo firewall-cmd --zone=public --add-port=514/udp
```