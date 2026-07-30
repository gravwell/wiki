---
myst:
  substitutions:
    package: "gravwell-llm-ingester"
    standalone: "gravwell_llm_ingester"
    dockername: ""
---
# LLM Ingester

The LLM ingester is an OpenAI-compatible HTTP proxy that sits between your LLM clients and the model provider. Clients are pointed at the ingester instead of at the provider; the ingester forwards each request upstream unmodified, streams the response back, and ingests the conversation as structured Gravwell entries along the way.

Because the ingester is a proxy rather than a log collector, it captures data that is not normally written anywhere: the user's prompt, the system prompt, the model's reply, any reasoning the provider exposes, every tool call the model makes along with its arguments and results, and the token accounting for each response. Each event is a separate entry with enumerated values describing the model, session, request, and timing, which makes it possible to answer questions like "which tools is this agent actually calling", "how many tokens did this team spend today", and "what did anyone ask the model about production credentials".

```{note}
The LLM ingester pairs naturally with the [vector preprocessor](/ingesters/preprocessors/vector) and the [semantic](/search/semantic/semantic) search module: attach an embedding to each prompt and response at ingest time, then search the conversations by meaning rather than by keyword.
```

## Installation

```{include} installation_instructions_template 
```

## Basic Configuration

The LLM ingester uses the unified global configuration block described in the [ingester section](ingesters_global_configuration_parameters). Like most other Gravwell ingesters, the LLM ingester supports multiple upstream indexers, TLS, cleartext, and named pipe connections, a local cache, and local logging.

The configuration file is at `/opt/gravwell/etc/gravwell_llm_ingester.conf`. The ingester will also read configuration snippets from its [configuration overlay directory](configuration_overlays) (`/opt/gravwell/etc/gravwell_llm_ingester.conf.d`).

At least one `Listener` block must be defined. Each listener binds a port, forwards to one upstream provider, and ingests into one tag. Here is a minimal configuration that proxies OpenAI and ingests into the tag `llm`:

```
[Global]
	Ingest-Secret = IngestSecrets
	Connection-Timeout = 0
	Insecure-Skip-TLS-Verify = false
	Pipe-Backend-Target=/opt/gravwell/comms/pipe #a named pipe connection, this should be used when the ingester is on the same machine as a backend
	Log-Level = INFO
	Log-File = /opt/gravwell/log/llm_ingester.log
	State-Store-Location = /opt/gravwell/etc/llm_ingester.state

[Listener "openai"]
	Bind = ":4180"
	Upstream-URL = "https://api.openai.com"
	Protocol = "openai-chat"
	Tag-Name = llm
	Log-Mode = "delta"
	Log-Tool-Calls = true
	Log-Usage = true
	Session-TTL = "30m"
```

Clients then use `http://<ingester>:4180/v1` as their API base URL. See [Client Configuration](llm_client_configuration) for examples.

### Additional Global Configuration Parameters

| Config Parameter     | Type    | Required | Default Value | Description |
|----------------------|---------|----------|---------------|-------------|
| State-Store-Location | string  | NO       |               | Path to a persistent state file for the session tracker. When set, session state survives an ingester restart. Only message hashes are written, never prompt content. Leaving this unset keeps session state in memory only. |
| Session-Match-Window | integer | NO       | 10            | How many of a request's most recent messages the session tracker compares when deciding whether a request continues an existing conversation. Larger values are more precise but do more work per request. |

## Listener Configuration

Listener blocks support the following configuration parameters:

