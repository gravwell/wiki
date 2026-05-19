# CoreDNS

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Ingester](/ingesters/integrations.md#coredns)
    Kit, [CoreDNS Kit](https://github.com/gravwell/coredns)
:::

## CoreDNS Configuration
[CoreDNS: Gravwell Integration Guide](https://coredns.io/explugins/gravwell/)

CoreDNS can be built with the Gravwell Plugin using the following shell code:
```shell
git clone https://github.com/coredns/coredns.git
pushd coredns
sed -i 's/metadata:metadata/metadata:metadata\ngravwell:github.com\/gravwell\/coredns/g' plugin.cfg
go generate
go get github.com/gravwell/coredns
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /tmp/coredns
popd
```

The CoreDNS binary will be located at `/tmp/coredns`. CoreDNS can then be started by providing a valid Corefile as the first argument. If you are running CoreDNS as a non-root user, you will need to give the binary the service bind capability.
```shell
setcap 'cap_net_bind_service=+ep' /tmp/coredns
```

Configuration is performed via the CoreDNS Corefile which has the basic syntax of directive value. Comments are preceded by the “#” character.  A basic Gravwell definition looks like so:
**Sample Configuration file
```
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
```

A unique Gravwell plugin section can be applied to each DNS listener. An example Corefile which listens to two different interfaces and applies a unique Gravwell configuration to each might look like so:
```
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
```

## Gravwell Configuration

### Gravwell Storage Well Configuration
**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/coredns.well`
```ini
[Storage-Well "coredns"]
    Location=/opt/gravwell/storage/coredns
    Tags=dns*
```
