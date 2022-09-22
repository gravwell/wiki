## gzip Preprocessor

The gzip preprocessor can uncompress entries which have been compressed with the GNU 'gzip' algorithm.

The GZIP preprocessor Type is `gzip`.

### Supported Options

* `Passthrough-Non-Gzip` (boolean, optional): if set to true, the preprocessor will pass through any entries whose contents cannot be uncompressed with gzip. By default, the preprocessor will drop any entries which are not gzip-compressed.

### Common Use Cases

Many cloud data bus providers will ship entries and/or package in a compressed form.  This preprocessor can decompress the data stream in the ingester rather than routing through a cloud lambda function can incur costs.

### Example: Decompressing compressed entries

Example config:

```
[Preprocessor "gz"]
	Type=gzip
	Passthrough-Non-Gzip=true
```
