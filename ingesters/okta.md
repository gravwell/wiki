---
myst:
  substitutions:
    package: "gravwell-hosted-runner"
    standalone: "gravwell_hosted_runner"
    dockername: "hosted_runner"
---
# Okta Ingester

The [Okta](https://www.okta.com/) ingester polls the Okta System Log and Users APIs.

* **System logs** (`okta` tag) all Okta system events: logins, MFA challenges, policy evaluations, admin actions, application access, and more. Events are ingested in order with timestamps preserved from the original records.
* **User records** (`okta-users` tag) snapshots of user profile changes, polled every 10 minutes. Useful for correlating identities against activity in the system log.

The tags `okta` and `okta-users` are fixed and expected by the Gravwell Okta kit; they cannot be changed.

This ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration). Multiple Okta stanzas can coexist alongside other Hosted Runner plugins in a single configuration file.

## Installation

```{include} installation_instructions_template 
```

If you already have the hosted runner installed, you can modify the config.

## Configuration

To configure the ingester you will need the following from Okta:

* **Domain**: Your Okta org domain, e.g. `myorg.okta.com`. This is the address of your Okta Admin Console without `-admin`.
* **Credentials**: An OAuth 2.0 service app (recommended) or an API token. See [Authentication](okta_authentication).

(okta_authentication)=
### Authentication

The ingester authenticates to Okta in one of two ways, and they are mutually exclusive: configure either `Client-ID` with `Private-Key` or `Private-Key-File`, or `Token`.

| Method                              | Parameters                                                                  | Okta setup                                                 |
|-------------------------------------|-----------------------------------------------------------------------------|------------------------------------------------------------|
| OAuth 2.0 service app (recommended) | `Client-ID`, `Private-Key` or `Private-Key-File`, and `Private-Key-ID`      | [Setting up an OAuth 2.0 Service App](okta_oauth_setup)    |
| API token                           | `Token`                                                                     | [Creating an Okta API Token](okta_api_token)               |

We recommend the OAuth 2.0 service app:

* It needs no administrator account. An API token acts with the permissions of the admin who created it, so it needs a dedicated (and licensed) admin user, or it ties log collection to a person's account.
* It is limited to the `okta.logs.read` and `okta.users.read` scopes and the Read-only Administrator role.
* Its access tokens are short-lived and renewed automatically. An API token stays valid for as long as it is used.
* Its requests appear in the Okta System Log as the app, apart from the activity of human admins.

Both methods use the same System Log checkpoint, so switching an existing stanza from `Token` to OAuth neither drops nor repeats events.

(okta_oauth_setup)=
### Setting up an OAuth 2.0 Service App

The ingester uses the OAuth 2.0 client credentials grant with `private_key_jwt` client authentication, as described in Okta's guide [Implement OAuth for Okta with a service app](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/main/). The setup needs an Okta administrator who can create app integrations and assign admin roles, normally a Super Administrator.

```{note}
A client secret is not supported. Okta accepts only `private_key_jwt` client authentication for access tokens with Okta API (`okta.*`) scopes.
```

#### Step 1: Create an API Services App

1. In the Okta Admin Console, go to **Applications > Applications** and click **Create App Integration**.
2. Choose **API Services** and click **Next**. If the dialog has two tabs, it is on the **Classic experience** tab.

   ![Create App Integration dialog with API Services selected](okta-oauth-create-app.png)

3. Name the app, for example `Gravwell Log Collection`. Keep **Use Okta-generated client ID** and click **Save**.

   ![New API Services App Integration with the app name and an Okta-generated client ID](okta-oauth-app-name.png)

(okta_oauth_key)=
#### Step 2: Add a Key and Set Client Authentication

The ingester signs its token requests with an RSA private key, and Okta checks them with the matching public key. You can generate the key pair on the ingester host and give Okta only the public key (see [Using Your Own Key](okta_own_key)), or let Okta generate it.

1. On the app's **General** tab, find **Public keys**, keep **Save keys in Okta** and click **Add key**.

   ![Public keys panel with Save keys in Okta selected and the Add key button](okta-oauth-public-keys.png)

2. In the **Add a public key** dialog, do one of the following, then click **Done**:
   * Paste the public JWK you built in [Using Your Own Key](okta_own_key).
   * Keep **Key Size** at **4096-bit (Recommended)** and click **Generate new key**. Okta shows the private key as a JWK only once, so copy it now. It goes in `Private-Key` or in the file named by `Private-Key-File`.

   ![Add a public key dialog with the Generate new key option](okta-oauth-add-key.png)