| Parameter                         | Type         | Required | Default Value    | Description |
|-----------------------------------|--------------|----------|------------------|-------------|
| Bind                              | string       | YES      |                  | Host:port pair the proxy listens on, e.g. `":4180"` or `"127.0.0.1:4180"`. Each listener must bind a unique address. |
| Upstream-URL                      | string       | YES      |                  | Base URL requests are forwarded to, e.g. `"https://api.openai.com"`. Must use the `http` or `https` scheme and include a host. |
| Protocol                          | string       | YES      |                  | Protocol module used to parse traffic. See [Protocols](llm_protocols). |
| Tag-Name                          | string       | NO       | `default`        | Tag assigned to ingested entries. |
| Log-Mode                          | string       | NO       | `delta`          | How much of each request is ingested: `delta`, `user`, or `full`. See [Log Modes](llm_log_modes). |
| Log-Tool-Calls                    | boolean      | NO       | false            | Capture tool invocations made by the model and the tool results sent back by the client. |
| Log-Usage                         | boolean      | NO       | false            | Ingest the token accounting record returned with each response. |
| Client-Authorization              | string       | NO       |                  | Bearer token that inbound clients must present as `Authorization: Bearer <token>`. Requests that do not match are rejected with a 401. When unset, no client authentication is required. |
| Upstream-Authorization            | string       | NO       |                  | Bearer token injected as the `Authorization` header on the upstream request, replacing whatever the client sent. When unset, the client's own `Authorization` header is passed through unchanged. |
| Session-TTL                       | string       | NO       | `30m`            | How long idle session-matching state is retained, as a Go duration string, e.g. `"90m"`. Must be positive. |
| Max-Body                          | integer      | NO       | 16777216 (16MB)  | Maximum size of an inbound request body. Larger requests are rejected with a 413. |
| TLS-Certificate-File              | string       | NO       |                  | Certificate PEM file used to run the proxy listener as HTTPS. Must be set together with `TLS-Key-File`. |
| TLS-Key-File                      | string       | NO       |                  | Key PEM file used to run the proxy listener as HTTPS. Must be set together with `TLS-Certificate-File`. |
| Insecure-Skip-TLS-Verify-Upstream | boolean      | NO       | false            | Do not verify the upstream provider's TLS certificate. Intended for lab use with self-signed certificates. |
| Preprocessor                      | string array | NO       |                  | Set of [preprocessors](/ingesters/preprocessors/preprocessors) to apply to ingested entries. |

Multiple listeners can be defined to proxy several providers, or to expose the same provider with different capture settings. Each listener maintains its own tag, log mode, credentials, and preprocessor chain:

```
[Listener "openai"]
	Bind = ":4180"
	Upstream-URL = "https://api.openai.com"
	Protocol = "openai-chat"
	Tag-Name = llm-openai
	Log-Mode = "delta"
	Log-Tool-Calls = true
	Log-Usage = true

[Listener "local"]
	Bind = ":4181"
	Upstream-URL = "http://127.0.0.1:11434"
	Protocol = "openai-chat"
	Tag-Name = llm-local
	Log-Mode = "full"
	Log-Usage = true
```

(llm_protocols)=
## Protocols

A protocol module tells the ingester how to parse the traffic flowing through a listener and which URL paths to expose. The following protocols are supported:

| Protocol      | Paths                   | Description |
|---------------|-------------------------|-------------|
| `openai-chat` | `/v1/chat/completions`  | OpenAI Chat Completions API, both buffered and streaming (server-sent events) responses. |

Requests to any path the protocol does not claim receive a 404; the proxy is not a general-purpose forwarder. The ingester logs the set of registered protocols to the `gravwell` tag at startup.

The `openai-chat` protocol is compatible with any provider that implements the OpenAI Chat Completions API, which includes most hosted providers and local model servers such as Ollama, vLLM, and llama.cpp. Provider-specific request and response fields are forwarded untouched even when the ingester does not parse them.

Both the string and the multimodal array forms of the `content` field are handled. For the array form, text parts are extracted and joined with newlines so that the ingested entry reads the way the user wrote it. Content arrays that carry no text parts at all, such as an image-only turn, are ingested as the raw JSON so that nothing is silently dropped.

(llm_log_modes)=
## Log Modes

