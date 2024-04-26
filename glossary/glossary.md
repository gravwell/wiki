# Gravwell Glossary

This section will include basic descriptions for various terms and items used within Gravwell,   as well as links to the location in the documentation where more detailed information can be found.

### Primary Components

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

[Search Agent](/scripting/searchagent.html)
  The Gravwell Search Agent is the component which runs Automations

[Datastore](/distributed/frontend.html#the-datastore-server)
  The Datastore is the server/process which allows for the syncing of data across multiple Gravwell Webservers.  

[Overwatch](/distributed/overwatch.html)
   Gravwell Overwatch is a Advanced Deployment configuration where you can have specialized webservers that can query multiple self contained Gravwell Clusters.   

[CBAC / Capability Based Access Controls](/cbac/cbac.html)
  

### Indexer Terms

### Ingesters

### Search


