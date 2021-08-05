# 環境変数

インデクサ、ウェブサーバ、インジェスタの各コンポーネントでは、一部のパラメータを設定ファイルではなく環境変数で設定できるようになりました。これは、大規模なデプロイメントのために、より一般的な設定ファイルを作成するのに役立ちます。複数のディレクティブを含む設定変数は、カンマ区切りのリストを使って環境変数で設定します。例えば、Federatorの起動時にインジェストシークレットを指定するには、次のようにします。

```
GRAVWELL_INGEST_SECRET=MyIngestSecret /opt/gravwell/bin/gravwell_federator
```

環境変数名の最後に"_FILE "を付けると、Gravwellはその環境変数にファイルへのパスが含まれており、そのファイルに目的のデータが含まれているとみなします。これは、[Dockerの"secrets"機能](https://docs.docker.com/engine/swarm/secrets/)と組み合わせると特に便利です。

```
GRAVWELL_INGEST_AUTH_FILE=/run/secrets/ingest_secret /opt/gravwell/bin/gravwell_indexer
```

注：環境変数の値は、対応するフィールドが適切な設定ファイル（gravwell.confまたはインジェスターの設定ファイル）で**明示的に設定されていない場合にのみ**使用されます 。

## インデクサーとウェブサーバー

下の表は、`gravwell.conf`のどのパラメータが、インデクサとウェブサーバの環境変数として設定できるかを示しています。これらの変数は、パラメータが `gravwell.conf` で**設定されていない場合にのみ**使用できることに注意してください。

| gravwell.conf 内の変数 | 環境変数 | 例 |
|:------|:----|:---|----:|
| Ingest-Auth | GRAVWELL_INGEST_AUTH | GRAVWELL_INGEST_AUTH=CE58DD3F22422C2E348FCE56FABA131A |
| Control-Auth | GRAVWELL_CONTROL_AUTH | GRAVWELL_CONTROL_AUTH=C2018569D613932A6BBD62A03A101E84 |
| Indexer-UUID | GRAVWELL_INDEXER_UUID | GRAVWELL_INDEXER_UUID=a6bb4386-3433-11e8-bc0b-b7a5a01a3120 |
| Webserver-UUID | GRAVWELL_WEBSERVER_UUID | GRAVWELL_WEBSERVER_UUID=b3191f54-3433-11e8-a0c2-afbff4695836 |
| Remote-Indexers | GRAVWELL_REMOTE_INDEXERS | GRAVWELL_REMOTE_INDEXERS=172.20.0.1:9404,172.20.0.2:9404|
| Replication-Peers | GRAVWELL_REPLICATION_PEERS | GRAVWELL_REPLICATION_PEERS=172.20.0.1:9406,172.20.0.2:9406 |
| Datastore | GRAVWELL_DATASTORE | GRAVWELL_DATASTORE=172.20.0.10:9405 |

## インジェスター

インジェスターでも同様に、一部のパラメータを設定ファイルで明示的に設定するのではなく、環境変数として設定することができます。

| Configファイル内の変数 | 環境変数 | 例 |
|:------|:----|:---|
| Ingest-Secret | GRAVWELL_INGEST_SECRET | GRAVWELL_INGEST_SECRET=CE58DD3F22422C2E348FCE56FABA131A |
| Log-Level | GRAVWELL_LOG_LEVEL | GRAVWELL_LOG_LEVEL=DEBUG |
| Cleartext-Backend-target | GRAVWELL_CLEARTEXT_TARGETS | GRAVWELL_CLEARTEXT_TARGETS=172.20.0.1:4023,172.20.0.2:4023 |
| Encrypted-Backend-target | GRAVWELL_ENCRYPTED_TARGETS | GRAVWELL_ENCRYPTED_TARGETS=172.20.0.1:4024,172.20.0.2:4024 |
| Pipe-Backend-target | GRAVWELL_PIPE_TARGETS | GRAVWELL_PIPE_TARGETS=/opt/gravwell/comms/pipe |


### フェデレーター固有の変数

フェデレータによって多くのリスナーを実行できますが、それぞれのリスナーにはそれぞれ異なるインジェスト・シークレットが関連付けられているため、実行時にそれらのリスナーのインジェスト・シークレットを個別に設定するための特別な環境変数のセットを使います。

各リスナーには名前があります。以下の例では、リスナーの名前は"base"です:

```
[IngestListener "base"]
	Cleartext-Bind = 0.0.0.0:4023
	Tags=syslog
```

実行時にこの"base"という名前のリスナーのインジェストシークレットを指定するのには、変数`FEDERATOR_base_INGEST_SECRET`を使います:

```
FEDERATOR_base_INGEST_SECRET=SuperSecret /opt/gravwell/bin/gravwell_federator
```

あるいは、他の環境変数の場合と同様、設定用ファイルを指定することもできます:

```
FEDERATOR_base_INGEST_SECRET_FILE=/run/secrets/federator_base_secret /opt/gravwell/bin/gravwell_federator
```

### データストア固有の変数

[データストア](#!distributed/frontend.md) も、実行時に環境変数によって設定できます。

| gravwell.conf 内の変数 | 環境変数 | 例 |
|------------------------|----------------------|---------|
| Datastore-Listen-Address | GRAVWELL_DATASTORE_LISTEN_ADDRESS | GRAVWELL_DATASTORE_LISTEN_ADDRESS=192.168.1.100 |
| Datastore-Port | GRAVWELL_DATASTORE_LISTEN_PORT | GRAVWELL_DATASTORE_LISTEN_PORT=9995 |
