# REST API

This section documents many parts of the REST API provided by the Gravwell webserver. We also host an interactive [Swagger OpenAPI documentation instance](https://api.docs.gravwell.io).


Complete open source clients are available for [Golang](https://github.com/gravwell/gravwell) with [hosted documentation](https://pkg.go.dev/github.com/gravwell/gravwell/v3/client).


The test API located at `/api/test` can be used to verify that the webserver is up and functioning. The test API is unauthenticated and always responds with a StatusOK 200 and an empty body if the webserver is available.

```{toctree}
---
maxdepth: 1
caption: Authenticating with the REST API
---
API Tokens System </tokens/tokens>
```

```{toctree}
---
maxdepth: 1
caption: REST APIs
---
Gravwell Direct Query API </search/directquery/directquery>
```

----

# Scripting API

This section documents the API available in [Script](#search-scripts) automations. The libraries and functions documented below can be used to accomplish complex tasks using Gravwell's scripting engine.

```{toctree}
---
maxdepth: 1
caption: Script Automations
---
Automation Script APIs & Examples </scripting/scriptingsearch>
```

