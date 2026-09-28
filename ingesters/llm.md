---
myst:
  substitutions:
    package: "gravwell-llm-ingester"
    standalone: "gravwell_llm_ingester"
    dockername: ""
---
# LLM Ingester

The LLM ingester is an HTTP proxy that sits between your LLM clients and the model provider. Clients are pointed at the ingester instead of at the provider; the ingester forwards each request upstream unmodified, streams the response back, and ingests the conversation as structured Gravwell entries along the way.

Two provider APIs are supported: the OpenAI Chat Completions API, spoken by most hosted providers and local model servers, and the Anthropic Messages API, spoken by Claude Code and the Anthropic SDKs. A listener speaks one of them, and both are normalized into the same set of ingested events.

Because the ingester is a proxy rather than a log collector, it captures data that is not normally written anywhere: the user's prompt, the system prompt, the model's reply, any reasoning the provider exposes, every tool call the model makes along with its arguments and results, and the token accounting for each response. Each event is a separate entry with enumerated values describing the model, session, request, and timing, which makes it possible to answer questions like "which tools is this agent actually calling", "how many tokens did this team spend today", and "what did anyone ask the model about production credentials".

```{note}
The LLM ingester pairs naturally with the [vector preprocessor](/ingesters/preprocessors/vector) and the [semantic](/search/semantic/semantic) search module: attach an embedding to each prompt and response at ingest time, then search the conversations by meaning rather than by keyword.
```

## Installation

```{include} installation_instructions_template 
```
(basic-configuration)=
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

A listener speaking the Anthropic Messages API is configured the same way, with a different `Protocol` and the auth style the Messages API expects:

```
[Listener "anthropic"]
	Bind = ":4181"
	Upstream-URL = "https://api.anthropic.com"
	Protocol = "anthropic-messages"
	Tag-Name = llm
	Auth-Style = "x-api-key"
	Anthropic-Version = "2023-06-01"
	Session-ID-Header = "x-claude-code-session-id"
	Log-Mode = "delta"
	Log-Tool-Calls = true
	Log-Usage = true
```

See [Anthropic Messages](llm_anthropic) for what each of those parameters does and [Claude Code](llm_claude_code) for pointing a client at it.

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
| Auth-Style                        | string       | NO       | `bearer`         | Which header carries the API credential: `bearer` for `Authorization: Bearer <token>`, or `x-api-key` for a bare `x-api-key: <token>`. See [Authorization](llm_authorization). |
| Client-Authorization              | string       | NO       |                  | Bare token that inbound clients must present in the header named by `Auth-Style`. Requests that do not match are rejected with a 401. When unset, no client authentication is required. |
| Upstream-Authorization            | string       | NO       |                  | Bare token injected into the header named by `Auth-Style` on the upstream request, replacing whatever credential the client sent. When unset, the client's own credential is passed through unchanged. |
| Anthropic-Version                 | string       | NO       |                  | Value injected as the `anthropic-version` header when a request arrives without one. Valid only on an `anthropic-messages` listener; setting it elsewhere prevents the ingester from starting. |
| Session-ID-Header                 | string       | NO       |                  | Name of a request header carrying the client's own conversation identifier, used as the `session_id` in place of a derived one. See [Session Tracking](llm_session_tracking). |
| Session-TTL                       | string       | NO       | `30m`            | How long idle session-matching state is retained, as a Go duration string, e.g. `"90m"`. Must be positive. |
| Allow-Unknown-Paths               | boolean      | NO       | false            | Forward every path to the upstream rather than answering 404 for paths the protocol neither parses nor declares a passthrough. Insecure; see [Protocols](llm_protocols). |
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

| Protocol             | Paths                   | Description |
|----------------------|-------------------------|-------------|
| `openai-chat`        | `/v1/chat/completions`  | OpenAI Chat Completions API, both buffered and streaming (server-sent events) responses. |
| `anthropic-messages` | `/v1/messages`          | Anthropic Messages API, both buffered and streaming (server-sent events) responses. This is what Claude Code and the Anthropic SDKs speak. |

The ingester logs the set of registered protocols to the `gravwell` tag at startup.

