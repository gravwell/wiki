# マスファイルインジェスター

マスファイルインジェスターは、多数のソースからの多数のログのアーカイブを取り込むための、非常に強力だが特殊なツールです。

##基本構成

マスファイルインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。他の多くのGravwellインジェスターと同様に、マスファイルインジェスターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートします。

## 使用例

Gravwellのユーザーは、ネットワーク侵害の可能性を調査する際にこのツールを使用しました。そのユーザーは50以上の異なるサーバーからのApacheログを持っており、それらすべてを検索する必要がありました。それらを次々と取り込むと、一時的にインデックス作成のパフォーマンスが悪くなります。このツールは、ログエントリの時間的性質を維持しながらファイルをインジェストし、確かなパフォーマンスを確保するために作成されました。マスファイルインジェスターは、インジェストするマシンが、インジェストする前にソースログを最適化するのに十分なスペース（ストレージとメモリ）を持っている場合に、最も効果的に機能します。最適化フェーズは、インジェスト時と検索時のGravwellストレージシステムへの負担を軽減し、インシデントレスポンダーが迅速に行動し、パフォーマンスの高いログデータへのアクセスを短時間で実現できるようにします。

## メモ

マスファイルインジェスターは、コマンドラインのパラメータで動作し、サービスとして実行するようには設計されていません。 コードは[Github](https://github.com/gravwell/ingesters)で公開されています。

```
Usage of ./massFile:
  -clear-conns string
        comma separated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -no-ingest
        Optimize logs but do not perform ingest
  -pipe-conn string
        Path to pipe connection
  -s string
        Source directory containing log files
  -skip-op
        Assume working directory already has optimized logs
  -tag-name string
        Tag name for ingested data (default "default")
  -timeout int
        Connection timeout in seconds (default 1)
  -tls-conns string
        comma separated server:port list of TLS connections
  -tls-private-key string
        Path to TLS private key
  -tls-public-key string
        Path to TLS public key
  -tls-remote-verify string
        Path to remote public key to verify against
  -w string
        Working directory for optimization
```
