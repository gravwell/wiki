# Path

The path module extracts and filters parts of a file path: the directory, filename, extension, and volume. It handles Windows paths (including UNC paths) and Linux/Unix paths. A path is treated as Windows if it starts with a drive letter (`C:`) or a UNC prefix (`\\server\share`); anything else is treated as Unix.

## Supported Options

* `-e <arg>`: Extract from an enumerated value instead of the entry data. Entries without that enumerated value are dropped.

## Arguments and syntax

The path module supports four extractions:

| Extraction | Description |
|------------|-------------|
| `base` | The last element of the path. For example, the base of `/opt/gravwell/etc/gravwell.conf` is `gravwell.conf`. |
| `dir` | All but the last element of the path. For example, the dir of `/opt/gravwell/etc/gravwell.conf` is `/opt/gravwell/etc`. Windows paths keep the trailing separator: the dir of `C:\Users\gravwell\foo.txt` is `C:\Users\gravwell\`. |
| `ext` | The extension of the basename, if any. The ext of `foo.txt` is `.txt`. |
| `volume` | The drive letter or UNC volume of a Windows path. For example, the volume of `\\Network\foo\bar` is `\\Network\foo`, and of `C:\Windows` is `C:`. Unix paths have an empty volume. |

Each extraction can be filtered with `==`, `!=`, `~` (contains), or `!~` (does not contain), and renamed with `as`. Entries that fail a filter are dropped. For example:

```gravwell
tag=filenames path base == "foo.txt" as foofiles
```

## Examples

Extract the filename of all ".txt" files:

```gravwell
tag=default path base ext == ".txt" | table
```

Extract directories that are not on the "C:" volume. Unix paths also pass this filter, since their volume is empty:

```gravwell
tag=default path dir volume != "C:" | table
```

