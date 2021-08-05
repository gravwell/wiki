# Gravwellシングルサインオン

GravwellのGUIは、SAMLを使用したシングルサインオンをサポートしています。 理論的には、SAML準拠のIDプロバイダーを使用するのであればログインできます。このページではGravwellのSSO構成オプションについて説明し、Windows AD FSサーバーで認証するようにGravwellを構成する方法の例を示します。

注：通常のユーザーがSSO経由でログインする場合でも、デフォルトの「admin」ユーザーはSSOを経由してのログインは行いません。 新しいシステムをセットアップするときは、SSOをすぐに設定する場合でも、必ず管理ユーザーのパスワードを変更してください。 Gravwell管理ユーザーは、必要に応じてGUIから新しい非SSOユーザーアカウントを作成できることにも注意してください。

## Gravwell SSO構成パラメーター

GravwellインスタンスでSSOを有効にするには、SSOセクションをウェブサーバーの`gravwell.conf`ファイルに追加する必要があります。 Windows AD FSサーバーで動作する最小限の構成を次に示します。

```
[SSO]
	Gravwell-Server-URL=https://gravwell.example.org
	Provider-Metadata-URL=https://sso.example.org/FederationMetadata/2007-06/FederationMetadata.xml
```

基本的なSSO構成パラメーターは次のとおりです。

* `Gravwell-Server-URL` (必須): SSOサーバーが認証した後、ユーザーがリダイレクトされるURLを指定します。 これは、Gravwellサーバーのユーザー向けのホスト名またはIPアドレスである必要があります。
* `Provider-Metadata-URL` (必須): SSOサーバーのXMLメタデータのURLを指定します。 上記のパス(`/FederationMetadata/2007-06/FederationMetadata.xml`) はAD FSサーバーで機能するはずですが、他のSSOプロバイダー用に調整が必要な場合があります。
* `Insecure-Skip-TLS-Verify` [デフォルト: false]: このパラメーターをtrueに設定した場合、GravwellはSSOサーバーとの通信時に無効なTLS証明書を無視します。このオプションの設定には注意してください。

以下は、SSOプロバイダーに基づいて調整する必要のある、より高度なパラメーターです。デフォルトは、Microsoft AD FSサーバーに適しています。

