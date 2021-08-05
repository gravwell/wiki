# カスタムDocker展開

ほとんどのGravwellコンポーネントは、最新のLinuxホストでの実行に適した静的にコンパイルされた実行可能ファイルとしてデプロイされます。これにより、Dockerへのデプロイも容易になります。 Gravwellのエンジニアは、開発、テスト、内部デプロイにdockerを幅広く使用しています。 また、Dockerを使用すると、大規模なGravwellのデプロイメントを迅速に立ち上げたり、解体したり、管理したりすることができます。 お客様がDockerを使ってすぐに導入できるように、クラスタ版とシングル版の両方のSKUのサンプルDockerfilesを提供しています。

Dockerfilesの完全なセットは [こちら](https://update.gravwell.io/files/docker_buildfiles_ad05723a547d31ee57ed8880bb5ef4e9.tar.bz2) にあり、ad5723a547d31ee57ed8880bb5ef4e9のMD5チェックサムがあります。

## Dockerコンテナーの構築

提供されているDockerfileを使ってDockerコンテナーを構築するのは非常に簡単です。 GravwellのDockerデプロイでは、非常に小さいベースコンテナであるbusyboxを使用しているため、非常に小さなコンテナにすることが可能です。

### Dockerfile

特定の展開要件に合わせてdockerファイルを変更するのに必要な作業はほとんどありません。 標準のgravwellドッカーファイルは、小さな起動スクリプトを使用して、起動時にX509証明書を再生成する必要があるかどうかを確認します。 展開に有効な証明書がある場合、Gravwellバイナリを直接起動し、インストーラーによって展開される/ opt / gravwell / binのユーティリティ（gencertおよびcrashreport）を削除できます。

単一インスタンスのベースDockerfile：
```
FROM busybox
MAINTAINER support@gravwell.io
ARG INSTALLER=gravwell_installer.sh
COPY $INSTALLER /tmp/installer.sh
COPY start.sh /tmp/start.sh
RUN /bin/sh /tmp/installer.sh --no-questions
RUN rm -f /tmp/installer.sh
RUN mv /tmp/start.sh /opt/gravwell/bin/
CMD ["/bin/sh", "/opt/gravwell/bin/start.sh"]
```

基本的な起動スクリプト：
```
#!/bin/sh

# unless environment variable says no, generate new SSL certs
if [ "$NO_SSL_GEN" == "true" ]; then
	echo "Skipping SSL certificate generation"
else
	/opt/gravwell/bin/gencert -key-file /opt/gravwell/etc/key.pem -cert-file /opt/gravwell/etc/cert.pem -host 127.0.0.1
	if [ "$?" != "0" ]; then
		echo "Failed to generate certificates"
		exit -1
	fi
fi

#fire up the indexer and webserver processes and wait
/opt/gravwell/bin/gravwell_indexer -config-override /opt/gravwell/etc/ &
/opt/gravwell/bin/gravwell_webserver -config-override /opt/gravwell/etc/ &
wait
```

## 環境変数を使用する

標準のdocker起動スクリプトは環境変数 `NO_SSL_GEN`を探し、「true」に設定されている場合はX509証明書の生成をスキップします。 デプロイメントが有効な証明書を注入している場合、コンテナを起動するときに引数 `-e NO_SSL_GEN = true`を必ず含めてください。

インデクサー、Webサーバー、およびインジェスターは、必要に応じて、構成ファイルではなく、環境変数を介して設定されるいくつかの構成パラメーターを持つこともできます。 詳細については、[環境変数](environment-variables.md)のドキュメントを参照してください。

## インジェスター用のサンプルDockerfile

Gravwellは新しいインジェスターとコンポーネントを継続的にリリースしていますが、すべての学期のDockerfileを常に持っているとは限りません。 ただし、Dockerfilesは非常に単純であり、簡単に変更できます。 以下は、SimpleRelay ingesterを介してDockerコンテナーを作成するDockerfileの例です。

```
FROM busybox
MAINTAINER support@gravwell.io
ARG INSTALLER=gravwell_installer.sh
COPY $INSTALLER /tmp/installer.sh
RUN /bin/sh /tmp/installer.sh --no-questions
RUN rm -f /tmp/installer.sh
CMD ["/opt/gravwell/bin/gravwell_simple_relay"]
```

コンテナを構築するには、単純なリレーインストーラーをDockerファイルと同じ作業ディレクトリにコピーし、次のコマンドを実行します。
```
docker build --ulimit nofile=32000:32000 --compress --build-arg INSTALLER=gravwell_simple_relay_installer_2.0.sh --no-cache --tag gravwell:simple_relay_2.0.0 .
```
