# Data Replication

Replication is included with all Gravwell Cluster Edition licenses, allowing for fault-tolerant high availability deployments.  The Gravwell replication engine transparently manages data replication across distributed indexers with automatic failover, load balanced data distribution, and compression.  Gravwell also provides fine tuned control over exactly which wells are included in replication and how the data is distributed across peers.  Customers can rapidly deploy a Gravwell cluster with uniform data distribution, or design a replication scheme that can tolerate entire data center failures using region-aware peer selection.  The online failover system also allows continued access to data even when some indexers are offline.

Attention: Gravwell's replication system is designed purely to replicate ingested data. It does not back up user accounts, dashboards, search history, resources, etc., which are stored on the webserver rather than the indexers. To add resiliency to a webserver, consider deploying [a datastore](#!distributed/frontend.md) for your webserver; the datastore will store a redundant live copy of your webserver's data.

The replication system is logically separated into "Clients" and "Peers", with each indexer potentially acting as both a peer and a client.  A client is responsible for reaching out to known replication peers and driving the replication transactions.  When deploying a Gravwell cluster in a replicating mode, it is important that indexers be able to initiate a TCP connection to any peers acting as replication storage nodes.

Replication connections are encrypted by default and require that indexers have functioning X509 certificates.  If the certificates are not signed by a valid certificate authority (CA) then `Insecure-Skip-TLS-Verify=true` must be added to the Replication configuration section.

Replication storage nodes (nodes which receive replicated data) are allotted a specific amount of storage and will not delete data until that storage is exhausted.  If a remote client node deletes data as part of normal ageout, the data shard is marked as deleted and prioritized for deletion when the replication node hits its storage limit.  The replication system prioritizes deleted shards first, cold shards second, and oldest shards last.  All replicated data is compressed; if a cold storage location is provided it is usually recommended that the replication storage location have the same storage capacity as the cold and hot storage combined.

Attention: Entries larger than 999MB will **not** be replicated. They can be ingested and searched as normal, but are omitted from replication.

## Basic Online Deployment

The most basic replication deployment is a uniform distribution where every Indexer can replicate against every other indexer.  A uniform deployment is configured by specifying every other indexer in the Replication Peer fields.

![Basic Replication](replicationOverview.png)

### Example Configuration

Given three indexers (192.168.100.50, 192.168.100.51, 192.168.100.52), the configuration for the indexer at 192.168.100.50 would look like this:

```
[Replication]
	Peer=192.168.100.51
	Peer=192.168.100.52
	Storage-Location=/opt/gravwell/replication_storage
	Insecure-Skip-TLS-Verify=true
	Connect-Wait-Timeout=60
```

Each node specifies the other nodes in its `Peer` fields.

## Region Aware Deployment

The Replication system can be configured to fine tune which peers an indexer is allowed to replicate data to.  By controlling replication peers, it is possible to set up availability regions where an entire region can be taken offline without losing data so long as no subsequent losses occur in the online availability zone.

![Region Aware Replication](RegionAwareReplication.png)

### Example Configuration

An example, an 8 node cluster may be divided into two availability zones (1 and 2).  If availability zone 1 had a subnet of 172.16.2.0/24 and availability zone 2 had a subnet of 172.20.1.0/24.

Nodes in Region 1 are configured to replicate to Region 2:

```
[Replication]
	Peer=172.20.1.101
	Peer=172.20.1.102
	Peer=172.20.1.103
	Peer=172.20.1.104
	Storage-Location=/opt/gravwell/replication_storage
	Connect-Wait-Timeout=60
```

Nodes in Region 2 are configured to replicate to Region 1:

```
[Replication]
	Peer=172.16.2.101
	Peer=172.16.2.102
	Peer=172.16.2.103
	Peer=172.16.2.104
	Storage-Location=/opt/gravwell/replication_storage
	Connect-Wait-Timeout=60
```

## Offline Deployment

Replication is not included with the standard Single Edition Gravwell license.  If your organization does not need a multi-node deployment of Gravwell but would like access to the replication engine for managed backups, contact <sales@gravwell.io> to upgrade a Single Edition license to a replicating Single Edition license.  Single edition replication is entirely offline, meaning that if the indexing goes offline the data cannot be searched until the indexer comes back online and completes recovery.

