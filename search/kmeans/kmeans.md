# kmeans

The `kmeans` module performs [k-means clustering](https://en.wikipedia.org/wiki/K-means_clustering) on numeric enumerated values. It groups entries into `k` clusters based on the similarity of one or more numeric fields, using Euclidean distance as the distance metric. Each entry is assigned a cluster ID, the distance from its cluster centroid, and the number of entries in its assigned cluster.

The module implements Lloyd's algorithm with [k-means++](http://ilpubs.stanford.edu:8090/778/1/2006-13.pdf) initialization for selecting initial centroids.

## Syntax

```
kmeans [options] <keys...>
```

One or more enumerated value names must be provided as keys. These form the dimensions of the feature vector used for clustering. All keys must contain numeric values; entries missing a key, or where a key cannot be read as a float, are dropped.

`kmeans` holds every entry until the search finishes, then clusters them and emits them all at once with their timestamps set to the end of the data. If there are fewer entries than clusters (`k`), no entries are emitted.

## Supported Options

* `-k <n>`: Set the number of clusters. Default is 3.
* `-maxtracked <n>`: Set the maximum number of entries to track. The query fails if this is reached. Default is 1000000.
* `-cluster <name>`: Set the name of the enumerated value for the cluster ID assigned to each entry. Defaults to `cluster`.
* `-distance <name>`: Set the name of the enumerated value for the Euclidean distance from the entry to its cluster centroid. Defaults to `distance`.
* `-count <name>`: Set the name of the enumerated value for the number of entries in the associated cluster. Defaults to `count`.
* `-centroids <name>`: Write the calculated centroids to a Gravwell resource as a CSV file, one centroid per row. An existing resource with that name is overwritten.
* `-r <name>`: Start from centroids saved previously with `-centroids` instead of picking initial centroids with k-means++. The resource must have one column per key, and `-k` must equal the number of centroids in it.

## Produced Enumerated Values

The `kmeans` module produces three enumerated values on each entry (names are configurable via flags):

| Enumerated Value | Default Name | Description |
|------------------|--------------|-------------|
| Cluster ID | `cluster` | Integer identifying which cluster the entry belongs to, starting at 0. |
| Distance | `distance` | Float representing the Euclidean distance from the entry to its cluster centroid. |
| Count | `count` | Integer representing the total number of entries in the entry's cluster. |

## Pre-calculating centroids

The `-centroids` flag writes the final calculated centroids to a Gravwell resource in CSV format. A later query can pass that resource to `-r` to use those centroids as its starting point, for example to compute centroids on a large training set and then cluster new data from the same starting point.

With `-r`, `kmeans` still runs the full clustering algorithm, so the centroids can move to fit the new data. Starting from saved centroids makes results more consistent between queries than random k-means++ initialization, but cluster assignments are not fixed to the saved centroids.

## Examples

### Cluster network connections by source and destination port

Extract numeric port values from netflow data and cluster into 5 groups:

```gravwell
tag=netflow netflow Src Dst SrcPort DstPort | kmeans -k 5 SrcPort DstPort | table Src Dst SrcPort DstPort cluster distance count
```

### Sort clusters by distance

Identify outliers by sorting entries by their distance from the cluster centroid:

```gravwell
tag=data json x y | kmeans -k 3 x y | sort by distance desc | table x y cluster distance
```

### Export centroids for later reuse

Compute centroids on a training set and save them to a resource:

```gravwell
tag=data json x y z | kmeans -k 4 -centroids cluster_model x y z | table x y z cluster
```

### Apply pre-computed centroids to new data

Cluster a new dataset starting from the saved centroids:

```gravwell
tag=data json x y z | kmeans -k 4 -r cluster_model x y z | table x y z cluster distance
```
