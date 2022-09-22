## CSV Router Preprocessor

The CSV router preprocessor can route entries to different tags based on the contents of a specified CSV column in the entry. The configuration specifies a column index (zero indexed), the contents of which are then tested against user-defined routing rules.

The CSV Router preprocessor Type is `csvrouter`.

### Supported Options

* `Route-Extraction` (int, required): This parameter specifies the column index (beginning with 0), of the CSV which will contain the string used to compare against routes.
* `Route` (string, required): At least one `Route` definition is required. This consists of two strings separated by a colon, e.g. `Route=sshd:sshlogtag`. The first string ('sshd') is matched against the value extracted via `Route-Extraction`, and the second string defines the name of the tag to which matching entries should be routed. If the second string is left blank, entries matching the first string *will be dropped*.
* `Drop-Misses` (boolean, optional): By default, entries which are not valid CSV do not contain enough columns to extract the `Route-Extraction` column will be passed through unmodified. Setting `Drop-Misses` to true will make the ingester drop any entries which 1) are not valid CSV, or 2) do not contain enough columns to extract the `Route-Extraction` column.

### Example: Routing based on the Third Column

Given the input:

```
0,fritz,employee
1,kris,boss
```

We can use the CSV router to route entries based on the user role, located in the 3rd column of the CSV.

Below is a partial Simple Relay configuration that will route the above CSV to two different tags, based on the contents of the 3rd column.

```
[Listener "csv"]
        Bind-String="0.0.0.0:2601" #we are binding to all interfaces, with TCP implied
        Tag-Name=default
        Preprocessor=role

[preprocessor "role"]
        Type = csvrouter
        Drop-Misses=false
        Route-Extraction=2 # csv is zero-indexed, so 2 is the 3rd column
        Route=employee:csvemployees
	Route=boss:csvbosses
```
