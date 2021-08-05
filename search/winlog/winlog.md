# winlog

winlogプロセッサは、XML形式のWindowsログデータのために特別に用意された抽出器です。より一般的な[xmlモジュール](/#!search/xml/xml.md)を必要とするのではなく、 Windowsログエントリから多くの一般的なフィールドを抽出するためのショートカットを提供します。

## サポートされているオプション

* `-e`: 「-e」オプションは、winlogモジュールが列挙された値で操作することを指定します。 列挙された値で操作することは、アップストリームモジュールを使用してログエントリを抽出している場合に便利です。
* `-or`: 「-or」フラグは、抽出やフィルタが成功した場合に、 winlogモジュールがエントリーをパイプラインの下で継続することを許可することを指定します。

## 演算子

各winlogフィールドは、高速フィルタとして動作する演算子のセットをサポートしています。winlogモジュールの場合、すべてのフィールドは文字列として抽出されるので、文字列フィルタのみが利用可能です。

| 演算子 | 名前 | 説明 |
|----------|------|-------------|
| == | 等しい | フィールドは等しくなければなりません
| != | 等しくない | フィールドは等しくてはいけません
| ~ | サブセット | フィールドはメンバーでなければなりません
| !~ | サブセットではない | フィールドはメンバーであってはいけません

## データ・フィールド

この形式のログエントリが与えられます。

```
<Event xmlns="http://schemas.microsoft.com/win/2004/08/events/event">
  <System>
    <Provider Name="Microsoft-Windows-Security-Auditing" Guid="{543496D5-5478-49A4-A5BA-3E3B0428E31D}"/>
    <EventID>4689</EventID>
    <Version>0</Version>
    <Level>0</Level>
    <Task>13313</Task>
    <Opcode>0</Opcode>
    <Keywords>0x8020000000000000</Keywords>
    <TimeCreated SystemTime="2018-11-26T20:42:07.323695200Z"/>
    <EventRecordID>1624709</EventRecordID>
    <Correlation/>
    <Execution ProcessID="4" ThreadID="4392"/>
    <Channel>Security</Channel>
    <Computer>MY-PC</Computer>
    <Security/>
  </System>
  <EventData>
    <Data Name="SubjectUserSid">S-1-2-14</Data>
    <Data Name="SubjectUserName">GRAVUSER$</Data>
    <Data Name="SubjectDomainName">WORKGROUP</Data>
    <Data Name="SubjectLogonId">0x3e3</Data>
    <Data Name="Status">0x0</Data>
    <Data Name="ProcessId">0x1384</Data>
    <Data Name="ProcessName">C:\Windows\servicing\TrustedInstaller.exe</Data>
  </EventData>
</Event>
```

以下のフィールドを抽出することができます。

| Field | XML spec | Type | Filter Options |
|-------|----------|------|----------------|
| System | Event.System | bytes | == != ~ !~ |
| EventData | Event.EventData | bytes | == != ~ !~ |
| UserData | Event.UserData | bytes | == != ~ !~ |
| Provider | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderName | Event.System.Provider[Name] | bytes | == != ~ !~ |
| ProviderGUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| GUID | Event.System.Provider[Guid] | bytes | == != ~ !~ |
| EventID | Event.System.EventID | uint | == != < <= > >= |
| Version | Event.System.Version | uint | == != < <= > >= |
| Level | Event.System.Level | uint | == != < <= > >= |
| Task | Event.System.Task | uint | == != < <= > >= |
| Opcode | Event.System.Opcode | bytes | == != ~ !~ |
| Keywords | Event.System.Keywords | bytes | == != ~ !~ |
| TimeCreated | Event.System.TimeCreated[SystemTime] | bytes | == != ~ !~ |
| EventRecordID | Event.System.EventRecordID | uint | == != < <= > >= |
| ProcessID | Event.System.Execution[ProcessID] | uint | == != < <= > >= |
| ThreadID | Event.System.Execution[ThreadID] | uint | == != < <= > >= |
| Channel | Event.System.Channel | bytes | == != ~ !~ |
| Computer | Event.System.Computer | bytes | == != ~ !~ |
| Correlation | Event.System.Correlation | bytes | == != ~ !~ |
| UserID | Event.System.Security[UserID] | uint | == != < <= > >= |

上記以外のフィールドを指定すると、winlogモジュールは `Event.Data[Name]==<field>` の抽出を試みます。例えば、上の例のSubjectLogonId(0x3e3)は、単に `SubjectLogonId` をwinlogモジュールに指定するだけで抽出できます。

### データ・フィールドのフィルタリング

Windows ログ内の抽出可能なすべてのフィールドは、フィールドのタイプに応じてさまざまな比較操作によるインラインフィルタリングをサポートしています。 イベントシステムフィールドの一部は整数であり、`EventID`、`Version`、および `Level` フィールドのように整数として比較することができます。 その他のシステムフィールドとすべてのデータフィールドはバイト配列として扱われます。 ほとんどのGravwell検索モジュールと同様に、列挙された値が生成された直後にフィルタリングを実行した方が、ほとんどの場合速くなります。


## 例

以下の例は、上記のサンプルログを参照しています。

process ID (4) とユーザー名 (GRAVUSER$) を抽出します。

```
winlog ProcessID SubjectUserName
```

Security チャンネル上の EventID == 4689 のイベントのみからプロセス名を抽出するには、以下のようにします。

```
winlog EventID==4689 Channel==Security ProcessName
```
