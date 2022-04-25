## Path

The path module extracts and filters path names.  It can process individual portions of a path, including the directory, filename, extension, and volume. It works with both Windows (including UNC paths) and Linux/Unix style paths. 

### Supported Options

* `-e`: Extract from an enumerated value instead of the DATA portion of the entry.

### Arguments and syntax

The path module supports the following four extractions:

| Extraction | Description |
|------------|-------------|
| `base` | The last element of the path. For example, the base of `/opt/gravwell/etc/gravwell.conf` is `gravwell.conf`. |
| `dir` | All but the last element of the path. For example, the dir of `C:\Users\gravwell\foo.txt` is `C:\Users\gravwell`. |
| `ext` | The extention of the basename, if any. The ext of `foo.txt` is `.txt`. |
| `volume` | (Windows only): The volume drive letter or UNC volume of a Windows path. For example, the volume of `\\Network\foo\bar` is `\\Network\foo`, and `C:\Windows` is `C:`. |

Each extraction can be filtered and renamed using the `as` keyword. For example:

```
tag=filenames path base == "foo.txt" as foofiles
```

### Examples

Extract the filename of all ".txt" files:

```
tag=default path base ext == ".txt" | table
```

Extract any Windows directories that are not on the "C:" volume:

```
tag=default path dir volume != "C:" | table
```