A provider's API is wider than the one endpoint a protocol parses, and a client will call the neighbors. Each protocol therefore declares the sibling endpoints its clients legitimately need, and the proxy forwards those untouched without ingesting them: the `anthropic-messages` protocol passes `/v1/messages/count_tokens` through, because Claude Code calls it on every request. Nothing needs configuring for this to work.

Requests to any other path receive a 404; the proxy is not a general-purpose forwarder. Setting `Allow-Unknown-Paths` on a listener forwards every path upstream instead.

```{warning}
`Allow-Unknown-Paths` is insecure. The proxy attaches the listener's `Upstream-Authorization` to whatever it forwards, so an open path list lets any client that can reach the listener aim it at a request the proxy never inspects and read the upstream credential back out of the result. Enable it only for a specific client that needs an endpoint the ingester does not know about, and give that client a listener of its own.
```

### openai-chat

The `openai-chat` protocol is compatible with any provider that implements the OpenAI Chat Completions API, which includes most hosted providers and local model servers such as Ollama, vLLM, and llama.cpp. Provider-specific request and response fields are forwarded untouched even when the ingester does not parse them.

Both the string and the multimodal array forms of the `content` field are handled. For the array form, text parts are extracted and joined with newlines so that the ingested entry reads the way the user wrote it. Content arrays that carry no text parts at all, such as an image-only turn, are ingested as the raw JSON so that nothing is silently dropped.

(llm_anthropic)=
### anthropic-messages

The `anthropic-messages` protocol handles Anthropic's Messages API. It is shaped differently from Chat Completions in several ways, all of which the ingester normalizes into the same [events](llm_ingested_events) produced by `openai-chat`:

* The system prompt is a top-level `system` field rather than a leading `system` message. It is ingested as a `request.system_message` event either way. Newer models also accept `system` entries inside the `messages` array, which Claude Code uses to inject tool and subagent context; those are ingested as system messages too.
* Message content is either a string or an array of typed content blocks. Text blocks are extracted and joined with newlines. Blocks the ingester does not model, such as images and documents, are folded into the entry as their raw JSON rather than dropped, which is bounded by `Max-Body` but can be bulky when a client inlines an image.
* There is no separate `tool` role. A tool call is a `tool_use` block on an assistant turn and a tool result is a `tool_result` block on a user turn; both are ingested as `response.tool_call` and `request.tool_result` respectively, correlated by `tool_call_id` as usual.
* Extended thinking arrives as `thinking` blocks and is ingested as `response.reasoning`, in generation order, so reasoning precedes the reply it belongs to.

Two listener parameters exist for this protocol. `Auth-Style` should be set to `x-api-key` for clients credentialed with an API key, since the Messages API authenticates with a bare `x-api-key` header rather than `Authorization: Bearer`; see [Authorization](llm_authorization). `Anthropic-Version` supplies the `anthropic-version` header the Messages API requires, and is injected only when a request arrives without one, which is what happens when the proxy holds the upstream key and the client was never configured to talk to Anthropic directly.

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
| `protocol`          | string  | Name of the protocol module that parsed the traffic, either `openai-chat` or `anthropic-messages`. Always present. |
| `listener`          | string  | Name of the listener config block that handled the request. Always present. |
| `upstream_status`   | int     | HTTP status code returned by the provider. Present on response-side events. |
| `duration_ms`       | int     | Milliseconds from receipt of the request to completion of the response. Present on response-side events. |
| `stream`            | boolean | Whether the client requested a streaming response. Always present. |
| `new_session`       | boolean | Attached, with the value `true`, to every event of a request that started a new session. Absent otherwise. |

Request-side events are ingested before the provider is contacted, so they carry no `upstream_status`, `request_id`, or `duration_ms`. Correlate them with the response using `session_id`.

### Token Usage and Streaming

Providers return the token accounting record at the end of a response. For buffered responses this is automatic. For streaming responses it depends on the protocol.

On an `openai-chat` listener, most providers omit the record from a stream unless the client asks for it. To capture usage on streamed requests, the client must set `stream_options.include_usage` to `true`; the ingester cannot add it because rewriting the request would change what the client receives.

```{note}
When `Log-Usage` is enabled but no usage records are appearing for streaming traffic from an OpenAI-compatible client, this is almost always the cause.
```

On an `anthropic-messages` listener the counts are part of the stream itself and need no opt-in: input and cache tokens arrive with the opening `message_start` event and the output count with the closing `message_delta`, and the ingester assembles them into a single `response.usage` event.