![Single Edition Offline Replication](SingleOfflineReplication.png)

Cluster Edition Gravwell licenses can choose to implement an offline replication configuration using the offline replicator.  The offline replicator acts exclusively as a replication peer, and does not provide automatic failover or act as an indexer.  Offline replication configurations can be useful in cloud environments where storage systems are already backed by a redundant store and loss is extremely unlikely.  By using an offline replication configuration, data can be replicated to a low cost instance that is attached to very low cost storage pools that would not perform well as an indexer.  In the unlikely event that an indexer is entirely lost, the low cost replication peer can restore the higher cost indexer instance.  Contact <sales@gravwell.io> for access to the offline replication package.

![Offline Replication](OfflineReplicationCluster.png)

## Configuration Options

Replication is controlled by the "Replication" configuration group in the gravwell.conf configuration file.  The Replication configuration group has the following configuration parameters.

| Parameter | Example | Description |
|:----------|:--------|------------:|
| Peer      | Peer=10.0.0.1:9406 | Designates a remote system acting as a replication storage node.  Multiple Peers can be specified. |
| Listen-Address | Listen-Address=10.0.0.101:9406 | Designates the address to which the replication system should bind.  Default is to listen on all addresses on TCP port 9406. |
| Storage-Location | Storage-Location=/mnt/storage/gravwell/replication | Designates the full path to use for replication storage. |
| Max-Replicated-Data-GB | Max-Replicated-Data-GB=4096 | Designates the maximum amount of storage the replication system will consume, in this case 4TB. |
| Replication-Secret-Override | Replication-Secret-Override=replicationsecret | Overrides the authentication token used when establishing connections to replication peers.  By default the "Control-Auth" token from the Global configuration group is used. |
| Disable-TLS | Disable-TLS=true | Disables TLS communication between replication peers. Defaults to false (TLS enabled) |
| Insecure-Skip-TLS-Verify | Insecure-Skip-TLS-Verify=true | Disables verification and validation of TLS public keys.  TLS is still enabled, but the system will accept any public key presented by a peer. |
| Key-File | Key-File=/opt/gravwell/etc/replicationkey.pem | Overrides the X509 private key used for negotiating a replication connection.  By default TLS connections use the Global key file. |
| Certificate-File | Certificate-File=/opt/gravwell/etc/replicationcert.pem | Overrides the X509 public key certificate used for negotiating a replication connection.  By default TLS connections use the Global certificate file. |
| Connect-Wait-Timeout | Connect-Wait-Timeout=30 | Specifies the number of seconds an Indexer should wait when attempting to connect to replication peers during startup. |
| Disable-Server | Disable-Server=true | Disable the indexer replication server, it will only act as a client.  This is important when using offline replication. | 
| Disable-Compression | Disable-Compression=true | Disable compression on the storage for the replicated data |
| Enable-Transparent-Compression | Enable-Transparent-Compression=true | Enable transparent compression on using the host file system for replicated data. |

## Replication Engine Behavior

The replication engine is a best effort asynchronous replication and restoration system designed to minimize impact on ingest and search.  The replication system attempts a best-effort data distribution but focuses on timely assignment and distribution.  This means that shards are assigned in a distributed first-come, first-serve order with some guidance based on previous distribution.  The system does not attempt a perfectly uniform data distribution and replication peers with higher throughput (either bandwidth, storage, or CPU) may take on a greater replication load than peers with less.  When designing a Gravwell cluster topology intended to support data replication, we recommend over-provisioning the replication storage by 10-15% to allow for unexpected bursts or data distribution that is not perfectly uniform.

The replication engine ensures backup of two core pieces of data: tags and the actual entries.  The mapping of printable tag names to storage IDs is maintained independently by each indexer and are critical for effective searching.  Because the tag to name maps are relatively small, every indexer replicates its entire map to every other replication peer.  Data on the other hand is only ever replicated once.