An LLM conversation is stateless on the wire: each request re-sends the entire conversation so far. Logging every request in full therefore re-ingests the whole conversation on every turn, which is occasionally what you want and usually not. `Log-Mode` controls that tradeoff.

| Mode    | Description |
|---------|-------------|
| `delta` | (Default) Ingest only what is new in each request: the latest user turn, plus any tool results the client appended after the model's last reply. The system prompt is ingested once, on the first request of a new session. This produces one entry per logical event and does not duplicate earlier turns. |
| `user`  | Ingest only the most recent user message from each request. Assistant replies, system prompts, and tool traffic on the request side are skipped. Response-side events still honor `Log-Tool-Calls` and `Log-Usage`. |
| `full`  | Ingest every message in every request body, including all previously-logged turns. Useful for capturing a complete conversation snapshot per request, at the cost of substantial duplication. |

Response-side events — the model's reply, its reasoning, its tool calls, and the usage record — are ingested in all three modes, subject to `Log-Tool-Calls` and `Log-Usage`.

```{note}
`delta` mode relies on session matching to know which turns have already been ingested. See [Session Tracking](llm_session_tracking).
```

(llm_ingested_events)=
## Ingested Events

Each logical event becomes its own entry. The entry's DATA field holds the text of the event, and the metadata is attached as [intrinsic enumerated values](intrinsic_enumerated_values). The SRC field is set to the client's IP address, and the timestamp is the time the event was ingested.

| `event_type`                 | DATA contents | Gated by |
|------------------------------|---------------|----------|
| `request.user_message`       | Text of the user's message. | |
| `request.system_message`     | Text of the system prompt. | |
| `request.tool_result`        | Result the client returned for a tool call. | `Log-Tool-Calls` |
| `response.reasoning`         | Reasoning text, when the provider exposes it. | |
| `response.assistant_message` | Text of the model's reply. | |
| `response.tool_call`         | JSON arguments the model passed to the tool. | `Log-Tool-Calls` |
| `response.usage`             | Empty; the counts are carried as enumerated values. | `Log-Usage` |

Which request-side events are ingested also depends on `Log-Mode`, as described under [Log Modes](llm_log_modes).

```{note}
Reasoning is captured from either the `reasoning` or the `reasoning_content` field, since providers disagree on the name. Reasoning events are emitted before the reply they precede.
```

### Enumerated Values

| Enumerated Value    | Type    | Description |
|---------------------|---------|-------------|
| `event_type`        | string  | The event type, as listed under [Ingested Events](llm_ingested_events). Always present. |
| `role`              | string  | Message role: `user`, `system`, `assistant`, or `tool`. |
| `tool_name`         | string  | Name of the tool the model invoked. Present on `response.tool_call`. |
| `tool_call_id`      | string  | Provider-assigned identifier correlating a tool call with its result. |
| `prompt_tokens`     | int     | Tokens consumed by the request. Present on `response.usage`. |
| `completion_tokens` | int     | Tokens generated in the response. Present on `response.usage`. |
| `total_tokens`      | int     | Total tokens billed for the exchange. Present on `response.usage`. |
| `session_id`        | string  | Identifier of the conversation this event belongs to. See [Session Tracking](llm_session_tracking). |
| `request_id`        | string  | Provider-assigned response identifier. Present on response-side events. |
| `model`             | string  | Model name, taken from the response when the provider reports one and from the request otherwise. |
| `protocol`          | string  | Name of the protocol module that parsed the traffic, e.g. `openai-chat`. Always present. |
| `listener`          | string  | Name of the listener config block that handled the request. Always present. |
| `upstream_status`   | int     | HTTP status code returned by the provider. Present on response-side events. |
| `duration_ms`       | int     | Milliseconds from receipt of the request to completion of the response. Present on response-side events. |
| `stream`            | boolean | Whether the client requested a streaming response. Always present. |
| `new_session`       | boolean | Attached, with the value `true`, to every event of a request that started a new session. Absent otherwise. |

