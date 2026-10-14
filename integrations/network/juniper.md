# Juniper

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [Simple Relay](/ingesters/simple_relay)
    Kit, [Juniper Kit](https://github.com/gravwell/kits/tree/main/juniper)
:::

## Juniper Configuration
Resources
* [System Log Explorer](https://apps.juniper.net/syslog-explorer/)
* [Overview of System Logging](https://www.juniper.net/documentation/us/en/software/junos/network-mgmt/topics/topic-map/system-logging.html)
* [Configure Syslog over TLS](https://www.juniper.net/documentation/us/en/software/junos/network-mgmt/topics/topic-map/syslog-over-tls.html)

To configure the device, follow these steps adapted from [Configure Syslog over TLS](https://www.juniper.net/documentation/us/en/software/junos/network-mgmt/topics/topic-map/syslog-over-tls.html):

1. Specify the syslog server that receives the system log messages. You can specify the IP address of the syslog server or a fully qualified hostname. In this example, use `10.102.70.233` as the IP address of the syslog server.

    ```
    user@host# set system syslog host 10.102.70.233 any any
    ```

2. Specify the port number of the syslog server.

    ```
    user@host# set system syslog host 10.102.70.233 port 10514
    ```

3. Specify the syslog transport protocol for the device. In this example, use TLS as the transport protocol.

    ```
    user@host# set system syslog host 10.102.70.233 transport tls
    ```

4. Specify the name of the trusted certificate authority (CA) group or specify the name of the CA profile to be used. In this example, use `example-ca` as the CA profile.

    ```
    user@host# set system syslog host 10.102.70.233 tlsdetails trusted-ca-group trusted-ca-group-name ca-profiles example-ca
    ```

5. Log all messages to the local `messages` file.

    ```
    user@host# set system syslog file messages any any
    ```

6. Commit the configuration.

    ```
    user@host# commit
    ```

7. Verify the configuration.

    ```
    user@host# show system syslog
    ```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/juniper-well.conf`
```ini
[Storage-Well "juniper"]
    Location=/opt/gravwell/storage/juniper
    Tags=juniper*
```

### Gravwell Ingester Configuration: Simple Relay
**Sample Juniper config:**  
Create or edit: `/opt/gravwell/etc/simple_relay.conf.d/juniper.conf`
```ini
[Listener "junipertcp"]
    Bind-String="tls://0.0.0.0:10514"
    Cert-File=/opt/gravwell/etc/cert.pem
    Key-File=/opt/gravwell/etc/key.pem
    Reader-Type=rfc5424
    Tag-Name=juniper
    Assume-Local-Timezone=true #if a time format does not have a timezone, assume local time
    Keep-Priority=true # leave the <nnn> priority tag at the start of each syslog entry
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart gravwell_simple_relay.service`
```