The Messages API reports cache reads and cache writes separately and reports no total. To keep `prompt_tokens`, `completion_tokens`, and `total_tokens` comparable across both protocols, the ingester folds the cache counts into the prompt count and computes the total:

```
prompt_tokens = input_tokens + cache_read_input_tokens + cache_creation_input_tokens
completion_tokens = output_tokens
total_tokens = prompt_tokens + completion_tokens
```

```{note}
Cached input tokens are billed at a different rate than uncached ones, so `prompt_tokens` from an Anthropic listener counts the tokens the model was given rather than tracking cost directly.
```

(llm_session_tracking)=
## Session Tracking

Neither the Chat Completions API nor the Messages API has a concept of a session, so the ingester derives one unless the client supplies it. Every event is stamped with a `session_id` that groups the turns of a single conversation, which is what makes it possible to reconstruct an exchange, follow an agent through a long chain of tool calls, or sum the token spend of one conversation.

### Client-Supplied Session IDs

Some clients stamp their own conversation identifier on every request. Naming that header with `Session-ID-Header` makes the ingester adopt the client's identifier as the `session_id` instead of deriving one, which is both more accurate and more useful: an identifier the client owns survives the history rewrites that prefix matching cannot follow, and it lines the ingested session up with the client's own transcript.

The parameter is not tied to a provider or a protocol. It works on any listener, with any header name, including one your own application sets. Claude Code sends `x-claude-code-session-id`:

```
[Listener "anthropic"]
	Bind = ":4181"
	Upstream-URL = "https://api.anthropic.com"
	Protocol = "anthropic-messages"
	Tag-Name = llm
	Auth-Style = "x-api-key"
	Session-ID-Header = "x-claude-code-session-id"
```

Values are held to what a session identifier plausibly is: at most 128 bytes of printable ASCII with no spaces. A request whose header is missing, empty, or fails that check falls back to prefix matching, and the listener logs a warning so a wrong header name degrades loudly rather than quietly.

### Derived Sessions

When no `Session-ID-Header` is configured, or a request arrives without a usable value, sessions are derived by matching each incoming request against the requests seen before it. The ingester hashes every message in a request and remembers the trailing window of hashes, bounded by `Session-Match-Window`. A later request continues an existing session when it carries that session's previous messages unchanged in the same positions; anything appended past that point — one new user turn, several stacked turns, tool results, or nothing at all in the case of a retry — still matches. When more than one session qualifies, the longest match wins as the most specific continuation. A request that matches nothing mints a new session, and its events are marked with `new_session`.

Session state, derived and client-supplied alike, is partitioned by client IP so that traffic from different clients never cross-matches, and is capped at 256 concurrent sessions per client, with the oldest dropped first. Entries idle for longer than `Session-TTL` are evicted. When several listeners are configured, the session store is shared and uses the longest `Session-TTL` among them.

Setting `State-Store-Location` in the `[Global]` block persists session state across restarts, so conversations in flight when the ingester restarts are not split into two sessions.

```{note}
Only message hashes are stored, in memory and on disk — never prompt or response content. The state file cannot be used to recover conversation text.
```

The client IP is taken from the leading entry of the `X-Forwarded-For` header when that value parses as an IP address, and from the connecting peer otherwise. This keeps sessions attributed to the real client when the ingester sits behind a load balancer, without letting a malformed header collapse every client onto one address. The ingester appends its own peer to the `X-Forwarded-For` chain on the upstream request.

(llm_authorization)=
## Authorization

The ingester handles two independent credentials. Both are configured as bare tokens, and `Auth-Style` decides which header carries them: `bearer` (the default) speaks `Authorization: Bearer <token>`, and `x-api-key` speaks a bare `x-api-key: <token>`. The style applies to both the client side and the upstream side of a listener.

`Client-Authorization` gates who may use the proxy. When set, inbound requests must present a matching credential in that header, and are rejected with a 401 before the request body is read. When unset, anyone who can reach the listener may use it.

`Upstream-Authorization` supplies the credential used against the provider. When set, it replaces whatever credential the client sent, which means the real provider API key never has to be distributed to clients — they authenticate to the proxy, and the proxy authenticates to the provider. When unset, the client's own credential is forwarded unchanged.

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

