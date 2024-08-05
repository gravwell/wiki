# Gravwell Glossary

This section defines common terms used in Gravwell, as well as links to documentation where more detailed information can be found.

## General

[API](/api/api)
  Gravwell uses a REST API (Application Programming Interface) which can control essentially every feature in the system. The direct search API is particularly useful in that it lets external systems query data from Gravwell in a single HTTP request. Gravwell also has the capability to interact with the APIs of various other applications within its automation systems.

[Automation](/automation)
Automations are searches, scripts, or flows which are executed at specified times or in response to specific triggers.

[CBAC / Capability Based Access Controls](/cbac/cbac)
   Capability Based Access Controls is a feature which allows configuration of fine-grained access (by user or by group) to individual Gravwell features, functionality, and/or tags.

[Cluster](/distributed/cluster)
   A multi-system deployment of Gravwell. This may range from a single webserver and a single indexer, all the way up to multiple webservers, multiple indexers, a datastore, and a load balancer.

[Community Edition / CE](https://www.gravwell.io/pricing)
    A free, ingest-limited license for personal or commercial use of Gravwell. Intended for small commercial projects, home labs, or experimentation of the Gravwell platform.

[Configuration / .conf file](configuration_overlays)
   All components of Gravwell use a `.conf` file to store and set the variables and options needed to function.  These are generally contained in the `/opt/gravwell/etc` directory and are read and applied when the application is started.

[Datastore](datastore_server)
  The Datastore is a component which synchronizes data across multiple Gravwell webservers.

[Entry](entries)
	An entry is the most basic unit of data in Gravwell. An entry may contain a syslog event, or a Netflow binary record, or a raw packet. Each entry contains four basic built-in fields: DATA, the actual data such as a syslog record; TAG, a tag assigned for organization; SRC, the IP address from which the entry originated; and TIMESTAMP, the time associated with the entry.

[Flow](/flows/flows)
Flows are a no-code system for developing advanced automations which are run by Gravwell on a schedule or in response to an alert trigger.

[Indexer](configuration_indexer)
  The indexer is the data storage component in a Gravwell deployment. It is responsible for the storage, retrieval, and processing of entries.

[Ingester](/ingesters/ingesters)
  The ingesters are the primary method of getting entries into the Gravwell indexers. They can reside on the same system as the indexer or webserver, or they can be installed on separate machines, depending on your needs.  Ingesters are generally designed to handle a specific data source or data type. Multiple ingesters can sit on a single system. 

[Kits](/kits/kits)
   Kits are Gravwell's way of bundling up a lot of related items (dashboards, queries, scheduled searches, autoextractors, scripts, etc.) for easy installation on other systems.  Gravwell Inc. provides pre-built kits for common use cases, but users can also build their own kits from within the Gravwell UI.

[License](/license/license)
   The license allows the application to run. Once uploaded to the webserver at install time, it will be automatically distributed to the indexers. License updates are generally automatic.
   
[Load Balancer](/distributed/loadbalancer)
   A load balancer is a system which will split an incoming web workload across multiple systems. Gravwell provides a custom load balancer designed for use with Gravwell webservers to simplify a clustered deployment.

[Overwatch](/distributed/overwatch)
   Overwatch is an advanced deployment where you can have specialized webservers that can query multiple Gravwell clusters.

[Scheduled Script](/scripting/scheduledsearch)
A user-authored script which is run automatically on a given schedule. The schedule is defined using cron format, meaning scripts may run as frequently as every minute.

[Scheduled Search](/scripting/scheduledsearch)
A Gravwell query which will be run automatically on a given schedule. The schedule is defined using cron format, meaning searches may run as frequently as every minute.

[Search Agent](/scripting/searchagent)
  The search agent is the component which runs automations.

[Self-Hosted / On-Prem](system_requirements)
     In the context of Gravwell, self-hosted systems are considered to be any Gravwell deployment where the management of the hardware & operating system is under the control of the customer, rather than Gravwell Inc. This can include owned/leased systems in a customer datacenter, or customer-managed VMs in a public cloud environment (AWS / Azure / GCP / Oracle / etc).  The application can be installed and managed via a standard Linux package manager (Redhat RPM or Debian .deb formats),  Docker containers,  or via a distribution-agnostic self-extracting installer.

[SSO / Single Sign-On](/configuration/sso)
   Gravwell can be configured to use a SAML-Compliant Identity Provider for user authentication.

[Structure-on-Read](/search/search)
  Gravwell is a "structure-on-read" data lake,  where data is ingested directly into the system with no pre-processing required; structure is then applied to that data at *search time*. This allows a lot of flexibility as you do not need to know how the data looks or "normalize" it before sending it into Gravwell, and if that structure changes for any reason (e.g.  errors or unexpected behavior in the source device, vendor format changes on an upgrade, enabling additional options on the source device to receive more detailed information) it will not prevent the data from being ingested or cause a data loss scenario.  By retaining all the data in its native format, it also makes it easier to ask questions retroactively on historical data knowing you have not lost any visibility since it was generated.

[Tags](tag_specification)
   Tags are used to logically organize entries. A tag is assigned to each entry at ingest time, thus incoming syslog data may be tagged "syslog", while webserver logs might be tagged "apache", Netflow v5 records are tagged "netflow", and Windows event logs get tagged "winlog". When running a search, the first part of the query string specifies which tags should be accessed, e.g. `tag=syslog grep Failed`.

[Webserver](configuration_webserver)
  The webserver is the component which coordinates searches across the indexers and serves Gravwell's web interface.

## GUI / Interface

[Actionables](/gui/actionables/actionables)
   Actionables provide a way to create custom menus that key on any text rendered in a query, enabling users to take different actions on that text via the menu. For instance, an actionable could be defined to provide a menu of options whenever an IP address is clicked. Actionables can be used to open external URLS, submit new Gravwell queries, launch dashboards, and execute templates.

[Dashboards](/gui/dashboards/dashboards)
   Dashboards are Gravwell's way of showing the results from multiple searches at the same time. A Dashboard contains many *tiles*, each associated with a Gravwell query.

[Email](/configuration/email)
   Gravwell can be configured to send emails via automated scripts and flows.  The outgoing SMTP server information can be specified by individual users, or the Gravwell administrator can define a system-wide server for all users.

<a href="/gui/dashboards/dashboards.html#investigative-dashboards">Investigative Dashboard</a>
   A special type of dashboard that leverages templates instead of standard queries. When an investigative dashboard is launched, it sets the variables for its templates, either by prompting the user directly or automatically if the user has accessed the dashboard through an actionable. Investigative dashboards get their name from the common use case: a dashboard which takes a specific IP/username/hostname/etc and runs multiple searches to investigate that variable.

[Labels](/gui/labels/labels)
   Gravwell allows the assignment of labels to various items for organizational purposes.  Labels can be used to help group or filter items within the GUI.

[Persistent Searches](/gui/persistent/persistent)
	A page showing a list of searches on the system. Searches will appear in Persistent Searches when they are actively running, were set into a background state, or have had their results saved.

[Playbooks](/gui/playbooks/playbooks)
    Playbooks are hypertext documents which help guide users through common tasks, describe functionality, and record information about data in the system. Playbooks can include a mix of text, images, links, resources, and executable queries. 

[Query Library](/gui/querylibrary/querylibrary)
    A collection of query strings saved by the user for later re-use (or distributed in a kit). A query library item can be used in a dashboard tile or referenced in a scheduled search or flow.

[Search History](search_history)
  A page showing a list of past searches run on the system.
  
[Search Query / Query Studio](/gui/queries/queries)
   A search query is the most basic function within Gravwell, where you can ask questions of the data within the system. A query is a string which specifies what data should be accessed, how it should be filtered and processed, and how it should be displayed to the user. Query Studio is the primary interface for building and submitting queries.

[Secrets](/gui/secrets/secrets)
  Gravwell can store secret strings (e.g. API tokens) for use in flows and scripts; this feature is intended to provide a basic level of defense against leaking secrets accidentally. Once entered, a secret's contents cannot be easily extracted/read by a user; however, they may change the value later.
  
[Template](/gui/templates/templates)
  Templates are essentially stored queries which contain a variable. Templates may be used in actionables or dashboards to provide basic investigative functionality, e.g. clicking an IP address and searching for all traffic it originated.

[User Files](/gui/files/files)
   Gravwell users can upload small files for use in playbooks, as cover images for custom kits, etc.
   
## Search

[Auto-Extractors](/configuration/autoextractors)
   Auto-extractors provide search-time information on the structure of data in a given tag. Because Gravwell is a structure-on-read system, the user generally needs to know what sort of structure to apply to any given entry. Once an appropriate structure has been determined for a given tag, it can be stored in an *auto-extractor* definition for later use.


[Compound Queries](compound_queries)
Multiple queries can be combined into a single query string to be run sequentially, with the results of earlier queries available to later queries. This can help fuse data from multiple sources or just to simplify complex queries.

[Enumerated Value / EV](enumerated_values)
   Enumerated Values are special data elements which can be attached to entries during a search.  All entries contain the DATA, TIMESTAMP, SRC, and TAG "special" EVs by default.  As the data moves through the pipeline, modules may attach additional enumerated values to the entry in the pipeline; the `netflow` module might be used to extract the `src` IP address from each record, then later the `stats` module can count how many times each different source IP was seen in the data.

[Extraction Modules](/search/extractionmodules)
   Extraction modules are used to pull specific information out of the raw underlying data. For example, the `netflow` module parses entries containing Netflow records; one might use the `netflow` module to extract the src and dst fields from a data set.

[Inline Filtering](/search/filtering)
	Filters can be applied directly in the extraction module invocation; this is called *inline filtering* and is typically the most efficient way to filter down the results in the pipeline.
	
[Macros](/search/macros)
   Macros are essentially string replacement rules for use in queries.  They allow the creation of pre-defined variables containing strings that can then be called within the search pipeline. For instance, a macro may be used to store a long and complex regular expression.

[Processing Modules](/search/processingmodules)
	Processing modules perform various actions on data, rather than extracting values as Extraction modules do. Processing modules can sort entries, filter based on extracted values, perform statistical operations on the data, or (in the case of the `eval` module) do arbitrary programmatic operations on entries.

[Renderers](/search/rendermodules)
	Renderers are the last modules invoked in a search query. They receive the outputs of the search modules and prepare it for display to the user.
	
[Resources](/resources/resources)
   Resources allow users to store persistent data for use in searches. They can be manually uploaded by users or automatically created by some search modules.  Common uses include storing lookup tables for data enrichment, whitelists/blacklists, and GeoIP databases.

[Search Pipeline](/search/search)
   The search pipeline is the core of Gravwell's functionality. Entries are pulled from the indexer disks, fields are extracted, filters are applied, statistics may be computed, and finally the results are presented to the user as a table, graph, or map.

[Search Timeframe](timeframe_selector)
   The period of time over which the search will run.
   
  
## Indexers and Storage

[Cloud Archive](/configuration/archive)
  The final tier of storage is called "cloud archive", where indexers can upload shards to a dedicated archive server before deleting them from the indexers. This is ideal for data that must be retained long-term but no longer needs to be actively searched.

[Cold Storage](/configuration/ageout)
	Wells may optionally have a "cold" storage location. Data may be sent from hot storage to cold storage based on a variety of ageout constraints. Cold storage frequently uses slower storage technology, as "cold" data is assumed to be accessed less frequently.

[Data Ageout](/configuration/ageout)
  Gravwell supports an ageout system whereby data management policies can be applied to individual wells. Entries can be evicted from a given well based on their age, the total storage consumed by the well, or the amount of free space remaining on the disk.

[Data Compression](/configuration/compression)
   Storage is compressed by default, which allows shifting some load from storage to CPU. For many log types, such as syslog data, it can also improve query speed as less data must be read from physical disks.

[Data Shard](configuration_tags_and_wells)
	Each individual well further segments its data into *shards*. Each shard contains approximately 1.5 days worth of entries.

[Hot Storage](/configuration/ageout)
   Every well has a "hot" storage location which holds the most recent data. This should be the fastest storage available, as it will generally be the most frequently accessed.

[Query Acceleration](/configuration/accelerators) 
   Gravwell supports the notion of "accelerators" for individual wells and/or tags. Accelerators can parse data at ingest time and build indexes based on the contents of the entries; at search time, these indexes can be used to narrow down which entries are actually processed, speeding up queries. The most basic accelerator, "fulltext", simply builds a full-text index of the entries; this allows very rapid text-based searching. Other accelerators can parse & index CSV, JSON, Netflow, packets, and other formats.

[Replication](/configuration/replication)
   Gravwell supports the ability to transparently replicate data across distributed indexers. If an indexer goes offline, the other indexers can serve their replicated copies of its data in queries until the indexer is repaired. If an indexer loses data due to e.g. a bad disk, it can restore from the replicated copies.

[Storage](indexer_storage_topology)
	"Storage" is the general term used for entries residing on disk. Gravwell organizes entries in several ways: they are logically grouped into *wells*, and then temporally grouped into *hot* and *cold* storage tiers.

[Storage Wells](configuration_tags_and_wells)
	Wells are the organizational structure used to group tagged entries on the disk. The user running Gravwell queries will be unaware of which wells his or her entries are coming from, but by thoughtful well configuration it is possible to achieve optimal query performance. A well configuration specifies where on the disk data should be stored, which tags should go into the well, and what data ageout strategy (if any) should be used.

   
## Ingesters

[Federator](/ingesters/federators/federator)
  A federator is an entry relay. It sits between ingesters and indexers and may be used for several reasons. If a deployment has many indexers, a federator can be deployed so ingesters only need to be pointed at the single federator rather than all the indexers. It can also be used to provide a restricted ingest point for less-trusted data sources, e.g. only allowing ingest to a subset of tags. It is also a useful way to allow incoming entries from the Internet without directly exposing the indexers.
  
[Preprocessor](/preprocessors/preprocessors)
Ingesters can be configured to apply additional preprocessing to entries before sending them to the indexer. For instance, the entry's tag may be modified based on the contents of the entry.

[Timegrinder](/ingesters/customtime/customtime)
Ingesters attempt to extract timestamps from entries using a library called timegrinder. The behavior of timegrinder can be modified in the ingester configuration.
