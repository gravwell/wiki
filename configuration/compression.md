# Compression

Storage is compressed by default, which helps shift some load from storage to CPU.  Gravwell is a highly asynchronous system built for modern hardware in a scale wide paradigm.  Modern systems are typically overprovisioned on CPU with mass storage lagging behind.  Compressing data allows Gravwell can reduce stress on storage links while employing excess CPU cycles to asynchronously compress and decompress data.  The result is that searches can be faster when employing compression on slower mass storage devices.  Compression can be configured independently for each well with different compression settings for hot and cold data.

A notable exception is data that will not compress much (if at all). In this situation, attempting to compress the data burns CPU time with no actual improvement in storage space or speed. Raw network traffic is a good example where encryption and high entropy prevent effective compression.  To disable compression for a well, add the "Disable-Compression=true" directive.

## Compression Settings

Gravwell supports two types of compression: default and transparent compression.  Default compression uses the [snappy](https://en.wikipedia.org/wiki/Snappy_%28compression%29) compression system to perform compression and decompression in userspace.  The default compression system is compatible with all filesystems.  The transparent compression system uses the underlying filesystem to provide transparent block level compression.

Transparent compression allows for offloading compression/decompression work to the host kernel while maintaining an uncompressed page cache.  Transparent compression can allow for very fast and efficient compression/decompression but requires that the underlying filesystem support transparent compression.  Currently the [BTRFS](https://btrfs.wiki.kernel.org/index.php/Main_Page) and [ZFS](https://wiki.archlinux.org/index.php/ZFS) filesystem are supported.

Attention: Transparent compression has important implications for ageout rules involving total storage. Please refer to the [ageout documentation](ageout.md) for more information.

**Disable-Compression**
Default Value: `false`
Example: `Disable-Compression=true`
Compression for the entire well is disabled, both hot and cold storage locations will not use compression

**Disable-Hot-Compression**
Default Value: `false`
Example: `Disable-Hot-Compression=true`
Compression for the hot storage location is disabled.

**Disable-Cold-Compression**
Default Value: `false`
Example: `Disable-Cold-Compression=true`
Compression for the cold storage location is disabled, if no cold storage location is specified the setting has no effect.

**Enable-Transparent-Compression**
Default Value: `false`
Example: `Enable-Transparent-Compression=true`
Gravwell will mark the storage data as compressable and rely on the kernel to perform the compression operations.

**Enable-Hot-Transparent-Compression**
Default Value: `false`
Example: `Enable-Hot-Transparent-Compression=true`
Gravwell will mark the hot storage data as compressable and rely on the kernel to perform the compression operations.

**Enable-Cold-Transparent-Compression**
Default Value: `false`
Example: `Enable-Cold-Transparent-Compression=true`
Gravwell will mark the Cold storage data as compressable and rely on the kernel to perform the compression operations.

Note: If transparent compression is enabled and the underlying filesytem is detected as incompatible with transparent compression, the data will effectivley be uncompressed and Gravwell will send a notification to users.

Warning: If hot and cold storage locations are not compatible with regards to compression, Gravwell must perform additional work to ageout data from hot to cold.  If acceleration is enabled, Gravwell will re-index the data as it performs the ageout.  Incompatible compression settings can incur significant overhead during ageout.  Uncompressed data is compatible with transparently compressed data, but default compression is not compatible with uncompressed or transparently compressed data.  Gravwell will still function perfectly fine with incompatible compression, the indexer will just work much harder during ageout.


## Compression and Replication

The [replication system][replication.md] adheres to the same rules as normal well storage.  Replicated data can be configured to use transparent compression, default compression, or no compression at all.  The same rules for compatibility between hot and cold storage locations in a well and compression also apply to replicated data and replication peers.  If a replication peer has configured an incompatible form of compression indexers will perform significantly more work when restoring after a failure.  For best performance, Gravwell reccomends that hot, cold, and replication stores use the same compression schemes.

Compression for replication storage locations is controlled by the `Disable-Compression` and `Enable-Transparent-Compression` directives.  The snappy compression systme is the default compression scheme.

## Compression Examples

An example storage well with compression disabled for an entire well:

```
[Storage-Well "network"]
	Location=/opt/gravwell/storage/network
	Cold-Location=/mnt/storage/gravwell_cold/network
	Tags=pcap
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Disable-Compression=true
```

An example storage well with compression disabled for the hot storage location and transparent compression enabled for the cold location.  This configuration is considered compatible and will not require additional work during ageout.

```
[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Cold-Location=/mnt/storage/gravwell_cold/syslog
	Tags=syslog
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Disable-Hot-Compression=true
	Enable-Cold-Transparent-Compression=true
```

An example storage well with transparent compression enabled on the hot storage location and default userspace compression on the cold well.  This configuration is considered incompatible and will incur additional overhead during data ageout.

```
[Storage-Well "windows"]
	Location=/opt/gravwell/storage/windows
	Cold-Location=/mnt/storage/gravwell_cold/windows
	Tags=windows
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Enable-Hot-Transparent-Compression=true
	Disable-Cold-Compression=true
```

An example replication configuration that uses transparent compression for the replication storage.

```
[Replication]
	Peer=indexer1
	Peer=indexer2
	Peer=indexer3
	Peer=indexer4
	Storage-Location=/mnt/storage/replication
	Enable-Transparent-Compression=true
```
