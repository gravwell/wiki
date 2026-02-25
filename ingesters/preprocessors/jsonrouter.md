# JSON Router Preprocessor

The JSON router preprocessor is a flexible tool for routing entries to different tags based on the value of a JSON field within the entries. The configuration specifies a JSON field path, the value of which is then tested against user-defined routing rules.

The JSON Router preprocessor Type is `jsonrouter`.

## Supported Options

* `Route-Key` (string, required): This parameter specifies the JSON field path to extract from the incoming entries. It supports nested paths using dot notation, e.g. `user.role` to access a nested field. For field names containing dots or special characters, wrap the field name in quotes, e.g. `"field.with.dots"` or `user."role.level"`.
* `Route` (string, required): At least one `Route` definition is required. This consists of two strings separated by a colon, e.g. `Route=admin:admintag`. The first string ('admin') is matched against the value extracted from the JSON field, and the second string defines the name of the tag to which matching entries should be routed. If the second string is left blank, entries matching the first string *will be dropped*.
* `Drop-Misses` (boolean, optional): By default, entries which do not contain the specified JSON field or do not match any route will be passed through unmodified. Setting `Drop-Misses` to true will make the ingester drop any entries which do not contain valid JSON or contain a matching field but do not match any key/value pairs in the Route parameter.

```{attention}
Route values with value keys and no tag name will drop data regardless of the Drop-Misses setting.  This allows for specific filtering of data based on JSON field values.
For example, `Route=debug:` will drop all entries where the specified JSON field has the value "debug", while allowing other values to be processed according to the Drop-Misses setting.
```

## Example: Routing to Tag Based on JSON Field Value

To illustrate the use of this preprocessor, consider a situation where many systems are sending JSON-formatted log entries to a Simple Relay ingester. We would like to separate logs based on the `action` field value. Incoming logs look like this:

```
{"action":"login","user":"alice","timestamp":1234567890}
{"action":"logout","user":"bob","timestamp":1234567891}
{"action":"error","message":"Connection failed","timestamp":1234567892}
```

We can apply a JSON router preprocessor to route these entries based on the `action` field:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=actionrouter

[preprocessor "actionrouter"]
        Type = jsonrouter
        Drop-Misses=false
        Route-Key=action
        Route=login:logintag
        Route=logout:logouttag
        Route=error:errortag
```

With the above configuration, logs with `action` set to "login" will be sent to the "logintag" tag, "logout" actions go to "logouttag", and "error" actions go to "errortag". All other logs will go straight to the "jsonlogs" tag.

## Example: Routing Based on Nested JSON Fields

The JSON router preprocessor supports nested field paths using dot notation. Consider logs with a nested structure:

```
{"user":{"role":"admin"},"message":"System configuration changed"}
{"user":{"role":"user"},"message":"Document viewed"}
{"user":{"role":"guest"},"message":"Login page accessed"}
```

You can route based on the nested `role` field:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=rolerouter

[preprocessor "rolerouter"]
        Type = jsonrouter
        Drop-Misses=false
        Route-Key=user.role
        Route=admin:admintag
        Route=user:usertag
```

This configuration routes entries based on the value of `user.role`, sending admin actions to "admintag" and user actions to "usertag", while guest logs will remain on the "jsonlogs" tag.

## Example: Dropping Specific Values and Misses

You can explicitly drop entries with certain values by leaving the tag name blank, and you can drop all entries that don't match any route by setting `Drop-Misses=true`:

```
[Listener "json"]
        Bind-String="0.0.0.0:7777"
        Tag-Name=jsonlogs
        Preprocessor=filterrouter

[preprocessor "filterrouter"]
        Type = jsonrouter
        Drop-Misses=true
        Route-Key=severity
        Route=critical:alerttag
        Route=warning:warntag
        Route=debug:
```

In this configuration:
- Critical severity logs go to "alerttag"
- Warning severity logs go to "warntag"  
- Debug severity logs are explicitly dropped (no tag specified after colon)
- All other logs (info, trace, etc.) are dropped because `Drop-Misses=true`

## Example: Field Names with Special Characters

Example Data:
```
{
    "data": { 
        "event.type": {
            "level":"high"
        }
    },
    "type:name": "info:log"
}
```


If your JSON field names contain dots, spaces, or other special characters, wrap them in quotes in the `Route-Key` parameter:

```
[preprocessor "specialfields"]
        Type = jsonrouter
        Route-Key=`"type:name"`
        Route=`"high:log":hightag`
        Route=`"info:log":` #drop info:log entries entirely
```

For nested paths where one segment contains special characters:

```
[preprocessor "nestedspecial"]
        Type = jsonrouter
        Route-Key=data."event.type".level
        Route=high:hightag
        Route=medium:mediumtag
```