3. Check that the key is listed with status **Active**, note its **KID**, and click **Save**.
4. Under **Client Credentials**, click **Edit**, set **Client authentication** to **Public key / Private key** and click **Save**. Okta allows this only once the app has an active key.

   ![Client Credentials with Public key / Private key selected for Client authentication](okta-oauth-client-auth.png)

You can leave **Require Demonstrating Proof of Possession (DPoP) header in token requests** on or off. The ingester handles DPoP automatically.

#### Step 3: Grant the API Scopes

1. Open the **Okta API Scopes** tab.
2. Click **Grant** next to `okta.logs.read` (System Log) and `okta.users.read` (Users). Grant no others.

![Okta API Scopes tab with okta.logs.read and okta.users.read granted](okta-oauth-scopes.png)

(okta_oauth_role)=
#### Step 4: Assign the Read-only Administrator Role

The scopes alone are not enough: Okta also checks the app's admin role. Read-only Administrator applies to the whole org and cannot be narrowed to a resource set.

1. Open the **Admin roles** tab and click **Edit assignments**.
2. Click **Add assignment**, choose **Read-only Administrator** as the **Role** and click **Save Changes**.

   ![Administrator assignment with Read-only Administrator chosen as the role](okta-oauth-role-assign.png)

3. Check that the **Admin roles** tab now lists **Read-only Administrator**.

   ![Admin roles tab listing Read-only Administrator](okta-oauth-roles.png)

#### Step 5: Copy the Client ID

Copy the **Client ID** from **Client Credentials** on the app's **General** tab. It goes in the `Client-ID` parameter. The client ID is not a secret.

#### Optional: Limit the App's API Rate

On the **Application Rate Limits** tab, click **Edit** to set the share of each Okta API rate limit that the app may use. Make sure `Request-Per-Minute` (default 60) stays within that share.

![Application Rate Limits tab](okta-oauth-rate-limits.png)

(okta_own_key)=
### Using Your Own Key

The key pair does not have to come from Okta. If you generate it on the ingester host, the private key never leaves that host and Okta receives only the public key. This is the preferred setup when someone other than the Okta administrator runs the ingester.

1. Generate a 4096-bit RSA key that only the `gravwell` user can read:

   ```
   sudo openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:4096 -out /opt/gravwell/etc/okta-private.pem
   sudo chown gravwell:gravwell /opt/gravwell/etc/okta-private.pem
   sudo chmod 600 /opt/gravwell/etc/okta-private.pem
   ```

2. Build the public key as a JWK. Okta requires a `kid`; any ID that is unique within the app will do. `AQAB` is the public exponent 65537, the `openssl` default.

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

   ```{warning}
   Check that `"n"` is not empty: for a 4096-bit key it is 683 characters long. If a command in the pipeline is missing, the shell prints only a "command not found" error and the JWK gets `"n":""`. Stock Debian has no `xxd`, so use the `basenc` form on Linux.
   ```

3. In [Step 2](okta_oauth_key), paste the JWK into the **Add a public key** dialog instead of generating a key.

4. Point the ingester at the private key and give it the same `kid`:

   ```
   Client-ID="0oa1b2c3d4e5f6g7h8i9"
   Private-Key-File="/opt/gravwell/etc/okta-private.pem"
   Private-Key-ID="gravwell-1"
   ```

(okta_api_token)=
### Creating an Okta API Token

Start by creating a dedicated Okta Service Account for logging purposes (something like "Okta Log User"). This user should be assigned to the "Read Only Admin" role. 

```{attention}
Do **not** use a token with write permissions to your Okta instance to the ingester. This gives significantly more access than is needed for monitoring. 
```

