# Search websocket

WebソケットのURL：/api/ws/search

このページでは、検索用のWebsocketプロトコルについて説明します。「grep foo」検索を開始したときにクライアントとサーバーとの間で転送されるJSONの完全な例は、[Websocket Search Example](websocket-search-example.md)ページで見ることができます。


## Ping/Pongキープアライブ

検索ウェブソケットは、クエリの確認、検索の送信、および検索結果と検索統計の受信に使用されます。 検索ウェブソケットは、RoutingWebsocketシステムを使用してメッセージの「タイプ」を処理することを想定しています。 `/api/ws/search` は、起動時に次のメッセージ「サブタイプ」または「タイプ」が登録されていることを期待しています：PONG、parse、search、attach。


<span style="color:red; ">注：メッセージの「タイプ」は、従来の命名法のために「SubProto」と呼ばれることがあります。これは将来変更される可能性がありますが、このAPIに対して開発している場合、 "SubProto"はRFC Websocketサブプロトコル仕様ではなくメッセージと共に送信される "type"値を指すことに注意してください。</span>

PONGタイプはキープアライブシステムであり、クライアントは定期的にPINGPONG要求を送信する必要があります。

これは、ユーザーが検索プロンプトに座っている場合や、後ろへの接続が正常であるかどうかをユーザーに知らせることができる場合に使用できます。WebSocket自体が生きているかどうかを調べるためだけに調べることができるので、これはまったく必要ないかもしれません。

## 構文解析検索

"parse" WebSocketタイプは、検索バックエンドを呼び出さずにクエリの有効性を迅速にテストするために使用されます。

有効なクエリと応答を含むリクエストの例は、次のJSONになります。

フロントエンドからのリクエスト：
```json
{
        SearchString: "tag=apache grep firefox | regex "Firefox(<version>[0-9]+) .+" | count by version""

}
```

バックエンドからの応答：
```
{
        GoodQuery: true,
        ParseQuery: "tag=apache grep firefox | regex "Firefox(<version>[0-9]+) .+" | count by version"",
        ModuleIndex: 0,
}
```

無効なクエリと応答を含むリクエストの例では、次のようなJSONになります。

フロントエンドからのリクエスト：
```
{
        SearchString: "tag=apache grep firefox | MakeRainbows",
}
```

バックエンドから応答する
```
{
        GoodQuery: false,
        ParseError: "ModuleError: MakeRainbows is not a valid module",
        ModuleIndex: 1,
}
```

## 検索を開始する
すべての検索はWebSocketを介して開始され、「parse」、「PONG」、「search」、および「attach」の各サブタイプが開始時に要求される必要があります。

これは、websocketの確立時に次のJSONを送信することによって行われます。
```
{"Subs":["PONG","parse","search","attach"]}
```


SearchStringメンバには、検索を実行する実際のクエリを含める必要があります。

SearchStartとSearchEndは、クエリが動作する時間範囲です。時間範囲は、 "2006-01-02T15：04：05.999999999Z07：00"のように見えるRFC3339Nano形式でフォーマットする必要があります。

適切なクエリを含む検索リクエストの例には、次のJSONが含まれます。
```
{
       SearchString: "tag=apache grep firefox | nosort",
       SearchStart:  "2015-01-01T12:01:00.0Z07:00",
       SearchEnd:    "2015-01-01T12:01:30.0Z07:00",
       Background:   false,
}
```

//検索がクールな場合、//サーバーはyay / nayと新しいサブタイプを応答します
// // searchStartとsearchEndは、RFC 3339ナノ形式の文字列にする必要があります。

適切なクエリに対する応答には、次のJSONが含まれます。
```
{
        SearchString: "tag=apache grep firefox | nosort",
        RenderModule: "text",
        RenderCmd:    "text",
        OutputSearchSubproto:  "searchSDF8973",
        OutputStatsSubproto:   "statsSDF8973",
        SearchID:              "skdlfjs9098",
		SearchStartRange:      "2015-01-01T12:01:00.0Z07:00",
        SearchEndRange:        "2015-01-01T12:01:30.0Z07:00",
        Background:            false,
}
```

エラーが発生した場合、JSONレスポンスは次のようになります。
```
{
        Error: "Search error: The parameter "ChuckTesta" is invalid",
}
```

良い検索要求応答では、クライアントは検索ACKで応答しなければなりません。Ackは真偽のどちらかで返答しなければなりません。フロントエンドが理解できないレンダリングモジュールをバックエンドが要求したときに誤った応答が使用されることがあります。これはフロントエンドとバックエンドの間にバージョンの不一致があるときに起こることがあります。

