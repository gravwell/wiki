# semantic

The `semantic` module performs LLM-based semantic similarity search. Rather than matching literal text like [grep](/search/grep/grep) or [words](/search/words/words), `semantic` compares the *meaning* of a search phrase against the meaning of each entry. This lets a query for "user could not log in" match an entry that reads "authentication failure for account jsmith" even though the two share no words.

The module works on [vector embeddings](https://en.wikipedia.org/wiki/Sentence_embedding). At query time, `semantic` sends your search phrase to a configured embedding model and receives a vector back. Each entry in the pipeline is expected to already carry its own embedding vector in an enumerated value. The module computes the [cosine similarity](https://en.wikipedia.org/wiki/Cosine_similarity) between the search vector and each entry's vector, attaches that similarity as a score, and drops any entry scoring below a threshold.

```{note}
The `semantic` module requires an embedding model to be configured on the webserver. See [Embedding Model Configuration](#semantic-embedding-config) below.
```

## Prerequisites

Two things must be in place before `semantic` can be used:

1. **An embedding model must be configured on the webserver.** The `[AI]` section of `gravwell.conf` provides the endpoint, model name, and optional token used to embed the search phrase. See [Embedding Model Configuration](#semantic-embedding-config).

2. **Entries must already have embedding vectors attached.** The `semantic` module does not generate embeddings for entries; it only generates one for your search phrase. Entry embeddings are normally produced at ingest time by the [vector preprocessor](/ingesters/preprocessors/vector), which calls an embedding model for each entry and attaches the result as an intrinsic enumerated value named `embeddings`.

Because `embeddings` is an intrinsic enumerated value created at ingest, it must be pulled into the pipeline with the [intrinsic](/search/intrinsic/intrinsic) module before `semantic` can read it:

```gravwell
tag=notes intrinsic embeddings | semantic "disk is full" | table
```

## Syntax

```
semantic [options] <search phrase>
```

The search phrase is a single string. Phrases containing spaces must be quoted.

## Supported Options

* `-vector <name>`: Read each entry's embedding vector from the named enumerated value. Defaults to `embeddings`, which is the name used by the [vector preprocessor](/ingesters/preprocessors/vector).
* `-score <name>`: Set the name of the enumerated value that receives the similarity score. Defaults to `score`.
* `-t <n>`: Set the score threshold as a whole percentage from 0 to 100. Entries whose similarity score is below `n/100` are dropped. Defaults to 75.
* `-p`: Permissive mode. Pass through entries that have no embedding vector, whose vector cannot be parsed, or which cannot be scored, instead of dropping them or failing the query. Passed-through entries do not receive a score enumerated value.

## Produced Enumerated Values

| Enumerated Value | Default Name | Description |
|------------------|--------------|-------------|
| Score | `score` | Float in the range -1.0 to 1.0 giving the cosine similarity between the search phrase and the entry. Higher is more similar. |

Only entries that meet or exceed the threshold receive a score, so every entry leaving `semantic` has a score at or above `-t` -- with the exception of entries passed through by `-p`, which have no score at all.

## Entry Handling

Without `-p`, the module handles problem entries as follows:

| Condition | Result |
|-----------|--------|
| Entry has no vector enumerated value | Entry is dropped |
| Score is below the threshold | Entry is dropped |
| Vector is not a valid JSON array of numbers | The query fails with an error |
| Vector cannot be compared to the search vector, e.g. because it has a different number of dimensions | The query fails with an error |

With `-p`, all four cases pass the entry down the pipeline unscored instead. This makes `-p` useful when only some of the data in a query carries embeddings, or when embeddings were generated inconsistently.

## Choosing a Threshold

Cosine similarity is a measure of the angle between two vectors, not a probability. A score of 1.0 means the two vectors point in exactly the same direction and 0.0 means they are unrelated. In practice, the useful range depends heavily on the embedding model in use; some models cluster all real text between 0.6 and 0.9, while others spread scores much wider.

The most reliable way to pick a threshold for your data and model is to start permissive, sort by score, and look at where the results stop being relevant:

```gravwell
tag=notes intrinsic embeddings | semantic -t 0 "disk is full" | sort by score desc | table score DATA
```

Then set `-t` to that cutoff for production queries, dashboards, and alerts.

## Performance Considerations

Embedding vectors can be large. A model producing 1024 dimensions yields an enumerated value on the order of 10KB to 20KB per entry, and because `semantic` runs on the webserver, every candidate entry's vector must be sent from the indexers to the webserver before it can be scored. This makes `semantic` considerably more expensive than a text search over the same data. Consider filtering as much data out as possible before invoking the module.

(semantic-embedding-config)=
## Embedding Model Configuration

`semantic` embeds the search phrase by calling an [OpenAI-compatible](https://platform.openai.com/docs/api-reference/embeddings) `/v1/embeddings` endpoint. The endpoint is configured in the `[AI]` section of `gravwell.conf` on the webserver:

```
[AI]
	Embedding-URL="https://ai.example.com/v1/embeddings"
	Embedding-Model="qwen3-embedding:8b"
	Embedding-Token="sk-abc123"
```

| Directive | Required | Description |
|-----------|----------|-------------|
| `Embedding-URL` | Yes | Full URL of the OpenAI-compatible embeddings endpoint. Setting this directive is what enables the `semantic` module. |
| `Embedding-Model` | Yes | Name of the embedding model to request. Required whenever `Embedding-URL` is set; the webserver will refuse to start without it. |
| `Embedding-Token` | No | Bearer token sent in the `Authorization` header. Omit for endpoints that do not require authentication. |

See [the configuration parameters documentation](/configuration/parameters) for the full `[AI]` section reference. These directives are independent of the `Enable` directive, which controls only [Logbot AI](/search/ai/ai) chat features; semantic search can be used with Logbot disabled.

```{attention}
The model configured here must be the same model that generated the entries' embeddings at ingest time. Embeddings from different models are not comparable, and vectors of differing dimensions cannot be scored at all -- a mismatch will cause the query to fail (or, with `-p`, to pass every entry through unscored).
```

Requests to the embedding endpoint are subject to a 10 second timeout. A single request is made per query when the query starts.

## Examples

### Find log entries about a concept

Search notes for anything semantically similar to a phrase, using the default threshold of 75:

```gravwell
tag=notes intrinsic embeddings | semantic "the customer wants a refund" | table DATA
```

### Rank results by relevance

Loosen the threshold and sort the strongest matches to the top:

```gravwell
tag=tickets intrinsic embeddings | semantic -t 40 "cannot connect to the VPN" | sort by score desc | table score DATA
```

### Use a custom vector and score name

Read vectors from an enumerated value named `vec` and write the similarity to `relevance`:

```gravwell
tag=docs intrinsic vec | semantic -vector vec -score relevance -t 60 "quarterly revenue growth" | sort by relevance desc | table relevance DATA
```

### Keep unembedded entries in the pipeline

In permissive mode, entries with no embedding vector continue down the pipeline unscored rather than being dropped:

```gravwell
tag=notes intrinsic embeddings | semantic -p -t 80 "authentication failure" | count
```

### Combine text filtering with semantic scoring

Narrow the candidate set with a cheap word match on the indexers, then score only what survives:

```gravwell
tag=syslog words error | intrinsic embeddings | semantic -t 65 "the service failed to start" | sort by score desc | table score DATA
```

### Chart semantic matches over time

```gravwell
tag=tickets intrinsic embeddings | semantic -t 70 "billing problem" | count | chart count
```
