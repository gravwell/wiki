# Okta

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Okta Hosted Ingester](/ingesters/okta)
    Kit, [Okta Kit](https://github.com/gravwell/kits/tree/main/okta)
:::

## Okta Configuration

### Okta System Log API
To configure Okta for ingestion with the API you will need the following:

* **Domain:** Your Okta org domain (e.g. `myorg.okta.com`). This is the address of your Okta Admin Console without `-admin`.
* **Credentials:** An OAuth 2.0 service app (recommended) or an API token.

The two authentication methods are mutually exclusive: configure either `Client-ID` with `Private-Key` or `Private-Key-File`, or `Token`. We recommend the OAuth 2.0 service app:

* It needs no administrator account. An API token acts with the permissions of the admin who created it, so it needs a dedicated (and licensed) admin user.
* It is limited to the `okta.logs.read` and `okta.users.read` scopes and the Read-only Administrator role.
* Its access tokens are short-lived and renewed automatically. An API token stays valid for as long as it is used.
* Its requests appear in the Okta System Log as the app, apart from the activity of human admins.

#### Option 1: OAuth 2.0 Service App (Recommended)

The setup needs an Okta administrator who can create app integrations and assign admin roles, normally a Super Administrator. See Okta's guide [Implement OAuth for Okta with a service app](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/main/) for background.

```{note}
A client secret is not supported. Okta accepts only `private_key_jwt` client authentication for access tokens with Okta API (`okta.*`) scopes.
```

1. In the Okta Admin Console, go to **Applications** > **Applications** and click **Create App Integration**. Choose **API Services** and click **Next**. If the dialog has two tabs, it is on the **Classic experience** tab.

   ```{image} ../../ingesters/okta-oauth-create-app.png
   :align: center
   ```

2. Name the app, for example `Gravwell Log Collection`. Keep **Use Okta-generated client ID** and click **Save**.

   ```{image} ../../ingesters/okta-oauth-app-name.png
   :align: center
   ```

3. On the app's **General** tab, find **Public keys**, keep **Save keys in Okta** and click **Add key**.

   ```{image} ../../ingesters/okta-oauth-public-keys.png
   :align: center
   ```

4. In the **Add a public key** dialog, do one of the following, then click **Done** and **Save**:
   * Paste the public JWK you built in [Using Your Own Key](okta_integration_own_key) (recommended).
   * Keep **Key Size** at **4096-bit (Recommended)** and click **Generate new key**. Okta shows the private key as a JWK only once, so copy it now for the ingester configuration.

   ```{image} ../../ingesters/okta-oauth-add-key.png
   :align: center
   ```

5. Under **Client Credentials**, click **Edit**, set **Client authentication** to **Public key / Private key** and click **Save**. Okta allows this only once the app has an active key. You can leave **Require Demonstrating Proof of Possession (DPoP) header in token requests** on or off; the ingester handles DPoP automatically.

   ```{image} ../../ingesters/okta-oauth-client-auth.png
   :align: center
   ```

6. On the **Okta API Scopes** tab, click **Grant** next to `okta.logs.read` and `okta.users.read`. Grant no others.

   ```{image} ../../ingesters/okta-oauth-scopes.png
   :align: center
   ```

7. On the **Admin roles** tab, click **Edit assignments**, then **Add assignment**. Choose **Read-only Administrator** as the **Role** and click **Save Changes**. The scopes alone are not enough: Okta also checks the app's admin role.

   ```{image} ../../ingesters/okta-oauth-role-assign.png
   :align: center
   ```

8. Check that the **Admin roles** tab now lists **Read-only Administrator**.

   ```{image} ../../ingesters/okta-oauth-roles.png
   :align: center
   ```

9. Copy the **Client ID** from **Client Credentials** on the **General** tab, for the `Client-ID` parameter.

10. (Optional) On the **Application Rate Limits** tab, click **Edit** to limit the share of each Okta API rate limit that the app may use. Make sure the ingester's `Request-Per-Minute` (default 60) stays within that share.

    ```{image} ../../ingesters/okta-oauth-rate-limits.png
    :align: center
    ```

(okta_integration_own_key)=
#### Using Your Own Key

The key pair does not have to come from Okta. If you generate it on the ingester host, the private key never leaves that host and Okta receives only the public key.

1. Generate a 4096-bit RSA key that only the `gravwell` user can read:

   ```
   sudo openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:4096 -out /opt/gravwell/etc/okta-private.pem
   sudo chown gravwell:gravwell /opt/gravwell/etc/okta-private.pem
   sudo chmod 600 /opt/gravwell/etc/okta-private.pem
   ```

2. Build the public key as a JWK. Okta requires a `kid`; any ID that is unique within the app will do.

   On Linux:

   ```
   n=$(sudo openssl rsa -in /opt/gravwell/etc/okta-private.pem -noout -modulus | cut -d= -f2 | basenc --base16 -d | basenc --base64url -w0 | tr -d '=')
   printf '{"kty":"RSA","use":"sig","alg":"RS256","kid":"gravwell-1","e":"AQAB","n":"%s"}\n' "$n"
   ```

   On macOS, which has `xxd` but no `basenc`:

   ```
   n=$(sudo openssl rsa -in /opt/gravwell/etc/okta-private.pem -noout -modulus | cut -d= -f2 | xxd -r -p | base64 | tr -d '=\n' | tr '+/' '-_')
   printf '{"kty":"RSA","use":"sig","alg":"RS256","kid":"gravwell-1","e":"AQAB","n":"%s"}\n' "$n"
   ```

   Check that `"n"` is not empty: for a 4096-bit key it is 683 characters long. Stock Debian has no `xxd`, so use the `basenc` form on Linux.

3. Paste the JWK into the **Add a public key** dialog in step 4 above, and set `Private-Key-File` and `Private-Key-ID` in the ingester configuration as shown in the [sample config](okta_integration_oauth_sample) below.

See the [Okta ingester documentation](/ingesters/okta) for the private key formats it accepts, how to give the key inline, and how to [rotate the key](okta_key_rotation).

#### Option 2: API Token

Start by creating a dedicated Okta Service Account for logging purposes (something like "Okta Log User"). This user should be assigned to the "Read Only Admin" role. See the [Okta documentation](https://developer.okta.com/docs/guides/create-an-api-token/main/) for instructions on generating an API token.

```{attention}
Do not use a token with write permissions to your Okta instance to the ingester. This gives significantly more access than is needed for monitoring.
```

1. Create an API token in the Okta Admin Console under **Security** > **API**.

```{image} ../../ingesters/okta-sidebar.png
:align: center
```

2. Create a new token.

```{image} ../../ingesters/okta-create-token.png
:align: center
```

3. Make sure to copy the token value to use in setting up the ingester.

```{image} ../../ingesters/okta-create-token-success.png
:align: center
```

4. Double-check the role is "Read Only Admin" and that it is attached to a dedicated service account.

```{image} ../../ingesters/okta-token-check.png
:align: center
```

#### Token Rate Limits

Okta is extremely sensitive to rate limits so double-check that your token rate limits align with the `Request-Per-Minute` config parameter. The ingester primarily hits the `/api/v1/logs` endpoint.

```{image} ../../ingesters/okta-token-rate-limits.png
:align: center
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

#### Sample Well Config
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/okta-well.conf`
```ini
[Storage-Well "okta"]
    Location=/opt/gravwell/storage/okta
    Tags=okta*
```

### Gravwell Ingester Configuration

If the Hosted Runner is not installed, follow the [configuration guide for Okta](/ingesters/okta) to create your own configuration.

Edit: `/opt/gravwell/etc/hosted_runner.conf`

(okta_integration_oauth_sample)=
#### Sample Okta Config: OAuth 2.0 Service App

With a key generated on the ingester host:

```ini
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Client-ID="0oa1b2c3d4e5f6g7h8i9"
    Private-Key-File="/opt/gravwell/etc/okta-private.pem"
    Private-Key-ID="gravwell-1"
```

With the JWK that the Okta console generated, given inline between backticks. `Private-Key-ID` defaults to the JWK's `kid`:

```ini
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Client-ID="0oa1b2c3d4e5f6g7h8i9"
    Private-Key=`{"kty":"RSA","kid":"...","n":"...","e":"AQAB","d":"...","p":"...","q":"...","dp":"...","dq":"...","qi":"..."}`
```

#### Sample Okta Config: API Token

```ini
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Token="your-okta-api-token"
```

#### Sample Okta Config: Rate Limiting Lowered for Lower-Tier Okta Plans
```ini
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Token="your-okta-api-token"
    Request-Per-Minute=10
    Request-Burst=5
    Request-Batch-Size=100
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_hosted_runner.service`
```

### Troubleshooting

* **`invalid_client: The client_assertion JWT kid is invalid.`** `Private-Key-ID` is wrong, or it is missing while the app has more than one active key. Okta logs these under "unknown client" with the reason `client_assertion_invalid_kid`, so a System Log filter on the app's client ID does not show them.
* **A failed token request every minute.** The Hosted Runner restarts a stopped ingester every minute, so a bad key, a removed key or a deactivated app leaves one failed token request per minute in the Okta System Log until it is fixed.
* **A missing scope.** Okta still issues a token without it. The ingester logs a warning naming the scope, and the Users API returns 403 `insufficient_scope`.
* **Three DPoP events per token request.** When the app requires DPoP, Okta logs a missing-proof failure, a nonce challenge and the grant for each token request. This is normal.

See [Troubleshooting](okta_troubleshooting) in the Okta ingester documentation for more detail.