次のJSONは、前の応答例に対する肯定的なACKを表します。
```
{
       Ok: True,
       OutputSearchSubproto: "searchSDF8973"
}
```

ACKが送信された後、バックエンドは検索を起動し、新しいサブタイプに関する検索結果の提供を開始します。元の検索、解析、およびPONGサブタイプはアクティブのままであり、フロントエンドが新しいクエリをチェックしたり、追加の検索を開始したりするために使用できます。ただし、アクティブなクエリとのやり取りはすべて、新しくネゴシエートされた検索固有のサブタイプを介して行われる必要があります。

## ノート
すべての検索は完全に非同期ですが、検索をバックグラウンド状態にすることを要求せずにクライアントが切断したり接続がクラッシュしたりすると、アクティブ検索は終了し、データはガベージコレクションされます。これはリソースの枯渇を防ぐためです。ユーザーはバックグラウンド検索を明示的に要求する必要があります。

検索は複数の消費者を持つことができます。たとえば、Bobが検索を開始し、Janetがそれにアタッチして結果を見ることがあります。非バックグラウンド検索は、すべてのコンシューマが切断された場合にのみ終了してクリーンアップします。そのため、Bobが検索を開始してJanetがアタッチしても、Bobがブラウザから離れたりブラウザを閉じたりした場合、検索は終了しません。ジャネットはそれと対話し続けることができます。ただし、Janetも自分のブラウザにアクセスするかブラウザを閉じると、検索は終了し、ガベージコレクトされます。

## アクティブ検索中の統計出力

統計は統計IDを介して要求されます

## 要求/応答IDの参照

要求および応答IDコードのリストは以下のとおりです。
```
{
    req: {
        REQ_CLOSE: 0x1,
        REQ_ENTRY_COUNT: 0x3,
        REQ_DETAILS: 0x4,
        REQ_TAGS: 0x5,
        REQ_STATS_SIZE: 0x7F000001, //gets backend "size" value of stats chunks. never used
        REQ_STATS_RANGE: 0x7F000002, //gets current time range covered by stats. rarely used
        REQ_STATS_GET: 0x7F000003, //gets stats sets over all time. may be used initially
        REQ_STATS_GET_RANGE: 0x7F000004, //gets stats in a specific range
        REQ_STATS_GET_SUMMARY: 0x7F000005, //gets stats summary for entire results
        REQ_STATS_GET_LOCATION: 0x7F000006, //get current timestamp for search progress
        REQ_GET_ENTRIES: 0x10, //1048578
        REQ_STREAMING: 0x11,
        REQ_TS_RANGE: 0x12,
	REQ_GET_EXPLORE_ENTRIES: 0xf010,
	REQ_EXPLORE_TS_RANGE: 0xf012,
        SEARCH_CTRL_CMD_DELETE: 'delete',
        SEARCH_CTRL_CMD_ARCHIVE: 'archive',
        SEARCH_CTRL_CMD_BACKGROUND: 'background',
        SEARCH_CTRL_CMD_STATUS: 'status'
    },
    rep: {
        RESP_CLOSE: 0x1,
        RESP_ENTRY_COUNT: 0x3,
        RESP_DETAILS: 0x4,
        RESP_TAGS: 0x5,
        RESP_STATS_SIZE: 0x7F000001, //2130706433
        RESP_STATS_RANGE: 0x7F000002, //2130706434
        RESP_STATS_GET: 0x7F000003, //2130706435
        RESP_STATS_GET_RANGE: 0x7F000004, //2130706436
        RESP_STATS_GET_SUMMARY: 0x7F000005,
        RESP_STATS_GET_LOCATION: 0x7F000006, //2130706438
        RESP_GET_ENTRIES: 0x10,
        RESP_STREAMING: 0x11,
        RESP_TS_RANGE: 0x12,
	RESP_GET_EXPLORE_ENTRIES: 0xf010,
	RESP_EXPLORE_TS_RANGE: 0xf012,
        RESP_ERROR: 0xFFFFFFFF,
        SEARCH_CTRL_CMD_DELETE: 'delete',
        SEARCH_CTRL_CMD_ARCHIVE: 'archive',
        SEARCH_CTRL_CMD_BACKGROUND: 'background',
        SEARCH_CTRL_CMD_STATUS: 'status'
    }
}
```
