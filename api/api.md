# API

このセクションでは、GUIと「フロントエンド」Webサーバー間で使用されるWebAPIについて説明します。

APIの大部分はRESTfulです。 このルールの例外は、検索からのデータの起動と監視に関連するデータ交換と転送の性質により、WebSocketを使用する検索APIです。

## 基本的なAPI

* [ログイン](login.md)
* [ユーザー設定](userprefs.md)
* [ユーザーアカウントコントロール](account.md)
* [ユーザーグループコントロール](groups.md)
* [通知](notifications.md)
* [検索コントロール](searchctrl.md)
* [検索結果のダウンロード](download.md)
* [検索履歴](searchhistory.md)
* [ロギング](loglevel.md)
* [エントリの取り込み](ingest.md)
* [その他のAPI](misc.md)
* [システムマネジメント](management.md)

## Gravwell内のオブジェクト

ユーザーが作成および変更できるさまざまなものがあります。それらのAPIはこのセクションにリストされています。

* [自動抽出機能](extractors.md)
* [ダッシュボード](dashboards.md)
* [キット](kits.md)
* [マクロ](macros.md)
* [プレイブック](playbooks.md)
* [リソース](resources.md)
* [スケジュールされた検索](scheduledsearches.md)
* [ライブラリの検索](searchlibrary.md)
* [テンプレート](templates.md)
* [ピボット（アクション可能）](pivots.md)
* [ユーザーファイル](userfiles.md)

## 検索と検索統計

[Websocketを検索](websocket-search.md)

[検索への再接続](websocket-search-attach.md)

[レンダラーとの対話](websocket-render.md)

## システム統計

システム統計は、通信にWebSocketも使用します。 これには、一般的なクラスターの状態を監視するために必要なすべての情報が含まれています。

[システム統計Websocket](websocket-stats.md)

他のいくつかの統計には、REST呼び出しを介してアクセスできます。

[REST Stats API](stats-json.md)

## テストAPI

システムには、_/api/test_ にあるテストAPIが含まれており、Webサーバーが動作しているかどうかをテストするために使用できます。テストAPIは完全に認証されておらず、常にStatusOK200と空の本文で応答します。
