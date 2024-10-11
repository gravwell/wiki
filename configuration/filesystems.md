# Gravwell Indexer Supported Filesystems

Gravwell Indexers require robust, seekable, and POSIX complaint filesystems in order to function properly.  The Gravwell system makes extensive use of memory mapping, madvise calls, and filesystem specific optimizations to maximize compression ratios and query throughput.  Picking a good filesystem for your deployment is critical to ensuring a manageable and fast Gravwell system.

Gravwell officially supports the following Linux filesystems.

| Filesystem | Minimum Kernel Version | Supports Transparent Compression |
|:-----------|:-----------------------|:--------------------------------:|
| EXT4       | 3.2                    |                                  |
| XFS        | 3.2                    |                                  |
| BTRFS      | 5.0                    | ✅                               |
| ZFS        | N/A                    | ✅                               |
| NFSv4      | N/A                    |                                  |




### Ext4

The Ext4 filesystem is well supported and an excellent default choice as a backing filesystem.  Ext4 supports volume sizes up to 1EiB and up to 4 Billion files, Gravwell extensively tests on Ext4.

### XFS

The XFS filesystem is extremely fast, well tested, and praised by kernel developers.  XFS supports a wide array of configuration options to optimize the filesystem for specific storage device topology.

### BTRFS

The BTRFS filesystem has been a core part of the Linux kernel for over a decade, but due to some rocky starts and conservative warnings about stability early on in its life cycle it gets a bad rap.  Gravwell extensively tests the BTRFS filesystem in transparent compression topology and has found it to be exceedingly fast, memory efficient, and well supported.  While BTRFS is supported all the way back to Linux Kernel 3.2, 5.X and newer kernels contain a highly optimized and stable code base.  Gravwell recommends BTRFS with ZSTD compression for a hot store when transparent compression is enabled and users want the best performance.

### ZFS

The ZFS filesystem has long been praised as **THE** next generation filesystem, it has a stable well maintained code base with robust documentation.  However, ZFS is in a bit of a strange situation in the Linux kernel in that many distributions do not natively support it and the Kernel maintainers believe it has an incompatible license.  ZFS also employs its own caching strategy that is not well blended with the Linux page cache; this means you need to manually tune the ZFS ARC cache and be aware that the ARC cache competes for memory with the Gravwell processes.  When memory gets tight, ZFS will not free memory in the same way that BTRFS may.  That being said, the additional configuration options available in ZFS make it a good choice for cold storage when compression ratios are of the utmost importance.

Gravwell recommends ZFS when transparent compression is desired for a cold storage tier and compression efficiency is more important than raw speed.  Setting the block size to 1MB and the compression system to zstd-10 can yield impression compression ratios that still perform well.  ZFS however is significantly slower than BTRFS when using transparent compression and a fast storage device.  ZFS also does not support the ability to disable COW and compression on a per file basis, so ZFS will attempt to compress and fragment highly orthogonal data structures like well indexes.

### NFSv4

Some customers desire storage arrays to be fully remote with dedicated storage appliances doing the dirty work of data management.  Gravwell tentatively supports NFSv4 with a few caveats.  The filesystem must be configured with all supporting daemons and mount options such that file permissions can be properly mapped to the NFS volume.  While it is possible to disable user/group management on NFS entirely, this is not recommended.

Gravwell Indexers also maintain long lived file handles with very high I/O requirements. NFS, being a network filesystem, suffers from network interruptions, which can cause process hangs, unexpected performance drops, and increased complexity of management.  Gravwell only tests on NFSv4 and generally does not recommend it.


## Unsupported Filesystems

Gravwell requires full robust POSIX compatibility. The following filesystems are not supported at all.  Gravwell may still function, but we make no guarantees about performance, reliability, or correctness.

* FAT32
* VFAT
* NTFS
* SMB/CIFS
* FUSE mounts

Other POSIX compliant filesystems like EXT2, EXT3, and ReiserFS are not tested.  Cluster filesystems such as GlusterFS, LusterFS, and CephFS are fully POSIX compliant and customers have reported good results, however Gravwell has not done extensive extensive testing and does not officially support them.
