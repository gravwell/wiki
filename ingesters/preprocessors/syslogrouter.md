# Syslog Router Preprocessor

The syslog router is designed to dynamically route a single Syslog data stream to multiple Gravwell tags by extracting fields and forming a tag name using those fields.  Common use cases for the syslog router preprocessor is consuming a unified data stream from a 3rd party syslog aggregator; the 3rd party aggregator may not support internal routing to specific IP:Port destinations, so the syslog router preprocessor can extract fields and route based on those fields.

The syslog router preprocessor uses named syslog fields and a template to create a Gravwell tag name.  Tag name templates can contain static values and components of the underlying data.

Templates reference extracted values by name using field definitions similar to bash.  For example, you can reference the syslog `Appname` in the template with `${Appname}`. The templates also support the following special keys:

* `${_SRC_}`, which will be replaced by the SRC field of the current entry.

The syslog router preprocessor Type is `syslogrouter`.

```{note}
The syslog router preprocessor requires properly formed RFC5424 or RFC3164 messages, it will not handle the wildly out of spec "syslog" that many vendors like to claim is compliant.
```

```{warning}
The syslog router preprocessor dynamically creates tags based on the content of data, this means that a data stream could easily exhaust all available tags if you choose a poorly formed template.  Make sure you know what you are doing when using the syslog router preprocessor.
```

## Supported Options

* `Template` (string, required): The `Template` directive is a simplified text template specification for creating a tag structure based on constant values and syslog data.
* `Drop-Misses` (Boolean, optional): By default, if an entry is not a valid RFC5424 or RFC3164 syslog entry, the syslog router preprocessor will not modify the tag and will pass the entry through on the default tag.  The `Drop-Misses` configuration directive causes the preprocessor to entirely drop the entry if it cannot accurately route to an appropriate tag.

### Supported Fields and Routing Rules

The syslog router supports the following Syslog field names for tag routing:

* `Priority`
* `Facility`
* `Severity`
* `Version`
* `Hostname`
* `Appname`
* `ProcID`
* `MsgId`

```{note}
Gravwell tag names may not contain control characters, non-printable characters, or any of the following special characters: `!@#$%^&*()=+<>,.:;\``"'{[}]|`.  If syslog router detects invalid characters in a formulated tag name, they will be replaced with the `_` (underscore) character.
```


```{note}
If a field does not exist (`-` in RFC5424) in a syslog message, the syslog router preprocessor will omit it entirely.
```

## Example: Routing based on Appname

```
[Listener "syslogtcp"]
        Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
        Reader-Type=rfc5424
        Tag-Name=syslog
        Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
        Preprocessor = apprtr

[Preprocessor "apprtr"]
        Type = syslogrouter
        Template=`syslog-${Appname}`
```

| Resulting Tag | Example Syslog Message |
|---------------|------------------------|
| syslog-foo    | <34>1 2003-10-11T22:14:15.003Z worker foo - ID47 - 'su root' |
| syslog-foo-to-the-bar    | <34>1 2003-10-11T22:14:15.003Z worker foo-to-the-bar - ID47 - 'su root' |
| syslog-su    | <34>Oct 11 22:14:15 mymachine su: 'su root' failed for BobFromAccounting on /dev/pts/8 |
| syslog-foo_bar    | <34>Oct 11 22:14:15 mymachine foo!bar: 'su root' failed for BobFromAccounting on /dev/pts/8 |
| syslog-    | <34>1 2003-10-11T22:14:15.003Z worker - - ID47 - 'su root' |

## Example: Routing based on Hostname

```
[Listener "syslogtcp"]
        Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
        Reader-Type=rfc5424
        Tag-Name=syslog
        Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
        Preprocessor = hostrtr

[Preprocessor "hostrtr"]
        Type = syslogrouter
        Template=`syslog-${Hostname}`
```

| Resulting Tag | Example Syslog Message |
|---------------|------------------------|
| syslog-worker    | <34>1 2003-10-11T22:14:15.003Z worker foo - ID47 - 'su root' |
| syslog-    | <34>1 2003-10-11T22:14:15.003Z - foo-to-the-bar - ID47 - 'su root' |
| syslog-192_168_1_1    | <34>Oct 11 22:14:15 192.168.1.1 foobar: 'su root' failed for BobFromAccounting on /dev/pts/8 |


## Example: Routing based on Priority and MsgID and SRC

```
[Listener "syslogtcp"]
        Bind-String="tcp://0.0.0.0:601" #standard RFC5424 reliable syslog
        Reader-Type=rfc5424
        Tag-Name=syslog
        Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
        Preprocessor = complexrouter

[Preprocessor "complexrouter"]
        Type = syslogrouter
        Template=`${Priority}-${MsgID}-${_SRC_}`
```

| Resulting Tag | Example Syslog Message |
|---------------|------------------------|
| 34-ID47-192_168_1_1 | <34>1 2003-10-11T22:14:15.003Z worker foo - ID47 - 'su root' |
| 34-ID47-feed_dead__beef | <34>1 2003-10-11T22:14:15.003Z worker foo - ID47 - 'su root' |

```{note}
Note that the dots and colons in the IPv4 and IPv6 addresses became underscores.  `192.168.1.1` became `192_168_1_1` and `feed:dead::beef` became `feed_dead__beef`.  IPv6 addresses will be rendered using shorthand notation.
```

```{warning}
It is generally recommended that the `${_SRC_}` variable **NOT** be used in the syslog preprocessor due to the generally un-authenticated nature of syslog receivers.  Consider yourself warned.
```
