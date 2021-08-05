# セッションインジェスター

セッションインジェスターは、より大きな単一のレコードをインジェストするために使用される特殊なツールです。 インジェスターは指定されたポートで待機し、クライアントからの接続を受信すると、受信したデータを1つのエントリに集約します。

これにより、Windowsのすべての実行ファイルをインデックス化するなどの動作が可能になります:

```
for i in `ls /path/to/windows/exes`; do cat $i | nc 192.168.1.1 7777 ; done
```

## 基本設定

セッション・インジェスターは、永続的な設定ファイルではなく、コマンドライン・パラメーターで動作します。

```
Usage of ./session:
  -bind string
        リッスンするIPとポートを指定するバインド文字列 (デフォルトでは "0.0.0.0:7777")
  -clear-conns string
        コンマで区切られたクリアテキストターゲットのサーバー：ポートのリスト
  -ingest-secret string
        インジェストキー (デフォルトでは "IngestSecrets")
  -max-session-mb int
        1つのセッションが受け入れる最大MB数 (デフォルトでは 8)
  -pipe-conns string
        名前付きパイプ接続のパスのコンマ区切りリスト
  -tag-name string
        インジェストデータのタグ名 (デフォルトでは "default")
  -timeout int
        接続タイムアウト(秒) (デフォルトでは 1)
  -tls-conns string
        TLS接続のコンマ区切りのサーバー：ポートリスト
  -tls-private-key string
        TLS秘密鍵のパス
  -tls-public-key string
        TLS公開鍵のパス
  -tls-remote-verify string
        照合するリモート公開鍵のパス
```