Request-side events are ingested before the provider is contacted, so they carry no `upstream_status`, `request_id`, or `duration_ms`. Correlate them with the response using `session_id`.

### Token Usage and Streaming

Providers return the token accounting record at the end of a response. For buffered responses this is automatic, but for streaming responses most providers omit it unless the client asks for it. To capture usage on streamed requests, the client must set `stream_options.include_usage` to `true`; the ingester cannot add it because rewriting the request would change what the client receives.

```{note}
When `Log-Usage` is enabled but no usage records are appearing for streaming traffic, this is almost always the cause.
```

(llm_session_tracking)=
## Session Tracking

The Chat Completions API has no concept of a session, so the ingester derives one. Every event is stamped with a `session_id` that groups the turns of a single conversation, which is what makes it possible to reconstruct an exchange, follow an agent through a long chain of tool calls, or sum the token spend of one conversation.

Sessions are derived by matching each incoming request against the requests seen before it. The ingester hashes every message in a request and remembers the trailing window of hashes, bounded by `Session-Match-Window`. A later request continues an existing session when it carries that session's previous messages unchanged in the same positions; anything appended past that point — one new user turn, several stacked turns, tool results, or nothing at all in the case of a retry — still matches. When more than one session qualifies, the longest match wins as the most specific continuation. A request that matches nothing mints a new session, and its events are marked with `new_session`.

Session state is partitioned by client IP so that traffic from different clients never cross-matches, and is capped at 256 concurrent sessions per client, with the oldest dropped first. Entries idle for longer than `Session-TTL` are evicted. When several listeners are configured, the session store is shared and uses the longest `Session-TTL` among them.

Setting `State-Store-Location` in the `[Global]` block persists session state across restarts, so conversations in flight when the ingester restarts are not split into two sessions.

```{note}
Only message hashes are stored, in memory and on disk — never prompt or response content. The state file cannot be used to recover conversation text.
```

The client IP is taken from the leading entry of the `X-Forwarded-For` header when that value parses as an IP address, and from the connecting peer otherwise. This keeps sessions attributed to the real client when the ingester sits behind a load balancer, without letting a malformed header collapse every client onto one address. The ingester appends its own peer to the `X-Forwarded-For` chain on the upstream request.

## Authorization

The ingester handles two independent credentials, and both are configured as bearer tokens; the proxy speaks `Bearer` to the client and to the provider on your behalf.

`Client-Authorization` gates who may use the proxy. When set, inbound requests must present `Authorization: Bearer <token>` matching the configured value, and are rejected with a 401 before the request body is read. When unset, anyone who can reach the listener may use it.

`Upstream-Authorization` supplies the credential used against the provider. When set, it replaces whatever `Authorization` header the client sent, which means the real provider API key never has to be distributed to clients — they authenticate to the proxy, and the proxy authenticates to the provider. When unset, the client's own `Authorization` header is forwarded unchanged.

Setting both turns the ingester into a gateway for the provider key:

```
[Listener "gateway"]
	Bind = ":4180"
	Upstream-URL = "https://api.openai.com"
	Protocol = "openai-chat"
	Tag-Name = llm
	Client-Authorization = "token-clients-must-present"
	Upstream-Authorization = "sk-your-real-provider-api-key"
```

## Configuring TLS

By default a listener runs a cleartext HTTP server. To run it as HTTPS, provide a certificate and key PEM file with the `TLS-Certificate-File` and `TLS-Key-File` parameters. Both must be set, and the keypair is validated at startup, so a bad path or a mismatched pair prevents the ingester from starting rather than failing later.

```
[Listener "openai"]
	Bind = ":4180"
	Upstream-URL = "https://api.openai.com"
	Protocol = "openai-chat"
	Tag-Name = llm
	TLS-Certificate-File = /opt/gravwell/etc/llm_ingester.crt
	TLS-Key-File = /opt/gravwell/etc/llm_ingester.key
```