### Choosing an Auth-Style

`bearer` is what OpenAI-compatible providers expect and is the default, so an `openai-chat` listener normally needs no `Auth-Style` at all.

The Anthropic Messages API authenticates with `x-api-key`, so an `anthropic-messages` listener normally sets `Auth-Style = "x-api-key"`. The exception is a client credentialed with a bearer token rather than an API key, which Claude Code does when it is configured with `ANTHROPIC_AUTH_TOKEN` instead of `ANTHROPIC_API_KEY`; such a listener wants `Auth-Style = "bearer"`. The style has to match how the client is credentialed, because it names the header the client gate reads.

```{note}
When `Upstream-Authorization` is set, every credential header the client sent is stripped before the request is forwarded, not just the one `Auth-Style` names. The Messages API accepts both `x-api-key` and `Authorization: Bearer`, so leaving the other one in place would forward the client's own key alongside the configured one and let the client choose which account is billed.
```

(llm_tls)=
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

A client is pointed at the ingester by overriding the API base URL it talks to. No other change is required, and the client continues to talk to the same models. For an OpenAI-compatible client the base URL includes `/v1` (`http://<ingester>:4180/v1`); for Claude Code it does not, because Claude Code appends the whole path itself.

Whichever client is in use, the key it presents depends on which side holds the provider credential:

| | Listener | Client's key |
|---|---|---|
| Client's key passes through (default) | neither `Client-Authorization` nor `Upstream-Authorization` set | the real provider key |
| Proxy holds the real key | `Upstream-Authorization` set to the real key, optionally with `Client-Authorization` as a gate token | the gate token, or any placeholder when `Client-Authorization` is unset |

```{warning}
Over plain HTTP both the prompts and the API key cross the wire in the clear. For anything but a proxy on localhost, set `TLS-Certificate-File` and `TLS-Key-File` on the listener and give clients an `https://` base URL. See [Configuring TLS](llm_tls).
```

### curl

```
# openai-chat listener
curl http://localhost:4180/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -d '{"model":"gpt-4o-mini","messages":[{"role":"user","content":"hello"}]}'

# anthropic-messages listener
curl http://localhost:4181/v1/messages \
  -H "Content-Type: application/json" \
  -H "anthropic-version: 2023-06-01" \
  -H "x-api-key: $YOUR_ANTHROPIC_API_KEY" \
  -d '{"model":"claude-opus-5","max_tokens":64,"messages":[{"role":"user","content":"hello"}]}'
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

### Codex

Codex supports pointing at an alternate endpoint through its own configuration; see the [Codex configuration documentation](https://developers.openai.com/codex/config-advanced).

(llm_claude_code)=
### Claude Code

Claude Code talks to the Messages API, so it needs an `anthropic-messages` listener like the one shown under [Basic Configuration](#basic-configuration). Point it at that listener with the `ANTHROPIC_BASE_URL` environment variable, either per invocation:

```
ANTHROPIC_BASE_URL=http://localhost:4181 \
ANTHROPIC_API_KEY=$YOUR_ANTHROPIC_API_KEY \
  claude
```

or persistently in an `env` block in Claude Code's settings, in `.claude/settings.json` inside a project or `~/.claude/settings.json` to cover every project:

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "http://localhost:4181",
    "ANTHROPIC_API_KEY": "sk-ant-your-key-or-the-gate-token"
  }
}
```

A few things are worth knowing on the client side:

* The base URL carries no `/v1`. Claude Code appends the full path itself, and it covers every API call the client makes, including the `/v1/messages/count_tokens` sibling the listener passes through without ingesting.
* Claude Code always wants some credential in its environment and falls back to a claude.ai login when it finds none. When the proxy holds the real key, give the client the gate token rather than nothing.
* `ANTHROPIC_API_KEY` sends `x-api-key`, which is what `Auth-Style = "x-api-key"` expects. `ANTHROPIC_AUTH_TOKEN` sends `Authorization: Bearer` instead and needs `Auth-Style = "bearer"` on the listener.
* Setting `Session-ID-Header = "x-claude-code-session-id"` on the listener stamps each entry with Claude Code's own session identifier, which is what makes an ingested conversation line up with the client's transcript. See [Session Tracking](llm_session_tracking).

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
