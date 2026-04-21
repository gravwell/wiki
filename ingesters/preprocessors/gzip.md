# gzip Preprocessor

The gzip preprocessor can uncompress entries which have been compressed with the GNU 'gzip' algorithm.

The GZIP preprocessor Type is `gzip`.

```{warning}
**Usage Constraints:**

- **Individual Entries Only:** This preprocessor only decompresses single, pre-gzipped entries. It **cannot** be used for bulk decompression.
- **Post-Ingest Processing:** The preprocessor operates **after** data has been successfully received by the ingester. 
- **Transport Agnostic:** It has no impact on the transport protocol used to deliver entries.
- **HTTP Compatibility:** This preprocessor does **not** enable gzipped uploads for HTTP ingesters that do not support `Content-Type: gzip`.
```

## Supported Options

* `Passthrough-Non-Gzip` (boolean, optional): if set to true, the preprocessor will pass through any entries whose contents cannot be uncompressed with gzip. By default, the preprocessor will drop any entries which are not gzip-compressed.

## Common Use Cases

Many cloud data bus providers will package data in a compressed form.  This preprocessor can decompress the data stream in the ingester rather than routing through a cloud lambda function (which can incur additional costs).

## Example: Decompressing compressed entries

Example config:

```
[Preprocessor "gz"]
	Type=gzip
	Passthrough-Non-Gzip=true
```