TLS on the listener is separate from TLS to the provider. Connections to the upstream URL are verified normally; `Insecure-Skip-TLS-Verify-Upstream` disables that verification and should be used only against a lab endpoint with a self-signed certificate.

(llm_client_configuration)=
## Client Configuration

Any OpenAI-compatible client can be pointed at the ingester by overriding its API base URL to `http://<ingester>:<port>/v1`. No other change is required, and the client continues to talk to the same models.

### curl

```
curl http://localhost:4180/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -d '{"model":"gpt-4o-mini","messages":[{"role":"user","content":"hello"}]}'
```

### opencode

Add a custom provider pointing at the listener to `opencode.json` or `opencode.jsonc`:

```
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "llm-ingester": {
      "name": "Gravwell LLM Ingester",
      "options": {
        "baseURL": "http://localhost:4180/v1"
      },
      "models": {
        "<your-model>": {}
      }
    }
  }
}
```

### crush

Add a custom provider to `crush.json` or `.crush.json`:

```
{
  "$schema": "https://charm.land/crush.json",
  "providers": {
    "llm-ingester": {
      "name": "Gravwell LLM Ingester",
      "type": "openai",
      "base_url": "http://localhost:4180/v1",
      "api_key": "$YOUR_API_KEY",
      "models": [
        { "id": "<your-model>", "name": "<your-model>" }
      ]
    }
  }
}
```

### Codex and Claude Code

Both tools support pointing at an alternate endpoint through their own configuration; see the [Codex configuration documentation](https://developers.openai.com/codex/config-advanced) and the [Claude Code network configuration documentation](https://code.claude.com/docs/en/network-config).

## Attaching Embeddings

Because every ingested event is a discrete entry of natural-language text, LLM traffic is a good fit for embedding at ingest time. Adding a [vector preprocessor](/ingesters/preprocessors/vector) to a listener attaches an `embeddings` enumerated value to each prompt, reply, and tool call, which can then be searched by meaning with the [semantic](/search/semantic/semantic) module:

```
[Listener "openai"]
	Bind = ":4180"
	Upstream-URL = "https://api.openai.com"
	Protocol = "openai-chat"
	Tag-Name = llm
	Preprocessor = embed

[Preprocessor "embed"]
	Type=vector
	Model="text-embedding-3-small"
	Endpoint="https://api.example.com/v1/embeddings"
	Token=`sk-example-token`
```

```{warning}
Embedding requests are made synchronously as entries flow through, so the throughput of the listener becomes bound to the throughput of the embeddings endpoint. Consider whether that endpoint meters requests or charges per token before enabling this on busy traffic.
```

## Example Queries

Show the conversation for a single session in order:

```gravwell
tag=llm intrinsic session_id == "0f0d5f13-2c9e-4a97-9b0e-0f2e0f19a37e" event_type role | sort by TIMESTAMP asc | table TIMESTAMP event_type role DATA
```

Count the tools an agent is calling:

```gravwell
tag=llm intrinsic event_type == "response.tool_call" tool_name | count by tool_name | sort by count desc | table tool_name count
```

Total token spend by model:

```gravwell
tag=llm intrinsic event_type == "response.usage" model total_tokens | stats sum(total_tokens) by model | table model sum
```

Chart prompt volume over time:

```gravwell
tag=llm intrinsic event_type == "request.user_message" | count | chart count
```

Find the slowest exchanges:

```gravwell
tag=llm intrinsic event_type == "response.assistant_message" duration_ms model | sort by duration_ms desc | table duration_ms model DATA
```

Search prompts by meaning rather than by keyword, using embeddings attached by the [vector preprocessor](/ingesters/preprocessors/vector):

```gravwell
tag=llm intrinsic event_type == "request.user_message" embeddings | semantic -t 70 "asking the model for production credentials" | sort by score desc | table score DATA
```
