# インテグレーション

Gravwellインジェストフレームワークは、BSD 2-clauseライセンスでオープンソース化されており、オープンソース製品と商用製品の両方に直接組み込むことができます。 データプロセッサーやジェネレーターは、インジェストフレームワークを直接組み込み、Gravwellとの統合を簡単に設定することができます。 このドキュメントページでは、Gravwellの統合機能を紹介し、必要に応じてドキュメント化しています。

## CoreDNS

CoreDNSは、高度な設定が可能で、プラグインにも対応したDNSサーバーであり、DNSサービスのベースとなるプラットフォームを提供します。 CoreDNSの基本機能は、リレー、プロキシ、または本格的なDNSサーバーとして機能する、堅牢で性能の高いDNSサーバーを提供します。CoreDNSのライセンスはApache-2.0で、[github](https://github.com/coredns/coredns)から入手可能です。 CoreDNSの詳細については、[https://coredns.io](https://coredns.io)をご覧ください。

### Gravwell CoreDNS プラグイン

CoreDNSにはGravwellプラグインが用意されており、インジェストフレームワークをCoreDNSに直接組み込むことができます。このプラグインを使用すると、静的にコンパイルされた高性能なDNSサーバが、DNS監査データをGravwellインスタンスに直接送信することができます。プラグインはBSD 2-Clauseでライセンスされており、[github](https://github.com/gravwell/coredns)から入手可能です。

このプラグインは、信頼性の高いローカルキャッシュ、ロードバランシング、フェイルオーバーなど、通常の機能をすべてサポートする完全なインジェストシステムを提供します。 Gravwellプラグインについての詳細は、CoreDNS [外部プラグイン](https://coredns.io/explugins/gravwell/)のページを参照してください。

#### GravwellでCoreDNSを構築する

GravwellプラグインでCoreDNSを構築するには、Golangツールチェーンとコンパイラがインストールされている必要があります。詳細は[こちら](https://golang.org/)を参照してください。

```
go get github.com/coredns/coredns
pushd $GOPATH/src/github.com/coredns/coredns/
sed -i 's/metadata:metadata/metadata:metadata\ngravwell:github.com\/gravwell\/coredns/g' plugin.cfg
go generate
CGO_ENABLED=0 go build -o /tmp/coredns
popd
```

静的にコンパイルされたバイナリは、_/tmp/coredns_に格納されます。 CoreDNS は、有効な Corefile を最初の引数として与えることで起動できます。 非rootユーザーでCoreDNSを実行する場合は、バイナリに[service bind capability](https://wiki.apache.org/httpd/NonRootPortBinding)を与える必要があります。

```
setcap 'cap_net_bind_service=+ep' /tmp/coredns
```

#### Gravwellプラグインの設定

設定はCoreDNS Corefileを介して行われ、基本的な構文は**directive** **value**となります。 コメントの前には「#」文字が付きます。

設定可能なパラメータは以下のとおりです:

* **Ingest-Secret** は、インデクサーとの認証に使用するトークンを定義します。 **Ingest-Secret**は必須です。
* **Cleartext-Target** は、TCP接続を使用するリモートインデクサーのアドレスとポートを定義します。 IPv4とIPv6のアドレス、およびホスト名がサポートされています。
* **Ciphertext-Target** は、TLS 接続を使用するリモートインデクサーのアドレスとポートを定義します。 IPv4とIPv6のアドレス、およびホスト名がサポートされています。
* **Tag** は、DNS 監査ログに割り当てられるタグを指定します。 特殊文字やスペースを含まない任意の英数字を指定できます。 有効なTagの値が必要です。
* **Encoding** はDNS監査ログを転送する際の形式を指定します。オプションは_json_または_text_です。 デフォルトは_json_です。
* **Insecure-Novalidate-TLS** は、TLS接続時の証明書検証の有無を切り替えます。 デフォルトでは検証はオンになっています。
* **Log-Level** は、統合されたgravwellタグを使ったロギングの冗長性を指定します。 オプションは、_OFF_ _INFO_ _WARN_ _ERROR_ です。 デフォルトは _ERROR_ です。
* **Ingest-Cache-Path** は、インデクサの接続が失われたときに作動するキャッシュシステムのファイルパスを指定します。 パスは書き込み可能なファイルの絶対パスでなければなりません。
* **Max-Cache-Size-MB** は、キャッシュファイルの最大サイズをメガバイトで指定します。これはセーフティーネットとして使用されます。ゼロがデフォルトで、無制限を意味します。
* **Cache-Depth** は、メモリ内インジェスト・バッファのサイズをエントリ数で指定します。デフォルトは128です。
* **Cache-Mode** は、インデクサの接続の状態に基づいてバッキング・キャッシュの動作を指定します。デフォルトのモードは「always」です。


基本的なGravwellの設定は次のようになります。:

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Tag dns
    Encoding json
    Log-Level INFO
    #Cleartext-Target 192.168.1.2:4023 #second indexer
    #Ciphertext-Target 192.168.1.1:4024
    #Insecure-Novalidate-TLS true #disable TLS certificate validation
    #Ingest-Cache-Path /tmp/coredns_ingest.cache #enable the local ingest cache
    #Max-Cache-Size-MB 1024
}
~~~

DNSリスナーごとに固有のGravwellプラグインセクションを適用することができます。2つの異なるインターフェースをリッスンし、それぞれに固有のGravwell設定を適用するCorefileの例は以下のようになります:

~~~
.:53 {
  forward . 8.8.8.8:53 8.8.4.4:53 9.9.9.9:53
  errors stdout
  bind 172.20.0.1
  cache 240
  whoami
  gravwell {
   Ingest-Secret SecretSetOne
   Cleartext-Target 172.20.0.1:4023
   Tag dns
   Encoding json
  }
}

.:53 {
  forward	. tls://1.1.1.1
  errors stdout
  bind 192.168.1.1
  hosts
  cache 60s
  gravwell {
   Ingest-Secret SecretSetTwo
   Cleartext-Target 192.168.1.100:4023
   Cleartext-Target 192.168.1.101:4023
   Cleartext-Target 192.168.1.102:4023
   Tag dns
   Encoding json
  }
}
~~~

DNSリスナーのそれぞれを、完全に独立した2つのGravwellインストレーション（1つは単一のインデクサー、もう1つは分散クラスタ）に送信していることに注目してください。

暗号化されていない接続を介して単一のインデクサにDNSリクエストを送信するGravwell Corefileセクションのサンプルです。 ローカルキャッシュは無効です。

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Tag dns
  }
~~~

TLS接続を介して2つのインデクサにDNSリクエストを送信し、署名されていない証明書を受け付けるGravwell Corefileセクションのサンプルです。ローカルキャッシュは無効です。
IPv4とIPv6のアドレスはCleartextとCiphertextの両方のターゲットに対応しています。IPv6アドレスは括弧で囲む必要があります。

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Ciphertext-Target 192.168.1.1:4024
    Ciphertext-Target [fe80::dead:beef:feed:febe%p1p1]:4024 #connecting to link local address via device p1p1
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~

TLS接続を介して2つのインデクサにDNSリクエストを送信し、署名されていない証明書を受け付けるGravwell Corefileセクションのサンプルです。ローカルキャッシュは無効です。

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Ciphertext-Target 192.168.1.1:4024
    Ciphertext-Target [dead::beef]:4024
    Insecure-Novalidate-TLS true
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~

2つのインデクサーにDNSリクエストを送信し、インデクサーとの通信に失敗した場合はローカルキャッシュを有効にするGravwell Corefileセクションのサンプルです。最大1GBのデータをローカルにキャッシュすることができます。

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Ciphertext-Target 192.168.1.2:4024
    Insecure-Novalidate-TLS true
    Ingest-Cache-Path /tmp/coredns_ingest.cache
    Max-Cache-Size-MB 1024
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~
