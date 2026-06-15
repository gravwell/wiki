# kmeans

The `kmeans` module performs [k-means clustering](https://en.wikipedia.org/wiki/K-means_clustering) on numeric enumerated values. It groups entries into `k` clusters based on the similarity of one or more numeric fields, using Euclidean distance as the distance metric. Each entry is assigned a cluster ID, the distance from its cluster centroid, and the number of entries in its assigned cluster.

The module implements Lloyd's algorithm with [k-means++](http://ilpubs.stanford.edu:8090/778/1/2006-13.pdf) initialization for selecting initial centroids.

## Syntax

```
kmeans [options] <keys...>
```

One or more enumerated value names must be provided as keys. These form the dimensions of the feature vector used for clustering. All keys must contain numeric (floating point) values; entries where any key cannot be read as a float are dropped.

## Supported Options

* `-k <n>`: Set the number of clusters. Default is 3.
* `-maxtracked <n>`: Set the maximum number of entries to track. Default is 1000000.
* `-cluster <name>`: Set the name of the enumerated value for the cluster ID assigned to each entry. Defaults to `cluster`.
* `-distance <name>`: Set the name of the enumerated value for the Euclidean distance from the entry to its cluster centroid. Defaults to `distance`.
* `-count <name>`: Set the name of the enumerated value for the number of entries in the associated cluster. Defaults to `count`.
* `-centroids <name>`: Write the calculated centroids to a Gravwell resource as a CSV file. 
* `-r <name>`: Use pre-calculated centroids generated previously with `-centroids`. Enables offline or stable centroid reuse.

## Produced Enumerated Values

The `kmeans` module produces three enumerated values on each entry (names are configurable via flags):

| Enumerated Value | Default Name | Description |
|------------------|--------------|-------------|
| Cluster ID | `cluster` | Integer identifying which cluster the entry belongs to. |
| Distance | `distance` | Float representing the Euclidean distance from the entry to its cluster centroid. |
| Count | `count` | Integer representing the total number of entries in the entry's cluster. |

## Pre-calculating centroids

The `-centroids` flag writes the final calculated centroids to a Gravwell resource in CSV format. This is useful for:

- **Stable clustering**: apply the same centroids consistently across different queries.
- **Offline training**: compute centroids on a large training set, then apply `-r` to new queries for near-instant assignment.

To use pre-calculated centroids, use the `-r` flag with the same resource name created when calculating the centroids.

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
tag=data json x y z | kmeans -k 4 -centroids cluster_model x y z |
```

### Apply pre-computed centroids to new data

Use saved centroids to cluster a new dataset consistently:

```gravwell
tag=data json x y z | kmeans -k 4 -r cluster_model x y z | table x y z cluster distance
```
