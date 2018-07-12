# Integrations

The gravwell ingest framework is open sourced via the BSD 2-clause license, which enables it to be directly included in both open source and commercial products.  Data processors and/or generators can directly embed the ingest framework and provide an easy to configure integration with Gravwell.  This documentation page is used to highlight and where appropriate document some Gravwell integrations.

## CoreDNS

CoreDNS is a highly configurable and plugin friendly DNS server meant to provide a base platform for DNS services.   The base functionality of CoreDNS provides a robust and performant DNS server that can act as a relay, proxy, or full blown DNS server.  CoreDNS is licensed under the Apache-2.0 and is available on [github](https://github.com/coredns/coredns).  To learn more about CoreDNS visit [https://coredns.io](https://coredns.io).

### Gravwell CoreDNS Plugin

A Gravwell plugin is available for CoreDNS which directly embeds the ingest framework into CoreDNS.  Using the plugin, a statically compiled and high performance DNS server can directly transmit DNS audit data to a Gravwell instance.  The plugin is licensed under BSD 2-Clause and available on [github](https://github.com/gravwell/coredns).

The plugin provides a complete ingest system which supports all the usual features: local caching for high reliability, load balancing, failover, etc...  Additional information about the Gravwell plugin can be found on the CoreDNS [External Plugins](https://coredns.io/explugins/gravwell/) page.

#### Building CoreDNS with Gravwell

Building CoreDNS with the Gravwell plugin requires that the Golang toolchain and compiler is installed, more information is available [here](https://golang.org/).

```
go get github.com/coredns/coredns
pushd $GOPATH/src/github.com/coredns/coredns/
sed -i 's/metadata:metadata/metadata:metadata\ngravwell:github.com\/gravwell\/coredns/g' plugin.cfg
go generate
CGO_ENABLED=0 go build -o /tmp/coredns
popd
```

The resulting statically compiled binary will be located at _/tmp/coredns_.  CoreDNS can then be started by providing a valid Corefile as the first argument.  If you are running CoreDNS as a non-root user, you will need to give the binary the [service bind capability](https://wiki.apache.org/httpd/NonRootPortBinding).

```
setcap 'cap_net_bind_service=+ep' /tmp/coredns
```

#### Configuring Gravwell Plugin

Configuration is performed via the CoreDNS Corefile which has the basic syntax of **directive** **value**.  Comments are preceeded by the "#" character.

The following configuration parameters are available:

* **Ingest-Secret** defines the token used to authenticate with indexers.  **Ingest-Secret** is required.
* **Cleartext-Target** defines the address and port for a remote indexer using a TCP connection.  IPv4 and IPv6 addresses as well as host names are supported.
* **Ciphertext-Target** defines the address and port for a remote indexer using a TLS connection.  IPv4 and IPv6 addresses as well as host names are supported.
* **Tag** specifies the tag that DNS audit logs are assigned.  Can be any alphanumeric value without special characters or spaces.  A valid Tag value is required.
* **Encoding** specifies the format of transported DNS audit logs.  Options are _json_ or _text_.  Deafult is _json_.
* **Insecure-Novalidate-TLS** toggles certificate validation on TLS connections.  Validation is on by default.
* **Log-Level** specifies the logging verbosity over the integrated gravwell tag.  Options are _OFF_ _INFO_ _WARN_ _ERROR_.  Default is _ERROR_.
* **Ingest-Cache-Path** specifies a file path for the cache system which engages when indexer connectivity is lost.  Path must be an absolute path to a writable file.
* **Max-Cache-Size-MB** specifies in megabytes the maximum size of the cache file.  This is used as a safty net.  Zero value is the default and represents unlimited.


A basic Gravwell definition looks like so:

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Tag dns
    Encoding json
    Log-Level INFO
    #Cleartext-Target 192.168.1.2:4023 #second indexer
    #Ciphertext-Target 192.168.1.1:4024
    #Insecure-Novalidate-TLS true #disable TLS certificate validation
    #Ingest-Cache-Path /tmp/coredns_ingest.cache #enable the local ingest cache
    #Max-Cache-Size-MB 1024
}
~~~

A unique Gravwell plugin section can be applied to each DNS listener.  An example Corefile which listens to two different interfaces and applies a unique Gravwell configuration to each might look like so:

~~~
.:53 {
  forward . 8.8.8.8:53 8.8.4.4:53 9.9.9.9:53
  errors stdout
  bind 172.20.0.1
  cache 240
  whoami
  gravwell {
   Ingest-Secret SecretSetOne
   Cleartext-Target 172.20.0.1:4023
   Tag dns
   Encoding json
  }
}

.:53 {
  forward	. tls://1.1.1.1
  errors stdout
  bind 192.168.1.1
  hosts
  cache 60s
  gravwell {
   Ingest-Secret SecretSetTwo
   Cleartext-Target 192.168.1.100:4023
   Cleartext-Target 192.168.1.101:4023
   Cleartext-Target 192.168.1.102:4023
   Tag dns
   Encoding json
  }
}
~~~

Notice that we are sending each of the DNS listeners to two completely independent Gravwell installations, one a single indexer the other a distributed cluster.

A sample Gravwell Corefile section which sends DNS requests to a single indexer over an unencrypted connection.  Local cache is disabled.

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Tag dns
  }
~~~

A sample Gravwell Corefile section which sends DNS requests to two indexers over a TLS connection and accepts unsigned certificates. Local cache is disabled.
IPv4 and IPv6 addresses are supported for both the Cleartext and Ciphertext targets.  IPv6 addresses must be enclosed in brackets.

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Ciphertext-Target 192.168.1.1:4024
    Ciphertext-Target [fe80::dead:beef:feed:febe%p1p1]:4024 #connecting to link local address via device p1p1
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~

A sample Gravwell Corefile section which sends DNS requests to two indexers over a TLS connection and accepts unsigned certificates. Local cache is disabled.

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Ciphertext-Target 192.168.1.1:4024
    Ciphertext-Target [dead::beef]:4024
    Insecure-Novalidate-TLS true
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~

A sample Gravwell Corefile section which sends DNS requests to two indexers and enables a local cache should indexer communication fail.  Up to 1GB of data can be locally cached.

~~~
gravwell {
    Ingest-Secret IngestSecretToken
    Cleartext-Target 192.168.1.1:4023
    Ciphertext-Target 192.168.1.2:4024
    Insecure-Novalidate-TLS true
    Ingest-Cache-Path /tmp/coredns_ingest.cache
    Max-Cache-Size-MB 1024
    Tag dns
    Encoding json
    Log-Level INFO
  }
~~~
