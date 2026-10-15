# Claude Compliance

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, Claude Compliance Hosted Ingester
:::

## Claude Compliance Configuration

Anthropic provides a **Compliance API** for Claude Enterprise and Claude Console customers. See [Compliance API](https://platform.claude.com/docs/en/manage-claude/compliance-api). The Claude Compliance ingester polls this API and ingests each JSON record as a Gravwell entry.

To configure the ingester you will need the following from Anthropic:

* **API key**: Either a **Compliance Access Key** (created in claude.ai, reaches every endpoint) or an **Admin API key** (reaches the Activity Feed only). Configure it inline with `Credential` or, preferably, in a protected file with `Credential-File`.

What each key can read determines which datasets you can ingest:

* **Activity Feed** (`activities` dataset): admin/system and resource events (workspace changes, API key creation, file operations). Available with either key type.
* **Everything else** (directory, organization settings, chat/file/project content, and session transcripts): Claude Enterprise only, and requires a Compliance Access Key.

```{attention}
The chat, project, and session datasets contain conversation content. Ingest only the datasets you need, and restrict access to the resulting well with [CBAC](/cbac/cbac).
```

Use a dedicated key for Gravwell. The ingester is read-only: it does not download binary files, delete vendor content, or acknowledge vendor records.

A separate, narrower **manual export** also exists: Organization Owners and Primary Owners can export audit logs via **claude.ai > Organization settings > Data and Privacy**, capped at a 180-day lookback and delivered as a downloadable file via an emailed link. This excludes chat and project content and title text and cannot be automated, so it is not a feed for continuous ingestion.

There is no webhook or push mechanism for the Compliance API; everything is pull-based, which is why the ingester polls.

## Gravwell Configuration

### Gravwell Storage Well Configuration

Set up the well configuration in your Gravwell indexers.

#### Sample Well Config
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/claude_compliance-well.conf`
```ini
[Storage-Well "claude_compliance"]
    Location=/opt/gravwell/storage/claude_compliance
    Tags=claude-compliance*
```

### Gravwell Ingester Configuration: Claude Compliance Hosted Ingester

The ingester runs as a plugin inside the [Gravwell Hosted Runner](hosted_runner_configuration). If the Hosted Runner is not installed, install it first. The `[Global]` and `[State]` blocks common to all Hosted Runner plugins are described in [Hosted Runner Configuration](hosted_runner_configuration).

Each stanza gets its checkpoint namespace from its quoted name and `Ingester-UUID`; each selected dataset keeps separate state within that namespace. No scope label is configured.

```{note}
The Claude Compliance plugin is still in development (see [gravwell/gravwell#2783](https://github.com/gravwell/gravwell/pull/2783)) and is not yet available in a released Hosted Runner. The options below come from that development version and may change before release.
```

Prefer a protected credential file containing the API key on one line, readable only by the Hosted Runner user (mode `0600` is recommended). The file must be a regular file of 16 KB or less, and the key must not contain line breaks. `Credential="REPLACE_WITH_COMPLIANCE_KEY"` is the supported inline alternative; configure only one credential form.

Edit: `/opt/gravwell/etc/hosted_runner.conf`

#### Sample Claude Compliance Config: Activity Feed

```ini
[ClaudeCompliance "activities"]
    Ingester-UUID="99f00000-0000-0000-0000-000000000000"
    Dataset="activities"
    Credential-File="/opt/gravwell/secrets/claude-compliance-key"
```

#### Sample Claude Compliance Config: Parameterized Dataset
Datasets that address a specific resource need a `Parameter` for each placeholder in their API path, written as `name:id`. For example, the users of a single organization:

```ini
[ClaudeCompliance "org-users"]
    Ingester-UUID="99f10000-0000-0000-0000-000000000000"
    Dataset="organization-users"
    Parameter="organization_id:your-organization-id"
    Credential-File="/opt/gravwell/secrets/claude-compliance-key"
```

#### Datasets and tags

Repeat `Dataset` in a stanza to select one or more of the 26 operations below. Each operation keeps its own checkpoint and routes to its family's default tag. `Tag-Prefix` defaults to `claude-compliance`; `Tag-Name` instead pins every selected operation in that stanza to one explicit tag.

| Family and default tag | Configure without `Parameter` | Collected automatically when `Follow-Children="enabled"` | Configure explicitly with `Parameter` |
| --- | --- | --- | --- |
| Activity Feed — `claude-compliance-activities` | `activities` | None | None |
| Directory — `claude-compliance-directory` | `organizations`, `groups` | `organization-users`, `organization-roles`, `role-permissions`, `group-members` | `organization-role` (`organization_id`, `role_id`); `organization-settings` (`organization_id`); `group` (`group_id`) |
| Conversations — `claude-compliance-conversations` | `chats` | `chat-messages` | `file-metadata` (`file_id`); `generated-file-metadata` (`file_id`) |
| Projects — `claude-compliance-projects` | `projects` | `project-attachments`, `project-collaborators` | `project` (`project_id`); `project-document` (`document_id`); `project-document-metadata` (`document_id`) |
| Artifacts — `claude-compliance-artifacts` | None | None | `artifact-metadata` (`artifact_version_id`) |
| Sessions — `claude-compliance-sessions` | `local-sessions`, `remote-sessions` | `local-session-messages`, `remote-session-messages` | `local-session` (`session_id`) |

The seven operations in the second column are roots. With `Follow-Children` at its default of `enabled`, five root families discover the nine child operations in the third column. The ten operations in the last column are not discovered because the API does not provide a safe global inventory for their IDs; configure them only when you have the required ID and key scope.

Every native JSON record remains unwrapped. The ingester adds only these intrinsic fields:

* `_source`: the exact dataset operation, which distinguishes operations sharing a family tag.
* `_parent`: the bound parent parameters for parameterized and discovered child requests.
* `_session`: the compact session envelope on session message entries when the API provides it.

```{note}
Give every `[ClaudeCompliance]` stanza its own `Ingester-UUID`, and keep it unchanged after collection begins. One stanza may repeat `Dataset`; all selected operations share that stanza's polling settings and request-rate budget.
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_hosted_runner.service`
```

## Additional Resources
* [Compliance API](https://platform.claude.com/docs/en/manage-claude/compliance-api)
* [Compliance API access and key scopes](https://platform.claude.com/docs/en/manage-claude/compliance-api-access)
* [Activity Feed](https://platform.claude.com/docs/en/manage-claude/compliance-activity-feed)
* [Content data](https://platform.claude.com/docs/en/manage-claude/compliance-content-data)
* [Sessions](https://platform.claude.com/docs/en/manage-claude/compliance-sessions)
* [Admin API](https://platform.claude.com/docs/en/manage-claude/admin-api)
* [Usage & Cost API](https://platform.claude.com/docs/en/manage-claude/usage-cost-api)
* [Access audit logs (manual export)](https://support.claude.com/en/articles/9970975-access-audit-logs)
