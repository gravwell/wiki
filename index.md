# 

![](logo-name.png)

# Gravwell

このサイトには、Gravwellのドキュメントと、その他にChangelogsなどのリソースが含まれています。

Gravwellを使い始めたばかりの方は、まず[クイックスタート](quickstart/quickstart.md)を読んでから、[パイプライン検索](search/search.md)を読んで詳細を学ぶことをお勧めします。

Gravwellは、無料の[コミュニティ版](https://www.gravwell.io/download)を発表できて嬉しく思います。

## クイックスタートとダウンロード

  * [クイックスタート](quickstart/quickstart.md)

  * [ダウンロード](quickstart/downloads.md)

## Gravwellで検索

  * [検索の概要](search/search.md)

  * [抽出モジュール](search/extractionmodules.md)

  * [処理モジュール](search/processingmodules.md)

  * [レンダーモジュール](search/rendermodules.md)

  * [全パイプラインモジュールのアルファベット順リスト](search/complete-module-list.md)

## システムアーキテクチャ

  * [Gravwellシステムアーキテクチャ](architecture/architecture.md)

    * [Gravwellが使用するネットワークポート](configuration/networking.md)


  * [リソースシステム](resources/resources.md)

## インジェスターの設定 : Gravwellへのデータ取り込み

  * [インジェスターの設定](ingesters/ingesters.md)

    * [ファイルフォロワーインジェスター](ingesters/file_follow.md)

    * [シンプルリレーインジェスター](ingesters/simple_relay.md)
    
    * [Windows イベントインジェスター](ingesters/ingesters.md#Windows_Event_Service)

    * [Netflow/IPFIX インジェスター](ingesters/ingesters.md#Netflow_Ingester)

    * [Collectd インジェスター](ingesters/ingesters.md#collectd_Ingester)

  * [プリプロセッサー](ingesters/preprocessors/preprocessors.md)

  * [カスタムタイムフォーマット](ingesters/customtime/customtime.md)

  * [サービスの統合](ingesters/integrations.md)

## 高度なGravwellのインストールと設定

  * [Gravwellのインストールと設定](configuration/configuration.md)

  * [Dockerデプロイメント](configuration/docker.md)

  * [TLS/HTTPSの設定](configuration/certificates.md)

  * [Gravwellクラスター](distributed/cluster.md)

  * [分散Gravwellウェブサーバー](distributed/frontend.md)

    * [Gravwellオーバーウォッチ](distributed/overwatch.md)


  * [環境変数](configuration/environment-variables.md)

  * [詳細設定パラメータ](configuration/parameters.md)

  * [シングルサインオン](configuration/sso.md)

  * [Gravwellの堅牢化](configuration/hardening.md)

  * [一般的な問題と警告](configuration/caveats.md)

## クエリの高速化、自動抽出、データ管理
  
  * [自動抽出器の設定](configuration/autoextractors.md)
  
  * [クエリの高速化（インデックス化とブルームフィルタ）](configuration/accelerators.md)

  * [データ複製](configuration/replication.md)

  * [データエイジアウト](configuration/ageout.md)

  * [データ圧縮](configuration/compression.md)

  * [データアーカイブ](configuration/archive.md)

## 自動化

  * [スケジュール検索とスクリプト](scripting/scheduledsearch.md)

    * [自動化スクリプトのAPIと例](scripting/scriptingsearch.md)


  * [スクリプトの概要](scripting/scripting.md)

	* [Ankoモジュール](scripting/anko.md)

	* [Evalモジュール](scripting/eval.md)

## ユーザーインターフェース

  * [Gravwell Web GUI](gui/gui.md)

    * [検索インターフェイス](gui/queries/queries.md)

    * [ラベルとフィルタリング](gui/labels/labels.md)

	* [キット](kits/kits.md)

  * [コマンドラインクライアント](cli/cli.md)

## API

  * [API](api/api.md)

## その他

  * [ライセンス](license/license.md)

  * [メトリクスとクラッシュレポート](metrics.md)

  * [Changelogs](changelog/list.md)

  * [Gravwell EULA](eula.md)

  * [オープンソースライセンス](open_source.md)

Documentation version 2.0
