# 

![](logo-name.png)

# Gravwell

This site contains documentation for Gravwell, plus other resources such as Changelogs.

If you're just starting out with Gravwell, we recommend reading the [Quickstart](quickstart/quickstart.md) first, then moving on to the [Search pipeline](search/search.md) documentation to learn more.

Gravwell is pleased to announce our free [Community Edition](https://www.gravwell.io/download)!

## Quickstart and Downloads

  * [Quickstart](quickstart/quickstart.md)

  * [Downloads](quickstart/downloads.md)

## Searching with Gravwell

  * [Search overview](search/search.md)

  * [Search Extraction modules](search/extractionmodules.md)

  * [Search Processing modules](search/processingmodules.md)

  * [Search Render modules](search/rendermodules.md)

  * [Alphabetical List of All Pipeline Modules](search/complete-module-list.md)

## System Architecture

  * [Gravwell System Architecture](architecture/architecture.md)

    * [Network Ports Used by Gravwell](configuration/networking.md)


  * [The Resource System](resources/resources.md)

## Ingester Configuration: Getting Data Into Gravwell

  * [Setting Up Ingesters](ingesters/ingesters.md)

    * [File Follower Ingester](ingesters/file_follow.md)

    * [Simple Relay Ingester](ingesters/simple_relay.md)
    
    * [Windows Events Ingester](ingesters/ingesters.md#Windows_Event_Service)

    * [Netflow/IPFIX Ingester](ingesters/ingesters.md#Netflow_Ingester)

    * [Collectd](ingesters/ingesters.md#collectd_Ingester)

  * [Ingester Preprocessors](ingesters/preprocessors/preprocessors.md)

  * [Service Integrations](ingesters/integrations.md)

## Advanced Gravwell Installation and Configuration

  * [Installing and Configuring Gravwell](configuration/configuration.md)

  * [Docker Deployment](configuration/docker.md)

  * [Setting up TLS/HTTPS](configuration/certificates.md)

  * [Building a Gravwell Cluster](distributed/cluster.md)

  * [Distributed Frontends](distributed/frontend.md)

    * [Overwatch](distributed/overwatch.md)


  * [Environment Variables](configuration/environment-variables.md)

  * [Detailed configuration parameters](configuration/parameters.md)

  * [Single Sign-On](configuration/sso.md)

  * [Hardening Gravwell](configuration/hardening.md)

  * [Common Problems & Caveats](configuration/caveats.md)

## Query Acceleration, Auto-Extraction, and Data Management
  
  * [Setting up Auto-extractors](configuration/autoextractors.md)
  
  * [Query Acceleration (indexing and bloom filters)](configuration/accelerators.md)

  * [Data Replication](configuration/replication.md)

  * [Data Ageout](configuration/ageout.md)

  * [Data Compression](configuration/compression.md)

  * [Data Archiving](configuration/archive.md)

## Automation

  * [Scheduled Searches & Scripts](scripting/scheduledsearch.md)

  * [Scripting Overview](scripting/scripting.md)

    * [Automation Script APIs & Examples](scripting/scriptingsearch.md)

## User Interfaces

  * [Gravwell Web GUI](gui/gui.md)

    * [The Search Interface](gui/queries/queries.md)

    * [Labels and Filtering](gui/labels/labels.md)

    * [Kits](kits/kits.md)

  * [Command-Line Client](cli/cli.md)

## API

  * [API](api/api.md)

## Misc

  * [Licensing](license/license.md)

  * [Metrics & Crash Reporting](metrics.md)

  * [Changelogs](changelog/list.md)

  * [Gravwell EULA](eula.md)

  * [Open-source licenses](open_source.md)

Documentation version 2.0