* `Username-Attribute` [デフォルト: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"]: ユーザー名を含むSAML属性を定義します。 Shibbolethサーバーでは、代わりに "uid" に設定する必要があります。
* `Common-Name-Attribute` [デフォルト: "http://schemas.xmlsoap.org/claims/CommonName"]: ユーザーの"共通名"を含むSAML属性を定義します。Shibbolethサーバーでは、代わりに"cn"に設定する必要があります。
* `Given-Name-Attribute` [デフォルト: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname"]: ユーザーの名を含むSAML属性を定義します。 Shibbolethサーバーでは、代わりに"givenName"に設定する必要があります。
* `Surname-Attribute` [デフォルト: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"]: ユーザーの姓を含むSAML属性を定義します。
* `Email-Attribute` [デフォルト: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"]: ユーザーのメールアドレスを含むSAML属性を定義します。 Shibbolethサーバーでは、代わりに"mail"に設定する必要があります。

ユーザーのログイン応答とともにグループメンバーシップのリストを受け取り、必要なグループを自動生成し、ユーザーをそれらのグループに追加するようにGravwellを設定することができます。これを有効にするには`Groups-Attribute`および、少なくとも1つの`Group-Mapping`を定義しなければなりません。

* `Groups-Attribute` [デフォルト: "http://schemas.microsoft.com/ws/2008/06/identity/claims/groups"]: ユーザが属するグループのリストを含む SAML 属性を定義します。通常、グループリストを送信するために、SSO プロバイダーを明示的に構成する必要があります。
* `Group-Mapping`: ユーザーのグループメンバーシップにリストされている場合に自動的に作成される可能性のあるグループの1つを定義します。これは、複数のグループを許可するために複数回指定することができます。引数はコロンで区切られた2つの名前で構成されなければなりません。1つ目はグループのSSOサーバーサイドの名前(通常はAD FSの名前、AzureのUUIDなど)で、2つ目はGravwellが使用すべき名前です。したがって、`Group-Mapping=Gravwell Users:gravwell-users`を定義すると、グループ「Gravwell Users」のメンバーであるユーザーのログイントークンを受け取ると、「gravwell-users」という名前のローカルグループを作成し、そこにユーザーを追加することになります。

## 例: Azure Active Directoryの設定

Azure Active DirectoryでのSSO設定については、専用ページに分かれています。[ここ](sso-azure/azure.md)で読むことができます。

## 例：Windows Server 2016のセットアップ

Gravwell SSOは、Windows Serverで提供されているMicrosoftのAD FS（Active Directory Federation Services）とうまく連携します。ここでは、SSO認証のためにAD FSとGravwellを設定する方法を説明します。

開始する前に、Active Directory と AD FS をサーバーにインストールしておく必要があります。これらのサービスの基本的なインストールとセットアップは、このドキュメントの範囲外です。SSOを設定する場合は、おそらくActive Directoryを既に構成してあると仮定しています。

重要: Gravwellで使用するユーザーアカウントは、Active Directoryで設定された電子メールアドレスを持っていなければなりません。これは、Gravwellの内部的なユーザー名として使用されます。EventID 364のイベントログにエラーが表示される場合、これが原因です。

### Gravwellをセットアップする

AD FSを構成するには、GravwellインスタンスのSSOメタデータファイルが必要です。したがって、最初にGravwellをセットアップします。 Gravwell GUIにSSOボタンが表示されますが、AD FSを構成するまで無効になります。SSOを有効にするには、Gravwell ウェブサーバー上でTLS証明書（自己署名またはその他）を設定する必要があります。 TLSの設定方法については、[ここ](certificates.md)を参照してください。

`gravwell.conf`を開き、`[Global]`セクションの下に`[SSO]`セクションを追加します。 AD FSサーバーが「sso.example.org」にあり、Gravwell ウェブサーバーが「gravwell.example.org」にある場合、設定は次のようになります。

```
[SSO]
	Gravwell-Server-URL=https://gravwell.example.org
	Provider-Metadata-URL=https://sso.example.org/FederationMetadata/2007-06/FederationMetadata.xml
```

何らかの理由でAD FSサーバーで自己署名証明書を使用している場合、セクションに`Insecure-Skip-TLS-Verify=true`を追加する必要があります。

Gravwell ウェブサーバーを再起動します (`systemctl restart gravwell_webserver.service`)。そのまましばらく待つと再表示されるはずです。表示されない場合は、設定に間違いがないかをチェックし、`/dev/shm/gravwell_webserver.service` および `/opt/gravwell/log/web/` でエラーがないか確認してください。

### 証明書利用者を追加

ここで、Gravwellからの認証要求を受け入れるようにAD FSサーバーを構成する必要があります。 AD FS管理ツールを開き、"証明書利用者信頼の追加"を選択します。

![](sso-trust1.png)

これにより、証明書利用者信頼の追加ウィザードが開きます。 最初に、Claims-AwareかNon-Claims-Awareのどちらのアプリケーションを使用するかを尋ねてきますので、Claims-Awareを選択し、Startをクリックします。

Gravwellサーバーに関するメタデータ情報をAD FSに取得するには、2つの方法のいずれかを選択する必要があります。Gravwellサーバーが適切に署名されたTLS証明書を使用している場合は、最初のオプションの「Federation metadata address」フィールドにURLを入力して「Next」をクリックするだけです。

![](sso-trust2.png)

ただし、Gravwellが自己署名証明書を使用している場合は、最初にメタデータを手動でダウンロードする必要があります。 Webブラウザを開き、`https://gravwell.example.org/saml/metadata`に移動します。プロンプトが表示されたらファイルを保存し、ウィザードに戻って2番目のオプションで適切なパスを設定します。

![](sso-trust3.png)

ウィザードの次のページで、表示名を設定するように求められます。 "Gravwell"またはそれに類似したものが良いでしょう。ウィザードの次のページでは、デフォルトのままにすることができます。

### 証明書利用者のクレーム発行ポリシーの編集

ここで、いくつかのクレーム発行変換ルールを依存ポリシーに追加する必要があります。新しく作成された証明書利用者の"クレーム発行ポリシーの編集"を選択します。

![](sso-trusts.png)

3つのルールを作成する必要があります。 最初に、[ルールの追加]をクリックしてウィザードを開きます。"LDAP属性をクレームとして送信"を選択し、[次へ]をクリックして、次のように入力します。

![](sso-ldap.png)

次に、別のルールを作成します。今回は、"着信クレームの変換"を選択し、以下に示すように入力します。

![](sso-transform.png)

最後に、"カスタムルールを使用してクレームを送信する"を選択し、次のテキストをフィールドに貼り付けて、別のルールを作成します。

```
c:[Type == "http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsaccountname", Issuer == "AD AUTHORITY"]
=> issue(store = "Active Directory",
types = ("http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname",
"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"),
query = ";mail,givenName,sn;{0}", param = c.Value);
```

![](sso-custom.png)

完了すると、次の3つのルールができているはずです。

![](sso-policy.png)

### (オプション) グループ情報の送信

Gravwellは、このドキュメントで前述したように、グループを自動的に作成し、SSOユーザーをこれらのグループに追加することができます。以下のスクリーンショットに示すように、「Token-Groups - Unqualified Names」を「Group」にマップするLDAP要求規則を作成することで、グループ名のリストを含むクレームを送信するようにActive Directoryを設定できます。

![](sso-groups.png)

gravwell.confでは、どの属性にグループのリストが含まれているかを示すために、`Groups-Attribute` フィールドを追加する必要があります(上記のように発信クレームタイプを「Group」に設定した場合は `http://schemas.microsoft.com/ws/2008/06/identity/claims/groups`)。また、Active Directoryグループ名をGravwell内の希望のグループ名にマッピングするために、少なくとも1つの `Groups-Mapping`フィールドが必要です。以下の例では、「Gravwell Users」という名前のADグループを「gravwell-users」という名前のGravwellグループにマップしています。

```
	Groups-Attribute=http://schemas.xmlsoap.org/claims/Group
	Group-Mapping=Gravwell Users:gravwell-users
```

### テスト設定

AD FSとGravwellの両方を構成すると、GravwellログインページでSSOログインボタンが有効になります。

![](sso-login.png)

クリックすると、WindowsサーバーのAD FS認証ページが表示されます。 ユーザー名にドメインを含む、Active Directoryユーザーの1人として認証します（"jfloren"ではなく"jfloren@gravwell.io"と入力します）。

![](sso-page.png)

[サインイン]をクリックすると、Gravwell UIに戻ります。今回は適切なユーザーとしてログインしました。
