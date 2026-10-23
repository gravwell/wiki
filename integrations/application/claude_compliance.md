# Claude Enterprise Compliance

| Integration Details | |
| --- | --- |
| Audience | Claude organization owner and Gravwell operator |
| Ingester | [Gravwell Hosted Runner](https://docs.gravwell.io/configuration/configuration.html#hosted-runner-configuration) with the ClaudeCompliance plugin |
| Data source | [Claude Enterprise Compliance API](https://platform.claude.com/docs/en/manage-claude/compliance-api) |
| Transport | Outbound HTTPS polling of `https://api.anthropic.com/v1/compliance`; no inbound listener or webhook |
| Framing | One compact native JSON object per Gravwell entry |

Send Claude Enterprise activity, organization directory records, conversations, projects, artifact metadata and session transcripts to Gravwell. The ingester polls the Compliance API and preserves each native JSON record. `_source` identifies the dataset, `_parent` records only discovery-derived parent request parameters, and `_session` carries a returned session envelope on session-message entries. Seven root datasets discover related records automatically; other object operations take explicit IDs. The ingester supports 26 JSON read operations mapped to six tags. The core kit includes a 26-row operation-to-endpoint-to-default-tag resource.

The ingester does not download binary file, artifact or project-document content through the `/content` endpoints. It does not call any delete endpoint, collect Enterprise Analytics, or list Claude Code Artifacts (`GET /v1/compliance/apps/code/artifacts`). It does not enumerate file, document or artifact IDs automatically.

This guide covers both roles. The Claude organization owner completes the Vendor setup section and hands the result to the Gravwell operator, who completes the remaining sections.

## Prerequisites and permissions

- A Claude Enterprise organization with Compliance API access. The parent organization's primary owner enables access for the parent and its linked organizations. An organization owner can create a key limited to that owner's organization after enablement.
- A Compliance Access Key with only the read scopes needed for the agreed coverage. No delete scope is required.
- A Hosted Runner build that includes the ClaudeCompliance plugin. The plugin is an upstream contribution and is not yet confirmed in a released Gravwell package. Confirm availability and the supported Gravwell version with your distribution provider before you enable these stanzas.
- Outbound HTTPS from the Hosted Runner host to `api.anthropic.com` on port 443, with normal certificate verification.
- Access from the Hosted Runner to the Gravwell ingest listener, preferably the encrypted listener, and an ingest secret for that deployment.
- Persistent, durable storage for the Hosted Runner state database and ingest cache.

| Collection | Required key scope |
| --- | --- |
| Activity Feed (`activities`) | `read:compliance_activities` |
| Organizations, roles, role permissions, groups and organization settings | `read:compliance_org_data` |
| Organization users and group members | `read:compliance_user_data` |
| Chats, messages, files, projects, project documents, artifacts, local and remote sessions | `read:compliance_user_data` |
| The complete seven-root configuration in this guide | All three read scopes above |

Directory child discovery lists users and group members, so a directory configuration needs both `read:compliance_org_data` and `read:compliance_user_data`. The organization settings endpoint requires `read:compliance_org_data`; the older `read:compliance_org_settings` scope was retired and no longer authorizes it. An eligible Console Admin API key created after Compliance API enablement can read the Activity Feed only. Ordinary Claude API keys and Analytics-only access do not grant Compliance access. See [Set up the Compliance API](https://platform.claude.com/docs/en/manage-claude/compliance-api-access).

The vendor administrator who creates the key and the Hosted Runner that reads with it hold separate privileges. The runner needs only the key file; it never needs a claude.ai administrator session.

## Claude Enterprise Configuration

The Claude organization owner performs these steps in claude.ai. They do not require Gravwell access.

### Enable access and create the key

1. Sign in to claude.ai as the parent organization's primary owner. Open **Organization settings → API** at [API access](https://claude.ai/admin-settings/api-access) and enable the Compliance API. Parent enablement covers linked organizations. Recording starts at enablement: activity and local-session content from before enablement, or while access is disabled, cannot be recovered later.
2. Decide whether collection covers the parent and all linked organizations or a single organization, and sign in with the matching owner role. An organization owner's key is limited to that organization.
3. On **Organization settings → API**, find **Keys**, choose **Create key**, and give it an organization-managed name such as `Gravwell Compliance`. Select `read:compliance_activities` for activity-only collection. Also select `read:compliance_org_data` and `read:compliance_user_data` for the full seven-root configuration. Choose **Create**.
4. Store the displayed secret immediately in your organization's secrets manager. It is shown only once. Key scopes cannot be edited after creation, and these keys do not expire automatically.
5. Confirm the **Scopes** column in the Keys table and record the key owner, the organizations the key can read, and who owns rotation.

For a standalone Claude Console organization, an admin enables **Settings → Security → Compliance API** and then creates a new Admin API key; that path provides Activity Feed coverage only.

A key with `read:compliance_user_data` can read conversation and transcript content across its organization boundary. Select that boundary deliberately.

### Confirm source coverage

- Local-session transcripts require users signed in to the Enterprise organization. Sessions using Console API keys or third-party model platforms are outside the source. In HIPAA-ready organizations, local-session coverage is limited to Cowork and Claude Code and captured content is retained for 30 days. Sessions under zero data retention are excluded.
- Remote-session endpoints cover the vendor's documented remote-session products, currently Cowork sessions in Anthropic-managed cloud environments. They do not cover every cloud-hosted Claude activity. Review [session coverage](https://platform.claude.com/docs/en/manage-claude/compliance-sessions).
- The Activity Feed retains activity for six years. Chat, file and project content follows the organization's claude.ai retention policy unless a user deletes it sooner. Local-session transcripts follow the custom conversation retention period when a finite one is set. Remote-session transcripts disappear when a user deletes the session.
- Transcript tool inputs and results are truncated by the source at 10,000 bytes per item by default. The API's `truncated`, `content_unavailable` and `provenance` indicators are source conditions, not ingestion failures.

Do not change retention, encryption or privacy policies solely to make a validation query succeed.

### Validate at the source

Use the new key in an approved API client for read-only checks against each family you enabled. Send the key in the `x-api-key` header with `anthropic-version: 2023-06-01`.

- `GET /v1/compliance/activities` for activity access.
- `GET /v1/compliance/organizations` or `/groups` for organization data.
- `/users` or `/members` if directory children are required. A successful organization-data request does not prove user-data access.
- `GET /v1/compliance/apps/chats` or a session list for content access.

Record the HTTP status and a sanitized record ID and time for each check. An empty successful list proves authorization only. A 403 means a missing scope. A bare `Not found` 404 means the request did not authenticate. See [Compliance API errors](https://platform.claude.com/docs/en/manage-claude/compliance-errors).

### Hand off to the Gravwell operator

Give the operator:

- The organization boundary: parent and linked organizations, or one organization.
- The enabled read scopes and the enablement time.
- The selected roots and any additional object operations.
- Retention and coverage constraints.
- Other services that share the parent organization's API budget.
- The key and rotation owners.

For explicit object operations, supply IDs returned by the API: the bare organization UUID, and the group, chat, project, session, file, document or artifact-version ID. A claude.ai browser URL identifier is not a substitute.

Transfer the key only through the approved secrets manager, as a file containing only the key with no header name or quotation marks. Never send it by email, ticket or chat. When a key is replaced, state whether the new key keeps the same visibility boundary so the operator can preserve the existing stanza name, UUID and checkpoints.

### Rotate and revoke

1. Create a same-scope replacement key and transfer it securely.
2. The operator replaces the credential file atomically. The ingester rereads it on each collection cycle.
3. The operator confirms successful collection with the new key.
4. Delete the previous key. Deletion takes effect on the next request.

To roll back a newly introduced key, coordinate with the operator before disabling it, and leave organization policies and other keys unchanged.

## Gravwell Configuration

Use a Hosted Runner distribution that includes the ClaudeCompliance plugin, and install or update it through its supported service or container procedure; see [Hosted Runner configuration](https://docs.gravwell.io/configuration/configuration.html#hosted-runner-configuration). An existing Hosted Runner can host these stanzas alongside other plugins. Keep its `[Global]` and `[State]` blocks.

Create the Claude key file and the Gravwell ingest-secret file as two separate protected files, for example `/opt/gravwell/secrets/claude-compliance-key` and `/opt/gravwell/secrets/gravwell-ingest.secret`. Give the Hosted Runner service account read access, restrict group access to an authorized service group, and deny access to other users. The same account needs write access to the state and cache directories.

### Gravwell Storage Well Configuration

On each intended indexer, add or merge this stanza in `/opt/gravwell/etc/gravwell.conf`, confirm the storage path, and apply it through your deployment's indexer configuration procedure. Avoid an overlapping route in another well.

```ini
[Storage-Well "claude-compliance"]
Location=/opt/gravwell/storage/claude-compliance
Tags=claude-compliance*
```

The wildcard routes the six default tags. If you change `Tag-Name`, update the well route and the kit's tag macros to match.

### Gravwell Ingester Configuration

Edit `/opt/gravwell/etc/hosted_runner.conf`, or an overlay in `/opt/gravwell/etc/hosted_runner.conf.d`. `[Global]` and `[State]` are shared by every stanza. A `[ClaudeCompliance "name"]` stanza can repeat `Dataset` to collect multiple operations. Each selected dataset receives its own checkpoint and its family tag, while the stanza owns the shared polling and rate-limit settings.

### Complete example

This example enables all seven roots and their discovered children in one stanza; it requires all three read scopes. Generate two different UUIDs with `uuidgen` and replace both zero UUIDs: one for `[Global]` and one for the Claude Compliance stanza. Keep them unchanged through restarts and upgrades.

```ini
[Global]
Label="Claude Compliance"
Ingester-UUID="00000000-0000-0000-0000-000000000000"
Ingest-Secret-File="/opt/gravwell/secrets/gravwell-ingest.secret"
# Ingest-Secret="REPLACE_WITH_INGEST_SECRET" # Alternative to Ingest-Secret-File
Encrypted-Backend-Target="gravwell.example.com:4024"
# Cleartext-Backend-Target="127.0.0.1:4023"
# Pipe-Backend-Target="/opt/gravwell/comms/pipe"
Connection-Timeout=30s
Insecure-Skip-TLS-Verify=false
Ingest-Cache-Path="/opt/gravwell/cache/claude-compliance.cache"
Max-Ingest-Cache=1024
Cache-Mode="always"
Log-Level="INFO"

[State]
Path="/opt/gravwell/state/claude-compliance.state"
Sync=true

[ClaudeCompliance "claude"]
Ingester-UUID="00000000-0000-0000-0000-000000000000"
Dataset="activities"
Dataset="organizations"
Dataset="groups"
Dataset="chats"
Dataset="projects"
Dataset="local-sessions"
Dataset="remote-sessions"
Credential-File="/opt/gravwell/secrets/claude-compliance-key"
Tag-Prefix="claude-compliance"
Follow-Children="enabled"
Max-Children=100
Max-Pending=10000
Lookback=24
Overlap-Seconds=300
Requests-Per-Minute=30
Request-Interval=300
Page-Size=100
Max-Pages=1000
Max-Retries=4
```

Use file-backed secrets and a verified encrypted connection. The commented Global alternatives remain accepted by the Hosted Runner for deployments that use Gravwell's other ingest transports; see [Ingester configuration](https://docs.gravwell.io/ingesters/ingesters.html). For an existing deployment, keep the Global UUID and each existing stanza's quoted name, UUID, tag settings, state path and cache path. Add a new UUID only when adding a new stanza.

### Collect one source

For one API source only, keep `[Global]` and `[State]` and add a single stanza. This stanza calls only `GET /v1/compliance/activities` and emits only `claude-compliance-activities`; every omitted setting takes the default in the parameter table:

```ini
[ClaudeCompliance "activities"]
Ingester-UUID="REPLACE_WITH_DISTINCT_PERSISTENT_UUID"
Dataset="activities"
Credential-File="/opt/gravwell/secrets/claude-compliance-key"
```

The seven-root example is the curated common configuration. It omits organization settings, the single-object detail operations (`organization-role`, `group`, `project`, `local-session`), file and generated-file metadata, project documents and artifact metadata. Those need explicit IDs.

### Select an explicit object

For an operation that needs an ID, add a stanza with a new UUID, the operation's `Dataset` and tag, and one `Parameter` for each placeholder in its endpoint. This example reads one artifact version:

```ini
[ClaudeCompliance "artifact-version"]
Ingester-UUID="00000000-0000-0000-0000-000000000000"
Dataset="artifact-metadata"
Credential-File="/opt/gravwell/secrets/claude-compliance-key"
Parameter="artifact_version_id:REPLACE_WITH_ARTIFACT_VERSION_ID"
Follow-Children="disabled"
Requests-Per-Minute=30
Request-Interval=300
```

Where the IDs come from:

- `organization_id`: the bare `uuid` from `GET /v1/compliance/organizations`, not the `org_`-prefixed value carried by some activity records.
- `role_id`: from the organization's role list.
- `group_id`: the group `id`.
- Chat, project and session IDs: from their list endpoints.
- File IDs and artifact versions: within chat messages, `files[].id` identifies uploaded-file metadata, `generated_files[].id` generated-file metadata, and `artifacts[].version_id` artifact-version metadata.
- `document_id`: the `id` of a project attachment whose `type` is `project_doc`.

A `role-permissions` stanza needs both `organization_id` and `role_id`. An `organization-settings` stanza uses `Parameter="organization_id:REPLACE_WITH_BARE_ORGANIZATION_UUID"` and the directory tag.

### Collect all sources for one tag

There is no family-wide selector. To collect everything available for `claude-compliance-directory`:

1. Keep the `organizations` and `groups` roots with `Follow-Children="enabled"`. Together they collect organizations, users, roles, role permissions, groups and group members.
2. Add an explicit `organization-role`, `organization-settings` or `group` stanza for each object ID whose single-object record is required.

Explicit selectors that share the same object IDs and polling settings may be grouped by repeating `Dataset` in one stanza and supplying the union of their required `Parameter` values. Otherwise use separate stanzas with distinct UUIDs. The same pattern applies to the other tags: keep the root that owns the tag and add the explicit selectors required for known object IDs.

### Parameters

`[Global]` and `[State]` use the standard Hosted Runner settings documented in [Hosted Runner configuration](https://docs.gravwell.io/configuration/configuration.html#hosted-runner-configuration). Every `[ClaudeCompliance "name"]` stanza accepts these parameters:

| Parameter | Type | Required | Default | Valid values | Effect |
| --- | --- | --- | --- | --- | --- |
| `Ingester-UUID` | UUID | No | plugin default | Nonzero UUID; unique per stanza when more than one stanza exists | Identifies the stanza. Its state bucket is keyed by the plugin, stanza name and UUID, so keep it stable. |
| `Host` | URL | No | `https://api.anthropic.com` | Bare HTTPS origin without credentials, path, query or fragment | Overrides the API origin for an approved proxy, mock or future regional endpoint. |
| `Dataset` | string, repeatable | Yes | none | One or more of the 26 values in the source table; no duplicates | Selects operations. Each dataset receives its own checkpoint and resolves to its family tag. |
| `Credential` | string | One credential form | none | One nonempty line | Inline Compliance key. It takes precedence when both forms are present; prefer a protected file. |
| `Credential-File` | path | One credential form | none | Regular file, at most 16 KiB, containing one nonempty line | File holding only the Compliance key. It is reread during collection. |
| `Parameter` | `name:id`, repeatable | Only for operations whose endpoint has a placeholder | none | Each required placeholder exactly once; the ID may not contain `/ \ ? #` | Fills endpoint placeholders. Missing, duplicate and unused parameters are rejected. |
| `Tag-Prefix` | tag prefix | No | `claude-compliance` | Valid tag prefix; mutually exclusive with `Tag-Name` | Produces `<prefix>-<family>` for the six semantic families. |
| `Tag-Name` | tag | No | unset | One valid Gravwell tag; mutually exclusive with `Tag-Prefix` | Sends every selected dataset in the stanza to one tag instead of family tags. |
| `Follow-Children` | string | No | `enabled` | `enabled`, `disabled` | Whether supported root datasets discover and collect child operations. |
| `Max-Children` | integer | No | 100 | 1 to 1000 | Child collections attempted per cycle. |
| `Max-Pending` | integer | No | 10000 | `Max-Children` to 100000 | Remembered child work items. |
| `Lookback` | whole hours or days | No | 24 hours | `24`, `24h`, `1d`, `1d12h` | Initial lower bound for time-filtered operations when no checkpoint exists. |
| `Overlap-Seconds` | integer | No | 300 | 1 to 86400 | How far each new window reaches back before the previous window's end. |
| `Requests-Per-Minute` | integer | No | 60 | 1 to 600 | Request rate shared by all datasets and discovered children in this stanza. |
| `Request-Interval` | integer seconds | No | 300 | 60 or more, within the platform duration limit | Delay between completed poll cycles. |
| `Page-Size` | integer | No | 100 | Positive; clamped to each endpoint's documented maximum | `limit` sent to list endpoints. |
| `Max-Pages` | integer | No | 1000 | 1 to 10000 | Pages per call before a checkpointed continuation. |
| `Max-Retries` | integer | No | 4 | 1 to 10; 0 selects the default | Retries for HTTP 429, transient 5xx, and connection or read failures. |

### Validate and activate

Validate on the ingester host with the installed binary, under an identity that can read both secret files and reach the state path:

```sh
gravwell_hosted_runner -validate \
  -config-file /opt/gravwell/etc/hosted_runner.conf \
  -config-overlays /opt/gravwell/etc/hosted_runner.conf.d
```

Validation loads the file and overlays, checks that the credential files are readable, and opens the state database. It does not call Anthropic or prove backend delivery. For a systemd installation, activate with:

```sh
sudo systemctl restart gravwell_hosted_runner.service
```

For another service manager or a container, use that deployment's activation procedure and keep its persistent volumes. A restart affects every plugin in a shared process. The Hosted Runner also supports SIGHUP for plugin configuration reloads.

## Tags and data sources

Every request is a `GET` relative to `https://api.anthropic.com/v1/compliance` with the Compliance Access Key; the Activity Feed also accepts an eligible Console Admin API key. For every entry the ingester sets `_source` to the exact `Dataset` value. Discovered child entries also receive `_parent` containing only the parent parameters derived from vendor records, such as `group_id:rbac_group_…`; operator-configured `Parameter` values are never copied into it. Session-message entries carry `_session` when the API response includes a session envelope. Native JSON remains unwrapped; Gravwell-owned keys inside it, including `__meta` and `__disabled`, are preserved. Endpoint, vendor, product, record-type and API-version context is not duplicated into intrinsic fields.

"Default root" rows are in the seven-root example, "Default child" rows are discovered by those roots, and "Optional" rows need an explicit stanza.

| Status | Tag | Exact `Dataset` | Method and endpoint | Contents | Extra identifier / permission | Timestamp · identity |
| --- | --- | --- | --- | --- | --- | --- |
| Default root | `claude-compliance-activities` | `activities` | `GET /activities` | Audit and activity events of every documented type | `read:compliance_activities`; none; `Page-Size` ≤ 5000 | `created_at` · `id` |
| Default root | `claude-compliance-directory` | `organizations` | `GET /organizations` | Parent and linked organizations | `read:compliance_org_data`; none | collection time · `uuid` |
| Default child | `claude-compliance-directory` | `organization-users` | `GET /organizations/{organization_id}/users` | Current organization members | `read:compliance_user_data`; `organization_id` | collection time · `id` |
| Default child | `claude-compliance-directory` | `organization-roles` | `GET /organizations/{organization_id}/roles` | Custom roles | `read:compliance_org_data`; `organization_id` | `updated_at`, else collection time · `id` |
| Optional | `claude-compliance-directory` | `organization-role` | `GET /organizations/{organization_id}/roles/{role_id}` | One role | `read:compliance_org_data`; `organization_id`, `role_id` | `updated_at`, else collection time · `id` |
| Default child | `claude-compliance-directory` | `role-permissions` | `GET /organizations/{organization_id}/roles/{role_id}/permissions` | Permissions granted by a role | `read:compliance_org_data`; `organization_id`, `role_id` | collection time · `resource_type`+`resource_id`+`action` |
| Optional | `claude-compliance-directory` | `organization-settings` | `GET /organizations/{organization_id}/settings` | Effective settings and Compliance key inventory, no secrets | `read:compliance_org_data`; `organization_id` | collection time · `organization_id` |
| Default root | `claude-compliance-directory` | `groups` | `GET /groups` | Groups | `read:compliance_org_data`; none | `updated_at`, else collection time · `id` |
| Optional | `claude-compliance-directory` | `group` | `GET /groups/{group_id}` | One group | `read:compliance_org_data`; `group_id` | `updated_at`, else collection time · `id` |
| Default child | `claude-compliance-directory` | `group-members` | `GET /groups/{group_id}/members` | Group membership | `read:compliance_user_data`; `group_id` | `updated_at`, else collection time · `user_id` |
| Default root | `claude-compliance-conversations` | `chats` | `GET /apps/chats` | Chat metadata, including `deleted_at` | `read:compliance_user_data`; none | `updated_at` · `id` |
| Default child | `claude-compliance-conversations` | `chat-messages` | `GET /apps/chats/{chat_id}/messages` | Messages with file, generated-file and artifact references | `read:compliance_user_data`; `chat_id` | `created_at` · `id` |
| Optional | `claude-compliance-conversations` | `file-metadata` | `GET /apps/chats/files/{file_id}` | Uploaded-file metadata | `read:compliance_user_data`; `file_id` | `created_at` · `id` |
| Optional | `claude-compliance-conversations` | `generated-file-metadata` | `GET /apps/chats/generated-files/{file_id}` | Generated-file metadata | `read:compliance_user_data`; `file_id` | `created_at` · `id` |
| Default root | `claude-compliance-projects` | `projects` | `GET /apps/projects` | Project metadata, including `deleted_at` | `read:compliance_user_data`; none; `Page-Size` ≤ 100 | `updated_at` · `id` |
| Optional | `claude-compliance-projects` | `project` | `GET /apps/projects/{project_id}` | One project with description, instructions and counts | `read:compliance_user_data`; `project_id` | `updated_at` · `id` |
| Default child | `claude-compliance-projects` | `project-attachments` | `GET /apps/projects/{project_id}/attachments` | Attached files and project documents | `read:compliance_user_data`; `project_id`; `Page-Size` ≤ 100 | `created_at` · `id` |
| Default child | `claude-compliance-projects` | `project-collaborators` | `GET /apps/projects/{project_id}/collaborators` | User, group, organization and role grants | `read:compliance_user_data`; `project_id`; `Page-Size` ≤ 100 | collection time · `type` + principal |
| Optional | `claude-compliance-projects` | `project-document` | `GET /apps/projects/documents/{document_id}` | Project document including its text content | `read:compliance_user_data`; `document_id` | `created_at` · `id` |
| Optional | `claude-compliance-projects` | `project-document-metadata` | `GET /apps/projects/documents/{document_id}/metadata` | Document metadata with MD5 and size | `read:compliance_user_data`; `document_id` | `created_at` · `id` |
| Optional | `claude-compliance-artifacts` | `artifact-metadata` | `GET /apps/artifacts/{artifact_version_id}` | Artifact-version metadata | `read:compliance_user_data`; `artifact_version_id` | `created_at` · `version_id` |
| Default root | `claude-compliance-sessions` | `local-sessions` | `GET /apps/sessions/local` | Local session metadata | `read:compliance_user_data`; none; `Page-Size` ≤ 500 | `updated_at` · `id` |
| Optional | `claude-compliance-sessions` | `local-session` | `GET /apps/sessions/local/{session_id}` | One local session | `read:compliance_user_data`; `session_id` | `updated_at` · `id` |
| Default child | `claude-compliance-sessions` | `local-session-messages` | `GET /apps/sessions/local/{session_id}/messages` | Local transcript turns | `read:compliance_user_data`; `session_id` | `created_at` · `id` |
| Default root | `claude-compliance-sessions` | `remote-sessions` | `GET /apps/sessions/remote` | Remote session metadata | `read:compliance_user_data`; none; `Page-Size` ≤ 500 | `updated_at` · `id` |
| Default child | `claude-compliance-sessions` | `remote-session-messages` | `GET /apps/sessions/remote/{session_id}/messages` | Remote transcript turns | `read:compliance_user_data`; `session_id` | `created_at` · `id` |

The seven roots populate five tags. `claude-compliance-artifacts` stays empty until `artifact-metadata` is selected with an explicit artifact version ID. Discovered children remain in their root's resolved family tag. To combine all data under one tag, set one `Tag-Name` on the stanza. When records share a tag, `_source` tells the operations apart; `_parent` is present only when the entry came from discovery-derived child context.

## Operating behavior

### Records and timestamps

Each entry is one compact, unwrapped native JSON object. List responses are split at their record array; single-object endpoints emit the whole response. The entry timestamp comes from the field shown in the source table.

Records that have no change time, such as users, organizations, role permissions, collaborators and settings, use collection time. They do not use account or organization creation time, which would misdate a later change. The native fields remain in the JSON. Use the intrinsic module, not `json`, to read the underscore-prefixed context fields.

### Windows, checkpoints and restart

When no checkpoint exists, `Lookback` sets the first lower bound for time-filtered datasets. An existing checkpoint always takes precedence, and changing `Lookback` alone never overrides it.

Each window's upper bound trails the current time by one minute, the vendor's documented indexing delay. The next window starts `Overlap-Seconds` before that bound, so a late-indexed record is not skipped.

The time-filtered operations are:

- `activities`: `created_at`.
- `chats`: `updated_at`, with `order_by=updated_at`.
- `local-sessions`: `updated_at.gte`.
- Refreshes of discovered `chat-messages`: `created_at`.

The `projects` list is a bounded full scan. Its documented filters use creation time while projects remain mutable, so narrowing that inventory could hide edited projects.

The `chats` root stays incremental with child discovery enabled, because the vendor returns a chat again whenever it receives a new message, changes project or is deleted. With child discovery enabled, `projects` and `local-sessions` list their full inventories so an unchanged parent cannot hide changed children. The directory, remote-session and transcript operations have no time filter and are bounded full scans. `Lookback` is therefore not a history or volume limit for the whole integration. A newly discovered chat collects its complete message history once before refreshes use its window.

A page cursor is checkpointed after every completed page, so a restart resumes the same walk. Vendor cursors expire or are re-evaluated after 24 hours. A stored walk older than 23 hours, or one whose cursor the API rejects with HTTP 400, therefore restarts from its first page. The restarted walk keeps the records it already wrote, so they are deduplicated rather than written again.

The Hosted Runner gives each stanza a state bucket keyed by plugin, stanza name and `Ingester-UUID`; the plugin separates selected datasets by their resolved endpoint path. Keep the stanza name and UUID stable across restarts, upgrades and same-boundary key rotation. Adding a dataset or changing endpoint parameters creates or selects a different dataset checkpoint and can collect retained source records.

### Deduplication and delivery

Each checkpoint keeps a bounded manifest of record identities and content digests. An unchanged record is not written again; a changed record produces a new entry with the same identity. For time-filtered operations, an identity last seen before the next window is dropped once the walk completes. Otherwise a manifest holds at most 100,000 identities and forgets the least recently changed first, which can cause at most one rewrite of a settled record.

Missing inventory objects do not produce deletion events. Preserve vendor deletion fields such as `deleted_at`.

After every page the ingester validates the whole page before writing any of it. It then writes the page's new records and waits up to two minutes for the ingest muxer's synchronization. Finally it checks for cancellation and stores the checkpoint.

Synchronization drains exposed ingest channels and current connections; it is not an acknowledgement that the whole disk cache reached an indexer. State and cache are not one atomic transaction, so a failure can replay the affected page. Exactly-once delivery is not claimed. `State.Sync=true` flushes each state transaction.

Keep the ingest cache and state on durable storage. Never delete a cache while keeping a checkpoint that may depend on its queued entries.

### Pagination, limits and child work

- `after_id` endpoints (`activities`, `chats` and `chat-messages`) continue while `has_more` is true; an absent `has_more` means false.
- Page-token endpoints stop when `has_more` is false.
- The session endpoints, which omit `has_more`, stop when `next_page` is null.
- A repeated cursor, a malformed response or an exhausted bound fails the walk and keeps the last completed page's checkpoint.
- `Max-Pages` and a fixed ceiling of 100,000 records per call split large walks into checkpointed continuations that run immediately.

Child discovery works as follows:

- Organizations discover users and roles, then role permissions. Groups discover members. Chats and sessions discover messages. Projects discover attachments and collaborators.
- Deleted chats and projects retire their child work. A `pending` remote session is not queued until it starts.
- Children are revisited when the parent record changes, and otherwise hourly. A local session's transcript is revisited only when its `updated_at`, the last inference call, changes. Active remote sessions are revisited every poll.
- `Max-Children` limits child collections per cycle, not records or requests; one child can span many pages.
- `Max-Pending` bounds remembered child work. When it is full, the current root page waits for existing work instead of being partly written. Completed work is evicted first, and a child that has failed ten consecutive times can be evicted after that.
- A failing child is retried with exponential backoff and does not block other children. A temporary credential-file or configuration read failure preserves the queued child and its checkpoint for the next cycle.
- A parent identifier containing a path, query or fragment delimiter is rejected for child expansion with a warning; the root record and other children continue.
- A child that returns HTTP 404 after its root listed successfully has been removed at the source. It is logged and completed rather than retried.

### Rate limits and retries

Anthropic allows 600 requests per minute per parent organization. That budget is shared across every key, linked organization and Compliance endpoint. Remote-session endpoints also have a second budget of their own.

Each Claude Compliance stanza owns one rate limiter shared by all of its selected datasets and discovered children. Separate stanzas, processes and other clients do not coordinate automatically, so divide the parent organization's budget among them. The example's 30 requests per minute is an operating choice, not a vendor quota.

HTTP 429, transient 5xx, and connection or read failures are retried up to `Max-Retries` times. The delay starts at one second, doubles to a 60-second cap, and honors a longer `retry-after` only up to the configured `Request-Interval`. A value beyond that bound ends the current cycle without advancing the affected checkpoint. A response with `x-should-retry: false` is not retried. Authentication, permission and other 4xx failures are not retried within the poll. Resolve them without deleting state.

### Volume and persistence

Full-scan operations and first-time chat histories can be large regardless of `Lookback`. If a required source cannot drain, raise only the applicable bound (`Max-Pages`, `Max-Pending`, `Max-Children` or the request rate), or narrow the configured scope. Shrinking `Lookback` cannot narrow a full-parent scan.

Keep `/opt/gravwell/etc/hosted_runner.conf`, both secret files, `/opt/gravwell/state` and `/opt/gravwell/cache` through host maintenance and upgrades, under the same service identity.

## Validation

1. Validate and activate the configuration, then check the Hosted Runner diagnostics for every enabled stanza.
2. Select a known record, or perform a normal, approved action in a covered organization. Activities become queryable about a minute after they occur, then wait for the next poll. Use a search range that includes the source timestamp and any initial history.
3. Confirm tag arrival and well assignment:

```gravwell
tag=claude-compliance-activities,claude-compliance-directory,claude-compliance-conversations,claude-compliance-projects,claude-compliance-sessions,claude-compliance-artifacts
| stats count by TAG
| table TAG count
```

4. Confirm provenance and native fields:

```gravwell
tag=claude-compliance-activities
| intrinsic _source _parent
| json id type created_at
| limit 20
| table TIMESTAMP _source _parent id type created_at
```

```gravwell
tag=claude-compliance-directory,claude-compliance-conversations,claude-compliance-projects,claude-compliance-sessions
| intrinsic _source _parent
| stats count by TAG _source
| table TAG _source count
```

5. Match a native ID and timestamp with the record the source administrator provided. Each entry must be one JSON object.
6. Check for duplicates. Over a completed poll, this query should return no rows except records that genuinely changed:

```gravwell
tag=claude-compliance-activities
| json id
| stats count by id
| eval count > 1
| table id count
```

7. After a completed poll, record counts, then restart through the normal service procedure with the same stanza names, UUIDs, cache and state. New records must continue without a repeated full history.

An empty source can be valid, but it proves nothing about another source's access. The artifacts tag stays empty until explicitly configured.

## Troubleshooting

| Symptom | Cause | Fix |
| --- | --- | --- |
| Configuration rejected | Invalid UUID, unknown or duplicate `Dataset`, conflicting `Tag-Name` and `Tag-Prefix`, wrong `Parameter` names, or an unreadable credential file | Correct the named field. Keep UUIDs distinct when using multiple stanzas and use the exact placeholder names from the source table. |
| Permission denied opening a secret or state | File ownership or directory access | Check with `ls -l`, including parent directories. The service user needs read access to both secrets and write access to state and cache. Never make secrets world-readable. |
| HTTP 401 | Key unrecognized, deactivated or expired | Confirm the key type and the complete file contents, and check that the key was not revoked. |
| HTTP 404 `Not found` on every request | Request not authenticated: wrong key type, such as a Claude API key | Use a Compliance Access Key (or an Admin API key for activities only). |
| HTTP 403 | Missing scope | Check the key's **Scopes** column. Users and members need `read:compliance_user_data`. Scopes cannot be added, so create a replacement key. |
| Organization settings 404 | Organization outside the key's hierarchy, or not a bare UUID | Use the bare `uuid` from `/organizations`. |
| Child logged as not found | The chat, session or object was deleted or retained out, or a remote session was not yet running | No action is needed. The child is revisited when its parent changes. |
| Local-session 503 recurring for one organization | The source cannot return that content, for example because of an unusable customer-managed key or pending retention evaluation | Resolve it at the source. Other organizations and session metadata continue. |
| HTTP 429 or growing lag | The shared parent budget is exhausted by this or other clients | Coordinate with other clients, lower `Requests-Per-Minute`, and review collection lag. |
| Page or child-work limit reached repeatedly | A large full scan or backlog | Check which dataset is limited and raise only that bound. `Lookback` does not narrow discovery inventories. |
| Response exceeds the internal safety limit | A source page is unexpectedly larger than the ingester's fixed in-memory guard | Lower `Page-Size`. If one record alone remains too large, preserve state and contact the source or Gravwell support instead of discarding the record. |
| TLS or backend error | DNS, port, certificate trust or ingest secret | Repair trust and routing. Do not disable certificate verification. |
| Service runs but wells or dashboards stay empty | Source access, well routing, kit macros or query timeframe | Verify each independently with the validation queries. |
| Repeated records | Overlap or retry replay, changed source content, or a changed stanza name, UUID, selected dataset or endpoint parameter | Keep the checkpoint identity stable, and compare native `id` and content before investigating further. |
| Ingest or state synchronization failure | Backend, cache capacity or state disk | Repair it, keep both cache and state, and confirm accepted entries after recovery. |

To disable collection, stop the Hosted Runner or remove the Claude stanzas, and keep state and cache for recovery. To roll back an upgrade, restore the previous approved Hosted Runner release and configuration, then repeat validation. Restore an older state snapshot only as a deliberate replay decision.

## Related kit

The companion Claude Compliance kit provides the six tag macros, per-family AutoExtractors, an Investigations dashboard, overview dashboards, a source pivot and synthetic demonstration data. It is being prepared as `claude-compliance` in the [Gravwell kits repository](https://github.com/gravwell/kits). Until it appears there or in your Gravwell kit catalog, it is not available for installation.

When you install it, confirm these tag macros match the `Tag-Name` values in your configuration:

- `CLAUDE_COMPLIANCE` must resolve to the bounded Claude Compliance prefix covering all six tags.
- `CLAUDE_COMPLIANCE_ACTIVITIES`, `CLAUDE_COMPLIANCE_DIRECTORY`, `CLAUDE_COMPLIANCE_CONVERSATIONS`, `CLAUDE_COMPLIANCE_PROJECTS`, `CLAUDE_COMPLIANCE_SESSIONS` and `CLAUDE_COMPLIANCE_ARTIFACTS` must each match one tag.

The kit's category macro reads the `_source` intrinsic field. The ingester emits `_source` automatically; `_parent` and `_session` appear only where their discovery or envelope context exists. Use the demo variant only in a separate demonstration environment; its records are synthetic and carry no ingestion context.

## References

- [Claude Compliance API overview](https://platform.claude.com/docs/en/manage-claude/compliance-api)
- [Set up the Compliance API and create scoped keys](https://platform.claude.com/docs/en/manage-claude/compliance-api-access)
- [Activity Feed filters and pagination](https://platform.claude.com/docs/en/manage-claude/compliance-activity-feed)
- [Organization, user, role, group and settings endpoints](https://platform.claude.com/docs/en/manage-claude/compliance-org-data)
- [Chats, files, projects and artifacts](https://platform.claude.com/docs/en/manage-claude/compliance-content-data)
- [Local and remote session transcripts](https://platform.claude.com/docs/en/manage-claude/compliance-sessions)
- [Compliance integration patterns and retention](https://platform.claude.com/docs/en/manage-claude/compliance-integration-patterns)
- [Compliance API errors, rate limits and retry guidance](https://platform.claude.com/docs/en/manage-claude/compliance-errors)
- [Compliance API reference](https://platform.claude.com/docs/en/api/compliance)
- [Gravwell Hosted Runner configuration](https://docs.gravwell.io/configuration/configuration.html#hosted-runner-configuration)
- [Gravwell ingester configuration](https://docs.gravwell.io/ingesters/ingesters.html)
- [Gravwell storage wells](https://docs.gravwell.io/configuration/configuration.html)
- [Gravwell intrinsic module](https://docs.gravwell.io/search/intrinsic/intrinsic.html)
- [Gravwell kits](https://docs.gravwell.io/kits/kits.html)