You can create an API token in the Okta Admin Console > Security > Api section. See the [Okta documentation](https://developer.okta.com/docs/guides/create-an-api-token/main/) for details.

![](okta-sidebar.png)

![](okta-create-token.png)

Once created the token value will be used in the `Token` config parameter. Make sure to copy it. 

![](okta-create-token-success.png)

Once create you should see this token in your list and should double-check the role is "Read Only Admin", and that it is attached to a dedicated service account.

![](okta-token-check.png)

#### Token Rate Limits
Okta is extremely sensitive to rate limits so double-check your token rate limits align to the `Request-Per-Minute` Config Parameter. The Ingester primarily hits the `/api/v1/logs` endpoint.

![](okta-token-rate-limits.png)

(okta_parameters)=
### Okta Stanza Parameters

The Okta ingester is configured via `[Okta "name"]` stanzas in the Hosted Runner configuration file, typically `/opt/gravwell/etc/hosted_runner.conf`. The `[Global]` and `[State]` blocks common to all Hosted Runner plugins are described in [Hosted Runner Configuration](hosted_runner_configuration).

| Config Parameter   | Type    | Required             | Default         | Description |
|--------------------|---------|----------------------|-----------------|-------------|
| Ingester-UUID      | UUID    | yes                  |                 | A unique UUID for this ingester instance. Used for state tracking. |
| Domain             | string  | yes                  |                 | Your Okta org domain, e.g. `myorg.okta.com`. Must be a host under `okta.com`, `oktapreview.com`, `okta-emea.com`, `okta-gov.com` or `okta.mil`. Custom URL domains are not accepted. |
| Token              | string  | one of               |                 | Okta SSWS API token. Set either `Token` or `Client-ID`, not both. |
| Client-ID          | string  | one of               |                 | Client ID of the OAuth 2.0 service app. Requires `Private-Key` or `Private-Key-File`. |
| Private-Key        | string  | with `Client-ID`     |                 | The service app's private key, as a JWK or as an unencrypted RSA PEM (PKCS#1 or PKCS#8). Set either `Private-Key` or `Private-Key-File`, not both. See [Private Key Formats](okta_private_key). |
| Private-Key-File   | string  | with `Client-ID`     |                 | Path to a file holding the private key, in either form. |
| Private-Key-ID     | string  | with 2+ active keys  | the JWK's `kid` | Key ID (`kid`) sent with each token request. Required whenever the app has more than one active key, which includes the middle of every [key rotation](okta_key_rotation). With a single active key, Okta accepts a request without one. A PEM key carries no `kid`, so set this to the KID shown under the app's **Public keys**. |
| Request-Batch-Size | integer | no                   | 100             | Number of log entries to request per API call. |
| Request-Per-Minute | integer | no                   | 60              | Maximum number of API requests per minute. |
| Request-Burst      | integer | no                   | 10              | Burst capacity for the request rate limiter. |

(okta_private_key)=
### Private Key Formats

The private key can be either of these:

* The JWK JSON that the Okta console shows when it generates a key. A JWK set holding a single key also works.
* An unencrypted RSA PEM, either PKCS#1 (`-----BEGIN RSA PRIVATE KEY-----`) or PKCS#8 (`-----BEGIN PRIVATE KEY-----`), such as the file `openssl genpkey` writes.

Encrypted keys and public JWKs are rejected when the configuration is loaded.

#### In a File

`Private-Key-File` keeps the key out of the configuration file, and is the recommended form.

* The file is read when the configuration is loaded, at startup and on every [reload](hosted_runner_reload). A reload picks up a changed key, even if the configuration file itself is unchanged.
* A missing, unreadable or invalid key file fails the load of the whole configuration. The Hosted Runner does not start, or, on a reload, it logs `failed to reload config` and keeps running with the previous configuration.
* When the ingester starts, it logs a warning if users other than the file's owner can access the file, and keeps running:

  ```
  the okta Private-Key-File is accessible to other users, restrict it to its owner (chmod 600)
  ```

  The warning includes the file's path and mode. Run `chmod 600` on the file to clear it.

#### Inline

`Private-Key` takes the key itself. These forms are all accepted:

```
Private-Key=`{"kty":"RSA","kid":"...","n":"...","e":"AQAB","d":"...","p":"...","q":"...","dp":"...","dq":"...","qi":"..."}`
Private-Key="{\"kty\":\"RSA\",\"kid\":\"...\",\"n\":\"...\", ... }"
Private-Key="-----BEGIN PRIVATE KEY-----\nMIIEv...\n-----END PRIVATE KEY-----"
Private-Key=`-----BEGIN PRIVATE KEY-----
MIIEv...
-----END PRIVATE KEY-----`
```

A value between backticks is raw, so a JWK can be pasted as it is, and so can a PEM with its line breaks. Between double quotes, escape each `"` as `\"` and write the PEM's line breaks as `\n`.

(okta_key_rotation)=
### Rotating the Private Key

To rotate the key without interrupting collection, keep both keys active in the app while the ingester switches over:

1. If `Private-Key-ID` is not set, set it to the KID of the current key, as listed under **Public keys** on the app's **General** tab, and [reload](hosted_runner_reload) the Hosted Runner. Once the app has two active keys, Okta rejects a token request without a `kid`.
2. Generate a new key pair next to the old one, and build its public JWK under a new `kid`, such as `gravwell-2`, as in [Using Your Own Key](okta_own_key):

   ```
   sudo openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:4096 -out /opt/gravwell/etc/okta-private.pem.new
   sudo chown gravwell:gravwell /opt/gravwell/etc/okta-private.pem.new
   sudo chmod 600 /opt/gravwell/etc/okta-private.pem.new
   ```

3. Add the new public JWK to the app with **Add key**. The app now has two active keys.
4. Rename the new key file over the old one. A rename replaces the file in one step, so the Hosted Runner never reads a half-written key.

   ```
   sudo mv /opt/gravwell/etc/okta-private.pem.new /opt/gravwell/etc/okta-private.pem
   ```

5. Set `Private-Key-ID` to the new `kid` and reload the Hosted Runner. It logs `restarted ingester on reload`, and the ingester requests a new token signed with the new key.
6. Once Okta has granted a token with the new key, deactivate and delete the old key in the app.

With an inline `Private-Key`, replace the value in the configuration in step 4 instead of the file.

## Example Configuration

OAuth 2.0 service app, with a key generated on the ingester host:

```
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Client-ID="0oa1b2c3d4e5f6g7h8i9"
    Private-Key-File="/opt/gravwell/etc/okta-private.pem"
    Private-Key-ID="gravwell-1"
```

OAuth 2.0 service app, with the JWK that the Okta console generated, given inline. `Private-Key-ID` defaults to the JWK's `kid`:

```
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Client-ID="0oa1b2c3d4e5f6g7h8i9"
    Private-Key=`{"kty":"RSA","kid":"...","n":"...","e":"AQAB","d":"...","p":"...","q":"...","dp":"...","dq":"...","qi":"..."}`
```

API token:

```
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Token="your-okta-api-token"
```

With optional rate limiting tuned down for lower-tier Okta plans:

```
[Okta "myorg"]
    Ingester-UUID="99c00000-0000-0000-0000-000000000000"
    Domain="myorg.okta.com"
    Token="your-okta-api-token"
    Request-Per-Minute=10
    Request-Burst=5
    Request-Batch-Size=100
```

(okta_troubleshooting)=
## Troubleshooting

### `invalid_client: The client_assertion JWT kid is invalid.`

The token request carried a `kid` that matches no active key of the app, or carried no `kid` while the app has more than one active key. Set `Private-Key-ID` to the KID of the key that the ingester holds.

Okta records these failures in the System Log under "unknown client" with the reason `client_assertion_invalid_kid`, so a System Log filter on the app's client ID does not show them.

### A Failed Token Request Every Minute

When a token request fails, the ingester stops, and the Hosted Runner restarts stopped ingesters every minute. A wrong key, a key removed from the app, or a deactivated app therefore leaves one failed token request per minute in the Okta System Log until it is fixed.

Once the problem is fixed, the ingester resumes from its saved checkpoint, so no System Log events are lost. A fix on the Okta side takes effect at the next restart, within a minute; a fix in the configuration takes effect when the Hosted Runner is [reloaded](hosted_runner_reload).

### A Scope Is Missing

Okta still issues a token when a requested scope is not granted to the app; it just leaves the scope out. The ingester logs a warning that names the missing scope:

```
okta access token is missing a scope, grant it on the service app's Okta API Scopes tab
```

Without `okta.users.read`, the Users API returns 403 `insufficient_scope`, the ingester logs `the okta.users.read scope is not granted to the service app, users cannot be collected`, and System Log collection continues. Grant the scope on the app's **Okta API Scopes** tab. The running ingester picks it up at its next users poll, without a restart.

### Other 403 Responses

In OAuth mode, the ingester adds `check that the service app is assigned the Read-only Administrator role and granted the okta.logs.read and okta.users.read scopes` to any other 403 error from Okta. Check the app's role assignment in [Step 4](okta_oauth_role).

### DPoP Events in the Okta System Log

When the app requires DPoP, each token request appears as three System Log events: a failure for the missing DPoP proof, a nonce challenge (`use_dpop_nonce`), then the grant. This is the normal DPoP handshake, not an error. The ingester reuses each access token until it expires, so this happens about once an hour.

### `"<domain>" is not an okta domain`

`Domain` must be your org's Okta domain, a host under `okta.com`, `oktapreview.com`, `okta-emea.com`, `okta-gov.com` or `okta.mil`. A custom URL domain, such as `login.example.com`, is not accepted.

## Additional Resources

* [Okta System Log API Reference](https://developer.okta.com/docs/reference/api/system-log/)
* [Okta Users API Reference](https://developer.okta.com/docs/reference/api/users/)
* [Implement OAuth for Okta with a Service App](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/main/)
* [Creating an Okta API Token](https://developer.okta.com/docs/guides/create-an-api-token/main/)
