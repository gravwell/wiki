# Gravwell Glossary

This section will include basic descriptions for various terms and items used within Gravwell,   as well as links to the location in the documentation where more detailed information can be found.

### Primary Components and General terms

[Indexer](/configuration/configuration.html#indexer-configuration)
  The Indexer is the Primary component in a Gravwell Deployment. It is responsible for the Storage, retrieval, and processing of data.

[Webserver](/configuration/configuration.html#webserver-configuration)
  The Webserver acts as the focusing point for all searches, and provides an interactive interface into Gravwell. (The GUI)

[Cluster](/distributed/cluster.html)
   A Multi-System Deployment of Gravwell

[Load Balancer](/distributed/loadbalancer.html)
   A Load Balancer is a system which will split an incoming workload across multiple systems.   Gravwell provides a custom load balancer designed for use with Gravwell webservers to simplify your clustered deployment

[Ingester](/ingesters/ingesters.html)
  The Ingester is the primary method of getting data into the Gravwell Indexers. They can be homed on the same system as the Indexer, or remote systems depending on your needs.  Ingesters are generally designed to handle a specific data source or data type.  Multiple Ingesters can sit on a single system. 

[Federator](/ingesters/federators/federator.html)
  A Federator is an entry Relay.  It can optionally be deployed between Ingesters and the Indexer data flow.  It's most common use is as a trust boundary.

[API](/api/api.html)
  An API is an Application Programing Interface.  Gravwell uses a REST API which can be leveraged for direct queries and some scripting.  Gravwell also has the capability to interact with the API's of various other applications within it's automation systems.

[License](/license/license.html)
   In Gravwell the License is used to enable the application and enable the features based off the License tier which you have subscribed.
   
[Search Agent](/scripting/searchagent.html)
  The Gravwell Search Agent is the component which runs Automations

[Datastore](/distributed/frontend.html#the-datastore-server)
  The Datastore is the server/process which allows for the syncing of data across multiple Gravwell Webservers.  

[Overwatch](/distributed/overwatch.html)
   Gravwell Overwatch is a Advanced Deployment configuration where you can have specialized webservers that can query multiple self contained Gravwell Clusters.   

[CBAC / Capability Based Access Controls](/cbac/cbac.html)
   Capability Based Access Controls allows configuring fine grained access by user or group to various Gravwell features, functionality, and/or tags

[SSO / Single Sign-On](/configuration/sso.html)
   The ability to leverage any SAML-Compliant Identity Provider to log into Gravwell

[Docker](/configuration/docker.html)
   A Containerized Environment that allows the installation, management, and configuration of an application abstracted outside of the host operating system(OS).  Gravwell publishes containers to Dockerhub for a number of components.

[Community Edition / CE](https://www.gravwell.io/pricing)
    The Free License for Personal or Commercial use of Gravwell perfect for small commercial projects,  Community home labs,  or experimentation of the Gravwell Platform.

[Self-Hosted / On-Prem](/quickstart/quickstart.html#system-requirements)
    Application run on your own servers.  In the context of Gravwell these systems can be owned/leased systems in a location you control,  or a system you manage in a public cloud environment (AWS / Azure / GCP / Oracle / etc).  The application can be installed and managed via a standard linux package manager (Redhat RPM or Debian .deb formats),  Docker containers,  or via a distribution-agnostic self-extracting installer.

[Cloud / Managed Hosting](https://www.gravwell.io/pricing/enterprise-cloud-edition)
    The Application run on servers Managed by Gravwell.   In the Context of Gravwell,  We will handle the hardware, Operating System, and software management,  allowing you to focus on your data and application.

[Configuration / .Conf file](/configuration/configuration.html#configuration-files-overlay-directories)
   All components of Gravwell will use a `.conf` file to store and set the variables and options needed to function.  These are generally contained in the `/opt/gravwell/etc` directory, and are read and applied when the application is started.

[TLS / SSL](/configuration/certificates.html)
   Encryption protocols that allow you to securely authenticate and transport data across a network.  SSL (Secure Socket Layer) is the older standard which is being depreciated in favor of the more modern TLS (Transport Layer Security) standards. In addition to enabling HTTPS access to the Webserver GUI,  Gravwell supports TLS listeners on several Ingesters,  as well as TLS Communication between the Ingesters, Federators, and Indexers.

[Structure on Read](/search/search.html)
  Gravwell is a Structure on Read data lake,  where data is ingested in it's raw form directly into the system and any structure is applied to that data when its read. This allows for a lot of flexibility as you do not need to know how the data looks or "normalize" it before you send it into Gravwell,  and if that structure changes for any reason (Ex.  Errors or unexpected behavior in the source device,   Vendor format changes on an upgrade,  enabling additional options on the source device to receive more detailed information) it will not prevent the data from being ingested or cause a data loss scenario.   By retaining all the data in it's native format,  it also makes it easier to ask questions retroactively on historical data knowing you have not lost any visibility since it was generated.
  

### GUI / Interface

[Search Query / Query Studio](/gui/queries/queries.html)
   A Search Query is the most basic function within Gravwell, where you can ask questions of the data within the system.   Query Studio is the primary interface for building and submitting the queries you may have.

[Query Syntax / grav.y](/search/spec.html)
   The basic syntax/grammer/rules upon which queries are written and inputted into the system.  We designed the language to be as easy and flexible as possible using some basic frameworks and concepts well established within the technical community.  From a high level,  if you are familiar with Linux CLI concepts such as the use of Pipes `|` to build a command pipeline, quoting `"` and escaping special characters,   or C style opperators `==` ``!=` `/*  */`,  you should be have no issues  understanding how to build a search query.

[Search History](/gui/queries/queries.html#accessing-search-history)
  A list of past searches run on the system.

[Query Library / Saved Search](/gui/queries/queries.html#accessing-the-query-library)
    A collection of previously saved searches

[Search Timeframe](/gui/queries/queries.html#selecting-a-timeframe)
   The Period of time you wish your search to cover.

[Search Results](/gui/queries/queries.html#search-results-page)
   The results of the search that was run

[Scheduled Search](/scripting/scheduledsearch.html#managing-scheduled-searches)   
     An automated Search scheduled using standard CRON notation.  Scheduled searches can either be a manually inputted search, or utilize a saved search from the Query Library.

[Template](/gui/templates/templates.html)
  Templates are essentially stored queries which contain a Variable.

[Macros](/search/macros.html)
   Macros are essentially string replacement rules.  They allow the creation of pre-defined variables containing strings that can then be called within the search pipeline.  Common uses might be a long regex string that you'd like to easily and neatly call within a search,  or a commonly used set of processing you do on some data before asking your question,  or even a list of tags/ips/users/etc which you may commonly need to insert into your search string.

[Dashboards](/gui/dashboards/dashboards.html)
   Dashboards are Gravwell's way of showing the results from multiple searches at the same time. A Dashboard contains many *tiles*, each associated with a Gravwell query.

[Investigative Dashboard](/gui/dashboards/dashboards.html#investigative-dashboards)
   A Special type of Dashboard that leverages Templates instead of standard queries.  Unlike traditional Dashboards, the Template's variables will give the different tiles a more dynamic quality.  They get their name from the common use case of people able to plug in a specific IP/Username/Hostname/etc and run multiple searches against that data in a single pane to investigate what you can find around the specific pivot point.

[Persistent Search](/gui/persistent/persistent.html)
    A Persistent Search allows a User to send a long-running search to the background,  or save search results for later analysis or sharing.

[Labels](/gui/labels/labels.html)
   Gravwell allows the assignment of "Labels" to various items for organizational purposes.  Labels can be used to help group or filter items within the GUI.

[Resources](/resources/resources.html)
   Resources allow users to store persistent data for use in searches. They can be manually uploaded by users, or automatically created by some search modules.  Common uses could include lookup tables data enhancement,  whitelist/blacklists,  regex strings, and/or anko scripts.

[Auto-Extractors](/configuration/autoextractors.html)
   Auto Extractors enable configuration of pre-defined extraction and definition to the unstructured data and data formats that are not self-describing.  They are a huge Quality of Life and daily usage improvement that will greatly simplify making use of the non-normalized data you work with on a daily basis.

[Kits](/kits/kits.htm)
   Kits are Gravwell's way of bundling up a lot of related items (Dashboards, Queries, Scheduled Searches, Autoextractors, scripts, etc) for easy installation to other systems.  Gravwell Inc provides pre-built kits for common use cases,  but users can also built their own kits from within the Gravwell UI.

[User Files](/gui/files/files.html)
   Gravwell Users can upload small files for use in Playbooks,  as cover images for custom kits,  etc.

[Playbooks](/gui/playbooks/playbooks.html)
    Playbooks are hypertext documents within Gravwell  which help guide users through common tasks,  describe functionality, and record information about data in the system.  Playbooks can include a mix of text, images, links, resources,  and executable queries. 

[Actionables](/gui/actionables/actionables.html)
   Actionals provide a way to create custom menus that key on any text rendered in a query, enabling users to take different actions on that text via the menu. Actionables can be used to open external URLS that key on data,  submit new gravwell queries,  launch dashboards, and execute templates.

[Secrets](/gui/secrets/secrets.html)
  Gravwell can store secret strings for use in flows.  This can be used to securely store items like API tokens within Gravwell which can then be used in a variety of Flows and Alerts without exposing the underlying secret data to users.  Once entered a secret's contents cannot be extracted/read by a user,   however they may change the value later.

[Email](/configuration/email.html)
   Gravwell can be configured to send emails via automated scripts and flows.   The outgoing smtp server information must be configured by the user or the system admin to allow the functionality.

[Topology](/gui/systems-health.html#topology)
   The Topology is the layout of the system's deployment.  Gravwell's Systems and Health options will allow you to view the basic topology layout of the system so you can see how the different physical and logical components are layed out within your deployment.

[CLI / Command Line Interface](/cli/cli.html)
    Gravwell has a command line client that can be used to remotely manage gravwell or perform basic searches.



    
### Indexer Terms

[Storage](/architecture/architecture.html#indexer-storage-topology)
   Storage, in it's most basic form, is the physical location where your data will be placed.  ie.  Your SSD and Disk based drives.  Due to the nature of ingesting large amounts of data and also searching through that data, Disc performance will be a primary factor in the system's overall performance.

[Tags](/search/search.html#tag-specification)
   A Tag is used as a method to logically seperate data of different types. A Tag is assigned to an entry at Ingest time. It can also be used along with CBAC to limit access in the system to data.  A Tag is also the primary way of informing and filtering the search pipeline to the specific data you wish to query,  so it can be useful to apply unique tags to different types of data.
   
[Storage Wells](/configuration/configuration.html#tags-and-wells)
  A Storage Well is where you can configure a variety of parameters informing the system how to store and manage the raw data being ingested into the system.  Keeping with the "Data Lake" analogy where every piece of data brought in is a drop of water,   a Well would be where you would tell the system which Storage pool and where its located, to  place that drop of data.

[Data Shard](/configuration/configuration.html#tags-and-wells)
  Wells store the data on-disk in Shards,  with each shard generally containing 1.5 days worth of data. 
  
[Data Ageout](/configuration.html#data-ageout)
  Gravwell supports an ageout system whereby data management policies can be applied to individual wells. This allows for the automated management of data at the shard level, for example allowing for a hands off ability to optomize physical storage and well usage or compliance with data retention policies.    Despite the name,  Ageout can be configured around Total Storage,  Storage Availability, or Time based criteria.

[Hot Storage](/configuration/ageout.html#basic-configuration)
   Every Well has a Hot storage location that would be defined by the `Location` directive.  This is the physical location (based off the unix path) where data is stored when it is first ingested. 

[Cold Storage](/configuration/ageout.html#basic-configuration)
    Cold storage is an optional location that can be defined within a Well Configuration, and is used in conjunction with ageout rules to move data from the Hot to Cold location.   A common use case with Cold storage would be regularly accessed data would initially go into a Hot Storage location housed on high speed NVME type storage,   but after its no longer needed for regular searches could be moved to less expensive and slower Spinner type hard disc drives.

[Cloud Archive](/configuration/archive.html)
  Gravwell supports an ageout mechanism called Cloud Archive,  where indexers can upload shards to a cloud archive server before deletion from the indexers.   This is ideal for data that must be retained long term but no longer needs to be actively searched.

[Replication](/configuration/replication.html)
   Gravwell supports the ability to transparently replicate data across distributed indexers allowing for a fault-tolerant high availability system that protects access to your data.

[Query Acceleration](/configuration.html#query-acceleration) 
   Gravwell supports the notion of "accelerators" for individual wells, which allow you to apply parsers to data at ingest to generate optimization blocks. Accelerators are extremely useful for needle-in-haystack style queries, where you need to zero in on data that has specific field values very quickly.

[UUID](/configuration/parameters.html#indexer-uuid)
  A Unique Identifier for system within the Gravwell cluster.   No two devices should have the same UUID or you may see unexpected behaviour

[Pipe Ingest](/configuration/parameters.html#pipe-ingest-path)
  While the "normal" method of transporting data from an ingester to the indexer involves transporting over a network connection, The system also supports the use of a unix named pipe that ingesters located on the same system can utilize for a high speed and low latency transport.   Names pipes can also be utilized to facilitate transport over unusual network transports or non-IP based interconnects.

[Ingest Secret](/configuration/parameters.html#ingest-auth)
  The Ingest Secret is a shared token that is used to authenticate Ingesters to the Indexer (or Federator)

[Control Secret](/configuration/parameters.html#control-auth)
  The Control Secret is a shared token that is used to Authenticate Ingesters to Webservers and visa versa

[Search Agent Secret](/configuration/parameters.html#search-agent-auth)
   The Search Agent Secrect is a shared token that is used to Authenticate the Search Agent to the Webserver

[Data Compression](/configuration/compression.html)
   Storage is compressed by default which allows shifting of some load from Storage to CPU. For many log types, Such as syslog data which compresses effectively, it can improve performance speed as less data must be read from physical disks,  and allow for more efficient use of the mass storage devices.

### Search

[Entry](/search/search.html#entries)
   Entries are your basic unit of data within Gravwell.  Entries are generated by the Ingester and sent to the indexer where they are stored until queried.  Every Entry contains 4 fields:  DATA, TIMESTAMP, SRC, and TAG.   The DATA field is the raw data sent to the Ingester,  the TIMESTAMP is the time extracted or applied to the entry which will be used to determine it's shard placement,  SRC will by default be the Source IP which the ingester received the data from,  and the TAG is the Gravwell Tag that will be used to determine where to place the data by the ingester.

[Search Pipeline](/search/search.html#search)
   The Search Pipeline is the core of Gravwell's functionality and operates in a similar fasion to the Linux/Unix Command line.   At it's most basic,  the Pipe will consist of at least 1 Extraction Module,  1 or more Filtering/processing modules,  and a Rendering Module.   As the data flows from the initial extraction module that pulls the data into the pipeline, it will be modified and transformed as it moves through the pipeline until it reaches the render module and the results are displayed.

[Extraction Modules](/search/search.html#extraction-modules)
   Extraction Modules are used to pull specific information out of the raw underlying data

[Enumerated Value / EV](/search/search.html#enumerated-values)
   Enumerated Values are special data elements which are created and used within the search pipeline.  All Entries contain the DATA, TIMESTAMP, SRC, and TAG EV's by default.  As the data moves through the pipeline additional Fields/EV's may be extracted and attached to the entry in the pipeline to facilitate additional processing and the eventual rendering of the results.
   
### Ingesters

[Attach](

### Search


