# Gravwellコミュニティエディション

注意 このドキュメントは、[universal quickstart](#!quickstart/quickstart.md)に移行したため、非推奨となりました。既存のリンクの機能を維持するためにこの文書を残していますが、更新は行いません。

GravwellのCommunity Editionは、個人使用を目的とした無料のライセンスプログラムです。通常のGravwellのライセンスとは異なり、Community Editionのライセンスは、1日あたりの取り込みデータ量が2GBに制限されています。Gravwell Community Editionは、通常のGravwellライセンスとは異なり、1日あたりの取り込みデータ量が2GBに制限されているが、これはホームネットワークでの使用には十分すぎる量である（ただし、すべてのパケットをキャプチャしてNetflixのストリーミングを開始するような場合は除く）。

Gravwell Community Editionの入手方法は簡単だ。まず、Debianパッケージリポジトリからソフトウェアをインストールするか、Dockerコンテナを実行するか、ディストリビューションに依存しない自己完結型のインストーラを使用します。次に、電子メールで送られてくる無料ライセンスにサインアップします。最後に、新しくインストールされたGravwellインスタンスがライセンスファイルのアップロードを促します。

## ソフトウェアのインストール

Gravwell Community Editionは、Dockerコンテナ、ディストリビューションに依存しない自己解凍型インストーラ、Debianパッケージリポジトリの3つの方法で配布される。DebianやUbuntuを使用している場合はDebianリポジトリ、Dockerを設定している場合はDockerコンテナ、それ以外の場合は自己解凍型インストーラの使用をお勧めします。

### Debian リポジトリ

Debian リポジトリからのインストールは非常に簡単です。:

```
# Get our signing key
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
# Add the repository
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debian/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get install apt-transport-https
sudo apt-get update
# Install the package
sudo apt-get install gravwell
```

インストール時には、Gravwellの各コンポーネントが使用する共有秘密の値を設定するよう促されます。セキュリティのために、インストーラーがランダムな値を生成する（デフォルト）ようにすることを強くお勧めします。

![Read the EULA](eula.png)

![Accept the EULA](eula2.png)

![Generate secrets](secret-prompt.png)

### Dockerコンテナ

GravwellはDockerhub上でWebサーバとインデクサを含む単一のコンテナとして提供されています。GravwellをDockerにインストールする詳細な手順については、[Dockerインストール手順書](#!configuration/docker.md)を参照してください。

## 自己完結型のインストーラ

Debian 以外のシステムでは、[self-contained installer](https://update.gravwell.io/files/gravwell_2.2.4.sh)をダウンロードして検証します:

```
curl -O https://update.gravwell.io/files/gravwell_2.2.4.sh
md5sum gravwell_2.2.4.sh #should be f549d11ed30b1ca1f71a511e2454b07b
```

その後、インストーラーを実行します。:

```
sudo bash gravwell_2.2.4.sh
```

画面の指示に従って操作すると、Gravwellのインスタンスが起動します。

## ライセンスの取得

ライセンスファイルを入手するには、[https://www.gravwell.io/download](https://www.gravwell.io/download)にアクセスし、フォームに必要事項を記入してください。後日、ログボからメールでライセンスファイルが送られてきます。

Gravwellがインストールされたら、Webブラウザを起動してサーバにアクセスします（例：[https://localhost/](https://localhost/)）。ライセンスファイルをアップロードする画面が表示されます。

![ライセンスのアップロード](upload-license.png)

注意してください。Gravwellのデフォルトのユーザー名/パスワードはadmin/changemeです。早急にadminのパスワードを変更することをお勧めします。

## インジェストを開始しよう

インストールしたばかりのGravwellインスタンスはそれだけではつまらない。データを提供するためのインジェスターが必要になります。Debianリポジトリからインストールするか、[ダウンロードページ](downloads.md)から各インジェスターの自己解凍型インストーラーを取得します。

Debian リポジトリで利用可能なインジェスターは、`apt-cache search gravwell` を実行することで見ることができます:

```
root@debian:~# apt-cache search gravwell
gravwell - Gravwell community edition (gravwell.io)
gravwell-federator - Gravwell ingest federator
gravwell-file-follow - Gravwell file follow ingester
gravwell-netflow-capture - Gravwell netflow ingester
gravwell-network-capture - Gravwell packet ingester
gravwell-simple-relay - Gravwell simple relay ingester
```

Gravwellのメインインスタンスと同じノードにインストールすれば、自動的にインデクサーに接続するように設定されるはずですが、ほとんどの場合はデータソースを設定する必要があります。これについては[ingester configuration documents](#!ingesters/ingesters.md)を参照してください。

最初の実験として、File Followインゲスター(gravwell-file-follow)をインストールすることを強くお勧めします。これは、Linuxのログファイルを取り込むように事前に設定されているので、`tag=auth`のような検索を行うことで、いくつかのエントリーをすぐに見ることができるはずです。

![Auth entries](auth.png)

debian ベースのリポジトリを使用していない場合は、[downloads section](downloads.md)に自己完結型のインストーラーがあります。

### インゲスターの設定

各インゲスターのインストールと設定についての詳細は、[Setting Up Ingesters](/ingesters/ingesters.md)セクションを参照してください。

## 次のステップ

Gravwellはパワフルで複雑な製品です。専門知識を身につけるには時間がかかりますが、簡単なクエリから始めて、必要に応じてより複雑な概念を調べることで、有用な質問にすぐに答えられるようになります。

始めるためのアイデアとして、[標準版クイックスタートドキュメント](quickstart.md#Feeding_Data)の継続セクション、特に[Searchingセクション](quickstart.md#Searching)から始めることをお勧めします。欲しいデータをシステムに取り込むために、[ingester configuration documents](#!ingesters/ingesters.md)を参照する必要があるかもしれません。

検索モジュール](#!search/searchmodules.md)と[レンダーモジュール](#!search/rendermodules.md)には、たくさんの例と各モジュールのオプションについての詳細な説明があります。

最後に、[Gravwell blog](https://www.gravwell.io/blog)には、Gravwellを実際に使用したケーススタディや事例が掲載されており、インスピレーションを得ることができます。

サポートが必要な場合は、[オープンコミュニティ on Keybase](https://keybase.io/team/gravwell.community)に参加するか、support@gravwell.io にメールをお送りください。私たちは、Gravwellを使って他の人がデータから価値を得るお手伝いができることを楽しみにしています。