Replication is designed to coordinate with data ageout, migration, and well isolation.  When an indexer ages out data to a cold storage pool or deletes it entirely, the data regions are marked as either cold or deleted on remote storage peers.  The remote storage peers use deletion, cold storage, and shard age when determining which data to keep and/or restore on a node failure.  If data has been marked as deleted by an indexer, the data will not be restored should the indexer fail and recover via replication.  Data that has previously been marked as cold will be put directly back into the cold storage pool during restoration.  Indexers should restore themselves to the exact same state they were in pre-failure when recovering using replication.

### Best Practices

Designing and deploying a high availability Gravwell cluster can be extremely simple as long as a few basic best practices are followed.  The following list calls out some guidelines every Gravwell administrator should strive to follow when deploying and recovering a Gravwell cluster instance.
 
1. `Indexer-UUID` represents an indexer's global identity.  The identity must be maintained for the lifetime of the node and appropriately restored after a failure.  If a failed indexer comes up with a different UUID than it previously used, it is interpreted as a wholly new member in the replication cluster and its previous data will not be restored. We recommend noting down the indexer UUIDs somewhere safe in case an indexer suffers catastrophic failure.
2. Changing well configurations can impact replication states.  Adding additional wells or deleting wells is perfectly acceptable but changing the well configurations *after* a failure but *before* restoration will prevent the replication engine from appropriately restoring data.
3. If an indexer fails, it is critically important that it be allowed to establish connections with replication peers and perform a first-level tag synchronization prior to ingesting new data.  It can be a good idea to set the `Connect-Wait-Timeout` config parameter to zero, ensuring the failed indexer will not start until it has established replication connections and performed a tag restoration.
4. Replication storage locations should be reserved exclusively for a single replication system.  For example, using the same network attached storage location for multiple indexers' `Storage-Location` will cause replication failures and data corruption.
5. Match the compression scheme for replicated and primary data.  If you are using host based transparent compression on the indexers, it is best to mimic that behavior on the replication stores.  If compression schemes match between indexers and replication peers, the restoration process is dramatically faster.

## Troubleshooting

Potential problems and solutions when debugging a replication problem

#### After a failure, an indexer did not restore its data
Ensure that the indexer maintained its original `Indexer-UUID` value when coming back online.  If the UUID changed, put it back to the original value and ensure the indexer has adequate time to restore all data.  Restoration after changing the `Indexer-UUID` may require significantly more time as the replication system merges the two disparate data stores.

#### Data is not showing up in the replication Storage-Location
Ensure that all replication peers have a common `Control-Auth` (or `Replication-Secret-Override`) value.  If peers cannot authenticate with each other they will not exchange data.

Ensure that X509 certificates are signed by a valid Certificate Authority (CA) that is respected by the keystore on the host systems.  If the certificate stores are not valid, either install the public keys into the host machines certificate store, or disable TLS validation via the `Insecure-Skip-TLS-Verify` option.

Attention: Disabling TLS verification via `Insecure-Skip-TLS-Verify` opens up replication to man-in-the-middle attacks.  An attacker could modify data in flight, potentially corrupting logs or hiding activity in replicated data.

Check firewall rules and/or routing ACLs to ensure that indexers are allowed to communicate with one another on the specified port.

#### After a failure an indexer is refusing to start due to a failed tag merge
If an indexer starts ingesting after a failure prior to restoring its tag mapping, it is possible to enter a state where the tag maps on replication nodes cannot be merged.  If you encounter an unmergeable tag error, contact <support@gravwell.io> for assistance in manually restoring the failed node.
 
#### After a failure an indexer did not restore all its data
Replication peers may not have been able to keep up with an indexer due to poor storage performance, poor network performance, or storage failures on the replication node.  Ensure that replication peers have adequate bandwidth and storage capacity to keep up with ingestion.  If a storage node is ingesting at hundreds of megabytes per second, the replication peers must be able to compute, transfer, and store the data at the same rate.

Also ensure that there was adequate storage on replication peers.  If a storage node is configured to keep 10TB of cold data and 1TB of hot data, replication peers should be capable of storing at least 11TB of data.  If a replication node was overloaded or misconfigured it may have been removing old data.

Ensure that the system times on replication nodes and indexers are consistent.  Both systems use the wall-clock time to determine eligibility of data for removal.  If an indexer has an incorrect system time, its data my be prioritized for deletion in the event a replication peer runs out of storage.
