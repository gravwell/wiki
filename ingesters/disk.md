# ディスクモニター

ディスクモニターインジェスターは、ディスクアクティビティのサンプルを定期的に採取し、Gravwellに送信するように設計されています。 ディスクモニターは、ストレージのレイテンシーの問題、ディスク障害の危険性、その他の潜在的なストレージの問題を特定するのに非常に役立ちます。 Gravwellでは、ディスクモニターを使って自社のストレージインフラを積極的に監視し、クエリがどのように動作しているかを調べたり、ストレージインフラの動作が悪いときにそれを特定したりしています。一例として、RAIDコントローラが診断ログで言及していなくても、レイテンシープロットでライトスルーモードに移行したRAIDアレイを特定できました。

ディスクモニターインジェスターは[github](https://github.com/gravwell/ingesters)で公開されています。

![diskmonitor](diskmonitor.png)

## 基本設定

ディスクモニターインジェスターは、[インジェスター](#!ingesters/ingesters.md#Global_Configuration_Parameters)で説明されている統一されたグローバル設定ブロックを使用します。 他の多くのGravwellインジェスターと同様に、ディスクモニターは複数のアップストリームインデクサー、TLS、クリアテキスト、名前付きパイプ接続、ローカルキャッシュ、ローカルロギングをサポートしています。

## セッションインジェスター

セッションインジェスターは、より大きな単一のレコードをインジェストするために使用される特殊なツールです。インジェスターは指定されたポートで待機し、クライアントからの接続を受信すると、受信したすべてのデータを1つのエントリに集約します。

これにより、すべてのWindows実行ファイルにインデックスを付けるなどの動作が可能になります。:

```
for i in `ls /path/to/windows/exes`; do cat $i | nc 192.168.1.1 7777 ; done
```

セッション・インジェスターは、永続的な設定ファイルではなく、コマンドライン・パラメーターによって駆動されます。

```
Usage of ./session:
  -bind string
        Bind string specifying optional IP and port to listen on (default "0.0.0.0:7777")
  -clear-conns string
        comma separated server:port list of cleartext targets
  -ingest-secret string
        Ingest key (default "IngestSecrets")
  -max-session-mb int
        Maximum MBs a single session will accept (default 8)
  -pipe-conns string
        comma separated list of paths for named pie connection
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
```

## 注

セッションインジェスターは正式にはサポートされておらず、インストーラーもありません。 セッションインジェスターのソースコードは[github](https://github.com/gravwell/ingesters)で公開されています。

